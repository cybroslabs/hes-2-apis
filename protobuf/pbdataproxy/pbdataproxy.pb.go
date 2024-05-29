// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: pbdataproxy.proto

package pbdataproxy

import (
	pbdriver "github.com/cybroslabs/hes-2-apis/protobuf/pbdriver"
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

// Taskmaster -> Dataproxy job/action progress update message
type BulkJobEventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobId []byte `protobuf:"bytes,1,opt,name=job_id,json=i,proto3" json:"job_id,omitempty"` // The job identifier.
}

func (x *BulkJobEventData) Reset() {
	*x = BulkJobEventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbdataproxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkJobEventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkJobEventData) ProtoMessage() {}

func (x *BulkJobEventData) ProtoReflect() protoreflect.Message {
	mi := &file_pbdataproxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BulkJobEventData.ProtoReflect.Descriptor instead.
func (*BulkJobEventData) Descriptor() ([]byte, []int) {
	return file_pbdataproxy_proto_rawDescGZIP(), []int{0}
}

func (x *BulkJobEventData) GetJobId() []byte {
	if x != nil {
		return x.JobId
	}
	return nil
}

// RestApi -> DataProxy
type CreateBulksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spec *BulkSpec `protobuf:"bytes,1,opt,name=spec,json=s,proto3" json:"spec,omitempty"` // The bulk-job spec.
}

func (x *CreateBulksRequest) Reset() {
	*x = CreateBulksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbdataproxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBulksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBulksRequest) ProtoMessage() {}

func (x *CreateBulksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pbdataproxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBulksRequest.ProtoReflect.Descriptor instead.
func (*CreateBulksRequest) Descriptor() ([]byte, []int) {
	return file_pbdataproxy_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBulksRequest) GetSpec() *BulkSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

// Sub-message - holds the bulk job specification.
type BulkSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BulkId           string                `protobuf:"bytes,1,opt,name=bulk_id,json=i,proto3" json:"bulk_id,omitempty"`                       // The bulk identifier.
	CorrelationId    string                `protobuf:"bytes,2,opt,name=correlation_id,json=c,proto3" json:"correlation_id,omitempty"`         // The correlation identifier, e.g. to define relation to non-homogenous group.
	OrgId            string                `protobuf:"bytes,3,opt,name=org_id,json=o,proto3" json:"org_id,omitempty"`                         // The organization identifier.
	DeviceDriverType string                `protobuf:"bytes,4,opt,name=device_driver_type,json=t,proto3" json:"device_driver_type,omitempty"` // The device driver type.
	Devices          []*JobDevice          `protobuf:"bytes,5,rep,name=devices,json=d,proto3" json:"devices,omitempty"`                       // The list of devices in the bulk.
	Settings         *pbdriver.JobSettings `protobuf:"bytes,6,opt,name=settings,json=s,proto3" json:"settings,omitempty"`                     // The bulk-shared job settings.
	JobActions       []*pbdriver.JobAction `protobuf:"bytes,7,rep,name=job_actions,json=a,proto3" json:"job_actions,omitempty"`               // The list actions to be executed.
}

func (x *BulkSpec) Reset() {
	*x = BulkSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbdataproxy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BulkSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BulkSpec) ProtoMessage() {}

func (x *BulkSpec) ProtoReflect() protoreflect.Message {
	mi := &file_pbdataproxy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_pbdataproxy_proto_rawDescGZIP(), []int{2}
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

func (x *BulkSpec) GetDeviceDriverType() string {
	if x != nil {
		return x.DeviceDriverType
	}
	return ""
}

func (x *BulkSpec) GetDevices() []*JobDevice {
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

// DataProxy -> RestApi - the message holds the bulk info and it's status.
type GetBulkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spec   *BulkSpec           `protobuf:"bytes,1,opt,name=spec,json=s,proto3" json:"spec,omitempty"`     // The bulk-job spec.
	Status *pbdriver.JobStatus `protobuf:"bytes,2,opt,name=status,json=t,proto3" json:"status,omitempty"` // The bulk-job status/data.
}

func (x *GetBulkResponse) Reset() {
	*x = GetBulkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbdataproxy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBulkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBulkResponse) ProtoMessage() {}

func (x *GetBulkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pbdataproxy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_pbdataproxy_proto_rawDescGZIP(), []int{3}
}

func (x *GetBulkResponse) GetSpec() *BulkSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *GetBulkResponse) GetStatus() *pbdriver.JobStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// Sub-message representing a single device in a bulk.
type JobDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                   `protobuf:"bytes,1,opt,name=id,json=i,proto3" json:"id,omitempty"`                           // The device (job) identifier.
	ExternalId     string                   `protobuf:"bytes,3,opt,name=external_id,json=e,proto3" json:"external_id,omitempty"`         // The external identifier.
	ConnectionInfo *pbdriver.ConnectionInfo `protobuf:"bytes,4,opt,name=connection_info,json=c,proto3" json:"connection_info,omitempty"` // The conenction (device) parameters.
}

func (x *JobDevice) Reset() {
	*x = JobDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbdataproxy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JobDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobDevice) ProtoMessage() {}

func (x *JobDevice) ProtoReflect() protoreflect.Message {
	mi := &file_pbdataproxy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JobDevice.ProtoReflect.Descriptor instead.
func (*JobDevice) Descriptor() ([]byte, []int) {
	return file_pbdataproxy_proto_rawDescGZIP(), []int{4}
}

func (x *JobDevice) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *JobDevice) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *JobDevice) GetConnectionInfo() *pbdriver.ConnectionInfo {
	if x != nil {
		return x.ConnectionInfo
	}
	return nil
}

// RestApi -> DataProxy - the message holds data for bulk listing request.
type GetBulksReuqest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tfrom       *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=tfrom,json=tf,proto3,oneof" json:"tfrom,omitempty"`               // The time range filter - left limit. Optional.
	Tto         *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=tto,json=tt,proto3,oneof" json:"tto,omitempty"`                   // The time range filter - right limit. Optional.
	IncludeData *bool                  `protobuf:"varint,3,opt,name=include_data,json=d,proto3,oneof" json:"include_data,omitempty"` // The result content filter - if true then data are included in the response.
	// FIXME: Should not not be optional!
	OrgId *string `protobuf:"bytes,4,opt,name=org_id,json=o,proto3,oneof" json:"org_id,omitempty"` // The organization identifier filter.
}

func (x *GetBulksReuqest) Reset() {
	*x = GetBulksReuqest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbdataproxy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBulksReuqest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBulksReuqest) ProtoMessage() {}

func (x *GetBulksReuqest) ProtoReflect() protoreflect.Message {
	mi := &file_pbdataproxy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
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

func (x *GetBulksReuqest) GetOrgId() string {
	if x != nil && x.OrgId != nil {
		return *x.OrgId
	}
	return ""
}

// DataProxy -> RestApi - the message holds list of bulks.
type GetBulksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bulks []*GetBulkResponse `protobuf:"bytes,1,rep,name=bulks,json=b,proto3" json:"bulks,omitempty"` // The list of bulk jobs related to the original request.
}

func (x *GetBulksResponse) Reset() {
	*x = GetBulksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbdataproxy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBulksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBulksResponse) ProtoMessage() {}

func (x *GetBulksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pbdataproxy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BulkId string `protobuf:"bytes,1,opt,name=bulk_id,json=i,proto3" json:"bulk_id,omitempty"` // The bulk identifier to be retrieved.
}

func (x *GetBulkRequest) Reset() {
	*x = GetBulkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbdataproxy_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBulkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBulkRequest) ProtoMessage() {}

func (x *GetBulkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pbdataproxy_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
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

var File_pbdataproxy_proto protoreflect.FileDescriptor

var file_pbdataproxy_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e,
	0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25,
	0x0a, 0x10, 0x42, 0x75, 0x6c, 0x6b, 0x4a, 0x6f, 0x62, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x11, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x01, 0x69, 0x22, 0x4c, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x75, 0x6c, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6f, 0x2e, 0x63,
	0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61,
	0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x01, 0x73, 0x22, 0xa0, 0x02, 0x0a, 0x08, 0x42, 0x75, 0x6c, 0x6b, 0x53, 0x70, 0x65, 0x63,
	0x12, 0x12, 0x0a, 0x07, 0x62, 0x75, 0x6c, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x01, 0x69, 0x12, 0x19, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x63, 0x12,
	0x11, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x01, 0x6f, 0x12, 0x1d, 0x0a, 0x12, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01,
	0x74, 0x12, 0x3a, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x4a, 0x6f, 0x62, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x01, 0x64, 0x12, 0x3a, 0x0a,
	0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x25, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65,
	0x73, 0x2e, 0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x01, 0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x6a, 0x6f, 0x62,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73,
	0x2e, 0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x01, 0x61, 0x22, 0x81, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x75,
	0x6c, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x73, 0x70,
	0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c,
	0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74,
	0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x52,
	0x01, 0x73, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x4a, 0x6f,
	0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x01, 0x74, 0x22, 0x78, 0x0a, 0x09, 0x4a, 0x6f,
	0x62, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0d, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x01, 0x69, 0x12, 0x16, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x65, 0x12, 0x44,
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62,
	0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x01, 0x63, 0x22, 0xdb, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b,
	0x73, 0x52, 0x65, 0x75, 0x71, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x05, 0x74, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x02, 0x74, 0x66, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x03,
	0x74, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x02, 0x74, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1c,
	0x0a, 0x0c, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x01, 0x64, 0x88, 0x01, 0x01, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x01,
	0x6f, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x66, 0x72, 0x6f, 0x6d, 0x42, 0x06,
	0x0a, 0x04, 0x5f, 0x74, 0x74, 0x6f, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x69, 0x6e, 0x63, 0x6c, 0x75,
	0x64, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x72, 0x67, 0x5f,
	0x69, 0x64, 0x22, 0x52, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x62, 0x75, 0x6c, 0x6b, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x01, 0x62, 0x22, 0x24, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x07, 0x62, 0x75, 0x6c, 0x6b,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x69, 0x32, 0x94, 0x03, 0x0a,
	0x10, 0x44, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x5a, 0x0a, 0x11, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4a, 0x6f, 0x62, 0x46, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x2d, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73,
	0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x42, 0x75, 0x6c, 0x6b, 0x4a, 0x6f, 0x62, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x55, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x6c, 0x6b, 0x12, 0x2f, 0x2e, 0x69, 0x6f,
	0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x75, 0x6c, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x67, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x73,
	0x12, 0x2c, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68,
	0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x73, 0x52, 0x65, 0x75, 0x71, 0x65, 0x73, 0x74, 0x1a, 0x2d,
	0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73,
	0x2e, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x75, 0x6c, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x12, 0x2b, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c,
	0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74,
	0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x79, 0x62, 0x72, 0x6f, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x68, 0x65, 0x73,
	0x2d, 0x32, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
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

var file_pbdataproxy_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pbdataproxy_proto_goTypes = []interface{}{
	(*BulkJobEventData)(nil),        // 0: io.clbs.openhes.pbdataproxy.BulkJobEventData
	(*CreateBulksRequest)(nil),      // 1: io.clbs.openhes.pbdataproxy.CreateBulksRequest
	(*BulkSpec)(nil),                // 2: io.clbs.openhes.pbdataproxy.BulkSpec
	(*GetBulkResponse)(nil),         // 3: io.clbs.openhes.pbdataproxy.GetBulkResponse
	(*JobDevice)(nil),               // 4: io.clbs.openhes.pbdataproxy.JobDevice
	(*GetBulksReuqest)(nil),         // 5: io.clbs.openhes.pbdataproxy.GetBulksReuqest
	(*GetBulksResponse)(nil),        // 6: io.clbs.openhes.pbdataproxy.GetBulksResponse
	(*GetBulkRequest)(nil),          // 7: io.clbs.openhes.pbdataproxy.GetBulkRequest
	(*pbdriver.JobSettings)(nil),    // 8: io.clbs.openhes.pbdriver.JobSettings
	(*pbdriver.JobAction)(nil),      // 9: io.clbs.openhes.pbdriver.JobAction
	(*pbdriver.JobStatus)(nil),      // 10: io.clbs.openhes.pbdriver.JobStatus
	(*pbdriver.ConnectionInfo)(nil), // 11: io.clbs.openhes.pbdriver.ConnectionInfo
	(*timestamppb.Timestamp)(nil),   // 12: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),           // 13: google.protobuf.Empty
}
var file_pbdataproxy_proto_depIdxs = []int32{
	2,  // 0: io.clbs.openhes.pbdataproxy.CreateBulksRequest.spec:type_name -> io.clbs.openhes.pbdataproxy.BulkSpec
	4,  // 1: io.clbs.openhes.pbdataproxy.BulkSpec.devices:type_name -> io.clbs.openhes.pbdataproxy.JobDevice
	8,  // 2: io.clbs.openhes.pbdataproxy.BulkSpec.settings:type_name -> io.clbs.openhes.pbdriver.JobSettings
	9,  // 3: io.clbs.openhes.pbdataproxy.BulkSpec.job_actions:type_name -> io.clbs.openhes.pbdriver.JobAction
	2,  // 4: io.clbs.openhes.pbdataproxy.GetBulkResponse.spec:type_name -> io.clbs.openhes.pbdataproxy.BulkSpec
	10, // 5: io.clbs.openhes.pbdataproxy.GetBulkResponse.status:type_name -> io.clbs.openhes.pbdriver.JobStatus
	11, // 6: io.clbs.openhes.pbdataproxy.JobDevice.connection_info:type_name -> io.clbs.openhes.pbdriver.ConnectionInfo
	12, // 7: io.clbs.openhes.pbdataproxy.GetBulksReuqest.tfrom:type_name -> google.protobuf.Timestamp
	12, // 8: io.clbs.openhes.pbdataproxy.GetBulksReuqest.tto:type_name -> google.protobuf.Timestamp
	3,  // 9: io.clbs.openhes.pbdataproxy.GetBulksResponse.bulks:type_name -> io.clbs.openhes.pbdataproxy.GetBulkResponse
	0,  // 10: io.clbs.openhes.pbdataproxy.DataproxyService.NotifyJobFinished:input_type -> io.clbs.openhes.pbdataproxy.BulkJobEventData
	1,  // 11: io.clbs.openhes.pbdataproxy.DataproxyService.CreateBulk:input_type -> io.clbs.openhes.pbdataproxy.CreateBulksRequest
	5,  // 12: io.clbs.openhes.pbdataproxy.DataproxyService.GetBulks:input_type -> io.clbs.openhes.pbdataproxy.GetBulksReuqest
	7,  // 13: io.clbs.openhes.pbdataproxy.DataproxyService.GetBulk:input_type -> io.clbs.openhes.pbdataproxy.GetBulkRequest
	13, // 14: io.clbs.openhes.pbdataproxy.DataproxyService.NotifyJobFinished:output_type -> google.protobuf.Empty
	13, // 15: io.clbs.openhes.pbdataproxy.DataproxyService.CreateBulk:output_type -> google.protobuf.Empty
	6,  // 16: io.clbs.openhes.pbdataproxy.DataproxyService.GetBulks:output_type -> io.clbs.openhes.pbdataproxy.GetBulksResponse
	3,  // 17: io.clbs.openhes.pbdataproxy.DataproxyService.GetBulk:output_type -> io.clbs.openhes.pbdataproxy.GetBulkResponse
	14, // [14:18] is the sub-list for method output_type
	10, // [10:14] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_pbdataproxy_proto_init() }
func file_pbdataproxy_proto_init() {
	if File_pbdataproxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbdataproxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkJobEventData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbdataproxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBulksRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbdataproxy_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BulkSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbdataproxy_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBulkResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbdataproxy_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JobDevice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbdataproxy_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBulksReuqest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbdataproxy_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBulksResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbdataproxy_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBulkRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_pbdataproxy_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbdataproxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pbdataproxy_proto_goTypes,
		DependencyIndexes: file_pbdataproxy_proto_depIdxs,
		MessageInfos:      file_pbdataproxy_proto_msgTypes,
	}.Build()
	File_pbdataproxy_proto = out.File
	file_pbdataproxy_proto_rawDesc = nil
	file_pbdataproxy_proto_goTypes = nil
	file_pbdataproxy_proto_depIdxs = nil
}
