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
// Caleb: We should have a "block header" which is small and can be
// copied into an Accusation
type Block struct {
	// header — data to identify this block
	// Caleb: Block header should contain the height of the block
	Header *BlockHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// justification — one of justification structure with payload information.
	// TODO: After implement type of justification, change Any type to oneof
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

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
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
	Url                  string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
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

func (m *Token) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

// BlockHeader contain reference on previus block and creation time.
// Hash of this structure will be used as a key in key-value storage
type BlockHeader struct {
	// prev_block — hash of previous block
	PrevBlock string `protobuf:"bytes,1,opt,name=prev_block,json=prevBlock,proto3" json:"prev_block,omitempty"`
	// creation — date of block creation in unix time format
	Creation int64 `protobuf:"varint,2,opt,name=creation,proto3" json:"creation,omitempty"`
	// issuer_address — issuer wallet address to transaction.
	// Update information about token will be by URL.
	IssuerAddress        string   `protobuf:"bytes,3,opt,name=issuer_address,json=issuerAddress,proto3" json:"issuer_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_2899384ae3d1f0cf, []int{4}
}

func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetPrevBlock() string {
	if m != nil {
		return m.PrevBlock
	}
	return ""
}

func (m *BlockHeader) GetCreation() int64 {
	if m != nil {
		return m.Creation
	}
	return 0
}

func (m *BlockHeader) GetIssuerAddress() string {
	if m != nil {
		return m.IssuerAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Block)(nil), "tokendb.Block")
	proto.RegisterType((*Lock)(nil), "tokendb.Lock")
	proto.RegisterType((*Owner)(nil), "tokendb.Owner")
	proto.RegisterType((*Token)(nil), "tokendb.Token")
	proto.RegisterType((*BlockHeader)(nil), "tokendb.BlockHeader")
}

func init() { proto.RegisterFile("protos/DB/db.proto", fileDescriptor_2899384ae3d1f0cf) }

var fileDescriptor_2899384ae3d1f0cf = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x4f, 0x8f, 0xd3, 0x30,
	0x10, 0xc5, 0x95, 0x64, 0xd3, 0xa5, 0x53, 0x5a, 0x90, 0xb5, 0x87, 0xb0, 0xe2, 0x4f, 0x95, 0x15,
	0xd0, 0xc3, 0x92, 0x48, 0xcb, 0x8d, 0xdb, 0x16, 0x0e, 0x1c, 0x90, 0x90, 0x22, 0x24, 0x24, 0x2e,
	0x95, 0x63, 0xbb, 0x8d, 0x89, 0x89, 0x23, 0xdb, 0xd9, 0xb2, 0x9f, 0x85, 0x2f, 0x8b, 0xfc, 0x87,
	0x34, 0x20, 0x0e, 0x95, 0x66, 0x7e, 0x6f, 0x26, 0x79, 0xcf, 0x75, 0x00, 0xf5, 0x4a, 0x1a, 0xa9,
	0xcb, 0x0f, 0xdb, 0x92, 0xd6, 0x85, 0x6b, 0xd0, 0xb9, 0x91, 0x2d, 0xeb, 0x68, 0x7d, 0xf9, 0xe4,
	0x20, 0xe5, 0x41, 0xb0, 0xd2, 0xe1, 0x7a, 0xd8, 0x97, 0xb8, 0xbb, 0xf7, 0x33, 0xf9, 0xaf, 0x18,
	0xd2, 0xad, 0x90, 0xa4, 0x45, 0xd7, 0x30, 0x6b, 0x18, 0xa6, 0x4c, 0x65, 0xd1, 0x3a, 0xda, 0x2c,
	0x6e, 0x2e, 0x8a, 0xb0, 0x5e, 0x38, 0xfd, 0xa3, 0xd3, 0xaa, 0x30, 0x83, 0xde, 0xc1, 0xf2, 0xfb,
	0xa0, 0x0d, 0xdf, 0x73, 0x82, 0x0d, 0x97, 0x5d, 0x16, 0x87, 0x25, 0xff, 0xaa, 0xe2, 0xcf, 0xab,
	0x8a, 0xdb, 0xee, 0xbe, 0xfa, 0x7b, 0x14, 0x3d, 0x85, 0xb9, 0xe6, 0x87, 0x0e, 0x9b, 0x41, 0xb1,
	0x2c, 0x59, 0x47, 0x9b, 0x79, 0x75, 0x02, 0xe8, 0x02, 0x52, 0x6d, 0xb0, 0x61, 0xd9, 0x99, 0x53,
	0x7c, 0x83, 0x5e, 0xc3, 0x23, 0x7c, 0x87, 0xb9, 0xc0, 0xb5, 0x60, 0x3b, 0x22, 0x87, 0xce, 0x64,
	0xe9, 0x3a, 0xda, 0x24, 0xd5, 0x6a, 0xc4, 0xef, 0x2d, 0x45, 0x57, 0x90, 0x5a, 0xbb, 0x3a, 0x9b,
	0xad, 0x93, 0xcd, 0xe2, 0x66, 0x39, 0xa6, 0xf8, 0x24, 0x49, 0x5b, 0x79, 0x0d, 0xbd, 0x82, 0x99,
	0x3c, 0x76, 0x4c, 0xe9, 0xec, 0xdc, 0x4d, 0xad, 0xc6, 0xa9, 0xcf, 0x16, 0x57, 0x41, 0xcd, 0x8f,
	0x70, 0x66, 0xd7, 0xd0, 0x0a, 0x62, 0x4e, 0xdd, 0xb9, 0xcc, 0xab, 0x98, 0x53, 0xeb, 0xd1, 0x7b,
	0x88, 0x9d, 0x07, 0xdf, 0x58, 0xea, 0xf6, 0x42, 0x26, 0xdf, 0x20, 0x04, 0x67, 0x8d, 0x11, 0x24,
	0xc4, 0x71, 0x35, 0x7a, 0x01, 0x8b, 0x5e, 0x49, 0xb9, 0x9f, 0x24, 0x49, 0x2b, 0x70, 0xc8, 0xa5,
	0xc8, 0xb7, 0x90, 0x3a, 0x27, 0xe8, 0x0a, 0x96, 0x8d, 0x14, 0x94, 0xa9, 0xdd, 0x11, 0x0b, 0xc1,
	0x4c, 0x30, 0xf1, 0xd0, 0xc3, 0xaf, 0x8e, 0xfd, 0xdf, 0x4e, 0xde, 0x42, 0xfa, 0xc5, 0xa6, 0x3a,
	0xc9, 0xd1, 0xd4, 0xed, 0x73, 0x00, 0xf6, 0xb3, 0xe7, 0xea, 0xf4, 0xf7, 0xa5, 0xd5, 0x84, 0xa0,
	0x4b, 0x78, 0x40, 0x14, 0xf3, 0x6a, 0xe2, 0x16, 0xc7, 0x1e, 0x3d, 0x86, 0x64, 0x50, 0x22, 0x44,
	0xb2, 0x65, 0x2e, 0x61, 0x31, 0xb9, 0x26, 0xe8, 0x19, 0x40, 0xaf, 0xd8, 0xdd, 0xae, 0xb6, 0x2c,
	0x78, 0x9e, 0x5b, 0xe2, 0xef, 0xda, 0xf4, 0xd9, 0xf1, 0x3f, 0xcf, 0x7e, 0x09, 0x2b, 0xae, 0xf5,
	0xc0, 0xd4, 0x0e, 0x53, 0xaa, 0x98, 0xd6, 0xe1, 0x38, 0x97, 0x9e, 0xde, 0x7a, 0xb8, 0x2d, 0xbe,
	0x5d, 0x1f, 0xb8, 0x69, 0x86, 0xba, 0x20, 0xf2, 0x47, 0xd9, 0xb7, 0xe6, 0x0d, 0xc1, 0xba, 0xb1,
	0x05, 0x2d, 0x45, 0x67, 0x7f, 0xaa, 0x27, 0xe5, 0xf8, 0x55, 0xd4, 0x33, 0x57, 0xbe, 0xfd, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x16, 0x3d, 0xc5, 0xf9, 0x29, 0x03, 0x00, 0x00,
}
