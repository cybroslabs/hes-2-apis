// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: common/metadata.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The metadata fields managed by user and system.
type MetadataFields struct {
	state                    protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Id            *string                `protobuf:"bytes,1,opt,name=id"`
	xxx_hidden_Generation    int32                  `protobuf:"varint,2,opt,name=generation"`
	xxx_hidden_Fields        map[string]*anypb.Any  `protobuf:"bytes,3,rep,name=fields" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	xxx_hidden_ManagedFields map[string]*anypb.Any  `protobuf:"bytes,4,rep,name=managed_fields,json=managedFields" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_raceDetectHookData   protoimpl.RaceDetectHookData
	XXX_presence             [1]uint32
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *MetadataFields) Reset() {
	*x = MetadataFields{}
	mi := &file_common_metadata_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetadataFields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataFields) ProtoMessage() {}

func (x *MetadataFields) ProtoReflect() protoreflect.Message {
	mi := &file_common_metadata_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MetadataFields) GetId() string {
	if x != nil {
		if x.xxx_hidden_Id != nil {
			return *x.xxx_hidden_Id
		}
		return ""
	}
	return ""
}

func (x *MetadataFields) GetGeneration() int32 {
	if x != nil {
		return x.xxx_hidden_Generation
	}
	return 0
}

func (x *MetadataFields) GetFields() map[string]*anypb.Any {
	if x != nil {
		return x.xxx_hidden_Fields
	}
	return nil
}

func (x *MetadataFields) GetManagedFields() map[string]*anypb.Any {
	if x != nil {
		return x.xxx_hidden_ManagedFields
	}
	return nil
}

func (x *MetadataFields) SetId(v string) {
	x.xxx_hidden_Id = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 4)
}

func (x *MetadataFields) SetGeneration(v int32) {
	x.xxx_hidden_Generation = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 4)
}

func (x *MetadataFields) SetFields(v map[string]*anypb.Any) {
	x.xxx_hidden_Fields = v
}

func (x *MetadataFields) SetManagedFields(v map[string]*anypb.Any) {
	x.xxx_hidden_ManagedFields = v
}

func (x *MetadataFields) HasId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *MetadataFields) HasGeneration() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *MetadataFields) ClearId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Id = nil
}

func (x *MetadataFields) ClearGeneration() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_Generation = 0
}

type MetadataFields_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Id            *string
	Generation    *int32
	Fields        map[string]*anypb.Any
	ManagedFields map[string]*anypb.Any
}

func (b0 MetadataFields_builder) Build() *MetadataFields {
	m0 := &MetadataFields{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Id != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 4)
		x.xxx_hidden_Id = b.Id
	}
	if b.Generation != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 4)
		x.xxx_hidden_Generation = *b.Generation
	}
	x.xxx_hidden_Fields = b.Fields
	x.xxx_hidden_ManagedFields = b.ManagedFields
	return m0
}

var File_common_metadata_proto protoreflect.FileDescriptor

var file_common_metadata_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa5, 0x03, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x51, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x67, 0x0a, 0x0e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x40, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65,
	0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x4d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x1a, 0x4f, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x56, 0x0a, 0x12, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2a, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x79, 0x62, 0x72, 0x6f, 0x73, 0x6c, 0x61,
	0x62, 0x73, 0x2f, 0x68, 0x65, 0x73, 0x2d, 0x32, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x08, 0x65, 0x64, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
}

var file_common_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_common_metadata_proto_goTypes = []any{
	(*MetadataFields)(nil), // 0: io.clbs.openhes.models.common.MetadataFields
	nil,                    // 1: io.clbs.openhes.models.common.MetadataFields.FieldsEntry
	nil,                    // 2: io.clbs.openhes.models.common.MetadataFields.ManagedFieldsEntry
	(*anypb.Any)(nil),      // 3: google.protobuf.Any
}
var file_common_metadata_proto_depIdxs = []int32{
	1, // 0: io.clbs.openhes.models.common.MetadataFields.fields:type_name -> io.clbs.openhes.models.common.MetadataFields.FieldsEntry
	2, // 1: io.clbs.openhes.models.common.MetadataFields.managed_fields:type_name -> io.clbs.openhes.models.common.MetadataFields.ManagedFieldsEntry
	3, // 2: io.clbs.openhes.models.common.MetadataFields.FieldsEntry.value:type_name -> google.protobuf.Any
	3, // 3: io.clbs.openhes.models.common.MetadataFields.ManagedFieldsEntry.value:type_name -> google.protobuf.Any
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_common_metadata_proto_init() }
func file_common_metadata_proto_init() {
	if File_common_metadata_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_metadata_proto_goTypes,
		DependencyIndexes: file_common_metadata_proto_depIdxs,
		MessageInfos:      file_common_metadata_proto_msgTypes,
	}.Build()
	File_common_metadata_proto = out.File
	file_common_metadata_proto_rawDesc = nil
	file_common_metadata_proto_goTypes = nil
	file_common_metadata_proto_depIdxs = nil
}
