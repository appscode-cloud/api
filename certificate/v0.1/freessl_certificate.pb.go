// Code generated by protoc-gen-go.
// source: freessl_certificate.proto
// DO NOT EDIT!

package certificate

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

type FreeSSLCertificateRegisterRequest struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *FreeSSLCertificateRegisterRequest) Reset()         { *m = FreeSSLCertificateRegisterRequest{} }
func (m *FreeSSLCertificateRegisterRequest) String() string { return proto.CompactTextString(m) }
func (*FreeSSLCertificateRegisterRequest) ProtoMessage()    {}
func (*FreeSSLCertificateRegisterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{0}
}

type FreeSSLCertificateRegisterResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Phid   string         `protobuf:"bytes,2,opt,name=phid" json:"phid,omitempty"`
}

func (m *FreeSSLCertificateRegisterResponse) Reset()         { *m = FreeSSLCertificateRegisterResponse{} }
func (m *FreeSSLCertificateRegisterResponse) String() string { return proto.CompactTextString(m) }
func (*FreeSSLCertificateRegisterResponse) ProtoMessage()    {}
func (*FreeSSLCertificateRegisterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{1}
}

func (m *FreeSSLCertificateRegisterResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type FreeSSLCertificateListResponse struct {
	Status       *dtypes.Status        `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Certificates []*FreeSSLCertificate `protobuf:"bytes,2,rep,name=certificates" json:"certificates,omitempty"`
}

func (m *FreeSSLCertificateListResponse) Reset()                    { *m = FreeSSLCertificateListResponse{} }
func (m *FreeSSLCertificateListResponse) String() string            { return proto.CompactTextString(m) }
func (*FreeSSLCertificateListResponse) ProtoMessage()               {}
func (*FreeSSLCertificateListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *FreeSSLCertificateListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *FreeSSLCertificateListResponse) GetCertificates() []*FreeSSLCertificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

type FreeSSLCertificateDescribeResponse struct {
	Status      *dtypes.Status      `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Certificate *FreeSSLCertificate `protobuf:"bytes,2,opt,name=certificate" json:"certificate,omitempty"`
}

func (m *FreeSSLCertificateDescribeResponse) Reset()         { *m = FreeSSLCertificateDescribeResponse{} }
func (m *FreeSSLCertificateDescribeResponse) String() string { return proto.CompactTextString(m) }
func (*FreeSSLCertificateDescribeResponse) ProtoMessage()    {}
func (*FreeSSLCertificateDescribeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{3}
}

func (m *FreeSSLCertificateDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *FreeSSLCertificateDescribeResponse) GetCertificate() *FreeSSLCertificate {
	if m != nil {
		return m.Certificate
	}
	return nil
}

type FreeSSLCertificate struct {
	Phid       string `protobuf:"bytes,1,opt,name=phid" json:"phid,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	CommonName string `protobuf:"bytes,3,opt,name=common_name,json=commonName" json:"common_name,omitempty"`
	IssuedBy   string `protobuf:"bytes,4,opt,name=issued_by,json=issuedBy" json:"issued_by,omitempty"`
	ValidFrom  string `protobuf:"bytes,5,opt,name=valid_from,json=validFrom" json:"valid_from,omitempty"`
	ExpireDate string `protobuf:"bytes,6,opt,name=expire_date,json=expireDate" json:"expire_date,omitempty"`
	// those feilds will not included into list response.
	// only describe response will include the underlying
	// feilds.
	Sans         []string `protobuf:"bytes,7,rep,name=sans" json:"sans,omitempty"`
	Cert         string   `protobuf:"bytes,8,opt,name=cert" json:"cert,omitempty"`
	Key          string   `protobuf:"bytes,9,opt,name=key" json:"key,omitempty"`
	Version      int32    `protobuf:"varint,10,opt,name=version" json:"version,omitempty"`
	SerialNumber string   `protobuf:"bytes,11,opt,name=serial_number,json=serialNumber" json:"serial_number,omitempty"`
}

func (m *FreeSSLCertificate) Reset()                    { *m = FreeSSLCertificate{} }
func (m *FreeSSLCertificate) String() string            { return proto.CompactTextString(m) }
func (*FreeSSLCertificate) ProtoMessage()               {}
func (*FreeSSLCertificate) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

type FreeSSLCertificateCreateRequest struct {
	AccountPhid string   `protobuf:"bytes,1,opt,name=account_phid,json=accountPhid" json:"account_phid,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	CommonName  string   `protobuf:"bytes,3,opt,name=common_name,json=commonName" json:"common_name,omitempty"`
	San         []string `protobuf:"bytes,4,rep,name=san" json:"san,omitempty"`
	Bundle      bool     `protobuf:"varint,5,opt,name=bundle" json:"bundle,omitempty"`
	KeyData     string   `protobuf:"bytes,6,opt,name=key_data,json=keyData" json:"key_data,omitempty"`
}

func (m *FreeSSLCertificateCreateRequest) Reset()                    { *m = FreeSSLCertificateCreateRequest{} }
func (m *FreeSSLCertificateCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*FreeSSLCertificateCreateRequest) ProtoMessage()               {}
func (*FreeSSLCertificateCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type FreeSSLCertificateCreateResponse struct {
	Status      *dtypes.Status      `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Certificate *FreeSSLCertificate `protobuf:"bytes,2,opt,name=certificate" json:"certificate,omitempty"`
}

func (m *FreeSSLCertificateCreateResponse) Reset()         { *m = FreeSSLCertificateCreateResponse{} }
func (m *FreeSSLCertificateCreateResponse) String() string { return proto.CompactTextString(m) }
func (*FreeSSLCertificateCreateResponse) ProtoMessage()    {}
func (*FreeSSLCertificateCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{6}
}

func (m *FreeSSLCertificateCreateResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *FreeSSLCertificateCreateResponse) GetCertificate() *FreeSSLCertificate {
	if m != nil {
		return m.Certificate
	}
	return nil
}

type FreeSSLCertificateRevokeRequest struct {
	AccountPhid string `protobuf:"bytes,1,opt,name=account_phid,json=accountPhid" json:"account_phid,omitempty"`
	CertPhid    string `protobuf:"bytes,2,opt,name=cert_phid,json=certPhid" json:"cert_phid,omitempty"`
}

func (m *FreeSSLCertificateRevokeRequest) Reset()                    { *m = FreeSSLCertificateRevokeRequest{} }
func (m *FreeSSLCertificateRevokeRequest) String() string            { return proto.CompactTextString(m) }
func (*FreeSSLCertificateRevokeRequest) ProtoMessage()               {}
func (*FreeSSLCertificateRevokeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

type FreeSSLCertificateRenewRequest struct {
	AccountPhid string `protobuf:"bytes,1,opt,name=account_phid,json=accountPhid" json:"account_phid,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	CertPhid    string `protobuf:"bytes,3,opt,name=cert_phid,json=certPhid" json:"cert_phid,omitempty"`
}

func (m *FreeSSLCertificateRenewRequest) Reset()                    { *m = FreeSSLCertificateRenewRequest{} }
func (m *FreeSSLCertificateRenewRequest) String() string            { return proto.CompactTextString(m) }
func (*FreeSSLCertificateRenewRequest) ProtoMessage()               {}
func (*FreeSSLCertificateRenewRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

type FreeSSLCertificateRenewResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Cert   string         `protobuf:"bytes,2,opt,name=cert" json:"cert,omitempty"`
}

func (m *FreeSSLCertificateRenewResponse) Reset()                    { *m = FreeSSLCertificateRenewResponse{} }
func (m *FreeSSLCertificateRenewResponse) String() string            { return proto.CompactTextString(m) }
func (*FreeSSLCertificateRenewResponse) ProtoMessage()               {}
func (*FreeSSLCertificateRenewResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func (m *FreeSSLCertificateRenewResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type FreeSSLCertificateDeleteRequest struct {
	Cert string `protobuf:"bytes,1,opt,name=cert" json:"cert,omitempty"`
}

func (m *FreeSSLCertificateDeleteRequest) Reset()         { *m = FreeSSLCertificateDeleteRequest{} }
func (m *FreeSSLCertificateDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*FreeSSLCertificateDeleteRequest) ProtoMessage()    {}
func (*FreeSSLCertificateDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{10}
}

type FreeSSLCertificateDescribeRequest struct {
	Cert string `protobuf:"bytes,1,opt,name=cert" json:"cert,omitempty"`
}

func (m *FreeSSLCertificateDescribeRequest) Reset()         { *m = FreeSSLCertificateDescribeRequest{} }
func (m *FreeSSLCertificateDescribeRequest) String() string { return proto.CompactTextString(m) }
func (*FreeSSLCertificateDescribeRequest) ProtoMessage()    {}
func (*FreeSSLCertificateDescribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{11}
}

type FreeSSLCertificateDeployRequest struct {
	Cert        string `protobuf:"bytes,1,opt,name=cert" json:"cert,omitempty"`
	SecretName  string `protobuf:"bytes,2,opt,name=secret_name,json=secretName" json:"secret_name,omitempty"`
	ClusterName string `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	Namespace   string `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *FreeSSLCertificateDeployRequest) Reset()         { *m = FreeSSLCertificateDeployRequest{} }
func (m *FreeSSLCertificateDeployRequest) String() string { return proto.CompactTextString(m) }
func (*FreeSSLCertificateDeployRequest) ProtoMessage()    {}
func (*FreeSSLCertificateDeployRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor1, []int{12}
}

func init() {
	proto.RegisterType((*FreeSSLCertificateRegisterRequest)(nil), "certificate.FreeSSLCertificateRegisterRequest")
	proto.RegisterType((*FreeSSLCertificateRegisterResponse)(nil), "certificate.FreeSSLCertificateRegisterResponse")
	proto.RegisterType((*FreeSSLCertificateListResponse)(nil), "certificate.FreeSSLCertificateListResponse")
	proto.RegisterType((*FreeSSLCertificateDescribeResponse)(nil), "certificate.FreeSSLCertificateDescribeResponse")
	proto.RegisterType((*FreeSSLCertificate)(nil), "certificate.FreeSSLCertificate")
	proto.RegisterType((*FreeSSLCertificateCreateRequest)(nil), "certificate.FreeSSLCertificateCreateRequest")
	proto.RegisterType((*FreeSSLCertificateCreateResponse)(nil), "certificate.FreeSSLCertificateCreateResponse")
	proto.RegisterType((*FreeSSLCertificateRevokeRequest)(nil), "certificate.FreeSSLCertificateRevokeRequest")
	proto.RegisterType((*FreeSSLCertificateRenewRequest)(nil), "certificate.FreeSSLCertificateRenewRequest")
	proto.RegisterType((*FreeSSLCertificateRenewResponse)(nil), "certificate.FreeSSLCertificateRenewResponse")
	proto.RegisterType((*FreeSSLCertificateDeleteRequest)(nil), "certificate.FreeSSLCertificateDeleteRequest")
	proto.RegisterType((*FreeSSLCertificateDescribeRequest)(nil), "certificate.FreeSSLCertificateDescribeRequest")
	proto.RegisterType((*FreeSSLCertificateDeployRequest)(nil), "certificate.FreeSSLCertificateDeployRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for FreeSSLCertificates service

type FreeSSLCertificatesClient interface {
	Register(ctx context.Context, in *FreeSSLCertificateRegisterRequest, opts ...grpc.CallOption) (*FreeSSLCertificateRegisterResponse, error)
	List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*FreeSSLCertificateListResponse, error)
	Describe(ctx context.Context, in *FreeSSLCertificateDescribeRequest, opts ...grpc.CallOption) (*FreeSSLCertificateDescribeResponse, error)
	Create(ctx context.Context, in *FreeSSLCertificateCreateRequest, opts ...grpc.CallOption) (*FreeSSLCertificateCreateResponse, error)
	Revoke(ctx context.Context, in *FreeSSLCertificateRevokeRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Renew(ctx context.Context, in *FreeSSLCertificateRenewRequest, opts ...grpc.CallOption) (*FreeSSLCertificateRenewResponse, error)
	Delete(ctx context.Context, in *FreeSSLCertificateDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Deploy(ctx context.Context, in *FreeSSLCertificateDeployRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type freeSSLCertificatesClient struct {
	cc *grpc.ClientConn
}

func NewFreeSSLCertificatesClient(cc *grpc.ClientConn) FreeSSLCertificatesClient {
	return &freeSSLCertificatesClient{cc}
}

func (c *freeSSLCertificatesClient) Register(ctx context.Context, in *FreeSSLCertificateRegisterRequest, opts ...grpc.CallOption) (*FreeSSLCertificateRegisterResponse, error) {
	out := new(FreeSSLCertificateRegisterResponse)
	err := grpc.Invoke(ctx, "/certificate.FreeSSLCertificates/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freeSSLCertificatesClient) List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*FreeSSLCertificateListResponse, error) {
	out := new(FreeSSLCertificateListResponse)
	err := grpc.Invoke(ctx, "/certificate.FreeSSLCertificates/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freeSSLCertificatesClient) Describe(ctx context.Context, in *FreeSSLCertificateDescribeRequest, opts ...grpc.CallOption) (*FreeSSLCertificateDescribeResponse, error) {
	out := new(FreeSSLCertificateDescribeResponse)
	err := grpc.Invoke(ctx, "/certificate.FreeSSLCertificates/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freeSSLCertificatesClient) Create(ctx context.Context, in *FreeSSLCertificateCreateRequest, opts ...grpc.CallOption) (*FreeSSLCertificateCreateResponse, error) {
	out := new(FreeSSLCertificateCreateResponse)
	err := grpc.Invoke(ctx, "/certificate.FreeSSLCertificates/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freeSSLCertificatesClient) Revoke(ctx context.Context, in *FreeSSLCertificateRevokeRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/certificate.FreeSSLCertificates/Revoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freeSSLCertificatesClient) Renew(ctx context.Context, in *FreeSSLCertificateRenewRequest, opts ...grpc.CallOption) (*FreeSSLCertificateRenewResponse, error) {
	out := new(FreeSSLCertificateRenewResponse)
	err := grpc.Invoke(ctx, "/certificate.FreeSSLCertificates/Renew", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freeSSLCertificatesClient) Delete(ctx context.Context, in *FreeSSLCertificateDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/certificate.FreeSSLCertificates/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *freeSSLCertificatesClient) Deploy(ctx context.Context, in *FreeSSLCertificateDeployRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/certificate.FreeSSLCertificates/Deploy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FreeSSLCertificates service

type FreeSSLCertificatesServer interface {
	Register(context.Context, *FreeSSLCertificateRegisterRequest) (*FreeSSLCertificateRegisterResponse, error)
	List(context.Context, *dtypes.VoidRequest) (*FreeSSLCertificateListResponse, error)
	Describe(context.Context, *FreeSSLCertificateDescribeRequest) (*FreeSSLCertificateDescribeResponse, error)
	Create(context.Context, *FreeSSLCertificateCreateRequest) (*FreeSSLCertificateCreateResponse, error)
	Revoke(context.Context, *FreeSSLCertificateRevokeRequest) (*dtypes.VoidResponse, error)
	Renew(context.Context, *FreeSSLCertificateRenewRequest) (*FreeSSLCertificateRenewResponse, error)
	Delete(context.Context, *FreeSSLCertificateDeleteRequest) (*dtypes.VoidResponse, error)
	Deploy(context.Context, *FreeSSLCertificateDeployRequest) (*dtypes.VoidResponse, error)
}

func RegisterFreeSSLCertificatesServer(s *grpc.Server, srv FreeSSLCertificatesServer) {
	s.RegisterService(&_FreeSSLCertificates_serviceDesc, srv)
}

func _FreeSSLCertificates_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreeSSLCertificateRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreeSSLCertificatesServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.FreeSSLCertificates/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreeSSLCertificatesServer).Register(ctx, req.(*FreeSSLCertificateRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreeSSLCertificates_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreeSSLCertificatesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.FreeSSLCertificates/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreeSSLCertificatesServer).List(ctx, req.(*dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreeSSLCertificates_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreeSSLCertificateDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreeSSLCertificatesServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.FreeSSLCertificates/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreeSSLCertificatesServer).Describe(ctx, req.(*FreeSSLCertificateDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreeSSLCertificates_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreeSSLCertificateCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreeSSLCertificatesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.FreeSSLCertificates/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreeSSLCertificatesServer).Create(ctx, req.(*FreeSSLCertificateCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreeSSLCertificates_Revoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreeSSLCertificateRevokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreeSSLCertificatesServer).Revoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.FreeSSLCertificates/Revoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreeSSLCertificatesServer).Revoke(ctx, req.(*FreeSSLCertificateRevokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreeSSLCertificates_Renew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreeSSLCertificateRenewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreeSSLCertificatesServer).Renew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.FreeSSLCertificates/Renew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreeSSLCertificatesServer).Renew(ctx, req.(*FreeSSLCertificateRenewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreeSSLCertificates_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreeSSLCertificateDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreeSSLCertificatesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.FreeSSLCertificates/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreeSSLCertificatesServer).Delete(ctx, req.(*FreeSSLCertificateDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FreeSSLCertificates_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FreeSSLCertificateDeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FreeSSLCertificatesServer).Deploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.FreeSSLCertificates/Deploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FreeSSLCertificatesServer).Deploy(ctx, req.(*FreeSSLCertificateDeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FreeSSLCertificates_serviceDesc = grpc.ServiceDesc{
	ServiceName: "certificate.FreeSSLCertificates",
	HandlerType: (*FreeSSLCertificatesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _FreeSSLCertificates_Register_Handler,
		},
		{
			MethodName: "List",
			Handler:    _FreeSSLCertificates_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _FreeSSLCertificates_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _FreeSSLCertificates_Create_Handler,
		},
		{
			MethodName: "Revoke",
			Handler:    _FreeSSLCertificates_Revoke_Handler,
		},
		{
			MethodName: "Renew",
			Handler:    _FreeSSLCertificates_Renew_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FreeSSLCertificates_Delete_Handler,
		},
		{
			MethodName: "Deploy",
			Handler:    _FreeSSLCertificates_Deploy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor1 = []byte{
	// 862 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x06, 0x25, 0x4b, 0x96, 0x46, 0x6e, 0x51, 0xac, 0x5d, 0x97, 0x66, 0xfd, 0x4b, 0x17, 0x85,
	0x6b, 0xbb, 0x62, 0xeb, 0xa2, 0x28, 0xda, 0x5b, 0x6b, 0xc3, 0x27, 0xc3, 0x28, 0x68, 0x20, 0xb7,
	0x80, 0xa1, 0xc8, 0xb1, 0x42, 0x98, 0x22, 0x19, 0x2e, 0x25, 0x47, 0x48, 0x7c, 0xc9, 0x25, 0x40,
	0x80, 0x20, 0x08, 0x72, 0xc8, 0x0b, 0xe4, 0x0d, 0x72, 0xcf, 0x4b, 0xe4, 0x15, 0xf2, 0x20, 0xd9,
	0x1f, 0x2a, 0xa2, 0x2c, 0x4a, 0x94, 0x12, 0x20, 0x17, 0x63, 0xf9, 0xed, 0xee, 0x7c, 0x33, 0xdf,
	0x8c, 0xbe, 0x35, 0xac, 0x5d, 0xc6, 0x88, 0x94, 0xfa, 0x96, 0x83, 0x71, 0xe2, 0x5d, 0x7a, 0x8e,
	0x9d, 0x60, 0x33, 0x8a, 0xc3, 0x24, 0x24, 0x8d, 0x0c, 0xa4, 0xad, 0xb7, 0xc3, 0xb0, 0xed, 0xa3,
	0x61, 0x47, 0x9e, 0x61, 0x07, 0x41, 0x98, 0xd8, 0x89, 0x17, 0x06, 0x54, 0x1e, 0xd5, 0x56, 0x39,
	0xec, 0x26, 0xfd, 0x08, 0xa9, 0x21, 0xfe, 0x4a, 0x5c, 0xff, 0x1b, 0x76, 0x4e, 0x59, 0xfc, 0x8b,
	0x8b, 0xb3, 0xe3, 0x61, 0x2c, 0x13, 0xdb, 0x1e, 0x4d, 0x30, 0x36, 0xf1, 0x41, 0x17, 0x69, 0x42,
	0x56, 0xa0, 0x82, 0x1d, 0xdb, 0xf3, 0x55, 0x65, 0x5b, 0xd9, 0xab, 0x9b, 0xf2, 0x43, 0xbf, 0x07,
	0xfa, 0xb4, 0xab, 0x34, 0x62, 0xec, 0x48, 0x7e, 0x86, 0x2a, 0x65, 0xa9, 0x74, 0xa9, 0xb8, 0xdc,
	0x38, 0xfa, 0xb6, 0x29, 0xb3, 0x68, 0x5e, 0x08, 0xd4, 0x4c, 0x77, 0x09, 0x81, 0x85, 0xe8, 0xbe,
	0xe7, 0xaa, 0x25, 0x41, 0x21, 0xd6, 0xfa, 0x73, 0x05, 0x36, 0xc7, 0x29, 0xce, 0x18, 0xc1, 0xdc,
	0xe1, 0x8f, 0x61, 0x29, 0x23, 0x16, 0x65, 0x34, 0x65, 0x76, 0x7a, 0xab, 0x99, 0x15, 0x35, 0xa7,
	0x9a, 0x91, 0x4b, 0xfa, 0x0b, 0x25, 0xaf, 0xe4, 0x13, 0xa4, 0x4e, 0xec, 0xb5, 0x70, 0xee, 0x9c,
	0xfe, 0x85, 0x6c, 0x03, 0x45, 0xe5, 0x33, 0xa4, 0x94, 0xbd, 0xa3, 0xbf, 0x2d, 0x01, 0x19, 0x3f,
	0xf3, 0x49, 0x4c, 0x65, 0x28, 0x26, 0xc7, 0x02, 0xbb, 0x83, 0x03, 0x81, 0xf9, 0x9a, 0x6c, 0xb1,
	0x0c, 0xc2, 0x4e, 0x27, 0x0c, 0x2c, 0xb1, 0x55, 0x16, 0x5b, 0x20, 0xa1, 0x73, 0x7e, 0xe0, 0x47,
	0xa8, 0x7b, 0x94, 0x76, 0xd1, 0xb5, 0x5a, 0x7d, 0x75, 0x41, 0x6c, 0xd7, 0x24, 0xf0, 0x5f, 0x9f,
	0x6c, 0x00, 0xf4, 0x6c, 0xdf, 0x73, 0xad, 0xcb, 0x38, 0xec, 0xa8, 0x15, 0xb1, 0x5b, 0x17, 0xc8,
	0x29, 0x03, 0x78, 0x70, 0x7c, 0x18, 0x79, 0x31, 0x5a, 0x2e, 0x2f, 0xaf, 0x2a, 0x83, 0x4b, 0xe8,
	0x24, 0xcd, 0x92, 0xda, 0x01, 0x55, 0x17, 0x59, 0x2f, 0x58, 0x46, 0x7c, 0xcd, 0x31, 0x5e, 0x9f,
	0x5a, 0x93, 0x59, 0xf2, 0x35, 0xf9, 0x0e, 0xca, 0x57, 0xd8, 0x57, 0xeb, 0x02, 0xe2, 0x4b, 0xa2,
	0xc2, 0x62, 0x0f, 0x63, 0xca, 0xe6, 0x5b, 0x05, 0x86, 0x56, 0xcc, 0xc1, 0x27, 0xd9, 0x85, 0x6f,
	0x28, 0xc6, 0x9e, 0xed, 0x5b, 0x41, 0xb7, 0xd3, 0xc2, 0x58, 0x6d, 0x88, 0x5b, 0x4b, 0x12, 0x3c,
	0x17, 0x98, 0xfe, 0x4e, 0x81, 0xad, 0x71, 0xd5, 0x8e, 0x63, 0x14, 0x03, 0x2c, 0x67, 0x7e, 0x07,
	0x96, 0x6c, 0xc7, 0x09, 0xbb, 0x41, 0x62, 0x65, 0xa4, 0x6c, 0xa4, 0xd8, 0xff, 0x9f, 0xad, 0x28,
	0x2b, 0x86, 0x15, 0xca, 0xb4, 0xe4, 0x35, 0xf3, 0x25, 0x59, 0x85, 0x6a, 0xab, 0x1b, 0xb8, 0x3e,
	0x0a, 0x09, 0x6b, 0x66, 0xfa, 0x45, 0xd6, 0xa0, 0xc6, 0x6a, 0xe5, 0xe2, 0xd9, 0xa9, 0x78, 0x8b,
	0xec, 0x9b, 0x29, 0x67, 0xf3, 0x1f, 0xc6, 0xf6, 0xe4, 0x02, 0xbe, 0xfe, 0x18, 0xda, 0x79, 0x7a,
	0x9a, 0xd8, 0x0b, 0xaf, 0xe6, 0xd1, 0x93, 0x0d, 0x1b, 0x0f, 0x6a, 0x65, 0x7c, 0xa0, 0xc6, 0x01,
	0xbe, 0xa9, 0x27, 0x79, 0x56, 0x60, 0x62, 0x80, 0xd7, 0x5f, 0xd8, 0xb1, 0x11, 0xd6, 0xf2, 0x2d,
	0xd6, 0xbb, 0xf9, 0x85, 0x09, 0xd6, 0xf9, 0x0d, 0x4e, 0x4c, 0x76, 0x69, 0x38, 0xd9, 0xfa, 0x9f,
	0x79, 0xe1, 0x4f, 0xd0, 0xc7, 0xe1, 0x1c, 0x0e, 0xae, 0x29, 0x99, 0x6b, 0x7f, 0xe5, 0x99, 0xf6,
	0xd0, 0x86, 0x26, 0x5f, 0x7c, 0xad, 0xe4, 0x13, 0x46, 0x7e, 0xd8, 0x9f, 0x72, 0x8f, 0x4f, 0x35,
	0x45, 0x27, 0xc6, 0xc4, 0xca, 0xc8, 0x07, 0x12, 0x12, 0x53, 0xcd, 0xb4, 0x77, 0xfc, 0x2e, 0x37,
	0xfe, 0xec, 0xdc, 0x37, 0x52, 0x4c, 0x1c, 0x59, 0x87, 0x3a, 0xdf, 0xa2, 0x91, 0xed, 0x60, 0x6a,
	0x25, 0x43, 0xe0, 0xe8, 0x4d, 0x0d, 0x96, 0xc7, 0x33, 0xa3, 0xe4, 0x99, 0x02, 0xb5, 0xc1, 0x9b,
	0x42, 0x9a, 0x45, 0x43, 0x39, 0xfa, 0x6e, 0x69, 0xc6, 0xcc, 0xe7, 0x65, 0x2f, 0xf5, 0x8d, 0x27,
	0xef, 0x3f, 0xbc, 0x2a, 0xfd, 0x40, 0xbe, 0x17, 0xaf, 0x68, 0xfa, 0xf0, 0x1a, 0xbd, 0xdf, 0x9a,
	0xbf, 0x1b, 0x89, 0x4f, 0x49, 0x1b, 0x16, 0xf8, 0xe3, 0x43, 0x96, 0x07, 0x2d, 0xbe, 0x13, 0x7a,
	0xee, 0x80, 0xec, 0xa0, 0x80, 0x2c, 0xfb, 0x6c, 0x15, 0x11, 0xbd, 0x64, 0x55, 0x0f, 0xfa, 0x59,
	0x58, 0xf5, 0xad, 0xc6, 0x17, 0x56, 0x7d, 0xfb, 0xbd, 0xd2, 0x7f, 0x12, 0xc9, 0x6c, 0x92, 0xf5,
	0xdc, 0x64, 0x8c, 0x47, 0x3c, 0xdc, 0x0d, 0x79, 0xaa, 0x40, 0x55, 0x3a, 0x0c, 0x39, 0x2c, 0x60,
	0x18, 0x71, 0x52, 0xed, 0xd7, 0x19, 0x4f, 0xa7, 0xd9, 0x6c, 0x8b, 0x6c, 0x34, 0x2d, 0x5f, 0x9a,
	0x7f, 0x94, 0x7d, 0x72, 0x0d, 0x55, 0xe9, 0x2d, 0x85, 0x89, 0x8c, 0x58, 0x90, 0xb6, 0x32, 0xda,
	0xb6, 0x94, 0xef, 0x17, 0xc1, 0xb7, 0xbb, 0xbf, 0x33, 0xa5, 0x7a, 0x61, 0x0e, 0x37, 0x84, 0xd9,
	0x6e, 0x45, 0xfc, 0xf8, 0xc9, 0x41, 0x21, 0xf1, 0xd0, 0x98, 0xb4, 0xc3, 0xd9, 0x0e, 0xcf, 0x9f,
	0x4f, 0x0c, 0x55, 0x69, 0x16, 0x85, 0x42, 0x8c, 0x78, 0xca, 0x04, 0x21, 0xd2, 0x31, 0xd8, 0x9f,
	0x3e, 0x06, 0x8f, 0x39, 0x27, 0xf7, 0x8b, 0x19, 0x38, 0x33, 0xb6, 0x32, 0x81, 0xb3, 0x29, 0x38,
	0xf7, 0xb4, 0xdd, 0x69, 0x9c, 0x86, 0x2b, 0x22, 0xb1, 0xd6, 0xb7, 0xaa, 0xe2, 0xbf, 0xd6, 0x3f,
	0x3e, 0x06, 0x00, 0x00, 0xff, 0xff, 0x86, 0x67, 0x5e, 0xea, 0x15, 0x0b, 0x00, 0x00,
}
