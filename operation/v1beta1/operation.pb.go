// Code generated by protoc-gen-go. DO NOT EDIT.
// source: operation.proto

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	operation.proto

It has these top-level messages:
	Log
	DescribeRequest
	DescribeResponse
	LogDescribeRequest
	LogDescribeResponse
	Auth
	Metadata
	Operation
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_cloud_v1alpha1 "appscode.com/api/cloud/v1alpha1"
import appscode_namespace_v1beta1 "appscode.com/api/namespace/v1beta1"

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

// Next Id: 18
type OperationType int32

const (
	OperationType_UNKNOWN          OperationType = 0
	OperationType_CLUSTER_APPLY    OperationType = 1
	OperationType_NAMESPACE_CREATE OperationType = 11
)

var OperationType_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "CLUSTER_APPLY",
	11: "NAMESPACE_CREATE",
}
var OperationType_value = map[string]int32{
	"UNKNOWN":          0,
	"CLUSTER_APPLY":    1,
	"NAMESPACE_CREATE": 11,
}

func (x OperationType) String() string {
	return proto.EnumName(OperationType_name, int32(x))
}
func (OperationType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Log struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Message   string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *Log) Reset()                    { *m = Log{} }
func (m *Log) String() string            { return proto.CompactTextString(m) }
func (*Log) ProtoMessage()               {}
func (*Log) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Log) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Log) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Log) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type DescribeRequest struct {
	Phid      string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *DescribeRequest) Reset()                    { *m = DescribeRequest{} }
func (m *DescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*DescribeRequest) ProtoMessage()               {}
func (*DescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DescribeRequest) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *DescribeRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type DescribeResponse struct {
	Op   *Operation `protobuf:"bytes,1,opt,name=op" json:"op,omitempty"`
	Logs []*Log     `protobuf:"bytes,2,rep,name=logs" json:"logs,omitempty"`
}

func (m *DescribeResponse) Reset()                    { *m = DescribeResponse{} }
func (m *DescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*DescribeResponse) ProtoMessage()               {}
func (*DescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DescribeResponse) GetOp() *Operation {
	if m != nil {
		return m.Op
	}
	return nil
}

func (m *DescribeResponse) GetLogs() []*Log {
	if m != nil {
		return m.Logs
	}
	return nil
}

type LogDescribeRequest struct {
	Phid  string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	LogId string `protobuf:"bytes,2,opt,name=log_id,json=logId" json:"log_id,omitempty"`
}

func (m *LogDescribeRequest) Reset()                    { *m = LogDescribeRequest{} }
func (m *LogDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*LogDescribeRequest) ProtoMessage()               {}
func (*LogDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LogDescribeRequest) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *LogDescribeRequest) GetLogId() string {
	if m != nil {
		return m.LogId
	}
	return ""
}

type LogDescribeResponse struct {
	Log *Log `protobuf:"bytes,1,opt,name=log" json:"log,omitempty"`
}

func (m *LogDescribeResponse) Reset()                    { *m = LogDescribeResponse{} }
func (m *LogDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*LogDescribeResponse) ProtoMessage()               {}
func (*LogDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *LogDescribeResponse) GetLog() *Log {
	if m != nil {
		return m.Log
	}
	return nil
}

type Auth struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Secret    string `protobuf:"bytes,3,opt,name=secret" json:"secret,omitempty"`
	AuthType  string `protobuf:"bytes,4,opt,name=auth_type,json=authType" json:"auth_type,omitempty"`
}

func (m *Auth) Reset()                    { *m = Auth{} }
func (m *Auth) String() string            { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()               {}
func (*Auth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Auth) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Auth) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Auth) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *Auth) GetAuthType() string {
	if m != nil {
		return m.AuthType
	}
	return ""
}

// Metadata holds other on request or operation specific data
// that could be used inside that operation.
// An welldefined message instead of a map is used
// so that the data fields can be explicitly defined with
// its own data type. Resolves data convertions.
type Metadata struct {
	// Contains PurchasePHID is this is a purchase request.
	PurchasePhids []string `protobuf:"bytes,1,rep,name=purchase_phids,json=purchasePhids" json:"purchase_phids,omitempty"`
	// PHID of the user who requested this operation.
	AuthorPhid string `protobuf:"bytes,2,opt,name=author_phid,json=authorPhid" json:"author_phid,omitempty"`
	AuthorName string `protobuf:"bytes,3,opt,name=author_name,json=authorName" json:"author_name,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Metadata) GetPurchasePhids() []string {
	if m != nil {
		return m.PurchasePhids
	}
	return nil
}

func (m *Metadata) GetAuthorPhid() string {
	if m != nil {
		return m.AuthorPhid
	}
	return ""
}

func (m *Metadata) GetAuthorName() string {
	if m != nil {
		return m.AuthorName
	}
	return ""
}

// Next Id: 22
type Operation struct {
	// Types that are valid to be assigned to Request:
	//	*Operation_ClusterApplyRequest
	//	*Operation_NamespaceCreateRequest
	Request  isOperation_Request `protobuf_oneof:"request"`
	Type     OperationType       `protobuf:"varint,14,opt,name=type,enum=appscode.operation.v1beta1.OperationType" json:"type,omitempty"`
	Phid     string              `protobuf:"bytes,15,opt,name=phid" json:"phid,omitempty"`
	Auth     *Auth               `protobuf:"bytes,16,opt,name=auth" json:"auth,omitempty"`
	Metadata *Metadata           `protobuf:"bytes,17,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *Operation) Reset()                    { *m = Operation{} }
func (m *Operation) String() string            { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()               {}
func (*Operation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isOperation_Request interface {
	isOperation_Request()
}

type Operation_ClusterApplyRequest struct {
	ClusterApplyRequest *appscode_cloud_v1alpha1.ClusterApplyRequest `protobuf:"bytes,1,opt,name=cluster_apply_request,json=clusterApplyRequest,oneof"`
}
type Operation_NamespaceCreateRequest struct {
	NamespaceCreateRequest *appscode_namespace_v1beta1.CreateRequest `protobuf:"bytes,11,opt,name=namespace_create_request,json=namespaceCreateRequest,oneof"`
}

func (*Operation_ClusterApplyRequest) isOperation_Request()    {}
func (*Operation_NamespaceCreateRequest) isOperation_Request() {}

func (m *Operation) GetRequest() isOperation_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Operation) GetClusterApplyRequest() *appscode_cloud_v1alpha1.ClusterApplyRequest {
	if x, ok := m.GetRequest().(*Operation_ClusterApplyRequest); ok {
		return x.ClusterApplyRequest
	}
	return nil
}

func (m *Operation) GetNamespaceCreateRequest() *appscode_namespace_v1beta1.CreateRequest {
	if x, ok := m.GetRequest().(*Operation_NamespaceCreateRequest); ok {
		return x.NamespaceCreateRequest
	}
	return nil
}

func (m *Operation) GetType() OperationType {
	if m != nil {
		return m.Type
	}
	return OperationType_UNKNOWN
}

func (m *Operation) GetPhid() string {
	if m != nil {
		return m.Phid
	}
	return ""
}

func (m *Operation) GetAuth() *Auth {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *Operation) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Operation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Operation_OneofMarshaler, _Operation_OneofUnmarshaler, _Operation_OneofSizer, []interface{}{
		(*Operation_ClusterApplyRequest)(nil),
		(*Operation_NamespaceCreateRequest)(nil),
	}
}

func _Operation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Operation)
	// request
	switch x := m.Request.(type) {
	case *Operation_ClusterApplyRequest:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClusterApplyRequest); err != nil {
			return err
		}
	case *Operation_NamespaceCreateRequest:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NamespaceCreateRequest); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Operation.Request has unexpected type %T", x)
	}
	return nil
}

func _Operation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Operation)
	switch tag {
	case 1: // request.cluster_apply_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(appscode_cloud_v1alpha1.ClusterApplyRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_ClusterApplyRequest{msg}
		return true, err
	case 11: // request.namespace_create_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(appscode_namespace_v1beta1.CreateRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_NamespaceCreateRequest{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Operation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Operation)
	// request
	switch x := m.Request.(type) {
	case *Operation_ClusterApplyRequest:
		s := proto.Size(x.ClusterApplyRequest)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_NamespaceCreateRequest:
		s := proto.Size(x.NamespaceCreateRequest)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Log)(nil), "appscode.operation.v1beta1.Log")
	proto.RegisterType((*DescribeRequest)(nil), "appscode.operation.v1beta1.DescribeRequest")
	proto.RegisterType((*DescribeResponse)(nil), "appscode.operation.v1beta1.DescribeResponse")
	proto.RegisterType((*LogDescribeRequest)(nil), "appscode.operation.v1beta1.LogDescribeRequest")
	proto.RegisterType((*LogDescribeResponse)(nil), "appscode.operation.v1beta1.LogDescribeResponse")
	proto.RegisterType((*Auth)(nil), "appscode.operation.v1beta1.Auth")
	proto.RegisterType((*Metadata)(nil), "appscode.operation.v1beta1.Metadata")
	proto.RegisterType((*Operation)(nil), "appscode.operation.v1beta1.Operation")
	proto.RegisterEnum("appscode.operation.v1beta1.OperationType", OperationType_name, OperationType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Operations service

type OperationsClient interface {
	Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*DescribeResponse, error)
	DescribeLog(ctx context.Context, in *LogDescribeRequest, opts ...grpc.CallOption) (*LogDescribeResponse, error)
}

type operationsClient struct {
	cc *grpc.ClientConn
}

func NewOperationsClient(cc *grpc.ClientConn) OperationsClient {
	return &operationsClient{cc}
}

func (c *operationsClient) Describe(ctx context.Context, in *DescribeRequest, opts ...grpc.CallOption) (*DescribeResponse, error) {
	out := new(DescribeResponse)
	err := grpc.Invoke(ctx, "/appscode.operation.v1beta1.Operations/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationsClient) DescribeLog(ctx context.Context, in *LogDescribeRequest, opts ...grpc.CallOption) (*LogDescribeResponse, error) {
	out := new(LogDescribeResponse)
	err := grpc.Invoke(ctx, "/appscode.operation.v1beta1.Operations/DescribeLog", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Operations service

type OperationsServer interface {
	Describe(context.Context, *DescribeRequest) (*DescribeResponse, error)
	DescribeLog(context.Context, *LogDescribeRequest) (*LogDescribeResponse, error)
}

func RegisterOperationsServer(s *grpc.Server, srv OperationsServer) {
	s.RegisterService(&_Operations_serviceDesc, srv)
}

func _Operations_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.operation.v1beta1.Operations/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).Describe(ctx, req.(*DescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Operations_DescribeLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationsServer).DescribeLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.operation.v1beta1.Operations/DescribeLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationsServer).DescribeLog(ctx, req.(*LogDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Operations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.operation.v1beta1.Operations",
	HandlerType: (*OperationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Describe",
			Handler:    _Operations_Describe_Handler,
		},
		{
			MethodName: "DescribeLog",
			Handler:    _Operations_DescribeLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "operation.proto",
}

func init() { proto.RegisterFile("operation.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 780 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x8e, 0xdb, 0x44,
	0x18, 0xad, 0xed, 0xb0, 0x49, 0xbe, 0x68, 0xb3, 0xe9, 0x94, 0x56, 0x56, 0xa8, 0x68, 0x64, 0x51,
	0x29, 0xd0, 0x12, 0x2b, 0x5b, 0x10, 0xe2, 0x02, 0x95, 0x6c, 0xb0, 0x28, 0x22, 0x9b, 0x46, 0xde,
	0xad, 0x10, 0xdc, 0x58, 0x13, 0x67, 0xe4, 0x18, 0xd9, 0x99, 0xc1, 0x33, 0xae, 0xb4, 0xaa, 0x40,
	0xa2, 0xaf, 0xc0, 0x6b, 0xf4, 0x96, 0x27, 0xe0, 0x11, 0x78, 0x05, 0x1e, 0x04, 0xcd, 0xd8, 0x1e,
	0x27, 0xbb, 0xda, 0xec, 0xaa, 0x77, 0x9e, 0x6f, 0xce, 0x39, 0x73, 0xbe, 0x9f, 0x19, 0xc3, 0x11,
	0x65, 0x24, 0xc3, 0x22, 0xa6, 0x9b, 0x11, 0xcb, 0xa8, 0xa0, 0xa8, 0x8f, 0x19, 0xe3, 0x21, 0x5d,
	0x91, 0x51, 0xbd, 0xf3, 0x7a, 0xbc, 0x24, 0x02, 0x8f, 0xfb, 0x0f, 0x23, 0x4a, 0xa3, 0x84, 0xb8,
	0x98, 0xc5, 0x2e, 0xde, 0x6c, 0xa8, 0x50, 0xdb, 0xbc, 0x60, 0xf6, 0x3f, 0xae, 0x98, 0xd7, 0xec,
	0x3f, 0xd1, 0xca, 0x21, 0x4d, 0x15, 0x26, 0x4c, 0x68, 0xbe, 0x72, 0x5f, 0x8f, 0x71, 0xc2, 0xd6,
	0x78, 0x5c, 0x2c, 0x4b, 0xf0, 0xe7, 0x57, 0xc0, 0x1b, 0x9c, 0x12, 0xce, 0x70, 0x48, 0xdc, 0xd2,
	0x8e, 0x2b, 0x08, 0x4e, 0x0b, 0xb8, 0x73, 0x0a, 0xd6, 0x8c, 0x46, 0xa8, 0x0b, 0x66, 0xbc, 0xb2,
	0x8d, 0x81, 0x31, 0x6c, 0xfb, 0x66, 0xbc, 0x42, 0x0f, 0xa1, 0x2d, 0xe2, 0x94, 0x70, 0x81, 0x53,
	0x66, 0x9b, 0x03, 0x63, 0x68, 0xf9, 0x75, 0x00, 0xd9, 0xd0, 0x4c, 0x09, 0xe7, 0x38, 0x22, 0xb6,
	0xa5, 0x28, 0xd5, 0xd2, 0x99, 0xc2, 0xd1, 0x77, 0x84, 0x87, 0x59, 0xbc, 0x24, 0x3e, 0xf9, 0x2d,
	0x27, 0x5c, 0x20, 0x04, 0x0d, 0xb6, 0xd6, 0xe2, 0xea, 0x7b, 0xbf, 0xbc, 0xf3, 0x07, 0xf4, 0x6a,
	0x11, 0xce, 0xe8, 0x86, 0x13, 0xf4, 0x25, 0x98, 0x94, 0x29, 0x8d, 0xce, 0xf1, 0xe3, 0xd1, 0xf5,
	0xa5, 0x1e, 0xbd, 0xac, 0x22, 0xbe, 0x49, 0x19, 0x7a, 0x06, 0x8d, 0x84, 0x46, 0xdc, 0x36, 0x07,
	0xd6, 0xb0, 0x73, 0xfc, 0x68, 0x1f, 0x71, 0x46, 0x23, 0x5f, 0x81, 0x9d, 0xe7, 0x80, 0x66, 0x34,
	0xba, 0x4d, 0x1e, 0xf7, 0xe1, 0x20, 0xa1, 0x51, 0x10, 0xaf, 0x54, 0x12, 0x6d, 0xff, 0x83, 0x84,
	0x46, 0x3f, 0xac, 0x9c, 0x17, 0x70, 0x6f, 0x47, 0xa0, 0xcc, 0x61, 0x0c, 0x56, 0x42, 0xa3, 0x32,
	0x89, 0x1b, 0xbd, 0x48, 0xac, 0x93, 0x43, 0x63, 0x92, 0x8b, 0xb5, 0x2c, 0x98, 0x6e, 0x63, 0xe9,
	0xa0, 0x0e, 0xa0, 0x3e, 0xb4, 0x72, 0x4e, 0x32, 0x19, 0x28, 0x8d, 0xe8, 0x35, 0x7a, 0x00, 0x07,
	0x9c, 0x84, 0x19, 0x11, 0x65, 0xab, 0xca, 0x15, 0xfa, 0x08, 0xda, 0x38, 0x17, 0xeb, 0x40, 0x5c,
	0x30, 0x62, 0x37, 0x0a, 0x92, 0x0c, 0x9c, 0x5f, 0x30, 0xe2, 0x70, 0x68, 0x9d, 0x12, 0x81, 0x57,
	0x58, 0x60, 0xf4, 0x18, 0xba, 0x2c, 0xcf, 0xc2, 0x35, 0xe6, 0x24, 0x90, 0x49, 0x73, 0xdb, 0x18,
	0x58, 0xc3, 0xb6, 0x7f, 0x58, 0x45, 0x17, 0x32, 0x88, 0x1e, 0x41, 0x47, 0xd2, 0x69, 0xa6, 0x40,
	0xa5, 0x0d, 0x28, 0x42, 0x12, 0xb1, 0x05, 0x50, 0x3e, 0xad, 0x6d, 0xc0, 0x1c, 0xa7, 0xc4, 0x79,
	0x67, 0x41, 0x5b, 0x77, 0x0f, 0x2d, 0xe1, 0x7e, 0x98, 0xe4, 0x5c, 0x90, 0x2c, 0xc0, 0x8c, 0x25,
	0x17, 0x41, 0x56, 0xf4, 0xa1, 0x2c, 0xdf, 0xd3, 0xba, 0x7c, 0xc5, 0xf4, 0x57, 0x97, 0x61, 0x34,
	0x2d, 0x58, 0x13, 0x49, 0x2a, 0x7b, 0xf7, 0xe2, 0x8e, 0x7f, 0x2f, 0xbc, 0x1a, 0x46, 0x04, 0x6c,
	0x5d, 0xc4, 0x20, 0xcc, 0x08, 0x16, 0x44, 0x1f, 0xd3, 0x51, 0xc7, 0x7c, 0x5a, 0x1f, 0xa3, 0x91,
	0xba, 0x4b, 0x53, 0xc5, 0xa8, 0xcf, 0x78, 0xa0, 0x21, 0x3b, 0x3b, 0xe8, 0x1b, 0x68, 0xa8, 0x2a,
	0x77, 0x07, 0xc6, 0xb0, 0xbb, 0x2d, 0xb9, 0x67, 0x7a, 0x65, 0x1b, 0x7c, 0x45, 0xd3, 0x83, 0x77,
	0xb4, 0x35, 0x78, 0x5f, 0x40, 0x43, 0x56, 0xce, 0xee, 0x29, 0x97, 0x83, 0x7d, 0x92, 0x72, 0x7e,
	0x7c, 0x85, 0x46, 0xdf, 0x42, 0x2b, 0x2d, 0xdb, 0x6a, 0xdf, 0x55, 0xcc, 0x4f, 0xf6, 0x31, 0xab,
	0x11, 0xf0, 0x35, 0xeb, 0xa4, 0x0d, 0xcd, 0xb2, 0x40, 0x9f, 0x79, 0x70, 0xb8, 0xe3, 0x16, 0x75,
	0xa0, 0xf9, 0x6a, 0xfe, 0xe3, 0xfc, 0xe5, 0x4f, 0xf3, 0xde, 0x1d, 0x74, 0x17, 0x0e, 0xa7, 0xb3,
	0x57, 0x67, 0xe7, 0x9e, 0x1f, 0x4c, 0x16, 0x8b, 0xd9, 0xcf, 0x3d, 0x03, 0x7d, 0x08, 0xbd, 0xf9,
	0xe4, 0xd4, 0x3b, 0x5b, 0x4c, 0xa6, 0x5e, 0x30, 0xf5, 0xbd, 0xc9, 0xb9, 0xd7, 0xeb, 0x1c, 0xff,
	0x69, 0x01, 0x68, 0x1d, 0x8e, 0xde, 0x19, 0xd0, 0xaa, 0x2e, 0x0e, 0x7a, 0xb2, 0xcf, 0xdd, 0xa5,
	0xfb, 0xd9, 0x7f, 0x7a, 0x3b, 0x70, 0x71, 0x17, 0x1d, 0xef, 0xed, 0xdf, 0xb6, 0xd9, 0x32, 0xde,
	0xfe, 0xfb, 0xdf, 0x5f, 0xe6, 0xd7, 0xe8, 0x2b, 0x37, 0xd8, 0x79, 0x83, 0xb5, 0x80, 0x7e, 0x32,
	0x75, 0x84, 0xbb, 0x6f, 0x64, 0x07, 0x7e, 0x77, 0x7f, 0xe5, 0x74, 0x83, 0xfe, 0x31, 0xa0, 0x53,
	0x69, 0xcb, 0x77, 0x74, 0x74, 0xc3, 0xad, 0xbe, 0x6c, 0xda, 0xbd, 0x35, 0xbe, 0xf4, 0x7d, 0xb6,
	0xe5, 0xfb, 0x7b, 0xe4, 0xbd, 0x87, 0x6f, 0xf9, 0xc4, 0xb9, 0x6f, 0x8a, 0x67, 0xab, 0xc8, 0xe2,
	0xe4, 0x39, 0x38, 0x21, 0x4d, 0x6b, 0x2b, 0x98, 0xc5, 0x57, 0xed, 0x9c, 0x74, 0x75, 0x9b, 0x16,
	0xf2, 0xd7, 0xb1, 0x30, 0x7e, 0x69, 0x96, 0x5b, 0xcb, 0x03, 0xf5, 0x33, 0x79, 0xf6, 0x7f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xfe, 0x34, 0xbe, 0xd0, 0x15, 0x07, 0x00, 0x00,
}
