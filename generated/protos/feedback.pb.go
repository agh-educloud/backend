// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feedback.proto

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

type Feedback_Rate int32

const (
	Feedback_TERRIBLE  Feedback_Rate = 0
	Feedback_BAD       Feedback_Rate = 1
	Feedback_OK        Feedback_Rate = 2
	Feedback_GOOD      Feedback_Rate = 3
	Feedback_EXCELLENT Feedback_Rate = 4
)

var Feedback_Rate_name = map[int32]string{
	0: "TERRIBLE",
	1: "BAD",
	2: "OK",
	3: "GOOD",
	4: "EXCELLENT",
}

var Feedback_Rate_value = map[string]int32{
	"TERRIBLE":  0,
	"BAD":       1,
	"OK":        2,
	"GOOD":      3,
	"EXCELLENT": 4,
}

func (x Feedback_Rate) String() string {
	return proto.EnumName(Feedback_Rate_name, int32(x))
}

func (Feedback_Rate) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{0, 0}
}

type Feedback struct {
	Rate                 Feedback_Rate `protobuf:"varint,1,opt,name=rate,proto3,enum=Feedback_Rate" json:"rate,omitempty"`
	WhatUserLike         string        `protobuf:"bytes,2,opt,name=whatUserLike,proto3" json:"whatUserLike,omitempty"`
	WhatUserDislike      string        `protobuf:"bytes,3,opt,name=whatUserDislike,proto3" json:"whatUserDislike,omitempty"`
	AdditionalComment    string        `protobuf:"bytes,4,opt,name=additionalComment,proto3" json:"additionalComment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Feedback) Reset()         { *m = Feedback{} }
func (m *Feedback) String() string { return proto.CompactTextString(m) }
func (*Feedback) ProtoMessage()    {}
func (*Feedback) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b189e8c8330c03e, []int{0}
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

func (m *Feedback) GetRate() Feedback_Rate {
	if m != nil {
		return m.Rate
	}
	return Feedback_TERRIBLE
}

func (m *Feedback) GetWhatUserLike() string {
	if m != nil {
		return m.WhatUserLike
	}
	return ""
}

func (m *Feedback) GetWhatUserDislike() string {
	if m != nil {
		return m.WhatUserDislike
	}
	return ""
}

func (m *Feedback) GetAdditionalComment() string {
	if m != nil {
		return m.AdditionalComment
	}
	return ""
}

func init() {
	proto.RegisterEnum("Feedback_Rate", Feedback_Rate_name, Feedback_Rate_value)
	proto.RegisterType((*Feedback)(nil), "Feedback")
}

func init() { proto.RegisterFile("feedback.proto", fileDescriptor_7b189e8c8330c03e) }

var fileDescriptor_7b189e8c8330c03e = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x4b, 0x4d, 0x4d,
	0x49, 0x4a, 0x4c, 0xce, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x7a, 0xc5, 0xc8, 0xc5, 0xe1,
	0x06, 0x15, 0x12, 0x52, 0xe2, 0x62, 0x29, 0x4a, 0x2c, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x33, 0xe2, 0xd3, 0x83, 0x49, 0xe8, 0x05, 0x25, 0x96, 0xa4, 0x06, 0x81, 0xe5, 0x84, 0x94, 0xb8,
	0x78, 0xca, 0x33, 0x12, 0x4b, 0x42, 0x8b, 0x53, 0x8b, 0x7c, 0x32, 0xb3, 0x53, 0x25, 0x98, 0x14,
	0x18, 0x35, 0x38, 0x83, 0x50, 0xc4, 0x84, 0x34, 0xb8, 0xf8, 0x61, 0x7c, 0x97, 0xcc, 0xe2, 0x1c,
	0x90, 0x32, 0x66, 0xb0, 0x32, 0x74, 0x61, 0x21, 0x1d, 0x2e, 0xc1, 0xc4, 0x94, 0x94, 0xcc, 0x92,
	0xcc, 0xfc, 0xbc, 0xc4, 0x1c, 0xe7, 0xfc, 0xdc, 0xdc, 0xd4, 0xbc, 0x12, 0x09, 0x16, 0xb0, 0x5a,
	0x4c, 0x09, 0x25, 0x3b, 0x2e, 0x16, 0x90, 0x4b, 0x84, 0x78, 0xb8, 0x38, 0x42, 0x5c, 0x83, 0x82,
	0x3c, 0x9d, 0x7c, 0x5c, 0x05, 0x18, 0x84, 0xd8, 0xb9, 0x98, 0x9d, 0x1c, 0x5d, 0x04, 0x18, 0x85,
	0xd8, 0xb8, 0x98, 0xfc, 0xbd, 0x05, 0x98, 0x84, 0x38, 0xb8, 0x58, 0xdc, 0xfd, 0xfd, 0x5d, 0x04,
	0x98, 0x85, 0x78, 0xb9, 0x38, 0x5d, 0x23, 0x9c, 0x5d, 0x7d, 0x7c, 0x5c, 0xfd, 0x42, 0x04, 0x58,
	0x92, 0xd8, 0xc0, 0x7e, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x16, 0xb6, 0x0e, 0x05,
	0x01, 0x00, 0x00,
}
