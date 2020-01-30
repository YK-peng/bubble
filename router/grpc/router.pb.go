// Code generated by protoc-gen-go. DO NOT EDIT.
// source: router.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type BaseRequest struct {
	SessionKey           []byte   `protobuf:"bytes,1,opt,name=sessionKey,proto3" json:"sessionKey,omitempty"`
	Uin                  uint32   `protobuf:"varint,2,opt,name=uin,proto3" json:"uin,omitempty"`
	DeviceID             []byte   `protobuf:"bytes,3,opt,name=deviceID,proto3" json:"deviceID,omitempty"`
	ClientVersion        int32    `protobuf:"varint,4,opt,name=clientVersion,proto3" json:"clientVersion,omitempty"`
	DeviceType           []byte   `protobuf:"bytes,5,opt,name=deviceType,proto3" json:"deviceType,omitempty"`
	Scene                uint32   `protobuf:"varint,6,opt,name=scene,proto3" json:"scene,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseRequest) Reset()         { *m = BaseRequest{} }
func (m *BaseRequest) String() string { return proto.CompactTextString(m) }
func (*BaseRequest) ProtoMessage()    {}
func (*BaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{0}
}

func (m *BaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseRequest.Unmarshal(m, b)
}
func (m *BaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseRequest.Marshal(b, m, deterministic)
}
func (m *BaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseRequest.Merge(m, src)
}
func (m *BaseRequest) XXX_Size() int {
	return xxx_messageInfo_BaseRequest.Size(m)
}
func (m *BaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BaseRequest proto.InternalMessageInfo

func (m *BaseRequest) GetSessionKey() []byte {
	if m != nil {
		return m.SessionKey
	}
	return nil
}

func (m *BaseRequest) GetUin() uint32 {
	if m != nil {
		return m.Uin
	}
	return 0
}

func (m *BaseRequest) GetDeviceID() []byte {
	if m != nil {
		return m.DeviceID
	}
	return nil
}

func (m *BaseRequest) GetClientVersion() int32 {
	if m != nil {
		return m.ClientVersion
	}
	return 0
}

func (m *BaseRequest) GetDeviceType() []byte {
	if m != nil {
		return m.DeviceType
	}
	return nil
}

func (m *BaseRequest) GetScene() uint32 {
	if m != nil {
		return m.Scene
	}
	return 0
}

type BaseResponse struct {
	Ret                  int32    `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=errMsg,proto3" json:"errMsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseResponse) Reset()         { *m = BaseResponse{} }
func (m *BaseResponse) String() string { return proto.CompactTextString(m) }
func (*BaseResponse) ProtoMessage()    {}
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{1}
}

func (m *BaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseResponse.Unmarshal(m, b)
}
func (m *BaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseResponse.Marshal(b, m, deterministic)
}
func (m *BaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseResponse.Merge(m, src)
}
func (m *BaseResponse) XXX_Size() int {
	return xxx_messageInfo_BaseResponse.Size(m)
}
func (m *BaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BaseResponse proto.InternalMessageInfo

func (m *BaseResponse) GetRet() int32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *BaseResponse) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type CliReq struct {
	BaseRequest          *BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	Op                   int32        `protobuf:"varint,2,opt,name=op,proto3" json:"op,omitempty"`
	Body                 []byte       `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CliReq) Reset()         { *m = CliReq{} }
func (m *CliReq) String() string { return proto.CompactTextString(m) }
func (*CliReq) ProtoMessage()    {}
func (*CliReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{2}
}

func (m *CliReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CliReq.Unmarshal(m, b)
}
func (m *CliReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CliReq.Marshal(b, m, deterministic)
}
func (m *CliReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CliReq.Merge(m, src)
}
func (m *CliReq) XXX_Size() int {
	return xxx_messageInfo_CliReq.Size(m)
}
func (m *CliReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CliReq.DiscardUnknown(m)
}

var xxx_messageInfo_CliReq proto.InternalMessageInfo

func (m *CliReq) GetBaseRequest() *BaseRequest {
	if m != nil {
		return m.BaseRequest
	}
	return nil
}

func (m *CliReq) GetOp() int32 {
	if m != nil {
		return m.Op
	}
	return 0
}

func (m *CliReq) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type CliRep struct {
	BaseResponse         *BaseResponse `protobuf:"bytes,1,opt,name=baseResponse,proto3" json:"baseResponse,omitempty"`
	Op                   int32         `protobuf:"varint,2,opt,name=op,proto3" json:"op,omitempty"`
	Body                 []byte        `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CliRep) Reset()         { *m = CliRep{} }
func (m *CliRep) String() string { return proto.CompactTextString(m) }
func (*CliRep) ProtoMessage()    {}
func (*CliRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{3}
}

func (m *CliRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CliRep.Unmarshal(m, b)
}
func (m *CliRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CliRep.Marshal(b, m, deterministic)
}
func (m *CliRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CliRep.Merge(m, src)
}
func (m *CliRep) XXX_Size() int {
	return xxx_messageInfo_CliRep.Size(m)
}
func (m *CliRep) XXX_DiscardUnknown() {
	xxx_messageInfo_CliRep.DiscardUnknown(m)
}

var xxx_messageInfo_CliRep proto.InternalMessageInfo

func (m *CliRep) GetBaseResponse() *BaseResponse {
	if m != nil {
		return m.BaseResponse
	}
	return nil
}

func (m *CliRep) GetOp() int32 {
	if m != nil {
		return m.Op
	}
	return 0
}

func (m *CliRep) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterType((*BaseRequest)(nil), "bubble.router.BaseRequest")
	proto.RegisterType((*BaseResponse)(nil), "bubble.router.BaseResponse")
	proto.RegisterType((*CliReq)(nil), "bubble.router.CliReq")
	proto.RegisterType((*CliRep)(nil), "bubble.router.CliRep")
}

func init() { proto.RegisterFile("router.proto", fileDescriptor_367072455c71aedc) }

var fileDescriptor_367072455c71aedc = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x49, 0xdb, 0x44, 0x9d, 0xa6, 0x22, 0x83, 0xca, 0x52, 0x41, 0x4a, 0xf0, 0xd0, 0x53,
	0x0e, 0xf5, 0xa0, 0x07, 0x41, 0xac, 0xbd, 0x88, 0x78, 0x59, 0xc4, 0x83, 0x37, 0x37, 0x1d, 0xca,
	0x4a, 0xcc, 0x6e, 0x77, 0x93, 0x42, 0x7f, 0x99, 0x7f, 0x4f, 0xb2, 0x1b, 0x6c, 0x22, 0x3d, 0x78,
	0x9b, 0x79, 0xcc, 0xcc, 0xf7, 0xf6, 0xb1, 0x10, 0x1b, 0x55, 0x95, 0x64, 0x52, 0x6d, 0x54, 0xa9,
	0x70, 0x24, 0x2a, 0x21, 0x72, 0x4a, 0xbd, 0x98, 0x7c, 0x07, 0x30, 0x9c, 0x7f, 0x58, 0xe2, 0xb4,
	0xae, 0xc8, 0x96, 0x78, 0x09, 0x60, 0xc9, 0x5a, 0xa9, 0x8a, 0x67, 0xda, 0xb2, 0x60, 0x12, 0x4c,
	0x63, 0xde, 0x52, 0xf0, 0x04, 0xfa, 0x95, 0x2c, 0x58, 0x6f, 0x12, 0x4c, 0x47, 0xbc, 0x2e, 0x71,
	0x0c, 0x87, 0x4b, 0xda, 0xc8, 0x8c, 0x9e, 0x16, 0xac, 0xef, 0xe6, 0x7f, 0x7b, 0xbc, 0x82, 0x51,
	0x96, 0x4b, 0x2a, 0xca, 0x37, 0x32, 0xf5, 0x05, 0x36, 0x98, 0x04, 0xd3, 0x90, 0x77, 0xc5, 0x9a,
	0xe9, 0x37, 0x5e, 0xb7, 0x9a, 0x58, 0xe8, 0x99, 0x3b, 0x05, 0x4f, 0x21, 0xb4, 0x19, 0x15, 0xc4,
	0x22, 0x47, 0xf5, 0x4d, 0x72, 0x0b, 0xb1, 0x37, 0x6e, 0xb5, 0x2a, 0x2c, 0xd5, 0xce, 0x0c, 0x95,
	0xce, 0x72, 0xc8, 0xeb, 0x12, 0xcf, 0x21, 0x22, 0x63, 0x5e, 0xec, 0xca, 0xd9, 0x3d, 0xe2, 0x4d,
	0x97, 0x7c, 0x42, 0xf4, 0x98, 0x4b, 0x4e, 0x6b, 0xbc, 0x83, 0xa1, 0xd8, 0x3d, 0xde, 0xed, 0x0e,
	0x67, 0xe3, 0xb4, 0x13, 0x51, 0xda, 0x8a, 0x87, 0xb7, 0xc7, 0xf1, 0x18, 0x7a, 0x4a, 0xbb, 0xdb,
	0x21, 0xef, 0x29, 0x8d, 0x08, 0x03, 0xa1, 0x96, 0xdb, 0x26, 0x05, 0x57, 0x27, 0x5f, 0x0d, 0x4b,
	0xe3, 0x3d, 0xc4, 0xa2, 0xe5, 0xb7, 0x81, 0x5d, 0xec, 0x85, 0xf9, 0x11, 0xde, 0x59, 0xf8, 0x0f,
	0x6e, 0xf6, 0x00, 0x91, 0x3f, 0x84, 0x37, 0x70, 0xb0, 0xa0, 0x5c, 0x6e, 0xc8, 0xe0, 0xd9, 0x1f,
	0x86, 0x7f, 0xfc, 0x78, 0xaf, 0xac, 0xe7, 0xd1, 0xfb, 0x60, 0x65, 0x74, 0x26, 0x22, 0xf7, 0x5f,
	0xae, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x39, 0x28, 0xca, 0x30, 0x3f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouterClient is the client API for Router service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouterClient interface {
	Deliver(ctx context.Context, in *CliReq, opts ...grpc.CallOption) (*CliRep, error)
}

type routerClient struct {
	cc *grpc.ClientConn
}

func NewRouterClient(cc *grpc.ClientConn) RouterClient {
	return &routerClient{cc}
}

func (c *routerClient) Deliver(ctx context.Context, in *CliReq, opts ...grpc.CallOption) (*CliRep, error) {
	out := new(CliRep)
	err := c.cc.Invoke(ctx, "/bubble.router.router/Deliver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RouterServer is the server API for Router service.
type RouterServer interface {
	Deliver(context.Context, *CliReq) (*CliRep, error)
}

// UnimplementedRouterServer can be embedded to have forward compatible implementations.
type UnimplementedRouterServer struct {
}

func (*UnimplementedRouterServer) Deliver(ctx context.Context, req *CliReq) (*CliRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deliver not implemented")
}

func RegisterRouterServer(s *grpc.Server, srv RouterServer) {
	s.RegisterService(&_Router_serviceDesc, srv)
}

func _Router_Deliver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CliReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouterServer).Deliver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bubble.router.router/Deliver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouterServer).Deliver(ctx, req.(*CliReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Router_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bubble.router.router",
	HandlerType: (*RouterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deliver",
			Handler:    _Router_Deliver_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "router.proto",
}
