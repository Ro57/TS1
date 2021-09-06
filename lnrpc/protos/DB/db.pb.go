// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/DB/db.proto

package DB

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

// Block is struct of block in blockchain
// Caleb: We should have a "block header" which is small and can be
// copied into an Accusation
type Block struct {
	// header — data to identify this block
	// Caleb: Block header should contain the height of the block
	Header *BlockHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// signature — issuer ID, needed for validate. If signature incorrect block
	// is not valid
	Signature string `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	// available_count — number available tokens for buying
	AvailableCount int64 `protobuf:"varint,4,opt,name=available_count,json=availableCount,proto3" json:"available_count,omitempty"`
	// locks — set of lock structures
	Locks []*Lock `protobuf:"bytes,5,rep,name=locks,proto3" json:"locks,omitempty"`
	// owners — a set of structures with addresses and their balances
	Owners []*Owner `protobuf:"bytes,6,rep,name=owners,proto3" json:"owners,omitempty"`
	// justification — one of justification structure with payload information.
	//
	// Types that are valid to be assigned to Justification:
	//	*Block_Lock
	//	*Block_Transfer
	//	*Block_LockOver
	Justification        isBlock_Justification `protobuf_oneof:"justification"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
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

func (m *Block) GetSignature() string {
	if m != nil {
		return m.Signature
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

type isBlock_Justification interface {
	isBlock_Justification()
}

type Block_Lock struct {
	Lock *justifications.LockToken `protobuf:"bytes,7,opt,name=lock,proto3,oneof"`
}

type Block_Transfer struct {
	Transfer *justifications.TranferToken `protobuf:"bytes,8,opt,name=transfer,proto3,oneof"`
}

type Block_LockOver struct {
	LockOver *justifications.LockTimeOver `protobuf:"bytes,9,opt,name=lock_over,json=lockOver,proto3,oneof"`
}

func (*Block_Lock) isBlock_Justification() {}

func (*Block_Transfer) isBlock_Justification() {}

func (*Block_LockOver) isBlock_Justification() {}

func (m *Block) GetJustification() isBlock_Justification {
	if m != nil {
		return m.Justification
	}
	return nil
}

func (m *Block) GetLock() *justifications.LockToken {
	if x, ok := m.GetJustification().(*Block_Lock); ok {
		return x.Lock
	}
	return nil
}

func (m *Block) GetTransfer() *justifications.TranferToken {
	if x, ok := m.GetJustification().(*Block_Transfer); ok {
		return x.Transfer
	}
	return nil
}

func (m *Block) GetLockOver() *justifications.LockTimeOver {
	if x, ok := m.GetJustification().(*Block_LockOver); ok {
		return x.LockOver
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Block) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Block_Lock)(nil),
		(*Block_Transfer)(nil),
		(*Block_LockOver)(nil),
	}
}

// Lock contain information about tokens and contract for transferring
type Lock struct {
	// id — hash identifactor of lock. It will be send as response on
	// lock_tokens_for_transfer method, It will be used as signature generated
	// from LockToken justification.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// count — number of sending tokens
	Count int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	// owner — wallet addres of new owner of tokens
	Owner string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	// htlc — hash of contract
	Htlc string `protobuf:"bytes,4,opt,name=htlc,proto3" json:"htlc,omitempty"`
	// proof_count — lock expiration time in PKT blocks
	ProofCount int32 `protobuf:"varint,5,opt,name=proof_count,json=proofCount,proto3" json:"proof_count,omitempty"`
	// timeover_pkt — block height of PKT block when this lock will over
	TimeoverPkt int32 `protobuf:"varint,6,opt,name=timeover_pkt,json=timeoverPkt,proto3" json:"timeover_pkt,omitempty"`
	// creation_height — creation height in token blockchain
	CreationHeight       uint64   `protobuf:"varint,7,opt,name=creation_height,json=creationHeight,proto3" json:"creation_height,omitempty"`
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

func (m *Lock) GetTimeoverPkt() int32 {
	if m != nil {
		return m.TimeoverPkt
	}
	return 0
}

func (m *Lock) GetCreationHeight() uint64 {
	if m != nil {
		return m.CreationHeight
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
	IssuerAddress string `protobuf:"bytes,3,opt,name=issuer_address,json=issuerAddress,proto3" json:"issuer_address,omitempty"`
	// state — hash of current state of database
	State string `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	// pkt_block_hash —  the hash of the most recent PKT block
	PktBlockHash string `protobuf:"bytes,5,opt,name=pkt_block_hash,json=pktBlockHash,proto3" json:"pkt_block_hash,omitempty"`
	// pkt_block_height — the height of the most recent PKT block
	PktBlockHeight int32 `protobuf:"varint,6,opt,name=pkt_block_height,json=pktBlockHeight,proto3" json:"pkt_block_height,omitempty"`
	// height — the current height of this TokenStrike chain
	Height               uint64   `protobuf:"varint,7,opt,name=height,proto3" json:"height,omitempty"`
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

func (m *BlockHeader) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *BlockHeader) GetPktBlockHash() string {
	if m != nil {
		return m.PktBlockHash
	}
	return ""
}

func (m *BlockHeader) GetPktBlockHeight() int32 {
	if m != nil {
		return m.PktBlockHeight
	}
	return 0
}

func (m *BlockHeader) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
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
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x26, 0x6d, 0xd3, 0x2d, 0xa7, 0x6b, 0x36, 0x59, 0x13, 0x0a, 0xd3, 0x80, 0xd2, 0xf1, 0x53,
	0xa1, 0xd1, 0x4a, 0xe3, 0x0e, 0xae, 0x28, 0x5c, 0xec, 0x02, 0x69, 0xc8, 0x9a, 0x84, 0xc4, 0x4d,
	0xe4, 0x26, 0xee, 0x62, 0x92, 0xc5, 0x91, 0xed, 0x74, 0x3c, 0x00, 0x6f, 0xc6, 0xf3, 0xf0, 0x0e,
	0xc8, 0xc7, 0x69, 0x16, 0xaa, 0x5d, 0x54, 0xf2, 0xf9, 0xce, 0xf7, 0x1d, 0x9f, 0x73, 0xbe, 0xb8,
	0x40, 0x2a, 0x25, 0x8d, 0xd4, 0x8b, 0x2f, 0xcb, 0x45, 0xba, 0x9a, 0x63, 0x40, 0xf6, 0x8c, 0xcc,
	0x79, 0x99, 0xae, 0x4e, 0xde, 0x36, 0xc9, 0x9f, 0xb5, 0x36, 0x62, 0x2d, 0x12, 0x66, 0x84, 0x2c,
	0x77, 0x43, 0x27, 0x9a, 0xfe, 0xee, 0x83, 0xbf, 0x2c, 0x64, 0x92, 0x93, 0x73, 0x18, 0x66, 0x9c,
	0xa5, 0x5c, 0x45, 0xde, 0xc4, 0x9b, 0x8d, 0x2e, 0x8e, 0xe7, 0x4d, 0xbd, 0x39, 0xe6, 0x2f, 0x31,
	0x47, 0x1b, 0x0e, 0x39, 0x85, 0x40, 0x8b, 0x9b, 0x92, 0x99, 0x5a, 0xf1, 0xa8, 0x37, 0xf1, 0x66,
	0x01, 0xbd, 0x07, 0xc8, 0x1b, 0x38, 0x64, 0x1b, 0x26, 0x0a, 0xb6, 0x2a, 0x78, 0x9c, 0xc8, 0xba,
	0x34, 0xd1, 0x60, 0xe2, 0xcd, 0xfa, 0x34, 0x6c, 0xe1, 0xcf, 0x16, 0x25, 0x67, 0xe0, 0xdb, 0xe2,
	0x3a, 0xf2, 0x27, 0xfd, 0xd9, 0xe8, 0x62, 0xdc, 0xde, 0xf9, 0x55, 0x26, 0x39, 0x75, 0x39, 0xf2,
	0x1a, 0x86, 0xf2, 0xae, 0xe4, 0x4a, 0x47, 0x43, 0x64, 0x85, 0x2d, 0xeb, 0xca, 0xc2, 0xb4, 0xc9,
	0x92, 0x05, 0x0c, 0xac, 0x20, 0xda, 0xc3, 0xfe, 0x9f, 0xcc, 0x77, 0x06, 0xb6, 0x25, 0xaf, 0xad,
	0xf0, 0xf2, 0x11, 0x45, 0x22, 0xf9, 0x00, 0xfb, 0x46, 0xb1, 0x52, 0xaf, 0xb9, 0x8a, 0xf6, 0x51,
	0x74, 0xba, 0x2b, 0xba, 0x56, 0xac, 0x5c, 0x73, 0xb5, 0xd5, 0xb5, 0x7c, 0xf2, 0x11, 0x02, 0x5b,
	0x23, 0x96, 0x1b, 0xae, 0xa2, 0xe0, 0x61, 0x31, 0xde, 0x28, 0x6e, 0xf9, 0xd5, 0x86, 0x2b, 0x2b,
	0xb6, 0x02, 0x7b, 0x5e, 0x1e, 0xc2, 0xf8, 0x3f, 0xea, 0xf4, 0x8f, 0x07, 0x03, 0xcb, 0x26, 0x21,
	0xf4, 0x44, 0x8a, 0x0e, 0x04, 0xb4, 0x27, 0x52, 0x72, 0x0c, 0xbe, 0xdb, 0x5f, 0x0f, 0xf7, 0xe7,
	0x02, 0x8b, 0xe2, 0xcc, 0x51, 0x1f, 0x89, 0x2e, 0x20, 0x04, 0x06, 0x99, 0x29, 0x12, 0x5c, 0x75,
	0x40, 0xf1, 0x4c, 0x9e, 0xc3, 0xa8, 0x52, 0x52, 0xae, 0x1b, 0x17, 0xfc, 0x89, 0x37, 0xf3, 0x29,
	0x20, 0xe4, 0x1c, 0x78, 0x01, 0x07, 0x46, 0xdc, 0x72, 0x3b, 0x46, 0x5c, 0xe5, 0x26, 0x1a, 0x22,
	0x63, 0xb4, 0xc5, 0xbe, 0xe5, 0xc6, 0xba, 0x99, 0x28, 0x8e, 0x8d, 0xc6, 0x19, 0x17, 0x37, 0x99,
	0xc1, 0x15, 0x0f, 0x68, 0xb8, 0x85, 0x2f, 0x11, 0x9d, 0x2e, 0xc1, 0x47, 0x47, 0xc8, 0x19, 0x8c,
	0x33, 0x59, 0xa4, 0x5c, 0xc5, 0x77, 0xac, 0x28, 0xb8, 0x69, 0x06, 0x3a, 0x70, 0xe0, 0x77, 0xc4,
	0x1e, 0x1e, 0x6d, 0x9a, 0x83, 0x8f, 0xcb, 0xbe, 0x4f, 0x7b, 0xdd, 0xc9, 0x9f, 0x01, 0xf0, 0x5f,
	0x95, 0x50, 0x78, 0x2d, 0x2a, 0x7d, 0xda, 0x41, 0xc8, 0x09, 0xec, 0x6f, 0x9b, 0xc2, 0xe5, 0xf4,
	0x69, 0x1b, 0x93, 0x23, 0xe8, 0xd7, 0xaa, 0x68, 0xd6, 0x63, 0x8f, 0xd3, 0xbf, 0x1e, 0x8c, 0x3a,
	0x5f, 0x37, 0x79, 0x0a, 0x50, 0x29, 0xbe, 0x89, 0x57, 0xf8, 0x1d, 0xb9, 0xa6, 0x03, 0x8b, 0xb8,
	0x27, 0xd2, 0x2d, 0xde, 0xdb, 0x29, 0xfe, 0x0a, 0x42, 0xa1, 0x75, 0xcd, 0x55, 0xcc, 0xd2, 0x54,
	0x71, 0xad, 0x1b, 0x6f, 0xc6, 0x0e, 0xfd, 0xe4, 0x40, 0x3b, 0x95, 0x36, 0xcc, 0xf0, 0xa6, 0x0b,
	0x17, 0x90, 0x97, 0x10, 0x56, 0xb9, 0x71, 0xd7, 0xc6, 0x19, 0xd3, 0x19, 0x1a, 0x15, 0xd0, 0x83,
	0x2a, 0x37, 0xae, 0x3f, 0xa6, 0x33, 0x32, 0x83, 0xa3, 0x0e, 0xcb, 0x19, 0xe1, 0xec, 0x0a, 0x5b,
	0x1e, 0xa2, 0xe4, 0xb1, 0x7d, 0xcb, 0x1d, 0xa3, 0x9a, 0x68, 0x39, 0xff, 0x71, 0x7e, 0x23, 0x4c,
	0x56, 0xaf, 0xe6, 0x89, 0xbc, 0x5d, 0x54, 0xb9, 0x79, 0x97, 0x30, 0x9d, 0xd9, 0x43, 0xba, 0x28,
	0x4a, 0xfb, 0x53, 0x55, 0xb2, 0x68, 0xff, 0x5b, 0x56, 0x43, 0x3c, 0xbe, 0xff, 0x17, 0x00, 0x00,
	0xff, 0xff, 0x96, 0x28, 0xae, 0xc4, 0x6f, 0x04, 0x00, 0x00,
}
