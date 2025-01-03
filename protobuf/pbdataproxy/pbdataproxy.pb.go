// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: pbdataproxy.proto

package pbdataproxy

import (
	pbdriver "github.com/cybroslabs/hes-2-apis/protobuf/pbdriver"
	pbtaskmaster "github.com/cybroslabs/hes-2-apis/protobuf/pbtaskmaster"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Bulk statuses
type BulkStatusCode int32

const (
	// The job is waiting in the queue
	BulkStatusCode_BULK_STATUS_QUEUED BulkStatusCode = 0
	// The job is running
	BulkStatusCode_BULK_STATUS_RUNNING BulkStatusCode = 1
	// The job is completed
	BulkStatusCode_BULK_STATUS_COMPLETED BulkStatusCode = 2
	// The job is cancelled
	BulkStatusCode_BULK_STATUS_CANCELLED BulkStatusCode = 3
	// The job has expired
	BulkStatusCode_BULK_STATUS_EXPIRED BulkStatusCode = 4
)

// Enum value maps for BulkStatusCode.
var (
	BulkStatusCode_name = map[int32]string{
		0: "BULK_STATUS_QUEUED",
		1: "BULK_STATUS_RUNNING",
		2: "BULK_STATUS_COMPLETED",
		3: "BULK_STATUS_CANCELLED",
		4: "BULK_STATUS_EXPIRED",
	}
	BulkStatusCode_value = map[string]int32{
		"BULK_STATUS_QUEUED":    0,
		"BULK_STATUS_RUNNING":   1,
		"BULK_STATUS_COMPLETED": 2,
		"BULK_STATUS_CANCELLED": 3,
		"BULK_STATUS_EXPIRED":   4,
	}
)

func (x BulkStatusCode) Enum() *BulkStatusCode {
	p := new(BulkStatusCode)
	*p = x
	return p
}

func (x BulkStatusCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BulkStatusCode) Descriptor() protoreflect.EnumDescriptor {
	return file_pbdataproxy_proto_enumTypes[0].Descriptor()
}

func (BulkStatusCode) Type() protoreflect.EnumType {
	return &file_pbdataproxy_proto_enumTypes[0]
}

func (x BulkStatusCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BulkStatusCode.Descriptor instead.
func (BulkStatusCode) EnumDescriptor() ([]byte, []int) {
	return file_pbdataproxy_proto_rawDescGZIP(), []int{0}
}

// RestApi -> DataProxy
type CreateBulkRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Spec          *BulkSpec              `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"` // The bulk-job spec.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateBulkRequest) Reset() {
	*x = CreateBulkRequest{}
	mi := &file_pbdataproxy_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateBulkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBulkRequest) ProtoMessage() {}

func (x *CreateBulkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pbdataproxy_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBulkRequest.ProtoReflect.Descriptor instead.
func (*CreateBulkRequest) Descriptor() ([]byte, []int) {
	return file_pbdataproxy_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBulkRequest) GetSpec() *BulkSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

// Sub-message - holds the bulk job specification.
type BulkSpec struct {
	state         protoimpl.MessageState    `protogen:"open.v1"`
	BulkId        string                    `protobuf:"bytes,1,opt,name=bulk_id,json=bulkId,proto3" json:"bulk_id,omitempty"`                      // The bulk identifier.
	CorrelationId string                    `protobuf:"bytes,2,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"` // The correlation identifier, e.g. to define relation to non-homogenous group.
	OrgId         string                    `protobuf:"bytes,3,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`                         // The organization identifier.
	DriverType    string                    `protobuf:"bytes,4,opt,name=driver_type,json=driverType,proto3" json:"driver_type,omitempty"`          // The device (driver) type.
	Devices       []*pbtaskmaster.JobDevice `protobuf:"bytes,5,rep,name=devices,proto3" json:"devices,omitempty"`                                  // The list of devices in the bulk.
	Settings      *pbdriver.JobSettings     `protobuf:"bytes,6,opt,name=settings,proto3" json:"settings,omitempty"`                                // The bulk-shared job settings.
	JobActions    []*pbdriver.JobAction     `protobuf:"bytes,7,rep,name=job_actions,json=jobActions,proto3" json:"job_actions,omitempty"`          // The list actions to be executed.
	WebhookUrl    *string                   `protobuf:"bytes,8,opt,name=webhook_url,json=webhookUrl,proto3,oneof" json:"webhook_url,omitempty"`    // The webhook URL to call when the bulk is completed.
	UserId        string                    `protobuf:"bytes,9,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                      // The user identifier.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BulkSpec) Reset() {
	*x = BulkSpec{}
	mi := &file_pbdataproxy_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BulkSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkSpec) ProtoMessage() {}

func (x *BulkSpec) ProtoReflect() protoreflect.Message {
	mi := &file_pbdataproxy_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkSpec.ProtoReflect.Descriptor instead.
func (*BulkSpec) Descriptor() ([]byte, []int) {
	return file_pbdataproxy_proto_rawDescGZIP(), []int{1}
}

func (x *BulkSpec) GetBulkId() string {
	if x != nil {
		return x.BulkId
	}
	return ""
}

func (x *BulkSpec) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *BulkSpec) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *BulkSpec) GetDriverType() string {
	if x != nil {
		return x.DriverType
	}
	return ""
}

func (x *BulkSpec) GetDevices() []*pbtaskmaster.JobDevice {
	if x != nil {
		return x.Devices
	}
	return nil
}

func (x *BulkSpec) GetSettings() *pbdriver.JobSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *BulkSpec) GetJobActions() []*pbdriver.JobAction {
	if x != nil {
		return x.JobActions
	}
	return nil
}

func (x *BulkSpec) GetWebhookUrl() string {
	if x != nil && x.WebhookUrl != nil {
		return *x.WebhookUrl
	}
	return ""
}

func (x *BulkSpec) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Sub-message - holds the bulk job status.
type BulkStatus struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        BulkStatusCode         `protobuf:"varint,1,opt,name=status,proto3,enum=io.clbs.openhes.pbdataproxy.BulkStatusCode" json:"status,omitempty"` // The job status.
	Jobs          []*BulkJobStatus       `protobuf:"bytes,2,rep,name=jobs,proto3" json:"jobs,omitempty"`                                                      // The list of job statuses.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BulkStatus) Reset() {
	*x = BulkStatus{}
	mi := &file_pbdataproxy_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BulkStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkStatus) ProtoMessage() {}

func (x *BulkStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pbdataproxy_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkStatus.ProtoReflect.Descriptor instead.
func (*BulkStatus) Descriptor() ([]byte, []int) {
	return file_pbdataproxy_proto_rawDescGZIP(), []int{2}
}

func (x *BulkStatus) GetStatus() BulkStatusCode {
	if x != nil {
		return x.Status
	}
	return BulkStatusCode_BULK_STATUS_QUEUED
}

func (x *BulkStatus) GetJobs() []*BulkJobStatus {
	if x != nil {
		return x.Jobs
	}
	return nil
}

type BulkJobStatus struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	JobId         string                  `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"` // The job identifier.
	Status        *pbtaskmaster.JobStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`            // The job status.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BulkJobStatus) Reset() {
	*x = BulkJobStatus{}
	mi := &file_pbdataproxy_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BulkJobStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkJobStatus) ProtoMessage() {}

func (x *BulkJobStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pbdataproxy_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkJobStatus.ProtoReflect.Descriptor instead.
func (*BulkJobStatus) Descriptor() ([]byte, []int) {
	return file_pbdataproxy_proto_rawDescGZIP(), []int{3}
}

func (x *BulkJobStatus) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *BulkJobStatus) GetStatus() *pbtaskmaster.JobStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// DataProxy -> RestApi - the message holds the bulk info and it's status.
type GetBulkResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Spec          *BulkSpec              `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`     // The bulk-job spec.
	Status        *BulkStatus            `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"` // The bulk-job status/data.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBulkResponse) Reset() {
	*x = GetBulkResponse{}
	mi := &file_pbdataproxy_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBulkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBulkResponse) ProtoMessage() {}

func (x *GetBulkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pbdataproxy_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBulkResponse.ProtoReflect.Descriptor instead.
func (*GetBulkResponse) Descriptor() ([]byte, []int) {
	return file_pbdataproxy_proto_rawDescGZIP(), []int{4}
}

func (x *GetBulkResponse) GetSpec() *BulkSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *GetBulkResponse) GetStatus() *BulkStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// RestApi -> DataProxy - the message holds data for bulk listing request.
type GetBulksReuqest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tfrom         *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=tfrom,proto3,oneof" json:"tfrom,omitempty"`                                 // The time range filter - left limit. Optional.
	Tto           *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=tto,proto3,oneof" json:"tto,omitempty"`                                     // The time range filter - right limit. Optional.
	IncludeData   *bool                  `protobuf:"varint,3,opt,name=include_data,json=includeData,proto3,oneof" json:"include_data,omitempty"` // The result content filter - if true then data are included in the response.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBulksReuqest) Reset() {
	*x = GetBulksReuqest{}
	mi := &file_pbdataproxy_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBulksReuqest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBulksReuqest) ProtoMessage() {}

func (x *GetBulksReuqest) ProtoReflect() protoreflect.Message {
	mi := &file_pbdataproxy_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBulksReuqest.ProtoReflect.Descriptor instead.
func (*GetBulksReuqest) Descriptor() ([]byte, []int) {
	return file_pbdataproxy_proto_rawDescGZIP(), []int{5}
}

func (x *GetBulksReuqest) GetTfrom() *timestamppb.Timestamp {
	if x != nil {
		return x.Tfrom
	}
	return nil
}

func (x *GetBulksReuqest) GetTto() *timestamppb.Timestamp {
	if x != nil {
		return x.Tto
	}
	return nil
}

func (x *GetBulksReuqest) GetIncludeData() bool {
	if x != nil && x.IncludeData != nil {
		return *x.IncludeData
	}
	return false
}

// DataProxy -> RestApi - the message holds list of bulks.
type GetBulksResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Bulks         []*GetBulkResponse     `protobuf:"bytes,1,rep,name=bulks,proto3" json:"bulks,omitempty"` // The list of bulk jobs related to the original request.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBulksResponse) Reset() {
	*x = GetBulksResponse{}
	mi := &file_pbdataproxy_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBulksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBulksResponse) ProtoMessage() {}

func (x *GetBulksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pbdataproxy_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBulksResponse.ProtoReflect.Descriptor instead.
func (*GetBulksResponse) Descriptor() ([]byte, []int) {
	return file_pbdataproxy_proto_rawDescGZIP(), []int{6}
}

func (x *GetBulksResponse) GetBulks() []*GetBulkResponse {
	if x != nil {
		return x.Bulks
	}
	return nil
}

// RestApi -> DataProxt - the message holds single bulk request.
type GetBulkRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	BulkId        string                 `protobuf:"bytes,1,opt,name=bulk_id,json=bulkId,proto3" json:"bulk_id,omitempty"` // The bulk identifier to be retrieved.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetBulkRequest) Reset() {
	*x = GetBulkRequest{}
	mi := &file_pbdataproxy_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetBulkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBulkRequest) ProtoMessage() {}

func (x *GetBulkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pbdataproxy_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBulkRequest.ProtoReflect.Descriptor instead.
func (*GetBulkRequest) Descriptor() ([]byte, []int) {
	return file_pbdataproxy_proto_rawDescGZIP(), []int{7}
}

func (x *GetBulkRequest) GetBulkId() string {
	if x != nil {
		return x.BulkId
	}
	return ""
}

// RestApi -> DataProxy - request to get the job status.
type GetJobStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	JobId         string                 `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"` // The job identifier.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetJobStatusRequest) Reset() {
	*x = GetJobStatusRequest{}
	mi := &file_pbdataproxy_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetJobStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobStatusRequest) ProtoMessage() {}

func (x *GetJobStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pbdataproxy_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobStatusRequest.ProtoReflect.Descriptor instead.
func (*GetJobStatusRequest) Descriptor() ([]byte, []int) {
	return file_pbdataproxy_proto_rawDescGZIP(), []int{8}
}

func (x *GetJobStatusRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

// DataProxy -> RestApi - the message holds the job status.
type GetJobStatusResponse struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Status        *pbtaskmaster.JobStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // The status of the job
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetJobStatusResponse) Reset() {
	*x = GetJobStatusResponse{}
	mi := &file_pbdataproxy_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetJobStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobStatusResponse) ProtoMessage() {}

func (x *GetJobStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pbdataproxy_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobStatusResponse.ProtoReflect.Descriptor instead.
func (*GetJobStatusResponse) Descriptor() ([]byte, []int) {
	return file_pbdataproxy_proto_rawDescGZIP(), []int{9}
}

func (x *GetJobStatusResponse) GetStatus() *pbtaskmaster.JobStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_pbdataproxy_proto protoreflect.FileDescriptor

var file_pbdataproxy_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x70, 0x62, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0e, 0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x6c, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x22, 0x9d, 0x03, 0x0a, 0x08, 0x42, 0x75, 0x6c, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x17, 0x0a, 0x07, 0x62, 0x75, 0x6c, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x62, 0x75, 0x6c, 0x6b, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72,
	0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x41, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c,
	0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x74, 0x61, 0x73,
	0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x08, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69,
	0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70,
	0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x44, 0x0a,
	0x0b, 0x6a, 0x6f, 0x62, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x4a, 0x6f,
	0x62, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6a, 0x6f, 0x62, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0b, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x77, 0x65, 0x62, 0x68,
	0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x75,
	0x72, 0x6c, 0x22, 0x91, 0x01, 0x0a, 0x0a, 0x42, 0x75, 0x6c, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x43, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2b, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x42, 0x75, 0x6c, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3e, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x22, 0x67, 0x0a, 0x0d, 0x42, 0x75, 0x6c, 0x6b, 0x4a, 0x6f,
	0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x3f,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73,
	0x2e, 0x70, 0x62, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4a, 0x6f,
	0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x8d, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x42, 0x75, 0x6c, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x12, 0x3f,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73,
	0x2e, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x42, 0x75, 0x6c,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0xc6, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x73, 0x52, 0x65, 0x75, 0x71,
	0x65, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x05, 0x74, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00,
	0x52, 0x05, 0x74, 0x66, 0x72, 0x6f, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x31, 0x0a, 0x03, 0x74, 0x74,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x03, 0x74, 0x74, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a,
	0x0c, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x0b, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x66, 0x72, 0x6f, 0x6d, 0x42,
	0x06, 0x0a, 0x04, 0x5f, 0x74, 0x74, 0x6f, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x69, 0x6e, 0x63, 0x6c,
	0x75, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0x56, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42,
	0x75, 0x6c, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x05,
	0x62, 0x75, 0x6c, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x69, 0x6f,
	0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x62, 0x75, 0x6c, 0x6b, 0x73,
	0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x62, 0x75, 0x6c, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x6c, 0x6b, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2a, 0x90, 0x01, 0x0a, 0x0e, 0x42, 0x75, 0x6c, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x42, 0x55, 0x4c, 0x4b, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a,
	0x13, 0x42, 0x55, 0x4c, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x55, 0x4e,
	0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x42, 0x55, 0x4c, 0x4b, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10,
	0x02, 0x12, 0x19, 0x0a, 0x15, 0x42, 0x55, 0x4c, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13,
	0x42, 0x55, 0x4c, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x45, 0x58, 0x50, 0x49,
	0x52, 0x45, 0x44, 0x10, 0x04, 0x32, 0xac, 0x03, 0x0a, 0x10, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x6c, 0x6b, 0x12, 0x2e, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c,
	0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74,
	0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x6c,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x67, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x73, 0x12, 0x2c, 0x2e, 0x69,
	0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70,
	0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75,
	0x6c, 0x6b, 0x73, 0x52, 0x65, 0x75, 0x71, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x69, 0x6f, 0x2e,
	0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64,
	0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x42, 0x75, 0x6c, 0x6b, 0x12, 0x2b, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2c, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x73, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x30, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65,
	0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x47, 0x65,
	0x74, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x31, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x63, 0x79, 0x62, 0x72, 0x6f, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x68, 0x65,
	0x73, 0x2d, 0x32, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbdataproxy_proto_rawDescOnce sync.Once
	file_pbdataproxy_proto_rawDescData = file_pbdataproxy_proto_rawDesc
)

func file_pbdataproxy_proto_rawDescGZIP() []byte {
	file_pbdataproxy_proto_rawDescOnce.Do(func() {
		file_pbdataproxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbdataproxy_proto_rawDescData)
	})
	return file_pbdataproxy_proto_rawDescData
}

var file_pbdataproxy_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pbdataproxy_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_pbdataproxy_proto_goTypes = []any{
	(BulkStatusCode)(0),            // 0: io.clbs.openhes.pbdataproxy.BulkStatusCode
	(*CreateBulkRequest)(nil),      // 1: io.clbs.openhes.pbdataproxy.CreateBulkRequest
	(*BulkSpec)(nil),               // 2: io.clbs.openhes.pbdataproxy.BulkSpec
	(*BulkStatus)(nil),             // 3: io.clbs.openhes.pbdataproxy.BulkStatus
	(*BulkJobStatus)(nil),          // 4: io.clbs.openhes.pbdataproxy.BulkJobStatus
	(*GetBulkResponse)(nil),        // 5: io.clbs.openhes.pbdataproxy.GetBulkResponse
	(*GetBulksReuqest)(nil),        // 6: io.clbs.openhes.pbdataproxy.GetBulksReuqest
	(*GetBulksResponse)(nil),       // 7: io.clbs.openhes.pbdataproxy.GetBulksResponse
	(*GetBulkRequest)(nil),         // 8: io.clbs.openhes.pbdataproxy.GetBulkRequest
	(*GetJobStatusRequest)(nil),    // 9: io.clbs.openhes.pbdataproxy.GetJobStatusRequest
	(*GetJobStatusResponse)(nil),   // 10: io.clbs.openhes.pbdataproxy.GetJobStatusResponse
	(*pbtaskmaster.JobDevice)(nil), // 11: io.clbs.openhes.pbtaskmaster.JobDevice
	(*pbdriver.JobSettings)(nil),   // 12: io.clbs.openhes.pbdriver.JobSettings
	(*pbdriver.JobAction)(nil),     // 13: io.clbs.openhes.pbdriver.JobAction
	(*pbtaskmaster.JobStatus)(nil), // 14: io.clbs.openhes.pbtaskmaster.JobStatus
	(*timestamppb.Timestamp)(nil),  // 15: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),          // 16: google.protobuf.Empty
}
var file_pbdataproxy_proto_depIdxs = []int32{
	2,  // 0: io.clbs.openhes.pbdataproxy.CreateBulkRequest.spec:type_name -> io.clbs.openhes.pbdataproxy.BulkSpec
	11, // 1: io.clbs.openhes.pbdataproxy.BulkSpec.devices:type_name -> io.clbs.openhes.pbtaskmaster.JobDevice
	12, // 2: io.clbs.openhes.pbdataproxy.BulkSpec.settings:type_name -> io.clbs.openhes.pbdriver.JobSettings
	13, // 3: io.clbs.openhes.pbdataproxy.BulkSpec.job_actions:type_name -> io.clbs.openhes.pbdriver.JobAction
	0,  // 4: io.clbs.openhes.pbdataproxy.BulkStatus.status:type_name -> io.clbs.openhes.pbdataproxy.BulkStatusCode
	4,  // 5: io.clbs.openhes.pbdataproxy.BulkStatus.jobs:type_name -> io.clbs.openhes.pbdataproxy.BulkJobStatus
	14, // 6: io.clbs.openhes.pbdataproxy.BulkJobStatus.status:type_name -> io.clbs.openhes.pbtaskmaster.JobStatus
	2,  // 7: io.clbs.openhes.pbdataproxy.GetBulkResponse.spec:type_name -> io.clbs.openhes.pbdataproxy.BulkSpec
	3,  // 8: io.clbs.openhes.pbdataproxy.GetBulkResponse.status:type_name -> io.clbs.openhes.pbdataproxy.BulkStatus
	15, // 9: io.clbs.openhes.pbdataproxy.GetBulksReuqest.tfrom:type_name -> google.protobuf.Timestamp
	15, // 10: io.clbs.openhes.pbdataproxy.GetBulksReuqest.tto:type_name -> google.protobuf.Timestamp
	5,  // 11: io.clbs.openhes.pbdataproxy.GetBulksResponse.bulks:type_name -> io.clbs.openhes.pbdataproxy.GetBulkResponse
	14, // 12: io.clbs.openhes.pbdataproxy.GetJobStatusResponse.status:type_name -> io.clbs.openhes.pbtaskmaster.JobStatus
	1,  // 13: io.clbs.openhes.pbdataproxy.DataproxyService.CreateBulk:input_type -> io.clbs.openhes.pbdataproxy.CreateBulkRequest
	6,  // 14: io.clbs.openhes.pbdataproxy.DataproxyService.GetBulks:input_type -> io.clbs.openhes.pbdataproxy.GetBulksReuqest
	8,  // 15: io.clbs.openhes.pbdataproxy.DataproxyService.GetBulk:input_type -> io.clbs.openhes.pbdataproxy.GetBulkRequest
	9,  // 16: io.clbs.openhes.pbdataproxy.DataproxyService.GetJobStatus:input_type -> io.clbs.openhes.pbdataproxy.GetJobStatusRequest
	16, // 17: io.clbs.openhes.pbdataproxy.DataproxyService.CreateBulk:output_type -> google.protobuf.Empty
	7,  // 18: io.clbs.openhes.pbdataproxy.DataproxyService.GetBulks:output_type -> io.clbs.openhes.pbdataproxy.GetBulksResponse
	5,  // 19: io.clbs.openhes.pbdataproxy.DataproxyService.GetBulk:output_type -> io.clbs.openhes.pbdataproxy.GetBulkResponse
	10, // 20: io.clbs.openhes.pbdataproxy.DataproxyService.GetJobStatus:output_type -> io.clbs.openhes.pbdataproxy.GetJobStatusResponse
	17, // [17:21] is the sub-list for method output_type
	13, // [13:17] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_pbdataproxy_proto_init() }
func file_pbdataproxy_proto_init() {
	if File_pbdataproxy_proto != nil {
		return
	}
	file_pbdataproxy_proto_msgTypes[1].OneofWrappers = []any{}
	file_pbdataproxy_proto_msgTypes[5].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbdataproxy_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pbdataproxy_proto_goTypes,
		DependencyIndexes: file_pbdataproxy_proto_depIdxs,
		EnumInfos:         file_pbdataproxy_proto_enumTypes,
		MessageInfos:      file_pbdataproxy_proto_msgTypes,
	}.Build()
	File_pbdataproxy_proto = out.File
	file_pbdataproxy_proto_rawDesc = nil
	file_pbdataproxy_proto_goTypes = nil
	file_pbdataproxy_proto_depIdxs = nil
}
