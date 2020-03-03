// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types/access_path.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AccessPath struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Path                 []byte   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccessPath) Reset()         { *m = AccessPath{} }
func (m *AccessPath) String() string { return proto.CompactTextString(m) }
func (*AccessPath) ProtoMessage()    {}
func (*AccessPath) Descriptor() ([]byte, []int) {
	return fileDescriptor_12123bd84945f471, []int{0}
}

func (m *AccessPath) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessPath.Unmarshal(m, b)
}
func (m *AccessPath) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessPath.Marshal(b, m, deterministic)
}
func (m *AccessPath) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessPath.Merge(m, src)
}
func (m *AccessPath) XXX_Size() int {
	return xxx_messageInfo_AccessPath.Size(m)
}
func (m *AccessPath) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessPath.DiscardUnknown(m)
}

var xxx_messageInfo_AccessPath proto.InternalMessageInfo

func (m *AccessPath) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *AccessPath) GetPath() []byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func init() {
	proto.RegisterType((*AccessPath)(nil), "types.AccessPath")
}

func init() { proto.RegisterFile("types/access_path.proto", fileDescriptor_12123bd84945f471) }

var fileDescriptor_12123bd84945f471 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0xa9, 0x2c, 0x48,
	0x2d, 0xd6, 0x4f, 0x4c, 0x4e, 0x4e, 0x2d, 0x2e, 0x8e, 0x2f, 0x48, 0x2c, 0xc9, 0xd0, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x4b, 0x28, 0x59, 0x71, 0x71, 0x39, 0x82, 0xe5, 0x02, 0x12,
	0x4b, 0x32, 0x84, 0x24, 0xb8, 0xd8, 0x13, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x78, 0x82, 0x60, 0x5c, 0x21, 0x21, 0x2e, 0x16, 0x90, 0x66, 0x09, 0x26, 0xb0, 0x30,
	0x98, 0xed, 0x64, 0x12, 0x65, 0x94, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab,
	0x9f, 0x93, 0x99, 0x54, 0x94, 0xa8, 0x5b, 0x9c, 0x99, 0x5e, 0x0c, 0x63, 0xa6, 0x64, 0xeb, 0xa6,
	0xe7, 0x43, 0x38, 0xfa, 0x45, 0x05, 0xc9, 0xfa, 0x60, 0x1b, 0x93, 0xd8, 0xc0, 0xf6, 0x1b, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x47, 0x32, 0xf8, 0x9a, 0x00, 0x00, 0x00,
}