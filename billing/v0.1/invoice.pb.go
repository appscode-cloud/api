// Code generated by protoc-gen-go.
// source: invoice.proto
// DO NOT EDIT!

package billing

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

type InvoiceCreateRequest struct {
	TimeUnit string `protobuf:"bytes,1,opt,name=time_unit,json=timeUnit" json:"time_unit,omitempty"`
}

func (m *InvoiceCreateRequest) Reset()                    { *m = InvoiceCreateRequest{} }
func (m *InvoiceCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*InvoiceCreateRequest) ProtoMessage()               {}
func (*InvoiceCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type InvoiceCreateResponse struct {
	Status  *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	CostUsd string         `protobuf:"bytes,2,opt,name=cost_usd,json=costUsd" json:"cost_usd,omitempty"`
}

func (m *InvoiceCreateResponse) Reset()                    { *m = InvoiceCreateResponse{} }
func (m *InvoiceCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*InvoiceCreateResponse) ProtoMessage()               {}
func (*InvoiceCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *InvoiceCreateResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*InvoiceCreateRequest)(nil), "billing.InvoiceCreateRequest")
	proto.RegisterType((*InvoiceCreateResponse)(nil), "billing.InvoiceCreateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Invoice service

type InvoiceClient interface {
	Create(ctx context.Context, in *InvoiceCreateRequest, opts ...grpc.CallOption) (*InvoiceCreateResponse, error)
}

type invoiceClient struct {
	cc *grpc.ClientConn
}

func NewInvoiceClient(cc *grpc.ClientConn) InvoiceClient {
	return &invoiceClient{cc}
}

func (c *invoiceClient) Create(ctx context.Context, in *InvoiceCreateRequest, opts ...grpc.CallOption) (*InvoiceCreateResponse, error) {
	out := new(InvoiceCreateResponse)
	err := grpc.Invoke(ctx, "/billing.Invoice/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Invoice service

type InvoiceServer interface {
	Create(context.Context, *InvoiceCreateRequest) (*InvoiceCreateResponse, error)
}

func RegisterInvoiceServer(s *grpc.Server, srv InvoiceServer) {
	s.RegisterService(&_Invoice_serviceDesc, srv)
}

func _Invoice_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InvoiceCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InvoiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/billing.Invoice/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InvoiceServer).Create(ctx, req.(*InvoiceCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Invoice_serviceDesc = grpc.ServiceDesc{
	ServiceName: "billing.Invoice",
	HandlerType: (*InvoiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Invoice_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("invoice.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0xcc, 0x2b, 0xcb,
	0xcf, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xca, 0xcc, 0xc9, 0xc9,
	0xcc, 0x4b, 0x97, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f,
	0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x28, 0x93, 0x12, 0x03, 0x09,
	0xa7, 0x94, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0x83, 0x49, 0x88, 0xb8, 0x92, 0x31, 0x97, 0x88, 0x27,
	0xc4, 0x3c, 0xe7, 0xa2, 0xd4, 0xc4, 0x92, 0xd4, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21,
	0x69, 0x2e, 0xce, 0x92, 0xcc, 0xdc, 0xd4, 0xf8, 0xd2, 0xbc, 0xcc, 0x12, 0x09, 0x46, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x0e, 0x90, 0x40, 0x28, 0x90, 0xaf, 0x14, 0xc5, 0x25, 0x8a, 0xa6, 0xa9, 0xb8,
	0x00, 0x68, 0x55, 0xaa, 0x90, 0x1a, 0x17, 0x5b, 0x31, 0xd0, 0xde, 0xd2, 0x62, 0xb0, 0x16, 0x6e,
	0x23, 0x3e, 0x3d, 0x88, 0x95, 0x7a, 0xc1, 0x60, 0xd1, 0x20, 0xa8, 0xac, 0x90, 0x24, 0x17, 0x47,
	0x72, 0x7e, 0x71, 0x49, 0x7c, 0x69, 0x71, 0x8a, 0x04, 0x13, 0xd8, 0x70, 0x76, 0x10, 0x3f, 0xb4,
	0x38, 0xc5, 0xa8, 0x8e, 0x8b, 0x1d, 0x6a, 0xb6, 0x50, 0x31, 0x17, 0x1b, 0xc4, 0x7c, 0x21, 0x59,
	0x3d, 0xa8, 0x2f, 0xf5, 0xb0, 0x39, 0x56, 0x4a, 0x0e, 0x97, 0x34, 0xc4, 0x59, 0x4a, 0x5a, 0x4d,
	0x97, 0x9f, 0x4c, 0x66, 0x52, 0x51, 0x52, 0x02, 0x06, 0x4e, 0x41, 0x71, 0x72, 0x7e, 0x0a, 0x24,
	0x94, 0xa0, 0x9a, 0xf4, 0xcb, 0x0c, 0xf4, 0x0c, 0xf5, 0xa1, 0xa1, 0x9a, 0xc4, 0x06, 0x0e, 0x17,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x6d, 0x3c, 0xda, 0x67, 0x01, 0x00, 0x00,
}
