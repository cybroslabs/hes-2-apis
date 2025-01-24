// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
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

var File_acquisition_internal_proto protoreflect.FileDescriptor

var file_acquisition_internal_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x69, 0x6f,
	0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18,
	0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a,
	0x0f, 0x53, 0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x3f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x23, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x61, 0x63, 0x68, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x53, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x69, 0x6f,
	0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xb5, 0x01,
	0x0a, 0x10, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x5b, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x69, 0x6f,
	0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x44, 0x0a, 0x04, 0x6a, 0x6f, 0x62, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e,
	0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4a, 0x6f, 0x62, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x04, 0x6a, 0x6f, 0x62, 0x73, 0x22, 0xbb, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x49, 0x0a, 0x03, 0x6a, 0x6f, 0x62, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e,
	0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61,
	0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4a, 0x6f, 0x62, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x03,
	0x6a, 0x6f, 0x62, 0x12, 0x52, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70,
	0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71,
	0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x8c, 0x01, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x55, 0x70, 0x67,
	0x72, 0x61, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x12, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x11, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x55,
	0x72, 0x6c, 0x22, 0x53, 0x0a, 0x10, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4a, 0x6f, 0x62, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x69, 0x6f, 0x2e, 0x63, 0x6c, 0x62, 0x73, 0x2e, 0x6f,
	0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x61, 0x63,
	0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x70, 0x65,
	0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x22, 0x57, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4a, 0x6f,
	0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x69, 0x6f, 0x2e, 0x63,
	0x6c, 0x62, 0x73, 0x2e, 0x6f, 0x70, 0x65, 0x6e, 0x68, 0x65, 0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4a,
	0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x79, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x63, 0x61,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x61, 0x73,
	0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x22, 0x5d, 0x0a, 0x15, 0x47,
	0x65, 0x74, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x61, 0x73,
	0x6b, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x79, 0x62, 0x72, 0x6f, 0x73, 0x6c,
	0x61, 0x62, 0x73, 0x2f, 0x68, 0x65, 0x73, 0x2d, 0x32, 0x2d, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x63, 0x71, 0x75, 0x69, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
}

var file_acquisition_internal_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_acquisition_internal_proto_goTypes = []any{
	(*SetCacheRequest)(nil),       // 0: io.clbs.openhes.models.acquisition.SetCacheRequest
	(*GetCacheRequest)(nil),       // 1: io.clbs.openhes.models.acquisition.GetCacheRequest
	(*GetCacheResponse)(nil),      // 2: io.clbs.openhes.models.acquisition.GetCacheResponse
	(*StartJobsRequest)(nil),      // 3: io.clbs.openhes.models.acquisition.StartJobsRequest
	(*ProgressUpdate)(nil),        // 4: io.clbs.openhes.models.acquisition.ProgressUpdate
	(*StartUpgradeRequest)(nil),   // 5: io.clbs.openhes.models.acquisition.StartUpgradeRequest
	(*QueueJobsRequest)(nil),      // 6: io.clbs.openhes.models.acquisition.QueueJobsRequest
	(*GetJobResponse)(nil),        // 7: io.clbs.openhes.models.acquisition.GetJobResponse
	(*SetDriverScaleRequest)(nil), // 8: io.clbs.openhes.models.acquisition.SetDriverScaleRequest
	(*GetDriverScaleRequest)(nil), // 9: io.clbs.openhes.models.acquisition.GetDriverScaleRequest
	(*common.FieldValue)(nil),     // 10: io.clbs.openhes.models.common.FieldValue
	(*ConnectionInfo)(nil),        // 11: io.clbs.openhes.models.acquisition.ConnectionInfo
	(*StartJobData)(nil),          // 12: io.clbs.openhes.models.acquisition.StartJobData
	(*JobProgressUpdate)(nil),     // 13: io.clbs.openhes.models.acquisition.JobProgressUpdate
	(*ActionProgressUpdate)(nil),  // 14: io.clbs.openhes.models.acquisition.ActionProgressUpdate
	(*JobSpec)(nil),               // 15: io.clbs.openhes.models.acquisition.JobSpec
	(*JobStatus)(nil),             // 16: io.clbs.openhes.models.acquisition.JobStatus
}
var file_acquisition_internal_proto_depIdxs = []int32{
	10, // 0: io.clbs.openhes.models.acquisition.SetCacheRequest.value:type_name -> io.clbs.openhes.models.common.FieldValue
	10, // 1: io.clbs.openhes.models.acquisition.GetCacheResponse.value:type_name -> io.clbs.openhes.models.common.FieldValue
	11, // 2: io.clbs.openhes.models.acquisition.StartJobsRequest.connection_info:type_name -> io.clbs.openhes.models.acquisition.ConnectionInfo
	12, // 3: io.clbs.openhes.models.acquisition.StartJobsRequest.jobs:type_name -> io.clbs.openhes.models.acquisition.StartJobData
	13, // 4: io.clbs.openhes.models.acquisition.ProgressUpdate.job:type_name -> io.clbs.openhes.models.acquisition.JobProgressUpdate
	14, // 5: io.clbs.openhes.models.acquisition.ProgressUpdate.action:type_name -> io.clbs.openhes.models.acquisition.ActionProgressUpdate
	15, // 6: io.clbs.openhes.models.acquisition.QueueJobsRequest.spec:type_name -> io.clbs.openhes.models.acquisition.JobSpec
	16, // 7: io.clbs.openhes.models.acquisition.GetJobResponse.status:type_name -> io.clbs.openhes.models.acquisition.JobStatus
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
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
			RawDescriptor: file_acquisition_internal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_acquisition_internal_proto_goTypes,
		DependencyIndexes: file_acquisition_internal_proto_depIdxs,
		MessageInfos:      file_acquisition_internal_proto_msgTypes,
	}.Build()
	File_acquisition_internal_proto = out.File
	file_acquisition_internal_proto_rawDesc = nil
	file_acquisition_internal_proto_goTypes = nil
	file_acquisition_internal_proto_depIdxs = nil
}
