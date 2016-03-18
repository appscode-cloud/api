// Code generated by protoc-gen-go.
// source: master.proto
// DO NOT EDIT!

package ci

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

type MasterCreateRequest struct {
	VolumeId string `protobuf:"bytes,1,opt,name=volume_id" json:"volume_id,omitempty"`
}

func (m *MasterCreateRequest) Reset()                    { *m = MasterCreateRequest{} }
func (m *MasterCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*MasterCreateRequest) ProtoMessage()               {}
func (*MasterCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type MasterCreateResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *MasterCreateResponse) Reset()                    { *m = MasterCreateResponse{} }
func (m *MasterCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*MasterCreateResponse) ProtoMessage()               {}
func (*MasterCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *MasterCreateResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*MasterCreateRequest)(nil), "ci.MasterCreateRequest")
	proto.RegisterType((*MasterCreateResponse)(nil), "ci.MasterCreateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for Master service

type MasterClient interface {
	Create(ctx context.Context, in *MasterCreateRequest, opts ...grpc.CallOption) (*MasterCreateResponse, error)
}

type masterClient struct {
	cc *grpc.ClientConn
}

func NewMasterClient(cc *grpc.ClientConn) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) Create(ctx context.Context, in *MasterCreateRequest, opts ...grpc.CallOption) (*MasterCreateResponse, error) {
	out := new(MasterCreateResponse)
	err := grpc.Invoke(ctx, "/ci.Master/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Master service

type MasterServer interface {
	Create(context.Context, *MasterCreateRequest) (*MasterCreateResponse, error)
}

func RegisterMasterServer(s *grpc.Server, srv MasterServer) {
	s.RegisterService(&_Master_serviceDesc, srv)
}

func _Master_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(MasterCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(MasterServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Master_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ci.Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Master_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor2 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x4d, 0x2c, 0x2e,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4a, 0xce, 0x94, 0x92, 0x49, 0xcf,
	0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c,
	0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0xa8, 0x90, 0x12, 0x03, 0x09, 0xa7, 0x94, 0x54, 0x16, 0xa4, 0x16,
	0xeb, 0x83, 0x49, 0x88, 0xb8, 0x92, 0x06, 0x97, 0xb0, 0x2f, 0xd8, 0x24, 0xe7, 0xa2, 0xd4, 0xc4,
	0x92, 0xd4, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x41, 0x2e, 0xce, 0xb2, 0xfc, 0x9c,
	0xd2, 0xdc, 0xd4, 0xf8, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x4e, 0x25, 0x33, 0x2e, 0x11,
	0x54, 0x95, 0xc5, 0x05, 0x40, 0xe3, 0x53, 0x85, 0xe4, 0xb8, 0xd8, 0x8a, 0x81, 0x76, 0x95, 0x16,
	0x83, 0xd5, 0x71, 0x1b, 0xf1, 0xe9, 0x41, 0xac, 0xd1, 0x0b, 0x06, 0x8b, 0x1a, 0x25, 0x73, 0xb1,
	0x41, 0xf4, 0x09, 0x45, 0x72, 0xb1, 0x41, 0xf4, 0x0a, 0x89, 0xeb, 0x25, 0x67, 0xea, 0x61, 0xb1,
	0x57, 0x4a, 0x02, 0x53, 0x02, 0x62, 0x8d, 0x92, 0x4c, 0xd3, 0xe5, 0x27, 0x93, 0x99, 0xc4, 0xa4,
	0x04, 0xc1, 0x1e, 0x4c, 0xce, 0xd4, 0x2f, 0x33, 0xd0, 0x87, 0x84, 0x81, 0x15, 0xa3, 0x56, 0x12,
	0x1b, 0xd8, 0x37, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x68, 0xac, 0xfe, 0x17, 0x01,
	0x00, 0x00,
}
