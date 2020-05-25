// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users.proto

package blog

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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
	// ID auto incremented but not used externally
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// UUID for user (Primary Reference ID)
	ApiKey string `protobuf:"bytes,2,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	// Customer internal Reference ID (account, member ID, etc.)
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Active indicates if a user is still active or not
	Active bool `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`
	// Created By Reference
	CreatedById int32 `protobuf:"varint,20,opt,name=created_by_id,json=createdById,proto3" json:"created_by_id,omitempty"`
	// Timestamp of creation time
	CreatedAt string `protobuf:"bytes,21,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Updated By Reference
	UpdatedById int32 `protobuf:"varint,22,opt,name=updated_by_id,json=updatedById,proto3" json:"updated_by_id,omitempty"`
	// Timestamp of last update
	UpdatedAt            string   `protobuf:"bytes,23,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{0}
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

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func (m *User) GetCreatedById() int32 {
	if m != nil {
		return m.CreatedById
	}
	return 0
}

func (m *User) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *User) GetUpdatedById() int32 {
	if m != nil {
		return m.UpdatedById
	}
	return 0
}

func (m *User) GetUpdatedAt() string {
	if m != nil {
		return m.UpdatedAt
	}
	return ""
}

type UserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Ids                  []string `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
	Limit                int32    `protobuf:"varint,10,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,11,opt,name=offset,proto3" json:"offset,omitempty"`
	Total                int32    `protobuf:"varint,12,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserRequest) Reset()         { *m = UserRequest{} }
func (m *UserRequest) String() string { return proto.CompactTextString(m) }
func (*UserRequest) ProtoMessage()    {}
func (*UserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{1}
}

func (m *UserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserRequest.Unmarshal(m, b)
}
func (m *UserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserRequest.Marshal(b, m, deterministic)
}
func (m *UserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserRequest.Merge(m, src)
}
func (m *UserRequest) XXX_Size() int {
	return xxx_messageInfo_UserRequest.Size(m)
}
func (m *UserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserRequest proto.InternalMessageInfo

func (m *UserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UserRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *UserRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *UserRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *UserRequest) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

type UserResponse struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	Limit                int32    `protobuf:"varint,10,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,11,opt,name=offset,proto3" json:"offset,omitempty"`
	Total                int32    `protobuf:"varint,12,opt,name=total,proto3" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserResponse) Reset()         { *m = UserResponse{} }
func (m *UserResponse) String() string { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()    {}
func (*UserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{2}
}

func (m *UserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserResponse.Unmarshal(m, b)
}
func (m *UserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserResponse.Marshal(b, m, deterministic)
}
func (m *UserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserResponse.Merge(m, src)
}
func (m *UserResponse) XXX_Size() int {
	return xxx_messageInfo_UserResponse.Size(m)
}
func (m *UserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserResponse proto.InternalMessageInfo

func (m *UserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UserResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *UserResponse) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *UserResponse) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *UserResponse) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "blog.User")
	proto.RegisterType((*UserRequest)(nil), "blog.UserRequest")
	proto.RegisterType((*UserResponse)(nil), "blog.UserResponse")
}

func init() {
	proto.RegisterFile("users.proto", fileDescriptor_030765f334c86cea)
}

var fileDescriptor_030765f334c86cea = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x41, 0x8b, 0xd3, 0x40,
	0x18, 0x65, 0xda, 0xb4, 0xda, 0xc9, 0x2a, 0x32, 0xac, 0xeb, 0xb8, 0xee, 0xea, 0x30, 0xa7, 0x1e,
	0x4c, 0xab, 0xd5, 0x83, 0xee, 0x2d, 0x15, 0x04, 0xd9, 0x93, 0x41, 0x4f, 0x22, 0x65, 0xda, 0x7c,
	0x0d, 0x83, 0x6d, 0x26, 0x66, 0xbe, 0x28, 0x61, 0x29, 0x08, 0xde, 0x3c, 0x2e, 0xfe, 0x3b, 0xff,
	0x8c, 0xcc, 0x64, 0x0a, 0xab, 0x08, 0x22, 0x78, 0xca, 0xf7, 0xde, 0xf7, 0xf2, 0xe6, 0xf1, 0xf1,
	0x68, 0xdc, 0x58, 0xa8, 0xed, 0xa4, 0xaa, 0x0d, 0x1a, 0x16, 0x2d, 0x37, 0xa6, 0x38, 0x3e, 0x29,
	0x8c, 0x29, 0x36, 0x30, 0x55, 0x95, 0x9e, 0xaa, 0xb2, 0x34, 0xa8, 0x50, 0x9b, 0x32, 0x68, 0x8e,
	0xef, 0x86, 0xad, 0x47, 0xcb, 0x66, 0x3d, 0x55, 0x65, 0x1b, 0x56, 0x0f, 0xfd, 0x67, 0x95, 0x14,
	0x50, 0x26, 0xf6, 0xb3, 0x2a, 0x0a, 0xa8, 0xa7, 0xa6, 0xf2, 0x3f, 0xff, 0xc1, 0xe8, 0xe4, 0x77,
	0x23, 0x8b, 0x75, 0xb3, 0xc2, 0x6e, 0x2b, 0x7f, 0xf4, 0x68, 0xf4, 0xd6, 0x42, 0xcd, 0x6e, 0xd2,
	0x9e, 0xce, 0x39, 0x11, 0x64, 0x3c, 0xc8, 0x7a, 0x3a, 0x67, 0x5b, 0x7a, 0x4d, 0x55, 0x7a, 0xf1,
	0x01, 0x5a, 0xde, 0x13, 0x64, 0x3c, 0x9a, 0xbf, 0xb9, 0x4c, 0x5f, 0x7f, 0x21, 0xe4, 0x1b, 0x39,
	0x7f, 0xa7, 0x92, 0x75, 0x9a, 0xbc, 0x7c, 0x94, 0x3c, 0x7f, 0x7f, 0xf1, 0x6c, 0x97, 0x5c, 0x85,
	0x4f, 0xff, 0x05, 0x3e, 0x9e, 0xed, 0xb2, 0xa1, 0xaa, 0xf4, 0x39, 0xb4, 0x8c, 0xd1, 0xa8, 0x54,
	0x5b, 0xe0, 0x7d, 0xf7, 0x56, 0xe6, 0x67, 0x76, 0x44, 0x87, 0x6a, 0x85, 0xfa, 0x13, 0xf0, 0x48,
	0x90, 0xf1, 0xf5, 0x2c, 0x20, 0x26, 0xe9, 0x8d, 0x55, 0x0d, 0x0a, 0x21, 0x5f, 0x2c, 0xdb, 0x85,
	0xce, 0xf9, 0xa1, 0x4f, 0x1d, 0x07, 0x72, 0xde, 0xbe, 0xca, 0xd9, 0x29, 0xa5, 0x7b, 0x8d, 0x42,
	0x7e, 0xdb, 0xbb, 0x8e, 0x02, 0x93, 0xa2, 0xb3, 0x68, 0xaa, 0xfc, 0x8a, 0xc5, 0x51, 0x67, 0x11,
	0xc8, 0xbd, 0xc5, 0x5e, 0xa3, 0x90, 0xdf, 0xe9, 0x2c, 0x02, 0x93, 0xe2, 0x99, 0xbc, 0x4c, 0x1f,
	0xcc, 0x4e, 0xd9, 0xbd, 0x0b, 0x21, 0x5d, 0x5a, 0x79, 0x26, 0xe4, 0x8b, 0x1a, 0x20, 0x17, 0xf3,
	0x5a, 0x21, 0x9a, 0x52, 0x8a, 0x9d, 0xfc, 0x4a, 0x68, 0xec, 0xae, 0x9b, 0xc1, 0xc7, 0x06, 0x2c,
	0xb2, 0xfb, 0x34, 0x72, 0x3d, 0xf0, 0x67, 0x8e, 0x67, 0x74, 0xe2, 0x7a, 0x30, 0xf1, 0x02, 0xcf,
	0xb3, 0x5b, 0xb4, 0xaf, 0x73, 0xcb, 0x7b, 0xa2, 0x3f, 0x1e, 0x65, 0x6e, 0x64, 0x87, 0x74, 0xb0,
	0xd1, 0x5b, 0x8d, 0x9c, 0xfa, 0x80, 0x1d, 0x70, 0x97, 0x31, 0xeb, 0xb5, 0x05, 0xe4, 0xb1, 0xa7,
	0x03, 0x72, 0x6a, 0x34, 0xa8, 0x36, 0xfc, 0xa0, 0x53, 0x7b, 0x20, 0xbf, 0x13, 0x7a, 0xd0, 0xa5,
	0xb0, 0x95, 0x29, 0x2d, 0xfc, 0x35, 0x86, 0xa0, 0x03, 0x5f, 0x57, 0x1f, 0xe4, 0x57, 0x41, 0xb7,
	0xf8, 0x1f, 0xb1, 0x96, 0x43, 0xdf, 0xc0, 0x27, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x74, 0xde,
	0xe9, 0x1c, 0x1b, 0x03, 0x00, 0x00,
}
