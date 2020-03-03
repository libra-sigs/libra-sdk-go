// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types/transaction_info.proto

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

// `TransactionInfo` is the object we store in the transaction accumulator. It
// consists of the transaction as well as the execution result of this
// transaction. This are later returned to the client so that a client can
// validate the tree
type TransactionInfo struct {
	// Hash of the transaction that is stored.
	TransactionHash []byte `protobuf:"bytes,1,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"`
	// The root hash of Sparse Merkle Tree describing the world state at the end
	// of this transaction
	StateRootHash []byte `protobuf:"bytes,2,opt,name=state_root_hash,json=stateRootHash,proto3" json:"state_root_hash,omitempty"`
	// The root hash of Merkle Accumulator storing all events emitted during this
	// transaction.
	EventRootHash []byte `protobuf:"bytes,3,opt,name=event_root_hash,json=eventRootHash,proto3" json:"event_root_hash,omitempty"`
	// The amount of gas used by this transaction.
	GasUsed uint64 `protobuf:"varint,4,opt,name=gas_used,json=gasUsed,proto3" json:"gas_used,omitempty"`
	// The major status of executing this transaction.
	MajorStatus          uint64   `protobuf:"varint,5,opt,name=major_status,json=majorStatus,proto3" json:"major_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionInfo) Reset()         { *m = TransactionInfo{} }
func (m *TransactionInfo) String() string { return proto.CompactTextString(m) }
func (*TransactionInfo) ProtoMessage()    {}
func (*TransactionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_af679f6db9619bd9, []int{0}
}

func (m *TransactionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionInfo.Unmarshal(m, b)
}
func (m *TransactionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionInfo.Marshal(b, m, deterministic)
}
func (m *TransactionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionInfo.Merge(m, src)
}
func (m *TransactionInfo) XXX_Size() int {
	return xxx_messageInfo_TransactionInfo.Size(m)
}
func (m *TransactionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionInfo proto.InternalMessageInfo

func (m *TransactionInfo) GetTransactionHash() []byte {
	if m != nil {
		return m.TransactionHash
	}
	return nil
}

func (m *TransactionInfo) GetStateRootHash() []byte {
	if m != nil {
		return m.StateRootHash
	}
	return nil
}

func (m *TransactionInfo) GetEventRootHash() []byte {
	if m != nil {
		return m.EventRootHash
	}
	return nil
}

func (m *TransactionInfo) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *TransactionInfo) GetMajorStatus() uint64 {
	if m != nil {
		return m.MajorStatus
	}
	return 0
}

func init() {
	proto.RegisterType((*TransactionInfo)(nil), "types.TransactionInfo")
}

func init() { proto.RegisterFile("types/transaction_info.proto", fileDescriptor_af679f6db9619bd9) }

var fileDescriptor_af679f6db9619bd9 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x15, 0x68, 0x01, 0x99, 0xa2, 0xa0, 0x4c, 0x41, 0x62, 0x28, 0x0c, 0xa8, 0x0c, 0xad,
	0x25, 0xe0, 0x09, 0x98, 0x60, 0x0d, 0xb0, 0xb0, 0x58, 0x97, 0xc4, 0xb5, 0x0d, 0xd4, 0x17, 0xf9,
	0x2e, 0x48, 0xbc, 0x22, 0x4f, 0x85, 0xb8, 0x28, 0x90, 0xed, 0xee, 0xfb, 0xbf, 0xbb, 0xe1, 0x57,
	0xe7, 0xfc, 0xd5, 0x59, 0xd2, 0x9c, 0x20, 0x12, 0x34, 0x1c, 0x30, 0x9a, 0x10, 0xb7, 0xb8, 0xe9,
	0x12, 0x32, 0x16, 0x73, 0x49, 0x2f, 0xbf, 0x33, 0x95, 0x3f, 0xff, 0x1b, 0x8f, 0x71, 0x8b, 0xc5,
	0xb5, 0x3a, 0x9d, 0x1e, 0x79, 0x20, 0x5f, 0x66, 0xcb, 0x6c, 0xb5, 0xa8, 0xf2, 0x09, 0x7f, 0x00,
	0xf2, 0xc5, 0x95, 0xca, 0x89, 0x81, 0xad, 0x49, 0x88, 0x3c, 0x98, 0x7b, 0x62, 0x9e, 0x08, 0xae,
	0x10, 0x79, 0xf4, 0xec, 0xa7, 0x8d, 0x3c, 0xf1, 0xf6, 0x07, 0x4f, 0xf0, 0x9f, 0x77, 0xa6, 0x8e,
	0x1c, 0x90, 0xe9, 0xc9, 0xb6, 0xe5, 0x6c, 0x99, 0xad, 0x66, 0xd5, 0xa1, 0x03, 0x7a, 0x21, 0xdb,
	0x16, 0x17, 0x6a, 0xb1, 0x83, 0x37, 0x4c, 0xe6, 0xf7, 0x73, 0x4f, 0xe5, 0x5c, 0xe2, 0x63, 0x61,
	0x4f, 0x82, 0xee, 0xef, 0x5e, 0x6f, 0x5c, 0x60, 0xdf, 0xd7, 0x9b, 0x06, 0x77, 0xfa, 0x23, 0xd4,
	0x09, 0xd6, 0x14, 0x1c, 0x8d, 0x63, 0xfb, 0xbe, 0x76, 0x38, 0x2c, 0x3a, 0x75, 0x8d, 0x96, 0x0a,
	0xea, 0x03, 0x29, 0xe4, 0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x4a, 0x65, 0xcd, 0x30, 0x01,
	0x00, 0x00,
}