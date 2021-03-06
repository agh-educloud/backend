// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/grpc/chat.proto

package educloud

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

type ChatMessage struct {
	Sender               *User    `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Message              *Message `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Code                 string   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatMessage) Reset()         { *m = ChatMessage{} }
func (m *ChatMessage) String() string { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()    {}
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_d927a3aef9bca0b3, []int{0}
}

func (m *ChatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatMessage.Unmarshal(m, b)
}
func (m *ChatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatMessage.Marshal(b, m, deterministic)
}
func (m *ChatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMessage.Merge(m, src)
}
func (m *ChatMessage) XXX_Size() int {
	return xxx_messageInfo_ChatMessage.Size(m)
}
func (m *ChatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMessage proto.InternalMessageInfo

func (m *ChatMessage) GetSender() *User {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *ChatMessage) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *ChatMessage) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type Message struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	TimeStamp            string   `protobuf:"bytes,2,opt,name=timeStamp,proto3" json:"timeStamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_d927a3aef9bca0b3, []int{1}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Message) GetTimeStamp() string {
	if m != nil {
		return m.TimeStamp
	}
	return ""
}

func init() {
	proto.RegisterType((*ChatMessage)(nil), "ChatMessage")
	proto.RegisterType((*Message)(nil), "Message")
}

func init() { proto.RegisterFile("protos/grpc/chat.proto", fileDescriptor_d927a3aef9bca0b3) }

var fileDescriptor_d927a3aef9bca0b3 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0x87, 0xa9, 0x7f, 0xb6, 0xdb, 0xa9, 0x20, 0xe4, 0x20, 0xa1, 0x28, 0x2c, 0xc5, 0x43, 0x4f,
	0x59, 0x59, 0x9f, 0x40, 0x3d, 0x7b, 0x49, 0xf1, 0x22, 0x78, 0x88, 0xc9, 0xb0, 0x5b, 0xb0, 0x4d,
	0xc9, 0x4c, 0xf7, 0xf9, 0x85, 0xb4, 0x61, 0x77, 0x6f, 0xc9, 0xf7, 0x25, 0xf3, 0x9b, 0x19, 0x78,
	0x18, 0x83, 0x67, 0x4f, 0xdb, 0x7d, 0x18, 0xed, 0xd6, 0x1e, 0x0c, 0xab, 0x08, 0xaa, 0x0b, 0x3e,
	0x11, 0x86, 0x85, 0xcb, 0x73, 0x4e, 0x6c, 0x78, 0xa2, 0xd9, 0xd4, 0x0e, 0xca, 0x8f, 0x83, 0xe1,
	0x4f, 0x24, 0x32, 0x7b, 0x14, 0x4f, 0xb0, 0x22, 0x1c, 0x1c, 0x06, 0x99, 0x6d, 0xb2, 0xa6, 0xdc,
	0xdd, 0xaa, 0x2f, 0xc2, 0xa0, 0x17, 0x28, 0x6a, 0xc8, 0xfb, 0xf9, 0xa5, 0xbc, 0x8a, 0x7e, 0xad,
	0x96, 0x9f, 0x3a, 0x09, 0x21, 0xe0, 0xc6, 0x7a, 0x87, 0xf2, 0x7a, 0x93, 0x35, 0x85, 0x8e, 0xe7,
	0xfa, 0x0d, 0xf2, 0x94, 0x20, 0x21, 0xb7, 0x7e, 0x60, 0x1c, 0x38, 0x46, 0x14, 0x3a, 0x5d, 0xc5,
	0x23, 0x14, 0xdc, 0xf5, 0xd8, 0xb2, 0xe9, 0xc7, 0x58, 0xbe, 0xd0, 0x27, 0xb0, 0xfb, 0x99, 0x1b,
	0x6d, 0x31, 0x1c, 0x3b, 0x8b, 0xa2, 0x81, 0x7b, 0x8d, 0x16, 0xbb, 0x23, 0x2e, 0x85, 0x49, 0xcc,
	0xbd, 0x56, 0x77, 0xea, 0x6c, 0xa0, 0x97, 0x4c, 0x3c, 0x43, 0xd9, 0xe2, 0xe0, 0x52, 0xfe, 0x85,
	0xae, 0x72, 0xd5, 0xc6, 0x6d, 0xbc, 0xc3, 0xf7, 0x1a, 0xdd, 0x64, 0xff, 0xfc, 0xe4, 0x7e, 0x57,
	0x71, 0x35, 0xaf, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x04, 0xb9, 0xd5, 0x66, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	ReceiveMessages(ctx context.Context, in *User, opts ...grpc.CallOption) (ChatService_ReceiveMessagesClient, error)
	SendMessage(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*Status, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) ReceiveMessages(ctx context.Context, in *User, opts ...grpc.CallOption) (ChatService_ReceiveMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[0], "/ChatService/ReceiveMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceReceiveMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_ReceiveMessagesClient interface {
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type chatServiceReceiveMessagesClient struct {
	grpc.ClientStream
}

func (x *chatServiceReceiveMessagesClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) SendMessage(ctx context.Context, in *ChatMessage, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ChatService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	ReceiveMessages(*User, ChatService_ReceiveMessagesServer) error
	SendMessage(context.Context, *ChatMessage) (*Status, error)
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) ReceiveMessages(req *User, srv ChatService_ReceiveMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveMessages not implemented")
}
func (*UnimplementedChatServiceServer) SendMessage(ctx context.Context, req *ChatMessage) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_ReceiveMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(User)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).ReceiveMessages(m, &chatServiceReceiveMessagesServer{stream})
}

type ChatService_ReceiveMessagesServer interface {
	Send(*ChatMessage) error
	grpc.ServerStream
}

type chatServiceReceiveMessagesServer struct {
	grpc.ServerStream
}

func (x *chatServiceReceiveMessagesServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ChatService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SendMessage(ctx, req.(*ChatMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMessage",
			Handler:    _ChatService_SendMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReceiveMessages",
			Handler:       _ChatService_ReceiveMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/grpc/chat.proto",
}
