// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file services/svcapi/api.proto (package io.clbs.openhes.services.svcapi, edition 2023)
/* eslint-disable */

import type { GenFile, GenService } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import type { EmptySchema, StringValueSchema } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_empty, file_google_protobuf_wrappers } from "@bufbuild/protobuf/wkt";
import type { ListOfFieldDescriptorSchema, ListSelectorSchema } from "../../common/fields_pb";
import { file_common_fields } from "../../common/fields_pb";
import { file_common_types } from "../../common/types_pb";
import type { AddCommunicationUnitsToCommunicationBusRequestSchema, AddDeviceConfigurationRegisterToDeviceConfigurationTemplateRequestSchema, AddDevicesToGroupRequestSchema, BulkJobSchema, BulkSchema, CreateBulkRequestSchema, CreateCommunicationBusRequestSchema, CreateCommunicationUnitRequestSchema, CreateDeviceConfigurationRegisterRequestSchema, CreateDeviceConfigurationTemplateRequestSchema, CreateDeviceGroupRequestSchema, CreateDeviceRequestSchema, CreateProxyBulkRequestSchema, DeviceConfigurationRegisterSchema, DeviceConfigurationTemplateSchema, DeviceGroupSchema, DeviceSchema, DriverSchema, ListOfBulkSchema, ListOfCommunicationBusSchema, ListOfCommunicationUnitSchema, ListOfDeviceCommunicationUnitSchema, ListOfDeviceConfigurationRegisterSchema, ListOfDeviceConfigurationTemplateSchema, ListOfDeviceGroupSchema, ListOfDeviceSchema, ListOfDriverSchema, ListOfModemPoolSchema, ModemPoolSchema, RemoveCommunicationUnitsFromCommunicationBusRequestSchema, RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplateRequestSchema, RemoveDevicesFromGroupRequestSchema, SetDeviceCommunicationUnitsRequestSchema, SetModemPoolRequestSchema, SetModemRequestSchema } from "../../acquisition/main_pb";
import { file_acquisition_main } from "../../acquisition/main_pb";
import type { CommunicationUnitSchema } from "../../acquisition/shared_pb";
import { file_acquisition_shared } from "../../acquisition/shared_pb";
import type { SystemConfigSchema } from "../../system/main_pb";
import { file_system_main } from "../../system/main_pb";

/**
 * Describes the file services/svcapi/api.proto.
 */
export const file_services_svcapi_api: GenFile = /*@__PURE__*/
  fileDesc("ChlzZXJ2aWNlcy9zdmNhcGkvYXBpLnByb3RvEh9pby5jbGJzLm9wZW5oZXMuc2VydmljZXMuc3ZjYXBpMvorCgpBcGlTZXJ2aWNlEo8BCiFDcmVhdGVEZXZpY2VDb25maWd1cmF0aW9uUmVnaXN0ZXISTC5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkNyZWF0ZURldmljZUNvbmZpZ3VyYXRpb25SZWdpc3RlclJlcXVlc3QaHC5nb29nbGUucHJvdG9idWYuU3RyaW5nVmFsdWUSlgEKIExpc3REZXZpY2VDb25maWd1cmF0aW9uUmVnaXN0ZXJzEisuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5jb21tb24uTGlzdFNlbGVjdG9yGkUuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5MaXN0T2ZEZXZpY2VDb25maWd1cmF0aW9uUmVnaXN0ZXISfwoeR2V0RGV2aWNlQ29uZmlndXJhdGlvblJlZ2lzdGVyEhwuZ29vZ2xlLnByb3RvYnVmLlN0cmluZ1ZhbHVlGj8uaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5EZXZpY2VDb25maWd1cmF0aW9uUmVnaXN0ZXISfAohVXBkYXRlRGV2aWNlQ29uZmlndXJhdGlvblJlZ2lzdGVyEj8uaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5EZXZpY2VDb25maWd1cmF0aW9uUmVnaXN0ZXIaFi5nb29nbGUucHJvdG9idWYuRW1wdHkSWQohRGVsZXRlRGV2aWNlQ29uZmlndXJhdGlvblJlZ2lzdGVyEhwuZ29vZ2xlLnByb3RvYnVmLlN0cmluZ1ZhbHVlGhYuZ29vZ2xlLnByb3RvYnVmLkVtcHR5Eo8BCiFDcmVhdGVEZXZpY2VDb25maWd1cmF0aW9uVGVtcGxhdGUSTC5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkNyZWF0ZURldmljZUNvbmZpZ3VyYXRpb25UZW1wbGF0ZVJlcXVlc3QaHC5nb29nbGUucHJvdG9idWYuU3RyaW5nVmFsdWUSlgEKIExpc3REZXZpY2VDb25maWd1cmF0aW9uVGVtcGxhdGVzEisuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5jb21tb24uTGlzdFNlbGVjdG9yGkUuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5MaXN0T2ZEZXZpY2VDb25maWd1cmF0aW9uVGVtcGxhdGUSfwoeR2V0RGV2aWNlQ29uZmlndXJhdGlvblRlbXBsYXRlEhwuZ29vZ2xlLnByb3RvYnVmLlN0cmluZ1ZhbHVlGj8uaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5EZXZpY2VDb25maWd1cmF0aW9uVGVtcGxhdGUSfAohVXBkYXRlRGV2aWNlQ29uZmlndXJhdGlvblRlbXBsYXRlEj8uaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5EZXZpY2VDb25maWd1cmF0aW9uVGVtcGxhdGUaFi5nb29nbGUucHJvdG9idWYuRW1wdHkSWQohRGVsZXRlRGV2aWNlQ29uZmlndXJhdGlvblRlbXBsYXRlEhwuZ29vZ2xlLnByb3RvYnVmLlN0cmluZ1ZhbHVlGhYuZ29vZ2xlLnByb3RvYnVmLkVtcHR5Er0BCjtBZGREZXZpY2VDb25maWd1cmF0aW9uUmVnaXN0ZXJUb0RldmljZUNvbmZpZ3VyYXRpb25UZW1wbGF0ZRJmLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uQWRkRGV2aWNlQ29uZmlndXJhdGlvblJlZ2lzdGVyVG9EZXZpY2VDb25maWd1cmF0aW9uVGVtcGxhdGVSZXF1ZXN0GhYuZ29vZ2xlLnByb3RvYnVmLkVtcHR5EscBCkBSZW1vdmVEZXZpY2VDb25maWd1cmF0aW9uUmVnaXN0ZXJGcm9tRGV2aWNlQ29uZmlndXJhdGlvblRlbXBsYXRlEmsuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5SZW1vdmVEZXZpY2VDb25maWd1cmF0aW9uUmVnaXN0ZXJGcm9tRGV2aWNlQ29uZmlndXJhdGlvblRlbXBsYXRlUmVxdWVzdBoWLmdvb2dsZS5wcm90b2J1Zi5FbXB0eRJkChRMaXN0RmllbGREZXNjcmlwdG9ycxIWLmdvb2dsZS5wcm90b2J1Zi5FbXB0eRo0LmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLkxpc3RPZkZpZWxkRGVzY3JpcHRvchJrCg9DcmVhdGVQcm94eUJ1bGsSOi5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkNyZWF0ZVByb3h5QnVsa1JlcXVlc3QaHC5nb29nbGUucHJvdG9idWYuU3RyaW5nVmFsdWUSYQoKQ3JlYXRlQnVsaxI1LmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uQ3JlYXRlQnVsa1JlcXVlc3QaHC5nb29nbGUucHJvdG9idWYuU3RyaW5nVmFsdWUSaAoJTGlzdEJ1bGtzEisuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5jb21tb24uTGlzdFNlbGVjdG9yGi4uaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5MaXN0T2ZCdWxrElEKB0dldEJ1bGsSHC5nb29nbGUucHJvdG9idWYuU3RyaW5nVmFsdWUaKC5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkJ1bGsSQgoKQ2FuY2VsQnVsaxIcLmdvb2dsZS5wcm90b2J1Zi5TdHJpbmdWYWx1ZRoWLmdvb2dsZS5wcm90b2J1Zi5FbXB0eRJXCgpHZXRCdWxrSm9iEhwuZ29vZ2xlLnByb3RvYnVmLlN0cmluZ1ZhbHVlGisuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5CdWxrSm9iEmwKC0xpc3REcml2ZXJzEisuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5jb21tb24uTGlzdFNlbGVjdG9yGjAuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5MaXN0T2ZEcml2ZXISVQoJR2V0RHJpdmVyEhwuZ29vZ2xlLnByb3RvYnVmLlN0cmluZ1ZhbHVlGiouaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5Ecml2ZXISewoXQ3JlYXRlQ29tbXVuaWNhdGlvblVuaXQSQi5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkNyZWF0ZUNvbW11bmljYXRpb25Vbml0UmVxdWVzdBocLmdvb2dsZS5wcm90b2J1Zi5TdHJpbmdWYWx1ZRKCAQoWTGlzdENvbW11bmljYXRpb25Vbml0cxIrLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLkxpc3RTZWxlY3Rvcho7LmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uTGlzdE9mQ29tbXVuaWNhdGlvblVuaXQSawoUR2V0Q29tbXVuaWNhdGlvblVuaXQSHC5nb29nbGUucHJvdG9idWYuU3RyaW5nVmFsdWUaNS5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkNvbW11bmljYXRpb25Vbml0EnkKFkNyZWF0ZUNvbW11bmljYXRpb25CdXMSQS5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkNyZWF0ZUNvbW11bmljYXRpb25CdXNSZXF1ZXN0GhwuZ29vZ2xlLnByb3RvYnVmLlN0cmluZ1ZhbHVlEoEBChZMaXN0Q29tbXVuaWNhdGlvbkJ1c2VzEisuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5jb21tb24uTGlzdFNlbGVjdG9yGjouaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5MaXN0T2ZDb21tdW5pY2F0aW9uQnVzEpUBCidBZGRDb21tdW5pY2F0aW9uVW5pdHNUb0NvbW11bmljYXRpb25CdXMSUi5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkFkZENvbW11bmljYXRpb25Vbml0c1RvQ29tbXVuaWNhdGlvbkJ1c1JlcXVlc3QaFi5nb29nbGUucHJvdG9idWYuRW1wdHkSnwEKLFJlbW92ZUNvbW11bmljYXRpb25Vbml0c0Zyb21Db21tdW5pY2F0aW9uQnVzElcuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5SZW1vdmVDb21tdW5pY2F0aW9uVW5pdHNGcm9tQ29tbXVuaWNhdGlvbkJ1c1JlcXVlc3QaFi5nb29nbGUucHJvdG9idWYuRW1wdHkSZQoMQ3JlYXRlRGV2aWNlEjcuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5DcmVhdGVEZXZpY2VSZXF1ZXN0GhwuZ29vZ2xlLnByb3RvYnVmLlN0cmluZ1ZhbHVlEmwKC0xpc3REZXZpY2VzEisuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5jb21tb24uTGlzdFNlbGVjdG9yGjAuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5MaXN0T2ZEZXZpY2USVQoJR2V0RGV2aWNlEhwuZ29vZ2xlLnByb3RvYnVmLlN0cmluZ1ZhbHVlGiouaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5EZXZpY2USfQobU2V0RGV2aWNlQ29tbXVuaWNhdGlvblVuaXRzEkYuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5TZXREZXZpY2VDb21tdW5pY2F0aW9uVW5pdHNSZXF1ZXN0GhYuZ29vZ2xlLnByb3RvYnVmLkVtcHR5En4KG0dldERldmljZUNvbW11bmljYXRpb25Vbml0cxIcLmdvb2dsZS5wcm90b2J1Zi5TdHJpbmdWYWx1ZRpBLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uTGlzdE9mRGV2aWNlQ29tbXVuaWNhdGlvblVuaXQSbwoRQ3JlYXRlRGV2aWNlR3JvdXASPC5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkNyZWF0ZURldmljZUdyb3VwUmVxdWVzdBocLmdvb2dsZS5wcm90b2J1Zi5TdHJpbmdWYWx1ZRJ2ChBMaXN0RGV2aWNlR3JvdXBzEisuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5jb21tb24uTGlzdFNlbGVjdG9yGjUuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5MaXN0T2ZEZXZpY2VHcm91cBJfCg5HZXREZXZpY2VHcm91cBIcLmdvb2dsZS5wcm90b2J1Zi5TdHJpbmdWYWx1ZRovLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uRGV2aWNlR3JvdXASaQoRQWRkRGV2aWNlc1RvR3JvdXASPC5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkFkZERldmljZXNUb0dyb3VwUmVxdWVzdBoWLmdvb2dsZS5wcm90b2J1Zi5FbXB0eRJzChZSZW1vdmVEZXZpY2VzRnJvbUdyb3VwEkEuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5SZW1vdmVEZXZpY2VzRnJvbUdyb3VwUmVxdWVzdBoWLmdvb2dsZS5wcm90b2J1Zi5FbXB0eRJ3ChZMaXN0RGV2aWNlR3JvdXBEZXZpY2VzEisuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5jb21tb24uTGlzdFNlbGVjdG9yGjAuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5MaXN0T2ZEZXZpY2UScgoOTGlzdE1vZGVtUG9vbHMSKy5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmNvbW1vbi5MaXN0U2VsZWN0b3IaMy5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLkxpc3RPZk1vZGVtUG9vbBJbCgxHZXRNb2RlbVBvb2wSHC5nb29nbGUucHJvdG9idWYuU3RyaW5nVmFsdWUaLS5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLk1vZGVtUG9vbBJoCg9DcmVhdGVNb2RlbVBvb2wSNy5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmFjcXVpc2l0aW9uLlNldE1vZGVtUG9vbFJlcXVlc3QaHC5nb29nbGUucHJvdG9idWYuU3RyaW5nVmFsdWUSYgoPVXBkYXRlTW9kZW1Qb29sEjcuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5TZXRNb2RlbVBvb2xSZXF1ZXN0GhYuZ29vZ2xlLnByb3RvYnVmLkVtcHR5EkcKD0RlbGV0ZU1vZGVtUG9vbBIcLmdvb2dsZS5wcm90b2J1Zi5TdHJpbmdWYWx1ZRoWLmdvb2dsZS5wcm90b2J1Zi5FbXB0eRJgCgtDcmVhdGVNb2RlbRIzLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuYWNxdWlzaXRpb24uU2V0TW9kZW1SZXF1ZXN0GhwuZ29vZ2xlLnByb3RvYnVmLlN0cmluZ1ZhbHVlEloKC1VwZGF0ZU1vZGVtEjMuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5hY3F1aXNpdGlvbi5TZXRNb2RlbVJlcXVlc3QaFi5nb29nbGUucHJvdG9idWYuRW1wdHkSQwoLRGVsZXRlTW9kZW0SHC5nb29nbGUucHJvdG9idWYuU3RyaW5nVmFsdWUaFi5nb29nbGUucHJvdG9idWYuRW1wdHkSUAoJR2V0Q29uZmlnEhYuZ29vZ2xlLnByb3RvYnVmLkVtcHR5GisuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5zeXN0ZW0uU3lzdGVtQ29uZmlnElAKCVNldENvbmZpZxIrLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuc3lzdGVtLlN5c3RlbUNvbmZpZxoWLmdvb2dsZS5wcm90b2J1Zi5FbXB0eUI5WjdnaXRodWIuY29tL2N5YnJvc2xhYnMvaGVzLTItYXBpcy9nZW4vZ28vc2VydmljZXMvc3ZjYXBpYghlZGl0aW9uc3DoBw", [file_google_protobuf_empty, file_google_protobuf_wrappers, file_common_fields, file_common_types, file_acquisition_main, file_acquisition_shared, file_system_main]);

/**
 * The Dataproxy related service definition.
 *
 * @generated from service io.clbs.openhes.services.svcapi.ApiService
 */
export const ApiService: GenService<{
  /**
   * @group: Device Configuration Register
   * Creates a new register. The register object holds the information about the single device register.
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.CreateDeviceConfigurationRegister
   */
  createDeviceConfigurationRegister: {
    methodKind: "unary";
    input: typeof CreateDeviceConfigurationRegisterRequestSchema;
    output: typeof StringValueSchema;
  },
  /**
   * @group: Device Configuration Register
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.ListDeviceConfigurationRegisters
   */
  listDeviceConfigurationRegisters: {
    methodKind: "unary";
    input: typeof ListSelectorSchema;
    output: typeof ListOfDeviceConfigurationRegisterSchema;
  },
  /**
   * @group: Device Configuration Register
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.GetDeviceConfigurationRegister
   */
  getDeviceConfigurationRegister: {
    methodKind: "unary";
    input: typeof StringValueSchema;
    output: typeof DeviceConfigurationRegisterSchema;
  },
  /**
   * @group: Device Configuration Register
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.UpdateDeviceConfigurationRegister
   */
  updateDeviceConfigurationRegister: {
    methodKind: "unary";
    input: typeof DeviceConfigurationRegisterSchema;
    output: typeof EmptySchema;
  },
  /**
   * @group: Device Configuration Register
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.DeleteDeviceConfigurationRegister
   */
  deleteDeviceConfigurationRegister: {
    methodKind: "unary";
    input: typeof StringValueSchema;
    output: typeof EmptySchema;
  },
  /**
   * @group: Device Configuration Templates
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.CreateDeviceConfigurationTemplate
   */
  createDeviceConfigurationTemplate: {
    methodKind: "unary";
    input: typeof CreateDeviceConfigurationTemplateRequestSchema;
    output: typeof StringValueSchema;
  },
  /**
   * @group: Device Configuration Templates
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.ListDeviceConfigurationTemplates
   */
  listDeviceConfigurationTemplates: {
    methodKind: "unary";
    input: typeof ListSelectorSchema;
    output: typeof ListOfDeviceConfigurationTemplateSchema;
  },
  /**
   * @group: Device Configuration Templates
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.GetDeviceConfigurationTemplate
   */
  getDeviceConfigurationTemplate: {
    methodKind: "unary";
    input: typeof StringValueSchema;
    output: typeof DeviceConfigurationTemplateSchema;
  },
  /**
   * @group: Device Configuration Templates
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.UpdateDeviceConfigurationTemplate
   */
  updateDeviceConfigurationTemplate: {
    methodKind: "unary";
    input: typeof DeviceConfigurationTemplateSchema;
    output: typeof EmptySchema;
  },
  /**
   * @group: Device Configuration Templates
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.DeleteDeviceConfigurationTemplate
   */
  deleteDeviceConfigurationTemplate: {
    methodKind: "unary";
    input: typeof StringValueSchema;
    output: typeof EmptySchema;
  },
  /**
   * @group: Device Configuration Templates
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.AddDeviceConfigurationRegisterToDeviceConfigurationTemplate
   */
  addDeviceConfigurationRegisterToDeviceConfigurationTemplate: {
    methodKind: "unary";
    input: typeof AddDeviceConfigurationRegisterToDeviceConfigurationTemplateRequestSchema;
    output: typeof EmptySchema;
  },
  /**
   * @group: Device Configuration Templates
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplate
   */
  removeDeviceConfigurationRegisterFromDeviceConfigurationTemplate: {
    methodKind: "unary";
    input: typeof RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplateRequestSchema;
    output: typeof EmptySchema;
  },
  /**
   * @group: Fields
   * The method to get the list of fields.
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.ListFieldDescriptors
   */
  listFieldDescriptors: {
    methodKind: "unary";
    input: typeof EmptySchema;
    output: typeof ListOfFieldDescriptorSchema;
  },
  /**
   * @group: Bulks
   * @tag: acquisition
   * @tag: action
   * Starts a new proxy bulk. The proxy bolk is a collection of jobs where each job represents a single device. Devices must be fully defined in the request.
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.CreateProxyBulk
   */
  createProxyBulk: {
    methodKind: "unary";
    input: typeof CreateProxyBulkRequestSchema;
    output: typeof StringValueSchema;
  },
  /**
   * @group: Bulks
   * @tag: acquisition
   * @tag: action
   * Starts a new bulk. The bulk is a collection of jobs where each jobs represents a single device. Devices that are part of the bulk are identified either as a list of registered device identifiers or as a group identifier.
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.CreateBulk
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.ListBulks
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.GetBulk
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.CancelBulk
   */
  cancelBulk: {
    methodKind: "unary";
    input: typeof StringValueSchema;
    output: typeof EmptySchema;
  },
  /**
   * @group: Bulks
   * Retrieves the job status.
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.GetBulkJob
   */
  getBulkJob: {
    methodKind: "unary";
    input: typeof StringValueSchema;
    output: typeof BulkJobSchema;
  },
  /**
   * @group: Driver Info
   * Retrieves the list of drivers.
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.ListDrivers
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.GetDriver
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.CreateCommunicationUnit
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.ListCommunicationUnits
   */
  listCommunicationUnits: {
    methodKind: "unary";
    input: typeof ListSelectorSchema;
    output: typeof ListOfCommunicationUnitSchema;
  },
  /**
   * @group: Devices
   * @tag: communicationunit
   * The method called by the RestAPI to get the information about the communication unit. The parameter contains the search criteria.
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.GetCommunicationUnit
   */
  getCommunicationUnit: {
    methodKind: "unary";
    input: typeof StringValueSchema;
    output: typeof CommunicationUnitSchema;
  },
  /**
   * @group: Devices
   * @tag: communicationbus
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.CreateCommunicationBus
   */
  createCommunicationBus: {
    methodKind: "unary";
    input: typeof CreateCommunicationBusRequestSchema;
    output: typeof StringValueSchema;
  },
  /**
   * @group: Devices
   * @tag: communicationbus
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.ListCommunicationBuses
   */
  listCommunicationBuses: {
    methodKind: "unary";
    input: typeof ListSelectorSchema;
    output: typeof ListOfCommunicationBusSchema;
  },
  /**
   * @group: Devices
   * @tag: communicationbus
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.AddCommunicationUnitsToCommunicationBus
   */
  addCommunicationUnitsToCommunicationBus: {
    methodKind: "unary";
    input: typeof AddCommunicationUnitsToCommunicationBusRequestSchema;
    output: typeof EmptySchema;
  },
  /**
   * @group: Devices
   * @tag: communicationbus
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.RemoveCommunicationUnitsFromCommunicationBus
   */
  removeCommunicationUnitsFromCommunicationBus: {
    methodKind: "unary";
    input: typeof RemoveCommunicationUnitsFromCommunicationBusRequestSchema;
    output: typeof EmptySchema;
  },
  /**
   * @group: Devices
   * @tag: device
   * The method called by the RestAPI to register a new device. The parameter contains the device specification.
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.CreateDevice
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.ListDevices
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.GetDevice
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.SetDeviceCommunicationUnits
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.GetDeviceCommunicationUnits
   */
  getDeviceCommunicationUnits: {
    methodKind: "unary";
    input: typeof StringValueSchema;
    output: typeof ListOfDeviceCommunicationUnitSchema;
  },
  /**
   * @group: Devices
   * @tag: devicegroup
   * The method called by the RestAPI to create a new device group. The parameter contains the device group specification.
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.CreateDeviceGroup
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.ListDeviceGroups
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.GetDeviceGroup
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.AddDevicesToGroup
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.RemoveDevicesFromGroup
   */
  removeDevicesFromGroup: {
    methodKind: "unary";
    input: typeof RemoveDevicesFromGroupRequestSchema;
    output: typeof EmptySchema;
  },
  /**
   * @group: Devices
   * @tag: devicegroup
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.ListDeviceGroupDevices
   */
  listDeviceGroupDevices: {
    methodKind: "unary";
    input: typeof ListSelectorSchema;
    output: typeof ListOfDeviceSchema;
  },
  /**
   * @group: Devices
   * @tag: modempool
   * The method to get list of the modem pools.
   *
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.ListModemPools
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.GetModemPool
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.CreateModemPool
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.UpdateModemPool
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.DeleteModemPool
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.CreateModem
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.UpdateModem
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.DeleteModem
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.GetConfig
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
   * @generated from rpc io.clbs.openhes.services.svcapi.ApiService.SetConfig
   */
  setConfig: {
    methodKind: "unary";
    input: typeof SystemConfigSchema;
    output: typeof EmptySchema;
  },
}> = /*@__PURE__*/
  serviceDesc(file_services_svcapi_api, 0);

