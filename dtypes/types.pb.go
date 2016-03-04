// Code generated by protoc-gen-go.
// source: types.proto
// DO NOT EDIT!

/*
Package dtypes is a generated protocol buffer package.

It is generated from these files:
	types.proto

It has these top-level messages:
	Status
	Help
	ErrorDetails
	LongRunningResponse
	Any
	VoidRequest
	VoidResponse
*/
package dtypes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// Available Next Id 9.
type StatusCode int32

const (
	// Not an error. Returned on success.
	// Similar to HTTP Status code 2**.
	StatusCode_OK StatusCode = 0
	// The request does not have valid authentication credentials for the
	// operation. Similar to HTTP status code 401.
	StatusCode_UNAUTHORIZED StatusCode = 1
	// The request contains invalid arguments.
	// Similar to HTTP status code 400
	StatusCode_BADREQUEST StatusCode = 2
	// The caller does not have permission to execute the specified
	// operation. Reserved For further use.
	// Similar to HTTP Status code 403
	StatusCode_PERMISSION_DENIED StatusCode = 3
	// Requested entity not found.
	// Similar to HTTP status code 404
	StatusCode_NOT_FOUND StatusCode = 4
	// The operation is not implemented or is not supported/enabled
	// Similar to HTTP status code 501
	StatusCode_UNIMPLEMENTED StatusCode = 5
	// Internal errors. This means that some invariants expected by the
	// underlying system have been broken.
	// Similar to HTTP status code 500
	StatusCode_INTERNAL StatusCode = 6
	// External Server Refusus Connection or Sends Invalid Data.
	StatusCode_EXTERNAL StatusCode = 7
	// Unknown error.
	// Errors raised by APIs that do not return enough error information
	// may be converted to this error.
	// Similar to HTTP status code 500
	StatusCode_UNKNOWN_ERROR StatusCode = 8
)

var StatusCode_name = map[int32]string{
	0: "OK",
	1: "UNAUTHORIZED",
	2: "BADREQUEST",
	3: "PERMISSION_DENIED",
	4: "NOT_FOUND",
	5: "UNIMPLEMENTED",
	6: "INTERNAL",
	7: "EXTERNAL",
	8: "UNKNOWN_ERROR",
}
var StatusCode_value = map[string]int32{
	"OK":                0,
	"UNAUTHORIZED":      1,
	"BADREQUEST":        2,
	"PERMISSION_DENIED": 3,
	"NOT_FOUND":         4,
	"UNIMPLEMENTED":     5,
	"INTERNAL":          6,
	"EXTERNAL":          7,
	"UNKNOWN_ERROR":     8,
}

func (x StatusCode) String() string {
	return proto.EnumName(StatusCode_name, int32(x))
}
func (StatusCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// Next Id 5.
type Status struct {
	// Response status code of type [api.StatusCode]
	Code string `protobuf:"bytes,1,opt,name=code" json:"code,omitempty"`
	// User facing message.
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	// Optional. Help link if there is an error.
	Help *Help `protobuf:"bytes,3,opt,name=help" json:"help,omitempty"`
	// A list of messages that carry the error details.  There will be a
	// common set of message types for APIs to use.
	Details []*Any `protobuf:"bytes,4,rep,name=details" json:"details,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Status) GetHelp() *Help {
	if m != nil {
		return m.Help
	}
	return nil
}

func (m *Status) GetDetails() []*Any {
	if m != nil {
		return m.Details
	}
	return nil
}

// Provides links to documentation or for performing an out of band action.
// Next Id 3;
type Help struct {
	// Describe what link offers
	Description string `protobuf:"bytes,1,opt,name=description" json:"description,omitempty"`
	// The URL of The link.
	Url string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
}

func (m *Help) Reset()                    { *m = Help{} }
func (m *Help) String() string            { return proto.CompactTextString(m) }
func (*Help) ProtoMessage()               {}
func (*Help) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// Basic Error details message to send in response. Application specific
// error messages can be provided.
type ErrorDetails struct {
	RequestedResource string `protobuf:"bytes,1,opt,name=requested_resource" json:"requested_resource,omitempty"`
	Stacktrace        string `protobuf:"bytes,2,opt,name=stacktrace" json:"stacktrace,omitempty"`
}

func (m *ErrorDetails) Reset()                    { *m = ErrorDetails{} }
func (m *ErrorDetails) String() string            { return proto.CompactTextString(m) }
func (*ErrorDetails) ProtoMessage()               {}
func (*ErrorDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Available Next ID: 6
type LongRunningResponse struct {
	Status  *Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	TypeId  string  `protobuf:"bytes,2,opt,name=type_id" json:"type_id,omitempty"`
	JobId   string  `protobuf:"bytes,3,opt,name=job_id" json:"job_id,omitempty"`
	JobType string  `protobuf:"bytes,4,opt,name=job_type" json:"job_type,omitempty"`
}

func (m *LongRunningResponse) Reset()                    { *m = LongRunningResponse{} }
func (m *LongRunningResponse) String() string            { return proto.CompactTextString(m) }
func (*LongRunningResponse) ProtoMessage()               {}
func (*LongRunningResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LongRunningResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

// `Any` contains an arbitrary serialized message along with a URL
// that describes the type of the serialized message.
//
// JSON
// ====
// The JSON representation of an `Any` value uses the regular
// representation of the deserialized, embedded message, with an
// additional field `@type` which contains the type URL. Example:
//
//     package google.profile;
//     message Person {
//       string first_name = 1;
//       string last_name = 2;
//     }
//
//     {
//       "@type": "type.googleapis.com/google.profile.Person",
//       "firstName": <string>,
//       "lastName": <string>
//     }
//
// If the embedded message type is well-known and has a custom JSON
// representation, that representation will be embedded adding a field
// `value` which holds the custom JSON in addition to the the `@type`
// field. Example (for message [google.protobuf.Duration][google.protobuf.Duration]):
//
//     {
//       "@type": "type.googleapis.com/google.protobuf.Duration",
//       "value": "1.212s"
//     }
//
type Any struct {
	// A URL/resource name whose content describes the type of the
	// serialized message.
	//
	// For URLs which use the schema `http`, `https`, or no schema, the
	// following restrictions and interpretations apply:
	//
	// * If no schema is provided, `https` is assumed.
	// * The last segment of the URL's path must represent the fully
	//   qualified name of the type (as in `path/google.protobuf.Duration`).
	// * An HTTP GET on the URL must yield a [google.protobuf.Type][google.protobuf.Type]
	//   value in binary format, or produce an error.
	// * Applications are allowed to cache lookup results based on the
	//   URL, or have them precompiled into a binary to avoid any
	//   lookup. Therefore, binary compatibility needs to be preserved
	//   on changes to types. (Use versioned type names to manage
	//   breaking changes.)
	//
	// Schemas other than `http`, `https` (or the empty schema) might be
	// used with implementation specific semantics.
	//
	TypeUrl string `protobuf:"bytes,1,opt,name=type_url" json:"type_url,omitempty"`
	// Must be valid serialized data of the above specified type.
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Any) Reset()                    { *m = Any{} }
func (m *Any) String() string            { return proto.CompactTextString(m) }
func (*Any) ProtoMessage()               {}
func (*Any) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// Void Requests and response to use with other types.
type VoidRequest struct {
}

func (m *VoidRequest) Reset()                    { *m = VoidRequest{} }
func (m *VoidRequest) String() string            { return proto.CompactTextString(m) }
func (*VoidRequest) ProtoMessage()               {}
func (*VoidRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type VoidResponse struct {
	Status *Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *VoidResponse) Reset()                    { *m = VoidResponse{} }
func (m *VoidResponse) String() string            { return proto.CompactTextString(m) }
func (*VoidResponse) ProtoMessage()               {}
func (*VoidResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *VoidResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*Status)(nil), "dtypes.Status")
	proto.RegisterType((*Help)(nil), "dtypes.Help")
	proto.RegisterType((*ErrorDetails)(nil), "dtypes.ErrorDetails")
	proto.RegisterType((*LongRunningResponse)(nil), "dtypes.LongRunningResponse")
	proto.RegisterType((*Any)(nil), "dtypes.Any")
	proto.RegisterType((*VoidRequest)(nil), "dtypes.VoidRequest")
	proto.RegisterType((*VoidResponse)(nil), "dtypes.VoidResponse")
	proto.RegisterEnum("dtypes.StatusCode", StatusCode_name, StatusCode_value)
}

var fileDescriptor0 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x69, 0x93, 0xa5, 0xdd, 0x4b, 0x3a, 0x32, 0x4f, 0x48, 0x15, 0x42, 0x68, 0xca, 0x01,
	0x4d, 0x1c, 0x7a, 0x18, 0x77, 0xa4, 0x42, 0x8c, 0x16, 0xad, 0x75, 0x86, 0xdb, 0x00, 0xe2, 0x12,
	0x65, 0x89, 0xb5, 0x05, 0x4a, 0x1c, 0x6c, 0x07, 0x69, 0xff, 0x0c, 0x7f, 0x2b, 0x2f, 0x71, 0x7a,
	0xe7, 0x54, 0xff, 0xfc, 0xbe, 0xd7, 0xef, 0xfb, 0x14, 0x83, 0x6f, 0x9e, 0x5a, 0xa1, 0x57, 0xad,
	0x92, 0x46, 0x12, 0xaf, 0x1a, 0x28, 0x2a, 0xc0, 0xdb, 0x99, 0xc2, 0x74, 0x9a, 0x04, 0xe0, 0x96,
	0xb2, 0x12, 0xcb, 0xc9, 0xe5, 0xe4, 0xea, 0x94, 0x3c, 0x87, 0xd9, 0x2f, 0xa1, 0x75, 0xf1, 0x20,
	0x96, 0xd3, 0xe1, 0xe2, 0x25, 0xb8, 0x8f, 0xe2, 0xd0, 0x2e, 0x1d, 0x24, 0xff, 0x3a, 0x58, 0xd9,
	0xfd, 0xd5, 0x0d, 0xde, 0x91, 0x57, 0x30, 0xab, 0x84, 0x29, 0xea, 0x83, 0x5e, 0xba, 0x97, 0x0e,
	0x8e, 0xfd, 0xe3, 0x78, 0xdd, 0x3c, 0x45, 0x57, 0xe0, 0x0e, 0xaa, 0x0b, 0xf0, 0x2b, 0xa1, 0x4b,
	0x55, 0xb7, 0xa6, 0x96, 0xcd, 0xe8, 0xe3, 0x83, 0xd3, 0xa9, 0x83, 0xf5, 0x88, 0xde, 0x43, 0x40,
	0x95, 0x92, 0x2a, 0xb6, 0x7f, 0x86, 0x9e, 0x44, 0x89, 0xdf, 0x9d, 0xd0, 0x46, 0x54, 0xb9, 0x12,
	0x5a, 0x76, 0xaa, 0x3c, 0x06, 0x24, 0x00, 0xda, 0x14, 0xe5, 0x4f, 0xa3, 0x8a, 0x72, 0xcc, 0x18,
	0x3d, 0xc2, 0xc5, 0x46, 0x36, 0x0f, 0xbc, 0x6b, 0x9a, 0x1a, 0x7f, 0x84, 0x6e, 0x65, 0xa3, 0x05,
	0x79, 0x0d, 0x9e, 0x1e, 0x3a, 0x0e, 0xab, 0xfe, 0xf5, 0xd9, 0x31, 0xdd, 0xd8, 0x1c, 0xbb, 0xf6,
	0x9c, 0xd7, 0xd5, 0xd8, 0xf5, 0x0c, 0xbc, 0x1f, 0xf2, 0xbe, 0x67, 0x67, 0xe0, 0x10, 0xe6, 0x3d,
	0xf7, 0x22, 0x2c, 0xd8, 0x3b, 0xbd, 0x01, 0x07, 0xab, 0xf5, 0x83, 0x61, 0xb3, 0xaf, 0x60, 0x63,
	0x2d, 0xe0, 0xe4, 0x4f, 0x71, 0xe8, 0x6c, 0xa2, 0x20, 0x5a, 0x80, 0xff, 0x45, 0xd6, 0x15, 0xb7,
	0x2d, 0xa2, 0x15, 0x04, 0x16, 0xff, 0x2f, 0xd9, 0xdb, 0xbf, 0x13, 0x00, 0x7b, 0xfc, 0x88, 0x9f,
	0x86, 0x78, 0x30, 0x4d, 0x6f, 0xc3, 0x67, 0x68, 0x1b, 0x64, 0x6c, 0x9d, 0xed, 0x6f, 0x52, 0x9e,
	0x7c, 0xa7, 0x71, 0x38, 0xc1, 0xc4, 0xf0, 0x61, 0x1d, 0x73, 0xfa, 0x39, 0xa3, 0xbb, 0x7d, 0x38,
	0x25, 0x2f, 0xe0, 0xfc, 0x8e, 0xf2, 0x6d, 0xb2, 0xdb, 0x25, 0x29, 0xcb, 0x63, 0xca, 0x12, 0x94,
	0x39, 0x98, 0xee, 0x94, 0xa5, 0xfb, 0xfc, 0x53, 0x9a, 0xb1, 0x38, 0x74, 0xc9, 0x39, 0x2c, 0x32,
	0x96, 0x6c, 0xef, 0x36, 0x74, 0x4b, 0xd9, 0x1e, 0x15, 0x27, 0xf8, 0x0a, 0xe6, 0x09, 0x1e, 0x39,
	0x5b, 0x6f, 0x42, 0xaf, 0x27, 0xfa, 0x6d, 0xa4, 0x99, 0x95, 0xdf, 0xb2, 0xf4, 0x2b, 0xcb, 0x29,
	0xe7, 0x29, 0x0f, 0xe7, 0xf7, 0xde, 0xf0, 0x9a, 0xde, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x36,
	0x75, 0x4a, 0xb4, 0x5c, 0x02, 0x00, 0x00,
}
