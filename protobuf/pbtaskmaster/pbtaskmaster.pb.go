// Editions version of proto3 file

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.2
// source: pbtaskmaster.proto

package pbtaskmaster

import (
	pbdrivermodels "github.com/cybroslabs/hes-2-apis/protobuf/pbdrivermodels"
	pbtaskmastermodels "github.com/cybroslabs/hes-2-apis/protobuf/pbtaskmastermodels"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_pbtaskmaster_proto protoreflect.FileDescriptor

var file_pbtaskmaster_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x62, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x62, 0x74, 0x61, 0x73, 0x6b, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x88, 0x07, 0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x2e, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61,
	0x73, 0x74, 0x65, 0x72, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x63, 0x0a,
	0x06, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x12, 0x2b, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62,
	0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x74, 0x61, 0x73, 0x6b,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3b, 0x0a, 0x09, 0x50, 0x75, 0x72, 0x67, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12,
	0x55, 0x0a, 0x0a, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x2f, 0x2e,
	0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e,
	0x70, 0x62, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x55, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2e, 0x2e, 0x69, 0x6f,
	0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62,
	0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x44, 0x61, 0x74, 0x61, 0x30, 0x01, 0x12, 0x66, 0x0a,
	0x0e, 0x4e, 0x65, 0x67, 0x6f, 0x74, 0x69, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12,
	0x2a, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65,
	0x73, 0x2e, 0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x4e, 0x65, 0x67, 0x6f, 0x74,
	0x69, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x69, 0x6f,
	0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x08, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x65,
	0x74, 0x12, 0x2d, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x69, 0x0a, 0x08, 0x43, 0x61, 0x63, 0x68,
	0x65, 0x47, 0x65, 0x74, 0x12, 0x2d, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x43, 0x61, 0x63, 0x68, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x32, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c,
	0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x74, 0x61, 0x73,
	0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x09,
	0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2a, 0x2e, 0x69, 0x6f, 0x2e, 0x63,
	0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x74, 0x61,
	0x73, 0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x38, 0x5a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x79, 0x62, 0x72,
	0x6f, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x68, 0x65, 0x73, 0x2d, 0x32, 0x2d, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x62, 0x74, 0x61, 0x73,
	0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x70, 0xe8, 0x07,
}

var file_pbtaskmaster_proto_goTypes = []any{
	(*pbtaskmastermodels.QueueJobsRequest)(nil),     // 0: io.clbs.openhes.pbtaskmaster.QueueJobsRequest
	(*pbtaskmastermodels.GetJobRequest)(nil),        // 1: io.clbs.openhes.pbtaskmaster.GetJobRequest
	(*emptypb.Empty)(nil),                           // 2: google.protobuf.Empty
	(*pbtaskmastermodels.CancelJobsRequest)(nil),    // 3: io.clbs.openhes.pbtaskmaster.CancelJobsRequest
	(*pbdrivermodels.NegotiateRequest)(nil),         // 4: io.clbs.openhes.pbdriver.NegotiateRequest
	(*pbtaskmastermodels.CacheSetRequest)(nil),      // 5: io.clbs.openhes.pbtaskmaster.CacheSetRequest
	(*pbtaskmastermodels.CacheGetRequest)(nil),      // 6: io.clbs.openhes.pbtaskmaster.CacheGetRequest
	(*pbtaskmastermodels.SystemConfig)(nil),         // 7: io.clbs.openhes.pbtaskmaster.SystemConfig
	(*pbtaskmastermodels.GetJobResponse)(nil),       // 8: io.clbs.openhes.pbtaskmaster.GetJobResponse
	(*pbtaskmastermodels.StreamEventsData)(nil),     // 9: io.clbs.openhes.pbtaskmaster.StreamEventsData
	(*pbdrivermodels.CommonResponse)(nil),           // 10: io.clbs.openhes.pbdriver.CommonResponse
	(*pbtaskmastermodels.CacheGetResponse)(nil),     // 11: io.clbs.openhes.pbtaskmaster.CacheGetResponse
	(*pbtaskmastermodels.SystemConfigResponse)(nil), // 12: io.clbs.openhes.pbtaskmaster.SystemConfigResponse
}
var file_pbtaskmaster_proto_depIdxs = []int32{
	0,  // 0: io.clbs.openhes.pbtaskmaster.TaskmasterService.QueueJobs:input_type -> io.clbs.openhes.pbtaskmaster.QueueJobsRequest
	1,  // 1: io.clbs.openhes.pbtaskmaster.TaskmasterService.GetJob:input_type -> io.clbs.openhes.pbtaskmaster.GetJobRequest
	2,  // 2: io.clbs.openhes.pbtaskmaster.TaskmasterService.PurgeJobs:input_type -> google.protobuf.Empty
	3,  // 3: io.clbs.openhes.pbtaskmaster.TaskmasterService.CancelJobs:input_type -> io.clbs.openhes.pbtaskmaster.CancelJobsRequest
	2,  // 4: io.clbs.openhes.pbtaskmaster.TaskmasterService.Subscribe:input_type -> google.protobuf.Empty
	4,  // 5: io.clbs.openhes.pbtaskmaster.TaskmasterService.NegotiateStart:input_type -> io.clbs.openhes.pbdriver.NegotiateRequest
	5,  // 6: io.clbs.openhes.pbtaskmaster.TaskmasterService.CacheSet:input_type -> io.clbs.openhes.pbtaskmaster.CacheSetRequest
	6,  // 7: io.clbs.openhes.pbtaskmaster.TaskmasterService.CacheGet:input_type -> io.clbs.openhes.pbtaskmaster.CacheGetRequest
	2,  // 8: io.clbs.openhes.pbtaskmaster.TaskmasterService.GetConfig:input_type -> google.protobuf.Empty
	7,  // 9: io.clbs.openhes.pbtaskmaster.TaskmasterService.SetConfig:input_type -> io.clbs.openhes.pbtaskmaster.SystemConfig
	2,  // 10: io.clbs.openhes.pbtaskmaster.TaskmasterService.QueueJobs:output_type -> google.protobuf.Empty
	8,  // 11: io.clbs.openhes.pbtaskmaster.TaskmasterService.GetJob:output_type -> io.clbs.openhes.pbtaskmaster.GetJobResponse
	2,  // 12: io.clbs.openhes.pbtaskmaster.TaskmasterService.PurgeJobs:output_type -> google.protobuf.Empty
	2,  // 13: io.clbs.openhes.pbtaskmaster.TaskmasterService.CancelJobs:output_type -> google.protobuf.Empty
	9,  // 14: io.clbs.openhes.pbtaskmaster.TaskmasterService.Subscribe:output_type -> io.clbs.openhes.pbtaskmaster.StreamEventsData
	10, // 15: io.clbs.openhes.pbtaskmaster.TaskmasterService.NegotiateStart:output_type -> io.clbs.openhes.pbdriver.CommonResponse
	2,  // 16: io.clbs.openhes.pbtaskmaster.TaskmasterService.CacheSet:output_type -> google.protobuf.Empty
	11, // 17: io.clbs.openhes.pbtaskmaster.TaskmasterService.CacheGet:output_type -> io.clbs.openhes.pbtaskmaster.CacheGetResponse
	12, // 18: io.clbs.openhes.pbtaskmaster.TaskmasterService.GetConfig:output_type -> io.clbs.openhes.pbtaskmaster.SystemConfigResponse
	2,  // 19: io.clbs.openhes.pbtaskmaster.TaskmasterService.SetConfig:output_type -> google.protobuf.Empty
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_pbtaskmaster_proto_init() }
func file_pbtaskmaster_proto_init() {
	if File_pbtaskmaster_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbtaskmaster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pbtaskmaster_proto_goTypes,
		DependencyIndexes: file_pbtaskmaster_proto_depIdxs,
	}.Build()
	File_pbtaskmaster_proto = out.File
	file_pbtaskmaster_proto_rawDesc = nil
	file_pbtaskmaster_proto_goTypes = nil
	file_pbtaskmaster_proto_depIdxs = nil
}
