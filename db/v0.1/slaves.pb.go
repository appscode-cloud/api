// Code generated by protoc-gen-go.
// source: slaves.proto
// DO NOT EDIT!

package db

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gengo/grpc-gateway/third_party/googleapis/google/api"
import dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SlaveAddRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Config  string `protobuf:"bytes,3,opt,name=config" json:"config,omitempty"`
}

func (m *SlaveAddRequest) Reset()                    { *m = SlaveAddRequest{} }
func (m *SlaveAddRequest) String() string            { return proto.CompactTextString(m) }
func (*SlaveAddRequest) ProtoMessage()               {}
func (*SlaveAddRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterType((*SlaveAddRequest)(nil), "db.SlaveAddRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Slaves service

type SlavesClient interface {
	Add(ctx context.Context, in *SlaveAddRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type slavesClient struct {
	cc *grpc.ClientConn
}

func NewSlavesClient(cc *grpc.ClientConn) SlavesClient {
	return &slavesClient{cc}
}

func (c *slavesClient) Add(ctx context.Context, in *SlaveAddRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/db.Slaves/Add", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Slaves service

type SlavesServer interface {
	Add(context.Context, *SlaveAddRequest) (*dtypes.VoidResponse, error)
}

func RegisterSlavesServer(s *grpc.Server, srv SlavesServer) {
	s.RegisterService(&_Slaves_serviceDesc, srv)
}

func _Slaves_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlaveAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlavesServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/db.Slaves/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlavesServer).Add(ctx, req.(*SlaveAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Slaves_serviceDesc = grpc.ServiceDesc{
	ServiceName: "db.Slaves",
	HandlerType: (*SlavesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Slaves_Add_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func init() { proto.RegisterFile("slaves.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xce, 0x49, 0x2c,
	0x4b, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4a, 0x49, 0x92, 0x92, 0x49, 0xcf,
	0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c,
	0xc9, 0xcc, 0xcf, 0x83, 0xaa, 0x90, 0x12, 0x03, 0x09, 0xa7, 0x94, 0x54, 0x16, 0xa4, 0x16, 0xeb,
	0x83, 0x49, 0x88, 0xb8, 0x92, 0x03, 0x17, 0x7f, 0x30, 0xc8, 0x24, 0xc7, 0x94, 0x94, 0xa0, 0xd4,
	0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x7e, 0x2e, 0xf6, 0xe4, 0x9c, 0xd2, 0xe2, 0x92, 0xd4, 0x22,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0x4e, 0x21, 0x1e, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26,
	0x30, 0x8f, 0x8f, 0x8b, 0x2d, 0x39, 0x3f, 0x2f, 0x2d, 0x33, 0x5d, 0x82, 0x19, 0xc4, 0x37, 0xca,
	0xe1, 0x62, 0x03, 0x9b, 0x50, 0x2c, 0x94, 0xc4, 0xc5, 0x0c, 0x34, 0x46, 0x48, 0x58, 0x2f, 0x25,
	0x49, 0x0f, 0xcd, 0x50, 0x29, 0x11, 0x3d, 0x88, 0xe5, 0x7a, 0x61, 0xf9, 0x99, 0x40, 0xc1, 0xe2,
	0x02, 0xa0, 0xd3, 0x52, 0x95, 0x0c, 0x9a, 0x2e, 0x3f, 0x99, 0xcc, 0xa4, 0x25, 0xa5, 0x0a, 0x74,
	0x75, 0x41, 0x71, 0x72, 0x7e, 0x0a, 0xc4, 0xf9, 0x29, 0x49, 0xfa, 0x65, 0x06, 0x7a, 0x86, 0xfa,
	0x60, 0x3f, 0xea, 0x57, 0x43, 0x5d, 0x53, 0x6b, 0xc5, 0xa8, 0x95, 0xc4, 0x06, 0x76, 0xb6, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x68, 0x1f, 0xeb, 0x41, 0x00, 0x01, 0x00, 0x00,
}
