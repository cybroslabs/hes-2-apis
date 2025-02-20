// @generated by protoc-gen-es v2.2.3 with parameter "target=ts"
// @generated from file common/fields.proto (package io.clbs.openhes.models.common, edition 2023)
/* eslint-disable */

import type { GenEnum, GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import { enumDesc, fileDesc, messageDesc } from "@bufbuild/protobuf/codegenv1";
import type { Timestamp } from "@bufbuild/protobuf/wkt";
import { file_google_protobuf_timestamp } from "@bufbuild/protobuf/wkt";
import type { Message } from "@bufbuild/protobuf";

/**
 * Describes the file common/fields.proto.
 */
export const file_common_fields: GenFile = /*@__PURE__*/
  fileDesc("ChNjb21tb24vZmllbGRzLnByb3RvEh1pby5jbGJzLm9wZW5oZXMubW9kZWxzLmNvbW1vbiLNAQoMTGlzdFNlbGVjdG9yEhEKCXBhZ2Vfc2l6ZRgBIAEoDRIOCgZvZmZzZXQYAiABKA0SQgoHc29ydF9ieRgDIAMoCzIxLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLkxpc3RTZWxlY3RvclNvcnRCeRJGCglmaWx0ZXJfYnkYBCADKAsyMy5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmNvbW1vbi5MaXN0U2VsZWN0b3JGaWx0ZXJCeRIOCgZmaWVsZHMYBSADKAkiOwoSTGlzdFNlbGVjdG9yU29ydEJ5EhAKCGZpZWxkX2lkGAEgASgJEhMKBGRlc2MYAiABKAg6BWZhbHNlIpQCChRMaXN0U2VsZWN0b3JGaWx0ZXJCeRIQCghmaWVsZF9pZBgBIAEoCRI/CghvcGVyYXRvchgCIAEoDjItLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLkZpbHRlck9wZXJhdG9yEj8KCWRhdGFfdHlwZRgDIAEoDjIsLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLkZpZWxkRGF0YVR5cGUSDAoEdGV4dBgEIAMoCRIPCgdpbnRlZ2VyGAUgAygSEg4KBm51bWJlchgGIAMoARIPCgdib29sZWFuGAcgAygIEigKBGRhdGUYCCADKAsyGi5nb29nbGUucHJvdG9idWYuVGltZXN0YW1wItsDCg9GaWVsZERlc2NyaXB0b3ISEAoIZmllbGRfaWQYASABKAkSDQoFbGFiZWwYAiABKAkSEAoIZ3JvdXBfaWQYBiABKAkSPwoJZGF0YV90eXBlGAMgASgOMiwuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5jb21tb24uRmllbGREYXRhVHlwZRJBCgZmb3JtYXQYBCABKA4yMS5pby5jbGJzLm9wZW5oZXMubW9kZWxzLmNvbW1vbi5GaWVsZERpc3BsYXlGb3JtYXQSDAoEdW5pdBgFIAEoCRIRCglwcmVjaXNpb24YCCABKAUSDwoHdG9vbHRpcBgJIAEoCRIQCghyZXF1aXJlZBgLIAEoCBIQCghlZGl0YWJsZRgMIAEoCBIPCgd2aXNpYmxlGA0gASgIEhMKC211bHRpX3ZhbHVlGA4gASgIEg8KB3NlY3VyZWQYDyABKAgSQgoKdmFsaWRhdGlvbhgQIAEoCzIuLmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLkZpZWxkVmFsaWRhdGlvbhJACg1kZWZhdWx0X3ZhbHVlGBEgASgLMikuaW8uY2xicy5vcGVuaGVzLm1vZGVscy5jb21tb24uRmllbGRWYWx1ZSKVAgoPRmllbGRWYWxpZGF0aW9uEgoKAnJlGAEgASgJEhIKCm1pbl9sZW5ndGgYAiABKAUSEgoKbWF4X2xlbmd0aBgDIAEoBRITCgttaW5faW50ZWdlchgEIAEoEhITCgttYXhfaW50ZWdlchgFIAEoEhISCgptaW5fbnVtYmVyGAYgASgBEhIKCm1heF9udW1iZXIYByABKAESTAoHb3B0aW9ucxgIIAMoCzI7LmlvLmNsYnMub3Blbmhlcy5tb2RlbHMuY29tbW9uLkZpZWxkVmFsaWRhdGlvbi5PcHRpb25zRW50cnkaLgoMT3B0aW9uc0VudHJ5EgsKA2tleRgBIAEoCRINCgV2YWx1ZRgCIAEoCToCOAEivQEKCkZpZWxkVmFsdWUSFgoMc3RyaW5nX3ZhbHVlGAEgASgJSAASFwoNaW50ZWdlcl92YWx1ZRgCIAEoA0gAEhYKDGRvdWJsZV92YWx1ZRgDIAEoAUgAEhYKDGJpbmFyeV92YWx1ZRgEIAEoDEgAEhQKCmJvb2xfdmFsdWUYBSABKAhIABIwCgpkYXRlX3ZhbHVlGAYgASgLMhouZ29vZ2xlLnByb3RvYnVmLlRpbWVzdGFtcEgAQgYKBGtpbmQiawoVTGlzdE9mRmllbGREZXNjcmlwdG9yEj0KBWl0ZW1zGAEgAygLMi4uaW8uY2xicy5vcGVuaGVzLm1vZGVscy5jb21tb24uRmllbGREZXNjcmlwdG9yEhMKC3RvdGFsX2NvdW50GAIgASgFKv0BCg5GaWx0ZXJPcGVyYXRvchIJCgVFUVVBTBAAEg0KCU5PVF9FUVVBTBABEhAKDEdSRUFURVJfVEhBThACEhkKFUdSRUFURVJfVEhBTl9PUl9FUVVBTBADEg0KCUxFU1NfVEhBThAEEhYKEkxFU1NfVEhBTl9PUl9FUVVBTBAFEgwKCENPTlRBSU5TEAYSEAoMTk9UX0NPTlRBSU5TEAcSDwoLU1RBUlRTX1dJVEgQCBINCglFTkRTX1dJVEgQCRIGCgJJThAKEgoKBk5PVF9JThALEgsKB0JFVFdFRU4QDBILCgdJU19OVUxMEA0SDwoLSVNfTk9UX05VTEwQDipaCg1GaWVsZERhdGFUeXBlEggKBFRFWFQQABILCgdJTlRFR0VSEAESCgoGRE9VQkxFEAISCwoHQk9PTEVBThADEg0KCVRJTUVTVEFNUBAEEgoKBkJJTkFSWRAFKqYBChJGaWVsZERpc3BsYXlGb3JtYXQSCwoHREVGQVVMVBAAEgwKCERVUkFUSU9OEAESDAoISU5URVJWQUwQAhIICgREQVRFEAMSDAoIVVRDX0RBVEUQBBIJCgVNT05USBAFEg0KCURBWU9GV0VFSxAGEg0KCVRJTUVPRkRBWRAHEgkKBU1PTkVZEAgSDAoIUEFTU1dPUkQQCRINCglNVUxUSUxJTkUQCkIwWi5naXRodWIuY29tL2N5YnJvc2xhYnMvaGVzLTItYXBpcy9nZW4vZ28vY29tbW9uYghlZGl0aW9uc3DoBw", [file_google_protobuf_timestamp]);

/**
 * The listing selector.
 *
 * @generated from message io.clbs.openhes.models.common.ListSelector
 */
export type ListSelector = Message<"io.clbs.openhes.models.common.ListSelector"> & {
  /**
   * The number of items per page.
   *
   * @generated from field: uint32 page_size = 1;
   */
  pageSize: number;

  /**
   * The offset of the first item to return, zero based.
   *
   * @generated from field: uint32 offset = 2;
   */
  offset: number;

  /**
   * The sorting criteria.
   *
   * @generated from field: repeated io.clbs.openhes.models.common.ListSelectorSortBy sort_by = 3;
   */
  sortBy: ListSelectorSortBy[];

  /**
   * The filtering criteria.
   *
   * @generated from field: repeated io.clbs.openhes.models.common.ListSelectorFilterBy filter_by = 4;
   */
  filterBy: ListSelectorFilterBy[];

  /**
   * FIXME: This needs to be designed properly.
   *
   * The list of additional fields to be returned.
   *
   * @generated from field: repeated string fields = 5;
   */
  fields: string[];
};

/**
 * Describes the message io.clbs.openhes.models.common.ListSelector.
 * Use `create(ListSelectorSchema)` to create a new message.
 */
export const ListSelectorSchema: GenMessage<ListSelector> = /*@__PURE__*/
  messageDesc(file_common_fields, 0);

/**
 * The sorting criteria.
 *
 * @generated from message io.clbs.openhes.models.common.ListSelectorSortBy
 */
export type ListSelectorSortBy = Message<"io.clbs.openhes.models.common.ListSelectorSortBy"> & {
  /**
   * Field id.
   *
   * @generated from field: string field_id = 1;
   */
  fieldId: string;

  /**
   * Set to true to sort in descending order.
   *
   * @generated from field: bool desc = 2 [default = false];
   */
  desc: boolean;
};

/**
 * Describes the message io.clbs.openhes.models.common.ListSelectorSortBy.
 * Use `create(ListSelectorSortBySchema)` to create a new message.
 */
export const ListSelectorSortBySchema: GenMessage<ListSelectorSortBy> = /*@__PURE__*/
  messageDesc(file_common_fields, 1);

/**
 * The filtering criteria.
 *
 * Depending on the operator, the 'text', 'integer', 'number', 'boolean' or 'date' field should be used.
 * - No value must be set for operators: 'IS_NULL', 'IS_NOT_NULL'.
 * - One value must be set for single operand operators: 'EQUAL', 'NOT_EQUAL', 'GREATER_THAN', 'GREATER_THAN_OR_EQUAL', 'LESS_THAN', 'LESS_THAN_OR_EQUAL', 'CONTAINS', 'NOT_CONTAINS', 'STARTS_WITH', 'ENDS_WITH'.
 * - Two values must be set for two operand operators: 'BETWEEN'.
 * - Any number of values can be set for generic operators: 'IN', 'NOT_IN'.
 *
 * Field type determines the data type and only related field should be used. Other fields shall not be set and will be ignored by the system.
 *
 * @generated from message io.clbs.openhes.models.common.ListSelectorFilterBy
 */
export type ListSelectorFilterBy = Message<"io.clbs.openhes.models.common.ListSelectorFilterBy"> & {
  /**
   * Field id.
   *
   * @generated from field: string field_id = 1;
   */
  fieldId: string;

  /**
   * The filter operator.
   *
   * @generated from field: io.clbs.openhes.models.common.FilterOperator operator = 2;
   */
  operator: FilterOperator;

  /**
   * The data type of the field.
   *
   * @generated from field: io.clbs.openhes.models.common.FieldDataType data_type = 3;
   */
  dataType: FieldDataType;

  /**
   * The text-typed value(s) used for filtering.
   *
   * @generated from field: repeated string text = 4;
   */
  text: string[];

  /**
   * The integer-typed value(s) used for filtering.
   *
   * @generated from field: repeated sint64 integer = 5;
   */
  integer: bigint[];

  /**
   * The number-typed value(s) used for filtering.
   *
   * @generated from field: repeated double number = 6;
   */
  number: number[];

  /**
   * The boolean-typed value(s) used for filtering.
   *
   * @generated from field: repeated bool boolean = 7;
   */
  boolean: boolean[];

  /**
   * The date-typed value(s) used for filtering.
   *
   * @generated from field: repeated google.protobuf.Timestamp date = 8;
   */
  date: Timestamp[];
};

/**
 * Describes the message io.clbs.openhes.models.common.ListSelectorFilterBy.
 * Use `create(ListSelectorFilterBySchema)` to create a new message.
 */
export const ListSelectorFilterBySchema: GenMessage<ListSelectorFilterBy> = /*@__PURE__*/
  messageDesc(file_common_fields, 2);

/**
 * The field descriptor.
 *
 * @generated from message io.clbs.openhes.models.common.FieldDescriptor
 */
export type FieldDescriptor = Message<"io.clbs.openhes.models.common.FieldDescriptor"> & {
  /**
   * Unique identifier for the field
   *
   * @generated from field: string field_id = 1;
   */
  fieldId: string;

  /**
   * Label displayed for the field
   *
   * @generated from field: string label = 2;
   */
  label: string;

  /**
   * Group (section) identifier for the field
   *
   * @generated from field: string group_id = 6;
   */
  groupId: string;

  /**
   * Data type of the field (e.g., text, double)
   *
   * @generated from field: io.clbs.openhes.models.common.FieldDataType data_type = 3;
   */
  dataType: FieldDataType;

  /**
   * Display format (e.g., 1h 30m)
   *
   * @generated from field: io.clbs.openhes.models.common.FieldDisplayFormat format = 4;
   */
  format: FieldDisplayFormat;

  /**
   * Unit to display (e.g., kWh, USD)
   *
   * @generated from field: string unit = 5;
   */
  unit: string;

  /**
   * Decimal precision for double numbers
   *
   * @generated from field: int32 precision = 8;
   */
  precision: number;

  /**
   * Tooltip or hint text
   *
   * @generated from field: string tooltip = 9;
   */
  tooltip: string;

  /**
   * Whether the field is mandatory
   *
   * @generated from field: bool required = 11;
   */
  required: boolean;

  /**
   * Whether the field is editable
   *
   * @generated from field: bool editable = 12;
   */
  editable: boolean;

  /**
   * Whether the field is visible
   *
   * @generated from field: bool visible = 13;
   */
  visible: boolean;

  /**
   * Whether the field can have multiple values
   *
   * @generated from field: bool multi_value = 14;
   */
  multiValue: boolean;

  /**
   * Whether the field shall be handled as a security fields (e.g., password, certificate input area, ...)
   *
   * @generated from field: bool secured = 15;
   */
  secured: boolean;

  /**
   * Validation rules for the field
   *
   * @generated from field: io.clbs.openhes.models.common.FieldValidation validation = 16;
   */
  validation?: FieldValidation;

  /**
   * The default value of the attribute, it does not support multi-value fields
   *
   * @generated from field: io.clbs.openhes.models.common.FieldValue default_value = 17;
   */
  defaultValue?: FieldValue;
};

/**
 * Describes the message io.clbs.openhes.models.common.FieldDescriptor.
 * Use `create(FieldDescriptorSchema)` to create a new message.
 */
export const FieldDescriptorSchema: GenMessage<FieldDescriptor> = /*@__PURE__*/
  messageDesc(file_common_fields, 3);

/**
 * Validation rules for the field.
 *
 * @generated from message io.clbs.openhes.models.common.FieldValidation
 */
export type FieldValidation = Message<"io.clbs.openhes.models.common.FieldValidation"> & {
  /**
   * Regular expression describing input format. If not set then any value of given type can be used. It can be used for string, int or double fields only.
   *
   * @generated from field: string re = 1;
   */
  re: string;

  /**
   * The minimum length. It's used for string fields only.
   *
   * @generated from field: int32 min_length = 2;
   */
  minLength: number;

  /**
   * The maximum length. It's used for string fields only.
   *
   * @generated from field: int32 max_length = 3;
   */
  maxLength: number;

  /**
   * The minimum value. It's used for integer fields only.
   *
   * @generated from field: sint64 min_integer = 4;
   */
  minInteger: bigint;

  /**
   * The maximum value. It's used for integer fields only.
   *
   * @generated from field: sint64 max_integer = 5;
   */
  maxInteger: bigint;

  /**
   * The minimum value. It's used for number fields only.
   *
   * @generated from field: double min_number = 6;
   */
  minNumber: number;

  /**
   * The maximum value. It's used for number fields only.
   *
   * @generated from field: double max_number = 7;
   */
  maxNumber: number;

  /**
   * The list of allowed values to be set (key-value pairs). The key here represents the field value to be set and the value here represents the label to be displayed.
   *
   * @generated from field: map<string, string> options = 8;
   */
  options: { [key: string]: string };
};

/**
 * Describes the message io.clbs.openhes.models.common.FieldValidation.
 * Use `create(FieldValidationSchema)` to create a new message.
 */
export const FieldValidationSchema: GenMessage<FieldValidation> = /*@__PURE__*/
  messageDesc(file_common_fields, 4);

/**
 * @generated from message io.clbs.openhes.models.common.FieldValue
 */
export type FieldValue = Message<"io.clbs.openhes.models.common.FieldValue"> & {
  /**
   * @generated from oneof io.clbs.openhes.models.common.FieldValue.kind
   */
  kind: {
    /**
     * Represents a string-typed value.
     *
     * @generated from field: string string_value = 1;
     */
    value: string;
    case: "stringValue";
  } | {
    /**
     * Represents a 64-bit integer-typed value.
     *
     * @generated from field: int64 integer_value = 2;
     */
    value: bigint;
    case: "integerValue";
  } | {
    /**
     * Represents a 64-bit double-typed value.
     *
     * @generated from field: double double_value = 3;
     */
    value: number;
    case: "doubleValue";
  } | {
    /**
     * Represents a binary-typed value.
     *
     * @generated from field: bytes binary_value = 4;
     */
    value: Uint8Array;
    case: "binaryValue";
  } | {
    /**
     * Represents a boolean-typed value.
     *
     * @generated from field: bool bool_value = 5;
     */
    value: boolean;
    case: "boolValue";
  } | {
    /**
     * Represents a date-typed value.
     *
     * @generated from field: google.protobuf.Timestamp date_value = 6;
     */
    value: Timestamp;
    case: "dateValue";
  } | { case: undefined; value?: undefined };
};

/**
 * Describes the message io.clbs.openhes.models.common.FieldValue.
 * Use `create(FieldValueSchema)` to create a new message.
 */
export const FieldValueSchema: GenMessage<FieldValue> = /*@__PURE__*/
  messageDesc(file_common_fields, 5);

/**
 * @generated from message io.clbs.openhes.models.common.ListOfFieldDescriptor
 */
export type ListOfFieldDescriptor = Message<"io.clbs.openhes.models.common.ListOfFieldDescriptor"> & {
  /**
   * The list of field descriptors.
   *
   * @generated from field: repeated io.clbs.openhes.models.common.FieldDescriptor items = 1;
   */
  items: FieldDescriptor[];

  /**
   * The total number of items.
   *
   * @generated from field: int32 total_count = 2;
   */
  totalCount: number;
};

/**
 * Describes the message io.clbs.openhes.models.common.ListOfFieldDescriptor.
 * Use `create(ListOfFieldDescriptorSchema)` to create a new message.
 */
export const ListOfFieldDescriptorSchema: GenMessage<ListOfFieldDescriptor> = /*@__PURE__*/
  messageDesc(file_common_fields, 6);

/**
 * The filter operator.
 *
 * @generated from enum io.clbs.openhes.models.common.FilterOperator
 */
export enum FilterOperator {
  /**
   * Single operand operator for text, integer, number, boolean, date fields.
   *
   * @generated from enum value: EQUAL = 0;
   */
  EQUAL = 0,

  /**
   * Single operand operator for text, integer, number, boolean, date fields.
   *
   * @generated from enum value: NOT_EQUAL = 1;
   */
  NOT_EQUAL = 1,

  /**
   * Single operand operator for integer, number, date fields.
   *
   * @generated from enum value: GREATER_THAN = 2;
   */
  GREATER_THAN = 2,

  /**
   * Single operand operator for integer, number, date fields.
   *
   * @generated from enum value: GREATER_THAN_OR_EQUAL = 3;
   */
  GREATER_THAN_OR_EQUAL = 3,

  /**
   * Single operand operator for integer, number, date fields.
   *
   * @generated from enum value: LESS_THAN = 4;
   */
  LESS_THAN = 4,

  /**
   * Single operand operator for integer, number, date fields.
   *
   * @generated from enum value: LESS_THAN_OR_EQUAL = 5;
   */
  LESS_THAN_OR_EQUAL = 5,

  /**
   * Single operand operator for text fields.
   *
   * @generated from enum value: CONTAINS = 6;
   */
  CONTAINS = 6,

  /**
   * Single operand operator for text fields.
   *
   * @generated from enum value: NOT_CONTAINS = 7;
   */
  NOT_CONTAINS = 7,

  /**
   * Single operand operator for text fields.
   *
   * @generated from enum value: STARTS_WITH = 8;
   */
  STARTS_WITH = 8,

  /**
   * Single operand operator for text fields.
   *
   * @generated from enum value: ENDS_WITH = 9;
   */
  ENDS_WITH = 9,

  /**
   * Multiple operand operator for text, integer, number, boolean fields.
   *
   * @generated from enum value: IN = 10;
   */
  IN = 10,

  /**
   * Multiple operand operator for text, integer, number, boolean fields.
   *
   * @generated from enum value: NOT_IN = 11;
   */
  NOT_IN = 11,

  /**
   * Two operand operator for integer, number, date fields.
   *
   * @generated from enum value: BETWEEN = 12;
   */
  BETWEEN = 12,

  /**
   * No operand operator. For both null and empty string.
   *
   * @generated from enum value: IS_NULL = 13;
   */
  IS_NULL = 13,

  /**
   * No operand operator. For both null and empty string.
   *
   * @generated from enum value: IS_NOT_NULL = 14;
   */
  IS_NOT_NULL = 14,
}

/**
 * Describes the enum io.clbs.openhes.models.common.FilterOperator.
 */
export const FilterOperatorSchema: GenEnum<FilterOperator> = /*@__PURE__*/
  enumDesc(file_common_fields, 0);

/**
 * Enum representing the field data type.
 *
 * @generated from enum io.clbs.openhes.models.common.FieldDataType
 */
export enum FieldDataType {
  /**
   * The text data type.
   *
   * @generated from enum value: TEXT = 0;
   */
  TEXT = 0,

  /**
   * The integer data type.
   *
   * @generated from enum value: INTEGER = 1;
   */
  INTEGER = 1,

  /**
   * The double data type.
   *
   * @generated from enum value: DOUBLE = 2;
   */
  DOUBLE = 2,

  /**
   * The boolean data type.
   *
   * @generated from enum value: BOOLEAN = 3;
   */
  BOOLEAN = 3,

  /**
   * The timestamp data type.
   *
   * @generated from enum value: TIMESTAMP = 4;
   */
  TIMESTAMP = 4,

  /**
   * The binary data type.
   *
   * @generated from enum value: BINARY = 5;
   */
  BINARY = 5,
}

/**
 * Describes the enum io.clbs.openhes.models.common.FieldDataType.
 */
export const FieldDataTypeSchema: GenEnum<FieldDataType> = /*@__PURE__*/
  enumDesc(file_common_fields, 1);

/**
 * Enum representing the field display format.
 *
 * @generated from enum io.clbs.openhes.models.common.FieldDisplayFormat
 */
export enum FieldDisplayFormat {
  /**
   * The default display format.
   *
   * @generated from enum value: DEFAULT = 0;
   */
  DEFAULT = 0,

  /**
   * The duration display format, e.g., 1h 30m. Data type must be integer or number; unit must be milliseconds.
   *
   * @generated from enum value: DURATION = 1;
   */
  DURATION = 1,

  /**
   * The interval display format.
   *
   * @generated from enum value: INTERVAL = 2;
   */
  INTERVAL = 2,

  /**
   * The date display format with local timezone. Data type must be timestamp or string.
   *
   * @generated from enum value: DATE = 3;
   */
  DATE = 3,

  /**
   * The date display format with in UTC timezone. Data type must be timestamp or string.
   *
   * @generated from enum value: UTC_DATE = 4;
   */
  UTC_DATE = 4,

  /**
   * The month display format.
   *
   * @generated from enum value: MONTH = 5;
   */
  MONTH = 5,

  /**
   * The day of week display format.
   *
   * @generated from enum value: DAYOFWEEK = 6;
   */
  DAYOFWEEK = 6,

  /**
   * The time of day display format.
   *
   * @generated from enum value: TIMEOFDAY = 7;
   */
  TIMEOFDAY = 7,

  /**
   * The money display format.
   *
   * @generated from enum value: MONEY = 8;
   */
  MONEY = 8,

  /**
   * The password display format, e.g., ********.
   *
   * @generated from enum value: PASSWORD = 9;
   */
  PASSWORD = 9,

  /**
   * The multiline-string display format.
   *
   * @generated from enum value: MULTILINE = 10;
   */
  MULTILINE = 10,
}

/**
 * Describes the enum io.clbs.openhes.models.common.FieldDisplayFormat.
 */
export const FieldDisplayFormatSchema: GenEnum<FieldDisplayFormat> = /*@__PURE__*/
  enumDesc(file_common_fields, 2);

