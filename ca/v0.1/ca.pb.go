// Code generated by protoc-gen-go.
// source: ca.proto
// DO NOT EDIT!

/*
Package ca is a generated protocol buffer package.

It is generated from these files:
	ca.proto

It has these top-level messages:
	CertificateCreateRequest
	CertificateCreateResponse
*/
package ca

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

// Use specific requests for protos
type CertificateCreateRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Csr  string `protobuf:"bytes,2,opt,name=csr" json:"csr,omitempty"`
}

func (m *CertificateCreateRequest) Reset()                    { *m = CertificateCreateRequest{} }
func (m *CertificateCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateCreateRequest) ProtoMessage()               {}
func (*CertificateCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CertificateCreateResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Phid   string         `protobuf:"bytes,2,opt,name=phid" json:"phid,omitempty"`
	Cert   []byte         `protobuf:"bytes,3,opt,name=cert,proto3" json:"cert,omitempty"`
	Key    []byte         `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	Root   []byte         `protobuf:"bytes,5,opt,name=root,proto3" json:"root,omitempty"`
}

func (m *CertificateCreateResponse) Reset()                    { *m = CertificateCreateResponse{} }
func (m *CertificateCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CertificateCreateResponse) ProtoMessage()               {}
func (*CertificateCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CertificateCreateResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*CertificateCreateRequest)(nil), "ca.CertificateCreateRequest")
	proto.RegisterType((*CertificateCreateResponse)(nil), "ca.CertificateCreateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for AuthorityCertificate service

type AuthorityCertificateClient interface {
	Create(ctx context.Context, in *CertificateCreateRequest, opts ...grpc.CallOption) (*CertificateCreateResponse, error)
}

type authorityCertificateClient struct {
	cc *grpc.ClientConn
}

func NewAuthorityCertificateClient(cc *grpc.ClientConn) AuthorityCertificateClient {
	return &authorityCertificateClient{cc}
}

func (c *authorityCertificateClient) Create(ctx context.Context, in *CertificateCreateRequest, opts ...grpc.CallOption) (*CertificateCreateResponse, error) {
	out := new(CertificateCreateResponse)
	err := grpc.Invoke(ctx, "/ca.AuthorityCertificate/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthorityCertificate service

type AuthorityCertificateServer interface {
	Create(context.Context, *CertificateCreateRequest) (*CertificateCreateResponse, error)
}

func RegisterAuthorityCertificateServer(s *grpc.Server, srv AuthorityCertificateServer) {
	s.RegisterService(&_AuthorityCertificate_serviceDesc, srv)
}

func _AuthorityCertificate_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CertificateCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AuthorityCertificateServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _AuthorityCertificate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ca.AuthorityCertificate",
	HandlerType: (*AuthorityCertificateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AuthorityCertificate_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

// Client API for ClientCertificate service

type ClientCertificateClient interface {
	Create(ctx context.Context, in *CertificateCreateRequest, opts ...grpc.CallOption) (*CertificateCreateResponse, error)
}

type clientCertificateClient struct {
	cc *grpc.ClientConn
}

func NewClientCertificateClient(cc *grpc.ClientConn) ClientCertificateClient {
	return &clientCertificateClient{cc}
}

func (c *clientCertificateClient) Create(ctx context.Context, in *CertificateCreateRequest, opts ...grpc.CallOption) (*CertificateCreateResponse, error) {
	out := new(CertificateCreateResponse)
	err := grpc.Invoke(ctx, "/ca.ClientCertificate/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ClientCertificate service

type ClientCertificateServer interface {
	Create(context.Context, *CertificateCreateRequest) (*CertificateCreateResponse, error)
}

func RegisterClientCertificateServer(s *grpc.Server, srv ClientCertificateServer) {
	s.RegisterService(&_ClientCertificate_serviceDesc, srv)
}

func _ClientCertificate_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CertificateCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClientCertificateServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _ClientCertificate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ca.ClientCertificate",
	HandlerType: (*ClientCertificateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ClientCertificate_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x91, 0x41, 0x4a, 0xc4, 0x30,
	0x14, 0x86, 0xe9, 0xcc, 0x38, 0x68, 0x14, 0xd1, 0x28, 0x43, 0xac, 0xa3, 0x68, 0x17, 0x22, 0x2e,
	0x5a, 0xad, 0x3b, 0x57, 0x4a, 0x6f, 0x50, 0x4f, 0x10, 0x33, 0xcf, 0x69, 0x70, 0x6c, 0x6a, 0x92,
	0x0a, 0x75, 0x25, 0xde, 0x40, 0x3c, 0x9a, 0x57, 0xf0, 0x20, 0x26, 0x2f, 0x23, 0x8c, 0xe0, 0xec,
	0xdc, 0x84, 0x9f, 0x3f, 0x8f, 0xff, 0xfb, 0xf3, 0x42, 0x56, 0x05, 0x4f, 0x1b, 0xad, 0xac, 0xa2,
	0x3d, 0xc1, 0xe3, 0xf1, 0x54, 0xa9, 0xe9, 0x0c, 0x32, 0xde, 0xc8, 0x8c, 0xd7, 0xb5, 0xb2, 0xdc,
	0x4a, 0x55, 0x9b, 0x30, 0x11, 0x8f, 0xbc, 0x3d, 0xb1, 0x5d, 0x03, 0x26, 0xc3, 0x33, 0xf8, 0xc9,
	0x35, 0x61, 0x05, 0x68, 0x2b, 0xef, 0xa5, 0xe0, 0x16, 0x0a, 0x0d, 0xee, 0x2c, 0xe1, 0xa9, 0x05,
	0x63, 0x29, 0x25, 0x83, 0x9a, 0x3f, 0x02, 0x8b, 0x8e, 0xa2, 0xd3, 0xb5, 0x12, 0x35, 0xdd, 0x22,
	0x7d, 0x61, 0x34, 0xeb, 0xa1, 0xe5, 0x65, 0xf2, 0x1e, 0x91, 0xbd, 0x3f, 0x22, 0x4c, 0xe3, 0xe0,
	0x40, 0x4f, 0xc8, 0xd0, 0xb8, 0x26, 0xad, 0xc1, 0x94, 0xf5, 0x7c, 0x33, 0x0d, 0x25, 0xd2, 0x5b,
	0x74, 0xcb, 0xf9, 0xad, 0x67, 0x35, 0x95, 0x9c, 0xcc, 0x83, 0x51, 0x7b, 0x4f, 0xb8, 0x60, 0xd6,
	0x77, 0xde, 0x46, 0x89, 0xda, 0xf3, 0x1f, 0xa0, 0x63, 0x03, 0xb4, 0xbc, 0xf4, 0x53, 0x5a, 0x29,
	0xcb, 0x56, 0xc2, 0x94, 0xd7, 0xf9, 0x6b, 0x44, 0x76, 0x6f, 0x5a, 0x5b, 0x29, 0x2d, 0x6d, 0xb7,
	0x50, 0x8e, 0x56, 0x64, 0x18, 0x0a, 0xd2, 0x71, 0xea, 0xb6, 0xb7, 0xec, 0xe9, 0xf1, 0xc1, 0x92,
	0xdb, 0xf0, 0xaa, 0xe4, 0xf8, 0xed, 0xf3, 0xeb, 0xa3, 0xb7, 0x1f, 0x8f, 0x70, 0xdb, 0x82, 0x67,
	0xcf, 0xe7, 0xe9, 0x45, 0xc6, 0x7f, 0x88, 0x57, 0xd1, 0x59, 0xfe, 0x42, 0xb6, 0x8b, 0x99, 0x84,
	0xda, 0x2e, 0xe2, 0xe1, 0x7f, 0xf0, 0x87, 0x88, 0x67, 0xf1, 0xce, 0x2f, 0xbc, 0x40, 0x9c, 0x63,
	0xdf, 0x0d, 0xf1, 0x6f, 0x2f, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x40, 0x7f, 0x18, 0x21,
	0x02, 0x00, 0x00,
}
