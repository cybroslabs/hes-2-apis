# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: services/svcdriveroperator/driveroperator.proto
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
    'services/svcdriveroperator/driveroperator.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2
from acquisition import internal_pb2 as acquisition_dot_internal__pb2
from acquisition import main_pb2 as acquisition_dot_main__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n/services/svcdriveroperator/driveroperator.proto\x12*io.clbs.openhes.services.svcdriveroperator\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x1a\x61\x63quisition/internal.proto\x1a\x16\x61\x63quisition/main.proto2\xc9\x04\n\x15\x44riverOperatorService\x12W\n\x0bListDrivers\x12\x16.google.protobuf.Empty\x1a\x30.io.clbs.openhes.models.acquisition.ListOfDriver\x12O\n\tSetDriver\x12*.io.clbs.openhes.models.acquisition.Driver\x1a\x16.google.protobuf.Empty\x12U\n\tGetDriver\x12\x1c.google.protobuf.StringValue\x1a*.io.clbs.openhes.models.acquisition.Driver\x12\x63\n\x0eSetDriverScale\x12\x39.io.clbs.openhes.models.acquisition.SetDriverScaleRequest\x1a\x16.google.protobuf.Empty\x12i\n\x0eGetDriverScale\x12\x39.io.clbs.openhes.models.acquisition.GetDriverScaleRequest\x1a\x1c.google.protobuf.UInt32Value\x12_\n\x0cStartUpgrade\x12\x37.io.clbs.openhes.models.acquisition.StartUpgradeRequest\x1a\x16.google.protobuf.EmptyBDZBgithub.com/cybroslabs/hes-2-apis/gen/go/services/svcdriveroperatorb\x08\x65\x64itionsp\xe8\x07')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'services.svcdriveroperator.driveroperator_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZBgithub.com/cybroslabs/hes-2-apis/gen/go/services/svcdriveroperator'
  _globals['_DRIVEROPERATORSERVICE']._serialized_start=209
  _globals['_DRIVEROPERATORSERVICE']._serialized_end=794
# @@protoc_insertion_point(module_scope)
