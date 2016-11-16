// Code generated by protoc-gen-go.
// source: alert.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	alert.proto
	incident.proto

It has these top-level messages:
	IcingaParam
	IcingaState
	NotifierParam
	Alert
	AlertListRequest
	AlertListResponse
	AlertCreateRequest
	AlertDescribeRequest
	AlertDescribeResponse
	AlertUpdateRequest
	AlertSyncRequest
	AlertDeleteRequest
	Incident
	IncidentListRequest
	IncidentListResponse
	IncidentDescribeRequest
	IncidentDescribeResponse
	IncidentNotifyRequest
	IncidentEventCreateRequest
*/
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type IcingaParam struct {
	CheckIntervalSec int64 `protobuf:"varint,1,opt,name=check_interval_sec,json=checkIntervalSec" json:"check_interval_sec,omitempty"`
	AlertIntervalSec int64 `protobuf:"varint,2,opt,name=alert_interval_sec,json=alertIntervalSec" json:"alert_interval_sec,omitempty"`
}

func (m *IcingaParam) Reset()                    { *m = IcingaParam{} }
func (m *IcingaParam) String() string            { return proto.CompactTextString(m) }
func (*IcingaParam) ProtoMessage()               {}
func (*IcingaParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// This one is used in kubernetes/v1beta1/client.proto
// To send information if alert is set for app/pod
type IcingaState struct {
	OK       int32 `protobuf:"varint,1,opt,name=OK,json=oK" json:"OK,omitempty"`
	Warning  int32 `protobuf:"varint,2,opt,name=Warning,json=warning" json:"Warning,omitempty"`
	Critical int32 `protobuf:"varint,3,opt,name=Critical,json=critical" json:"Critical,omitempty"`
	Unknown  int32 `protobuf:"varint,4,opt,name=Unknown,json=unknown" json:"Unknown,omitempty"`
}

func (m *IcingaState) Reset()                    { *m = IcingaState{} }
func (m *IcingaState) String() string            { return proto.CompactTextString(m) }
func (*IcingaState) ProtoMessage()               {}
func (*IcingaState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type NotifierParam struct {
	State   string `protobuf:"bytes,1,opt,name=state" json:"state,omitempty"`
	UserUid string `protobuf:"bytes,2,opt,name=user_uid,json=userUid" json:"user_uid,omitempty"`
	Method  string `protobuf:"bytes,3,opt,name=method" json:"method,omitempty"`
}

func (m *NotifierParam) Reset()                    { *m = NotifierParam{} }
func (m *NotifierParam) String() string            { return proto.CompactTextString(m) }
func (*NotifierParam) ProtoMessage()               {}
func (*NotifierParam) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Alert struct {
	Phid                 string            `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	KubernetesCluster    string            `protobuf:"bytes,2,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
	KubernetesNamespace  string            `protobuf:"bytes,3,opt,name=kubernetes_namespace,json=kubernetesNamespace" json:"kubernetes_namespace,omitempty"`
	KubernetesObjectType string            `protobuf:"bytes,4,opt,name=kubernetes_objectType,json=kubernetesObjectType" json:"kubernetes_objectType,omitempty"`
	KubernetesObjectName string            `protobuf:"bytes,5,opt,name=kubernetes_objectName,json=kubernetesObjectName" json:"kubernetes_objectName,omitempty"`
	IcingaService        string            `protobuf:"bytes,7,opt,name=icinga_service,json=icingaService" json:"icinga_service,omitempty"`
	IcingaParam          *IcingaParam      `protobuf:"bytes,8,opt,name=icinga_param,json=icingaParam" json:"icinga_param,omitempty"`
	CheckCommand         string            `protobuf:"bytes,9,opt,name=check_command,json=checkCommand" json:"check_command,omitempty"`
	NotifierParams       []*NotifierParam  `protobuf:"bytes,11,rep,name=notifier_params,json=notifierParams" json:"notifier_params,omitempty"`
	IcingawebUrl         string            `protobuf:"bytes,12,opt,name=icingaweb_url,json=icingawebUrl" json:"icingaweb_url,omitempty"`
	Vars                 map[string]string `protobuf:"bytes,13,rep,name=vars" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Alert) Reset()                    { *m = Alert{} }
func (m *Alert) String() string            { return proto.CompactTextString(m) }
func (*Alert) ProtoMessage()               {}
func (*Alert) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Alert) GetIcingaParam() *IcingaParam {
	if m != nil {
		return m.IcingaParam
	}
	return nil
}

func (m *Alert) GetNotifierParams() []*NotifierParam {
	if m != nil {
		return m.NotifierParams
	}
	return nil
}

func (m *Alert) GetVars() map[string]string {
	if m != nil {
		return m.Vars
	}
	return nil
}

// Next Id: 5
type AlertListRequest struct {
	KubernetesCluster    string `protobuf:"bytes,1,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
	KubernetesNamespace  string `protobuf:"bytes,2,opt,name=kubernetes_namespace,json=kubernetesNamespace" json:"kubernetes_namespace,omitempty"`
	KubernetesObjectType string `protobuf:"bytes,3,opt,name=kubernetes_objectType,json=kubernetesObjectType" json:"kubernetes_objectType,omitempty"`
	KubernetesObjectName string `protobuf:"bytes,4,opt,name=kubernetes_objectName,json=kubernetesObjectName" json:"kubernetes_objectName,omitempty"`
}

func (m *AlertListRequest) Reset()                    { *m = AlertListRequest{} }
func (m *AlertListRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertListRequest) ProtoMessage()               {}
func (*AlertListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type AlertListResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Alerts []*Alert                `protobuf:"bytes,2,rep,name=alerts" json:"alerts,omitempty"`
}

func (m *AlertListResponse) Reset()                    { *m = AlertListResponse{} }
func (m *AlertListResponse) String() string            { return proto.CompactTextString(m) }
func (*AlertListResponse) ProtoMessage()               {}
func (*AlertListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AlertListResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AlertListResponse) GetAlerts() []*Alert {
	if m != nil {
		return m.Alerts
	}
	return nil
}

// Next Id: 12
type AlertCreateRequest struct {
	KubernetesCluster    string            `protobuf:"bytes,1,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
	KubernetesNamespace  string            `protobuf:"bytes,2,opt,name=kubernetes_namespace,json=kubernetesNamespace" json:"kubernetes_namespace,omitempty"`
	KubernetesObjectType string            `protobuf:"bytes,3,opt,name=kubernetes_objectType,json=kubernetesObjectType" json:"kubernetes_objectType,omitempty"`
	KubernetesObjectName string            `protobuf:"bytes,4,opt,name=kubernetes_objectName,json=kubernetesObjectName" json:"kubernetes_objectName,omitempty"`
	IcingaService        string            `protobuf:"bytes,6,opt,name=icinga_service,json=icingaService" json:"icinga_service,omitempty"`
	IcingaParam          *IcingaParam      `protobuf:"bytes,7,opt,name=icinga_param,json=icingaParam" json:"icinga_param,omitempty"`
	CheckCommand         string            `protobuf:"bytes,8,opt,name=check_command,json=checkCommand" json:"check_command,omitempty"`
	NotifierParams       []*NotifierParam  `protobuf:"bytes,10,rep,name=notifier_params,json=notifierParams" json:"notifier_params,omitempty"`
	Vars                 map[string]string `protobuf:"bytes,11,rep,name=vars" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *AlertCreateRequest) Reset()                    { *m = AlertCreateRequest{} }
func (m *AlertCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertCreateRequest) ProtoMessage()               {}
func (*AlertCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AlertCreateRequest) GetIcingaParam() *IcingaParam {
	if m != nil {
		return m.IcingaParam
	}
	return nil
}

func (m *AlertCreateRequest) GetNotifierParams() []*NotifierParam {
	if m != nil {
		return m.NotifierParams
	}
	return nil
}

func (m *AlertCreateRequest) GetVars() map[string]string {
	if m != nil {
		return m.Vars
	}
	return nil
}

// Next Id: 2
type AlertDescribeRequest struct {
	Phid string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
}

func (m *AlertDescribeRequest) Reset()                    { *m = AlertDescribeRequest{} }
func (m *AlertDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertDescribeRequest) ProtoMessage()               {}
func (*AlertDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type AlertDescribeResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Alerts *Alert                  `protobuf:"bytes,2,opt,name=alerts" json:"alerts,omitempty"`
}

func (m *AlertDescribeResponse) Reset()                    { *m = AlertDescribeResponse{} }
func (m *AlertDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*AlertDescribeResponse) ProtoMessage()               {}
func (*AlertDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AlertDescribeResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AlertDescribeResponse) GetAlerts() *Alert {
	if m != nil {
		return m.Alerts
	}
	return nil
}

// Next Id: 6
type AlertUpdateRequest struct {
	Phid           string            `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	IcingaParam    *IcingaParam      `protobuf:"bytes,2,opt,name=icinga_param,json=icingaParam" json:"icinga_param,omitempty"`
	NotifierParams []*NotifierParam  `protobuf:"bytes,4,rep,name=notifier_params,json=notifierParams" json:"notifier_params,omitempty"`
	Vars           map[string]string `protobuf:"bytes,5,rep,name=vars" json:"vars,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *AlertUpdateRequest) Reset()                    { *m = AlertUpdateRequest{} }
func (m *AlertUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertUpdateRequest) ProtoMessage()               {}
func (*AlertUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AlertUpdateRequest) GetIcingaParam() *IcingaParam {
	if m != nil {
		return m.IcingaParam
	}
	return nil
}

func (m *AlertUpdateRequest) GetNotifierParams() []*NotifierParam {
	if m != nil {
		return m.NotifierParams
	}
	return nil
}

func (m *AlertUpdateRequest) GetVars() map[string]string {
	if m != nil {
		return m.Vars
	}
	return nil
}

// Next Id: 6
type AlertSyncRequest struct {
	KubernetesCluster    string                          `protobuf:"bytes,1,opt,name=kubernetes_cluster,json=kubernetesCluster" json:"kubernetes_cluster,omitempty"`
	KubernetesNamespace  string                          `protobuf:"bytes,2,opt,name=kubernetes_namespace,json=kubernetesNamespace" json:"kubernetes_namespace,omitempty"`
	KubernetesObjectType string                          `protobuf:"bytes,3,opt,name=kubernetes_objectType,json=kubernetesObjectType" json:"kubernetes_objectType,omitempty"`
	KubernetesObjectName string                          `protobuf:"bytes,4,opt,name=kubernetes_objectName,json=kubernetesObjectName" json:"kubernetes_objectName,omitempty"`
	PodAncestors         []*AlertSyncRequest_PodAncestor `protobuf:"bytes,5,rep,name=pod_ancestors,json=podAncestors" json:"pod_ancestors,omitempty"`
}

func (m *AlertSyncRequest) Reset()                    { *m = AlertSyncRequest{} }
func (m *AlertSyncRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertSyncRequest) ProtoMessage()               {}
func (*AlertSyncRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *AlertSyncRequest) GetPodAncestors() []*AlertSyncRequest_PodAncestor {
	if m != nil {
		return m.PodAncestors
	}
	return nil
}

// Next Id: 3
type AlertSyncRequest_PodAncestor struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *AlertSyncRequest_PodAncestor) Reset()         { *m = AlertSyncRequest_PodAncestor{} }
func (m *AlertSyncRequest_PodAncestor) String() string { return proto.CompactTextString(m) }
func (*AlertSyncRequest_PodAncestor) ProtoMessage()    {}
func (*AlertSyncRequest_PodAncestor) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{10, 0}
}

// Next Id: 2
type AlertDeleteRequest struct {
	Phid string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
}

func (m *AlertDeleteRequest) Reset()                    { *m = AlertDeleteRequest{} }
func (m *AlertDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*AlertDeleteRequest) ProtoMessage()               {}
func (*AlertDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*IcingaParam)(nil), "appscode.alert.v1beta1.IcingaParam")
	proto.RegisterType((*IcingaState)(nil), "appscode.alert.v1beta1.IcingaState")
	proto.RegisterType((*NotifierParam)(nil), "appscode.alert.v1beta1.NotifierParam")
	proto.RegisterType((*Alert)(nil), "appscode.alert.v1beta1.Alert")
	proto.RegisterType((*AlertListRequest)(nil), "appscode.alert.v1beta1.AlertListRequest")
	proto.RegisterType((*AlertListResponse)(nil), "appscode.alert.v1beta1.AlertListResponse")
	proto.RegisterType((*AlertCreateRequest)(nil), "appscode.alert.v1beta1.AlertCreateRequest")
	proto.RegisterType((*AlertDescribeRequest)(nil), "appscode.alert.v1beta1.AlertDescribeRequest")
	proto.RegisterType((*AlertDescribeResponse)(nil), "appscode.alert.v1beta1.AlertDescribeResponse")
	proto.RegisterType((*AlertUpdateRequest)(nil), "appscode.alert.v1beta1.AlertUpdateRequest")
	proto.RegisterType((*AlertSyncRequest)(nil), "appscode.alert.v1beta1.AlertSyncRequest")
	proto.RegisterType((*AlertSyncRequest_PodAncestor)(nil), "appscode.alert.v1beta1.AlertSyncRequest.PodAncestor")
	proto.RegisterType((*AlertDeleteRequest)(nil), "appscode.alert.v1beta1.AlertDeleteRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Alerts service

type AlertsClient interface {
	List(ctx context.Context, in *AlertListRequest, opts ...grpc.CallOption) (*AlertListResponse, error)
	Create(ctx context.Context, in *AlertCreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Describe(ctx context.Context, in *AlertDescribeRequest, opts ...grpc.CallOption) (*AlertDescribeResponse, error)
	Update(ctx context.Context, in *AlertUpdateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Sync(ctx context.Context, in *AlertSyncRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *AlertDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error)
}

type alertsClient struct {
	cc *grpc.ClientConn
}

func NewAlertsClient(cc *grpc.ClientConn) AlertsClient {
	return &alertsClient{cc}
}

func (c *alertsClient) List(ctx context.Context, in *AlertListRequest, opts ...grpc.CallOption) (*AlertListResponse, error) {
	out := new(AlertListResponse)
	err := grpc.Invoke(ctx, "/appscode.alert.v1beta1.Alerts/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) Create(ctx context.Context, in *AlertCreateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.alert.v1beta1.Alerts/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) Describe(ctx context.Context, in *AlertDescribeRequest, opts ...grpc.CallOption) (*AlertDescribeResponse, error) {
	out := new(AlertDescribeResponse)
	err := grpc.Invoke(ctx, "/appscode.alert.v1beta1.Alerts/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) Update(ctx context.Context, in *AlertUpdateRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.alert.v1beta1.Alerts/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) Sync(ctx context.Context, in *AlertSyncRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.alert.v1beta1.Alerts/Sync", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertsClient) Delete(ctx context.Context, in *AlertDeleteRequest, opts ...grpc.CallOption) (*appscode_dtypes.VoidResponse, error) {
	out := new(appscode_dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/appscode.alert.v1beta1.Alerts/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Alerts service

type AlertsServer interface {
	List(context.Context, *AlertListRequest) (*AlertListResponse, error)
	Create(context.Context, *AlertCreateRequest) (*appscode_dtypes.VoidResponse, error)
	Describe(context.Context, *AlertDescribeRequest) (*AlertDescribeResponse, error)
	Update(context.Context, *AlertUpdateRequest) (*appscode_dtypes.VoidResponse, error)
	Sync(context.Context, *AlertSyncRequest) (*appscode_dtypes.VoidResponse, error)
	Delete(context.Context, *AlertDeleteRequest) (*appscode_dtypes.VoidResponse, error)
}

func RegisterAlertsServer(s *grpc.Server, srv AlertsServer) {
	s.RegisterService(&_Alerts_serviceDesc, srv)
}

func _Alerts_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.alert.v1beta1.Alerts/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).List(ctx, req.(*AlertListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerts_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.alert.v1beta1.Alerts/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).Create(ctx, req.(*AlertCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerts_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.alert.v1beta1.Alerts/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).Describe(ctx, req.(*AlertDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerts_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.alert.v1beta1.Alerts/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).Update(ctx, req.(*AlertUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerts_Sync_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertSyncRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).Sync(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.alert.v1beta1.Alerts/Sync",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).Sync(ctx, req.(*AlertSyncRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Alerts_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlertDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.alert.v1beta1.Alerts/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertsServer).Delete(ctx, req.(*AlertDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Alerts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.alert.v1beta1.Alerts",
	HandlerType: (*AlertsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Alerts_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Alerts_Create_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Alerts_Describe_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Alerts_Update_Handler,
		},
		{
			MethodName: "Sync",
			Handler:    _Alerts_Sync_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Alerts_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("alert.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1070 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xdc, 0x57, 0xcb, 0x6e, 0x1c, 0x45,
	0x17, 0x56, 0xb7, 0xe7, 0x7a, 0xc6, 0xf6, 0xef, 0xd4, 0xef, 0x84, 0x66, 0x44, 0x82, 0xd5, 0x51,
	0x14, 0xc7, 0xb2, 0xa7, 0xb1, 0x8d, 0x95, 0x28, 0x20, 0xa1, 0xc4, 0x01, 0x11, 0x05, 0x39, 0x56,
	0x1b, 0x27, 0xc0, 0x66, 0x54, 0x53, 0x5d, 0xd8, 0x85, 0x67, 0xaa, 0x9b, 0xae, 0x1a, 0x5b, 0x23,
	0x2b, 0x42, 0xca, 0x9e, 0x05, 0x97, 0x57, 0xe0, 0x0d, 0x40, 0x3c, 0x00, 0x0b, 0x1e, 0x80, 0x3d,
	0x2b, 0xc4, 0x9a, 0x47, 0x40, 0x75, 0xb1, 0xa7, 0xdb, 0x19, 0x7b, 0xda, 0xc6, 0xab, 0x6c, 0x46,
	0x5d, 0xe7, 0xfa, 0xd5, 0x39, 0xdf, 0x39, 0x3d, 0x0d, 0x0d, 0xdc, 0xa5, 0xa9, 0x6c, 0x25, 0x69,
	0x2c, 0x63, 0x74, 0x0d, 0x27, 0x89, 0x20, 0x71, 0x44, 0x5b, 0x46, 0xba, 0xbf, 0xdc, 0xa1, 0x12,
	0x2f, 0x37, 0xdf, 0xda, 0x89, 0xe3, 0x9d, 0x2e, 0x0d, 0x70, 0xc2, 0x02, 0xcc, 0x79, 0x2c, 0xb1,
	0x64, 0x31, 0x17, 0xc6, 0xab, 0x79, 0xe3, 0xc8, 0xeb, 0x14, 0xfd, 0xdb, 0x39, 0x7d, 0x24, 0x07,
	0x09, 0x15, 0x81, 0xfe, 0x35, 0x06, 0x3e, 0x83, 0xc6, 0x63, 0xc2, 0xf8, 0x0e, 0xde, 0xc4, 0x29,
	0xee, 0xa1, 0x45, 0x40, 0x64, 0x97, 0x92, 0xbd, 0x36, 0xe3, 0x92, 0xa6, 0xfb, 0xb8, 0xdb, 0x16,
	0x94, 0x78, 0xce, 0x9c, 0x33, 0x3f, 0x11, 0xce, 0x68, 0xcd, 0x63, 0xab, 0xd8, 0xa2, 0x44, 0x59,
	0x6b, 0xb0, 0x79, 0x6b, 0xd7, 0x58, 0x6b, 0x4d, 0xc6, 0xda, 0xef, 0x1d, 0xa5, 0xda, 0x92, 0x58,
	0x52, 0x34, 0x0d, 0xee, 0xd3, 0x27, 0x3a, 0x74, 0x39, 0x74, 0xe3, 0x27, 0xc8, 0x83, 0xea, 0x73,
	0x9c, 0x72, 0xc6, 0x77, 0x74, 0x84, 0x72, 0x58, 0x3d, 0x30, 0x47, 0xd4, 0x84, 0xda, 0x7a, 0xca,
	0x24, 0x23, 0xb8, 0xeb, 0x4d, 0x68, 0x55, 0x8d, 0xd8, 0xb3, 0xf2, 0xda, 0xe6, 0x7b, 0x3c, 0x3e,
	0xe0, 0x5e, 0xc9, 0x78, 0xf5, 0xcd, 0xd1, 0xff, 0x0c, 0xa6, 0x36, 0x62, 0xc9, 0xbe, 0x64, 0x34,
	0x35, 0x77, 0x9b, 0x85, 0xb2, 0x50, 0x99, 0x75, 0xce, 0x7a, 0x68, 0x0e, 0xe8, 0x4d, 0xa8, 0xf5,
	0x05, 0x4d, 0xdb, 0x7d, 0x16, 0xe9, 0xbc, 0xf5, 0xb0, 0xaa, 0xce, 0xdb, 0x2c, 0x42, 0xd7, 0xa0,
	0xd2, 0xa3, 0x72, 0x37, 0x8e, 0x74, 0xd6, 0x7a, 0x68, 0x4f, 0xfe, 0xdf, 0x25, 0x28, 0x3f, 0x50,
	0xb7, 0x43, 0x08, 0x4a, 0xc9, 0x2e, 0x8b, 0x6c, 0x44, 0xfd, 0x8c, 0x96, 0x00, 0xed, 0xf5, 0x3b,
	0x34, 0xe5, 0x54, 0x52, 0xd1, 0x26, 0xdd, 0xbe, 0x90, 0x34, 0xb5, 0xa1, 0xaf, 0x0c, 0x35, 0xeb,
	0x46, 0x81, 0x96, 0x61, 0x36, 0x63, 0xce, 0x71, 0x8f, 0x8a, 0x04, 0x13, 0x6a, 0x53, 0xfe, 0x7f,
	0xa8, 0xdb, 0x38, 0x52, 0xa1, 0x55, 0xb8, 0x9a, 0x71, 0x89, 0x3b, 0x5f, 0x51, 0x22, 0x3f, 0x1d,
	0x24, 0x54, 0x57, 0xa0, 0x1e, 0x66, 0xe2, 0x3d, 0x3d, 0xd6, 0x8d, 0x74, 0x52, 0x21, 0xbd, 0xf2,
	0x68, 0x27, 0xa5, 0x43, 0xb7, 0x60, 0x9a, 0xe9, 0x96, 0xb5, 0x05, 0x4d, 0xf7, 0x19, 0xa1, 0x5e,
	0x55, 0x5b, 0x4f, 0x19, 0xe9, 0x96, 0x11, 0xa2, 0x8f, 0x60, 0xd2, 0x9a, 0x25, 0xaa, 0xd2, 0x5e,
	0x6d, 0xce, 0x99, 0x6f, 0xac, 0xdc, 0x6c, 0x8d, 0xa6, 0x74, 0x2b, 0x43, 0xb8, 0xb0, 0xc1, 0x32,
	0xec, 0xbb, 0x09, 0x53, 0x86, 0x7d, 0x24, 0xee, 0xf5, 0x30, 0x8f, 0xbc, 0xba, 0xce, 0x36, 0xa9,
	0x85, 0xeb, 0x46, 0x86, 0x36, 0xe0, 0x7f, 0xdc, 0xf6, 0xd5, 0xa4, 0x13, 0x5e, 0x63, 0x6e, 0x62,
	0xbe, 0xb1, 0x72, 0xeb, 0xb4, 0x7c, 0x39, 0x1a, 0x84, 0xd3, 0x3c, 0x7b, 0x14, 0x2a, 0xa9, 0xc1,
	0x70, 0x40, 0x3b, 0xed, 0x7e, 0xda, 0xf5, 0x26, 0x4d, 0xd2, 0x63, 0xe1, 0x76, 0xda, 0x45, 0xef,
	0x41, 0x69, 0x1f, 0xa7, 0xc2, 0x9b, 0xd2, 0x99, 0x6e, 0x9f, 0x96, 0x49, 0xb3, 0xa2, 0xf5, 0x0c,
	0xa7, 0xe2, 0x43, 0x2e, 0xd3, 0x41, 0xa8, 0x9d, 0x9a, 0x77, 0xa1, 0x7e, 0x2c, 0x42, 0x33, 0x30,
	0xb1, 0x47, 0x07, 0x96, 0x31, 0xea, 0x51, 0xf1, 0x72, 0x1f, 0x77, 0xfb, 0xd4, 0x72, 0xc4, 0x1c,
	0xee, 0xbb, 0xf7, 0x1c, 0xff, 0x4f, 0x07, 0x66, 0x74, 0xc8, 0x4f, 0x98, 0x90, 0x21, 0xfd, 0xba,
	0x4f, 0x85, 0x3c, 0x85, 0x5f, 0xce, 0x79, 0xf9, 0xe5, 0x5e, 0x80, 0x5f, 0x13, 0x17, 0xe1, 0x57,
	0xe9, 0x74, 0x7e, 0xf9, 0x87, 0x70, 0x25, 0x73, 0x3f, 0x91, 0xc4, 0x5c, 0x50, 0x14, 0x40, 0x45,
	0x8d, 0x66, 0x5f, 0xe8, 0x4b, 0x35, 0x56, 0xde, 0x18, 0x56, 0xdb, 0x2c, 0xb0, 0xd6, 0x96, 0x56,
	0x87, 0xd6, 0x0c, 0xad, 0x41, 0x45, 0xb7, 0x41, 0x78, 0xae, 0x6e, 0xcf, 0xf5, 0x33, 0xdb, 0x13,
	0x5a, 0x63, 0xff, 0xf7, 0x12, 0x20, 0x2d, 0x59, 0x4f, 0x29, 0x96, 0xf4, 0x35, 0xac, 0xef, 0x88,
	0xf9, 0xad, 0x14, 0x99, 0xdf, 0xea, 0x65, 0xcd, 0x6f, 0xad, 0xd8, 0xfc, 0xc2, 0x7f, 0x99, 0xdf,
	0x8f, 0xed, 0x68, 0x9a, 0x25, 0xf0, 0xee, 0x99, 0xbd, 0xcf, 0x75, 0xfa, 0xf2, 0xe6, 0x74, 0x01,
	0x66, 0x75, 0xf8, 0x47, 0x54, 0x90, 0x94, 0x75, 0x8e, 0xa9, 0x34, 0xe2, 0xf5, 0xe0, 0x7f, 0x03,
	0x57, 0x4f, 0xd8, 0x5e, 0x06, 0xed, 0x9d, 0xe2, 0xb4, 0xff, 0xcd, 0xb5, 0xb4, 0xdf, 0x4e, 0xa2,
	0x0c, 0xed, 0x47, 0xbd, 0xca, 0x4e, 0xf2, 0xc2, 0xbd, 0x20, 0x2f, 0x46, 0xb4, 0xbc, 0x74, 0x19,
	0x2d, 0x2f, 0x17, 0x68, 0x79, 0xee, 0x96, 0x97, 0xd7, 0xf2, 0x7f, 0x5c, 0xbb, 0x9a, 0xb7, 0x06,
	0x9c, 0xbc, 0x8e, 0xab, 0xe3, 0x73, 0x98, 0x4a, 0xe2, 0xa8, 0x8d, 0x39, 0xa1, 0x42, 0xc6, 0x05,
	0x8b, 0x9d, 0x29, 0x46, 0x6b, 0x33, 0x8e, 0x1e, 0x58, 0xe7, 0x70, 0x32, 0x19, 0x1e, 0x44, 0x73,
	0x0d, 0x1a, 0x19, 0xa5, 0x62, 0x9e, 0xe2, 0xf7, 0x11, 0xf3, 0xd4, 0xb3, 0x92, 0xa9, 0x7a, 0xd8,
	0x52, 0xe8, 0x67, 0x7f, 0xde, 0xf2, 0xf6, 0x11, 0xed, 0xd2, 0x33, 0x79, 0xbb, 0xf2, 0x6b, 0x15,
	0x2a, 0xda, 0x54, 0xa0, 0xef, 0x1d, 0x28, 0xa9, 0xb7, 0x0b, 0x9a, 0x3f, 0x13, 0x78, 0xe6, 0x05,
	0xdb, 0xbc, 0x53, 0xc0, 0xd2, 0xcc, 0xac, 0xbf, 0xf6, 0xf2, 0x17, 0xcf, 0xad, 0x39, 0x2f, 0xff,
	0xf8, 0xeb, 0x07, 0xf7, 0x0e, 0xba, 0x1d, 0xe4, 0xfe, 0x6d, 0x0f, 0xcb, 0x1a, 0xd8, 0x08, 0x81,
	0x19, 0x41, 0xf4, 0xad, 0x03, 0x15, 0xb3, 0x8a, 0xd0, 0x42, 0xf1, 0x7d, 0xd5, 0xbc, 0xfe, 0xca,
	0x46, 0x78, 0x16, 0xb3, 0xe8, 0x18, 0xcc, 0xbd, 0x0c, 0x98, 0x45, 0xbf, 0x28, 0x98, 0xfb, 0xce,
	0x02, 0xfa, 0xc9, 0x81, 0xda, 0xd1, 0x3e, 0x42, 0x8b, 0x67, 0x22, 0x3a, 0xb1, 0xe2, 0x9a, 0x4b,
	0x05, 0xad, 0x2d, 0xc6, 0xf7, 0x33, 0x18, 0xdf, 0x41, 0xad, 0x82, 0x18, 0x83, 0x43, 0xd5, 0xd6,
	0x17, 0xe8, 0x47, 0x07, 0x2a, 0x66, 0x9e, 0xc7, 0xd4, 0x2d, 0x37, 0xf4, 0xe3, 0xea, 0xf6, 0x41,
	0x06, 0xd3, 0x6a, 0xf3, 0x9c, 0x98, 0x54, 0xf9, 0x7e, 0x76, 0xa0, 0xa4, 0x98, 0x3f, 0x86, 0x63,
	0x99, 0xe1, 0x18, 0x07, 0x89, 0x64, 0x20, 0x3d, 0xf7, 0xc3, 0xb1, 0x90, 0xec, 0xaa, 0x11, 0xc1,
	0xe1, 0xab, 0xfb, 0xe7, 0x45, 0x80, 0x89, 0xfe, 0x24, 0x0c, 0xc4, 0x80, 0x93, 0xa5, 0x61, 0xd7,
	0xbf, 0x73, 0xa0, 0x62, 0x66, 0x69, 0x4c, 0x35, 0x73, 0x03, 0x37, 0x0e, 0x7a, 0xae, 0xc3, 0x0b,
	0xe7, 0xac, 0xe6, 0xc3, 0xbb, 0x70, 0x83, 0xc4, 0xbd, 0x0c, 0x9a, 0x84, 0xe5, 0x11, 0x3d, 0x04,
	0x0d, 0x69, 0x53, 0x7d, 0xbc, 0x6e, 0x3a, 0x5f, 0x54, 0xad, 0xb8, 0x53, 0xd1, 0x9f, 0xb3, 0xab,
	0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x36, 0x72, 0xa7, 0x54, 0x0f, 0x00, 0x00,
}
