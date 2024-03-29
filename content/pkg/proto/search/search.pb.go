// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: search.proto

package search

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

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           uint64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title        string  `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Description  string  `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Rating       float64 `protobuf:"fixed64,4,opt,name=Rating,proto3" json:"Rating,omitempty"`
	SumRatings   float64 `protobuf:"fixed64,5,opt,name=SumRatings,proto3" json:"SumRatings,omitempty"`
	CountRatings uint64  `protobuf:"varint,6,opt,name=CountRatings,proto3" json:"CountRatings,omitempty"`
	Year         int32   `protobuf:"varint,7,opt,name=Year,proto3" json:"Year,omitempty"`
	IsFree       bool    `protobuf:"varint,8,opt,name=IsFree,proto3" json:"IsFree,omitempty"`
	AgeLimit     int32   `protobuf:"varint,9,opt,name=AgeLimit,proto3" json:"AgeLimit,omitempty"`
	TrailerURL   string  `protobuf:"bytes,10,opt,name=TrailerURL,proto3" json:"TrailerURL,omitempty"`
	PreviewURL   string  `protobuf:"bytes,11,opt,name=PreviewURL,proto3" json:"PreviewURL,omitempty"`
	Type         string  `protobuf:"bytes,12,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{0}
}

func (x *Content) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Content) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Content) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Content) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Content) GetSumRatings() float64 {
	if x != nil {
		return x.SumRatings
	}
	return 0
}

func (x *Content) GetCountRatings() uint64 {
	if x != nil {
		return x.CountRatings
	}
	return 0
}

func (x *Content) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Content) GetIsFree() bool {
	if x != nil {
		return x.IsFree
	}
	return false
}

func (x *Content) GetAgeLimit() int32 {
	if x != nil {
		return x.AgeLimit
	}
	return 0
}

func (x *Content) GetTrailerURL() string {
	if x != nil {
		return x.TrailerURL
	}
	return ""
}

func (x *Content) GetPreviewURL() string {
	if x != nil {
		return x.PreviewURL
	}
	return ""
}

func (x *Content) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         uint64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name       string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Gender     string  `protobuf:"bytes,3,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Growth     *int32  `protobuf:"varint,4,opt,name=Growth,proto3,oneof" json:"Growth,omitempty"`
	Birthplace *string `protobuf:"bytes,5,opt,name=Birthplace,proto3,oneof" json:"Birthplace,omitempty"`
	AvatarURL  string  `protobuf:"bytes,6,opt,name=AvatarURL,proto3" json:"AvatarURL,omitempty"`
	Age        int32   `protobuf:"varint,7,opt,name=Age,proto3" json:"Age,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{1}
}

func (x *Person) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Person) GetGrowth() int32 {
	if x != nil && x.Growth != nil {
		return *x.Growth
	}
	return 0
}

func (x *Person) GetBirthplace() string {
	if x != nil && x.Birthplace != nil {
		return *x.Birthplace
	}
	return ""
}

func (x *Person) GetAvatarURL() string {
	if x != nil {
		return x.AvatarURL
	}
	return ""
}

func (x *Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type SearchContent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total   uint32     `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	Content []*Content `protobuf:"bytes,2,rep,name=Content,proto3" json:"Content,omitempty"`
}

func (x *SearchContent) Reset() {
	*x = SearchContent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchContent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchContent) ProtoMessage() {}

func (x *SearchContent) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchContent.ProtoReflect.Descriptor instead.
func (*SearchContent) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{2}
}

func (x *SearchContent) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SearchContent) GetContent() []*Content {
	if x != nil {
		return x.Content
	}
	return nil
}

type SearchPerson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total   uint32    `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	Persons []*Person `protobuf:"bytes,2,rep,name=Persons,proto3" json:"Persons,omitempty"`
}

func (x *SearchPerson) Reset() {
	*x = SearchPerson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchPerson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchPerson) ProtoMessage() {}

func (x *SearchPerson) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchPerson.ProtoReflect.Descriptor instead.
func (*SearchPerson) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{3}
}

func (x *SearchPerson) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *SearchPerson) GetPersons() []*Person {
	if x != nil {
		return x.Persons
	}
	return nil
}

type SearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content *SearchContent `protobuf:"bytes,1,opt,name=Content,proto3" json:"Content,omitempty"`
	Persons *SearchPerson  `protobuf:"bytes,2,opt,name=Persons,proto3" json:"Persons,omitempty"`
}

func (x *SearchResponse) Reset() {
	*x = SearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchResponse) ProtoMessage() {}

func (x *SearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchResponse.ProtoReflect.Descriptor instead.
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{4}
}

func (x *SearchResponse) GetContent() *SearchContent {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *SearchResponse) GetPersons() *SearchPerson {
	if x != nil {
		return x.Persons
	}
	return nil
}

type SearchParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query      string `protobuf:"bytes,1,opt,name=Query,proto3" json:"Query,omitempty"`
	TargetSlug string `protobuf:"bytes,2,opt,name=TargetSlug,proto3" json:"TargetSlug,omitempty"`
	Limit      uint32 `protobuf:"varint,3,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset     uint32 `protobuf:"varint,4,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

func (x *SearchParams) Reset() {
	*x = SearchParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_search_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchParams) ProtoMessage() {}

func (x *SearchParams) ProtoReflect() protoreflect.Message {
	mi := &file_search_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchParams.ProtoReflect.Descriptor instead.
func (*SearchParams) Descriptor() ([]byte, []int) {
	return file_search_proto_rawDescGZIP(), []int{5}
}

func (x *SearchParams) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *SearchParams) GetTargetSlug() string {
	if x != nil {
		return x.TargetSlug
	}
	return ""
}

func (x *SearchParams) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *SearchParams) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_search_proto protoreflect.FileDescriptor

var file_search_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0xc9, 0x02, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x52, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x53, 0x75, 0x6d, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x61, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x49, 0x73,
	0x46, 0x72, 0x65, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x49, 0x73, 0x46, 0x72,
	0x65, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x67, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x41, 0x67, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x12, 0x1e,
	0x0a, 0x0a, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x55, 0x52, 0x4c, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x55, 0x52, 0x4c, 0x12, 0x12,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x22, 0xd0, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x06, 0x47, 0x72, 0x6f,
	0x77, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x06, 0x47, 0x72, 0x6f,
	0x77, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a, 0x0a, 0x42, 0x69, 0x72, 0x74, 0x68, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x42, 0x69,
	0x72, 0x74, 0x68, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1c, 0x0a, 0x09, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x52, 0x4c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x52, 0x4c, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x41, 0x67, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x5f,
	0x47, 0x72, 0x6f, 0x77, 0x74, 0x68, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x42, 0x69, 0x72, 0x74, 0x68,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x22, 0x50, 0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x29, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x4e, 0x0a, 0x0c, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x28, 0x0a,
	0x07, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x07,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0x71, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x52, 0x07, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0x72, 0x0a, 0x0c, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x6c, 0x75, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x6c, 0x75, 0x67,
	0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x32, 0x49,
	0x0a, 0x0d, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x38, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x16, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x3b,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_search_proto_rawDescOnce sync.Once
	file_search_proto_rawDescData = file_search_proto_rawDesc
)

func file_search_proto_rawDescGZIP() []byte {
	file_search_proto_rawDescOnce.Do(func() {
		file_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_search_proto_rawDescData)
	})
	return file_search_proto_rawDescData
}

var file_search_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_search_proto_goTypes = []interface{}{
	(*Content)(nil),        // 0: search.Content
	(*Person)(nil),         // 1: search.Person
	(*SearchContent)(nil),  // 2: search.SearchContent
	(*SearchPerson)(nil),   // 3: search.SearchPerson
	(*SearchResponse)(nil), // 4: search.SearchResponse
	(*SearchParams)(nil),   // 5: search.SearchParams
}
var file_search_proto_depIdxs = []int32{
	0, // 0: search.SearchContent.Content:type_name -> search.Content
	1, // 1: search.SearchPerson.Persons:type_name -> search.Person
	2, // 2: search.SearchResponse.Content:type_name -> search.SearchContent
	3, // 3: search.SearchResponse.Persons:type_name -> search.SearchPerson
	5, // 4: search.SearchService.Search:input_type -> search.SearchParams
	4, // 5: search.SearchService.Search:output_type -> search.SearchResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_search_proto_init() }
func file_search_proto_init() {
	if File_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
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
		file_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_search_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchContent); i {
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
		file_search_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchPerson); i {
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
		file_search_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchResponse); i {
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
		file_search_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchParams); i {
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
	file_search_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_search_proto_goTypes,
		DependencyIndexes: file_search_proto_depIdxs,
		MessageInfos:      file_search_proto_msgTypes,
	}.Build()
	File_search_proto = out.File
	file_search_proto_rawDesc = nil
	file_search_proto_goTypes = nil
	file_search_proto_depIdxs = nil
}
