// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/grpc/feedback.proto

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
	return fileDescriptor_18aaec75d046cef9, []int{0}
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

func init() { proto.RegisterFile("protos/grpc/feedback.proto", fileDescriptor_18aaec75d046cef9) }

var fileDescriptor_18aaec75d046cef9 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x4f, 0x4b, 0x4d, 0x4d, 0x49, 0x4a, 0x4c, 0xce,
	0xd6, 0x03, 0x0b, 0x4a, 0x49, 0x20, 0xcb, 0x15, 0x97, 0x24, 0x96, 0x94, 0x16, 0x43, 0x64, 0x94,
	0x7c, 0xb8, 0x38, 0xdc, 0xa0, 0x6a, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0xf2, 0x4b, 0x52, 0x25, 0x18,
	0x15, 0x18, 0x35, 0x18, 0x83, 0xc0, 0x6c, 0xb0, 0x58, 0x62, 0x6e, 0xaa, 0x04, 0x93, 0x02, 0xa3,
	0x06, 0x67, 0x10, 0x98, 0x2d, 0x24, 0xc1, 0xc5, 0x9e, 0x9c, 0x9f, 0x9b, 0x9b, 0x9a, 0x57, 0x22,
	0xc1, 0x0c, 0x16, 0x86, 0x71, 0x8d, 0x4c, 0xb9, 0xf8, 0x61, 0xa6, 0x05, 0xa7, 0x16, 0x95, 0x65,
	0x26, 0xa7, 0x0a, 0x29, 0x71, 0xf1, 0x04, 0xa7, 0xe6, 0xa5, 0xc0, 0x2d, 0xe1, 0xd4, 0x83, 0x31,
	0xa5, 0xd8, 0xf5, 0x82, 0xc1, 0x4e, 0x71, 0xe2, 0x8a, 0xe2, 0x48, 0x4d, 0x29, 0x4d, 0xce, 0xc9,
	0x2f, 0x4d, 0x49, 0x62, 0x03, 0xbb, 0xcb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x45, 0x08, 0xf5,
	0xd1, 0xcf, 0x00, 0x00, 0x00,
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
	Metadata: "protos/grpc/feedback.proto",
}