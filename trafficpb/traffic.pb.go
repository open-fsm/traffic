// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0-devel
// 	protoc        v3.7.1
// source: traffic.proto

package trafficpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_traffic_proto protoreflect.FileDescriptor

var file_traffic_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x70, 0x62, 0x1a, 0x0a, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x3d, 0x0a, 0x07, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x12, 0x32, 0x0a, 0x05, 0x57, 0x72, 0x69, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x74, 0x72, 0x61,
	0x66, 0x66, 0x69, 0x63, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_traffic_proto_goTypes = []interface{}{
	(*Request)(nil),  // 0: trafficpb.Request
	(*Response)(nil), // 1: trafficpb.Response
}
var file_traffic_proto_depIdxs = []int32{
	0, // 0: trafficpb.Traffic.Write:input_type -> trafficpb.Request
	1, // 1: trafficpb.Traffic.Write:output_type -> trafficpb.Response
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_traffic_proto_init() }
func file_traffic_proto_init() {
	if File_traffic_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_traffic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_traffic_proto_goTypes,
		DependencyIndexes: file_traffic_proto_depIdxs,
	}.Build()
	File_traffic_proto = out.File
	file_traffic_proto_rawDesc = nil
	file_traffic_proto_goTypes = nil
	file_traffic_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TrafficClient is the client API for Traffic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TrafficClient interface {
	Write(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type trafficClient struct {
	cc grpc.ClientConnInterface
}

func NewTrafficClient(cc grpc.ClientConnInterface) TrafficClient {
	return &trafficClient{cc}
}

func (c *trafficClient) Write(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/trafficpb.Traffic/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TrafficServer is the server API for Traffic service.
type TrafficServer interface {
	Write(context.Context, *Request) (*Response, error)
}

// UnimplementedTrafficServer can be embedded to have forward compatible implementations.
type UnimplementedTrafficServer struct {
}

func (*UnimplementedTrafficServer) Write(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}

func RegisterTrafficServer(s *grpc.Server, srv TrafficServer) {
	s.RegisterService(&_Traffic_serviceDesc, srv)
}

func _Traffic_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrafficServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/trafficpb.Traffic/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrafficServer).Write(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Traffic_serviceDesc = grpc.ServiceDesc{
	ServiceName: "trafficpb.Traffic",
	HandlerType: (*TrafficServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _Traffic_Write_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "traffic.proto",
}
