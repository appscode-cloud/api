// Code generated by protoc-gen-go.
// source: pvc.proto
// DO NOT EDIT!

package pv

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

type PVCRegisterRequest struct {
	Cluster   string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Size      int64  `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
	Namespace string `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *PVCRegisterRequest) Reset()                    { *m = PVCRegisterRequest{} }
func (m *PVCRegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*PVCRegisterRequest) ProtoMessage()               {}
func (*PVCRegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type PVCUnregisterRequest struct {
	Cluster   string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *PVCUnregisterRequest) Reset()                    { *m = PVCUnregisterRequest{} }
func (m *PVCUnregisterRequest) String() string            { return proto.CompactTextString(m) }
func (*PVCUnregisterRequest) ProtoMessage()               {}
func (*PVCUnregisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func init() {
	proto.RegisterType((*PVCRegisterRequest)(nil), "pv.PVCRegisterRequest")
	proto.RegisterType((*PVCUnregisterRequest)(nil), "pv.PVCUnregisterRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for PersistentVolumeClaims service

type PersistentVolumeClaimsClient interface {
	Register(ctx context.Context, in *PVCRegisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Unregister(ctx context.Context, in *PVCUnregisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type persistentVolumeClaimsClient struct {
	cc *grpc.ClientConn
}

func NewPersistentVolumeClaimsClient(cc *grpc.ClientConn) PersistentVolumeClaimsClient {
	return &persistentVolumeClaimsClient{cc}
}

func (c *persistentVolumeClaimsClient) Register(ctx context.Context, in *PVCRegisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/pv.PersistentVolumeClaims/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *persistentVolumeClaimsClient) Unregister(ctx context.Context, in *PVCUnregisterRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/pv.PersistentVolumeClaims/Unregister", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PersistentVolumeClaims service

type PersistentVolumeClaimsServer interface {
	Register(context.Context, *PVCRegisterRequest) (*dtypes.VoidResponse, error)
	Unregister(context.Context, *PVCUnregisterRequest) (*dtypes.VoidResponse, error)
}

func RegisterPersistentVolumeClaimsServer(s *grpc.Server, srv PersistentVolumeClaimsServer) {
	s.RegisterService(&_PersistentVolumeClaims_serviceDesc, srv)
}

func _PersistentVolumeClaims_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(PVCRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PersistentVolumeClaimsServer).Register(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _PersistentVolumeClaims_Unregister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(PVCUnregisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(PersistentVolumeClaimsServer).Unregister(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _PersistentVolumeClaims_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pv.PersistentVolumeClaims",
	HandlerType: (*PersistentVolumeClaimsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _PersistentVolumeClaims_Register_Handler,
		},
		{
			MethodName: "Unregister",
			Handler:    _PersistentVolumeClaims_Unregister_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor1 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x91, 0xdf, 0x4a, 0xfb, 0x30,
	0x14, 0xc7, 0xe9, 0x3a, 0x7e, 0x3f, 0x9b, 0xcb, 0x30, 0x4a, 0x28, 0xbb, 0xd0, 0x0a, 0xa2, 0x43,
	0x12, 0xff, 0xdc, 0x79, 0xbb, 0x17, 0x18, 0x05, 0x7b, 0x9f, 0x75, 0xc7, 0x12, 0x68, 0x93, 0xd8,
	0xa4, 0x05, 0x1d, 0xbb, 0xf1, 0x15, 0x7c, 0x34, 0x5f, 0xc1, 0x1b, 0xdf, 0xc2, 0x26, 0xdd, 0x18,
	0x3a, 0xc4, 0x0b, 0x6f, 0xda, 0xd3, 0x4f, 0x0f, 0xf9, 0x9c, 0xf3, 0x0d, 0x8a, 0x74, 0x57, 0x50,
	0xdd, 0x28, 0xab, 0xf0, 0x48, 0x77, 0xc9, 0xb4, 0x54, 0xaa, 0xac, 0x80, 0x71, 0x2d, 0x18, 0x97,
	0x52, 0x59, 0x6e, 0x85, 0x92, 0x66, 0xe8, 0x48, 0x62, 0x87, 0x57, 0xf6, 0x49, 0x83, 0x61, 0xfe,
	0x39, 0xf0, 0xd4, 0x22, 0xbc, 0xc8, 0xe7, 0x19, 0x94, 0xc2, 0x58, 0x68, 0x32, 0x78, 0x6c, 0xc1,
	0x58, 0x4c, 0xd0, 0xff, 0xa2, 0x6a, 0x1d, 0x21, 0xc1, 0x71, 0x70, 0x1e, 0x65, 0xbb, 0x4f, 0x8c,
	0xd1, 0x58, 0xf2, 0x1a, 0xc8, 0xc8, 0x63, 0x5f, 0x3b, 0x66, 0xc4, 0x33, 0x90, 0xb0, 0x67, 0x61,
	0xe6, 0x6b, 0x3c, 0x45, 0x91, 0xfb, 0x67, 0x34, 0x2f, 0x80, 0x8c, 0x7d, 0xf3, 0x1e, 0xa4, 0x4b,
	0x34, 0xe9, 0xad, 0xf7, 0xb2, 0xf9, 0x93, 0xf7, 0x8b, 0x23, 0xfc, 0xe6, 0xb8, 0xf9, 0x08, 0x50,
	0xbc, 0x80, 0xc6, 0x38, 0x81, 0xb4, 0xb9, 0xaa, 0xda, 0x1a, 0xe6, 0x15, 0x17, 0xb5, 0xc1, 0x0f,
	0xe8, 0x68, 0xb7, 0x31, 0x8e, 0xa9, 0xee, 0xe8, 0x61, 0x04, 0xc9, 0x84, 0x0e, 0x69, 0xd1, 0x5c,
	0x89, 0x55, 0xd6, 0x9f, 0xd9, 0x67, 0x09, 0xe9, 0xe5, 0xcb, 0xdb, 0xfb, 0xeb, 0xe8, 0x2c, 0x39,
	0xf1, 0x31, 0xeb, 0x8e, 0x75, 0x57, 0xf4, 0xba, 0x7f, 0x17, 0x6c, 0xbd, 0x9d, 0x74, 0xc3, 0xd6,
	0x6e, 0x86, 0xcd, 0x5d, 0x30, 0xc3, 0x25, 0x42, 0xfb, 0x1d, 0x31, 0xd9, 0x9a, 0x0e, 0xd6, 0xfe,
	0xc1, 0x75, 0xe1, 0x5d, 0xa7, 0xb3, 0xdf, 0x5d, 0xcb, 0x7f, 0xfe, 0x32, 0x6f, 0x3f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x7e, 0x1e, 0x00, 0xaf, 0x13, 0x02, 0x00, 0x00,
}
