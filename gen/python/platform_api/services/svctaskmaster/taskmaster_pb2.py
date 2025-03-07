# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: services/svctaskmaster/taskmaster.proto
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
    'services/svctaskmaster/taskmaster.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2
from acquisition import main_pb2 as acquisition_dot_main__pb2
from acquisition import internal_pb2 as acquisition_dot_internal__pb2
from common import types_pb2 as common_dot_types__pb2
from system import main_pb2 as system_dot_main__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\'services/svctaskmaster/taskmaster.proto\x12&io.clbs.openhes.services.svctaskmaster\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x16\x61\x63quisition/main.proto\x1a\x1a\x61\x63quisition/internal.proto\x1a\x12\x63ommon/types.proto\x1a\x11system/main.proto2\x9e\x06\n\x11TaskmasterService\x12Y\n\tQueueJobs\x12\x34.io.clbs.openhes.models.acquisition.QueueJobsRequest\x1a\x16.google.protobuf.Empty\x12Z\n\x06GetJob\x12\x1c.google.protobuf.StringValue\x1a\x32.io.clbs.openhes.models.acquisition.GetJobResponse\x12;\n\tPurgeJobs\x12\x16.google.protobuf.Empty\x1a\x16.google.protobuf.Empty\x12M\n\nCancelJobs\x12\'.io.clbs.openhes.models.common.ListOfId\x1a\x16.google.protobuf.Empty\x12R\n\tSetDriver\x12-.io.clbs.openhes.models.acquisition.SetDriver\x1a\x16.google.protobuf.Empty\x12W\n\x08SetCache\x12\x33.io.clbs.openhes.models.acquisition.SetCacheRequest\x1a\x16.google.protobuf.Empty\x12u\n\x08GetCache\x12\x33.io.clbs.openhes.models.acquisition.GetCacheRequest\x1a\x34.io.clbs.openhes.models.acquisition.GetCacheResponse\x12P\n\tGetConfig\x12\x16.google.protobuf.Empty\x1a+.io.clbs.openhes.models.system.SystemConfig\x12P\n\tSetConfig\x12+.io.clbs.openhes.models.system.SystemConfig\x1a\x16.google.protobuf.EmptyB@Z>github.com/cybroslabs/hes-2-apis/gen/go/services/svctaskmasterb\x08\x65\x64itionsp\xe8\x07')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'services.svctaskmaster.taskmaster_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z>github.com/cybroslabs/hes-2-apis/gen/go/services/svctaskmaster'
  _globals['_TASKMASTERSERVICE']._serialized_start=236
  _globals['_TASKMASTERSERVICE']._serialized_end=1034
# @@protoc_insertion_point(module_scope)
