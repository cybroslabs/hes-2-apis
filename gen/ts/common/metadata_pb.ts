// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file common/metadata.proto (package io.clbs.openhes.models.common, edition 2023)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Value } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_struct } from "@bufbuild/protobuf/wkt";
import type { FieldValue } from "./fields_pb";
import { file_common_fields } from "./fields_pb";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file common/metadata.proto.
 */
export const file_common_metadata: GenFile = /*@__PURE__*/
  fileDesc("ChVjb21tb24vbWV0YWRhdGEucHJvdG8SHWlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uIpIDCg5NZXRhZGF0YUZpZWxkcxIKCgJpZBgBIAEoCRISCgpnZW5lcmF0aW9uGAIgASgFEkkKBmZpZWxkcxgDIAMoCzI5LmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLk1ldGFkYXRhRmllbGRzLkZpZWxkc0VudHJ5ElgKDm1hbmFnZWRfZmllbGRzGAQgAygLMkAuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5jb21tb24uTWV0YWRhdGFGaWVsZHMuTWFuYWdlZEZpZWxkc0VudHJ5EhMKBG5hbWUYBSABKAlCBaoBAggDGlgKC0ZpZWxkc0VudHJ5EgsKA2tleRgBIAEoCRI4CgV2YWx1ZRgCIAEoCzIpLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLkZpZWxkVmFsdWU6AjgBGkwKEk1hbmFnZWRGaWVsZHNFbnRyeRILCgNrZXkYASABKAkSJQoFdmFsdWUYAiABKAsyFi5nb29nbGUucHJvdG9idWYuVmFsdWU6AjgBQjBaLmdpdGh1Yi5jb20vY3licm9zbGFicy9oZXMtMi1hcGlzL2dlbi9nby9jb21tb25iCGVkaXRpb25zcOgH", [file_google_protobuf_struct, file_common_fields]);

/**
 * The metadata fields managed by user and system.
 *
 * @generated from message io.clbs.openhes.models.common.MetadataFields
 */
export type MetadataFields = Message<"io.clbs.openhes.models.common.MetadataFields"> & {
  /**
   * The UUID of the resource. It serves as the unique identifier of the resource. It's immutable and typically auto-generated during Create operations.
   *
   * @generated from field: string id = 1;
   */
  id: string;

  /**
   * The generation of the resource.
   *
   * @generated from field: int32 generation = 2;
   */
  generation: number;

  /**
   * The additional fields managed by user.
   *
   * @generated from field: map<string, io.clbs.openhes.models.common.FieldValue> fields = 3;
   */
  fields: { [key: string]: FieldValue };

  /**
   * The additional fields managed by system.
   *
   * @generated from field: map<string, google.protobuf.Value> managed_fields = 4;
   */
  managedFields: { [key: string]: Value };

  /**
   * The name of the resource. It's mutable and typically set by user. Must be set.
   *
   * @generated from field: string name = 5 [features.field_presence = LEGACY_REQUIRED];
   */
  name: string;
};

/**
 * Describes the message io.clbs.openhes.models.common.MetadataFields.
 * Use `create(MetadataFieldsSchema)` to create a new message.
 */
export const MetadataFieldsSchema: GenMessage<MetadataFields> = /*@__PURE__*/
  messageDesc(file_common_metadata, 0);

