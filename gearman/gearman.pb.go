// Code generated by protoc-gen-go.
// source: gearman.proto
// DO NOT EDIT!

/*
Package gearman is a generated protocol buffer package.

It is generated from these files:
	gearman.proto

It has these top-level messages:
	Auth
	Metadata
	Operation
*/
package gearman

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import certificate "github.com/appscode/api/certificate/v0.1"
import ci "github.com/appscode/api/ci/v0.1"
import kubernetes "github.com/appscode/api/kubernetes/v0.1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// Next Id: 11
type OperationType int32

const (
	OperationType_UNKNOWN          OperationType = 0
	OperationType_CLUSTER_CREATE   OperationType = 1
	OperationType_CLUSTER_SCALE    OperationType = 2
	OperationType_CLUSTER_DELETE   OperationType = 3
	OperationType_CLUSTER_UPDATE   OperationType = 4
	OperationType_SLAVE_CREATE     OperationType = 5
	OperationType_SLAVE_DELETE     OperationType = 6
	OperationType_CHECK_DNS        OperationType = 7
	OperationType_NAMESPACE_CREATE OperationType = 8
)

var OperationType_name = map[int32]string{
	0: "UNKNOWN",
	1: "CLUSTER_CREATE",
	2: "CLUSTER_SCALE",
	3: "CLUSTER_DELETE",
	4: "CLUSTER_UPDATE",
	5: "SLAVE_CREATE",
	6: "SLAVE_DELETE",
	7: "CHECK_DNS",
	8: "NAMESPACE_CREATE",
}
var OperationType_value = map[string]int32{
	"UNKNOWN":          0,
	"CLUSTER_CREATE":   1,
	"CLUSTER_SCALE":    2,
	"CLUSTER_DELETE":   3,
	"CLUSTER_UPDATE":   4,
	"SLAVE_CREATE":     5,
	"SLAVE_DELETE":     6,
	"CHECK_DNS":        7,
	"NAMESPACE_CREATE": 8,
}

func (x OperationType) String() string {
	return proto.EnumName(OperationType_name, int32(x))
}
func (OperationType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Auth struct {
	Namespace string `protobuf:"bytes,1,opt,name=namespace" json:"namespace,omitempty"`
	Username  string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Token     string `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
}

func (m *Auth) Reset()                    { *m = Auth{} }
func (m *Auth) String() string            { return proto.CompactTextString(m) }
func (*Auth) ProtoMessage()               {}
func (*Auth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Metadata struct {
	KubeProductPhid string `protobuf:"bytes,1,opt,name=kube_product_phid" json:"kube_product_phid,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Next Id: 10
type Operation struct {
	Phid        string        `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Auth        *Auth         `protobuf:"bytes,2,opt,name=auth" json:"auth,omitempty"`
	RequestTime int64         `protobuf:"varint,3,opt,name=request_time" json:"request_time,omitempty"`
	Metadata    *Metadata     `protobuf:"bytes,4,opt,name=metadata" json:"metadata,omitempty"`
	Type        OperationType `protobuf:"varint,5,opt,name=type,enum=gearman.OperationType" json:"type,omitempty"`
	// Types that are valid to be assigned to Request:
	//	*Operation_ClusterCreateRequest
	//	*Operation_ClusterScaleRequest
	//	*Operation_ClusterDeleteRequest
	//	*Operation_ClusterUpdateRequest
	//	*Operation_CiSlaveCreateRequest
	//	*Operation_CiSlaveDeleteRequest
	//	*Operation_DnsCheckRequest
	Request isOperation_Request `protobuf_oneof:"request"`
}

func (m *Operation) Reset()                    { *m = Operation{} }
func (m *Operation) String() string            { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()               {}
func (*Operation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type isOperation_Request interface {
	isOperation_Request()
}

type Operation_ClusterCreateRequest struct {
	ClusterCreateRequest *kubernetes.ClusterCreateRequest `protobuf:"bytes,6,opt,name=cluster_create_request,oneof"`
}
type Operation_ClusterScaleRequest struct {
	ClusterScaleRequest *kubernetes.ClusterScaleRequest `protobuf:"bytes,7,opt,name=cluster_scale_request,oneof"`
}
type Operation_ClusterDeleteRequest struct {
	ClusterDeleteRequest *kubernetes.ClusterDeleteRequest `protobuf:"bytes,8,opt,name=cluster_delete_request,oneof"`
}
type Operation_ClusterUpdateRequest struct {
	ClusterUpdateRequest *kubernetes.ClusterUpdateRequest `protobuf:"bytes,9,opt,name=cluster_update_request,oneof"`
}
type Operation_CiSlaveCreateRequest struct {
	CiSlaveCreateRequest *ci.SlaveCreateRequest `protobuf:"bytes,10,opt,name=ci_slave_create_request,oneof"`
}
type Operation_CiSlaveDeleteRequest struct {
	CiSlaveDeleteRequest *ci.SlaveDeleteRequest `protobuf:"bytes,11,opt,name=ci_slave_delete_request,oneof"`
}
type Operation_DnsCheckRequest struct {
	DnsCheckRequest *certificate.DNSCheckRequest `protobuf:"bytes,12,opt,name=dns_check_request,oneof"`
}

func (*Operation_ClusterCreateRequest) isOperation_Request() {}
func (*Operation_ClusterScaleRequest) isOperation_Request()  {}
func (*Operation_ClusterDeleteRequest) isOperation_Request() {}
func (*Operation_ClusterUpdateRequest) isOperation_Request() {}
func (*Operation_CiSlaveCreateRequest) isOperation_Request() {}
func (*Operation_CiSlaveDeleteRequest) isOperation_Request() {}
func (*Operation_DnsCheckRequest) isOperation_Request()      {}

func (m *Operation) GetRequest() isOperation_Request {
	if m != nil {
		return m.Request
	}
	return nil
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

func (m *Operation) GetClusterCreateRequest() *kubernetes.ClusterCreateRequest {
	if x, ok := m.GetRequest().(*Operation_ClusterCreateRequest); ok {
		return x.ClusterCreateRequest
	}
	return nil
}

func (m *Operation) GetClusterScaleRequest() *kubernetes.ClusterScaleRequest {
	if x, ok := m.GetRequest().(*Operation_ClusterScaleRequest); ok {
		return x.ClusterScaleRequest
	}
	return nil
}

func (m *Operation) GetClusterDeleteRequest() *kubernetes.ClusterDeleteRequest {
	if x, ok := m.GetRequest().(*Operation_ClusterDeleteRequest); ok {
		return x.ClusterDeleteRequest
	}
	return nil
}

func (m *Operation) GetClusterUpdateRequest() *kubernetes.ClusterUpdateRequest {
	if x, ok := m.GetRequest().(*Operation_ClusterUpdateRequest); ok {
		return x.ClusterUpdateRequest
	}
	return nil
}

func (m *Operation) GetCiSlaveCreateRequest() *ci.SlaveCreateRequest {
	if x, ok := m.GetRequest().(*Operation_CiSlaveCreateRequest); ok {
		return x.CiSlaveCreateRequest
	}
	return nil
}

func (m *Operation) GetCiSlaveDeleteRequest() *ci.SlaveDeleteRequest {
	if x, ok := m.GetRequest().(*Operation_CiSlaveDeleteRequest); ok {
		return x.CiSlaveDeleteRequest
	}
	return nil
}

func (m *Operation) GetDnsCheckRequest() *certificate.DNSCheckRequest {
	if x, ok := m.GetRequest().(*Operation_DnsCheckRequest); ok {
		return x.DnsCheckRequest
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Operation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Operation_OneofMarshaler, _Operation_OneofUnmarshaler, _Operation_OneofSizer, []interface{}{
		(*Operation_ClusterCreateRequest)(nil),
		(*Operation_ClusterScaleRequest)(nil),
		(*Operation_ClusterDeleteRequest)(nil),
		(*Operation_ClusterUpdateRequest)(nil),
		(*Operation_CiSlaveCreateRequest)(nil),
		(*Operation_CiSlaveDeleteRequest)(nil),
		(*Operation_DnsCheckRequest)(nil),
	}
}

func _Operation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Operation)
	// request
	switch x := m.Request.(type) {
	case *Operation_ClusterCreateRequest:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClusterCreateRequest); err != nil {
			return err
		}
	case *Operation_ClusterScaleRequest:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClusterScaleRequest); err != nil {
			return err
		}
	case *Operation_ClusterDeleteRequest:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClusterDeleteRequest); err != nil {
			return err
		}
	case *Operation_ClusterUpdateRequest:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ClusterUpdateRequest); err != nil {
			return err
		}
	case *Operation_CiSlaveCreateRequest:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CiSlaveCreateRequest); err != nil {
			return err
		}
	case *Operation_CiSlaveDeleteRequest:
		b.EncodeVarint(11<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CiSlaveDeleteRequest); err != nil {
			return err
		}
	case *Operation_DnsCheckRequest:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.DnsCheckRequest); err != nil {
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
	case 6: // request.cluster_create_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kubernetes.ClusterCreateRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_ClusterCreateRequest{msg}
		return true, err
	case 7: // request.cluster_scale_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kubernetes.ClusterScaleRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_ClusterScaleRequest{msg}
		return true, err
	case 8: // request.cluster_delete_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kubernetes.ClusterDeleteRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_ClusterDeleteRequest{msg}
		return true, err
	case 9: // request.cluster_update_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(kubernetes.ClusterUpdateRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_ClusterUpdateRequest{msg}
		return true, err
	case 10: // request.ci_slave_create_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ci.SlaveCreateRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_CiSlaveCreateRequest{msg}
		return true, err
	case 11: // request.ci_slave_delete_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ci.SlaveDeleteRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_CiSlaveDeleteRequest{msg}
		return true, err
	case 12: // request.dns_check_request
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(certificate.DNSCheckRequest)
		err := b.DecodeMessage(msg)
		m.Request = &Operation_DnsCheckRequest{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Operation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Operation)
	// request
	switch x := m.Request.(type) {
	case *Operation_ClusterCreateRequest:
		s := proto.Size(x.ClusterCreateRequest)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_ClusterScaleRequest:
		s := proto.Size(x.ClusterScaleRequest)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_ClusterDeleteRequest:
		s := proto.Size(x.ClusterDeleteRequest)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_ClusterUpdateRequest:
		s := proto.Size(x.ClusterUpdateRequest)
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_CiSlaveCreateRequest:
		s := proto.Size(x.CiSlaveCreateRequest)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_CiSlaveDeleteRequest:
		s := proto.Size(x.CiSlaveDeleteRequest)
		n += proto.SizeVarint(11<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Operation_DnsCheckRequest:
		s := proto.Size(x.DnsCheckRequest)
		n += proto.SizeVarint(12<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Auth)(nil), "gearman.Auth")
	proto.RegisterType((*Metadata)(nil), "gearman.Metadata")
	proto.RegisterType((*Operation)(nil), "gearman.Operation")
	proto.RegisterEnum("gearman.OperationType", OperationType_name, OperationType_value)
}

var fileDescriptor0 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x93, 0xcd, 0x6e, 0x9b, 0x4e,
	0x14, 0xc5, 0xe3, 0x18, 0xdb, 0x70, 0x6d, 0x22, 0x3c, 0xfa, 0xff, 0x1d, 0x9a, 0x56, 0x6a, 0xe4,
	0x7e, 0xa8, 0xea, 0x82, 0xb4, 0xe9, 0x2a, 0x5d, 0x95, 0x60, 0xa4, 0x48, 0x71, 0x48, 0x64, 0xec,
	0x76, 0x89, 0x26, 0xc3, 0xb4, 0x46, 0xb6, 0x31, 0x85, 0x21, 0x52, 0xdf, 0xab, 0x2f, 0xd4, 0x37,
	0xe9, 0x30, 0x7c, 0x16, 0x25, 0x4b, 0xee, 0x3d, 0xe7, 0x77, 0xcf, 0x5c, 0x66, 0x40, 0xfd, 0x41,
	0x71, 0xbc, 0xc3, 0xa1, 0x11, 0xc5, 0x7b, 0xb6, 0x47, 0x83, 0xe2, 0xf3, 0xe4, 0x2d, 0x8e, 0x82,
	0x33, 0x42, 0x63, 0x16, 0x7c, 0x0f, 0x08, 0x66, 0xf4, 0xec, 0xe1, 0x83, 0xf1, 0xb1, 0x59, 0xc8,
	0x0d, 0x27, 0xc7, 0x42, 0x17, 0xe4, 0xed, 0x64, 0x8b, 0x1f, 0xca, 0xc6, 0x34, 0x6b, 0x6c, 0xd2,
	0x7b, 0x1a, 0x87, 0x94, 0xd1, 0xa4, 0xf0, 0x6f, 0xd3, 0x84, 0xd1, 0x38, 0xc9, 0x35, 0xd3, 0xcf,
	0x20, 0x99, 0x29, 0x5b, 0xa3, 0x31, 0x28, 0x21, 0xde, 0xd1, 0x24, 0xc2, 0x84, 0xea, 0x9d, 0xd3,
	0xce, 0x3b, 0x05, 0x69, 0x20, 0xa7, 0x09, 0xf7, 0xf2, 0xb2, 0x7e, 0x28, 0x2a, 0x2a, 0xf4, 0xd8,
	0x7e, 0x43, 0x43, 0xbd, 0x9b, 0x7d, 0x4e, 0xdf, 0x80, 0x7c, 0x43, 0x19, 0xf6, 0x31, 0xc3, 0xe8,
	0x19, 0x8c, 0xb3, 0x49, 0x1e, 0xa7, 0xfa, 0x29, 0x61, 0x5e, 0xb4, 0x0e, 0xfc, 0x9c, 0x33, 0xfd,
	0x23, 0x81, 0x72, 0x1b, 0xd1, 0x18, 0xb3, 0x60, 0x1f, 0xa2, 0x11, 0x48, 0x75, 0x0f, 0x3d, 0x07,
	0x09, 0xf3, 0xf1, 0x82, 0x3f, 0x3c, 0x57, 0x8d, 0x72, 0x15, 0x22, 0xd3, 0x7f, 0x30, 0x8a, 0xe9,
	0xcf, 0x94, 0x26, 0xcc, 0x63, 0x01, 0x0f, 0x91, 0x4d, 0xed, 0xa2, 0x57, 0x20, 0xef, 0x8a, 0xa9,
	0xba, 0x24, 0x6c, 0xe3, 0xca, 0x56, 0xc5, 0x79, 0x0d, 0x12, 0xfb, 0x15, 0x51, 0xbd, 0xc7, 0x05,
	0x47, 0xe7, 0x93, 0x4a, 0x50, 0xe5, 0x58, 0xf2, 0x2e, 0xba, 0x84, 0x49, 0xb1, 0x0e, 0x8f, 0xc4,
	0x94, 0x6f, 0xd4, 0x2b, 0xe6, 0xe9, 0x7d, 0x01, 0x3e, 0x35, 0xea, 0xed, 0x19, 0x56, 0xae, 0xb4,
	0x84, 0x70, 0x91, 0xeb, 0xae, 0x0e, 0xd0, 0x17, 0xf8, 0xbf, 0x64, 0x24, 0x04, 0x6f, 0x6b, 0xc4,
	0x40, 0x20, 0x5e, 0x3e, 0x82, 0x70, 0x33, 0x5d, 0x4d, 0x68, 0xa4, 0xf0, 0xe9, 0x96, 0x36, 0x52,
	0xc8, 0x4f, 0xa6, 0x98, 0x09, 0xe1, 0xa3, 0x8c, 0x34, 0xf2, 0x9b, 0x27, 0x51, 0x9e, 0x64, 0xac,
	0x84, 0xb0, 0x66, 0x5c, 0xc0, 0x31, 0x09, 0x3c, 0x71, 0x81, 0xda, 0xeb, 0x00, 0x01, 0x99, 0x18,
	0x24, 0x30, 0xdc, 0xac, 0xdf, 0x5e, 0x42, 0xd3, 0xda, 0x3a, 0xc3, 0xb0, 0x65, 0x6d, 0x27, 0xbf,
	0x80, 0xb1, 0x1f, 0x26, 0x1e, 0x59, 0x53, 0xb2, 0xa9, 0x4c, 0x23, 0x61, 0x7a, 0x61, 0x34, 0x2f,
	0xfb, 0xcc, 0x71, 0xad, 0x4c, 0x54, 0x59, 0x2f, 0x15, 0x18, 0x14, 0x86, 0xf7, 0xbf, 0x3b, 0xa0,
	0xfe, 0xfb, 0x6f, 0x87, 0x30, 0x58, 0x39, 0xd7, 0xce, 0xed, 0x37, 0x47, 0x3b, 0x40, 0x08, 0x8e,
	0xac, 0xf9, 0xca, 0x5d, 0xda, 0x0b, 0xcf, 0x5a, 0xd8, 0xe6, 0xd2, 0xd6, 0x3a, 0xfc, 0xc6, 0xab,
	0x65, 0xcd, 0xb5, 0xcc, 0xb9, 0xad, 0x1d, 0x36, 0x65, 0x33, 0x7b, 0x6e, 0x73, 0x59, 0xb7, 0x59,
	0x5b, 0xdd, 0xcd, 0x32, 0xab, 0xc4, 0x5f, 0xc6, 0xc8, 0x9d, 0x9b, 0x5f, 0xed, 0x12, 0xd6, 0xab,
	0x2b, 0x85, 0xaf, 0xcf, 0xdf, 0x8a, 0x62, 0x5d, 0xd9, 0xd6, 0xb5, 0xc7, 0x73, 0x6b, 0x03, 0x7e,
	0x97, 0x35, 0xc7, 0xbc, 0xb1, 0xdd, 0x3b, 0xd3, 0xaa, 0x6c, 0xf2, 0x7d, 0x5f, 0x3c, 0xc2, 0x4f,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x80, 0x80, 0x5e, 0x03, 0x04, 0x00, 0x00,
}
