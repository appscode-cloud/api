// Code generated by protoc-gen-go.
// source: certificate.proto
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

type CertificateListResponse struct {
	Status       *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Certificates []*Certificate `protobuf:"bytes,2,rep,name=certificates" json:"certificates,omitempty"`
}

func (m *CertificateListResponse) Reset()                    { *m = CertificateListResponse{} }
func (m *CertificateListResponse) String() string            { return proto.CompactTextString(m) }
func (*CertificateListResponse) ProtoMessage()               {}
func (*CertificateListResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *CertificateListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CertificateListResponse) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

type CertificateDescribeRequest struct {
	Uid string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *CertificateDescribeRequest) Reset()                    { *m = CertificateDescribeRequest{} }
func (m *CertificateDescribeRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateDescribeRequest) ProtoMessage()               {}
func (*CertificateDescribeRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type CertificateDescribeResponse struct {
	Status      *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Certificate *Certificate   `protobuf:"bytes,2,opt,name=certificate" json:"certificate,omitempty"`
}

func (m *CertificateDescribeResponse) Reset()                    { *m = CertificateDescribeResponse{} }
func (m *CertificateDescribeResponse) String() string            { return proto.CompactTextString(m) }
func (*CertificateDescribeResponse) ProtoMessage()               {}
func (*CertificateDescribeResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *CertificateDescribeResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CertificateDescribeResponse) GetCertificate() *Certificate {
	if m != nil {
		return m.Certificate
	}
	return nil
}

type CertificateCreateRequest struct {
	AccountPhid      string       `protobuf:"bytes,1,opt,name=account_phid,json=accountPhid" json:"account_phid,omitempty"`
	Name             string       `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	CommonName       string       `protobuf:"bytes,3,opt,name=common_name,json=commonName" json:"common_name,omitempty"`
	San              []string     `protobuf:"bytes,4,rep,name=san" json:"san,omitempty"`
	Bundle           bool         `protobuf:"varint,5,opt,name=bundle" json:"bundle,omitempty"`
	KeyData          string       `protobuf:"bytes,6,opt,name=key_data,json=keyData" json:"key_data,omitempty"`
	SubjectInfo      *SubjectInfo `protobuf:"bytes,7,opt,name=subject_info,json=subjectInfo" json:"subject_info,omitempty"`
	IssuePrivateCert bool         `protobuf:"varint,8,opt,name=issue_private_cert,json=issuePrivateCert" json:"issue_private_cert,omitempty"`
}

func (m *CertificateCreateRequest) Reset()                    { *m = CertificateCreateRequest{} }
func (m *CertificateCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateCreateRequest) ProtoMessage()               {}
func (*CertificateCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *CertificateCreateRequest) GetSubjectInfo() *SubjectInfo {
	if m != nil {
		return m.SubjectInfo
	}
	return nil
}

// A Name contains the SubjectInfo fields.
type SubjectInfo struct {
	C  string `protobuf:"bytes,1,opt,name=C,json=c" json:"C,omitempty"`
	ST string `protobuf:"bytes,2,opt,name=ST,json=sT" json:"ST,omitempty"`
	L  string `protobuf:"bytes,3,opt,name=L,json=l" json:"L,omitempty"`
	O  string `protobuf:"bytes,4,opt,name=O,json=o" json:"O,omitempty"`
	OU string `protobuf:"bytes,5,opt,name=OU,json=oU" json:"OU,omitempty"`
}

func (m *SubjectInfo) Reset()                    { *m = SubjectInfo{} }
func (m *SubjectInfo) String() string            { return proto.CompactTextString(m) }
func (*SubjectInfo) ProtoMessage()               {}
func (*SubjectInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

type CertificateCreateResponse struct {
	Status      *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Certificate *Certificate   `protobuf:"bytes,2,opt,name=certificate" json:"certificate,omitempty"`
	Phid        string         `protobuf:"bytes,3,opt,name=phid" json:"phid,omitempty"`
	Cert        []byte         `protobuf:"bytes,4,opt,name=cert,proto3" json:"cert,omitempty"`
	Key         []byte         `protobuf:"bytes,5,opt,name=key,proto3" json:"key,omitempty"`
	Root        []byte         `protobuf:"bytes,6,opt,name=root,proto3" json:"root,omitempty"`
}

func (m *CertificateCreateResponse) Reset()                    { *m = CertificateCreateResponse{} }
func (m *CertificateCreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CertificateCreateResponse) ProtoMessage()               {}
func (*CertificateCreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *CertificateCreateResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CertificateCreateResponse) GetCertificate() *Certificate {
	if m != nil {
		return m.Certificate
	}
	return nil
}

type Certificate struct {
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

func (m *Certificate) Reset()                    { *m = Certificate{} }
func (m *Certificate) String() string            { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()               {}
func (*Certificate) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

type CertificateImportRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	CertData string `protobuf:"bytes,2,opt,name=cert_data,json=certData" json:"cert_data,omitempty"`
	KeyData  string `protobuf:"bytes,3,opt,name=key_data,json=keyData" json:"key_data,omitempty"`
}

func (m *CertificateImportRequest) Reset()                    { *m = CertificateImportRequest{} }
func (m *CertificateImportRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateImportRequest) ProtoMessage()               {}
func (*CertificateImportRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

type CertificateDeleteRequest struct {
	Uid string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
}

func (m *CertificateDeleteRequest) Reset()                    { *m = CertificateDeleteRequest{} }
func (m *CertificateDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateDeleteRequest) ProtoMessage()               {}
func (*CertificateDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

type CertificateRenewRequest struct {
	AccountPhid string `protobuf:"bytes,1,opt,name=account_phid,json=accountPhid" json:"account_phid,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Uid         string `protobuf:"bytes,3,opt,name=uid" json:"uid,omitempty"`
}

func (m *CertificateRenewRequest) Reset()                    { *m = CertificateRenewRequest{} }
func (m *CertificateRenewRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateRenewRequest) ProtoMessage()               {}
func (*CertificateRenewRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

type CertificateRenewResponse struct {
	Status *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Cert   string         `protobuf:"bytes,2,opt,name=cert" json:"cert,omitempty"`
}

func (m *CertificateRenewResponse) Reset()                    { *m = CertificateRenewResponse{} }
func (m *CertificateRenewResponse) String() string            { return proto.CompactTextString(m) }
func (*CertificateRenewResponse) ProtoMessage()               {}
func (*CertificateRenewResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *CertificateRenewResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type CertificateRevokeRequest struct {
	AccountPhid string `protobuf:"bytes,1,opt,name=account_phid,json=accountPhid" json:"account_phid,omitempty"`
	Uid         string `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *CertificateRevokeRequest) Reset()                    { *m = CertificateRevokeRequest{} }
func (m *CertificateRevokeRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateRevokeRequest) ProtoMessage()               {}
func (*CertificateRevokeRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

type CertificateDeployRequest struct {
	Uid         string `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	SecretName  string `protobuf:"bytes,2,opt,name=secret_name,json=secretName" json:"secret_name,omitempty"`
	ClusterName string `protobuf:"bytes,3,opt,name=cluster_name,json=clusterName" json:"cluster_name,omitempty"`
	Namespace   string `protobuf:"bytes,4,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *CertificateDeployRequest) Reset()                    { *m = CertificateDeployRequest{} }
func (m *CertificateDeployRequest) String() string            { return proto.CompactTextString(m) }
func (*CertificateDeployRequest) ProtoMessage()               {}
func (*CertificateDeployRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

type DNSCheckRequest struct {
	Domain  string `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	KeyAuth string `protobuf:"bytes,2,opt,name=key_auth,json=keyAuth" json:"key_auth,omitempty"`
}

func (m *DNSCheckRequest) Reset()                    { *m = DNSCheckRequest{} }
func (m *DNSCheckRequest) String() string            { return proto.CompactTextString(m) }
func (*DNSCheckRequest) ProtoMessage()               {}
func (*DNSCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{13} }

func init() {
	proto.RegisterType((*CertificateListResponse)(nil), "certificate.CertificateListResponse")
	proto.RegisterType((*CertificateDescribeRequest)(nil), "certificate.CertificateDescribeRequest")
	proto.RegisterType((*CertificateDescribeResponse)(nil), "certificate.CertificateDescribeResponse")
	proto.RegisterType((*CertificateCreateRequest)(nil), "certificate.CertificateCreateRequest")
	proto.RegisterType((*SubjectInfo)(nil), "certificate.SubjectInfo")
	proto.RegisterType((*CertificateCreateResponse)(nil), "certificate.CertificateCreateResponse")
	proto.RegisterType((*Certificate)(nil), "certificate.Certificate")
	proto.RegisterType((*CertificateImportRequest)(nil), "certificate.CertificateImportRequest")
	proto.RegisterType((*CertificateDeleteRequest)(nil), "certificate.CertificateDeleteRequest")
	proto.RegisterType((*CertificateRenewRequest)(nil), "certificate.CertificateRenewRequest")
	proto.RegisterType((*CertificateRenewResponse)(nil), "certificate.CertificateRenewResponse")
	proto.RegisterType((*CertificateRevokeRequest)(nil), "certificate.CertificateRevokeRequest")
	proto.RegisterType((*CertificateDeployRequest)(nil), "certificate.CertificateDeployRequest")
	proto.RegisterType((*DNSCheckRequest)(nil), "certificate.DNSCheckRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Certificates service

type CertificatesClient interface {
	List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*CertificateListResponse, error)
	Describe(ctx context.Context, in *CertificateDescribeRequest, opts ...grpc.CallOption) (*CertificateDescribeResponse, error)
	Create(ctx context.Context, in *CertificateCreateRequest, opts ...grpc.CallOption) (*CertificateCreateResponse, error)
	Import(ctx context.Context, in *CertificateImportRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *CertificateDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Renew(ctx context.Context, in *CertificateRenewRequest, opts ...grpc.CallOption) (*CertificateRenewResponse, error)
	Revoke(ctx context.Context, in *CertificateRevokeRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Deploy(ctx context.Context, in *CertificateDeployRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type certificatesClient struct {
	cc *grpc.ClientConn
}

func NewCertificatesClient(cc *grpc.ClientConn) CertificatesClient {
	return &certificatesClient{cc}
}

func (c *certificatesClient) List(ctx context.Context, in *dtypes.VoidRequest, opts ...grpc.CallOption) (*CertificateListResponse, error) {
	out := new(CertificateListResponse)
	err := grpc.Invoke(ctx, "/certificate.Certificates/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Describe(ctx context.Context, in *CertificateDescribeRequest, opts ...grpc.CallOption) (*CertificateDescribeResponse, error) {
	out := new(CertificateDescribeResponse)
	err := grpc.Invoke(ctx, "/certificate.Certificates/Describe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Create(ctx context.Context, in *CertificateCreateRequest, opts ...grpc.CallOption) (*CertificateCreateResponse, error) {
	out := new(CertificateCreateResponse)
	err := grpc.Invoke(ctx, "/certificate.Certificates/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Import(ctx context.Context, in *CertificateImportRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/certificate.Certificates/Import", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Delete(ctx context.Context, in *CertificateDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/certificate.Certificates/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Renew(ctx context.Context, in *CertificateRenewRequest, opts ...grpc.CallOption) (*CertificateRenewResponse, error) {
	out := new(CertificateRenewResponse)
	err := grpc.Invoke(ctx, "/certificate.Certificates/Renew", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Revoke(ctx context.Context, in *CertificateRevokeRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/certificate.Certificates/Revoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificatesClient) Deploy(ctx context.Context, in *CertificateDeployRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/certificate.Certificates/Deploy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Certificates service

type CertificatesServer interface {
	List(context.Context, *dtypes.VoidRequest) (*CertificateListResponse, error)
	Describe(context.Context, *CertificateDescribeRequest) (*CertificateDescribeResponse, error)
	Create(context.Context, *CertificateCreateRequest) (*CertificateCreateResponse, error)
	Import(context.Context, *CertificateImportRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *CertificateDeleteRequest) (*dtypes.VoidResponse, error)
	Renew(context.Context, *CertificateRenewRequest) (*CertificateRenewResponse, error)
	Revoke(context.Context, *CertificateRevokeRequest) (*dtypes.VoidResponse, error)
	Deploy(context.Context, *CertificateDeployRequest) (*dtypes.VoidResponse, error)
}

func RegisterCertificatesServer(s *grpc.Server, srv CertificatesServer) {
	s.RegisterService(&_Certificates_serviceDesc, srv)
}

func _Certificates_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.Certificates/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).List(ctx, req.(*dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateDescribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.Certificates/Describe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Describe(ctx, req.(*CertificateDescribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.Certificates/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Create(ctx, req.(*CertificateCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_Import_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateImportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).Import(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.Certificates/Import",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Import(ctx, req.(*CertificateImportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.Certificates/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Delete(ctx, req.(*CertificateDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_Renew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRenewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).Renew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.Certificates/Renew",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Renew(ctx, req.(*CertificateRenewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_Revoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRevokeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).Revoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.Certificates/Revoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Revoke(ctx, req.(*CertificateRevokeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Certificates_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateDeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificatesServer).Deploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/certificate.Certificates/Deploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificatesServer).Deploy(ctx, req.(*CertificateDeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Certificates_serviceDesc = grpc.ServiceDesc{
	ServiceName: "certificate.Certificates",
	HandlerType: (*CertificatesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Certificates_List_Handler,
		},
		{
			MethodName: "Describe",
			Handler:    _Certificates_Describe_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Certificates_Create_Handler,
		},
		{
			MethodName: "Import",
			Handler:    _Certificates_Import_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Certificates_Delete_Handler,
		},
		{
			MethodName: "Renew",
			Handler:    _Certificates_Renew_Handler,
		},
		{
			MethodName: "Revoke",
			Handler:    _Certificates_Revoke_Handler,
		},
		{
			MethodName: "Deploy",
			Handler:    _Certificates_Deploy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor2 = []byte{
	// 1036 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x96, 0xdd, 0x6e, 0xdc, 0x44,
	0x14, 0xc7, 0xe5, 0x4d, 0xb2, 0xd9, 0x3d, 0xbb, 0x94, 0x32, 0xa0, 0xe2, 0xba, 0x45, 0x14, 0x43,
	0x3f, 0xa8, 0xd2, 0x35, 0xa4, 0x04, 0x89, 0xb4, 0x12, 0x1f, 0xbb, 0x42, 0xaa, 0x54, 0x25, 0x95,
	0xd3, 0x96, 0x4b, 0xcb, 0x6b, 0x4f, 0x1a, 0x13, 0xaf, 0xc7, 0x78, 0xec, 0x85, 0x15, 0x54, 0x08,
	0x90, 0xe0, 0x02, 0x09, 0x09, 0xc1, 0x13, 0x70, 0xcd, 0xab, 0x70, 0x03, 0xaf, 0xc0, 0x83, 0x30,
	0x73, 0x66, 0xb6, 0x3b, 0x8e, 0xe2, 0x24, 0x2b, 0xa4, 0xde, 0xac, 0x66, 0xce, 0x1c, 0xcf, 0xf9,
	0xcd, 0xff, 0x9c, 0x39, 0xb3, 0xf0, 0x52, 0x44, 0x8b, 0x32, 0xd9, 0x4f, 0xa2, 0xb0, 0xa4, 0x83,
	0xbc, 0x60, 0x25, 0x23, 0x3d, 0xc3, 0xe4, 0x5c, 0x7e, 0xc2, 0xd8, 0x93, 0x94, 0x7a, 0x61, 0x9e,
	0x78, 0x61, 0x96, 0xb1, 0x32, 0x2c, 0x13, 0x96, 0x71, 0xe5, 0xea, 0x5c, 0x90, 0xe6, 0xb8, 0x9c,
	0xe5, 0x94, 0x7b, 0xf8, 0xab, 0xec, 0xee, 0xb7, 0xf0, 0xea, 0x70, 0xb1, 0xc9, 0xfd, 0x84, 0x97,
	0x3e, 0xe5, 0xb9, 0xf8, 0x8e, 0x92, 0x6b, 0xd0, 0xe6, 0x62, 0x93, 0x8a, 0xdb, 0xd6, 0x15, 0xeb,
	0x46, 0x6f, 0xf3, 0xdc, 0x40, 0x7d, 0x3f, 0xd8, 0x43, 0xab, 0xaf, 0x57, 0xc9, 0x5d, 0xe8, 0x1b,
	0x1c, 0xdc, 0x6e, 0x5d, 0x59, 0x11, 0xde, 0xf6, 0xc0, 0xe4, 0x35, 0x62, 0xf8, 0x35, 0x6f, 0x77,
	0x00, 0x8e, 0xb1, 0x38, 0xa2, 0x3c, 0x2a, 0x92, 0x31, 0xf5, 0xe9, 0x17, 0x15, 0xe5, 0x25, 0x39,
	0x0f, 0x2b, 0x55, 0x12, 0x23, 0x40, 0xd7, 0x97, 0x43, 0xf7, 0x3b, 0x0b, 0x2e, 0x1d, 0xfb, 0xc1,
	0x92, 0xd4, 0xdb, 0x60, 0xaa, 0x27, 0xa0, 0xad, 0x13, 0xa1, 0x4d, 0x67, 0xf7, 0x8f, 0x16, 0xd8,
	0xc6, 0xe2, 0xb0, 0xa0, 0xd2, 0x45, 0x23, 0xbf, 0x01, 0xfd, 0x30, 0x8a, 0x58, 0x95, 0x95, 0x41,
	0x7e, 0xf0, 0x8c, 0xbd, 0xa7, 0x6d, 0x0f, 0x84, 0x89, 0x10, 0x58, 0xcd, 0xc2, 0x89, 0x0a, 0xda,
	0xf5, 0x71, 0x4c, 0x5e, 0x17, 0x3c, 0x6c, 0x32, 0x61, 0x59, 0x80, 0x4b, 0x2b, 0xb8, 0x04, 0xca,
	0xb4, 0x23, 0x1d, 0x84, 0x14, 0x3c, 0xcc, 0xec, 0x55, 0xa1, 0xae, 0x90, 0x42, 0x0c, 0xc9, 0x05,
	0x68, 0x8f, 0xab, 0x2c, 0x4e, 0xa9, 0xbd, 0x26, 0xbc, 0x3b, 0xbe, 0x9e, 0x91, 0x8b, 0xd0, 0x39,
	0xa4, 0xb3, 0x20, 0x0e, 0xcb, 0xd0, 0x6e, 0xe3, 0x3e, 0xeb, 0x62, 0x3e, 0x12, 0x53, 0x72, 0x07,
	0xfa, 0xbc, 0x1a, 0x7f, 0x4e, 0xa3, 0x32, 0x48, 0xb2, 0x7d, 0x66, 0xaf, 0x1f, 0x73, 0xec, 0x3d,
	0xe5, 0x70, 0x4f, 0xac, 0xfb, 0x3d, 0xbe, 0x98, 0x90, 0x0d, 0x20, 0x09, 0xe7, 0x15, 0x0d, 0xf2,
	0x22, 0x99, 0x0a, 0xcf, 0x40, 0x7e, 0x65, 0x77, 0x30, 0xf6, 0x79, 0x5c, 0x79, 0xa0, 0x16, 0xa4,
	0x36, 0xee, 0x67, 0xd0, 0x33, 0x76, 0x22, 0x7d, 0xb0, 0x86, 0x5a, 0x0b, 0x2b, 0x22, 0xe7, 0xa0,
	0xb5, 0xf7, 0x50, 0x9f, 0xbf, 0xc5, 0x1f, 0xca, 0xd5, 0xfb, 0xfa, 0xcc, 0x56, 0x2a, 0x67, 0xbb,
	0xe2, 0xa0, 0x38, 0x63, 0xd2, 0x77, 0xf7, 0x11, 0x1e, 0x51, 0xf8, 0xb2, 0x47, 0xee, 0x5f, 0x16,
	0x5c, 0x3c, 0x46, 0xfd, 0xe7, 0x97, 0x7f, 0x99, 0x3f, 0x4c, 0xad, 0x02, 0xc6, 0xb1, 0xb4, 0xa1,
	0x1c, 0x12, 0xbb, 0xef, 0xe3, 0x58, 0xa6, 0x4c, 0x08, 0x8f, 0xe8, 0x7d, 0x5f, 0x0e, 0xa5, 0x57,
	0xc1, 0x58, 0x89, 0x69, 0x11, 0x5e, 0x72, 0xec, 0xfe, 0xd9, 0x82, 0xde, 0xf0, 0x98, 0xdd, 0xad,
	0xfa, 0xee, 0xcb, 0x57, 0xcc, 0x25, 0xe8, 0x62, 0x56, 0xe2, 0x60, 0x3c, 0xd3, 0x72, 0x76, 0x94,
	0xe1, 0x93, 0x19, 0x79, 0x0d, 0x60, 0x1a, 0xa6, 0x49, 0x1c, 0xec, 0x17, 0x6c, 0xa2, 0xd5, 0xed,
	0xa2, 0xe5, 0x53, 0x61, 0x90, 0x9b, 0xd3, 0xaf, 0xf2, 0xa4, 0xa0, 0xb2, 0x8c, 0xa8, 0x2e, 0x23,
	0x50, 0xa6, 0x91, 0xa6, 0x14, 0x35, 0xc8, 0x45, 0x05, 0xc9, 0x7a, 0xc4, 0xf1, 0x33, 0x0d, 0x3a,
	0x8a, 0xd2, 0xd4, 0xa0, 0xab, 0x6e, 0xb0, 0xd4, 0xc0, 0x86, 0xf5, 0x29, 0x2d, 0xb8, 0x68, 0x4e,
	0x36, 0x08, 0xeb, 0x9a, 0x3f, 0x9f, 0x92, 0x37, 0xe1, 0x05, 0x4e, 0x8b, 0x24, 0x4c, 0x83, 0xac,
	0x9a, 0x8c, 0x69, 0x61, 0xf7, 0xf0, 0xab, 0xbe, 0x32, 0xee, 0xa0, 0xcd, 0xdd, 0xaf, 0xdd, 0xbd,
	0x7b, 0x93, 0x9c, 0x15, 0xe5, 0xfc, 0xee, 0xcd, 0x65, 0xb2, 0x0c, 0x99, 0x84, 0x0a, 0x12, 0x44,
	0x5d, 0x07, 0xa5, 0x5f, 0x47, 0x1a, 0xf0, 0x3e, 0x98, 0x57, 0x65, 0xa5, 0x76, 0x55, 0xdc, 0x8d,
	0x5a, 0x9c, 0x11, 0x4d, 0x69, 0x79, 0x42, 0x5b, 0x1a, 0xd7, 0xfa, 0xa8, 0x4f, 0x33, 0xfa, 0xe5,
	0xff, 0x6c, 0x08, 0x3a, 0xc6, 0xca, 0x22, 0xc6, 0xe3, 0x1a, 0x91, 0x8e, 0xb1, 0x64, 0xd9, 0xcf,
	0x53, 0xd4, 0x5a, 0xa4, 0xc8, 0xdd, 0x3d, 0xb2, 0xef, 0x94, 0x1d, 0x2e, 0xd3, 0xcd, 0x34, 0x68,
	0x6b, 0x01, 0xfa, 0x8b, 0x75, 0x44, 0xbb, 0x3c, 0x65, 0xb3, 0x46, 0xed, 0x64, 0xad, 0x71, 0x1a,
	0x15, 0xb4, 0x0c, 0x0c, 0x11, 0x40, 0x99, 0xb0, 0x90, 0x05, 0x44, 0x94, 0x56, 0xbc, 0xa4, 0x85,
	0x59, 0xea, 0x3d, 0x6d, 0x43, 0x97, 0xcb, 0xd0, 0x95, 0x4b, 0x3c, 0x0f, 0x23, 0xaa, 0x6b, 0x7d,
	0x61, 0x70, 0x47, 0xf0, 0xe2, 0x68, 0x67, 0x6f, 0x78, 0x40, 0xa3, 0xc3, 0x39, 0x86, 0x68, 0x9e,
	0x31, 0x9b, 0x84, 0x49, 0xa6, 0x49, 0xf4, 0x6c, 0x5e, 0x11, 0x61, 0x55, 0x1e, 0x68, 0x12, 0x59,
	0x11, 0x1f, 0x8b, 0xe9, 0xe6, 0xdf, 0x1d, 0xe8, 0x1b, 0xc7, 0xe2, 0xa4, 0x80, 0x55, 0xf9, 0x62,
	0x92, 0x97, 0xe7, 0x62, 0x3f, 0x66, 0x49, 0xac, 0x03, 0x38, 0x6f, 0x35, 0xf5, 0x12, 0xf3, 0x91,
	0x75, 0xbd, 0xef, 0xff, 0xf9, 0xf7, 0xb7, 0xd6, 0xdb, 0xe4, 0xba, 0x78, 0xb7, 0x73, 0x1e, 0xb1,
	0x58, 0x3d, 0xe0, 0xc6, 0xa7, 0xde, 0xf4, 0x9d, 0xc1, 0xbb, 0x5e, 0x99, 0xf2, 0x5b, 0x78, 0x9f,
	0x7e, 0xb7, 0xa0, 0x33, 0x7f, 0xf4, 0xc8, 0xf5, 0xa6, 0x18, 0x47, 0xde, 0x51, 0xe7, 0xc6, 0xe9,
	0x8e, 0x1a, 0x68, 0x0b, 0x81, 0x3c, 0x72, 0xeb, 0x8c, 0x40, 0xde, 0xd7, 0x22, 0x85, 0x4f, 0xc9,
	0xcf, 0x16, 0xb4, 0x55, 0x27, 0x26, 0x57, 0x9b, 0x62, 0xd5, 0xde, 0x49, 0xe7, 0xda, 0x69, 0x6e,
	0x1a, 0x68, 0x13, 0x81, 0x36, 0x9c, 0xb3, 0x2a, 0xb4, 0x6d, 0xdd, 0x24, 0x3f, 0x09, 0x1a, 0xd5,
	0x19, 0x9a, 0x69, 0x6a, 0x9d, 0xc3, 0x79, 0xa5, 0x9e, 0x42, 0x1d, 0xfb, 0x23, 0x8c, 0xbd, 0xed,
	0x6c, 0x9d, 0x59, 0x0c, 0x59, 0x73, 0x4f, 0xbd, 0x04, 0xf7, 0x96, 0x24, 0xdf, 0x40, 0x5b, 0xb5,
	0x8e, 0x66, 0x90, 0x5a, 0x6b, 0x69, 0x00, 0xd1, 0x59, 0xb9, 0xb9, 0x64, 0x56, 0x7e, 0xb5, 0x60,
	0x0d, 0xfb, 0x04, 0x69, 0xac, 0x46, 0xb3, 0x55, 0x39, 0x57, 0x4f, 0xf1, 0xd2, 0x34, 0x77, 0x90,
	0x66, 0xcb, 0xbd, 0xbd, 0x14, 0x8d, 0x57, 0x20, 0xc9, 0x0f, 0x22, 0x37, 0xaa, 0xc7, 0x90, 0x13,
	0xc2, 0x19, 0x3d, 0xa8, 0x41, 0x92, 0xbb, 0x08, 0xf1, 0xbe, 0xf3, 0xde, 0xb2, 0x10, 0x18, 0xfa,
	0x47, 0x4b, 0x26, 0x46, 0xf6, 0xa5, 0x93, 0x12, 0x63, 0xf4, 0xad, 0x06, 0x8a, 0x0f, 0x91, 0xe2,
	0x03, 0x77, 0x49, 0x8a, 0x18, 0xb7, 0x16, 0x05, 0x32, 0x6e, 0xe3, 0xff, 0xf0, 0xdb, 0xff, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x27, 0xb7, 0xd4, 0x52, 0xdf, 0x0b, 0x00, 0x00,
}