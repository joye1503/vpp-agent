// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: l4.proto

package l4

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Layer 4 VPP features settings
type L4Features struct {
	Enabled              bool     `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *L4Features) Reset()         { *m = L4Features{} }
func (m *L4Features) String() string { return proto.CompactTextString(m) }
func (*L4Features) ProtoMessage()    {}
func (*L4Features) Descriptor() ([]byte, []int) {
	return fileDescriptor_l4_3288b4ffe3b5800c, []int{0}
}
func (m *L4Features) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_L4Features.Unmarshal(m, b)
}
func (m *L4Features) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_L4Features.Marshal(b, m, deterministic)
}
func (dst *L4Features) XXX_Merge(src proto.Message) {
	xxx_messageInfo_L4Features.Merge(dst, src)
}
func (m *L4Features) XXX_Size() int {
	return xxx_messageInfo_L4Features.Size(m)
}
func (m *L4Features) XXX_DiscardUnknown() {
	xxx_messageInfo_L4Features.DiscardUnknown(m)
}

var xxx_messageInfo_L4Features proto.InternalMessageInfo

func (m *L4Features) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

// Application namespaces
type AppNamespaces struct {
	AppNamespaces        []*AppNamespaces_AppNamespace `protobuf:"bytes,100,rep,name=app_namespaces,json=appNamespaces,proto3" json:"app_namespaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *AppNamespaces) Reset()         { *m = AppNamespaces{} }
func (m *AppNamespaces) String() string { return proto.CompactTextString(m) }
func (*AppNamespaces) ProtoMessage()    {}
func (*AppNamespaces) Descriptor() ([]byte, []int) {
	return fileDescriptor_l4_3288b4ffe3b5800c, []int{1}
}
func (m *AppNamespaces) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppNamespaces.Unmarshal(m, b)
}
func (m *AppNamespaces) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppNamespaces.Marshal(b, m, deterministic)
}
func (dst *AppNamespaces) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppNamespaces.Merge(dst, src)
}
func (m *AppNamespaces) XXX_Size() int {
	return xxx_messageInfo_AppNamespaces.Size(m)
}
func (m *AppNamespaces) XXX_DiscardUnknown() {
	xxx_messageInfo_AppNamespaces.DiscardUnknown(m)
}

var xxx_messageInfo_AppNamespaces proto.InternalMessageInfo

func (m *AppNamespaces) GetAppNamespaces() []*AppNamespaces_AppNamespace {
	if m != nil {
		return m.AppNamespaces
	}
	return nil
}

type AppNamespaces_AppNamespace struct {
	NamespaceId          string   `protobuf:"bytes,1,opt,name=namespace_id,json=namespaceId,proto3" json:"namespace_id,omitempty"`
	Secret               uint64   `protobuf:"varint,2,opt,name=secret,proto3" json:"secret,omitempty"`
	Interface            string   `protobuf:"bytes,3,opt,name=interface,proto3" json:"interface,omitempty"`
	Ipv4FibId            uint32   `protobuf:"varint,4,opt,name=ipv4_fib_id,json=ipv4FibId,proto3" json:"ipv4_fib_id,omitempty"`
	Ipv6FibId            uint32   `protobuf:"varint,6,opt,name=ipv6_fib_id,json=ipv6FibId,proto3" json:"ipv6_fib_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppNamespaces_AppNamespace) Reset()         { *m = AppNamespaces_AppNamespace{} }
func (m *AppNamespaces_AppNamespace) String() string { return proto.CompactTextString(m) }
func (*AppNamespaces_AppNamespace) ProtoMessage()    {}
func (*AppNamespaces_AppNamespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_l4_3288b4ffe3b5800c, []int{1, 0}
}
func (m *AppNamespaces_AppNamespace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppNamespaces_AppNamespace.Unmarshal(m, b)
}
func (m *AppNamespaces_AppNamespace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppNamespaces_AppNamespace.Marshal(b, m, deterministic)
}
func (dst *AppNamespaces_AppNamespace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppNamespaces_AppNamespace.Merge(dst, src)
}
func (m *AppNamespaces_AppNamespace) XXX_Size() int {
	return xxx_messageInfo_AppNamespaces_AppNamespace.Size(m)
}
func (m *AppNamespaces_AppNamespace) XXX_DiscardUnknown() {
	xxx_messageInfo_AppNamespaces_AppNamespace.DiscardUnknown(m)
}

var xxx_messageInfo_AppNamespaces_AppNamespace proto.InternalMessageInfo

func (m *AppNamespaces_AppNamespace) GetNamespaceId() string {
	if m != nil {
		return m.NamespaceId
	}
	return ""
}

func (m *AppNamespaces_AppNamespace) GetSecret() uint64 {
	if m != nil {
		return m.Secret
	}
	return 0
}

func (m *AppNamespaces_AppNamespace) GetInterface() string {
	if m != nil {
		return m.Interface
	}
	return ""
}

func (m *AppNamespaces_AppNamespace) GetIpv4FibId() uint32 {
	if m != nil {
		return m.Ipv4FibId
	}
	return 0
}

func (m *AppNamespaces_AppNamespace) GetIpv6FibId() uint32 {
	if m != nil {
		return m.Ipv6FibId
	}
	return 0
}

func init() {
	proto.RegisterType((*L4Features)(nil), "l4.L4Features")
	proto.RegisterType((*AppNamespaces)(nil), "l4.AppNamespaces")
	proto.RegisterType((*AppNamespaces_AppNamespace)(nil), "l4.AppNamespaces.AppNamespace")
}

func init() { proto.RegisterFile("l4.proto", fileDescriptor_l4_3288b4ffe3b5800c) }

var fileDescriptor_l4_3288b4ffe3b5800c = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0x31, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xca, 0x31, 0x51, 0x52, 0xe3, 0xe2, 0xf2, 0x31, 0x71, 0x4b,
	0x4d, 0x2c, 0x29, 0x2d, 0x4a, 0x2d, 0x16, 0x92, 0xe0, 0x62, 0x4f, 0xcd, 0x4b, 0x4c, 0xca, 0x49,
	0x4d, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x08, 0x82, 0x71, 0x95, 0x1a, 0x98, 0xb8, 0x78, 0x1d,
	0x0b, 0x0a, 0xfc, 0x12, 0x73, 0x53, 0x8b, 0x0b, 0x12, 0x93, 0x53, 0x8b, 0x85, 0x5c, 0xb9, 0xf8,
	0x12, 0x0b, 0x0a, 0xe2, 0xf3, 0xe0, 0x22, 0x12, 0x29, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x72, 0x7a,
	0x39, 0x26, 0x7a, 0x28, 0x4a, 0x51, 0x78, 0x41, 0xbc, 0x89, 0xc8, 0x72, 0x52, 0xcb, 0x19, 0xb9,
	0x78, 0x90, 0xe5, 0x85, 0x14, 0xb9, 0x78, 0xe0, 0x66, 0xc6, 0x67, 0x42, 0x1c, 0xc2, 0x19, 0xc4,
	0x0d, 0x17, 0xf3, 0x4c, 0x11, 0x12, 0xe3, 0x62, 0x2b, 0x4e, 0x4d, 0x2e, 0x4a, 0x2d, 0x91, 0x60,
	0x52, 0x60, 0xd4, 0x60, 0x09, 0x82, 0xf2, 0x84, 0x64, 0xb8, 0x38, 0x33, 0xf3, 0x4a, 0x52, 0x8b,
	0xd2, 0x12, 0x93, 0x53, 0x25, 0x98, 0xc1, 0xfa, 0x10, 0x02, 0x42, 0x72, 0x5c, 0xdc, 0x99, 0x05,
	0x65, 0x26, 0xf1, 0x69, 0x99, 0x49, 0x20, 0x73, 0x59, 0x14, 0x18, 0x35, 0x78, 0x83, 0x38, 0x41,
	0x42, 0x6e, 0x99, 0x49, 0x9e, 0x29, 0x50, 0x79, 0x33, 0x98, 0x3c, 0x1b, 0x5c, 0xde, 0x0c, 0x2c,
	0x9f, 0xc4, 0x06, 0x0e, 0x35, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x07, 0xd5, 0xf9,
	0x41, 0x01, 0x00, 0x00,
}
