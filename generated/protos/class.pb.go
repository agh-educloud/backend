// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/class.proto

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

type Class struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"1,omitempty"`
	Topic                string          `protobuf:"bytes,2,opt,name=topic,proto3" json:"2,omitempty"`
	QuizQuestion         []*QuizQuestion `protobuf:"bytes,3,rep,name=quizQuestion,proto3" json:"3,omitempty"`
	Homework             []*Homework     `protobuf:"bytes,4,rep,name=homework,proto3" json:"4,omitempty"`
	Presentation         []byte          `protobuf:"bytes,5,opt,name=presentation,proto3" json:"5,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Class) Reset()         { *m = Class{} }
func (m *Class) String() string { return proto.CompactTextString(m) }
func (*Class) ProtoMessage()    {}
func (*Class) Descriptor() ([]byte, []int) {
	return fileDescriptor_f129ab25ee84b681, []int{0}
}

func (m *Class) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Class.Unmarshal(m, b)
}
func (m *Class) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Class.Marshal(b, m, deterministic)
}
func (m *Class) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Class.Merge(m, src)
}
func (m *Class) XXX_Size() int {
	return xxx_messageInfo_Class.Size(m)
}
func (m *Class) XXX_DiscardUnknown() {
	xxx_messageInfo_Class.DiscardUnknown(m)
}

var xxx_messageInfo_Class proto.InternalMessageInfo

func (m *Class) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Class) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Class) GetQuizQuestion() []*QuizQuestion {
	if m != nil {
		return m.QuizQuestion
	}
	return nil
}

func (m *Class) GetHomework() []*Homework {
	if m != nil {
		return m.Homework
	}
	return nil
}

func (m *Class) GetPresentation() []byte {
	if m != nil {
		return m.Presentation
	}
	return nil
}

type ClassCreationRequest struct {
	Class                *Class   `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClassCreationRequest) Reset()         { *m = ClassCreationRequest{} }
func (m *ClassCreationRequest) String() string { return proto.CompactTextString(m) }
func (*ClassCreationRequest) ProtoMessage()    {}
func (*ClassCreationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f129ab25ee84b681, []int{1}
}

func (m *ClassCreationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassCreationRequest.Unmarshal(m, b)
}
func (m *ClassCreationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassCreationRequest.Marshal(b, m, deterministic)
}
func (m *ClassCreationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassCreationRequest.Merge(m, src)
}
func (m *ClassCreationRequest) XXX_Size() int {
	return xxx_messageInfo_ClassCreationRequest.Size(m)
}
func (m *ClassCreationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassCreationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClassCreationRequest proto.InternalMessageInfo

func (m *ClassCreationRequest) GetClass() *Class {
	if m != nil {
		return m.Class
	}
	return nil
}

type ClassCreationResponse struct {
	ClassUuid            int32    `protobuf:"varint,1,opt,name=classUuid,proto3" json:"1,omitempty"`
	SecretCode           int32    `protobuf:"varint,2,opt,name=secretCode,proto3" json:"2,omitempty"`
	Error                *Status  `protobuf:"bytes,3,opt,name=error,proto3" json:"3,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClassCreationResponse) Reset()         { *m = ClassCreationResponse{} }
func (m *ClassCreationResponse) String() string { return proto.CompactTextString(m) }
func (*ClassCreationResponse) ProtoMessage()    {}
func (*ClassCreationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f129ab25ee84b681, []int{2}
}

func (m *ClassCreationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassCreationResponse.Unmarshal(m, b)
}
func (m *ClassCreationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassCreationResponse.Marshal(b, m, deterministic)
}
func (m *ClassCreationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassCreationResponse.Merge(m, src)
}
func (m *ClassCreationResponse) XXX_Size() int {
	return xxx_messageInfo_ClassCreationResponse.Size(m)
}
func (m *ClassCreationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassCreationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClassCreationResponse proto.InternalMessageInfo

func (m *ClassCreationResponse) GetClassUuid() int32 {
	if m != nil {
		return m.ClassUuid
	}
	return 0
}

func (m *ClassCreationResponse) GetSecretCode() int32 {
	if m != nil {
		return m.SecretCode
	}
	return 0
}

func (m *ClassCreationResponse) GetError() *Status {
	if m != nil {
		return m.Error
	}
	return nil
}

type ClassUpdateRequest struct {
	ClassUuid            int32    `protobuf:"varint,1,opt,name=classUuid,proto3" json:"1,omitempty"`
	Class                *Class   `protobuf:"bytes,2,opt,name=class,proto3" json:"2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClassUpdateRequest) Reset()         { *m = ClassUpdateRequest{} }
func (m *ClassUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*ClassUpdateRequest) ProtoMessage()    {}
func (*ClassUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f129ab25ee84b681, []int{3}
}

func (m *ClassUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassUpdateRequest.Unmarshal(m, b)
}
func (m *ClassUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassUpdateRequest.Marshal(b, m, deterministic)
}
func (m *ClassUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassUpdateRequest.Merge(m, src)
}
func (m *ClassUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_ClassUpdateRequest.Size(m)
}
func (m *ClassUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClassUpdateRequest proto.InternalMessageInfo

func (m *ClassUpdateRequest) GetClassUuid() int32 {
	if m != nil {
		return m.ClassUuid
	}
	return 0
}

func (m *ClassUpdateRequest) GetClass() *Class {
	if m != nil {
		return m.Class
	}
	return nil
}

type ClassUuid struct {
	ClassUuid            int32    `protobuf:"varint,1,opt,name=classUuid,proto3" json:"1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClassUuid) Reset()         { *m = ClassUuid{} }
func (m *ClassUuid) String() string { return proto.CompactTextString(m) }
func (*ClassUuid) ProtoMessage()    {}
func (*ClassUuid) Descriptor() ([]byte, []int) {
	return fileDescriptor_f129ab25ee84b681, []int{4}
}

func (m *ClassUuid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassUuid.Unmarshal(m, b)
}
func (m *ClassUuid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassUuid.Marshal(b, m, deterministic)
}
func (m *ClassUuid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassUuid.Merge(m, src)
}
func (m *ClassUuid) XXX_Size() int {
	return xxx_messageInfo_ClassUuid.Size(m)
}
func (m *ClassUuid) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassUuid.DiscardUnknown(m)
}

var xxx_messageInfo_ClassUuid proto.InternalMessageInfo

func (m *ClassUuid) GetClassUuid() int32 {
	if m != nil {
		return m.ClassUuid
	}
	return 0
}

type GetClassesResponse struct {
	Classes              []*ClassWithUuid `protobuf:"bytes,1,rep,name=classes,proto3" json:"1,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetClassesResponse) Reset()         { *m = GetClassesResponse{} }
func (m *GetClassesResponse) String() string { return proto.CompactTextString(m) }
func (*GetClassesResponse) ProtoMessage()    {}
func (*GetClassesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f129ab25ee84b681, []int{5}
}

func (m *GetClassesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClassesResponse.Unmarshal(m, b)
}
func (m *GetClassesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClassesResponse.Marshal(b, m, deterministic)
}
func (m *GetClassesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClassesResponse.Merge(m, src)
}
func (m *GetClassesResponse) XXX_Size() int {
	return xxx_messageInfo_GetClassesResponse.Size(m)
}
func (m *GetClassesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClassesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetClassesResponse proto.InternalMessageInfo

func (m *GetClassesResponse) GetClasses() []*ClassWithUuid {
	if m != nil {
		return m.Classes
	}
	return nil
}

type ClassWithUuid struct {
	ClassUuid            int32    `protobuf:"varint,1,opt,name=classUuid,proto3" json:"1,omitempty"`
	Class                *Class   `protobuf:"bytes,2,opt,name=class,proto3" json:"2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClassWithUuid) Reset()         { *m = ClassWithUuid{} }
func (m *ClassWithUuid) String() string { return proto.CompactTextString(m) }
func (*ClassWithUuid) ProtoMessage()    {}
func (*ClassWithUuid) Descriptor() ([]byte, []int) {
	return fileDescriptor_f129ab25ee84b681, []int{6}
}

func (m *ClassWithUuid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassWithUuid.Unmarshal(m, b)
}
func (m *ClassWithUuid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassWithUuid.Marshal(b, m, deterministic)
}
func (m *ClassWithUuid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassWithUuid.Merge(m, src)
}
func (m *ClassWithUuid) XXX_Size() int {
	return xxx_messageInfo_ClassWithUuid.Size(m)
}
func (m *ClassWithUuid) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassWithUuid.DiscardUnknown(m)
}

var xxx_messageInfo_ClassWithUuid proto.InternalMessageInfo

func (m *ClassWithUuid) GetClassUuid() int32 {
	if m != nil {
		return m.ClassUuid
	}
	return 0
}

func (m *ClassWithUuid) GetClass() *Class {
	if m != nil {
		return m.Class
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_f129ab25ee84b681, []int{7}
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

type JoinClassRequest struct {
	SecretCode           string   `protobuf:"bytes,1,opt,name=secretCode,proto3" json:"1,omitempty"`
	User                 *User    `protobuf:"bytes,2,opt,name=user,proto3" json:"2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinClassRequest) Reset()         { *m = JoinClassRequest{} }
func (m *JoinClassRequest) String() string { return proto.CompactTextString(m) }
func (*JoinClassRequest) ProtoMessage()    {}
func (*JoinClassRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f129ab25ee84b681, []int{8}
}

func (m *JoinClassRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinClassRequest.Unmarshal(m, b)
}
func (m *JoinClassRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinClassRequest.Marshal(b, m, deterministic)
}
func (m *JoinClassRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinClassRequest.Merge(m, src)
}
func (m *JoinClassRequest) XXX_Size() int {
	return xxx_messageInfo_JoinClassRequest.Size(m)
}
func (m *JoinClassRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinClassRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinClassRequest proto.InternalMessageInfo

func (m *JoinClassRequest) GetSecretCode() string {
	if m != nil {
		return m.SecretCode
	}
	return ""
}

func (m *JoinClassRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type Assigment struct {
	// Types that are valid to be assigned to Event:
	//	*Assigment_QuizQuestion
	//	*Assigment_Homework
	Event                isAssigment_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Assigment) Reset()         { *m = Assigment{} }
func (m *Assigment) String() string { return proto.CompactTextString(m) }
func (*Assigment) ProtoMessage()    {}
func (*Assigment) Descriptor() ([]byte, []int) {
	return fileDescriptor_f129ab25ee84b681, []int{9}
}

func (m *Assigment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Assigment.Unmarshal(m, b)
}
func (m *Assigment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Assigment.Marshal(b, m, deterministic)
}
func (m *Assigment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Assigment.Merge(m, src)
}
func (m *Assigment) XXX_Size() int {
	return xxx_messageInfo_Assigment.Size(m)
}
func (m *Assigment) XXX_DiscardUnknown() {
	xxx_messageInfo_Assigment.DiscardUnknown(m)
}

var xxx_messageInfo_Assigment proto.InternalMessageInfo

type isAssigment_Event interface {
	isAssigment_Event()
}

type Assigment_QuizQuestion struct {
	QuizQuestion *QuizQuestion `protobuf:"bytes,1,opt,name=quizQuestion,proto3,oneof"`
}

type Assigment_Homework struct {
	Homework *Homework `protobuf:"bytes,2,opt,name=homework,proto3,oneof"`
}

func (*Assigment_QuizQuestion) isAssigment_Event() {}

func (*Assigment_Homework) isAssigment_Event() {}

func (m *Assigment) GetEvent() isAssigment_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Assigment) GetQuizQuestion() *QuizQuestion {
	if x, ok := m.GetEvent().(*Assigment_QuizQuestion); ok {
		return x.QuizQuestion
	}
	return nil
}

func (m *Assigment) GetHomework() *Homework {
	if x, ok := m.GetEvent().(*Assigment_Homework); ok {
		return x.Homework
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Assigment) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Assigment_QuizQuestion)(nil),
		(*Assigment_Homework)(nil),
	}
}

type JoinAssigmentLoopRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinAssigmentLoopRequest) Reset()         { *m = JoinAssigmentLoopRequest{} }
func (m *JoinAssigmentLoopRequest) String() string { return proto.CompactTextString(m) }
func (*JoinAssigmentLoopRequest) ProtoMessage()    {}
func (*JoinAssigmentLoopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f129ab25ee84b681, []int{10}
}

func (m *JoinAssigmentLoopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinAssigmentLoopRequest.Unmarshal(m, b)
}
func (m *JoinAssigmentLoopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinAssigmentLoopRequest.Marshal(b, m, deterministic)
}
func (m *JoinAssigmentLoopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinAssigmentLoopRequest.Merge(m, src)
}
func (m *JoinAssigmentLoopRequest) XXX_Size() int {
	return xxx_messageInfo_JoinAssigmentLoopRequest.Size(m)
}
func (m *JoinAssigmentLoopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinAssigmentLoopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinAssigmentLoopRequest proto.InternalMessageInfo

func (m *JoinAssigmentLoopRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*Class)(nil), "Class")
	proto.RegisterType((*ClassCreationRequest)(nil), "ClassCreationRequest")
	proto.RegisterType((*ClassCreationResponse)(nil), "ClassCreationResponse")
	proto.RegisterType((*ClassUpdateRequest)(nil), "ClassUpdateRequest")
	proto.RegisterType((*ClassUuid)(nil), "ClassUuid")
	proto.RegisterType((*GetClassesResponse)(nil), "GetClassesResponse")
	proto.RegisterType((*ClassWithUuid)(nil), "ClassWithUuid")
	proto.RegisterType((*Empty)(nil), "Empty")
	proto.RegisterType((*JoinClassRequest)(nil), "JoinClassRequest")
	proto.RegisterType((*Assigment)(nil), "Assigment")
	proto.RegisterType((*JoinAssigmentLoopRequest)(nil), "JoinAssigmentLoopRequest")
}

func init() { proto.RegisterFile("protos/class.proto", fileDescriptor_f129ab25ee84b681) }

var fileDescriptor_f129ab25ee84b681 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdf, 0x6e, 0xd3, 0x4c,
	0x10, 0xc5, 0xeb, 0xb6, 0x4e, 0xe2, 0x49, 0xfa, 0xa9, 0xdd, 0x34, 0x9f, 0x4c, 0x54, 0x4a, 0x64,
	0x09, 0x35, 0x15, 0xd2, 0x02, 0x29, 0x5c, 0x81, 0x90, 0x68, 0xa8, 0x88, 0x80, 0x4a, 0xad, 0xa3,
	0x08, 0x89, 0x3b, 0xd7, 0x1e, 0x1a, 0x8b, 0xc4, 0xeb, 0xee, 0xae, 0x8b, 0xe0, 0xa5, 0x78, 0x01,
	0x1e, 0x0e, 0x79, 0xd7, 0x7f, 0x62, 0x37, 0xf4, 0x86, 0xbb, 0xfa, 0xec, 0x99, 0x9d, 0x33, 0xb3,
	0xbf, 0x06, 0x48, 0xcc, 0x99, 0x64, 0xe2, 0xa9, 0xbf, 0xf0, 0x84, 0xa0, 0xea, 0xa3, 0xdf, 0xcd,
	0x34, 0x21, 0x3d, 0x99, 0xe4, 0xe2, 0x5e, 0x26, 0xde, 0x24, 0xe1, 0xcf, 0x4c, 0xea, 0x65, 0xd2,
	0x9c, 0x2d, 0xf1, 0x3b, 0xe3, 0xdf, 0x6a, 0x4e, 0x7f, 0xee, 0xc9, 0x9a, 0x94, 0x08, 0xe4, 0xb5,
	0xe2, 0xaf, 0x88, 0xc1, 0x95, 0xe7, 0x67, 0xc5, 0xce, 0x2f, 0x03, 0xcc, 0x71, 0x9a, 0x85, 0x10,
	0xd8, 0x8e, 0xbc, 0x25, 0xda, 0xc6, 0xc0, 0x18, 0x5a, 0xae, 0xfa, 0x9b, 0xec, 0x83, 0x29, 0x59,
	0x1c, 0xfa, 0xf6, 0xa6, 0x12, 0xf5, 0x07, 0x79, 0x0e, 0x9d, 0x34, 0xd5, 0x65, 0x82, 0x42, 0x86,
	0x2c, 0xb2, 0xb7, 0x06, 0x5b, 0xc3, 0xf6, 0x68, 0x87, 0x5e, 0xae, 0x88, 0x6e, 0xc5, 0x42, 0x1e,
	0x43, 0x2b, 0x4f, 0x6d, 0x6f, 0x2b, 0xbb, 0x45, 0x27, 0x99, 0xe0, 0x16, 0x47, 0xc4, 0x81, 0x4e,
	0xcc, 0x51, 0x60, 0x24, 0x3d, 0x75, 0xb3, 0x39, 0x30, 0x86, 0x1d, 0xb7, 0xa2, 0x39, 0x2f, 0x60,
	0x5f, 0x05, 0x1e, 0x73, 0x54, 0x82, 0x8b, 0x37, 0x69, 0x17, 0x72, 0x00, 0xa6, 0x5a, 0xaa, 0x1a,
	0xa0, 0x3d, 0x6a, 0x50, 0xe5, 0x72, 0xb5, 0xe8, 0x48, 0xe8, 0xd5, 0xaa, 0x44, 0xcc, 0x22, 0x81,
	0xe4, 0x00, 0x2c, 0xe5, 0x98, 0x25, 0x61, 0xa0, 0x4a, 0x4d, 0xb7, 0x14, 0xc8, 0x21, 0x80, 0x40,
	0x9f, 0xa3, 0x1c, 0xb3, 0x00, 0xd5, 0x16, 0x4c, 0x77, 0x45, 0x21, 0x0f, 0xc1, 0x44, 0xce, 0x19,
	0xb7, 0xb7, 0x54, 0xd3, 0x26, 0x9d, 0xaa, 0x37, 0x74, 0xb5, 0xea, 0x5c, 0x00, 0x51, 0x5d, 0x67,
	0x71, 0xe0, 0x49, 0x2c, 0x93, 0xde, 0xd7, 0xb2, 0x98, 0x63, 0x73, 0xdd, 0x1c, 0xc7, 0x60, 0x8d,
	0x57, 0xac, 0xf7, 0x5c, 0xe4, 0xbc, 0x01, 0xf2, 0x1e, 0xa5, 0x72, 0xa3, 0x28, 0xe6, 0x1d, 0x42,
	0xd3, 0xd7, 0x92, 0x6d, 0xa8, 0x87, 0xf8, 0x4f, 0x37, 0xf8, 0x1c, 0xca, 0x79, 0x5a, 0xe6, 0xe6,
	0xc7, 0xce, 0x47, 0xd8, 0xa9, 0x9c, 0xfc, 0x53, 0xee, 0x26, 0x98, 0x67, 0xcb, 0x58, 0xfe, 0x70,
	0xce, 0x61, 0xf7, 0x03, 0x0b, 0x23, 0x7d, 0x98, 0x2d, 0xa4, 0xba, 0x65, 0x0d, 0xe0, 0xea, 0x96,
	0x1f, 0xc0, 0x76, 0x4a, 0x72, 0x76, 0xb3, 0x49, 0x67, 0x02, 0xb9, 0xab, 0x24, 0x87, 0x83, 0xf5,
	0x56, 0x88, 0xf0, 0x7a, 0x89, 0x91, 0x24, 0x27, 0x35, 0x30, 0x35, 0x09, 0x55, 0x30, 0x27, 0x1b,
	0x35, 0x34, 0x8f, 0x56, 0xd0, 0xd4, 0x0d, 0x4a, 0x34, 0x27, 0x1b, 0x25, 0x9c, 0xa7, 0x4d, 0x30,
	0xf1, 0x16, 0x23, 0xe9, 0xbc, 0x04, 0x3b, 0x1d, 0xa1, 0xe8, 0xfb, 0x89, 0xb1, 0x38, 0x1f, 0x25,
	0x8f, 0x6a, 0xdc, 0x89, 0x3a, 0xfa, 0xbd, 0x09, 0xbd, 0x0b, 0x4d, 0x32, 0x72, 0x35, 0xff, 0x14,
	0xf9, 0x6d, 0xe8, 0x23, 0x79, 0x0d, 0x6d, 0xc5, 0x25, 0xea, 0xff, 0xc4, 0x1e, 0x5d, 0x07, 0x78,
	0xff, 0x7f, 0xba, 0x9e, 0xe0, 0x27, 0xd0, 0xd6, 0x7c, 0xe9, 0xea, 0x2e, 0xbd, 0x8b, 0x5c, 0x3f,
	0x07, 0x93, 0x38, 0xd0, 0x7e, 0x87, 0x0b, 0xcc, 0xcd, 0x40, 0x0b, 0x9a, 0x4a, 0xcf, 0x21, 0xb4,
	0x72, 0x70, 0x2a, 0x86, 0xec, 0x49, 0xc9, 0x31, 0x40, 0x09, 0x16, 0x69, 0x50, 0xf5, 0xb0, 0xfd,
	0x2e, 0x5d, 0x4b, 0x1b, 0x4c, 0xa5, 0xc7, 0xd7, 0x5c, 0xd6, 0xa1, 0xe3, 0xb9, 0x27, 0xcf, 0x51,
	0x08, 0xef, 0x1a, 0x9f, 0x19, 0xe4, 0x11, 0xb4, 0xce, 0xa2, 0xe0, 0xef, 0xa9, 0x46, 0xaf, 0x60,
	0x37, 0x5d, 0x66, 0x65, 0x71, 0x47, 0x60, 0x15, 0x30, 0x91, 0x3d, 0x5a, 0x07, 0xab, 0x28, 0x3e,
	0x85, 0x2f, 0x2d, 0x0c, 0x12, 0x7f, 0xc1, 0x92, 0xe0, 0xaa, 0xa1, 0x7e, 0xf9, 0x4e, 0xfe, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x74, 0x20, 0x6f, 0x3d, 0x8b, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PresenterClassServiceClient is the client API for PresenterClassService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PresenterClassServiceClient interface {
	CreateClass(ctx context.Context, in *ClassCreationRequest, opts ...grpc.CallOption) (*ClassCreationResponse, error)
	UpdateClass(ctx context.Context, in *ClassUpdateRequest, opts ...grpc.CallOption) (*Status, error)
	DeleteClass(ctx context.Context, in *ClassUuid, opts ...grpc.CallOption) (*Status, error)
	GetClass(ctx context.Context, in *ClassUuid, opts ...grpc.CallOption) (*Class, error)
	GetClasses(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetClassesResponse, error)
	StartClass(ctx context.Context, in *ClassUuid, opts ...grpc.CallOption) (PresenterClassService_StartClassClient, error)
	EndClass(ctx context.Context, in *ClassUuid, opts ...grpc.CallOption) (*Status, error)
}

type presenterClassServiceClient struct {
	cc *grpc.ClientConn
}

func NewPresenterClassServiceClient(cc *grpc.ClientConn) PresenterClassServiceClient {
	return &presenterClassServiceClient{cc}
}

func (c *presenterClassServiceClient) CreateClass(ctx context.Context, in *ClassCreationRequest, opts ...grpc.CallOption) (*ClassCreationResponse, error) {
	out := new(ClassCreationResponse)
	err := c.cc.Invoke(ctx, "/PresenterClassService/CreateClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presenterClassServiceClient) UpdateClass(ctx context.Context, in *ClassUpdateRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/PresenterClassService/UpdateClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presenterClassServiceClient) DeleteClass(ctx context.Context, in *ClassUuid, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/PresenterClassService/DeleteClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presenterClassServiceClient) GetClass(ctx context.Context, in *ClassUuid, opts ...grpc.CallOption) (*Class, error) {
	out := new(Class)
	err := c.cc.Invoke(ctx, "/PresenterClassService/GetClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presenterClassServiceClient) GetClasses(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetClassesResponse, error) {
	out := new(GetClassesResponse)
	err := c.cc.Invoke(ctx, "/PresenterClassService/GetClasses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *presenterClassServiceClient) StartClass(ctx context.Context, in *ClassUuid, opts ...grpc.CallOption) (PresenterClassService_StartClassClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PresenterClassService_serviceDesc.Streams[0], "/PresenterClassService/StartClass", opts...)
	if err != nil {
		return nil, err
	}
	x := &presenterClassServiceStartClassClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PresenterClassService_StartClassClient interface {
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type presenterClassServiceStartClassClient struct {
	grpc.ClientStream
}

func (x *presenterClassServiceStartClassClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *presenterClassServiceClient) EndClass(ctx context.Context, in *ClassUuid, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/PresenterClassService/EndClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PresenterClassServiceServer is the server API for PresenterClassService service.
type PresenterClassServiceServer interface {
	CreateClass(context.Context, *ClassCreationRequest) (*ClassCreationResponse, error)
	UpdateClass(context.Context, *ClassUpdateRequest) (*Status, error)
	DeleteClass(context.Context, *ClassUuid) (*Status, error)
	GetClass(context.Context, *ClassUuid) (*Class, error)
	GetClasses(context.Context, *Empty) (*GetClassesResponse, error)
	StartClass(*ClassUuid, PresenterClassService_StartClassServer) error
	EndClass(context.Context, *ClassUuid) (*Status, error)
}

// UnimplementedPresenterClassServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPresenterClassServiceServer struct {
}

func (*UnimplementedPresenterClassServiceServer) CreateClass(ctx context.Context, req *ClassCreationRequest) (*ClassCreationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClass not implemented")
}
func (*UnimplementedPresenterClassServiceServer) UpdateClass(ctx context.Context, req *ClassUpdateRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClass not implemented")
}
func (*UnimplementedPresenterClassServiceServer) DeleteClass(ctx context.Context, req *ClassUuid) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteClass not implemented")
}
func (*UnimplementedPresenterClassServiceServer) GetClass(ctx context.Context, req *ClassUuid) (*Class, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClass not implemented")
}
func (*UnimplementedPresenterClassServiceServer) GetClasses(ctx context.Context, req *Empty) (*GetClassesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClasses not implemented")
}
func (*UnimplementedPresenterClassServiceServer) StartClass(req *ClassUuid, srv PresenterClassService_StartClassServer) error {
	return status.Errorf(codes.Unimplemented, "method StartClass not implemented")
}
func (*UnimplementedPresenterClassServiceServer) EndClass(ctx context.Context, req *ClassUuid) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndClass not implemented")
}

func RegisterPresenterClassServiceServer(s *grpc.Server, srv PresenterClassServiceServer) {
	s.RegisterService(&_PresenterClassService_serviceDesc, srv)
}

func _PresenterClassService_CreateClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClassCreationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenterClassServiceServer).CreateClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PresenterClassService/CreateClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenterClassServiceServer).CreateClass(ctx, req.(*ClassCreationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresenterClassService_UpdateClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClassUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenterClassServiceServer).UpdateClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PresenterClassService/UpdateClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenterClassServiceServer).UpdateClass(ctx, req.(*ClassUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresenterClassService_DeleteClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClassUuid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenterClassServiceServer).DeleteClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PresenterClassService/DeleteClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenterClassServiceServer).DeleteClass(ctx, req.(*ClassUuid))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresenterClassService_GetClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClassUuid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenterClassServiceServer).GetClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PresenterClassService/GetClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenterClassServiceServer).GetClass(ctx, req.(*ClassUuid))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresenterClassService_GetClasses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenterClassServiceServer).GetClasses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PresenterClassService/GetClasses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenterClassServiceServer).GetClasses(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _PresenterClassService_StartClass_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ClassUuid)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PresenterClassServiceServer).StartClass(m, &presenterClassServiceStartClassServer{stream})
}

type PresenterClassService_StartClassServer interface {
	Send(*ChatMessage) error
	grpc.ServerStream
}

type presenterClassServiceStartClassServer struct {
	grpc.ServerStream
}

func (x *presenterClassServiceStartClassServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _PresenterClassService_EndClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClassUuid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PresenterClassServiceServer).EndClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PresenterClassService/EndClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PresenterClassServiceServer).EndClass(ctx, req.(*ClassUuid))
	}
	return interceptor(ctx, in, info, handler)
}

var _PresenterClassService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PresenterClassService",
	HandlerType: (*PresenterClassServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClass",
			Handler:    _PresenterClassService_CreateClass_Handler,
		},
		{
			MethodName: "UpdateClass",
			Handler:    _PresenterClassService_UpdateClass_Handler,
		},
		{
			MethodName: "DeleteClass",
			Handler:    _PresenterClassService_DeleteClass_Handler,
		},
		{
			MethodName: "GetClass",
			Handler:    _PresenterClassService_GetClass_Handler,
		},
		{
			MethodName: "GetClasses",
			Handler:    _PresenterClassService_GetClasses_Handler,
		},
		{
			MethodName: "EndClass",
			Handler:    _PresenterClassService_EndClass_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartClass",
			Handler:       _PresenterClassService_StartClass_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/class.proto",
}

// UserClassServiceClient is the client API for UserClassService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClassServiceClient interface {
	JoinClass(ctx context.Context, in *JoinClassRequest, opts ...grpc.CallOption) (*Status, error)
}

type userClassServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserClassServiceClient(cc *grpc.ClientConn) UserClassServiceClient {
	return &userClassServiceClient{cc}
}

func (c *userClassServiceClient) JoinClass(ctx context.Context, in *JoinClassRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/UserClassService/JoinClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserClassServiceServer is the server API for UserClassService service.
type UserClassServiceServer interface {
	JoinClass(context.Context, *JoinClassRequest) (*Status, error)
}

// UnimplementedUserClassServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserClassServiceServer struct {
}

func (*UnimplementedUserClassServiceServer) JoinClass(ctx context.Context, req *JoinClassRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinClass not implemented")
}

func RegisterUserClassServiceServer(s *grpc.Server, srv UserClassServiceServer) {
	s.RegisterService(&_UserClassService_serviceDesc, srv)
}

func _UserClassService_JoinClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserClassServiceServer).JoinClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserClassService/JoinClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserClassServiceServer).JoinClass(ctx, req.(*JoinClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserClassService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UserClassService",
	HandlerType: (*UserClassServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinClass",
			Handler:    _UserClassService_JoinClass_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/class.proto",
}
