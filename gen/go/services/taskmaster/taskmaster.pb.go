// Editions version of proto3 file

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: services/taskmaster/taskmaster.proto

package taskmaster

import (
	acquisition "github.com/cybroslabs/hes-2-apis/gen/go/acquisition"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_services_taskmaster_taskmaster_proto protoreflect.FileDescriptor

var file_services_taskmaster_taskmaster_proto_rawDesc = []byte{
	0x0a, 0x24, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xb4, 0x07, 0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4a, 0x6f,
	0x62, 0x73, 0x12, 0x34, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75,
	0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4a, 0x6f, 0x62,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x6f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x12, 0x31, 0x2e, 0x69, 0x6f, 0x2e,
	0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e,
	0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3b, 0x0a, 0x09, 0x50, 0x75, 0x72, 0x67, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x5b,
	0x0a, 0x0a, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x35, 0x2e, 0x69,
	0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x5b, 0x0a, 0x09, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x34, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68,
	0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x30, 0x01, 0x12, 0x5e, 0x0a, 0x0e, 0x4e, 0x65, 0x67, 0x6f,
	0x74, 0x69, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x34, 0x2e, 0x69, 0x6f, 0x2e,
	0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x57, 0x0a, 0x08, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x53, 0x65, 0x74, 0x12, 0x33, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63,
	0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x75, 0x0a, 0x08, 0x43, 0x61, 0x63, 0x68, 0x65, 0x47, 0x65, 0x74, 0x12, 0x33, 0x2e,
	0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x34, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75,
	0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x30, 0x2e,
	0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x55, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x30, 0x2e, 0x69,
	0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x79, 0x62, 0x72, 0x6f, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f,
	0x68, 0x65, 0x73, 0x2d, 0x32, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70,
	0xe8, 0x07,
}

var file_services_taskmaster_taskmaster_proto_goTypes = []any{
	(*acquisition.QueueJobsRequest)(nil),  // 0: io.clbs.openhes.models.acquisition.QueueJobsRequest
	(*acquisition.GetJobRequest)(nil),     // 1: io.clbs.openhes.models.acquisition.GetJobRequest
	(*emptypb.Empty)(nil),                 // 2: google.protobuf.Empty
	(*acquisition.CancelJobsRequest)(nil), // 3: io.clbs.openhes.models.acquisition.CancelJobsRequest
	(*acquisition.NegotiateRequest)(nil),  // 4: io.clbs.openhes.models.acquisition.NegotiateRequest
	(*acquisition.CacheSetRequest)(nil),   // 5: io.clbs.openhes.models.acquisition.CacheSetRequest
	(*acquisition.CacheGetRequest)(nil),   // 6: io.clbs.openhes.models.acquisition.CacheGetRequest
	(*acquisition.SystemConfig)(nil),      // 7: io.clbs.openhes.models.acquisition.SystemConfig
	(*acquisition.GetJobResponse)(nil),    // 8: io.clbs.openhes.models.acquisition.GetJobResponse
	(*acquisition.StreamEventsData)(nil),  // 9: io.clbs.openhes.models.acquisition.StreamEventsData
	(*acquisition.CacheGetResponse)(nil),  // 10: io.clbs.openhes.models.acquisition.CacheGetResponse
}
var file_services_taskmaster_taskmaster_proto_depIdxs = []int32{
	0,  // 0: io.clbs.openhes.services.taskmaster.TaskmasterService.QueueJobs:input_type -> io.clbs.openhes.models.acquisition.QueueJobsRequest
	1,  // 1: io.clbs.openhes.services.taskmaster.TaskmasterService.GetJob:input_type -> io.clbs.openhes.models.acquisition.GetJobRequest
	2,  // 2: io.clbs.openhes.services.taskmaster.TaskmasterService.PurgeJobs:input_type -> google.protobuf.Empty
	3,  // 3: io.clbs.openhes.services.taskmaster.TaskmasterService.CancelJobs:input_type -> io.clbs.openhes.models.acquisition.CancelJobsRequest
	2,  // 4: io.clbs.openhes.services.taskmaster.TaskmasterService.Subscribe:input_type -> google.protobuf.Empty
	4,  // 5: io.clbs.openhes.services.taskmaster.TaskmasterService.NegotiateStart:input_type -> io.clbs.openhes.models.acquisition.NegotiateRequest
	5,  // 6: io.clbs.openhes.services.taskmaster.TaskmasterService.CacheSet:input_type -> io.clbs.openhes.models.acquisition.CacheSetRequest
	6,  // 7: io.clbs.openhes.services.taskmaster.TaskmasterService.CacheGet:input_type -> io.clbs.openhes.models.acquisition.CacheGetRequest
	2,  // 8: io.clbs.openhes.services.taskmaster.TaskmasterService.GetConfig:input_type -> google.protobuf.Empty
	7,  // 9: io.clbs.openhes.services.taskmaster.TaskmasterService.SetConfig:input_type -> io.clbs.openhes.models.acquisition.SystemConfig
	2,  // 10: io.clbs.openhes.services.taskmaster.TaskmasterService.QueueJobs:output_type -> google.protobuf.Empty
	8,  // 11: io.clbs.openhes.services.taskmaster.TaskmasterService.GetJob:output_type -> io.clbs.openhes.models.acquisition.GetJobResponse
	2,  // 12: io.clbs.openhes.services.taskmaster.TaskmasterService.PurgeJobs:output_type -> google.protobuf.Empty
	2,  // 13: io.clbs.openhes.services.taskmaster.TaskmasterService.CancelJobs:output_type -> google.protobuf.Empty
	9,  // 14: io.clbs.openhes.services.taskmaster.TaskmasterService.Subscribe:output_type -> io.clbs.openhes.models.acquisition.StreamEventsData
	2,  // 15: io.clbs.openhes.services.taskmaster.TaskmasterService.NegotiateStart:output_type -> google.protobuf.Empty
	2,  // 16: io.clbs.openhes.services.taskmaster.TaskmasterService.CacheSet:output_type -> google.protobuf.Empty
	10, // 17: io.clbs.openhes.services.taskmaster.TaskmasterService.CacheGet:output_type -> io.clbs.openhes.models.acquisition.CacheGetResponse
	7,  // 18: io.clbs.openhes.services.taskmaster.TaskmasterService.GetConfig:output_type -> io.clbs.openhes.models.acquisition.SystemConfig
	2,  // 19: io.clbs.openhes.services.taskmaster.TaskmasterService.SetConfig:output_type -> google.protobuf.Empty
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_services_taskmaster_taskmaster_proto_init() }
func file_services_taskmaster_taskmaster_proto_init() {
	if File_services_taskmaster_taskmaster_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_services_taskmaster_taskmaster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_taskmaster_taskmaster_proto_goTypes,
		DependencyIndexes: file_services_taskmaster_taskmaster_proto_depIdxs,
	}.Build()
	File_services_taskmaster_taskmaster_proto = out.File
	file_services_taskmaster_taskmaster_proto_rawDesc = nil
	file_services_taskmaster_taskmaster_proto_goTypes = nil
	file_services_taskmaster_taskmaster_proto_depIdxs = nil
}
