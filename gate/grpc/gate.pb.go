// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gate.proto

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

//多播
type PushMsgReq struct {
	Keys                 []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	Op                   int32    `protobuf:"varint,2,opt,name=op,proto3" json:"op,omitempty"`
	Body                 []byte   `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushMsgReq) Reset()         { *m = PushMsgReq{} }
func (m *PushMsgReq) String() string { return proto.CompactTextString(m) }
func (*PushMsgReq) ProtoMessage()    {}
func (*PushMsgReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{0}
}

func (m *PushMsgReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushMsgReq.Unmarshal(m, b)
}
func (m *PushMsgReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushMsgReq.Marshal(b, m, deterministic)
}
func (m *PushMsgReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushMsgReq.Merge(m, src)
}
func (m *PushMsgReq) XXX_Size() int {
	return xxx_messageInfo_PushMsgReq.Size(m)
}
func (m *PushMsgReq) XXX_DiscardUnknown() {
	xxx_messageInfo_PushMsgReq.DiscardUnknown(m)
}

var xxx_messageInfo_PushMsgReq proto.InternalMessageInfo

func (m *PushMsgReq) GetKeys() []string {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *PushMsgReq) GetOp() int32 {
	if m != nil {
		return m.Op
	}
	return 0
}

func (m *PushMsgReq) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type PushMsgReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushMsgReply) Reset()         { *m = PushMsgReply{} }
func (m *PushMsgReply) String() string { return proto.CompactTextString(m) }
func (*PushMsgReply) ProtoMessage()    {}
func (*PushMsgReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{1}
}

func (m *PushMsgReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushMsgReply.Unmarshal(m, b)
}
func (m *PushMsgReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushMsgReply.Marshal(b, m, deterministic)
}
func (m *PushMsgReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushMsgReply.Merge(m, src)
}
func (m *PushMsgReply) XXX_Size() int {
	return xxx_messageInfo_PushMsgReply.Size(m)
}
func (m *PushMsgReply) XXX_DiscardUnknown() {
	xxx_messageInfo_PushMsgReply.DiscardUnknown(m)
}

var xxx_messageInfo_PushMsgReply proto.InternalMessageInfo

//广播
type BroadcastReq struct {
	Op                   int32    `protobuf:"varint,1,opt,name=op,proto3" json:"op,omitempty"`
	Body                 []byte   `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Speed                int32    `protobuf:"varint,3,opt,name=speed,proto3" json:"speed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BroadcastReq) Reset()         { *m = BroadcastReq{} }
func (m *BroadcastReq) String() string { return proto.CompactTextString(m) }
func (*BroadcastReq) ProtoMessage()    {}
func (*BroadcastReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{2}
}

func (m *BroadcastReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadcastReq.Unmarshal(m, b)
}
func (m *BroadcastReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadcastReq.Marshal(b, m, deterministic)
}
func (m *BroadcastReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadcastReq.Merge(m, src)
}
func (m *BroadcastReq) XXX_Size() int {
	return xxx_messageInfo_BroadcastReq.Size(m)
}
func (m *BroadcastReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadcastReq.DiscardUnknown(m)
}

var xxx_messageInfo_BroadcastReq proto.InternalMessageInfo

func (m *BroadcastReq) GetOp() int32 {
	if m != nil {
		return m.Op
	}
	return 0
}

func (m *BroadcastReq) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *BroadcastReq) GetSpeed() int32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

type BroadcastReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BroadcastReply) Reset()         { *m = BroadcastReply{} }
func (m *BroadcastReply) String() string { return proto.CompactTextString(m) }
func (*BroadcastReply) ProtoMessage()    {}
func (*BroadcastReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_743bb58a714d8b7d, []int{3}
}

func (m *BroadcastReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadcastReply.Unmarshal(m, b)
}
func (m *BroadcastReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadcastReply.Marshal(b, m, deterministic)
}
func (m *BroadcastReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadcastReply.Merge(m, src)
}
func (m *BroadcastReply) XXX_Size() int {
	return xxx_messageInfo_BroadcastReply.Size(m)
}
func (m *BroadcastReply) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadcastReply.DiscardUnknown(m)
}

var xxx_messageInfo_BroadcastReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PushMsgReq)(nil), "bubble.gate.PushMsgReq")
	proto.RegisterType((*PushMsgReply)(nil), "bubble.gate.PushMsgReply")
	proto.RegisterType((*BroadcastReq)(nil), "bubble.gate.BroadcastReq")
	proto.RegisterType((*BroadcastReply)(nil), "bubble.gate.BroadcastReply")
}

func init() { proto.RegisterFile("gate.proto", fileDescriptor_743bb58a714d8b7d) }

var fileDescriptor_743bb58a714d8b7d = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd0, 0x31, 0x4b, 0xc4, 0x30,
	0x14, 0x07, 0x70, 0x92, 0x6b, 0x4f, 0xee, 0x59, 0x8a, 0x3c, 0x04, 0xeb, 0xb9, 0x94, 0x4c, 0x9d,
	0x32, 0xe8, 0xec, 0x72, 0x0a, 0xba, 0x08, 0x92, 0xd1, 0x2d, 0xb9, 0x86, 0x0a, 0x06, 0xf2, 0x6c,
	0x72, 0x43, 0xbe, 0x82, 0x9f, 0x5a, 0x9a, 0x93, 0xbb, 0x8a, 0xde, 0xf6, 0x92, 0xf7, 0xf8, 0x25,
	0xff, 0x07, 0x30, 0xe8, 0x68, 0x25, 0x8d, 0x3e, 0x7a, 0x3c, 0x37, 0x3b, 0x63, 0x9c, 0x95, 0xd3,
	0x95, 0x78, 0x04, 0x78, 0xdd, 0x85, 0xf7, 0x97, 0x30, 0x28, 0xfb, 0x89, 0x08, 0xc5, 0x87, 0x4d,
	0xa1, 0x61, 0xed, 0xa2, 0x5b, 0xa9, 0x5c, 0x63, 0x0d, 0xdc, 0x53, 0xc3, 0x5b, 0xd6, 0x95, 0x8a,
	0x7b, 0x9a, 0x66, 0x8c, 0xef, 0x53, 0xb3, 0x68, 0x59, 0x57, 0xa9, 0x5c, 0x8b, 0x1a, 0xaa, 0x83,
	0x42, 0x2e, 0x89, 0x67, 0xa8, 0x36, 0xa3, 0xd7, 0xfd, 0x56, 0x87, 0x38, 0xb9, 0x7b, 0x83, 0xfd,
	0x31, 0xf8, 0xd1, 0xc0, 0x4b, 0x28, 0x03, 0x59, 0xdb, 0x67, 0xb8, 0x54, 0xfb, 0x83, 0xb8, 0x80,
	0x7a, 0x26, 0x91, 0x4b, 0xb7, 0x5f, 0x0c, 0x8a, 0x27, 0x1d, 0x2d, 0xde, 0xc3, 0xd9, 0xcf, 0xa3,
	0x78, 0x25, 0x67, 0x99, 0xe4, 0x31, 0xd0, 0xfa, 0xfa, 0xff, 0x06, 0xb9, 0x84, 0x0f, 0xb0, 0x3a,
	0xc8, 0xf8, 0x7b, 0x6e, 0xfe, 0xf7, 0xf5, 0xcd, 0xa9, 0x16, 0xb9, 0xb4, 0x59, 0xbe, 0x15, 0xc3,
	0x48, 0x5b, 0xb3, 0xcc, 0xab, 0xbd, 0xfb, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x32, 0xf3, 0x2b,
	0x68, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GateClient is the client API for Gate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GateClient interface {
	PushMsg(ctx context.Context, in *PushMsgReq, opts ...grpc.CallOption) (*PushMsgReply, error)
	Broadcast(ctx context.Context, in *BroadcastReq, opts ...grpc.CallOption) (*BroadcastReply, error)
}

type gateClient struct {
	cc *grpc.ClientConn
}

func NewGateClient(cc *grpc.ClientConn) GateClient {
	return &gateClient{cc}
}

func (c *gateClient) PushMsg(ctx context.Context, in *PushMsgReq, opts ...grpc.CallOption) (*PushMsgReply, error) {
	out := new(PushMsgReply)
	err := c.cc.Invoke(ctx, "/bubble.gate.Gate/PushMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gateClient) Broadcast(ctx context.Context, in *BroadcastReq, opts ...grpc.CallOption) (*BroadcastReply, error) {
	out := new(BroadcastReply)
	err := c.cc.Invoke(ctx, "/bubble.gate.Gate/Broadcast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GateServer is the server API for Gate service.
type GateServer interface {
	PushMsg(context.Context, *PushMsgReq) (*PushMsgReply, error)
	Broadcast(context.Context, *BroadcastReq) (*BroadcastReply, error)
}

// UnimplementedGateServer can be embedded to have forward compatible implementations.
type UnimplementedGateServer struct {
}

func (*UnimplementedGateServer) PushMsg(ctx context.Context, req *PushMsgReq) (*PushMsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushMsg not implemented")
}
func (*UnimplementedGateServer) Broadcast(ctx context.Context, req *BroadcastReq) (*BroadcastReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}

func RegisterGateServer(s *grpc.Server, srv GateServer) {
	s.RegisterService(&_Gate_serviceDesc, srv)
}

func _Gate_PushMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMsgReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateServer).PushMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bubble.gate.Gate/PushMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateServer).PushMsg(ctx, req.(*PushMsgReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gate_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GateServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bubble.gate.Gate/Broadcast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GateServer).Broadcast(ctx, req.(*BroadcastReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Gate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bubble.gate.Gate",
	HandlerType: (*GateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushMsg",
			Handler:    _Gate_PushMsg_Handler,
		},
		{
			MethodName: "Broadcast",
			Handler:    _Gate_Broadcast_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gate.proto",
}
