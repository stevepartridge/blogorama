// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.8.0
// source: posts.proto

package blog

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUID for post
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Title of the blog post
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// Content body of the blog post
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	// Flag for if post is private
	Private bool `protobuf:"varint,5,opt,name=private,proto3" json:"private,omitempty"`
	// Created By Reference
	CreatedById int32 `protobuf:"varint,20,opt,name=created_by_id,json=createdById,proto3" json:"created_by_id,omitempty"`
	// Timestamp of creation time
	CreatedAt string `protobuf:"bytes,21,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Updated By Reference
	UpdatedById int32 `protobuf:"varint,22,opt,name=updated_by_id,json=updatedById,proto3" json:"updated_by_id,omitempty"`
	// Timestamp of last update
	UpdatedAt string `protobuf:"bytes,23,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Post) GetPrivate() bool {
	if x != nil {
		return x.Private
	}
	return false
}

func (x *Post) GetCreatedById() int32 {
	if x != nil {
		return x.CreatedById
	}
	return 0
}

func (x *Post) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Post) GetUpdatedById() int32 {
	if x != nil {
		return x.UpdatedById
	}
	return 0
}

func (x *Post) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type PostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post   *Post   `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	Ids    []int32 `protobuf:"varint,2,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	UserId int32   `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Limit  int32   `protobuf:"varint,10,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32   `protobuf:"varint,11,opt,name=offset,proto3" json:"offset,omitempty"`
	Total  int32   `protobuf:"varint,12,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *PostRequest) Reset() {
	*x = PostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostRequest) ProtoMessage() {}

func (x *PostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostRequest.ProtoReflect.Descriptor instead.
func (*PostRequest) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{1}
}

func (x *PostRequest) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

func (x *PostRequest) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *PostRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PostRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PostRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *PostRequest) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type PostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post   *Post   `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	Posts  []*Post `protobuf:"bytes,2,rep,name=posts,proto3" json:"posts,omitempty"`
	Limit  int32   `protobuf:"varint,10,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32   `protobuf:"varint,11,opt,name=offset,proto3" json:"offset,omitempty"`
	Total  int32   `protobuf:"varint,12,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *PostResponse) Reset() {
	*x = PostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostResponse) ProtoMessage() {}

func (x *PostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_posts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostResponse.ProtoReflect.Descriptor instead.
func (*PostResponse) Descriptor() ([]byte, []int) {
	return file_posts_proto_rawDescGZIP(), []int{2}
}

func (x *PostResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

func (x *PostResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

func (x *PostResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *PostResponse) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *PostResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_posts_proto protoreflect.FileDescriptor

var file_posts_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x62,
	0x6c, 0x6f, 0x67, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77,
	0x61, 0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa8, 0x02, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x3a, 0x40, 0x92, 0x41, 0x3d, 0x32, 0x3b, 0x12,
	0x39, 0x7b, 0x20, 0x22, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4d, 0x79, 0x20,
	0x46, 0x69, 0x72, 0x73, 0x74, 0x20, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x2c, 0x20, 0x22, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x3a, 0x20, 0x22, 0x4d, 0x79, 0x20, 0x70, 0x6f, 0x73, 0x74,
	0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x7d, 0x22, 0x9c, 0x01, 0x0a, 0x0b, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x70, 0x6f,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x94, 0x01, 0x0a, 0x0c, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x70, 0x6f,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x70, 0x6f,
	0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x6c, 0x6f, 0x67,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x74, 0x65, 0x76, 0x65, 0x70, 0x61, 0x72, 0x74, 0x72, 0x69, 0x64, 0x67, 0x65, 0x2f, 0x62, 0x6c,
	0x6f, 0x67, 0x6f, 0x72, 0x61, 0x6d, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67,
	0x6f, 0x3b, 0x62, 0x6c, 0x6f, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_posts_proto_rawDescOnce sync.Once
	file_posts_proto_rawDescData = file_posts_proto_rawDesc
)

func file_posts_proto_rawDescGZIP() []byte {
	file_posts_proto_rawDescOnce.Do(func() {
		file_posts_proto_rawDescData = protoimpl.X.CompressGZIP(file_posts_proto_rawDescData)
	})
	return file_posts_proto_rawDescData
}

var file_posts_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_posts_proto_goTypes = []interface{}{
	(*Post)(nil),         // 0: blog.Post
	(*PostRequest)(nil),  // 1: blog.PostRequest
	(*PostResponse)(nil), // 2: blog.PostResponse
}
var file_posts_proto_depIdxs = []int32{
	0, // 0: blog.PostRequest.post:type_name -> blog.Post
	0, // 1: blog.PostResponse.post:type_name -> blog.Post
	0, // 2: blog.PostResponse.posts:type_name -> blog.Post
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_posts_proto_init() }
func file_posts_proto_init() {
	if File_posts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_posts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_posts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_posts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_posts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_posts_proto_goTypes,
		DependencyIndexes: file_posts_proto_depIdxs,
		MessageInfos:      file_posts_proto_msgTypes,
	}.Build()
	File_posts_proto = out.File
	file_posts_proto_rawDesc = nil
	file_posts_proto_goTypes = nil
	file_posts_proto_depIdxs = nil
}
