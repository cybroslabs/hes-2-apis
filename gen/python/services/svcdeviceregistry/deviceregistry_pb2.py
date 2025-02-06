# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: services/svcdeviceregistry/deviceregistry.proto
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
    'services/svcdeviceregistry/deviceregistry.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2
from acquisition import main_pb2 as acquisition_dot_main__pb2
from acquisition import internal_pb2 as acquisition_dot_internal__pb2
from acquisition import shared_pb2 as acquisition_dot_shared__pb2
from common import fields_pb2 as common_dot_fields__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n/services/svcdeviceregistry/deviceregistry.proto\x12*io.clbs.openhes.services.svcdeviceregistry\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x16\x61\x63quisition/main.proto\x1a\x1a\x61\x63quisition/internal.proto\x1a\x18\x61\x63quisition/shared.proto\x1a\x13\x63ommon/fields.proto2\x8f\x18\n\x15\x44\x65viceRegistryService\x12[\n\x12SetDriverTemplates\x12-.io.clbs.openhes.models.acquisition.SetDriver\x1a\x16.google.protobuf.Empty\x12{\n\x17\x43reateCommunicationUnit\x12\x42.io.clbs.openhes.models.acquisition.CreateCommunicationUnitRequest\x1a\x1c.google.protobuf.StringValue\x12\x82\x01\n\x16ListCommunicationUnits\x12+.io.clbs.openhes.models.common.ListSelector\x1a;.io.clbs.openhes.models.acquisition.ListOfCommunicationUnit\x12k\n\x14GetCommunicationUnit\x12\x1c.google.protobuf.StringValue\x1a\x35.io.clbs.openhes.models.acquisition.CommunicationUnit\x12\x81\x01\n\x1a\x43reateCommunicationUnitBus\x12\x45.io.clbs.openhes.models.acquisition.CreateCommunicationUnitBusRequest\x1a\x1c.google.protobuf.StringValue\x12\x89\x01\n\x1aListCommunicationUnitBuses\x12+.io.clbs.openhes.models.common.ListSelector\x1a>.io.clbs.openhes.models.acquisition.ListOfCommunicationUnitBus\x12{\n\x1a\x41\x64\x64\x43ommunicationUnitsToBus\x12\x45.io.clbs.openhes.models.acquisition.AddCommunicationUnitsToBusRequest\x1a\x16.google.protobuf.Empty\x12\x85\x01\n\x1fRemoveCommunicationUnitsFromBus\x12J.io.clbs.openhes.models.acquisition.RemoveCommunicationUnitsFromBusRequest\x1a\x16.google.protobuf.Empty\x12\x65\n\x0c\x43reateDevice\x12\x37.io.clbs.openhes.models.acquisition.CreateDeviceRequest\x1a\x1c.google.protobuf.StringValue\x12l\n\x0bListDevices\x12+.io.clbs.openhes.models.common.ListSelector\x1a\x30.io.clbs.openhes.models.acquisition.ListOfDevice\x12U\n\tGetDevice\x12\x1c.google.protobuf.StringValue\x1a*.io.clbs.openhes.models.acquisition.Device\x12}\n\x1bSetDeviceCommunicationUnits\x12\x46.io.clbs.openhes.models.acquisition.SetDeviceCommunicationUnitsRequest\x1a\x16.google.protobuf.Empty\x12~\n\x1bGetDeviceCommunicationUnits\x12\x1c.google.protobuf.StringValue\x1a\x41.io.clbs.openhes.models.acquisition.ListOfDeviceCommunicationUnit\x12r\n\x17GetDeviceConnectionInfo\x12\x1a.google.protobuf.ListValue\x1a;.io.clbs.openhes.models.acquisition.MapDeviceConnectionInfo\x12\x61\n\rSetDeviceInfo\x12\x38.io.clbs.openhes.models.acquisition.SetDeviceInfoRequest\x1a\x16.google.protobuf.Empty\x12o\n\x11\x43reateDeviceGroup\x12<.io.clbs.openhes.models.acquisition.CreateDeviceGroupRequest\x1a\x1c.google.protobuf.StringValue\x12v\n\x10ListDeviceGroups\x12+.io.clbs.openhes.models.common.ListSelector\x1a\x35.io.clbs.openhes.models.acquisition.ListOfDeviceGroup\x12_\n\x0eGetDeviceGroup\x12\x1c.google.protobuf.StringValue\x1a/.io.clbs.openhes.models.acquisition.DeviceGroup\x12i\n\x11\x41\x64\x64\x44\x65vicesToGroup\x12<.io.clbs.openhes.models.acquisition.AddDevicesToGroupRequest\x1a\x16.google.protobuf.Empty\x12s\n\x16RemoveDevicesFromGroup\x12\x41.io.clbs.openhes.models.acquisition.RemoveDevicesFromGroupRequest\x1a\x16.google.protobuf.Empty\x12r\n\x0eListModemPools\x12+.io.clbs.openhes.models.common.ListSelector\x1a\x33.io.clbs.openhes.models.acquisition.ListOfModemPool\x12[\n\x0cGetModemPool\x12\x1c.google.protobuf.StringValue\x1a-.io.clbs.openhes.models.acquisition.ModemPool\x12h\n\x0f\x43reateModemPool\x12\x37.io.clbs.openhes.models.acquisition.SetModemPoolRequest\x1a\x1c.google.protobuf.StringValue\x12\x62\n\x0fUpdateModemPool\x12\x37.io.clbs.openhes.models.acquisition.SetModemPoolRequest\x1a\x16.google.protobuf.Empty\x12G\n\x0f\x44\x65leteModemPool\x12\x1c.google.protobuf.StringValue\x1a\x16.google.protobuf.Empty\x12`\n\x0b\x43reateModem\x12\x33.io.clbs.openhes.models.acquisition.SetModemRequest\x1a\x1c.google.protobuf.StringValue\x12Z\n\x0bUpdateModem\x12\x33.io.clbs.openhes.models.acquisition.SetModemRequest\x1a\x16.google.protobuf.Empty\x12\x43\n\x0b\x44\x65leteModem\x12\x1c.google.protobuf.StringValue\x1a\x16.google.protobuf.EmptyBDZBgithub.com/cybroslabs/hes-2-apis/gen/go/services/svcdeviceregistryb\x08\x65\x64itionsp\xe8\x07')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'services.svcdeviceregistry.deviceregistry_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'ZBgithub.com/cybroslabs/hes-2-apis/gen/go/services/svcdeviceregistry'
  _globals['_DEVICEREGISTRYSERVICE']._serialized_start=286
  _globals['_DEVICEREGISTRYSERVICE']._serialized_end=3373
# @@protoc_insertion_point(module_scope)
