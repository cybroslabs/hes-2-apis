# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: common/fields.proto
# Protobuf Python Version: 5.29.3
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    29,
    3,
    '',
    'common/fields.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x13\x63ommon/fields.proto\x12\x1dio.clbs.openhes.models.common\x1a\x1fgoogle/protobuf/timestamp.proto\"\xf9\x01\n\x0cListSelector\x12\x1b\n\tpage_size\x18\x01 \x01(\rR\x08pageSize\x12\x16\n\x06offset\x18\x02 \x01(\rR\x06offset\x12J\n\x07sort_by\x18\x03 \x03(\x0b\x32\x31.io.clbs.openhes.models.common.ListSelectorSortByR\x06sortBy\x12P\n\tfilter_by\x18\x04 \x03(\x0b\x32\x33.io.clbs.openhes.models.common.ListSelectorFilterByR\x08\x66ilterBy\x12\x16\n\x06\x66ields\x18\x05 \x03(\tR\x06\x66ields\"J\n\x12ListSelectorSortBy\x12\x19\n\x08\x66ield_id\x18\x01 \x01(\tR\x07\x66ieldId\x12\x19\n\x04\x64\x65sc\x18\x02 \x01(\x08:\x05\x66\x61lseR\x04\x64\x65sc\"\x8c\x02\n\x14ListSelectorFilterBy\x12\x19\n\x08\x66ield_id\x18\x01 \x01(\tR\x07\x66ieldId\x12I\n\x08operator\x18\x02 \x01(\x0e\x32-.io.clbs.openhes.models.common.FilterOperatorR\x08operator\x12\x12\n\x04text\x18\x03 \x03(\tR\x04text\x12\x18\n\x07integer\x18\x04 \x03(\x12R\x07integer\x12\x16\n\x06number\x18\x05 \x03(\x01R\x06number\x12\x18\n\x07\x62oolean\x18\x06 \x03(\x08R\x07\x62oolean\x12.\n\x04\x64\x61te\x18\x07 \x03(\x0b\x32\x1a.google.protobuf.TimestampR\x04\x64\x61te\"X\n\x0eListDescriptor\x12\x46\n\x06\x66ields\x18\x01 \x03(\x0b\x32..io.clbs.openhes.models.common.FieldDescriptorR\x06\x66ields\"\xec\x04\n\x0f\x46ieldDescriptor\x12\x19\n\x08\x66ield_id\x18\x01 \x01(\tR\x07\x66ieldId\x12\x14\n\x05label\x18\x02 \x01(\tR\x05label\x12\x19\n\x08group_id\x18\x06 \x01(\tR\x07groupId\x12I\n\tdata_type\x18\x03 \x01(\x0e\x32,.io.clbs.openhes.models.common.FieldDataTypeR\x08\x64\x61taType\x12I\n\x06\x66ormat\x18\x04 \x01(\x0e\x32\x31.io.clbs.openhes.models.common.FieldDisplayFormatR\x06\x66ormat\x12\x12\n\x04unit\x18\x05 \x01(\tR\x04unit\x12\x1c\n\tprecision\x18\x08 \x01(\x05R\tprecision\x12\x18\n\x07tooltip\x18\t \x01(\tR\x07tooltip\x12\x1a\n\x08required\x18\x0b \x01(\x08R\x08required\x12\x1a\n\x08\x65\x64itable\x18\x0c \x01(\x08R\x08\x65\x64itable\x12\x18\n\x07visible\x18\r \x01(\x08R\x07visible\x12\x1f\n\x0bmulti_value\x18\x0e \x01(\x08R\nmultiValue\x12\x18\n\x07secured\x18\x0f \x01(\x08R\x07secured\x12N\n\nvalidation\x18\x10 \x01(\x0b\x32..io.clbs.openhes.models.common.FieldValidationR\nvalidation\x12N\n\rdefault_value\x18\x11 \x01(\x0b\x32).io.clbs.openhes.models.common.FieldValueR\x0c\x64\x65\x66\x61ultValue\"\xf2\x02\n\x0f\x46ieldValidation\x12\x0e\n\x02re\x18\x01 \x01(\tR\x02re\x12\x1d\n\nmin_length\x18\x02 \x01(\x05R\tminLength\x12\x1d\n\nmax_length\x18\x03 \x01(\x05R\tmaxLength\x12\x1f\n\x0bmin_integer\x18\x04 \x01(\x12R\nminInteger\x12\x1f\n\x0bmax_integer\x18\x05 \x01(\x12R\nmaxInteger\x12\x1d\n\nmin_number\x18\x06 \x01(\x01R\tminNumber\x12\x1d\n\nmax_number\x18\x07 \x01(\x01R\tmaxNumber\x12U\n\x07options\x18\x08 \x03(\x0b\x32;.io.clbs.openhes.models.common.FieldValidation.OptionsEntryR\x07options\x1a:\n\x0cOptionsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\"\x88\x02\n\nFieldValue\x12#\n\x0cstring_value\x18\x01 \x01(\tH\x00R\x0bstringValue\x12%\n\rinteger_value\x18\x02 \x01(\x03H\x00R\x0cintegerValue\x12#\n\x0c\x64ouble_value\x18\x03 \x01(\x01H\x00R\x0b\x64oubleValue\x12#\n\x0c\x62inary_value\x18\x04 \x01(\x0cH\x00R\x0b\x62inaryValue\x12\x1f\n\nbool_value\x18\x05 \x01(\x08H\x00R\tboolValue\x12;\n\ndate_value\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampH\x00R\tdateValueB\x06\n\x04kind*\xfd\x01\n\x0e\x46ilterOperator\x12\t\n\x05\x45QUAL\x10\x00\x12\r\n\tNOT_EQUAL\x10\x01\x12\x10\n\x0cGREATER_THAN\x10\x02\x12\x19\n\x15GREATER_THAN_OR_EQUAL\x10\x03\x12\r\n\tLESS_THAN\x10\x04\x12\x16\n\x12LESS_THAN_OR_EQUAL\x10\x05\x12\x0c\n\x08\x43ONTAINS\x10\x06\x12\x10\n\x0cNOT_CONTAINS\x10\x07\x12\x0f\n\x0bSTARTS_WITH\x10\x08\x12\r\n\tENDS_WITH\x10\t\x12\x06\n\x02IN\x10\n\x12\n\n\x06NOT_IN\x10\x0b\x12\x0b\n\x07\x42\x45TWEEN\x10\x0c\x12\x0b\n\x07IS_NULL\x10\r\x12\x0f\n\x0bIS_NOT_NULL\x10\x0e*Z\n\rFieldDataType\x12\x08\n\x04TEXT\x10\x00\x12\x0b\n\x07INTEGER\x10\x01\x12\n\n\x06\x44OUBLE\x10\x02\x12\x0b\n\x07\x42OOLEAN\x10\x03\x12\r\n\tTIMESTAMP\x10\x04\x12\n\n\x06\x42INARY\x10\x05*\xa6\x01\n\x12\x46ieldDisplayFormat\x12\x0b\n\x07\x44\x45\x46\x41ULT\x10\x00\x12\x0c\n\x08\x44URATION\x10\x01\x12\x0c\n\x08INTERVAL\x10\x02\x12\x08\n\x04\x44\x41TE\x10\x03\x12\x0c\n\x08UTC_DATE\x10\x04\x12\t\n\x05MONTH\x10\x05\x12\r\n\tDAYOFWEEK\x10\x06\x12\r\n\tTIMEOFDAY\x10\x07\x12\t\n\x05MONEY\x10\x08\x12\x0c\n\x08PASSWORD\x10\t\x12\r\n\tMULTILINE\x10\nB0Z.github.com/cybroslabs/hes-2-apis/gen/go/commonb\x08\x65\x64itionsp\xe8\x07')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'common.fields_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z.github.com/cybroslabs/hes-2-apis/gen/go/common'
  _globals['_FIELDVALIDATION_OPTIONSENTRY']._loaded_options = None
  _globals['_FIELDVALIDATION_OPTIONSENTRY']._serialized_options = b'8\001'
  _globals['_FILTEROPERATOR']._serialized_start=2040
  _globals['_FILTEROPERATOR']._serialized_end=2293
  _globals['_FIELDDATATYPE']._serialized_start=2295
  _globals['_FIELDDATATYPE']._serialized_end=2385
  _globals['_FIELDDISPLAYFORMAT']._serialized_start=2388
  _globals['_FIELDDISPLAYFORMAT']._serialized_end=2554
  _globals['_LISTSELECTOR']._serialized_start=88
  _globals['_LISTSELECTOR']._serialized_end=337
  _globals['_LISTSELECTORSORTBY']._serialized_start=339
  _globals['_LISTSELECTORSORTBY']._serialized_end=413
  _globals['_LISTSELECTORFILTERBY']._serialized_start=416
  _globals['_LISTSELECTORFILTERBY']._serialized_end=684
  _globals['_LISTDESCRIPTOR']._serialized_start=686
  _globals['_LISTDESCRIPTOR']._serialized_end=774
  _globals['_FIELDDESCRIPTOR']._serialized_start=777
  _globals['_FIELDDESCRIPTOR']._serialized_end=1397
  _globals['_FIELDVALIDATION']._serialized_start=1400
  _globals['_FIELDVALIDATION']._serialized_end=1770
  _globals['_FIELDVALIDATION_OPTIONSENTRY']._serialized_start=1712
  _globals['_FIELDVALIDATION_OPTIONSENTRY']._serialized_end=1770
  _globals['_FIELDVALUE']._serialized_start=1773
  _globals['_FIELDVALUE']._serialized_end=2037
# @@protoc_insertion_point(module_scope)
