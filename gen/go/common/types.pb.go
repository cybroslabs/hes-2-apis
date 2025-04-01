// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: common/types.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The list of UUID identifiers.
type ListOfId struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Id          []string               `protobuf:"bytes,1,rep,name=id"`
	xxx_hidden_TotalCount  int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *ListOfId) Reset() {
	*x = ListOfId{}
	mi := &file_common_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOfId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOfId) ProtoMessage() {}

func (x *ListOfId) ProtoReflect() protoreflect.Message {
	mi := &file_common_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ListOfId) GetId() []string {
	if x != nil {
		return x.xxx_hidden_Id
	}
	return nil
}

func (x *ListOfId) GetTotalCount() int32 {
	if x != nil {
		return x.xxx_hidden_TotalCount
	}
	return 0
}

func (x *ListOfId) SetId(v []string) {
	x.xxx_hidden_Id = v
}

func (x *ListOfId) SetTotalCount(v int32) {
	x.xxx_hidden_TotalCount = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 2)
}

func (x *ListOfId) HasTotalCount() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *ListOfId) ClearTotalCount() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_TotalCount = 0
}

type ListOfId_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Id         []string
	TotalCount *int32
}

func (b0 ListOfId_builder) Build() *ListOfId {
	m0 := &ListOfId{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Id = b.Id
	if b.TotalCount != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 2)
		x.xxx_hidden_TotalCount = *b.TotalCount
	}
	return m0
}

// The list of string items.
type ListOfString struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Items       []string               `protobuf:"bytes,1,rep,name=items"`
	xxx_hidden_TotalCount  int32                  `protobuf:"varint,2,opt,name=total_count,json=totalCount"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *ListOfString) Reset() {
	*x = ListOfString{}
	mi := &file_common_types_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOfString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOfString) ProtoMessage() {}

func (x *ListOfString) ProtoReflect() protoreflect.Message {
	mi := &file_common_types_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ListOfString) GetItems() []string {
	if x != nil {
		return x.xxx_hidden_Items
	}
	return nil
}

func (x *ListOfString) GetTotalCount() int32 {
	if x != nil {
		return x.xxx_hidden_TotalCount
	}
	return 0
}

func (x *ListOfString) SetItems(v []string) {
	x.xxx_hidden_Items = v
}

func (x *ListOfString) SetTotalCount(v int32) {
	x.xxx_hidden_TotalCount = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 2)
}

func (x *ListOfString) HasTotalCount() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *ListOfString) ClearTotalCount() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_TotalCount = 0
}

type ListOfString_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Items      []string
	TotalCount *int32
}

func (b0 ListOfString_builder) Build() *ListOfString {
	m0 := &ListOfString{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Items = b.Items
	if b.TotalCount != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 2)
		x.xxx_hidden_TotalCount = *b.TotalCount
	}
	return m0
}

var File_common_types_proto protoreflect.FileDescriptor

const file_common_types_proto_rawDesc = "" +
	"\n" +
	"\x12common/types.proto\x12\x1dio.clbs.openhes.models.common\";\n" +
	"\bListOfId\x12\x0e\n" +
	"\x02id\x18\x01 \x03(\tR\x02id\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCount\"E\n" +
	"\fListOfString\x12\x14\n" +
	"\x05items\x18\x01 \x03(\tR\x05items\x12\x1f\n" +
	"\vtotal_count\x18\x02 \x01(\x05R\n" +
	"totalCountB0Z.github.com/cybroslabs/hes-2-apis/gen/go/commonb\beditionsp\xe8\a"

var file_common_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_types_proto_goTypes = []any{
	(*ListOfId)(nil),     // 0: io.clbs.openhes.models.common.ListOfId
	(*ListOfString)(nil), // 1: io.clbs.openhes.models.common.ListOfString
}
var file_common_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_common_types_proto_init() }
func file_common_types_proto_init() {
	if File_common_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_common_types_proto_rawDesc), len(file_common_types_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_types_proto_goTypes,
		DependencyIndexes: file_common_types_proto_depIdxs,
		MessageInfos:      file_common_types_proto_msgTypes,
	}.Build()
	File_common_types_proto = out.File
	file_common_types_proto_goTypes = nil
	file_common_types_proto_depIdxs = nil
}
