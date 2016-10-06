// Code generated by protoc-gen-go.
// source: persistentvolume.proto
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

type PVRegisterRequest struct {
	Cluster    string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Identifier string `protobuf:"bytes,3,opt,name=identifier" json:"identifier,omitempty"`
	Plugin     string `protobuf:"bytes,4,opt,name=plugin" json:"plugin,omitempty"`
	SizeGb     int64  `protobuf:"varint,5,opt,name=size_gb,json=sizeGb" json:"size_gb,omitempty"`
	Endpoint   string `protobuf:"bytes,6,opt,name=endpoint" json:"endpoint,omitempty"`
}

func (m *PVRegisterRequest) Reset()                    { *m = PVRegisterRequest{} }
func (m *PVRegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*PVRegisterRequest) ProtoMessage()               {}
func (*PVRegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type PVUnregisterRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *PVUnregisterRequest) Reset()                    { *m = PVUnregisterRequest{} }
func (m *PVUnregisterRequest) String() string            { return proto.CompactTextString(m) }
func (*PVUnregisterRequest) ProtoMessage()               {}
func (*PVUnregisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type PVDescribeRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *PVDescribeRequest) Reset()                    { *m = PVDescribeRequest{} }
func (m *PVDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*PVDescribeRequest) ProtoMessage()               {}
func (*PVDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type PVInfo struct {
	Name        string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	SizeGb      int64    `protobuf:"varint,2,opt,name=size_gb,json=sizeGb" json:"size_gb,omitempty"`
	Status      string   `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	Volume      string   `protobuf:"bytes,4,opt,name=volume" json:"volume,omitempty"`
	Claim       string   `protobuf:"bytes,5,opt,name=claim" json:"claim,omitempty"`
	Plugin      string   `protobuf:"bytes,6,opt,name=plugin" json:"plugin,omitempty"`
	AccessModes []string `protobuf:"bytes,7,rep,name=accessModes" json:"accessModes,omitempty"`
}

func (m *PVInfo) Reset()                    { *m = PVInfo{} }
func (m *PVInfo) String() string            { return proto.CompactTextString(m) }
func (*PVInfo) ProtoMessage()               {}
func (*PVInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type PVDescribeResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Pv     *PVInfo        `protobuf:"bytes,2,opt,name=pv" json:"pv,omitempty"`
}

func (m *PVDescribeResponse) Reset()                    { *m = PVDescribeResponse{} }
func (m *PVDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*PVDescribeResponse) ProtoMessage()               {}
func (*PVDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *PVDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *PVDescribeResponse) GetPv() *PVInfo {
	if m != nil {
		return m.Pv
	}
	return nil
}

func init() {
	proto.RegisterType((*PVRegisterRequest)(nil), "pv.v1beta1.PVRegisterRequest")
	proto.RegisterType((*PVUnregisterRequest)(nil), "pv.v1beta1.PVUnregisterRequest")
	proto.RegisterType((*PVDescribeRequest)(nil), "pv.v1beta1.PVDescribeRequest")
	proto.RegisterType((*PVInfo)(nil), "pv.v1beta1.PVInfo")
	proto.RegisterType((*PVDescribeResponse)(nil), "pv.v1beta1.PVDescribeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for PersistentVolumes service

type PersistentVolumesClient interface {
	Describe(ctx context.Context, in *PVDescribeRequest, opts ...grpc.CallOption) (*PVDescribeResponse, error)
	Register(ctx context.Context, in *PVRegisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Unregister(ctx context.Context, in *PVUnregisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type persistentVolumesClient struct {
	cc *grpc.ClientConn
}

func NewPersistentVolumesClient(cc *grpc.ClientConn) PersistentVolumesClient {
	return &persistentVolumesClient{cc}
}

func (c *persistentVolumesClient) Describe(ctx context.Context, in *PVDescribeRequest, opts ...grpc.CallOption) (*PVDescribeResponse, error) {
	out := new(PVDescribeResponse)
	err := grpc.Invoke(ctx, "/pv.v1beta1.PersistentVolumes/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumesClient) Register(ctx context.Context, in *PVRegisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/pv.v1beta1.PersistentVolumes/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumesClient) Unregister(ctx context.Context, in *PVUnregisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/pv.v1beta1.PersistentVolumes/Unregister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PersistentVolumes service

type PersistentVolumesServer interface {
	Describe(context.Context, *PVDescribeRequest) (*PVDescribeResponse, error)
	Register(context.Context, *PVRegisterRequest) (*dtypes.VoidResponse, error)
	Unregister(context.Context, *PVUnregisterRequest) (*dtypes.VoidResponse, error)
}

func RegisterPersistentVolumesServer(s *grpc.Server, srv PersistentVolumesServer) {
	s.RegisterService(&_PersistentVolumes_serviceDesc, srv)
}

func _PersistentVolumes_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PVDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumesServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.v1beta1.PersistentVolumes/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumesServer).Describe(ctx, req.(*PVDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersistentVolumes_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PVRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumesServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.v1beta1.PersistentVolumes/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumesServer).Register(ctx, req.(*PVRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PersistentVolumes_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PVUnregisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersistentVolumesServer).Unregister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pv.v1beta1.PersistentVolumes/Unregister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersistentVolumesServer).Unregister(ctx, req.(*PVUnregisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PersistentVolumes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pv.v1beta1.PersistentVolumes",
	HandlerType: (*PersistentVolumesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Describe",
			Handler:    _PersistentVolumes_Describe_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _PersistentVolumes_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _PersistentVolumes_Unregister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor2,
}

func init() { proto.RegisterFile("persistentvolume.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x8a, 0xd4, 0x40,
	0x10, 0xa6, 0x67, 0x76, 0x33, 0x3b, 0x35, 0x20, 0x6c, 0xbb, 0x8c, 0x21, 0xe8, 0x3a, 0xe4, 0x20,
	0xc3, 0x1e, 0x12, 0x76, 0xbc, 0x79, 0xf3, 0x07, 0x44, 0x50, 0x19, 0xa3, 0x8e, 0xe0, 0x45, 0xf3,
	0x53, 0x1b, 0x1a, 0x33, 0xdd, 0x6d, 0xba, 0x13, 0xd0, 0x65, 0x2f, 0xbe, 0x82, 0x17, 0x3d, 0xf9,
	0x06, 0xde, 0x7c, 0x12, 0x5f, 0xc1, 0xa3, 0x0f, 0x21, 0xe9, 0xfc, 0x6c, 0x66, 0x64, 0x2f, 0xba,
	0x97, 0xd0, 0xf5, 0x55, 0xaa, 0xf3, 0x7d, 0x5f, 0x55, 0x05, 0xa6, 0x12, 0x73, 0xc5, 0x94, 0x46,
	0xae, 0x4b, 0x91, 0x15, 0x6b, 0xf4, 0x64, 0x2e, 0xb4, 0xa0, 0x20, 0x4b, 0xaf, 0x3c, 0x8e, 0x50,
	0x87, 0xc7, 0xce, 0xf5, 0x54, 0x88, 0x34, 0x43, 0x3f, 0x94, 0xcc, 0x0f, 0x39, 0x17, 0x3a, 0xd4,
	0x4c, 0x70, 0x55, 0xbf, 0xe9, 0x4c, 0x2b, 0x38, 0xd1, 0x1f, 0x24, 0x2a, 0xdf, 0x3c, 0x6b, 0xdc,
	0xfd, 0x4e, 0x60, 0x7f, 0xb9, 0x0a, 0x30, 0xad, 0x6e, 0xcf, 0x03, 0x7c, 0x5f, 0xa0, 0xd2, 0xd4,
	0x86, 0x51, 0x9c, 0x15, 0x15, 0x62, 0x93, 0x19, 0x99, 0x8f, 0x83, 0x36, 0xa4, 0x14, 0x76, 0x78,
	0xb8, 0x46, 0x7b, 0x60, 0x60, 0x73, 0xa6, 0x87, 0x00, 0x2c, 0x41, 0xae, 0xd9, 0x09, 0xc3, 0xdc,
	0x1e, 0x9a, 0x4c, 0x0f, 0xa1, 0x53, 0xb0, 0x64, 0x56, 0xa4, 0x8c, 0xdb, 0x3b, 0x26, 0xd7, 0x44,
	0xf4, 0x1a, 0x8c, 0x14, 0xfb, 0x88, 0x6f, 0xd2, 0xc8, 0xde, 0x9d, 0x91, 0xf9, 0x30, 0xb0, 0xaa,
	0xf0, 0x61, 0x44, 0x1d, 0xd8, 0x43, 0x9e, 0x48, 0xc1, 0xb8, 0xb6, 0x2d, 0x53, 0xd2, 0xc5, 0xee,
	0x7d, 0xb8, 0xba, 0x5c, 0xbd, 0xe4, 0xf9, 0xff, 0x30, 0x76, 0xef, 0x56, 0xa2, 0x1f, 0xa0, 0x8a,
	0x73, 0x16, 0xe1, 0xbf, 0x5d, 0xf1, 0x83, 0x80, 0xb5, 0x5c, 0x3d, 0xe2, 0x27, 0xa2, 0x4b, 0x93,
	0x9e, 0x27, 0x3d, 0x6d, 0x83, 0x0d, 0x6d, 0x53, 0xb0, 0x94, 0x0e, 0x75, 0xa1, 0x1a, 0xa3, 0x9a,
	0xa8, 0xc2, 0xeb, 0xd6, 0xb6, 0x26, 0xd5, 0x11, 0x3d, 0x80, 0xdd, 0x38, 0x0b, 0xd9, 0xda, 0x58,
	0x34, 0x0e, 0xea, 0xa0, 0x67, 0xa9, 0xb5, 0x61, 0xe9, 0x0c, 0x26, 0x61, 0x1c, 0xa3, 0x52, 0x4f,
	0x44, 0x82, 0xca, 0x1e, 0xcd, 0x86, 0xf3, 0x71, 0xd0, 0x87, 0xdc, 0xb7, 0x40, 0xfb, 0xd2, 0x95,
	0x14, 0x5c, 0x21, 0xbd, 0xd5, 0xb1, 0xaa, 0x44, 0x4c, 0x16, 0x57, 0xbc, 0x7a, 0x56, 0xbc, 0xe7,
	0x06, 0xed, 0x58, 0xba, 0x30, 0x90, 0xa5, 0x51, 0x34, 0x59, 0x50, 0xef, 0x7c, 0xfa, 0xbc, 0xda,
	0x8a, 0x60, 0x20, 0xcb, 0xc5, 0xef, 0x21, 0xec, 0x2f, 0xbb, 0x79, 0x5d, 0x19, 0x19, 0x8a, 0x7e,
	0x23, 0xb0, 0xd7, 0x7e, 0x96, 0xde, 0xd8, 0x2c, 0xdd, 0xea, 0x84, 0x73, 0x78, 0x51, 0xba, 0x66,
	0xeb, 0xbe, 0xf8, 0xf4, 0xf3, 0xd7, 0xe7, 0xc1, 0x53, 0xfa, 0xd8, 0x0f, 0xa5, 0x54, 0xb1, 0x48,
	0xea, 0xa9, 0x7f, 0x57, 0x44, 0x98, 0x73, 0xd4, 0xa8, 0xfc, 0xa6, 0xd8, 0x6f, 0x3a, 0xa8, 0xfc,
	0xd3, 0xe6, 0x74, 0xe6, 0x6f, 0xaf, 0x92, 0xf2, 0x4f, 0xab, 0x8e, 0x9d, 0xd1, 0x2f, 0x04, 0xf6,
	0xda, 0x45, 0xd8, 0x66, 0xb8, 0xb5, 0x20, 0xce, 0x41, 0xeb, 0xcf, 0x4a, 0xb0, 0xa4, 0xe3, 0xf5,
	0xca, 0xf0, 0x7a, 0xe6, 0x5c, 0x2a, 0xaf, 0x3b, 0xe4, 0x88, 0x7e, 0x25, 0x00, 0xe7, 0x33, 0x4f,
	0x6f, 0x6e, 0x92, 0xfb, 0x6b, 0x1b, 0x2e, 0xa0, 0xd7, 0xd8, 0x76, 0x74, 0xa9, 0xf4, 0xee, 0x8d,
	0x5f, 0x8f, 0x9a, 0xb2, 0xc8, 0x32, 0xff, 0x94, 0xdb, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9c,
	0xb7, 0x12, 0x64, 0xaf, 0x04, 0x00, 0x00,
}
