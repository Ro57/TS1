// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/descredit/descredit.proto

package descredit

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	justifications "github.com/pkt-cash/pktd/lnd/lnrpc/protos/justifications"
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

// DuplicateBlock 2 identical blocks were generated (Updated the structure)
type DuplicateBlock struct {
	// first_block_header — copied block
	FirstBlockHeader string `protobuf:"bytes,1,opt,name=first_block_header,json=firstBlockHeader,proto3" json:"first_block_header,omitempty"`
	// second_block_header — block copy
	SecondBlockHeader    string   `protobuf:"bytes,2,opt,name=second_block_header,json=secondBlockHeader,proto3" json:"second_block_header,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DuplicateBlock) Reset()         { *m = DuplicateBlock{} }
func (m *DuplicateBlock) String() string { return proto.CompactTextString(m) }
func (*DuplicateBlock) ProtoMessage()    {}
func (*DuplicateBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d1aba1ea580c354, []int{0}
}

func (m *DuplicateBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DuplicateBlock.Unmarshal(m, b)
}
func (m *DuplicateBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DuplicateBlock.Marshal(b, m, deterministic)
}
func (m *DuplicateBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DuplicateBlock.Merge(m, src)
}
func (m *DuplicateBlock) XXX_Size() int {
	return xxx_messageInfo_DuplicateBlock.Size(m)
}
func (m *DuplicateBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_DuplicateBlock.DiscardUnknown(m)
}

var xxx_messageInfo_DuplicateBlock proto.InternalMessageInfo

func (m *DuplicateBlock) GetFirstBlockHeader() string {
	if m != nil {
		return m.FirstBlockHeader
	}
	return ""
}

func (m *DuplicateBlock) GetSecondBlockHeader() string {
	if m != nil {
		return m.SecondBlockHeader
	}
	return ""
}

// DenialService the issuer accepts a lock_tokens request from a sender but
// then ignores the transfer_tokens request from the recipient
type DenialService struct {
	// lock — information about ignored lock
	Lock *justifications.LockToken `protobuf:"bytes,1,opt,name=lock,proto3" json:"lock,omitempty"`
	// htlc_secret_hash — htlc genereted issuer
	HtlcSecretHash string `protobuf:"bytes,2,opt,name=htlc_secret_hash,json=htlcSecretHash,proto3" json:"htlc_secret_hash,omitempty"`
	// proof_mercle_branch — proof from PKT blockchain
	ProofMercleBranch    string   `protobuf:"bytes,3,opt,name=proof_mercle_branch,json=proofMercleBranch,proto3" json:"proof_mercle_branch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DenialService) Reset()         { *m = DenialService{} }
func (m *DenialService) String() string { return proto.CompactTextString(m) }
func (*DenialService) ProtoMessage()    {}
func (*DenialService) Descriptor() ([]byte, []int) {
	return fileDescriptor_4d1aba1ea580c354, []int{1}
}

func (m *DenialService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DenialService.Unmarshal(m, b)
}
func (m *DenialService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DenialService.Marshal(b, m, deterministic)
}
func (m *DenialService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenialService.Merge(m, src)
}
func (m *DenialService) XXX_Size() int {
	return xxx_messageInfo_DenialService.Size(m)
}
func (m *DenialService) XXX_DiscardUnknown() {
	xxx_messageInfo_DenialService.DiscardUnknown(m)
}

var xxx_messageInfo_DenialService proto.InternalMessageInfo

func (m *DenialService) GetLock() *justifications.LockToken {
	if m != nil {
		return m.Lock
	}
	return nil
}

func (m *DenialService) GetHtlcSecretHash() string {
	if m != nil {
		return m.HtlcSecretHash
	}
	return ""
}

func (m *DenialService) GetProofMercleBranch() string {
	if m != nil {
		return m.ProofMercleBranch
	}
	return ""
}

func init() {
	proto.RegisterType((*DuplicateBlock)(nil), "descredit.DuplicateBlock")
	proto.RegisterType((*DenialService)(nil), "descredit.DenialService")
}

func init() { proto.RegisterFile("protos/descredit/descredit.proto", fileDescriptor_4d1aba1ea580c354) }

var fileDescriptor_4d1aba1ea580c354 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xcd, 0x4e, 0xeb, 0x30,
	0x10, 0x85, 0x95, 0x7b, 0x11, 0x52, 0x8d, 0xa8, 0x8a, 0xd9, 0x14, 0x56, 0x55, 0x57, 0x15, 0xa2,
	0x89, 0x44, 0xc5, 0x0b, 0x44, 0x5d, 0x74, 0x01, 0x9b, 0x96, 0x15, 0x1b, 0xcb, 0x99, 0x4c, 0xb0,
	0x89, 0x6b, 0x47, 0xb6, 0xc3, 0xb3, 0xf0, 0xb8, 0xc8, 0x93, 0x8a, 0x9f, 0x2c, 0x22, 0xcd, 0x9c,
	0xef, 0x8b, 0xe6, 0x48, 0x66, 0x8b, 0xce, 0xbb, 0xe8, 0x42, 0x51, 0x63, 0x00, 0x8f, 0xb5, 0x8e,
	0x3f, 0x53, 0x4e, 0x88, 0x4f, 0xbe, 0x83, 0xdb, 0xbb, 0x93, 0xfc, 0xde, 0x87, 0xa8, 0x1b, 0x0d,
	0x32, 0x6a, 0x67, 0xc7, 0xeb, 0xf0, 0xdb, 0xd2, 0xb2, 0xe9, 0xb6, 0xef, 0x4c, 0x4a, 0xb1, 0x34,
	0x0e, 0x5a, 0x7e, 0xcf, 0x78, 0xa3, 0x7d, 0x88, 0xa2, 0x4a, 0xab, 0x50, 0x28, 0x6b, 0xf4, 0xf3,
	0x6c, 0x91, 0xad, 0x26, 0xfb, 0x19, 0x11, 0xf2, 0x76, 0x94, 0xf3, 0x9c, 0x5d, 0x07, 0x04, 0x67,
	0xeb, 0xbf, 0xfa, 0x3f, 0xd2, 0xaf, 0x06, 0xf4, 0xcb, 0x5f, 0x7e, 0x66, 0xec, 0x72, 0x8b, 0x56,
	0x4b, 0x73, 0x40, 0xff, 0xa1, 0x01, 0xf9, 0x9a, 0x9d, 0x25, 0x4e, 0x17, 0x2e, 0x1e, 0x6e, 0xf2,
	0x51, 0xcd, 0x27, 0x07, 0xed, 0x8b, 0x6b, 0xd1, 0xee, 0x49, 0xe3, 0x2b, 0x36, 0x53, 0xd1, 0x80,
	0x08, 0x08, 0x1e, 0xa3, 0x50, 0x32, 0xa8, 0xd3, 0xb5, 0x69, 0xca, 0x0f, 0x14, 0xef, 0x64, 0x50,
	0xa9, 0x5a, 0xe7, 0x9d, 0x6b, 0xc4, 0x11, 0x3d, 0x18, 0x14, 0x95, 0x97, 0x16, 0xd4, 0xfc, 0xff,
	0x50, 0x8d, 0xd0, 0x33, 0x91, 0x92, 0x40, 0xf9, 0xf8, 0xba, 0x79, 0xd3, 0x51, 0xf5, 0x55, 0x0e,
	0xee, 0x58, 0x74, 0x6d, 0x5c, 0x83, 0x0c, 0x2a, 0x0d, 0x75, 0x61, 0x6c, 0xfa, 0x7c, 0x07, 0xc5,
	0xf8, 0x21, 0xaa, 0x73, 0x4a, 0x36, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xce, 0xa6, 0x97,
	0xa3, 0x01, 0x00, 0x00,
}
