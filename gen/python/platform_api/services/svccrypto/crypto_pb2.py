# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: services/svccrypto/crypto.proto
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
    'services/svccrypto/crypto.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from crypto import crypto_pb2 as crypto_dot_crypto__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1fservices/svccrypto/crypto.proto\x12\"io.clbs.openhes.services.svccrypto\x1a\x13\x63rypto/crypto.proto2j\n\rCryproService\x12Y\n\x04\x44lms\x12%.io.clbs.openhes.models.crypto.DlmsIn\x1a&.io.clbs.openhes.models.crypto.DlmsOut(\x01\x30\x01\x42<Z:github.com/cybroslabs/hes-2-apis/gen/go/services/svccryptob\x08\x65\x64itionsp\xe8\x07')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'services.svccrypto.crypto_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z:github.com/cybroslabs/hes-2-apis/gen/go/services/svccrypto'
  _globals['_CRYPROSERVICE']._serialized_start=92
  _globals['_CRYPROSERVICE']._serialized_end=198
# @@protoc_insertion_point(module_scope)
