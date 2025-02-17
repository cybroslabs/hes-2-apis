// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file acquisition/main.proto (package io.clbs.openhes.models.acquisition, edition 2023)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_google_protobuf_empty, file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import type { MetadataFields } from "../common/metadata_pb";
import { file_common_metadata } from "../common/metadata_pb";
import type { ListSelector } from "../common/fields_pb";
import { file_common_fields } from "../common/fields_pb";
import type { BulkJob, BulkStatusCode, CommunicationBus, CommunicationUnit, CommunicationUnitSpec, DeviceCommunicationUnit, DeviceSpec, DeviceStatus, DriverTemplates, JobAction, JobSettings, ListOfJobDevice, ListOfJobDeviceId, ModemInfo } from "./shared_pb";
import { file_acquisition_shared } from "./shared_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file acquisition/main.proto.
 */
export const file_acquisition_main: GenFile = /*@__PURE__*/
  fileDesc("ChZhY3F1aXNpdGlvbi9tYWluLnByb3RvEiJpby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uIrEBCh5DcmVhdGVDb21tdW5pY2F0aW9uVW5pdFJlcXVlc3QSTgoEc3BlYxgBIAEoCzI5LmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uQ29tbXVuaWNhdGlvblVuaXRTcGVjQgWqAQIIAxI/CghtZXRhZGF0YRgCIAEoCzItLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLk1ldGFkYXRhRmllbGRzIl8KF0xpc3RPZkNvbW11bmljYXRpb25Vbml0EkQKBWl0ZW1zGAEgAygLMjUuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5Db21tdW5pY2F0aW9uVW5pdCJmCh1DcmVhdGVDb21tdW5pY2F0aW9uQnVzUmVxdWVzdBI/CghtZXRhZGF0YRgCIAEoCzItLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLk1ldGFkYXRhRmllbGRzSgQIARACIl0KFkxpc3RPZkNvbW11bmljYXRpb25CdXMSQwoFaXRlbXMYASADKAsyNC5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkNvbW11bmljYXRpb25CdXMibQouQWRkQ29tbXVuaWNhdGlvblVuaXRzVG9Db21tdW5pY2F0aW9uQnVzUmVxdWVzdBIcChRjb21tdW5pY2F0aW9uX2J1c19pZBgBIAEoCRIdChVjb21tdW5pY2F0aW9uX3VuaXRfaWQYAiADKAkicgozUmVtb3ZlQ29tbXVuaWNhdGlvblVuaXRzRnJvbUNvbW11bmljYXRpb25CdXNSZXF1ZXN0EhwKFGNvbW11bmljYXRpb25fYnVzX2lkGAEgASgJEh0KFWNvbW11bmljYXRpb25fdW5pdF9pZBgCIAMoCSKbAQoTQ3JlYXRlRGV2aWNlUmVxdWVzdBJDCgRzcGVjGAEgASgLMi4uaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5EZXZpY2VTcGVjQgWqAQIIAxI/CghtZXRhZGF0YRgCIAEoCzItLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLk1ldGFkYXRhRmllbGRzIkkKDExpc3RPZkRldmljZRI5CgVpdGVtcxgBIAMoCzIqLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uRGV2aWNlItABCgZEZXZpY2USQwoEc3BlYxgBIAEoCzIuLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uRGV2aWNlU3BlY0IFqgECCAMSQAoGc3RhdHVzGAIgASgLMjAuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5EZXZpY2VTdGF0dXMSPwoIbWV0YWRhdGEYAyABKAsyLS5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmNvbW1vbi5NZXRhZGF0YUZpZWxkcyKlAQoYQ3JlYXRlRGV2aWNlR3JvdXBSZXF1ZXN0EkgKBHNwZWMYASABKAsyMy5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkRldmljZUdyb3VwU3BlY0IFqgECCAMSPwoIbWV0YWRhdGEYAiABKAsyLS5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmNvbW1vbi5NZXRhZGF0YUZpZWxkcyJTChFMaXN0T2ZEZXZpY2VHcm91cBI+CgVpdGVtcxgBIAMoCzIvLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uRGV2aWNlR3JvdXAi7QEKEVN0cmVhbURldmljZUdyb3VwEkMKBHNwZWMYASABKAsyMy5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkRldmljZUdyb3VwU3BlY0gAEkcKBnN0YXR1cxgCIAEoCzI1LmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uRGV2aWNlR3JvdXBTdGF0dXNIABJBCghtZXRhZGF0YRgDIAEoCzItLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLk1ldGFkYXRhRmllbGRzSABCBwoFcGFydHMi3wEKC0RldmljZUdyb3VwEkgKBHNwZWMYASABKAsyMy5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkRldmljZUdyb3VwU3BlY0IFqgECCAMSRQoGc3RhdHVzGAIgASgLMjUuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5EZXZpY2VHcm91cFN0YXR1cxI/CghtZXRhZGF0YRgDIAEoCzItLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLk1ldGFkYXRhRmllbGRzImsKD0RldmljZUdyb3VwU3BlYxITCgtleHRlcm5hbF9pZBgBIAEoCRJDCg5keW5hbWljX2ZpbHRlchgCIAEoCzIrLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLkxpc3RTZWxlY3RvciLVAQoRRGV2aWNlR3JvdXBTdGF0dXMSUwoHZGV2aWNlcxgEIAMoCzJCLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uRGV2aWNlR3JvdXBTdGF0dXMuRGV2aWNlc0VudHJ5GmsKDERldmljZXNFbnRyeRILCgNrZXkYASABKAkSSgoFdmFsdWUYAiABKAsyOy5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkRldmljZUdyb3VwU3RhdHVzRGV2aWNlOgI4ASIuChdEZXZpY2VHcm91cFN0YXR1c0RldmljZRITCgtkcml2ZXJfdHlwZRgBIAEoCSKRAQoiU2V0RGV2aWNlQ29tbXVuaWNhdGlvblVuaXRzUmVxdWVzdBIRCglkZXZpY2VfaWQYASABKAkSWAoTY29tbXVuaWNhdGlvbl91bml0cxgCIAMoCzI7LmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uRGV2aWNlQ29tbXVuaWNhdGlvblVuaXQiawodTGlzdE9mRGV2aWNlQ29tbXVuaWNhdGlvblVuaXQSSgoFaXRlbXMYASADKAsyOy5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkRldmljZUNvbW11bmljYXRpb25Vbml0Ij8KGEFkZERldmljZXNUb0dyb3VwUmVxdWVzdBIQCghncm91cF9pZBgBIAEoCRIRCglkZXZpY2VfaWQYAiADKAkiRAodUmVtb3ZlRGV2aWNlc0Zyb21Hcm91cFJlcXVlc3QSEAoIZ3JvdXBfaWQYASABKAkSEQoJZGV2aWNlX2lkGAIgAygJIpcBChFDcmVhdGVCdWxrUmVxdWVzdBJBCgRzcGVjGAEgASgLMiwuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5CdWxrU3BlY0IFqgECCAMSPwoIbWV0YWRhdGEYAiABKAsyLS5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmNvbW1vbi5NZXRhZGF0YUZpZWxkcyJFCgpMaXN0T2ZCdWxrEjcKBWl0ZW1zGAEgAygLMiguaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5CdWxrIsoBCgRCdWxrEkEKBHNwZWMYASABKAsyLC5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkJ1bGtTcGVjQgWqAQIIAxI+CgZzdGF0dXMYAiABKAsyLi5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkJ1bGtTdGF0dXMSPwoIbWV0YWRhdGEYAyABKAsyLS5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmNvbW1vbi5NZXRhZGF0YUZpZWxkcyKNAwoIQnVsa1NwZWMSFgoOY29ycmVsYXRpb25faWQYAiABKAkSEwoLZHJpdmVyX3R5cGUYAyABKAkSSAoHZGV2aWNlcxgEIAEoCzI1LmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uTGlzdE9mSm9iRGV2aWNlSWRIABJNCg5jdXN0b21fZGV2aWNlcxgFIAEoCzIzLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uTGlzdE9mSm9iRGV2aWNlSAASGQoPZGV2aWNlX2dyb3VwX2lkGAkgASgJSAASQQoIc2V0dGluZ3MYBiABKAsyLy5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkpvYlNldHRpbmdzEj4KB2FjdGlvbnMYByADKAsyLS5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkpvYkFjdGlvbhITCgt3ZWJob29rX3VybBgIIAEoCUIICgZkZXZpY2UiiwEKCkJ1bGtTdGF0dXMSQgoGc3RhdHVzGAEgASgOMjIuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5CdWxrU3RhdHVzQ29kZRI5CgRqb2JzGAIgAygLMisuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5CdWxrSm9iIqQBChNTZXRNb2RlbVBvb2xSZXF1ZXN0EkYKBHNwZWMYASABKAsyMS5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLk1vZGVtUG9vbFNwZWNCBaoBAggDEj8KCG1ldGFkYXRhGAMgASgLMi0uaW8uY2xicy5vcGVuaGVzLm1vZGVscy5jb21tb24uTWV0YWRhdGFGaWVsZHNKBAgCEAMiTwoPTGlzdE9mTW9kZW1Qb29sEjwKBWl0ZW1zGAEgAygLMi0uaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5Nb2RlbVBvb2wiDwoNTW9kZW1Qb29sU3BlYyJQCg9Nb2RlbVBvb2xTdGF0dXMSPQoGbW9kZW1zGAEgAygLMi0uaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5Nb2RlbUluZm8i2QEKCU1vZGVtUG9vbBJGCgRzcGVjGAEgASgLMjEuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5Nb2RlbVBvb2xTcGVjQgWqAQIIAxJDCgZzdGF0dXMYAiABKAsyMy5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLk1vZGVtUG9vbFN0YXR1cxI/CghtZXRhZGF0YRgDIAEoCzItLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLk1ldGFkYXRhRmllbGRzImAKD1NldE1vZGVtUmVxdWVzdBIPCgdwb29sX2lkGAEgASgJEjwKBW1vZGVtGAIgASgLMi0uaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5Nb2RlbUluZm8iXAoJU2V0RHJpdmVyEkMKBHNwZWMYASABKAsyLi5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkRyaXZlclNwZWNCBaoBAggDSgQIAhADSgQIAxAEIkkKDExpc3RPZkRyaXZlchI5CgVpdGVtcxgBIAMoCzIqLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uRHJpdmVyIlkKBkRyaXZlchJDCgRzcGVjGAEgASgLMi4uaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5Ecml2ZXJTcGVjQgWqAQIIA0oECAIQA0oECAMQBCL7AQoKRHJpdmVyU3BlYxIPCgd2ZXJzaW9uGAEgASgJEhYKDmxpc3RlbmluZ19wb3J0GAIgASgNEhMKC2RyaXZlcl90eXBlGAMgASgJEhsKE21heF9jb25jdXJyZW50X2pvYnMYBCABKAUSGQoRbWF4X2Nhc2NhZGVfZGVwdGgYBSABKA0SGQoRdHlwaWNhbF9tZW1fdXNhZ2UYBiABKAUSRgoJdGVtcGxhdGVzGAcgASgLMjMuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5Ecml2ZXJUZW1wbGF0ZXMSFAoMZGlzcGxheV9uYW1lGAggASgJQjVaM2dpdGh1Yi5jb20vY3licm9zbGFicy9oZXMtMi1hcGlzL2dlbi9nby9hY3F1aXNpdGlvbmIIZWRpdGlvbnNw6Ac", [file_google_protobuf_empty, file_google_protobuf_timestamp, file_common_metadata, file_common_fields, file_acquisition_shared]);

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
 * @generated from message io.clbs.openhes.models.acquisition.CreateCommunicationBusRequest
 */
export type CreateCommunicationBusRequest = Message<"io.clbs.openhes.models.acquisition.CreateCommunicationBusRequest"> & {
  /**
   * The metadata fields.
   *
   * @generated from field: io.clbs.openhes.models.common.MetadataFields metadata = 2;
   */
  metadata?: MetadataFields;
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.CreateCommunicationBusRequest.
 * Use `create(CreateCommunicationBusRequestSchema)` to create a new message.
 */
export const CreateCommunicationBusRequestSchema: GenMessage<CreateCommunicationBusRequest> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 2);

/**
 * @generated from message io.clbs.openhes.models.acquisition.ListOfCommunicationBus
 */
export type ListOfCommunicationBus = Message<"io.clbs.openhes.models.acquisition.ListOfCommunicationBus"> & {
  /**
   * The list of communication unit buses.
   *
   * @generated from field: repeated io.clbs.openhes.models.acquisition.CommunicationBus items = 1;
   */
  items: CommunicationBus[];
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.ListOfCommunicationBus.
 * Use `create(ListOfCommunicationBusSchema)` to create a new message.
 */
export const ListOfCommunicationBusSchema: GenMessage<ListOfCommunicationBus> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 3);

/**
 * @generated from message io.clbs.openhes.models.acquisition.AddCommunicationUnitsToCommunicationBusRequest
 */
export type AddCommunicationUnitsToCommunicationBusRequest = Message<"io.clbs.openhes.models.acquisition.AddCommunicationUnitsToCommunicationBusRequest"> & {
  /**
   * The unique identifier of the communication bus.
   *
   * @generated from field: string communication_bus_id = 1;
   */
  communicationBusId: string;

  /**
   * The unique identifier of the communication unit.
   *
   * @generated from field: repeated string communication_unit_id = 2;
   */
  communicationUnitId: string[];
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.AddCommunicationUnitsToCommunicationBusRequest.
 * Use `create(AddCommunicationUnitsToCommunicationBusRequestSchema)` to create a new message.
 */
export const AddCommunicationUnitsToCommunicationBusRequestSchema: GenMessage<AddCommunicationUnitsToCommunicationBusRequest> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 4);

/**
 * @generated from message io.clbs.openhes.models.acquisition.RemoveCommunicationUnitsFromCommunicationBusRequest
 */
export type RemoveCommunicationUnitsFromCommunicationBusRequest = Message<"io.clbs.openhes.models.acquisition.RemoveCommunicationUnitsFromCommunicationBusRequest"> & {
  /**
   * The unique identifier of the communication bus.
   *
   * @generated from field: string communication_bus_id = 1;
   */
  communicationBusId: string;

  /**
   * The unique identifier of the communication unit.
   *
   * @generated from field: repeated string communication_unit_id = 2;
   */
  communicationUnitId: string[];
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.RemoveCommunicationUnitsFromCommunicationBusRequest.
 * Use `create(RemoveCommunicationUnitsFromCommunicationBusRequestSchema)` to create a new message.
 */
export const RemoveCommunicationUnitsFromCommunicationBusRequestSchema: GenMessage<RemoveCommunicationUnitsFromCommunicationBusRequest> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 5);

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
  messageDesc(file_acquisition_main, 6);

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
  messageDesc(file_acquisition_main, 7);

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
   * The device status.
   *
   * @generated from field: io.clbs.openhes.models.acquisition.DeviceStatus status = 2;
   */
  status?: DeviceStatus;

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
  messageDesc(file_acquisition_main, 8);

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
  messageDesc(file_acquisition_main, 9);

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
  messageDesc(file_acquisition_main, 10);

/**
 * @generated from message io.clbs.openhes.models.acquisition.StreamDeviceGroup
 */
export type StreamDeviceGroup = Message<"io.clbs.openhes.models.acquisition.StreamDeviceGroup"> & {
  /**
   * The oneof field containing the device group partial data.
   *
   * @generated from oneof io.clbs.openhes.models.acquisition.StreamDeviceGroup.parts
   */
  parts: {
    /**
     * The device group specification.
     *
     * @generated from field: io.clbs.openhes.models.acquisition.DeviceGroupSpec spec = 1;
     */
    value: DeviceGroupSpec;
    case: "spec";
  } | {
    /**
     * The device group status.
     *
     * @generated from field: io.clbs.openhes.models.acquisition.DeviceGroupStatus status = 2;
     */
    value: DeviceGroupStatus;
    case: "status";
  } | {
    /**
     * The metadata fields.
     *
     * @generated from field: io.clbs.openhes.models.common.MetadataFields metadata = 3;
     */
    value: MetadataFields;
    case: "metadata";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.StreamDeviceGroup.
 * Use `create(StreamDeviceGroupSchema)` to create a new message.
 */
export const StreamDeviceGroupSchema: GenMessage<StreamDeviceGroup> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 11);

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
  messageDesc(file_acquisition_main, 12);

/**
 * @generated from message io.clbs.openhes.models.acquisition.DeviceGroupSpec
 */
export type DeviceGroupSpec = Message<"io.clbs.openhes.models.acquisition.DeviceGroupSpec"> & {
  /**
   * The external identifier of the communication unit.
   *
   * @generated from field: string external_id = 1;
   */
  externalId: string;

  /**
   * The list selector used over Device object to dynamically select devices.
   *
   * @generated from field: io.clbs.openhes.models.common.ListSelector dynamic_filter = 2;
   */
  dynamicFilter?: ListSelector;
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.DeviceGroupSpec.
 * Use `create(DeviceGroupSpecSchema)` to create a new message.
 */
export const DeviceGroupSpecSchema: GenMessage<DeviceGroupSpec> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 13);

/**
 * @generated from message io.clbs.openhes.models.acquisition.DeviceGroupStatus
 */
export type DeviceGroupStatus = Message<"io.clbs.openhes.models.acquisition.DeviceGroupStatus"> & {
  /**
   * The list of devices that are part of the group. The key represents the device identifier, the value contains additional information.
   *
   * @generated from field: map<string, io.clbs.openhes.models.acquisition.DeviceGroupStatusDevice> devices = 4;
   */
  devices: { [key: string]: DeviceGroupStatusDevice };
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.DeviceGroupStatus.
 * Use `create(DeviceGroupStatusSchema)` to create a new message.
 */
export const DeviceGroupStatusSchema: GenMessage<DeviceGroupStatus> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 14);

/**
 * @generated from message io.clbs.openhes.models.acquisition.DeviceGroupStatusDevice
 */
export type DeviceGroupStatusDevice = Message<"io.clbs.openhes.models.acquisition.DeviceGroupStatusDevice"> & {
  /**
   * The driver type.
   *
   * @generated from field: string driver_type = 1;
   */
  driverType: string;
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.DeviceGroupStatusDevice.
 * Use `create(DeviceGroupStatusDeviceSchema)` to create a new message.
 */
export const DeviceGroupStatusDeviceSchema: GenMessage<DeviceGroupStatusDevice> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 15);

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
  messageDesc(file_acquisition_main, 16);

/**
 * @generated from message io.clbs.openhes.models.acquisition.ListOfDeviceCommunicationUnit
 */
export type ListOfDeviceCommunicationUnit = Message<"io.clbs.openhes.models.acquisition.ListOfDeviceCommunicationUnit"> & {
  /**
   * The list of linked communication units.
   *
   * @generated from field: repeated io.clbs.openhes.models.acquisition.DeviceCommunicationUnit items = 1;
   */
  items: DeviceCommunicationUnit[];
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.ListOfDeviceCommunicationUnit.
 * Use `create(ListOfDeviceCommunicationUnitSchema)` to create a new message.
 */
export const ListOfDeviceCommunicationUnitSchema: GenMessage<ListOfDeviceCommunicationUnit> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 17);

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
  messageDesc(file_acquisition_main, 18);

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
  messageDesc(file_acquisition_main, 19);

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
  messageDesc(file_acquisition_main, 20);

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
  messageDesc(file_acquisition_main, 21);

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
  messageDesc(file_acquisition_main, 22);

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
  } | {
    /**
     * The device group identifier.
     *
     * @generated from field: string device_group_id = 9;
     */
    value: string;
    case: "deviceGroupId";
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
  messageDesc(file_acquisition_main, 23);

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
  messageDesc(file_acquisition_main, 24);

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

  /**
   * The metadata fields.
   *
   * @generated from field: io.clbs.openhes.models.common.MetadataFields metadata = 3;
   */
  metadata?: MetadataFields;
};

/**
 * Describes the message io.clbs.openhes.models.acquisition.SetModemPoolRequest.
 * Use `create(SetModemPoolRequestSchema)` to create a new message.
 */
export const SetModemPoolRequestSchema: GenMessage<SetModemPoolRequest> = /*@__PURE__*/
  messageDesc(file_acquisition_main, 25);

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
  messageDesc(file_acquisition_main, 26);

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
  messageDesc(file_acquisition_main, 27);

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
  messageDesc(file_acquisition_main, 28);

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
  messageDesc(file_acquisition_main, 29);

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
  messageDesc(file_acquisition_main, 30);

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
  messageDesc(file_acquisition_main, 31);

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
  messageDesc(file_acquisition_main, 32);

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
  messageDesc(file_acquisition_main, 33);

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
  messageDesc(file_acquisition_main, 34);

