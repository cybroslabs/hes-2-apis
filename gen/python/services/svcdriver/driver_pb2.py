# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: services/svcdriver/driver.proto
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
    'services/svcdriver/driver.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from acquisition import internal_pb2 as acquisition_dot_internal__pb2
from acquisition import main_pb2 as acquisition_dot_main__pb2
from common import types_pb2 as common_dot_types__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1fservices/svcdriver/driver.proto\x12\"io.clbs.openhes.services.svcdriver\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1a\x61\x63quisition/internal.proto\x1a\x16\x61\x63quisition/main.proto\x1a\x12\x63ommon/types.proto2\xd5\x01\n\rDriverService\x12v\n\x08StartJob\x12\x34.io.clbs.openhes.models.acquisition.StartJobsRequest\x1a\x32.io.clbs.openhes.models.acquisition.ProgressUpdate0\x01\x12L\n\tCancelJob\x12\'.io.clbs.openhes.models.common.ListOfId\x1a\x16.google.protobuf.EmptyB<Z:github.com/cybroslabs/hes-2-apis/gen/go/services/svcdriverb\x08\x65\x64itionsp\xe8\x07')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'services.svcdriver.driver_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z:github.com/cybroslabs/hes-2-apis/gen/go/services/svcdriver'
  _globals['_DRIVERSERVICE']._serialized_start=173
  _globals['_DRIVERSERVICE']._serialized_end=386
# @@protoc_insertion_point(module_scope)
