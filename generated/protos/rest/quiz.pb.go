// Code generated by protoc-gen-go. DO NOT EDIT.
// source: quiz.proto

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

type RestQuizQuestionUuid struct {
	Uuid                 int32    `protobuf:"varint,1,opt,name=uuid,proto3" json:"1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RestQuizQuestionUuid) Reset()         { *m = RestQuizQuestionUuid{} }
func (m *RestQuizQuestionUuid) String() string { return proto.CompactTextString(m) }
func (*RestQuizQuestionUuid) ProtoMessage()    {}
func (*RestQuizQuestionUuid) Descriptor() ([]byte, []int) {
	return fileDescriptor_93138adfabfb5582, []int{0}
}

func (m *RestQuizQuestionUuid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RestQuizQuestionUuid.Unmarshal(m, b)
}
func (m *RestQuizQuestionUuid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RestQuizQuestionUuid.Marshal(b, m, deterministic)
}
func (m *RestQuizQuestionUuid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RestQuizQuestionUuid.Merge(m, src)
}
func (m *RestQuizQuestionUuid) XXX_Size() int {
	return xxx_messageInfo_RestQuizQuestionUuid.Size(m)
}
func (m *RestQuizQuestionUuid) XXX_DiscardUnknown() {
	xxx_messageInfo_RestQuizQuestionUuid.DiscardUnknown(m)
}

var xxx_messageInfo_RestQuizQuestionUuid proto.InternalMessageInfo

func (m *RestQuizQuestionUuid) GetUuid() int32 {
	if m != nil {
		return m.Uuid
	}
	return 0
}

type QuizQuestionStatistics struct {
	QuestionUuid               int32    `protobuf:"varint,1,opt,name=questionUuid,proto3" json:"1,omitempty"`
	PercentageOfCorrectAnswers float32  `protobuf:"fixed32,2,opt,name=percentageOfCorrectAnswers,proto3" json:"2,omitempty"`
	Participants               int32    `protobuf:"varint,3,opt,name=participants,proto3" json:"3,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *QuizQuestionStatistics) Reset()         { *m = QuizQuestionStatistics{} }
func (m *QuizQuestionStatistics) String() string { return proto.CompactTextString(m) }
func (*QuizQuestionStatistics) ProtoMessage()    {}
func (*QuizQuestionStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_93138adfabfb5582, []int{1}
}

func (m *QuizQuestionStatistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuizQuestionStatistics.Unmarshal(m, b)
}
func (m *QuizQuestionStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuizQuestionStatistics.Marshal(b, m, deterministic)
}
func (m *QuizQuestionStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuizQuestionStatistics.Merge(m, src)
}
func (m *QuizQuestionStatistics) XXX_Size() int {
	return xxx_messageInfo_QuizQuestionStatistics.Size(m)
}
func (m *QuizQuestionStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_QuizQuestionStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_QuizQuestionStatistics proto.InternalMessageInfo

func (m *QuizQuestionStatistics) GetQuestionUuid() int32 {
	if m != nil {
		return m.QuestionUuid
	}
	return 0
}

func (m *QuizQuestionStatistics) GetPercentageOfCorrectAnswers() float32 {
	if m != nil {
		return m.PercentageOfCorrectAnswers
	}
	return 0
}

func (m *QuizQuestionStatistics) GetParticipants() int32 {
	if m != nil {
		return m.Participants
	}
	return 0
}

type OpenQuizQuestionAnswers struct {
	Url                  []string `protobuf:"bytes,1,rep,name=url,proto3" json:"1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OpenQuizQuestionAnswers) Reset()         { *m = OpenQuizQuestionAnswers{} }
func (m *OpenQuizQuestionAnswers) String() string { return proto.CompactTextString(m) }
func (*OpenQuizQuestionAnswers) ProtoMessage()    {}
func (*OpenQuizQuestionAnswers) Descriptor() ([]byte, []int) {
	return fileDescriptor_93138adfabfb5582, []int{2}
}

func (m *OpenQuizQuestionAnswers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpenQuizQuestionAnswers.Unmarshal(m, b)
}
func (m *OpenQuizQuestionAnswers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpenQuizQuestionAnswers.Marshal(b, m, deterministic)
}
func (m *OpenQuizQuestionAnswers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpenQuizQuestionAnswers.Merge(m, src)
}
func (m *OpenQuizQuestionAnswers) XXX_Size() int {
	return xxx_messageInfo_OpenQuizQuestionAnswers.Size(m)
}
func (m *OpenQuizQuestionAnswers) XXX_DiscardUnknown() {
	xxx_messageInfo_OpenQuizQuestionAnswers.DiscardUnknown(m)
}

var xxx_messageInfo_OpenQuizQuestionAnswers proto.InternalMessageInfo

func (m *OpenQuizQuestionAnswers) GetUrl() []string {
	if m != nil {
		return m.Url
	}
	return nil
}

func init() {
	proto.RegisterType((*RestQuizQuestionUuid)(nil), "RestQuizQuestionUuid")
	proto.RegisterType((*QuizQuestionStatistics)(nil), "QuizQuestionStatistics")
	proto.RegisterType((*OpenQuizQuestionAnswers)(nil), "OpenQuizQuestionAnswers")
}

func init() { proto.RegisterFile("quiz.proto", fileDescriptor_93138adfabfb5582) }

var fileDescriptor_93138adfabfb5582 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xd9, 0xe6, 0xff, 0x17, 0x1c, 0x7b, 0x90, 0x25, 0xda, 0x98, 0x53, 0xc8, 0x29, 0x28,
	0x04, 0xd1, 0x9b, 0x07, 0x41, 0x14, 0x7a, 0xd2, 0xda, 0x04, 0x1f, 0x60, 0x4d, 0xc7, 0xb2, 0x50,
	0x36, 0xe9, 0xcc, 0xac, 0x42, 0x9f, 0xc6, 0x57, 0xf4, 0x0d, 0x24, 0xa9, 0x42, 0x84, 0x46, 0xf0,
	0x36, 0xfb, 0xf1, 0xfd, 0xbe, 0xd9, 0xfd, 0x16, 0x60, 0xed, 0xed, 0x26, 0x6f, 0xa8, 0x96, 0x3a,
	0x1e, 0xb3, 0x18, 0xf1, 0xbc, 0x3d, 0xa5, 0xa7, 0x10, 0x16, 0xc8, 0x32, 0xf7, 0x76, 0x33, 0xf7,
	0xc8, 0x62, 0x6b, 0xf7, 0xe4, 0xed, 0x42, 0x6b, 0xf8, 0xe7, 0xbd, 0x5d, 0x44, 0x2a, 0x51, 0xd9,
	0xff, 0xa2, 0x9b, 0xd3, 0x77, 0x05, 0xc7, 0x7d, 0x63, 0x29, 0x46, 0x2c, 0x8b, 0xad, 0x58, 0xa7,
	0x30, 0x5e, 0xf7, 0xf0, 0x2f, 0xec, 0x87, 0xa6, 0xaf, 0x21, 0x6e, 0x90, 0x2a, 0x74, 0x62, 0x96,
	0x38, 0x7b, 0xb9, 0xad, 0x89, 0xb0, 0x92, 0x1b, 0xc7, 0x6f, 0x48, 0x1c, 0x8d, 0x12, 0x95, 0x8d,
	0x8a, 0x5f, 0x1c, 0xed, 0x8e, 0xc6, 0x90, 0xd8, 0xca, 0x36, 0xc6, 0x09, 0x47, 0xc1, 0x76, 0x47,
	0x5f, 0x4b, 0xcf, 0x60, 0x32, 0x6b, 0xd0, 0xf5, 0x6f, 0xf9, 0x8d, 0x1f, 0x42, 0xe0, 0x69, 0x15,
	0xa9, 0x24, 0xc8, 0xf6, 0x8b, 0x76, 0xbc, 0xf8, 0x50, 0x10, 0x3e, 0x12, 0x32, 0x3a, 0x41, 0x6a,
	0x91, 0x12, 0xe9, 0xd5, 0x56, 0xa8, 0xaf, 0x20, 0xbc, 0xc3, 0x15, 0x2e, 0x8d, 0x60, 0x3f, 0x49,
	0x1f, 0xe5, 0xbb, 0xba, 0x8a, 0x0f, 0x3a, 0xb9, 0xec, 0x6a, 0xd5, 0xf7, 0x70, 0x32, 0x45, 0x19,
	0xa8, 0x69, 0x20, 0x60, 0x92, 0xef, 0xf6, 0x9f, 0x2b, 0xfd, 0x00, 0xf1, 0x14, 0x65, 0xe8, 0x4d,
	0x7f, 0xce, 0x7b, 0xde, 0xeb, 0xbe, 0xfd, 0xf2, 0x33, 0x00, 0x00, 0xff, 0xff, 0x59, 0x68, 0x08,
	0x5a, 0x12, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PresenterQuizServiceClient is the client API for PresenterQuizService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PresenterQuizServiceClient interface {
	DelegateQuizQuestion(ctx context.Context, in *RestQuizQuestionUuid, opts ...grpc.CallOption) (*RestStatus, error)
	GetQuizQuestionStatistics(ctx context.Context, in *RestQuizQuestionUuid, opts ...grpc.CallOption) (PresenterQuizService_GetQuizQuestionStatisticsClient, error)
	GetOpenQuizQuestionAnswers(ctx context.Context, in *RestQuizQuestionUuid, opts ...grpc.CallOption) (PresenterQuizService_GetOpenQuizQuestionAnswersClient, error)
}

type presenterQuizServiceClient struct {
	cc *grpc.ClientConn
}

func NewPresenterQuizServiceClient(cc *grpc.ClientConn) PresenterQuizServiceClient {
	return &presenterQuizServiceClient{cc}
}

func (c *presenterQuizServiceClient) DelegateQuizQuestion(ctx context.Context, in *RestQuizQuestionUuid, opts ...grpc.CallOption) (*RestStatus, error) {
	out := new(RestStatus)
	err := c.cc.Invoke(ctx, "/PresenterQuizService/DelegateQuizQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presenterQuizServiceClient) GetQuizQuestionStatistics(ctx context.Context, in *RestQuizQuestionUuid, opts ...grpc.CallOption) (PresenterQuizService_GetQuizQuestionStatisticsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PresenterQuizService_serviceDesc.Streams[0], "/PresenterQuizService/GetQuizQuestionStatistics", opts...)
	if err != nil {
		return nil, err
	}
	x := &presenterQuizServiceGetQuizQuestionStatisticsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PresenterQuizService_GetQuizQuestionStatisticsClient interface {
	Recv() (*QuizQuestionStatistics, error)
	grpc.ClientStream
}

type presenterQuizServiceGetQuizQuestionStatisticsClient struct {
	grpc.ClientStream
}

func (x *presenterQuizServiceGetQuizQuestionStatisticsClient) Recv() (*QuizQuestionStatistics, error) {
	m := new(QuizQuestionStatistics)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *presenterQuizServiceClient) GetOpenQuizQuestionAnswers(ctx context.Context, in *RestQuizQuestionUuid, opts ...grpc.CallOption) (PresenterQuizService_GetOpenQuizQuestionAnswersClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PresenterQuizService_serviceDesc.Streams[1], "/PresenterQuizService/GetOpenQuizQuestionAnswers", opts...)
	if err != nil {
		return nil, err
	}
	x := &presenterQuizServiceGetOpenQuizQuestionAnswersClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PresenterQuizService_GetOpenQuizQuestionAnswersClient interface {
	Recv() (*QuizQuestionStatistics, error)
	grpc.ClientStream
}

type presenterQuizServiceGetOpenQuizQuestionAnswersClient struct {
	grpc.ClientStream
}

func (x *presenterQuizServiceGetOpenQuizQuestionAnswersClient) Recv() (*QuizQuestionStatistics, error) {
	m := new(QuizQuestionStatistics)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PresenterQuizServiceServer is the server API for PresenterQuizService service.
type PresenterQuizServiceServer interface {
	DelegateQuizQuestion(context.Context, *RestQuizQuestionUuid) (*RestStatus, error)
	GetQuizQuestionStatistics(*RestQuizQuestionUuid, PresenterQuizService_GetQuizQuestionStatisticsServer) error
	GetOpenQuizQuestionAnswers(*RestQuizQuestionUuid, PresenterQuizService_GetOpenQuizQuestionAnswersServer) error
}

// UnimplementedPresenterQuizServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPresenterQuizServiceServer struct {
}

func (*UnimplementedPresenterQuizServiceServer) DelegateQuizQuestion(ctx context.Context, req *RestQuizQuestionUuid) (*RestStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelegateQuizQuestion not implemented")
}
func (*UnimplementedPresenterQuizServiceServer) GetQuizQuestionStatistics(req *RestQuizQuestionUuid, srv PresenterQuizService_GetQuizQuestionStatisticsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetQuizQuestionStatistics not implemented")
}
func (*UnimplementedPresenterQuizServiceServer) GetOpenQuizQuestionAnswers(req *RestQuizQuestionUuid, srv PresenterQuizService_GetOpenQuizQuestionAnswersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetOpenQuizQuestionAnswers not implemented")
}

func RegisterPresenterQuizServiceServer(s *grpc.Server, srv PresenterQuizServiceServer) {
	s.RegisterService(&_PresenterQuizService_serviceDesc, srv)
}

func _PresenterQuizService_DelegateQuizQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestQuizQuestionUuid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenterQuizServiceServer).DelegateQuizQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PresenterQuizService/DelegateQuizQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenterQuizServiceServer).DelegateQuizQuestion(ctx, req.(*RestQuizQuestionUuid))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresenterQuizService_GetQuizQuestionStatistics_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RestQuizQuestionUuid)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PresenterQuizServiceServer).GetQuizQuestionStatistics(m, &presenterQuizServiceGetQuizQuestionStatisticsServer{stream})
}

type PresenterQuizService_GetQuizQuestionStatisticsServer interface {
	Send(*QuizQuestionStatistics) error
	grpc.ServerStream
}

type presenterQuizServiceGetQuizQuestionStatisticsServer struct {
	grpc.ServerStream
}

func (x *presenterQuizServiceGetQuizQuestionStatisticsServer) Send(m *QuizQuestionStatistics) error {
	return x.ServerStream.SendMsg(m)
}

func _PresenterQuizService_GetOpenQuizQuestionAnswers_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RestQuizQuestionUuid)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PresenterQuizServiceServer).GetOpenQuizQuestionAnswers(m, &presenterQuizServiceGetOpenQuizQuestionAnswersServer{stream})
}

type PresenterQuizService_GetOpenQuizQuestionAnswersServer interface {
	Send(*QuizQuestionStatistics) error
	grpc.ServerStream
}

type presenterQuizServiceGetOpenQuizQuestionAnswersServer struct {
	grpc.ServerStream
}

func (x *presenterQuizServiceGetOpenQuizQuestionAnswersServer) Send(m *QuizQuestionStatistics) error {
	return x.ServerStream.SendMsg(m)
}

var _PresenterQuizService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PresenterQuizService",
	HandlerType: (*PresenterQuizServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DelegateQuizQuestion",
			Handler:    _PresenterQuizService_DelegateQuizQuestion_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetQuizQuestionStatistics",
			Handler:       _PresenterQuizService_GetQuizQuestionStatistics_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetOpenQuizQuestionAnswers",
			Handler:       _PresenterQuizService_GetOpenQuizQuestionAnswers_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "quiz.proto",
}
