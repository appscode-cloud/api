// Code generated by protoc-gen-go.
// source: cloudcredential.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	cloudcredential.proto

It has these top-level messages:
	CloudCredentialCreateRequest
	CloudCredentialUpdateRequest
	CloudCredentialDeleteRequest
	CloudCredentialListRequest
	CloudCredentialListResponse
	Credential
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CloudCredentialCreateRequest struct {
	Name     string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Provider string            `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	Data     map[string]string `protobuf:"bytes,3,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CloudCredentialCreateRequest) Reset()                    { *m = CloudCredentialCreateRequest{} }
func (m *CloudCredentialCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CloudCredentialCreateRequest) ProtoMessage()               {}
func (*CloudCredentialCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CloudCredentialCreateRequest) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type CloudCredentialUpdateRequest struct {
	Name     string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Provider string            `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	Data     map[string]string `protobuf:"bytes,3,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CloudCredentialUpdateRequest) Reset()                    { *m = CloudCredentialUpdateRequest{} }
func (m *CloudCredentialUpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*CloudCredentialUpdateRequest) ProtoMessage()               {}
func (*CloudCredentialUpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CloudCredentialUpdateRequest) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type CloudCredentialDeleteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CloudCredentialDeleteRequest) Reset()                    { *m = CloudCredentialDeleteRequest{} }
func (m *CloudCredentialDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*CloudCredentialDeleteRequest) ProtoMessage()               {}
func (*CloudCredentialDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type CloudCredentialListRequest struct {
}

func (m *CloudCredentialListRequest) Reset()                    { *m = CloudCredentialListRequest{} }
func (m *CloudCredentialListRequest) String() string            { return proto.CompactTextString(m) }
func (*CloudCredentialListRequest) ProtoMessage()               {}
func (*CloudCredentialListRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type CloudCredentialListResponse struct {
	Status      *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Credentials []*Credential  `protobuf:"bytes,2,rep,name=credentials" json:"credentials,omitempty"`
}

func (m *CloudCredentialListResponse) Reset()                    { *m = CloudCredentialListResponse{} }
func (m *CloudCredentialListResponse) String() string            { return proto.CompactTextString(m) }
func (*CloudCredentialListResponse) ProtoMessage()               {}
func (*CloudCredentialListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CloudCredentialListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CloudCredentialListResponse) GetCredentials() []*Credential {
	if m != nil {
		return m.Credentials
	}
	return nil
}

type Credential struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Provider    string `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	Information string `protobuf:"bytes,3,opt,name=information" json:"information,omitempty"`
	ModifiedAt  int64  `protobuf:"varint,4,opt,name=modified_at,json=modifiedAt" json:"modified_at,omitempty"`
}

func (m *Credential) Reset()                    { *m = Credential{} }
func (m *Credential) String() string            { return proto.CompactTextString(m) }
func (*Credential) ProtoMessage()               {}
func (*Credential) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*CloudCredentialCreateRequest)(nil), "credential.v1beta1.CloudCredentialCreateRequest")
	proto.RegisterType((*CloudCredentialUpdateRequest)(nil), "credential.v1beta1.CloudCredentialUpdateRequest")
	proto.RegisterType((*CloudCredentialDeleteRequest)(nil), "credential.v1beta1.CloudCredentialDeleteRequest")
	proto.RegisterType((*CloudCredentialListRequest)(nil), "credential.v1beta1.CloudCredentialListRequest")
	proto.RegisterType((*CloudCredentialListResponse)(nil), "credential.v1beta1.CloudCredentialListResponse")
	proto.RegisterType((*Credential)(nil), "credential.v1beta1.Credential")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for CloudCredentials service

type CloudCredentialsClient interface {
	List(ctx context.Context, in *CloudCredentialListRequest, opts ...grpc.CallOption) (*CloudCredentialListResponse, error)
	Create(ctx context.Context, in *CloudCredentialCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Update(ctx context.Context, in *CloudCredentialUpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *CloudCredentialDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type cloudCredentialsClient struct {
	cc *grpc.ClientConn
}

func NewCloudCredentialsClient(cc *grpc.ClientConn) CloudCredentialsClient {
	return &cloudCredentialsClient{cc}
}

func (c *cloudCredentialsClient) List(ctx context.Context, in *CloudCredentialListRequest, opts ...grpc.CallOption) (*CloudCredentialListResponse, error) {
	out := new(CloudCredentialListResponse)
	err := grpc.Invoke(ctx, "/credential.v1beta1.CloudCredentials/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudCredentialsClient) Create(ctx context.Context, in *CloudCredentialCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/credential.v1beta1.CloudCredentials/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudCredentialsClient) Update(ctx context.Context, in *CloudCredentialUpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/credential.v1beta1.CloudCredentials/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudCredentialsClient) Delete(ctx context.Context, in *CloudCredentialDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/credential.v1beta1.CloudCredentials/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CloudCredentials service

type CloudCredentialsServer interface {
	List(context.Context, *CloudCredentialListRequest) (*CloudCredentialListResponse, error)
	Create(context.Context, *CloudCredentialCreateRequest) (*dtypes.VoidResponse, error)
	Update(context.Context, *CloudCredentialUpdateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *CloudCredentialDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterCloudCredentialsServer(s *grpc.Server, srv CloudCredentialsServer) {
	s.RegisterService(&_CloudCredentials_serviceDesc, srv)
}

func _CloudCredentials_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudCredentialListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCredentialsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.v1beta1.CloudCredentials/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCredentialsServer).List(ctx, req.(*CloudCredentialListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudCredentials_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudCredentialCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCredentialsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.v1beta1.CloudCredentials/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCredentialsServer).Create(ctx, req.(*CloudCredentialCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudCredentials_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudCredentialUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCredentialsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.v1beta1.CloudCredentials/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCredentialsServer).Update(ctx, req.(*CloudCredentialUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudCredentials_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloudCredentialDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudCredentialsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/credential.v1beta1.CloudCredentials/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudCredentialsServer).Delete(ctx, req.(*CloudCredentialDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CloudCredentials_serviceDesc = grpc.ServiceDesc{
	ServiceName: "credential.v1beta1.CloudCredentials",
	HandlerType: (*CloudCredentialsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _CloudCredentials_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CloudCredentials_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CloudCredentials_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CloudCredentials_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("cloudcredential.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x94, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x99, 0x26, 0x56, 0xfb, 0x02, 0xb2, 0x0c, 0xab, 0x84, 0x58, 0xb4, 0xe4, 0x20, 0x65,
	0xc1, 0xc4, 0x8d, 0xc2, 0x6a, 0x45, 0x70, 0xed, 0x7a, 0x13, 0x0f, 0x11, 0x3d, 0x78, 0x91, 0xd9,
	0xce, 0xdb, 0x25, 0x98, 0x66, 0x62, 0x66, 0x5a, 0x28, 0x22, 0x82, 0x17, 0x8f, 0x1e, 0xf6, 0x23,
	0xf8, 0x91, 0x3c, 0xf8, 0x05, 0xfc, 0x04, 0x7e, 0x02, 0xc9, 0x4c, 0x36, 0x4d, 0x77, 0x6b, 0x69,
	0x17, 0xd9, 0x4b, 0x48, 0xfe, 0x2f, 0xef, 0xcd, 0xef, 0xfd, 0x67, 0xde, 0xc0, 0x8d, 0x51, 0x2a,
	0x26, 0x7c, 0x54, 0x20, 0xc7, 0x4c, 0x25, 0x2c, 0x0d, 0xf2, 0x42, 0x28, 0x41, 0x69, 0x43, 0x99,
	0xee, 0x1e, 0xa2, 0x62, 0xbb, 0x5e, 0xf7, 0x58, 0x88, 0xe3, 0x14, 0x43, 0x96, 0x27, 0x21, 0xcb,
	0x32, 0xa1, 0x98, 0x4a, 0x44, 0x26, 0x4d, 0x86, 0x77, 0xb3, 0x94, 0xb9, 0x9a, 0xe5, 0x28, 0x43,
	0xfd, 0x34, 0xba, 0xff, 0x8b, 0x40, 0x77, 0x58, 0xae, 0x31, 0xac, 0x2b, 0x0e, 0x0b, 0x64, 0x0a,
	0x63, 0xfc, 0x38, 0x41, 0xa9, 0x28, 0x05, 0x3b, 0x63, 0x63, 0x74, 0x49, 0x8f, 0xf4, 0x3b, 0xb1,
	0x7e, 0xa7, 0x1e, 0x5c, 0xcb, 0x0b, 0x31, 0x4d, 0x38, 0x16, 0x6e, 0x4b, 0xeb, 0xf5, 0x37, 0x7d,
	0x05, 0x36, 0x67, 0x8a, 0xb9, 0x56, 0xcf, 0xea, 0x3b, 0xd1, 0x20, 0x38, 0x4f, 0x1a, 0xac, 0x5a,
	0x2f, 0x38, 0x60, 0x8a, 0xbd, 0xc8, 0x54, 0x31, 0x8b, 0x75, 0x1d, 0x6f, 0x0f, 0x3a, 0xb5, 0x44,
	0xb7, 0xc0, 0xfa, 0x80, 0xb3, 0x8a, 0xa5, 0x7c, 0xa5, 0xdb, 0x70, 0x65, 0xca, 0xd2, 0x09, 0x56,
	0x1c, 0xe6, 0x63, 0xd0, 0x7a, 0x44, 0x96, 0x75, 0xf6, 0x26, 0xe7, 0x97, 0xda, 0xd9, 0xc2, 0x7a,
	0xff, 0xaf, 0xb3, 0xe8, 0x5c, 0x63, 0x07, 0x98, 0xe2, 0xca, 0xc6, 0xfc, 0x2e, 0x78, 0x67, 0x72,
	0x5e, 0x26, 0x52, 0x55, 0x19, 0xfe, 0x37, 0x02, 0xb7, 0x96, 0x86, 0x65, 0x2e, 0x32, 0x89, 0xf4,
	0x2e, 0xb4, 0xa5, 0x62, 0x6a, 0x22, 0x75, 0x4d, 0x27, 0xba, 0x1e, 0x98, 0xa3, 0x14, 0xbc, 0xd6,
	0x6a, 0x5c, 0x45, 0xe9, 0x33, 0x70, 0xe6, 0xae, 0x48, 0xb7, 0xa5, 0x9d, 0xba, 0xbd, 0xd4, 0xa9,
	0x5a, 0x8a, 0x9b, 0x29, 0xfe, 0x17, 0x80, 0x79, 0x68, 0xe3, 0x2d, 0xea, 0x81, 0x93, 0x64, 0x47,
	0xa2, 0x18, 0xeb, 0xb3, 0xef, 0x5a, 0x3a, 0xdc, 0x94, 0xe8, 0x1d, 0x70, 0xc6, 0x82, 0x27, 0x47,
	0x09, 0xf2, 0xf7, 0x4c, 0xb9, 0x76, 0x8f, 0xf4, 0xad, 0x18, 0x4e, 0xa5, 0x7d, 0x15, 0xfd, 0xb1,
	0x61, 0xeb, 0x8c, 0x15, 0x92, 0xfe, 0x20, 0x60, 0x97, 0x86, 0xd0, 0x60, 0x8d, 0x5d, 0x6f, 0x18,
	0xeb, 0x85, 0x6b, 0xff, 0x6f, 0x9c, 0xf6, 0x07, 0x5f, 0x7f, 0xfe, 0x3e, 0x69, 0x3d, 0xa4, 0x51,
	0xc8, 0xf2, 0x5c, 0x8e, 0x04, 0x37, 0x03, 0x3d, 0xaf, 0x12, 0x56, 0x55, 0x42, 0x7d, 0x33, 0xdc,
	0x6b, 0x78, 0x47, 0xbf, 0x13, 0x68, 0x9b, 0x61, 0xa2, 0xf7, 0x37, 0x9d, 0x3b, 0x6f, 0xfb, 0x74,
	0x4b, 0xdf, 0x8a, 0x84, 0xd7, 0x38, 0x4f, 0x35, 0xce, 0x9e, 0x7f, 0x01, 0x9c, 0x01, 0xd9, 0xd1,
	0x44, 0x66, 0x08, 0xd6, 0x22, 0x5a, 0x98, 0x97, 0xd5, 0x44, 0xde, 0x05, 0x89, 0x4e, 0x08, 0xb4,
	0xcd, 0xb4, 0xac, 0x45, 0xb4, 0x30, 0x58, 0xff, 0x20, 0xda, 0xd7, 0x44, 0x4f, 0x76, 0x1e, 0x6f,
	0x4e, 0x14, 0x7e, 0x2a, 0x8f, 0xf4, 0xe7, 0xe7, 0x9d, 0x77, 0x57, 0xab, 0x7f, 0x0e, 0xdb, 0xfa,
	0x5e, 0x7e, 0xf0, 0x37, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x17, 0x5f, 0x2f, 0xfa, 0x05, 0x00, 0x00,
}
