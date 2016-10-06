// Code generated by protoc-gen-go.
// source: job.proto
// DO NOT EDIT!

package v1beta1

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

type JobListRequest struct {
	Parents string `protobuf:"bytes,1,opt,name=parents" json:"parents,omitempty"`
}

func (m *JobListRequest) Reset()                    { *m = JobListRequest{} }
func (m *JobListRequest) String() string            { return proto.CompactTextString(m) }
func (*JobListRequest) ProtoMessage()               {}
func (*JobListRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type JobListResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Jobs   []*Job         `protobuf:"bytes,2,rep,name=jobs" json:"jobs,omitempty"`
}

func (m *JobListResponse) Reset()                    { *m = JobListResponse{} }
func (m *JobListResponse) String() string            { return proto.CompactTextString(m) }
func (*JobListResponse) ProtoMessage()               {}
func (*JobListResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *JobListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *JobListResponse) GetJobs() []*Job {
	if m != nil {
		return m.Jobs
	}
	return nil
}

type JobBuildRequest struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Parents string `protobuf:"bytes,2,opt,name=parents" json:"parents,omitempty"`
	Param   string `protobuf:"bytes,3,opt,name=param" json:"param,omitempty"`
}

func (m *JobBuildRequest) Reset()                    { *m = JobBuildRequest{} }
func (m *JobBuildRequest) String() string            { return proto.CompactTextString(m) }
func (*JobBuildRequest) ProtoMessage()               {}
func (*JobBuildRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type JobDescribeRequest struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Parents string `protobuf:"bytes,2,opt,name=parents" json:"parents,omitempty"`
}

func (m *JobDescribeRequest) Reset()                    { *m = JobDescribeRequest{} }
func (m *JobDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*JobDescribeRequest) ProtoMessage()               {}
func (*JobDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

type JobDescribeResponse struct {
	Status                *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Name                  string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Parents               string         `protobuf:"bytes,3,opt,name=parents" json:"parents,omitempty"`
	Description           string         `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	DisplayName           string         `protobuf:"bytes,5,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
	JobType               string         `protobuf:"bytes,6,opt,name=job_type,json=jobType" json:"job_type,omitempty"`
	JobColor              string         `protobuf:"bytes,7,opt,name=job_color,json=jobColor" json:"job_color,omitempty"`
	Buildable             bool           `protobuf:"varint,8,opt,name=buildable" json:"buildable,omitempty"`
	BuildCount            int64          `protobuf:"varint,9,opt,name=build_count,json=buildCount" json:"build_count,omitempty"`
	BuildIds              []int64        `protobuf:"varint,10,rep,packed,name=build_ids,json=buildIds" json:"build_ids,omitempty"`
	FirstBuildId          int64          `protobuf:"varint,11,opt,name=first_build_id,json=firstBuildId" json:"first_build_id,omitempty"`
	LastBuildId           int64          `protobuf:"varint,12,opt,name=last_build_id,json=lastBuildId" json:"last_build_id,omitempty"`
	LastCompletedBuildId  int64          `protobuf:"varint,13,opt,name=last_completed_build_id,json=lastCompletedBuildId" json:"last_completed_build_id,omitempty"`
	LastFailedBuildId     int64          `protobuf:"varint,14,opt,name=last_failed_build_id,json=lastFailedBuildId" json:"last_failed_build_id,omitempty"`
	LastSuccessfulBuildId int64          `protobuf:"varint,15,opt,name=last_successful_build_id,json=lastSuccessfulBuildId" json:"last_successful_build_id,omitempty"`
	NextBuildNumber       int64          `protobuf:"varint,16,opt,name=next_build_number,json=nextBuildNumber" json:"next_build_number,omitempty"`
	HealthReport          []*JobHealth   `protobuf:"bytes,17,rep,name=health_report,json=healthReport" json:"health_report,omitempty"`
}

func (m *JobDescribeResponse) Reset()                    { *m = JobDescribeResponse{} }
func (m *JobDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*JobDescribeResponse) ProtoMessage()               {}
func (*JobDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *JobDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *JobDescribeResponse) GetHealthReport() []*JobHealth {
	if m != nil {
		return m.HealthReport
	}
	return nil
}

type JobHealth struct {
	Description string `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	Score       int64  `protobuf:"varint,4,opt,name=score" json:"score,omitempty"`
}

func (m *JobHealth) Reset()                    { *m = JobHealth{} }
func (m *JobHealth) String() string            { return proto.CompactTextString(m) }
func (*JobHealth) ProtoMessage()               {}
func (*JobHealth) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

type JobDeleteRequest struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Parents string `protobuf:"bytes,2,opt,name=parents" json:"parents,omitempty"`
}

func (m *JobDeleteRequest) Reset()                    { *m = JobDeleteRequest{} }
func (m *JobDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*JobDeleteRequest) ProtoMessage()               {}
func (*JobDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

type JobCreateRequest struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Parents string `protobuf:"bytes,2,opt,name=parents" json:"parents,omitempty"`
	ShFile  string `protobuf:"bytes,3,opt,name=sh_file,json=shFile" json:"sh_file,omitempty"`
}

func (m *JobCreateRequest) Reset()                    { *m = JobCreateRequest{} }
func (m *JobCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*JobCreateRequest) ProtoMessage()               {}
func (*JobCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

type JobCopyRequest struct {
	Source      string `protobuf:"bytes,1,opt,name=source" json:"source,omitempty"`
	Destination string `protobuf:"bytes,2,opt,name=destination" json:"destination,omitempty"`
}

func (m *JobCopyRequest) Reset()                    { *m = JobCopyRequest{} }
func (m *JobCopyRequest) String() string            { return proto.CompactTextString(m) }
func (*JobCopyRequest) ProtoMessage()               {}
func (*JobCopyRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

type Job struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Parents string `protobuf:"bytes,2,opt,name=parents" json:"parents,omitempty"`
	JobType string `protobuf:"bytes,3,opt,name=job_type,json=jobType" json:"job_type,omitempty"`
	Color   string `protobuf:"bytes,4,opt,name=color" json:"color,omitempty"`
	Url     string `protobuf:"bytes,5,opt,name=url" json:"url,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func init() {
	proto.RegisterType((*JobListRequest)(nil), "ci.v1beta1.JobListRequest")
	proto.RegisterType((*JobListResponse)(nil), "ci.v1beta1.JobListResponse")
	proto.RegisterType((*JobBuildRequest)(nil), "ci.v1beta1.JobBuildRequest")
	proto.RegisterType((*JobDescribeRequest)(nil), "ci.v1beta1.JobDescribeRequest")
	proto.RegisterType((*JobDescribeResponse)(nil), "ci.v1beta1.JobDescribeResponse")
	proto.RegisterType((*JobHealth)(nil), "ci.v1beta1.JobHealth")
	proto.RegisterType((*JobDeleteRequest)(nil), "ci.v1beta1.JobDeleteRequest")
	proto.RegisterType((*JobCreateRequest)(nil), "ci.v1beta1.JobCreateRequest")
	proto.RegisterType((*JobCopyRequest)(nil), "ci.v1beta1.JobCopyRequest")
	proto.RegisterType((*Job)(nil), "ci.v1beta1.Job")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Jobs service

type JobsClient interface {
	List(ctx context.Context, in *JobListRequest, opts ...grpc.CallOption) (*JobListResponse, error)
	Describe(ctx context.Context, in *JobDescribeRequest, opts ...grpc.CallOption) (*JobDescribeResponse, error)
	Create(ctx context.Context, in *JobCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Copy(ctx context.Context, in *JobCopyRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Build(ctx context.Context, in *JobBuildRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *JobDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type jobsClient struct {
	cc *grpc.ClientConn
}

func NewJobsClient(cc *grpc.ClientConn) JobsClient {
	return &jobsClient{cc}
}

func (c *jobsClient) List(ctx context.Context, in *JobListRequest, opts ...grpc.CallOption) (*JobListResponse, error) {
	out := new(JobListResponse)
	err := grpc.Invoke(ctx, "/ci.v1beta1.Jobs/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) Describe(ctx context.Context, in *JobDescribeRequest, opts ...grpc.CallOption) (*JobDescribeResponse, error) {
	out := new(JobDescribeResponse)
	err := grpc.Invoke(ctx, "/ci.v1beta1.Jobs/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) Create(ctx context.Context, in *JobCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/ci.v1beta1.Jobs/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) Copy(ctx context.Context, in *JobCopyRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/ci.v1beta1.Jobs/Copy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) Build(ctx context.Context, in *JobBuildRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/ci.v1beta1.Jobs/Build", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) Delete(ctx context.Context, in *JobDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/ci.v1beta1.Jobs/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Jobs service

type JobsServer interface {
	List(context.Context, *JobListRequest) (*JobListResponse, error)
	Describe(context.Context, *JobDescribeRequest) (*JobDescribeResponse, error)
	Create(context.Context, *JobCreateRequest) (*dtypes.VoidResponse, error)
	Copy(context.Context, *JobCopyRequest) (*dtypes.VoidResponse, error)
	Build(context.Context, *JobBuildRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *JobDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterJobsServer(s *grpc.Server, srv JobsServer) {
	s.RegisterService(&_Jobs_serviceDesc, srv)
}

func _Jobs_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.v1beta1.Jobs/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).List(ctx, req.(*JobListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.v1beta1.Jobs/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).Describe(ctx, req.(*JobDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.v1beta1.Jobs/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).Create(ctx, req.(*JobCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_Copy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobCopyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).Copy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.v1beta1.Jobs/Copy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).Copy(ctx, req.(*JobCopyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_Build_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobBuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).Build(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.v1beta1.Jobs/Build",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).Build(ctx, req.(*JobBuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.v1beta1.Jobs/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).Delete(ctx, req.(*JobDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Jobs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ci.v1beta1.Jobs",
	HandlerType: (*JobsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Jobs_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Jobs_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Jobs_Create_Handler,
		},
		{
			MethodName: "Copy",
			Handler:    _Jobs_Copy_Handler,
		},
		{
			MethodName: "Build",
			Handler:    _Jobs_Build_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Jobs_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor2,
}

func init() { proto.RegisterFile("job.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 857 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x56, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0x56, 0xd6, 0xf9, 0x3d, 0xc9, 0x26, 0xbb, 0x43, 0xda, 0x9a, 0x74, 0x61, 0x83, 0x69, 0xab,
	0x28, 0x42, 0x31, 0x0d, 0x42, 0x95, 0x2a, 0x21, 0xa1, 0x04, 0x55, 0x10, 0xa1, 0x5e, 0xb8, 0x80,
	0x04, 0x17, 0x58, 0x63, 0x7b, 0xd2, 0x38, 0x72, 0x3c, 0xc6, 0x33, 0xae, 0x88, 0x10, 0x37, 0xbc,
	0x02, 0x17, 0xbc, 0x00, 0x6f, 0xc4, 0x2b, 0xf0, 0x20, 0xd5, 0x9c, 0xb1, 0x13, 0x27, 0xdb, 0x54,
	0xed, 0xde, 0xac, 0x32, 0xdf, 0xf9, 0xce, 0xf7, 0xcd, 0xcf, 0x39, 0xc7, 0x0b, 0xad, 0x35, 0xf7,
	0x26, 0x49, 0xca, 0x25, 0x27, 0xe0, 0x87, 0x93, 0x57, 0x8f, 0x3d, 0x26, 0xe9, 0xe3, 0xc1, 0xd5,
	0x4b, 0xce, 0x5f, 0x46, 0xcc, 0xa6, 0x49, 0x68, 0xd3, 0x38, 0xe6, 0x92, 0xca, 0x90, 0xc7, 0x42,
	0x33, 0x07, 0x77, 0x15, 0x1c, 0xc8, 0x6d, 0xc2, 0x84, 0x8d, 0x7f, 0x35, 0x6e, 0x8d, 0xa1, 0xbb,
	0xe0, 0xde, 0xf7, 0xa1, 0x90, 0x0e, 0xfb, 0x2d, 0x63, 0x42, 0x12, 0x13, 0x1a, 0x09, 0x4d, 0x59,
	0x2c, 0x85, 0x59, 0x19, 0x56, 0x46, 0x2d, 0xa7, 0x58, 0x5a, 0xbf, 0x42, 0x6f, 0xc7, 0x15, 0x09,
	0x8f, 0x05, 0x23, 0x8f, 0xa0, 0x2e, 0x24, 0x95, 0x99, 0xe6, 0xb6, 0xa7, 0xdd, 0x89, 0xf6, 0x98,
	0xbc, 0x40, 0xd4, 0xc9, 0xa3, 0xe4, 0x53, 0xa8, 0xae, 0xb9, 0x27, 0xcc, 0xb3, 0xa1, 0x31, 0x6a,
	0x4f, 0x7b, 0x93, 0xfd, 0xbe, 0x27, 0x0b, 0xee, 0x39, 0x18, 0xb4, 0x7e, 0x44, 0xfd, 0x59, 0x16,
	0x46, 0x41, 0xb1, 0x19, 0x02, 0xd5, 0x98, 0x6e, 0x58, 0xbe, 0x13, 0xfc, 0x5d, 0xde, 0xe0, 0xd9,
	0xc1, 0x06, 0x49, 0x1f, 0x6a, 0x09, 0x4d, 0xe9, 0xc6, 0x34, 0x10, 0xd7, 0x0b, 0x6b, 0x06, 0x64,
	0xc1, 0xbd, 0x6f, 0x98, 0xf0, 0xd3, 0xd0, 0x63, 0xb7, 0x52, 0xb6, 0xfe, 0xa9, 0xc1, 0x07, 0x07,
	0x22, 0xef, 0x79, 0xfe, 0xc2, 0xed, 0xec, 0xcd, 0x6e, 0xc6, 0xe1, 0x39, 0x86, 0xd0, 0x0e, 0xd0,
	0x29, 0x51, 0x4f, 0x68, 0x56, 0x31, 0x5a, 0x86, 0xc8, 0x27, 0xd0, 0x09, 0x42, 0x91, 0x44, 0x74,
	0xeb, 0xa2, 0x6e, 0x2d, 0xa7, 0x68, 0xec, 0xb9, 0x92, 0xff, 0x10, 0x9a, 0x6b, 0xee, 0xb9, 0x6a,
	0x3b, 0x66, 0x5d, 0xeb, 0xaf, 0xb9, 0xf7, 0xc3, 0x36, 0x61, 0xe4, 0x3e, 0xd6, 0x90, 0xeb, 0xf3,
	0x88, 0xa7, 0x66, 0x03, 0x63, 0x8a, 0x3b, 0x57, 0x6b, 0x72, 0x05, 0x2d, 0x4f, 0x3d, 0x01, 0xf5,
	0x22, 0x66, 0x36, 0x87, 0x95, 0x51, 0xd3, 0xd9, 0x03, 0xe4, 0x1a, 0xda, 0xb8, 0x70, 0x7d, 0x9e,
	0xc5, 0xd2, 0x6c, 0x0d, 0x2b, 0x23, 0xc3, 0x01, 0x84, 0xe6, 0x0a, 0x51, 0xda, 0x9a, 0x10, 0x06,
	0xc2, 0x84, 0xa1, 0x31, 0x32, 0x9c, 0x26, 0x02, 0xdf, 0x05, 0x82, 0x3c, 0x80, 0xee, 0x32, 0x4c,
	0x85, 0x74, 0x0b, 0x8a, 0xd9, 0x46, 0x81, 0x0e, 0xa2, 0x33, 0x4d, 0x23, 0x16, 0x9c, 0x47, 0xb4,
	0x4c, 0xea, 0x20, 0xa9, 0xad, 0xc0, 0x82, 0xf3, 0x25, 0xdc, 0x43, 0x8e, 0xcf, 0x37, 0x49, 0xc4,
	0x24, 0x0b, 0xf6, 0xec, 0x73, 0x64, 0xf7, 0x55, 0x78, 0x5e, 0x44, 0x8b, 0x34, 0x1b, 0x10, 0x77,
	0x97, 0x34, 0x8c, 0xca, 0x39, 0x5d, 0xcc, 0xb9, 0x54, 0xb1, 0x67, 0x18, 0x2a, 0x12, 0x9e, 0x80,
	0x89, 0x09, 0x22, 0xf3, 0x7d, 0x26, 0xc4, 0x32, 0x8b, 0xf6, 0x49, 0x3d, 0x4c, 0xba, 0xa3, 0xe2,
	0x2f, 0x76, 0xe1, 0x22, 0x71, 0x0c, 0x97, 0x31, 0xfb, 0xbd, 0x38, 0x44, 0x9c, 0x6d, 0x3c, 0x96,
	0x9a, 0x17, 0x98, 0xd1, 0x53, 0x01, 0xe4, 0x3d, 0x47, 0x98, 0x3c, 0x85, 0xf3, 0x15, 0xa3, 0x91,
	0x5c, 0xb9, 0x29, 0x4b, 0x78, 0x2a, 0xcd, 0x4b, 0x6c, 0x93, 0x3b, 0x47, 0x6d, 0xf2, 0x2d, 0x72,
	0x9c, 0x8e, 0xe6, 0x3a, 0x48, 0xb5, 0xe6, 0xd0, 0xda, 0x85, 0x8e, 0x0b, 0xa7, 0x72, 0xb3, 0x70,
	0xfa, 0x50, 0x13, 0x3e, 0x4f, 0x19, 0x16, 0x95, 0xe1, 0xe8, 0x85, 0xf5, 0x35, 0x5c, 0x60, 0x75,
	0xab, 0xbb, 0xba, 0x5d, 0x83, 0xfc, 0x8c, 0x0a, 0xf3, 0x94, 0xd1, 0x5b, 0x2a, 0x90, 0x7b, 0xd0,
	0x10, 0x2b, 0x77, 0x19, 0x46, 0x2c, 0x6f, 0x87, 0xba, 0x58, 0x3d, 0x0b, 0x23, 0x66, 0x2d, 0x70,
	0x44, 0xcd, 0x79, 0xb2, 0x2d, 0x84, 0xef, 0x42, 0x5d, 0xf0, 0x2c, 0xf5, 0x0b, 0xe9, 0x7c, 0x95,
	0x1f, 0x5f, 0x86, 0x31, 0x8e, 0xbe, 0xdc, 0xa0, 0x0c, 0x59, 0xaf, 0xc0, 0x58, 0x70, 0xef, 0x3d,
	0x77, 0x56, 0xee, 0x24, 0xe3, 0xb0, 0x93, 0xfa, 0x50, 0xd3, 0x5d, 0xa4, 0x7b, 0x54, 0x2f, 0xc8,
	0x05, 0x18, 0x59, 0x1a, 0xe5, 0x4d, 0xa9, 0x7e, 0x4e, 0xff, 0xad, 0x41, 0x75, 0xc1, 0x3d, 0x41,
	0x96, 0x50, 0x55, 0x03, 0x94, 0x0c, 0x8e, 0xde, 0xb6, 0x34, 0x81, 0x07, 0xf7, 0xdf, 0x18, 0xd3,
	0x13, 0xc7, 0x7a, 0xf8, 0xd7, 0x7f, 0xff, 0xff, 0x7d, 0x76, 0x4d, 0x3e, 0xb2, 0x69, 0x92, 0x08,
	0x9f, 0x07, 0x7a, 0xe2, 0xfb, 0xa1, 0x9d, 0x67, 0xd8, 0x6a, 0x96, 0x92, 0x2d, 0x34, 0x8b, 0x61,
	0x45, 0x3e, 0x3e, 0xd2, 0x3b, 0x1a, 0x85, 0x83, 0xeb, 0x93, 0xf1, 0xdc, 0xf3, 0x33, 0xf4, 0x7c,
	0x44, 0x1e, 0xbc, 0xd5, 0xd3, 0xfe, 0x43, 0xdd, 0xe3, 0x9f, 0x64, 0x09, 0x75, 0x5d, 0x07, 0xe4,
	0xea, 0x48, 0xf8, 0xa0, 0x3c, 0x06, 0xfd, 0x62, 0x56, 0xfe, 0xc4, 0xc3, 0x60, 0xe7, 0x35, 0x42,
	0x2f, 0xcb, 0x7a, 0xfb, 0xf9, 0x9e, 0x56, 0xc6, 0x64, 0x0d, 0x55, 0x55, 0x14, 0x37, 0xae, 0xb2,
	0x54, 0x29, 0x27, 0x3c, 0x3e, 0x47, 0x8f, 0xf1, 0xe0, 0xe1, 0x49, 0x0f, 0xea, 0xe3, 0xc7, 0xd3,
	0xf6, 0x79, 0xb2, 0x55, 0x5e, 0x5b, 0xa8, 0x61, 0xc3, 0x92, 0xe3, 0xb7, 0x29, 0x7f, 0xad, 0x4e,
	0xb8, 0x7d, 0x85, 0x6e, 0x4f, 0xac, 0xe9, 0xbb, 0xdc, 0xde, 0xce, 0x19, 0xe7, 0x86, 0x3e, 0x66,
	0x5d, 0x37, 0xe6, 0x8d, 0xeb, 0x3c, 0xe8, 0xd7, 0x13, 0xe6, 0xf9, 0xd3, 0x8d, 0xdf, 0xe9, 0xe9,
	0x66, 0xad, 0x5f, 0x1a, 0x39, 0xec, 0xd5, 0xf1, 0xff, 0x83, 0x2f, 0x5e, 0x07, 0x00, 0x00, 0xff,
	0xff, 0x40, 0x06, 0xf5, 0xd4, 0x6e, 0x08, 0x00, 0x00,
}
