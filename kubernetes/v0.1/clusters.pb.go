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
	GraphanaUrl string `protobuf:"bytes,7,opt,name=graphana_url,json=graphanaUrl" json:"graphana_url,omitempty"`
	KibanaUrl   string `protobuf:"bytes,8,opt,name=kibana_url,json=kibanaUrl" json:"kibana_url,omitempty"`
	IcingaUrl   string `protobuf:"bytes,9,opt,name=icinga_url,json=icingaUrl" json:"icinga_url,omitempty"`
	Nodes       int32  `protobuf:"varint,10,opt,name=nodes" json:"nodes,omitempty"`
	CreatedAt   string `protobuf:"bytes,11,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
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
	CpuCore     int64 `protobuf:"varint,5,opt,name=cpu_core,json=cpuCore" json:"cpu_core,omitempty"`
	TotalMemory int64 `protobuf:"varint,6,opt,name=total_memory,json=totalMemory" json:"total_memory,omitempty"`
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
	CloudCredential     string            `protobuf:"bytes,4,opt,name=cloud_credential,json=cloudCredential" json:"cloud_credential,omitempty"`
	CloudCredentialData map[string]string `protobuf:"bytes,5,rep,name=cloud_credential_data,json=cloudCredentialData" json:"cloud_credential_data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	NodeSet             map[string]int64  `protobuf:"bytes,6,rep,name=node_set,json=nodeSet" json:"node_set,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	SaltbaseVersion     string            `protobuf:"bytes,7,opt,name=saltbase_version,json=saltbaseVersion" json:"saltbase_version,omitempty"`
	KubeStarterVersion  string            `protobuf:"bytes,8,opt,name=kube_starter_version,json=kubeStarterVersion" json:"kube_starter_version,omitempty"`
	KubeVersion         string            `protobuf:"bytes,9,opt,name=kube_version,json=kubeVersion" json:"kube_version,omitempty"`
}

func (m *ClusterCreateRequest) Reset()                    { *m = ClusterCreateRequest{} }
func (m *ClusterCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterCreateRequest) ProtoMessage()               {}
func (*ClusterCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ClusterCreateRequest) GetCloudCredentialData() map[string]string {
	if m != nil {
		return m.CloudCredentialData
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
	NodeChanges map[string]int64 `protobuf:"bytes,2,rep,name=node_changes,json=nodeChanges" json:"node_changes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
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
	ReleaseReservedIp bool   `protobuf:"varint,2,opt,name=release_reserved_ip,json=releaseReservedIp" json:"release_reserved_ip,omitempty"`
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

type ClusterContainerRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	DiskName string `protobuf:"bytes,2,opt,name=disk_name,json=diskName" json:"disk_name,omitempty"`
}

func (m *ClusterContainerRequest) Reset()                    { *m = ClusterContainerRequest{} }
func (m *ClusterContainerRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterContainerRequest) ProtoMessage()               {}
func (*ClusterContainerRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

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
	proto.RegisterType((*ClusterContainerRequest)(nil), "kubernetes.ClusterContainerRequest")
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
	// 1079 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x56, 0x4f, 0x73, 0xdb, 0x44,
	0x14, 0x1f, 0xd9, 0xb1, 0xe3, 0x3c, 0xa7, 0x69, 0x58, 0xa7, 0x41, 0x28, 0x2d, 0x6d, 0x45, 0x29,
	0x69, 0xa8, 0xed, 0xd4, 0xe5, 0x00, 0x39, 0xc0, 0x30, 0x0e, 0x0c, 0x30, 0xa5, 0x07, 0x79, 0xe8,
	0x81, 0x8b, 0x46, 0x96, 0x16, 0x57, 0x13, 0x65, 0xa5, 0xee, 0xae, 0x3c, 0x13, 0x3a, 0x3d, 0xd0,
	0xaf, 0xc0, 0x1d, 0x8e, 0x7c, 0x0b, 0x0e, 0x7c, 0x05, 0x8e, 0x5c, 0xf9, 0x08, 0x7c, 0x00, 0xf6,
	0x9f, 0x64, 0x3b, 0x75, 0xe4, 0xe4, 0xd0, 0x8b, 0x66, 0xf7, 0xf7, 0xde, 0xdb, 0xf7, 0xff, 0x3d,
	0xc1, 0x56, 0x98, 0xe4, 0x8c, 0x63, 0xca, 0x7a, 0x19, 0x4d, 0x79, 0x8a, 0xe0, 0x24, 0x1f, 0x63,
	0x4a, 0x30, 0xc7, 0xcc, 0xb9, 0x39, 0x49, 0xd3, 0x49, 0x82, 0xfb, 0x41, 0x16, 0xf7, 0x03, 0x42,
	0x52, 0x1e, 0xf0, 0x38, 0x25, 0x86, 0xd3, 0xd9, 0x95, 0x70, 0xc4, 0xcf, 0x32, 0xcc, 0xfa, 0xea,
	0xab, 0x71, 0xf7, 0xf7, 0x1a, 0xac, 0x0f, 0xf5, 0xa3, 0x08, 0xc1, 0x5a, 0xf6, 0x3c, 0x8e, 0x6c,
	0xeb, 0x8e, 0xb5, 0xbf, 0xe1, 0xa9, 0xb3, 0xc4, 0x48, 0x70, 0x8a, 0xed, 0x9a, 0xc6, 0xe4, 0x19,
	0x39, 0xd0, 0x12, 0xc2, 0xd3, 0x38, 0xc2, 0xd4, 0xae, 0x2b, 0xbc, 0xbc, 0x4b, 0xfe, 0x9f, 0x53,
	0x82, 0xed, 0x35, 0xcd, 0x2f, 0xcf, 0x68, 0x17, 0x9a, 0x14, 0x4f, 0x84, 0x31, 0x76, 0x43, 0xa1,
	0xe6, 0x86, 0xb6, 0xa0, 0x96, 0x32, 0xbb, 0xa9, 0x30, 0x71, 0x42, 0x77, 0x61, 0x73, 0x42, 0x83,
	0xec, 0x79, 0x40, 0x02, 0x3f, 0xa7, 0x89, 0xbd, 0xae, 0x28, 0xed, 0x02, 0xfb, 0x81, 0x26, 0xe8,
	0x16, 0xc0, 0x49, 0x3c, 0x2e, 0x18, 0x5a, 0x8a, 0x61, 0x43, 0x23, 0x86, 0x1c, 0x87, 0x31, 0x99,
	0x68, 0xf2, 0x86, 0x26, 0x6b, 0x44, 0x92, 0x77, 0xa0, 0x41, 0xd2, 0x08, 0x33, 0x1b, 0x04, 0xa5,
	0xe1, 0xe9, 0x8b, 0x14, 0x0a, 0x29, 0x0e, 0x38, 0x8e, 0xfc, 0x80, 0xdb, 0x6d, 0x2d, 0x64, 0x90,
	0x2f, 0xb9, 0xfb, 0x10, 0x76, 0x4d, 0x80, 0x8e, 0x31, 0x0b, 0x69, 0x3c, 0xc6, 0x1e, 0x7e, 0x91,
	0x63, 0xc6, 0xcb, 0xd8, 0x58, 0xb3, 0xd8, 0xb8, 0xff, 0xd4, 0xe0, 0xdd, 0x37, 0xd8, 0x59, 0x26,
	0x12, 0x81, 0xd1, 0x7d, 0x68, 0x32, 0x91, 0x95, 0x9c, 0x29, 0x89, 0xf6, 0x60, 0xab, 0xa7, 0x13,
	0xd2, 0x1b, 0x29, 0xd4, 0x33, 0x54, 0xd4, 0x85, 0x75, 0x93, 0x67, 0x15, 0xf6, 0xf6, 0xa0, 0xd3,
	0x9b, 0xe5, 0xb9, 0x67, 0x5e, 0xf7, 0x0a, 0x1e, 0xf4, 0x05, 0x34, 0x58, 0x86, 0x43, 0xa6, 0x72,
	0xd1, 0x1e, 0x3c, 0x58, 0xc2, 0x7c, 0xde, 0x94, 0xde, 0x48, 0x0a, 0x78, 0x5a, 0xce, 0xf9, 0xc3,
	0x82, 0x86, 0x02, 0xd0, 0xfb, 0x22, 0x14, 0x29, 0xe1, 0x41, 0x4c, 0x44, 0x8d, 0x29, 0x2b, 0x1b,
	0xde, 0x1c, 0xa2, 0x2a, 0x24, 0x8d, 0x98, 0x32, 0xab, 0xe1, 0xa9, 0xb3, 0xac, 0x06, 0x86, 0xe9,
	0x34, 0x0e, 0xb1, 0xb6, 0xa0, 0xe1, 0x95, 0x77, 0xb4, 0x0d, 0x75, 0x2a, 0x0c, 0x5b, 0x53, 0xb0,
	0x3c, 0xa2, 0xf7, 0xa0, 0x15, 0x66, 0xb9, 0x1f, 0xa6, 0x14, 0xab, 0x6a, 0xa8, 0x0b, 0x3f, 0xb2,
	0x7c, 0x28, 0xae, 0x32, 0xfd, 0x5c, 0x54, 0x6d, 0xe2, 0x9f, 0xe2, 0xd3, 0x94, 0x9e, 0xa9, 0xc2,
	0xa8, 0x7b, 0x6d, 0x85, 0x7d, 0xaf, 0x20, 0x97, 0x40, 0xc7, 0x78, 0xf4, 0x24, 0x66, 0xfc, 0xca,
	0x81, 0xed, 0x0b, 0xe5, 0xa6, 0x81, 0x84, 0x0b, 0xf5, 0x8b, 0x22, 0x5b, 0x32, 0xb9, 0x7f, 0xad,
	0xc1, 0x8e, 0x41, 0x87, 0xaa, 0x20, 0x2a, 0x52, 0xbf, 0xd0, 0x16, 0xb5, 0x0b, 0xda, 0xa2, 0x3e,
	0xd7, 0x16, 0x0f, 0x60, 0x3b, 0x4c, 0xd2, 0x3c, 0xf2, 0x45, 0xad, 0x45, 0x98, 0xf0, 0x38, 0x48,
	0x4c, 0xdb, 0x5c, 0x57, 0xf8, 0xb0, 0x84, 0xd1, 0x29, 0xdc, 0x38, 0xcf, 0xea, 0x47, 0x01, 0x0f,
	0x44, 0x08, 0xa5, 0x17, 0x9f, 0x2d, 0xf1, 0x62, 0xc1, 0x5e, 0x01, 0x2e, 0x3c, 0x78, 0x2c, 0x64,
	0xbf, 0x22, 0x9c, 0x9e, 0x79, 0x9d, 0xf0, 0x4d, 0x0a, 0xfa, 0x06, 0x5a, 0xb2, 0x35, 0x7c, 0x86,
	0xb9, 0xc8, 0x82, 0xd4, 0xd0, 0x5d, 0xa9, 0xe1, 0xa9, 0x10, 0x18, 0x61, 0xae, 0x5f, 0x5d, 0x27,
	0xfa, 0x26, 0x7d, 0x64, 0x41, 0xc2, 0xc7, 0x01, 0xc3, 0xfe, 0x54, 0x44, 0x54, 0x0e, 0x01, 0xdd,
	0xd6, 0xd7, 0x0b, 0xfc, 0x99, 0x86, 0xd1, 0x21, 0xec, 0x48, 0x1d, 0xbe, 0xc8, 0x15, 0x15, 0xaf,
	0x97, 0xec, 0xba, 0xc9, 0x91, 0xa4, 0x8d, 0x34, 0xa9, 0x90, 0x10, 0x05, 0xa3, 0x24, 0x0a, 0x4e,
	0xdd, 0xef, 0x6d, 0x89, 0x19, 0x16, 0xe7, 0x6b, 0xb0, 0x2f, 0x72, 0x5d, 0x16, 0xe7, 0x09, 0x3e,
	0x33, 0x29, 0x94, 0x47, 0x39, 0x1f, 0xa6, 0x41, 0x92, 0x17, 0xd3, 0x4e, 0x5f, 0x8e, 0x6a, 0x9f,
	0x5a, 0xce, 0x11, 0x6c, 0xce, 0x3b, 0xb8, 0x4a, 0xb6, 0x3e, 0x27, 0xeb, 0xfe, 0x69, 0x95, 0x55,
	0x3b, 0x0a, 0x83, 0xa4, 0xb2, 0x86, 0x46, 0xb0, 0xa9, 0x22, 0x1f, 0x8a, 0x81, 0x37, 0xc1, 0x45,
	0x95, 0x1e, 0x2e, 0x89, 0xfe, 0xfc, 0x53, 0x2a, 0xf8, 0x43, 0x2d, 0xa2, 0x13, 0xd0, 0x26, 0x33,
	0xc4, 0xf9, 0x1c, 0xb6, 0xcf, 0x33, 0x5c, 0xc9, 0x81, 0x1f, 0xcb, 0x26, 0x38, 0xc6, 0x09, 0xae,
	0x6e, 0x82, 0x1e, 0x74, 0xa8, 0x60, 0x92, 0xf9, 0xa6, 0x58, 0xce, 0x01, 0x31, 0x55, 0xe3, 0x4c,
	0xbd, 0xd9, 0xf2, 0xde, 0x31, 0x24, 0xcf, 0x50, 0xbe, 0xcd, 0xdc, 0x47, 0xb0, 0x57, 0x38, 0x24,
	0x93, 0x9b, 0x67, 0x23, 0x31, 0xa8, 0x32, 0x3e, 0xa7, 0x82, 0xa6, 0x49, 0xa9, 0x42, 0x9e, 0xdd,
	0x04, 0x6e, 0x2e, 0x17, 0xb9, 0xe2, 0x34, 0xb8, 0x07, 0xd7, 0xc4, 0x68, 0xfb, 0x29, 0x9e, 0xe4,
	0x54, 0xad, 0x4a, 0x93, 0xf5, 0x45, 0xd0, 0x3d, 0x98, 0x4d, 0x00, 0x85, 0x57, 0x0d, 0x7f, 0x0c,
	0x37, 0xce, 0xf1, 0xbe, 0x15, 0x93, 0xbe, 0x2b, 0x57, 0xcc, 0xb0, 0x18, 0xcd, 0x55, 0x29, 0xd9,
	0x83, 0x8d, 0x28, 0x66, 0x27, 0xfe, 0xdc, 0x1e, 0x6f, 0x49, 0xe0, 0xa9, 0xb8, 0x0f, 0xfe, 0x6b,
	0x42, 0xcb, 0x3c, 0xc6, 0xd0, 0x2f, 0x16, 0xb4, 0x8a, 0x55, 0x81, 0xdc, 0xca, 0x3d, 0xa2, 0xd4,
	0x39, 0x1f, 0x5c, 0x62, 0xd7, 0xb8, 0x0f, 0x5f, 0xff, 0xfd, 0xef, 0xaf, 0xb5, 0xfb, 0xe8, 0x9e,
	0xfa, 0x35, 0x99, 0x09, 0xf4, 0xa7, 0x87, 0xbd, 0x47, 0xfd, 0x62, 0xd6, 0xf6, 0x5f, 0x4a, 0xfb,
	0x5e, 0xa1, 0x17, 0xd0, 0xd4, 0x83, 0x05, 0xdd, 0x59, 0x35, 0x73, 0x9c, 0xbd, 0x22, 0x8c, 0x4f,
	0x52, 0x32, 0xf1, 0x72, 0x42, 0xc4, 0xa2, 0x2f, 0xd5, 0xee, 0x2b, 0xb5, 0xae, 0x73, 0xab, 0x52,
	0xed, 0x91, 0x75, 0x80, 0x52, 0xb1, 0xfe, 0x64, 0x37, 0xa1, 0xdb, 0x2b, 0xfa, 0xec, 0x52, 0x0a,
	0xdd, 0xd5, 0x0a, 0x73, 0x68, 0xea, 0x4e, 0x5a, 0xea, 0xe3, 0x42, 0x93, 0x55, 0xab, 0x34, 0xa1,
	0x3d, 0xb8, 0x5c, 0x68, 0x43, 0x58, 0x93, 0x6b, 0x13, 0x75, 0x8a, 0x27, 0x9f, 0xa5, 0x71, 0x54,
	0xe8, 0x59, 0xe6, 0xfb, 0xfc, 0x92, 0x75, 0x3f, 0x54, 0xba, 0x6e, 0xa3, 0x6a, 0xf7, 0xd0, 0x6f,
	0x16, 0x5c, 0x5b, 0xe8, 0x4b, 0xf4, 0xd1, 0xb2, 0xa8, 0x2e, 0x69, 0x76, 0x67, 0x7f, 0x35, 0xa3,
	0xb1, 0xe5, 0x48, 0xd9, 0xf2, 0x09, 0x1a, 0x54, 0xd9, 0xd2, 0x65, 0x5a, 0xb6, 0xcb, 0x94, 0x70,
	0xff, 0xa5, 0x9c, 0x1e, 0xaf, 0xd0, 0x6b, 0x4b, 0x54, 0x98, 0xea, 0xa7, 0xe5, 0x15, 0x36, 0xdf,
	0xe5, 0xce, 0xdd, 0x0a, 0x0e, 0x63, 0xcb, 0x63, 0x65, 0x4b, 0x17, 0x7d, 0x5c, 0x9d, 0x03, 0xdd,
	0xc2, 0x26, 0x15, 0xe3, 0xa6, 0xfa, 0xfb, 0x7e, 0xfc, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6d,
	0xf3, 0x33, 0x0e, 0xd1, 0x0b, 0x00, 0x00,
}
