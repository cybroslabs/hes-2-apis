// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file pbdeviceregistry-models.proto (package io.clbs.openhes.pbdeviceregistry, edition 2023)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_google_protobuf_empty, file_google_protobuf_wrappers } from "@bufbuild/protobuf/wkt";
import type { ApplicationProtocol, ArrayOfConnectionInfo, AttributeValue, ConnectionInfo, ModemInfo } from "./pbdriver-models_pb";
import { file_pbdriver_models } from "./pbdriver-models_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file pbdeviceregistry-models.proto.
 */
export const file_pbdeviceregistry_models: GenFile = /*@__PURE__*/
  fileDesc("Ch1wYmRldmljZXJlZ2lzdHJ5LW1vZGVscy5wcm90bxIgaW8uY2xicy5vcGVuaGVzLnBiZGV2aWNlcmVnaXN0cnkiZwoeQ3JlYXRlQ29tbXVuaWNhdGlvblVuaXRSZXF1ZXN0EkUKBHNwZWMYASABKAsyNy5pby5jbGJzLm9wZW5oZXMucGJkZXZpY2VyZWdpc3RyeS5Db21tdW5pY2F0aW9uVW5pdFNwZWMiTQocR2V0Q29tbXVuaWNhdGlvblVuaXRzUmVxdWVzdBIKCgJpZBgBIAEoCRITCgtleHRlcm5hbF9pZBgCIAEoCRIMCgRuYW1lGAMgASgJImYKHUdldENvbW11bmljYXRpb25Vbml0c1Jlc3BvbnNlEkUKBHNwZWMYASADKAsyNy5pby5jbGJzLm9wZW5oZXMucGJkZXZpY2VyZWdpc3RyeS5Db21tdW5pY2F0aW9uVW5pdFNwZWMiiQEKFUNvbW11bmljYXRpb25Vbml0U3BlYxIKCgJpZBgBIAEoCRITCgtleHRlcm5hbF9pZBgCIAEoCRIMCgRuYW1lGAMgASgJEkEKD2Nvbm5lY3Rpb25faW5mbxgEIAEoCzIoLmlvLmNsYnMub3Blbmhlcy5wYmRyaXZlci5Db25uZWN0aW9uSW5mbyJRChNDcmVhdGVEZXZpY2VSZXF1ZXN0EjoKBHNwZWMYASABKAsyLC5pby5jbGJzLm9wZW5oZXMucGJkZXZpY2VyZWdpc3RyeS5EZXZpY2VTcGVjIkIKEUdldERldmljZXNSZXF1ZXN0EgoKAmlkGAEgASgJEhMKC2V4dGVybmFsX2lkGAIgASgJEgwKBG5hbWUYAyABKAkiUAoSR2V0RGV2aWNlc1Jlc3BvbnNlEjoKBHNwZWMYASADKAsyLC5pby5jbGJzLm9wZW5oZXMucGJkZXZpY2VyZWdpc3RyeS5EZXZpY2VTcGVjItgCCgpEZXZpY2VTcGVjEgoKAmlkGAEgASgJEhMKC2V4dGVybmFsX2lkGAIgASgJEgwKBG5hbWUYAyABKAkSUAoKYXR0cmlidXRlcxgEIAMoCzI8LmlvLmNsYnMub3Blbmhlcy5wYmRldmljZXJlZ2lzdHJ5LkRldmljZVNwZWMuQXR0cmlidXRlc0VudHJ5EloKF2NvbW11bmljYXRpb25fdW5pdF9saW5rGAUgAygLMjkuaW8uY2xicy5vcGVuaGVzLnBiZGV2aWNlcmVnaXN0cnkuRGV2aWNlQ29tbXVuaWNhdGlvblVuaXQSEAoIdGltZXpvbmUYBiABKAkaWwoPQXR0cmlidXRlc0VudHJ5EgsKA2tleRgBIAEoCRI3CgV2YWx1ZRgCIAEoCzIoLmlvLmNsYnMub3Blbmhlcy5wYmRyaXZlci5BdHRyaWJ1dGVWYWx1ZToCOAEiUwoPRGV2aWNlR3JvdXBTcGVjEgoKAmlkGAEgASgJEhMKC2V4dGVybmFsX2lkGAIgASgJEgwKBG5hbWUYAyABKAkSEQoJZGV2aWNlX2lkGAQgAygJIlsKGENyZWF0ZURldmljZUdyb3VwUmVxdWVzdBI/CgRzcGVjGAEgASgLMjEuaW8uY2xicy5vcGVuaGVzLnBiZGV2aWNlcmVnaXN0cnkuRGV2aWNlR3JvdXBTcGVjItoBChdHZXREZXZpY2VHcm91cHNSZXNwb25zZRJVCgZncm91cHMYASADKAsyRS5pby5jbGJzLm9wZW5oZXMucGJkZXZpY2VyZWdpc3RyeS5HZXREZXZpY2VHcm91cHNSZXNwb25zZS5Hcm91cHNFbnRyeRpoCgtHcm91cHNFbnRyeRILCgNrZXkYASABKAkSSAoFdmFsdWUYAiABKAsyOS5pby5jbGJzLm9wZW5oZXMucGJkZXZpY2VyZWdpc3RyeS5EZXZpY2VHcm91cE92ZXJ2aWV3U3BlYzoCOAEiSAoXRGV2aWNlR3JvdXBPdmVydmlld1NwZWMSCgoCaWQYASABKAkSEwoLZXh0ZXJuYWxfaWQYAiABKAkSDAoEbmFtZRgDIAEoCSJZChZHZXREZXZpY2VHcm91cFJlc3BvbnNlEj8KBHNwZWMYASABKAsyMS5pby5jbGJzLm9wZW5oZXMucGJkZXZpY2VyZWdpc3RyeS5EZXZpY2VHcm91cFNwZWMiPwoYQWRkRGV2aWNlc1RvR3JvdXBSZXF1ZXN0EhAKCGdyb3VwX2lkGAEgASgJEhEKCWRldmljZV9pZBgCIAMoCSJECh1SZW1vdmVEZXZpY2VzRnJvbUdyb3VwUmVxdWVzdBIQCghncm91cF9pZBgBIAEoCRIRCglkZXZpY2VfaWQYAiADKAkijwEKIlNldERldmljZUNvbW11bmljYXRpb25Vbml0c1JlcXVlc3QSEQoJZGV2aWNlX2lkGAEgASgJElYKE2NvbW11bmljYXRpb25fdW5pdHMYAiADKAsyOS5pby5jbGJzLm9wZW5oZXMucGJkZXZpY2VyZWdpc3RyeS5EZXZpY2VDb21tdW5pY2F0aW9uVW5pdCJ9ChdEZXZpY2VDb21tdW5pY2F0aW9uVW5pdBIdChVjb21tdW5pY2F0aW9uX3VuaXRfaWQYASABKAkSQwoMYXBwX3Byb3RvY29sGAIgASgOMi0uaW8uY2xicy5vcGVuaGVzLnBiZHJpdmVyLkFwcGxpY2F0aW9uUHJvdG9jb2wiOAojR2V0RGV2aWNlc0NvbW11bmljYXRpb25Vbml0c1JlcXVlc3QSEQoJZGV2aWNlX2lkGAEgAygJIu0BCiRHZXREZXZpY2VzQ29tbXVuaWNhdGlvblVuaXRzUmVzcG9uc2USZAoHZGV2aWNlcxgBIAMoCzJTLmlvLmNsYnMub3Blbmhlcy5wYmRldmljZXJlZ2lzdHJ5LkdldERldmljZXNDb21tdW5pY2F0aW9uVW5pdHNSZXNwb25zZS5EZXZpY2VzRW50cnkaXwoMRGV2aWNlc0VudHJ5EgsKA2tleRgBIAEoCRI+CgV2YWx1ZRgCIAEoCzIvLmlvLmNsYnMub3Blbmhlcy5wYmRyaXZlci5BcnJheU9mQ29ubmVjdGlvbkluZm86AjgBIi4KDU1vZGVtUG9vbFNwZWMSDwoHcG9vbF9pZBgBIAEoCRIMCgRuYW1lGAIgASgJIlcKFUdldE1vZGVtUG9vbHNSZXNwb25zZRI+CgVwb29scxgBIAMoCzIvLmlvLmNsYnMub3Blbmhlcy5wYmRldmljZXJlZ2lzdHJ5Lk1vZGVtUG9vbFNwZWMiJgoTR2V0TW9kZW1Qb29sUmVxdWVzdBIPCgdwb29sX2lkGAEgASgJIlkKFEdldE1vZGVtUG9vbFJlc3BvbnNlEjMKBm1vZGVtcxgBIAMoCzIjLmlvLmNsYnMub3Blbmhlcy5wYmRyaXZlci5Nb2RlbUluZm8SDAoEbmFtZRgCIAEoCSI0ChNTZXRNb2RlbVBvb2xSZXF1ZXN0Eg8KB3Bvb2xfaWQYASABKAkSDAoEbmFtZRgCIAEoCSJWCg9TZXRNb2RlbVJlcXVlc3QSDwoHcG9vbF9pZBgBIAEoCRIyCgVtb2RlbRgCIAEoCzIjLmlvLmNsYnMub3Blbmhlcy5wYmRyaXZlci5Nb2RlbUluZm9CQlpAZ2l0aHViLmNvbS9jeWJyb3NsYWJzL2hlcy0yLWFwaXMvcHJvdG9idWYvcGJkZXZpY2VyZWdpc3RyeW1vZGVsc2IIZWRpdGlvbnNw6Ac", [file_google_protobuf_empty, file_google_protobuf_wrappers, file_pbdriver_models]);

/**
 * RestApi -> DriverRegistry - The communication unit specification.
 *
 * @generated from message io.clbs.openhes.pbdeviceregistry.CreateCommunicationUnitRequest
 */
export type CreateCommunicationUnitRequest = Message<"io.clbs.openhes.pbdeviceregistry.CreateCommunicationUnitRequest"> & {
  /**
   * The communication unit specification.
   *
   * @generated from field: io.clbs.openhes.pbdeviceregistry.CommunicationUnitSpec spec = 1;
   */
  spec?: CommunicationUnitSpec;
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.CreateCommunicationUnitRequest.
 * Use `create(CreateCommunicationUnitRequestSchema)` to create a new message.
 */
export const CreateCommunicationUnitRequestSchema: GenMessage<CreateCommunicationUnitRequest> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 0);

/**
 * RestApi -> DriverRegistry - the request message to get the information about the communication unit.
 *
 * @generated from message io.clbs.openhes.pbdeviceregistry.GetCommunicationUnitsRequest
 */
export type GetCommunicationUnitsRequest = Message<"io.clbs.openhes.pbdeviceregistry.GetCommunicationUnitsRequest"> & {
  /**
   * The UUID of the communication unit.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * The external identifier of the communication unit.
   *
   * @generated from field: string external_id = 2;
   */
  externalId: string;

  /**
   * The name of the communication unit.
   *
   * @generated from field: string name = 3;
   */
  name: string;
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.GetCommunicationUnitsRequest.
 * Use `create(GetCommunicationUnitsRequestSchema)` to create a new message.
 */
export const GetCommunicationUnitsRequestSchema: GenMessage<GetCommunicationUnitsRequest> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 1);

/**
 * DriverRegistry -> RestApi - the message holds the information about the communication unit.
 *
 * @generated from message io.clbs.openhes.pbdeviceregistry.GetCommunicationUnitsResponse
 */
export type GetCommunicationUnitsResponse = Message<"io.clbs.openhes.pbdeviceregistry.GetCommunicationUnitsResponse"> & {
  /**
   * The communication unit specification.
   *
   * @generated from field: repeated io.clbs.openhes.pbdeviceregistry.CommunicationUnitSpec spec = 1;
   */
  spec: CommunicationUnitSpec[];
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.GetCommunicationUnitsResponse.
 * Use `create(GetCommunicationUnitsResponseSchema)` to create a new message.
 */
export const GetCommunicationUnitsResponseSchema: GenMessage<GetCommunicationUnitsResponse> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 2);

/**
 * Sub-message - the communication unit specification.
 *
 * @generated from message io.clbs.openhes.pbdeviceregistry.CommunicationUnitSpec
 */
export type CommunicationUnitSpec = Message<"io.clbs.openhes.pbdeviceregistry.CommunicationUnitSpec"> & {
  /**
   * The UUID of the communication unit.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * The external identifier of the communication unit.
   *
   * @generated from field: string external_id = 2;
   */
  externalId: string;

  /**
   * The name of the communication unit.
   *
   * @generated from field: string name = 3;
   */
  name: string;

  /**
   * The connection info.
   *
   * @generated from field: io.clbs.openhes.pbdriver.ConnectionInfo connection_info = 4;
   */
  connectionInfo?: ConnectionInfo;
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.CommunicationUnitSpec.
 * Use `create(CommunicationUnitSpecSchema)` to create a new message.
 */
export const CommunicationUnitSpecSchema: GenMessage<CommunicationUnitSpec> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 3);

/**
 * RestApi -> DriverRegistry - the request message to create a new device.
 *
 * @generated from message io.clbs.openhes.pbdeviceregistry.CreateDeviceRequest
 */
export type CreateDeviceRequest = Message<"io.clbs.openhes.pbdeviceregistry.CreateDeviceRequest"> & {
  /**
   * The device specification.
   *
   * @generated from field: io.clbs.openhes.pbdeviceregistry.DeviceSpec spec = 1;
   */
  spec?: DeviceSpec;
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.CreateDeviceRequest.
 * Use `create(CreateDeviceRequestSchema)` to create a new message.
 */
export const CreateDeviceRequestSchema: GenMessage<CreateDeviceRequest> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 4);

/**
 * RestApi -> DriverRegistry - the request message to get the information about the device.
 *
 * @generated from message io.clbs.openhes.pbdeviceregistry.GetDevicesRequest
 */
export type GetDevicesRequest = Message<"io.clbs.openhes.pbdeviceregistry.GetDevicesRequest"> & {
  /**
   * The UUID of the device.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * The external identifier of the device.
   *
   * @generated from field: string external_id = 2;
   */
  externalId: string;

  /**
   * The name of the device.
   *
   * @generated from field: string name = 3;
   */
  name: string;
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.GetDevicesRequest.
 * Use `create(GetDevicesRequestSchema)` to create a new message.
 */
export const GetDevicesRequestSchema: GenMessage<GetDevicesRequest> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 5);

/**
 * DriverRegistry -> RestApi - the message holds the information about the device.
 *
 * @generated from message io.clbs.openhes.pbdeviceregistry.GetDevicesResponse
 */
export type GetDevicesResponse = Message<"io.clbs.openhes.pbdeviceregistry.GetDevicesResponse"> & {
  /**
   * The device specification.
   *
   * @generated from field: repeated io.clbs.openhes.pbdeviceregistry.DeviceSpec spec = 1;
   */
  spec: DeviceSpec[];
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.GetDevicesResponse.
 * Use `create(GetDevicesResponseSchema)` to create a new message.
 */
export const GetDevicesResponseSchema: GenMessage<GetDevicesResponse> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 6);

/**
 * Sub-message - the device specification.
 *
 * @generated from message io.clbs.openhes.pbdeviceregistry.DeviceSpec
 */
export type DeviceSpec = Message<"io.clbs.openhes.pbdeviceregistry.DeviceSpec"> & {
  /**
   * The UUID of the device.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * The external identifier of the device.
   *
   * @generated from field: string external_id = 2;
   */
  externalId: string;

  /**
   * The name of the device.
   *
   * @generated from field: string name = 3;
   */
  name: string;

  /**
   * The template of the action attributes. It is represented as a list of attribute definitions.
   *
   * @generated from field: map<string, io.clbs.openhes.pbdriver.AttributeValue> attributes = 4;
   */
  attributes: { [key: string]: AttributeValue };

  /**
   * The list of communication unit identifiers (and additional info) that set CUs usable to communicate with the device. It's an ordered set where the first element is the primary communication unit with the highest priority.
   *
   * @generated from field: repeated io.clbs.openhes.pbdeviceregistry.DeviceCommunicationUnit communication_unit_link = 5;
   */
  communicationUnitLink: DeviceCommunicationUnit[];

  /**
   * The timezone related to the device, e.g. "America/New_York", "Europe/Prague", "CET", "GMT", "Etc/GMT+2".
   *
   * @generated from field: string timezone = 6;
   */
  timezone: string;
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.DeviceSpec.
 * Use `create(DeviceSpecSchema)` to create a new message.
 */
export const DeviceSpecSchema: GenMessage<DeviceSpec> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 7);

/**
 * Sub-message that represents the device group.
 *
 * @generated from message io.clbs.openhes.pbdeviceregistry.DeviceGroupSpec
 */
export type DeviceGroupSpec = Message<"io.clbs.openhes.pbdeviceregistry.DeviceGroupSpec"> & {
  /**
   * The UUID of the device group.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * The external identifier of the communication unit.
   *
   * @generated from field: string external_id = 2;
   */
  externalId: string;

  /**
   * The name of the device group.
   *
   * @generated from field: string name = 3;
   */
  name: string;

  /**
   * The list of device identifiers that are part of the group.
   *
   * @generated from field: repeated string device_id = 4;
   */
  deviceId: string[];
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.DeviceGroupSpec.
 * Use `create(DeviceGroupSpecSchema)` to create a new message.
 */
export const DeviceGroupSpecSchema: GenMessage<DeviceGroupSpec> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 8);

/**
 * RestApi -> DriverRegistry - the request message to get the information about the device group.
 *
 * @generated from message io.clbs.openhes.pbdeviceregistry.CreateDeviceGroupRequest
 */
export type CreateDeviceGroupRequest = Message<"io.clbs.openhes.pbdeviceregistry.CreateDeviceGroupRequest"> & {
  /**
   * The device group specification.
   *
   * @generated from field: io.clbs.openhes.pbdeviceregistry.DeviceGroupSpec spec = 1;
   */
  spec?: DeviceGroupSpec;
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.CreateDeviceGroupRequest.
 * Use `create(CreateDeviceGroupRequestSchema)` to create a new message.
 */
export const CreateDeviceGroupRequestSchema: GenMessage<CreateDeviceGroupRequest> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 9);

/**
 * DriverRegistry -> RestApi - the message holds the information about the device group.
 *
 * @generated from message io.clbs.openhes.pbdeviceregistry.GetDeviceGroupsResponse
 */
export type GetDeviceGroupsResponse = Message<"io.clbs.openhes.pbdeviceregistry.GetDeviceGroupsResponse"> & {
  /**
   * The list of device groups.
   *
   * @generated from field: map<string, io.clbs.openhes.pbdeviceregistry.DeviceGroupOverviewSpec> groups = 1;
   */
  groups: { [key: string]: DeviceGroupOverviewSpec };
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.GetDeviceGroupsResponse.
 * Use `create(GetDeviceGroupsResponseSchema)` to create a new message.
 */
export const GetDeviceGroupsResponseSchema: GenMessage<GetDeviceGroupsResponse> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 10);

/**
 * @generated from message io.clbs.openhes.pbdeviceregistry.DeviceGroupOverviewSpec
 */
export type DeviceGroupOverviewSpec = Message<"io.clbs.openhes.pbdeviceregistry.DeviceGroupOverviewSpec"> & {
  /**
   * The UUID of the device group.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * The external identifier of the communication unit.
   *
   * @generated from field: string external_id = 2;
   */
  externalId: string;

  /**
   * The name of the device group.
   *
   * @generated from field: string name = 3;
   */
  name: string;
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.DeviceGroupOverviewSpec.
 * Use `create(DeviceGroupOverviewSpecSchema)` to create a new message.
 */
export const DeviceGroupOverviewSpecSchema: GenMessage<DeviceGroupOverviewSpec> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 11);

/**
 * DriverRegistry -> RestApi - the message holds the information about the device group.
 *
 * @generated from message io.clbs.openhes.pbdeviceregistry.GetDeviceGroupResponse
 */
export type GetDeviceGroupResponse = Message<"io.clbs.openhes.pbdeviceregistry.GetDeviceGroupResponse"> & {
  /**
   * The device group specification.
   *
   * @generated from field: io.clbs.openhes.pbdeviceregistry.DeviceGroupSpec spec = 1;
   */
  spec?: DeviceGroupSpec;
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.GetDeviceGroupResponse.
 * Use `create(GetDeviceGroupResponseSchema)` to create a new message.
 */
export const GetDeviceGroupResponseSchema: GenMessage<GetDeviceGroupResponse> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 12);

/**
 * RestApi -> DriverRegistry - the request message to add a new device to the device group.
 *
 * @generated from message io.clbs.openhes.pbdeviceregistry.AddDevicesToGroupRequest
 */
export type AddDevicesToGroupRequest = Message<"io.clbs.openhes.pbdeviceregistry.AddDevicesToGroupRequest"> & {
  /**
   * The unique identifier of the device group.
   *
   * @generated from field: string group_id = 1;
   */
  groupId: string;

  /**
   * The unique identifier of the device.
   *
   * @generated from field: repeated string device_id = 2;
   */
  deviceId: string[];
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.AddDevicesToGroupRequest.
 * Use `create(AddDevicesToGroupRequestSchema)` to create a new message.
 */
export const AddDevicesToGroupRequestSchema: GenMessage<AddDevicesToGroupRequest> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 13);

/**
 * RestApi -> DriverRegistry - the request message to remove a device from the device group.
 *
 * @generated from message io.clbs.openhes.pbdeviceregistry.RemoveDevicesFromGroupRequest
 */
export type RemoveDevicesFromGroupRequest = Message<"io.clbs.openhes.pbdeviceregistry.RemoveDevicesFromGroupRequest"> & {
  /**
   * The unique identifier of the device group.
   *
   * @generated from field: string group_id = 1;
   */
  groupId: string;

  /**
   * The unique identifier of the device.
   *
   * @generated from field: repeated string device_id = 2;
   */
  deviceId: string[];
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.RemoveDevicesFromGroupRequest.
 * Use `create(RemoveDevicesFromGroupRequestSchema)` to create a new message.
 */
export const RemoveDevicesFromGroupRequestSchema: GenMessage<RemoveDevicesFromGroupRequest> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 14);

/**
 * RestApi -> DriverRegistry - the request message to add a new device to the communication unit.
 *
 * @generated from message io.clbs.openhes.pbdeviceregistry.SetDeviceCommunicationUnitsRequest
 */
export type SetDeviceCommunicationUnitsRequest = Message<"io.clbs.openhes.pbdeviceregistry.SetDeviceCommunicationUnitsRequest"> & {
  /**
   * The unique identifier of the device.
   *
   * @generated from field: string device_id = 1;
   */
  deviceId: string;

  /**
   * The list of linked communication units.
   *
   * @generated from field: repeated io.clbs.openhes.pbdeviceregistry.DeviceCommunicationUnit communication_units = 2;
   */
  communicationUnits: DeviceCommunicationUnit[];
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.SetDeviceCommunicationUnitsRequest.
 * Use `create(SetDeviceCommunicationUnitsRequestSchema)` to create a new message.
 */
export const SetDeviceCommunicationUnitsRequestSchema: GenMessage<SetDeviceCommunicationUnitsRequest> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 15);

/**
 * @generated from message io.clbs.openhes.pbdeviceregistry.DeviceCommunicationUnit
 */
export type DeviceCommunicationUnit = Message<"io.clbs.openhes.pbdeviceregistry.DeviceCommunicationUnit"> & {
  /**
   * The unique identifier of the communication unit.
   *
   * @generated from field: string communication_unit_id = 1;
   */
  communicationUnitId: string;

  /**
   * The application protocol to be used for the communication over the communication unit.
   *
   * @generated from field: io.clbs.openhes.pbdriver.ApplicationProtocol app_protocol = 2;
   */
  appProtocol: ApplicationProtocol;
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.DeviceCommunicationUnit.
 * Use `create(DeviceCommunicationUnitSchema)` to create a new message.
 */
export const DeviceCommunicationUnitSchema: GenMessage<DeviceCommunicationUnit> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 16);

/**
 * RestApi -> DriverRegistry - the request message to get
 *
 * @generated from message io.clbs.openhes.pbdeviceregistry.GetDevicesCommunicationUnitsRequest
 */
export type GetDevicesCommunicationUnitsRequest = Message<"io.clbs.openhes.pbdeviceregistry.GetDevicesCommunicationUnitsRequest"> & {
  /**
   * The unique identifier of the device.
   *
   * @generated from field: repeated string device_id = 1;
   */
  deviceId: string[];
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.GetDevicesCommunicationUnitsRequest.
 * Use `create(GetDevicesCommunicationUnitsRequestSchema)` to create a new message.
 */
export const GetDevicesCommunicationUnitsRequestSchema: GenMessage<GetDevicesCommunicationUnitsRequest> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 17);

/**
 * DriverRegistry -> RestApi - the message holds the information about the communication units linked to the device.
 *
 * @generated from message io.clbs.openhes.pbdeviceregistry.GetDevicesCommunicationUnitsResponse
 */
export type GetDevicesCommunicationUnitsResponse = Message<"io.clbs.openhes.pbdeviceregistry.GetDevicesCommunicationUnitsResponse"> & {
  /**
   * The list of devices with their communication units.
   *
   * @generated from field: map<string, io.clbs.openhes.pbdriver.ArrayOfConnectionInfo> devices = 1;
   */
  devices: { [key: string]: ArrayOfConnectionInfo };
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.GetDevicesCommunicationUnitsResponse.
 * Use `create(GetDevicesCommunicationUnitsResponseSchema)` to create a new message.
 */
export const GetDevicesCommunicationUnitsResponseSchema: GenMessage<GetDevicesCommunicationUnitsResponse> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 18);

/**
 * Sub-message that represents the modem pool.
 *
 * @generated from message io.clbs.openhes.pbdeviceregistry.ModemPoolSpec
 */
export type ModemPoolSpec = Message<"io.clbs.openhes.pbdeviceregistry.ModemPoolSpec"> & {
  /**
   * The modem pool identifier.
   *
   * @generated from field: string pool_id = 1;
   */
  poolId: string;

  /**
   * The name of the modem pool.
   *
   * @generated from field: string name = 2;
   */
  name: string;
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.ModemPoolSpec.
 * Use `create(ModemPoolSpecSchema)` to create a new message.
 */
export const ModemPoolSpecSchema: GenMessage<ModemPoolSpec> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 19);

/**
 * RestApi -> DriverRegistry - the request message to get the modem pools.
 *
 * @generated from message io.clbs.openhes.pbdeviceregistry.GetModemPoolsResponse
 */
export type GetModemPoolsResponse = Message<"io.clbs.openhes.pbdeviceregistry.GetModemPoolsResponse"> & {
  /**
   * The list of modem pools.
   *
   * @generated from field: repeated io.clbs.openhes.pbdeviceregistry.ModemPoolSpec pools = 1;
   */
  pools: ModemPoolSpec[];
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.GetModemPoolsResponse.
 * Use `create(GetModemPoolsResponseSchema)` to create a new message.
 */
export const GetModemPoolsResponseSchema: GenMessage<GetModemPoolsResponse> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 20);

/**
 * RestApi -> DriverRegistry - the request message to get the modem pool info.
 *
 * @generated from message io.clbs.openhes.pbdeviceregistry.GetModemPoolRequest
 */
export type GetModemPoolRequest = Message<"io.clbs.openhes.pbdeviceregistry.GetModemPoolRequest"> & {
  /**
   * The modem pool identifier.
   *
   * @generated from field: string pool_id = 1;
   */
  poolId: string;
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.GetModemPoolRequest.
 * Use `create(GetModemPoolRequestSchema)` to create a new message.
 */
export const GetModemPoolRequestSchema: GenMessage<GetModemPoolRequest> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 21);

/**
 * DriverRegistry -> RestApi - the message holds the information about the modem pool.
 *
 * @generated from message io.clbs.openhes.pbdeviceregistry.GetModemPoolResponse
 */
export type GetModemPoolResponse = Message<"io.clbs.openhes.pbdeviceregistry.GetModemPoolResponse"> & {
  /**
   * The modems registered within the pool.
   *
   * @generated from field: repeated io.clbs.openhes.pbdriver.ModemInfo modems = 1;
   */
  modems: ModemInfo[];

  /**
   * The name of the modem pool.
   *
   * @generated from field: string name = 2;
   */
  name: string;
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.GetModemPoolResponse.
 * Use `create(GetModemPoolResponseSchema)` to create a new message.
 */
export const GetModemPoolResponseSchema: GenMessage<GetModemPoolResponse> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 22);

/**
 * @generated from message io.clbs.openhes.pbdeviceregistry.SetModemPoolRequest
 */
export type SetModemPoolRequest = Message<"io.clbs.openhes.pbdeviceregistry.SetModemPoolRequest"> & {
  /**
   * The modem pool identifier. It must be unique within the system.
   *
   * @generated from field: string pool_id = 1;
   */
  poolId: string;

  /**
   * The name of the modem pool to be created. It must be unique within the system.
   *
   * @generated from field: string name = 2;
   */
  name: string;
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.SetModemPoolRequest.
 * Use `create(SetModemPoolRequestSchema)` to create a new message.
 */
export const SetModemPoolRequestSchema: GenMessage<SetModemPoolRequest> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 23);

/**
 * @generated from message io.clbs.openhes.pbdeviceregistry.SetModemRequest
 */
export type SetModemRequest = Message<"io.clbs.openhes.pbdeviceregistry.SetModemRequest"> & {
  /**
   * The modem pool identifier, required for update operation.
   *
   * @generated from field: string pool_id = 1;
   */
  poolId: string;

  /**
   * The modem specification.
   *
   * @generated from field: io.clbs.openhes.pbdriver.ModemInfo modem = 2;
   */
  modem?: ModemInfo;
};

/**
 * Describes the message io.clbs.openhes.pbdeviceregistry.SetModemRequest.
 * Use `create(SetModemRequestSchema)` to create a new message.
 */
export const SetModemRequestSchema: GenMessage<SetModemRequest> = /*@__PURE__*/
  messageDesc(file_pbdeviceregistry_models, 24);

