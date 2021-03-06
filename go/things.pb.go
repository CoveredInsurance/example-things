// Code generated by protoc-gen-go. DO NOT EDIT.
// source: things.proto

/*
Package things is a generated protocol buffer package.

It is generated from these files:
	things.proto

It has these top-level messages:
	Thing
*/
package things

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Thing
type Thing struct {
	Id   uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Thing) Reset()                    { *m = Thing{} }
func (m *Thing) String() string            { return proto.CompactTextString(m) }
func (*Thing) ProtoMessage()               {}
func (*Thing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Thing) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Thing) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Thing)(nil), "things.Thing")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Things service

type ThingsClient interface {
	Get(ctx context.Context, in *Thing, opts ...grpc.CallOption) (*Thing, error)
	Post(ctx context.Context, in *Thing, opts ...grpc.CallOption) (*Thing, error)
}

type thingsClient struct {
	cc *grpc.ClientConn
}

func NewThingsClient(cc *grpc.ClientConn) ThingsClient {
	return &thingsClient{cc}
}

func (c *thingsClient) Get(ctx context.Context, in *Thing, opts ...grpc.CallOption) (*Thing, error) {
	out := new(Thing)
	err := grpc.Invoke(ctx, "/things.Things/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *thingsClient) Post(ctx context.Context, in *Thing, opts ...grpc.CallOption) (*Thing, error) {
	out := new(Thing)
	err := grpc.Invoke(ctx, "/things.Things/Post", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Things service

type ThingsServer interface {
	Get(context.Context, *Thing) (*Thing, error)
	Post(context.Context, *Thing) (*Thing, error)
}

func RegisterThingsServer(s *grpc.Server, srv ThingsServer) {
	s.RegisterService(&_Things_serviceDesc, srv)
}

func _Things_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Thing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThingsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/things.Things/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThingsServer).Get(ctx, req.(*Thing))
	}
	return interceptor(ctx, in, info, handler)
}

func _Things_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Thing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ThingsServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/things.Things/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ThingsServer).Post(ctx, req.(*Thing))
	}
	return interceptor(ctx, in, info, handler)
}

var _Things_serviceDesc = grpc.ServiceDesc{
	ServiceName: "things.Things",
	HandlerType: (*ThingsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Things_Get_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _Things_Post_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "things.proto",
}

func init() { proto.RegisterFile("things.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xc9, 0xc8, 0xcc,
	0x4b, 0x2f, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0xa4, 0x64, 0xd2, 0xf3,
	0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b,
	0x32, 0xf3, 0xf3, 0xa0, 0xaa, 0x94, 0xb4, 0xb9, 0x58, 0x43, 0x40, 0xea, 0x84, 0xf8, 0xb8, 0x98,
	0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x82, 0x98, 0x32, 0x53, 0x84, 0x84, 0xb8, 0x58,
	0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0xa3, 0x3c, 0x2e,
	0x36, 0xb0, 0xe2, 0x62, 0x21, 0x1b, 0x2e, 0x66, 0xf7, 0xd4, 0x12, 0x21, 0x5e, 0x3d, 0xa8, 0x95,
	0x60, 0x61, 0x29, 0x54, 0xae, 0x92, 0x78, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0x04, 0x85, 0xf8, 0xf5,
	0xcb, 0x0c, 0xf5, 0x21, 0x32, 0xfa, 0xd5, 0x99, 0x29, 0xb5, 0x42, 0x6a, 0x5c, 0x2c, 0x01, 0xf9,
	0xc5, 0x84, 0xb4, 0x33, 0x38, 0xc9, 0x73, 0xc9, 0x26, 0xe7, 0xe7, 0xea, 0x65, 0x96, 0x14, 0x27,
	0xe7, 0x97, 0xa5, 0x16, 0xa5, 0xa6, 0xe8, 0xa5, 0x56, 0x24, 0xe6, 0x16, 0xe4, 0xa4, 0x42, 0x15,
	0x26, 0xb1, 0x81, 0x3d, 0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x39, 0x0c, 0x58, 0x01, 0xfa,
	0x00, 0x00, 0x00,
}
