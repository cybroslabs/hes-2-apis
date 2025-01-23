// Editions version of proto3 file

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: services/driveroperator.proto

package driveroperator

import (
	acquisition "github.com/cybroslabs/hes-2-apis/gen/go/acquisition"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_services_driveroperator_proto protoreflect.FileDescriptor

var file_services_driveroperator_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x27, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf1, 0x04, 0x0a,
	0x15, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x34, 0x2e, 0x69,
	0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x62, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x34, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c,
	0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x65,
	0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x67, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x33, 0x2e, 0x69, 0x6f, 0x2e,
	0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12,
	0x63, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x63, 0x61, 0x6c,
	0x65, 0x12, 0x39, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x53, 0x63, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x69, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x39, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x5f, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12,
	0x37, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65,
	0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x79, 0x62, 0x72, 0x6f, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x68, 0x65, 0x73, 0x2d, 0x32, 0x2d,
	0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
}

var file_services_driveroperator_proto_goTypes = []any{
	(*emptypb.Empty)(nil),                     // 0: google.protobuf.Empty
	(*acquisition.NegotiateRequest)(nil),      // 1: io.clbs.openhes.models.acquisition.NegotiateRequest
	(*wrapperspb.StringValue)(nil),            // 2: google.protobuf.StringValue
	(*acquisition.SetDriverScaleRequest)(nil), // 3: io.clbs.openhes.models.acquisition.SetDriverScaleRequest
	(*acquisition.GetDriverScaleRequest)(nil), // 4: io.clbs.openhes.models.acquisition.GetDriverScaleRequest
	(*acquisition.StartUpgradeRequest)(nil),   // 5: io.clbs.openhes.models.acquisition.StartUpgradeRequest
	(*acquisition.ListOfDriverInfo)(nil),      // 6: io.clbs.openhes.models.acquisition.ListOfDriverInfo
	(*acquisition.DriverTemplates)(nil),       // 7: io.clbs.openhes.models.acquisition.DriverTemplates
	(*wrapperspb.UInt32Value)(nil),            // 8: google.protobuf.UInt32Value
}
var file_services_driveroperator_proto_depIdxs = []int32{
	0, // 0: io.clbs.openhes.services.driveroperator.DriverOperatorService.GetDrivers:input_type -> google.protobuf.Empty
	1, // 1: io.clbs.openhes.services.driveroperator.DriverOperatorService.SetDriverTemplates:input_type -> io.clbs.openhes.models.acquisition.NegotiateRequest
	2, // 2: io.clbs.openhes.services.driveroperator.DriverOperatorService.GetDriverTemplates:input_type -> google.protobuf.StringValue
	3, // 3: io.clbs.openhes.services.driveroperator.DriverOperatorService.SetDriverScale:input_type -> io.clbs.openhes.models.acquisition.SetDriverScaleRequest
	4, // 4: io.clbs.openhes.services.driveroperator.DriverOperatorService.GetDriverScale:input_type -> io.clbs.openhes.models.acquisition.GetDriverScaleRequest
	5, // 5: io.clbs.openhes.services.driveroperator.DriverOperatorService.StartUpgrade:input_type -> io.clbs.openhes.models.acquisition.StartUpgradeRequest
	6, // 6: io.clbs.openhes.services.driveroperator.DriverOperatorService.GetDrivers:output_type -> io.clbs.openhes.models.acquisition.ListOfDriverInfo
	0, // 7: io.clbs.openhes.services.driveroperator.DriverOperatorService.SetDriverTemplates:output_type -> google.protobuf.Empty
	7, // 8: io.clbs.openhes.services.driveroperator.DriverOperatorService.GetDriverTemplates:output_type -> io.clbs.openhes.models.acquisition.DriverTemplates
	0, // 9: io.clbs.openhes.services.driveroperator.DriverOperatorService.SetDriverScale:output_type -> google.protobuf.Empty
	8, // 10: io.clbs.openhes.services.driveroperator.DriverOperatorService.GetDriverScale:output_type -> google.protobuf.UInt32Value
	0, // 11: io.clbs.openhes.services.driveroperator.DriverOperatorService.StartUpgrade:output_type -> google.protobuf.Empty
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_services_driveroperator_proto_init() }
func file_services_driveroperator_proto_init() {
	if File_services_driveroperator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_driveroperator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_driveroperator_proto_goTypes,
		DependencyIndexes: file_services_driveroperator_proto_depIdxs,
	}.Build()
	File_services_driveroperator_proto = out.File
	file_services_driveroperator_proto_rawDesc = nil
	file_services_driveroperator_proto_goTypes = nil
	file_services_driveroperator_proto_depIdxs = nil
}
