// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file acquisition/main.proto (package io.clbs.openhes.models.acquisition, edition 2023)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_google_protobuf_empty, file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import type { MetadataFields } from "../common/metadata_pb";
import { file_common_metadata } from "../common/metadata_pb";
import type { BulkJob, BulkStatusCode, ConnectionInfo, DeviceCommunicationUnit, DeviceSpec, DriverTemplates, JobAction, JobSettings, ListOfJobDevice, ListOfJobDeviceId, ModemInfo } from "./shared_pb";
import { file_acquisition_shared } from "./shared_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file acquisition/main.proto.
 */
export const file_acquisition_main: GenFile = /*@__PURE__*/
  fileDesc("ChZhY3F1aXNpdGlvbi9tYWluLnByb3RvEiJpby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uIrEBCh5DcmVhdGVDb21tdW5pY2F0aW9uVW5pdFJlcXVlc3QSTgoEc3BlYxgBIAEoCzI5LmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uQ29tbXVuaWNhdGlvblVuaXRTcGVjQgWqAQIIAxI/CghtZXRhZGF0YRgCIAEoCzItLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLk1ldGFkYXRhRmllbGRzIl8KF0xpc3RPZkNvbW11bmljYXRpb25Vbml0EkQKBWl0ZW1zGAEgAygLMjUuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5Db21tdW5pY2F0aW9uVW5pdCKqAQoRQ29tbXVuaWNhdGlvblVuaXQSTgoEc3BlYxgBIAEoCzI5LmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uQ29tbXVuaWNhdGlvblVuaXRTcGVjQgWqAQIIAxI/CghtZXRhZGF0YRgDIAEoCzItLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLk1ldGFkYXRhRmllbGRzSgQIAhADInkKFUNvbW11bmljYXRpb25Vbml0U3BlYxITCgtleHRlcm5hbF9pZBgBIAEoCRJLCg9jb25uZWN0aW9uX2luZm8YAiABKAsyMi5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkNvbm5lY3Rpb25JbmZvIpsBChNDcmVhdGVEZXZpY2VSZXF1ZXN0EkMKBHNwZWMYASABKAsyLi5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkRldmljZVNwZWNCBaoBAggDEj8KCG1ldGFkYXRhGAIgASgLMi0uaW8uY2xicy5vcGVuaGVzLm1vZGVscy5jb21tb24uTWV0YWRhdGFGaWVsZHMiSQoMTGlzdE9mRGV2aWNlEjkKBWl0ZW1zGAEgAygLMiouaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5EZXZpY2UilAEKBkRldmljZRJDCgRzcGVjGAEgASgLMi4uaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5EZXZpY2VTcGVjQgWqAQIIAxI/CghtZXRhZGF0YRgDIAEoCzItLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLk1ldGFkYXRhRmllbGRzSgQIAhADIqUBChhDcmVhdGVEZXZpY2VHcm91cFJlcXVlc3QSSAoEc3BlYxgBIAEoCzIzLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uRGV2aWNlR3JvdXBTcGVjQgWqAQIIAxI/CghtZXRhZGF0YRgCIAEoCzItLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLk1ldGFkYXRhRmllbGRzIlMKEUxpc3RPZkRldmljZUdyb3VwEj4KBWl0ZW1zGAEgAygLMi8uaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5EZXZpY2VHcm91cCLfAQoLRGV2aWNlR3JvdXASSAoEc3BlYxgBIAEoCzIzLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uRGV2aWNlR3JvdXBTcGVjQgWqAQIIAxJFCgZzdGF0dXMYAiABKAsyNS5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkRldmljZUdyb3VwU3RhdHVzEj8KCG1ldGFkYXRhGAMgASgLMi0uaW8uY2xicy5vcGVuaGVzLm1vZGVscy5jb21tb24uTWV0YWRhdGFGaWVsZHMiJgoPRGV2aWNlR3JvdXBTcGVjEhMKC2V4dGVybmFsX2lkGAIgASgJIiYKEURldmljZUdyb3VwU3RhdHVzEhEKCWRldmljZV9pZBgEIAMoCSKRAQoiU2V0RGV2aWNlQ29tbXVuaWNhdGlvblVuaXRzUmVxdWVzdBIRCglkZXZpY2VfaWQYASABKAkSWAoTY29tbXVuaWNhdGlvbl91bml0cxgCIAMoCzI7LmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uRGV2aWNlQ29tbXVuaWNhdGlvblVuaXQiPwoYQWRkRGV2aWNlc1RvR3JvdXBSZXF1ZXN0EhAKCGdyb3VwX2lkGAEgASgJEhEKCWRldmljZV9pZBgCIAMoCSJECh1SZW1vdmVEZXZpY2VzRnJvbUdyb3VwUmVxdWVzdBIQCghncm91cF9pZBgBIAEoCRIRCglkZXZpY2VfaWQYAiADKAkilwEKEUNyZWF0ZUJ1bGtSZXF1ZXN0EkEKBHNwZWMYASABKAsyLC5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkJ1bGtTcGVjQgWqAQIIAxI/CghtZXRhZGF0YRgCIAEoCzItLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLk1ldGFkYXRhRmllbGRzIkUKCkxpc3RPZkJ1bGsSNwoFaXRlbXMYASADKAsyKC5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkJ1bGsiygEKBEJ1bGsSQQoEc3BlYxgBIAEoCzIsLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uQnVsa1NwZWNCBaoBAggDEj4KBnN0YXR1cxgCIAEoCzIuLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uQnVsa1N0YXR1cxI/CghtZXRhZGF0YRgDIAEoCzItLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLk1ldGFkYXRhRmllbGRzIvICCghCdWxrU3BlYxIWCg5jb3JyZWxhdGlvbl9pZBgCIAEoCRITCgtkcml2ZXJfdHlwZRgDIAEoCRJICgdkZXZpY2VzGAQgASgLMjUuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5MaXN0T2ZKb2JEZXZpY2VJZEgAEk0KDmN1c3RvbV9kZXZpY2VzGAUgASgLMjMuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5MaXN0T2ZKb2JEZXZpY2VIABJBCghzZXR0aW5ncxgGIAEoCzIvLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uSm9iU2V0dGluZ3MSPgoHYWN0aW9ucxgHIAMoCzItLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uSm9iQWN0aW9uEhMKC3dlYmhvb2tfdXJsGAggASgJQggKBmRldmljZSKLAQoKQnVsa1N0YXR1cxJCCgZzdGF0dXMYASABKA4yMi5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkJ1bGtTdGF0dXNDb2RlEjkKBGpvYnMYAiADKAsyKy5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkJ1bGtKb2IiXQoTU2V0TW9kZW1Qb29sUmVxdWVzdBJGCgRzcGVjGAEgASgLMjEuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5Nb2RlbVBvb2xTcGVjQgWqAQIIAyJPCg9MaXN0T2ZNb2RlbVBvb2wSPAoFaXRlbXMYASADKAsyLS5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLk1vZGVtUG9vbCIPCg1Nb2RlbVBvb2xTcGVjIlAKD01vZGVtUG9vbFN0YXR1cxI9CgZtb2RlbXMYASADKAsyLS5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLk1vZGVtSW5mbyLZAQoJTW9kZW1Qb29sEkYKBHNwZWMYASABKAsyMS5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLk1vZGVtUG9vbFNwZWNCBaoBAggDEkMKBnN0YXR1cxgCIAEoCzIzLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uTW9kZW1Qb29sU3RhdHVzEj8KCG1ldGFkYXRhGAMgASgLMi0uaW8uY2xicy5vcGVuaGVzLm1vZGVscy5jb21tb24uTWV0YWRhdGFGaWVsZHMiYAoPU2V0TW9kZW1SZXF1ZXN0Eg8KB3Bvb2xfaWQYASABKAkSPAoFbW9kZW0YAiABKAsyLS5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLk1vZGVtSW5mbyJcCglTZXREcml2ZXISQwoEc3BlYxgBIAEoCzIuLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uRHJpdmVyU3BlY0IFqgECCANKBAgCEANKBAgDEAQiSQoMTGlzdE9mRHJpdmVyEjkKBWl0ZW1zGAEgAygLMiouaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5Ecml2ZXIiWQoGRHJpdmVyEkMKBHNwZWMYASABKAsyLi5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkRyaXZlclNwZWNCBaoBAggDSgQIAhADSgQIAxAEIvsBCgpEcml2ZXJTcGVjEg8KB3ZlcnNpb24YASABKAkSFgoObGlzdGVuaW5nX3BvcnQYAiABKA0SEwoLZHJpdmVyX3R5cGUYAyABKAkSGwoTbWF4X2NvbmN1cnJlbnRfam9icxgEIAEoBRIZChFtYXhfY2FzY2FkZV9kZXB0aBgFIAEoDRIZChF0eXBpY2FsX21lbV91c2FnZRgGIAEoBRJGCgl0ZW1wbGF0ZXMYByABKAsyMy5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkRyaXZlclRlbXBsYXRlcxIUCgxkaXNwbGF5X25hbWUYCCABKAlCNVozZ2l0aHViLmNvbS9jeWJyb3NsYWJzL2hlcy0yLWFwaXMvZ2VuL2dvL2FjcXVpc2l0aW9uYghlZGl0aW9uc3DoBw", [file_google_protobuf_empty, file_google_protobuf_timestamp, file_common_metadata, file_acquisition_shared]);

/**
 * @generated from message io.clbs.openhes.models.acquisition.CreateCommunicationUnitRequest
 */
export type CreateCommunicationUnitRequest = Message<"io.clbs.openhes.models.acquisition.CreateCommunicationUnitRequest"> & {
  /**
   * The communication unit specification.
   *
   * @generated from field: io.clbs.openhes.models.acquisition.CommunicationUnitSpec spec = 1 [features.field_presence = LEGACY_REQUIRED];
   */
  spec?: CommunicationUnitSpec;

  /**
   * The metadata fields.
   *
   * @generated from field: io.clbs.openhes.models.common.MetadataFields metadata = 2;
   */
  metadata?: MetadataFields;
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.CreateCommunicationUnitRequest.
 * Use `create(CreateCommunicationUnitRequestSchema)` to create a new message.
 */
export const CreateCommunicationUnitRequestSchema: GenMessage<CreateCommunicationUnitRequest> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 0);

/**
 * @generated from message io.clbs.openhes.models.acquisition.ListOfCommunicationUnit
 */
export type ListOfCommunicationUnit = Message<"io.clbs.openhes.models.acquisition.ListOfCommunicationUnit"> & {
  /**
   * The communication unit specification.
   *
   * @generated from field: repeated io.clbs.openhes.models.acquisition.CommunicationUnit items = 1;
   */
  items: CommunicationUnit[];
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.ListOfCommunicationUnit.
 * Use `create(ListOfCommunicationUnitSchema)` to create a new message.
 */
export const ListOfCommunicationUnitSchema: GenMessage<ListOfCommunicationUnit> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 1);

/**
 * @generated from message io.clbs.openhes.models.acquisition.CommunicationUnit
 */
export type CommunicationUnit = Message<"io.clbs.openhes.models.acquisition.CommunicationUnit"> & {
  /**
   * The communication unit specification.
   *
   * @generated from field: io.clbs.openhes.models.acquisition.CommunicationUnitSpec spec = 1 [features.field_presence = LEGACY_REQUIRED];
   */
  spec?: CommunicationUnitSpec;

  /**
   * The metadata fields.
   *
   * @generated from field: io.clbs.openhes.models.common.MetadataFields metadata = 3;
   */
  metadata?: MetadataFields;
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.CommunicationUnit.
 * Use `create(CommunicationUnitSchema)` to create a new message.
 */
export const CommunicationUnitSchema: GenMessage<CommunicationUnit> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 2);

/**
 * @generated from message io.clbs.openhes.models.acquisition.CommunicationUnitSpec
 */
export type CommunicationUnitSpec = Message<"io.clbs.openhes.models.acquisition.CommunicationUnitSpec"> & {
  /**
   * The external identifier of the communication unit.
   *
   * @generated from field: string external_id = 1;
   */
  externalId: string;

  /**
   * The connection info.
   *
   * @generated from field: io.clbs.openhes.models.acquisition.ConnectionInfo connection_info = 2;
   */
  connectionInfo?: ConnectionInfo;
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.CommunicationUnitSpec.
 * Use `create(CommunicationUnitSpecSchema)` to create a new message.
 */
export const CommunicationUnitSpecSchema: GenMessage<CommunicationUnitSpec> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 3);

/**
 * @generated from message io.clbs.openhes.models.acquisition.CreateDeviceRequest
 */
export type CreateDeviceRequest = Message<"io.clbs.openhes.models.acquisition.CreateDeviceRequest"> & {
  /**
   * The device specification.
   *
   * @generated from field: io.clbs.openhes.models.acquisition.DeviceSpec spec = 1 [features.field_presence = LEGACY_REQUIRED];
   */
  spec?: DeviceSpec;

  /**
   * The metadata fields.
   *
   * @generated from field: io.clbs.openhes.models.common.MetadataFields metadata = 2;
   */
  metadata?: MetadataFields;
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.CreateDeviceRequest.
 * Use `create(CreateDeviceRequestSchema)` to create a new message.
 */
export const CreateDeviceRequestSchema: GenMessage<CreateDeviceRequest> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 4);

/**
 * @generated from message io.clbs.openhes.models.acquisition.ListOfDevice
 */
export type ListOfDevice = Message<"io.clbs.openhes.models.acquisition.ListOfDevice"> & {
  /**
   * The list of devices.
   *
   * @generated from field: repeated io.clbs.openhes.models.acquisition.Device items = 1;
   */
  items: Device[];
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.ListOfDevice.
 * Use `create(ListOfDeviceSchema)` to create a new message.
 */
export const ListOfDeviceSchema: GenMessage<ListOfDevice> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 5);

/**
 * @generated from message io.clbs.openhes.models.acquisition.Device
 */
export type Device = Message<"io.clbs.openhes.models.acquisition.Device"> & {
  /**
   * The device specification.
   *
   * @generated from field: io.clbs.openhes.models.acquisition.DeviceSpec spec = 1 [features.field_presence = LEGACY_REQUIRED];
   */
  spec?: DeviceSpec;

  /**
   * The metadata fields.
   *
   * @generated from field: io.clbs.openhes.models.common.MetadataFields metadata = 3;
   */
  metadata?: MetadataFields;
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.Device.
 * Use `create(DeviceSchema)` to create a new message.
 */
export const DeviceSchema: GenMessage<Device> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 6);

/**
 * @generated from message io.clbs.openhes.models.acquisition.CreateDeviceGroupRequest
 */
export type CreateDeviceGroupRequest = Message<"io.clbs.openhes.models.acquisition.CreateDeviceGroupRequest"> & {
  /**
   * The device group specification.
   *
   * @generated from field: io.clbs.openhes.models.acquisition.DeviceGroupSpec spec = 1 [features.field_presence = LEGACY_REQUIRED];
   */
  spec?: DeviceGroupSpec;

  /**
   * The metadata fields.
   *
   * @generated from field: io.clbs.openhes.models.common.MetadataFields metadata = 2;
   */
  metadata?: MetadataFields;
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.CreateDeviceGroupRequest.
 * Use `create(CreateDeviceGroupRequestSchema)` to create a new message.
 */
export const CreateDeviceGroupRequestSchema: GenMessage<CreateDeviceGroupRequest> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 7);

/**
 * @generated from message io.clbs.openhes.models.acquisition.ListOfDeviceGroup
 */
export type ListOfDeviceGroup = Message<"io.clbs.openhes.models.acquisition.ListOfDeviceGroup"> & {
  /**
   * The list of device groups.
   *
   * @generated from field: repeated io.clbs.openhes.models.acquisition.DeviceGroup items = 1;
   */
  items: DeviceGroup[];
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.ListOfDeviceGroup.
 * Use `create(ListOfDeviceGroupSchema)` to create a new message.
 */
export const ListOfDeviceGroupSchema: GenMessage<ListOfDeviceGroup> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 8);

/**
 * @generated from message io.clbs.openhes.models.acquisition.DeviceGroup
 */
export type DeviceGroup = Message<"io.clbs.openhes.models.acquisition.DeviceGroup"> & {
  /**
   * The device group specification.
   *
   * @generated from field: io.clbs.openhes.models.acquisition.DeviceGroupSpec spec = 1 [features.field_presence = LEGACY_REQUIRED];
   */
  spec?: DeviceGroupSpec;

  /**
   * The device group status.
   *
   * @generated from field: io.clbs.openhes.models.acquisition.DeviceGroupStatus status = 2;
   */
  status?: DeviceGroupStatus;

  /**
   * The metadata fields.
   *
   * @generated from field: io.clbs.openhes.models.common.MetadataFields metadata = 3;
   */
  metadata?: MetadataFields;
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.DeviceGroup.
 * Use `create(DeviceGroupSchema)` to create a new message.
 */
export const DeviceGroupSchema: GenMessage<DeviceGroup> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 9);

/**
 * @generated from message io.clbs.openhes.models.acquisition.DeviceGroupSpec
 */
export type DeviceGroupSpec = Message<"io.clbs.openhes.models.acquisition.DeviceGroupSpec"> & {
  /**
   * The external identifier of the communication unit.
   *
   * @generated from field: string external_id = 2;
   */
  externalId: string;
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.DeviceGroupSpec.
 * Use `create(DeviceGroupSpecSchema)` to create a new message.
 */
export const DeviceGroupSpecSchema: GenMessage<DeviceGroupSpec> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 10);

/**
 * @generated from message io.clbs.openhes.models.acquisition.DeviceGroupStatus
 */
export type DeviceGroupStatus = Message<"io.clbs.openhes.models.acquisition.DeviceGroupStatus"> & {
  /**
   * The list of device identifiers that are part of the group.
   *
   * @generated from field: repeated string device_id = 4;
   */
  deviceId: string[];
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.DeviceGroupStatus.
 * Use `create(DeviceGroupStatusSchema)` to create a new message.
 */
export const DeviceGroupStatusSchema: GenMessage<DeviceGroupStatus> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 11);

/**
 * @generated from message io.clbs.openhes.models.acquisition.SetDeviceCommunicationUnitsRequest
 */
export type SetDeviceCommunicationUnitsRequest = Message<"io.clbs.openhes.models.acquisition.SetDeviceCommunicationUnitsRequest"> & {
  /**
   * The unique identifier of the device.
   *
   * @generated from field: string device_id = 1;
   */
  deviceId: string;

  /**
   * The list of linked communication units.
   *
   * @generated from field: repeated io.clbs.openhes.models.acquisition.DeviceCommunicationUnit communication_units = 2;
   */
  communicationUnits: DeviceCommunicationUnit[];
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.SetDeviceCommunicationUnitsRequest.
 * Use `create(SetDeviceCommunicationUnitsRequestSchema)` to create a new message.
 */
export const SetDeviceCommunicationUnitsRequestSchema: GenMessage<SetDeviceCommunicationUnitsRequest> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 12);

/**
 * @generated from message io.clbs.openhes.models.acquisition.AddDevicesToGroupRequest
 */
export type AddDevicesToGroupRequest = Message<"io.clbs.openhes.models.acquisition.AddDevicesToGroupRequest"> & {
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
 * Describes the message io.clbs.openhes.models.acquisition.AddDevicesToGroupRequest.
 * Use `create(AddDevicesToGroupRequestSchema)` to create a new message.
 */
export const AddDevicesToGroupRequestSchema: GenMessage<AddDevicesToGroupRequest> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 13);

/**
 * @generated from message io.clbs.openhes.models.acquisition.RemoveDevicesFromGroupRequest
 */
export type RemoveDevicesFromGroupRequest = Message<"io.clbs.openhes.models.acquisition.RemoveDevicesFromGroupRequest"> & {
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
 * Describes the message io.clbs.openhes.models.acquisition.RemoveDevicesFromGroupRequest.
 * Use `create(RemoveDevicesFromGroupRequestSchema)` to create a new message.
 */
export const RemoveDevicesFromGroupRequestSchema: GenMessage<RemoveDevicesFromGroupRequest> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 14);

/**
 * @generated from message io.clbs.openhes.models.acquisition.CreateBulkRequest
 */
export type CreateBulkRequest = Message<"io.clbs.openhes.models.acquisition.CreateBulkRequest"> & {
  /**
   * The bulk-job spec.
   *
   * @generated from field: io.clbs.openhes.models.acquisition.BulkSpec spec = 1 [features.field_presence = LEGACY_REQUIRED];
   */
  spec?: BulkSpec;

  /**
   * The metadata fields.
   *
   * @generated from field: io.clbs.openhes.models.common.MetadataFields metadata = 2;
   */
  metadata?: MetadataFields;
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.CreateBulkRequest.
 * Use `create(CreateBulkRequestSchema)` to create a new message.
 */
export const CreateBulkRequestSchema: GenMessage<CreateBulkRequest> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 15);

/**
 * @generated from message io.clbs.openhes.models.acquisition.ListOfBulk
 */
export type ListOfBulk = Message<"io.clbs.openhes.models.acquisition.ListOfBulk"> & {
  /**
   * The list of bulks.
   *
   * @generated from field: repeated io.clbs.openhes.models.acquisition.Bulk items = 1;
   */
  items: Bulk[];
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.ListOfBulk.
 * Use `create(ListOfBulkSchema)` to create a new message.
 */
export const ListOfBulkSchema: GenMessage<ListOfBulk> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 16);

/**
 * @generated from message io.clbs.openhes.models.acquisition.Bulk
 */
export type Bulk = Message<"io.clbs.openhes.models.acquisition.Bulk"> & {
  /**
   * The bulk-job spec.
   *
   * @generated from field: io.clbs.openhes.models.acquisition.BulkSpec spec = 1 [features.field_presence = LEGACY_REQUIRED];
   */
  spec?: BulkSpec;

  /**
   * The bulk-job status/data.
   *
   * @generated from field: io.clbs.openhes.models.acquisition.BulkStatus status = 2;
   */
  status?: BulkStatus;

  /**
   * The metadata fields.
   *
   * @generated from field: io.clbs.openhes.models.common.MetadataFields metadata = 3;
   */
  metadata?: MetadataFields;
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.Bulk.
 * Use `create(BulkSchema)` to create a new message.
 */
export const BulkSchema: GenMessage<Bulk> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 17);

/**
 * @generated from message io.clbs.openhes.models.acquisition.BulkSpec
 */
export type BulkSpec = Message<"io.clbs.openhes.models.acquisition.BulkSpec"> & {
  /**
   * @gqltype: UUID
   *
   * The correlation identifier, e.g. to define relation to non-homogenous group.
   *
   * @generated from field: string correlation_id = 2;
   */
  correlationId: string;

  /**
   * The device (driver) type.
   *
   * @generated from field: string driver_type = 3;
   */
  driverType: string;

  /**
   * @generated from oneof io.clbs.openhes.models.acquisition.BulkSpec.device
   */
  device: {
    /**
     * The list of devices in the bulk.
     *
     * @generated from field: io.clbs.openhes.models.acquisition.ListOfJobDeviceId devices = 4;
     */
    value: ListOfJobDeviceId;
    case: "devices";
  } | {
    /**
     * The list of custom devices in the bulk.
     *
     * @generated from field: io.clbs.openhes.models.acquisition.ListOfJobDevice custom_devices = 5;
     */
    value: ListOfJobDevice;
    case: "customDevices";
  } | { case: undefined; value?: undefined };

  /**
   * The bulk-shared job settings.
   *
   * @generated from field: io.clbs.openhes.models.acquisition.JobSettings settings = 6;
   */
  settings?: JobSettings;

  /**
   * The list actions to be executed.
   *
   * @generated from field: repeated io.clbs.openhes.models.acquisition.JobAction actions = 7;
   */
  actions: JobAction[];

  /**
   * The webhook URL to call when the bulk is completed.
   *
   * @generated from field: string webhook_url = 8;
   */
  webhookUrl: string;
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.BulkSpec.
 * Use `create(BulkSpecSchema)` to create a new message.
 */
export const BulkSpecSchema: GenMessage<BulkSpec> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 18);

/**
 * @generated from message io.clbs.openhes.models.acquisition.BulkStatus
 */
export type BulkStatus = Message<"io.clbs.openhes.models.acquisition.BulkStatus"> & {
  /**
   * The job status.
   *
   * @generated from field: io.clbs.openhes.models.acquisition.BulkStatusCode status = 1;
   */
  status: BulkStatusCode;

  /**
   * The list of job statuses.
   *
   * @generated from field: repeated io.clbs.openhes.models.acquisition.BulkJob jobs = 2;
   */
  jobs: BulkJob[];
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.BulkStatus.
 * Use `create(BulkStatusSchema)` to create a new message.
 */
export const BulkStatusSchema: GenMessage<BulkStatus> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 19);

/**
 * @generated from message io.clbs.openhes.models.acquisition.SetModemPoolRequest
 */
export type SetModemPoolRequest = Message<"io.clbs.openhes.models.acquisition.SetModemPoolRequest"> & {
  /**
   * The modem pool specification.
   *
   * @generated from field: io.clbs.openhes.models.acquisition.ModemPoolSpec spec = 1 [features.field_presence = LEGACY_REQUIRED];
   */
  spec?: ModemPoolSpec;
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.SetModemPoolRequest.
 * Use `create(SetModemPoolRequestSchema)` to create a new message.
 */
export const SetModemPoolRequestSchema: GenMessage<SetModemPoolRequest> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 20);

/**
 * @generated from message io.clbs.openhes.models.acquisition.ListOfModemPool
 */
export type ListOfModemPool = Message<"io.clbs.openhes.models.acquisition.ListOfModemPool"> & {
  /**
   * The list of modem pools.
   *
   * @generated from field: repeated io.clbs.openhes.models.acquisition.ModemPool items = 1;
   */
  items: ModemPool[];
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.ListOfModemPool.
 * Use `create(ListOfModemPoolSchema)` to create a new message.
 */
export const ListOfModemPoolSchema: GenMessage<ListOfModemPool> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 21);

/**
 * @generated from message io.clbs.openhes.models.acquisition.ModemPoolSpec
 */
export type ModemPoolSpec = Message<"io.clbs.openhes.models.acquisition.ModemPoolSpec"> & {
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.ModemPoolSpec.
 * Use `create(ModemPoolSpecSchema)` to create a new message.
 */
export const ModemPoolSpecSchema: GenMessage<ModemPoolSpec> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 22);

/**
 * @generated from message io.clbs.openhes.models.acquisition.ModemPoolStatus
 */
export type ModemPoolStatus = Message<"io.clbs.openhes.models.acquisition.ModemPoolStatus"> & {
  /**
   * The list of modems in the pool.
   *
   * @generated from field: repeated io.clbs.openhes.models.acquisition.ModemInfo modems = 1;
   */
  modems: ModemInfo[];
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.ModemPoolStatus.
 * Use `create(ModemPoolStatusSchema)` to create a new message.
 */
export const ModemPoolStatusSchema: GenMessage<ModemPoolStatus> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 23);

/**
 * @generated from message io.clbs.openhes.models.acquisition.ModemPool
 */
export type ModemPool = Message<"io.clbs.openhes.models.acquisition.ModemPool"> & {
  /**
   * The modem pool specification.
   *
   * @generated from field: io.clbs.openhes.models.acquisition.ModemPoolSpec spec = 1 [features.field_presence = LEGACY_REQUIRED];
   */
  spec?: ModemPoolSpec;

  /**
   * The modem pool status.
   *
   * @generated from field: io.clbs.openhes.models.acquisition.ModemPoolStatus status = 2;
   */
  status?: ModemPoolStatus;

  /**
   * The metadata fields.
   *
   * @generated from field: io.clbs.openhes.models.common.MetadataFields metadata = 3;
   */
  metadata?: MetadataFields;
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.ModemPool.
 * Use `create(ModemPoolSchema)` to create a new message.
 */
export const ModemPoolSchema: GenMessage<ModemPool> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 24);

/**
 * @generated from message io.clbs.openhes.models.acquisition.SetModemRequest
 */
export type SetModemRequest = Message<"io.clbs.openhes.models.acquisition.SetModemRequest"> & {
  /**
   * The modem pool identifier, required for update operation.
   *
   * @generated from field: string pool_id = 1;
   */
  poolId: string;

  /**
   * The modem specification.
   *
   * @generated from field: io.clbs.openhes.models.acquisition.ModemInfo modem = 2;
   */
  modem?: ModemInfo;
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.SetModemRequest.
 * Use `create(SetModemRequestSchema)` to create a new message.
 */
export const SetModemRequestSchema: GenMessage<SetModemRequest> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 25);

/**
 * @generated from message io.clbs.openhes.models.acquisition.SetDriver
 */
export type SetDriver = Message<"io.clbs.openhes.models.acquisition.SetDriver"> & {
  /**
   * The driver specification.
   *
   * @generated from field: io.clbs.openhes.models.acquisition.DriverSpec spec = 1 [features.field_presence = LEGACY_REQUIRED];
   */
  spec?: DriverSpec;
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.SetDriver.
 * Use `create(SetDriverSchema)` to create a new message.
 */
export const SetDriverSchema: GenMessage<SetDriver> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 26);

/**
 * @generated from message io.clbs.openhes.models.acquisition.ListOfDriver
 */
export type ListOfDriver = Message<"io.clbs.openhes.models.acquisition.ListOfDriver"> & {
  /**
   * The list of drivers.
   *
   * @generated from field: repeated io.clbs.openhes.models.acquisition.Driver items = 1;
   */
  items: Driver[];
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.ListOfDriver.
 * Use `create(ListOfDriverSchema)` to create a new message.
 */
export const ListOfDriverSchema: GenMessage<ListOfDriver> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 27);

/**
 * @generated from message io.clbs.openhes.models.acquisition.Driver
 */
export type Driver = Message<"io.clbs.openhes.models.acquisition.Driver"> & {
  /**
   * The driver specification.
   *
   * @generated from field: io.clbs.openhes.models.acquisition.DriverSpec spec = 1 [features.field_presence = LEGACY_REQUIRED];
   */
  spec?: DriverSpec;
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.Driver.
 * Use `create(DriverSchema)` to create a new message.
 */
export const DriverSchema: GenMessage<Driver> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 28);

/**
 * @generated from message io.clbs.openhes.models.acquisition.DriverSpec
 */
export type DriverSpec = Message<"io.clbs.openhes.models.acquisition.DriverSpec"> & {
  /**
   * The driver version. The format is not defined. The driver itself is versioned by the docker image tags so this value shall be either the same (set during the image build) or any useful user-readable version string.
   *
   * @generated from field: string version = 1;
   */
  version: string;

  /**
   * The port the driver's gRPC will listen on.
   *
   * @generated from field: uint32 listening_port = 2;
   */
  listeningPort: number;

  /**
   * The technical/internal ID of the driver.
   *
   * @generated from field: string driver_type = 3;
   */
  driverType: string;

  /**
   * The maximum number of concurrent jobs the driver can handle. The value 0 is not allowed, the maximum number respect typical_mem_usage not to overgrow the memory resources!
   *
   * @generated from field: int32 max_concurrent_jobs = 4;
   */
  maxConcurrentJobs: number;

  /**
   * The maximum cascade depth the driver can handle. Number 1 means that the driver cannot handle cascading jobs, 2 means that the driver can handle cascading jobs with one level of depth, etc.
   * The value 0 means that the driver can handle any number of cascading jobs.
   *
   * @generated from field: uint32 max_cascade_depth = 5;
   */
  maxCascadeDepth: number;

  /**
   * The typical memory usage of the driver in MB.
   *
   * @generated from field: int32 typical_mem_usage = 6;
   */
  typicalMemUsage: number;

  /**
   * The connection and action templates.
   *
   * @generated from field: io.clbs.openhes.models.acquisition.DriverTemplates templates = 7;
   */
  templates?: DriverTemplates;

  /**
   * The display name of the driver. Must be in format '<manufacturer> <device_type> [<device_type_version>]'.
   * It must respect upper/lower characters.
   * The generic drivers, such as 'cybros labs generic', must be named as '<driver_company_name> generic'.
   *
   * Examples: 'Addax NP73E', 'cybros labs generic', 'Landis+Gyr S650 v2'
   *
   * @generated from field: string display_name = 8;
   */
  displayName: string;
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.DriverSpec.
 * Use `create(DriverSpecSchema)` to create a new message.
 */
export const DriverSpecSchema: GenMessage<DriverSpec> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 29);

