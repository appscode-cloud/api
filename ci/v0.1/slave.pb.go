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
	Protocol string `protobuf:"bytes,1,opt,name=protocol" json:"protocol,omitempty"`
	Port     int32  `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
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
const _ = grpc.SupportPackageIsVersion1

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

func _Slaves_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SlavesServer).List(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Slaves_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SlaveDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SlavesServer).Describe(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Slaves_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SlaveCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SlavesServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Slaves_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SlaveDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SlavesServer).Delete(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Slaves_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(SlaveRestartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(SlavesServer).Restart(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
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
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x56, 0x92, 0xda, 0x4d, 0x26, 0xa5, 0x6a, 0x37, 0xb4, 0x6c, 0xdd, 0xaa, 0x2a, 0x2e, 0x48,
	0xd0, 0x83, 0x0d, 0x45, 0xe2, 0xc0, 0x95, 0x5e, 0x90, 0x2a, 0x40, 0xad, 0xc4, 0x0d, 0x59, 0x8e,
	0x33, 0x89, 0x56, 0x18, 0xaf, 0xf1, 0xae, 0x23, 0x10, 0xe2, 0xc2, 0x2b, 0xf0, 0x10, 0x3c, 0x10,
	0x0f, 0xc0, 0x85, 0x0b, 0x6f, 0xc1, 0xee, 0xf8, 0x47, 0x4d, 0x48, 0xab, 0x5e, 0x12, 0x7b, 0x66,
	0xbe, 0x6f, 0xfc, 0x7d, 0xdf, 0x2e, 0x0c, 0x55, 0x1a, 0xcf, 0x31, 0xc8, 0x0b, 0xa9, 0x25, 0xeb,
	0x26, 0xc2, 0x3b, 0x98, 0x49, 0x39, 0x4b, 0x31, 0x8c, 0x73, 0x11, 0xc6, 0x59, 0x26, 0x75, 0xac,
	0x85, 0xcc, 0x54, 0x35, 0xe1, 0xed, 0xda, 0xf2, 0x44, 0x7f, 0xc9, 0x51, 0x85, 0xf4, 0x5b, 0xd5,
	0xfd, 0xd7, 0xb0, 0x7d, 0x69, 0x89, 0xce, 0x85, 0xd2, 0x17, 0xa8, 0x72, 0x83, 0x40, 0x76, 0x08,
	0xae, 0x32, 0xf0, 0x52, 0xf1, 0xce, 0x51, 0xe7, 0xd1, 0xf0, 0x74, 0x33, 0xa8, 0x90, 0xc1, 0x25,
	0x55, 0xd9, 0x9e, 0xe9, 0x5b, 0x90, 0xe2, 0xdd, 0xa3, 0x9e, 0xe9, 0x0f, 0x82, 0x44, 0x04, 0x44,
	0xe3, 0x3f, 0x04, 0x87, 0x1e, 0xd8, 0x06, 0xac, 0x65, 0xf1, 0x47, 0x24, 0x86, 0x01, 0xdb, 0x6c,
	0x19, 0xbb, 0xf6, 0xdd, 0xff, 0xd9, 0x01, 0x46, 0x73, 0x2f, 0x0b, 0x8c, 0x35, 0x5e, 0xe0, 0xa7,
	0x12, 0x95, 0x66, 0xdb, 0x30, 0xc0, 0xcf, 0x98, 0x94, 0x5a, 0x16, 0xd5, 0x6e, 0xc7, 0x22, 0xd3,
	0x78, 0x8c, 0x69, 0x8d, 0x64, 0xfb, 0x30, 0x2a, 0x15, 0x16, 0x91, 0xa1, 0x2b, 0x74, 0x99, 0x47,
	0x2a, 0x29, 0x44, 0xae, 0x79, 0x8f, 0x9a, 0x1c, 0xb6, 0x54, 0x9c, 0xea, 0x71, 0xac, 0x30, 0x9a,
	0x63, 0xa1, 0x8c, 0x01, 0x7c, 0x8d, 0x3a, 0x1e, 0xb0, 0x44, 0x54, 0x20, 0x03, 0x6e, 0x7a, 0x4e,
	0x4d, 0xe9, 0xe4, 0xb2, 0xd0, 0x8a, 0xbb, 0xa4, 0x66, 0xc3, 0xaa, 0x79, 0x6b, 0x0a, 0xaf, 0xb2,
	0xa9, 0xf4, 0x4f, 0xa0, 0xdf, 0x3c, 0xb3, 0x2d, 0xe8, 0x93, 0x6b, 0x89, 0x4c, 0x6b, 0x5d, 0x46,
	0xa5, 0x85, 0xd2, 0xb7, 0x39, 0xfe, 0x03, 0xb8, 0x4b, 0xa2, 0xce, 0xd0, 0x7e, 0xd5, 0xb8, 0x95,
	0xb5, 0xe0, 0x85, 0xff, 0xbb, 0x03, 0x3b, 0x4b, 0x63, 0xb7, 0xf4, 0xbd, 0xe1, 0xa9, 0x9c, 0x58,
	0x30, 0xcb, 0xea, 0xef, 0xb1, 0x11, 0x0c, 0x33, 0x39, 0xc1, 0xa8, 0x66, 0xa9, 0xa4, 0xef, 0xc2,
	0xa6, 0x9c, 0x4e, 0x53, 0x91, 0x61, 0x64, 0xcc, 0x56, 0xad, 0xec, 0x3b, 0xe0, 0x90, 0xb3, 0x46,
	0xb6, 0x7d, 0xad, 0xc4, 0xcd, 0xc5, 0x04, 0x0b, 0xbe, 0x4e, 0x95, 0x21, 0xf4, 0xd4, 0x87, 0x92,
	0xf7, 0x1b, 0x96, 0x25, 0xcb, 0x07, 0x54, 0x67, 0x00, 0x09, 0x65, 0x38, 0x89, 0x62, 0xcd, 0x81,
	0x14, 0xfa, 0x75, 0xb8, 0x67, 0x98, 0xa2, 0xbe, 0xc6, 0x85, 0x63, 0x18, 0xd1, 0x8c, 0x11, 0x6f,
	0x69, 0x57, 0x0f, 0x3d, 0xaf, 0x0d, 0x6d, 0x87, 0x6e, 0x67, 0xd4, 0xe9, 0xdf, 0x1e, 0xb8, 0x04,
	0x54, 0xec, 0x0d, 0xac, 0xd9, 0xb3, 0xcd, 0x46, 0xcd, 0xc8, 0x3b, 0x29, 0x26, 0xf5, 0x36, 0x6f,
	0xa7, 0x3d, 0xb8, 0x57, 0xcf, 0xbf, 0xbf, 0xff, 0xfd, 0xd7, 0x9f, 0x1f, 0xdd, 0x1d, 0x36, 0xa2,
	0xcb, 0x94, 0x88, 0x70, 0xfe, 0x24, 0x78, 0x1a, 0x56, 0x47, 0x9e, 0x25, 0xd0, 0x6f, 0x82, 0x63,
	0xbc, 0xc5, 0x2f, 0x45, 0xee, 0xed, 0xad, 0xe8, 0xd4, 0xec, 0x3e, 0xb1, 0x1f, 0x30, 0x6f, 0x05,
	0x7b, 0xf8, 0xd5, 0xaa, 0xff, 0xc6, 0xde, 0x83, 0x5b, 0xdd, 0x0c, 0xb6, 0xdb, 0x12, 0x2d, 0x5c,
	0x15, 0x6f, 0xbf, 0xd1, 0x73, 0x2e, 0xb3, 0xd9, 0x45, 0x99, 0x65, 0xc2, 0xfc, 0x35, 0x2b, 0x0e,
	0x69, 0x05, 0xf7, 0x56, 0x09, 0x78, 0xd1, 0x39, 0x61, 0x31, 0xb8, 0x55, 0x36, 0x57, 0xe8, 0x17,
	0xc2, 0xba, 0x99, 0xbe, 0x56, 0x70, 0x72, 0x93, 0x82, 0x19, 0xac, 0xd7, 0xa9, 0xb1, 0x7b, 0xed,
	0x8e, 0xc5, 0xb0, 0x3d, 0xfe, 0x7f, 0xa3, 0xde, 0xf0, 0x98, 0x36, 0x1c, 0xb3, 0xfb, 0xd7, 0x6f,
	0x08, 0x0b, 0x1c, 0x4b, 0xa9, 0xc7, 0x2e, 0x5d, 0xc9, 0x67, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xf6, 0x58, 0xc6, 0xbb, 0x11, 0x05, 0x00, 0x00,
}
