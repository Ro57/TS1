// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/DB/db.proto

package DB

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// Block is struct of block in blockchain
type Block struct {
	// prev_block — hash of previous block
	PrevBlock string `protobuf:"bytes,1,opt,name=prev_block,json=prevBlock,proto3" json:"prev_block,omitempty"`
	// justification — one of justification structure with payload information.
	// TODO: After implement types of justifications, change from
	// google.protobuf.Any to type oneof
	Justification *any.Any `protobuf:"bytes,2,opt,name=justification,proto3" json:"justification,omitempty"`
	// signature — issuer ID, needed for validate. If signature incorrect block
	// is not valid
	Signature string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	// state — hash of current state of database
	State string `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	// available_count — number available tokens for buying
	AvailableCount int64 `protobuf:"varint,5,opt,name=available_count,json=availableCount,proto3" json:"available_count,omitempty"`
	// locks — set of lock structures
	Locks []*Lock `protobuf:"bytes,6,rep,name=locks,proto3" json:"locks,omitempty"`
	// owners — a set of structures with addresses and their balances
	Owners               []*Owner `protobuf:"bytes,7,rep,name=owners,proto3" json:"owners,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_2899384ae3d1f0cf, []int{0}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetPrevBlock() string {
	if m != nil {
		return m.PrevBlock
	}
	return ""
}

func (m *Block) GetJustification() *any.Any {
	if m != nil {
		return m.Justification
	}
	return nil
}

func (m *Block) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *Block) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Block) GetAvailableCount() int64 {
	if m != nil {
		return m.AvailableCount
	}
	return 0
}

func (m *Block) GetLocks() []*Lock {
	if m != nil {
		return m.Locks
	}
	return nil
}

func (m *Block) GetOwners() []*Owner {
	if m != nil {
		return m.Owners
	}
	return nil
}

// Lock contain information about tokens and contract for transferring
type Lock struct {
	// id — identifactor of lock. It will be send as response on
	// lock_tokens_for_transfer method
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// count — number of sending tokens
	Count int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	// owner — wallet addres of new owner of tokens
	Owner string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	// htlc — hash of contract
	Htlc string `protobuf:"bytes,4,opt,name=htlc,proto3" json:"htlc,omitempty"`
	// proof_count — lock expiration time in PKT blocks
	ProofCount           int32    `protobuf:"varint,5,opt,name=proof_count,json=proofCount,proto3" json:"proof_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Lock) Reset()         { *m = Lock{} }
func (m *Lock) String() string { return proto.CompactTextString(m) }
func (*Lock) ProtoMessage()    {}
func (*Lock) Descriptor() ([]byte, []int) {
	return fileDescriptor_2899384ae3d1f0cf, []int{1}
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

func (m *Lock) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Lock) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Lock) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Lock) GetHtlc() string {
	if m != nil {
		return m.Htlc
	}
	return ""
}

func (m *Lock) GetProofCount() int32 {
	if m != nil {
		return m.ProofCount
	}
	return 0
}

// Owner contains information about the holders' wallets and their balances
type Owner struct {
	// holder_wallet — hash of wallet address of holder
	HolderWallet string `protobuf:"bytes,1,opt,name=holder_wallet,json=holderWallet,proto3" json:"holder_wallet,omitempty"`
	// count — number of tokens held on wallet
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Owner) Reset()         { *m = Owner{} }
func (m *Owner) String() string { return proto.CompactTextString(m) }
func (*Owner) ProtoMessage()    {}
func (*Owner) Descriptor() ([]byte, []int) {
	return fileDescriptor_2899384ae3d1f0cf, []int{2}
}

func (m *Owner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Owner.Unmarshal(m, b)
}
func (m *Owner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Owner.Marshal(b, m, deterministic)
}
func (m *Owner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Owner.Merge(m, src)
}
func (m *Owner) XXX_Size() int {
	return xxx_messageInfo_Owner.Size(m)
}
func (m *Owner) XXX_DiscardUnknown() {
	xxx_messageInfo_Owner.DiscardUnknown(m)
}

var xxx_messageInfo_Owner proto.InternalMessageInfo

func (m *Owner) GetHolderWallet() string {
	if m != nil {
		return m.HolderWallet
	}
	return ""
}

func (m *Owner) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

// Token contain information about token
type Token struct {
	// count — number of issued tokens;
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// expiration — number of PKT block after which the token expires
	Expiration int32 `protobuf:"varint,2,opt,name=expiration,proto3" json:"expiration,omitempty"`
	// creation — date of token creation in unix time format
	Creation int64 `protobuf:"varint,3,opt,name=creation,proto3" json:"creation,omitempty"`
	// urls — set of urls for access to blockchain
	Urls                 []string `protobuf:"bytes,4,rep,name=urls,proto3" json:"urls,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_2899384ae3d1f0cf, []int{3}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Token) GetExpiration() int32 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

func (m *Token) GetCreation() int64 {
	if m != nil {
		return m.Creation
	}
	return 0
}

func (m *Token) GetUrls() []string {
	if m != nil {
		return m.Urls
	}
	return nil
}

func init() {
	proto.RegisterType((*Block)(nil), "db.Block")
	proto.RegisterType((*Lock)(nil), "db.Lock")
	proto.RegisterType((*Owner)(nil), "db.Owner")
	proto.RegisterType((*Token)(nil), "db.Token")
}

func init() { proto.RegisterFile("protos/DB/db.proto", fileDescriptor_2899384ae3d1f0cf) }

var fileDescriptor_2899384ae3d1f0cf = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xc1, 0x8b, 0x13, 0x31,
	0x14, 0xc6, 0x99, 0x99, 0xa6, 0x6e, 0x5f, 0xdd, 0x15, 0xc2, 0x1e, 0xe2, 0xa2, 0xeb, 0x58, 0x0f,
	0xf6, 0xa0, 0x19, 0x58, 0x6f, 0xde, 0xac, 0x1e, 0x05, 0x61, 0x10, 0x04, 0x2f, 0x25, 0xc9, 0xa4,
	0x6d, 0x6c, 0x36, 0x19, 0x92, 0xcc, 0xd6, 0xbd, 0xfa, 0x97, 0x4b, 0x92, 0xb1, 0x5b, 0x61, 0x0f,
	0x85, 0xf7, 0xfd, 0xbe, 0xf7, 0x5e, 0xbe, 0x47, 0x07, 0x70, 0xef, 0x6c, 0xb0, 0xbe, 0xf9, 0xb2,
	0x6a, 0x3a, 0x4e, 0x93, 0xc0, 0x65, 0xc7, 0xaf, 0x9e, 0x6f, 0xad, 0xdd, 0x6a, 0xd9, 0x24, 0xc2,
	0x87, 0x4d, 0xc3, 0xcc, 0x7d, 0xb6, 0x17, 0x7f, 0x4a, 0x40, 0x2b, 0x6d, 0xc5, 0x1e, 0xbf, 0x04,
	0xe8, 0x9d, 0xbc, 0x5b, 0xf3, 0xa8, 0x48, 0x51, 0x17, 0xcb, 0x59, 0x3b, 0x8b, 0x24, 0xdb, 0x1f,
	0xe1, 0xfc, 0xd7, 0xe0, 0x83, 0xda, 0x28, 0xc1, 0x82, 0xb2, 0x86, 0x94, 0x75, 0xb1, 0x9c, 0xdf,
	0x5c, 0xd2, 0xbc, 0x9b, 0xfe, 0xdb, 0x4d, 0x3f, 0x99, 0xfb, 0xf6, 0xff, 0x56, 0xfc, 0x02, 0x66,
	0x5e, 0x6d, 0x0d, 0x0b, 0x83, 0x93, 0xa4, 0xca, 0x9b, 0x8f, 0x00, 0x5f, 0x02, 0xf2, 0x81, 0x05,
	0x49, 0x26, 0xc9, 0xc9, 0x02, 0xbf, 0x85, 0x67, 0xec, 0x8e, 0x29, 0xcd, 0xb8, 0x96, 0x6b, 0x61,
	0x07, 0x13, 0x08, 0xaa, 0x8b, 0x65, 0xd5, 0x5e, 0x1c, 0xf1, 0xe7, 0x48, 0xf1, 0x35, 0xa0, 0x18,
	0xd0, 0x93, 0x69, 0x5d, 0x2d, 0xe7, 0x37, 0x67, 0xb4, 0xe3, 0xf4, 0xab, 0x15, 0xfb, 0x36, 0x63,
	0xfc, 0x1a, 0xa6, 0xf6, 0x60, 0xa4, 0xf3, 0xe4, 0x49, 0x6a, 0x98, 0xc5, 0x86, 0x6f, 0x91, 0xb4,
	0xa3, 0xb1, 0x38, 0xc0, 0x24, 0x4e, 0xe0, 0x0b, 0x28, 0x55, 0x37, 0x9e, 0x5e, 0xaa, 0x2e, 0x26,
	0xcb, 0x2f, 0x97, 0xe9, 0xe5, 0x2c, 0x22, 0x4d, 0x73, 0xe3, 0x25, 0x59, 0x60, 0x0c, 0x93, 0x5d,
	0xd0, 0x62, 0x3c, 0x22, 0xd5, 0xf8, 0x15, 0xcc, 0x7b, 0x67, 0xed, 0xe6, 0x24, 0x3f, 0x6a, 0x21,
	0xa1, 0x94, 0x7d, 0xb1, 0x02, 0x94, 0x92, 0xe0, 0x37, 0x70, 0xbe, 0xb3, 0xba, 0x93, 0x6e, 0x7d,
	0x60, 0x5a, 0xcb, 0x30, 0x86, 0x78, 0x9a, 0xe1, 0x8f, 0xc4, 0x1e, 0x8f, 0xb3, 0xb8, 0x05, 0xf4,
	0xdd, 0xee, 0xa5, 0x79, 0xb0, 0x8b, 0xd3, 0xb4, 0xd7, 0x00, 0xf2, 0x77, 0xaf, 0xdc, 0xc3, 0x9f,
	0x86, 0xda, 0x13, 0x82, 0xaf, 0xe0, 0x4c, 0x38, 0x99, 0xdd, 0x2a, 0x0d, 0x1e, 0x75, 0xbc, 0x69,
	0x70, 0xda, 0x93, 0x49, 0x5d, 0xc5, 0x9b, 0x62, 0xbd, 0xa2, 0x3f, 0xdf, 0x6d, 0x55, 0xd8, 0x0d,
	0x9c, 0x0a, 0x7b, 0xdb, 0xf4, 0xfb, 0xf0, 0x5e, 0x30, 0xbf, 0x8b, 0x45, 0xd7, 0x68, 0x13, 0x7f,
	0xae, 0x17, 0xcd, 0xf1, 0x43, 0xe4, 0xd3, 0x54, 0x7e, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xcb,
	0x09, 0xaf, 0x82, 0x9c, 0x02, 0x00, 0x00,
}
