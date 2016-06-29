// Code generated by protoc-gen-go.
// source: client.proto
// DO NOT EDIT!

/*
Package backup is a generated protocol buffer package.

It is generated from these files:
	client.proto
	server.proto

It has these top-level messages:
	ClientReConfigureRequest
	MountRequest
	UnMountRequest
	ServerCreateRequest
	ServerDeleteRequest
*/
package backup

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClientReConfigureRequest struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Namespace   string `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	ClusterName string `protobuf:"bytes,3,opt,name=cluster_name" json:"cluster_name,omitempty"`
}

func (m *ClientReConfigureRequest) Reset()                    { *m = ClientReConfigureRequest{} }
func (m *ClientReConfigureRequest) String() string            { return proto.CompactTextString(m) }
func (*ClientReConfigureRequest) ProtoMessage()               {}
func (*ClientReConfigureRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type MountRequest struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Namespace   string `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	ClusterName string `protobuf:"bytes,3,opt,name=cluster_name" json:"cluster_name,omitempty"`
	DiskType    string `protobuf:"bytes,4,opt,name=disk_type" json:"disk_type,omitempty"`
	DiskName    string `protobuf:"bytes,5,opt,name=disk_name" json:"disk_name,omitempty"`
	MountPath   string `protobuf:"bytes,6,opt,name=mount_path" json:"mount_path,omitempty"`
}

func (m *MountRequest) Reset()                    { *m = MountRequest{} }
func (m *MountRequest) String() string            { return proto.CompactTextString(m) }
func (*MountRequest) ProtoMessage()               {}
func (*MountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type UnMountRequest struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Namespace   string `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	ClusterName string `protobuf:"bytes,3,opt,name=cluster_name" json:"cluster_name,omitempty"`
	DiskName    string `protobuf:"bytes,5,opt,name=disk_name" json:"disk_name,omitempty"`
}

func (m *UnMountRequest) Reset()                    { *m = UnMountRequest{} }
func (m *UnMountRequest) String() string            { return proto.CompactTextString(m) }
func (*UnMountRequest) ProtoMessage()               {}
func (*UnMountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*ClientReConfigureRequest)(nil), "backup.ClientReConfigureRequest")
	proto.RegisterType((*MountRequest)(nil), "backup.MountRequest")
	proto.RegisterType((*UnMountRequest)(nil), "backup.UnMountRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Client service

type ClientClient interface {
	ReConfigure(ctx context.Context, in *ClientReConfigureRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Mount(ctx context.Context, in *MountRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	UnMount(ctx context.Context, in *UnMountRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type clientClient struct {
	cc *grpc.ClientConn
}

func NewClientClient(cc *grpc.ClientConn) ClientClient {
	return &clientClient{cc}
}

func (c *clientClient) ReConfigure(ctx context.Context, in *ClientReConfigureRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/backup.Client/ReConfigure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) Mount(ctx context.Context, in *MountRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/backup.Client/Mount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientClient) UnMount(ctx context.Context, in *UnMountRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/backup.Client/UnMount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Client service

type ClientServer interface {
	ReConfigure(context.Context, *ClientReConfigureRequest) (*dtypes.VoidResponse, error)
	Mount(context.Context, *MountRequest) (*dtypes.VoidResponse, error)
	UnMount(context.Context, *UnMountRequest) (*dtypes.VoidResponse, error)
}

func RegisterClientServer(s *grpc.Server, srv ClientServer) {
	s.RegisterService(&_Client_serviceDesc, srv)
}

func _Client_ReConfigure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientReConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).ReConfigure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup.Client/ReConfigure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).ReConfigure(ctx, req.(*ClientReConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_Mount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).Mount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup.Client/Mount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).Mount(ctx, req.(*MountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Client_UnMount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnMountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServer).UnMount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup.Client/UnMount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServer).UnMount(ctx, req.(*UnMountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Client_serviceDesc = grpc.ServiceDesc{
	ServiceName: "backup.Client",
	HandlerType: (*ClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReConfigure",
			Handler:    _Client_ReConfigure_Handler,
		},
		{
			MethodName: "Mount",
			Handler:    _Client_Mount_Handler,
		},
		{
			MethodName: "UnMount",
			Handler:    _Client_UnMount_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func init() { proto.RegisterFile("client.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x92, 0x3f, 0x4b, 0x03, 0x41,
	0x10, 0xc5, 0x49, 0x4c, 0x4e, 0x32, 0x06, 0xc1, 0x25, 0x84, 0xe3, 0xb0, 0x08, 0x29, 0x24, 0xf8,
	0xe7, 0xce, 0x98, 0x42, 0xb0, 0x4d, 0x6d, 0x61, 0x40, 0x0b, 0x9b, 0xb0, 0xb9, 0x9b, 0xc4, 0x23,
	0xc9, 0xee, 0x7a, 0xbb, 0x27, 0xd8, 0x28, 0xf8, 0x15, 0xfc, 0x68, 0xf6, 0x56, 0x7e, 0x10, 0xb3,
	0xb3, 0x1e, 0x26, 0xc5, 0x91, 0x42, 0x9b, 0x83, 0x7b, 0x0c, 0xef, 0x37, 0xf3, 0xde, 0x42, 0x33,
	0x5e, 0xa4, 0x28, 0x4c, 0xa8, 0x32, 0x69, 0x24, 0xf3, 0x26, 0x3c, 0x9e, 0xe7, 0x2a, 0x38, 0x9c,
	0x49, 0x39, 0x5b, 0x60, 0xc4, 0x55, 0x1a, 0x71, 0x21, 0xa4, 0xe1, 0x26, 0x95, 0x42, 0xbb, 0xa9,
	0xa0, 0x6d, 0xe5, 0xc4, 0x3c, 0x2b, 0xd4, 0x11, 0x7d, 0x9d, 0xde, 0xbd, 0x01, 0x7f, 0x48, 0x6e,
	0x23, 0x1c, 0x4a, 0x31, 0x4d, 0x67, 0x79, 0x86, 0x23, 0x7c, 0xcc, 0x51, 0x1b, 0xd6, 0x84, 0x9a,
	0xe0, 0x4b, 0xf4, 0x2b, 0x9d, 0x4a, 0xaf, 0xc1, 0x0e, 0xa0, 0x61, 0xff, 0xb4, 0xe2, 0x31, 0xfa,
	0x55, 0x92, 0x5a, 0x76, 0x95, 0x5c, 0x1b, 0xcc, 0xc6, 0x34, 0xb8, 0x63, 0xd5, 0xee, 0x2b, 0x34,
	0xaf, 0x65, 0x6e, 0x1d, 0xff, 0x62, 0x63, 0x07, 0x93, 0x54, 0xcf, 0xc7, 0x76, 0x5b, 0xbf, 0xb6,
	0x21, 0xd1, 0x54, 0x9d, 0x24, 0x06, 0xb0, 0xb4, 0xb0, 0xb1, 0xe2, 0xe6, 0xc1, 0xf7, 0x68, 0x81,
	0x7b, 0xd8, 0xbf, 0x15, 0xff, 0xb9, 0xc2, 0x2f, 0xef, 0xe2, 0xb3, 0x0a, 0x9e, 0x0b, 0x8c, 0xbd,
	0xc0, 0xde, 0x5a, 0x68, 0xac, 0x13, 0xba, 0x22, 0xc2, 0xb2, 0x3c, 0x83, 0x56, 0xe8, 0x0a, 0x08,
	0xef, 0x64, 0x9a, 0x8c, 0x56, 0xfc, 0x55, 0x3d, 0xd8, 0xbd, 0x7c, 0xfb, 0xf8, 0x7a, 0xaf, 0xf6,
	0x83, 0xd3, 0x55, 0x73, 0x4a, 0xc7, 0x32, 0x71, 0x15, 0x3a, 0xb3, 0xe8, 0xe9, 0x3c, 0xec, 0x47,
	0xae, 0xef, 0x28, 0xc3, 0xb8, 0xb0, 0xbc, 0xaa, 0x1c, 0xb3, 0x29, 0xd4, 0xe9, 0x48, 0xd6, 0x2a,
	0xc8, 0xeb, 0x37, 0x97, 0xd0, 0xfa, 0x44, 0x3b, 0x09, 0x8e, 0xb6, 0xd2, 0x28, 0x56, 0xcb, 0x59,
	0xc0, 0xee, 0x4f, 0x9c, 0xac, 0x5d, 0x90, 0x36, 0xf3, 0x2d, 0x61, 0x0d, 0x88, 0x75, 0x16, 0xf4,
	0xb6, 0xb2, 0x72, 0x51, 0xd0, 0x26, 0x1e, 0xbd, 0xcb, 0xc1, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x21, 0xfd, 0x22, 0x10, 0xe5, 0x02, 0x00, 0x00,
}
