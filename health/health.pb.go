// Code generated by protoc-gen-go.
// source: health.proto
// DO NOT EDIT!

/*
Package health is a generated protocol buffer package.

It is generated from these files:
	health.proto

It has these top-level messages:
	Response
	Version
*/
package health

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

type Response struct {
	Status  *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Version *Version       `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Response) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Response) GetVersion() *Version {
	if m != nil {
		return m.Version
	}
	return nil
}

type Version struct {
	Version    string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Status     string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	Endpoint   string `protobuf:"bytes,3,opt,name=endpoint" json:"endpoint,omitempty"`
	CommitHash string `protobuf:"bytes,4,opt,name=commit_hash,json=commitHash" json:"commit_hash,omitempty"`
	// Only development mode informations.
	Name            string `protobuf:"bytes,5,opt,name=Name,json=name" json:"Name,omitempty"`
	GitBranch       string `protobuf:"bytes,6,opt,name=git_branch,json=gitBranch" json:"git_branch,omitempty"`
	GitTag          string `protobuf:"bytes,7,opt,name=git_tag,json=gitTag" json:"git_tag,omitempty"`
	CommitTimestamp string `protobuf:"bytes,8,opt,name=commit_timestamp,json=commitTimestamp" json:"commit_timestamp,omitempty"`
	BuildTimestamp  string `protobuf:"bytes,9,opt,name=build_timestamp,json=buildTimestamp" json:"build_timestamp,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Response)(nil), "health.Response")
	proto.RegisterType((*Version)(nil), "health.Version")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Health service

type HealthClient interface {
	Status(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*Response, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Status(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/health.Health/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Health service

type HealthServer interface {
	Status(context.Context, *dtypes.VoidRequest) (*Response, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health.Health/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Status(ctx, req.(*dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "health.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Health_Status_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x4c, 0x91, 0xcf, 0x4e, 0xe3, 0x30,
	0x10, 0xc6, 0xd5, 0x6c, 0x37, 0x49, 0xa7, 0xab, 0xb6, 0x9a, 0x95, 0xba, 0x56, 0xb4, 0x08, 0xd4,
	0x03, 0x7f, 0x2e, 0x89, 0x54, 0x9e, 0x00, 0x4e, 0x95, 0x90, 0x38, 0x84, 0xaa, 0x37, 0x84, 0xdc,
	0xc6, 0x4a, 0x22, 0x35, 0x76, 0xa8, 0x5d, 0x24, 0x38, 0xf2, 0x0a, 0x5c, 0x78, 0x2f, 0x5e, 0x81,
	0x07, 0xc1, 0x1d, 0x27, 0x81, 0x4b, 0x94, 0xf9, 0x7d, 0xdf, 0x78, 0xc6, 0x9f, 0xe1, 0x4f, 0x21,
	0xf8, 0xd6, 0x14, 0x71, 0xbd, 0x53, 0x46, 0xa1, 0xef, 0xaa, 0xe8, 0x7f, 0xae, 0x54, 0xbe, 0x15,
	0x09, 0xaf, 0xcb, 0x84, 0x4b, 0xa9, 0x0c, 0x37, 0xa5, 0x92, 0xda, 0xb9, 0xa2, 0xe9, 0x01, 0x67,
	0xe6, 0xb9, 0x16, 0x3a, 0xa1, 0xaf, 0xe3, 0xb3, 0x7b, 0x08, 0x53, 0xa1, 0x6b, 0x6b, 0x14, 0x78,
	0x0a, 0xbe, 0xb6, 0x5d, 0x7b, 0xcd, 0x7a, 0x27, 0xbd, 0xf3, 0xe1, 0x7c, 0x14, 0xbb, 0x86, 0xf8,
	0x8e, 0x68, 0xda, 0xa8, 0x78, 0x01, 0xc1, 0x93, 0xd8, 0x69, 0x7b, 0x3a, 0xf3, 0xc8, 0x38, 0x8e,
	0x9b, 0x8d, 0x56, 0x0e, 0xa7, 0xad, 0x3e, 0x7b, 0xf7, 0x20, 0x68, 0x20, 0xb2, 0xef, 0xb6, 0xc3,
	0xf9, 0x83, 0xce, 0x85, 0xd3, 0x6e, 0xb0, 0x47, 0x42, 0x3b, 0x28, 0x82, 0x50, 0xc8, 0xac, 0x56,
	0xa5, 0x34, 0xec, 0x17, 0x29, 0x5d, 0x8d, 0xc7, 0x30, 0xdc, 0xa8, 0xaa, 0x2a, 0xcd, 0x43, 0xc1,
	0x75, 0xc1, 0xfa, 0x24, 0x83, 0x43, 0x0b, 0x4b, 0x10, 0xa1, 0x7f, 0xcb, 0x2b, 0xc1, 0x7e, 0x93,
	0xd2, 0x97, 0xf6, 0x1f, 0x8f, 0x00, 0x72, 0xdb, 0xb1, 0xde, 0x71, 0xb9, 0x29, 0x98, 0x4f, 0xca,
	0xc0, 0x92, 0x6b, 0x02, 0xf8, 0x0f, 0x82, 0x83, 0x6c, 0x78, 0xce, 0x02, 0xb7, 0x88, 0x2d, 0x97,
	0x3c, 0xb7, 0x37, 0x9e, 0x34, 0xc3, 0x4c, 0x59, 0x09, 0xbb, 0x5d, 0x55, 0xb3, 0x90, 0x1c, 0x63,
	0xc7, 0x97, 0x2d, 0xc6, 0x33, 0x18, 0xaf, 0xf7, 0xe5, 0x36, 0xfb, 0xe1, 0x1c, 0x90, 0x73, 0x44,
	0xb8, 0x33, 0xce, 0x6f, 0xc0, 0x5f, 0x50, 0x6a, 0x78, 0x05, 0xbe, 0x4b, 0x18, 0xff, 0xb6, 0x89,
	0xaf, 0x54, 0x99, 0xa5, 0xe2, 0x71, 0x6f, 0xbd, 0xd1, 0xa4, 0x4d, 0xb7, 0x7d, 0xa8, 0xd9, 0xe4,
	0xf5, 0xe3, 0xf3, 0xcd, 0x03, 0x0c, 0x13, 0xa7, 0xbc, 0xac, 0x7d, 0x7a, 0xcd, 0xcb, 0xaf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xb8, 0x6f, 0x64, 0x12, 0x1b, 0x02, 0x00, 0x00,
}
