// Code generated by protoc-gen-go.
// source: volume.proto
// DO NOT EDIT!

package glusterfs

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

type VolumeCreateRequest struct {
	KubeCluster      string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	GlusterfsCluster string `protobuf:"bytes,2,opt,name=glusterfs_cluster,json=glusterfsCluster" json:"glusterfs_cluster,omitempty"`
	Volume           string `protobuf:"bytes,3,opt,name=volume" json:"volume,omitempty"`
}

func (m *VolumeCreateRequest) Reset()                    { *m = VolumeCreateRequest{} }
func (m *VolumeCreateRequest) String() string            { return proto.CompactTextString(m) }
func (*VolumeCreateRequest) ProtoMessage()               {}
func (*VolumeCreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type VolumeDeleteRequest struct {
	KubeCluster      string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	KubeNamespace    string `protobuf:"bytes,2,opt,name=kube_namespace,json=kubeNamespace" json:"kube_namespace,omitempty"`
	GlusterfsCluster string `protobuf:"bytes,3,opt,name=glusterfs_cluster,json=glusterfsCluster" json:"glusterfs_cluster,omitempty"`
	Volume           string `protobuf:"bytes,4,opt,name=volume" json:"volume,omitempty"`
}

func (m *VolumeDeleteRequest) Reset()                    { *m = VolumeDeleteRequest{} }
func (m *VolumeDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*VolumeDeleteRequest) ProtoMessage()               {}
func (*VolumeDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type VolumeListRequest struct {
	KubeCluster      string `protobuf:"bytes,1,opt,name=kube_cluster,json=kubeCluster" json:"kube_cluster,omitempty"`
	KubeNamespace    string `protobuf:"bytes,2,opt,name=kube_namespace,json=kubeNamespace" json:"kube_namespace,omitempty"`
	GlusterfsCluster string `protobuf:"bytes,3,opt,name=glusterfs_cluster,json=glusterfsCluster" json:"glusterfs_cluster,omitempty"`
}

func (m *VolumeListRequest) Reset()                    { *m = VolumeListRequest{} }
func (m *VolumeListRequest) String() string            { return proto.CompactTextString(m) }
func (*VolumeListRequest) ProtoMessage()               {}
func (*VolumeListRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type VolumeListResponse struct {
	Status  *dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Volumes []*Volume      `protobuf:"bytes,2,rep,name=volumes" json:"volumes,omitempty"`
}

func (m *VolumeListResponse) Reset()                    { *m = VolumeListResponse{} }
func (m *VolumeListResponse) String() string            { return proto.CompactTextString(m) }
func (*VolumeListResponse) ProtoMessage()               {}
func (*VolumeListResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *VolumeListResponse) GetStatus() *dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VolumeListResponse) GetVolumes() []*Volume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

type Volume struct {
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
}

func (m *Volume) Reset()                    { *m = Volume{} }
func (m *Volume) String() string            { return proto.CompactTextString(m) }
func (*Volume) ProtoMessage()               {}
func (*Volume) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func init() {
	proto.RegisterType((*VolumeCreateRequest)(nil), "glusterfs.VolumeCreateRequest")
	proto.RegisterType((*VolumeDeleteRequest)(nil), "glusterfs.VolumeDeleteRequest")
	proto.RegisterType((*VolumeListRequest)(nil), "glusterfs.VolumeListRequest")
	proto.RegisterType((*VolumeListResponse)(nil), "glusterfs.VolumeListResponse")
	proto.RegisterType((*Volume)(nil), "glusterfs.Volume")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Volumes service

type VolumesClient interface {
	List(ctx context.Context, in *VolumeListRequest, opts ...grpc.CallOption) (*VolumeListResponse, error)
	Create(ctx context.Context, in *VolumeCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
	Delete(ctx context.Context, in *VolumeDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error)
}

type volumesClient struct {
	cc *grpc.ClientConn
}

func NewVolumesClient(cc *grpc.ClientConn) VolumesClient {
	return &volumesClient{cc}
}

func (c *volumesClient) List(ctx context.Context, in *VolumeListRequest, opts ...grpc.CallOption) (*VolumeListResponse, error) {
	out := new(VolumeListResponse)
	err := grpc.Invoke(ctx, "/glusterfs.Volumes/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumesClient) Create(ctx context.Context, in *VolumeCreateRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/glusterfs.Volumes/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *volumesClient) Delete(ctx context.Context, in *VolumeDeleteRequest, opts ...grpc.CallOption) (*dtypes.VoidResponse, error) {
	out := new(dtypes.VoidResponse)
	err := grpc.Invoke(ctx, "/glusterfs.Volumes/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Volumes service

type VolumesServer interface {
	List(context.Context, *VolumeListRequest) (*VolumeListResponse, error)
	Create(context.Context, *VolumeCreateRequest) (*dtypes.VoidResponse, error)
	Delete(context.Context, *VolumeDeleteRequest) (*dtypes.VoidResponse, error)
}

func RegisterVolumesServer(s *grpc.Server, srv VolumesServer) {
	s.RegisterService(&_Volumes_serviceDesc, srv)
}

func _Volumes_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glusterfs.Volumes/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumesServer).List(ctx, req.(*VolumeListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Volumes_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glusterfs.Volumes/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumesServer).Create(ctx, req.(*VolumeCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Volumes_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VolumeDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VolumesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glusterfs.Volumes/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VolumesServer).Delete(ctx, req.(*VolumeDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Volumes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "glusterfs.Volumes",
	HandlerType: (*VolumesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Volumes_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Volumes_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Volumes_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("volume.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x94, 0x41, 0x4b, 0xe3, 0x40,
	0x14, 0xc7, 0x49, 0x53, 0x52, 0x3a, 0xed, 0x96, 0xed, 0xec, 0x52, 0x4a, 0xe8, 0x2e, 0xbb, 0x81,
	0x5d, 0x96, 0x2e, 0x24, 0xda, 0x5e, 0xc4, 0x83, 0x97, 0x7a, 0x14, 0x0f, 0x11, 0x8a, 0x78, 0xd1,
	0x69, 0x3a, 0xd6, 0x60, 0x9a, 0x89, 0x9d, 0x49, 0x41, 0x4a, 0x3c, 0x78, 0xf4, 0xea, 0x57, 0xd0,
	0x83, 0x9f, 0xc7, 0xb3, 0x37, 0x3f, 0x88, 0xc9, 0xbc, 0x24, 0xb4, 0xb4, 0x45, 0x14, 0xc4, 0x4b,
	0xe9, 0xfc, 0xdf, 0x7f, 0xde, 0xfb, 0xcd, 0x9b, 0x37, 0x41, 0xd5, 0x29, 0xf3, 0xc2, 0x31, 0x35,
	0x83, 0x09, 0x13, 0x0c, 0x97, 0x47, 0x5e, 0xc8, 0x05, 0x9d, 0x9c, 0x72, 0xbd, 0x35, 0x62, 0x6c,
	0xe4, 0x51, 0x8b, 0x04, 0xae, 0x45, 0x7c, 0x9f, 0x09, 0x22, 0x5c, 0xe6, 0x73, 0x30, 0xea, 0x8d,
	0x44, 0x1e, 0x8a, 0xcb, 0x80, 0x72, 0x4b, 0xfe, 0x82, 0x6e, 0x44, 0xe8, 0x5b, 0x5f, 0x26, 0xec,
	0x4d, 0x28, 0x11, 0xd4, 0xa6, 0x17, 0x21, 0xe5, 0x02, 0xff, 0x46, 0xd5, 0xf3, 0x70, 0x40, 0x8f,
	0x1d, 0x48, 0xdf, 0x54, 0x7e, 0x29, 0xff, 0xca, 0x76, 0x25, 0xd1, 0x7a, 0x20, 0xe1, 0xff, 0xa8,
	0x9e, 0x17, 0xcf, 0x7d, 0x05, 0xe9, 0xfb, 0x9a, 0x07, 0x32, 0x73, 0x03, 0x69, 0xc0, 0xdd, 0x54,
	0xa5, 0x23, 0x5d, 0x19, 0x77, 0x4a, 0x56, 0x7f, 0x97, 0x7a, 0xf4, 0x4d, 0xf5, 0xff, 0xa0, 0x9a,
	0xb4, 0xf8, 0x64, 0x4c, 0x79, 0x40, 0x1c, 0x9a, 0x16, 0xff, 0x92, 0xa8, 0xfb, 0x99, 0xb8, 0x1a,
	0x53, 0x7d, 0x15, 0xb3, 0xb8, 0x80, 0x79, 0xa3, 0xa0, 0x3a, 0x60, 0xee, 0xb9, 0x5c, 0x7c, 0x2e,
	0xa4, 0xe1, 0x22, 0x3c, 0xcf, 0xc2, 0x83, 0xf8, 0x96, 0x29, 0xfe, 0x8b, 0x34, 0x1e, 0x5f, 0x79,
	0xc8, 0x25, 0x46, 0xa5, 0x53, 0x33, 0xe1, 0xb6, 0xcd, 0x03, 0xa9, 0xda, 0x69, 0x34, 0x2e, 0x55,
	0x82, 0x43, 0xf1, 0x18, 0x45, 0x8d, 0x8d, 0x75, 0x33, 0xaf, 0x60, 0x42, 0x5e, 0x3b, 0x73, 0x18,
	0x2d, 0xa4, 0x81, 0x84, 0x31, 0x2a, 0x06, 0x44, 0x9c, 0xa5, 0x67, 0x94, 0xff, 0x3b, 0x4f, 0x2a,
	0x2a, 0x41, 0x98, 0xe3, 0x7b, 0x05, 0x15, 0x13, 0x1e, 0xdc, 0x5a, 0x4a, 0x37, 0xd7, 0x32, 0xfd,
	0xc7, 0x9a, 0x28, 0x1c, 0xc2, 0x38, 0xb9, 0x7e, 0x7c, 0xbe, 0x2d, 0x1c, 0xe1, 0xc3, 0x78, 0x8a,
	0x03, 0xee, 0xb0, 0x21, 0x8c, 0x73, 0xbe, 0xc7, 0x9a, 0x6e, 0x98, 0x9b, 0x56, 0x4a, 0x67, 0xcd,
	0xe6, 0xbb, 0x1f, 0xa5, 0xcb, 0xbc, 0xd3, 0xb1, 0xb0, 0xd4, 0xd3, 0x08, 0x5f, 0x21, 0x0d, 0x26,
	0x1d, 0xff, 0x5c, 0x42, 0x59, 0x78, 0x02, 0xfa, 0xf7, 0xac, 0x81, 0x7d, 0xe6, 0x0e, 0x73, 0xc2,
	0x1d, 0x49, 0xb8, 0xa5, 0x77, 0xdf, 0x41, 0xb8, 0xad, 0xb4, 0xf1, 0x83, 0x82, 0x34, 0x18, 0xf5,
	0x15, 0x00, 0x0b, 0x6f, 0x60, 0x0d, 0x80, 0x2b, 0x01, 0x9c, 0x36, 0xf9, 0xa8, 0x16, 0x59, 0x33,
	0xd8, 0x1c, 0x0d, 0x34, 0xf9, 0x89, 0xe8, 0xbe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x58, 0xf9, 0xc9,
	0x45, 0x73, 0x04, 0x00, 0x00,
}
