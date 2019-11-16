// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/grpc/homework.proto

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

type Homework struct {
	HomeworkUuid         int32    `protobuf:"varint,1,opt,name=homeworkUuid,proto3" json:"homeworkUuid,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	MailTo               string   `protobuf:"bytes,3,opt,name=mailTo,proto3" json:"mailTo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Homework) Reset()         { *m = Homework{} }
func (m *Homework) String() string { return proto.CompactTextString(m) }
func (*Homework) ProtoMessage()    {}
func (*Homework) Descriptor() ([]byte, []int) {
	return fileDescriptor_04748ef99c8ecbba, []int{0}
}

func (m *Homework) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Homework.Unmarshal(m, b)
}
func (m *Homework) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Homework.Marshal(b, m, deterministic)
}
func (m *Homework) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Homework.Merge(m, src)
}
func (m *Homework) XXX_Size() int {
	return xxx_messageInfo_Homework.Size(m)
}
func (m *Homework) XXX_DiscardUnknown() {
	xxx_messageInfo_Homework.DiscardUnknown(m)
}

var xxx_messageInfo_Homework proto.InternalMessageInfo

func (m *Homework) GetHomeworkUuid() int32 {
	if m != nil {
		return m.HomeworkUuid
	}
	return 0
}

func (m *Homework) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Homework) GetMailTo() string {
	if m != nil {
		return m.MailTo
	}
	return ""
}

type Homeworks struct {
	Homework             []*Homework `protobuf:"bytes,1,rep,name=homework,proto3" json:"homework,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Homeworks) Reset()         { *m = Homeworks{} }
func (m *Homeworks) String() string { return proto.CompactTextString(m) }
func (*Homeworks) ProtoMessage()    {}
func (*Homeworks) Descriptor() ([]byte, []int) {
	return fileDescriptor_04748ef99c8ecbba, []int{1}
}

func (m *Homeworks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Homeworks.Unmarshal(m, b)
}
func (m *Homeworks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Homeworks.Marshal(b, m, deterministic)
}
func (m *Homeworks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Homeworks.Merge(m, src)
}
func (m *Homeworks) XXX_Size() int {
	return xxx_messageInfo_Homeworks.Size(m)
}
func (m *Homeworks) XXX_DiscardUnknown() {
	xxx_messageInfo_Homeworks.DiscardUnknown(m)
}

var xxx_messageInfo_Homeworks proto.InternalMessageInfo

func (m *Homeworks) GetHomework() []*Homework {
	if m != nil {
		return m.Homework
	}
	return nil
}

func init() {
	proto.RegisterType((*Homework)(nil), "Homework")
	proto.RegisterType((*Homeworks)(nil), "Homeworks")
}

func init() { proto.RegisterFile("protos/grpc/homework.proto", fileDescriptor_04748ef99c8ecbba) }

var fileDescriptor_04748ef99c8ecbba = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcd, 0x4a, 0xc5, 0x30,
	0x10, 0x85, 0x89, 0xc5, 0x6b, 0x3b, 0xde, 0x85, 0x44, 0x91, 0x90, 0x55, 0xa9, 0x08, 0xc5, 0x45,
	0x2f, 0xd4, 0x37, 0x70, 0xa3, 0xeb, 0xd6, 0x82, 0xb8, 0xab, 0x4d, 0xd4, 0x60, 0xda, 0x29, 0xf9,
	0xd1, 0xd7, 0x17, 0x63, 0xd3, 0xea, 0x2e, 0xf9, 0xbe, 0x93, 0x13, 0x66, 0x80, 0xcf, 0x06, 0x1d,
	0xda, 0xc3, 0x9b, 0x99, 0x87, 0xc3, 0x3b, 0x8e, 0xf2, 0x0b, 0xcd, 0x47, 0x15, 0x20, 0x67, 0x7f,
	0x9d, 0x75, 0xbd, 0xf3, 0xf6, 0xd7, 0x14, 0x4f, 0x90, 0x3e, 0x2c, 0x59, 0x5a, 0xc0, 0x3e, 0xbe,
	0xeb, 0xbc, 0x12, 0x8c, 0xe4, 0xa4, 0x3c, 0x6e, 0xfe, 0x31, 0x7a, 0x06, 0x89, 0x37, 0x9a, 0x1d,
	0xe5, 0xa4, 0xcc, 0x9a, 0x9f, 0x23, 0xbd, 0x84, 0xdd, 0xd8, 0x2b, 0xfd, 0x88, 0x2c, 0x09, 0x70,
	0xb9, 0x15, 0x35, 0x64, 0xb1, 0xd9, 0xd2, 0x6b, 0x48, 0x63, 0x0d, 0x23, 0x79, 0x52, 0x9e, 0xd6,
	0x59, 0x15, 0x6d, 0xb3, 0xaa, 0xfa, 0x15, 0xce, 0x3b, 0x2b, 0x4d, 0x34, 0xad, 0x34, 0x9f, 0x6a,
	0x90, 0xf4, 0x06, 0x2e, 0x5a, 0x39, 0x89, 0x15, 0xa3, 0xf6, 0x4e, 0xe1, 0x44, 0xb7, 0x0e, 0x7e,
	0x52, 0xb5, 0x61, 0x2c, 0x7a, 0x05, 0xfb, 0x7b, 0xe9, 0xb6, 0x9f, 0xa3, 0xe0, 0xb0, 0x86, 0xed,
	0x1d, 0x3c, 0xa7, 0x52, 0xf8, 0x41, 0xa3, 0x17, 0x2f, 0xbb, 0xb0, 0x88, 0xdb, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x52, 0xa8, 0x2c, 0x4f, 0x40, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserHomeworkServiceClient is the client API for UserHomeworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserHomeworkServiceClient interface {
	SendHomeworkSolution(ctx context.Context, in *Homework, opts ...grpc.CallOption) (*Status, error)
	GetHomeworks(ctx context.Context, in *Status, opts ...grpc.CallOption) (*Homeworks, error)
}

type userHomeworkServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserHomeworkServiceClient(cc *grpc.ClientConn) UserHomeworkServiceClient {
	return &userHomeworkServiceClient{cc}
}

func (c *userHomeworkServiceClient) SendHomeworkSolution(ctx context.Context, in *Homework, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/UserHomeworkService/SendHomeworkSolution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userHomeworkServiceClient) GetHomeworks(ctx context.Context, in *Status, opts ...grpc.CallOption) (*Homeworks, error) {
	out := new(Homeworks)
	err := c.cc.Invoke(ctx, "/UserHomeworkService/GetHomeworks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserHomeworkServiceServer is the server API for UserHomeworkService service.
type UserHomeworkServiceServer interface {
	SendHomeworkSolution(context.Context, *Homework) (*Status, error)
	GetHomeworks(context.Context, *Status) (*Homeworks, error)
}

// UnimplementedUserHomeworkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserHomeworkServiceServer struct {
}

func (*UnimplementedUserHomeworkServiceServer) SendHomeworkSolution(ctx context.Context, req *Homework) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendHomeworkSolution not implemented")
}
func (*UnimplementedUserHomeworkServiceServer) GetHomeworks(ctx context.Context, req *Status) (*Homeworks, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHomeworks not implemented")
}

func RegisterUserHomeworkServiceServer(s *grpc.Server, srv UserHomeworkServiceServer) {
	s.RegisterService(&_UserHomeworkService_serviceDesc, srv)
}

func _UserHomeworkService_SendHomeworkSolution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Homework)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserHomeworkServiceServer).SendHomeworkSolution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserHomeworkService/SendHomeworkSolution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserHomeworkServiceServer).SendHomeworkSolution(ctx, req.(*Homework))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserHomeworkService_GetHomeworks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Status)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserHomeworkServiceServer).GetHomeworks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserHomeworkService/GetHomeworks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserHomeworkServiceServer).GetHomeworks(ctx, req.(*Status))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserHomeworkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UserHomeworkService",
	HandlerType: (*UserHomeworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendHomeworkSolution",
			Handler:    _UserHomeworkService_SendHomeworkSolution_Handler,
		},
		{
			MethodName: "GetHomeworks",
			Handler:    _UserHomeworkService_GetHomeworks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/grpc/homework.proto",
}