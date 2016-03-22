// Code generated by protoc-gen-go.
// source: clusters.proto
// DO NOT EDIT!

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

type Cluster struct {
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

func (m *Cluster) Reset()                    { *m = Cluster{} }
func (m *Cluster) String() string            { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()               {}
func (*Cluster) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ClusterDescribeRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ClusterDescribeRequest) Reset()                    { *m = ClusterDescribeRequest{} }
func (m *ClusterDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterDescribeRequest) ProtoMessage()               {}
func (*ClusterDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type ClusterDescribeResponse struct {
	Status  *dtypes.Status                 `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Cluster *Cluster                       `protobuf:"bytes,2,opt,name=cluster" json:"cluster,omitempty"`
	Specs   *ClusterDescribeResponse_Specs `protobuf:"bytes,3,opt,name=specs" json:"specs,omitempty"`
}

func (m *ClusterDescribeResponse) Reset()                    { *m = ClusterDescribeResponse{} }
func (m *ClusterDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterDescribeResponse) ProtoMessage()               {}
func (*ClusterDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ClusterDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ClusterDescribeResponse) GetCluster() *Cluster {
	if m != nil {
		return m.Cluster
	}
	return nil
}

func (m *ClusterDescribeResponse) GetSpecs() *ClusterDescribeResponse_Specs {
	if m != nil {
		return m.Specs
	}
	return nil
}

type ClusterDescribeResponse_Specs struct {
	Containers  int32 `protobuf:"varint,1,opt,name=containers" json:"containers,omitempty"`
	Pods        int32 `protobuf:"varint,2,opt,name=pods" json:"pods,omitempty"`
	Services    int32 `protobuf:"varint,3,opt,name=services" json:"services,omitempty"`
	Rcs         int32 `protobuf:"varint,4,opt,name=rcs" json:"rcs,omitempty"`
	CpuCore     int64 `protobuf:"varint,5,opt,name=cpu_core" json:"cpu_core,omitempty"`
	TotalMemory int64 `protobuf:"varint,6,opt,name=total_memory" json:"total_memory,omitempty"`
}

func (m *ClusterDescribeResponse_Specs) Reset()         { *m = ClusterDescribeResponse_Specs{} }
func (m *ClusterDescribeResponse_Specs) String() string { return proto.CompactTextString(m) }
func (*ClusterDescribeResponse_Specs) ProtoMessage()    {}
func (*ClusterDescribeResponse_Specs) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{2, 0}
}

type ClusterListResponse struct {
	Status   *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Clusters []*Cluster     `protobuf:"bytes,2,rep,name=clusters" json:"clusters,omitempty"`
}

func (m *ClusterListResponse) Reset()                    { *m = ClusterListResponse{} }
func (m *ClusterListResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterListResponse) ProtoMessage()               {}
func (*ClusterListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *ClusterListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ClusterListResponse) GetClusters() []*Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type ClusterCreateRequest struct {
	Name                string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Provider            string            `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	Zone                string            `protobuf:"bytes,3,opt,name=zone" json:"zone,omitempty"`
	CloudCredentialPhid string            `protobuf:"bytes,4,opt,name=cloud_credential_phid" json:"cloud_credential_phid,omitempty"`
	CloudCredential     map[string]string `protobuf:"bytes,5,rep,name=cloud_credential" json:"cloud_credential,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	NodeSet             map[string]int64  `protobuf:"bytes,6,rep,name=node_set" json:"node_set,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	SaltbaseVersion     string            `protobuf:"bytes,7,opt,name=saltbase_version" json:"saltbase_version,omitempty"`
	KubeStarterVersion  string            `protobuf:"bytes,8,opt,name=kube_starter_version" json:"kube_starter_version,omitempty"`
	KubeVersion         string            `protobuf:"bytes,9,opt,name=kube_version" json:"kube_version,omitempty"`
}

func (m *ClusterCreateRequest) Reset()                    { *m = ClusterCreateRequest{} }
func (m *ClusterCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterCreateRequest) ProtoMessage()               {}
func (*ClusterCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

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

type ClusterScaleRequest struct {
	Name        string           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	NodeChanges map[string]int64 `protobuf:"bytes,2,rep,name=node_changes" json:"node_changes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *ClusterScaleRequest) Reset()                    { *m = ClusterScaleRequest{} }
func (m *ClusterScaleRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterScaleRequest) ProtoMessage()               {}
func (*ClusterScaleRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

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
func (*ClusterDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

type ClusterStartupScriptRequest struct {
	Role string `protobuf:"bytes,1,opt,name=role" json:"role,omitempty"`
}

func (m *ClusterStartupScriptRequest) Reset()                    { *m = ClusterStartupScriptRequest{} }
func (m *ClusterStartupScriptRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterStartupScriptRequest) ProtoMessage()               {}
func (*ClusterStartupScriptRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

type ClusterStartupScriptResponse struct {
	Status        *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Configuration string         `protobuf:"bytes,2,opt,name=configuration" json:"configuration,omitempty"`
}

func (m *ClusterStartupScriptResponse) Reset()                    { *m = ClusterStartupScriptResponse{} }
func (m *ClusterStartupScriptResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterStartupScriptResponse) ProtoMessage()               {}
func (*ClusterStartupScriptResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *ClusterStartupScriptResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type ClusterConfigRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ClusterConfigRequest) Reset()                    { *m = ClusterConfigRequest{} }
func (m *ClusterConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterConfigRequest) ProtoMessage()               {}
func (*ClusterConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

type ClusterConfigResponse struct {
	Status        *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Configuration string         `protobuf:"bytes,2,opt,name=configuration" json:"configuration,omitempty"`
}

func (m *ClusterConfigResponse) Reset()                    { *m = ClusterConfigResponse{} }
func (m *ClusterConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*ClusterConfigResponse) ProtoMessage()               {}
func (*ClusterConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

func (m *ClusterConfigResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*Cluster)(nil), "kubernetes.Cluster")
	proto.RegisterType((*ClusterDescribeRequest)(nil), "kubernetes.ClusterDescribeRequest")
	proto.RegisterType((*ClusterDescribeResponse)(nil), "kubernetes.ClusterDescribeResponse")
	proto.RegisterType((*ClusterDescribeResponse_Specs)(nil), "kubernetes.ClusterDescribeResponse.Specs")
	proto.RegisterType((*ClusterListResponse)(nil), "kubernetes.ClusterListResponse")
	proto.RegisterType((*ClusterCreateRequest)(nil), "kubernetes.ClusterCreateRequest")
	proto.RegisterType((*ClusterScaleRequest)(nil), "kubernetes.ClusterScaleRequest")
	proto.RegisterType((*ClusterDeleteRequest)(nil), "kubernetes.ClusterDeleteRequest")
	proto.RegisterType((*ClusterStartupScriptRequest)(nil), "kubernetes.ClusterStartupScriptRequest")
	proto.RegisterType((*ClusterStartupScriptResponse)(nil), "kubernetes.ClusterStartupScriptResponse")
	proto.RegisterType((*ClusterConfigRequest)(nil), "kubernetes.ClusterConfigRequest")
	proto.RegisterType((*ClusterConfigResponse)(nil), "kubernetes.ClusterConfigResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for Clusters service

type ClustersClient interface {
	Describe(ctx context.Context, in *ClusterDescribeRequest, opts ...grpc.CallOption) (*ClusterDescribeResponse, error)
	Create(ctx context.Context, in *ClusterCreateRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error)
	Scale(ctx context.Context, in *ClusterScaleRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error)
	Delete(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error)
	List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*ClusterListResponse, error)
	StartupScript(ctx context.Context, in *ClusterStartupScriptRequest, opts ...grpc.CallOption) (*ClusterStartupScriptResponse, error)
	// TODO(sanjid): rename to ClientConfig
	Config(ctx context.Context, in *ClusterConfigRequest, opts ...grpc.CallOption) (*ClusterConfigResponse, error)
}

type clustersClient struct {
	cc *grpc.ClientConn
}

func NewClustersClient(cc *grpc.ClientConn) ClustersClient {
	return &clustersClient{cc}
}

func (c *clustersClient) Describe(ctx context.Context, in *ClusterDescribeRequest, opts ...grpc.CallOption) (*ClusterDescribeResponse, error) {
	out := new(ClusterDescribeResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Clusters/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Create(ctx context.Context, in *ClusterCreateRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error) {
	out := new(dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Clusters/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Scale(ctx context.Context, in *ClusterScaleRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error) {
	out := new(dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Clusters/Scale", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Delete(ctx context.Context, in *ClusterDeleteRequest, opts ...grpc.CallOption) (*dtypes.LongRunningResponse, error) {
	out := new(dtypes.LongRunningResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Clusters/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*ClusterListResponse, error) {
	out := new(ClusterListResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Clusters/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) StartupScript(ctx context.Context, in *ClusterStartupScriptRequest, opts ...grpc.CallOption) (*ClusterStartupScriptResponse, error) {
	out := new(ClusterStartupScriptResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Clusters/StartupScript", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Config(ctx context.Context, in *ClusterConfigRequest, opts ...grpc.CallOption) (*ClusterConfigResponse, error) {
	out := new(ClusterConfigResponse)
	err := grpc.Invoke(ctx, "/kubernetes.Clusters/Config", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Clusters service

type ClustersServer interface {
	Describe(context.Context, *ClusterDescribeRequest) (*ClusterDescribeResponse, error)
	Create(context.Context, *ClusterCreateRequest) (*dtypes.LongRunningResponse, error)
	Scale(context.Context, *ClusterScaleRequest) (*dtypes.LongRunningResponse, error)
	Delete(context.Context, *ClusterDeleteRequest) (*dtypes.LongRunningResponse, error)
	List(context.Context, *dtypes.VoidRequest) (*ClusterListResponse, error)
	StartupScript(context.Context, *ClusterStartupScriptRequest) (*ClusterStartupScriptResponse, error)
	// TODO(sanjid): rename to ClientConfig
	Config(context.Context, *ClusterConfigRequest) (*ClusterConfigResponse, error)
}

func RegisterClustersServer(s *grpc.Server, srv ClustersServer) {
	s.RegisterService(&_Clusters_serviceDesc, srv)
}

func _Clusters_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClustersServer).Describe(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Clusters_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClustersServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Clusters_Scale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterScaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClustersServer).Scale(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Clusters_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClustersServer).Delete(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Clusters_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClustersServer).List(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Clusters_StartupScript_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterStartupScriptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClustersServer).StartupScript(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Clusters_Config_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ClusterConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(ClustersServer).Config(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Clusters_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kubernetes.Clusters",
	HandlerType: (*ClustersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Describe",
			Handler:    _Clusters_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Clusters_Create_Handler,
		},
		{
			MethodName: "Scale",
			Handler:    _Clusters_Scale_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Clusters_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Clusters_List_Handler,
		},
		{
			MethodName: "StartupScript",
			Handler:    _Clusters_StartupScript_Handler,
		},
		{
			MethodName: "Config",
			Handler:    _Clusters_Config_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor1 = []byte{
	// 922 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x56, 0x92, 0x26, 0xcd, 0x9e, 0xb4, 0x55, 0x35, 0x6d, 0x17, 0xcb, 0xed, 0x42, 0xf1, 0x96,
	0x65, 0xb7, 0xab, 0x26, 0x25, 0x08, 0x54, 0xed, 0x0d, 0x42, 0x5d, 0xb8, 0x5a, 0x2d, 0xd2, 0x46,
	0x70, 0x85, 0x64, 0x39, 0xf6, 0xc1, 0x6b, 0xd5, 0x9d, 0x31, 0x33, 0xe3, 0x48, 0x65, 0xb5, 0x02,
	0x71, 0xcb, 0x25, 0xb7, 0x48, 0x3c, 0x00, 0x6f, 0xc1, 0x2b, 0xf0, 0x0a, 0xf0, 0x1e, 0xcc, 0x1c,
	0xdb, 0x69, 0x13, 0x9c, 0x1f, 0x89, 0x9b, 0xca, 0x73, 0xe6, 0xcc, 0xf9, 0xf9, 0xce, 0xf7, 0x9d,
	0x14, 0x76, 0xc2, 0x34, 0x57, 0x1a, 0xa5, 0xea, 0x67, 0x52, 0x68, 0xc1, 0xe0, 0x2a, 0x1f, 0xa3,
	0xe4, 0xa8, 0x51, 0xb9, 0x47, 0xb1, 0x10, 0x71, 0x8a, 0x83, 0x20, 0x4b, 0x06, 0x01, 0xe7, 0x42,
	0x07, 0x3a, 0x11, 0xbc, 0xf4, 0x74, 0xef, 0x5b, 0x73, 0xa4, 0x6f, 0x32, 0x54, 0x03, 0xfa, 0x5b,
	0xd8, 0xbd, 0x3f, 0x1b, 0xb0, 0x79, 0x59, 0x04, 0x65, 0x5b, 0xb0, 0x91, 0xbd, 0x4e, 0x22, 0xa7,
	0x71, 0xdc, 0x78, 0x7c, 0xcf, 0x9e, 0x78, 0x70, 0x8d, 0x4e, 0x93, 0x4e, 0xbb, 0xd0, 0x35, 0x0f,
	0x26, 0x49, 0x84, 0xd2, 0x69, 0x55, 0xf7, 0x3f, 0x08, 0x8e, 0xce, 0x06, 0x9d, 0x76, 0xa0, 0x23,
	0x31, 0x36, 0x09, 0x9d, 0x36, 0x9d, 0x01, 0x9a, 0x42, 0x39, 0x1d, 0xfa, 0xde, 0x87, 0xad, 0x58,
	0x06, 0xd9, 0xeb, 0x80, 0x07, 0x7e, 0x2e, 0x53, 0x67, 0x93, 0xac, 0xc6, 0xe5, 0x2a, 0x19, 0x57,
	0xb6, 0x6e, 0x65, 0x4b, 0xc2, 0x84, 0xc7, 0x85, 0xed, 0x1e, 0xd9, 0xb6, 0xa1, 0xcd, 0x45, 0x84,
	0xca, 0x01, 0x73, 0x6c, 0x5b, 0x97, 0x50, 0x62, 0xa0, 0x31, 0xf2, 0x03, 0xed, 0xf4, 0xac, 0x8b,
	0xf7, 0x08, 0xee, 0x97, 0x3d, 0x3c, 0x47, 0x15, 0xca, 0x64, 0x8c, 0xaf, 0xf0, 0xfb, 0x1c, 0x95,
	0x9e, 0x36, 0x41, 0x2d, 0x79, 0xbf, 0x34, 0xe1, 0x9d, 0xff, 0x38, 0xaa, 0xcc, 0xa0, 0x84, 0xec,
	0x5d, 0xe8, 0x28, 0x03, 0x59, 0xae, 0xc8, 0xb7, 0x37, 0xdc, 0xe9, 0x17, 0x68, 0xf5, 0x47, 0x64,
	0x65, 0x27, 0xb0, 0x59, 0x82, 0x4f, 0x88, 0xf4, 0x86, 0x7b, 0xfd, 0x5b, 0xf0, 0xfb, 0x15, 0x84,
	0x17, 0xd0, 0x56, 0x19, 0x86, 0x8a, 0x30, 0xea, 0x0d, 0x9f, 0xd4, 0xf8, 0xcc, 0x67, 0xee, 0x8f,
	0xec, 0x03, 0x37, 0x83, 0x36, 0x7d, 0x50, 0x83, 0x82, 0xeb, 0x20, 0xe1, 0x66, 0xce, 0x54, 0x4c,
	0x9b, 0x26, 0x23, 0x22, 0x45, 0x99, 0xdb, 0x76, 0x16, 0x0a, 0xe5, 0x24, 0x09, 0xb1, 0xc8, 0xd3,
	0x66, 0x3d, 0x68, 0x49, 0x93, 0x74, 0xa3, 0xba, 0x0e, 0xb3, 0xdc, 0x0f, 0x85, 0x44, 0x1a, 0x46,
	0xcb, 0x0e, 0x40, 0x1b, 0x3e, 0xa4, 0xfe, 0x35, 0x5e, 0x0b, 0x79, 0x43, 0x63, 0x69, 0x79, 0xdf,
	0xc2, 0x5e, 0x59, 0xd2, 0x8b, 0x44, 0xe9, 0xb5, 0x81, 0xf8, 0xc0, 0x84, 0x2f, 0x59, 0x68, 0xea,
	0x69, 0x2d, 0x40, 0xc2, 0xfb, 0xa3, 0x05, 0xfb, 0xe5, 0xf7, 0x25, 0xcd, 0xab, 0x76, 0x24, 0x33,
	0xbc, 0x6a, 0xce, 0xf0, 0xaa, 0x60, 0xd9, 0x03, 0x38, 0x08, 0x53, 0x91, 0x47, 0xbe, 0x19, 0x7a,
	0x84, 0x5c, 0x27, 0xa6, 0x0b, 0x22, 0x69, 0x41, 0xbb, 0xaf, 0x60, 0x77, 0xfe, 0xda, 0xf4, 0x6c,
	0x8b, 0xfa, 0xa4, 0xa6, 0xa8, 0x99, 0x42, 0x8c, 0xd1, 0x3c, 0xbc, 0x9c, 0xbe, 0xfb, 0x82, 0x6b,
	0x79, 0xc3, 0x3e, 0x83, 0xae, 0x65, 0x9b, 0xaf, 0x50, 0x1b, 0x98, 0x6c, 0xa0, 0xb3, 0x95, 0x81,
	0x5e, 0x9a, 0x07, 0x23, 0xd4, 0x45, 0x00, 0x07, 0x76, 0x55, 0x90, 0xea, 0x71, 0xa0, 0xd0, 0x9f,
	0x18, 0x8c, 0xac, 0x24, 0x0a, 0xc2, 0x1f, 0xc1, 0xbe, 0x8d, 0xe4, 0x1b, 0x74, 0xa5, 0x89, 0x31,
	0xbd, 0xed, 0x56, 0x22, 0xa1, 0xdb, 0xca, 0x4a, 0xe4, 0x77, 0x3f, 0xb5, 0x20, 0xd6, 0x94, 0x69,
	0x06, 0x7e, 0x85, 0x37, 0x25, 0x86, 0x46, 0x21, 0x93, 0x20, 0xcd, 0x4b, 0xa9, 0x3e, 0x6b, 0x5e,
	0x34, 0xdc, 0x3e, 0x6c, 0xcd, 0x54, 0xb5, 0xd8, 0xbf, 0x65, 0xfd, 0xbd, 0xdf, 0x1b, 0x53, 0x32,
	0x8c, 0xc2, 0x20, 0x5d, 0x30, 0xac, 0x2f, 0x61, 0x8b, 0xc0, 0x09, 0x8d, 0x94, 0x63, 0xac, 0xc6,
	0x7f, 0x5e, 0x03, 0xd0, 0xdd, 0x20, 0x84, 0xcf, 0x65, 0xf1, 0x84, 0xaa, 0x71, 0x87, 0xb0, 0x3b,
	0x6f, 0x5b, 0x59, 0xe1, 0xe7, 0x53, 0x3a, 0x3d, 0xc7, 0x14, 0x17, 0xd1, 0xe9, 0x10, 0xf6, 0xa4,
	0xb9, 0xb6, 0xe0, 0x4b, 0xb4, 0x22, 0x31, 0x6b, 0x22, 0xc9, 0x28, 0x4c, 0xd7, 0x7b, 0x0a, 0x87,
	0x55, 0x79, 0x76, 0x04, 0x79, 0x36, 0x32, 0x42, 0xcc, 0xf4, 0x9d, 0x48, 0x52, 0xa4, 0xd5, 0xae,
	0xf8, 0x1a, 0x8e, 0xea, 0x9d, 0xd7, 0x94, 0xc9, 0x01, 0x6c, 0x1b, 0x19, 0x7f, 0x97, 0xc4, 0xb9,
	0xa4, 0x45, 0x5c, 0x0c, 0xc7, 0x3b, 0xb9, 0x55, 0x05, 0xdd, 0xd6, 0x2f, 0xaa, 0x97, 0x70, 0x30,
	0xe7, 0xf5, 0xbf, 0xb2, 0x0e, 0xff, 0xe9, 0x40, 0xb7, 0x0c, 0xa8, 0xd8, 0x8f, 0xd0, 0xad, 0x76,
	0x10, 0xf3, 0x96, 0x2e, 0x28, 0x2a, 0xcd, 0x7d, 0xb8, 0xc6, 0x12, 0xf3, 0x4e, 0x7f, 0xfe, 0xeb,
	0xef, 0x5f, 0x9b, 0x27, 0xcc, 0xa3, 0xdf, 0x9f, 0xdb, 0x07, 0x83, 0xc9, 0xf9, 0xa0, 0x5a, 0x17,
	0x83, 0x37, 0xb6, 0xc5, 0xb7, 0x4c, 0x40, 0xa7, 0x10, 0x10, 0x3b, 0x5e, 0xa5, 0x2d, 0xf7, 0xb0,
	0x6a, 0xf0, 0x85, 0xe0, 0xf1, 0xab, 0x9c, 0x73, 0xf3, 0xeb, 0x30, 0x4d, 0xfa, 0x88, 0x92, 0x1e,
	0xbb, 0x87, 0x4b, 0x92, 0x3e, 0x6b, 0x9c, 0xb2, 0x6b, 0xb3, 0x5b, 0x2d, 0x21, 0xd9, 0x7b, 0x2b,
	0xa8, 0xba, 0x56, 0x3a, 0x6f, 0x55, 0x3a, 0x05, 0x9d, 0x82, 0xa3, 0xb5, 0xfd, 0xcd, 0xd0, 0x77,
	0x79, 0xc2, 0x12, 0xd4, 0xd3, 0x75, 0x40, 0x0d, 0x60, 0xc3, 0xae, 0x71, 0xb6, 0x57, 0x05, 0xfc,
	0x46, 0x24, 0x51, 0x95, 0xa5, 0xae, 0xef, 0xbb, 0x4b, 0xdf, 0x7b, 0x48, 0x99, 0x1e, 0xb0, 0x65,
	0xad, 0xb1, 0xdf, 0x1a, 0xb0, 0x3d, 0x23, 0x06, 0xf6, 0x61, 0x1d, 0x9e, 0x35, 0xda, 0x72, 0x1f,
	0xaf, 0x76, 0x2c, 0x2b, 0xb9, 0xa0, 0x4a, 0x86, 0xec, 0x7c, 0x71, 0x25, 0x67, 0xaa, 0x78, 0x79,
	0xa6, 0xe8, 0xe9, 0xe0, 0x8d, 0x95, 0xed, 0x5b, 0xf6, 0x53, 0xc3, 0xf0, 0x8a, 0xc8, 0x5f, 0xcf,
	0xab, 0xbb, 0x7a, 0x73, 0xdf, 0x5f, 0xe2, 0x51, 0x56, 0xf2, 0x11, 0x55, 0xf2, 0x94, 0x3d, 0x59,
	0x86, 0x7e, 0xa1, 0xb6, 0x72, 0x08, 0xe3, 0x0e, 0xfd, 0x53, 0xf5, 0xf1, 0xbf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x70, 0xd5, 0x41, 0x37, 0xa8, 0x09, 0x00, 0x00,
}
