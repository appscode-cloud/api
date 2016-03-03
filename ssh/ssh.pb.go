// Code generated by protoc-gen-go.
// source: ssh.proto
// DO NOT EDIT!

/*
Package ssh is a generated protocol buffer package.

It is generated from these files:
	ssh.proto

It has these top-level messages:
	SecureShellGetRequest
	SecureShellGetResponse
	SSHKey
*/
package ssh

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
type SecureShellGetRequest struct {
	JenkinsNamespace string `protobuf:"bytes,1,opt,name=jenkins_namespace" json:"jenkins_namespace,omitempty"`
	ClusterNamespace string `protobuf:"bytes,2,opt,name=cluster_namespace" json:"cluster_namespace,omitempty"`
	ClusterName      string `protobuf:"bytes,3,opt,name=cluster_name" json:"cluster_name,omitempty"`
	InstanceName     string `protobuf:"bytes,4,opt,name=instance_name" json:"instance_name,omitempty"`
}

func (m *SecureShellGetRequest) Reset()                    { *m = SecureShellGetRequest{} }
func (m *SecureShellGetRequest) String() string            { return proto.CompactTextString(m) }
func (*SecureShellGetRequest) ProtoMessage()               {}
func (*SecureShellGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// return phid ?
type SecureShellGetResponse struct {
	Status       *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	SshKey       *SSHKey        `protobuf:"bytes,2,opt,name=ssh_key" json:"ssh_key,omitempty"`
	InstanceAddr string         `protobuf:"bytes,3,opt,name=instance_addr" json:"instance_addr,omitempty"`
	InstancePort int32          `protobuf:"varint,4,opt,name=instance_port" json:"instance_port,omitempty"`
	User         string         `protobuf:"bytes,5,opt,name=user" json:"user,omitempty"`
}

func (m *SecureShellGetResponse) Reset()                    { *m = SecureShellGetResponse{} }
func (m *SecureShellGetResponse) String() string            { return proto.CompactTextString(m) }
func (*SecureShellGetResponse) ProtoMessage()               {}
func (*SecureShellGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SecureShellGetResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *SecureShellGetResponse) GetSshKey() *SSHKey {
	if m != nil {
		return m.SshKey
	}
	return nil
}

type SSHKey struct {
	PublicKey          []byte `protobuf:"bytes,1,opt,name=public_key,proto3" json:"public_key,omitempty"`
	PrivateKey         []byte `protobuf:"bytes,2,opt,name=private_key,proto3" json:"private_key,omitempty"`
	AwsFingerprint     string `protobuf:"bytes,3,opt,name=aws_fingerprint" json:"aws_fingerprint,omitempty"`
	OpensshFingerprint string `protobuf:"bytes,4,opt,name=openssh_fingerprint" json:"openssh_fingerprint,omitempty"`
}

func (m *SSHKey) Reset()                    { *m = SSHKey{} }
func (m *SSHKey) String() string            { return proto.CompactTextString(m) }
func (*SSHKey) ProtoMessage()               {}
func (*SSHKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*SecureShellGetRequest)(nil), "ssh.SecureShellGetRequest")
	proto.RegisterType((*SecureShellGetResponse)(nil), "ssh.SecureShellGetResponse")
	proto.RegisterType((*SSHKey)(nil), "ssh.SSHKey")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for SecureShell service

type SecureShellClient interface {
	SecureShellGet(ctx context.Context, in *SecureShellGetRequest, opts ...grpc.CallOption) (*SecureShellGetResponse, error)
}

type secureShellClient struct {
	cc *grpc.ClientConn
}

func NewSecureShellClient(cc *grpc.ClientConn) SecureShellClient {
	return &secureShellClient{cc}
}

func (c *secureShellClient) SecureShellGet(ctx context.Context, in *SecureShellGetRequest, opts ...grpc.CallOption) (*SecureShellGetResponse, error) {
	out := new(SecureShellGetResponse)
	err := grpc.Invoke(ctx, "/ssh.SecureShell/SecureShellGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SecureShell service

type SecureShellServer interface {
	SecureShellGet(context.Context, *SecureShellGetRequest) (*SecureShellGetResponse, error)
}

func RegisterSecureShellServer(s *grpc.Server, srv SecureShellServer) {
	s.RegisterService(&_SecureShell_serviceDesc, srv)
}

func _SecureShell_SecureShellGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SecureShellGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SecureShellServer).SecureShellGet(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _SecureShell_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ssh.SecureShell",
	HandlerType: (*SecureShellServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SecureShellGet",
			Handler:    _SecureShell_SecureShellGet_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x4e, 0xf3, 0x30,
	0x10, 0xc5, 0xd5, 0xbf, 0x9f, 0x3a, 0xc9, 0x57, 0xc0, 0xa5, 0xa5, 0xa4, 0x15, 0x42, 0x5d, 0xb1,
	0x4a, 0x50, 0x39, 0x04, 0x48, 0xec, 0xc8, 0x01, 0x2a, 0x37, 0x1d, 0xd2, 0xd0, 0x60, 0x1b, 0xdb,
	0x29, 0x2a, 0x4b, 0x8e, 0x00, 0x47, 0xe3, 0x0a, 0x1c, 0x04, 0xc7, 0xae, 0x50, 0x8a, 0xb2, 0x89,
	0xa2, 0xf7, 0x46, 0x7e, 0xbf, 0x37, 0x36, 0xf4, 0x94, 0x5a, 0x87, 0x42, 0x72, 0xcd, 0x49, 0xcb,
	0xfc, 0x06, 0xd3, 0x94, 0xf3, 0x34, 0xc7, 0x88, 0x8a, 0x2c, 0xa2, 0x8c, 0x71, 0x4d, 0x75, 0xc6,
	0x99, 0x72, 0x23, 0xc1, 0xa8, 0x94, 0x57, 0x7a, 0x27, 0x50, 0x45, 0xf6, 0xeb, 0xf4, 0xd9, 0x1b,
	0x0c, 0x63, 0x4c, 0x0a, 0x89, 0xf1, 0x1a, 0xf3, 0xfc, 0x16, 0xf5, 0x03, 0xbe, 0x14, 0xa8, 0x34,
	0x39, 0x87, 0x93, 0x27, 0x64, 0x9b, 0x8c, 0xa9, 0x05, 0xa3, 0xcf, 0xa8, 0x04, 0x4d, 0x70, 0xdc,
	0xb8, 0x6c, 0x5c, 0xf5, 0x4a, 0x2b, 0xc9, 0x0b, 0xa5, 0x51, 0x56, 0xac, 0xa6, 0xb5, 0x4e, 0xc1,
	0xaf, 0x5a, 0xe3, 0x96, 0x55, 0x87, 0xf0, 0xdf, 0x9c, 0xa3, 0x29, 0x4b, 0xd0, 0xc9, 0xed, 0x52,
	0x9e, 0x7d, 0x34, 0x60, 0xf4, 0x37, 0x5c, 0x09, 0xc3, 0x8c, 0xe4, 0x02, 0xba, 0x66, 0x5e, 0x17,
	0xca, 0x46, 0x7a, 0xf3, 0x7e, 0xe8, 0xd8, 0xc3, 0xd8, 0xaa, 0x64, 0x0a, 0xff, 0x4c, 0xe7, 0xc5,
	0x06, 0x77, 0x36, 0xd8, 0x9b, 0x7b, 0x61, 0xb9, 0x8e, 0x38, 0xbe, 0xbb, 0xc7, 0xdd, 0x41, 0x1e,
	0x5d, 0xad, 0x64, 0x0d, 0x86, 0xe0, 0x52, 0x5b, 0x8c, 0x0e, 0xf1, 0xa1, 0x5d, 0x28, 0x94, 0xe3,
	0x8e, 0x85, 0x4a, 0xa1, 0xbb, 0x3f, 0x85, 0x00, 0x88, 0x62, 0x99, 0x67, 0x89, 0x8d, 0x29, 0x39,
	0x7c, 0x32, 0x00, 0x4f, 0xc8, 0x6c, 0x4b, 0x35, 0xfe, 0x66, 0xfb, 0xe4, 0x0c, 0x8e, 0xe8, 0xab,
	0x5a, 0x3c, 0x66, 0x2c, 0x45, 0x69, 0x6c, 0xa6, 0xf7, 0x81, 0x13, 0x18, 0x70, 0x81, 0xac, 0x24,
	0xad, 0x9a, 0xb6, 0xfd, 0x5c, 0x80, 0x57, 0x29, 0x4f, 0x28, 0xf4, 0x0f, 0x77, 0x41, 0x02, 0x57,
	0xa9, 0xee, 0x76, 0x82, 0x49, 0xad, 0xe7, 0x96, 0x37, 0x1b, 0xbd, 0x7f, 0x7d, 0x7f, 0x36, 0x8f,
	0x49, 0xdf, 0xbe, 0x05, 0x33, 0x68, 0x50, 0xa3, 0xed, 0xf5, 0xb2, 0x6b, 0xaf, 0xfc, 0xe6, 0x27,
	0x00, 0x00, 0xff, 0xff, 0x77, 0xf7, 0xca, 0x54, 0x3a, 0x02, 0x00, 0x00,
}
