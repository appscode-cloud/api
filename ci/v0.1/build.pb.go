// Code generated by protoc-gen-go.
// source: build.proto
// DO NOT EDIT!

/*
Package ci is a generated protocol buffer package.

It is generated from these files:
	build.proto
	job.proto
	master.proto
	slave.proto

It has these top-level messages:
	BuildDescribeRequest
	BuildDescribeResponse
	BuildListRequest
	BuildListResponse
	Build
	JobListResponse
	JobBuildRequest
	JobDescribeRequest
	JobDescribeResponse
	JobDeleteRequest
	JobCreateRequest
	JobCopyRequest
	Job
	MasterCreateRequest
	MasterCreateResponse
	SlaveListResponse
	Slave
	SlaveCreateRequest
	SlaveDescribeRequest
	SlaveDescribeResponse
	SlaveDeleteRequest
	SlaveRestartRequest
	SlaveRestartResponse
*/
package ci

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

type BuildDescribeRequest struct {
	JobName string `protobuf:"bytes,1,opt,name=job_name,json=jobName" json:"job_name,omitempty"`
	Number  int64  `protobuf:"varint,2,opt,name=number" json:"number,omitempty"`
	Console string `protobuf:"bytes,3,opt,name=console" json:"console,omitempty"`
}

func (m *BuildDescribeRequest) Reset()                    { *m = BuildDescribeRequest{} }
func (m *BuildDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*BuildDescribeRequest) ProtoMessage()               {}
func (*BuildDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type BuildDescribeResponse struct {
	Status        *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	BuildNumber   int64          `protobuf:"varint,2,opt,name=build_number,json=buildNumber" json:"build_number,omitempty"`
	JobName       string         `protobuf:"bytes,3,opt,name=job_name,json=jobName" json:"job_name,omitempty"`
	BaseUrl       string         `protobuf:"bytes,4,opt,name=base_url,json=baseUrl" json:"base_url,omitempty"`
	Result        string         `protobuf:"bytes,5,opt,name=result" json:"result,omitempty"`
	ConsoleOutput string         `protobuf:"bytes,6,opt,name=console_output,json=consoleOutput" json:"console_output,omitempty"`
}

func (m *BuildDescribeResponse) Reset()                    { *m = BuildDescribeResponse{} }
func (m *BuildDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*BuildDescribeResponse) ProtoMessage()               {}
func (*BuildDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BuildDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type BuildListRequest struct {
	JobName string `protobuf:"bytes,1,opt,name=job_name,json=jobName" json:"job_name,omitempty"`
}

func (m *BuildListRequest) Reset()                    { *m = BuildListRequest{} }
func (m *BuildListRequest) String() string            { return proto.CompactTextString(m) }
func (*BuildListRequest) ProtoMessage()               {}
func (*BuildListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type BuildListResponse struct {
	Status  *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	JobName string         `protobuf:"bytes,2,opt,name=job_name,json=jobName" json:"job_name,omitempty"`
	Builds  []*Build       `protobuf:"bytes,3,rep,name=builds" json:"builds,omitempty"`
}

func (m *BuildListResponse) Reset()                    { *m = BuildListResponse{} }
func (m *BuildListResponse) String() string            { return proto.CompactTextString(m) }
func (*BuildListResponse) ProtoMessage()               {}
func (*BuildListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BuildListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *BuildListResponse) GetBuilds() []*Build {
	if m != nil {
		return m.Builds
	}
	return nil
}

type Build struct {
	BuildNumber int64  `protobuf:"varint,1,opt,name=build_number,json=buildNumber" json:"build_number,omitempty"`
	Result      string `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
}

func (m *Build) Reset()                    { *m = Build{} }
func (m *Build) String() string            { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()               {}
func (*Build) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*BuildDescribeRequest)(nil), "ci.BuildDescribeRequest")
	proto.RegisterType((*BuildDescribeResponse)(nil), "ci.BuildDescribeResponse")
	proto.RegisterType((*BuildListRequest)(nil), "ci.BuildListRequest")
	proto.RegisterType((*BuildListResponse)(nil), "ci.BuildListResponse")
	proto.RegisterType((*Build)(nil), "ci.Build")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for Builds service

type BuildsClient interface {
	Describe(ctx context.Context, in *BuildDescribeRequest, opts ...grpc.CallOption) (*BuildDescribeResponse, error)
	List(ctx context.Context, in *BuildListRequest, opts ...grpc.CallOption) (*BuildListResponse, error)
}

type buildsClient struct {
	cc *grpc.ClientConn
}

func NewBuildsClient(cc *grpc.ClientConn) BuildsClient {
	return &buildsClient{cc}
}

func (c *buildsClient) Describe(ctx context.Context, in *BuildDescribeRequest, opts ...grpc.CallOption) (*BuildDescribeResponse, error) {
	out := new(BuildDescribeResponse)
	err := grpc.Invoke(ctx, "/ci.Builds/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *buildsClient) List(ctx context.Context, in *BuildListRequest, opts ...grpc.CallOption) (*BuildListResponse, error) {
	out := new(BuildListResponse)
	err := grpc.Invoke(ctx, "/ci.Builds/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Builds service

type BuildsServer interface {
	Describe(context.Context, *BuildDescribeRequest) (*BuildDescribeResponse, error)
	List(context.Context, *BuildListRequest) (*BuildListResponse, error)
}

func RegisterBuildsServer(s *grpc.Server, srv BuildsServer) {
	s.RegisterService(&_Builds_serviceDesc, srv)
}

func _Builds_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(BuildDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BuildsServer).Describe(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Builds_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(BuildListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BuildsServer).List(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Builds_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ci.Builds",
	HandlerType: (*BuildsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Describe",
			Handler:    _Builds_Describe_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Builds_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xcf, 0x8a, 0xd3, 0x40,
	0x18, 0x27, 0xc9, 0x6e, 0x76, 0xf7, 0x5b, 0x5d, 0x74, 0xd8, 0x96, 0x69, 0x28, 0x62, 0x83, 0x55,
	0x2f, 0x4d, 0xa4, 0xbe, 0x41, 0xf1, 0x28, 0x15, 0x22, 0x1e, 0x3c, 0x85, 0x24, 0x1d, 0xca, 0x48,
	0x9a, 0x89, 0x99, 0x19, 0xa1, 0x94, 0x5e, 0x7c, 0x05, 0x1f, 0xad, 0x07, 0x5f, 0xc0, 0x07, 0x71,
	0x32, 0x33, 0xad, 0x69, 0x2b, 0xca, 0x5e, 0x02, 0xdf, 0xef, 0x37, 0xf9, 0x7e, 0x7f, 0x32, 0x81,
	0xdb, 0x5c, 0xd2, 0x72, 0x11, 0xd5, 0x0d, 0x13, 0x0c, 0xb9, 0x05, 0x0d, 0x86, 0x4b, 0xc6, 0x96,
	0x25, 0x89, 0xb3, 0x9a, 0xc6, 0x59, 0x55, 0x31, 0x91, 0x09, 0xca, 0x2a, 0x6e, 0x4e, 0x04, 0xfd,
	0x16, 0x5e, 0x88, 0x75, 0x4d, 0x78, 0xac, 0x9f, 0x06, 0x0f, 0x0b, 0xb8, 0x9f, 0xb5, 0x8b, 0xde,
	0x11, 0x5e, 0x34, 0x34, 0x27, 0x09, 0xf9, 0x2a, 0x09, 0x17, 0x68, 0x00, 0xd7, 0x5f, 0x58, 0x9e,
	0x56, 0xd9, 0x8a, 0x60, 0xe7, 0xb9, 0xf3, 0xfa, 0x26, 0xb9, 0x52, 0xf3, 0x5c, 0x8d, 0xa8, 0x0f,
	0x7e, 0x25, 0x57, 0x39, 0x69, 0xb0, 0xab, 0x08, 0x2f, 0xb1, 0x13, 0xc2, 0x70, 0x55, 0x28, 0x41,
	0x56, 0x12, 0xec, 0x99, 0x37, 0xec, 0x18, 0xfe, 0x74, 0xa0, 0x77, 0xa2, 0xc2, 0x6b, 0xc5, 0x11,
	0xf4, 0x12, 0x7c, 0xae, 0x8c, 0x4a, 0xae, 0x45, 0x6e, 0xa7, 0x77, 0x91, 0xf1, 0x18, 0x7d, 0xd4,
	0x68, 0x62, 0x59, 0x34, 0x82, 0x47, 0x3a, 0x6f, 0x7a, 0xa4, 0x6c, 0x3a, 0x98, 0x1b, 0xf9, 0xae,
	0x63, 0xef, 0xd8, 0xb1, 0xa2, 0xf2, 0x8c, 0x93, 0x54, 0x36, 0x25, 0xbe, 0x30, 0x54, 0x3b, 0x7f,
	0x6a, 0xca, 0x36, 0x4c, 0x43, 0xb8, 0x2c, 0x05, 0xbe, 0xd4, 0x84, 0x9d, 0xd0, 0x18, 0xee, 0xac,
	0xfb, 0x94, 0x49, 0x51, 0x4b, 0x81, 0x7d, 0xcd, 0x3f, 0xb6, 0xe8, 0x07, 0x0d, 0x86, 0x13, 0x78,
	0xa2, 0x83, 0xbd, 0xa7, 0x5c, 0xfc, 0xbf, 0xba, 0x70, 0x0d, 0x4f, 0x3b, 0xc7, 0x1f, 0xd8, 0x41,
	0x77, 0xaf, 0x7b, 0x1c, 0x70, 0x04, 0xbe, 0xae, 0x82, 0xab, 0xe4, 0x9e, 0x5a, 0x71, 0x13, 0x15,
	0x34, 0xd2, 0x4a, 0x89, 0x25, 0xc2, 0x19, 0x5c, 0x6a, 0xe0, 0xac, 0x4a, 0xe7, 0xbc, 0xca, 0x3f,
	0xa5, 0xb8, 0xdd, 0x52, 0xa6, 0x3b, 0x07, 0x7c, 0xbd, 0x84, 0xa3, 0x0a, 0xae, 0xf7, 0x1f, 0x13,
	0xe1, 0x83, 0xda, 0xc9, 0x2d, 0x0a, 0x06, 0x7f, 0x61, 0x4c, 0xea, 0x70, 0xf2, 0x7d, 0xf7, 0xeb,
	0x87, 0xfb, 0x0a, 0x8d, 0xf5, 0x85, 0x2d, 0x68, 0xfc, 0xed, 0x4d, 0x6c, 0xac, 0xc6, 0x9b, 0x7d,
	0xcc, 0x6d, 0xbc, 0x31, 0x26, 0xb7, 0xe8, 0x33, 0x5c, 0xb4, 0xa5, 0xa1, 0xfb, 0xc3, 0xc6, 0x4e,
	0xe5, 0x41, 0xef, 0x04, 0xb5, 0x1a, 0x2f, 0xb4, 0xc6, 0x33, 0x34, 0xfc, 0x97, 0x46, 0xee, 0xeb,
	0x3f, 0xe1, 0xed, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x2d, 0xaa, 0x29, 0x52, 0x03, 0x00,
	0x00,
}
