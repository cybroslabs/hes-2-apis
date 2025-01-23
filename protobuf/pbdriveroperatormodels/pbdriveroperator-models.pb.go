// Editions version of proto3 file

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.3
// source: pbdriveroperator-models.proto

package pbdriveroperatormodels

import (
	_ "github.com/cybroslabs/hes-2-apis/protobuf/pbdrivermodels"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Taskmaster -> Driver Operator driver-scale change request message
type SetDriverScaleRequest struct {
	state                   protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_TaskmasterId *string                `protobuf:"bytes,1,opt,name=taskmaster_id,json=taskmasterId" json:"taskmaster_id,omitempty"`
	xxx_hidden_DriverType   *string                `protobuf:"bytes,2,opt,name=driver_type,json=driverType" json:"driver_type,omitempty"`
	xxx_hidden_Replicas     uint32                 `protobuf:"varint,3,opt,name=replicas" json:"replicas,omitempty"`
	XXX_raceDetectHookData  protoimpl.RaceDetectHookData
	XXX_presence            [1]uint32
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *SetDriverScaleRequest) Reset() {
	*x = SetDriverScaleRequest{}
	mi := &file_pbdriveroperator_models_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetDriverScaleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDriverScaleRequest) ProtoMessage() {}

func (x *SetDriverScaleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pbdriveroperator_models_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *SetDriverScaleRequest) GetTaskmasterId() string {
	if x != nil {
		if x.xxx_hidden_TaskmasterId != nil {
			return *x.xxx_hidden_TaskmasterId
		}
		return ""
	}
	return ""
}

func (x *SetDriverScaleRequest) GetDriverType() string {
	if x != nil {
		if x.xxx_hidden_DriverType != nil {
			return *x.xxx_hidden_DriverType
		}
		return ""
	}
	return ""
}

func (x *SetDriverScaleRequest) GetReplicas() uint32 {
	if x != nil {
		return x.xxx_hidden_Replicas
	}
	return 0
}

func (x *SetDriverScaleRequest) SetTaskmasterId(v string) {
	x.xxx_hidden_TaskmasterId = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 3)
}

func (x *SetDriverScaleRequest) SetDriverType(v string) {
	x.xxx_hidden_DriverType = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 3)
}

func (x *SetDriverScaleRequest) SetReplicas(v uint32) {
	x.xxx_hidden_Replicas = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 3)
}

func (x *SetDriverScaleRequest) HasTaskmasterId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *SetDriverScaleRequest) HasDriverType() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *SetDriverScaleRequest) HasReplicas() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *SetDriverScaleRequest) ClearTaskmasterId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_TaskmasterId = nil
}

func (x *SetDriverScaleRequest) ClearDriverType() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_DriverType = nil
}

func (x *SetDriverScaleRequest) ClearReplicas() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_Replicas = 0
}

type SetDriverScaleRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	TaskmasterId *string
	DriverType   *string
	Replicas     *uint32
}

func (b0 SetDriverScaleRequest_builder) Build() *SetDriverScaleRequest {
	m0 := &SetDriverScaleRequest{}
	b, x := &b0, m0
	_, _ = b, x
	if b.TaskmasterId != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 3)
		x.xxx_hidden_TaskmasterId = b.TaskmasterId
	}
	if b.DriverType != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 3)
		x.xxx_hidden_DriverType = b.DriverType
	}
	if b.Replicas != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 3)
		x.xxx_hidden_Replicas = *b.Replicas
	}
	return m0
}

// Taskmaster -> Driver Operator get driver-scale message
type GetDriverScaleRequest struct {
	state                   protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_TaskmasterId *string                `protobuf:"bytes,1,opt,name=taskmaster_id,json=taskmasterId" json:"taskmaster_id,omitempty"`
	xxx_hidden_DriverType   *string                `protobuf:"bytes,2,opt,name=driver_type,json=driverType" json:"driver_type,omitempty"`
	XXX_raceDetectHookData  protoimpl.RaceDetectHookData
	XXX_presence            [1]uint32
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *GetDriverScaleRequest) Reset() {
	*x = GetDriverScaleRequest{}
	mi := &file_pbdriveroperator_models_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDriverScaleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDriverScaleRequest) ProtoMessage() {}

func (x *GetDriverScaleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pbdriveroperator_models_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetDriverScaleRequest) GetTaskmasterId() string {
	if x != nil {
		if x.xxx_hidden_TaskmasterId != nil {
			return *x.xxx_hidden_TaskmasterId
		}
		return ""
	}
	return ""
}

func (x *GetDriverScaleRequest) GetDriverType() string {
	if x != nil {
		if x.xxx_hidden_DriverType != nil {
			return *x.xxx_hidden_DriverType
		}
		return ""
	}
	return ""
}

func (x *GetDriverScaleRequest) SetTaskmasterId(v string) {
	x.xxx_hidden_TaskmasterId = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 2)
}

func (x *GetDriverScaleRequest) SetDriverType(v string) {
	x.xxx_hidden_DriverType = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 2)
}

func (x *GetDriverScaleRequest) HasTaskmasterId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *GetDriverScaleRequest) HasDriverType() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *GetDriverScaleRequest) ClearTaskmasterId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_TaskmasterId = nil
}

func (x *GetDriverScaleRequest) ClearDriverType() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_DriverType = nil
}

type GetDriverScaleRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	TaskmasterId *string
	DriverType   *string
}

func (b0 GetDriverScaleRequest_builder) Build() *GetDriverScaleRequest {
	m0 := &GetDriverScaleRequest{}
	b, x := &b0, m0
	_, _ = b, x
	if b.TaskmasterId != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 2)
		x.xxx_hidden_TaskmasterId = b.TaskmasterId
	}
	if b.DriverType != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 2)
		x.xxx_hidden_DriverType = b.DriverType
	}
	return m0
}

type StartUpgradeRequest struct {
	state                        protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_DriverType        *string                `protobuf:"bytes,1,opt,name=driver_type,json=driverType" json:"driver_type,omitempty"`
	xxx_hidden_DriverVersion     *string                `protobuf:"bytes,2,opt,name=driver_version,json=driverVersion" json:"driver_version,omitempty"`
	xxx_hidden_DeviceregistryUrl *string                `protobuf:"bytes,3,opt,name=deviceregistry_url,json=deviceregistryUrl" json:"deviceregistry_url,omitempty"`
	XXX_raceDetectHookData       protoimpl.RaceDetectHookData
	XXX_presence                 [1]uint32
	unknownFields                protoimpl.UnknownFields
	sizeCache                    protoimpl.SizeCache
}

func (x *StartUpgradeRequest) Reset() {
	*x = StartUpgradeRequest{}
	mi := &file_pbdriveroperator_models_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartUpgradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartUpgradeRequest) ProtoMessage() {}

func (x *StartUpgradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pbdriveroperator_models_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *StartUpgradeRequest) GetDriverType() string {
	if x != nil {
		if x.xxx_hidden_DriverType != nil {
			return *x.xxx_hidden_DriverType
		}
		return ""
	}
	return ""
}

func (x *StartUpgradeRequest) GetDriverVersion() string {
	if x != nil {
		if x.xxx_hidden_DriverVersion != nil {
			return *x.xxx_hidden_DriverVersion
		}
		return ""
	}
	return ""
}

func (x *StartUpgradeRequest) GetDeviceregistryUrl() string {
	if x != nil {
		if x.xxx_hidden_DeviceregistryUrl != nil {
			return *x.xxx_hidden_DeviceregistryUrl
		}
		return ""
	}
	return ""
}

func (x *StartUpgradeRequest) SetDriverType(v string) {
	x.xxx_hidden_DriverType = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 3)
}

func (x *StartUpgradeRequest) SetDriverVersion(v string) {
	x.xxx_hidden_DriverVersion = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 3)
}

func (x *StartUpgradeRequest) SetDeviceregistryUrl(v string) {
	x.xxx_hidden_DeviceregistryUrl = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 2, 3)
}

func (x *StartUpgradeRequest) HasDriverType() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *StartUpgradeRequest) HasDriverVersion() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *StartUpgradeRequest) HasDeviceregistryUrl() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 2)
}

func (x *StartUpgradeRequest) ClearDriverType() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_DriverType = nil
}

func (x *StartUpgradeRequest) ClearDriverVersion() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_DriverVersion = nil
}

func (x *StartUpgradeRequest) ClearDeviceregistryUrl() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 2)
	x.xxx_hidden_DeviceregistryUrl = nil
}

type StartUpgradeRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	DriverType        *string
	DriverVersion     *string
	DeviceregistryUrl *string
}

func (b0 StartUpgradeRequest_builder) Build() *StartUpgradeRequest {
	m0 := &StartUpgradeRequest{}
	b, x := &b0, m0
	_, _ = b, x
	if b.DriverType != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 3)
		x.xxx_hidden_DriverType = b.DriverType
	}
	if b.DriverVersion != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 3)
		x.xxx_hidden_DriverVersion = b.DriverVersion
	}
	if b.DeviceregistryUrl != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 2, 3)
		x.xxx_hidden_DeviceregistryUrl = b.DeviceregistryUrl
	}
	return m0
}

// API -> Taskmaster request to get driver templates
type GetDriverTemplatesRequest struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_DriverType  *string                `protobuf:"bytes,1,opt,name=driver_type,json=driverType" json:"driver_type,omitempty"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *GetDriverTemplatesRequest) Reset() {
	*x = GetDriverTemplatesRequest{}
	mi := &file_pbdriveroperator_models_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDriverTemplatesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDriverTemplatesRequest) ProtoMessage() {}

func (x *GetDriverTemplatesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pbdriveroperator_models_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetDriverTemplatesRequest) GetDriverType() string {
	if x != nil {
		if x.xxx_hidden_DriverType != nil {
			return *x.xxx_hidden_DriverType
		}
		return ""
	}
	return ""
}

func (x *GetDriverTemplatesRequest) SetDriverType(v string) {
	x.xxx_hidden_DriverType = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *GetDriverTemplatesRequest) HasDriverType() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *GetDriverTemplatesRequest) ClearDriverType() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_DriverType = nil
}

type GetDriverTemplatesRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	DriverType *string
}

func (b0 GetDriverTemplatesRequest_builder) Build() *GetDriverTemplatesRequest {
	m0 := &GetDriverTemplatesRequest{}
	b, x := &b0, m0
	_, _ = b, x
	if b.DriverType != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_DriverType = b.DriverType
	}
	return m0
}

// Sub-message containing driver info
type DriverInfo struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_DriverType  *string                `protobuf:"bytes,1,opt,name=driver_type,json=driverType" json:"driver_type,omitempty"`
	xxx_hidden_Version     *string                `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *DriverInfo) Reset() {
	*x = DriverInfo{}
	mi := &file_pbdriveroperator_models_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DriverInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverInfo) ProtoMessage() {}

func (x *DriverInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pbdriveroperator_models_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *DriverInfo) GetDriverType() string {
	if x != nil {
		if x.xxx_hidden_DriverType != nil {
			return *x.xxx_hidden_DriverType
		}
		return ""
	}
	return ""
}

func (x *DriverInfo) GetVersion() string {
	if x != nil {
		if x.xxx_hidden_Version != nil {
			return *x.xxx_hidden_Version
		}
		return ""
	}
	return ""
}

func (x *DriverInfo) SetDriverType(v string) {
	x.xxx_hidden_DriverType = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 2)
}

func (x *DriverInfo) SetVersion(v string) {
	x.xxx_hidden_Version = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 1, 2)
}

func (x *DriverInfo) HasDriverType() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *DriverInfo) HasVersion() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 1)
}

func (x *DriverInfo) ClearDriverType() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_DriverType = nil
}

func (x *DriverInfo) ClearVersion() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 1)
	x.xxx_hidden_Version = nil
}

type DriverInfo_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	DriverType *string
	Version    *string
}

func (b0 DriverInfo_builder) Build() *DriverInfo {
	m0 := &DriverInfo{}
	b, x := &b0, m0
	_, _ = b, x
	if b.DriverType != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 2)
		x.xxx_hidden_DriverType = b.DriverType
	}
	if b.Version != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 1, 2)
		x.xxx_hidden_Version = b.Version
	}
	return m0
}

// Taskmaster -> API get drivers response message
type ListOfDriverInfo struct {
	state            protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Items *[]*DriverInfo         `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *ListOfDriverInfo) Reset() {
	*x = ListOfDriverInfo{}
	mi := &file_pbdriveroperator_models_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListOfDriverInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOfDriverInfo) ProtoMessage() {}

func (x *ListOfDriverInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pbdriveroperator_models_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ListOfDriverInfo) GetItems() []*DriverInfo {
	if x != nil {
		if x.xxx_hidden_Items != nil {
			return *x.xxx_hidden_Items
		}
	}
	return nil
}

func (x *ListOfDriverInfo) SetItems(v []*DriverInfo) {
	x.xxx_hidden_Items = &v
}

type ListOfDriverInfo_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Items []*DriverInfo
}

func (b0 ListOfDriverInfo_builder) Build() *ListOfDriverInfo {
	m0 := &ListOfDriverInfo{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Items = &b.Items
	return m0
}

var File_pbdriveroperator_models_proto protoreflect.FileDescriptor

var file_pbdriveroperator_models_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x6f, 0x72, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x20, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73,
	0x2e, 0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x70, 0x62, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73,
	0x22, 0x5d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x63, 0x61,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x61, 0x73,
	0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x8c, 0x01, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x2d, 0x0a, 0x12, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x55, 0x72, 0x6c, 0x22, 0x3c,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x54, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x47, 0x0a, 0x0a,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x56, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x44,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x42, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c,
	0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x70, 0x62, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x44, 0x72, 0x69, 0x76,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x42, 0x5a,
	0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x79, 0x62, 0x72,
	0x6f, 0x73, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x68, 0x65, 0x73, 0x2d, 0x32, 0x2d, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x62, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
}

var file_pbdriveroperator_models_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pbdriveroperator_models_proto_goTypes = []any{
	(*SetDriverScaleRequest)(nil),     // 0: io.clbs.openhes.pbdriveroperator.SetDriverScaleRequest
	(*GetDriverScaleRequest)(nil),     // 1: io.clbs.openhes.pbdriveroperator.GetDriverScaleRequest
	(*StartUpgradeRequest)(nil),       // 2: io.clbs.openhes.pbdriveroperator.StartUpgradeRequest
	(*GetDriverTemplatesRequest)(nil), // 3: io.clbs.openhes.pbdriveroperator.GetDriverTemplatesRequest
	(*DriverInfo)(nil),                // 4: io.clbs.openhes.pbdriveroperator.DriverInfo
	(*ListOfDriverInfo)(nil),          // 5: io.clbs.openhes.pbdriveroperator.ListOfDriverInfo
}
var file_pbdriveroperator_models_proto_depIdxs = []int32{
	4, // 0: io.clbs.openhes.pbdriveroperator.ListOfDriverInfo.items:type_name -> io.clbs.openhes.pbdriveroperator.DriverInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pbdriveroperator_models_proto_init() }
func file_pbdriveroperator_models_proto_init() {
	if File_pbdriveroperator_models_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbdriveroperator_models_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbdriveroperator_models_proto_goTypes,
		DependencyIndexes: file_pbdriveroperator_models_proto_depIdxs,
		MessageInfos:      file_pbdriveroperator_models_proto_msgTypes,
	}.Build()
	File_pbdriveroperator_models_proto = out.File
	file_pbdriveroperator_models_proto_rawDesc = nil
	file_pbdriveroperator_models_proto_goTypes = nil
	file_pbdriveroperator_models_proto_depIdxs = nil
}
