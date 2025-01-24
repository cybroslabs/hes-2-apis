// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type AccessLevelTemplate struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type ActionData struct {
	Nodata   *Empty         `json:"nodata,omitempty"`
	Billings *BillingValues `json:"billings,omitempty"`
	Profile  *ProfileValues `json:"profile,omitempty"`
}

type ActionFwUpdate struct {
	Empty *bool `json:"_empty,omitempty"`
}

type ActionGetClock struct {
	Empty *bool `json:"_empty,omitempty"`
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
	ActionID *string           `json:"action_id,omitempty"`
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
	GroupID  *string   `json:"group_id,omitempty"`
	DeviceID []*string `json:"device_id,omitempty"`
}

type Any struct {
	TypeURL *string `json:"type_url,omitempty"`
	Value   *string `json:"value,omitempty"`
}

type ApplicationProtocolTemplate struct {
	ID         *string                `json:"id,omitempty"`
	Protocol   *ApplicationProtocol   `json:"protocol,omitempty"`
	Attributes []*AttributeDefinition `json:"attributes,omitempty"`
}

type AttributeDefinition struct {
	Name         *string         `json:"name,omitempty"`
	Description  *string         `json:"description,omitempty"`
	Type         *AttributeType  `json:"type,omitempty"`
	Mandatory    *bool           `json:"mandatory,omitempty"`
	DefaultValue *AttributeValue `json:"default_value,omitempty"`
	Options      []*Mapstring    `json:"options,omitempty"`
}

type AttributeValue struct {
	StrValue    *string  `json:"str_value,omitempty"`
	IntValue    *int64   `json:"int_value,omitempty"`
	DoubleValue *float64 `json:"double_value,omitempty"`
	BinaryValue *string  `json:"binary_value,omitempty"`
	BoolValue   *bool    `json:"bool_value,omitempty"`
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
	JobID  *string    `json:"job_id,omitempty"`
	Status *JobStatus `json:"status,omitempty"`
}

type BulkSpec struct {
	BulkID        *string      `json:"bulk_id,omitempty"`
	CorrelationID *string      `json:"correlation_id,omitempty"`
	DriverType    *string      `json:"driver_type,omitempty"`
	Devices       []*JobDevice `json:"devices,omitempty"`
	Settings      *JobSettings `json:"settings,omitempty"`
	JobActions    []*JobAction `json:"job_actions,omitempty"`
	WebhookURL    *string      `json:"webhook_url,omitempty"`
	UserID        *string      `json:"user_id,omitempty"`
}

type BulkStatus struct {
	Status *BulkStatusCode `json:"status,omitempty"`
	Jobs   []*BulkJob      `json:"jobs,omitempty"`
}

type CancelJobsRequest struct {
	JobID []*string `json:"job_id,omitempty"`
}

type CommunicationTemplate struct {
	Type      *CommunicationType  `json:"type,omitempty"`
	Datalinks []*DataLinkTemplate `json:"datalinks,omitempty"`
}

type CommunicationUnitSpec struct {
	ID             *string         `json:"id,omitempty"`
	ExternalID     *string         `json:"external_id,omitempty"`
	Name           *string         `json:"name,omitempty"`
	ConnectionInfo *ConnectionInfo `json:"connection_info,omitempty"`
}

type ConnectionInfo struct {
	Tcpip            *ConnectionTypeDirectTCPIP      `json:"tcpip,omitempty"`
	ModemPool        *ConnectionTypeModemPool        `json:"modem_pool,omitempty"`
	SerialOverIP     *ConnectionTypeControlledSerial `json:"serial_over_ip,omitempty"`
	LinkProtocol     *DataLinkProtocol               `json:"link_protocol,omitempty"`
	CustomGroupingID *string                         `json:"custom_grouping_id,omitempty"`
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
	PoolID *string    `json:"pool_id,omitempty"`
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
	Spec *CommunicationUnitSpec `json:"spec,omitempty"`
}

type CreateDeviceGroupRequest struct {
	Spec *DeviceGroupSpec `json:"spec,omitempty"`
}

type CreateDeviceRequest struct {
	Spec *DeviceSpec `json:"spec,omitempty"`
}

type DataLinkTemplate struct {
	LinkProtocol    *DataLinkProtocol      `json:"link_protocol,omitempty"`
	AppProtocolRefs []*ApplicationProtocol `json:"app_protocol_refs,omitempty"`
	Attributes      []*AttributeDefinition `json:"attributes,omitempty"`
}

type DeviceCommunicationUnit struct {
	CommunicationUnitID *string              `json:"communication_unit_id,omitempty"`
	AppProtocol         *ApplicationProtocol `json:"app_protocol,omitempty"`
}

type DeviceConnectionInfo struct {
	CommunicationUnit *ConnectionInfo      `json:"communication_unit,omitempty"`
	AppProtocol       *ApplicationProtocol `json:"app_protocol,omitempty"`
	DeviceAttributes  []*MapAttributeValue `json:"device_attributes,omitempty"`
}

type DeviceGroupOverviewSpec struct {
	ID         *string `json:"id,omitempty"`
	ExternalID *string `json:"external_id,omitempty"`
	Name       *string `json:"name,omitempty"`
}

type DeviceGroupSpec struct {
	ID         *string   `json:"id,omitempty"`
	ExternalID *string   `json:"external_id,omitempty"`
	Name       *string   `json:"name,omitempty"`
	DeviceID   []*string `json:"device_id,omitempty"`
}

type DeviceSpec struct {
	ID                    *string                    `json:"id,omitempty"`
	ExternalID            *string                    `json:"external_id,omitempty"`
	Name                  *string                    `json:"name,omitempty"`
	Attributes            []*MapAttributeValue       `json:"attributes,omitempty"`
	CommunicationUnitLink []*DeviceCommunicationUnit `json:"communication_unit_link,omitempty"`
	Timezone              *string                    `json:"timezone,omitempty"`
}

type DriverInfo struct {
	DriverType *string `json:"driver_type,omitempty"`
	Version    *string `json:"version,omitempty"`
}

type DriverTemplates struct {
	CommunicationTemplates []*CommunicationTemplate       `json:"communication_templates,omitempty"`
	AppProtocols           []*ApplicationProtocolTemplate `json:"app_protocols,omitempty"`
	ActionAttributes       []*JobActionAttributes         `json:"action_attributes,omitempty"`
	AccessTemplates        []*AccessLevelTemplate         `json:"access_templates,omitempty"`
	ActionConstraints      *JobActionContraints           `json:"action_constraints,omitempty"`
}

type Empty struct {
	Empty *bool `json:"_empty,omitempty"`
}

type GetDeviceGroupResponse struct {
	Spec *DeviceGroupSpec `json:"spec,omitempty"`
}

type GetDeviceGroupsResponse struct {
	Groups []*MapDeviceGroupOverviewSpec `json:"groups,omitempty"`
}

type GetModemPoolResponse struct {
	Modems []*ModemInfo `json:"modems,omitempty"`
	Name   *string      `json:"name,omitempty"`
}

type GetModemPoolsResponse struct {
	Pools []*ModemPoolSpec `json:"pools,omitempty"`
}

type JobAction struct {
	ActionID             *string                     `json:"action_id,omitempty"`
	Attributes           []*MapAttributeValue        `json:"attributes,omitempty"`
	GetRegister          *ActionGetRegister          `json:"get_register,omitempty"`
	GetPeriodicalProfile *ActionGetPeriodicalProfile `json:"get_periodical_profile,omitempty"`
	GetIrregularProfile  *ActionGetIrregularProfile  `json:"get_irregular_profile,omitempty"`
	GetEvents            *ActionGetEvents            `json:"get_events,omitempty"`
	GetClock             *ActionGetClock             `json:"get_clock,omitempty"`
	SyncClock            *ActionSyncClock            `json:"sync_clock,omitempty"`
	GetRelayState        *ActionGetRelayState        `json:"get_relay_state,omitempty"`
	SetRelayState        *ActionSetRelayState        `json:"set_relay_state,omitempty"`
	GetDisconnectorState *ActionGetDisconnectorState `json:"get_disconnector_state,omitempty"`
	SetDisconnectorState *ActionSetDisconnectorState `json:"set_disconnector_state,omitempty"`
	GetTou               *ActionGetTou               `json:"get_tou,omitempty"`
	SetTou               *ActionSetTou               `json:"set_tou,omitempty"`
	GetLimiter           *ActionGetLimiter           `json:"get_limiter,omitempty"`
	SetLimiter           *ActionSetLimiter           `json:"set_limiter,omitempty"`
	ResetBillingPeriod   *ActionResetBillingPeriod   `json:"reset_billing_period,omitempty"`
	FwUpdate             *ActionFwUpdate             `json:"fw_update,omitempty"`
}

type JobActionAttributes struct {
	Type       *ActionType            `json:"type,omitempty"`
	Attributes []*AttributeDefinition `json:"attributes,omitempty"`
}

type JobActionContraints struct {
	GetRegisterTypeName       []*Mapstring     `json:"get_register_type_name,omitempty"`
	GetRegisterTypeAttributes []*MapStringList `json:"get_register_type_attributes,omitempty"`
}

type JobDevice struct {
	ID               *string              `json:"id,omitempty"`
	DeviceID         *string              `json:"device_id,omitempty"`
	ExternalID       *string              `json:"external_id,omitempty"`
	DeviceAttributes []*MapAttributeValue `json:"device_attributes,omitempty"`
	ConnectionInfo   []*ConnectionInfo    `json:"connection_info,omitempty"`
	AppProtocol      *ApplicationProtocol `json:"app_protocol,omitempty"`
	Timezone         *string              `json:"timezone,omitempty"`
}

type JobSettings struct {
	MaxDuration *int64       `json:"max_duration,omitempty"`
	Priority    *JobPriority `json:"priority,omitempty"`
	Attempts    []*int32     `json:"attempts,omitempty"`
	RetryDelay  *int64       `json:"retry_delay,omitempty"`
	DeferStart  *string      `json:"defer_start,omitempty"`
	ExpiresAt   *string      `json:"expires_at,omitempty"`
}

type JobStatus struct {
	Status       *JobStatusCode  `json:"status,omitempty"`
	Code         *JobErrorCode   `json:"code,omitempty"`
	Results      []*ActionResult `json:"results,omitempty"`
	CreatedAt    *string         `json:"created_at,omitempty"`
	StartedAt    *string         `json:"started_at,omitempty"`
	FinishedAt   *string         `json:"finished_at,omitempty"`
	AttemptsDone *int32          `json:"attempts_done,omitempty"`
}

type ListOfBulk struct {
	Items []*Bulk `json:"items,omitempty"`
}

type ListOfCommunicationUnitSpec struct {
	Items []*CommunicationUnitSpec `json:"items,omitempty"`
}

type ListOfConnectionInfo struct {
	Items []*DeviceConnectionInfo `json:"items,omitempty"`
}

type ListOfDriverInfo struct {
	Items []*DriverInfo `json:"items,omitempty"`
}

type ListSelector struct {
	PageSize *int32                  `json:"page_size,omitempty"`
	Offset   *int32                  `json:"offset,omitempty"`
	SortBy   []*ListSelectorSortBy   `json:"sort_by,omitempty"`
	FilterBy []*ListSelectorFilterBy `json:"filter_by,omitempty"`
	Fields   []*string               `json:"fields,omitempty"`
}

type ListSelectorFilterBy struct {
	FieldID  *string         `json:"field_id,omitempty"`
	Operator *FilterOperator `json:"operator,omitempty"`
	Text     []*string       `json:"text,omitempty"`
	Integer  []*int64        `json:"integer,omitempty"`
	Number   []*float64      `json:"number,omitempty"`
	Boolean  []*bool         `json:"boolean,omitempty"`
	Date     []*string       `json:"date,omitempty"`
}

type ListSelectorSortBy struct {
	FieldID *string `json:"field_id,omitempty"`
	Desc    *bool   `json:"desc,omitempty"`
}

type MeasuredValue struct {
	Status           *int64   `json:"status,omitempty"`
	Exponent         *int32   `json:"exponent,omitempty"`
	DoubleValue      *float64 `json:"double_value,omitempty"`
	IntValue         *int64   `json:"int_value,omitempty"`
	StrValue         *string  `json:"str_value,omitempty"`
	TimestampValue   *string  `json:"timestamp_value,omitempty"`
	TimestampTzValue *string  `json:"timestamp_tz_value,omitempty"`
	BoolValue        *bool    `json:"bool_value,omitempty"`
}

type MetadataFields struct {
	Fields        []*MapAny `json:"fields,omitempty"`
	ManagedFields []*MapAny `json:"managed_fields,omitempty"`
}

type ModemInfo struct {
	ID             *string                    `json:"id,omitempty"`
	Name           *string                    `json:"name,omitempty"`
	AtInit         *string                    `json:"at_init,omitempty"`
	AtTest         *string                    `json:"at_test,omitempty"`
	AtConfig       *string                    `json:"at_config,omitempty"`
	AtDial         *string                    `json:"at_dial,omitempty"`
	AtHangup       *string                    `json:"at_hangup,omitempty"`
	AtEscape       *string                    `json:"at_escape,omitempty"`
	AtDsr          *bool                      `json:"at_dsr,omitempty"`
	ConnectTimeout *int32                     `json:"connect_timeout,omitempty"`
	Tcpip          *ConnectionTypeDirectTCPIP `json:"tcpip,omitempty"`
}

type ModemPoolSpec struct {
	PoolID *string `json:"pool_id,omitempty"`
	Name   *string `json:"name,omitempty"`
}

type ProfileBlok struct {
	StartTimestamp *string          `json:"start_timestamp,omitempty"`
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
	GroupID  *string   `json:"group_id,omitempty"`
	DeviceID []*string `json:"device_id,omitempty"`
}

type SetDeviceCommunicationUnitsRequest struct {
	DeviceID           *string                    `json:"device_id,omitempty"`
	CommunicationUnits []*DeviceCommunicationUnit `json:"communication_units,omitempty"`
}

type SetModemPoolRequest struct {
	PoolID *string `json:"pool_id,omitempty"`
	Name   *string `json:"name,omitempty"`
}

type SetModemRequest struct {
	PoolID *string    `json:"pool_id,omitempty"`
	Modem  *ModemInfo `json:"modem,omitempty"`
}

type StringList struct {
	Items []*string `json:"items,omitempty"`
}

type StringValue struct {
	Value *string `json:"value,omitempty"`
}

type SystemConfig struct {
	MaxReplicas           *int32 `json:"max_replicas,omitempty"`
	MaxCascadeDeviceCount *int32 `json:"max_cascade_device_count,omitempty"`
	MaxSlotsPerDriver     *int32 `json:"max_slots_per_driver,omitempty"`
}

type MapAny struct {
	Key   string `json:"key"`
	Value *Any   `json:"value,omitempty"`
}

type MapAttributeValue struct {
	Key   string          `json:"key"`
	Value *AttributeValue `json:"value,omitempty"`
}

type MapDeviceGroupOverviewSpec struct {
	Key   string                   `json:"key"`
	Value *DeviceGroupOverviewSpec `json:"value,omitempty"`
}

type MapStringList struct {
	Key   string      `json:"key"`
	Value *StringList `json:"value,omitempty"`
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
	ActionTypeActionTypeGetClock             ActionType = "ACTION_TYPE_GET_CLOCK"
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
	ActionTypeActionTypeGetClock,
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
	case ActionTypeActionTypeGetRegister, ActionTypeActionTypeGetPeriodicalProfile, ActionTypeActionTypeGetIrregularProfile, ActionTypeActionTypeGetEvents, ActionTypeActionTypeGetClock, ActionTypeActionTypeSyncClock, ActionTypeActionTypeGetRelayState, ActionTypeActionTypeSetRelayState, ActionTypeActionTypeGetDisconnectorState, ActionTypeActionTypeSetDisconnectorState, ActionTypeActionTypeGetTou, ActionTypeActionTypeSetTou, ActionTypeActionTypeGetLimiter, ActionTypeActionTypeSetLimiter, ActionTypeActionTypeResetBillingPeriod, ActionTypeActionTypeFwUpdate:
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

type AttributeType string

const (
	AttributeTypeAttributeTypeInt         AttributeType = "ATTRIBUTE_TYPE_INT"
	AttributeTypeAttributeTypeString      AttributeType = "ATTRIBUTE_TYPE_STRING"
	AttributeTypeAttributeTypeDouble      AttributeType = "ATTRIBUTE_TYPE_DOUBLE"
	AttributeTypeAttributeTypeBinary      AttributeType = "ATTRIBUTE_TYPE_BINARY"
	AttributeTypeAttributeTypeTimestamp   AttributeType = "ATTRIBUTE_TYPE_TIMESTAMP"
	AttributeTypeAttributeTypeTimestampTz AttributeType = "ATTRIBUTE_TYPE_TIMESTAMP_TZ"
	AttributeTypeAttributeTypeBool        AttributeType = "ATTRIBUTE_TYPE_BOOL"
)

var AllAttributeType = []AttributeType{
	AttributeTypeAttributeTypeInt,
	AttributeTypeAttributeTypeString,
	AttributeTypeAttributeTypeDouble,
	AttributeTypeAttributeTypeBinary,
	AttributeTypeAttributeTypeTimestamp,
	AttributeTypeAttributeTypeTimestampTz,
	AttributeTypeAttributeTypeBool,
}

func (e AttributeType) IsValid() bool {
	switch e {
	case AttributeTypeAttributeTypeInt, AttributeTypeAttributeTypeString, AttributeTypeAttributeTypeDouble, AttributeTypeAttributeTypeBinary, AttributeTypeAttributeTypeTimestamp, AttributeTypeAttributeTypeTimestampTz, AttributeTypeAttributeTypeBool:
		return true
	}
	return false
}

func (e AttributeType) String() string {
	return string(e)
}

func (e *AttributeType) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AttributeType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AttributeType", str)
	}
	return nil
}

func (e AttributeType) MarshalGQL(w io.Writer) {
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
