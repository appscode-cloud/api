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

type SlaveListResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Slaves []*Slave       `protobuf:"bytes,2,rep,name=slaves" json:"slaves,omitempty"`
}

func (m *SlaveListResponse) Reset()                    { *m = SlaveListResponse{} }
func (m *SlaveListResponse) String() string            { return proto.CompactTextString(m) }
func (*SlaveListResponse) ProtoMessage()               {}
func (*SlaveListResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

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
	Name   string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
}

func (m *Slave) Reset()                    { *m = Slave{} }
func (m *Slave) String() string            { return proto.CompactTextString(m) }
func (*Slave) ProtoMessage()               {}
func (*Slave) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

type SlaveCreateRequest struct {
	Executors         int32       `protobuf:"varint,1,opt,name=executors" json:"executors,omitempty"`
	Labels            string      `protobuf:"bytes,2,opt,name=labels" json:"labels,omitempty"`
	UserStartupScript string      `protobuf:"bytes,3,opt,name=user_startup_script" json:"user_startup_script,omitempty"`
	SaltbaseVersion   string      `protobuf:"bytes,4,opt,name=saltbase_version" json:"saltbase_version,omitempty"`
	CiStarterVersion  string      `protobuf:"bytes,5,opt,name=ci_starter_version" json:"ci_starter_version,omitempty"`
	Ports             []*PortInfo `protobuf:"bytes,6,rep,name=ports" json:"ports,omitempty"`
}

func (m *SlaveCreateRequest) Reset()                    { *m = SlaveCreateRequest{} }
func (m *SlaveCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*SlaveCreateRequest) ProtoMessage()               {}
func (*SlaveCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

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
func (*PortInfo) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

type SlaveDescribeRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *SlaveDescribeRequest) Reset()                    { *m = SlaveDescribeRequest{} }
func (m *SlaveDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*SlaveDescribeRequest) ProtoMessage()               {}
func (*SlaveDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

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
}

func (m *SlaveDescribeResponse) Reset()                    { *m = SlaveDescribeResponse{} }
func (m *SlaveDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*SlaveDescribeResponse) ProtoMessage()               {}
func (*SlaveDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

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
func (*SlaveDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

type SlaveRestartRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *SlaveRestartRequest) Reset()                    { *m = SlaveRestartRequest{} }
func (m *SlaveRestartRequest) String() string            { return proto.CompactTextString(m) }
func (*SlaveRestartRequest) ProtoMessage()               {}
func (*SlaveRestartRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

type SlaveRestartResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *SlaveRestartResponse) Reset()                    { *m = SlaveRestartResponse{} }
func (m *SlaveRestartResponse) String() string            { return proto.CompactTextString(m) }
func (*SlaveRestartResponse) ProtoMessage()               {}
func (*SlaveRestartResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{8} }

func (m *SlaveRestartResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*SlaveListResponse)(nil), "ci.SlaveListResponse")
	proto.RegisterType((*Slave)(nil), "ci.Slave")
	proto.RegisterType((*SlaveCreateRequest)(nil), "ci.SlaveCreateRequest")
	proto.RegisterType((*PortInfo)(nil), "ci.PortInfo")
	proto.RegisterType((*SlaveDescribeRequest)(nil), "ci.SlaveDescribeRequest")
	proto.RegisterType((*SlaveDescribeResponse)(nil), "ci.SlaveDescribeResponse")
	proto.RegisterType((*SlaveDeleteRequest)(nil), "ci.SlaveDeleteRequest")
	proto.RegisterType((*SlaveRestartRequest)(nil), "ci.SlaveRestartRequest")
	proto.RegisterType((*SlaveRestartResponse)(nil), "ci.SlaveRestartResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Slaves service

type SlavesClient interface {
	List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*SlaveListResponse, error)
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

func (c *slavesClient) List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*SlaveListResponse, error) {
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
	List(context.Context, *dtypes.VoidRequest) (*SlaveListResponse, error)
	Describe(context.Context, *SlaveDescribeRequest) (*SlaveDescribeResponse, error)
	Create(context.Context, *SlaveCreateRequest) (*dtypes.LongRunningResponse, error)
	Delete(context.Context, *SlaveDeleteRequest) (*dtypes.LongRunningResponse, error)
	Restart(context.Context, *SlaveRestartRequest) (*SlaveRestartResponse, error)
}

func RegisterSlavesServer(s *grpc.Server, srv SlavesServer) {
	s.RegisterService(&_Slaves_serviceDesc, srv)
}

func _Slaves_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtypes.VoidRequest)
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
		return srv.(SlavesServer).List(ctx, req.(*dtypes.VoidRequest))
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

var fileDescriptor3 = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6e, 0xd3, 0x4c,
	0x14, 0xc5, 0x95, 0x7f, 0x6e, 0x72, 0xd3, 0x2f, 0x5f, 0x33, 0x69, 0xd3, 0xa9, 0x0b, 0xa5, 0x32,
	0x54, 0x42, 0x5d, 0xd8, 0xa5, 0x48, 0x2c, 0xd8, 0xd2, 0x0d, 0xa2, 0x42, 0x28, 0x95, 0x90, 0x58,
	0x20, 0xcb, 0x71, 0x6e, 0xa2, 0x11, 0xc6, 0x63, 0x3c, 0xe3, 0x08, 0x84, 0xd8, 0xf0, 0x0a, 0x3c,
	0x04, 0x0f, 0x84, 0x58, 0xb3, 0xe1, 0x41, 0x98, 0x19, 0x8f, 0xad, 0x26, 0xa4, 0xa8, 0x9b, 0xc4,
	0x9e, 0x7b, 0xcf, 0xcf, 0x73, 0xce, 0x9d, 0x81, 0xbe, 0x48, 0xa2, 0x25, 0xfa, 0x59, 0xce, 0x25,
	0x27, 0xcd, 0x98, 0xb9, 0x77, 0x16, 0x9c, 0x2f, 0x12, 0x0c, 0xa2, 0x8c, 0x05, 0x51, 0x9a, 0x72,
	0x19, 0x49, 0xc6, 0x53, 0x51, 0x76, 0xb8, 0x63, 0xbd, 0x3c, 0x93, 0x9f, 0x32, 0x14, 0x81, 0xf9,
	0x2d, 0xd7, 0xbd, 0x97, 0x30, 0xbc, 0xd2, 0xa0, 0x4b, 0x26, 0xe4, 0x04, 0x45, 0xa6, 0x14, 0x48,
	0x8e, 0xc0, 0x11, 0x4a, 0x5e, 0x08, 0xda, 0x38, 0x6e, 0x3c, 0xec, 0x9f, 0x0f, 0xfc, 0x52, 0xe9,
	0x5f, 0x99, 0x55, 0x72, 0xa0, 0xea, 0x5a, 0x24, 0x68, 0xf3, 0xb8, 0xa5, 0xea, 0x3d, 0x3f, 0x66,
	0xbe, 0xc1, 0x78, 0x27, 0xd0, 0x31, 0x0f, 0x64, 0x1b, 0xda, 0x69, 0xf4, 0x1e, 0x0d, 0xa1, 0x47,
	0x06, 0x35, 0xb1, 0xa9, 0xdf, 0xbd, 0xef, 0x0d, 0x20, 0xa6, 0xef, 0x59, 0x8e, 0x91, 0xc4, 0x09,
	0x7e, 0x28, 0x50, 0x48, 0x32, 0x84, 0x1e, 0x7e, 0xc4, 0xb8, 0x90, 0x3c, 0x2f, 0xbf, 0xdd, 0xd1,
	0xca, 0x24, 0x9a, 0x62, 0x62, 0x95, 0xe4, 0x10, 0x46, 0x85, 0xc0, 0x3c, 0x54, 0xb8, 0x5c, 0x16,
	0x59, 0x28, 0xe2, 0x9c, 0x65, 0x92, 0xb6, 0x4c, 0x91, 0xc2, 0x8e, 0x88, 0x12, 0x39, 0x8d, 0x04,
	0x86, 0x4b, 0xcc, 0x85, 0x0a, 0x80, 0xb6, 0x4d, 0xc5, 0x05, 0x12, 0xb3, 0x52, 0xa4, 0xc4, 0x55,
	0xad, 0x63, 0x91, 0x9d, 0x8c, 0xe7, 0x52, 0x50, 0xc7, 0xb8, 0xd9, 0xd6, 0x6e, 0x5e, 0xa9, 0x85,
	0xe7, 0xe9, 0x9c, 0x7b, 0x67, 0xd0, 0xad, 0x9e, 0xc9, 0x0e, 0x74, 0x4d, 0x6a, 0x31, 0x4f, 0xac,
	0x2f, 0x02, 0xa0, 0xa5, 0x61, 0x1e, 0xa5, 0x0b, 0xb4, 0xde, 0x1e, 0xc0, 0xae, 0xb1, 0x76, 0x81,
	0x7a, 0x6f, 0xd3, 0xda, 0xdc, 0x4a, 0x22, 0xde, 0xaf, 0x06, 0xec, 0xad, 0xb5, 0xdd, 0x32, 0xfd,
	0x8a, 0x53, 0xe6, 0xb1, 0x12, 0x99, 0x4e, 0xa1, 0x45, 0x46, 0xd0, 0x4f, 0xf9, 0x0c, 0x43, 0x4b,
	0x29, 0x03, 0x18, 0xc3, 0x80, 0xcf, 0xe7, 0x09, 0x4b, 0x31, 0x54, 0x91, 0x8b, 0xda, 0xfc, 0x7f,
	0xd0, 0x31, 0xf9, 0x2a, 0xf3, 0xfa, 0xb5, 0xb4, 0xb8, 0x64, 0x33, 0xcc, 0xe9, 0x96, 0x59, 0xe9,
	0x43, 0x4b, 0xbc, 0x2b, 0x68, 0xb7, 0xa2, 0xac, 0x05, 0xdf, 0xab, 0x72, 0x88, 0xcd, 0x24, 0x67,
	0x61, 0x24, 0x29, 0x18, 0x87, 0x9e, 0x1d, 0xf1, 0x05, 0x26, 0x28, 0x6f, 0x48, 0xe1, 0x3e, 0x8c,
	0x4c, 0x8f, 0x32, 0xaf, 0xb1, 0x9b, 0x9b, 0x9e, 0xd8, 0x40, 0xeb, 0xa6, 0xdb, 0x05, 0x75, 0xfe,
	0xb3, 0x05, 0x8e, 0x11, 0x0a, 0xf2, 0x02, 0xda, 0xfa, 0x84, 0x93, 0x51, 0xd5, 0xf2, 0x9a, 0xb3,
	0x99, 0xfd, 0x9a, 0xbb, 0x57, 0x1f, 0xdf, 0xeb, 0xb7, 0xc0, 0xdb, 0xff, 0xfa, 0xe3, 0xf7, 0xb7,
	0xe6, 0x90, 0xfc, 0x1f, 0xc4, 0x2c, 0x58, 0x9e, 0xf9, 0x8f, 0x82, 0xf2, 0xd0, 0x93, 0x10, 0xba,
	0xd5, 0xd0, 0x08, 0xad, 0xb5, 0x6b, 0xe3, 0x76, 0x0f, 0x36, 0x54, 0x2c, 0xf9, 0xc8, 0x90, 0x29,
	0x19, 0xaf, 0x91, 0x83, 0xcf, 0xda, 0xf5, 0x17, 0xf2, 0x06, 0x9c, 0xf2, 0x5e, 0x90, 0x71, 0x0d,
	0x59, 0xb9, 0x28, 0xee, 0x61, 0xe5, 0xe3, 0x92, 0xa7, 0x8b, 0x49, 0x91, 0xa6, 0x4c, 0xfd, 0x55,
	0x78, 0xd7, 0xe0, 0x77, 0xdd, 0xf5, 0x8d, 0x3f, 0x6d, 0x9c, 0x92, 0xb7, 0xe0, 0x94, 0xf3, 0xb8,
	0x86, 0x5e, 0x19, 0xd0, 0xbf, 0xd1, 0x76, 0xe7, 0xa7, 0x37, 0xed, 0x3c, 0x86, 0x2d, 0x3b, 0x25,
	0xb2, 0x5f, 0xf3, 0x57, 0x87, 0xeb, 0xd2, 0xbf, 0x0b, 0x96, 0x7e, 0x62, 0xe8, 0xf7, 0xc8, 0xdd,
	0xcd, 0xf4, 0x20, 0xc7, 0x29, 0xe7, 0x72, 0xea, 0x98, 0x4b, 0xf8, 0xf8, 0x4f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xaa, 0xf2, 0xdd, 0xaa, 0x03, 0x05, 0x00, 0x00,
}
