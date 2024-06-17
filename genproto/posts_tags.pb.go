// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: posts_tags.proto

package genproto

import (
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

type PostTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostTagId string `protobuf:"bytes,1,opt,name=PostTagId,proto3" json:"PostTagId,omitempty"`
	PostId    string `protobuf:"bytes,2,opt,name=PostId,proto3" json:"PostId,omitempty"`
	TagId     string `protobuf:"bytes,3,opt,name=TagId,proto3" json:"TagId,omitempty"`
}

func (x *PostTag) Reset() {
	*x = PostTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_tags_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTag) ProtoMessage() {}

func (x *PostTag) ProtoReflect() protoreflect.Message {
	mi := &file_posts_tags_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTag.ProtoReflect.Descriptor instead.
func (*PostTag) Descriptor() ([]byte, []int) {
	return file_posts_tags_proto_rawDescGZIP(), []int{0}
}

func (x *PostTag) GetPostTagId() string {
	if x != nil {
		return x.PostTagId
	}
	return ""
}

func (x *PostTag) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *PostTag) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

type PostTagModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string `protobuf:"bytes,2,opt,name=PostId,proto3" json:"PostId,omitempty"`
	TagId  string `protobuf:"bytes,3,opt,name=TagId,proto3" json:"TagId,omitempty"`
}

func (x *PostTagModel) Reset() {
	*x = PostTagModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_tags_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostTagModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTagModel) ProtoMessage() {}

func (x *PostTagModel) ProtoReflect() protoreflect.Message {
	mi := &file_posts_tags_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTagModel.ProtoReflect.Descriptor instead.
func (*PostTagModel) Descriptor() ([]byte, []int) {
	return file_posts_tags_proto_rawDescGZIP(), []int{1}
}

func (x *PostTagModel) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *PostTagModel) GetTagId() string {
	if x != nil {
		return x.TagId
	}
	return ""
}

type ByPost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string `protobuf:"bytes,1,opt,name=PostId,proto3" json:"PostId,omitempty"`
}

func (x *ByPost) Reset() {
	*x = ByPost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_tags_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByPost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByPost) ProtoMessage() {}

func (x *ByPost) ProtoReflect() protoreflect.Message {
	mi := &file_posts_tags_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByPost.ProtoReflect.Descriptor instead.
func (*ByPost) Descriptor() ([]byte, []int) {
	return file_posts_tags_proto_rawDescGZIP(), []int{2}
}

func (x *ByPost) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

type PostTags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostTags []*PostTag `protobuf:"bytes,1,rep,name=PostTags,proto3" json:"PostTags,omitempty"`
}

func (x *PostTags) Reset() {
	*x = PostTags{}
	if protoimpl.UnsafeEnabled {
		mi := &file_posts_tags_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostTags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostTags) ProtoMessage() {}

func (x *PostTags) ProtoReflect() protoreflect.Message {
	mi := &file_posts_tags_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostTags.ProtoReflect.Descriptor instead.
func (*PostTags) Descriptor() ([]byte, []int) {
	return file_posts_tags_proto_rawDescGZIP(), []int{3}
}

func (x *PostTags) GetPostTags() []*PostTag {
	if x != nil {
		return x.PostTags
	}
	return nil
}

var File_posts_tags_proto protoreflect.FileDescriptor

var file_posts_tags_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x55,
	0x0a, 0x07, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x6f, 0x73,
	0x74, 0x54, 0x61, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x6f,
	0x73, 0x74, 0x54, 0x61, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x54, 0x61, 0x67, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x54, 0x61, 0x67, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x0c, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x54, 0x61, 0x67, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x61,
	0x67, 0x49, 0x64, 0x22, 0x20, 0x0a, 0x06, 0x42, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50,
	0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67,
	0x73, 0x12, 0x31, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x52, 0x08, 0x50, 0x6f, 0x73, 0x74,
	0x54, 0x61, 0x67, 0x73, 0x32, 0xc1, 0x02, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x54, 0x61, 0x67, 0x73, 0x12, 0x14, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x42, 0x79, 0x50, 0x6f, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x66,
	0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x54, 0x61, 0x67, 0x73, 0x12, 0x3d, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x54, 0x61, 0x67, 0x12, 0x15, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x1a, 0x15, 0x2e, 0x66,
	0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x54, 0x61, 0x67, 0x12, 0x37, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61,
	0x67, 0x12, 0x12, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x15, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x12, 0x3b, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x12, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x12, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x3d, 0x0a, 0x0d, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x12, 0x15, 0x2e, 0x66, 0x6f, 0x72,
	0x75, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61,
	0x67, 0x1a, 0x15, 0x2e, 0x66, 0x6f, 0x72, 0x75, 0x6d, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x54, 0x61, 0x67, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_posts_tags_proto_rawDescOnce sync.Once
	file_posts_tags_proto_rawDescData = file_posts_tags_proto_rawDesc
)

func file_posts_tags_proto_rawDescGZIP() []byte {
	file_posts_tags_proto_rawDescOnce.Do(func() {
		file_posts_tags_proto_rawDescData = protoimpl.X.CompressGZIP(file_posts_tags_proto_rawDescData)
	})
	return file_posts_tags_proto_rawDescData
}

var file_posts_tags_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_posts_tags_proto_goTypes = []interface{}{
	(*PostTag)(nil),      // 0: forum_protos.PostTag
	(*PostTagModel)(nil), // 1: forum_protos.PostTagModel
	(*ByPost)(nil),       // 2: forum_protos.ByPost
	(*PostTags)(nil),     // 3: forum_protos.PostTags
	(*ById)(nil),         // 4: forum_protos.ById
	(*Void)(nil),         // 5: forum_protos.Void
}
var file_posts_tags_proto_depIdxs = []int32{
	0, // 0: forum_protos.PostTags.PostTags:type_name -> forum_protos.PostTag
	2, // 1: forum_protos.PostTagService.GetPostTags:input_type -> forum_protos.ByPost
	0, // 2: forum_protos.PostTagService.CreatePostTag:input_type -> forum_protos.PostTag
	4, // 3: forum_protos.PostTagService.GetPostTag:input_type -> forum_protos.ById
	4, // 4: forum_protos.PostTagService.DeletePostTagById:input_type -> forum_protos.ById
	0, // 5: forum_protos.PostTagService.UpdatePostTag:input_type -> forum_protos.PostTag
	3, // 6: forum_protos.PostTagService.GetPostTags:output_type -> forum_protos.PostTags
	0, // 7: forum_protos.PostTagService.CreatePostTag:output_type -> forum_protos.PostTag
	0, // 8: forum_protos.PostTagService.GetPostTag:output_type -> forum_protos.PostTag
	5, // 9: forum_protos.PostTagService.DeletePostTagById:output_type -> forum_protos.Void
	0, // 10: forum_protos.PostTagService.UpdatePostTag:output_type -> forum_protos.PostTag
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_posts_tags_proto_init() }
func file_posts_tags_proto_init() {
	if File_posts_tags_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_posts_tags_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostTag); i {
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
		file_posts_tags_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostTagModel); i {
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
		file_posts_tags_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByPost); i {
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
		file_posts_tags_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostTags); i {
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
			RawDescriptor: file_posts_tags_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_posts_tags_proto_goTypes,
		DependencyIndexes: file_posts_tags_proto_depIdxs,
		MessageInfos:      file_posts_tags_proto_msgTypes,
	}.Build()
	File_posts_tags_proto = out.File
	file_posts_tags_proto_rawDesc = nil
	file_posts_tags_proto_goTypes = nil
	file_posts_tags_proto_depIdxs = nil
}
