// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/justifications/justifications.proto

package justifications

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	lock "github.com/pkt-cash/pktd/lnd/lnrpc/protos/lock"
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

// LockToken the token locking
type LockToken struct {
	// lock — information about lock
	Lock                 *lock.Lock `protobuf:"bytes,1,opt,name=lock,proto3" json:"lock,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LockToken) Reset()         { *m = LockToken{} }
func (m *LockToken) String() string { return proto.CompactTextString(m) }
func (*LockToken) ProtoMessage()    {}
func (*LockToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee571a5de13873a7, []int{0}
}

func (m *LockToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockToken.Unmarshal(m, b)
}
func (m *LockToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockToken.Marshal(b, m, deterministic)
}
func (m *LockToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockToken.Merge(m, src)
}
func (m *LockToken) XXX_Size() int {
	return xxx_messageInfo_LockToken.Size(m)
}
func (m *LockToken) XXX_DiscardUnknown() {
	xxx_messageInfo_LockToken.DiscardUnknown(m)
}

var xxx_messageInfo_LockToken proto.InternalMessageInfo

func (m *LockToken) GetLock() *lock.Lock {
	if m != nil {
		return m.Lock
	}
	return nil
}

// TranferToken receiving funds for tokens and unlcok them
type TranferToken struct {
	// htlc_secret — htlc genereted issuer
	HtlcSecret string `protobuf:"bytes,1,opt,name=htlc_secret,json=htlcSecret,proto3" json:"htlc_secret,omitempty"`
	// lock — hash information about lock
	Lock                 string   `protobuf:"bytes,2,opt,name=lock,proto3" json:"lock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TranferToken) Reset()         { *m = TranferToken{} }
func (m *TranferToken) String() string { return proto.CompactTextString(m) }
func (*TranferToken) ProtoMessage()    {}
func (*TranferToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee571a5de13873a7, []int{1}
}

func (m *TranferToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranferToken.Unmarshal(m, b)
}
func (m *TranferToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranferToken.Marshal(b, m, deterministic)
}
func (m *TranferToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranferToken.Merge(m, src)
}
func (m *TranferToken) XXX_Size() int {
	return xxx_messageInfo_TranferToken.Size(m)
}
func (m *TranferToken) XXX_DiscardUnknown() {
	xxx_messageInfo_TranferToken.DiscardUnknown(m)
}

var xxx_messageInfo_TranferToken proto.InternalMessageInfo

func (m *TranferToken) GetHtlcSecret() string {
	if m != nil {
		return m.HtlcSecret
	}
	return ""
}

func (m *TranferToken) GetLock() string {
	if m != nil {
		return m.Lock
	}
	return ""
}

// LockTimeOver timeout for token locking
type LockTimeOver struct {
	// proof_elapsed — PKT block hash confirming expiration lock
	ProofElapsed string `protobuf:"bytes,1,opt,name=proof_elapsed,json=proofElapsed,proto3" json:"proof_elapsed,omitempty"`
	// lock_id — hash with information about lock justification
	Lock                 string   `protobuf:"bytes,2,opt,name=lock,proto3" json:"lock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LockTimeOver) Reset()         { *m = LockTimeOver{} }
func (m *LockTimeOver) String() string { return proto.CompactTextString(m) }
func (*LockTimeOver) ProtoMessage()    {}
func (*LockTimeOver) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee571a5de13873a7, []int{2}
}

func (m *LockTimeOver) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockTimeOver.Unmarshal(m, b)
}
func (m *LockTimeOver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockTimeOver.Marshal(b, m, deterministic)
}
func (m *LockTimeOver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockTimeOver.Merge(m, src)
}
func (m *LockTimeOver) XXX_Size() int {
	return xxx_messageInfo_LockTimeOver.Size(m)
}
func (m *LockTimeOver) XXX_DiscardUnknown() {
	xxx_messageInfo_LockTimeOver.DiscardUnknown(m)
}

var xxx_messageInfo_LockTimeOver proto.InternalMessageInfo

func (m *LockTimeOver) GetProofElapsed() string {
	if m != nil {
		return m.ProofElapsed
	}
	return ""
}

func (m *LockTimeOver) GetLock() string {
	if m != nil {
		return m.Lock
	}
	return ""
}

// Genesis initial block justification
type Genesis struct {
	// token — token identification by name
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Genesis) Reset()         { *m = Genesis{} }
func (m *Genesis) String() string { return proto.CompactTextString(m) }
func (*Genesis) ProtoMessage()    {}
func (*Genesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_ee571a5de13873a7, []int{3}
}

func (m *Genesis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Genesis.Unmarshal(m, b)
}
func (m *Genesis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Genesis.Marshal(b, m, deterministic)
}
func (m *Genesis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Genesis.Merge(m, src)
}
func (m *Genesis) XXX_Size() int {
	return xxx_messageInfo_Genesis.Size(m)
}
func (m *Genesis) XXX_DiscardUnknown() {
	xxx_messageInfo_Genesis.DiscardUnknown(m)
}

var xxx_messageInfo_Genesis proto.InternalMessageInfo

func (m *Genesis) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*LockToken)(nil), "justifications.LockToken")
	proto.RegisterType((*TranferToken)(nil), "justifications.TranferToken")
	proto.RegisterType((*LockTimeOver)(nil), "justifications.LockTimeOver")
	proto.RegisterType((*Genesis)(nil), "justifications.Genesis")
}

func init() {
	proto.RegisterFile("protos/justifications/justifications.proto", fileDescriptor_ee571a5de13873a7)
}

var fileDescriptor_ee571a5de13873a7 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4b, 0x4b, 0x03, 0x31,
	0x10, 0xc7, 0xa9, 0xf8, 0xa0, 0xd3, 0xd5, 0x43, 0x10, 0x11, 0x0f, 0x56, 0xea, 0x45, 0x14, 0x77,
	0x41, 0x2f, 0xe2, 0x51, 0x91, 0x5e, 0x04, 0xa1, 0xf6, 0xe4, 0xa5, 0x6c, 0xd3, 0x59, 0x37, 0xee,
	0x36, 0x13, 0x92, 0xa9, 0x9f, 0x5f, 0x32, 0xc9, 0xc5, 0xd2, 0x43, 0x42, 0xe6, 0xf7, 0x7f, 0x24,
	0x04, 0x6e, 0x9d, 0x27, 0xa6, 0x50, 0xfd, 0x6c, 0x02, 0x9b, 0xc6, 0xe8, 0x9a, 0x0d, 0xd9, 0xed,
	0xb1, 0x14, 0x93, 0x3a, 0xf9, 0x4f, 0x2f, 0xce, 0x72, 0xb6, 0x27, 0xdd, 0xc9, 0x96, 0x7c, 0x93,
	0x3b, 0x18, 0xbe, 0x93, 0xee, 0xe6, 0xd4, 0xa1, 0x55, 0x97, 0xb0, 0x1f, 0xa5, 0xf3, 0xc1, 0xd5,
	0xe0, 0x66, 0xf4, 0x00, 0xa5, 0xf8, 0xa2, 0x3c, 0x13, 0x3e, 0x79, 0x85, 0x62, 0xee, 0x6b, 0xdb,
	0xa0, 0x4f, 0xfe, 0x31, 0x8c, 0x5a, 0xee, 0xf5, 0x22, 0xa0, 0xf6, 0xc8, 0x12, 0x1b, 0xce, 0x20,
	0xa2, 0x4f, 0x21, 0x4a, 0xe5, 0xc2, 0x3d, 0x51, 0x52, 0xc9, 0x14, 0x0a, 0xb9, 0xd1, 0xac, 0xf1,
	0xe3, 0x17, 0xbd, 0xba, 0x86, 0x63, 0xe7, 0x89, 0x9a, 0x05, 0xf6, 0xb5, 0x0b, 0xb8, 0xca, 0x35,
	0x85, 0xc0, 0xb7, 0xc4, 0x76, 0x16, 0x8d, 0xe1, 0x68, 0x8a, 0x16, 0x83, 0x09, 0xea, 0x14, 0x0e,
	0x38, 0xbe, 0x28, 0x67, 0xd3, 0xf0, 0xf2, 0xfc, 0xf5, 0xf4, 0x6d, 0xb8, 0xdd, 0x2c, 0x4b, 0x4d,
	0xeb, 0xca, 0x75, 0x7c, 0xaf, 0xeb, 0xd0, 0xc6, 0xc3, 0xaa, 0xea, 0x6d, 0x5c, 0xde, 0xe9, 0x6a,
	0xe7, 0xa7, 0x2e, 0x0f, 0x05, 0x3f, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xdf, 0xad, 0xb5, 0x5a,
	0x74, 0x01, 0x00, 0x00,
}
