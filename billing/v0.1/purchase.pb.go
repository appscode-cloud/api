// Code generated by protoc-gen-go.
// source: purchase.proto
// DO NOT EDIT!

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

type PurchaseBeginRequest struct {
	ProductSku string `protobuf:"bytes,1,opt,name=product_sku,json=productSku" json:"product_sku,omitempty"`
	Count      int64  `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *PurchaseBeginRequest) Reset()                    { *m = PurchaseBeginRequest{} }
func (m *PurchaseBeginRequest) String() string            { return proto.CompactTextString(m) }
func (*PurchaseBeginRequest) ProtoMessage()               {}
func (*PurchaseBeginRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type PurchaseBeginResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Phid   string         `protobuf:"bytes,2,opt,name=phid" json:"phid,omitempty"`
}

func (m *PurchaseBeginResponse) Reset()                    { *m = PurchaseBeginResponse{} }
func (m *PurchaseBeginResponse) String() string            { return proto.CompactTextString(m) }
func (*PurchaseBeginResponse) ProtoMessage()               {}
func (*PurchaseBeginResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *PurchaseBeginResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type PurchaseConfirmRequest struct {
	Phid       string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	ObjectPhid string `protobuf:"bytes,2,opt,name=object_phid,json=objectPhid" json:"object_phid,omitempty"`
}

func (m *PurchaseConfirmRequest) Reset()                    { *m = PurchaseConfirmRequest{} }
func (m *PurchaseConfirmRequest) String() string            { return proto.CompactTextString(m) }
func (*PurchaseConfirmRequest) ProtoMessage()               {}
func (*PurchaseConfirmRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

type PurchaseCloseRequest struct {
	ObjectPhid string `protobuf:"bytes,1,opt,name=object_phid,json=objectPhid" json:"object_phid,omitempty"`
}

func (m *PurchaseCloseRequest) Reset()                    { *m = PurchaseCloseRequest{} }
func (m *PurchaseCloseRequest) String() string            { return proto.CompactTextString(m) }
func (*PurchaseCloseRequest) ProtoMessage()               {}
func (*PurchaseCloseRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func init() {
	proto.RegisterType((*PurchaseBeginRequest)(nil), "billing.PurchaseBeginRequest")
	proto.RegisterType((*PurchaseBeginResponse)(nil), "billing.PurchaseBeginResponse")
	proto.RegisterType((*PurchaseConfirmRequest)(nil), "billing.PurchaseConfirmRequest")
	proto.RegisterType((*PurchaseCloseRequest)(nil), "billing.PurchaseCloseRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Purchase service

type PurchaseClient interface {
	Begin(ctx context.Context, in *PurchaseBeginRequest, opts ...grpc.CallOption) (*PurchaseBeginResponse, error)
	Confirm(ctx context.Context, in *PurchaseConfirmRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Close(ctx context.Context, in *PurchaseCloseRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type purchaseClient struct {
	cc *grpc.ClientConn
}

func NewPurchaseClient(cc *grpc.ClientConn) PurchaseClient {
	return &purchaseClient{cc}
}

func (c *purchaseClient) Begin(ctx context.Context, in *PurchaseBeginRequest, opts ...grpc.CallOption) (*PurchaseBeginResponse, error) {
	out := new(PurchaseBeginResponse)
	err := grpc.Invoke(ctx, "/billing.Purchase/Begin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *purchaseClient) Confirm(ctx context.Context, in *PurchaseConfirmRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/billing.Purchase/Confirm", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *purchaseClient) Close(ctx context.Context, in *PurchaseCloseRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/billing.Purchase/Close", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Purchase service

type PurchaseServer interface {
	Begin(context.Context, *PurchaseBeginRequest) (*PurchaseBeginResponse, error)
	Confirm(context.Context, *PurchaseConfirmRequest) (*dtypes.VoidResponse, error)
	Close(context.Context, *PurchaseCloseRequest) (*dtypes.VoidResponse, error)
}

func RegisterPurchaseServer(s *grpc.Server, srv PurchaseServer) {
	s.RegisterService(&_Purchase_serviceDesc, srv)
}

func _Purchase_Begin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseBeginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PurchaseServer).Begin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/billing.Purchase/Begin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PurchaseServer).Begin(ctx, req.(*PurchaseBeginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Purchase_Confirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseConfirmRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PurchaseServer).Confirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/billing.Purchase/Confirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PurchaseServer).Confirm(ctx, req.(*PurchaseConfirmRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Purchase_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseCloseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PurchaseServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/billing.Purchase/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PurchaseServer).Close(ctx, req.(*PurchaseCloseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Purchase_serviceDesc = grpc.ServiceDesc{
	ServiceName: "billing.Purchase",
	HandlerType: (*PurchaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Begin",
			Handler:    _Purchase_Begin_Handler,
		},
		{
			MethodName: "Confirm",
			Handler:    _Purchase_Confirm_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Purchase_Close_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor3 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x52, 0xc1, 0x4a, 0xeb, 0x40,
	0x14, 0x25, 0x7d, 0xaf, 0xed, 0xeb, 0x14, 0xba, 0x18, 0xfa, 0x4a, 0x09, 0xef, 0x59, 0x89, 0x28,
	0x52, 0x21, 0xd1, 0x8a, 0x08, 0x2e, 0xed, 0x5a, 0x28, 0x29, 0xb8, 0x95, 0x34, 0x19, 0xd3, 0xb1,
	0x71, 0x66, 0xec, 0xcc, 0x28, 0x22, 0x6e, 0xfc, 0x05, 0x3f, 0xcd, 0x4f, 0xd0, 0x0f, 0x71, 0x32,
	0x33, 0xb1, 0xb1, 0xb6, 0xd0, 0x4d, 0x48, 0xce, 0xbd, 0xf7, 0xdc, 0x73, 0xce, 0x0d, 0x68, 0x31,
	0x39, 0x8f, 0xa7, 0x11, 0x47, 0x3e, 0x9b, 0x53, 0x41, 0x61, 0x7d, 0x82, 0xb3, 0x0c, 0x93, 0xd4,
	0xfd, 0x97, 0x52, 0x9a, 0x66, 0x28, 0x88, 0x18, 0x0e, 0x22, 0x42, 0xa8, 0x88, 0x04, 0xa6, 0x84,
	0x9b, 0x36, 0xb7, 0x93, 0xc3, 0x89, 0x78, 0x64, 0x88, 0x07, 0xfa, 0x69, 0x70, 0xef, 0x02, 0xb4,
	0x47, 0x96, 0xf0, 0x1c, 0xa5, 0x98, 0x84, 0xe8, 0x4e, 0x22, 0x2e, 0x60, 0x0f, 0x34, 0x55, 0x43,
	0x22, 0x63, 0x71, 0xc5, 0x67, 0xb2, 0xeb, 0x6c, 0x3b, 0xfb, 0x8d, 0x10, 0x58, 0x68, 0x3c, 0x93,
	0xb0, 0x0d, 0xaa, 0x31, 0x95, 0x44, 0x74, 0x2b, 0xaa, 0xf4, 0x2b, 0x34, 0x1f, 0xde, 0x18, 0xfc,
	0x5d, 0xa2, 0xe3, 0x4c, 0x89, 0x40, 0x70, 0x0f, 0xd4, 0xb8, 0x52, 0x24, 0xb9, 0xa6, 0x6a, 0x0e,
	0x5a, 0xbe, 0x11, 0xe3, 0x8f, 0x35, 0x1a, 0xda, 0x2a, 0x84, 0xe0, 0x37, 0x9b, 0xe2, 0x44, 0xb3,
	0x36, 0x42, 0xfd, 0xae, 0x34, 0x76, 0x0a, 0xd2, 0x21, 0x25, 0xd7, 0x78, 0x7e, 0x5b, 0xa8, 0x2c,
	0xba, 0x9d, 0x45, 0x77, 0xae, 0x9c, 0x4e, 0x6e, 0x90, 0x12, 0x5e, 0x22, 0x02, 0x06, 0x1a, 0xe5,
	0x74, 0xa7, 0x0b, 0xcb, 0xc3, 0x8c, 0x72, 0x54, 0xb2, 0x5c, 0x1e, 0x74, 0x96, 0x07, 0x07, 0xef,
	0x15, 0xf0, 0xa7, 0x98, 0x84, 0x1c, 0x54, 0xb5, 0x43, 0xf8, 0xdf, 0xb7, 0x17, 0xf0, 0x57, 0x05,
	0xe9, 0x6e, 0xad, 0x2b, 0x9b, 0x60, 0xbc, 0x83, 0x97, 0xb7, 0x8f, 0xd7, 0xca, 0xae, 0xb7, 0xa3,
	0x0e, 0xc7, 0x78, 0x4c, 0x13, 0x73, 0x41, 0x3b, 0x14, 0xdc, 0x1f, 0xfa, 0x47, 0x41, 0x71, 0x72,
	0xf8, 0x00, 0xea, 0x36, 0x01, 0xd8, 0xfb, 0xc1, 0xfb, 0x3d, 0x1b, 0xb7, 0x5d, 0x24, 0x7c, 0x49,
	0x71, 0xf2, 0xb5, 0xee, 0x44, 0xaf, 0x0b, 0xdc, 0xfe, 0x06, 0xeb, 0x82, 0xa7, 0x3c, 0x88, 0xe7,
	0x33, 0xa7, 0x0f, 0x31, 0xa8, 0xea, 0xac, 0x56, 0xb8, 0x2d, 0x67, 0xb8, 0x66, 0xa9, 0xf5, 0xd8,
	0xdf, 0xc4, 0xe3, 0xa4, 0xa6, 0x7f, 0xcc, 0xe3, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xbb,
	0xba, 0x0b, 0xe9, 0x02, 0x00, 0x00,
}
