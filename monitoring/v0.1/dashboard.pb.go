// Code generated by protoc-gen-go.
// source: dashboard.proto
// DO NOT EDIT!

/*
Package monitoring is a generated protocol buffer package.

It is generated from these files:
	dashboard.proto

It has these top-level messages:
	DashboardCreateRequest
	DashboardCreateResponse
*/
package monitoring

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

type DashboardCreateRequest struct {
	Type        string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	ClusterName string `protobuf:"bytes,2,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	ObjectName  string `protobuf:"bytes,3,opt,name=object_name,json=objectName" json:"object_name,omitempty"`
	Namespace   string `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *DashboardCreateRequest) Reset()                    { *m = DashboardCreateRequest{} }
func (m *DashboardCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*DashboardCreateRequest) ProtoMessage()               {}
func (*DashboardCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DashboardCreateResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Url    string         `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
}

func (m *DashboardCreateResponse) Reset()                    { *m = DashboardCreateResponse{} }
func (m *DashboardCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*DashboardCreateResponse) ProtoMessage()               {}
func (*DashboardCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DashboardCreateResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*DashboardCreateRequest)(nil), "monitoring.DashboardCreateRequest")
	proto.RegisterType((*DashboardCreateResponse)(nil), "monitoring.DashboardCreateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Dashboard service

type DashboardClient interface {
	Create(ctx context.Context, in *DashboardCreateRequest, opts ...grpc.CallOption) (*DashboardCreateResponse, error)
}

type dashboardClient struct {
	cc *grpc.ClientConn
}

func NewDashboardClient(cc *grpc.ClientConn) DashboardClient {
	return &dashboardClient{cc}
}

func (c *dashboardClient) Create(ctx context.Context, in *DashboardCreateRequest, opts ...grpc.CallOption) (*DashboardCreateResponse, error) {
	out := new(DashboardCreateResponse)
	err := grpc.Invoke(ctx, "/monitoring.Dashboard/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Dashboard service

type DashboardServer interface {
	Create(context.Context, *DashboardCreateRequest) (*DashboardCreateResponse, error)
}

func RegisterDashboardServer(s *grpc.Server, srv DashboardServer) {
	s.RegisterService(&_Dashboard_serviceDesc, srv)
}

func _Dashboard_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DashboardCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DashboardServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/monitoring.Dashboard/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DashboardServer).Create(ctx, req.(*DashboardCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dashboard_serviceDesc = grpc.ServiceDesc{
	ServiceName: "monitoring.Dashboard",
	HandlerType: (*DashboardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Dashboard_Create_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xc9, 0xee, 0x52, 0xe8, 0x54, 0x54, 0x72, 0x58, 0x4b, 0x59, 0x50, 0x2b, 0xa8, 0x88,
	0xa4, 0x5a, 0x6f, 0x5e, 0xf5, 0xec, 0xa1, 0xfb, 0x00, 0x92, 0xb6, 0xa1, 0x56, 0xba, 0x49, 0x4c,
	0x52, 0xc1, 0xab, 0xe0, 0xd5, 0xcb, 0x3e, 0x9a, 0xaf, 0xe0, 0x83, 0xd8, 0x26, 0x75, 0x7b, 0x50,
	0xf6, 0x12, 0x86, 0xef, 0xff, 0x33, 0xfc, 0x33, 0x03, 0x7b, 0x25, 0xd5, 0x4f, 0xb9, 0xa0, 0xaa,
	0x24, 0x52, 0x09, 0x23, 0x30, 0xac, 0x04, 0xaf, 0x8d, 0x50, 0x35, 0xaf, 0xa2, 0x45, 0x25, 0x44,
	0xd5, 0xb0, 0x84, 0xca, 0x3a, 0xa1, 0x9c, 0x0b, 0x43, 0x4d, 0x2d, 0xb8, 0x76, 0xce, 0x68, 0xde,
	0xe3, 0xd2, 0xbc, 0x49, 0xa6, 0x13, 0xfb, 0x3a, 0x1e, 0x7f, 0x22, 0x98, 0xdf, 0xff, 0x76, 0xbd,
	0x53, 0x8c, 0x1a, 0x96, 0xb1, 0x97, 0x96, 0x69, 0x83, 0x31, 0xcc, 0x7a, 0x67, 0x88, 0x8e, 0xd0,
	0xb9, 0x9f, 0xd9, 0x1a, 0x1f, 0xc3, 0x4e, 0xd1, 0xb4, 0xda, 0x30, 0xf5, 0xc8, 0xe9, 0x8a, 0x85,
	0x13, 0xab, 0x05, 0x03, 0x7b, 0xe8, 0x10, 0x3e, 0x84, 0x40, 0xe4, 0xcf, 0xac, 0x30, 0xce, 0x31,
	0xb5, 0x0e, 0x70, 0xc8, 0x1a, 0x16, 0xe0, 0xf7, 0x8a, 0x96, 0xb4, 0x60, 0xe1, 0xcc, 0xca, 0x23,
	0x88, 0x97, 0x70, 0xf0, 0x27, 0x8f, 0x96, 0xdd, 0x20, 0x0c, 0x9f, 0x82, 0xa7, 0xbb, 0xa9, 0x5a,
	0x6d, 0x23, 0x05, 0xe9, 0x2e, 0x71, 0x03, 0x91, 0xa5, 0xa5, 0xd9, 0xa0, 0xe2, 0x7d, 0x98, 0xb6,
	0xaa, 0x19, 0xb2, 0xf5, 0x65, 0xba, 0x46, 0xe0, 0x6f, 0xba, 0xe2, 0x0f, 0x04, 0x9e, 0x6b, 0x8d,
	0x63, 0x32, 0x6e, 0x90, 0xfc, 0xbf, 0x87, 0xe8, 0x64, 0xab, 0xc7, 0x65, 0x8b, 0xd3, 0xf7, 0xaf,
	0xef, 0xf5, 0xe4, 0x32, 0x3a, 0xeb, 0xf6, 0x2f, 0x75, 0x21, 0x4a, 0x77, 0x88, 0xf1, 0x67, 0xf2,
	0x7a, 0x45, 0xae, 0x93, 0xcd, 0x01, 0x6f, 0xd1, 0x45, 0xee, 0xd9, 0x13, 0xdc, 0xfc, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x2e, 0x6c, 0x58, 0xfb, 0xd7, 0x01, 0x00, 0x00,
}
