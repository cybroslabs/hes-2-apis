// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: services/dataproxy/dataproxy.proto

package dataproxy

import (
	acquisition "github.com/cybroslabs/hes-2-apis/gen/go/acquisition"
	common "github.com/cybroslabs/hes-2-apis/gen/go/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_services_dataproxy_dataproxy_proto protoreflect.FileDescriptor

var file_services_dataproxy_dataproxy_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xf6, 0x02, 0x0a, 0x10, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x75, 0x6c, 0x6b, 0x12, 0x35, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63,
	0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x75, 0x6c, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x68, 0x0a, 0x09, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x75, 0x6c, 0x6b, 0x73, 0x12, 0x2b, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x1a, 0x2e, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71,
	0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x42,
	0x75, 0x6c, 0x6b, 0x12, 0x51, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x12, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x28, 0x2e, 0x69,
	0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x12, 0x42, 0x0a, 0x0a, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x42, 0x75, 0x6c, 0x6b, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x79, 0x62, 0x72, 0x6f, 0x73, 0x6c,
	0x61, 0x62, 0x73, 0x2f, 0x68, 0x65, 0x73, 0x2d, 0x32, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x70, 0xe8, 0x07,
}

var file_services_dataproxy_dataproxy_proto_goTypes = []any{
	(*acquisition.CreateBulkRequest)(nil), // 0: io.clbs.openhes.models.acquisition.CreateBulkRequest
	(*common.ListSelector)(nil),           // 1: io.clbs.openhes.models.common.ListSelector
	(*wrapperspb.StringValue)(nil),        // 2: google.protobuf.StringValue
	(*acquisition.ListOfBulk)(nil),        // 3: io.clbs.openhes.models.acquisition.ListOfBulk
	(*acquisition.Bulk)(nil),              // 4: io.clbs.openhes.models.acquisition.Bulk
	(*emptypb.Empty)(nil),                 // 5: google.protobuf.Empty
}
var file_services_dataproxy_dataproxy_proto_depIdxs = []int32{
	0, // 0: io.clbs.openhes.services.dataproxy.DataproxyService.CreateBulk:input_type -> io.clbs.openhes.models.acquisition.CreateBulkRequest
	1, // 1: io.clbs.openhes.services.dataproxy.DataproxyService.ListBulks:input_type -> io.clbs.openhes.models.common.ListSelector
	2, // 2: io.clbs.openhes.services.dataproxy.DataproxyService.GetBulk:input_type -> google.protobuf.StringValue
	2, // 3: io.clbs.openhes.services.dataproxy.DataproxyService.CancelBulk:input_type -> google.protobuf.StringValue
	2, // 4: io.clbs.openhes.services.dataproxy.DataproxyService.CreateBulk:output_type -> google.protobuf.StringValue
	3, // 5: io.clbs.openhes.services.dataproxy.DataproxyService.ListBulks:output_type -> io.clbs.openhes.models.acquisition.ListOfBulk
	4, // 6: io.clbs.openhes.services.dataproxy.DataproxyService.GetBulk:output_type -> io.clbs.openhes.models.acquisition.Bulk
	5, // 7: io.clbs.openhes.services.dataproxy.DataproxyService.CancelBulk:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_dataproxy_dataproxy_proto_init() }
func file_services_dataproxy_dataproxy_proto_init() {
	if File_services_dataproxy_dataproxy_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_dataproxy_dataproxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_dataproxy_dataproxy_proto_goTypes,
		DependencyIndexes: file_services_dataproxy_dataproxy_proto_depIdxs,
	}.Build()
	File_services_dataproxy_dataproxy_proto = out.File
	file_services_dataproxy_dataproxy_proto_rawDesc = nil
	file_services_dataproxy_dataproxy_proto_goTypes = nil
	file_services_dataproxy_dataproxy_proto_depIdxs = nil
}
