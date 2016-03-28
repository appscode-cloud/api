// Code generated by protoc-gen-go.
// source: cloud.proto
// DO NOT EDIT!

/*
Package credential is a generated protocol buffer package.

It is generated from these files:
	cloud.proto

It has these top-level messages:
	CloudCredentialCreateRequest
	CloudCredentialUpdateRequest
	CloudCredentialDeleteRequest
	CloudCredentialListResponse
	CloudCredential
*/
package credential

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

type CloudCredentialListResponse struct {
	Status      *dtypes.Status     `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Credentials []*CloudCredential `protobuf:"bytes,2,rep,name=credentials" json:"credentials,omitempty"`
}

func (m *CloudCredentialListResponse) Reset()                    { *m = CloudCredentialListResponse{} }
func (m *CloudCredentialListResponse) String() string            { return proto.CompactTextString(m) }
func (*CloudCredentialListResponse) ProtoMessage()               {}
func (*CloudCredentialListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CloudCredentialListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CloudCredentialListResponse) GetCredentials() []*CloudCredential {
	if m != nil {
		return m.Credentials
	}
	return nil
}

type CloudCredential struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Provider    string `protobuf:"bytes,2,opt,name=provider" json:"provider,omitempty"`
	Information string `protobuf:"bytes,3,opt,name=information" json:"information,omitempty"`
}

func (m *CloudCredential) Reset()                    { *m = CloudCredential{} }
func (m *CloudCredential) String() string            { return proto.CompactTextString(m) }
func (*CloudCredential) ProtoMessage()               {}
func (*CloudCredential) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*CloudCredentialCreateRequest)(nil), "credential.CloudCredentialCreateRequest")
	proto.RegisterType((*CloudCredentialUpdateRequest)(nil), "credential.CloudCredentialUpdateRequest")
	proto.RegisterType((*CloudCredentialDeleteRequest)(nil), "credential.CloudCredentialDeleteRequest")
	proto.RegisterType((*CloudCredentialListResponse)(nil), "credential.CloudCredentialListResponse")
	proto.RegisterType((*CloudCredential)(nil), "credential.CloudCredential")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion1

// Client API for Cloud service

type CloudClient interface {
	List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*CloudCredentialListResponse, error)
	Create(ctx context.Context, in *CloudCredentialCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Update(ctx context.Context, in *CloudCredentialUpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *CloudCredentialDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type cloudClient struct {
	cc *grpc.ClientConn
}

func NewCloudClient(cc *grpc.ClientConn) CloudClient {
	return &cloudClient{cc}
}

func (c *cloudClient) List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*CloudCredentialListResponse, error) {
	out := new(CloudCredentialListResponse)
	err := grpc.Invoke(ctx, "/credential.Cloud/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudClient) Create(ctx context.Context, in *CloudCredentialCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/credential.Cloud/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudClient) Update(ctx context.Context, in *CloudCredentialUpdateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/credential.Cloud/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudClient) Delete(ctx context.Context, in *CloudCredentialDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/credential.Cloud/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cloud service

type CloudServer interface {
	List(context.Context, *dtypes.VoidRequest) (*CloudCredentialListResponse, error)
	Create(context.Context, *CloudCredentialCreateRequest) (*dtypes.VoidResponse, error)
	Update(context.Context, *CloudCredentialUpdateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *CloudCredentialDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterCloudServer(s *grpc.Server, srv CloudServer) {
	s.RegisterService(&_Cloud_serviceDesc, srv)
}

func _Cloud_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CloudServer).List(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Cloud_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CloudCredentialCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CloudServer).Create(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Cloud_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CloudCredentialUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CloudServer).Update(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Cloud_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CloudCredentialDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(CloudServer).Delete(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Cloud_serviceDesc = grpc.ServiceDesc{
	ServiceName: "credential.Cloud",
	HandlerType: (*CloudServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Cloud_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Cloud_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Cloud_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Cloud_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x94, 0x4f, 0x6a, 0xdb, 0x40,
	0x14, 0xc6, 0x91, 0x65, 0x8b, 0xfa, 0xa9, 0xff, 0x18, 0x97, 0xe2, 0xca, 0xc6, 0xb8, 0xa2, 0x50,
	0xe1, 0x16, 0xc9, 0xa8, 0x9b, 0xe2, 0x55, 0xc1, 0x6e, 0x57, 0x5d, 0x39, 0x24, 0xfb, 0x89, 0x35,
	0x31, 0x22, 0xf2, 0x8c, 0xa2, 0x19, 0x19, 0x4c, 0xc8, 0x26, 0x57, 0xc8, 0x39, 0x72, 0x8b, 0xdc,
	0x20, 0x57, 0xc8, 0x41, 0x32, 0x1a, 0x39, 0x8e, 0xe5, 0x44, 0x4a, 0x08, 0x64, 0x63, 0xac, 0x8f,
	0x79, 0xef, 0xf7, 0xbd, 0x79, 0x9f, 0x04, 0xe6, 0x2c, 0x62, 0x69, 0xe0, 0xc6, 0x09, 0x13, 0x0c,
	0xc1, 0x2c, 0x21, 0x01, 0xa1, 0x22, 0xc4, 0x91, 0xd5, 0x9d, 0x33, 0x36, 0x8f, 0x88, 0x87, 0xe3,
	0xd0, 0xc3, 0x94, 0x32, 0x81, 0x45, 0xc8, 0x28, 0xcf, 0x4f, 0x5a, 0x9f, 0x33, 0x39, 0x10, 0xab,
	0x98, 0x70, 0x4f, 0xfd, 0xe6, 0xba, 0x7d, 0xa9, 0x41, 0x77, 0x9c, 0x75, 0x1c, 0x6f, 0x3a, 0xc9,
	0x7f, 0x58, 0x90, 0x29, 0x39, 0x49, 0x09, 0x17, 0xe8, 0x2d, 0xd4, 0x29, 0x5e, 0x90, 0xb6, 0xd6,
	0xd7, 0x9c, 0x26, 0xfa, 0x08, 0x6f, 0x64, 0xdd, 0x32, 0x0c, 0x48, 0xd2, 0xae, 0x29, 0xe5, 0x0f,
	0xd4, 0x03, 0x2c, 0x70, 0x5b, 0xef, 0xeb, 0x8e, 0xe9, 0xfb, 0xee, 0xbd, 0x23, 0xb7, 0xaa, 0xaf,
	0x3b, 0x91, 0x45, 0x7f, 0xa9, 0x48, 0x56, 0xd6, 0x0f, 0x68, 0x6e, 0x1e, 0x90, 0x09, 0xfa, 0x31,
	0x59, 0xad, 0x69, 0xef, 0xa0, 0xb1, 0xc4, 0x51, 0x4a, 0x72, 0xd4, 0xa8, 0xf6, 0x5b, 0x7b, 0xcc,
	0xef, 0x7e, 0x1c, 0xbc, 0x8a, 0xdf, 0x42, 0xdf, 0x97, 0xfa, 0xfd, 0xf9, 0xc0, 0xee, 0x84, 0x44,
	0xa4, 0xc4, 0xae, 0xcd, 0xa0, 0xb3, 0x73, 0xfa, 0x7f, 0xc8, 0xc5, 0x94, 0xf0, 0x58, 0x6e, 0x92,
	0xa0, 0x1e, 0x18, 0x5c, 0xae, 0x35, 0xe5, 0xea, 0xb8, 0xe9, 0xbf, 0x77, 0xf3, 0x8d, 0xba, 0x7b,
	0x4a, 0x45, 0x43, 0x99, 0x8e, 0x4d, 0x25, 0x97, 0x2e, 0xb2, 0x11, 0x3b, 0x15, 0x23, 0xda, 0xff,
	0xe0, 0xc3, 0x8e, 0xf4, 0xe4, 0x05, 0xb6, 0xc0, 0x0c, 0xe9, 0x11, 0x4b, 0x16, 0x2a, 0x5f, 0xf2,
	0x1e, 0xa5, 0xe8, 0x5f, 0xe9, 0xd0, 0x50, 0x8d, 0xd0, 0x1c, 0xea, 0x99, 0x67, 0xd4, 0xba, 0xf3,
	0x76, 0xc0, 0xc2, 0x60, 0x3d, 0xad, 0xf5, 0xbd, 0xc2, 0xcb, 0xf6, 0xa4, 0xf6, 0xd7, 0xf3, 0xeb,
	0x9b, 0x8b, 0x5a, 0x07, 0x7d, 0x51, 0x71, 0xde, 0x1a, 0xca, 0x5b, 0x0e, 0x3d, 0xf5, 0x06, 0x20,
	0x0a, 0x46, 0x9e, 0x28, 0xe4, 0x3c, 0x37, 0x74, 0xd6, 0xa7, 0xa2, 0xa9, 0x35, 0xec, 0x9b, 0x82,
	0xf5, 0xac, 0x72, 0xd8, 0x48, 0x1b, 0x64, 0xbc, 0x3c, 0x11, 0x95, 0xbc, 0x42, 0x68, 0xaa, 0x79,
	0x76, 0x35, 0x2f, 0x01, 0x23, 0x8f, 0x4a, 0x25, 0xaf, 0x90, 0xa6, 0x12, 0x9e, 0xa3, 0x78, 0xf6,
	0xa0, 0x5f, 0xca, 0xf3, 0x4e, 0xb3, 0x95, 0x9f, 0x1d, 0x1a, 0xea, 0xa3, 0xf0, 0xeb, 0x36, 0x00,
	0x00, 0xff, 0xff, 0xa5, 0xe3, 0xe4, 0xbf, 0x65, 0x04, 0x00, 0x00,
}
