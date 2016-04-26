// Code generated by protoc-gen-go.
// source: charge.proto
// DO NOT EDIT!

/*
Package billing is a generated protocol buffer package.

It is generated from these files:
	charge.proto
	subscription.proto

It has these top-level messages:
	ChargeCalculateRequest
	ChargeCalculateResponse
	SubscriptionDescribeRequest
	SubscriptionDescribeResponse
	Subscription
	SubscriptionOpenRequest
	SubscriptionCloseRequest
	SubscriptionCreateRequest
	Quota
	PhabricatorQuota
	CIQuota
	ArtifactoryQuota
	ClusterQuota
	DBQuota
	SubscriptionQutaRequest
	SubscriptionQutaResponse
*/
package billing

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
const _ = proto.ProtoPackageIsVersion1

type ChargeCalculateRequest struct {
	ChargeType string `protobuf:"bytes,1,opt,name=charge_type,json=chargeType" json:"charge_type,omitempty"`
}

func (m *ChargeCalculateRequest) Reset()                    { *m = ChargeCalculateRequest{} }
func (m *ChargeCalculateRequest) String() string            { return proto.CompactTextString(m) }
func (*ChargeCalculateRequest) ProtoMessage()               {}
func (*ChargeCalculateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ChargeCalculateResponse struct {
	Status  *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	CostUsd string         `protobuf:"bytes,2,opt,name=cost_usd,json=costUsd" json:"cost_usd,omitempty"`
}

func (m *ChargeCalculateResponse) Reset()                    { *m = ChargeCalculateResponse{} }
func (m *ChargeCalculateResponse) String() string            { return proto.CompactTextString(m) }
func (*ChargeCalculateResponse) ProtoMessage()               {}
func (*ChargeCalculateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ChargeCalculateResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*ChargeCalculateRequest)(nil), "billing.ChargeCalculateRequest")
	proto.RegisterType((*ChargeCalculateResponse)(nil), "billing.ChargeCalculateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Charge service

type ChargeClient interface {
	Calculate(ctx context.Context, in *ChargeCalculateRequest, opts ...grpc.CallOption) (*ChargeCalculateResponse, error)
}

type chargeClient struct {
	cc *grpc.ClientConn
}

func NewChargeClient(cc *grpc.ClientConn) ChargeClient {
	return &chargeClient{cc}
}

func (c *chargeClient) Calculate(ctx context.Context, in *ChargeCalculateRequest, opts ...grpc.CallOption) (*ChargeCalculateResponse, error) {
	out := new(ChargeCalculateResponse)
	err := grpc.Invoke(ctx, "/billing.Charge/Calculate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Charge service

type ChargeServer interface {
	Calculate(context.Context, *ChargeCalculateRequest) (*ChargeCalculateResponse, error)
}

func RegisterChargeServer(s *grpc.Server, srv ChargeServer) {
	s.RegisterService(&_Charge_serviceDesc, srv)
}

func _Charge_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChargeCalculateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChargeServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/billing.Charge/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChargeServer).Calculate(ctx, req.(*ChargeCalculateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Charge_serviceDesc = grpc.ServiceDesc{
	ServiceName: "billing.Charge",
	HandlerType: (*ChargeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _Charge_Calculate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0x48, 0x2c,
	0x4a, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xca, 0xcc, 0xc9, 0xc9, 0xcc,
	0x4b, 0x97, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc,
	0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x28, 0x93, 0x12, 0x03, 0x09, 0xa7,
	0x94, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0x83, 0x49, 0x88, 0xb8, 0x92, 0x25, 0x97, 0x98, 0x33, 0xd8,
	0x38, 0xe7, 0xc4, 0x9c, 0xe4, 0xd2, 0x9c, 0xc4, 0x92, 0xd4, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2,
	0x12, 0x21, 0x79, 0x2e, 0x6e, 0x88, 0x45, 0xf1, 0x20, 0xf5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x5c, 0x10, 0xa1, 0x10, 0xa0, 0x88, 0x52, 0x0c, 0x97, 0x38, 0x86, 0xd6, 0xe2, 0x02, 0xa0,
	0x95, 0xa9, 0x42, 0x6a, 0x5c, 0x6c, 0xc5, 0x40, 0xfb, 0x4b, 0x8b, 0xc1, 0xda, 0xb8, 0x8d, 0xf8,
	0xf4, 0x20, 0x56, 0xeb, 0x05, 0x83, 0x45, 0x83, 0xa0, 0xb2, 0x42, 0x92, 0x5c, 0x1c, 0xc9, 0xf9,
	0xc5, 0x25, 0xf1, 0xa5, 0xc5, 0x29, 0x12, 0x4c, 0x60, 0x0b, 0xd8, 0x41, 0xfc, 0xd0, 0xe2, 0x14,
	0xa3, 0x76, 0x46, 0x2e, 0x36, 0x88, 0xf1, 0x42, 0xb5, 0x5c, 0x9c, 0x70, 0x2b, 0x84, 0xe4, 0xf5,
	0xa0, 0x1e, 0xd6, 0xc3, 0xee, 0x6e, 0x29, 0x05, 0xdc, 0x0a, 0x20, 0xae, 0x53, 0xd2, 0x6d, 0xba,
	0xfc, 0x64, 0x32, 0x93, 0xba, 0x94, 0x92, 0x3e, 0x54, 0xa5, 0x7e, 0x99, 0x81, 0x9e, 0xa1, 0x3e,
	0xc4, 0x6b, 0xfa, 0xd5, 0x48, 0xbe, 0xae, 0xb5, 0x62, 0xd4, 0x4a, 0x62, 0x03, 0x87, 0x94, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x76, 0x44, 0xc3, 0x72, 0x78, 0x01, 0x00, 0x00,
}
