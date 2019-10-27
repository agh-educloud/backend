// Code generated by protoc-gen-go. DO NOT EDIT.
// source: homework.proto

package protos

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Document_DocumentType int32

const (
	Document_DOC Document_DocumentType = 0
	Document_PDF Document_DocumentType = 1
)

var Document_DocumentType_name = map[int32]string{
	0: "DOC",
	1: "PDF",
}

var Document_DocumentType_value = map[string]int32{
	"DOC": 0,
	"PDF": 1,
}

func (x Document_DocumentType) String() string {
	return proto.EnumName(Document_DocumentType_name, int32(x))
}

func (Document_DocumentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_395c96247398ffa4, []int{1, 0}
}

type Photo_PhotoFormat int32

const (
	Photo_JPG Photo_PhotoFormat = 0
	Photo_PNG Photo_PhotoFormat = 1
)

var Photo_PhotoFormat_name = map[int32]string{
	0: "JPG",
	1: "PNG",
}

var Photo_PhotoFormat_value = map[string]int32{
	"JPG": 0,
	"PNG": 1,
}

func (x Photo_PhotoFormat) String() string {
	return proto.EnumName(Photo_PhotoFormat_name, int32(x))
}

func (Photo_PhotoFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_395c96247398ffa4, []int{2, 0}
}

type Homework struct {
	HomeworkUuid int32 `protobuf:"varint,1,opt,name=homeworkUuid,proto3" json:"homeworkUuid,omitempty"`
	// Types that are valid to be assigned to HomeworkInstruction:
	//	*Homework_Document
	//	*Homework_Photo
	HomeworkInstruction  isHomework_HomeworkInstruction `protobuf_oneof:"homeworkInstruction"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *Homework) Reset()         { *m = Homework{} }
func (m *Homework) String() string { return proto.CompactTextString(m) }
func (*Homework) ProtoMessage()    {}
func (*Homework) Descriptor() ([]byte, []int) {
	return fileDescriptor_395c96247398ffa4, []int{0}
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

type isHomework_HomeworkInstruction interface {
	isHomework_HomeworkInstruction()
}

type Homework_Document struct {
	Document *Document `protobuf:"bytes,2,opt,name=document,proto3,oneof"`
}

type Homework_Photo struct {
	Photo *Photo `protobuf:"bytes,3,opt,name=photo,proto3,oneof"`
}

func (*Homework_Document) isHomework_HomeworkInstruction() {}

func (*Homework_Photo) isHomework_HomeworkInstruction() {}

func (m *Homework) GetHomeworkInstruction() isHomework_HomeworkInstruction {
	if m != nil {
		return m.HomeworkInstruction
	}
	return nil
}

func (m *Homework) GetDocument() *Document {
	if x, ok := m.GetHomeworkInstruction().(*Homework_Document); ok {
		return x.Document
	}
	return nil
}

func (m *Homework) GetPhoto() *Photo {
	if x, ok := m.GetHomeworkInstruction().(*Homework_Photo); ok {
		return x.Photo
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Homework) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Homework_Document)(nil),
		(*Homework_Photo)(nil),
	}
}

type Document struct {
	DocumentType         Document_DocumentType `protobuf:"varint,1,opt,name=documentType,proto3,enum=Document_DocumentType" json:"documentType,omitempty"`
	Document             []byte                `protobuf:"bytes,2,opt,name=document,proto3" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Document) Reset()         { *m = Document{} }
func (m *Document) String() string { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()    {}
func (*Document) Descriptor() ([]byte, []int) {
	return fileDescriptor_395c96247398ffa4, []int{1}
}

func (m *Document) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Document.Unmarshal(m, b)
}
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Document.Marshal(b, m, deterministic)
}
func (m *Document) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document.Merge(m, src)
}
func (m *Document) XXX_Size() int {
	return xxx_messageInfo_Document.Size(m)
}
func (m *Document) XXX_DiscardUnknown() {
	xxx_messageInfo_Document.DiscardUnknown(m)
}

var xxx_messageInfo_Document proto.InternalMessageInfo

func (m *Document) GetDocumentType() Document_DocumentType {
	if m != nil {
		return m.DocumentType
	}
	return Document_DOC
}

func (m *Document) GetDocument() []byte {
	if m != nil {
		return m.Document
	}
	return nil
}

type Photo struct {
	PhotoFormat          Photo_PhotoFormat `protobuf:"varint,1,opt,name=photoFormat,proto3,enum=Photo_PhotoFormat" json:"photoFormat,omitempty"`
	Photo                []byte            `protobuf:"bytes,2,opt,name=Photo,proto3" json:"Photo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Photo) Reset()         { *m = Photo{} }
func (m *Photo) String() string { return proto.CompactTextString(m) }
func (*Photo) ProtoMessage()    {}
func (*Photo) Descriptor() ([]byte, []int) {
	return fileDescriptor_395c96247398ffa4, []int{2}
}

func (m *Photo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Photo.Unmarshal(m, b)
}
func (m *Photo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Photo.Marshal(b, m, deterministic)
}
func (m *Photo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Photo.Merge(m, src)
}
func (m *Photo) XXX_Size() int {
	return xxx_messageInfo_Photo.Size(m)
}
func (m *Photo) XXX_DiscardUnknown() {
	xxx_messageInfo_Photo.DiscardUnknown(m)
}

var xxx_messageInfo_Photo proto.InternalMessageInfo

func (m *Photo) GetPhotoFormat() Photo_PhotoFormat {
	if m != nil {
		return m.PhotoFormat
	}
	return Photo_JPG
}

func (m *Photo) GetPhoto() []byte {
	if m != nil {
		return m.Photo
	}
	return nil
}

type CreateHomeworkRequest struct {
	ClassUuid            int32     `protobuf:"varint,1,opt,name=classUuid,proto3" json:"classUuid,omitempty"`
	Homework             *Homework `protobuf:"bytes,2,opt,name=homework,proto3" json:"homework,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateHomeworkRequest) Reset()         { *m = CreateHomeworkRequest{} }
func (m *CreateHomeworkRequest) String() string { return proto.CompactTextString(m) }
func (*CreateHomeworkRequest) ProtoMessage()    {}
func (*CreateHomeworkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_395c96247398ffa4, []int{3}
}

func (m *CreateHomeworkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateHomeworkRequest.Unmarshal(m, b)
}
func (m *CreateHomeworkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateHomeworkRequest.Marshal(b, m, deterministic)
}
func (m *CreateHomeworkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateHomeworkRequest.Merge(m, src)
}
func (m *CreateHomeworkRequest) XXX_Size() int {
	return xxx_messageInfo_CreateHomeworkRequest.Size(m)
}
func (m *CreateHomeworkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateHomeworkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateHomeworkRequest proto.InternalMessageInfo

func (m *CreateHomeworkRequest) GetClassUuid() int32 {
	if m != nil {
		return m.ClassUuid
	}
	return 0
}

func (m *CreateHomeworkRequest) GetHomework() *Homework {
	if m != nil {
		return m.Homework
	}
	return nil
}

type CreateHomeworkResponse struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	HomeworkUuid         int32    `protobuf:"varint,2,opt,name=homeworkUuid,proto3" json:"homeworkUuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateHomeworkResponse) Reset()         { *m = CreateHomeworkResponse{} }
func (m *CreateHomeworkResponse) String() string { return proto.CompactTextString(m) }
func (*CreateHomeworkResponse) ProtoMessage()    {}
func (*CreateHomeworkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_395c96247398ffa4, []int{4}
}

func (m *CreateHomeworkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateHomeworkResponse.Unmarshal(m, b)
}
func (m *CreateHomeworkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateHomeworkResponse.Marshal(b, m, deterministic)
}
func (m *CreateHomeworkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateHomeworkResponse.Merge(m, src)
}
func (m *CreateHomeworkResponse) XXX_Size() int {
	return xxx_messageInfo_CreateHomeworkResponse.Size(m)
}
func (m *CreateHomeworkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateHomeworkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateHomeworkResponse proto.InternalMessageInfo

func (m *CreateHomeworkResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CreateHomeworkResponse) GetHomeworkUuid() int32 {
	if m != nil {
		return m.HomeworkUuid
	}
	return 0
}

type UpdateHomeworkRequest struct {
	ClassUuid            int32     `protobuf:"varint,1,opt,name=classUuid,proto3" json:"classUuid,omitempty"`
	HomeworkUuid         int32     `protobuf:"varint,2,opt,name=homeworkUuid,proto3" json:"homeworkUuid,omitempty"`
	Homework             *Homework `protobuf:"bytes,3,opt,name=homework,proto3" json:"homework,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UpdateHomeworkRequest) Reset()         { *m = UpdateHomeworkRequest{} }
func (m *UpdateHomeworkRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateHomeworkRequest) ProtoMessage()    {}
func (*UpdateHomeworkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_395c96247398ffa4, []int{5}
}

func (m *UpdateHomeworkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateHomeworkRequest.Unmarshal(m, b)
}
func (m *UpdateHomeworkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateHomeworkRequest.Marshal(b, m, deterministic)
}
func (m *UpdateHomeworkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateHomeworkRequest.Merge(m, src)
}
func (m *UpdateHomeworkRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateHomeworkRequest.Size(m)
}
func (m *UpdateHomeworkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateHomeworkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateHomeworkRequest proto.InternalMessageInfo

func (m *UpdateHomeworkRequest) GetClassUuid() int32 {
	if m != nil {
		return m.ClassUuid
	}
	return 0
}

func (m *UpdateHomeworkRequest) GetHomeworkUuid() int32 {
	if m != nil {
		return m.HomeworkUuid
	}
	return 0
}

func (m *UpdateHomeworkRequest) GetHomework() *Homework {
	if m != nil {
		return m.Homework
	}
	return nil
}

type DeleteHomeworkRequest struct {
	ClassUuid            int32    `protobuf:"varint,1,opt,name=classUuid,proto3" json:"classUuid,omitempty"`
	HomeworkUuid         int32    `protobuf:"varint,2,opt,name=homeworkUuid,proto3" json:"homeworkUuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteHomeworkRequest) Reset()         { *m = DeleteHomeworkRequest{} }
func (m *DeleteHomeworkRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteHomeworkRequest) ProtoMessage()    {}
func (*DeleteHomeworkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_395c96247398ffa4, []int{6}
}

func (m *DeleteHomeworkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteHomeworkRequest.Unmarshal(m, b)
}
func (m *DeleteHomeworkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteHomeworkRequest.Marshal(b, m, deterministic)
}
func (m *DeleteHomeworkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteHomeworkRequest.Merge(m, src)
}
func (m *DeleteHomeworkRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteHomeworkRequest.Size(m)
}
func (m *DeleteHomeworkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteHomeworkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteHomeworkRequest proto.InternalMessageInfo

func (m *DeleteHomeworkRequest) GetClassUuid() int32 {
	if m != nil {
		return m.ClassUuid
	}
	return 0
}

func (m *DeleteHomeworkRequest) GetHomeworkUuid() int32 {
	if m != nil {
		return m.HomeworkUuid
	}
	return 0
}

type HomeworkSolution struct {
	HomeworkUuid int32 `protobuf:"varint,1,opt,name=homeworkUuid,proto3" json:"homeworkUuid,omitempty"`
	// Types that are valid to be assigned to HomeworkSolution:
	//	*HomeworkSolution_Document
	//	*HomeworkSolution_Photo
	HomeworkSolution     isHomeworkSolution_HomeworkSolution `protobuf_oneof:"homeworkSolution"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *HomeworkSolution) Reset()         { *m = HomeworkSolution{} }
func (m *HomeworkSolution) String() string { return proto.CompactTextString(m) }
func (*HomeworkSolution) ProtoMessage()    {}
func (*HomeworkSolution) Descriptor() ([]byte, []int) {
	return fileDescriptor_395c96247398ffa4, []int{7}
}

func (m *HomeworkSolution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HomeworkSolution.Unmarshal(m, b)
}
func (m *HomeworkSolution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HomeworkSolution.Marshal(b, m, deterministic)
}
func (m *HomeworkSolution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HomeworkSolution.Merge(m, src)
}
func (m *HomeworkSolution) XXX_Size() int {
	return xxx_messageInfo_HomeworkSolution.Size(m)
}
func (m *HomeworkSolution) XXX_DiscardUnknown() {
	xxx_messageInfo_HomeworkSolution.DiscardUnknown(m)
}

var xxx_messageInfo_HomeworkSolution proto.InternalMessageInfo

func (m *HomeworkSolution) GetHomeworkUuid() int32 {
	if m != nil {
		return m.HomeworkUuid
	}
	return 0
}

type isHomeworkSolution_HomeworkSolution interface {
	isHomeworkSolution_HomeworkSolution()
}

type HomeworkSolution_Document struct {
	Document *Document `protobuf:"bytes,2,opt,name=document,proto3,oneof"`
}

type HomeworkSolution_Photo struct {
	Photo *Photo `protobuf:"bytes,3,opt,name=photo,proto3,oneof"`
}

func (*HomeworkSolution_Document) isHomeworkSolution_HomeworkSolution() {}

func (*HomeworkSolution_Photo) isHomeworkSolution_HomeworkSolution() {}

func (m *HomeworkSolution) GetHomeworkSolution() isHomeworkSolution_HomeworkSolution {
	if m != nil {
		return m.HomeworkSolution
	}
	return nil
}

func (m *HomeworkSolution) GetDocument() *Document {
	if x, ok := m.GetHomeworkSolution().(*HomeworkSolution_Document); ok {
		return x.Document
	}
	return nil
}

func (m *HomeworkSolution) GetPhoto() *Photo {
	if x, ok := m.GetHomeworkSolution().(*HomeworkSolution_Photo); ok {
		return x.Photo
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HomeworkSolution) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HomeworkSolution_Document)(nil),
		(*HomeworkSolution_Photo)(nil),
	}
}

func init() {
	proto.RegisterEnum("Document_DocumentType", Document_DocumentType_name, Document_DocumentType_value)
	proto.RegisterEnum("Photo_PhotoFormat", Photo_PhotoFormat_name, Photo_PhotoFormat_value)
	proto.RegisterType((*Homework)(nil), "Homework")
	proto.RegisterType((*Document)(nil), "Document")
	proto.RegisterType((*Photo)(nil), "Photo")
	proto.RegisterType((*CreateHomeworkRequest)(nil), "CreateHomeworkRequest")
	proto.RegisterType((*CreateHomeworkResponse)(nil), "CreateHomeworkResponse")
	proto.RegisterType((*UpdateHomeworkRequest)(nil), "UpdateHomeworkRequest")
	proto.RegisterType((*DeleteHomeworkRequest)(nil), "DeleteHomeworkRequest")
	proto.RegisterType((*HomeworkSolution)(nil), "HomeworkSolution")
}

func init() { proto.RegisterFile("homework.proto", fileDescriptor_395c96247398ffa4) }

var fileDescriptor_395c96247398ffa4 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4d, 0x8f, 0x12, 0x41,
	0x10, 0xa5, 0x97, 0xc0, 0xb2, 0xc5, 0x84, 0x8c, 0xb5, 0xcb, 0x48, 0x88, 0x71, 0x49, 0x27, 0xea,
	0x9e, 0x3a, 0x71, 0xf4, 0xe4, 0x4d, 0x97, 0xec, 0xb2, 0x1e, 0x94, 0x0c, 0x72, 0x30, 0xd1, 0x03,
	0x42, 0x45, 0x88, 0x30, 0x3d, 0x4e, 0xf7, 0x68, 0xbc, 0x79, 0xf0, 0xec, 0xc5, 0x9f, 0xe9, 0x9f,
	0x30, 0xf3, 0xd5, 0xcc, 0x30, 0x90, 0xe8, 0xc1, 0xbd, 0x90, 0xae, 0xaf, 0x57, 0xaf, 0x78, 0x0f,
	0xa0, 0xb3, 0x94, 0x1b, 0xfa, 0x2a, 0xc3, 0x4f, 0x22, 0x08, 0xa5, 0x96, 0x7d, 0x4b, 0xe9, 0x99,
	0x8e, 0x54, 0x1a, 0xf1, 0x9f, 0x0c, 0x5a, 0xa3, 0xac, 0x01, 0x39, 0x58, 0x79, 0xf3, 0x34, 0x5a,
	0x2d, 0x7a, 0x6c, 0xc0, 0x2e, 0x1a, 0x5e, 0x29, 0x87, 0x8f, 0xa0, 0xb5, 0x90, 0xf3, 0x68, 0x43,
	0xbe, 0xee, 0x1d, 0x0d, 0xd8, 0x45, 0xdb, 0x3d, 0x11, 0xc3, 0x2c, 0x31, 0xaa, 0x79, 0xa6, 0x88,
	0xf7, 0xa1, 0x11, 0x2c, 0xa5, 0x96, 0xbd, 0x7a, 0xd2, 0xd5, 0x14, 0xe3, 0x38, 0x1a, 0xd5, 0xbc,
	0x34, 0xfd, 0xa2, 0x0b, 0xa7, 0x39, 0xf0, 0x8d, 0xaf, 0x74, 0x18, 0xcd, 0xf5, 0x4a, 0xfa, 0xfc,
	0x07, 0x83, 0x56, 0x8e, 0x87, 0xcf, 0xc0, 0xca, 0xf1, 0xde, 0x7c, 0x0b, 0x28, 0x21, 0xd4, 0x71,
	0x1d, 0xb3, 0xd0, 0x3c, 0xe2, 0xaa, 0x57, 0xea, 0xc5, 0xfe, 0x0e, 0x51, 0x6b, 0xcb, 0x8d, 0x0f,
	0xc0, 0x2a, 0x4e, 0xe2, 0x31, 0xd4, 0x87, 0xaf, 0x2f, 0xed, 0x5a, 0xfc, 0x18, 0x0f, 0xaf, 0x6c,
	0xc6, 0x35, 0x34, 0x12, 0xbe, 0xf8, 0x14, 0xda, 0x09, 0xdf, 0x2b, 0x19, 0x6e, 0x66, 0x3a, 0x63,
	0x80, 0xe9, 0x31, 0xe9, 0x67, 0x5a, 0xf1, 0x8a, 0x6d, 0x78, 0x96, 0x8d, 0x67, 0x9b, 0xd3, 0x80,
	0x9f, 0x43, 0xbb, 0x30, 0x11, 0x2f, 0x7b, 0x39, 0xbe, 0xce, 0xb6, 0xbe, 0xba, 0xb6, 0x19, 0x7f,
	0x07, 0xdd, 0xcb, 0x90, 0x66, 0x9a, 0x72, 0x49, 0x3c, 0xfa, 0x1c, 0x91, 0xd2, 0x78, 0x0f, 0x4e,
	0xe6, 0xeb, 0x99, 0x52, 0x05, 0x59, 0xb6, 0x09, 0x7c, 0x00, 0xad, 0xfc, 0xab, 0x34, 0x9a, 0x18,
	0x04, 0x53, 0xe2, 0xef, 0xc1, 0xd9, 0x45, 0x57, 0x81, 0xf4, 0x15, 0xe1, 0x39, 0x34, 0x53, 0x57,
	0x24, 0xd8, 0x6d, 0xf7, 0x58, 0x4c, 0x92, 0xd0, 0xcb, 0xd2, 0x15, 0x67, 0x1c, 0x55, 0x9d, 0xc1,
	0xbf, 0x33, 0xe8, 0x4e, 0x83, 0xc5, 0x3f, 0xb3, 0xff, 0x0b, 0xec, 0xd2, 0x85, 0xf5, 0xc3, 0x17,
	0xbe, 0x85, 0xee, 0x90, 0xd6, 0xf4, 0x1f, 0x18, 0xf0, 0x5f, 0x0c, 0xec, 0x1c, 0x75, 0x22, 0xd7,
	0x51, 0x6c, 0xd6, 0xdb, 0xfd, 0xc1, 0x20, 0xd8, 0xcb, 0x1d, 0x02, 0xee, 0x6f, 0x06, 0xbd, 0x71,
	0x48, 0x8a, 0x7c, 0x4d, 0xa1, 0xa1, 0x47, 0xe1, 0x97, 0xd5, 0x9c, 0xf0, 0x39, 0x74, 0xca, 0x7a,
	0xa3, 0x23, 0xf6, 0xda, 0xab, 0x7f, 0x57, 0x1c, 0x30, 0xc6, 0x63, 0xe8, 0x94, 0x25, 0x45, 0x47,
	0xec, 0xd5, 0xb8, 0x9f, 0x5b, 0x26, 0x1e, 0x29, 0x6b, 0x80, 0x8e, 0xd8, 0x2b, 0xca, 0x76, 0xe4,
	0x21, 0xd8, 0x71, 0xc7, 0xc7, 0xe2, 0x9e, 0xad, 0xbe, 0xa6, 0xcf, 0xbd, 0x81, 0xd3, 0xa9, 0xaa,
	0xde, 0xe9, 0xc2, 0xd9, 0x84, 0xfc, 0x45, 0x45, 0x9d, 0x3b, 0x62, 0x37, 0x65, 0xa0, 0x3e, 0x34,
	0x93, 0xbf, 0xbf, 0x27, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x34, 0xf8, 0xad, 0x1e, 0x05,
	0x00, 0x00,
}
