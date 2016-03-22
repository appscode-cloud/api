// Code generated by protoc-gen-go.
// source: authentication.proto
// DO NOT EDIT!

/*
Package authentication is a generated protocol buffer package.

It is generated from these files:
	authentication.proto

It has these top-level messages:
	ValidateRequest
	ValidateResponse
*/
package authentication

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

// Next Id 4
type ValidateRequest struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Secret    string `protobuf:"bytes,3,opt,name=secret" json:"secret,omitempty"`
}

func (m *ValidateRequest) Reset()                    { *m = ValidateRequest{} }
func (m *ValidateRequest) String() string            { return proto.CompactTextString(m) }
func (*ValidateRequest) ProtoMessage()               {}
func (*ValidateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ValidateResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *ValidateResponse) Reset()                    { *m = ValidateResponse{} }
func (m *ValidateResponse) String() string            { return proto.CompactTextString(m) }
func (*ValidateResponse) ProtoMessage()               {}
func (*ValidateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ValidateResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*ValidateRequest)(nil), "authentication.ValidateRequest")
	proto.RegisterType((*ValidateResponse)(nil), "authentication.ValidateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for Authentication service

type AuthenticationClient interface {
	// This rpc is used to check a valid user from other applications.
	User(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
	// appctl used this to validates the user token with phabricator.
	Token(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
}

type authenticationClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationClient(cc *grpc.ClientConn) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) User(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := grpc.Invoke(ctx, "/authentication.Authentication/User", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) Token(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := grpc.Invoke(ctx, "/authentication.Authentication/Token", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Authentication service

type AuthenticationServer interface {
	// This rpc is used to check a valid user from other applications.
	User(context.Context, *ValidateRequest) (*ValidateResponse, error)
	// appctl used this to validates the user token with phabricator.
	Token(context.Context, *ValidateRequest) (*ValidateResponse, error)
}

func RegisterAuthenticationServer(s *grpc.Server, srv AuthenticationServer) {
	s.RegisterService(&_Authentication_serviceDesc, srv)
}

func _Authentication_User_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AuthenticationServer).User(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Authentication_Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(AuthenticationServer).Token(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Authentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "authentication.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "User",
			Handler:    _Authentication_User_Handler,
		},
		{
			MethodName: "Token",
			Handler:    _Authentication_Token_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x10, 0xc7, 0xd9, 0xaa, 0xc5, 0x8e, 0xb0, 0xd6, 0x28, 0xba, 0x2c, 0x5a, 0x4b, 0x4e, 0xe2, 0x61,
	0x23, 0xeb, 0xcd, 0x9b, 0x17, 0x1f, 0xc0, 0xaf, 0x7b, 0xdc, 0x0e, 0xed, 0x6a, 0x4d, 0xe2, 0x66,
	0x22, 0x78, 0xf5, 0x15, 0x7c, 0x34, 0x5f, 0x41, 0x7c, 0x0e, 0xf3, 0x71, 0x90, 0x2d, 0x78, 0xf3,
	0x12, 0xc8, 0x8f, 0x3f, 0xf3, 0x9b, 0xff, 0xc0, 0x9e, 0x74, 0xb4, 0x40, 0x45, 0x6d, 0x23, 0xa9,
	0xd5, 0xaa, 0x32, 0x9d, 0x26, 0xcd, 0xf2, 0x3e, 0x2d, 0x0f, 0xe7, 0x5a, 0xcf, 0x97, 0x28, 0xa4,
	0x69, 0x85, 0x54, 0x4a, 0x53, 0xc4, 0x36, 0xa5, 0xcb, 0xfd, 0x80, 0x67, 0xf4, 0x66, 0xd0, 0x8a,
	0xf8, 0x26, 0xce, 0xaf, 0x60, 0xfb, 0x5e, 0x2e, 0xdb, 0x99, 0x24, 0xbc, 0xc6, 0x17, 0x87, 0x96,
	0xd8, 0x0e, 0x8c, 0x94, 0x7c, 0x46, 0x6b, 0x64, 0x83, 0x45, 0x36, 0xcd, 0x4e, 0x46, 0x6c, 0x0c,
	0x9b, 0xce, 0x62, 0x17, 0x70, 0x31, 0x88, 0x24, 0x87, 0xa1, 0xc5, 0xa6, 0x43, 0x2a, 0xd6, 0xc2,
	0x9f, 0xd7, 0x30, 0xfe, 0x9d, 0x63, 0x8d, 0x17, 0x23, 0x9b, 0xf8, 0x8c, 0xdf, 0xc2, 0xd9, 0x38,
	0x65, 0xab, 0xce, 0xab, 0xb4, 0x40, 0x75, 0x13, 0x69, 0xfd, 0x9d, 0x41, 0x7e, 0xd9, 0x2b, 0xc1,
	0x16, 0xb0, 0x7e, 0xe7, 0x45, 0xec, 0xb8, 0x5a, 0xe9, 0xbc, 0xb2, 0x64, 0x39, 0xfd, 0x3b, 0x90,
	0xec, 0xfc, 0xe8, 0xfd, 0xf3, 0xeb, 0x63, 0x70, 0xc0, 0x59, 0xba, 0x88, 0x4f, 0x2b, 0xf1, 0x7a,
	0x26, 0x42, 0x8f, 0x8b, 0xec, 0x94, 0x3d, 0xc2, 0xc6, 0xad, 0x7e, 0x42, 0xf5, 0x1f, 0xaa, 0x49,
	0x54, 0x15, 0x7c, 0xb7, 0xaf, 0xa2, 0x30, 0xdf, 0xbb, 0x1e, 0x86, 0xf1, 0xd6, 0xe7, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x32, 0xb1, 0xa4, 0x11, 0xc9, 0x01, 0x00, 0x00,
}
