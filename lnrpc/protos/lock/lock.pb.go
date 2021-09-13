// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/lock/lock.proto

package lock

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

// Lock contain information about tokens and contract for transferring
type Lock struct {
	// count — number of sending tokens
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// recipient — wallet addres of new owner of tokens
	Recipient string `protobuf:"bytes,2,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// sender — owner of the wallet address to which tokens will be returned
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	// htlc_secret_hash — hash of contract
	HtlcSecretHash string `protobuf:"bytes,4,opt,name=htlc_secret_hash,json=htlcSecretHash,proto3" json:"htlc_secret_hash,omitempty"`
	// proof_count — lock expiration time in PKT blocks
	ProofCount int32 `protobuf:"varint,5,opt,name=proof_count,json=proofCount,proto3" json:"proof_count,omitempty"`
	// creation_height — creation height in token blockchain
	CreationHeight uint64 `protobuf:"varint,6,opt,name=creation_height,json=creationHeight,proto3" json:"creation_height,omitempty"`
	// signature generated with old owner private key
	Signature            string   `protobuf:"bytes,7,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Lock) Reset()         { *m = Lock{} }
func (m *Lock) String() string { return proto.CompactTextString(m) }
func (*Lock) ProtoMessage()    {}
func (*Lock) Descriptor() ([]byte, []int) {
	return fileDescriptor_d340bab4d79d59c9, []int{0}
}

func (m *Lock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Lock.Unmarshal(m, b)
}
func (m *Lock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Lock.Marshal(b, m, deterministic)
}
func (m *Lock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lock.Merge(m, src)
}
func (m *Lock) XXX_Size() int {
	return xxx_messageInfo_Lock.Size(m)
}
func (m *Lock) XXX_DiscardUnknown() {
	xxx_messageInfo_Lock.DiscardUnknown(m)
}

var xxx_messageInfo_Lock proto.InternalMessageInfo

func (m *Lock) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Lock) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *Lock) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *Lock) GetHtlcSecretHash() string {
	if m != nil {
		return m.HtlcSecretHash
	}
	return ""
}

func (m *Lock) GetProofCount() int32 {
	if m != nil {
		return m.ProofCount
	}
	return 0
}

func (m *Lock) GetCreationHeight() uint64 {
	if m != nil {
		return m.CreationHeight
	}
	return 0
}

func (m *Lock) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func init() {
	proto.RegisterType((*Lock)(nil), "lock.Lock")
}

func init() { proto.RegisterFile("protos/lock/lock.proto", fileDescriptor_d340bab4d79d59c9) }

var fileDescriptor_d340bab4d79d59c9 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xd0, 0xcd, 0x4e, 0xc4, 0x20,
	0x14, 0x05, 0xe0, 0xe0, 0xb4, 0x35, 0x83, 0xc9, 0x68, 0x88, 0x99, 0xb0, 0x30, 0xb1, 0x71, 0x23,
	0x1b, 0x5b, 0x13, 0xdf, 0x40, 0x37, 0xb3, 0x70, 0x55, 0x77, 0x6e, 0x1a, 0xe6, 0x16, 0x0b, 0x69,
	0x05, 0x02, 0xb7, 0x6f, 0xec, 0x83, 0x18, 0xa8, 0x3f, 0xb3, 0x80, 0xdc, 0xf3, 0x41, 0xc2, 0x09,
	0x74, 0xef, 0x83, 0x43, 0x17, 0xdb, 0xd9, 0xc1, 0x94, 0xb7, 0x26, 0x03, 0x2b, 0xd2, 0x7c, 0xf7,
	0x45, 0x68, 0xf1, 0xea, 0x60, 0x62, 0xd7, 0xb4, 0x04, 0xb7, 0x58, 0xe4, 0xa4, 0x26, 0x62, 0xd3,
	0xad, 0x81, 0xdd, 0xd0, 0x6d, 0x50, 0x60, 0xbc, 0x51, 0x16, 0xf9, 0x59, 0x4d, 0xc4, 0xb6, 0xfb,
	0x07, 0xb6, 0xa7, 0x55, 0x54, 0x76, 0x50, 0x81, 0x6f, 0xf2, 0xd1, 0x4f, 0x62, 0x82, 0x5e, 0x69,
	0x9c, 0xa1, 0x8f, 0x0a, 0x82, 0xc2, 0x5e, 0xcb, 0xa8, 0x79, 0x91, 0x6f, 0xec, 0x92, 0xbf, 0x65,
	0x3e, 0xc8, 0xa8, 0xd9, 0x2d, 0xbd, 0xf0, 0xc1, 0xb9, 0x8f, 0x7e, 0x7d, 0xbb, 0xac, 0x89, 0x28,
	0x3b, 0x9a, 0xe9, 0x25, 0x17, 0xb8, 0xa7, 0x97, 0x10, 0x94, 0x44, 0xe3, 0x6c, 0xaf, 0x95, 0x19,
	0x35, 0xf2, 0xaa, 0x26, 0xa2, 0xe8, 0x76, 0xbf, 0x7c, 0xc8, 0x9a, 0x9a, 0x46, 0x33, 0x5a, 0x89,
	0x4b, 0x50, 0xfc, 0x7c, 0x6d, 0xfa, 0x07, 0xcf, 0x8f, 0xef, 0xcd, 0x68, 0x50, 0x2f, 0xc7, 0x06,
	0xdc, 0x67, 0xeb, 0x27, 0x7c, 0x00, 0x19, 0x75, 0x1a, 0x86, 0x76, 0xb6, 0x69, 0x05, 0x0f, 0xed,
	0xc9, 0x4f, 0x1d, 0xab, 0x1c, 0x9e, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x04, 0x59, 0xec, 0x94,
	0x3f, 0x01, 0x00, 0x00,
}
