// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file common/metadata.proto (package io.clbs.openhes.models.common, edition 2023)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import { file_google_protobuf_struct } from "@bufbuild/protobuf/wkt";
import { file_common_fields } from "./fields_pb";
import type { JsonObject, Message } from "@bufbuild/protobuf";

/**
 * Describes the file common/metadata.proto.
 */
export const file_common_metadata: GenFile = /*@__PURE__*/
  fileDesc("ChVjb21tb24vbWV0YWRhdGEucHJvdG8SHWlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uIp8BCg5NZXRhZGF0YUZpZWxkcxIKCgJpZBgBIAEoCRISCgpnZW5lcmF0aW9uGAIgASgFEicKBmZpZWxkcxgDIAEoCzIXLmdvb2dsZS5wcm90b2J1Zi5TdHJ1Y3QSLwoObWFuYWdlZF9maWVsZHMYBCABKAsyFy5nb29nbGUucHJvdG9idWYuU3RydWN0EhMKBG5hbWUYBSABKAlCBaoBAggDQjBaLmdpdGh1Yi5jb20vY3licm9zbGFicy9oZXMtMi1hcGlzL2dlbi9nby9jb21tb25iCGVkaXRpb25zcOgH", [file_google_protobuf_struct, file_common_fields]);

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
   * @generated from field: google.protobuf.Struct fields = 3;
   */
  fields?: JsonObject;

  /**
   * The additional fields managed by system.
   *
   * @generated from field: google.protobuf.Struct managed_fields = 4;
   */
  managedFields?: JsonObject;

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

