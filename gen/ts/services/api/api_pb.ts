// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file services/api/api.proto (package io.clbs.openhes.services.api, edition 2023)
/* eslint-disable */

import type { GenFile, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { EmptySchema, StringValueSchema } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_empty, file_google_protobuf_wrappers } from "@bufbuild/protobuf/wkt";
import type { ListSelectorSchema } from "../../common/fields_pb";
import { file_common_fields } from "../../common/fields_pb";
import { file_common_types } from "../../common/types_pb";
import type { AddDevicesToGroupRequestSchema, BulkSchema, CommunicationUnitSchema, CreateBulkRequestSchema, CreateCommunicationUnitRequestSchema, CreateDeviceGroupRequestSchema, CreateDeviceRequestSchema, DeviceGroupSchema, DeviceSchema, DriverSchema, ListOfBulkSchema, ListOfCommunicationUnitSchema, ListOfDeviceGroupSchema, ListOfDeviceSchema, ListOfDriverSchema, ListOfModemPoolSchema, ModemPoolSchema, RemoveDevicesFromGroupRequestSchema, SetDeviceCommunicationUnitsRequestSchema, SetModemPoolRequestSchema, SetModemRequestSchema } from "../../acquisition/main_pb";
import { file_acquisition_main } from "../../acquisition/main_pb";
import type { SystemConfigSchema } from "../../system/main_pb";
import { file_system_main } from "../../system/main_pb";

/**
 * Describes the file services/api/api.proto.
 */
export const file_services_api_api: GenFile = /*@__PURE__*/
  fileDesc("ChZzZXJ2aWNlcy9hcGkvYXBpLnByb3RvEhxpby5jbGJzLm9wZW5oZXMuc2VydmljZXMuYXBpMoIXCgpBcGlTZXJ2aWNlEmEKCkNyZWF0ZUJ1bGsSNS5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkNyZWF0ZUJ1bGtSZXF1ZXN0GhwuZ29vZ2xlLnByb3RvYnVmLlN0cmluZ1ZhbHVlEmgKCUxpc3RCdWxrcxIrLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLkxpc3RTZWxlY3RvchouLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uTGlzdE9mQnVsaxJRCgdHZXRCdWxrEhwuZ29vZ2xlLnByb3RvYnVmLlN0cmluZ1ZhbHVlGiguaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5CdWxrEkIKCkNhbmNlbEJ1bGsSHC5nb29nbGUucHJvdG9idWYuU3RyaW5nVmFsdWUaFi5nb29nbGUucHJvdG9idWYuRW1wdHkSbAoLTGlzdERyaXZlcnMSKy5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmNvbW1vbi5MaXN0U2VsZWN0b3IaMC5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkxpc3RPZkRyaXZlchJVCglHZXREcml2ZXISHC5nb29nbGUucHJvdG9idWYuU3RyaW5nVmFsdWUaKi5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkRyaXZlchJ7ChdDcmVhdGVDb21tdW5pY2F0aW9uVW5pdBJCLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uQ3JlYXRlQ29tbXVuaWNhdGlvblVuaXRSZXF1ZXN0GhwuZ29vZ2xlLnByb3RvYnVmLlN0cmluZ1ZhbHVlEoEBChVMaXN0Q29tbXVuaWNhdGlvblVuaXQSKy5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmNvbW1vbi5MaXN0U2VsZWN0b3IaOy5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkxpc3RPZkNvbW11bmljYXRpb25Vbml0EmsKFEdldENvbW11bmljYXRpb25Vbml0EhwuZ29vZ2xlLnByb3RvYnVmLlN0cmluZ1ZhbHVlGjUuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5Db21tdW5pY2F0aW9uVW5pdBJlCgxDcmVhdGVEZXZpY2USNy5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkNyZWF0ZURldmljZVJlcXVlc3QaHC5nb29nbGUucHJvdG9idWYuU3RyaW5nVmFsdWUSbAoLTGlzdERldmljZXMSKy5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmNvbW1vbi5MaXN0U2VsZWN0b3IaMC5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkxpc3RPZkRldmljZRJVCglHZXREZXZpY2USHC5nb29nbGUucHJvdG9idWYuU3RyaW5nVmFsdWUaKi5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkRldmljZRJ9ChtTZXREZXZpY2VDb21tdW5pY2F0aW9uVW5pdHMSRi5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLlNldERldmljZUNvbW11bmljYXRpb25Vbml0c1JlcXVlc3QaFi5nb29nbGUucHJvdG9idWYuRW1wdHkSeQocR2V0RGV2aWNlc0NvbW11bmljYXRpb25Vbml0cxIcLmdvb2dsZS5wcm90b2J1Zi5TdHJpbmdWYWx1ZRo7LmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uTGlzdE9mQ29tbXVuaWNhdGlvblVuaXQSbwoRQ3JlYXRlRGV2aWNlR3JvdXASPC5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkNyZWF0ZURldmljZUdyb3VwUmVxdWVzdBocLmdvb2dsZS5wcm90b2J1Zi5TdHJpbmdWYWx1ZRJ2ChBMaXN0RGV2aWNlR3JvdXBzEisuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5jb21tb24uTGlzdFNlbGVjdG9yGjUuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5MaXN0T2ZEZXZpY2VHcm91cBJfCg5HZXREZXZpY2VHcm91cBIcLmdvb2dsZS5wcm90b2J1Zi5TdHJpbmdWYWx1ZRovLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uRGV2aWNlR3JvdXASaQoRQWRkRGV2aWNlc1RvR3JvdXASPC5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkFkZERldmljZXNUb0dyb3VwUmVxdWVzdBoWLmdvb2dsZS5wcm90b2J1Zi5FbXB0eRJzChZSZW1vdmVEZXZpY2VzRnJvbUdyb3VwEkEuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5SZW1vdmVEZXZpY2VzRnJvbUdyb3VwUmVxdWVzdBoWLmdvb2dsZS5wcm90b2J1Zi5FbXB0eRJyCg5MaXN0TW9kZW1Qb29scxIrLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLkxpc3RTZWxlY3RvchozLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uTGlzdE9mTW9kZW1Qb29sElsKDEdldE1vZGVtUG9vbBIcLmdvb2dsZS5wcm90b2J1Zi5TdHJpbmdWYWx1ZRotLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uTW9kZW1Qb29sEmgKD0NyZWF0ZU1vZGVtUG9vbBI3LmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uU2V0TW9kZW1Qb29sUmVxdWVzdBocLmdvb2dsZS5wcm90b2J1Zi5TdHJpbmdWYWx1ZRJiCg9VcGRhdGVNb2RlbVBvb2wSNy5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLlNldE1vZGVtUG9vbFJlcXVlc3QaFi5nb29nbGUucHJvdG9idWYuRW1wdHkSRwoPRGVsZXRlTW9kZW1Qb29sEhwuZ29vZ2xlLnByb3RvYnVmLlN0cmluZ1ZhbHVlGhYuZ29vZ2xlLnByb3RvYnVmLkVtcHR5EmAKC0NyZWF0ZU1vZGVtEjMuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5TZXRNb2RlbVJlcXVlc3QaHC5nb29nbGUucHJvdG9idWYuU3RyaW5nVmFsdWUSWgoLVXBkYXRlTW9kZW0SMy5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLlNldE1vZGVtUmVxdWVzdBoWLmdvb2dsZS5wcm90b2J1Zi5FbXB0eRJDCgtEZWxldGVNb2RlbRIcLmdvb2dsZS5wcm90b2J1Zi5TdHJpbmdWYWx1ZRoWLmdvb2dsZS5wcm90b2J1Zi5FbXB0eRJQCglHZXRDb25maWcSFi5nb29nbGUucHJvdG9idWYuRW1wdHkaKy5pby5jbGJzLm9wZW5oZXMubW9kZWxzLnN5c3RlbS5TeXN0ZW1Db25maWcSUAoJU2V0Q29uZmlnEisuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5zeXN0ZW0uU3lzdGVtQ29uZmlnGhYuZ29vZ2xlLnByb3RvYnVmLkVtcHR5QjZaNGdpdGh1Yi5jb20vY3licm9zbGFicy9oZXMtMi1hcGlzL2dlbi9nby9zZXJ2aWNlcy9hcGliCGVkaXRpb25zcOgH", [file_google_protobuf_empty, file_google_protobuf_wrappers, file_common_fields, file_common_types, file_acquisition_main, file_system_main]);

/**
 * The Dataproxy related service definition.
 *
 * @generated from service io.clbs.openhes.services.api.ApiService
 */
export const ApiService: GenService<{
  /**
   * @group: Bulks
   * @tag: acquisition
   * @tag: action
   * Starts a new bulk of jobs.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.CreateBulk
   */
  createBulk: {
    methodKind: "unary";
    input: typeof CreateBulkRequestSchema;
    output: typeof StringValueSchema;
  },
  /**
   * @group: Bulks
   * Retrieves the list of bulks.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.ListBulks
   */
  listBulks: {
    methodKind: "unary";
    input: typeof ListSelectorSchema;
    output: typeof ListOfBulkSchema;
  },
  /**
   * @group: Bulks
   * Retrieves the bulk info and status.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.GetBulk
   */
  getBulk: {
    methodKind: "unary";
    input: typeof StringValueSchema;
    output: typeof BulkSchema;
  },
  /**
   * @group: Bulks
   * Cancels the bulk of jobs.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.CancelBulk
   */
  cancelBulk: {
    methodKind: "unary";
    input: typeof StringValueSchema;
    output: typeof EmptySchema;
  },
  /**
   * @group: Driver Info
   * Retrieves the list of drivers.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.ListDrivers
   */
  listDrivers: {
    methodKind: "unary";
    input: typeof ListSelectorSchema;
    output: typeof ListOfDriverSchema;
  },
  /**
   * @group: Driver Info
   * Retrieves the driver.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.GetDriver
   */
  getDriver: {
    methodKind: "unary";
    input: typeof StringValueSchema;
    output: typeof DriverSchema;
  },
  /**
   * @group: Devices
   * @tag: communicationunit
   * The method called by the RestAPI to register a new communication unit. The parameter contains the communication unit specification.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.CreateCommunicationUnit
   */
  createCommunicationUnit: {
    methodKind: "unary";
    input: typeof CreateCommunicationUnitRequestSchema;
    output: typeof StringValueSchema;
  },
  /**
   * @group: Devices
   * @tag: communicationunit
   * The method called by the RestAPI to get the information about the communication unit. The parameter contains the search criteria.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.ListCommunicationUnit
   */
  listCommunicationUnit: {
    methodKind: "unary";
    input: typeof ListSelectorSchema;
    output: typeof ListOfCommunicationUnitSchema;
  },
  /**
   * @group: Devices
   * @tag: communicationunit
   * The method called by the RestAPI to get the information about the communication unit. The parameter contains the search criteria.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.GetCommunicationUnit
   */
  getCommunicationUnit: {
    methodKind: "unary";
    input: typeof StringValueSchema;
    output: typeof CommunicationUnitSchema;
  },
  /**
   * @group: Devices
   * @tag: device
   * The method called by the RestAPI to register a new device. The parameter contains the device specification.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.CreateDevice
   */
  createDevice: {
    methodKind: "unary";
    input: typeof CreateDeviceRequestSchema;
    output: typeof StringValueSchema;
  },
  /**
   * @group: Devices
   * @tag: device
   * The method called by the RestAPI to get the information about the device. The parameter contains the search criteria.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.ListDevices
   */
  listDevices: {
    methodKind: "unary";
    input: typeof ListSelectorSchema;
    output: typeof ListOfDeviceSchema;
  },
  /**
   * @group: Devices
   * @tag: device
   * The method called by the RestAPI to get the information about the device. The parameter contains the search criteria.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.GetDevice
   */
  getDevice: {
    methodKind: "unary";
    input: typeof StringValueSchema;
    output: typeof DeviceSchema;
  },
  /**
   * @group: Devices
   * @tag: device
   * The method called by the RestAPI to replace ordered set of linked communication units.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.SetDeviceCommunicationUnits
   */
  setDeviceCommunicationUnits: {
    methodKind: "unary";
    input: typeof SetDeviceCommunicationUnitsRequestSchema;
    output: typeof EmptySchema;
  },
  /**
   * @group: Devices
   * @tag: device
   * The method called by the RestAPI to get communication units definitions linked to the device(s).
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.GetDevicesCommunicationUnits
   */
  getDevicesCommunicationUnits: {
    methodKind: "unary";
    input: typeof StringValueSchema;
    output: typeof ListOfCommunicationUnitSchema;
  },
  /**
   * @group: Devices
   * @tag: devicegroup
   * The method called by the RestAPI to create a new device group. The parameter contains the device group specification.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.CreateDeviceGroup
   */
  createDeviceGroup: {
    methodKind: "unary";
    input: typeof CreateDeviceGroupRequestSchema;
    output: typeof StringValueSchema;
  },
  /**
   * @group: Devices
   * @tag: devicegroup
   * The method returns a list of device groups.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.ListDeviceGroups
   */
  listDeviceGroups: {
    methodKind: "unary";
    input: typeof ListSelectorSchema;
    output: typeof ListOfDeviceGroupSchema;
  },
  /**
   * @group: Devices
   * @tag: devicegroup
   * The method returns single device group.
   * @param The device group identifier.
   * @return The device group specification.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.GetDeviceGroup
   */
  getDeviceGroup: {
    methodKind: "unary";
    input: typeof StringValueSchema;
    output: typeof DeviceGroupSchema;
  },
  /**
   * @group: Devices
   * @tag: devicegroup
   * The method called by the RestAPI to add a new device to the device group. The parameter contains the device group specification.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.AddDevicesToGroup
   */
  addDevicesToGroup: {
    methodKind: "unary";
    input: typeof AddDevicesToGroupRequestSchema;
    output: typeof EmptySchema;
  },
  /**
   * @group: Devices
   * @tag: devicegroup
   * The method called by the RestAPI to remove a device from the device group. The parameter contains the device group specification.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.RemoveDevicesFromGroup
   */
  removeDevicesFromGroup: {
    methodKind: "unary";
    input: typeof RemoveDevicesFromGroupRequestSchema;
    output: typeof EmptySchema;
  },
  /**
   * @group: Devices
   * @tag: modempool
   * The method to get list of the modem pools.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.ListModemPools
   */
  listModemPools: {
    methodKind: "unary";
    input: typeof ListSelectorSchema;
    output: typeof ListOfModemPoolSchema;
  },
  /**
   * @group: Devices
   * @tag: modempool
   * The method to get the information about the modem pool. The method returns the modem pool information.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.GetModemPool
   */
  getModemPool: {
    methodKind: "unary";
    input: typeof StringValueSchema;
    output: typeof ModemPoolSchema;
  },
  /**
   * @group: Devices
   * @tag: modempool
   * The method to create a new modem pool.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.CreateModemPool
   */
  createModemPool: {
    methodKind: "unary";
    input: typeof SetModemPoolRequestSchema;
    output: typeof StringValueSchema;
  },
  /**
   * @group: Devices
   * @tag: modempool
   * The method to update the modem pool.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.UpdateModemPool
   */
  updateModemPool: {
    methodKind: "unary";
    input: typeof SetModemPoolRequestSchema;
    output: typeof EmptySchema;
  },
  /**
   * @group: Devices
   * @tag: modempool
   * The method to delete the modem pool.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.DeleteModemPool
   */
  deleteModemPool: {
    methodKind: "unary";
    input: typeof StringValueSchema;
    output: typeof EmptySchema;
  },
  /**
   * @group: Devices
   * @tag: modem
   * The method to create a new modem within the pool.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.CreateModem
   */
  createModem: {
    methodKind: "unary";
    input: typeof SetModemRequestSchema;
    output: typeof StringValueSchema;
  },
  /**
   * @group: Devices
   * @tag: modem
   * The method to update the modem within the pool.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.UpdateModem
   */
  updateModem: {
    methodKind: "unary";
    input: typeof SetModemRequestSchema;
    output: typeof EmptySchema;
  },
  /**
   * @group: Devices
   * @tag: modem
   * The method to delete the modem.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.DeleteModem
   */
  deleteModem: {
    methodKind: "unary";
    input: typeof StringValueSchema;
    output: typeof EmptySchema;
  },
  /**
   * @group: Configuration
   * The method to get the system configuration.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.GetConfig
   */
  getConfig: {
    methodKind: "unary";
    input: typeof EmptySchema;
    output: typeof SystemConfigSchema;
  },
  /**
   * @group: Configuration
   * The method to set the system configuration.
   *
   * @generated from rpc io.clbs.openhes.services.api.ApiService.SetConfig
   */
  setConfig: {
    methodKind: "unary";
    input: typeof SystemConfigSchema;
    output: typeof EmptySchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_services_api_api, 0);

