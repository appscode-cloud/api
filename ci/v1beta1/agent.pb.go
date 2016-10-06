// Code generated by protoc-gen-go.
// source: agent.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	agent.proto
	build.proto
	job.proto
	master.proto

It has these top-level messages:
	AgentListRequest
	AgentListResponse
	Agent
	AgentCreateRequest
	PortInfo
	AgentDescribeRequest
	AgentDescribeResponse
	AgentDeleteRequest
	AgentRestartRequest
	AgentRestartResponse
	BuildDescribeRequest
	BuildDescribeResponse
	BuildListRequest
	BuildListResponse
	Build
	JobListRequest
	JobListResponse
	JobBuildRequest
	JobDescribeRequest
	JobDescribeResponse
	JobHealth
	JobDeleteRequest
	JobCreateRequest
	JobCopyRequest
	Job
	MasterCreateRequest
	MasterDeleteRequest
*/
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AgentListRequest struct {
	// List of status to get the agent filterd on the status
	// values in
	//   PENDING
	//   FAILED
	//   ONLINE
	//   OFFLINE
	//   DELETED
	Status []string `protobuf:"bytes,1,rep,name=status" json:"status,omitempty"`
}

func (m *AgentListRequest) Reset()                    { *m = AgentListRequest{} }
func (m *AgentListRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentListRequest) ProtoMessage()               {}
func (*AgentListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AgentListResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Agents []*Agent       `protobuf:"bytes,2,rep,name=agents" json:"agents,omitempty"`
}

func (m *AgentListResponse) Reset()                    { *m = AgentListResponse{} }
func (m *AgentListResponse) String() string            { return proto.CompactTextString(m) }
func (*AgentListResponse) ProtoMessage()               {}
func (*AgentListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AgentListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AgentListResponse) GetAgents() []*Agent {
	if m != nil {
		return m.Agents
	}
	return nil
}

type Agent struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Status      string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	AgentStatus string `protobuf:"bytes,3,opt,name=agent_status,json=agentStatus" json:"agent_status,omitempty"`
	IsRefreshed int32  `protobuf:"varint,4,opt,name=is_refreshed,json=isRefreshed" json:"is_refreshed,omitempty"`
	CreatedAt   int64  `protobuf:"varint,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt   int64  `protobuf:"varint,6,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *Agent) Reset()                    { *m = Agent{} }
func (m *Agent) String() string            { return proto.CompactTextString(m) }
func (*Agent) ProtoMessage()               {}
func (*Agent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type AgentCreateRequest struct {
	Sku               string      `protobuf:"bytes,1,opt,name=sku" json:"sku,omitempty"`
	Executors         int32       `protobuf:"varint,2,opt,name=executors" json:"executors,omitempty"`
	Labels            string      `protobuf:"bytes,3,opt,name=labels" json:"labels,omitempty"`
	UserStartupScript string      `protobuf:"bytes,4,opt,name=user_startup_script,json=userStartupScript" json:"user_startup_script,omitempty"`
	SaltbaseVersion   string      `protobuf:"bytes,5,opt,name=saltbase_version,json=saltbaseVersion" json:"saltbase_version,omitempty"`
	CiStarterVersion  string      `protobuf:"bytes,6,opt,name=ci_starter_version,json=ciStarterVersion" json:"ci_starter_version,omitempty"`
	Ports             []*PortInfo `protobuf:"bytes,7,rep,name=ports" json:"ports,omitempty"`
	Role              string      `protobuf:"bytes,8,opt,name=role" json:"role,omitempty"`
}

func (m *AgentCreateRequest) Reset()                    { *m = AgentCreateRequest{} }
func (m *AgentCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentCreateRequest) ProtoMessage()               {}
func (*AgentCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AgentCreateRequest) GetPorts() []*PortInfo {
	if m != nil {
		return m.Ports
	}
	return nil
}

type PortInfo struct {
	Protocol  string `protobuf:"bytes,1,opt,name=protocol" json:"protocol,omitempty"`
	PortRange string `protobuf:"bytes,2,opt,name=port_range,json=portRange" json:"port_range,omitempty"`
}

func (m *PortInfo) Reset()                    { *m = PortInfo{} }
func (m *PortInfo) String() string            { return proto.CompactTextString(m) }
func (*PortInfo) ProtoMessage()               {}
func (*PortInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type AgentDescribeRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *AgentDescribeRequest) Reset()                    { *m = AgentDescribeRequest{} }
func (m *AgentDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentDescribeRequest) ProtoMessage()               {}
func (*AgentDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type AgentDescribeResponse struct {
	Status           *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Name             string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Executors        int64          `protobuf:"varint,3,opt,name=executors" json:"executors,omitempty"`
	AgentStatus      string         `protobuf:"bytes,4,opt,name=agent_status,json=agentStatus" json:"agent_status,omitempty"`
	AgentStatusCause string         `protobuf:"bytes,5,opt,name=agent_status_cause,json=agentStatusCause" json:"agent_status_cause,omitempty"`
	Label            string         `protobuf:"bytes,6,opt,name=label" json:"label,omitempty"`
	Provider         string         `protobuf:"bytes,7,opt,name=provider" json:"provider,omitempty"`
	Sku              string         `protobuf:"bytes,8,opt,name=sku" json:"sku,omitempty"`
	StartupScript    string         `protobuf:"bytes,9,opt,name=startup_script,json=startupScript" json:"startup_script,omitempty"`
	CreatedAt        int64          `protobuf:"varint,10,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	UpdatedAt        int64          `protobuf:"varint,11,opt,name=updated_at,json=updatedAt" json:"updated_at,omitempty"`
}

func (m *AgentDescribeResponse) Reset()                    { *m = AgentDescribeResponse{} }
func (m *AgentDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*AgentDescribeResponse) ProtoMessage()               {}
func (*AgentDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AgentDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type AgentDeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *AgentDeleteRequest) Reset()                    { *m = AgentDeleteRequest{} }
func (m *AgentDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentDeleteRequest) ProtoMessage()               {}
func (*AgentDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type AgentRestartRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *AgentRestartRequest) Reset()                    { *m = AgentRestartRequest{} }
func (m *AgentRestartRequest) String() string            { return proto.CompactTextString(m) }
func (*AgentRestartRequest) ProtoMessage()               {}
func (*AgentRestartRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type AgentRestartResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *AgentRestartResponse) Reset()                    { *m = AgentRestartResponse{} }
func (m *AgentRestartResponse) String() string            { return proto.CompactTextString(m) }
func (*AgentRestartResponse) ProtoMessage()               {}
func (*AgentRestartResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AgentRestartResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*AgentListRequest)(nil), "ci.v1beta1.AgentListRequest")
	proto.RegisterType((*AgentListResponse)(nil), "ci.v1beta1.AgentListResponse")
	proto.RegisterType((*Agent)(nil), "ci.v1beta1.Agent")
	proto.RegisterType((*AgentCreateRequest)(nil), "ci.v1beta1.AgentCreateRequest")
	proto.RegisterType((*PortInfo)(nil), "ci.v1beta1.PortInfo")
	proto.RegisterType((*AgentDescribeRequest)(nil), "ci.v1beta1.AgentDescribeRequest")
	proto.RegisterType((*AgentDescribeResponse)(nil), "ci.v1beta1.AgentDescribeResponse")
	proto.RegisterType((*AgentDeleteRequest)(nil), "ci.v1beta1.AgentDeleteRequest")
	proto.RegisterType((*AgentRestartRequest)(nil), "ci.v1beta1.AgentRestartRequest")
	proto.RegisterType((*AgentRestartResponse)(nil), "ci.v1beta1.AgentRestartResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Agents service

type AgentsClient interface {
	List(ctx context.Context, in *AgentListRequest, opts ...grpc.CallOption) (*AgentListResponse, error)
	Describe(ctx context.Context, in *AgentDescribeRequest, opts ...grpc.CallOption) (*AgentDescribeResponse, error)
	Create(ctx context.Context, in *AgentCreateRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error)
	Delete(ctx context.Context, in *AgentDeleteRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error)
	Restart(ctx context.Context, in *AgentRestartRequest, opts ...grpc.CallOption) (*AgentRestartResponse, error)
}

type agentsClient struct {
	cc *grpc.ClientConn
}

func NewAgentsClient(cc *grpc.ClientConn) AgentsClient {
	return &agentsClient{cc}
}

func (c *agentsClient) List(ctx context.Context, in *AgentListRequest, opts ...grpc.CallOption) (*AgentListResponse, error) {
	out := new(AgentListResponse)
	err := grpc.Invoke(ctx, "/ci.v1beta1.Agents/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) Describe(ctx context.Context, in *AgentDescribeRequest, opts ...grpc.CallOption) (*AgentDescribeResponse, error) {
	out := new(AgentDescribeResponse)
	err := grpc.Invoke(ctx, "/ci.v1beta1.Agents/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) Create(ctx context.Context, in *AgentCreateRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error) {
	out := new(dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/ci.v1beta1.Agents/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) Delete(ctx context.Context, in *AgentDeleteRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error) {
	out := new(dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/ci.v1beta1.Agents/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentsClient) Restart(ctx context.Context, in *AgentRestartRequest, opts ...grpc.CallOption) (*AgentRestartResponse, error) {
	out := new(AgentRestartResponse)
	err := grpc.Invoke(ctx, "/ci.v1beta1.Agents/Restart", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Agents service

type AgentsServer interface {
	List(context.Context, *AgentListRequest) (*AgentListResponse, error)
	Describe(context.Context, *AgentDescribeRequest) (*AgentDescribeResponse, error)
	Create(context.Context, *AgentCreateRequest) (*dtypes.LongRunningResponse, error)
	Delete(context.Context, *AgentDeleteRequest) (*dtypes.LongRunningResponse, error)
	Restart(context.Context, *AgentRestartRequest) (*AgentRestartResponse, error)
}

func RegisterAgentsServer(s *grpc.Server, srv AgentsServer) {
	s.RegisterService(&_Agents_serviceDesc, srv)
}

func _Agents_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.v1beta1.Agents/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).List(ctx, req.(*AgentListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.v1beta1.Agents/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).Describe(ctx, req.(*AgentDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.v1beta1.Agents/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).Create(ctx, req.(*AgentCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.v1beta1.Agents/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).Delete(ctx, req.(*AgentDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agents_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentRestartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.v1beta1.Agents/Restart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).Restart(ctx, req.(*AgentRestartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agents_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ci.v1beta1.Agents",
	HandlerType: (*AgentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Agents_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Agents_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Agents_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Agents_Delete_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _Agents_Restart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("agent.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 791 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x55, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0x56, 0xfe, 0x9c, 0xf8, 0xe4, 0xde, 0xde, 0x74, 0xda, 0x5b, 0x45, 0xb9, 0xed, 0x6d, 0x6a,
	0x41, 0x49, 0xa3, 0xca, 0x56, 0x8b, 0x10, 0x12, 0x12, 0x48, 0xa5, 0x65, 0x81, 0xd4, 0x05, 0x72,
	0x24, 0x16, 0x6c, 0xa2, 0x89, 0x33, 0x0d, 0x16, 0xc6, 0xe3, 0xce, 0x8c, 0x2b, 0x10, 0x42, 0x48,
	0x6c, 0x59, 0xf2, 0x0c, 0x3c, 0x03, 0x0f, 0xc2, 0x2b, 0xb0, 0x63, 0xc1, 0x2b, 0xa0, 0xf9, 0xb1,
	0xe3, 0x24, 0x6d, 0xda, 0x6e, 0x22, 0xcf, 0x77, 0x3e, 0x9f, 0xcf, 0xe7, 0x9c, 0x6f, 0x4e, 0xa0,
	0x89, 0x27, 0x24, 0x16, 0x6e, 0xc2, 0xa8, 0xa0, 0x08, 0x82, 0xd0, 0xbd, 0x38, 0x18, 0x11, 0x81,
	0x0f, 0x3a, 0x9b, 0x13, 0x4a, 0x27, 0x11, 0xf1, 0x70, 0x12, 0x7a, 0x38, 0x8e, 0xa9, 0xc0, 0x22,
	0xa4, 0x31, 0xd7, 0xcc, 0xce, 0x86, 0x84, 0xc7, 0xe2, 0x7d, 0x42, 0xb8, 0xa7, 0x7e, 0x35, 0xee,
	0xf4, 0xa1, 0x75, 0x24, 0x13, 0x9e, 0x86, 0x5c, 0xf8, 0xe4, 0x3c, 0x25, 0x5c, 0xa0, 0x0d, 0xb0,
	0xb8, 0xc0, 0x22, 0xe5, 0xed, 0x52, 0xb7, 0xd2, 0xb3, 0x7d, 0x73, 0x72, 0xce, 0x60, 0xb5, 0xc0,
	0xe5, 0x09, 0x8d, 0x39, 0x41, 0xbb, 0x05, 0x72, 0xa9, 0xd7, 0x3c, 0x5c, 0x71, 0xb5, 0x8a, 0x3b,
	0x50, 0x68, 0xf6, 0x32, 0xda, 0x03, 0x4b, 0x7d, 0x39, 0x6f, 0x97, 0xbb, 0x95, 0x5e, 0xf3, 0x70,
	0xd5, 0x9d, 0x7e, 0xbb, 0xab, 0xd2, 0xfa, 0x86, 0xe0, 0x7c, 0x2f, 0x41, 0x4d, 0x21, 0x08, 0x41,
	0x35, 0xc6, 0x6f, 0x89, 0x4a, 0x6d, 0xfb, 0xea, 0xb9, 0xf0, 0x75, 0x65, 0x85, 0x66, 0x02, 0x3b,
	0xf0, 0x97, 0x7a, 0x7f, 0x68, 0xa2, 0x15, 0x15, 0xd5, 0xed, 0x1a, 0xe4, 0x94, 0x90, 0x0f, 0x19,
	0x39, 0x63, 0x84, 0xbf, 0x26, 0xe3, 0x76, 0xb5, 0x5b, 0xea, 0xd5, 0xfc, 0x66, 0xc8, 0xfd, 0x0c,
	0x42, 0x5b, 0x00, 0x01, 0x23, 0x58, 0x90, 0xf1, 0x10, 0x8b, 0x76, 0xad, 0x5b, 0xea, 0x55, 0x7c,
	0xdb, 0x20, 0x47, 0x42, 0x86, 0xd3, 0x64, 0x9c, 0x85, 0x2d, 0x1d, 0x36, 0xc8, 0x91, 0x70, 0xbe,
	0x95, 0x01, 0xa9, 0x2f, 0x3f, 0x56, 0x6f, 0x64, 0x0d, 0x6d, 0x41, 0x85, 0xbf, 0x49, 0x4d, 0x15,
	0xf2, 0x11, 0x6d, 0x82, 0x4d, 0xde, 0x91, 0x20, 0x15, 0x94, 0xe9, 0x3a, 0x6a, 0xfe, 0x14, 0x90,
	0x25, 0x46, 0x78, 0x44, 0xa2, 0xac, 0x08, 0x73, 0x42, 0x2e, 0xac, 0xa5, 0x9c, 0x30, 0x59, 0x21,
	0x13, 0x69, 0x32, 0xe4, 0x01, 0x0b, 0x13, 0xa1, 0xca, 0xb0, 0xfd, 0x55, 0x19, 0x1a, 0xe8, 0xc8,
	0x40, 0x05, 0xd0, 0x1e, 0xb4, 0x38, 0x8e, 0xc4, 0x08, 0x73, 0x32, 0xbc, 0x20, 0x8c, 0x87, 0x34,
	0x56, 0x25, 0xd9, 0xfe, 0x3f, 0x19, 0xfe, 0x52, 0xc3, 0x68, 0x1f, 0x50, 0x10, 0xea, 0xc4, 0x84,
	0xe5, 0x64, 0x4b, 0x91, 0x5b, 0x41, 0x38, 0xd0, 0x81, 0x8c, 0xdd, 0x87, 0x5a, 0x42, 0x99, 0xe0,
	0xed, 0xba, 0x9a, 0xe5, 0x7a, 0x71, 0x96, 0x2f, 0x28, 0x13, 0xcf, 0xe3, 0x33, 0xea, 0x6b, 0x8a,
	0x9c, 0x21, 0xa3, 0x11, 0x69, 0x37, 0xf4, 0x0c, 0xe5, 0xb3, 0xf3, 0x0c, 0x1a, 0x19, 0x0d, 0x75,
	0xa0, 0xa1, 0xac, 0x18, 0xd0, 0xc8, 0x74, 0x28, 0x3f, 0xcb, 0x76, 0xcb, 0x24, 0x43, 0x86, 0xe3,
	0x09, 0x31, 0xf3, 0xb6, 0x25, 0xe2, 0x4b, 0xc0, 0xe9, 0xc3, 0xba, 0xea, 0xf6, 0x09, 0x91, 0x9d,
	0x18, 0xe5, 0xfd, 0xbe, 0xc4, 0x36, 0xce, 0xaf, 0x32, 0xfc, 0x3b, 0x47, 0xbe, 0xa5, 0x83, 0xb3,
	0xac, 0xe5, 0x82, 0x19, 0x67, 0xe6, 0x58, 0xd1, 0x76, 0x98, 0xce, 0x71, 0xde, 0x92, 0xd5, 0x45,
	0x4b, 0xee, 0x03, 0x2a, 0x52, 0x86, 0x01, 0x4e, 0x39, 0x31, 0x43, 0x6a, 0x15, 0x88, 0xc7, 0x12,
	0x47, 0xeb, 0x50, 0x53, 0x56, 0x30, 0x83, 0xd1, 0x07, 0xd3, 0xc1, 0x8b, 0x70, 0x4c, 0x58, 0xbb,
	0x9e, 0x77, 0x50, 0x9d, 0x33, 0xeb, 0x35, 0xa6, 0xd6, 0xbb, 0x0b, 0x2b, 0x73, 0xfe, 0xb1, 0x55,
	0xf0, 0x6f, 0x3e, 0xe3, 0x9d, 0xd9, 0x8b, 0x00, 0xcb, 0x2f, 0x42, 0x73, 0xfe, 0x22, 0xf4, 0xcc,
	0x3d, 0x38, 0x21, 0x11, 0x11, 0x4b, 0xe7, 0xb2, 0x07, 0x6b, 0xfa, 0xf6, 0x13, 0xa5, 0xbf, 0x8c,
	0xfa, 0xc4, 0x8c, 0x3b, 0xa7, 0xde, 0x6e, 0x80, 0x87, 0xbf, 0xab, 0x60, 0xa9, 0x04, 0x1c, 0x45,
	0x50, 0x95, 0x5b, 0x0c, 0x6d, 0x2e, 0x6c, 0xa1, 0xc2, 0x22, 0xec, 0x6c, 0x5d, 0x11, 0xd5, 0xba,
	0xce, 0xbd, 0xcf, 0x3f, 0x7e, 0x7e, 0x2d, 0xef, 0xa0, 0x6d, 0x0f, 0x27, 0x09, 0x0f, 0xe8, 0x58,
	0x2f, 0xdf, 0x20, 0xf4, 0xcc, 0x3b, 0x9e, 0x5e, 0x68, 0xe8, 0x13, 0x34, 0x32, 0xd7, 0xa1, 0xee,
	0x42, 0xce, 0x39, 0xf7, 0x76, 0x76, 0x96, 0x30, 0x8c, 0xb2, 0xab, 0x94, 0x7b, 0x68, 0xf7, 0x1a,
	0x65, 0xef, 0x83, 0x6c, 0xdc, 0x47, 0x74, 0x0e, 0x96, 0xde, 0x48, 0xe8, 0xff, 0x85, 0xe4, 0x33,
	0xab, 0xaa, 0xf3, 0x5f, 0xd6, 0xbb, 0x53, 0x1a, 0x4f, 0xfc, 0x34, 0x8e, 0xc3, 0x78, 0x92, 0xcb,
	0xf6, 0x95, 0xec, 0x1d, 0xe7, 0xba, 0x82, 0x1f, 0x95, 0xfa, 0x28, 0x05, 0x4b, 0x0f, 0xff, 0x12,
	0xc9, 0x19, 0x57, 0x2c, 0x97, 0x34, 0x95, 0xf6, 0x6f, 0x5a, 0xe9, 0x97, 0x12, 0xd4, 0x8d, 0x3f,
	0xd0, 0xf6, 0xe2, 0x5f, 0xcc, 0x8c, 0xc9, 0x3a, 0xdd, 0xab, 0x09, 0x46, 0xfe, 0xb1, 0x92, 0x7f,
	0xe8, 0x3c, 0xb8, 0x99, 0xbc, 0x87, 0x03, 0xf5, 0xa7, 0xeb, 0x31, 0x32, 0xa2, 0x54, 0x3c, 0xb5,
	0x5f, 0xd5, 0x0d, 0x6f, 0x64, 0xa9, 0xa5, 0x76, 0xff, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0e,
	0x42, 0xc5, 0x59, 0xc0, 0x07, 0x00, 0x00,
}
