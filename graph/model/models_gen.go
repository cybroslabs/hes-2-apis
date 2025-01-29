// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/google/uuid"
)

type AccessLevelTemplate struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ActionData struct {
	Nodata           *Empty         `json:"nodata,omitempty"`
	Billings         *BillingValues `json:"billings,omitempty"`
	Profile          *ProfileValues `json:"profile,omitempty"`
	IrregularProfile *BillingValues `json:"irregularProfile,omitempty"`
}

type ActionFwUpdate struct {
	Empty *bool `json:"_empty,omitempty"`
}

type ActionGetDeviceInfo struct {
	InfoTimestamp            *string  `json:"infoTimestamp,omitempty"`
	ManufacturerSerialNumber *string  `json:"manufacturerSerialNumber,omitempty"`
	DeviceSerialNumber       *string  `json:"deviceSerialNumber,omitempty"`
	FirmwareVersion          *string  `json:"firmwareVersion,omitempty"`
	ClockDelta               *float64 `json:"clockDelta,omitempty"`
	DeviceModel              *string  `json:"deviceModel,omitempty"`
	ErrorRegister            *int32   `json:"errorRegister,omitempty"`
}

type ActionGetDisconnectorState struct {
	Empty *bool `json:"_empty,omitempty"`
}

type ActionGetEvents struct {
	From *string `json:"from,omitempty"`
	To   *string `json:"to,omitempty"`
}

type ActionGetIrregularProfile struct {
	From *string `json:"from,omitempty"`
	To   *string `json:"to,omitempty"`
}

type ActionGetLimiter struct {
	Empty *bool `json:"_empty,omitempty"`
}

type ActionGetPeriodicalProfile struct {
	From *string `json:"from,omitempty"`
	To   *string `json:"to,omitempty"`
}

type ActionGetRegister struct {
	Empty *bool `json:"_empty,omitempty"`
}

type ActionGetRelayState struct {
	Empty *bool `json:"_empty,omitempty"`
}

type ActionGetTou struct {
	Empty *bool `json:"_empty,omitempty"`
}

type ActionResetBillingPeriod struct {
	Empty *bool `json:"_empty,omitempty"`
}

type ActionResult struct {
	ActionID *string           `json:"actionId,omitempty"`
	Status   *ActionResultCode `json:"status,omitempty"`
	Data     *ActionData       `json:"data,omitempty"`
}

type ActionSetDisconnectorState struct {
	Empty *bool `json:"_empty,omitempty"`
}

type ActionSetLimiter struct {
	Empty *bool `json:"_empty,omitempty"`
}

type ActionSetRelayState struct {
	Empty *bool `json:"_empty,omitempty"`
}

type ActionSetTou struct {
	Empty *bool `json:"_empty,omitempty"`
}

type ActionSyncClock struct {
	Empty *bool `json:"_empty,omitempty"`
}

type AddDevicesToGroupRequest struct {
	GroupID  *string   `json:"groupId,omitempty"`
	DeviceID []*string `json:"deviceId,omitempty"`
}

type ApplicationProtocolTemplate struct {
	ID         *string              `json:"id,omitempty"`
	Protocol   *ApplicationProtocol `json:"protocol,omitempty"`
	Attributes []*FieldDescriptor   `json:"attributes,omitempty"`
}

type BillingValue struct {
	Timestamp *string        `json:"timestamp,omitempty"`
	Unit      *string        `json:"unit,omitempty"`
	Value     *MeasuredValue `json:"value,omitempty"`
}

type BillingValues struct {
	Values []*BillingValue `json:"values,omitempty"`
}

type Bulk struct {
	Spec     *BulkSpec       `json:"spec,omitempty"`
	Status   *BulkStatus     `json:"status,omitempty"`
	Metadata *MetadataFields `json:"metadata,omitempty"`
}

type BulkJob struct {
	JobID  *string    `json:"jobId,omitempty"`
	Status *JobStatus `json:"status,omitempty"`
}

type BulkSpec struct {
	CorrelationID *uuid.UUID         `json:"correlationId,omitempty"`
	DriverType    *string            `json:"driverType,omitempty"`
	Devices       *ListOfJobDeviceID `json:"devices,omitempty"`
	CustomDevices *ListOfJobDevice   `json:"customDevices,omitempty"`
	Settings      *JobSettings       `json:"settings,omitempty"`
	Actions       []*JobAction       `json:"actions,omitempty"`
	WebhookURL    *string            `json:"webhookUrl,omitempty"`
	DeviceGroupID *string            `json:"deviceGroupId,omitempty"`
}

type BulkStatus struct {
	Status *BulkStatusCode `json:"status,omitempty"`
	Jobs   []*BulkJob      `json:"jobs,omitempty"`
}

type CommunicationTemplate struct {
	Type      *CommunicationType  `json:"type,omitempty"`
	Datalinks []*DataLinkTemplate `json:"datalinks,omitempty"`
}

type CommunicationUnit struct {
	Spec     *CommunicationUnitSpec `json:"spec,omitempty"`
	Metadata *MetadataFields        `json:"metadata,omitempty"`
}

type CommunicationUnitSpec struct {
	ExternalID     *string         `json:"externalId,omitempty"`
	ConnectionInfo *ConnectionInfo `json:"connectionInfo,omitempty"`
}

type ConnectionInfo struct {
	Tcpip            *ConnectionTypeDirectTCPIP      `json:"tcpip,omitempty"`
	ModemPool        *ConnectionTypeModemPool        `json:"modemPool,omitempty"`
	SerialOverIP     *ConnectionTypeControlledSerial `json:"serialOverIp,omitempty"`
	LinkProtocol     *DataLinkProtocol               `json:"linkProtocol,omitempty"`
	CustomGroupingID *string                         `json:"customGroupingId,omitempty"`
}

type ConnectionTypeControlledSerial struct {
	Direct *ConnectionTypeSerialDirect `json:"direct,omitempty"`
	Moxa   *ConnectionTypeSerialMoxa   `json:"moxa,omitempty"`
}

type ConnectionTypeDirectTCPIP struct {
	Host    *string `json:"host,omitempty"`
	Port    *int32  `json:"port,omitempty"`
	Timeout *int32  `json:"timeout,omitempty"`
}

type ConnectionTypeModemPool struct {
	Number *string    `json:"number,omitempty"`
	PoolID *string    `json:"poolId,omitempty"`
	Modem  *ModemInfo `json:"modem,omitempty"`
}

type ConnectionTypeSerialDirect struct {
	Host *string `json:"host,omitempty"`
	Port *int32  `json:"port,omitempty"`
}

type ConnectionTypeSerialMoxa struct {
	Host        *string `json:"host,omitempty"`
	DataPort    *int32  `json:"dataPort,omitempty"`
	CommandPort *int32  `json:"commandPort,omitempty"`
}

type CreateBulkRequest struct {
	Spec     *BulkSpec       `json:"spec,omitempty"`
	Metadata *MetadataFields `json:"metadata,omitempty"`
}

type CreateCommunicationUnitRequest struct {
	Spec     *CommunicationUnitSpec `json:"spec,omitempty"`
	Metadata *MetadataFields        `json:"metadata,omitempty"`
}

type CreateDeviceGroupRequest struct {
	Spec     *DeviceGroupSpec `json:"spec,omitempty"`
	Metadata *MetadataFields  `json:"metadata,omitempty"`
}

type CreateDeviceRequest struct {
	Spec     *DeviceSpec     `json:"spec,omitempty"`
	Metadata *MetadataFields `json:"metadata,omitempty"`
}

type DataLinkTemplate struct {
	LinkProtocol    *DataLinkProtocol      `json:"linkProtocol,omitempty"`
	AppProtocolRefs []*ApplicationProtocol `json:"appProtocolRefs,omitempty"`
	Attributes      []*FieldDescriptor     `json:"attributes,omitempty"`
}

type Device struct {
	Spec     *DeviceSpec     `json:"spec,omitempty"`
	Metadata *MetadataFields `json:"metadata,omitempty"`
}

type DeviceCommunicationUnit struct {
	CommunicationUnitID *string              `json:"communicationUnitId,omitempty"`
	AppProtocol         *ApplicationProtocol `json:"appProtocol,omitempty"`
}

type DeviceGroup struct {
	Spec     *DeviceGroupSpec   `json:"spec,omitempty"`
	Status   *DeviceGroupStatus `json:"status,omitempty"`
	Metadata *MetadataFields    `json:"metadata,omitempty"`
}

type DeviceGroupSpec struct {
	ExternalID    *string       `json:"externalId,omitempty"`
	DynamicFilter *ListSelector `json:"dynamicFilter,omitempty"`
}

type DeviceGroupStatus struct {
	DeviceID []*string `json:"deviceId,omitempty"`
}

type DeviceSpec struct {
	ExternalID            *string                    `json:"externalId,omitempty"`
	Attributes            []*MapFieldValue           `json:"attributes,omitempty"`
	CommunicationUnitLink []*DeviceCommunicationUnit `json:"communicationUnitLink,omitempty"`
	Timezone              *string                    `json:"timezone,omitempty"`
}

type Driver struct {
	Spec *DriverSpec `json:"spec,omitempty"`
}

type DriverSpec struct {
	Version           *string          `json:"version,omitempty"`
	ListeningPort     *int32           `json:"listeningPort,omitempty"`
	DriverType        *string          `json:"driverType,omitempty"`
	MaxConcurrentJobs *int32           `json:"maxConcurrentJobs,omitempty"`
	MaxCascadeDepth   *int32           `json:"maxCascadeDepth,omitempty"`
	TypicalMemUsage   *int32           `json:"typicalMemUsage,omitempty"`
	Templates         *DriverTemplates `json:"templates,omitempty"`
	DisplayName       *string          `json:"displayName,omitempty"`
}

type DriverTemplates struct {
	CommunicationTemplates []*CommunicationTemplate       `json:"communicationTemplates,omitempty"`
	AppProtocols           []*ApplicationProtocolTemplate `json:"appProtocols,omitempty"`
	ActionAttributes       []*JobActionAttributes         `json:"actionAttributes,omitempty"`
	AccessTemplates        []*AccessLevelTemplate         `json:"accessTemplates,omitempty"`
	ActionConstraints      *JobActionContraints           `json:"actionConstraints,omitempty"`
}

type Empty struct {
	Empty *bool `json:"_empty,omitempty"`
}

type FieldDescriptor struct {
	FieldID      *string             `json:"fieldId,omitempty"`
	Label        *string             `json:"label,omitempty"`
	DataType     *FieldDataType      `json:"dataType,omitempty"`
	Format       *FieldDisplayFormat `json:"format,omitempty"`
	Unit         *string             `json:"unit,omitempty"`
	GroupID      *string             `json:"groupId,omitempty"`
	Precision    *int32              `json:"precision,omitempty"`
	Tooltip      *string             `json:"tooltip,omitempty"`
	Required     *bool               `json:"required,omitempty"`
	Editable     *bool               `json:"editable,omitempty"`
	Visible      *bool               `json:"visible,omitempty"`
	MultiValue   *bool               `json:"multiValue,omitempty"`
	Secured      *bool               `json:"secured,omitempty"`
	Validation   *FieldValidation    `json:"validation,omitempty"`
	DefaultValue *FieldValue         `json:"defaultValue,omitempty"`
}

type FieldValidation struct {
	Re         *string      `json:"re,omitempty"`
	MinLength  *int32       `json:"minLength,omitempty"`
	MaxLength  *int32       `json:"maxLength,omitempty"`
	MinInteger *int64       `json:"minInteger,omitempty"`
	MaxInteger *int64       `json:"maxInteger,omitempty"`
	MinNumber  *float64     `json:"minNumber,omitempty"`
	MaxNumber  *float64     `json:"maxNumber,omitempty"`
	Options    []*Mapstring `json:"options,omitempty"`
}

type FieldValue struct {
	StringValue  *string  `json:"stringValue,omitempty"`
	IntegerValue *int64   `json:"integerValue,omitempty"`
	DoubleValue  *float64 `json:"doubleValue,omitempty"`
	BinaryValue  *string  `json:"binaryValue,omitempty"`
	BoolValue    *bool    `json:"boolValue,omitempty"`
	DateValue    *string  `json:"dateValue,omitempty"`
}

type JobAction struct {
	ActionID             *string                     `json:"actionId,omitempty"`
	Attributes           []*MapFieldValue            `json:"attributes,omitempty"`
	GetRegister          *ActionGetRegister          `json:"getRegister,omitempty"`
	GetPeriodicalProfile *ActionGetPeriodicalProfile `json:"getPeriodicalProfile,omitempty"`
	GetIrregularProfile  *ActionGetIrregularProfile  `json:"getIrregularProfile,omitempty"`
	GetEvents            *ActionGetEvents            `json:"getEvents,omitempty"`
	GetDeviceInfo        *ActionGetDeviceInfo        `json:"getDeviceInfo,omitempty"`
	SyncClock            *ActionSyncClock            `json:"syncClock,omitempty"`
	GetRelayState        *ActionGetRelayState        `json:"getRelayState,omitempty"`
	SetRelayState        *ActionSetRelayState        `json:"setRelayState,omitempty"`
	GetDisconnectorState *ActionGetDisconnectorState `json:"getDisconnectorState,omitempty"`
	SetDisconnectorState *ActionSetDisconnectorState `json:"setDisconnectorState,omitempty"`
	GetTou               *ActionGetTou               `json:"getTou,omitempty"`
	SetTou               *ActionSetTou               `json:"setTou,omitempty"`
	GetLimiter           *ActionGetLimiter           `json:"getLimiter,omitempty"`
	SetLimiter           *ActionSetLimiter           `json:"setLimiter,omitempty"`
	ResetBillingPeriod   *ActionResetBillingPeriod   `json:"resetBillingPeriod,omitempty"`
	FwUpdate             *ActionFwUpdate             `json:"fwUpdate,omitempty"`
}

type JobActionAttributes struct {
	Type       *ActionType        `json:"type,omitempty"`
	Attributes []*FieldDescriptor `json:"attributes,omitempty"`
}

type JobActionContraints struct {
	GetRegisterTypeName       []*Mapstring       `json:"getRegisterTypeName,omitempty"`
	GetRegisterTypeAttributes []*MapListOfString `json:"getRegisterTypeAttributes,omitempty"`
}

type JobDevice struct {
	JobID            *string              `json:"jobId,omitempty"`
	DeviceID         *string              `json:"deviceId,omitempty"`
	ExternalID       *string              `json:"externalId,omitempty"`
	DeviceAttributes []*MapFieldValue     `json:"deviceAttributes,omitempty"`
	ConnectionInfo   []*ConnectionInfo    `json:"connectionInfo,omitempty"`
	AppProtocol      *ApplicationProtocol `json:"appProtocol,omitempty"`
	Timezone         *string              `json:"timezone,omitempty"`
}

type JobDeviceID struct {
	JobID    *string `json:"jobId,omitempty"`
	DeviceID *string `json:"deviceId,omitempty"`
}

type JobSettings struct {
	MaxDuration *int64       `json:"maxDuration,omitempty"`
	Priority    *JobPriority `json:"priority,omitempty"`
	Attempts    []*int32     `json:"attempts,omitempty"`
	RetryDelay  *int64       `json:"retryDelay,omitempty"`
	DeferStart  *string      `json:"deferStart,omitempty"`
	ExpiresAt   *string      `json:"expiresAt,omitempty"`
}

type JobStatus struct {
	Status       *JobStatusCode  `json:"status,omitempty"`
	Code         *JobErrorCode   `json:"code,omitempty"`
	Results      []*ActionResult `json:"results,omitempty"`
	CreatedAt    *string         `json:"createdAt,omitempty"`
	StartedAt    *string         `json:"startedAt,omitempty"`
	FinishedAt   *string         `json:"finishedAt,omitempty"`
	AttemptsDone *int32          `json:"attemptsDone,omitempty"`
}

type ListOfBulk struct {
	Items []*Bulk `json:"items,omitempty"`
}

type ListOfCommunicationUnit struct {
	Items []*CommunicationUnit `json:"items,omitempty"`
}

type ListOfDevice struct {
	Items []*Device `json:"items,omitempty"`
}

type ListOfDeviceCommunicationUnit struct {
	CommunicationUnits []*DeviceCommunicationUnit `json:"communicationUnits,omitempty"`
}

type ListOfDeviceGroup struct {
	Items []*DeviceGroup `json:"items,omitempty"`
}

type ListOfDriver struct {
	Items []*Driver `json:"items,omitempty"`
}

type ListOfJobDevice struct {
	List []*JobDevice `json:"list,omitempty"`
}

type ListOfJobDeviceID struct {
	List []*JobDeviceID `json:"list,omitempty"`
}

type ListOfModemPool struct {
	Items []*ModemPool `json:"items,omitempty"`
}

type ListOfString struct {
	Items []*string `json:"items,omitempty"`
}

type ListSelector struct {
	PageSize *int32                  `json:"pageSize,omitempty"`
	Offset   *int32                  `json:"offset,omitempty"`
	SortBy   []*ListSelectorSortBy   `json:"sortBy,omitempty"`
	FilterBy []*ListSelectorFilterBy `json:"filterBy,omitempty"`
	Fields   []*string               `json:"fields,omitempty"`
}

type ListSelectorFilterBy struct {
	FieldID  *string         `json:"fieldId,omitempty"`
	Operator *FilterOperator `json:"operator,omitempty"`
	Text     []*string       `json:"text,omitempty"`
	Integer  []*int64        `json:"integer,omitempty"`
	Number   []*float64      `json:"number,omitempty"`
	Boolean  []*bool         `json:"boolean,omitempty"`
	Date     []*string       `json:"date,omitempty"`
}

type ListSelectorSortBy struct {
	FieldID *string `json:"fieldId,omitempty"`
	Desc    *bool   `json:"desc,omitempty"`
}

type MeasuredValue struct {
	Status           *int64   `json:"status,omitempty"`
	Exponent         *int32   `json:"exponent,omitempty"`
	DoubleValue      *float64 `json:"doubleValue,omitempty"`
	IntegerValue     *int64   `json:"integerValue,omitempty"`
	StringValue      *string  `json:"stringValue,omitempty"`
	TimestampValue   *string  `json:"timestampValue,omitempty"`
	TimestampTzValue *string  `json:"timestampTzValue,omitempty"`
	BoolValue        *bool    `json:"boolValue,omitempty"`
}

type MetadataFields struct {
	ID            *string          `json:"id,omitempty"`
	Generation    *int32           `json:"generation,omitempty"`
	Fields        []*MapFieldValue `json:"fields,omitempty"`
	ManagedFields []*MapFieldValue `json:"managedFields,omitempty"`
	Name          *string          `json:"name,omitempty"`
}

type ModemInfo struct {
	ModemID        *string                    `json:"modemId,omitempty"`
	Name           *string                    `json:"name,omitempty"`
	AtInit         *string                    `json:"atInit,omitempty"`
	AtTest         *string                    `json:"atTest,omitempty"`
	AtConfig       *string                    `json:"atConfig,omitempty"`
	AtDial         *string                    `json:"atDial,omitempty"`
	AtHangup       *string                    `json:"atHangup,omitempty"`
	AtEscape       *string                    `json:"atEscape,omitempty"`
	AtDsr          *bool                      `json:"atDsr,omitempty"`
	ConnectTimeout *int32                     `json:"connectTimeout,omitempty"`
	Tcpip          *ConnectionTypeDirectTCPIP `json:"tcpip,omitempty"`
}

type ModemPool struct {
	Spec     *ModemPoolSpec   `json:"spec,omitempty"`
	Status   *ModemPoolStatus `json:"status,omitempty"`
	Metadata *MetadataFields  `json:"metadata,omitempty"`
}

type ModemPoolSpec struct {
	Empty *bool `json:"_empty,omitempty"`
}

type ModemPoolStatus struct {
	Modems []*ModemInfo `json:"modems,omitempty"`
}

type ProfileBlok struct {
	StartTimestamp *string          `json:"startTimestamp,omitempty"`
	Values         []*MeasuredValue `json:"values,omitempty"`
}

type ProfileValues struct {
	Period *int32         `json:"period,omitempty"`
	Unit   *string        `json:"unit,omitempty"`
	Blocks []*ProfileBlok `json:"blocks,omitempty"`
}

type Query struct {
}

type RemoveDevicesFromGroupRequest struct {
	GroupID  *string   `json:"groupId,omitempty"`
	DeviceID []*string `json:"deviceId,omitempty"`
}

type SetDeviceCommunicationUnitsRequest struct {
	DeviceID           *string                    `json:"deviceId,omitempty"`
	CommunicationUnits []*DeviceCommunicationUnit `json:"communicationUnits,omitempty"`
}

type SetModemPoolRequest struct {
	Spec     *ModemPoolSpec  `json:"spec,omitempty"`
	Metadata *MetadataFields `json:"metadata,omitempty"`
}

type SetModemRequest struct {
	PoolID *string    `json:"poolId,omitempty"`
	Modem  *ModemInfo `json:"modem,omitempty"`
}

type StringValue struct {
	Value *string `json:"value,omitempty"`
}

type SystemConfig struct {
	MaxReplicas           *int32 `json:"maxReplicas,omitempty"`
	MaxCascadeDeviceCount *int32 `json:"maxCascadeDeviceCount,omitempty"`
	MaxSlotsPerDriver     *int32 `json:"maxSlotsPerDriver,omitempty"`
}

type MapFieldValue struct {
	Key   string      `json:"key"`
	Value *FieldValue `json:"value,omitempty"`
}

type MapListOfString struct {
	Key   string        `json:"key"`
	Value *ListOfString `json:"value,omitempty"`
}

type Mapstring struct {
	Key   string  `json:"key"`
	Value *string `json:"value,omitempty"`
}

type ActionResultCode string

const (
	ActionResultCodeErrorCodeActionOk          ActionResultCode = "ERROR_CODE_ACTION_OK"
	ActionResultCodeErrorCodeActionUnsupported ActionResultCode = "ERROR_CODE_ACTION_UNSUPPORTED"
	ActionResultCodeErrorCodeActionPending     ActionResultCode = "ERROR_CODE_ACTION_PENDING"
	ActionResultCodeErrorCodeActionError       ActionResultCode = "ERROR_CODE_ACTION_ERROR"
)

var AllActionResultCode = []ActionResultCode{
	ActionResultCodeErrorCodeActionOk,
	ActionResultCodeErrorCodeActionUnsupported,
	ActionResultCodeErrorCodeActionPending,
	ActionResultCodeErrorCodeActionError,
}

func (e ActionResultCode) IsValid() bool {
	switch e {
	case ActionResultCodeErrorCodeActionOk, ActionResultCodeErrorCodeActionUnsupported, ActionResultCodeErrorCodeActionPending, ActionResultCodeErrorCodeActionError:
		return true
	}
	return false
}

func (e ActionResultCode) String() string {
	return string(e)
}

func (e *ActionResultCode) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ActionResultCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ActionResultCode", str)
	}
	return nil
}

func (e ActionResultCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ActionType string

const (
	ActionTypeActionTypeGetRegister          ActionType = "ACTION_TYPE_GET_REGISTER"
	ActionTypeActionTypeGetPeriodicalProfile ActionType = "ACTION_TYPE_GET_PERIODICAL_PROFILE"
	ActionTypeActionTypeGetIrregularProfile  ActionType = "ACTION_TYPE_GET_IRREGULAR_PROFILE"
	ActionTypeActionTypeGetEvents            ActionType = "ACTION_TYPE_GET_EVENTS"
	ActionTypeActionTypeGetDeviceInfo        ActionType = "ACTION_TYPE_GET_DEVICE_INFO"
	ActionTypeActionTypeSyncClock            ActionType = "ACTION_TYPE_SYNC_CLOCK"
	ActionTypeActionTypeGetRelayState        ActionType = "ACTION_TYPE_GET_RELAY_STATE"
	ActionTypeActionTypeSetRelayState        ActionType = "ACTION_TYPE_SET_RELAY_STATE"
	ActionTypeActionTypeGetDisconnectorState ActionType = "ACTION_TYPE_GET_DISCONNECTOR_STATE"
	ActionTypeActionTypeSetDisconnectorState ActionType = "ACTION_TYPE_SET_DISCONNECTOR_STATE"
	ActionTypeActionTypeGetTou               ActionType = "ACTION_TYPE_GET_TOU"
	ActionTypeActionTypeSetTou               ActionType = "ACTION_TYPE_SET_TOU"
	ActionTypeActionTypeGetLimiter           ActionType = "ACTION_TYPE_GET_LIMITER"
	ActionTypeActionTypeSetLimiter           ActionType = "ACTION_TYPE_SET_LIMITER"
	ActionTypeActionTypeResetBillingPeriod   ActionType = "ACTION_TYPE_RESET_BILLING_PERIOD"
	ActionTypeActionTypeFwUpdate             ActionType = "ACTION_TYPE_FW_UPDATE"
)

var AllActionType = []ActionType{
	ActionTypeActionTypeGetRegister,
	ActionTypeActionTypeGetPeriodicalProfile,
	ActionTypeActionTypeGetIrregularProfile,
	ActionTypeActionTypeGetEvents,
	ActionTypeActionTypeGetDeviceInfo,
	ActionTypeActionTypeSyncClock,
	ActionTypeActionTypeGetRelayState,
	ActionTypeActionTypeSetRelayState,
	ActionTypeActionTypeGetDisconnectorState,
	ActionTypeActionTypeSetDisconnectorState,
	ActionTypeActionTypeGetTou,
	ActionTypeActionTypeSetTou,
	ActionTypeActionTypeGetLimiter,
	ActionTypeActionTypeSetLimiter,
	ActionTypeActionTypeResetBillingPeriod,
	ActionTypeActionTypeFwUpdate,
}

func (e ActionType) IsValid() bool {
	switch e {
	case ActionTypeActionTypeGetRegister, ActionTypeActionTypeGetPeriodicalProfile, ActionTypeActionTypeGetIrregularProfile, ActionTypeActionTypeGetEvents, ActionTypeActionTypeGetDeviceInfo, ActionTypeActionTypeSyncClock, ActionTypeActionTypeGetRelayState, ActionTypeActionTypeSetRelayState, ActionTypeActionTypeGetDisconnectorState, ActionTypeActionTypeSetDisconnectorState, ActionTypeActionTypeGetTou, ActionTypeActionTypeSetTou, ActionTypeActionTypeGetLimiter, ActionTypeActionTypeSetLimiter, ActionTypeActionTypeResetBillingPeriod, ActionTypeActionTypeFwUpdate:
		return true
	}
	return false
}

func (e ActionType) String() string {
	return string(e)
}

func (e *ActionType) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ActionType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ActionType", str)
	}
	return nil
}

func (e ActionType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ApplicationProtocol string

const (
	ApplicationProtocolAppprotoIec62056_21 ApplicationProtocol = "APPPROTO_IEC_62056_21"
	ApplicationProtocolAppprotoDlmsSn      ApplicationProtocol = "APPPROTO_DLMS_SN"
	ApplicationProtocolAppprotoDlmsLn      ApplicationProtocol = "APPPROTO_DLMS_LN"
	ApplicationProtocolAppprotoSctm        ApplicationProtocol = "APPPROTO_SCTM"
	ApplicationProtocolAppprotoLis200      ApplicationProtocol = "APPPROTO_LIS200"
	ApplicationProtocolAppprotoAnsiC12     ApplicationProtocol = "APPPROTO_ANSI_C12"
	ApplicationProtocolAppprotoMqtt        ApplicationProtocol = "APPPROTO_MQTT"
)

var AllApplicationProtocol = []ApplicationProtocol{
	ApplicationProtocolAppprotoIec62056_21,
	ApplicationProtocolAppprotoDlmsSn,
	ApplicationProtocolAppprotoDlmsLn,
	ApplicationProtocolAppprotoSctm,
	ApplicationProtocolAppprotoLis200,
	ApplicationProtocolAppprotoAnsiC12,
	ApplicationProtocolAppprotoMqtt,
}

func (e ApplicationProtocol) IsValid() bool {
	switch e {
	case ApplicationProtocolAppprotoIec62056_21, ApplicationProtocolAppprotoDlmsSn, ApplicationProtocolAppprotoDlmsLn, ApplicationProtocolAppprotoSctm, ApplicationProtocolAppprotoLis200, ApplicationProtocolAppprotoAnsiC12, ApplicationProtocolAppprotoMqtt:
		return true
	}
	return false
}

func (e ApplicationProtocol) String() string {
	return string(e)
}

func (e *ApplicationProtocol) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ApplicationProtocol(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ApplicationProtocol", str)
	}
	return nil
}

func (e ApplicationProtocol) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type BulkStatusCode string

const (
	BulkStatusCodeBulkStatusQueued    BulkStatusCode = "BULK_STATUS_QUEUED"
	BulkStatusCodeBulkStatusRunning   BulkStatusCode = "BULK_STATUS_RUNNING"
	BulkStatusCodeBulkStatusCompleted BulkStatusCode = "BULK_STATUS_COMPLETED"
	BulkStatusCodeBulkStatusCancelled BulkStatusCode = "BULK_STATUS_CANCELLED"
	BulkStatusCodeBulkStatusExpired   BulkStatusCode = "BULK_STATUS_EXPIRED"
)

var AllBulkStatusCode = []BulkStatusCode{
	BulkStatusCodeBulkStatusQueued,
	BulkStatusCodeBulkStatusRunning,
	BulkStatusCodeBulkStatusCompleted,
	BulkStatusCodeBulkStatusCancelled,
	BulkStatusCodeBulkStatusExpired,
}

func (e BulkStatusCode) IsValid() bool {
	switch e {
	case BulkStatusCodeBulkStatusQueued, BulkStatusCodeBulkStatusRunning, BulkStatusCodeBulkStatusCompleted, BulkStatusCodeBulkStatusCancelled, BulkStatusCodeBulkStatusExpired:
		return true
	}
	return false
}

func (e BulkStatusCode) String() string {
	return string(e)
}

func (e *BulkStatusCode) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BulkStatusCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BulkStatusCode", str)
	}
	return nil
}

func (e BulkStatusCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CommunicationType string

const (
	CommunicationTypeCommunicationTypeTCPIP            CommunicationType = "COMMUNICATION_TYPE_TCPIP"
	CommunicationTypeCommunicationTypeModemPool        CommunicationType = "COMMUNICATION_TYPE_MODEM_POOL"
	CommunicationTypeCommunicationTypeSerialLineDirect CommunicationType = "COMMUNICATION_TYPE_SERIAL_LINE_DIRECT"
	CommunicationTypeCommunicationTypeSerialLineMoxa   CommunicationType = "COMMUNICATION_TYPE_SERIAL_LINE_MOXA"
	CommunicationTypeCommunicationTypeListening        CommunicationType = "COMMUNICATION_TYPE_LISTENING"
)

var AllCommunicationType = []CommunicationType{
	CommunicationTypeCommunicationTypeTCPIP,
	CommunicationTypeCommunicationTypeModemPool,
	CommunicationTypeCommunicationTypeSerialLineDirect,
	CommunicationTypeCommunicationTypeSerialLineMoxa,
	CommunicationTypeCommunicationTypeListening,
}

func (e CommunicationType) IsValid() bool {
	switch e {
	case CommunicationTypeCommunicationTypeTCPIP, CommunicationTypeCommunicationTypeModemPool, CommunicationTypeCommunicationTypeSerialLineDirect, CommunicationTypeCommunicationTypeSerialLineMoxa, CommunicationTypeCommunicationTypeListening:
		return true
	}
	return false
}

func (e CommunicationType) String() string {
	return string(e)
}

func (e *CommunicationType) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CommunicationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CommunicationType", str)
	}
	return nil
}

func (e CommunicationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DataLinkProtocol string

const (
	DataLinkProtocolLinkprotoIec62056_21   DataLinkProtocol = "LINKPROTO_IEC_62056_21"
	DataLinkProtocolLinkprotoHdlc          DataLinkProtocol = "LINKPROTO_HDLC"
	DataLinkProtocolLinkprotoCosemWrapper  DataLinkProtocol = "LINKPROTO_COSEM_WRAPPER"
	DataLinkProtocolLinkprotoModbus        DataLinkProtocol = "LINKPROTO_MODBUS"
	DataLinkProtocolLinkprotoMbus          DataLinkProtocol = "LINKPROTO_MBUS"
	DataLinkProtocolLinkprotoNotApplicable DataLinkProtocol = "LINKPROTO_NOT_APPLICABLE"
)

var AllDataLinkProtocol = []DataLinkProtocol{
	DataLinkProtocolLinkprotoIec62056_21,
	DataLinkProtocolLinkprotoHdlc,
	DataLinkProtocolLinkprotoCosemWrapper,
	DataLinkProtocolLinkprotoModbus,
	DataLinkProtocolLinkprotoMbus,
	DataLinkProtocolLinkprotoNotApplicable,
}

func (e DataLinkProtocol) IsValid() bool {
	switch e {
	case DataLinkProtocolLinkprotoIec62056_21, DataLinkProtocolLinkprotoHdlc, DataLinkProtocolLinkprotoCosemWrapper, DataLinkProtocolLinkprotoModbus, DataLinkProtocolLinkprotoMbus, DataLinkProtocolLinkprotoNotApplicable:
		return true
	}
	return false
}

func (e DataLinkProtocol) String() string {
	return string(e)
}

func (e *DataLinkProtocol) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DataLinkProtocol(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DataLinkProtocol", str)
	}
	return nil
}

func (e DataLinkProtocol) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FieldDataType string

const (
	FieldDataTypeText      FieldDataType = "TEXT"
	FieldDataTypeInteger   FieldDataType = "INTEGER"
	FieldDataTypeDouble    FieldDataType = "DOUBLE"
	FieldDataTypeBoolean   FieldDataType = "BOOLEAN"
	FieldDataTypeTimestamp FieldDataType = "TIMESTAMP"
	FieldDataTypeBinary    FieldDataType = "BINARY"
)

var AllFieldDataType = []FieldDataType{
	FieldDataTypeText,
	FieldDataTypeInteger,
	FieldDataTypeDouble,
	FieldDataTypeBoolean,
	FieldDataTypeTimestamp,
	FieldDataTypeBinary,
}

func (e FieldDataType) IsValid() bool {
	switch e {
	case FieldDataTypeText, FieldDataTypeInteger, FieldDataTypeDouble, FieldDataTypeBoolean, FieldDataTypeTimestamp, FieldDataTypeBinary:
		return true
	}
	return false
}

func (e FieldDataType) String() string {
	return string(e)
}

func (e *FieldDataType) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FieldDataType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FieldDataType", str)
	}
	return nil
}

func (e FieldDataType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FieldDisplayFormat string

const (
	FieldDisplayFormatDefault   FieldDisplayFormat = "DEFAULT"
	FieldDisplayFormatDuration  FieldDisplayFormat = "DURATION"
	FieldDisplayFormatInterval  FieldDisplayFormat = "INTERVAL"
	FieldDisplayFormatDate      FieldDisplayFormat = "DATE"
	FieldDisplayFormatUtcDate   FieldDisplayFormat = "UTC_DATE"
	FieldDisplayFormatMonth     FieldDisplayFormat = "MONTH"
	FieldDisplayFormatDayofweek FieldDisplayFormat = "DAYOFWEEK"
	FieldDisplayFormatTimeofday FieldDisplayFormat = "TIMEOFDAY"
	FieldDisplayFormatMoney     FieldDisplayFormat = "MONEY"
	FieldDisplayFormatPassword  FieldDisplayFormat = "PASSWORD"
	FieldDisplayFormatMultiline FieldDisplayFormat = "MULTILINE"
)

var AllFieldDisplayFormat = []FieldDisplayFormat{
	FieldDisplayFormatDefault,
	FieldDisplayFormatDuration,
	FieldDisplayFormatInterval,
	FieldDisplayFormatDate,
	FieldDisplayFormatUtcDate,
	FieldDisplayFormatMonth,
	FieldDisplayFormatDayofweek,
	FieldDisplayFormatTimeofday,
	FieldDisplayFormatMoney,
	FieldDisplayFormatPassword,
	FieldDisplayFormatMultiline,
}

func (e FieldDisplayFormat) IsValid() bool {
	switch e {
	case FieldDisplayFormatDefault, FieldDisplayFormatDuration, FieldDisplayFormatInterval, FieldDisplayFormatDate, FieldDisplayFormatUtcDate, FieldDisplayFormatMonth, FieldDisplayFormatDayofweek, FieldDisplayFormatTimeofday, FieldDisplayFormatMoney, FieldDisplayFormatPassword, FieldDisplayFormatMultiline:
		return true
	}
	return false
}

func (e FieldDisplayFormat) String() string {
	return string(e)
}

func (e *FieldDisplayFormat) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FieldDisplayFormat(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FieldDisplayFormat", str)
	}
	return nil
}

func (e FieldDisplayFormat) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FilterOperator string

const (
	FilterOperatorEqual              FilterOperator = "EQUAL"
	FilterOperatorNotEqual           FilterOperator = "NOT_EQUAL"
	FilterOperatorGreaterThan        FilterOperator = "GREATER_THAN"
	FilterOperatorGreaterThanOrEqual FilterOperator = "GREATER_THAN_OR_EQUAL"
	FilterOperatorLessThan           FilterOperator = "LESS_THAN"
	FilterOperatorLessThanOrEqual    FilterOperator = "LESS_THAN_OR_EQUAL"
	FilterOperatorContains           FilterOperator = "CONTAINS"
	FilterOperatorNotContains        FilterOperator = "NOT_CONTAINS"
	FilterOperatorStartsWith         FilterOperator = "STARTS_WITH"
	FilterOperatorEndsWith           FilterOperator = "ENDS_WITH"
	FilterOperatorIn                 FilterOperator = "IN"
	FilterOperatorNotIn              FilterOperator = "NOT_IN"
	FilterOperatorBetween            FilterOperator = "BETWEEN"
	FilterOperatorIsNull             FilterOperator = "IS_NULL"
	FilterOperatorIsNotNull          FilterOperator = "IS_NOT_NULL"
)

var AllFilterOperator = []FilterOperator{
	FilterOperatorEqual,
	FilterOperatorNotEqual,
	FilterOperatorGreaterThan,
	FilterOperatorGreaterThanOrEqual,
	FilterOperatorLessThan,
	FilterOperatorLessThanOrEqual,
	FilterOperatorContains,
	FilterOperatorNotContains,
	FilterOperatorStartsWith,
	FilterOperatorEndsWith,
	FilterOperatorIn,
	FilterOperatorNotIn,
	FilterOperatorBetween,
	FilterOperatorIsNull,
	FilterOperatorIsNotNull,
}

func (e FilterOperator) IsValid() bool {
	switch e {
	case FilterOperatorEqual, FilterOperatorNotEqual, FilterOperatorGreaterThan, FilterOperatorGreaterThanOrEqual, FilterOperatorLessThan, FilterOperatorLessThanOrEqual, FilterOperatorContains, FilterOperatorNotContains, FilterOperatorStartsWith, FilterOperatorEndsWith, FilterOperatorIn, FilterOperatorNotIn, FilterOperatorBetween, FilterOperatorIsNull, FilterOperatorIsNotNull:
		return true
	}
	return false
}

func (e FilterOperator) String() string {
	return string(e)
}

func (e *FilterOperator) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FilterOperator(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FilterOperator", str)
	}
	return nil
}

func (e FilterOperator) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type JobErrorCode string

const (
	JobErrorCodeJobErrorCodeNone          JobErrorCode = "JOB_ERROR_CODE_NONE"
	JobErrorCodeJobErrorCodeBusy          JobErrorCode = "JOB_ERROR_CODE_BUSY"
	JobErrorCodeJobErrorCodeError         JobErrorCode = "JOB_ERROR_CODE_ERROR"
	JobErrorCodeJobErrorCodeAlreadyExists JobErrorCode = "JOB_ERROR_CODE_ALREADY_EXISTS"
	JobErrorCodeJobErrorCodeFatal         JobErrorCode = "JOB_ERROR_CODE_FATAL"
)

var AllJobErrorCode = []JobErrorCode{
	JobErrorCodeJobErrorCodeNone,
	JobErrorCodeJobErrorCodeBusy,
	JobErrorCodeJobErrorCodeError,
	JobErrorCodeJobErrorCodeAlreadyExists,
	JobErrorCodeJobErrorCodeFatal,
}

func (e JobErrorCode) IsValid() bool {
	switch e {
	case JobErrorCodeJobErrorCodeNone, JobErrorCodeJobErrorCodeBusy, JobErrorCodeJobErrorCodeError, JobErrorCodeJobErrorCodeAlreadyExists, JobErrorCodeJobErrorCodeFatal:
		return true
	}
	return false
}

func (e JobErrorCode) String() string {
	return string(e)
}

func (e *JobErrorCode) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = JobErrorCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid JobErrorCode", str)
	}
	return nil
}

func (e JobErrorCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type JobPriority string

const (
	JobPriorityJobPriority0 JobPriority = "JOB_PRIORITY_0"
	JobPriorityJobPriority1 JobPriority = "JOB_PRIORITY_1"
	JobPriorityJobPriority2 JobPriority = "JOB_PRIORITY_2"
	JobPriorityJobPriority3 JobPriority = "JOB_PRIORITY_3"
	JobPriorityJobPriority4 JobPriority = "JOB_PRIORITY_4"
	JobPriorityJobPriority5 JobPriority = "JOB_PRIORITY_5"
	JobPriorityJobPriority6 JobPriority = "JOB_PRIORITY_6"
	JobPriorityJobPriority7 JobPriority = "JOB_PRIORITY_7"
	JobPriorityJobPriority8 JobPriority = "JOB_PRIORITY_8"
)

var AllJobPriority = []JobPriority{
	JobPriorityJobPriority0,
	JobPriorityJobPriority1,
	JobPriorityJobPriority2,
	JobPriorityJobPriority3,
	JobPriorityJobPriority4,
	JobPriorityJobPriority5,
	JobPriorityJobPriority6,
	JobPriorityJobPriority7,
	JobPriorityJobPriority8,
}

func (e JobPriority) IsValid() bool {
	switch e {
	case JobPriorityJobPriority0, JobPriorityJobPriority1, JobPriorityJobPriority2, JobPriorityJobPriority3, JobPriorityJobPriority4, JobPriorityJobPriority5, JobPriorityJobPriority6, JobPriorityJobPriority7, JobPriorityJobPriority8:
		return true
	}
	return false
}

func (e JobPriority) String() string {
	return string(e)
}

func (e *JobPriority) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = JobPriority(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid JobPriority", str)
	}
	return nil
}

func (e JobPriority) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type JobStatusCode string

const (
	JobStatusCodeJobStatusQueued    JobStatusCode = "JOB_STATUS_QUEUED"
	JobStatusCodeJobStatusRunning   JobStatusCode = "JOB_STATUS_RUNNING"
	JobStatusCodeJobStatusCompleted JobStatusCode = "JOB_STATUS_COMPLETED"
	JobStatusCodeJobStatusFailed    JobStatusCode = "JOB_STATUS_FAILED"
	JobStatusCodeJobStatusCancelled JobStatusCode = "JOB_STATUS_CANCELLED"
	JobStatusCodeJobStatusExpired   JobStatusCode = "JOB_STATUS_EXPIRED"
)

var AllJobStatusCode = []JobStatusCode{
	JobStatusCodeJobStatusQueued,
	JobStatusCodeJobStatusRunning,
	JobStatusCodeJobStatusCompleted,
	JobStatusCodeJobStatusFailed,
	JobStatusCodeJobStatusCancelled,
	JobStatusCodeJobStatusExpired,
}

func (e JobStatusCode) IsValid() bool {
	switch e {
	case JobStatusCodeJobStatusQueued, JobStatusCodeJobStatusRunning, JobStatusCodeJobStatusCompleted, JobStatusCodeJobStatusFailed, JobStatusCodeJobStatusCancelled, JobStatusCodeJobStatusExpired:
		return true
	}
	return false
}

func (e JobStatusCode) String() string {
	return string(e)
}

func (e *JobStatusCode) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = JobStatusCode(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid JobStatusCode", str)
	}
	return nil
}

func (e JobStatusCode) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
