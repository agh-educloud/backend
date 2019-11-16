// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/feedback.proto

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

type Feedback struct {
	Note                 float64  `protobuf:"fixed64,1,opt,name=note,proto3" json:"1,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"2,omitempty"`
	Comment              string   `protobuf:"bytes,3,opt,name=comment,proto3" json:"3,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Feedback) Reset()         { *m = Feedback{} }
func (m *Feedback) String() string { return proto.CompactTextString(m) }
func (*Feedback) ProtoMessage()    {}
func (*Feedback) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a46fba5d4f895ee, []int{0}
}

func (m *Feedback) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feedback.Unmarshal(m, b)
}
func (m *Feedback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feedback.Marshal(b, m, deterministic)
}
func (m *Feedback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feedback.Merge(m, src)
}
func (m *Feedback) XXX_Size() int {
	return xxx_messageInfo_Feedback.Size(m)
}
func (m *Feedback) XXX_DiscardUnknown() {
	xxx_messageInfo_Feedback.DiscardUnknown(m)
}

var xxx_messageInfo_Feedback proto.InternalMessageInfo

func (m *Feedback) GetNote() float64 {
	if m != nil {
		return m.Note
	}
	return 0
}

func (m *Feedback) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Feedback) GetComment() string {
	if m != nil {
		return m.Comment
	}
	return ""
}

func init() {
	proto.RegisterType((*Feedback)(nil), "Feedback")
}

func init() { proto.RegisterFile("protos/feedback.proto", fileDescriptor_8a46fba5d4f895ee) }

var fileDescriptor_8a46fba5d4f895ee = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0x4b, 0x4d, 0x4d, 0x49, 0x4a, 0x4c, 0xce, 0xd6, 0x03, 0xf3, 0xa5, 0x84,
	0xa1, 0xc2, 0xc5, 0x25, 0x89, 0x25, 0xa5, 0xc5, 0x10, 0x41, 0x25, 0x1f, 0x2e, 0x0e, 0x37, 0xa8,
	0x32, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xfc, 0x92, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xc6, 0x20,
	0x30, 0x1b, 0x2c, 0x96, 0x98, 0x9b, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x0b,
	0x49, 0x70, 0xb1, 0x27, 0xe7, 0xe7, 0xe6, 0xa6, 0xe6, 0x95, 0x48, 0x30, 0x83, 0x85, 0x61, 0x5c,
	0x23, 0x53, 0x2e, 0x7e, 0x98, 0x69, 0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0x4a, 0x5c,
	0x3c, 0xc1, 0xa9, 0x79, 0x29, 0x70, 0x4b, 0x38, 0xf5, 0x60, 0x4c, 0x29, 0x76, 0xbd, 0x60, 0xb0,
	0x53, 0x9c, 0xb8, 0xa2, 0x38, 0x52, 0x53, 0x4a, 0x93, 0x73, 0xf2, 0x4b, 0x53, 0x92, 0xd8, 0xc0,
	0xee, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xdb, 0x13, 0x89, 0x46, 0xc5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FeedbackServiceClient is the client API for FeedbackService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedbackServiceClient interface {
	SendFeedback(ctx context.Context, in *Feedback, opts ...grpc.CallOption) (*Status, error)
}

type feedbackServiceClient struct {
	cc *grpc.ClientConn
}

func NewFeedbackServiceClient(cc *grpc.ClientConn) FeedbackServiceClient {
	return &feedbackServiceClient{cc}
}

func (c *feedbackServiceClient) SendFeedback(ctx context.Context, in *Feedback, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/FeedbackService/SendFeedback", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedbackServiceServer is the server API for FeedbackService service.
type FeedbackServiceServer interface {
	SendFeedback(context.Context, *Feedback) (*Status, error)
}

// UnimplementedFeedbackServiceServer can be embedded to have forward compatible implementations.
type UnimplementedFeedbackServiceServer struct {
}

func (*UnimplementedFeedbackServiceServer) SendFeedback(ctx context.Context, req *Feedback) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendFeedback not implemented")
}

func RegisterFeedbackServiceServer(s *grpc.Server, srv FeedbackServiceServer) {
	s.RegisterService(&_FeedbackService_serviceDesc, srv)
}

func _FeedbackService_SendFeedback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Feedback)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedbackServiceServer).SendFeedback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/FeedbackService/SendFeedback",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedbackServiceServer).SendFeedback(ctx, req.(*Feedback))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedbackService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "FeedbackService",
	HandlerType: (*FeedbackServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendFeedback",
			Handler:    _FeedbackService_SendFeedback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/feedback.proto",
}