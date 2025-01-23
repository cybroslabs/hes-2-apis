// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file pbdataproxy-models.proto (package io.clbs.openhes.pbdataproxy, edition 2023)
/* eslint-disable */

import type { GenEnum, GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { enumDesc, fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_google_protobuf_empty, file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import type { JobDevice, JobStatus } from "./pbtaskmaster-models_pb";
import { file_pbtaskmaster_models } from "./pbtaskmaster-models_pb";
import type { JobAction, JobSettings } from "./pbdriver-models_pb";
import { file_pbdriver_models } from "./pbdriver-models_pb";
import type { MetadataFields } from "./models_pb";
import { file_models } from "./models_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file pbdataproxy-models.proto.
 */
export const file_pbdataproxy_models: GenFile = /*@__PURE__*/
  fileDesc("ChhwYmRhdGFwcm94eS1tb2RlbHMucHJvdG8SG2lvLmNsYnMub3Blbmhlcy5wYmRhdGFwcm94eSKCAQoRQ3JlYXRlQnVsa1JlcXVlc3QSMwoEc3BlYxgBIAEoCzIlLmlvLmNsYnMub3Blbmhlcy5wYmRhdGFwcm94eS5CdWxrU3BlYxI4CghtZXRhZGF0YRgDIAEoCzImLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuTWV0YWRhdGFGaWVsZHMiqwIKCEJ1bGtTcGVjEg8KB2J1bGtfaWQYASABKAkSFgoOY29ycmVsYXRpb25faWQYAiABKAkSDgoGb3JnX2lkGAMgASgJEhMKC2RyaXZlcl90eXBlGAQgASgJEjgKB2RldmljZXMYBSADKAsyJy5pby5jbGJzLm9wZW5oZXMucGJ0YXNrbWFzdGVyLkpvYkRldmljZRI3CghzZXR0aW5ncxgGIAEoCzIlLmlvLmNsYnMub3Blbmhlcy5wYmRyaXZlci5Kb2JTZXR0aW5ncxI4Cgtqb2JfYWN0aW9ucxgHIAMoCzIjLmlvLmNsYnMub3Blbmhlcy5wYmRyaXZlci5Kb2JBY3Rpb24SEwoLd2ViaG9va191cmwYCCABKAkSDwoHdXNlcl9pZBgJIAEoCSJ9CgpCdWxrU3RhdHVzEjsKBnN0YXR1cxgBIAEoDjIrLmlvLmNsYnMub3Blbmhlcy5wYmRhdGFwcm94eS5CdWxrU3RhdHVzQ29kZRIyCgRqb2JzGAIgAygLMiQuaW8uY2xicy5vcGVuaGVzLnBiZGF0YXByb3h5LkJ1bGtKb2IiUgoHQnVsa0pvYhIOCgZqb2JfaWQYASABKAkSNwoGc3RhdHVzGAIgASgLMicuaW8uY2xicy5vcGVuaGVzLnBidGFza21hc3Rlci5Kb2JTdGF0dXMirgEKBEJ1bGsSMwoEc3BlYxgBIAEoCzIlLmlvLmNsYnMub3Blbmhlcy5wYmRhdGFwcm94eS5CdWxrU3BlYxI3CgZzdGF0dXMYAiABKAsyJy5pby5jbGJzLm9wZW5oZXMucGJkYXRhcHJveHkuQnVsa1N0YXR1cxI4CghtZXRhZGF0YRgDIAEoCzImLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuTWV0YWRhdGFGaWVsZHMiPgoKTGlzdE9mQnVsaxIwCgVpdGVtcxgBIAMoCzIhLmlvLmNsYnMub3Blbmhlcy5wYmRhdGFwcm94eS5CdWxrKpABCg5CdWxrU3RhdHVzQ29kZRIWChJCVUxLX1NUQVRVU19RVUVVRUQQABIXChNCVUxLX1NUQVRVU19SVU5OSU5HEAESGQoVQlVMS19TVEFUVVNfQ09NUExFVEVEEAISGQoVQlVMS19TVEFUVVNfQ0FOQ0VMTEVEEAMSFwoTQlVMS19TVEFUVVNfRVhQSVJFRBAEQj1aO2dpdGh1Yi5jb20vY3licm9zbGFicy9oZXMtMi1hcGlzL3Byb3RvYnVmL3BiZGF0YXByb3h5bW9kZWxzYghlZGl0aW9uc3DoBw", [file_google_protobuf_empty, file_google_protobuf_timestamp, file_pbtaskmaster_models, file_pbdriver_models, file_models]);

/**
 * RestApi -> DataProxy
 *
 * @generated from message io.clbs.openhes.pbdataproxy.CreateBulkRequest
 */
export type CreateBulkRequest = Message<"io.clbs.openhes.pbdataproxy.CreateBulkRequest"> & {
  /**
   * The bulk-job spec.
   *
   * @generated from field: io.clbs.openhes.pbdataproxy.BulkSpec spec = 1;
   */
  spec?: BulkSpec;

  /**
   * The metadata fields.
   *
   * @generated from field: io.clbs.openhes.models.MetadataFields metadata = 3;
   */
  metadata?: MetadataFields;
};

/**
 * Describes the message io.clbs.openhes.pbdataproxy.CreateBulkRequest.
 * Use `create(CreateBulkRequestSchema)` to create a new message.
 */
export const CreateBulkRequestSchema: GenMessage<CreateBulkRequest> = /*@__PURE__*/
  messageDesc(file_pbdataproxy_models, 0);

/**
 * Sub-message - holds the bulk job specification.
 *
 * @generated from message io.clbs.openhes.pbdataproxy.BulkSpec
 */
export type BulkSpec = Message<"io.clbs.openhes.pbdataproxy.BulkSpec"> & {
  /**
   * The bulk identifier.
   *
   * @generated from field: string bulk_id = 1;
   */
  bulkId: string;

  /**
   * The correlation identifier, e.g. to define relation to non-homogenous group.
   *
   * @generated from field: string correlation_id = 2;
   */
  correlationId: string;

  /**
   * The organization identifier.
   *
   * @generated from field: string org_id = 3;
   */
  orgId: string;

  /**
   * The device (driver) type.
   *
   * @generated from field: string driver_type = 4;
   */
  driverType: string;

  /**
   * The list of devices in the bulk.
   *
   * @generated from field: repeated io.clbs.openhes.pbtaskmaster.JobDevice devices = 5;
   */
  devices: JobDevice[];

  /**
   * The bulk-shared job settings.
   *
   * @generated from field: io.clbs.openhes.pbdriver.JobSettings settings = 6;
   */
  settings?: JobSettings;

  /**
   * The list actions to be executed.
   *
   * @generated from field: repeated io.clbs.openhes.pbdriver.JobAction job_actions = 7;
   */
  jobActions: JobAction[];

  /**
   * The webhook URL to call when the bulk is completed.
   *
   * @generated from field: string webhook_url = 8;
   */
  webhookUrl: string;

  /**
   * The user identifier.
   *
   * @generated from field: string user_id = 9;
   */
  userId: string;
};

/**
 * Describes the message io.clbs.openhes.pbdataproxy.BulkSpec.
 * Use `create(BulkSpecSchema)` to create a new message.
 */
export const BulkSpecSchema: GenMessage<BulkSpec> = /*@__PURE__*/
  messageDesc(file_pbdataproxy_models, 1);

/**
 * Sub-message - holds the bulk job status.
 *
 * @generated from message io.clbs.openhes.pbdataproxy.BulkStatus
 */
export type BulkStatus = Message<"io.clbs.openhes.pbdataproxy.BulkStatus"> & {
  /**
   * The job status.
   *
   * @generated from field: io.clbs.openhes.pbdataproxy.BulkStatusCode status = 1;
   */
  status: BulkStatusCode;

  /**
   * The list of job statuses.
   *
   * @generated from field: repeated io.clbs.openhes.pbdataproxy.BulkJob jobs = 2;
   */
  jobs: BulkJob[];
};

/**
 * Describes the message io.clbs.openhes.pbdataproxy.BulkStatus.
 * Use `create(BulkStatusSchema)` to create a new message.
 */
export const BulkStatusSchema: GenMessage<BulkStatus> = /*@__PURE__*/
  messageDesc(file_pbdataproxy_models, 2);

/**
 * @generated from message io.clbs.openhes.pbdataproxy.BulkJob
 */
export type BulkJob = Message<"io.clbs.openhes.pbdataproxy.BulkJob"> & {
  /**
   * The job identifier.
   *
   * @generated from field: string job_id = 1;
   */
  jobId: string;

  /**
   * The job status.
   *
   * @generated from field: io.clbs.openhes.pbtaskmaster.JobStatus status = 2;
   */
  status?: JobStatus;
};

/**
 * Describes the message io.clbs.openhes.pbdataproxy.BulkJob.
 * Use `create(BulkJobSchema)` to create a new message.
 */
export const BulkJobSchema: GenMessage<BulkJob> = /*@__PURE__*/
  messageDesc(file_pbdataproxy_models, 3);

/**
 * DataProxy -> RestApi - the message holds the bulk info and it's status.
 *
 * @generated from message io.clbs.openhes.pbdataproxy.Bulk
 */
export type Bulk = Message<"io.clbs.openhes.pbdataproxy.Bulk"> & {
  /**
   * The bulk-job spec.
   *
   * @generated from field: io.clbs.openhes.pbdataproxy.BulkSpec spec = 1;
   */
  spec?: BulkSpec;

  /**
   * The bulk-job status/data.
   *
   * @generated from field: io.clbs.openhes.pbdataproxy.BulkStatus status = 2;
   */
  status?: BulkStatus;

  /**
   * The metadata fields.
   *
   * @generated from field: io.clbs.openhes.models.MetadataFields metadata = 3;
   */
  metadata?: MetadataFields;
};

/**
 * Describes the message io.clbs.openhes.pbdataproxy.Bulk.
 * Use `create(BulkSchema)` to create a new message.
 */
export const BulkSchema: GenMessage<Bulk> = /*@__PURE__*/
  messageDesc(file_pbdataproxy_models, 4);

/**
 * DataProxy -> RestApi - the message holds list of bulks.
 *
 * @generated from message io.clbs.openhes.pbdataproxy.ListOfBulk
 */
export type ListOfBulk = Message<"io.clbs.openhes.pbdataproxy.ListOfBulk"> & {
  /**
   * The list of bulks.
   *
   * @generated from field: repeated io.clbs.openhes.pbdataproxy.Bulk items = 1;
   */
  items: Bulk[];
};

/**
 * Describes the message io.clbs.openhes.pbdataproxy.ListOfBulk.
 * Use `create(ListOfBulkSchema)` to create a new message.
 */
export const ListOfBulkSchema: GenMessage<ListOfBulk> = /*@__PURE__*/
  messageDesc(file_pbdataproxy_models, 5);

/**
 * Bulk statuses
 *
 * @generated from enum io.clbs.openhes.pbdataproxy.BulkStatusCode
 */
export enum BulkStatusCode {
  /**
   * The job is waiting in the queue
   *
   * @generated from enum value: BULK_STATUS_QUEUED = 0;
   */
  BULK_STATUS_QUEUED = 0,

  /**
   * The job is running
   *
   * @generated from enum value: BULK_STATUS_RUNNING = 1;
   */
  BULK_STATUS_RUNNING = 1,

  /**
   * The job is completed
   *
   * @generated from enum value: BULK_STATUS_COMPLETED = 2;
   */
  BULK_STATUS_COMPLETED = 2,

  /**
   * The job is cancelled
   *
   * @generated from enum value: BULK_STATUS_CANCELLED = 3;
   */
  BULK_STATUS_CANCELLED = 3,

  /**
   * The job has expired
   *
   * @generated from enum value: BULK_STATUS_EXPIRED = 4;
   */
  BULK_STATUS_EXPIRED = 4,
}

/**
 * Describes the enum io.clbs.openhes.pbdataproxy.BulkStatusCode.
 */
export const BulkStatusCodeSchema: GenEnum<BulkStatusCode> = /*@__PURE__*/
  enumDesc(file_pbdataproxy_models, 0);

