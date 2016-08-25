// Code generated by protoc-gen-go.
// source: ssh.proto
// DO NOT EDIT!

/*
Package ssh is a generated protocol buffer package.

It is generated from these files:
	ssh.proto

It has these top-level messages:
	SSHGetRequest
	SSHGetResponse
	SSHKey
*/
package ssh

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Use specific requests for protos
type SSHGetRequest struct {
	Namespace    string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	ClusterName  string `protobuf:"bytes,2,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	InstanceName string `protobuf:"bytes,3,opt,name=instance_name,json=instanceName" json:"instance_name,omitempty"`
}

func (m *SSHGetRequest) Reset()                    { *m = SSHGetRequest{} }
func (m *SSHGetRequest) String() string            { return proto.CompactTextString(m) }
func (*SSHGetRequest) ProtoMessage()               {}
func (*SSHGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// return phid ?
type SSHGetResponse struct {
	Status       *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	SshKey       *SSHKey        `protobuf:"bytes,2,opt,name=ssh_key,json=sshKey" json:"ssh_key,omitempty"`
	InstanceAddr string         `protobuf:"bytes,3,opt,name=instance_addr,json=instanceAddr" json:"instance_addr,omitempty"`
	InstancePort int32          `protobuf:"varint,4,opt,name=instance_port,json=instancePort" json:"instance_port,omitempty"`
	User         string         `protobuf:"bytes,5,opt,name=user" json:"user,omitempty"`
	Command      string         `protobuf:"bytes,6,opt,name=command" json:"command,omitempty"`
}

func (m *SSHGetResponse) Reset()                    { *m = SSHGetResponse{} }
func (m *SSHGetResponse) String() string            { return proto.CompactTextString(m) }
func (*SSHGetResponse) ProtoMessage()               {}
func (*SSHGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SSHGetResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *SSHGetResponse) GetSshKey() *SSHKey {
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
	proto.RegisterType((*SSHGetRequest)(nil), "ssh.SSHGetRequest")
	proto.RegisterType((*SSHGetResponse)(nil), "ssh.SSHGetResponse")
	proto.RegisterType((*SSHKey)(nil), "ssh.SSHKey")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for SSH service

type SSHClient interface {
	Get(ctx context.Context, in *SSHGetRequest, opts ...grpc.CallOption) (*SSHGetResponse, error)
}

type sSHClient struct {
	cc *grpc.ClientConn
}

func NewSSHClient(cc *grpc.ClientConn) SSHClient {
	return &sSHClient{cc}
}

func (c *sSHClient) Get(ctx context.Context, in *SSHGetRequest, opts ...grpc.CallOption) (*SSHGetResponse, error) {
	out := new(SSHGetResponse)
	err := grpc.Invoke(ctx, "/ssh.SSH/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SSH service

type SSHServer interface {
	Get(context.Context, *SSHGetRequest) (*SSHGetResponse, error)
}

func RegisterSSHServer(s *grpc.Server, srv SSHServer) {
	s.RegisterService(&_SSH_serviceDesc, srv)
}

func _SSH_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SSHGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SSHServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ssh.SSH/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SSHServer).Get(ctx, req.(*SSHGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SSH_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ssh.SSH",
	HandlerType: (*SSHServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SSH_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("ssh.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x92, 0x5f, 0x8e, 0xd3, 0x30,
	0x10, 0x87, 0xd5, 0x6d, 0xb7, 0xab, 0x4e, 0xb2, 0x45, 0x9a, 0x95, 0x56, 0x51, 0xb5, 0xfc, 0x0b,
	0x08, 0x78, 0x4a, 0xa0, 0x9c, 0x80, 0x17, 0x40, 0x42, 0x5a, 0xa1, 0x54, 0x3c, 0x57, 0xde, 0x64,
	0xe8, 0x46, 0xb4, 0xb6, 0x89, 0x9d, 0x45, 0x7d, 0xe5, 0x0a, 0xdc, 0x80, 0x2b, 0x21, 0x71, 0x02,
	0x0e, 0xc2, 0xd8, 0x4e, 0x48, 0xcb, 0x4b, 0xe2, 0x7c, 0xfe, 0x62, 0xff, 0xc6, 0x1e, 0x98, 0x19,
	0x73, 0x9b, 0xe9, 0x46, 0x59, 0x85, 0x63, 0x1e, 0x2e, 0xae, 0x36, 0x4a, 0x6d, 0xb6, 0x94, 0x0b,
	0x5d, 0xe7, 0x42, 0x4a, 0x65, 0x85, 0xad, 0x95, 0x34, 0x41, 0x59, 0x5c, 0x3a, 0x5c, 0xd9, 0xbd,
	0x26, 0x93, 0xfb, 0x67, 0xe0, 0x69, 0x0b, 0xe7, 0xab, 0xd5, 0xfb, 0x77, 0x64, 0x0b, 0xfa, 0xda,
	0x92, 0xb1, 0x78, 0x05, 0x33, 0x29, 0x76, 0x64, 0xb4, 0x28, 0x29, 0x19, 0x3d, 0x1a, 0xbd, 0x98,
	0x15, 0x03, 0xc0, 0xc7, 0x10, 0x97, 0xdb, 0xd6, 0x58, 0x6a, 0xd6, 0x0e, 0x26, 0x27, 0x5e, 0x88,
	0x3a, 0x76, 0xcd, 0x08, 0x9f, 0xc0, 0x79, 0x2d, 0x8d, 0x15, 0xb2, 0xa4, 0xe0, 0x8c, 0xbd, 0x13,
	0xf7, 0xd0, 0x49, 0xe9, 0xef, 0x11, 0xcc, 0xfb, 0x7d, 0x8d, 0xe6, 0x98, 0x84, 0xcf, 0x60, 0xca,
	0x82, 0x6d, 0x8d, 0xdf, 0x35, 0x5a, 0xce, 0xb3, 0x10, 0x37, 0x5b, 0x79, 0x5a, 0x74, 0xb3, 0xf8,
	0x14, 0xce, 0xb8, 0xdc, 0xf5, 0x17, 0xda, 0xfb, 0xdd, 0xa3, 0x65, 0x94, 0xb9, 0x93, 0xe0, 0xd5,
	0x3e, 0xd0, 0x9e, 0x2d, 0x73, 0xcb, 0xef, 0xa3, 0x14, 0xa2, 0xaa, 0x9a, 0xff, 0x53, 0xbc, 0x61,
	0x76, 0x24, 0x69, 0xd5, 0xd8, 0x64, 0xc2, 0xd2, 0xe9, 0x20, 0x7d, 0x64, 0x86, 0x08, 0x93, 0xd6,
	0x50, 0x93, 0x9c, 0xfa, 0x05, 0xfc, 0x18, 0x13, 0x38, 0x2b, 0xd5, 0x6e, 0x27, 0x64, 0x95, 0x4c,
	0x3d, 0xee, 0x3f, 0xd3, 0x9f, 0x23, 0x98, 0x86, 0x28, 0x78, 0x1f, 0x40, 0xb7, 0x37, 0xdb, 0xba,
	0xf4, 0x59, 0x5d, 0x51, 0x71, 0x31, 0x0b, 0xc4, 0x4d, 0x3f, 0x84, 0x48, 0x37, 0xf5, 0x9d, 0xb0,
	0xf4, 0xaf, 0x96, 0xb8, 0x80, 0x0e, 0x39, 0xe1, 0x39, 0xdc, 0x13, 0xdf, 0xcc, 0xfa, 0x73, 0x2d,
	0x37, 0xd4, 0x30, 0x97, 0xb6, 0x2b, 0x62, 0xce, 0xf8, 0xed, 0x40, 0x31, 0x87, 0x0b, 0xa5, 0x49,
	0xba, 0x53, 0x39, 0x94, 0x27, 0x5e, 0xc6, 0x6e, 0xea, 0xe0, 0x87, 0xe5, 0x27, 0x18, 0x73, 0x46,
	0xbc, 0x86, 0x31, 0x5f, 0x00, 0x62, 0x7f, 0x7e, 0x43, 0x17, 0x2c, 0x2e, 0x8e, 0x58, 0xb8, 0xa1,
	0xf4, 0xc1, 0xf7, 0x5f, 0x7f, 0x7e, 0x9c, 0x24, 0x78, 0xc9, 0x3d, 0xa6, 0x4d, 0xa9, 0xaa, 0xd0,
	0x6c, 0x6c, 0xe6, 0x77, 0x2f, 0xb3, 0x57, 0x37, 0x53, 0xdf, 0x52, 0xaf, 0xff, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x23, 0x55, 0xcd, 0x22, 0x9a, 0x02, 0x00, 0x00,
}
