// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: favorites.proto

package content

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Favorite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    uint64               `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	ContentID uint64               `protobuf:"varint,2,opt,name=ContentID,proto3" json:"ContentID,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // string DateAdding = 3;
}

func (x *Favorite) Reset() {
	*x = Favorite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorites_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Favorite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Favorite) ProtoMessage() {}

func (x *Favorite) ProtoReflect() protoreflect.Message {
	mi := &file_favorites_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Favorite.ProtoReflect.Descriptor instead.
func (*Favorite) Descriptor() ([]byte, []int) {
	return file_favorites_proto_rawDescGZIP(), []int{0}
}

func (x *Favorite) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Favorite) GetContentID() uint64 {
	if x != nil {
		return x.ContentID
	}
	return 0
}

func (x *Favorite) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type FavoritesOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   uint64 `protobuf:"varint,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	SortDate string `protobuf:"bytes,2,opt,name=SortDate,proto3" json:"SortDate,omitempty"`
	Limit    uint32 `protobuf:"varint,3,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset   uint32 `protobuf:"varint,4,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

func (x *FavoritesOptions) Reset() {
	*x = FavoritesOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorites_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FavoritesOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FavoritesOptions) ProtoMessage() {}

func (x *FavoritesOptions) ProtoReflect() protoreflect.Message {
	mi := &file_favorites_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FavoritesOptions.ProtoReflect.Descriptor instead.
func (*FavoritesOptions) Descriptor() ([]byte, []int) {
	return file_favorites_proto_rawDescGZIP(), []int{1}
}

func (x *FavoritesOptions) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *FavoritesOptions) GetSortDate() string {
	if x != nil {
		return x.SortDate
	}
	return ""
}

func (x *FavoritesOptions) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FavoritesOptions) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type Favorites struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsLast    bool        `protobuf:"varint,1,opt,name=IsLast,proto3" json:"IsLast,omitempty"`
	Favorites []*Favorite `protobuf:"bytes,2,rep,name=Favorites,proto3" json:"Favorites,omitempty"`
}

func (x *Favorites) Reset() {
	*x = Favorites{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorites_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Favorites) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Favorites) ProtoMessage() {}

func (x *Favorites) ProtoReflect() protoreflect.Message {
	mi := &file_favorites_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Favorites.ProtoReflect.Descriptor instead.
func (*Favorites) Descriptor() ([]byte, []int) {
	return file_favorites_proto_rawDescGZIP(), []int{2}
}

func (x *Favorites) GetIsLast() bool {
	if x != nil {
		return x.IsLast
	}
	return false
}

func (x *Favorites) GetFavorites() []*Favorite {
	if x != nil {
		return x.Favorites
	}
	return nil
}

type Nothing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dummy bool `protobuf:"varint,1,opt,name=Dummy,proto3" json:"Dummy,omitempty"`
}

func (x *Nothing) Reset() {
	*x = Nothing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorites_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nothing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nothing) ProtoMessage() {}

func (x *Nothing) ProtoReflect() protoreflect.Message {
	mi := &file_favorites_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nothing.ProtoReflect.Descriptor instead.
func (*Nothing) Descriptor() ([]byte, []int) {
	return file_favorites_proto_rawDescGZIP(), []int{3}
}

func (x *Nothing) GetDummy() bool {
	if x != nil {
		return x.Dummy
	}
	return false
}

type HasFav struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasFav bool `protobuf:"varint,1,opt,name=HasFav,proto3" json:"HasFav,omitempty"`
}

func (x *HasFav) Reset() {
	*x = HasFav{}
	if protoimpl.UnsafeEnabled {
		mi := &file_favorites_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HasFav) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HasFav) ProtoMessage() {}

func (x *HasFav) ProtoReflect() protoreflect.Message {
	mi := &file_favorites_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HasFav.ProtoReflect.Descriptor instead.
func (*HasFav) Descriptor() ([]byte, []int) {
	return file_favorites_proto_rawDescGZIP(), []int{4}
}

func (x *HasFav) GetHasFav() bool {
	if x != nil {
		return x.HasFav
	}
	return false
}

var File_favorites_proto protoreflect.FileDescriptor

var file_favorites_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x08, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1c, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x09, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x74, 0x0a, 0x10, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x6f, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x54,
	0x0a, 0x09, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x49,
	0x73, 0x4c, 0x61, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x49, 0x73, 0x4c,
	0x61, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x09, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x52, 0x09, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x73, 0x22, 0x1f, 0x0a, 0x07, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x44, 0x75, 0x6d, 0x6d, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x44, 0x75, 0x6d, 0x6d, 0x79, 0x22, 0x20, 0x0a, 0x06, 0x48, 0x61, 0x73, 0x46, 0x61, 0x76, 0x12,
	0x16, 0x0a, 0x06, 0x48, 0x61, 0x73, 0x46, 0x61, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x48, 0x61, 0x73, 0x46, 0x61, 0x76, 0x32, 0xfc, 0x01, 0x0a, 0x17, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x46,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2e, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x0a, 0x41,
	0x64, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x1a, 0x10, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x19,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x73, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x35, 0x0a, 0x0d, 0x48, 0x61, 0x73, 0x46, 0x61, 0x76, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x12, 0x11, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x61,
	0x73, 0x46, 0x61, 0x76, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_favorites_proto_rawDescOnce sync.Once
	file_favorites_proto_rawDescData = file_favorites_proto_rawDesc
)

func file_favorites_proto_rawDescGZIP() []byte {
	file_favorites_proto_rawDescOnce.Do(func() {
		file_favorites_proto_rawDescData = protoimpl.X.CompressGZIP(file_favorites_proto_rawDescData)
	})
	return file_favorites_proto_rawDescData
}

var file_favorites_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_favorites_proto_goTypes = []interface{}{
	(*Favorite)(nil),            // 0: content.Favorite
	(*FavoritesOptions)(nil),    // 1: content.FavoritesOptions
	(*Favorites)(nil),           // 2: content.Favorites
	(*Nothing)(nil),             // 3: content.Nothing
	(*HasFav)(nil),              // 4: content.HasFav
	(*timestamp.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_favorites_proto_depIdxs = []int32{
	5, // 0: content.Favorite.created_at:type_name -> google.protobuf.Timestamp
	0, // 1: content.Favorites.Favorites:type_name -> content.Favorite
	0, // 2: content.FavoritesContentService.DeleteContent:input_type -> content.Favorite
	0, // 3: content.FavoritesContentService.AddContent:input_type -> content.Favorite
	1, // 4: content.FavoritesContentService.GetContent:input_type -> content.FavoritesOptions
	0, // 5: content.FavoritesContentService.HasFavContent:input_type -> content.Favorite
	3, // 6: content.FavoritesContentService.DeleteContent:output_type -> content.Nothing
	3, // 7: content.FavoritesContentService.AddContent:output_type -> content.Nothing
	2, // 8: content.FavoritesContentService.GetContent:output_type -> content.Favorites
	4, // 9: content.FavoritesContentService.HasFavContent:output_type -> content.HasFav
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_favorites_proto_init() }
func file_favorites_proto_init() {
	if File_favorites_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_favorites_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Favorite); i {
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
		file_favorites_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FavoritesOptions); i {
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
		file_favorites_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Favorites); i {
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
		file_favorites_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nothing); i {
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
		file_favorites_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HasFav); i {
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
			RawDescriptor: file_favorites_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_favorites_proto_goTypes,
		DependencyIndexes: file_favorites_proto_depIdxs,
		MessageInfos:      file_favorites_proto_msgTypes,
	}.Build()
	File_favorites_proto = out.File
	file_favorites_proto_rawDesc = nil
	file_favorites_proto_goTypes = nil
	file_favorites_proto_depIdxs = nil
}
