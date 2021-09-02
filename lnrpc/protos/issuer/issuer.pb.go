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
	return fileDescriptor_868a71152c014334, []int{0}
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
	return fileDescriptor_868a71152c014334, []int{1}
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
	proto.RegisterType((*SignTokenSellRequest)(nil), "issuer.SignTokenSellRequest")
	proto.RegisterType((*SignTokenSellResponse)(nil), "issuer.SignTokenSellResponse")
}

func init() { proto.RegisterFile("protos/issuer/issuer.proto", fileDescriptor_868a71152c014334) }

var fileDescriptor_868a71152c014334 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x51, 0x4b, 0xfb, 0x30,
	0x14, 0xc5, 0xe9, 0xc3, 0x7f, 0xf0, 0xbf, 0x3a, 0x94, 0xa0, 0x43, 0xaa, 0xc3, 0xd1, 0x27, 0x05,
	0x4d, 0xa4, 0x7e, 0x83, 0xa9, 0x88, 0x30, 0x10, 0x5b, 0x9f, 0x7c, 0x91, 0xb6, 0xde, 0x76, 0x61,
	0x5d, 0x13, 0x93, 0x54, 0xf0, 0x83, 0xfb, 0x2e, 0x4d, 0x5a, 0xcc, 0xc4, 0x3d, 0x94, 0xe6, 0x9e,
	0x7b, 0xfa, 0xe3, 0xf4, 0x10, 0x08, 0xa5, 0x12, 0x46, 0x68, 0xc6, 0xb5, 0x6e, 0x51, 0xf5, 0x2f,
	0x6a, 0x45, 0x32, 0x72, 0x53, 0x78, 0x5c, 0x09, 0x51, 0xd5, 0xc8, 0xac, 0x9a, 0xb7, 0x25, 0xc3,
	0xb5, 0x34, 0x9f, 0xce, 0x14, 0x46, 0x3d, 0x40, 0xa1, 0xac, 0x79, 0x91, 0x19, 0xa1, 0xbc, 0xa3,
	0xf3, 0x44, 0xb7, 0x70, 0x90, 0xf2, 0xaa, 0x79, 0x16, 0x2b, 0x6c, 0x52, 0xac, 0xeb, 0x04, 0xdf,
	0x5b, 0xd4, 0x86, 0x5c, 0xc0, 0x3f, 0x51, 0x96, 0xa8, 0x8e, 0x82, 0x59, 0x70, 0xb6, 0x13, 0x4f,
	0xa8, 0xf7, 0xa5, 0x35, 0x3f, 0x76, 0xdb, 0xc4, 0x99, 0xa2, 0x39, 0x1c, 0xfe, 0xa2, 0x68, 0x29,
	0x1a, 0x8d, 0xe4, 0x1c, 0xf6, 0x5d, 0xd2, 0x57, 0xcd, 0xab, 0x26, 0x33, 0xad, 0x42, 0x4b, 0xfc,
	0x9f, 0xec, 0x39, 0x3d, 0x1d, 0xe4, 0xf8, 0x2b, 0x80, 0xf1, 0x83, 0xd3, 0x50, 0x7d, 0xf0, 0x02,
	0xc9, 0x02, 0xc6, 0x1b, 0x54, 0x72, 0x42, 0xfb, 0x12, 0xfe, 0x8a, 0x1c, 0x4e, 0xb7, 0x6c, 0xfb,
	0x28, 0x37, 0x00, 0x16, 0x6f, 0x37, 0x64, 0xea, 0xff, 0xd0, 0x8f, 0x3e, 0xb0, 0x26, 0xd4, 0x15,
	0x4b, 0x87, 0x62, 0xe9, 0x5d, 0x57, 0x2c, 0x79, 0x82, 0xdd, 0x7b, 0x34, 0xd6, 0xba, 0xe0, 0xda,
	0x90, 0x53, 0x1f, 0xe3, 0x6f, 0x06, 0xd0, 0x6c, 0xbb, 0xc1, 0xe5, 0x9a, 0xc7, 0x2f, 0x57, 0x15,
	0x37, 0xcb, 0x36, 0xa7, 0x85, 0x58, 0x33, 0xb9, 0x32, 0x97, 0x45, 0xa6, 0x97, 0xdd, 0xe1, 0x8d,
	0xd5, 0x4d, 0xf7, 0x28, 0x59, 0xb0, 0x8d, 0xbb, 0x90, 0x8f, 0xec, 0x78, 0xfd, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0x79, 0x61, 0xac, 0x59, 0x23, 0x02, 0x00, 0x00,
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
	// GetTokenList — return list of issued token with infomation about
	// expiration time and fix price.
	GetTokenList(ctx context.Context, in *replicator.GetTokenListRequest, opts ...grpc.CallOption) (*replicator.GetTokenListResponse, error)
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

// IssuerServiceServer is the server API for IssuerService service.
type IssuerServiceServer interface {
	// SignTokenSell Returns token sell signature for further
	// registration along with offer via Replication Server
	SignTokenSell(context.Context, *SignTokenSellRequest) (*SignTokenSellResponse, error)
	// IssueToken — Issue new token with given data. Request data equal to
	// token purchase data, because it is token offer.
	IssueToken(context.Context, *replicator.IssueTokenRequest) (*empty.Empty, error)
	// GetTokenList — return list of issued token with infomation about
	// expiration time and fix price.
	GetTokenList(context.Context, *replicator.GetTokenListRequest) (*replicator.GetTokenListResponse, error)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/issuer/issuer.proto",
}
