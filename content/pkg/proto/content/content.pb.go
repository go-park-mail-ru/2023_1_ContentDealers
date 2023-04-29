// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: content.proto

package content

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ContentID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *ContentID) Reset() {
	*x = ContentID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentID) ProtoMessage() {}

func (x *ContentID) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentID.ProtoReflect.Descriptor instead.
func (*ContentID) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{0}
}

func (x *ContentID) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type ContentIDs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentIDs []*ContentID `protobuf:"bytes,1,rep,name=ContentIDs,proto3" json:"ContentIDs,omitempty"`
}

func (x *ContentIDs) Reset() {
	*x = ContentIDs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentIDs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentIDs) ProtoMessage() {}

func (x *ContentIDs) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentIDs.ProtoReflect.Descriptor instead.
func (*ContentIDs) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{1}
}

func (x *ContentIDs) GetContentIDs() []*ContentID {
	if x != nil {
		return x.ContentIDs
	}
	return nil
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[2]
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
	return file_content_proto_rawDescGZIP(), []int{2}
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

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{3}
}

func (x *Role) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Role) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type PersonRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person *Person `protobuf:"bytes,1,opt,name=Person,proto3" json:"Person,omitempty"`
	Role   *Role   `protobuf:"bytes,2,opt,name=Role,proto3" json:"Role,omitempty"`
}

func (x *PersonRole) Reset() {
	*x = PersonRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonRole) ProtoMessage() {}

func (x *PersonRole) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonRole.ProtoReflect.Descriptor instead.
func (*PersonRole) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{4}
}

func (x *PersonRole) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

func (x *PersonRole) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

type Genre struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *Genre) Reset() {
	*x = Genre{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Genre) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Genre) ProtoMessage() {}

func (x *Genre) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Genre.ProtoReflect.Descriptor instead.
func (*Genre) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{5}
}

func (x *Genre) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Genre) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Selection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
}

func (x *Selection) Reset() {
	*x = Selection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Selection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Selection) ProtoMessage() {}

func (x *Selection) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Selection.ProtoReflect.Descriptor instead.
func (*Selection) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{6}
}

func (x *Selection) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Selection) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type Country struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *Country) Reset() {
	*x = Country{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Country) ProtoMessage() {}

func (x *Country) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Country.ProtoReflect.Descriptor instead.
func (*Country) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{7}
}

func (x *Country) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Country) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           uint64        `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Title        string        `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Description  string        `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Rating       float64       `protobuf:"fixed64,4,opt,name=Rating,proto3" json:"Rating,omitempty"`
	Year         int32         `protobuf:"varint,5,opt,name=Year,proto3" json:"Year,omitempty"`
	IsFree       bool          `protobuf:"varint,6,opt,name=IsFree,proto3" json:"IsFree,omitempty"`
	AgeLimit     int32         `protobuf:"varint,7,opt,name=AgeLimit,proto3" json:"AgeLimit,omitempty"`
	TrailerURL   string        `protobuf:"bytes,8,opt,name=TrailerURL,proto3" json:"TrailerURL,omitempty"`
	PreviewURL   string        `protobuf:"bytes,9,opt,name=PreviewURL,proto3" json:"PreviewURL,omitempty"`
	Type         string        `protobuf:"bytes,10,opt,name=Type,proto3" json:"Type,omitempty"`
	PersonsRoles []*PersonRole `protobuf:"bytes,11,rep,name=PersonsRoles,proto3" json:"PersonsRoles,omitempty"`
	Genres       []*Genre      `protobuf:"bytes,12,rep,name=Genres,proto3" json:"Genres,omitempty"`
	Selections   []*Selection  `protobuf:"bytes,13,rep,name=Selections,proto3" json:"Selections,omitempty"`
	Countries    []*Country    `protobuf:"bytes,14,rep,name=Countries,proto3" json:"Countries,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[8]
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
	return file_content_proto_rawDescGZIP(), []int{8}
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

func (x *Content) GetPersonsRoles() []*PersonRole {
	if x != nil {
		return x.PersonsRoles
	}
	return nil
}

func (x *Content) GetGenres() []*Genre {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *Content) GetSelections() []*Selection {
	if x != nil {
		return x.Selections
	}
	return nil
}

func (x *Content) GetCountries() []*Country {
	if x != nil {
		return x.Countries
	}
	return nil
}

type ContentSeq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content []*Content `protobuf:"bytes,1,rep,name=Content,proto3" json:"Content,omitempty"`
}

func (x *ContentSeq) Reset() {
	*x = ContentSeq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentSeq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentSeq) ProtoMessage() {}

func (x *ContentSeq) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentSeq.ProtoReflect.Descriptor instead.
func (*ContentSeq) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{9}
}

func (x *ContentSeq) GetContent() []*Content {
	if x != nil {
		return x.Content
	}
	return nil
}

type Film struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ContentURL string   `protobuf:"bytes,2,opt,name=ContentURL,proto3" json:"ContentURL,omitempty"`
	Content    *Content `protobuf:"bytes,3,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *Film) Reset() {
	*x = Film{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Film) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Film) ProtoMessage() {}

func (x *Film) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Film.ProtoReflect.Descriptor instead.
func (*Film) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{10}
}

func (x *Film) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Film) GetContentURL() string {
	if x != nil {
		return x.ContentURL
	}
	return ""
}

func (x *Film) GetContent() *Content {
	if x != nil {
		return x.Content
	}
	return nil
}

type Episode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          uint64                 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	SeasonNum   uint32                 `protobuf:"varint,2,opt,name=SeasonNum,proto3" json:"SeasonNum,omitempty"`
	EpisodeNum  uint32                 `protobuf:"varint,3,opt,name=EpisodeNum,proto3" json:"EpisodeNum,omitempty"`
	ContentURL  string                 `protobuf:"bytes,4,opt,name=ContentURL,proto3" json:"ContentURL,omitempty"`
	Title       string                 `protobuf:"bytes,5,opt,name=Title,proto3" json:"Title,omitempty"`
	ReleaseDate *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=ReleaseDate,proto3" json:"ReleaseDate,omitempty"`
}

func (x *Episode) Reset() {
	*x = Episode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Episode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Episode) ProtoMessage() {}

func (x *Episode) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Episode.ProtoReflect.Descriptor instead.
func (*Episode) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{11}
}

func (x *Episode) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Episode) GetSeasonNum() uint32 {
	if x != nil {
		return x.SeasonNum
	}
	return 0
}

func (x *Episode) GetEpisodeNum() uint32 {
	if x != nil {
		return x.EpisodeNum
	}
	return 0
}

func (x *Episode) GetContentURL() string {
	if x != nil {
		return x.ContentURL
	}
	return ""
}

func (x *Episode) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Episode) GetReleaseDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ReleaseDate
	}
	return nil
}

type Series struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      uint64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Content *Content `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (x *Series) Reset() {
	*x = Series{}
	if protoimpl.UnsafeEnabled {
		mi := &file_content_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Series) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Series) ProtoMessage() {}

func (x *Series) ProtoReflect() protoreflect.Message {
	mi := &file_content_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Series.ProtoReflect.Descriptor instead.
func (*Series) Descriptor() ([]byte, []int) {
	return file_content_proto_rawDescGZIP(), []int{12}
}

func (x *Series) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Series) GetContent() *Content {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_content_proto protoreflect.FileDescriptor

var file_content_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1b, 0x0a, 0x09, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x40, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x49, 0x44, 0x73, 0x12, 0x32, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x52, 0x0a, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x73, 0x22, 0x2c, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x22, 0x58, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x6f,
	0x6c, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x04, 0x52,
	0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x2b,
	0x0a, 0x05, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x09, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x2d,
	0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xca, 0x03,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x49, 0x73, 0x46, 0x72, 0x65, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x49,
	0x73, 0x46, 0x72, 0x65, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x67, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x41, 0x67, 0x65, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x55, 0x52, 0x4c, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x55, 0x52,
	0x4c, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x55, 0x52, 0x4c, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x55, 0x52,
	0x4c, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73,
	0x52, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x6f, 0x6c, 0x65,
	0x52, 0x0c, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x26,
	0x0a, 0x06, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x6e, 0x72, 0x65, 0x52, 0x06,
	0x47, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x0a, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a, 0x09, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x09, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x38, 0x0a, 0x0a, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x71, 0x12, 0x2a, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x22, 0x62, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x55, 0x52, 0x4c, 0x12, 0x2a, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0xcb, 0x01, 0x0a, 0x07, 0x45, 0x70, 0x69,
	0x73, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4e, 0x75,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x53, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x4e,
	0x75, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x4e, 0x75, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x4e,
	0x75, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x55, 0x52, 0x4c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x55,
	0x52, 0x4c, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x44, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x2a, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x52, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0xd0, 0x01, 0x0a,
	0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x39, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x6d, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x1a, 0x0d, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x2e, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x16, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x73, 0x12, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x73, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x71, 0x22, 0x00, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_content_proto_rawDescOnce sync.Once
	file_content_proto_rawDescData = file_content_proto_rawDesc
)

func file_content_proto_rawDescGZIP() []byte {
	file_content_proto_rawDescOnce.Do(func() {
		file_content_proto_rawDescData = protoimpl.X.CompressGZIP(file_content_proto_rawDescData)
	})
	return file_content_proto_rawDescData
}

var file_content_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_content_proto_goTypes = []interface{}{
	(*ContentID)(nil),             // 0: content.ContentID
	(*ContentIDs)(nil),            // 1: content.ContentIDs
	(*Person)(nil),                // 2: content.Person
	(*Role)(nil),                  // 3: content.Role
	(*PersonRole)(nil),            // 4: content.PersonRole
	(*Genre)(nil),                 // 5: content.Genre
	(*Selection)(nil),             // 6: content.Selection
	(*Country)(nil),               // 7: content.Country
	(*Content)(nil),               // 8: content.Content
	(*ContentSeq)(nil),            // 9: content.ContentSeq
	(*Film)(nil),                  // 10: content.Film
	(*Episode)(nil),               // 11: content.Episode
	(*Series)(nil),                // 12: content.Series
	(*timestamppb.Timestamp)(nil), // 13: google.protobuf.Timestamp
}
var file_content_proto_depIdxs = []int32{
	0,  // 0: content.ContentIDs.ContentIDs:type_name -> content.ContentID
	2,  // 1: content.PersonRole.Person:type_name -> content.Person
	3,  // 2: content.PersonRole.Role:type_name -> content.Role
	4,  // 3: content.Content.PersonsRoles:type_name -> content.PersonRole
	5,  // 4: content.Content.Genres:type_name -> content.Genre
	6,  // 5: content.Content.Selections:type_name -> content.Selection
	7,  // 6: content.Content.Countries:type_name -> content.Country
	8,  // 7: content.ContentSeq.Content:type_name -> content.Content
	8,  // 8: content.Film.Content:type_name -> content.Content
	13, // 9: content.Episode.ReleaseDate:type_name -> google.protobuf.Timestamp
	8,  // 10: content.Series.Content:type_name -> content.Content
	0,  // 11: content.ContentService.GetFilmByContentID:input_type -> content.ContentID
	0,  // 12: content.ContentService.GetSeriesByContentID:input_type -> content.ContentID
	1,  // 13: content.ContentService.GetContentByContentIDs:input_type -> content.ContentIDs
	10, // 14: content.ContentService.GetFilmByContentID:output_type -> content.Film
	12, // 15: content.ContentService.GetSeriesByContentID:output_type -> content.Series
	9,  // 16: content.ContentService.GetContentByContentIDs:output_type -> content.ContentSeq
	14, // [14:17] is the sub-list for method output_type
	11, // [11:14] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_content_proto_init() }
func file_content_proto_init() {
	if File_content_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_content_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentID); i {
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
		file_content_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentIDs); i {
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
		file_content_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_content_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_content_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonRole); i {
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
		file_content_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Genre); i {
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
		file_content_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Selection); i {
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
		file_content_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Country); i {
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
		file_content_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_content_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentSeq); i {
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
		file_content_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Film); i {
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
		file_content_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Episode); i {
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
		file_content_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Series); i {
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
			RawDescriptor: file_content_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_content_proto_goTypes,
		DependencyIndexes: file_content_proto_depIdxs,
		MessageInfos:      file_content_proto_msgTypes,
	}.Build()
	File_content_proto = out.File
	file_content_proto_rawDesc = nil
	file_content_proto_goTypes = nil
	file_content_proto_depIdxs = nil
}
