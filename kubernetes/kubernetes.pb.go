// Code generated by protoc-gen-go.
// source: kubernetes.proto
// DO NOT EDIT!

/*
Package kubernetes is a generated protocol buffer package.

It is generated from these files:
	kubernetes.proto

It has these top-level messages:
	ClusterVoidRequest
	ClusterListResponse
	ClusterDetails
	ClusterCreateRequest
	ClusterResizeRequest
	ClusterScaleRequest
	ClusterDeleteRequest
	ClusterStartupScriptRequest
	ClusterStartupScriptResponse
	ClusterKubeConfigRequest
	ClusterKubeConfigResponse
*/
package kubernetes

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

type ClusterVoidRequest struct {
}

func (m *ClusterVoidRequest) Reset()                    { *m = ClusterVoidRequest{} }
func (m *ClusterVoidRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterVoidRequest) ProtoMessage()               {}
func (*ClusterVoidRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ClusterListResponse struct {
	Status   *dtypes.Status    `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Clusters []*ClusterDetails `protobuf:"bytes,2,rep,name=clusters" json:"clusters,omitempty"`
}

func (m *ClusterListResponse) Reset()                    { *m = ClusterListResponse{} }
func (m *ClusterListResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterListResponse) ProtoMessage()               {}
func (*ClusterListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ClusterListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ClusterListResponse) GetClusters() []*ClusterDetails {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type ClusterDetails struct {
	Phid        string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Provider    string `protobuf:"bytes,3,opt,name=provider" json:"provider,omitempty"`
	Zone        string `protobuf:"bytes,4,opt,name=zone" json:"zone,omitempty"`
	Region      string `protobuf:"bytes,5,opt,name=region" json:"region,omitempty"`
	Os          string `protobuf:"bytes,6,opt,name=os" json:"os,omitempty"`
	GraphanaUrl string `protobuf:"bytes,7,opt,name=graphana_url" json:"graphana_url,omitempty"`
	KibanaUrl   string `protobuf:"bytes,8,opt,name=kibana_url" json:"kibana_url,omitempty"`
	IcingaUrl   string `protobuf:"bytes,9,opt,name=icinga_url" json:"icinga_url,omitempty"`
	Nodes       int32  `protobuf:"varint,10,opt,name=nodes" json:"nodes,omitempty"`
	CreatedAt   string `protobuf:"bytes,11,opt,name=created_at" json:"created_at,omitempty"`
}

func (m *ClusterDetails) Reset()                    { *m = ClusterDetails{} }
func (m *ClusterDetails) String() string            { return proto.CompactTextString(m) }
func (*ClusterDetails) ProtoMessage()               {}
func (*ClusterDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ClusterCreateRequest struct {
	Name                string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Provider            string            `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	Zone                string            `protobuf:"bytes,3,opt,name=zone" json:"zone,omitempty"`
	CloudCredentialPhid string            `protobuf:"bytes,4,opt,name=cloud_credential_phid" json:"cloud_credential_phid,omitempty"`
	CloudCredential     map[string]string `protobuf:"bytes,5,rep,name=cloud_credential" json:"cloud_credential,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	NodeSet             map[string]int64  `protobuf:"bytes,6,rep,name=node_set" json:"node_set,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	SaltbaseVersion     string            `protobuf:"bytes,7,opt,name=saltbase_version" json:"saltbase_version,omitempty"`
}

func (m *ClusterCreateRequest) Reset()                    { *m = ClusterCreateRequest{} }
func (m *ClusterCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterCreateRequest) ProtoMessage()               {}
func (*ClusterCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ClusterCreateRequest) GetCloudCredential() map[string]string {
	if m != nil {
		return m.CloudCredential
	}
	return nil
}

func (m *ClusterCreateRequest) GetNodeSet() map[string]int64 {
	if m != nil {
		return m.NodeSet
	}
	return nil
}

type ClusterResizeRequest struct {
	Name    string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	NodeSet map[string]int64 `protobuf:"bytes,2,rep,name=node_set" json:"node_set,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *ClusterResizeRequest) Reset()                    { *m = ClusterResizeRequest{} }
func (m *ClusterResizeRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterResizeRequest) ProtoMessage()               {}
func (*ClusterResizeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ClusterResizeRequest) GetNodeSet() map[string]int64 {
	if m != nil {
		return m.NodeSet
	}
	return nil
}

type ClusterScaleRequest struct {
	Name        string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	NodeChanges map[string]int64 `protobuf:"bytes,2,rep,name=node_changes" json:"node_changes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *ClusterScaleRequest) Reset()                    { *m = ClusterScaleRequest{} }
func (m *ClusterScaleRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterScaleRequest) ProtoMessage()               {}
func (*ClusterScaleRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ClusterScaleRequest) GetNodeChanges() map[string]int64 {
	if m != nil {
		return m.NodeChanges
	}
	return nil
}

type ClusterDeleteRequest struct {
	Name              string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ReleaseReservedIp bool   `protobuf:"varint,2,opt,name=release_reserved_ip" json:"release_reserved_ip,omitempty"`
}

func (m *ClusterDeleteRequest) Reset()                    { *m = ClusterDeleteRequest{} }
func (m *ClusterDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterDeleteRequest) ProtoMessage()               {}
func (*ClusterDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type ClusterStartupScriptRequest struct {
	Role string `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
}

func (m *ClusterStartupScriptRequest) Reset()                    { *m = ClusterStartupScriptRequest{} }
func (m *ClusterStartupScriptRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterStartupScriptRequest) ProtoMessage()               {}
func (*ClusterStartupScriptRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type ClusterStartupScriptResponse struct {
	Status        *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Configuration string         `protobuf:"bytes,2,opt,name=configuration" json:"configuration,omitempty"`
}

func (m *ClusterStartupScriptResponse) Reset()                    { *m = ClusterStartupScriptResponse{} }
func (m *ClusterStartupScriptResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterStartupScriptResponse) ProtoMessage()               {}
func (*ClusterStartupScriptResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ClusterStartupScriptResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type ClusterKubeConfigRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ClusterKubeConfigRequest) Reset()                    { *m = ClusterKubeConfigRequest{} }
func (m *ClusterKubeConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterKubeConfigRequest) ProtoMessage()               {}
func (*ClusterKubeConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type ClusterKubeConfigResponse struct {
	Status        *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Configuration string         `protobuf:"bytes,2,opt,name=configuration" json:"configuration,omitempty"`
}

func (m *ClusterKubeConfigResponse) Reset()                    { *m = ClusterKubeConfigResponse{} }
func (m *ClusterKubeConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterKubeConfigResponse) ProtoMessage()               {}
func (*ClusterKubeConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ClusterKubeConfigResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*ClusterVoidRequest)(nil), "kubernetes.ClusterVoidRequest")
	proto.RegisterType((*ClusterListResponse)(nil), "kubernetes.ClusterListResponse")
	proto.RegisterType((*ClusterDetails)(nil), "kubernetes.ClusterDetails")
	proto.RegisterType((*ClusterCreateRequest)(nil), "kubernetes.ClusterCreateRequest")
	proto.RegisterType((*ClusterResizeRequest)(nil), "kubernetes.ClusterResizeRequest")
	proto.RegisterType((*ClusterScaleRequest)(nil), "kubernetes.ClusterScaleRequest")
	proto.RegisterType((*ClusterDeleteRequest)(nil), "kubernetes.ClusterDeleteRequest")
	proto.RegisterType((*ClusterStartupScriptRequest)(nil), "kubernetes.ClusterStartupScriptRequest")
	proto.RegisterType((*ClusterStartupScriptResponse)(nil), "kubernetes.ClusterStartupScriptResponse")
	proto.RegisterType((*ClusterKubeConfigRequest)(nil), "kubernetes.ClusterKubeConfigRequest")
	proto.RegisterType((*ClusterKubeConfigResponse)(nil), "kubernetes.ClusterKubeConfigResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Cluster service

type ClusterClient interface {
	ClusterCreate(ctx context.Context, in *ClusterCreateRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error)
	ClusterResize(ctx context.Context, in *ClusterResizeRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error)
	ClusterScale(ctx context.Context, in *ClusterScaleRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error)
	ClusterDelete(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error)
	ClusterList(ctx context.Context, in *ClusterVoidRequest, opts ...grpc.CallOption) (*ClusterListResponse, error)
	StartupScript(ctx context.Context, in *ClusterStartupScriptRequest, opts ...grpc.CallOption) (*ClusterStartupScriptResponse, error)
	ClusterKubeConfig(ctx context.Context, in *ClusterKubeConfigRequest, opts ...grpc.CallOption) (*ClusterKubeConfigResponse, error)
}

type clusterClient struct {
	cc *grpc.ClientConn
}

func NewClusterClient(cc *grpc.ClientConn) ClusterClient {
	return &clusterClient{cc}
}

func (c *clusterClient) ClusterCreate(ctx context.Context, in *ClusterCreateRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error) {
	out := new(dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Cluster/ClusterCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) ClusterResize(ctx context.Context, in *ClusterResizeRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error) {
	out := new(dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Cluster/ClusterResize", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) ClusterScale(ctx context.Context, in *ClusterScaleRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error) {
	out := new(dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Cluster/ClusterScale", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) ClusterDelete(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error) {
	out := new(dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Cluster/ClusterDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) ClusterList(ctx context.Context, in *ClusterVoidRequest, opts ...grpc.CallOption) (*ClusterListResponse, error) {
	out := new(ClusterListResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Cluster/ClusterList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) StartupScript(ctx context.Context, in *ClusterStartupScriptRequest, opts ...grpc.CallOption) (*ClusterStartupScriptResponse, error) {
	out := new(ClusterStartupScriptResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Cluster/StartupScript", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) ClusterKubeConfig(ctx context.Context, in *ClusterKubeConfigRequest, opts ...grpc.CallOption) (*ClusterKubeConfigResponse, error) {
	out := new(ClusterKubeConfigResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Cluster/ClusterKubeConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cluster service

type ClusterServer interface {
	ClusterCreate(context.Context, *ClusterCreateRequest) (*dtypes.LongRunningResponse, error)
	ClusterResize(context.Context, *ClusterResizeRequest) (*dtypes.LongRunningResponse, error)
	ClusterScale(context.Context, *ClusterScaleRequest) (*dtypes.LongRunningResponse, error)
	ClusterDelete(context.Context, *ClusterDeleteRequest) (*dtypes.LongRunningResponse, error)
	ClusterList(context.Context, *ClusterVoidRequest) (*ClusterListResponse, error)
	StartupScript(context.Context, *ClusterStartupScriptRequest) (*ClusterStartupScriptResponse, error)
	ClusterKubeConfig(context.Context, *ClusterKubeConfigRequest) (*ClusterKubeConfigResponse, error)
}

func RegisterClusterServer(s *grpc.Server, srv ClusterServer) {
	s.RegisterService(&_Cluster_serviceDesc, srv)
}

func _Cluster_ClusterCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClusterServer).ClusterCreate(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Cluster_ClusterResize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterResizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClusterServer).ClusterResize(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Cluster_ClusterScale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterScaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClusterServer).ClusterScale(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Cluster_ClusterDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClusterServer).ClusterDelete(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Cluster_ClusterList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterVoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClusterServer).ClusterList(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Cluster_StartupScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterStartupScriptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClusterServer).StartupScript(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Cluster_ClusterKubeConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterKubeConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClusterServer).ClusterKubeConfig(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Cluster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kubernetes.Cluster",
	HandlerType: (*ClusterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClusterCreate",
			Handler:    _Cluster_ClusterCreate_Handler,
		},
		{
			MethodName: "ClusterResize",
			Handler:    _Cluster_ClusterResize_Handler,
		},
		{
			MethodName: "ClusterScale",
			Handler:    _Cluster_ClusterScale_Handler,
		},
		{
			MethodName: "ClusterDelete",
			Handler:    _Cluster_ClusterDelete_Handler,
		},
		{
			MethodName: "ClusterList",
			Handler:    _Cluster_ClusterList_Handler,
		},
		{
			MethodName: "StartupScript",
			Handler:    _Cluster_StartupScript_Handler,
		},
		{
			MethodName: "ClusterKubeConfig",
			Handler:    _Cluster_ClusterKubeConfig_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 813 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x56, 0x41, 0x4f, 0xdb, 0x48,
	0x14, 0x96, 0x13, 0x08, 0xe1, 0x25, 0x41, 0xd9, 0x01, 0x76, 0x8d, 0xc3, 0x02, 0xb2, 0xd8, 0xdd,
	0x88, 0x5d, 0x12, 0xc4, 0x6a, 0x57, 0x15, 0x97, 0xaa, 0x0a, 0xed, 0xa5, 0xa8, 0x95, 0x12, 0xb5,
	0xd7, 0x68, 0x62, 0x4f, 0xcd, 0x28, 0xae, 0xed, 0xce, 0x8c, 0x23, 0x01, 0xe2, 0xd2, 0x73, 0x4f,
	0xad, 0x7a, 0xee, 0xef, 0xe9, 0xa1, 0xa7, 0xfe, 0x85, 0xfe, 0x87, 0x5e, 0x3b, 0x1e, 0xdb, 0xa9,
	0x4d, 0x9c, 0x86, 0x8a, 0x0b, 0x62, 0xde, 0xbc, 0xf9, 0xbe, 0xf7, 0x7d, 0xf3, 0xde, 0x38, 0xd0,
	0x1c, 0x87, 0x23, 0xc2, 0x3c, 0x22, 0x08, 0xef, 0x04, 0xcc, 0x17, 0x3e, 0x82, 0xef, 0x11, 0x63,
	0xdb, 0xf1, 0x7d, 0xc7, 0x25, 0x5d, 0x1c, 0xd0, 0x2e, 0xf6, 0x3c, 0x5f, 0x60, 0x41, 0x7d, 0x2f,
	0xc9, 0x34, 0x7e, 0x8d, 0xc2, 0xb6, 0xb8, 0x08, 0x08, 0xef, 0xaa, 0xbf, 0x71, 0xdc, 0xdc, 0x00,
	0xd4, 0x73, 0x43, 0x2e, 0x08, 0x7b, 0xee, 0x53, 0xbb, 0x4f, 0x5e, 0x85, 0x84, 0x0b, 0xd3, 0x82,
	0xf5, 0x24, 0x7a, 0x46, 0xb9, 0xe8, 0x13, 0x1e, 0x48, 0x24, 0x82, 0x76, 0xa0, 0xc2, 0x25, 0x6c,
	0xc8, 0x75, 0x6d, 0x4f, 0x6b, 0xd7, 0x8e, 0xd7, 0x3a, 0x31, 0x62, 0x67, 0xa0, 0xa2, 0xe8, 0x1f,
	0xa8, 0x5a, 0xf1, 0x31, 0xae, 0x97, 0xf6, 0xca, 0x32, 0xc3, 0xe8, 0x64, 0x6a, 0x4e, 0x20, 0x4f,
	0x89, 0xc0, 0xd4, 0xe5, 0xe6, 0x47, 0x0d, 0xd6, 0xf2, 0x21, 0x54, 0x87, 0xa5, 0xe0, 0x9c, 0xda,
	0x0a, 0x7e, 0x35, 0x5a, 0x79, 0xf8, 0x25, 0x91, 0x50, 0xd1, 0xaa, 0x09, 0x55, 0x59, 0xf2, 0x84,
	0xda, 0x84, 0xe9, 0xe5, 0x74, 0xff, 0xd2, 0xf7, 0x88, 0xbe, 0xa4, 0x56, 0x6b, 0x50, 0x61, 0xc4,
	0x91, 0x92, 0xf5, 0x65, 0xb5, 0x06, 0x28, 0xf9, 0x5c, 0xaf, 0xa8, 0xff, 0x37, 0xa0, 0xee, 0x30,
	0x1c, 0x9c, 0x63, 0x0f, 0x0f, 0x43, 0xe6, 0xea, 0x2b, 0x2a, 0x2a, 0x53, 0xc6, 0x74, 0x94, 0xc6,
	0xaa, 0x69, 0x8c, 0x5a, 0xd4, 0x73, 0xe2, 0xd8, 0xaa, 0x8a, 0x35, 0x60, 0xd9, 0xf3, 0x6d, 0xc2,
	0x75, 0x90, 0xcb, 0xe5, 0x28, 0xc5, 0x62, 0x04, 0x0b, 0x62, 0x0f, 0xb1, 0xd0, 0x6b, 0x51, 0x8a,
	0xf9, 0xb5, 0x04, 0x1b, 0x89, 0x96, 0x9e, 0xda, 0x4b, 0x9c, 0x9c, 0x6a, 0xd0, 0x66, 0x34, 0x94,
	0x72, 0x1a, 0x62, 0x45, 0xbf, 0xc3, 0xa6, 0xe5, 0xfa, 0xa1, 0x3d, 0x94, 0x04, 0x36, 0xf1, 0x04,
	0xc5, 0xee, 0x50, 0x19, 0x12, 0x4b, 0x7c, 0x0a, 0xcd, 0x9b, 0xdb, 0x52, 0x6c, 0xe4, 0xf3, 0x7f,
	0x05, 0x3e, 0xe7, 0x0a, 0x91, 0x41, 0x79, 0xb0, 0x37, 0x3d, 0xf7, 0xd0, 0x13, 0xec, 0x02, 0xdd,
	0x87, 0x6a, 0xa4, 0x6c, 0xc8, 0x89, 0x90, 0x4e, 0x45, 0x40, 0x87, 0x0b, 0x81, 0x9e, 0xc8, 0x03,
	0x03, 0x22, 0x62, 0x00, 0x1d, 0x9a, 0x1c, 0xbb, 0x62, 0x84, 0x39, 0x19, 0x4e, 0xe4, 0xb5, 0x47,
	0xf6, 0x2b, 0x73, 0x8d, 0xff, 0x23, 0x43, 0x0a, 0x28, 0x6b, 0x50, 0x1e, 0x93, 0x8b, 0xc4, 0x0f,
	0xe9, 0xec, 0x04, 0xbb, 0x61, 0x72, 0xc5, 0x27, 0xa5, 0x7b, 0x9a, 0xd1, 0x81, 0x7a, 0x8e, 0x61,
	0x7e, 0x7e, 0x39, 0xca, 0x37, 0xdf, 0x6b, 0x53, 0xe7, 0x65, 0x9f, 0xd2, 0xcb, 0x39, 0xce, 0x67,
	0x95, 0x96, 0xe6, 0x2a, 0xcd, 0x21, 0xe4, 0x94, 0xfe, 0x74, 0x5d, 0x1f, 0xb4, 0xe9, 0x0c, 0x0d,
	0x2c, 0xec, 0xce, 0x29, 0xeb, 0x11, 0xd4, 0x55, 0x59, 0x96, 0x6c, 0x4d, 0x87, 0xa4, 0x53, 0x73,
	0x54, 0x50, 0x5a, 0x16, 0x44, 0x55, 0xd6, 0x8b, 0x8f, 0xc4, 0xd5, 0x1d, 0x43, 0xf3, 0x66, 0x6c,
	0x61, 0x85, 0x0f, 0xa6, 0xc6, 0x9d, 0x12, 0x97, 0xcc, 0x6b, 0xd9, 0x16, 0xac, 0x33, 0xb9, 0x1d,
	0x5d, 0x30, 0x23, 0x9c, 0xb0, 0x89, 0x6c, 0x7b, 0x1a, 0x28, 0x98, 0xaa, 0xf9, 0x37, 0xb4, 0xd2,
	0xf2, 0x04, 0x66, 0x22, 0x0c, 0x06, 0x16, 0xa3, 0x81, 0xc8, 0x20, 0x31, 0xdf, 0x4d, 0x90, 0xcc,
	0x67, 0xb0, 0x5d, 0x9c, 0x7c, 0xcb, 0xd7, 0x65, 0x13, 0x1a, 0x96, 0xef, 0xbd, 0xa0, 0x4e, 0xc8,
	0xd4, 0xd3, 0x16, 0x37, 0x8d, 0xd9, 0x06, 0x3d, 0x81, 0x7d, 0x2c, 0x4d, 0xeb, 0xa9, 0x8c, 0x42,
	0x29, 0x66, 0x1f, 0xb6, 0x0a, 0x32, 0xef, 0xc4, 0x7e, 0xfc, 0xa9, 0x02, 0x2b, 0x09, 0x28, 0x1a,
	0x43, 0x23, 0x37, 0x31, 0x68, 0x6f, 0xd1, 0x30, 0x19, 0xad, 0x94, 0xe5, 0xcc, 0xf7, 0x9c, 0x7e,
	0xe8, 0x79, 0xf2, 0xe9, 0x49, 0x4b, 0x32, 0x5b, 0xaf, 0x3f, 0x7f, 0x79, 0x57, 0xda, 0x34, 0x9b,
	0xea, 0x4d, 0x4f, 0x5f, 0xd6, 0xee, 0xe4, 0xe8, 0x44, 0x3b, 0xc8, 0x90, 0xc5, 0x4d, 0x5b, 0x48,
	0x96, 0xeb, 0xe7, 0x5b, 0x91, 0x19, 0x85, 0x64, 0x14, 0xea, 0xd9, 0x36, 0x44, 0xbb, 0x0b, 0x1a,
	0xf4, 0x0e, 0x54, 0xfe, 0x54, 0x57, 0xdc, 0x95, 0x85, 0xba, 0x72, 0x0d, 0xfb, 0x63, 0xb2, 0x5d,
	0x45, 0xb6, 0x75, 0xf0, 0xdb, 0x4d, 0xb2, 0xee, 0x55, 0xd4, 0x1a, 0xd7, 0xe8, 0x1c, 0x6a, 0x99,
	0x6f, 0x1d, 0xda, 0x29, 0xa0, 0xcb, 0x7c, 0x1a, 0x8d, 0x22, 0xe9, 0xd9, 0x8f, 0xa4, 0xa9, 0x2b,
	0x42, 0x84, 0x66, 0xd4, 0xa1, 0xb7, 0x1a, 0x34, 0x72, 0xad, 0x8f, 0xfe, 0x2a, 0xf2, 0xb1, 0x60,
	0x92, 0x8c, 0xf6, 0xe2, 0xc4, 0x84, 0xbe, 0xa3, 0xe8, 0xdb, 0xe8, 0xcf, 0x19, 0xbd, 0x3c, 0xce,
	0x3f, 0xe4, 0xea, 0x40, 0xf7, 0x2a, 0x1a, 0xcd, 0x6b, 0xf4, 0x46, 0x83, 0x5f, 0x66, 0xa6, 0x02,
	0xed, 0x17, 0xf0, 0xcd, 0x8c, 0x97, 0xf1, 0xc7, 0x82, 0xac, 0xa4, 0xa4, 0x03, 0x55, 0xd2, 0x3e,
	0x32, 0x67, 0x4a, 0x8a, 0x8e, 0xc7, 0x53, 0x95, 0xdc, 0xc6, 0xa8, 0xa2, 0x7e, 0x96, 0xfc, 0xfb,
	0x2d, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x02, 0x56, 0xd3, 0xec, 0x08, 0x00, 0x00,
}
