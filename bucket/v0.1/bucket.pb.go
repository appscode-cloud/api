// Code generated by protoc-gen-go.
// source: bucket.proto
// DO NOT EDIT!

/*
Package bucket is a generated protocol buffer package.

It is generated from these files:
	bucket.proto

It has these top-level messages:
	BucketListRequest
	BucketListResponse
*/
package bucket

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

type BucketListRequest struct {
	CloudCredential string `protobuf:"bytes,1,opt,name=cloud_credential" json:"cloud_credential,omitempty"`
}

func (m *BucketListRequest) Reset()                    { *m = BucketListRequest{} }
func (m *BucketListRequest) String() string            { return proto.CompactTextString(m) }
func (*BucketListRequest) ProtoMessage()               {}
func (*BucketListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type BucketListResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Name   []string       `protobuf:"bytes,2,rep,name=name" json:"name,omitempty"`
}

func (m *BucketListResponse) Reset()                    { *m = BucketListResponse{} }
func (m *BucketListResponse) String() string            { return proto.CompactTextString(m) }
func (*BucketListResponse) ProtoMessage()               {}
func (*BucketListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BucketListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*BucketListRequest)(nil), "bucket.BucketListRequest")
	proto.RegisterType((*BucketListResponse)(nil), "bucket.BucketListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Buckets service

type BucketsClient interface {
	List(ctx context.Context, in *BucketListRequest, opts ...grpc.CallOption) (*BucketListResponse, error)
}

type bucketsClient struct {
	cc *grpc.ClientConn
}

func NewBucketsClient(cc *grpc.ClientConn) BucketsClient {
	return &bucketsClient{cc}
}

func (c *bucketsClient) List(ctx context.Context, in *BucketListRequest, opts ...grpc.CallOption) (*BucketListResponse, error) {
	out := new(BucketListResponse)
	err := grpc.Invoke(ctx, "/bucket.Buckets/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Buckets service

type BucketsServer interface {
	List(context.Context, *BucketListRequest) (*BucketListResponse, error)
}

func RegisterBucketsServer(s *grpc.Server, srv BucketsServer) {
	s.RegisterService(&_Buckets_serviceDesc, srv)
}

func _Buckets_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BucketListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BucketsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bucket.Buckets/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BucketsServer).List(ctx, req.(*BucketListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Buckets_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bucket.Buckets",
	HandlerType: (*BucketsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Buckets_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x2a, 0x4d, 0xce,
	0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0xa4, 0x64, 0xd2, 0xf3,
	0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b,
	0x32, 0xf3, 0xf3, 0x8a, 0x21, 0xaa, 0xa4, 0xc4, 0x40, 0xc2, 0x29, 0x25, 0x95, 0x05, 0xa9, 0xc5,
	0xfa, 0x60, 0x12, 0x22, 0xae, 0xa4, 0xcb, 0x25, 0xe8, 0x04, 0xd6, 0xef, 0x93, 0x59, 0x5c, 0x12,
	0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc1, 0x25, 0x90, 0x9c, 0x93, 0x5f, 0x9a, 0x12,
	0x9f, 0x5c, 0x94, 0x9a, 0x92, 0x9a, 0x57, 0x92, 0x99, 0x98, 0x23, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0xa9, 0xe4, 0xc4, 0x25, 0x84, 0xac, 0xbc, 0xb8, 0x00, 0x68, 0x43, 0xaa, 0x90, 0x1c, 0x17, 0x5b,
	0x31, 0xd0, 0xba, 0xd2, 0x62, 0xb0, 0x2a, 0x6e, 0x23, 0x3e, 0x3d, 0x88, 0x4d, 0x7a, 0xc1, 0x60,
	0x51, 0x21, 0x1e, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0x05, 0x66, 0x0d, 0x4e, 0xa3,
	0x64, 0x2e, 0x76, 0x88, 0x19, 0xc5, 0x42, 0x11, 0x5c, 0x2c, 0x20, 0x83, 0x84, 0x24, 0xf5, 0xa0,
	0x5e, 0xc2, 0x70, 0x8b, 0x94, 0x14, 0x36, 0x29, 0x88, 0xbd, 0x4a, 0x92, 0x4d, 0x97, 0x9f, 0x4c,
	0x66, 0x12, 0x16, 0x12, 0x04, 0x7b, 0x1a, 0xa2, 0xae, 0x58, 0xbf, 0xcc, 0x40, 0xcf, 0x30, 0x89,
	0x0d, 0xec, 0x3d, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0x51, 0x8b, 0x6f, 0x2c, 0x01,
	0x00, 0x00,
}
