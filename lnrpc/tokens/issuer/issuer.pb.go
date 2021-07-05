// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tokens/issuer/issuer.proto

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

type SignTokenPurchaseRequest struct {
	Offer                *replicator.TokenOffer `protobuf:"bytes,1,opt,name=offer,proto3" json:"offer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SignTokenPurchaseRequest) Reset()         { *m = SignTokenPurchaseRequest{} }
func (m *SignTokenPurchaseRequest) String() string { return proto.CompactTextString(m) }
func (*SignTokenPurchaseRequest) ProtoMessage()    {}
func (*SignTokenPurchaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8faa554d9d56586, []int{0}
}

func (m *SignTokenPurchaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignTokenPurchaseRequest.Unmarshal(m, b)
}
func (m *SignTokenPurchaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignTokenPurchaseRequest.Marshal(b, m, deterministic)
}
func (m *SignTokenPurchaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignTokenPurchaseRequest.Merge(m, src)
}
func (m *SignTokenPurchaseRequest) XXX_Size() int {
	return xxx_messageInfo_SignTokenPurchaseRequest.Size(m)
}
func (m *SignTokenPurchaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignTokenPurchaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignTokenPurchaseRequest proto.InternalMessageInfo

func (m *SignTokenPurchaseRequest) GetOffer() *replicator.TokenOffer {
	if m != nil {
		return m.Offer
	}
	return nil
}

type SignTokenPurchaseResponse struct {
	IssuerSignature      string   `protobuf:"bytes,1,opt,name=issuer_signature,json=issuerSignature,proto3" json:"issuer_signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignTokenPurchaseResponse) Reset()         { *m = SignTokenPurchaseResponse{} }
func (m *SignTokenPurchaseResponse) String() string { return proto.CompactTextString(m) }
func (*SignTokenPurchaseResponse) ProtoMessage()    {}
func (*SignTokenPurchaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8faa554d9d56586, []int{1}
}

func (m *SignTokenPurchaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignTokenPurchaseResponse.Unmarshal(m, b)
}
func (m *SignTokenPurchaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignTokenPurchaseResponse.Marshal(b, m, deterministic)
}
func (m *SignTokenPurchaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignTokenPurchaseResponse.Merge(m, src)
}
func (m *SignTokenPurchaseResponse) XXX_Size() int {
	return xxx_messageInfo_SignTokenPurchaseResponse.Size(m)
}
func (m *SignTokenPurchaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignTokenPurchaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignTokenPurchaseResponse proto.InternalMessageInfo

func (m *SignTokenPurchaseResponse) GetIssuerSignature() string {
	if m != nil {
		return m.IssuerSignature
	}
	return ""
}

type SignTokenSellRequest struct {
	Offer                *replicator.TokenOffer `protobuf:"bytes,1,opt,name=offer,proto3" json:"offer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SignTokenSellRequest) Reset()         { *m = SignTokenSellRequest{} }
func (m *SignTokenSellRequest) String() string { return proto.CompactTextString(m) }
func (*SignTokenSellRequest) ProtoMessage()    {}
func (*SignTokenSellRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8faa554d9d56586, []int{2}
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

type SignTokenSellResponse struct {
	IssuerSignature      string   `protobuf:"bytes,1,opt,name=issuer_signature,json=issuerSignature,proto3" json:"issuer_signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignTokenSellResponse) Reset()         { *m = SignTokenSellResponse{} }
func (m *SignTokenSellResponse) String() string { return proto.CompactTextString(m) }
func (*SignTokenSellResponse) ProtoMessage()    {}
func (*SignTokenSellResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8faa554d9d56586, []int{3}
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

type IssueTokenRequest struct {
	Offer                *replicator.TokenOffer `protobuf:"bytes,1,opt,name=offer,proto3" json:"offer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *IssueTokenRequest) Reset()         { *m = IssueTokenRequest{} }
func (m *IssueTokenRequest) String() string { return proto.CompactTextString(m) }
func (*IssueTokenRequest) ProtoMessage()    {}
func (*IssueTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8faa554d9d56586, []int{4}
}

func (m *IssueTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueTokenRequest.Unmarshal(m, b)
}
func (m *IssueTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueTokenRequest.Marshal(b, m, deterministic)
}
func (m *IssueTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueTokenRequest.Merge(m, src)
}
func (m *IssueTokenRequest) XXX_Size() int {
	return xxx_messageInfo_IssueTokenRequest.Size(m)
}
func (m *IssueTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IssueTokenRequest proto.InternalMessageInfo

func (m *IssueTokenRequest) GetOffer() *replicator.TokenOffer {
	if m != nil {
		return m.Offer
	}
	return nil
}

type UpdateTokenRequest struct {
	Offer                *replicator.TokenOffer `protobuf:"bytes,1,opt,name=offer,proto3" json:"offer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *UpdateTokenRequest) Reset()         { *m = UpdateTokenRequest{} }
func (m *UpdateTokenRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTokenRequest) ProtoMessage()    {}
func (*UpdateTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8faa554d9d56586, []int{5}
}

func (m *UpdateTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTokenRequest.Unmarshal(m, b)
}
func (m *UpdateTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTokenRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTokenRequest.Merge(m, src)
}
func (m *UpdateTokenRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTokenRequest.Size(m)
}
func (m *UpdateTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTokenRequest proto.InternalMessageInfo

func (m *UpdateTokenRequest) GetOffer() *replicator.TokenOffer {
	if m != nil {
		return m.Offer
	}
	return nil
}

type RevokeTokenRequest struct {
	TokenName            string   `protobuf:"bytes,1,opt,name=token_name,json=tokenName,proto3" json:"token_name,omitempty"`
	Login                string   `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeTokenRequest) Reset()         { *m = RevokeTokenRequest{} }
func (m *RevokeTokenRequest) String() string { return proto.CompactTextString(m) }
func (*RevokeTokenRequest) ProtoMessage()    {}
func (*RevokeTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8faa554d9d56586, []int{6}
}

func (m *RevokeTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeTokenRequest.Unmarshal(m, b)
}
func (m *RevokeTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeTokenRequest.Marshal(b, m, deterministic)
}
func (m *RevokeTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeTokenRequest.Merge(m, src)
}
func (m *RevokeTokenRequest) XXX_Size() int {
	return xxx_messageInfo_RevokeTokenRequest.Size(m)
}
func (m *RevokeTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeTokenRequest proto.InternalMessageInfo

func (m *RevokeTokenRequest) GetTokenName() string {
	if m != nil {
		return m.TokenName
	}
	return ""
}

func (m *RevokeTokenRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func init() {
	proto.RegisterType((*SignTokenPurchaseRequest)(nil), "issuer.SignTokenPurchaseRequest")
	proto.RegisterType((*SignTokenPurchaseResponse)(nil), "issuer.SignTokenPurchaseResponse")
	proto.RegisterType((*SignTokenSellRequest)(nil), "issuer.SignTokenSellRequest")
	proto.RegisterType((*SignTokenSellResponse)(nil), "issuer.SignTokenSellResponse")
	proto.RegisterType((*IssueTokenRequest)(nil), "issuer.IssueTokenRequest")
	proto.RegisterType((*UpdateTokenRequest)(nil), "issuer.UpdateTokenRequest")
	proto.RegisterType((*RevokeTokenRequest)(nil), "issuer.RevokeTokenRequest")
}

func init() { proto.RegisterFile("tokens/issuer/issuer.proto", fileDescriptor_e8faa554d9d56586) }

var fileDescriptor_e8faa554d9d56586 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xdf, 0x4b, 0xeb, 0x30,
	0x14, 0xc7, 0xd9, 0x60, 0x83, 0x9d, 0x71, 0xb9, 0x77, 0x61, 0x77, 0x6c, 0xd5, 0xc1, 0xec, 0x93,
	0x82, 0xb6, 0x32, 0xff, 0x00, 0xd9, 0x50, 0x71, 0x20, 0x2a, 0x9d, 0x82, 0xf8, 0x32, 0xba, 0xee,
	0xac, 0x2b, 0x6b, 0x9b, 0x9a, 0xa4, 0x82, 0x7f, 0xb9, 0xaf, 0xd2, 0xa4, 0x75, 0x9d, 0x75, 0x03,
	0xe7, 0x43, 0x69, 0xce, 0x8f, 0x7c, 0x38, 0xc9, 0xf7, 0x1b, 0xd0, 0x04, 0x5d, 0x62, 0xc8, 0x4d,
	0x8f, 0xf3, 0x18, 0x59, 0xfa, 0x33, 0x22, 0x46, 0x05, 0x25, 0x55, 0x15, 0x69, 0x7a, 0xda, 0xc3,
	0x30, 0xf2, 0x3d, 0xc7, 0x16, 0x94, 0xe5, 0x96, 0xaa, 0x57, 0xdb, 0x73, 0x29, 0x75, 0x7d, 0x34,
	0x65, 0x34, 0x8d, 0xe7, 0x26, 0x06, 0x91, 0x78, 0x53, 0x45, 0xfd, 0x1a, 0xda, 0x63, 0xcf, 0x0d,
	0x1f, 0x12, 0xcc, 0x7d, 0xcc, 0x9c, 0x85, 0xcd, 0xd1, 0xc2, 0x97, 0x18, 0xb9, 0x20, 0xc7, 0x50,
	0xa1, 0xf3, 0x39, 0xb2, 0x76, 0xa9, 0x57, 0x3a, 0xac, 0xf7, 0x5b, 0x46, 0x0e, 0x2d, 0x37, 0xdc,
	0x25, 0x55, 0x4b, 0x35, 0xe9, 0x57, 0xd0, 0xf9, 0x86, 0xc4, 0x23, 0x1a, 0x72, 0x24, 0x47, 0xf0,
	0x4f, 0x4d, 0x3c, 0xe1, 0x9e, 0x1b, 0xda, 0x22, 0x66, 0x28, 0xa9, 0x35, 0xeb, 0xaf, 0xca, 0x8f,
	0xb3, 0xb4, 0x7e, 0x01, 0xcd, 0x4f, 0xce, 0x18, 0x7d, 0x7f, 0xb7, 0x69, 0x86, 0xf0, 0xff, 0x0b,
	0xe5, 0xe7, 0x93, 0x0c, 0xa0, 0x31, 0x4a, 0x52, 0x12, 0xb2, 0xeb, 0x18, 0xe4, 0x31, 0x9a, 0xd9,
	0xe2, 0x37, 0x8c, 0x11, 0x10, 0x0b, 0x5f, 0xe9, 0x72, 0x9d, 0xd1, 0x05, 0x90, 0xda, 0x4f, 0x42,
	0x3b, 0xc8, 0x4e, 0x50, 0x93, 0x99, 0x5b, 0x3b, 0x40, 0xd2, 0x84, 0x8a, 0x4f, 0x5d, 0x2f, 0x6c,
	0x97, 0x65, 0x45, 0x05, 0xfd, 0xf7, 0x32, 0x54, 0xe5, 0x91, 0x18, 0x79, 0x82, 0x46, 0x41, 0x2e,
	0xd2, 0x33, 0x52, 0x97, 0x6d, 0xf2, 0x84, 0x76, 0xb0, 0xa5, 0x23, 0xbd, 0xe1, 0x1b, 0xf8, 0xb3,
	0x76, 0xf5, 0x64, 0xbf, 0xb0, 0x27, 0xa7, 0xab, 0xd6, 0xdd, 0x50, 0x4d, 0x69, 0xe7, 0x00, 0x2b,
	0x11, 0x48, 0x27, 0x6b, 0x2e, 0x08, 0xa3, 0xb5, 0x0c, 0xe5, 0x73, 0x23, 0xf3, 0xb9, 0x71, 0x99,
	0xf8, 0x9c, 0x0c, 0xa0, 0x9e, 0x93, 0x80, 0x68, 0x19, 0xa1, 0xa8, 0xcb, 0x36, 0x44, 0x4e, 0x81,
	0x15, 0xa2, 0x28, 0xcb, 0x26, 0xc4, 0xb0, 0xff, 0x7c, 0xea, 0x7a, 0x62, 0x11, 0x4f, 0x0d, 0x87,
	0x06, 0x66, 0xb4, 0x14, 0x27, 0x8e, 0xcd, 0x17, 0xc9, 0x62, 0x66, 0xfa, 0x61, 0xf2, 0xb1, 0xc8,
	0x31, 0xd7, 0x5e, 0xfc, 0xb4, 0x2a, 0x19, 0x67, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x79, 0x83,
	0x9a, 0x67, 0x09, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IssuerClient is the client API for Issuer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IssuerClient interface {
	// Returns token purchase signature for further registration along with
	// offer via Replication Server
	SignTokenPurchase(ctx context.Context, in *SignTokenPurchaseRequest, opts ...grpc.CallOption) (*SignTokenPurchaseResponse, error)
	// Returns token sell signature for further registration along with offer
	// via Replication Server
	SignTokenSell(ctx context.Context, in *SignTokenSellRequest, opts ...grpc.CallOption) (*SignTokenSellResponse, error)
	// Issue new token with given data. Request data equal to token purchase
	// data, because it is token offer.
	IssueToken(ctx context.Context, in *IssueTokenRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Update token inforamtion with given data. Request data equal to token purchase
	// data, because it is token offer.
	UpdateToken(ctx context.Context, in *UpdateTokenRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Revoke token - delete information about token by his name.
	RevokeToken(ctx context.Context, in *RevokeTokenRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type issuerClient struct {
	cc *grpc.ClientConn
}

func NewIssuerClient(cc *grpc.ClientConn) IssuerClient {
	return &issuerClient{cc}
}

func (c *issuerClient) SignTokenPurchase(ctx context.Context, in *SignTokenPurchaseRequest, opts ...grpc.CallOption) (*SignTokenPurchaseResponse, error) {
	out := new(SignTokenPurchaseResponse)
	err := c.cc.Invoke(ctx, "/issuer.Issuer/SignTokenPurchase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issuerClient) SignTokenSell(ctx context.Context, in *SignTokenSellRequest, opts ...grpc.CallOption) (*SignTokenSellResponse, error) {
	out := new(SignTokenSellResponse)
	err := c.cc.Invoke(ctx, "/issuer.Issuer/SignTokenSell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issuerClient) IssueToken(ctx context.Context, in *IssueTokenRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/issuer.Issuer/IssueToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issuerClient) UpdateToken(ctx context.Context, in *UpdateTokenRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/issuer.Issuer/UpdateToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *issuerClient) RevokeToken(ctx context.Context, in *RevokeTokenRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/issuer.Issuer/RevokeToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IssuerServer is the server API for Issuer service.
type IssuerServer interface {
	// Returns token purchase signature for further registration along with
	// offer via Replication Server
	SignTokenPurchase(context.Context, *SignTokenPurchaseRequest) (*SignTokenPurchaseResponse, error)
	// Returns token sell signature for further registration along with offer
	// via Replication Server
	SignTokenSell(context.Context, *SignTokenSellRequest) (*SignTokenSellResponse, error)
	// Issue new token with given data. Request data equal to token purchase
	// data, because it is token offer.
	IssueToken(context.Context, *IssueTokenRequest) (*empty.Empty, error)
	// Update token inforamtion with given data. Request data equal to token purchase
	// data, because it is token offer.
	UpdateToken(context.Context, *UpdateTokenRequest) (*empty.Empty, error)
	// Revoke token - delete information about token by his name.
	RevokeToken(context.Context, *RevokeTokenRequest) (*empty.Empty, error)
}

// UnimplementedIssuerServer can be embedded to have forward compatible implementations.
type UnimplementedIssuerServer struct {
}

func (*UnimplementedIssuerServer) SignTokenPurchase(ctx context.Context, req *SignTokenPurchaseRequest) (*SignTokenPurchaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignTokenPurchase not implemented")
}
func (*UnimplementedIssuerServer) SignTokenSell(ctx context.Context, req *SignTokenSellRequest) (*SignTokenSellResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignTokenSell not implemented")
}
func (*UnimplementedIssuerServer) IssueToken(ctx context.Context, req *IssueTokenRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IssueToken not implemented")
}
func (*UnimplementedIssuerServer) UpdateToken(ctx context.Context, req *UpdateTokenRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateToken not implemented")
}
func (*UnimplementedIssuerServer) RevokeToken(ctx context.Context, req *RevokeTokenRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeToken not implemented")
}

func RegisterIssuerServer(s *grpc.Server, srv IssuerServer) {
	s.RegisterService(&_Issuer_serviceDesc, srv)
}

func _Issuer_SignTokenPurchase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignTokenPurchaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssuerServer).SignTokenPurchase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issuer.Issuer/SignTokenPurchase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssuerServer).SignTokenPurchase(ctx, req.(*SignTokenPurchaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Issuer_SignTokenSell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignTokenSellRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssuerServer).SignTokenSell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issuer.Issuer/SignTokenSell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssuerServer).SignTokenSell(ctx, req.(*SignTokenSellRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Issuer_IssueToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssuerServer).IssueToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issuer.Issuer/IssueToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssuerServer).IssueToken(ctx, req.(*IssueTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Issuer_UpdateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssuerServer).UpdateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issuer.Issuer/UpdateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssuerServer).UpdateToken(ctx, req.(*UpdateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Issuer_RevokeToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IssuerServer).RevokeToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/issuer.Issuer/RevokeToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IssuerServer).RevokeToken(ctx, req.(*RevokeTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Issuer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "issuer.Issuer",
	HandlerType: (*IssuerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignTokenPurchase",
			Handler:    _Issuer_SignTokenPurchase_Handler,
		},
		{
			MethodName: "SignTokenSell",
			Handler:    _Issuer_SignTokenSell_Handler,
		},
		{
			MethodName: "IssueToken",
			Handler:    _Issuer_IssueToken_Handler,
		},
		{
			MethodName: "UpdateToken",
			Handler:    _Issuer_UpdateToken_Handler,
		},
		{
			MethodName: "RevokeToken",
			Handler:    _Issuer_RevokeToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tokens/issuer/issuer.proto",
}