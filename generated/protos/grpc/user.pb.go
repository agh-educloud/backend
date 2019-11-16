// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/grpc/user.proto

package educloud

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

type User struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c4ca6f17a95a046, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "User")
}

func init() { proto.RegisterFile("protos/grpc/user.proto", fileDescriptor_4c4ca6f17a95a046) }

var fileDescriptor_4c4ca6f17a95a046 = []byte{
	// 98 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x03, 0x0b, 0x28,
	0xe9, 0x71, 0xb1, 0x84, 0x16, 0xa7, 0x16, 0x09, 0x09, 0x71, 0xb1, 0x94, 0x96, 0x66, 0xa6, 0x48,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x20, 0xb1, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26,
	0x88, 0x18, 0x88, 0xed, 0xc4, 0x15, 0xc5, 0x91, 0x9a, 0x52, 0x9a, 0x9c, 0x93, 0x5f, 0x9a, 0x92,
	0xc4, 0x06, 0x36, 0xc2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xda, 0x6f, 0x82, 0x9e, 0x5c, 0x00,
	0x00, 0x00,
}