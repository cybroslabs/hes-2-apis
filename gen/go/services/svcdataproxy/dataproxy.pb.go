// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: services/svcdataproxy/dataproxy.proto

package svcdataproxy

import (
	acquisition "github.com/cybroslabs/hes-2-apis/gen/go/acquisition"
	common "github.com/cybroslabs/hes-2-apis/gen/go/common"
	system "github.com/cybroslabs/hes-2-apis/gen/go/system"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_services_svcdataproxy_dataproxy_proto protoreflect.FileDescriptor

const file_services_svcdataproxy_dataproxy_proto_rawDesc = "" +
	"\n" +
	"%services/svcdataproxy/dataproxy.proto\x12%io.clbs.openhes.services.svcdataproxy\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x16acquisition/main.proto\x1a\x18acquisition/shared.proto\x1a\x13common/fields.proto\x1a\x11system/main.proto2\xe3\v\n" +
	"\x10DataproxyService\x12h\n" +
	"\tListBulks\x12+.io.clbs.openhes.models.common.ListSelector\x1a..io.clbs.openhes.models.acquisition.ListOfBulk\x12z\n" +
	"\fListBulkJobs\x127.io.clbs.openhes.models.acquisition.ListBulkJobsRequest\x1a1.io.clbs.openhes.models.acquisition.ListOfBulkJob\x12W\n" +
	"\n" +
	"GetBulkJob\x12\x1c.google.protobuf.StringValue\x1a+.io.clbs.openhes.models.acquisition.BulkJob\x12B\n" +
	"\n" +
	"CancelBulk\x12\x1c.google.protobuf.StringValue\x1a\x16.google.protobuf.Empty\x12k\n" +
	"\x0fCreateProxyBulk\x12:.io.clbs.openhes.models.acquisition.CreateProxyBulkRequest\x1a\x1c.google.protobuf.StringValue\x12[\n" +
	"\fGetProxyBulk\x12\x1c.google.protobuf.StringValue\x1a-.io.clbs.openhes.models.acquisition.ProxyBulk\x12a\n" +
	"\n" +
	"CreateBulk\x125.io.clbs.openhes.models.acquisition.CreateBulkRequest\x1a\x1c.google.protobuf.StringValue\x12Q\n" +
	"\aGetBulk\x12\x1c.google.protobuf.StringValue\x1a(.io.clbs.openhes.models.acquisition.Bulk\x12P\n" +
	"\tGetConfig\x12\x16.google.protobuf.Empty\x1a+.io.clbs.openhes.models.system.SystemConfig\x12P\n" +
	"\tSetConfig\x12+.io.clbs.openhes.models.system.SystemConfig\x1a\x16.google.protobuf.Empty\x12\x86\x01\n" +
	"\x15GetMeterDataRegisters\x127.io.clbs.openhes.models.acquisition.GetMeterDataRequest\x1a2.io.clbs.openhes.models.acquisition.RegisterValues0\x01\x12\x84\x01\n" +
	"\x14GetMeterDataProfiles\x127.io.clbs.openhes.models.acquisition.GetMeterDataRequest\x1a1.io.clbs.openhes.models.acquisition.ProfileValues0\x01\x12\x96\x01\n" +
	"\x1dGetMeterDataIrregularProfiles\x127.io.clbs.openhes.models.acquisition.GetMeterDataRequest\x1a:.io.clbs.openhes.models.acquisition.IrregularProfileValues0\x01\x12\x7f\n" +
	"\x0eGetMeterEvents\x129.io.clbs.openhes.models.acquisition.GetMeterEventsRequest\x1a0.io.clbs.openhes.models.acquisition.EventRecords0\x01B?Z=github.com/cybroslabs/hes-2-apis/gen/go/services/svcdataproxyb\beditionsp\xe8\a"

var file_services_svcdataproxy_dataproxy_proto_goTypes = []any{
	(*common.ListSelector)(nil),                // 0: io.clbs.openhes.models.common.ListSelector
	(*acquisition.ListBulkJobsRequest)(nil),    // 1: io.clbs.openhes.models.acquisition.ListBulkJobsRequest
	(*wrapperspb.StringValue)(nil),             // 2: google.protobuf.StringValue
	(*acquisition.CreateProxyBulkRequest)(nil), // 3: io.clbs.openhes.models.acquisition.CreateProxyBulkRequest
	(*acquisition.CreateBulkRequest)(nil),      // 4: io.clbs.openhes.models.acquisition.CreateBulkRequest
	(*emptypb.Empty)(nil),                      // 5: google.protobuf.Empty
	(*system.SystemConfig)(nil),                // 6: io.clbs.openhes.models.system.SystemConfig
	(*acquisition.GetMeterDataRequest)(nil),    // 7: io.clbs.openhes.models.acquisition.GetMeterDataRequest
	(*acquisition.GetMeterEventsRequest)(nil),  // 8: io.clbs.openhes.models.acquisition.GetMeterEventsRequest
	(*acquisition.ListOfBulk)(nil),             // 9: io.clbs.openhes.models.acquisition.ListOfBulk
	(*acquisition.ListOfBulkJob)(nil),          // 10: io.clbs.openhes.models.acquisition.ListOfBulkJob
	(*acquisition.BulkJob)(nil),                // 11: io.clbs.openhes.models.acquisition.BulkJob
	(*acquisition.ProxyBulk)(nil),              // 12: io.clbs.openhes.models.acquisition.ProxyBulk
	(*acquisition.Bulk)(nil),                   // 13: io.clbs.openhes.models.acquisition.Bulk
	(*acquisition.RegisterValues)(nil),         // 14: io.clbs.openhes.models.acquisition.RegisterValues
	(*acquisition.ProfileValues)(nil),          // 15: io.clbs.openhes.models.acquisition.ProfileValues
	(*acquisition.IrregularProfileValues)(nil), // 16: io.clbs.openhes.models.acquisition.IrregularProfileValues
	(*acquisition.EventRecords)(nil),           // 17: io.clbs.openhes.models.acquisition.EventRecords
}
var file_services_svcdataproxy_dataproxy_proto_depIdxs = []int32{
	0,  // 0: io.clbs.openhes.services.svcdataproxy.DataproxyService.ListBulks:input_type -> io.clbs.openhes.models.common.ListSelector
	1,  // 1: io.clbs.openhes.services.svcdataproxy.DataproxyService.ListBulkJobs:input_type -> io.clbs.openhes.models.acquisition.ListBulkJobsRequest
	2,  // 2: io.clbs.openhes.services.svcdataproxy.DataproxyService.GetBulkJob:input_type -> google.protobuf.StringValue
	2,  // 3: io.clbs.openhes.services.svcdataproxy.DataproxyService.CancelBulk:input_type -> google.protobuf.StringValue
	3,  // 4: io.clbs.openhes.services.svcdataproxy.DataproxyService.CreateProxyBulk:input_type -> io.clbs.openhes.models.acquisition.CreateProxyBulkRequest
	2,  // 5: io.clbs.openhes.services.svcdataproxy.DataproxyService.GetProxyBulk:input_type -> google.protobuf.StringValue
	4,  // 6: io.clbs.openhes.services.svcdataproxy.DataproxyService.CreateBulk:input_type -> io.clbs.openhes.models.acquisition.CreateBulkRequest
	2,  // 7: io.clbs.openhes.services.svcdataproxy.DataproxyService.GetBulk:input_type -> google.protobuf.StringValue
	5,  // 8: io.clbs.openhes.services.svcdataproxy.DataproxyService.GetConfig:input_type -> google.protobuf.Empty
	6,  // 9: io.clbs.openhes.services.svcdataproxy.DataproxyService.SetConfig:input_type -> io.clbs.openhes.models.system.SystemConfig
	7,  // 10: io.clbs.openhes.services.svcdataproxy.DataproxyService.GetMeterDataRegisters:input_type -> io.clbs.openhes.models.acquisition.GetMeterDataRequest
	7,  // 11: io.clbs.openhes.services.svcdataproxy.DataproxyService.GetMeterDataProfiles:input_type -> io.clbs.openhes.models.acquisition.GetMeterDataRequest
	7,  // 12: io.clbs.openhes.services.svcdataproxy.DataproxyService.GetMeterDataIrregularProfiles:input_type -> io.clbs.openhes.models.acquisition.GetMeterDataRequest
	8,  // 13: io.clbs.openhes.services.svcdataproxy.DataproxyService.GetMeterEvents:input_type -> io.clbs.openhes.models.acquisition.GetMeterEventsRequest
	9,  // 14: io.clbs.openhes.services.svcdataproxy.DataproxyService.ListBulks:output_type -> io.clbs.openhes.models.acquisition.ListOfBulk
	10, // 15: io.clbs.openhes.services.svcdataproxy.DataproxyService.ListBulkJobs:output_type -> io.clbs.openhes.models.acquisition.ListOfBulkJob
	11, // 16: io.clbs.openhes.services.svcdataproxy.DataproxyService.GetBulkJob:output_type -> io.clbs.openhes.models.acquisition.BulkJob
	5,  // 17: io.clbs.openhes.services.svcdataproxy.DataproxyService.CancelBulk:output_type -> google.protobuf.Empty
	2,  // 18: io.clbs.openhes.services.svcdataproxy.DataproxyService.CreateProxyBulk:output_type -> google.protobuf.StringValue
	12, // 19: io.clbs.openhes.services.svcdataproxy.DataproxyService.GetProxyBulk:output_type -> io.clbs.openhes.models.acquisition.ProxyBulk
	2,  // 20: io.clbs.openhes.services.svcdataproxy.DataproxyService.CreateBulk:output_type -> google.protobuf.StringValue
	13, // 21: io.clbs.openhes.services.svcdataproxy.DataproxyService.GetBulk:output_type -> io.clbs.openhes.models.acquisition.Bulk
	6,  // 22: io.clbs.openhes.services.svcdataproxy.DataproxyService.GetConfig:output_type -> io.clbs.openhes.models.system.SystemConfig
	5,  // 23: io.clbs.openhes.services.svcdataproxy.DataproxyService.SetConfig:output_type -> google.protobuf.Empty
	14, // 24: io.clbs.openhes.services.svcdataproxy.DataproxyService.GetMeterDataRegisters:output_type -> io.clbs.openhes.models.acquisition.RegisterValues
	15, // 25: io.clbs.openhes.services.svcdataproxy.DataproxyService.GetMeterDataProfiles:output_type -> io.clbs.openhes.models.acquisition.ProfileValues
	16, // 26: io.clbs.openhes.services.svcdataproxy.DataproxyService.GetMeterDataIrregularProfiles:output_type -> io.clbs.openhes.models.acquisition.IrregularProfileValues
	17, // 27: io.clbs.openhes.services.svcdataproxy.DataproxyService.GetMeterEvents:output_type -> io.clbs.openhes.models.acquisition.EventRecords
	14, // [14:28] is the sub-list for method output_type
	0,  // [0:14] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_services_svcdataproxy_dataproxy_proto_init() }
func file_services_svcdataproxy_dataproxy_proto_init() {
	if File_services_svcdataproxy_dataproxy_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_svcdataproxy_dataproxy_proto_rawDesc), len(file_services_svcdataproxy_dataproxy_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_svcdataproxy_dataproxy_proto_goTypes,
		DependencyIndexes: file_services_svcdataproxy_dataproxy_proto_depIdxs,
	}.Build()
	File_services_svcdataproxy_dataproxy_proto = out.File
	file_services_svcdataproxy_dataproxy_proto_goTypes = nil
	file_services_svcdataproxy_dataproxy_proto_depIdxs = nil
}
