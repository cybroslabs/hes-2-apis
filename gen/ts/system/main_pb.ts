// @generated by protoc-gen-es v2.3.0 with parameter "target=ts,json_types=true"
// @generated from file system/main.proto (package io.clbs.openhes.models.system, edition 2023)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file system/main.proto.
 */
export const file_system_main: GenFile = /*@__PURE__*/
  fileDesc("ChFzeXN0ZW0vbWFpbi5wcm90bxIdaW8uY2xicy5vcGVuaGVzLm1vZGVscy5zeXN0ZW0ikwIKDFN5c3RlbUNvbmZpZxIUCgxtYXhfcmVwbGljYXMYASABKAUSIAoYbWF4X2Nhc2NhZGVfZGV2aWNlX2NvdW50GAIgASgFEhwKFG1heF9zbG90c19wZXJfZHJpdmVyGAMgASgFElIKDG1pbl9yZXBsaWNhcxgEIAMoCzI8LmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuc3lzdGVtLlN5c3RlbUNvbmZpZy5NaW5SZXBsaWNhc0VudHJ5EiUKHWRpc2FibGVfZGF0YV9wcm94eV9wcm9jZXNzaW5nGAUgASgIGjIKEE1pblJlcGxpY2FzRW50cnkSCwoDa2V5GAEgASgJEg0KBXZhbHVlGAIgASgFOgI4AUIwWi5naXRodWIuY29tL2N5YnJvc2xhYnMvaGVzLTItYXBpcy9nZW4vZ28vc3lzdGVtYghlZGl0aW9uc3DoBw");

/**
 * @generated from message io.clbs.openhes.models.system.SystemConfig
 */
export type SystemConfig = Message<"io.clbs.openhes.models.system.SystemConfig"> & {
  /**
   * The maximum number of replicas for all drivers.
   *    0 represents no active replicas will run, effectively disabling acquisition
   *   >0 represents the maximum number of replicas per driver
   *
   * @generated from field: int32 max_replicas = 1;
   */
  maxReplicas: number;

  /**
   * The maximum number of cascade devices for the driver. Minimum is 1.
   *
   * @generated from field: int32 max_cascade_device_count = 2;
   */
  maxCascadeDeviceCount: number;

  /**
   * The maximum number of slots per driver
   * @values:
   * - -1 represents unlimited number of slots, effecticaly using maximum number of slots supported by driver
   * -  0 represents no active slots will run, effectively disabling acquisition
   * - >0 represents the maximum number of slots per driver, the number of slots never exceeds the number of slots supported by driver
   *
   * @generated from field: int32 max_slots_per_driver = 3;
   */
  maxSlotsPerDriver: number;

  /**
   * The minimum number of replicas per type of driver.
   * The key is the driver type, the value is the minimum number of replicas.
   * The minimum replicas is guaranteed to be running at all times even if the total number of replicas exceeds the maximum number of replicas set in max_replicas.
   *
   * @generated from field: map<string, int32> min_replicas = 4;
   */
  minReplicas: { [key: string]: number };

  /**
   * Disable data proxy to process data from ouro temp tables.
   *
   * @generated from field: bool disable_data_proxy_processing = 5;
   */
  disableDataProxyProcessing: boolean;
};

/**
 * @generated from message io.clbs.openhes.models.system.SystemConfig
 */
export type SystemConfigJson = {
  /**
   * The maximum number of replicas for all drivers.
   *    0 represents no active replicas will run, effectively disabling acquisition
   *   >0 represents the maximum number of replicas per driver
   *
   * @generated from field: int32 max_replicas = 1;
   */
  maxReplicas?: number;

  /**
   * The maximum number of cascade devices for the driver. Minimum is 1.
   *
   * @generated from field: int32 max_cascade_device_count = 2;
   */
  maxCascadeDeviceCount?: number;

  /**
   * The maximum number of slots per driver
   * @values:
   * - -1 represents unlimited number of slots, effecticaly using maximum number of slots supported by driver
   * -  0 represents no active slots will run, effectively disabling acquisition
   * - >0 represents the maximum number of slots per driver, the number of slots never exceeds the number of slots supported by driver
   *
   * @generated from field: int32 max_slots_per_driver = 3;
   */
  maxSlotsPerDriver?: number;

  /**
   * The minimum number of replicas per type of driver.
   * The key is the driver type, the value is the minimum number of replicas.
   * The minimum replicas is guaranteed to be running at all times even if the total number of replicas exceeds the maximum number of replicas set in max_replicas.
   *
   * @generated from field: map<string, int32> min_replicas = 4;
   */
  minReplicas?: { [key: string]: number };

  /**
   * Disable data proxy to process data from ouro temp tables.
   *
   * @generated from field: bool disable_data_proxy_processing = 5;
   */
  disableDataProxyProcessing?: boolean;
};

/**
 * Describes the message io.clbs.openhes.models.system.SystemConfig.
 * Use `create(SystemConfigSchema)` to create a new message.
 */
export const SystemConfigSchema: GenMessage<SystemConfig, SystemConfigJson> = /*@__PURE__*/
  messageDesc(file_system_main, 0);

