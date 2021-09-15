// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/issuer/issuer.proto

package issuer

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	replicator "github.com/pkt-cash/pktd/lnd/lnrpc/protos/replicator"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// LockTokenRequest send information about token
type LockTokenRequest struct {
	// token — token name
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	// count — number of tokens to lock
	Count int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	// htlc — hash of preimagine
	Htlc string `protobuf:"bytes,3,opt,name=htlc,proto3" json:"htlc,omitempty"`
	// recipient — wallet addres of new owner of tokens
	Recipient string `protobuf:"bytes,4,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// proof_count — lock expiration time in PKT blocks
	ProofCount           int32    `protobuf:"varint,5,opt,name=proof_count,json=proofCount,proto3" json:"proof_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LockTokenRequest) Reset()         { *m = LockTokenRequest{} }
func (m *LockTokenRequest) String() string { return proto.CompactTextString(m) }
func (*LockTokenRequest) ProtoMessage()    {}
func (*LockTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_868a71152c014334, []int{0}
}

func (m *LockTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockTokenRequest.Unmarshal(m, b)
}
func (m *LockTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockTokenRequest.Marshal(b, m, deterministic)
}
func (m *LockTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockTokenRequest.Merge(m, src)
}
func (m *LockTokenRequest) XXX_Size() int {
	return xxx_messageInfo_LockTokenRequest.Size(m)
}
func (m *LockTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LockTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LockTokenRequest proto.InternalMessageInfo

func (m *LockTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LockTokenRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *LockTokenRequest) GetHtlc() string {
	if m != nil {
		return m.Htlc
	}
	return ""
}

func (m *LockTokenRequest) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *LockTokenRequest) GetProofCount() int32 {
	if m != nil {
		return m.ProofCount
	}
	return 0
}

// LockTokenResponse response with hash of lock
type LockTokenResponse struct {
	// lock_id — hash of lock
	LockId               string   `protobuf:"bytes,1,opt,name=lock_id,json=lockId,proto3" json:"lock_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LockTokenResponse) Reset()         { *m = LockTokenResponse{} }
func (m *LockTokenResponse) String() string { return proto.CompactTextString(m) }
func (*LockTokenResponse) ProtoMessage()    {}
func (*LockTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_868a71152c014334, []int{1}
}

func (m *LockTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockTokenResponse.Unmarshal(m, b)
}
func (m *LockTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockTokenResponse.Marshal(b, m, deterministic)
}
func (m *LockTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockTokenResponse.Merge(m, src)
}
func (m *LockTokenResponse) XXX_Size() int {
	return xxx_messageInfo_LockTokenResponse.Size(m)
}
func (m *LockTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LockTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LockTokenResponse proto.InternalMessageInfo

func (m *LockTokenResponse) GetLockId() string {
	if m != nil {
		return m.LockId
	}
	return ""
}

// SignTokenSellRequest — info about sell token
type SignTokenSellRequest struct {
	// offer on sell token
	Offer                *replicator.TokenOffer `protobuf:"bytes,1,opt,name=offer,proto3" json:"offer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SignTokenSellRequest) Reset()         { *m = SignTokenSellRequest{} }
func (m *SignTokenSellRequest) String() string { return proto.CompactTextString(m) }
func (*SignTokenSellRequest) ProtoMessage()    {}
func (*SignTokenSellRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_868a71152c014334, []int{2}
}

func (m *SignTokenSellRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignTokenSellRequest.Unmarshal(m, b)
}
func (m *SignTokenSellRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignTokenSellRequest.Marshal(b, m, deterministic)
}
func (m *SignTokenSellRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignTokenSellRequest.Merge(m, src)
}
func (m *SignTokenSellRequest) XXX_Size() int {
	return xxx_messageInfo_SignTokenSellRequest.Size(m)
}
func (m *SignTokenSellRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignTokenSellRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignTokenSellRequest proto.InternalMessageInfo

func (m *SignTokenSellRequest) GetOffer() *replicator.TokenOffer {
	if m != nil {
		return m.Offer
	}
	return nil
}

// SignTokenSellResponse — info about sign offer
type SignTokenSellResponse struct {
	// issuer_signature signature on sell offer
	IssuerSignature      string   `protobuf:"bytes,1,opt,name=issuer_signature,json=issuerSignature,proto3" json:"issuer_signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignTokenSellResponse) Reset()         { *m = SignTokenSellResponse{} }
func (m *SignTokenSellResponse) String() string { return proto.CompactTextString(m) }
func (*SignTokenSellResponse) ProtoMessage()    {}
func (*SignTokenSellResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_868a71152c014334, []int{3}
}

func (m *SignTokenSellResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignTokenSellResponse.Unmarshal(m, b)
}
func (m *SignTokenSellResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignTokenSellResponse.Marshal(b, m, deterministic)
}
func (m *SignTokenSellResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignTokenSellResponse.Merge(m, src)
}
func (m *SignTokenSellResponse) XXX_Size() int {
	return xxx_messageInfo_SignTokenSellResponse.Size(m)
}
func (m *SignTokenSellResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignTokenSellResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignTokenSellResponse proto.InternalMessageInfo

func (m *SignTokenSellResponse) GetIssuerSignature() string {
	if m != nil {
		return m.IssuerSignature
	}
	return ""
}

func init() {
	proto.RegisterType((*LockTokenRequest)(nil), "issuer.LockTokenRequest")
	proto.RegisterType((*LockTokenResponse)(nil), "issuer.LockTokenResponse")
	proto.RegisterType((*SignTokenSellRequest)(nil), "issuer.SignTokenSellRequest")
	proto.RegisterType((*SignTokenSellResponse)(nil), "issuer.SignTokenSellResponse")
}

func init() { proto.RegisterFile("protos/issuer/issuer.proto", fileDescriptor_868a71152c014334) }

var fileDescriptor_868a71152c014334 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0xb6, 0xa5, 0xa8, 0x67, 0x4c, 0x0c, 0x6b, 0x8c, 0x10, 0x36, 0x2d, 0xca, 0x55, 0x91,
	0x46, 0x82, 0xca, 0x0b, 0xa0, 0x0d, 0x84, 0x26, 0x55, 0x42, 0xa4, 0x5c, 0x71, 0x53, 0xb5, 0xee,
	0x49, 0x6a, 0x25, 0x8b, 0x8d, 0xed, 0x20, 0xf1, 0x18, 0x3c, 0x0d, 0xaf, 0x87, 0xfc, 0x93, 0x2d,
	0xab, 0xda, 0x8b, 0xaa, 0xfe, 0x7e, 0xf4, 0xe5, 0x3b, 0xf1, 0x09, 0xc4, 0x42, 0x72, 0xcd, 0x55,
	0xce, 0x94, 0xea, 0x50, 0xfa, 0xbf, 0xcc, 0x92, 0x64, 0xe4, 0x50, 0xfc, 0xb6, 0xe2, 0xbc, 0x6a,
	0x30, 0xb7, 0xec, 0xaa, 0x2b, 0x73, 0xbc, 0x17, 0xfa, 0x8f, 0x33, 0xc5, 0xa9, 0x0f, 0x90, 0x28,
	0x1a, 0x46, 0x97, 0x9a, 0xcb, 0xc1, 0xd1, 0x79, 0xd2, 0xbf, 0x01, 0x9c, 0xce, 0x38, 0xad, 0x7f,
	0xf0, 0x1a, 0xdb, 0x02, 0x7f, 0x75, 0xa8, 0x34, 0x39, 0x83, 0x50, 0x1b, 0x1c, 0x05, 0x49, 0x30,
	0x19, 0x17, 0x0e, 0x18, 0x96, 0xf2, 0xae, 0xd5, 0xd1, 0x41, 0x12, 0x4c, 0x0e, 0x0b, 0x07, 0x08,
	0x81, 0xa3, 0x8d, 0x6e, 0x68, 0x74, 0x68, 0xad, 0xf6, 0x4c, 0x2e, 0x60, 0x2c, 0x91, 0x32, 0xc1,
	0xb0, 0xd5, 0xd1, 0x91, 0x15, 0x1e, 0x09, 0x72, 0x05, 0xc7, 0x42, 0x72, 0x5e, 0x2e, 0x5c, 0x5a,
	0x98, 0x04, 0x93, 0xb0, 0x00, 0x4b, 0xdd, 0x1a, 0x26, 0xbd, 0x86, 0x97, 0x83, 0x4a, 0x4a, 0xf0,
	0x56, 0x21, 0x79, 0x0d, 0xcf, 0x1a, 0x4e, 0xeb, 0x05, 0x5b, 0xfb, 0x56, 0x23, 0x03, 0xef, 0xd6,
	0xe9, 0x67, 0x38, 0x9b, 0xb3, 0xaa, 0xb5, 0xee, 0x39, 0x36, 0x4d, 0x3f, 0xc4, 0x35, 0x84, 0xbc,
	0x2c, 0x51, 0x5a, 0xfb, 0xf1, 0xf4, 0x3c, 0x1b, 0xcc, 0x6e, 0xcd, 0xdf, 0x8c, 0x5a, 0x38, 0x53,
	0x7a, 0x03, 0xaf, 0xb6, 0x52, 0xfc, 0x73, 0xdf, 0xc1, 0xa9, 0x7b, 0xd7, 0x0b, 0xc5, 0xaa, 0x76,
	0xa9, 0x3b, 0x89, 0xbe, 0xc0, 0x0b, 0xc7, 0xcf, 0x7b, 0x7a, 0xfa, 0xef, 0x00, 0x4e, 0xee, 0x1c,
	0x87, 0xf2, 0x37, 0xa3, 0x48, 0x66, 0x70, 0xf2, 0x24, 0x95, 0x5c, 0x64, 0xfe, 0x1a, 0x77, 0x55,
	0x8e, 0x2f, 0xf7, 0xa8, 0xbe, 0xca, 0x2d, 0x80, 0x8d, 0xb7, 0x0a, 0xb9, 0x1c, 0x0e, 0xf4, 0xc8,
	0xf7, 0x59, 0xe7, 0x99, 0x5b, 0x8d, 0xac, 0x5f, 0x8d, 0xec, 0x8b, 0x59, 0x0d, 0xf2, 0x1d, 0x9e,
	0x7f, 0x45, 0x6d, 0xad, 0x33, 0xa6, 0x34, 0xb9, 0x1a, 0xc6, 0x0c, 0x95, 0x3e, 0x28, 0xd9, 0x6f,
	0xf0, 0xbd, 0x3e, 0xc1, 0xf8, 0xe1, 0xbe, 0x48, 0xd4, 0xcf, 0xb0, 0xbd, 0x55, 0xf1, 0x9b, 0x1d,
	0x8a, 0x4b, 0xb8, 0x99, 0xfe, 0xfc, 0x50, 0x31, 0xbd, 0xe9, 0x56, 0x19, 0xe5, 0xf7, 0xb9, 0xa8,
	0xf5, 0x7b, 0xba, 0x54, 0x1b, 0x73, 0x58, 0xe7, 0x4d, 0x6b, 0x7e, 0x52, 0xd0, 0xfc, 0xc9, 0xf7,
	0xb0, 0x1a, 0x59, 0xf8, 0xf1, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x05, 0x5d, 0xbd, 0xf5, 0x27,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IssuerServiceClient is the client API for IssuerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IssuerServiceClient interface {
	// SignTokenSell Returns token sell signature for further
	// registration along with offer via Replication Server
	SignTokenSell(ctx context.Context, in *SignTokenSellRequest, opts ...grpc.CallOption) (*SignTokenSellResponse, error)
	// IssueToken — Issue new token with given data. Request data equal to
	// token purchase data, because it is token offer.
	IssueToken(ctx context.Context, in *replicator.IssueTokenRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// GetTokenList — Return list of issued token with information about
	// expiration time and fix price.
	GetTokenList(ctx context.Context, in *replicator.GetTokenListRequest, opts ...grpc.CallOption) (*replicator.GetTokenListResponse, error)
	// LockToken — Return hash of lock token for verify htlc and information
	// about transaction
	LockToken(ctx context.Context, in *LockTokenRequest, opts ...grpc.CallOption) (*LockTokenResponse, error)
}

type issuerServiceClient struct {
	cc *grpc.ClientConn
}

func NewIssuerServiceClient(cc *grpc.ClientConn) IssuerServiceClient {
	return &issuerServiceClient{cc}
}

func (c *issuerServiceClient) SignTokenSell(ctx context.Context, in *SignTokenSellRequest, opts ...grpc.CallOption) (*SignTokenSellResponse, error) {
	out := new(SignTokenSellResponse)
	err := c.cc.Invoke(ctx, "/issuer.IssuerService/SignTokenSell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issuerServiceClient) IssueToken(ctx context.Context, in *replicator.IssueTokenRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/issuer.IssuerService/IssueToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issuerServiceClient) GetTokenList(ctx context.Context, in *replicator.GetTokenListRequest, opts ...grpc.CallOption) (*replicator.GetTokenListResponse, error) {
	out := new(replicator.GetTokenListResponse)
	err := c.cc.Invoke(ctx, "/issuer.IssuerService/GetTokenList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issuerServiceClient) LockToken(ctx context.Context, in *LockTokenRequest, opts ...grpc.CallOption) (*LockTokenResponse, error) {
	out := new(LockTokenResponse)
	err := c.cc.Invoke(ctx, "/issuer.IssuerService/LockToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IssuerServiceServer is the server API for IssuerService service.
type IssuerServiceServer interface {
	// SignTokenSell Returns token sell signature for further
	// registration along with offer via Replication Server
	SignTokenSell(context.Context, *SignTokenSellRequest) (*SignTokenSellResponse, error)
	// IssueToken — Issue new token with given data. Request data equal to
	// token purchase data, because it is token offer.
	IssueToken(context.Context, *replicator.IssueTokenRequest) (*empty.Empty, error)
	// GetTokenList — Return list of issued token with information about
	// expiration time and fix price.
	GetTokenList(context.Context, *replicator.GetTokenListRequest) (*replicator.GetTokenListResponse, error)
	// LockToken — Return hash of lock token for verify htlc and information
	// about transaction
	LockToken(context.Context, *LockTokenRequest) (*LockTokenResponse, error)
}

// UnimplementedIssuerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedIssuerServiceServer struct {
}

func (*UnimplementedIssuerServiceServer) SignTokenSell(ctx context.Context, req *SignTokenSellRequest) (*SignTokenSellResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignTokenSell not implemented")
}
func (*UnimplementedIssuerServiceServer) IssueToken(ctx context.Context, req *replicator.IssueTokenRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueToken not implemented")
}
func (*UnimplementedIssuerServiceServer) GetTokenList(ctx context.Context, req *replicator.GetTokenListRequest) (*replicator.GetTokenListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenList not implemented")
}
func (*UnimplementedIssuerServiceServer) LockToken(ctx context.Context, req *LockTokenRequest) (*LockTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LockToken not implemented")
}

func RegisterIssuerServiceServer(s *grpc.Server, srv IssuerServiceServer) {
	s.RegisterService(&_IssuerService_serviceDesc, srv)
}

func _IssuerService_SignTokenSell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignTokenSellRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssuerServiceServer).SignTokenSell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issuer.IssuerService/SignTokenSell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssuerServiceServer).SignTokenSell(ctx, req.(*SignTokenSellRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssuerService_IssueToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(replicator.IssueTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssuerServiceServer).IssueToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issuer.IssuerService/IssueToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssuerServiceServer).IssueToken(ctx, req.(*replicator.IssueTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssuerService_GetTokenList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(replicator.GetTokenListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssuerServiceServer).GetTokenList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issuer.IssuerService/GetTokenList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssuerServiceServer).GetTokenList(ctx, req.(*replicator.GetTokenListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IssuerService_LockToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssuerServiceServer).LockToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issuer.IssuerService/LockToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssuerServiceServer).LockToken(ctx, req.(*LockTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IssuerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "issuer.IssuerService",
	HandlerType: (*IssuerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignTokenSell",
			Handler:    _IssuerService_SignTokenSell_Handler,
		},
		{
			MethodName: "IssueToken",
			Handler:    _IssuerService_IssueToken_Handler,
		},
		{
			MethodName: "GetTokenList",
			Handler:    _IssuerService_GetTokenList_Handler,
		},
		{
			MethodName: "LockToken",
			Handler:    _IssuerService_LockToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/issuer/issuer.proto",
}
