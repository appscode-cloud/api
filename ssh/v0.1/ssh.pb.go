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
	JenkinsNamespace string `protobuf:"bytes,1,opt,name=jenkins_namespace,json=jenkinsNamespace" json:"jenkins_namespace,omitempty"`
	ClusterNamespace string `protobuf:"bytes,2,opt,name=cluster_namespace,json=clusterNamespace" json:"cluster_namespace,omitempty"`
	ClusterName      string `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	InstanceName     string `protobuf:"bytes,4,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
}

func (m *SecureShellGetRequest) Reset()                    { *m = SecureShellGetRequest{} }
func (m *SecureShellGetRequest) String() string            { return proto.CompactTextString(m) }
func (*SecureShellGetRequest) ProtoMessage()               {}
func (*SecureShellGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// return phid ?
type SecureShellGetResponse struct {
	Status       *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	SshKey       *SSHKey        `protobuf:"bytes,2,opt,name=ssh_key,json=sshKey" json:"ssh_key,omitempty"`
	InstanceAddr string         `protobuf:"bytes,3,opt,name=instance_addr,json=instanceAddr" json:"instance_addr,omitempty"`
	InstancePort int32          `protobuf:"varint,4,opt,name=instance_port,json=instancePort" json:"instance_port,omitempty"`
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
	PublicKey          []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	PrivateKey         []byte `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	AwsFingerprint     string `protobuf:"bytes,3,opt,name=aws_fingerprint,json=awsFingerprint" json:"aws_fingerprint,omitempty"`
	OpensshFingerprint string `protobuf:"bytes,4,opt,name=openssh_fingerprint,json=opensshFingerprint" json:"openssh_fingerprint,omitempty"`
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
	Get(ctx context.Context, in *SecureShellGetRequest, opts ...grpc.CallOption) (*SecureShellGetResponse, error)
}

type secureShellClient struct {
	cc *grpc.ClientConn
}

func NewSecureShellClient(cc *grpc.ClientConn) SecureShellClient {
	return &secureShellClient{cc}
}

func (c *secureShellClient) Get(ctx context.Context, in *SecureShellGetRequest, opts ...grpc.CallOption) (*SecureShellGetResponse, error) {
	out := new(SecureShellGetResponse)
	err := grpc.Invoke(ctx, "/ssh.SecureShell/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SecureShell service

type SecureShellServer interface {
	Get(context.Context, *SecureShellGetRequest) (*SecureShellGetResponse, error)
}

func RegisterSecureShellServer(s *grpc.Server, srv SecureShellServer) {
	s.RegisterService(&_SecureShell_serviceDesc, srv)
}

func _SecureShell_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SecureShellGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SecureShellServer).Get(ctx, in)
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
			MethodName: "Get",
			Handler:    _SecureShell_Get_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x8e, 0xd3, 0x30,
	0x14, 0x87, 0x95, 0xe9, 0x4c, 0x51, 0x5f, 0xc2, 0x30, 0x18, 0x31, 0x8a, 0x02, 0x08, 0x08, 0x08,
	0x90, 0x90, 0x12, 0x28, 0x27, 0x60, 0x03, 0x48, 0x48, 0x08, 0x25, 0x1b, 0x76, 0x95, 0x9b, 0x3e,
	0xd2, 0xd0, 0x60, 0x1b, 0xdb, 0x29, 0xea, 0x96, 0x2b, 0x70, 0x03, 0x6e, 0xc1, 0x9e, 0x1b, 0x70,
	0x05, 0x0e, 0x82, 0xff, 0xa4, 0x6d, 0x3a, 0xea, 0xa6, 0xb5, 0xbe, 0xf7, 0xc5, 0xf9, 0xbd, 0x97,
	0x07, 0x13, 0xa5, 0x96, 0x99, 0x90, 0x5c, 0x73, 0x32, 0x32, 0xc7, 0xe4, 0x6e, 0xcd, 0x79, 0xdd,
	0x62, 0x4e, 0x45, 0x93, 0x53, 0xc6, 0xb8, 0xa6, 0xba, 0xe1, 0x4c, 0x79, 0x25, 0xb9, 0xb4, 0x78,
	0xa1, 0x37, 0x02, 0x55, 0xee, 0x7e, 0x3d, 0x4f, 0x7f, 0x07, 0x70, 0xbb, 0xc4, 0xaa, 0x93, 0x58,
	0x2e, 0xb1, 0x6d, 0xdf, 0xa2, 0x2e, 0xf0, 0x5b, 0x87, 0x4a, 0x93, 0xe7, 0x70, 0xf3, 0x0b, 0xb2,
	0x55, 0xc3, 0xd4, 0x8c, 0xd1, 0xaf, 0xa8, 0x04, 0xad, 0x30, 0x0e, 0x1e, 0x04, 0xcf, 0x26, 0xc5,
	0x45, 0x5f, 0xf8, 0xb0, 0xe5, 0x56, 0xae, 0xda, 0x4e, 0x69, 0x94, 0x03, 0xf9, 0xc4, 0xcb, 0x7d,
	0x61, 0x2f, 0x3f, 0x84, 0x68, 0x28, 0xc7, 0x23, 0xe7, 0x85, 0x03, 0x8f, 0x3c, 0x82, 0xeb, 0xe6,
	0x7e, 0x4d, 0x59, 0x85, 0xde, 0x39, 0x75, 0x4e, 0xb4, 0x85, 0x56, 0x4a, 0xff, 0x04, 0x70, 0x79,
	0x35, 0xbb, 0x12, 0xa6, 0x67, 0x24, 0x4f, 0x60, 0x6c, 0x44, 0xdd, 0x29, 0x97, 0x38, 0x9c, 0x9e,
	0x67, 0xbe, 0xf7, 0xac, 0x74, 0xb4, 0xe8, 0xab, 0xe4, 0x31, 0x5c, 0x33, 0xb3, 0x9b, 0xad, 0x70,
	0xe3, 0xd2, 0x86, 0xd3, 0x30, 0xb3, 0x63, 0x2d, 0xcb, 0x77, 0xef, 0x71, 0x63, 0x2c, 0xb5, 0x34,
	0xff, 0x07, 0x69, 0xe8, 0x62, 0x21, 0xfb, 0xc4, 0xbb, 0x34, 0xaf, 0x0d, 0x3b, 0x90, 0x04, 0x97,
	0xda, 0x45, 0x3e, 0xdb, 0x4b, 0x1f, 0x0d, 0x23, 0x04, 0x4e, 0x3b, 0x85, 0x32, 0x3e, 0x73, 0x17,
	0xb8, 0x73, 0xfa, 0x2b, 0x80, 0xb1, 0x7f, 0x21, 0xb9, 0x07, 0x20, 0xba, 0x79, 0xdb, 0x54, 0x2e,
	0x91, 0x8d, 0x1e, 0x15, 0x13, 0x4f, 0x6c, 0xf9, 0x3e, 0x84, 0x42, 0x36, 0x6b, 0xaa, 0x71, 0x97,
	0x38, 0x2a, 0xa0, 0x47, 0x56, 0x78, 0x0a, 0x37, 0xe8, 0x77, 0x35, 0xfb, 0xdc, 0xb0, 0x1a, 0xa5,
	0xe1, 0x4c, 0xf7, 0x51, 0xcf, 0x0d, 0x7e, 0xb3, 0xa7, 0x24, 0x87, 0x5b, 0x5c, 0x20, 0xb3, 0xbd,
	0x0f, 0x65, 0x3f, 0x65, 0xd2, 0x97, 0x06, 0x0f, 0x4c, 0x6b, 0x08, 0x07, 0xa3, 0x26, 0x9f, 0x60,
	0x64, 0xc6, 0x4d, 0x12, 0x3f, 0xad, 0x63, 0xfb, 0x93, 0xdc, 0x39, 0x5a, 0xf3, 0xdf, 0x27, 0x8d,
	0x7f, 0xfc, 0xfd, 0xf7, 0xf3, 0x84, 0x90, 0x0b, 0xb7, 0xae, 0x46, 0x34, 0xfd, 0xe4, 0xeb, 0x17,
	0xd9, 0xcb, 0xf9, 0xd8, 0xed, 0xe5, 0xab, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3f, 0x82, 0x71,
	0x2f, 0xdf, 0x02, 0x00, 0x00,
}
