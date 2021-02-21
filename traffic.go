package traffic

import (
	"fmt"
	"log"
	"time"
	"sync"
	"context"
	"google.golang.org/grpc"
	pb "github.com/golang/protobuf/proto"
	"github.com/open-fsm/spec/proto"
	"github.com/open-fsm/traffic/trafficpb"
)

type VR interface {
	Call(context.Context, proto.Message) error
}

type Traffic interface {
	Send([]proto.Message)
	AddPeer(id int64, nodes []string)
	DelPeer(id int64)
	WaitError() chan error
}

type Config struct {
	Cid int64
}

func (c *Config) Validate() error {
	return nil
}

type traffic struct{
	sync.RWMutex

	id int64
	cid int64
	vr VR
	peers map[int64]*peer
	errorC chan error
}

func New(id, cid int64, vr VR, errorC chan error) Traffic {
	return &traffic{
		id: id,
		cid: cid,
		vr: vr,
		errorC: errorC,
	}
}

func (t *traffic) Send(msgs []proto.Message) {
	for _, m := range msgs {
		if m.To == 0 {
			continue
		}
		to := int64(m.To)

		peer, ok := t.peers[to]
		if ok {
			peer.send(m)
			continue
		}
		log.Printf("traffic: send message to unknown destination %t", to)
	}
}

func (t *traffic) WaitError() chan error {
	return t.errorC
}

func (t *traffic) AddPeer(id int64, nodes []string) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.peers[id]; ok {
		return
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure(), grpc.WithBlock())
	cc, err := grpc.Dial(nodes[0], opts...)
	if err != nil {
		cc.Close()
		log.Fatalf("traffic: fail to dial: %v", err)
	}
	t.peers[id] = NewPeer(cc, id, t.cid, t.vr, t.errorC)
}

func (t *traffic) DelPeer(id int64) {
	t.Lock()
	defer t.Unlock()
	t.delPeer(id)
}

func (t *traffic) delPeer(id int64) {
	if peer, ok := t.peers[id]; ok {
		peer.Stop()
	} else {
		log.Panicf("traffic: unexpected delete of unknown peer '%d'", id)
	}
	delete(t.peers, id)
}

func (t *traffic) GetPeer(id int64) *peer {
	t.RLock()
	defer t.RUnlock()
	return t.peers[id]
}

type server struct {
	sync.Mutex
	trafficpb.UnimplementedTrafficServer;VR
	cid int64
}

func NewServer(cfg *Config, vr VR) *server {
	ds := &server{
		cid: cfg.Cid,
		VR: vr,
	}
	return ds
}

func (s *server) Write(ctx context.Context, req *trafficpb.Request) (*trafficpb.Response, error) {
	if req.Data == nil {
		return nil ,fmt.Errorf("data is nil")
	}

	var m proto.Message
	codec.MustUnmarshal(&m, req.Data)
	if err := s.Call(context.TODO(), m); err != nil {
		return nil, err
	}
	return &trafficpb.Response{}, nil
}

const (
	connPerSender = 4
	senderBufSize = 64

	dialTimeout      = time.Second
	connReadTimeout  = 5 * time.Second
	connWriteTimeout = 5 * time.Second
)

type peer struct {
	trafficpb.TrafficClient
	sync.Mutex

	id       int64
	cid      int64
	vr       VR
	errorC   chan error
	messageC chan *proto.Message
	alive    bool
}

func NewPeer(cc *grpc.ClientConn, id, cid int64, vr VR, errorC chan error) *peer {
	p := &peer{
		id:            id,
		cid:           cid,
		vr:            vr,
		errorC:        errorC,
		TrafficClient: trafficpb.NewTrafficClient(cc),
		messageC:      make(chan *proto.Message, senderBufSize),
	}
	go p.handler()
	return p
}

func (p *peer) send(m proto.Message) error {
	select {
	case p.messageC <- &m:
		return nil
	default:
		log.Printf("triffic.peer sender: dropping %s because maximal number %d of sender buffer entries to %s has been reached",
			m.Type, senderBufSize, p.messageC)
		return fmt.Errorf("reach maximal serving")
	}
	return nil
}

func (p *peer) push(data []byte) error {
	res, err := p.Write(context.TODO(), &trafficpb.Request{
		Cid: p.cid,
		Data: data,
	})
	if err != nil {
		return err
	}

	p.Lock()
	switch res.Code {
	case trafficpb.Bar:
	case trafficpb.Foo:
	}
	p.Unlock()

	return nil
}

func (p *peer) handler() {
	for msg := range p.messageC {
		data, _ := pb.Marshal(msg)
		p.Lock()
		err := p.push(data)
		if err != nil {
			if p.alive {
				p.alive = false
			}
		} else {
			if !p.alive	{
				p.alive = true
			}
		}
		p.Unlock()
	}
}

func (p *peer) Stop() {
	close(p.messageC)
}