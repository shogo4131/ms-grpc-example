// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2-devel
// 	protoc        v5.27.3
// source: catalogue.proto

package book

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

// idを含むリクエスト
type GetBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBookRequest) Reset() {
	*x = GetBookRequest{}
	mi := &file_catalogue_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookRequest) ProtoMessage() {}

func (x *GetBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_catalogue_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookRequest.ProtoReflect.Descriptor instead.
func (*GetBookRequest) Descriptor() ([]byte, []int) {
	return file_catalogue_proto_rawDescGZIP(), []int{0}
}

func (x *GetBookRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// bookを含むレスポンス
type GetBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book *Book `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *GetBookResponse) Reset() {
	*x = GetBookResponse{}
	mi := &file_catalogue_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookResponse) ProtoMessage() {}

func (x *GetBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_catalogue_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookResponse.ProtoReflect.Descriptor instead.
func (*GetBookResponse) Descriptor() ([]byte, []int) {
	return file_catalogue_proto_rawDescGZIP(), []int{1}
}

func (x *GetBookResponse) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

// 書籍情報Bookの定義
type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Author string `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Price  int32  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	mi := &file_catalogue_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_catalogue_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_catalogue_proto_rawDescGZIP(), []int{2}
}

func (x *Book) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Book) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_catalogue_proto protoreflect.FileDescriptor

var file_catalogue_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x31, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04,
	0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6b, 0x22, 0x5a, 0x0a, 0x04,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x32, 0x45, 0x0a, 0x09, 0x43, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x75, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x12, 0x14, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x1c, 0x5a, 0x1a, 0x67, 0x69, 0x68, 0x79, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x75, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_catalogue_proto_rawDescOnce sync.Once
	file_catalogue_proto_rawDescData = file_catalogue_proto_rawDesc
)

func file_catalogue_proto_rawDescGZIP() []byte {
	file_catalogue_proto_rawDescOnce.Do(func() {
		file_catalogue_proto_rawDescData = protoimpl.X.CompressGZIP(file_catalogue_proto_rawDescData)
	})
	return file_catalogue_proto_rawDescData
}

var file_catalogue_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_catalogue_proto_goTypes = []any{
	(*GetBookRequest)(nil),  // 0: book.GetBookRequest
	(*GetBookResponse)(nil), // 1: book.GetBookResponse
	(*Book)(nil),            // 2: book.Book
}
var file_catalogue_proto_depIdxs = []int32{
	2, // 0: book.GetBookResponse.book:type_name -> book.Book
	0, // 1: book.Catalogue.GetBook:input_type -> book.GetBookRequest
	1, // 2: book.Catalogue.GetBook:output_type -> book.GetBookResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_catalogue_proto_init() }
func file_catalogue_proto_init() {
	if File_catalogue_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_catalogue_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_catalogue_proto_goTypes,
		DependencyIndexes: file_catalogue_proto_depIdxs,
		MessageInfos:      file_catalogue_proto_msgTypes,
	}.Build()
	File_catalogue_proto = out.File
	file_catalogue_proto_rawDesc = nil
	file_catalogue_proto_goTypes = nil
	file_catalogue_proto_depIdxs = nil
}
