// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/protobuf/user.proto

package rpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
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

type GetUserRequest struct {
	Email                string   `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bd554100d5a16d5, []int{0}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type GetUserResponse struct {
	User                 *User    `protobuf:"bytes,5,opt,name=User,proto3" json:"User,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bd554100d5a16d5, []int{1}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type User struct {
	ID                   uint64   `protobuf:"varint,4,opt,name=ID,proto3" json:"ID,omitempty"`
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=Password,proto3" json:"Password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bd554100d5a16d5, []int{2}
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

func (m *User) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUserRequest)(nil), "user.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "user.GetUserResponse")
	proto.RegisterType((*User)(nil), "user.User")
}

func init() { proto.RegisterFile("api/protobuf/user.proto", fileDescriptor_0bd554100d5a16d5) }

var fileDescriptor_0bd554100d5a16d5 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x2c, 0xc8, 0xd4,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x4f, 0x2a, 0x4d, 0xd3, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x03, 0xf3,
	0x84, 0x58, 0x40, 0x6c, 0x29, 0xe9, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0x84, 0x8a, 0xd4, 0xdc,
	0x82, 0x92, 0x4a, 0x88, 0x12, 0x25, 0x35, 0x2e, 0x3e, 0xf7, 0xd4, 0x92, 0xd0, 0xe2, 0xd4, 0xa2,
	0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x11, 0x2e, 0x56, 0xd7, 0xdc, 0xc4, 0xcc, 0x1c,
	0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0xc9, 0x90, 0x8b, 0x1f, 0xae, 0xae, 0xb8,
	0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x8e, 0x8b, 0x05, 0xc4, 0x97, 0x60, 0x55, 0x60, 0xd4, 0xe0,
	0x36, 0xe2, 0xd2, 0x03, 0x5b, 0x0c, 0x56, 0x01, 0x16, 0x57, 0x8a, 0x81, 0xc8, 0x0b, 0xf1, 0x71,
	0x31, 0x79, 0xba, 0x48, 0xb0, 0x28, 0x30, 0x6a, 0xb0, 0x04, 0x31, 0x79, 0xba, 0x08, 0x09, 0x71,
	0xb1, 0xf8, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x82, 0xcd, 0x07, 0xb3, 0xb1, 0x5b, 0x2a, 0x24, 0xc5,
	0xc5, 0x11, 0x90, 0x58, 0x5c, 0x5c, 0x9e, 0x5f, 0x94, 0x22, 0xc1, 0x0c, 0x96, 0x80, 0xf3, 0x8d,
	0x3c, 0xb8, 0x78, 0x40, 0xa6, 0x17, 0x07, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x59, 0x70,
	0xb1, 0x43, 0x1d, 0x28, 0x24, 0x02, 0x71, 0x0a, 0xaa, 0xbf, 0xa4, 0x44, 0xd1, 0x44, 0x21, 0xbe,
	0x50, 0x62, 0x70, 0x62, 0x8f, 0x62, 0xd5, 0xb3, 0x2e, 0x2a, 0x48, 0x4e, 0x62, 0x03, 0x07, 0x89,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x68, 0xe7, 0x99, 0x5d, 0x50, 0x01, 0x00, 0x00,
}
