// Code generated by protoc-gen-go. DO NOT EDIT.
// source: authentication.proto

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	authentication.proto
	conduit.proto
	project.proto

It has these top-level messages:
	LoginRequest
	LoginResponse
	CSRFTokenResponse
	ConduitWhoAmIResponse
	ConduitUsersResponse
	ConduitUser
	Preferences
	ProjectListRequest
	ProjectListResponse
	ProjectMemberListRequest
	ProjectMemberListResponse
	Project
	Member
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "appscode.com/api/dtypes"

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

type LoginRequest struct {
	Namespace  string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Username   string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password   string `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Token      string `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
	IssueToken bool   `protobuf:"varint,5,opt,name=issue_token,json=issueToken" json:"issue_token,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginRequest) GetIssueToken() bool {
	if m != nil {
		return m.IssueToken
	}
	return false
}

type LoginResponse struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type CSRFTokenResponse struct {
	CsrfToken string `protobuf:"bytes,1,opt,name=csrf_token,json=csrfToken" json:"csrf_token,omitempty"`
}

func (m *CSRFTokenResponse) Reset()                    { *m = CSRFTokenResponse{} }
func (m *CSRFTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*CSRFTokenResponse) ProtoMessage()               {}
func (*CSRFTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CSRFTokenResponse) GetCsrfToken() string {
	if m != nil {
		return m.CsrfToken
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "appscode.auth.v1beta1.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "appscode.auth.v1beta1.LoginResponse")
	proto.RegisterType((*CSRFTokenResponse)(nil), "appscode.auth.v1beta1.CSRFTokenResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Authentication service

type AuthenticationClient interface {
	// This rpc is used to check a valid user from other applications.
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	CSRFToken(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*CSRFTokenResponse, error)
}

type authenticationClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationClient(cc *grpc.ClientConn) AuthenticationClient {
	return &authenticationClient{cc}
}

func (c *authenticationClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/appscode.auth.v1beta1.Authentication/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) Logout(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.auth.v1beta1.Authentication/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationClient) CSRFToken(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*CSRFTokenResponse, error) {
	out := new(CSRFTokenResponse)
	err := grpc.Invoke(ctx, "/appscode.auth.v1beta1.Authentication/CSRFToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Authentication service

type AuthenticationServer interface {
	// This rpc is used to check a valid user from other applications.
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *appscode_dtypes.VoidRequest) (*appscode_dtypes.VoidResponse, error)
	CSRFToken(context.Context, *appscode_dtypes.VoidRequest) (*CSRFTokenResponse, error)
}

func RegisterAuthenticationServer(s *grpc.Server, srv AuthenticationServer) {
	s.RegisterService(&_Authentication_serviceDesc, srv)
}

func _Authentication_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.auth.v1beta1.Authentication/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.auth.v1beta1.Authentication/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).Logout(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authentication_CSRFToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServer).CSRFToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.auth.v1beta1.Authentication/CSRFToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServer).CSRFToken(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authentication_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.auth.v1beta1.Authentication",
	HandlerType: (*AuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Authentication_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Authentication_Logout_Handler,
		},
		{
			MethodName: "CSRFToken",
			Handler:    _Authentication_CSRFToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authentication.proto",
}

func init() { proto.RegisterFile("authentication.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xbb, 0x8e, 0xd3, 0x40,
	0x14, 0xd5, 0x64, 0xc9, 0x12, 0x5f, 0x1e, 0x12, 0xc3, 0x22, 0x59, 0x56, 0x02, 0x2b, 0x2f, 0x8f,
	0x88, 0x65, 0x6d, 0xed, 0x52, 0xf0, 0x68, 0x10, 0x8b, 0x44, 0xb5, 0x45, 0x64, 0x10, 0x05, 0x4d,
	0x34, 0xb1, 0x07, 0xc7, 0x90, 0xcc, 0x1d, 0x3c, 0x63, 0x10, 0x15, 0x52, 0x5a, 0x24, 0x1a, 0x4a,
	0x7e, 0x82, 0x86, 0x2f, 0xe1, 0x17, 0xf8, 0x10, 0x34, 0x63, 0xc7, 0xb1, 0xc5, 0x26, 0x69, 0x2c,
	0xcd, 0x3d, 0xe7, 0xdc, 0x73, 0x74, 0x8f, 0x61, 0x8f, 0x15, 0x7a, 0xca, 0x85, 0xce, 0x62, 0xa6,
	0x33, 0x14, 0x81, 0xcc, 0x51, 0x23, 0xbd, 0xc1, 0xa4, 0x54, 0x31, 0x26, 0x3c, 0x30, 0x70, 0xf0,
	0xe9, 0x78, 0xc2, 0x35, 0x3b, 0xf6, 0xfa, 0x29, 0x62, 0x3a, 0xe3, 0x21, 0x93, 0x59, 0xc8, 0x84,
	0x40, 0x6d, 0x35, 0xaa, 0x14, 0x79, 0x37, 0x97, 0xa2, 0x35, 0xf8, 0x41, 0xbd, 0x34, 0xc6, 0xb9,
	0xe5, 0x24, 0xfa, 0x8b, 0xe4, 0x2a, 0xb4, 0xdf, 0x92, 0xe4, 0xff, 0x24, 0x70, 0xf9, 0x0c, 0xd3,
	0x4c, 0x44, 0xfc, 0x63, 0xc1, 0x95, 0xa6, 0x7d, 0x70, 0x04, 0x9b, 0x73, 0x25, 0x59, 0xcc, 0x5d,
	0xb2, 0x4f, 0x86, 0x4e, 0xb4, 0x1a, 0x50, 0x0f, 0x7a, 0x85, 0xe2, 0xb9, 0x19, 0xb8, 0x1d, 0x0b,
	0xd6, 0x6f, 0x83, 0x49, 0xa6, 0xd4, 0x67, 0xcc, 0x13, 0x77, 0xa7, 0xc4, 0x96, 0x6f, 0xba, 0x07,
	0x5d, 0x8d, 0x1f, 0xb8, 0x70, 0x2f, 0x58, 0xa0, 0x7c, 0xd0, 0x5b, 0x70, 0x29, 0x53, 0xaa, 0xe0,
	0xe3, 0x12, 0xeb, 0xee, 0x93, 0x61, 0x2f, 0x02, 0x3b, 0x7a, 0x6d, 0x26, 0xfe, 0x1d, 0xb8, 0x52,
	0x85, 0x53, 0x12, 0x85, 0xe2, 0xab, 0x3d, 0xa4, 0xb1, 0xc7, 0x3f, 0x81, 0x6b, 0x2f, 0x5e, 0x45,
	0x2f, 0xad, 0xa6, 0xa6, 0x0e, 0x00, 0x62, 0x95, 0xbf, 0x1b, 0x37, 0xf9, 0x8e, 0x99, 0x58, 0xda,
	0xc9, 0xaf, 0x1d, 0xb8, 0xfa, 0xbc, 0xd5, 0x05, 0xfd, 0x46, 0xa0, 0x6b, 0xed, 0xe8, 0x41, 0x70,
	0x6e, 0x21, 0x41, 0xf3, 0x52, 0xde, 0xed, 0xcd, 0xa4, 0x32, 0x86, 0xff, 0x68, 0xf1, 0xdb, 0xed,
	0xf4, 0xc8, 0xe2, 0xcf, 0xdf, 0x1f, 0x9d, 0x43, 0xff, 0x6e, 0x38, 0x6e, 0xb7, 0x56, 0xe8, 0x69,
	0x58, 0x09, 0xc3, 0x99, 0x11, 0x86, 0xef, 0x15, 0x8a, 0xa7, 0xe4, 0x3e, 0xfd, 0x0a, 0xbb, 0x67,
	0x98, 0x62, 0xa1, 0x69, 0x7f, 0x65, 0x54, 0x36, 0x18, 0xbc, 0xc1, 0x2c, 0x59, 0xc6, 0x18, 0xac,
	0x41, 0x2b, 0xff, 0xc7, 0x0d, 0xff, 0x07, 0xfe, 0xbd, 0x2d, 0xfe, 0x58, 0xe8, 0x3a, 0xc0, 0x77,
	0x02, 0x4e, 0x7d, 0xd6, 0x2d, 0x21, 0x86, 0x6b, 0x6e, 0xf1, 0x5f, 0x2d, 0xfe, 0x93, 0x46, 0x9e,
	0x23, 0x7a, 0xb8, 0x29, 0x8f, 0x29, 0xeb, 0xc8, 0xd6, 0x67, 0x33, 0x9d, 0x3e, 0x83, 0x41, 0x8c,
	0xf3, 0x86, 0x93, 0xcc, 0x5a, 0x6e, 0xa7, 0xd7, 0xdb, 0x85, 0x8e, 0xcc, 0x1f, 0x3e, 0x22, 0x6f,
	0x2f, 0x56, 0xf8, 0x64, 0xd7, 0xfe, 0xf3, 0x0f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x97, 0xb7,
	0x0a, 0x4d, 0x85, 0x03, 0x00, 0x00,
}
