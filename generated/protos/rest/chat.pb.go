// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package rest

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type StudentQuestions struct {
	Message              []*RestChatMessage `protobuf:"bytes,1,rep,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StudentQuestions) Reset()         { *m = StudentQuestions{} }
func (m *StudentQuestions) String() string { return proto.CompactTextString(m) }
func (*StudentQuestions) ProtoMessage()    {}
func (*StudentQuestions) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}

func (m *StudentQuestions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StudentQuestions.Unmarshal(m, b)
}
func (m *StudentQuestions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StudentQuestions.Marshal(b, m, deterministic)
}
func (m *StudentQuestions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StudentQuestions.Merge(m, src)
}
func (m *StudentQuestions) XXX_Size() int {
	return xxx_messageInfo_StudentQuestions.Size(m)
}
func (m *StudentQuestions) XXX_DiscardUnknown() {
	xxx_messageInfo_StudentQuestions.DiscardUnknown(m)
}

var xxx_messageInfo_StudentQuestions proto.InternalMessageInfo

func (m *StudentQuestions) GetMessage() []*RestChatMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

type RestChatMessage struct {
	Message              *RestMessage `protobuf:"bytes,1,opt,name=message,proto3" json:"1,omitempty"`
	Code                 string       `protobuf:"bytes,2,opt,name=code,proto3" json:"2,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RestChatMessage) Reset()         { *m = RestChatMessage{} }
func (m *RestChatMessage) String() string { return proto.CompactTextString(m) }
func (*RestChatMessage) ProtoMessage()    {}
func (*RestChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{2}
}

func (m *RestChatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestChatMessage.Unmarshal(m, b)
}
func (m *RestChatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestChatMessage.Marshal(b, m, deterministic)
}
func (m *RestChatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestChatMessage.Merge(m, src)
}
func (m *RestChatMessage) XXX_Size() int {
	return xxx_messageInfo_RestChatMessage.Size(m)
}
func (m *RestChatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RestChatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RestChatMessage proto.InternalMessageInfo

func (m *RestChatMessage) GetMessage() *RestMessage {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *RestChatMessage) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type RestMessage struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"1,omitempty"`
	TimeStamp            string   `protobuf:"bytes,2,opt,name=timeStamp,proto3" json:"2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestMessage) Reset()         { *m = RestMessage{} }
func (m *RestMessage) String() string { return proto.CompactTextString(m) }
func (*RestMessage) ProtoMessage()    {}
func (*RestMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{3}
}

func (m *RestMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestMessage.Unmarshal(m, b)
}
func (m *RestMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestMessage.Marshal(b, m, deterministic)
}
func (m *RestMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestMessage.Merge(m, src)
}
func (m *RestMessage) XXX_Size() int {
	return xxx_messageInfo_RestMessage.Size(m)
}
func (m *RestMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RestMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RestMessage proto.InternalMessageInfo

func (m *RestMessage) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *RestMessage) GetTimeStamp() string {
	if m != nil {
		return m.TimeStamp
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*StudentQuestions)(nil), "StudentQuestions")
	proto.RegisterType((*RestChatMessage)(nil), "RestChatMessage")
	proto.RegisterType((*RestMessage)(nil), "RestMessage")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x50, 0x3d, 0x4b, 0xc6, 0x30,
	0x10, 0x26, 0x7e, 0xbc, 0xa5, 0x57, 0xc1, 0xd7, 0xb8, 0x14, 0x71, 0x78, 0xc9, 0x20, 0xc5, 0x21,
	0x48, 0xdc, 0x5d, 0xa4, 0x38, 0x75, 0x30, 0xfd, 0x05, 0x31, 0x3d, 0x6c, 0x87, 0x24, 0xa5, 0xb9,
	0x0a, 0xfe, 0x7b, 0x21, 0x5a, 0x5a, 0xea, 0x76, 0xf7, 0x7c, 0x71, 0xf7, 0x00, 0xd8, 0xde, 0x90,
	0x1c, 0xa7, 0x40, 0x41, 0x64, 0x70, 0x59, 0xbb, 0x91, 0xbe, 0xc5, 0x0b, 0x1c, 0x5b, 0x9a, 0x3b,
	0xf4, 0xf4, 0x3e, 0x63, 0xa4, 0x21, 0xf8, 0xc8, 0x1f, 0x21, 0x73, 0x18, 0xa3, 0xf9, 0xc4, 0x92,
	0x9d, 0xce, 0xab, 0x42, 0x1d, 0xa5, 0xc6, 0x48, 0xaf, 0xbd, 0xa1, 0xe6, 0x17, 0xd7, 0x8b, 0x40,
	0x34, 0x70, 0xbd, 0xe3, 0xf8, 0xc3, 0xd6, 0xce, 0xaa, 0x42, 0x5d, 0x25, 0xfb, 0xde, 0xca, 0x39,
	0x5c, 0xd8, 0xd0, 0x61, 0x79, 0x76, 0x62, 0x55, 0xae, 0xd3, 0x2c, 0x6a, 0x28, 0x36, 0x5a, 0x5e,
	0x42, 0x66, 0x83, 0x27, 0xf4, 0x94, 0xa2, 0x72, 0xbd, 0xac, 0xfc, 0x1e, 0x72, 0x1a, 0x1c, 0xb6,
	0x64, 0xdc, 0xf8, 0x97, 0xb0, 0x02, 0xaa, 0x5e, 0xaf, 0x6a, 0x71, 0xfa, 0x1a, 0x2c, 0x72, 0x05,
	0xb7, 0x6f, 0x48, 0xff, 0x7e, 0x3d, 0xc8, 0xd4, 0xc3, 0xdd, 0x8d, 0xdc, 0x53, 0x4f, 0xec, 0xe3,
	0x90, 0xca, 0x7a, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xf7, 0x54, 0x70, 0x3a, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RestChatServiceClient is the client API for RestChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RestChatServiceClient interface {
	GetStudentQuestions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (RestChatService_GetStudentQuestionsClient, error)
}

type restChatServiceClient struct {
	cc *grpc.ClientConn
}

func NewRestChatServiceClient(cc *grpc.ClientConn) RestChatServiceClient {
	return &restChatServiceClient{cc}
}

func (c *restChatServiceClient) GetStudentQuestions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (RestChatService_GetStudentQuestionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RestChatService_serviceDesc.Streams[0], "/RestChatService/GetStudentQuestions", opts...)
	if err != nil {
		return nil, err
	}
	x := &restChatServiceGetStudentQuestionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RestChatService_GetStudentQuestionsClient interface {
	Recv() (*StudentQuestions, error)
	grpc.ClientStream
}

type restChatServiceGetStudentQuestionsClient struct {
	grpc.ClientStream
}

func (x *restChatServiceGetStudentQuestionsClient) Recv() (*StudentQuestions, error) {
	m := new(StudentQuestions)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RestChatServiceServer is the server API for RestChatService service.
type RestChatServiceServer interface {
	GetStudentQuestions(*Empty, RestChatService_GetStudentQuestionsServer) error
}

// UnimplementedRestChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRestChatServiceServer struct {
}

func (*UnimplementedRestChatServiceServer) GetStudentQuestions(req *Empty, srv RestChatService_GetStudentQuestionsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStudentQuestions not implemented")
}

func RegisterRestChatServiceServer(s *grpc.Server, srv RestChatServiceServer) {
	s.RegisterService(&_RestChatService_serviceDesc, srv)
}

func _RestChatService_GetStudentQuestions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RestChatServiceServer).GetStudentQuestions(m, &restChatServiceGetStudentQuestionsServer{stream})
}

type RestChatService_GetStudentQuestionsServer interface {
	Send(*StudentQuestions) error
	grpc.ServerStream
}

type restChatServiceGetStudentQuestionsServer struct {
	grpc.ServerStream
}

func (x *restChatServiceGetStudentQuestionsServer) Send(m *StudentQuestions) error {
	return x.ServerStream.SendMsg(m)
}

var _RestChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RestChatService",
	HandlerType: (*RestChatServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStudentQuestions",
			Handler:       _RestChatService_GetStudentQuestions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat.proto",
}
