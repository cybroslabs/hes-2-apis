# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: common/metadata.proto
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
    'common/metadata.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from common import fields_pb2 as common_dot_fields__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x15\x63ommon/metadata.proto\x12\x1dio.clbs.openhes.models.common\x1a\x1cgoogle/protobuf/struct.proto\x1a\x13\x63ommon/fields.proto\"\xe3\x03\n\x0eMetadataFields\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x1e\n\ngeneration\x18\x02 \x01(\x05R\ngeneration\x12Q\n\x06\x66ields\x18\x03 \x03(\x0b\x32\x39.io.clbs.openhes.models.common.MetadataFields.FieldsEntryR\x06\x66ields\x12g\n\x0emanaged_fields\x18\x04 \x03(\x0b\x32@.io.clbs.openhes.models.common.MetadataFields.ManagedFieldsEntryR\rmanagedFields\x12\x12\n\x04name\x18\x05 \x01(\tR\x04name\x1a\x64\n\x0b\x46ieldsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12?\n\x05value\x18\x02 \x01(\x0b\x32).io.clbs.openhes.models.common.FieldValueR\x05value:\x02\x38\x01\x1ak\n\x12ManagedFieldsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12?\n\x05value\x18\x02 \x01(\x0b\x32).io.clbs.openhes.models.common.FieldValueR\x05value:\x02\x38\x01\x42\x30Z.github.com/cybroslabs/hes-2-apis/gen/go/commonb\x08\x65\x64itionsp\xe8\x07')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'common.metadata_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z.github.com/cybroslabs/hes-2-apis/gen/go/common'
  _globals['_METADATAFIELDS_FIELDSENTRY']._loaded_options = None
  _globals['_METADATAFIELDS_FIELDSENTRY']._serialized_options = b'8\001'
  _globals['_METADATAFIELDS_MANAGEDFIELDSENTRY']._loaded_options = None
  _globals['_METADATAFIELDS_MANAGEDFIELDSENTRY']._serialized_options = b'8\001'
  _globals['_METADATAFIELDS']._serialized_start=108
  _globals['_METADATAFIELDS']._serialized_end=591
  _globals['_METADATAFIELDS_FIELDSENTRY']._serialized_start=382
  _globals['_METADATAFIELDS_FIELDSENTRY']._serialized_end=482
  _globals['_METADATAFIELDS_MANAGEDFIELDSENTRY']._serialized_start=484
  _globals['_METADATAFIELDS_MANAGEDFIELDSENTRY']._serialized_end=591
# @@protoc_insertion_point(module_scope)
