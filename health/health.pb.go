// Code generated by protoc-gen-go.
// source: health.proto
// DO NOT EDIT!

/*
Package health is a generated protocol buffer package.

It is generated from these files:
	health.proto

It has these top-level messages:
	StatusResponse
	URLBase
	NetConfig
	Metadata
*/
package health

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"
import appscode_version "github.com/appscode/api/version"

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

type StatusResponse struct {
	Status   *appscode_dtypes.Status   `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Version  *appscode_version.Version `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Metadata *Metadata                 `protobuf:"bytes,3,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StatusResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *StatusResponse) GetVersion() *appscode_version.Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *StatusResponse) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type URLBase struct {
	Scheme   string `protobuf:"bytes,1,opt,name=scheme" json:"scheme,omitempty"`
	BaseAddr string `protobuf:"bytes,2,opt,name=base_addr,json=baseAddr" json:"base_addr,omitempty"`
}

func (m *URLBase) Reset()                    { *m = URLBase{} }
func (m *URLBase) String() string            { return proto.CompactTextString(m) }
func (*URLBase) ProtoMessage()               {}
func (*URLBase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *URLBase) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *URLBase) GetBaseAddr() string {
	if m != nil {
		return m.BaseAddr
	}
	return ""
}

type NetConfig struct {
	TeamAddr         string   `protobuf:"bytes,1,opt,name=team_addr,json=teamAddr" json:"team_addr,omitempty"`
	PublicUrls       *URLBase `protobuf:"bytes,2,opt,name=public_urls,json=publicUrls" json:"public_urls,omitempty"`
	TeamUrls         *URLBase `protobuf:"bytes,3,opt,name=team_urls,json=teamUrls" json:"team_urls,omitempty"`
	ClusterUrls      *URLBase `protobuf:"bytes,4,opt,name=cluster_urls,json=clusterUrls" json:"cluster_urls,omitempty"`
	InClusterUrls    *URLBase `protobuf:"bytes,5,opt,name=in_cluster_urls,json=inClusterUrls" json:"in_cluster_urls,omitempty"`
	URLShortenerUrls *URLBase `protobuf:"bytes,6,opt,name=URL_shortener_urls,json=URLShortenerUrls" json:"URL_shortener_urls,omitempty"`
	FileUrls         *URLBase `protobuf:"bytes,7,opt,name=file_urls,json=fileUrls" json:"file_urls,omitempty"`
}

func (m *NetConfig) Reset()                    { *m = NetConfig{} }
func (m *NetConfig) String() string            { return proto.CompactTextString(m) }
func (*NetConfig) ProtoMessage()               {}
func (*NetConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *NetConfig) GetTeamAddr() string {
	if m != nil {
		return m.TeamAddr
	}
	return ""
}

func (m *NetConfig) GetPublicUrls() *URLBase {
	if m != nil {
		return m.PublicUrls
	}
	return nil
}

func (m *NetConfig) GetTeamUrls() *URLBase {
	if m != nil {
		return m.TeamUrls
	}
	return nil
}

func (m *NetConfig) GetClusterUrls() *URLBase {
	if m != nil {
		return m.ClusterUrls
	}
	return nil
}

func (m *NetConfig) GetInClusterUrls() *URLBase {
	if m != nil {
		return m.InClusterUrls
	}
	return nil
}

func (m *NetConfig) GetURLShortenerUrls() *URLBase {
	if m != nil {
		return m.URLShortenerUrls
	}
	return nil
}

func (m *NetConfig) GetFileUrls() *URLBase {
	if m != nil {
		return m.FileUrls
	}
	return nil
}

type Metadata struct {
	Env       string     `protobuf:"bytes,1,opt,name=env" json:"env,omitempty"`
	TeamId    string     `protobuf:"bytes,2,opt,name=team_id,json=teamId" json:"team_id,omitempty"`
	NetConfig *NetConfig `protobuf:"bytes,3,opt,name=net_config,json=netConfig" json:"net_config,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Metadata) GetEnv() string {
	if m != nil {
		return m.Env
	}
	return ""
}

func (m *Metadata) GetTeamId() string {
	if m != nil {
		return m.TeamId
	}
	return ""
}

func (m *Metadata) GetNetConfig() *NetConfig {
	if m != nil {
		return m.NetConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*StatusResponse)(nil), "appscode.health.StatusResponse")
	proto.RegisterType((*URLBase)(nil), "appscode.health.URLBase")
	proto.RegisterType((*NetConfig)(nil), "appscode.health.NetConfig")
	proto.RegisterType((*Metadata)(nil), "appscode.health.Metadata")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Health service

type HealthClient interface {
	Status(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Status(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/appscode.health.Health/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Health service

type HealthServer interface {
	Status(context.Context, *appscode_dtypes.VoidRequest) (*StatusResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.health.Health/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Status(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.health.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Health_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "health.proto",
}

func init() { proto.RegisterFile("health.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x93, 0xcf, 0x6e, 0x13, 0x3f,
	0x10, 0xc7, 0x95, 0xe4, 0xf7, 0xdb, 0x24, 0x93, 0x42, 0x2b, 0x1f, 0x48, 0x08, 0x11, 0x45, 0x7b,
	0xea, 0x29, 0x8b, 0x5a, 0xf5, 0x50, 0x21, 0x21, 0x48, 0x25, 0x04, 0x52, 0x40, 0x95, 0xab, 0xf4,
	0xc0, 0x65, 0xe5, 0xac, 0xa7, 0x89, 0xd1, 0xc6, 0xde, 0xae, 0xbd, 0x95, 0x10, 0xb7, 0xbe, 0x02,
	0x17, 0x5e, 0x82, 0x23, 0x4f, 0xc2, 0x2b, 0xf0, 0x20, 0x68, 0x6d, 0x6f, 0xca, 0x12, 0xa1, 0xbd,
	0xec, 0x9f, 0xf1, 0xf7, 0xf3, 0x9d, 0xb1, 0xc7, 0x03, 0x7b, 0x6b, 0x64, 0xa9, 0x59, 0x4f, 0xb3,
	0x5c, 0x19, 0x45, 0xf6, 0x59, 0x96, 0xe9, 0x44, 0x71, 0x9c, 0xba, 0xf0, 0x78, 0xb2, 0x52, 0x6a,
	0x95, 0x62, 0xc4, 0x32, 0x11, 0x31, 0x29, 0x95, 0x61, 0x46, 0x28, 0xa9, 0x9d, 0x7c, 0xfc, 0xb4,
	0x92, 0xff, 0x63, 0xfd, 0xb0, 0xb6, 0xce, 0xcd, 0xe7, 0x0c, 0x75, 0x64, 0x9f, 0x5e, 0x10, 0xd6,
	0x04, 0xb7, 0x98, 0x6b, 0xa1, 0x64, 0xf5, 0x76, 0x9a, 0xf0, 0x7b, 0x0b, 0x1e, 0x5e, 0x1a, 0x66,
	0x0a, 0x4d, 0x51, 0x67, 0x4a, 0x6a, 0x24, 0x11, 0x04, 0xda, 0x46, 0x46, 0xad, 0x67, 0xad, 0xa3,
	0xc1, 0xf1, 0x70, 0xba, 0xad, 0xdb, 0x25, 0x99, 0x7a, 0xc0, 0xcb, 0xc8, 0x09, 0x74, 0xbd, 0xe9,
	0xa8, 0x6d, 0x89, 0xc7, 0xf7, 0x44, 0x95, 0xed, 0xca, 0xbd, 0x69, 0xa5, 0x24, 0xa7, 0xd0, 0xdb,
	0xa0, 0x61, 0x9c, 0x19, 0x36, 0xea, 0xfc, 0x4d, 0xf9, 0x63, 0x7b, 0xef, 0x05, 0x74, 0x2b, 0x0d,
	0x5f, 0x42, 0x77, 0x41, 0xe7, 0x33, 0xa6, 0x91, 0x3c, 0x82, 0x40, 0x27, 0x6b, 0xdc, 0xa0, 0xad,
	0xb3, 0x4f, 0xfd, 0x1f, 0x79, 0x02, 0xfd, 0x25, 0xd3, 0x18, 0x33, 0xce, 0x73, 0x5b, 0x50, 0x9f,
	0xf6, 0xca, 0xc0, 0x6b, 0xce, 0xf3, 0xf0, 0x5b, 0x07, 0xfa, 0x1f, 0xd0, 0x9c, 0x2b, 0x79, 0x2d,
	0x56, 0xa5, 0xd4, 0x20, 0xdb, 0x38, 0xa9, 0x73, 0xe9, 0x95, 0x81, 0x52, 0x4a, 0xce, 0x60, 0x90,
	0x15, 0xcb, 0x54, 0x24, 0x71, 0x91, 0xa7, 0xda, 0x6f, 0x6d, 0xb4, 0x53, 0xa4, 0x2f, 0x87, 0x82,
	0x13, 0x2f, 0xf2, 0x54, 0x93, 0x53, 0xef, 0x6b, 0xc1, 0x4e, 0x03, 0x68, 0x33, 0x5a, 0xec, 0x05,
	0xec, 0x25, 0x69, 0xa1, 0x0d, 0xe6, 0x8e, 0xfc, 0xaf, 0x81, 0x1c, 0x78, 0xb5, 0x85, 0x5f, 0xc1,
	0xbe, 0x90, 0x71, 0x8d, 0xff, 0xbf, 0x81, 0x7f, 0x20, 0xe4, 0xf9, 0x1f, 0x0e, 0x6f, 0x80, 0x2c,
	0xe8, 0x3c, 0xd6, 0x6b, 0x95, 0x1b, 0x94, 0x95, 0x49, 0xd0, 0x60, 0x72, 0xb0, 0xa0, 0xf3, 0xcb,
	0x0a, 0xa9, 0x76, 0x7f, 0x2d, 0x52, 0x74, 0x78, 0xb7, 0x69, 0xf7, 0xa5, 0xb4, 0xc4, 0xc2, 0x0c,
	0x7a, 0x55, 0xc3, 0xc9, 0x01, 0x74, 0x50, 0xde, 0xfa, 0x96, 0x94, 0x9f, 0x64, 0x08, 0x5d, 0x7b,
	0xa4, 0x82, 0xfb, 0x9e, 0x06, 0xe5, 0xef, 0x3b, 0x4e, 0xce, 0x00, 0x24, 0x9a, 0x38, 0xb1, 0x1d,
	0xf5, 0x87, 0x3d, 0xde, 0x49, 0xb7, 0xed, 0x39, 0xed, 0xcb, 0xea, 0xf3, 0xf8, 0x0b, 0x04, 0x6f,
	0xed, 0x32, 0xb9, 0x81, 0xc0, 0x5d, 0x6a, 0x32, 0xd9, 0xb9, 0xed, 0x57, 0x4a, 0x70, 0x8a, 0x37,
	0x05, 0x6a, 0x33, 0x3e, 0xdc, 0x31, 0xae, 0x0f, 0x4f, 0x78, 0x74, 0xf7, 0x63, 0xd4, 0xee, 0xb5,
	0xee, 0x7e, 0xfe, 0xfa, 0xda, 0x9e, 0x90, 0x71, 0x14, 0xd7, 0x66, 0xd0, 0x31, 0xd1, 0x27, 0xad,
	0xe4, 0xec, 0x39, 0x0c, 0x13, 0xb5, 0xb9, 0xf7, 0x63, 0x99, 0xf0, 0x9e, 0xb3, 0x81, 0xab, 0xea,
	0xa2, 0x9c, 0xd0, 0x8b, 0xd6, 0xc7, 0xc0, 0x85, 0x97, 0x81, 0x1d, 0xd9, 0x93, 0xdf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xa5, 0x39, 0x8d, 0x4f, 0x56, 0x04, 0x00, 0x00,
}
