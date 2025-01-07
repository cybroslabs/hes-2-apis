// Editions version of proto3 file

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v5.29.2
// source: pbapi.proto

package pbapi

import (
	pbdataproxymodels "github.com/cybroslabs/hes-2-apis/protobuf/pbdataproxymodels"
	pbdrivermodels "github.com/cybroslabs/hes-2-apis/protobuf/pbdrivermodels"
	pbdriveroperatormodels "github.com/cybroslabs/hes-2-apis/protobuf/pbdriveroperatormodels"
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

type PublicCreateBulkRequest struct {
	state                    protoimpl.MessageState           `protogen:"opaque.v1"`
	xxx_hidden_Id            *string                          `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	xxx_hidden_CorrelationId *string                          `protobuf:"bytes,2,opt,name=correlation_id,json=correlationId" json:"correlation_id,omitempty"`
	xxx_hidden_DriverType    *string                          `protobuf:"bytes,3,opt,name=driver_type,json=driverType" json:"driver_type,omitempty"`
	xxx_hidden_Device        isPublicCreateBulkRequest_Device `protobuf_oneof:"device"`
	xxx_hidden_Settings      *pbdrivermodels.JobSettings      `protobuf:"bytes,6,opt,name=settings" json:"settings,omitempty"`
	xxx_hidden_Actions       *[]*pbdrivermodels.JobAction     `protobuf:"bytes,7,rep,name=actions" json:"actions,omitempty"`
	xxx_hidden_WebhookUrl    *string                          `protobuf:"bytes,8,opt,name=webhook_url,json=webhookUrl" json:"webhook_url,omitempty"`
	XXX_raceDetectHookData   protoimpl.RaceDetectHookData
	XXX_presence             [1]uint32
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *PublicCreateBulkRequest) Reset() {
	*x = PublicCreateBulkRequest{}
	mi := &file_pbapi_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublicCreateBulkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicCreateBulkRequest) ProtoMessage() {}

func (x *PublicCreateBulkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pbapi_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *PublicCreateBulkRequest) GetId() string {
	if x != nil {
		if x.xxx_hidden_Id != nil {
			return *x.xxx_hidden_Id
		}
		return ""
	}
	return ""
}

func (x *PublicCreateBulkRequest) GetCorrelationId() string {
	if x != nil {
		if x.xxx_hidden_CorrelationId != nil {
			return *x.xxx_hidden_CorrelationId
		}
		return ""
	}
	return ""
}

func (x *PublicCreateBulkRequest) GetDriverType() string {
	if x != nil {
		if x.xxx_hidden_DriverType != nil {
			return *x.xxx_hidden_DriverType
		}
		return ""
	}
	return ""
}

func (x *PublicCreateBulkRequest) GetDevices() *JobDeviceList {
	if x != nil {
		if x, ok := x.xxx_hidden_Device.(*publicCreateBulkRequest_Devices); ok {
			return x.Devices
		}
	}
	return nil
}

func (x *PublicCreateBulkRequest) GetCustomDevices() *JobCustomDeviceList {
	if x != nil {
		if x, ok := x.xxx_hidden_Device.(*publicCreateBulkRequest_CustomDevices); ok {
			return x.CustomDevices
		}
	}
	return nil
}

func (x *PublicCreateBulkRequest) GetSettings() *pbdrivermodels.JobSettings {
	if x != nil {
		return x.xxx_hidden_Settings
	}
	return nil
}

func (x *PublicCreateBulkRequest) GetActions() []*pbdrivermodels.JobAction {
	if x != nil {
		if x.xxx_hidden_Actions != nil {
			return *x.xxx_hidden_Actions
		}
	}
	return nil
}

func (x *PublicCreateBulkRequest) GetWebhookUrl() string {
	if x != nil {
		if x.xxx_hidden_WebhookUrl != nil {
			return *x.xxx_hidden_WebhookUrl
		}
		return ""
	}
	return ""
}

func (x *PublicCreateBulkRequest) SetId(v string) {
	x.xxx_hidden_Id = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 7)
}

func (x *PublicCreateBulkRequest) SetCorrelationId(v string) {
	x.xxx_hidden_CorrelationId = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 7)
}

func (x *PublicCreateBulkRequest) SetDriverType(v string) {
	x.xxx_hidden_DriverType = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 7)
}

func (x *PublicCreateBulkRequest) SetDevices(v *JobDeviceList) {
	if v == nil {
		x.xxx_hidden_Device = nil
		return
	}
	x.xxx_hidden_Device = &publicCreateBulkRequest_Devices{v}
}

func (x *PublicCreateBulkRequest) SetCustomDevices(v *JobCustomDeviceList) {
	if v == nil {
		x.xxx_hidden_Device = nil
		return
	}
	x.xxx_hidden_Device = &publicCreateBulkRequest_CustomDevices{v}
}

func (x *PublicCreateBulkRequest) SetSettings(v *pbdrivermodels.JobSettings) {
	x.xxx_hidden_Settings = v
}

func (x *PublicCreateBulkRequest) SetActions(v []*pbdrivermodels.JobAction) {
	x.xxx_hidden_Actions = &v
}

func (x *PublicCreateBulkRequest) SetWebhookUrl(v string) {
	x.xxx_hidden_WebhookUrl = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 6, 7)
}

func (x *PublicCreateBulkRequest) HasId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *PublicCreateBulkRequest) HasCorrelationId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *PublicCreateBulkRequest) HasDriverType() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *PublicCreateBulkRequest) HasDevice() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Device != nil
}

func (x *PublicCreateBulkRequest) HasDevices() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_Device.(*publicCreateBulkRequest_Devices)
	return ok
}

func (x *PublicCreateBulkRequest) HasCustomDevices() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_Device.(*publicCreateBulkRequest_CustomDevices)
	return ok
}

func (x *PublicCreateBulkRequest) HasSettings() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Settings != nil
}

func (x *PublicCreateBulkRequest) HasWebhookUrl() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 6)
}

func (x *PublicCreateBulkRequest) ClearId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Id = nil
}

func (x *PublicCreateBulkRequest) ClearCorrelationId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_CorrelationId = nil
}

func (x *PublicCreateBulkRequest) ClearDriverType() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_DriverType = nil
}

func (x *PublicCreateBulkRequest) ClearDevice() {
	x.xxx_hidden_Device = nil
}

func (x *PublicCreateBulkRequest) ClearDevices() {
	if _, ok := x.xxx_hidden_Device.(*publicCreateBulkRequest_Devices); ok {
		x.xxx_hidden_Device = nil
	}
}

func (x *PublicCreateBulkRequest) ClearCustomDevices() {
	if _, ok := x.xxx_hidden_Device.(*publicCreateBulkRequest_CustomDevices); ok {
		x.xxx_hidden_Device = nil
	}
}

func (x *PublicCreateBulkRequest) ClearSettings() {
	x.xxx_hidden_Settings = nil
}

func (x *PublicCreateBulkRequest) ClearWebhookUrl() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 6)
	x.xxx_hidden_WebhookUrl = nil
}

const PublicCreateBulkRequest_Device_not_set_case case_PublicCreateBulkRequest_Device = 0
const PublicCreateBulkRequest_Devices_case case_PublicCreateBulkRequest_Device = 4
const PublicCreateBulkRequest_CustomDevices_case case_PublicCreateBulkRequest_Device = 5

func (x *PublicCreateBulkRequest) WhichDevice() case_PublicCreateBulkRequest_Device {
	if x == nil {
		return PublicCreateBulkRequest_Device_not_set_case
	}
	switch x.xxx_hidden_Device.(type) {
	case *publicCreateBulkRequest_Devices:
		return PublicCreateBulkRequest_Devices_case
	case *publicCreateBulkRequest_CustomDevices:
		return PublicCreateBulkRequest_CustomDevices_case
	default:
		return PublicCreateBulkRequest_Device_not_set_case
	}
}

type PublicCreateBulkRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Id            *string
	CorrelationId *string
	DriverType    *string
	// Fields of oneof xxx_hidden_Device:
	Devices       *JobDeviceList
	CustomDevices *JobCustomDeviceList
	// -- end of xxx_hidden_Device
	Settings   *pbdrivermodels.JobSettings
	Actions    []*pbdrivermodels.JobAction
	WebhookUrl *string
}

func (b0 PublicCreateBulkRequest_builder) Build() *PublicCreateBulkRequest {
	m0 := &PublicCreateBulkRequest{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Id != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 7)
		x.xxx_hidden_Id = b.Id
	}
	if b.CorrelationId != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 7)
		x.xxx_hidden_CorrelationId = b.CorrelationId
	}
	if b.DriverType != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 7)
		x.xxx_hidden_DriverType = b.DriverType
	}
	if b.Devices != nil {
		x.xxx_hidden_Device = &publicCreateBulkRequest_Devices{b.Devices}
	}
	if b.CustomDevices != nil {
		x.xxx_hidden_Device = &publicCreateBulkRequest_CustomDevices{b.CustomDevices}
	}
	x.xxx_hidden_Settings = b.Settings
	x.xxx_hidden_Actions = &b.Actions
	if b.WebhookUrl != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 6, 7)
		x.xxx_hidden_WebhookUrl = b.WebhookUrl
	}
	return m0
}

type case_PublicCreateBulkRequest_Device protoreflect.FieldNumber

func (x case_PublicCreateBulkRequest_Device) String() string {
	md := file_pbapi_proto_msgTypes[0].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isPublicCreateBulkRequest_Device interface {
	isPublicCreateBulkRequest_Device()
}

type publicCreateBulkRequest_Devices struct {
	Devices *JobDeviceList `protobuf:"bytes,4,opt,name=devices,oneof"` // The list of devices in the bulk.
}

type publicCreateBulkRequest_CustomDevices struct {
	CustomDevices *JobCustomDeviceList `protobuf:"bytes,5,opt,name=custom_devices,json=customDevices,oneof"` // The list of custom devices in the bulk.
}

func (*publicCreateBulkRequest_Devices) isPublicCreateBulkRequest_Device() {}

func (*publicCreateBulkRequest_CustomDevices) isPublicCreateBulkRequest_Device() {}

type JobDeviceList struct {
	state           protoimpl.MessageState           `protogen:"opaque.v1"`
	xxx_hidden_List *[]*pbtaskmastermodels.JobDevice `protobuf:"bytes,5,rep,name=list" json:"list,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *JobDeviceList) Reset() {
	*x = JobDeviceList{}
	mi := &file_pbapi_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobDeviceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobDeviceList) ProtoMessage() {}

func (x *JobDeviceList) ProtoReflect() protoreflect.Message {
	mi := &file_pbapi_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *JobDeviceList) GetList() []*pbtaskmastermodels.JobDevice {
	if x != nil {
		if x.xxx_hidden_List != nil {
			return *x.xxx_hidden_List
		}
	}
	return nil
}

func (x *JobDeviceList) SetList(v []*pbtaskmastermodels.JobDevice) {
	x.xxx_hidden_List = &v
}

type JobDeviceList_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	List []*pbtaskmastermodels.JobDevice
}

func (b0 JobDeviceList_builder) Build() *JobDeviceList {
	m0 := &JobDeviceList{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_List = &b.List
	return m0
}

type JobCustomDeviceList struct {
	state                    protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_CustomDevices *[]*JobCustomDevice    `protobuf:"bytes,5,rep,name=custom_devices,json=customDevices" json:"custom_devices,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *JobCustomDeviceList) Reset() {
	*x = JobCustomDeviceList{}
	mi := &file_pbapi_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobCustomDeviceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobCustomDeviceList) ProtoMessage() {}

func (x *JobCustomDeviceList) ProtoReflect() protoreflect.Message {
	mi := &file_pbapi_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *JobCustomDeviceList) GetCustomDevices() []*JobCustomDevice {
	if x != nil {
		if x.xxx_hidden_CustomDevices != nil {
			return *x.xxx_hidden_CustomDevices
		}
	}
	return nil
}

func (x *JobCustomDeviceList) SetCustomDevices(v []*JobCustomDevice) {
	x.xxx_hidden_CustomDevices = &v
}

type JobCustomDeviceList_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	CustomDevices []*JobCustomDevice
}

func (b0 JobCustomDeviceList_builder) Build() *JobCustomDeviceList {
	m0 := &JobCustomDeviceList{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_CustomDevices = &b.CustomDevices
	return m0
}

type JobCustomDevice struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Id          *string                `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *JobCustomDevice) Reset() {
	*x = JobCustomDevice{}
	mi := &file_pbapi_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *JobCustomDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JobCustomDevice) ProtoMessage() {}

func (x *JobCustomDevice) ProtoReflect() protoreflect.Message {
	mi := &file_pbapi_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *JobCustomDevice) GetId() string {
	if x != nil {
		if x.xxx_hidden_Id != nil {
			return *x.xxx_hidden_Id
		}
		return ""
	}
	return ""
}

func (x *JobCustomDevice) SetId(v string) {
	x.xxx_hidden_Id = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *JobCustomDevice) HasId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *JobCustomDevice) ClearId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Id = nil
}

type JobCustomDevice_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Id *string
}

func (b0 JobCustomDevice_builder) Build() *JobCustomDevice {
	m0 := &JobCustomDevice{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Id != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_Id = b.Id
	}
	return m0
}

var File_pbapi_proto protoreflect.FileDescriptor

var file_pbapi_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x62, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x69,
	0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70,
	0x62, 0x61, 0x70, 0x69, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x70, 0x62, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70,
	0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x03,
	0x0a, 0x17, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75,
	0x6c, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x40, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x61, 0x70, 0x69, 0x2e, 0x4a, 0x6f, 0x62, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x07, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x53, 0x0a, 0x0e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x69, 0x6f,
	0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62,
	0x61, 0x70, 0x69, 0x2e, 0x4a, 0x6f, 0x62, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x41, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x69, 0x6f, 0x2e,
	0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x3d, 0x0a, 0x07, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69,
	0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70,
	0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x41, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x42, 0x08, 0x0a, 0x06, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x4c, 0x0a, 0x0d, 0x4a, 0x6f, 0x62, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x4a, 0x6f, 0x62, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x22, 0x64, 0x0a, 0x13, 0x4a, 0x6f, 0x62, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x0e, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x61, 0x70, 0x69, 0x2e, 0x4a, 0x6f, 0x62, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x21, 0x0a, 0x0f, 0x4a, 0x6f, 0x62,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0x82, 0x04, 0x0a,
	0x0a, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x6c, 0x6b, 0x12, 0x2e, 0x2e, 0x69, 0x6f, 0x2e, 0x63,
	0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x61, 0x70,
	0x69, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75,
	0x6c, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x67, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x73, 0x12, 0x2c, 0x2e,
	0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e,
	0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x42,
	0x75, 0x6c, 0x6b, 0x73, 0x52, 0x65, 0x75, 0x71, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x69, 0x6f,
	0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62,
	0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c,
	0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x64, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x42, 0x75, 0x6c, 0x6b, 0x12, 0x2b, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x6c, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x73, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x30, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68,
	0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x47,
	0x65, 0x74, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x31, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65,
	0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x61, 0x74, 0x61, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x34, 0x2e, 0x69, 0x6f,
	0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x79, 0x62, 0x72, 0x6f, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x68, 0x65, 0x73, 0x2d, 0x32,
	0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70,
	0x62, 0x61, 0x70, 0x69, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8,
	0x07,
}

var file_pbapi_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pbapi_proto_goTypes = []any{
	(*PublicCreateBulkRequest)(nil),                   // 0: io.clbs.openhes.pbapi.PublicCreateBulkRequest
	(*JobDeviceList)(nil),                             // 1: io.clbs.openhes.pbapi.JobDeviceList
	(*JobCustomDeviceList)(nil),                       // 2: io.clbs.openhes.pbapi.JobCustomDeviceList
	(*JobCustomDevice)(nil),                           // 3: io.clbs.openhes.pbapi.JobCustomDevice
	(*pbdrivermodels.JobSettings)(nil),                // 4: io.clbs.openhes.pbdriver.JobSettings
	(*pbdrivermodels.JobAction)(nil),                  // 5: io.clbs.openhes.pbdriver.JobAction
	(*pbtaskmastermodels.JobDevice)(nil),              // 6: io.clbs.openhes.pbtaskmaster.JobDevice
	(*pbdataproxymodels.GetBulksReuqest)(nil),         // 7: io.clbs.openhes.pbdataproxy.GetBulksReuqest
	(*pbdataproxymodels.GetBulkRequest)(nil),          // 8: io.clbs.openhes.pbdataproxy.GetBulkRequest
	(*pbdataproxymodels.GetJobStatusRequest)(nil),     // 9: io.clbs.openhes.pbdataproxy.GetJobStatusRequest
	(*emptypb.Empty)(nil),                             // 10: google.protobuf.Empty
	(*pbdataproxymodels.GetBulksResponse)(nil),        // 11: io.clbs.openhes.pbdataproxy.GetBulksResponse
	(*pbdataproxymodels.GetBulkResponse)(nil),         // 12: io.clbs.openhes.pbdataproxy.GetBulkResponse
	(*pbdataproxymodels.GetJobStatusResponse)(nil),    // 13: io.clbs.openhes.pbdataproxy.GetJobStatusResponse
	(*pbdriveroperatormodels.GetDriversResponse)(nil), // 14: io.clbs.openhes.pbdriveroperator.GetDriversResponse
}
var file_pbapi_proto_depIdxs = []int32{
	1,  // 0: io.clbs.openhes.pbapi.PublicCreateBulkRequest.devices:type_name -> io.clbs.openhes.pbapi.JobDeviceList
	2,  // 1: io.clbs.openhes.pbapi.PublicCreateBulkRequest.custom_devices:type_name -> io.clbs.openhes.pbapi.JobCustomDeviceList
	4,  // 2: io.clbs.openhes.pbapi.PublicCreateBulkRequest.settings:type_name -> io.clbs.openhes.pbdriver.JobSettings
	5,  // 3: io.clbs.openhes.pbapi.PublicCreateBulkRequest.actions:type_name -> io.clbs.openhes.pbdriver.JobAction
	6,  // 4: io.clbs.openhes.pbapi.JobDeviceList.list:type_name -> io.clbs.openhes.pbtaskmaster.JobDevice
	3,  // 5: io.clbs.openhes.pbapi.JobCustomDeviceList.custom_devices:type_name -> io.clbs.openhes.pbapi.JobCustomDevice
	0,  // 6: io.clbs.openhes.pbapi.ApiService.CreateBulk:input_type -> io.clbs.openhes.pbapi.PublicCreateBulkRequest
	7,  // 7: io.clbs.openhes.pbapi.ApiService.GetBulks:input_type -> io.clbs.openhes.pbdataproxy.GetBulksReuqest
	8,  // 8: io.clbs.openhes.pbapi.ApiService.GetBulk:input_type -> io.clbs.openhes.pbdataproxy.GetBulkRequest
	9,  // 9: io.clbs.openhes.pbapi.ApiService.GetJobStatus:input_type -> io.clbs.openhes.pbdataproxy.GetJobStatusRequest
	10, // 10: io.clbs.openhes.pbapi.ApiService.GetDrivers:input_type -> google.protobuf.Empty
	10, // 11: io.clbs.openhes.pbapi.ApiService.CreateBulk:output_type -> google.protobuf.Empty
	11, // 12: io.clbs.openhes.pbapi.ApiService.GetBulks:output_type -> io.clbs.openhes.pbdataproxy.GetBulksResponse
	12, // 13: io.clbs.openhes.pbapi.ApiService.GetBulk:output_type -> io.clbs.openhes.pbdataproxy.GetBulkResponse
	13, // 14: io.clbs.openhes.pbapi.ApiService.GetJobStatus:output_type -> io.clbs.openhes.pbdataproxy.GetJobStatusResponse
	14, // 15: io.clbs.openhes.pbapi.ApiService.GetDrivers:output_type -> io.clbs.openhes.pbdriveroperator.GetDriversResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_pbapi_proto_init() }
func file_pbapi_proto_init() {
	if File_pbapi_proto != nil {
		return
	}
	file_pbapi_proto_msgTypes[0].OneofWrappers = []any{
		(*publicCreateBulkRequest_Devices)(nil),
		(*publicCreateBulkRequest_CustomDevices)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbapi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pbapi_proto_goTypes,
		DependencyIndexes: file_pbapi_proto_depIdxs,
		MessageInfos:      file_pbapi_proto_msgTypes,
	}.Build()
	File_pbapi_proto = out.File
	file_pbapi_proto_rawDesc = nil
	file_pbapi_proto_goTypes = nil
	file_pbapi_proto_depIdxs = nil
}
