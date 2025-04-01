// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: services/svcdriveroperator/driveroperator.proto

package svcdriveroperator

import (
	acquisition "github.com/cybroslabs/hes-2-apis/gen/go/acquisition"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

var File_services_svcdriveroperator_driveroperator_proto protoreflect.FileDescriptor

const file_services_svcdriveroperator_driveroperator_proto_rawDesc = "" +
	"\n" +
	"/services/svcdriveroperator/driveroperator.proto\x12*io.clbs.openhes.services.svcdriveroperator\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x1aacquisition/internal.proto\x1a\x16acquisition/main.proto2\xc9\x04\n" +
	"\x15DriverOperatorService\x12W\n" +
	"\vListDrivers\x12\x16.google.protobuf.Empty\x1a0.io.clbs.openhes.models.acquisition.ListOfDriver\x12O\n" +
	"\tSetDriver\x12*.io.clbs.openhes.models.acquisition.Driver\x1a\x16.google.protobuf.Empty\x12U\n" +
	"\tGetDriver\x12\x1c.google.protobuf.StringValue\x1a*.io.clbs.openhes.models.acquisition.Driver\x12c\n" +
	"\x0eSetDriverScale\x129.io.clbs.openhes.models.acquisition.SetDriverScaleRequest\x1a\x16.google.protobuf.Empty\x12i\n" +
	"\x0eGetDriverScale\x129.io.clbs.openhes.models.acquisition.GetDriverScaleRequest\x1a\x1c.google.protobuf.UInt32Value\x12_\n" +
	"\fStartUpgrade\x127.io.clbs.openhes.models.acquisition.StartUpgradeRequest\x1a\x16.google.protobuf.EmptyBDZBgithub.com/cybroslabs/hes-2-apis/gen/go/services/svcdriveroperatorb\beditionsp\xe8\a"

var file_services_svcdriveroperator_driveroperator_proto_goTypes = []any{
	(*emptypb.Empty)(nil),                     // 0: google.protobuf.Empty
	(*acquisition.Driver)(nil),                // 1: io.clbs.openhes.models.acquisition.Driver
	(*wrapperspb.StringValue)(nil),            // 2: google.protobuf.StringValue
	(*acquisition.SetDriverScaleRequest)(nil), // 3: io.clbs.openhes.models.acquisition.SetDriverScaleRequest
	(*acquisition.GetDriverScaleRequest)(nil), // 4: io.clbs.openhes.models.acquisition.GetDriverScaleRequest
	(*acquisition.StartUpgradeRequest)(nil),   // 5: io.clbs.openhes.models.acquisition.StartUpgradeRequest
	(*acquisition.ListOfDriver)(nil),          // 6: io.clbs.openhes.models.acquisition.ListOfDriver
	(*wrapperspb.UInt32Value)(nil),            // 7: google.protobuf.UInt32Value
}
var file_services_svcdriveroperator_driveroperator_proto_depIdxs = []int32{
	0, // 0: io.clbs.openhes.services.svcdriveroperator.DriverOperatorService.ListDrivers:input_type -> google.protobuf.Empty
	1, // 1: io.clbs.openhes.services.svcdriveroperator.DriverOperatorService.SetDriver:input_type -> io.clbs.openhes.models.acquisition.Driver
	2, // 2: io.clbs.openhes.services.svcdriveroperator.DriverOperatorService.GetDriver:input_type -> google.protobuf.StringValue
	3, // 3: io.clbs.openhes.services.svcdriveroperator.DriverOperatorService.SetDriverScale:input_type -> io.clbs.openhes.models.acquisition.SetDriverScaleRequest
	4, // 4: io.clbs.openhes.services.svcdriveroperator.DriverOperatorService.GetDriverScale:input_type -> io.clbs.openhes.models.acquisition.GetDriverScaleRequest
	5, // 5: io.clbs.openhes.services.svcdriveroperator.DriverOperatorService.StartUpgrade:input_type -> io.clbs.openhes.models.acquisition.StartUpgradeRequest
	6, // 6: io.clbs.openhes.services.svcdriveroperator.DriverOperatorService.ListDrivers:output_type -> io.clbs.openhes.models.acquisition.ListOfDriver
	0, // 7: io.clbs.openhes.services.svcdriveroperator.DriverOperatorService.SetDriver:output_type -> google.protobuf.Empty
	1, // 8: io.clbs.openhes.services.svcdriveroperator.DriverOperatorService.GetDriver:output_type -> io.clbs.openhes.models.acquisition.Driver
	0, // 9: io.clbs.openhes.services.svcdriveroperator.DriverOperatorService.SetDriverScale:output_type -> google.protobuf.Empty
	7, // 10: io.clbs.openhes.services.svcdriveroperator.DriverOperatorService.GetDriverScale:output_type -> google.protobuf.UInt32Value
	0, // 11: io.clbs.openhes.services.svcdriveroperator.DriverOperatorService.StartUpgrade:output_type -> google.protobuf.Empty
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_svcdriveroperator_driveroperator_proto_init() }
func file_services_svcdriveroperator_driveroperator_proto_init() {
	if File_services_svcdriveroperator_driveroperator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_svcdriveroperator_driveroperator_proto_rawDesc), len(file_services_svcdriveroperator_driveroperator_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_svcdriveroperator_driveroperator_proto_goTypes,
		DependencyIndexes: file_services_svcdriveroperator_driveroperator_proto_depIdxs,
	}.Build()
	File_services_svcdriveroperator_driveroperator_proto = out.File
	file_services_svcdriveroperator_driveroperator_proto_goTypes = nil
	file_services_svcdriveroperator_driveroperator_proto_depIdxs = nil
}
