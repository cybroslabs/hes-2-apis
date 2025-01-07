// Editions version of proto3 file

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: pbdriveroperator.proto

package pbdriveroperator

import (
	pbdrivermodels "github.com/cybroslabs/hes-2-apis/protobuf/pbdrivermodels"
	pbdriveroperatormodels "github.com/cybroslabs/hes-2-apis/protobuf/pbdriveroperatormodels"
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

var File_pbdriveroperator_proto protoreflect.FileDescriptor

var file_pbdriveroperator_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62,
	0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xf6, 0x04,
	0x0a, 0x15, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x34, 0x2e,
	0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e,
	0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x2a, 0x2e, 0x69, 0x6f, 0x2e, 0x63,
	0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x2e, 0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x7c, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x12, 0x3b, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x29, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68,
	0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x12, 0x61, 0x0a, 0x0e, 0x53,
	0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x37, 0x2e,
	0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e,
	0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x53, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x67,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x63, 0x61, 0x6c, 0x65,
	0x12, 0x37, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68,
	0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x63, 0x61,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x5d, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x35, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62,
	0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x79, 0x62, 0x72, 0x6f, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f,
	0x68, 0x65, 0x73, 0x2d, 0x32, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8,
	0x07,
}

var file_pbdriveroperator_proto_goTypes = []any{
	(*emptypb.Empty)(nil),                                    // 0: google.protobuf.Empty
	(*pbdrivermodels.NegotiateRequest)(nil),                  // 1: io.clbs.openhes.pbdriver.NegotiateRequest
	(*pbdriveroperatormodels.GetDriverTemplatesRequest)(nil), // 2: io.clbs.openhes.pbdriveroperator.GetDriverTemplatesRequest
	(*pbdriveroperatormodels.SetDriverScaleRequest)(nil),     // 3: io.clbs.openhes.pbdriveroperator.SetDriverScaleRequest
	(*pbdriveroperatormodels.GetDriverScaleRequest)(nil),     // 4: io.clbs.openhes.pbdriveroperator.GetDriverScaleRequest
	(*pbdriveroperatormodels.StartUpgradeRequest)(nil),       // 5: io.clbs.openhes.pbdriveroperator.StartUpgradeRequest
	(*pbdriveroperatormodels.GetDriversResponse)(nil),        // 6: io.clbs.openhes.pbdriveroperator.GetDriversResponse
	(*pbdrivermodels.DriverTemplates)(nil),                   // 7: io.clbs.openhes.pbdriver.DriverTemplates
	(*wrapperspb.UInt32Value)(nil),                           // 8: google.protobuf.UInt32Value
}
var file_pbdriveroperator_proto_depIdxs = []int32{
	0, // 0: io.clbs.openhes.pbdriveroperator.DriverOperatorService.GetDrivers:input_type -> google.protobuf.Empty
	1, // 1: io.clbs.openhes.pbdriveroperator.DriverOperatorService.SetDriverTemplates:input_type -> io.clbs.openhes.pbdriver.NegotiateRequest
	2, // 2: io.clbs.openhes.pbdriveroperator.DriverOperatorService.GetDriverTemplates:input_type -> io.clbs.openhes.pbdriveroperator.GetDriverTemplatesRequest
	3, // 3: io.clbs.openhes.pbdriveroperator.DriverOperatorService.SetDriverScale:input_type -> io.clbs.openhes.pbdriveroperator.SetDriverScaleRequest
	4, // 4: io.clbs.openhes.pbdriveroperator.DriverOperatorService.GetDriverScale:input_type -> io.clbs.openhes.pbdriveroperator.GetDriverScaleRequest
	5, // 5: io.clbs.openhes.pbdriveroperator.DriverOperatorService.StartUpgrade:input_type -> io.clbs.openhes.pbdriveroperator.StartUpgradeRequest
	6, // 6: io.clbs.openhes.pbdriveroperator.DriverOperatorService.GetDrivers:output_type -> io.clbs.openhes.pbdriveroperator.GetDriversResponse
	0, // 7: io.clbs.openhes.pbdriveroperator.DriverOperatorService.SetDriverTemplates:output_type -> google.protobuf.Empty
	7, // 8: io.clbs.openhes.pbdriveroperator.DriverOperatorService.GetDriverTemplates:output_type -> io.clbs.openhes.pbdriver.DriverTemplates
	0, // 9: io.clbs.openhes.pbdriveroperator.DriverOperatorService.SetDriverScale:output_type -> google.protobuf.Empty
	8, // 10: io.clbs.openhes.pbdriveroperator.DriverOperatorService.GetDriverScale:output_type -> google.protobuf.UInt32Value
	0, // 11: io.clbs.openhes.pbdriveroperator.DriverOperatorService.StartUpgrade:output_type -> google.protobuf.Empty
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pbdriveroperator_proto_init() }
func file_pbdriveroperator_proto_init() {
	if File_pbdriveroperator_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbdriveroperator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pbdriveroperator_proto_goTypes,
		DependencyIndexes: file_pbdriveroperator_proto_depIdxs,
	}.Build()
	File_pbdriveroperator_proto = out.File
	file_pbdriveroperator_proto_rawDesc = nil
	file_pbdriveroperator_proto_goTypes = nil
	file_pbdriveroperator_proto_depIdxs = nil
}
