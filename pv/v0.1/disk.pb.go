// Code generated by protoc-gen-go.
// source: disk.proto
// DO NOT EDIT!

/*
Package pv is a generated protocol buffer package.

It is generated from these files:
	disk.proto
	pvc.proto
	pv.proto

It has these top-level messages:
	DiskCreateRequest
	DiskDeleteRequest
	DiskListRequest
	DiskListResponse
	Result
	Disk
	PV
	PVC
	DiskDescribeRequest
	DiskDescribeResponse
	PVCRegisterRequest
	PVCUnregisterRequest
	PVCListRequest
	PVCInfo
	PVCListResponse
	PVRegisterRequest
	PVUnregisterRequest
	PVListRequest
	PVInfo
	PVListResponse
*/
package pv

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

type DiskCreateRequest struct {
	Cluster  string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Zone     string `protobuf:"bytes,3,opt,name=zone" json:"zone,omitempty"`
	DiskType string `protobuf:"bytes,4,opt,name=disk_type,json=diskType" json:"disk_type,omitempty"`
	Size     int64  `protobuf:"varint,5,opt,name=size" json:"size,omitempty"`
}

func (m *DiskCreateRequest) Reset()                    { *m = DiskCreateRequest{} }
func (m *DiskCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*DiskCreateRequest) ProtoMessage()               {}
func (*DiskCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type DiskDeleteRequest struct {
	Cluster    string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Identifier string `protobuf:"bytes,2,opt,name=identifier" json:"identifier,omitempty"`
}

func (m *DiskDeleteRequest) Reset()                    { *m = DiskDeleteRequest{} }
func (m *DiskDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DiskDeleteRequest) ProtoMessage()               {}
func (*DiskDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type DiskListRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
}

func (m *DiskListRequest) Reset()                    { *m = DiskListRequest{} }
func (m *DiskListRequest) String() string            { return proto.CompactTextString(m) }
func (*DiskListRequest) ProtoMessage()               {}
func (*DiskListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type DiskListResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Kube   string         `protobuf:"bytes,2,opt,name=kube" json:"kube,omitempty"`
	Result []*Result      `protobuf:"bytes,3,rep,name=result" json:"result,omitempty"`
}

func (m *DiskListResponse) Reset()                    { *m = DiskListResponse{} }
func (m *DiskListResponse) String() string            { return proto.CompactTextString(m) }
func (*DiskListResponse) ProtoMessage()               {}
func (*DiskListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DiskListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DiskListResponse) GetResult() []*Result {
	if m != nil {
		return m.Result
	}
	return nil
}

type Result struct {
	Disk *Disk `protobuf:"bytes,1,opt,name=disk" json:"disk,omitempty"`
	Pv   *PV   `protobuf:"bytes,2,opt,name=pv" json:"pv,omitempty"`
	Pvc  *PVC  `protobuf:"bytes,3,opt,name=pvc" json:"pvc,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Result) GetDisk() *Disk {
	if m != nil {
		return m.Disk
	}
	return nil
}

func (m *Result) GetPv() *PV {
	if m != nil {
		return m.Pv
	}
	return nil
}

func (m *Result) GetPvc() *PVC {
	if m != nil {
		return m.Pvc
	}
	return nil
}

type Disk struct {
	Name     string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id       string   `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Plugin   string   `protobuf:"bytes,3,opt,name=plugin" json:"plugin,omitempty"`
	Size     int64    `protobuf:"varint,4,opt,name=size" json:"size,omitempty"`
	Type     string   `protobuf:"bytes,5,opt,name=type" json:"type,omitempty"`
	Zone     string   `protobuf:"bytes,6,opt,name=zone" json:"zone,omitempty"`
	Status   string   `protobuf:"bytes,7,opt,name=status" json:"status,omitempty"`
	Users    []string `protobuf:"bytes,8,rep,name=users" json:"users,omitempty"`
	Kind     string   `protobuf:"bytes,9,opt,name=kind" json:"kind,omitempty"`
	Endpoint string   `protobuf:"bytes,10,opt,name=endpoint" json:"endpoint,omitempty"`
}

func (m *Disk) Reset()                    { *m = Disk{} }
func (m *Disk) String() string            { return proto.CompactTextString(m) }
func (*Disk) ProtoMessage()               {}
func (*Disk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type PV struct {
	Name   string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Size   int64    `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Status string   `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	Claim  string   `protobuf:"bytes,4,opt,name=claim" json:"claim,omitempty"`
	Users  []string `protobuf:"bytes,5,rep,name=users" json:"users,omitempty"`
}

func (m *PV) Reset()                    { *m = PV{} }
func (m *PV) String() string            { return proto.CompactTextString(m) }
func (*PV) ProtoMessage()               {}
func (*PV) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type PVC struct {
	Name      string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Size      int64    `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Namespace string   `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	Status    string   `protobuf:"bytes,4,opt,name=status" json:"status,omitempty"`
	Users     []string `protobuf:"bytes,5,rep,name=users" json:"users,omitempty"`
}

func (m *PVC) Reset()                    { *m = PVC{} }
func (m *PVC) String() string            { return proto.CompactTextString(m) }
func (*PVC) ProtoMessage()               {}
func (*PVC) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type DiskDescribeRequest struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster" json:"cluster,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Plugin  string `protobuf:"bytes,3,opt,name=plugin" json:"plugin,omitempty"`
}

func (m *DiskDescribeRequest) Reset()                    { *m = DiskDescribeRequest{} }
func (m *DiskDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*DiskDescribeRequest) ProtoMessage()               {}
func (*DiskDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type DiskDescribeResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Result *Result        `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
}

func (m *DiskDescribeResponse) Reset()                    { *m = DiskDescribeResponse{} }
func (m *DiskDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*DiskDescribeResponse) ProtoMessage()               {}
func (*DiskDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DiskDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DiskDescribeResponse) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*DiskCreateRequest)(nil), "pv.DiskCreateRequest")
	proto.RegisterType((*DiskDeleteRequest)(nil), "pv.DiskDeleteRequest")
	proto.RegisterType((*DiskListRequest)(nil), "pv.DiskListRequest")
	proto.RegisterType((*DiskListResponse)(nil), "pv.DiskListResponse")
	proto.RegisterType((*Result)(nil), "pv.Result")
	proto.RegisterType((*Disk)(nil), "pv.Disk")
	proto.RegisterType((*PV)(nil), "pv.PV")
	proto.RegisterType((*PVC)(nil), "pv.PVC")
	proto.RegisterType((*DiskDescribeRequest)(nil), "pv.DiskDescribeRequest")
	proto.RegisterType((*DiskDescribeResponse)(nil), "pv.DiskDescribeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Disks service

type DisksClient interface {
	List(ctx context.Context, in *DiskListRequest, opts ...grpc.CallOption) (*DiskListResponse, error)
	Describe(ctx context.Context, in *DiskDescribeRequest, opts ...grpc.CallOption) (*DiskDescribeResponse, error)
	Create(ctx context.Context, in *DiskCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *DiskDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type disksClient struct {
	cc *grpc.ClientConn
}

func NewDisksClient(cc *grpc.ClientConn) DisksClient {
	return &disksClient{cc}
}

func (c *disksClient) List(ctx context.Context, in *DiskListRequest, opts ...grpc.CallOption) (*DiskListResponse, error) {
	out := new(DiskListResponse)
	err := grpc.Invoke(ctx, "/pv.Disks/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disksClient) Describe(ctx context.Context, in *DiskDescribeRequest, opts ...grpc.CallOption) (*DiskDescribeResponse, error) {
	out := new(DiskDescribeResponse)
	err := grpc.Invoke(ctx, "/pv.Disks/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disksClient) Create(ctx context.Context, in *DiskCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/pv.Disks/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disksClient) Delete(ctx context.Context, in *DiskDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/pv.Disks/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Disks service

type DisksServer interface {
	List(context.Context, *DiskListRequest) (*DiskListResponse, error)
	Describe(context.Context, *DiskDescribeRequest) (*DiskDescribeResponse, error)
	Create(context.Context, *DiskCreateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *DiskDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterDisksServer(s *grpc.Server, srv DisksServer) {
	s.RegisterService(&_Disks_serviceDesc, srv)
}

func _Disks_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DiskListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DisksServer).List(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Disks_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DiskDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DisksServer).Describe(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Disks_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DiskCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DisksServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Disks_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DiskDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DisksServer).Delete(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Disks_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pv.Disks",
	HandlerType: (*DisksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Disks_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Disks_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Disks_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Disks_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 665 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x4e, 0xd4, 0x40,
	0x14, 0x4e, 0x7f, 0xb6, 0xec, 0x1e, 0x12, 0xd4, 0x01, 0xb1, 0xd6, 0x0d, 0x31, 0xd5, 0x18, 0x04,
	0xd3, 0xca, 0xea, 0x95, 0xb7, 0x70, 0xa9, 0x09, 0xa9, 0x06, 0x63, 0xbc, 0x30, 0xdd, 0xed, 0x88,
	0x13, 0x4a, 0x3b, 0x76, 0xa6, 0x9b, 0x00, 0xe1, 0xc6, 0x0b, 0x5f, 0xc0, 0x77, 0xf0, 0x85, 0xbc,
	0xf5, 0xd2, 0x07, 0x71, 0xfe, 0xba, 0xed, 0x22, 0xe0, 0xea, 0x0d, 0xcc, 0x7c, 0x67, 0x7a, 0xbe,
	0x6f, 0xbe, 0x73, 0xe6, 0x2c, 0x40, 0x46, 0xd8, 0x51, 0x44, 0xab, 0x92, 0x97, 0xc8, 0xa6, 0xd3,
	0x60, 0x78, 0x58, 0x96, 0x87, 0x39, 0x8e, 0x53, 0x4a, 0xe2, 0xb4, 0x28, 0x4a, 0x9e, 0x72, 0x52,
	0x16, 0x4c, 0x9f, 0x08, 0xd6, 0x25, 0x9c, 0xf1, 0x13, 0x8a, 0x59, 0xac, 0xfe, 0x6a, 0x3c, 0xfc,
	0x6a, 0xc1, 0xad, 0x3d, 0x91, 0x68, 0xb7, 0xc2, 0x29, 0xc7, 0x09, 0xfe, 0x5c, 0x63, 0xc6, 0x91,
	0x0f, 0x4b, 0x93, 0xbc, 0x66, 0x1c, 0x57, 0xbe, 0x75, 0xdf, 0xda, 0x1c, 0x24, 0xcd, 0x16, 0x21,
	0x70, 0x8b, 0xf4, 0x18, 0xfb, 0xb6, 0x82, 0xd5, 0x5a, 0x62, 0xa7, 0x65, 0x81, 0x7d, 0x47, 0x63,
	0x72, 0x8d, 0xee, 0xc1, 0x40, 0xea, 0xfb, 0x20, 0xb9, 0x7c, 0x57, 0x05, 0xfa, 0x12, 0x78, 0x23,
	0xf6, 0xf2, 0x03, 0x46, 0x4e, 0xb1, 0xdf, 0x13, 0xb8, 0x93, 0xa8, 0x75, 0xf8, 0x4a, 0xeb, 0xd8,
	0xc3, 0x39, 0x5e, 0x44, 0xc7, 0x06, 0x00, 0xc9, 0x70, 0xc1, 0xc9, 0x47, 0x22, 0x82, 0x5a, 0x4d,
	0x07, 0x09, 0xb7, 0xe1, 0x86, 0x4c, 0xf7, 0x92, 0x30, 0xfe, 0xd7, 0x64, 0x61, 0x05, 0x37, 0xdb,
	0xc3, 0x8c, 0x0a, 0xd7, 0x30, 0x7a, 0x04, 0x1e, 0x13, 0x16, 0xd6, 0x4c, 0x1d, 0x5e, 0x1e, 0xad,
	0x44, 0xda, 0xbd, 0xe8, 0xb5, 0x42, 0x13, 0x13, 0x95, 0x77, 0x39, 0xaa, 0xc7, 0x33, 0x43, 0xe4,
	0x1a, 0x85, 0xe0, 0x55, 0x98, 0xd5, 0x39, 0x17, 0x96, 0x38, 0xe2, 0x5b, 0x88, 0xe8, 0x34, 0x4a,
	0x14, 0x92, 0x98, 0x48, 0xf8, 0x0e, 0x3c, 0x8d, 0xa0, 0x21, 0xb8, 0xd2, 0x19, 0xc3, 0xd3, 0x97,
	0x67, 0xa5, 0x9a, 0x44, 0xa1, 0x68, 0x1d, 0x44, 0x71, 0x55, 0xf6, 0xe5, 0x91, 0x27, 0x63, 0xfb,
	0x07, 0x89, 0x40, 0xd0, 0x5d, 0x70, 0xe8, 0x74, 0xa2, 0x3c, 0x5f, 0x1e, 0x2d, 0xe9, 0xc0, 0x6e,
	0x22, 0xb1, 0xf0, 0xa7, 0x05, 0xae, 0xcc, 0x30, 0x2b, 0x96, 0xd5, 0x29, 0xd6, 0x0a, 0xd8, 0x24,
	0x33, 0x6a, 0xc5, 0x4a, 0xe4, 0xf7, 0x68, 0x5e, 0x1f, 0x92, 0xc2, 0x94, 0xcf, 0xec, 0x66, 0x35,
	0x72, 0xdb, 0x1a, 0x49, 0x4c, 0xd5, 0xb3, 0xa7, 0xf3, 0x71, 0x53, 0x4b, 0x55, 0x7c, 0xaf, 0x53,
	0xfc, 0xf5, 0x99, 0x77, 0x4b, 0x3a, 0xa7, 0xf1, 0x6a, 0x0d, 0x7a, 0x35, 0xc3, 0x15, 0xf3, 0xfb,
	0xc2, 0x96, 0x41, 0xa2, 0x37, 0xca, 0x41, 0x52, 0x64, 0xfe, 0xc0, 0x38, 0x28, 0xd6, 0x28, 0x80,
	0x3e, 0x2e, 0x32, 0x5a, 0x92, 0x82, 0xfb, 0xa0, 0xbb, 0xa7, 0xd9, 0x87, 0x14, 0xec, 0xfd, 0x83,
	0x4b, 0xef, 0xd6, 0x68, 0xb6, 0x3b, 0x9a, 0x5b, 0x2d, 0xce, 0x45, 0x2d, 0x93, 0x3c, 0x25, 0xc7,
	0xa6, 0x39, 0xf5, 0xa6, 0x55, 0xd8, 0xeb, 0x28, 0x0c, 0x4f, 0xc0, 0x11, 0xe6, 0x2e, 0x4c, 0x39,
	0x84, 0x81, 0x8c, 0x31, 0x9a, 0x4e, 0x9a, 0x47, 0xd1, 0x02, 0x1d, 0x41, 0xee, 0xe5, 0xe6, 0xcc,
	0x51, 0xbf, 0x87, 0x55, 0xfd, 0x2c, 0xd8, 0xa4, 0x22, 0xe3, 0xff, 0x7c, 0xa0, 0x57, 0xd4, 0x38,
	0x1c, 0xc3, 0xda, 0x7c, 0xf2, 0x7f, 0xec, 0xfd, 0xb6, 0xcf, 0x75, 0x7f, 0x5e, 0xd2, 0xe7, 0xa3,
	0xef, 0x0e, 0xf4, 0x24, 0x09, 0x43, 0x6f, 0xc1, 0x95, 0x2f, 0x0c, 0xad, 0x36, 0x1d, 0xde, 0x79,
	0x9c, 0xc1, 0xda, 0x3c, 0xa8, 0x85, 0x84, 0x0f, 0xbf, 0xfc, 0xf8, 0xf5, 0xcd, 0xde, 0x40, 0x43,
	0x35, 0xd5, 0xe8, 0x34, 0x9e, 0x3e, 0x8d, 0x76, 0x62, 0xf9, 0x2e, 0x58, 0x7c, 0x66, 0x6e, 0x7c,
	0x8e, 0x18, 0xf4, 0x9b, 0x2b, 0xa0, 0x3b, 0x4d, 0x9e, 0x0b, 0x8e, 0x05, 0xfe, 0x9f, 0x01, 0x43,
	0xf2, 0x5c, 0x91, 0x44, 0xe8, 0xc9, 0x75, 0x24, 0xf1, 0x99, 0x74, 0x52, 0xfc, 0xd3, 0xd6, 0x9d,
	0x23, 0x0c, 0x9e, 0x9e, 0x99, 0xe8, 0x76, 0x93, 0x79, 0x6e, 0x86, 0x8a, 0x1b, 0x19, 0xd3, 0x0e,
	0x4a, 0x92, 0xcd, 0xc8, 0x22, 0x45, 0xb6, 0x19, 0x3c, 0x58, 0x80, 0xec, 0x85, 0xb5, 0x85, 0x3e,
	0x81, 0xa7, 0x47, 0x62, 0x4b, 0x33, 0x37, 0x22, 0xaf, 0xa0, 0xd9, 0x51, 0x34, 0xdb, 0x5b, 0x8f,
	0xaf, 0xa7, 0x69, 0x07, 0xe6, 0xf9, 0xd8, 0x53, 0x3f, 0x08, 0xcf, 0x7e, 0x07, 0x00, 0x00, 0xff,
	0xff, 0x3f, 0x26, 0x29, 0xdd, 0x58, 0x06, 0x00, 0x00,
}
