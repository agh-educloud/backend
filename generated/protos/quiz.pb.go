// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/quiz.proto

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

type QuizRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuizRequest) Reset()         { *m = QuizRequest{} }
func (m *QuizRequest) String() string { return proto.CompactTextString(m) }
func (*QuizRequest) ProtoMessage()    {}
func (*QuizRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_99abdaeba5813822, []int{0}
}

func (m *QuizRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuizRequest.Unmarshal(m, b)
}
func (m *QuizRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuizRequest.Marshal(b, m, deterministic)
}
func (m *QuizRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuizRequest.Merge(m, src)
}
func (m *QuizRequest) XXX_Size() int {
	return xxx_messageInfo_QuizRequest.Size(m)
}
func (m *QuizRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuizRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuizRequest proto.InternalMessageInfo

func (m *QuizRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type QuizAnswer struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Answer               string   `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	Photo                []byte   `protobuf:"bytes,3,opt,name=photo,proto3" json:"photo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuizAnswer) Reset()         { *m = QuizAnswer{} }
func (m *QuizAnswer) String() string { return proto.CompactTextString(m) }
func (*QuizAnswer) ProtoMessage()    {}
func (*QuizAnswer) Descriptor() ([]byte, []int) {
	return fileDescriptor_99abdaeba5813822, []int{1}
}

func (m *QuizAnswer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuizAnswer.Unmarshal(m, b)
}
func (m *QuizAnswer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuizAnswer.Marshal(b, m, deterministic)
}
func (m *QuizAnswer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuizAnswer.Merge(m, src)
}
func (m *QuizAnswer) XXX_Size() int {
	return xxx_messageInfo_QuizAnswer.Size(m)
}
func (m *QuizAnswer) XXX_DiscardUnknown() {
	xxx_messageInfo_QuizAnswer.DiscardUnknown(m)
}

var xxx_messageInfo_QuizAnswer proto.InternalMessageInfo

func (m *QuizAnswer) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *QuizAnswer) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

func (m *QuizAnswer) GetPhoto() []byte {
	if m != nil {
		return m.Photo
	}
	return nil
}

type Question struct {
	ClassId              string   `protobuf:"bytes,1,opt,name=classId,proto3" json:"classId,omitempty"`
	GroupId              string   `protobuf:"bytes,2,opt,name=groupId,proto3" json:"groupId,omitempty"`
	Answers              []string `protobuf:"bytes,3,rep,name=answers,proto3" json:"answers,omitempty"`
	PhotoQuestion        bool     `protobuf:"varint,4,opt,name=photoQuestion,proto3" json:"photoQuestion,omitempty"`
	ClosedQuestion       bool     `protobuf:"varint,5,opt,name=closedQuestion,proto3" json:"closedQuestion,omitempty"`
	OpenQuestion         bool     `protobuf:"varint,6,opt,name=openQuestion,proto3" json:"openQuestion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Question) Reset()         { *m = Question{} }
func (m *Question) String() string { return proto.CompactTextString(m) }
func (*Question) ProtoMessage()    {}
func (*Question) Descriptor() ([]byte, []int) {
	return fileDescriptor_99abdaeba5813822, []int{2}
}

func (m *Question) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Question.Unmarshal(m, b)
}
func (m *Question) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Question.Marshal(b, m, deterministic)
}
func (m *Question) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Question.Merge(m, src)
}
func (m *Question) XXX_Size() int {
	return xxx_messageInfo_Question.Size(m)
}
func (m *Question) XXX_DiscardUnknown() {
	xxx_messageInfo_Question.DiscardUnknown(m)
}

var xxx_messageInfo_Question proto.InternalMessageInfo

func (m *Question) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

func (m *Question) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *Question) GetAnswers() []string {
	if m != nil {
		return m.Answers
	}
	return nil
}

func (m *Question) GetPhotoQuestion() bool {
	if m != nil {
		return m.PhotoQuestion
	}
	return false
}

func (m *Question) GetClosedQuestion() bool {
	if m != nil {
		return m.ClosedQuestion
	}
	return false
}

func (m *Question) GetOpenQuestion() bool {
	if m != nil {
		return m.OpenQuestion
	}
	return false
}

type QuizQuestion struct {
	Uuid                 int32     `protobuf:"varint,1,opt,name=uuid,proto3" json:"1,omitempty"`
	Question             string    `protobuf:"bytes,2,opt,name=question,proto3" json:"2,omitempty"`
	Hint                 string    `protobuf:"bytes,3,opt,name=hint,proto3" json:"3,omitempty"`
	Option               []*Option `protobuf:"bytes,4,rep,name=option,proto3" json:"4,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *QuizQuestion) Reset()         { *m = QuizQuestion{} }
func (m *QuizQuestion) String() string { return proto.CompactTextString(m) }
func (*QuizQuestion) ProtoMessage()    {}
func (*QuizQuestion) Descriptor() ([]byte, []int) {
	return fileDescriptor_99abdaeba5813822, []int{3}
}

func (m *QuizQuestion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuizQuestion.Unmarshal(m, b)
}
func (m *QuizQuestion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuizQuestion.Marshal(b, m, deterministic)
}
func (m *QuizQuestion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuizQuestion.Merge(m, src)
}
func (m *QuizQuestion) XXX_Size() int {
	return xxx_messageInfo_QuizQuestion.Size(m)
}
func (m *QuizQuestion) XXX_DiscardUnknown() {
	xxx_messageInfo_QuizQuestion.DiscardUnknown(m)
}

var xxx_messageInfo_QuizQuestion proto.InternalMessageInfo

func (m *QuizQuestion) GetUuid() int32 {
	if m != nil {
		return m.Uuid
	}
	return 0
}

func (m *QuizQuestion) GetQuestion() string {
	if m != nil {
		return m.Question
	}
	return ""
}

func (m *QuizQuestion) GetHint() string {
	if m != nil {
		return m.Hint
	}
	return ""
}

func (m *QuizQuestion) GetOption() []*Option {
	if m != nil {
		return m.Option
	}
	return nil
}

type Option struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Option) Reset()         { *m = Option{} }
func (m *Option) String() string { return proto.CompactTextString(m) }
func (*Option) ProtoMessage()    {}
func (*Option) Descriptor() ([]byte, []int) {
	return fileDescriptor_99abdaeba5813822, []int{4}
}

func (m *Option) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Option.Unmarshal(m, b)
}
func (m *Option) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Option.Marshal(b, m, deterministic)
}
func (m *Option) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Option.Merge(m, src)
}
func (m *Option) XXX_Size() int {
	return xxx_messageInfo_Option.Size(m)
}
func (m *Option) XXX_DiscardUnknown() {
	xxx_messageInfo_Option.DiscardUnknown(m)
}

var xxx_messageInfo_Option proto.InternalMessageInfo

func (m *Option) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type QuizQuestionCreation struct {
	QuizQuestion         *QuizQuestion `protobuf:"bytes,1,opt,name=quizQuestion,proto3" json:"quizQuestion,omitempty"`
	Answer               *Option       `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *QuizQuestionCreation) Reset()         { *m = QuizQuestionCreation{} }
func (m *QuizQuestionCreation) String() string { return proto.CompactTextString(m) }
func (*QuizQuestionCreation) ProtoMessage()    {}
func (*QuizQuestionCreation) Descriptor() ([]byte, []int) {
	return fileDescriptor_99abdaeba5813822, []int{5}
}

func (m *QuizQuestionCreation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuizQuestionCreation.Unmarshal(m, b)
}
func (m *QuizQuestionCreation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuizQuestionCreation.Marshal(b, m, deterministic)
}
func (m *QuizQuestionCreation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuizQuestionCreation.Merge(m, src)
}
func (m *QuizQuestionCreation) XXX_Size() int {
	return xxx_messageInfo_QuizQuestionCreation.Size(m)
}
func (m *QuizQuestionCreation) XXX_DiscardUnknown() {
	xxx_messageInfo_QuizQuestionCreation.DiscardUnknown(m)
}

var xxx_messageInfo_QuizQuestionCreation proto.InternalMessageInfo

func (m *QuizQuestionCreation) GetQuizQuestion() *QuizQuestion {
	if m != nil {
		return m.QuizQuestion
	}
	return nil
}

func (m *QuizQuestionCreation) GetAnswer() *Option {
	if m != nil {
		return m.Answer
	}
	return nil
}

type QuizQuestionAnswer struct {
	Uuid                 int32    `protobuf:"varint,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Answer               *Option  `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuizQuestionAnswer) Reset()         { *m = QuizQuestionAnswer{} }
func (m *QuizQuestionAnswer) String() string { return proto.CompactTextString(m) }
func (*QuizQuestionAnswer) ProtoMessage()    {}
func (*QuizQuestionAnswer) Descriptor() ([]byte, []int) {
	return fileDescriptor_99abdaeba5813822, []int{6}
}

func (m *QuizQuestionAnswer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuizQuestionAnswer.Unmarshal(m, b)
}
func (m *QuizQuestionAnswer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuizQuestionAnswer.Marshal(b, m, deterministic)
}
func (m *QuizQuestionAnswer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuizQuestionAnswer.Merge(m, src)
}
func (m *QuizQuestionAnswer) XXX_Size() int {
	return xxx_messageInfo_QuizQuestionAnswer.Size(m)
}
func (m *QuizQuestionAnswer) XXX_DiscardUnknown() {
	xxx_messageInfo_QuizQuestionAnswer.DiscardUnknown(m)
}

var xxx_messageInfo_QuizQuestionAnswer proto.InternalMessageInfo

func (m *QuizQuestionAnswer) GetUuid() int32 {
	if m != nil {
		return m.Uuid
	}
	return 0
}

func (m *QuizQuestionAnswer) GetAnswer() *Option {
	if m != nil {
		return m.Answer
	}
	return nil
}

type QuizQuestionResponse struct {
	QuestionUuid         int32    `protobuf:"varint,1,opt,name=questionUuid,proto3" json:"questionUuid,omitempty"`
	CorrectAnswer        *Option  `protobuf:"bytes,2,opt,name=correctAnswer,proto3" json:"correctAnswer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuizQuestionResponse) Reset()         { *m = QuizQuestionResponse{} }
func (m *QuizQuestionResponse) String() string { return proto.CompactTextString(m) }
func (*QuizQuestionResponse) ProtoMessage()    {}
func (*QuizQuestionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_99abdaeba5813822, []int{7}
}

func (m *QuizQuestionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuizQuestionResponse.Unmarshal(m, b)
}
func (m *QuizQuestionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuizQuestionResponse.Marshal(b, m, deterministic)
}
func (m *QuizQuestionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuizQuestionResponse.Merge(m, src)
}
func (m *QuizQuestionResponse) XXX_Size() int {
	return xxx_messageInfo_QuizQuestionResponse.Size(m)
}
func (m *QuizQuestionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuizQuestionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuizQuestionResponse proto.InternalMessageInfo

func (m *QuizQuestionResponse) GetQuestionUuid() int32 {
	if m != nil {
		return m.QuestionUuid
	}
	return 0
}

func (m *QuizQuestionResponse) GetCorrectAnswer() *Option {
	if m != nil {
		return m.CorrectAnswer
	}
	return nil
}

type QuizQuestionStatistics struct {
	QuestionUuid               int32    `protobuf:"varint,1,opt,name=questionUuid,proto3" json:"questionUuid,omitempty"`
	PercentageOfCorrectAnswers float32  `protobuf:"fixed32,2,opt,name=percentageOfCorrectAnswers,proto3" json:"percentageOfCorrectAnswers,omitempty"`
	Participants               int32    `protobuf:"varint,3,opt,name=participants,proto3" json:"participants,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *QuizQuestionStatistics) Reset()         { *m = QuizQuestionStatistics{} }
func (m *QuizQuestionStatistics) String() string { return proto.CompactTextString(m) }
func (*QuizQuestionStatistics) ProtoMessage()    {}
func (*QuizQuestionStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_99abdaeba5813822, []int{8}
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

func init() {
	proto.RegisterType((*QuizRequest)(nil), "QuizRequest")
	proto.RegisterType((*QuizAnswer)(nil), "QuizAnswer")
	proto.RegisterType((*Question)(nil), "Question")
	proto.RegisterType((*QuizQuestion)(nil), "QuizQuestion")
	proto.RegisterType((*Option)(nil), "Option")
	proto.RegisterType((*QuizQuestionCreation)(nil), "QuizQuestionCreation")
	proto.RegisterType((*QuizQuestionAnswer)(nil), "QuizQuestionAnswer")
	proto.RegisterType((*QuizQuestionResponse)(nil), "QuizQuestionResponse")
	proto.RegisterType((*QuizQuestionStatistics)(nil), "QuizQuestionStatistics")
}

func init() { proto.RegisterFile("protos/quiz.proto", fileDescriptor_99abdaeba5813822) }

var fileDescriptor_99abdaeba5813822 = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x55, 0xd6, 0xf5, 0xeb, 0xb6, 0xe5, 0xc3, 0x1b, 0x53, 0x94, 0x07, 0xa8, 0x22, 0x40, 0x7d,
	0x59, 0x80, 0xf2, 0x8e, 0x18, 0x93, 0x90, 0xfa, 0x34, 0xcd, 0xd5, 0x84, 0xc4, 0x9b, 0x49, 0xcc,
	0x66, 0x54, 0xc5, 0xa9, 0xaf, 0x3d, 0xa4, 0xfd, 0x1a, 0x7e, 0x0f, 0xbf, 0x0a, 0xd9, 0x4e, 0x52,
	0x07, 0x0a, 0xec, 0xcd, 0xe7, 0xdc, 0x73, 0xef, 0xb9, 0xb9, 0xd7, 0x0e, 0x3c, 0xae, 0x94, 0xd4,
	0x12, 0x5f, 0x6d, 0x8d, 0xb8, 0xcb, 0xdc, 0x39, 0x39, 0xaa, 0x29, 0xd4, 0x4c, 0x1b, 0xf4, 0x64,
	0xfa, 0x02, 0x26, 0x97, 0x46, 0xdc, 0x51, 0xbe, 0x35, 0x1c, 0x35, 0x39, 0x81, 0x81, 0x41, 0xae,
	0x56, 0x45, 0x1c, 0xcd, 0xa3, 0xc5, 0x98, 0xd6, 0x28, 0xa5, 0x00, 0x56, 0x76, 0x56, 0xe2, 0x77,
	0xae, 0xfe, 0xa6, 0xb2, 0x3c, 0x73, 0x8a, 0xf8, 0xc0, 0xf3, 0x1e, 0x91, 0x63, 0xe8, 0x57, 0x37,
	0x52, 0xcb, 0xb8, 0x37, 0x8f, 0x16, 0x53, 0xea, 0x41, 0xfa, 0x33, 0x82, 0xd1, 0xa5, 0x75, 0x15,
	0xb2, 0x24, 0x31, 0x0c, 0xf3, 0x0d, 0x43, 0x6c, 0x6b, 0x36, 0xd0, 0x46, 0xae, 0x95, 0x34, 0xd5,
	0xaa, 0xa8, 0xab, 0x36, 0xd0, 0x46, 0xbc, 0x01, 0xc6, 0xbd, 0x79, 0xcf, 0x46, 0x6a, 0x48, 0x9e,
	0xc3, 0xcc, 0x79, 0x34, 0xe5, 0xe3, 0xc3, 0x79, 0xb4, 0x18, 0xd1, 0x2e, 0x49, 0x5e, 0xc2, 0x83,
	0x7c, 0x23, 0x91, 0x17, 0xad, 0xac, 0xef, 0x64, 0xbf, 0xb1, 0x24, 0x85, 0xa9, 0xac, 0x78, 0xd9,
	0xaa, 0x06, 0x4e, 0xd5, 0xe1, 0x52, 0x84, 0xa9, 0x1d, 0x50, 0x9b, 0x43, 0xe0, 0xd0, 0x18, 0xe1,
	0x3f, 0xa6, 0x4f, 0xdd, 0x99, 0x24, 0x30, 0xda, 0x36, 0x35, 0xfc, 0xa7, 0xb4, 0xd8, 0xea, 0x6f,
	0x44, 0xa9, 0xdd, 0x84, 0xc6, 0xd4, 0x9d, 0xc9, 0x33, 0x18, 0xc8, 0xaa, 0x6e, 0xbf, 0xb7, 0x98,
	0x2c, 0x87, 0xd9, 0x85, 0x83, 0xb4, 0xa6, 0xd3, 0xa7, 0x30, 0xf0, 0x8c, 0x9d, 0xf0, 0x2d, 0xdb,
	0x18, 0x5e, 0x0f, 0xcf, 0x83, 0xf4, 0x1b, 0x1c, 0x87, 0x4d, 0x9d, 0x2b, 0xce, 0x9c, 0xfa, 0x0d,
	0x4c, 0xb7, 0x01, 0xef, 0x92, 0x26, 0xcb, 0x59, 0x16, 0x8a, 0x69, 0x47, 0x62, 0x7b, 0x09, 0x56,
	0x1b, 0xf6, 0xe2, 0xe9, 0x74, 0x05, 0x24, 0x4c, 0xaf, 0x6f, 0xca, 0xbe, 0x31, 0xfc, 0xb7, 0x94,
	0xe8, 0xb6, 0x4d, 0x39, 0x56, 0xb2, 0x44, 0x6e, 0xf7, 0xd0, 0xcc, 0xeb, 0x6a, 0x57, 0xb4, 0xc3,
	0x91, 0x53, 0x98, 0xe5, 0x52, 0x29, 0x9e, 0xeb, 0xb3, 0xbd, 0x1e, 0xdd, 0x68, 0xfa, 0x23, 0x82,
	0x93, 0xd0, 0x6b, 0xad, 0x99, 0x16, 0xa8, 0x45, 0x8e, 0xf7, 0x72, 0x7b, 0x07, 0x49, 0xc5, 0x55,
	0xce, 0x4b, 0xcd, 0xae, 0xf9, 0xc5, 0xd7, 0xf3, 0xb0, 0x36, 0x3a, 0xeb, 0x03, 0xfa, 0x0f, 0x85,
	0xf5, 0xa8, 0x98, 0xd2, 0x22, 0x17, 0x15, 0x2b, 0x35, 0xba, 0xed, 0xf7, 0x69, 0x87, 0x5b, 0x16,
	0xfe, 0x85, 0xae, 0xb9, 0xba, 0x15, 0x39, 0x27, 0xa7, 0xf0, 0xe8, 0x13, 0x13, 0xfa, 0xa3, 0x54,
	0x4d, 0xcf, 0x48, 0xa6, 0x59, 0xf0, 0x86, 0x93, 0x71, 0xd6, 0x44, 0x5e, 0x47, 0xf6, 0x8e, 0x7b,
	0xb3, 0x76, 0x93, 0x93, 0x6c, 0xf7, 0x92, 0x93, 0x61, 0xb6, 0x76, 0x7f, 0x83, 0xe5, 0x1a, 0x1e,
	0x5e, 0xa1, 0x55, 0xed, 0x9c, 0xde, 0x03, 0x69, 0x52, 0x83, 0x8b, 0x70, 0x94, 0xfd, 0xb9, 0xe6,
	0xe4, 0x49, 0xb6, 0x6f, 0x61, 0x1f, 0xe0, 0xf3, 0x88, 0x17, 0x26, 0xdf, 0x48, 0x53, 0x7c, 0x19,
	0xb8, 0xff, 0xcd, 0xdb, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x8b, 0xaf, 0x55, 0x99, 0x04,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QuizServiceClient is the client API for QuizService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QuizServiceClient interface {
	WaitForQuestions(ctx context.Context, in *QuizRequest, opts ...grpc.CallOption) (QuizService_WaitForQuestionsClient, error)
	AnswerQuestion(ctx context.Context, in *QuizAnswer, opts ...grpc.CallOption) (*Status, error)
}

type quizServiceClient struct {
	cc *grpc.ClientConn
}

func NewQuizServiceClient(cc *grpc.ClientConn) QuizServiceClient {
	return &quizServiceClient{cc}
}

func (c *quizServiceClient) WaitForQuestions(ctx context.Context, in *QuizRequest, opts ...grpc.CallOption) (QuizService_WaitForQuestionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_QuizService_serviceDesc.Streams[0], "/QuizService/WaitForQuestions", opts...)
	if err != nil {
		return nil, err
	}
	x := &quizServiceWaitForQuestionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QuizService_WaitForQuestionsClient interface {
	Recv() (*Question, error)
	grpc.ClientStream
}

type quizServiceWaitForQuestionsClient struct {
	grpc.ClientStream
}

func (x *quizServiceWaitForQuestionsClient) Recv() (*Question, error) {
	m := new(Question)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *quizServiceClient) AnswerQuestion(ctx context.Context, in *QuizAnswer, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/QuizService/AnswerQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuizServiceServer is the server API for QuizService service.
type QuizServiceServer interface {
	WaitForQuestions(*QuizRequest, QuizService_WaitForQuestionsServer) error
	AnswerQuestion(context.Context, *QuizAnswer) (*Status, error)
}

// UnimplementedQuizServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQuizServiceServer struct {
}

func (*UnimplementedQuizServiceServer) WaitForQuestions(req *QuizRequest, srv QuizService_WaitForQuestionsServer) error {
	return status.Errorf(codes.Unimplemented, "method WaitForQuestions not implemented")
}
func (*UnimplementedQuizServiceServer) AnswerQuestion(ctx context.Context, req *QuizAnswer) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnswerQuestion not implemented")
}

func RegisterQuizServiceServer(s *grpc.Server, srv QuizServiceServer) {
	s.RegisterService(&_QuizService_serviceDesc, srv)
}

func _QuizService_WaitForQuestions_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(QuizRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QuizServiceServer).WaitForQuestions(m, &quizServiceWaitForQuestionsServer{stream})
}

type QuizService_WaitForQuestionsServer interface {
	Send(*Question) error
	grpc.ServerStream
}

type quizServiceWaitForQuestionsServer struct {
	grpc.ServerStream
}

func (x *quizServiceWaitForQuestionsServer) Send(m *Question) error {
	return x.ServerStream.SendMsg(m)
}

func _QuizService_AnswerQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuizAnswer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuizServiceServer).AnswerQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/QuizService/AnswerQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuizServiceServer).AnswerQuestion(ctx, req.(*QuizAnswer))
	}
	return interceptor(ctx, in, info, handler)
}

var _QuizService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "QuizService",
	HandlerType: (*QuizServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AnswerQuestion",
			Handler:    _QuizService_AnswerQuestion_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WaitForQuestions",
			Handler:       _QuizService_WaitForQuestions_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protos/quiz.proto",
}

// UserQuizServiceClient is the client API for UserQuizService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserQuizServiceClient interface {
	AnswerQuizQuestion(ctx context.Context, in *QuizQuestionAnswer, opts ...grpc.CallOption) (*QuizQuestionResponse, error)
}

type userQuizServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserQuizServiceClient(cc *grpc.ClientConn) UserQuizServiceClient {
	return &userQuizServiceClient{cc}
}

func (c *userQuizServiceClient) AnswerQuizQuestion(ctx context.Context, in *QuizQuestionAnswer, opts ...grpc.CallOption) (*QuizQuestionResponse, error) {
	out := new(QuizQuestionResponse)
	err := c.cc.Invoke(ctx, "/UserQuizService/AnswerQuizQuestion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserQuizServiceServer is the server API for UserQuizService service.
type UserQuizServiceServer interface {
	AnswerQuizQuestion(context.Context, *QuizQuestionAnswer) (*QuizQuestionResponse, error)
}

// UnimplementedUserQuizServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUserQuizServiceServer struct {
}

func (*UnimplementedUserQuizServiceServer) AnswerQuizQuestion(ctx context.Context, req *QuizQuestionAnswer) (*QuizQuestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnswerQuizQuestion not implemented")
}

func RegisterUserQuizServiceServer(s *grpc.Server, srv UserQuizServiceServer) {
	s.RegisterService(&_UserQuizService_serviceDesc, srv)
}

func _UserQuizService_AnswerQuizQuestion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuizQuestionAnswer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserQuizServiceServer).AnswerQuizQuestion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserQuizService/AnswerQuizQuestion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserQuizServiceServer).AnswerQuizQuestion(ctx, req.(*QuizQuestionAnswer))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserQuizService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UserQuizService",
	HandlerType: (*UserQuizServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AnswerQuizQuestion",
			Handler:    _UserQuizService_AnswerQuizQuestion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/quiz.proto",
}