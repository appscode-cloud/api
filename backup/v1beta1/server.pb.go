// Code generated by protoc-gen-go.
// source: server.proto
// DO NOT EDIT!

package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ServerCreateRequest struct {
	Cluster    string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Credential string `protobuf:"bytes,2,opt,name=credential" json:"credential,omitempty"`
	BucketName string `protobuf:"bytes,3,opt,name=bucket_name,json=bucketName" json:"bucket_name,omitempty"`
	Disk       string `protobuf:"bytes,4,opt,name=disk" json:"disk,omitempty"`
	// TODO will be removed
	Namespace string `protobuf:"bytes,5,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *ServerCreateRequest) Reset()                    { *m = ServerCreateRequest{} }
func (m *ServerCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*ServerCreateRequest) ProtoMessage()               {}
func (*ServerCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ServerDeleteRequest struct {
	Cluster   string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *ServerDeleteRequest) Reset()                    { *m = ServerDeleteRequest{} }
func (m *ServerDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*ServerDeleteRequest) ProtoMessage()               {}
func (*ServerDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func init() {
	proto.RegisterType((*ServerCreateRequest)(nil), "backup.v1beta1.ServerCreateRequest")
	proto.RegisterType((*ServerDeleteRequest)(nil), "backup.v1beta1.ServerDeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Servers service

type ServersClient interface {
	Create(ctx context.Context, in *ServerCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *ServerDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type serversClient struct {
	cc *grpc.ClientConn
}

func NewServersClient(cc *grpc.ClientConn) ServersClient {
	return &serversClient{cc}
}

func (c *serversClient) Create(ctx context.Context, in *ServerCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/backup.v1beta1.Servers/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serversClient) Delete(ctx context.Context, in *ServerDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/backup.v1beta1.Servers/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Servers service

type ServersServer interface {
	Create(context.Context, *ServerCreateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *ServerDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterServersServer(s *grpc.Server, srv ServersServer) {
	s.RegisterService(&_Servers_serviceDesc, srv)
}

func _Servers_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup.v1beta1.Servers/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServersServer).Create(ctx, req.(*ServerCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Servers_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServersServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup.v1beta1.Servers/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServersServer).Delete(ctx, req.(*ServerDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Servers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "backup.v1beta1.Servers",
	HandlerType: (*ServersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Servers_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Servers_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("server.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x92, 0xc1, 0x4a, 0xf3, 0x40,
	0x14, 0x85, 0x49, 0xfe, 0xfe, 0x2d, 0x1d, 0xc5, 0xc5, 0x28, 0x12, 0x4a, 0x51, 0x89, 0x1b, 0xe9,
	0x22, 0x43, 0x75, 0xe7, 0x46, 0xd4, 0x82, 0x20, 0xe8, 0xa2, 0x82, 0x0b, 0x37, 0x32, 0x99, 0x5c,
	0x4a, 0x48, 0x3a, 0x33, 0xce, 0x9d, 0x14, 0x44, 0xdc, 0xb8, 0x16, 0x5c, 0xf8, 0x06, 0x2e, 0x7d,
	0x1d, 0x5f, 0xc1, 0x07, 0x91, 0xce, 0xa4, 0x68, 0xc5, 0x82, 0xe0, 0x26, 0x4c, 0xce, 0x3d, 0x39,
	0x7c, 0xf7, 0x64, 0xc8, 0x32, 0x82, 0x99, 0x80, 0x49, 0xb4, 0x51, 0x56, 0xd1, 0x95, 0x94, 0x8b,
	0xa2, 0xd2, 0xc9, 0xa4, 0x9f, 0x82, 0xe5, 0xfd, 0x4e, 0x77, 0xa4, 0xd4, 0xa8, 0x04, 0xc6, 0x75,
	0xce, 0xb8, 0x94, 0xca, 0x72, 0x9b, 0x2b, 0x89, 0xde, 0xdd, 0x59, 0x9f, 0xca, 0x99, 0xbd, 0xd5,
	0x80, 0xcc, 0x3d, 0xbd, 0x1e, 0xbf, 0x04, 0x64, 0xf5, 0xc2, 0xc5, 0x1e, 0x1b, 0xe0, 0x16, 0x86,
	0x70, 0x53, 0x01, 0x5a, 0x1a, 0x91, 0x96, 0x28, 0x2b, 0xb4, 0x60, 0xa2, 0x60, 0x2b, 0xd8, 0x69,
	0x0f, 0x67, 0xaf, 0x74, 0x83, 0x10, 0x61, 0x20, 0x03, 0x69, 0x73, 0x5e, 0x46, 0xa1, 0x1b, 0x7e,
	0x51, 0xe8, 0x26, 0x59, 0x4a, 0x2b, 0x51, 0x80, 0xbd, 0x96, 0x7c, 0x0c, 0xd1, 0x3f, 0x6f, 0xf0,
	0xd2, 0x39, 0x1f, 0x03, 0xa5, 0xa4, 0x91, 0xe5, 0x58, 0x44, 0x0d, 0x37, 0x71, 0x67, 0xda, 0x25,
	0xed, 0xa9, 0x1b, 0x35, 0x17, 0x10, 0xfd, 0x77, 0x83, 0x4f, 0x21, 0x3e, 0x9b, 0x31, 0x0e, 0xa0,
	0x84, 0xdf, 0x30, 0xce, 0xc5, 0x85, 0xdf, 0xe2, 0x76, 0x5f, 0x43, 0xd2, 0xf2, 0x79, 0x48, 0x9f,
	0x02, 0xd2, 0xf4, 0x9b, 0xd3, 0xed, 0x64, 0xbe, 0xd1, 0xe4, 0x87, 0x5e, 0x3a, 0x6b, 0x89, 0x2f,
	0x31, 0xb9, 0x54, 0x79, 0x36, 0x04, 0xd4, 0x4a, 0x22, 0xc4, 0xa7, 0x0f, 0x6f, 0xef, 0xcf, 0xe1,
	0x20, 0x3e, 0x60, 0x5c, 0x6b, 0x14, 0x2a, 0xf3, 0xbf, 0xa1, 0xa8, 0x52, 0x30, 0x12, 0x2c, 0x20,
	0xab, 0x33, 0x59, 0x4d, 0x88, 0xec, 0xae, 0x3e, 0xdd, 0xbb, 0x4f, 0x58, 0xca, 0x45, 0x55, 0xf2,
	0xfd, 0xa0, 0x47, 0x1f, 0x03, 0xd2, 0xf4, 0x7b, 0x2e, 0x22, 0x9a, 0x6b, 0x61, 0x01, 0xd1, 0x89,
	0x23, 0x3a, 0xec, 0xfd, 0x95, 0xe8, 0xa8, 0x7d, 0xd5, 0xaa, 0x9d, 0x69, 0xd3, 0x5d, 0x99, 0xbd,
	0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x58, 0xc1, 0xa5, 0x69, 0x88, 0x02, 0x00, 0x00,
}
