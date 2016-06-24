// Code generated by protoc-gen-go.
// source: slave.proto
// DO NOT EDIT!

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

type SlaveType int32

const (
	SlaveType_ALL       SlaveType = 0
	SlaveType_CRETAED   SlaveType = 1
	SlaveType_DELETED   SlaveType = 2
	SlaveType_REFRESHED SlaveType = 3
)

var SlaveType_name = map[int32]string{
	0: "ALL",
	1: "CRETAED",
	2: "DELETED",
	3: "REFRESHED",
}
var SlaveType_value = map[string]int32{
	"ALL":       0,
	"CRETAED":   1,
	"DELETED":   2,
	"REFRESHED": 3,
}

func (x SlaveType) String() string {
	return proto.EnumName(SlaveType_name, int32(x))
}
func (SlaveType) EnumDescriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type SlaveListRequest struct {
	From string    `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	To   string    `protobuf:"bytes,2,opt,name=to" json:"to,omitempty"`
	Type SlaveType `protobuf:"varint,3,opt,name=type,enum=ci.SlaveType" json:"type,omitempty"`
}

func (m *SlaveListRequest) Reset()                    { *m = SlaveListRequest{} }
func (m *SlaveListRequest) String() string            { return proto.CompactTextString(m) }
func (*SlaveListRequest) ProtoMessage()               {}
func (*SlaveListRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type SlaveListResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Slaves []*Slave       `protobuf:"bytes,2,rep,name=slaves" json:"slaves,omitempty"`
}

func (m *SlaveListResponse) Reset()                    { *m = SlaveListResponse{} }
func (m *SlaveListResponse) String() string            { return proto.CompactTextString(m) }
func (*SlaveListResponse) ProtoMessage()               {}
func (*SlaveListResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *SlaveListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *SlaveListResponse) GetSlaves() []*Slave {
	if m != nil {
		return m.Slaves
	}
	return nil
}

type Slave struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Status      string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	IsDeleted   int32  `protobuf:"varint,3,opt,name=is_deleted" json:"is_deleted,omitempty"`
	IsRefreshed int32  `protobuf:"varint,4,opt,name=is_refreshed" json:"is_refreshed,omitempty"`
	CreatedAt   string `protobuf:"bytes,5,opt,name=created_at" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,6,opt,name=updated_at" json:"updated_at,omitempty"`
}

func (m *Slave) Reset()                    { *m = Slave{} }
func (m *Slave) String() string            { return proto.CompactTextString(m) }
func (*Slave) ProtoMessage()               {}
func (*Slave) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

type SlaveCreateRequest struct {
	Sku               string      `protobuf:"bytes,1,opt,name=sku" json:"sku,omitempty"`
	Executors         int32       `protobuf:"varint,2,opt,name=executors" json:"executors,omitempty"`
	Labels            string      `protobuf:"bytes,3,opt,name=labels" json:"labels,omitempty"`
	UserStartupScript string      `protobuf:"bytes,4,opt,name=user_startup_script" json:"user_startup_script,omitempty"`
	SaltbaseVersion   string      `protobuf:"bytes,5,opt,name=saltbase_version" json:"saltbase_version,omitempty"`
	CiStarterVersion  string      `protobuf:"bytes,6,opt,name=ci_starter_version" json:"ci_starter_version,omitempty"`
	Ports             []*PortInfo `protobuf:"bytes,7,rep,name=ports" json:"ports,omitempty"`
	Role              string      `protobuf:"bytes,8,opt,name=role" json:"role,omitempty"`
}

func (m *SlaveCreateRequest) Reset()                    { *m = SlaveCreateRequest{} }
func (m *SlaveCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*SlaveCreateRequest) ProtoMessage()               {}
func (*SlaveCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *SlaveCreateRequest) GetPorts() []*PortInfo {
	if m != nil {
		return m.Ports
	}
	return nil
}

type PortInfo struct {
	Protocol  string `protobuf:"bytes,1,opt,name=protocol" json:"protocol,omitempty"`
	PortRange string `protobuf:"bytes,2,opt,name=port_range" json:"port_range,omitempty"`
}

func (m *PortInfo) Reset()                    { *m = PortInfo{} }
func (m *PortInfo) String() string            { return proto.CompactTextString(m) }
func (*PortInfo) ProtoMessage()               {}
func (*PortInfo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

type SlaveDescribeRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *SlaveDescribeRequest) Reset()                    { *m = SlaveDescribeRequest{} }
func (m *SlaveDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*SlaveDescribeRequest) ProtoMessage()               {}
func (*SlaveDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

type SlaveDescribeResponse struct {
	Status        *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Name          string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Executors     int64          `protobuf:"varint,3,opt,name=executors" json:"executors,omitempty"`
	NodeStatus    string         `protobuf:"bytes,4,opt,name=node_status" json:"node_status,omitempty"`
	OfflineReason string         `protobuf:"bytes,5,opt,name=offline_reason" json:"offline_reason,omitempty"`
	Label         string         `protobuf:"bytes,6,opt,name=label" json:"label,omitempty"`
	Provider      string         `protobuf:"bytes,7,opt,name=provider" json:"provider,omitempty"`
	Sku           string         `protobuf:"bytes,8,opt,name=sku" json:"sku,omitempty"`
	StartupScript string         `protobuf:"bytes,9,opt,name=startup_script" json:"startup_script,omitempty"`
	CreatedAt     string         `protobuf:"bytes,10,opt,name=created_at" json:"created_at,omitempty"`
	UpdatedAt     string         `protobuf:"bytes,11,opt,name=updated_at" json:"updated_at,omitempty"`
	IsDeleted     int32          `protobuf:"varint,12,opt,name=is_deleted" json:"is_deleted,omitempty"`
	IsRefreshed   int32          `protobuf:"varint,13,opt,name=is_refreshed" json:"is_refreshed,omitempty"`
}

func (m *SlaveDescribeResponse) Reset()                    { *m = SlaveDescribeResponse{} }
func (m *SlaveDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*SlaveDescribeResponse) ProtoMessage()               {}
func (*SlaveDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *SlaveDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type SlaveDeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *SlaveDeleteRequest) Reset()                    { *m = SlaveDeleteRequest{} }
func (m *SlaveDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*SlaveDeleteRequest) ProtoMessage()               {}
func (*SlaveDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

type SlaveRestartRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *SlaveRestartRequest) Reset()                    { *m = SlaveRestartRequest{} }
func (m *SlaveRestartRequest) String() string            { return proto.CompactTextString(m) }
func (*SlaveRestartRequest) ProtoMessage()               {}
func (*SlaveRestartRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

type SlaveRestartResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *SlaveRestartResponse) Reset()                    { *m = SlaveRestartResponse{} }
func (m *SlaveRestartResponse) String() string            { return proto.CompactTextString(m) }
func (*SlaveRestartResponse) ProtoMessage()               {}
func (*SlaveRestartResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{9} }

func (m *SlaveRestartResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*SlaveListRequest)(nil), "ci.SlaveListRequest")
	proto.RegisterType((*SlaveListResponse)(nil), "ci.SlaveListResponse")
	proto.RegisterType((*Slave)(nil), "ci.Slave")
	proto.RegisterType((*SlaveCreateRequest)(nil), "ci.SlaveCreateRequest")
	proto.RegisterType((*PortInfo)(nil), "ci.PortInfo")
	proto.RegisterType((*SlaveDescribeRequest)(nil), "ci.SlaveDescribeRequest")
	proto.RegisterType((*SlaveDescribeResponse)(nil), "ci.SlaveDescribeResponse")
	proto.RegisterType((*SlaveDeleteRequest)(nil), "ci.SlaveDeleteRequest")
	proto.RegisterType((*SlaveRestartRequest)(nil), "ci.SlaveRestartRequest")
	proto.RegisterType((*SlaveRestartResponse)(nil), "ci.SlaveRestartResponse")
	proto.RegisterEnum("ci.SlaveType", SlaveType_name, SlaveType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Slaves service

type SlavesClient interface {
	List(ctx context.Context, in *SlaveListRequest, opts ...grpc.CallOption) (*SlaveListResponse, error)
	Describe(ctx context.Context, in *SlaveDescribeRequest, opts ...grpc.CallOption) (*SlaveDescribeResponse, error)
	Create(ctx context.Context, in *SlaveCreateRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error)
	Delete(ctx context.Context, in *SlaveDeleteRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error)
	Restart(ctx context.Context, in *SlaveRestartRequest, opts ...grpc.CallOption) (*SlaveRestartResponse, error)
}

type slavesClient struct {
	cc *grpc.ClientConn
}

func NewSlavesClient(cc *grpc.ClientConn) SlavesClient {
	return &slavesClient{cc}
}

func (c *slavesClient) List(ctx context.Context, in *SlaveListRequest, opts ...grpc.CallOption) (*SlaveListResponse, error) {
	out := new(SlaveListResponse)
	err := grpc.Invoke(ctx, "/ci.Slaves/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slavesClient) Describe(ctx context.Context, in *SlaveDescribeRequest, opts ...grpc.CallOption) (*SlaveDescribeResponse, error) {
	out := new(SlaveDescribeResponse)
	err := grpc.Invoke(ctx, "/ci.Slaves/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slavesClient) Create(ctx context.Context, in *SlaveCreateRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error) {
	out := new(dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/ci.Slaves/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slavesClient) Delete(ctx context.Context, in *SlaveDeleteRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error) {
	out := new(dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/ci.Slaves/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *slavesClient) Restart(ctx context.Context, in *SlaveRestartRequest, opts ...grpc.CallOption) (*SlaveRestartResponse, error) {
	out := new(SlaveRestartResponse)
	err := grpc.Invoke(ctx, "/ci.Slaves/Restart", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Slaves service

type SlavesServer interface {
	List(context.Context, *SlaveListRequest) (*SlaveListResponse, error)
	Describe(context.Context, *SlaveDescribeRequest) (*SlaveDescribeResponse, error)
	Create(context.Context, *SlaveCreateRequest) (*dtypes.LongRunningResponse, error)
	Delete(context.Context, *SlaveDeleteRequest) (*dtypes.LongRunningResponse, error)
	Restart(context.Context, *SlaveRestartRequest) (*SlaveRestartResponse, error)
}

func RegisterSlavesServer(s *grpc.Server, srv SlavesServer) {
	s.RegisterService(&_Slaves_serviceDesc, srv)
}

func _Slaves_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlaveListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlavesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Slaves/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlavesServer).List(ctx, req.(*SlaveListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Slaves_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlaveDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlavesServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Slaves/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlavesServer).Describe(ctx, req.(*SlaveDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Slaves_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlaveCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlavesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Slaves/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlavesServer).Create(ctx, req.(*SlaveCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Slaves_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlaveDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlavesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Slaves/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlavesServer).Delete(ctx, req.(*SlaveDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Slaves_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SlaveRestartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SlavesServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ci.Slaves/Restart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SlavesServer).Restart(ctx, req.(*SlaveRestartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Slaves_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ci.Slaves",
	HandlerType: (*SlavesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Slaves_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Slaves_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Slaves_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Slaves_Delete_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _Slaves_Restart_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

func init() { proto.RegisterFile("slave.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 751 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0xf2, 0x1f, 0xdf, 0xa4, 0x51, 0x3a, 0xfd, 0xf9, 0x5c, 0xb7, 0xaa, 0xc0, 0x2d, 0x02,
	0x05, 0x14, 0x97, 0x20, 0xb1, 0x40, 0x62, 0x51, 0x35, 0x41, 0x54, 0x8a, 0x10, 0x4a, 0xbb, 0x61,
	0x65, 0x39, 0xce, 0x24, 0x18, 0x5c, 0x8f, 0x99, 0x19, 0x07, 0x10, 0x62, 0xc3, 0x2b, 0xb0, 0x60,
	0xc5, 0xcb, 0xf0, 0x0a, 0xbc, 0x02, 0x0f, 0xc2, 0xcc, 0x78, 0x6c, 0xd2, 0x34, 0xaa, 0xb2, 0x49,
	0x3c, 0x67, 0xee, 0xdc, 0x73, 0xee, 0xb9, 0x77, 0x06, 0x1a, 0x2c, 0xf4, 0xe6, 0xb8, 0x1b, 0x53,
	0xc2, 0x09, 0x2a, 0xfa, 0x81, 0x75, 0x30, 0x23, 0x64, 0x16, 0x62, 0xc7, 0x8b, 0x03, 0xc7, 0x8b,
	0x22, 0xc2, 0x3d, 0x1e, 0x90, 0x88, 0xa5, 0x11, 0xd6, 0xae, 0x84, 0x27, 0xfc, 0x73, 0x8c, 0x99,
	0xa3, 0x7e, 0x53, 0xdc, 0x3e, 0x87, 0xf6, 0x85, 0x4c, 0x34, 0x0c, 0x18, 0x1f, 0xe1, 0x0f, 0x09,
	0x66, 0x1c, 0x35, 0xa1, 0x3c, 0xa5, 0xe4, 0xca, 0x2c, 0xdc, 0x29, 0x3c, 0x30, 0x10, 0x40, 0x91,
	0x13, 0xb3, 0xa8, 0xbe, 0xf7, 0xa1, 0x2c, 0x0f, 0x9b, 0x25, 0xb1, 0x6a, 0xf5, 0x36, 0xba, 0x7e,
	0xd0, 0x55, 0xa7, 0x2f, 0x05, 0x68, 0xbf, 0x82, 0xcd, 0x85, 0x54, 0x2c, 0x16, 0xe4, 0x18, 0x1d,
	0x42, 0x95, 0x09, 0x25, 0x09, 0x53, 0xd9, 0x1a, 0xbd, 0x56, 0x37, 0x15, 0xd1, 0xbd, 0x50, 0x28,
	0xda, 0x13, 0xfb, 0xf2, 0x10, 0x13, 0x0c, 0x25, 0xb1, 0x6f, 0xe4, 0x39, 0xed, 0x8f, 0x50, 0x51,
	0x1f, 0x52, 0x4f, 0xe4, 0x5d, 0x61, 0xad, 0xa7, 0x95, 0x67, 0x4c, 0x35, 0x09, 0x81, 0x01, 0x73,
	0x27, 0x38, 0xc4, 0x1c, 0x4f, 0x94, 0xb2, 0x0a, 0xda, 0x86, 0xa6, 0xc0, 0x28, 0x9e, 0x52, 0xcc,
	0xde, 0x0a, 0xb4, 0xac, 0x50, 0x11, 0xe9, 0x53, 0xec, 0x89, 0x30, 0xd7, 0xe3, 0x66, 0x25, 0x3b,
	0x9d, 0xc4, 0x93, 0x0c, 0xab, 0x4a, 0xcc, 0xfe, 0x55, 0x00, 0xa4, 0x98, 0xcf, 0x54, 0x74, 0x66,
	0x4b, 0x03, 0x4a, 0xec, 0x7d, 0xa2, 0x55, 0x6c, 0x82, 0x81, 0x3f, 0x61, 0x3f, 0xe1, 0x84, 0xa6,
	0x42, 0x2a, 0x52, 0x58, 0xe8, 0x8d, 0x71, 0xc8, 0x94, 0x08, 0x69, 0xd6, 0x56, 0xc2, 0x30, 0x75,
	0x85, 0x5a, 0xca, 0x93, 0xd8, 0x65, 0x3e, 0x0d, 0x62, 0xae, 0xb4, 0x18, 0xc8, 0x84, 0x36, 0xf3,
	0x42, 0x3e, 0xf6, 0x18, 0x76, 0xe7, 0x98, 0x32, 0xd1, 0x2a, 0xad, 0xc8, 0x02, 0xe4, 0x07, 0xe9,
	0x21, 0x71, 0x38, 0xdb, 0xab, 0xea, 0x94, 0x95, 0x98, 0x50, 0xce, 0xcc, 0x9a, 0x32, 0xab, 0x29,
	0xcd, 0x7a, 0x2d, 0x80, 0xf3, 0x68, 0x4a, 0xa4, 0x4d, 0x94, 0x84, 0xd8, 0xac, 0xab, 0x22, 0x4e,
	0xa0, 0x9e, 0xef, 0xb4, 0xa1, 0xae, 0xba, 0xed, 0x93, 0x50, 0xcb, 0x17, 0x65, 0xcb, 0x44, 0x2e,
	0xf5, 0xa2, 0x19, 0x4e, 0x8d, 0xb4, 0x8f, 0x61, 0x5b, 0x55, 0xdd, 0xc7, 0x52, 0xe9, 0x18, 0x2f,
	0x8c, 0xc3, 0x3f, 0xfb, 0xed, 0x9f, 0x45, 0xd8, 0x59, 0x0a, 0x5b, 0xb3, 0xd5, 0x59, 0x9e, 0xe2,
	0x4d, 0x03, 0xa5, 0x61, 0x25, 0xb4, 0x05, 0x8d, 0x88, 0x4c, 0xb0, 0xab, 0xb3, 0xa4, 0x46, 0xed,
	0x42, 0x8b, 0x4c, 0xa7, 0x61, 0x10, 0x61, 0xd1, 0x4f, 0x8f, 0xe5, 0x36, 0x6d, 0x40, 0x45, 0xb9,
	0xad, 0x9d, 0x49, 0x4b, 0x9c, 0x07, 0x13, 0x4c, 0x85, 0x39, 0x12, 0xd1, 0xed, 0xaa, 0x67, 0x59,
	0x96, 0xda, 0x60, 0x64, 0x3e, 0x2c, 0x8c, 0x04, 0xac, 0x18, 0x89, 0xc6, 0x8a, 0x21, 0x6b, 0xae,
	0x1c, 0xb2, 0x0d, 0x89, 0xda, 0xb6, 0x9e, 0x9d, 0xbe, 0x8a, 0x5d, 0xed, 0xe1, 0x11, 0x6c, 0xa9,
	0x18, 0x61, 0x9d, 0x14, 0xb5, 0x3a, 0xe8, 0xa9, 0x6e, 0x47, 0x1e, 0xb4, 0x9e, 0xcd, 0x9d, 0xe7,
	0x60, 0xe4, 0x77, 0x12, 0xd5, 0xa0, 0x74, 0x3a, 0x1c, 0xb6, 0xff, 0x13, 0x6e, 0xd4, 0xce, 0x46,
	0x83, 0xcb, 0xd3, 0x41, 0xbf, 0x5d, 0x90, 0x8b, 0xfe, 0x60, 0x38, 0xb8, 0x14, 0x8b, 0xa2, 0x30,
	0xd2, 0x18, 0x0d, 0x5e, 0x8c, 0x06, 0x17, 0x2f, 0xc5, 0xb2, 0xd4, 0xfb, 0x51, 0x86, 0xaa, 0x3a,
	0xcf, 0xd0, 0x1b, 0x28, 0xcb, 0xbb, 0x8c, 0xb6, 0xf3, 0x3b, 0xb9, 0xf0, 0x4a, 0x58, 0x3b, 0x4b,
	0x68, 0x2a, 0xcf, 0x3e, 0xfe, 0xf6, 0xfb, 0xcf, 0xf7, 0xe2, 0x21, 0x3a, 0x10, 0x0f, 0x51, 0xcc,
	0x7c, 0xd1, 0x4f, 0xf5, 0x22, 0xf9, 0x81, 0x33, 0x3f, 0xe9, 0x3e, 0x76, 0xd2, 0xcb, 0x8e, 0x42,
	0xa8, 0x67, 0xf3, 0x83, 0xcc, 0x3c, 0xd1, 0xd2, 0xe4, 0x59, 0x7b, 0x2b, 0x76, 0x34, 0xcd, 0x43,
	0x45, 0x73, 0x0f, 0x1d, 0xdd, 0x46, 0xe3, 0x7c, 0x91, 0x7e, 0x7e, 0x45, 0x53, 0xa8, 0xa6, 0x57,
	0x19, 0xed, 0xe6, 0x19, 0xaf, 0xdd, 0x6d, 0x6b, 0x3f, 0x33, 0x71, 0x48, 0xa2, 0xd9, 0x28, 0x89,
	0xa2, 0x40, 0xfc, 0x65, 0x5c, 0xf7, 0x15, 0xd7, 0x5d, 0xeb, 0xd6, 0x92, 0x9e, 0x15, 0x3a, 0xe8,
	0x1d, 0x54, 0xd3, 0xb6, 0x2f, 0xf0, 0x5c, 0x9b, 0x83, 0xdb, 0x79, 0x74, 0x4d, 0x9d, 0xb5, 0x6a,
	0x8a, 0xa1, 0xa6, 0x27, 0x03, 0xfd, 0x9f, 0x93, 0x5d, 0x1f, 0x28, 0xcb, 0xbc, 0xb9, 0xa1, 0xa9,
	0x7a, 0x8a, 0xea, 0x11, 0xea, 0xac, 0x41, 0xe5, 0x50, 0x3c, 0x26, 0x84, 0x8f, 0xab, 0xea, 0x0d,
	0x79, 0xf2, 0x37, 0x00, 0x00, 0xff, 0xff, 0x17, 0xd3, 0x1c, 0x9f, 0x7a, 0x06, 0x00, 0x00,
}
