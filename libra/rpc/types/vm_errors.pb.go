// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types/vm_errors.proto

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

// The statuses and errors produced by the VM can be categorized into a
// couple different types:
// 1. Validation Statuses: all the errors that can (/should) be
//    the result of executing the prologue -- these are primarily used by
//    the vm validator and AC.
// 2. Verification Errors: errors that are the result of performing
//    bytecode verification (happens at the time of publishing).
// 3. VM Invariant Errors: errors that arise from an internal invariant of
//    the VM being violated. These signify a problem with either the VM or
//    bytecode verifier.
// 4. Binary Errors: errors that can occur during the process of
//    deserialization of a transaction.
// 5. Runtime Statuses: errors that can arise from the execution of a
//    transaction (assuming the prologue executes without error). These are
//    errors that can occur during execution due to things such as division
//    by zero, running out of gas, etc. These do not signify an issue with
//    the VM.
type VMStatus struct {
	// e.g. assertion violation, out of gas
	MajorStatus uint64 `protobuf:"varint,1,opt,name=major_status,json=majorStatus,proto3" json:"major_status,omitempty"`
	// Any substatus code. e.g. assertion error number
	HasSubStatus         bool     `protobuf:"varint,2,opt,name=has_sub_status,json=hasSubStatus,proto3" json:"has_sub_status,omitempty"`
	SubStatus            uint64   `protobuf:"varint,3,opt,name=sub_status,json=subStatus,proto3" json:"sub_status,omitempty"`
	HasMessage           bool     `protobuf:"varint,4,opt,name=has_message,json=hasMessage,proto3" json:"has_message,omitempty"`
	Message              string   `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VMStatus) Reset()         { *m = VMStatus{} }
func (m *VMStatus) String() string { return proto.CompactTextString(m) }
func (*VMStatus) ProtoMessage()    {}
func (*VMStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_42c5ee2f0a4a19be, []int{0}
}

func (m *VMStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMStatus.Unmarshal(m, b)
}
func (m *VMStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMStatus.Marshal(b, m, deterministic)
}
func (m *VMStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMStatus.Merge(m, src)
}
func (m *VMStatus) XXX_Size() int {
	return xxx_messageInfo_VMStatus.Size(m)
}
func (m *VMStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_VMStatus.DiscardUnknown(m)
}

var xxx_messageInfo_VMStatus proto.InternalMessageInfo

func (m *VMStatus) GetMajorStatus() uint64 {
	if m != nil {
		return m.MajorStatus
	}
	return 0
}

func (m *VMStatus) GetHasSubStatus() bool {
	if m != nil {
		return m.HasSubStatus
	}
	return false
}

func (m *VMStatus) GetSubStatus() uint64 {
	if m != nil {
		return m.SubStatus
	}
	return 0
}

func (m *VMStatus) GetHasMessage() bool {
	if m != nil {
		return m.HasMessage
	}
	return false
}

func (m *VMStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*VMStatus)(nil), "types.VMStatus")
}

func init() { proto.RegisterFile("types/vm_errors.proto", fileDescriptor_42c5ee2f0a4a19be) }

var fileDescriptor_42c5ee2f0a4a19be = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x89, 0xb6, 0xda, 0x4e, 0x8b, 0x87, 0x05, 0x61, 0x11, 0xc5, 0x55, 0x3c, 0xec, 0xa5,
	0x5d, 0x50, 0x9f, 0xc0, 0x7b, 0x2f, 0x5b, 0xf0, 0xe0, 0x25, 0x4c, 0x6a, 0x48, 0x56, 0x1b, 0xb3,
	0xcc, 0x24, 0x82, 0x0f, 0xe5, 0x3b, 0x8a, 0x49, 0x23, 0xde, 0xe6, 0xff, 0xff, 0x8f, 0x0f, 0x06,
	0xce, 0xc3, 0xd7, 0xa8, 0xb9, 0xfb, 0x74, 0x52, 0x13, 0x79, 0xe2, 0xf5, 0x48, 0x3e, 0xf8, 0x6a,
	0x9a, 0xea, 0x8b, 0xcb, 0xbc, 0xee, 0xf1, 0xc3, 0x44, 0x34, 0x5a, 0x72, 0xf0, 0x84, 0x46, 0x67,
	0xe8, 0xf6, 0x5b, 0xc0, 0xec, 0x79, 0xb3, 0x0d, 0x18, 0x22, 0x57, 0x37, 0xb0, 0x74, 0xf8, 0xe6,
	0x49, 0x72, 0xca, 0xb5, 0x68, 0x44, 0x3b, 0xe9, 0x17, 0xa9, 0x3b, 0x20, 0x77, 0x70, 0x66, 0x91,
	0x25, 0x47, 0x55, 0xa0, 0xa3, 0x46, 0xb4, 0xb3, 0x7e, 0x69, 0x91, 0xb7, 0x51, 0x1d, 0xa8, 0x2b,
	0x80, 0x7f, 0xc4, 0x71, 0xd2, 0xcc, 0xf9, 0x6f, 0xbe, 0x86, 0xc5, 0xaf, 0xc4, 0x69, 0x66, 0x34,
	0xba, 0x9e, 0x24, 0x03, 0x58, 0xe4, 0x4d, 0x6e, 0xaa, 0x1a, 0x4e, 0xcb, 0x38, 0x6d, 0x44, 0x3b,
	0xef, 0x4b, 0x7c, 0x7a, 0x7c, 0xb9, 0x37, 0x43, 0xb0, 0x51, 0xad, 0x77, 0xde, 0x75, 0xfb, 0x41,
	0x11, 0xae, 0x78, 0x30, 0x5c, 0xce, 0xd7, 0xf7, 0x95, 0xf1, 0x39, 0x74, 0x34, 0xee, 0xba, 0xf4,
	0xbc, 0x3a, 0x49, 0xcf, 0x3e, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0x26, 0x9d, 0xb6, 0x55, 0x2a,
	0x01, 0x00, 0x00,
}