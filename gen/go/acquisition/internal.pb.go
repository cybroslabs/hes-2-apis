// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: acquisition/internal.proto

package acquisition

import (
	common "github.com/cybroslabs/hes-2-apis/gen/go/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SetCacheRequest struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Key         *string                `protobuf:"bytes,1,opt,name=key"`
	xxx_hidden_Value       *common.FieldValue     `protobuf:"bytes,2,opt,name=value"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *SetCacheRequest) Reset() {
	*x = SetCacheRequest{}
	mi := &file_acquisition_internal_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetCacheRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetCacheRequest) ProtoMessage() {}

func (x *SetCacheRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acquisition_internal_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *SetCacheRequest) GetKey() string {
	if x != nil {
		if x.xxx_hidden_Key != nil {
			return *x.xxx_hidden_Key
		}
		return ""
	}
	return ""
}

func (x *SetCacheRequest) GetValue() *common.FieldValue {
	if x != nil {
		return x.xxx_hidden_Value
	}
	return nil
}

func (x *SetCacheRequest) SetKey(v string) {
	x.xxx_hidden_Key = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 2)
}

func (x *SetCacheRequest) SetValue(v *common.FieldValue) {
	x.xxx_hidden_Value = v
}

func (x *SetCacheRequest) HasKey() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *SetCacheRequest) HasValue() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Value != nil
}

func (x *SetCacheRequest) ClearKey() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Key = nil
}

func (x *SetCacheRequest) ClearValue() {
	x.xxx_hidden_Value = nil
}

type SetCacheRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Key   *string
	Value *common.FieldValue
}

func (b0 SetCacheRequest_builder) Build() *SetCacheRequest {
	m0 := &SetCacheRequest{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Key != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 2)
		x.xxx_hidden_Key = b.Key
	}
	x.xxx_hidden_Value = b.Value
	return m0
}

type GetCacheRequest struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Key         *string                `protobuf:"bytes,1,opt,name=key"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *GetCacheRequest) Reset() {
	*x = GetCacheRequest{}
	mi := &file_acquisition_internal_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCacheRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCacheRequest) ProtoMessage() {}

func (x *GetCacheRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acquisition_internal_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetCacheRequest) GetKey() string {
	if x != nil {
		if x.xxx_hidden_Key != nil {
			return *x.xxx_hidden_Key
		}
		return ""
	}
	return ""
}

func (x *GetCacheRequest) SetKey(v string) {
	x.xxx_hidden_Key = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *GetCacheRequest) HasKey() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *GetCacheRequest) ClearKey() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_Key = nil
}

type GetCacheRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Key *string
}

func (b0 GetCacheRequest_builder) Build() *GetCacheRequest {
	m0 := &GetCacheRequest{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Key != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_Key = b.Key
	}
	return m0
}

type GetCacheResponse struct {
	state            protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Value *common.FieldValue     `protobuf:"bytes,2,opt,name=value"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *GetCacheResponse) Reset() {
	*x = GetCacheResponse{}
	mi := &file_acquisition_internal_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCacheResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCacheResponse) ProtoMessage() {}

func (x *GetCacheResponse) ProtoReflect() protoreflect.Message {
	mi := &file_acquisition_internal_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetCacheResponse) GetValue() *common.FieldValue {
	if x != nil {
		return x.xxx_hidden_Value
	}
	return nil
}

func (x *GetCacheResponse) SetValue(v *common.FieldValue) {
	x.xxx_hidden_Value = v
}

func (x *GetCacheResponse) HasValue() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Value != nil
}

func (x *GetCacheResponse) ClearValue() {
	x.xxx_hidden_Value = nil
}

type GetCacheResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Value *common.FieldValue
}

func (b0 GetCacheResponse_builder) Build() *GetCacheResponse {
	m0 := &GetCacheResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Value = b.Value
	return m0
}

type StartJobsRequest struct {
	state                     protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_ConnectionInfo *ConnectionInfo        `protobuf:"bytes,2,opt,name=connection_info,json=connectionInfo"`
	xxx_hidden_Jobs           *[]*StartJobData       `protobuf:"bytes,1,rep,name=jobs"`
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *StartJobsRequest) Reset() {
	*x = StartJobsRequest{}
	mi := &file_acquisition_internal_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartJobsRequest) ProtoMessage() {}

func (x *StartJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acquisition_internal_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *StartJobsRequest) GetConnectionInfo() *ConnectionInfo {
	if x != nil {
		return x.xxx_hidden_ConnectionInfo
	}
	return nil
}

func (x *StartJobsRequest) GetJobs() []*StartJobData {
	if x != nil {
		if x.xxx_hidden_Jobs != nil {
			return *x.xxx_hidden_Jobs
		}
	}
	return nil
}

func (x *StartJobsRequest) SetConnectionInfo(v *ConnectionInfo) {
	x.xxx_hidden_ConnectionInfo = v
}

func (x *StartJobsRequest) SetJobs(v []*StartJobData) {
	x.xxx_hidden_Jobs = &v
}

func (x *StartJobsRequest) HasConnectionInfo() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_ConnectionInfo != nil
}

func (x *StartJobsRequest) ClearConnectionInfo() {
	x.xxx_hidden_ConnectionInfo = nil
}

type StartJobsRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	ConnectionInfo *ConnectionInfo
	Jobs           []*StartJobData
}

func (b0 StartJobsRequest_builder) Build() *StartJobsRequest {
	m0 := &StartJobsRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_ConnectionInfo = b.ConnectionInfo
	x.xxx_hidden_Jobs = &b.Jobs
	return m0
}

// Driver -> Taskmaster job/action progress update message
type ProgressUpdate struct {
	state               protoimpl.MessageState    `protogen:"opaque.v1"`
	xxx_hidden_Progress isProgressUpdate_Progress `protobuf_oneof:"progress"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *ProgressUpdate) Reset() {
	*x = ProgressUpdate{}
	mi := &file_acquisition_internal_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProgressUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProgressUpdate) ProtoMessage() {}

func (x *ProgressUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_acquisition_internal_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *ProgressUpdate) GetJob() *JobProgressUpdate {
	if x != nil {
		if x, ok := x.xxx_hidden_Progress.(*progressUpdate_Job); ok {
			return x.Job
		}
	}
	return nil
}

func (x *ProgressUpdate) GetAction() *ActionProgressUpdate {
	if x != nil {
		if x, ok := x.xxx_hidden_Progress.(*progressUpdate_Action); ok {
			return x.Action
		}
	}
	return nil
}

func (x *ProgressUpdate) SetJob(v *JobProgressUpdate) {
	if v == nil {
		x.xxx_hidden_Progress = nil
		return
	}
	x.xxx_hidden_Progress = &progressUpdate_Job{v}
}

func (x *ProgressUpdate) SetAction(v *ActionProgressUpdate) {
	if v == nil {
		x.xxx_hidden_Progress = nil
		return
	}
	x.xxx_hidden_Progress = &progressUpdate_Action{v}
}

func (x *ProgressUpdate) HasProgress() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Progress != nil
}

func (x *ProgressUpdate) HasJob() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_Progress.(*progressUpdate_Job)
	return ok
}

func (x *ProgressUpdate) HasAction() bool {
	if x == nil {
		return false
	}
	_, ok := x.xxx_hidden_Progress.(*progressUpdate_Action)
	return ok
}

func (x *ProgressUpdate) ClearProgress() {
	x.xxx_hidden_Progress = nil
}

func (x *ProgressUpdate) ClearJob() {
	if _, ok := x.xxx_hidden_Progress.(*progressUpdate_Job); ok {
		x.xxx_hidden_Progress = nil
	}
}

func (x *ProgressUpdate) ClearAction() {
	if _, ok := x.xxx_hidden_Progress.(*progressUpdate_Action); ok {
		x.xxx_hidden_Progress = nil
	}
}

const ProgressUpdate_Progress_not_set_case case_ProgressUpdate_Progress = 0
const ProgressUpdate_Job_case case_ProgressUpdate_Progress = 1
const ProgressUpdate_Action_case case_ProgressUpdate_Progress = 2

func (x *ProgressUpdate) WhichProgress() case_ProgressUpdate_Progress {
	if x == nil {
		return ProgressUpdate_Progress_not_set_case
	}
	switch x.xxx_hidden_Progress.(type) {
	case *progressUpdate_Job:
		return ProgressUpdate_Job_case
	case *progressUpdate_Action:
		return ProgressUpdate_Action_case
	default:
		return ProgressUpdate_Progress_not_set_case
	}
}

type ProgressUpdate_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	// The oneof field containing the progress update data - either for job or for action.

	// Fields of oneof xxx_hidden_Progress:
	Job    *JobProgressUpdate
	Action *ActionProgressUpdate
	// -- end of xxx_hidden_Progress
}

func (b0 ProgressUpdate_builder) Build() *ProgressUpdate {
	m0 := &ProgressUpdate{}
	b, x := &b0, m0
	_, _ = b, x
	if b.Job != nil {
		x.xxx_hidden_Progress = &progressUpdate_Job{b.Job}
	}
	if b.Action != nil {
		x.xxx_hidden_Progress = &progressUpdate_Action{b.Action}
	}
	return m0
}

type case_ProgressUpdate_Progress protoreflect.FieldNumber

func (x case_ProgressUpdate_Progress) String() string {
	md := file_acquisition_internal_proto_msgTypes[4].Descriptor()
	if x == 0 {
		return "not set"
	}
	return protoimpl.X.MessageFieldStringOf(md, protoreflect.FieldNumber(x))
}

type isProgressUpdate_Progress interface {
	isProgressUpdate_Progress()
}

type progressUpdate_Job struct {
	Job *JobProgressUpdate `protobuf:"bytes,1,opt,name=job,oneof"` // The job progress update data. It shall be called once and only once for each job. It shall be also called as the last message in the stream; other updates will be ignored after this.
}

type progressUpdate_Action struct {
	Action *ActionProgressUpdate `protobuf:"bytes,2,opt,name=action,oneof"` // The action progress update data. It shall be called for each action in the job.
}

func (*progressUpdate_Job) isProgressUpdate_Progress() {}

func (*progressUpdate_Action) isProgressUpdate_Progress() {}

type StartUpgradeRequest struct {
	state                        protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_DriverType        *string                `protobuf:"bytes,1,opt,name=driver_type,json=driverType"`
	xxx_hidden_DriverVersion     *string                `protobuf:"bytes,2,opt,name=driver_version,json=driverVersion"`
	xxx_hidden_DeviceregistryUrl *string                `protobuf:"bytes,3,opt,name=deviceregistry_url,json=deviceregistryUrl"`
	XXX_raceDetectHookData       protoimpl.RaceDetectHookData
	XXX_presence                 [1]uint32
	unknownFields                protoimpl.UnknownFields
	sizeCache                    protoimpl.SizeCache
}

func (x *StartUpgradeRequest) Reset() {
	*x = StartUpgradeRequest{}
	mi := &file_acquisition_internal_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartUpgradeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartUpgradeRequest) ProtoMessage() {}

func (x *StartUpgradeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acquisition_internal_proto_msgTypes[5]
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

type QueueJobsRequest struct {
	state           protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Spec *[]*JobSpec            `protobuf:"bytes,1,rep,name=spec"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *QueueJobsRequest) Reset() {
	*x = QueueJobsRequest{}
	mi := &file_acquisition_internal_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueueJobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueJobsRequest) ProtoMessage() {}

func (x *QueueJobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acquisition_internal_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *QueueJobsRequest) GetSpec() []*JobSpec {
	if x != nil {
		if x.xxx_hidden_Spec != nil {
			return *x.xxx_hidden_Spec
		}
	}
	return nil
}

func (x *QueueJobsRequest) SetSpec(v []*JobSpec) {
	x.xxx_hidden_Spec = &v
}

type QueueJobsRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Spec []*JobSpec
}

func (b0 QueueJobsRequest_builder) Build() *QueueJobsRequest {
	m0 := &QueueJobsRequest{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Spec = &b.Spec
	return m0
}

type GetJobResponse struct {
	state             protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Status *JobStatus             `protobuf:"bytes,1,opt,name=status"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *GetJobResponse) Reset() {
	*x = GetJobResponse{}
	mi := &file_acquisition_internal_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobResponse) ProtoMessage() {}

func (x *GetJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_acquisition_internal_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *GetJobResponse) GetStatus() *JobStatus {
	if x != nil {
		return x.xxx_hidden_Status
	}
	return nil
}

func (x *GetJobResponse) SetStatus(v *JobStatus) {
	x.xxx_hidden_Status = v
}

func (x *GetJobResponse) HasStatus() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Status != nil
}

func (x *GetJobResponse) ClearStatus() {
	x.xxx_hidden_Status = nil
}

type GetJobResponse_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Status *JobStatus
}

func (b0 GetJobResponse_builder) Build() *GetJobResponse {
	m0 := &GetJobResponse{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Status = b.Status
	return m0
}

type SetDriverScaleRequest struct {
	state                   protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_TaskmasterId *string                `protobuf:"bytes,1,opt,name=taskmaster_id,json=taskmasterId"`
	xxx_hidden_DriverType   *string                `protobuf:"bytes,2,opt,name=driver_type,json=driverType"`
	xxx_hidden_Replicas     uint32                 `protobuf:"varint,3,opt,name=replicas"`
	XXX_raceDetectHookData  protoimpl.RaceDetectHookData
	XXX_presence            [1]uint32
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *SetDriverScaleRequest) Reset() {
	*x = SetDriverScaleRequest{}
	mi := &file_acquisition_internal_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetDriverScaleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDriverScaleRequest) ProtoMessage() {}

func (x *SetDriverScaleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acquisition_internal_proto_msgTypes[8]
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

type GetDriverScaleRequest struct {
	state                   protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_TaskmasterId *string                `protobuf:"bytes,1,opt,name=taskmaster_id,json=taskmasterId"`
	xxx_hidden_DriverType   *string                `protobuf:"bytes,2,opt,name=driver_type,json=driverType"`
	XXX_raceDetectHookData  protoimpl.RaceDetectHookData
	XXX_presence            [1]uint32
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *GetDriverScaleRequest) Reset() {
	*x = GetDriverScaleRequest{}
	mi := &file_acquisition_internal_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDriverScaleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDriverScaleRequest) ProtoMessage() {}

func (x *GetDriverScaleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acquisition_internal_proto_msgTypes[9]
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

type MapDeviceConnectionInfo struct {
	state              protoimpl.MessageState           `protogen:"opaque.v1"`
	xxx_hidden_Devices map[string]*ListOfConnectionInfo `protobuf:"bytes,1,rep,name=devices" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *MapDeviceConnectionInfo) Reset() {
	*x = MapDeviceConnectionInfo{}
	mi := &file_acquisition_internal_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MapDeviceConnectionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapDeviceConnectionInfo) ProtoMessage() {}

func (x *MapDeviceConnectionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_acquisition_internal_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *MapDeviceConnectionInfo) GetDevices() map[string]*ListOfConnectionInfo {
	if x != nil {
		return x.xxx_hidden_Devices
	}
	return nil
}

func (x *MapDeviceConnectionInfo) SetDevices(v map[string]*ListOfConnectionInfo) {
	x.xxx_hidden_Devices = v
}

type MapDeviceConnectionInfo_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Devices map[string]*ListOfConnectionInfo
}

func (b0 MapDeviceConnectionInfo_builder) Build() *MapDeviceConnectionInfo {
	m0 := &MapDeviceConnectionInfo{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Devices = b.Devices
	return m0
}

type SetDeviceInfoRequest struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_DeviceId    *string                `protobuf:"bytes,1,opt,name=device_id,json=deviceId"`
	xxx_hidden_Info        *DeviceInfo            `protobuf:"bytes,2,opt,name=info"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *SetDeviceInfoRequest) Reset() {
	*x = SetDeviceInfoRequest{}
	mi := &file_acquisition_internal_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetDeviceInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDeviceInfoRequest) ProtoMessage() {}

func (x *SetDeviceInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_acquisition_internal_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *SetDeviceInfoRequest) GetDeviceId() string {
	if x != nil {
		if x.xxx_hidden_DeviceId != nil {
			return *x.xxx_hidden_DeviceId
		}
		return ""
	}
	return ""
}

func (x *SetDeviceInfoRequest) GetInfo() *DeviceInfo {
	if x != nil {
		return x.xxx_hidden_Info
	}
	return nil
}

func (x *SetDeviceInfoRequest) SetDeviceId(v string) {
	x.xxx_hidden_DeviceId = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 2)
}

func (x *SetDeviceInfoRequest) SetInfo(v *DeviceInfo) {
	x.xxx_hidden_Info = v
}

func (x *SetDeviceInfoRequest) HasDeviceId() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *SetDeviceInfoRequest) HasInfo() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Info != nil
}

func (x *SetDeviceInfoRequest) ClearDeviceId() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_DeviceId = nil
}

func (x *SetDeviceInfoRequest) ClearInfo() {
	x.xxx_hidden_Info = nil
}

type SetDeviceInfoRequest_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	DeviceId *string
	Info     *DeviceInfo
}

func (b0 SetDeviceInfoRequest_builder) Build() *SetDeviceInfoRequest {
	m0 := &SetDeviceInfoRequest{}
	b, x := &b0, m0
	_, _ = b, x
	if b.DeviceId != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 2)
		x.xxx_hidden_DeviceId = b.DeviceId
	}
	x.xxx_hidden_Info = b.Info
	return m0
}

var File_acquisition_internal_proto protoreflect.FileDescriptor

const file_acquisition_internal_proto_rawDesc = "" +
	"\n" +
	"\x1aacquisition/internal.proto\x12\"io.clbs.openhes.models.acquisition\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x18acquisition/shared.proto\x1a\x13common/fields.proto\"d\n" +
	"\x0fSetCacheRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12?\n" +
	"\x05value\x18\x02 \x01(\v2).io.clbs.openhes.models.common.FieldValueR\x05value\"#\n" +
	"\x0fGetCacheRequest\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\"S\n" +
	"\x10GetCacheResponse\x12?\n" +
	"\x05value\x18\x02 \x01(\v2).io.clbs.openhes.models.common.FieldValueR\x05value\"\xb5\x01\n" +
	"\x10StartJobsRequest\x12[\n" +
	"\x0fconnection_info\x18\x02 \x01(\v22.io.clbs.openhes.models.acquisition.ConnectionInfoR\x0econnectionInfo\x12D\n" +
	"\x04jobs\x18\x01 \x03(\v20.io.clbs.openhes.models.acquisition.StartJobDataR\x04jobs\"\xbb\x01\n" +
	"\x0eProgressUpdate\x12I\n" +
	"\x03job\x18\x01 \x01(\v25.io.clbs.openhes.models.acquisition.JobProgressUpdateH\x00R\x03job\x12R\n" +
	"\x06action\x18\x02 \x01(\v28.io.clbs.openhes.models.acquisition.ActionProgressUpdateH\x00R\x06actionB\n" +
	"\n" +
	"\bprogress\"\x8c\x01\n" +
	"\x13StartUpgradeRequest\x12\x1f\n" +
	"\vdriver_type\x18\x01 \x01(\tR\n" +
	"driverType\x12%\n" +
	"\x0edriver_version\x18\x02 \x01(\tR\rdriverVersion\x12-\n" +
	"\x12deviceregistry_url\x18\x03 \x01(\tR\x11deviceregistryUrl\"S\n" +
	"\x10QueueJobsRequest\x12?\n" +
	"\x04spec\x18\x01 \x03(\v2+.io.clbs.openhes.models.acquisition.JobSpecR\x04spec\"W\n" +
	"\x0eGetJobResponse\x12E\n" +
	"\x06status\x18\x01 \x01(\v2-.io.clbs.openhes.models.acquisition.JobStatusR\x06status\"y\n" +
	"\x15SetDriverScaleRequest\x12#\n" +
	"\rtaskmaster_id\x18\x01 \x01(\tR\ftaskmasterId\x12\x1f\n" +
	"\vdriver_type\x18\x02 \x01(\tR\n" +
	"driverType\x12\x1a\n" +
	"\breplicas\x18\x03 \x01(\rR\breplicas\"]\n" +
	"\x15GetDriverScaleRequest\x12#\n" +
	"\rtaskmaster_id\x18\x01 \x01(\tR\ftaskmasterId\x12\x1f\n" +
	"\vdriver_type\x18\x02 \x01(\tR\n" +
	"driverType\"\xf3\x01\n" +
	"\x17MapDeviceConnectionInfo\x12b\n" +
	"\adevices\x18\x01 \x03(\v2H.io.clbs.openhes.models.acquisition.MapDeviceConnectionInfo.DevicesEntryR\adevices\x1at\n" +
	"\fDevicesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12N\n" +
	"\x05value\x18\x02 \x01(\v28.io.clbs.openhes.models.acquisition.ListOfConnectionInfoR\x05value:\x028\x01\"w\n" +
	"\x14SetDeviceInfoRequest\x12\x1b\n" +
	"\tdevice_id\x18\x01 \x01(\tR\bdeviceId\x12B\n" +
	"\x04info\x18\x02 \x01(\v2..io.clbs.openhes.models.acquisition.DeviceInfoR\x04infoB5Z3github.com/cybroslabs/hes-2-apis/gen/go/acquisitionb\beditionsp\xe8\a"

var file_acquisition_internal_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_acquisition_internal_proto_goTypes = []any{
	(*SetCacheRequest)(nil),         // 0: io.clbs.openhes.models.acquisition.SetCacheRequest
	(*GetCacheRequest)(nil),         // 1: io.clbs.openhes.models.acquisition.GetCacheRequest
	(*GetCacheResponse)(nil),        // 2: io.clbs.openhes.models.acquisition.GetCacheResponse
	(*StartJobsRequest)(nil),        // 3: io.clbs.openhes.models.acquisition.StartJobsRequest
	(*ProgressUpdate)(nil),          // 4: io.clbs.openhes.models.acquisition.ProgressUpdate
	(*StartUpgradeRequest)(nil),     // 5: io.clbs.openhes.models.acquisition.StartUpgradeRequest
	(*QueueJobsRequest)(nil),        // 6: io.clbs.openhes.models.acquisition.QueueJobsRequest
	(*GetJobResponse)(nil),          // 7: io.clbs.openhes.models.acquisition.GetJobResponse
	(*SetDriverScaleRequest)(nil),   // 8: io.clbs.openhes.models.acquisition.SetDriverScaleRequest
	(*GetDriverScaleRequest)(nil),   // 9: io.clbs.openhes.models.acquisition.GetDriverScaleRequest
	(*MapDeviceConnectionInfo)(nil), // 10: io.clbs.openhes.models.acquisition.MapDeviceConnectionInfo
	(*SetDeviceInfoRequest)(nil),    // 11: io.clbs.openhes.models.acquisition.SetDeviceInfoRequest
	nil,                             // 12: io.clbs.openhes.models.acquisition.MapDeviceConnectionInfo.DevicesEntry
	(*common.FieldValue)(nil),       // 13: io.clbs.openhes.models.common.FieldValue
	(*ConnectionInfo)(nil),          // 14: io.clbs.openhes.models.acquisition.ConnectionInfo
	(*StartJobData)(nil),            // 15: io.clbs.openhes.models.acquisition.StartJobData
	(*JobProgressUpdate)(nil),       // 16: io.clbs.openhes.models.acquisition.JobProgressUpdate
	(*ActionProgressUpdate)(nil),    // 17: io.clbs.openhes.models.acquisition.ActionProgressUpdate
	(*JobSpec)(nil),                 // 18: io.clbs.openhes.models.acquisition.JobSpec
	(*JobStatus)(nil),               // 19: io.clbs.openhes.models.acquisition.JobStatus
	(*DeviceInfo)(nil),              // 20: io.clbs.openhes.models.acquisition.DeviceInfo
	(*ListOfConnectionInfo)(nil),    // 21: io.clbs.openhes.models.acquisition.ListOfConnectionInfo
}
var file_acquisition_internal_proto_depIdxs = []int32{
	13, // 0: io.clbs.openhes.models.acquisition.SetCacheRequest.value:type_name -> io.clbs.openhes.models.common.FieldValue
	13, // 1: io.clbs.openhes.models.acquisition.GetCacheResponse.value:type_name -> io.clbs.openhes.models.common.FieldValue
	14, // 2: io.clbs.openhes.models.acquisition.StartJobsRequest.connection_info:type_name -> io.clbs.openhes.models.acquisition.ConnectionInfo
	15, // 3: io.clbs.openhes.models.acquisition.StartJobsRequest.jobs:type_name -> io.clbs.openhes.models.acquisition.StartJobData
	16, // 4: io.clbs.openhes.models.acquisition.ProgressUpdate.job:type_name -> io.clbs.openhes.models.acquisition.JobProgressUpdate
	17, // 5: io.clbs.openhes.models.acquisition.ProgressUpdate.action:type_name -> io.clbs.openhes.models.acquisition.ActionProgressUpdate
	18, // 6: io.clbs.openhes.models.acquisition.QueueJobsRequest.spec:type_name -> io.clbs.openhes.models.acquisition.JobSpec
	19, // 7: io.clbs.openhes.models.acquisition.GetJobResponse.status:type_name -> io.clbs.openhes.models.acquisition.JobStatus
	12, // 8: io.clbs.openhes.models.acquisition.MapDeviceConnectionInfo.devices:type_name -> io.clbs.openhes.models.acquisition.MapDeviceConnectionInfo.DevicesEntry
	20, // 9: io.clbs.openhes.models.acquisition.SetDeviceInfoRequest.info:type_name -> io.clbs.openhes.models.acquisition.DeviceInfo
	21, // 10: io.clbs.openhes.models.acquisition.MapDeviceConnectionInfo.DevicesEntry.value:type_name -> io.clbs.openhes.models.acquisition.ListOfConnectionInfo
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_acquisition_internal_proto_init() }
func file_acquisition_internal_proto_init() {
	if File_acquisition_internal_proto != nil {
		return
	}
	file_acquisition_shared_proto_init()
	file_acquisition_internal_proto_msgTypes[4].OneofWrappers = []any{
		(*progressUpdate_Job)(nil),
		(*progressUpdate_Action)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_acquisition_internal_proto_rawDesc), len(file_acquisition_internal_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_acquisition_internal_proto_goTypes,
		DependencyIndexes: file_acquisition_internal_proto_depIdxs,
		MessageInfos:      file_acquisition_internal_proto_msgTypes,
	}.Build()
	File_acquisition_internal_proto = out.File
	file_acquisition_internal_proto_goTypes = nil
	file_acquisition_internal_proto_depIdxs = nil
}
