// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/descredit/descredit.proto

package descredit

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	DB "github.com/pkt-cash/pktd/lnd/lnrpc/protos/DB"
	_ "github.com/pkt-cash/pktd/lnd/lnrpc/protos/lock"
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
	/// Caleb: Should be `message Block` rather than string
	FirstBlockHeader *DB.Block `protobuf:"bytes,1,opt,name=first_block_header,json=firstBlockHeader,proto3" json:"first_block_header,omitempty"`
	// second_block_header — block copy
	/// Caleb: Should be `message Block` rather than string
	SecondBlockHeader    *DB.Block `protobuf:"bytes,2,opt,name=second_block_header,json=secondBlockHeader,proto3" json:"second_block_header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
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

func (m *DuplicateBlock) GetFirstBlockHeader() *DB.Block {
	if m != nil {
		return m.FirstBlockHeader
	}
	return nil
}

func (m *DuplicateBlock) GetSecondBlockHeader() *DB.Block {
	if m != nil {
		return m.SecondBlockHeader
	}
	return nil
}

// DenialService the issuer accepts a lock_tokens request from a sender but
// then ignores the transfer_tokens request from the recipient
type DenialService struct {
	// lock — hash of ignored lock
	Lock string `protobuf:"bytes,1,opt,name=lock,proto3" json:"lock,omitempty"`
	// htlc_secret — htlc genereted issuer
	HtlcSecret string `protobuf:"bytes,2,opt,name=htlc_secret,json=htlcSecret,proto3" json:"htlc_secret,omitempty"`
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

func (m *DenialService) GetLock() string {
	if m != nil {
		return m.Lock
	}
	return ""
}

func (m *DenialService) GetHtlcSecret() string {
	if m != nil {
		return m.HtlcSecret
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
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x3f, 0x6b, 0xf3, 0x30,
	0x10, 0xc6, 0xf1, 0xfb, 0x96, 0x82, 0x15, 0x1a, 0x1a, 0x05, 0xda, 0x90, 0xa5, 0x21, 0x53, 0x97,
	0xda, 0xd0, 0xd0, 0xad, 0x74, 0x30, 0x1e, 0xba, 0x74, 0x71, 0xb6, 0x2e, 0xc6, 0x3e, 0x5d, 0x6a,
	0x61, 0xc5, 0x12, 0xd2, 0xa5, 0x1f, 0xa3, 0x9f, 0xb9, 0xe8, 0x9c, 0xfe, 0xa5, 0x83, 0xcd, 0xa3,
	0xe7, 0xf7, 0xbb, 0x43, 0x20, 0xb1, 0x72, 0xde, 0x92, 0x0d, 0xb9, 0xc2, 0x00, 0x1e, 0x95, 0xa6,
	0xaf, 0x94, 0x31, 0x92, 0xe9, 0x67, 0xb1, 0xbc, 0x3c, 0xca, 0x65, 0x91, 0x93, 0xed, 0x71, 0x50,
	0xed, 0xe8, 0x2c, 0x2f, 0x8e, 0xc0, 0x58, 0xe8, 0xf9, 0x37, 0xf6, 0xeb, 0xb7, 0x44, 0x4c, 0xcb,
	0x83, 0x33, 0x1a, 0x1a, 0xc2, 0x22, 0x02, 0x79, 0x2f, 0xe4, 0x4e, 0xfb, 0x40, 0x75, 0x1b, 0x8f,
	0x75, 0x87, 0x8d, 0x42, 0xbf, 0x48, 0x56, 0xc9, 0xf5, 0xe4, 0x76, 0x9a, 0x7d, 0xac, 0x65, 0xb7,
	0x3a, 0x67, 0x93, 0xf3, 0x23, 0x7b, 0xf2, 0x41, 0xcc, 0x03, 0x82, 0x1d, 0xd4, 0xcf, 0xf1, 0x7f,
	0x7f, 0x8e, 0xcf, 0x46, 0xf5, 0xdb, 0xfc, 0x9a, 0xc4, 0x59, 0x89, 0x83, 0x6e, 0xcc, 0x16, 0xfd,
	0xab, 0x06, 0x94, 0x52, 0x9c, 0x44, 0xcc, 0x17, 0x48, 0x2b, 0xce, 0xf2, 0x4a, 0x4c, 0x3a, 0x32,
	0x50, 0x07, 0x04, 0x8f, 0xc4, 0xcb, 0xd3, 0x4a, 0xc4, 0x6a, 0xcb, 0x8d, 0xcc, 0xc4, 0xdc, 0x79,
	0x6b, 0x77, 0xf5, 0x1e, 0x3d, 0x18, 0xac, 0x5b, 0xdf, 0x0c, 0xd0, 0x2d, 0xfe, 0xb3, 0x38, 0x63,
	0xf4, 0xc4, 0xa4, 0x60, 0x50, 0xdc, 0x3d, 0x6f, 0x5e, 0x34, 0x75, 0x87, 0x36, 0x03, 0xbb, 0xcf,
	0x5d, 0x4f, 0x37, 0xd0, 0x84, 0x2e, 0x06, 0x95, 0x9b, 0x21, 0x7e, 0xde, 0x41, 0xfe, 0xfb, 0x25,
	0xda, 0x53, 0x6e, 0x36, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x5a, 0x8f, 0x01, 0xa4, 0x01,
	0x00, 0x00,
}
