// Code generated by protoc-gen-go.
// source: loadbalancer.proto
// DO NOT EDIT!

package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ListRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *ListRequest) Reset()                    { *m = ListRequest{} }
func (m *ListRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()               {}
func (*ListRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *ListRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

type ListResponse struct {
	Status        *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	LoadBalancers []*LoadBalancer         `protobuf:"bytes,2,rep,name=load_balancers,json=loadBalancers" json:"load_balancers,omitempty"`
}

func (m *ListResponse) Reset()                    { *m = ListResponse{} }
func (m *ListResponse) String() string            { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()               {}
func (*ListResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *ListResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListResponse) GetLoadBalancers() []*LoadBalancer {
	if m != nil {
		return m.LoadBalancers
	}
	return nil
}

type DescribeRequest struct {
	Kind      string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	Cluster   string `protobuf:"bytes,4,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *DescribeRequest) Reset()                    { *m = DescribeRequest{} }
func (m *DescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeRequest) ProtoMessage()               {}
func (*DescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *DescribeRequest) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *DescribeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DescribeRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DescribeRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

type DescribeResponse struct {
	Status       *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	LoadBalancer *LoadBalancer           `protobuf:"bytes,2,opt,name=load_balancer,json=loadBalancer" json:"load_balancer,omitempty"`
}

func (m *DescribeResponse) Reset()                    { *m = DescribeResponse{} }
func (m *DescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeResponse) ProtoMessage()               {}
func (*DescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *DescribeResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DescribeResponse) GetLoadBalancer() *LoadBalancer {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

type CreateRequest struct {
	Name         string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Namespace    string        `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	Cluster      string        `protobuf:"bytes,3,opt,name=cluster" json:"cluster,omitempty"`
	LoadBalancer *LoadBalancer `protobuf:"bytes,4,opt,name=load_balancer,json=loadBalancer" json:"load_balancer,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *CreateRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *CreateRequest) GetLoadBalancer() *LoadBalancer {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

type UpdateRequest struct {
	Name         string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Cluster      string        `protobuf:"bytes,2,opt,name=cluster" json:"cluster,omitempty"`
	LoadBalancer *LoadBalancer `protobuf:"bytes,3,opt,name=load_balancer,json=loadBalancer" json:"load_balancer,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func (m *UpdateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *UpdateRequest) GetLoadBalancer() *LoadBalancer {
	if m != nil {
		return m.LoadBalancer
	}
	return nil
}

type DeleteRequest struct {
	Kind      string `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	Cluster   string `protobuf:"bytes,4,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{6} }

func (m *DeleteRequest) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *DeleteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeleteRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *DeleteRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

type LoadBalancer struct {
	// 'kind' defines is it the regular kubernetes instance or the
	// appscode superset called Extended Ingress. This field will
	// strictly contains only those two values
	// 'ingress' - default kubernetes ingress object.
	// 'extendedIngress' - appscode superset of ingress.
	// when creating a Loadbalancer from UI this field will always
	// be only 'extendedIngress.' List, Describe, Update and Delete
	// will support both two modes.
	// Create will support only extendedIngress.
	// For Creating or Updating an regular ingress one must use the
	// kubectl or direct API calls directly to kubernetes.
	Kind              string            `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Name              string            `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Namespace         string            `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	CreationTimestamp int64             `protobuf:"varint,4,opt,name=creation_timestamp,json=creationTimestamp" json:"creation_timestamp,omitempty"`
	Options           map[string]string `protobuf:"bytes,5,rep,name=options" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Spec              *Spec             `protobuf:"bytes,6,opt,name=spec" json:"spec,omitempty"`
	Status            *Status           `protobuf:"bytes,7,opt,name=status" json:"status,omitempty"`
	Json              string            `protobuf:"bytes,8,opt,name=json" json:"json,omitempty"`
}

func (m *LoadBalancer) Reset()                    { *m = LoadBalancer{} }
func (m *LoadBalancer) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancer) ProtoMessage()               {}
func (*LoadBalancer) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{7} }

func (m *LoadBalancer) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *LoadBalancer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoadBalancer) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *LoadBalancer) GetCreationTimestamp() int64 {
	if m != nil {
		return m.CreationTimestamp
	}
	return 0
}

func (m *LoadBalancer) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *LoadBalancer) GetSpec() *Spec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *LoadBalancer) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *LoadBalancer) GetJson() string {
	if m != nil {
		return m.Json
	}
	return ""
}

type Spec struct {
	Backend *HTTPLoadBalancerRule `protobuf:"bytes,1,opt,name=backend" json:"backend,omitempty"`
	Rules   []*LoadBalancerRule   `protobuf:"bytes,2,rep,name=rules" json:"rules,omitempty"`
}

func (m *Spec) Reset()                    { *m = Spec{} }
func (m *Spec) String() string            { return proto.CompactTextString(m) }
func (*Spec) ProtoMessage()               {}
func (*Spec) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{8} }

func (m *Spec) GetBackend() *HTTPLoadBalancerRule {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (m *Spec) GetRules() []*LoadBalancerRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

type Status struct {
	Status []*LoadBalancerStatus `protobuf:"bytes,1,rep,name=status" json:"status,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{9} }

func (m *Status) GetStatus() []*LoadBalancerStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type LoadBalancerStatus struct {
	IP   string `protobuf:"bytes,1,opt,name=IP,json=iP" json:"IP,omitempty"`
	Host string `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
}

func (m *LoadBalancerStatus) Reset()                    { *m = LoadBalancerStatus{} }
func (m *LoadBalancerStatus) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancerStatus) ProtoMessage()               {}
func (*LoadBalancerStatus) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{10} }

func (m *LoadBalancerStatus) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *LoadBalancerStatus) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type LoadBalancerBackend struct {
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	ServicePort string `protobuf:"bytes,2,opt,name=service_port,json=servicePort" json:"service_port,omitempty"`
}

func (m *LoadBalancerBackend) Reset()                    { *m = LoadBalancerBackend{} }
func (m *LoadBalancerBackend) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancerBackend) ProtoMessage()               {}
func (*LoadBalancerBackend) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{11} }

func (m *LoadBalancerBackend) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *LoadBalancerBackend) GetServicePort() string {
	if m != nil {
		return m.ServicePort
	}
	return ""
}

type LoadBalancerRule struct {
	Host string `protobuf:"bytes,1,opt,name=host" json:"host,omitempty"`
	// ssl secret name to enable https on the host.
	// ssl secret must contain data with the certs pem file.
	SSLSecretName string                  `protobuf:"bytes,5,opt,name=SSL_secret_name,json=sSLSecretName" json:"SSL_secret_name,omitempty"`
	Http          []*HTTPLoadBalancerRule `protobuf:"bytes,2,rep,name=http" json:"http,omitempty"`
	Tcp           []*TCPLoadBalancerRule  `protobuf:"bytes,3,rep,name=tcp" json:"tcp,omitempty"`
}

func (m *LoadBalancerRule) Reset()                    { *m = LoadBalancerRule{} }
func (m *LoadBalancerRule) String() string            { return proto.CompactTextString(m) }
func (*LoadBalancerRule) ProtoMessage()               {}
func (*LoadBalancerRule) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{12} }

func (m *LoadBalancerRule) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *LoadBalancerRule) GetSSLSecretName() string {
	if m != nil {
		return m.SSLSecretName
	}
	return ""
}

func (m *LoadBalancerRule) GetHttp() []*HTTPLoadBalancerRule {
	if m != nil {
		return m.Http
	}
	return nil
}

func (m *LoadBalancerRule) GetTcp() []*TCPLoadBalancerRule {
	if m != nil {
		return m.Tcp
	}
	return nil
}

type HTTPLoadBalancerRule struct {
	Path         string               `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Backend      *LoadBalancerBackend `protobuf:"bytes,2,opt,name=backend" json:"backend,omitempty"`
	HeaderRules  []string             `protobuf:"bytes,3,rep,name=header_rules,json=headerRules" json:"header_rules,omitempty"`
	RewriteRules []string             `protobuf:"bytes,4,rep,name=rewrite_rules,json=rewriteRules" json:"rewrite_rules,omitempty"`
}

func (m *HTTPLoadBalancerRule) Reset()                    { *m = HTTPLoadBalancerRule{} }
func (m *HTTPLoadBalancerRule) String() string            { return proto.CompactTextString(m) }
func (*HTTPLoadBalancerRule) ProtoMessage()               {}
func (*HTTPLoadBalancerRule) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{13} }

func (m *HTTPLoadBalancerRule) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HTTPLoadBalancerRule) GetBackend() *LoadBalancerBackend {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (m *HTTPLoadBalancerRule) GetHeaderRules() []string {
	if m != nil {
		return m.HeaderRules
	}
	return nil
}

func (m *HTTPLoadBalancerRule) GetRewriteRules() []string {
	if m != nil {
		return m.RewriteRules
	}
	return nil
}

type TCPLoadBalancerRule struct {
	Port          string               `protobuf:"bytes,1,opt,name=port" json:"port,omitempty"`
	Backend       *LoadBalancerBackend `protobuf:"bytes,2,opt,name=backend" json:"backend,omitempty"`
	SSLSecretName string               `protobuf:"bytes,3,opt,name=SSL_secret_name,json=sSLSecretName" json:"SSL_secret_name,omitempty"`
	SecretPemName string               `protobuf:"bytes,4,opt,name=secret_pem_name,json=secretPemName" json:"secret_pem_name,omitempty"`
}

func (m *TCPLoadBalancerRule) Reset()                    { *m = TCPLoadBalancerRule{} }
func (m *TCPLoadBalancerRule) String() string            { return proto.CompactTextString(m) }
func (*TCPLoadBalancerRule) ProtoMessage()               {}
func (*TCPLoadBalancerRule) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{14} }

func (m *TCPLoadBalancerRule) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *TCPLoadBalancerRule) GetBackend() *LoadBalancerBackend {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (m *TCPLoadBalancerRule) GetSSLSecretName() string {
	if m != nil {
		return m.SSLSecretName
	}
	return ""
}

func (m *TCPLoadBalancerRule) GetSecretPemName() string {
	if m != nil {
		return m.SecretPemName
	}
	return ""
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "appscode.kubernetes.v1beta1.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "appscode.kubernetes.v1beta1.ListResponse")
	proto.RegisterType((*DescribeRequest)(nil), "appscode.kubernetes.v1beta1.DescribeRequest")
	proto.RegisterType((*DescribeResponse)(nil), "appscode.kubernetes.v1beta1.DescribeResponse")
	proto.RegisterType((*CreateRequest)(nil), "appscode.kubernetes.v1beta1.CreateRequest")
	proto.RegisterType((*UpdateRequest)(nil), "appscode.kubernetes.v1beta1.UpdateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "appscode.kubernetes.v1beta1.DeleteRequest")
	proto.RegisterType((*LoadBalancer)(nil), "appscode.kubernetes.v1beta1.LoadBalancer")
	proto.RegisterType((*Spec)(nil), "appscode.kubernetes.v1beta1.Spec")
	proto.RegisterType((*Status)(nil), "appscode.kubernetes.v1beta1.Status")
	proto.RegisterType((*LoadBalancerStatus)(nil), "appscode.kubernetes.v1beta1.LoadBalancerStatus")
	proto.RegisterType((*LoadBalancerBackend)(nil), "appscode.kubernetes.v1beta1.LoadBalancerBackend")
	proto.RegisterType((*LoadBalancerRule)(nil), "appscode.kubernetes.v1beta1.LoadBalancerRule")
	proto.RegisterType((*HTTPLoadBalancerRule)(nil), "appscode.kubernetes.v1beta1.HTTPLoadBalancerRule")
	proto.RegisterType((*TCPLoadBalancerRule)(nil), "appscode.kubernetes.v1beta1.TCPLoadBalancerRule")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LoadBalancers service

type LoadBalancersClient interface {
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*DescribeResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type loadBalancersClient struct {
	cc *grpc.ClientConn
}

func NewLoadBalancersClient(cc *grpc.ClientConn) LoadBalancersClient {
	return &loadBalancersClient{cc}
}

func (c *loadBalancersClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := grpc.Invoke(ctx, "/appscode.kubernetes.v1beta1.LoadBalancers/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*DescribeResponse, error) {
	out := new(DescribeResponse)
	err := grpc.Invoke(ctx, "/appscode.kubernetes.v1beta1.LoadBalancers/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.kubernetes.v1beta1.LoadBalancers/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.kubernetes.v1beta1.LoadBalancers/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loadBalancersClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.kubernetes.v1beta1.LoadBalancers/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LoadBalancers service

type LoadBalancersServer interface {
	List(context.Context, *ListRequest) (*ListResponse, error)
	Describe(context.Context, *DescribeRequest) (*DescribeResponse, error)
	Create(context.Context, *CreateRequest) (*appscode_dtypes.VoidResponse, error)
	Update(context.Context, *UpdateRequest) (*appscode_dtypes.VoidResponse, error)
	Delete(context.Context, *DeleteRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterLoadBalancersServer(s *grpc.Server, srv LoadBalancersServer) {
	s.RegisterService(&_LoadBalancers_serviceDesc, srv)
}

func _LoadBalancers_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.kubernetes.v1beta1.LoadBalancers/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancers_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.kubernetes.v1beta1.LoadBalancers/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).Describe(ctx, req.(*DescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancers_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.kubernetes.v1beta1.LoadBalancers/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancers_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.kubernetes.v1beta1.LoadBalancers/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LoadBalancers_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoadBalancersServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.kubernetes.v1beta1.LoadBalancers/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoadBalancersServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoadBalancers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.kubernetes.v1beta1.LoadBalancers",
	HandlerType: (*LoadBalancersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _LoadBalancers_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _LoadBalancers_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _LoadBalancers_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _LoadBalancers_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _LoadBalancers_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loadbalancer.proto",
}

func init() { proto.RegisterFile("loadbalancer.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 1020 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x57, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0xd7, 0x78, 0x37, 0x76, 0xf3, 0x6c, 0x37, 0xe9, 0xb4, 0x12, 0x2b, 0x53, 0xa0, 0xd9, 0x8a,
	0x92, 0x46, 0xc4, 0xdb, 0x04, 0x81, 0xaa, 0x70, 0x73, 0x1a, 0x95, 0x82, 0x15, 0x16, 0x3b, 0x45,
	0x08, 0x90, 0xac, 0xf1, 0x7a, 0xd4, 0x2c, 0xb1, 0x77, 0xa6, 0x3b, 0xe3, 0xa0, 0xa8, 0xea, 0xa5,
	0x77, 0x0e, 0xa8, 0x17, 0xf8, 0x04, 0x1c, 0x10, 0x07, 0x10, 0x5c, 0xb9, 0xf5, 0x13, 0x70, 0xe2,
	0xce, 0x07, 0x41, 0xf3, 0x67, 0xe3, 0xdd, 0x26, 0xb5, 0x4d, 0x6b, 0x71, 0x49, 0x66, 0xe7, 0xfd,
	0x99, 0xdf, 0xef, 0xf7, 0xde, 0xac, 0xdf, 0x02, 0x1e, 0x32, 0x32, 0xe8, 0x93, 0x21, 0x49, 0x22,
	0x9a, 0x36, 0x79, 0xca, 0x24, 0xc3, 0xaf, 0x13, 0xce, 0x45, 0xc4, 0x06, 0xb4, 0x79, 0x34, 0xee,
	0xd3, 0x34, 0xa1, 0x92, 0x8a, 0xe6, 0xf1, 0x56, 0x9f, 0x4a, 0xb2, 0xd5, 0xb8, 0xfa, 0x80, 0xb1,
	0x07, 0x43, 0x1a, 0x10, 0x1e, 0x07, 0x24, 0x49, 0x98, 0x24, 0x32, 0x66, 0x89, 0x30, 0xa1, 0x8d,
	0x37, 0xb3, 0xd0, 0x17, 0xd8, 0xdf, 0x2a, 0xd8, 0x07, 0xf2, 0x84, 0x53, 0x11, 0xe8, 0xbf, 0xc6,
	0xc1, 0x7f, 0x07, 0xaa, 0xed, 0x58, 0xc8, 0x0e, 0x7d, 0x38, 0xa6, 0x42, 0x62, 0x0f, 0x2a, 0xd1,
	0x70, 0x2c, 0x24, 0x4d, 0x3d, 0x74, 0x0d, 0xad, 0x2f, 0x77, 0xb2, 0x47, 0xff, 0x7b, 0x04, 0x35,
	0xe3, 0x29, 0x38, 0x4b, 0x04, 0xc5, 0x01, 0x94, 0x85, 0x24, 0x72, 0x2c, 0xb4, 0x67, 0x75, 0xfb,
	0xb5, 0xe6, 0x29, 0x0d, 0x73, 0x4e, 0xb3, 0xab, 0xcd, 0x1d, 0xeb, 0x86, 0x43, 0xb8, 0xa8, 0xc8,
	0xf7, 0x32, 0xf6, 0xc2, 0x2b, 0x5d, 0x73, 0xd6, 0xab, 0xdb, 0x37, 0x9b, 0x53, 0xf8, 0x37, 0xdb,
	0x8c, 0x0c, 0x5a, 0x36, 0xa2, 0x53, 0x1f, 0xe6, 0x9e, 0x84, 0xff, 0x10, 0x56, 0xee, 0x50, 0x11,
	0xa5, 0x71, 0x9f, 0x66, 0x04, 0x30, 0xb8, 0x47, 0x71, 0x32, 0xb0, 0xe8, 0xf5, 0x5a, 0xed, 0x25,
	0x64, 0x44, 0xbd, 0x92, 0xd9, 0x53, 0x6b, 0x7c, 0x15, 0x96, 0xd5, 0x7f, 0xc1, 0x49, 0x44, 0x3d,
	0x47, 0x1b, 0x26, 0x1b, 0x79, 0x19, 0xdc, 0xa2, 0x0c, 0x4f, 0x11, 0xac, 0x4e, 0xce, 0x7c, 0x59,
	0x29, 0xf6, 0xa1, 0x5e, 0x90, 0x42, 0x43, 0xfb, 0x4f, 0x4a, 0xd4, 0xf2, 0x4a, 0xf8, 0x3f, 0x23,
	0xa8, 0xef, 0xa6, 0x94, 0xc8, 0xbc, 0x0e, 0x9a, 0x33, 0x7a, 0x11, 0xe7, 0xd2, 0x14, 0xce, 0x4e,
	0x81, 0xf3, 0x59, 0xb4, 0xee, 0xab, 0xa1, 0xfd, 0x0e, 0x41, 0xfd, 0x3e, 0x1f, 0xcc, 0x40, 0x9b,
	0xc3, 0x53, 0x9a, 0x81, 0xc7, 0x79, 0x35, 0x3c, 0x0c, 0xea, 0x77, 0xe8, 0x90, 0xca, 0xff, 0xad,
	0x89, 0x7e, 0x74, 0xa0, 0x96, 0xc7, 0xb3, 0xa0, 0x03, 0x37, 0x01, 0x47, 0xaa, 0x09, 0x62, 0x96,
	0xf4, 0x64, 0x3c, 0xa2, 0x42, 0x92, 0x11, 0xd7, 0x67, 0x3b, 0x9d, 0x4b, 0x99, 0xe5, 0x20, 0x33,
	0xe0, 0x10, 0x2a, 0x8c, 0xeb, 0x97, 0x85, 0xb7, 0xa4, 0x2f, 0xe2, 0x07, 0x73, 0x0b, 0xd8, 0xfc,
	0xd4, 0x04, 0xee, 0x25, 0x32, 0x3d, 0xe9, 0x64, 0x69, 0xf0, 0xfb, 0xe0, 0x0a, 0x4e, 0x23, 0xaf,
	0xac, 0xeb, 0xb1, 0x36, 0x35, 0x5d, 0x97, 0xd3, 0xa8, 0xa3, 0xdd, 0xf1, 0x87, 0xa7, 0xd7, 0xa7,
	0xa2, 0x03, 0xaf, 0x4f, 0x0f, 0x2c, 0x5e, 0x25, 0x0c, 0xee, 0x37, 0x82, 0x25, 0xde, 0x05, 0x23,
	0x93, 0x5a, 0x37, 0x76, 0xa0, 0x96, 0x07, 0x88, 0x57, 0xc1, 0x39, 0xa2, 0x27, 0x56, 0x5d, 0xb5,
	0xc4, 0x57, 0x60, 0xe9, 0x98, 0x0c, 0xc7, 0x99, 0xba, 0xe6, 0x61, 0xa7, 0x74, 0x1b, 0xf9, 0x3f,
	0x20, 0x70, 0x15, 0x36, 0xfc, 0x09, 0x54, 0xfa, 0x24, 0x3a, 0xa2, 0xb6, 0x2c, 0xd5, 0xed, 0xad,
	0xa9, 0xb0, 0x3e, 0x3a, 0x38, 0x08, 0x0b, 0x3d, 0x36, 0x1e, 0xd2, 0x4e, 0x96, 0x01, 0xef, 0xc2,
	0x52, 0x3a, 0x1e, 0xd2, 0xec, 0x95, 0xb7, 0x39, 0x7f, 0xab, 0xaa, 0x34, 0x26, 0xd6, 0xff, 0x0c,
	0xca, 0x86, 0x3c, 0xbe, 0x9b, 0x7b, 0xe1, 0xa8, 0x7c, 0xc1, 0xdc, 0xf9, 0x8a, 0xea, 0xf9, 0xb7,
	0x01, 0x9f, 0xb5, 0xe2, 0x8b, 0x50, 0xba, 0x17, 0x5a, 0xb9, 0x4a, 0x71, 0xa8, 0x34, 0x3e, 0x64,
	0x42, 0x66, 0xad, 0xa8, 0xd6, 0xfe, 0x57, 0x70, 0x39, 0x1f, 0xd9, 0xb2, 0x44, 0xd7, 0xa0, 0x26,
	0x68, 0x7a, 0x1c, 0x47, 0xb4, 0x97, 0xbb, 0xd1, 0x55, 0xbb, 0xb7, 0xaf, 0x9a, 0x38, 0xe7, 0xc2,
	0x59, 0x9a, 0x65, 0xcd, 0x5c, 0x42, 0x96, 0x4a, 0xff, 0x6f, 0x04, 0xab, 0xcf, 0xab, 0x70, 0x8a,
	0x02, 0x4d, 0x50, 0xe0, 0x1b, 0xb0, 0xd2, 0xed, 0xb6, 0x7b, 0x82, 0x46, 0x29, 0x95, 0xe6, 0xc4,
	0x25, 0x6d, 0xae, 0x8b, 0x6e, 0xbb, 0xab, 0x77, 0xf5, 0x99, 0x7b, 0xe0, 0x1e, 0x4a, 0xc9, 0xad,
	0xfc, 0x2f, 0x51, 0x49, 0x1d, 0x8e, 0x5b, 0xe0, 0xc8, 0x88, 0x7b, 0x8e, 0xce, 0x72, 0x6b, 0x6a,
	0x96, 0x83, 0xdd, 0xb3, 0x49, 0x54, 0xb0, 0xff, 0x27, 0x82, 0x2b, 0xe7, 0x1d, 0xa1, 0xf8, 0x71,
	0x22, 0x0f, 0x33, 0x7e, 0x6a, 0x8d, 0x3f, 0x9e, 0x34, 0xa1, 0xf9, 0x89, 0xb8, 0x35, 0x77, 0xa5,
	0x6d, 0x45, 0x26, 0x3d, 0xb8, 0x06, 0xb5, 0x43, 0x4a, 0x06, 0x34, 0xed, 0x99, 0x56, 0x54, 0x2c,
	0x96, 0x3b, 0x55, 0xb3, 0xa7, 0x10, 0x08, 0x7c, 0x1d, 0xea, 0x29, 0xfd, 0x36, 0x8d, 0x25, 0xb5,
	0x3e, 0xae, 0xf6, 0xa9, 0xd9, 0x4d, 0xed, 0xe4, 0x3f, 0x43, 0x70, 0xf9, 0x1c, 0x76, 0x1a, 0xbf,
	0xaa, 0x67, 0x86, 0x9f, 0xa5, 0x72, 0xa1, 0xf8, 0xcf, 0xa9, 0xb5, 0x73, 0x5e, 0xad, 0x6f, 0xc0,
	0x8a, 0xf5, 0xe1, 0x74, 0x64, 0xfc, 0x5c, 0xeb, 0xa7, 0xb7, 0x43, 0x3a, 0x52, 0x7e, 0xdb, 0x3f,
	0x55, 0xa0, 0x9e, 0x3f, 0x50, 0xe0, 0xdf, 0x10, 0xb8, 0x6a, 0xc6, 0xc1, 0xeb, 0xd3, 0x51, 0x4e,
	0x06, 0xa6, 0xc6, 0xcd, 0x39, 0x3c, 0xcd, 0x94, 0xe0, 0xdf, 0x7f, 0xf2, 0x87, 0x57, 0xba, 0x80,
	0x9e, 0xfc, 0xf5, 0xcf, 0xd3, 0xd2, 0x3d, 0x7c, 0x37, 0xe8, 0x15, 0x66, 0xb3, 0x49, 0x74, 0x60,
	0xa3, 0x03, 0xfb, 0x73, 0x21, 0x82, 0x47, 0x76, 0xf5, 0x38, 0xc8, 0x4f, 0x8f, 0x22, 0x50, 0x2f,
	0x3b, 0xfc, 0x0c, 0xc1, 0x85, 0x6c, 0x22, 0xc1, 0xef, 0x4e, 0x85, 0xf3, 0xdc, 0xb0, 0xd4, 0xd8,
	0x9c, 0xd3, 0xdb, 0x12, 0xf8, 0x3a, 0x47, 0x20, 0xc4, 0xfb, 0x0b, 0x20, 0xf0, 0x48, 0x55, 0xe7,
	0xb1, 0xe1, 0xf1, 0x0b, 0x82, 0xb2, 0x99, 0x61, 0xf0, 0xc6, 0x54, 0x5c, 0x85, 0x41, 0xa7, 0xf1,
	0xc6, 0x99, 0x59, 0xeb, 0x73, 0x16, 0x0f, 0x4e, 0x31, 0x7f, 0x91, 0xc3, 0xdc, 0xf6, 0x17, 0x25,
	0xfa, 0x0e, 0xda, 0xc0, 0xbf, 0x23, 0x28, 0x9b, 0x29, 0x66, 0x06, 0xde, 0xc2, 0xa8, 0x33, 0x0b,
	0x6f, 0x2f, 0x87, 0xb7, 0xdb, 0x58, 0xb0, 0xc6, 0x0a, 0xf6, 0xaf, 0x08, 0xca, 0x66, 0xda, 0x99,
	0x01, 0xbb, 0x30, 0x12, 0xcd, 0x82, 0x5d, 0x68, 0x8d, 0x8d, 0x05, 0xc3, 0x6e, 0xed, 0xc1, 0xdb,
	0x11, 0x1b, 0x4d, 0x10, 0x10, 0x1e, 0x37, 0x0b, 0xdf, 0x51, 0x36, 0x6b, 0xeb, 0x52, 0xfe, 0x3e,
	0x87, 0xea, 0x03, 0x27, 0x44, 0x5f, 0x56, 0xac, 0xb5, 0x5f, 0xd6, 0x9f, 0x3c, 0xef, 0xfd, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0x88, 0xec, 0xa2, 0xe1, 0x84, 0x0d, 0x00, 0x00,
}
