# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: services/svcapi/api.proto
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
    'services/svcapi/api.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2
from common import fields_pb2 as common_dot_fields__pb2
from common import types_pb2 as common_dot_types__pb2
from acquisition import main_pb2 as acquisition_dot_main__pb2
from acquisition import shared_pb2 as acquisition_dot_shared__pb2
from system import main_pb2 as system_dot_main__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x19services/svcapi/api.proto\x12\x1fio.clbs.openhes.services.svcapi\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x13\x63ommon/fields.proto\x1a\x12\x63ommon/types.proto\x1a\x16\x61\x63quisition/main.proto\x1a\x18\x61\x63quisition/shared.proto\x1a\x11system/main.proto2\xfa+\n\nApiService\x12\x8f\x01\n!CreateDeviceConfigurationRegister\x12L.io.clbs.openhes.models.acquisition.CreateDeviceConfigurationRegisterRequest\x1a\x1c.google.protobuf.StringValue\x12\x96\x01\n ListDeviceConfigurationRegisters\x12+.io.clbs.openhes.models.common.ListSelector\x1a\x45.io.clbs.openhes.models.acquisition.ListOfDeviceConfigurationRegister\x12\x7f\n\x1eGetDeviceConfigurationRegister\x12\x1c.google.protobuf.StringValue\x1a?.io.clbs.openhes.models.acquisition.DeviceConfigurationRegister\x12|\n!UpdateDeviceConfigurationRegister\x12?.io.clbs.openhes.models.acquisition.DeviceConfigurationRegister\x1a\x16.google.protobuf.Empty\x12Y\n!DeleteDeviceConfigurationRegister\x12\x1c.google.protobuf.StringValue\x1a\x16.google.protobuf.Empty\x12\x8f\x01\n!CreateDeviceConfigurationTemplate\x12L.io.clbs.openhes.models.acquisition.CreateDeviceConfigurationTemplateRequest\x1a\x1c.google.protobuf.StringValue\x12\x96\x01\n ListDeviceConfigurationTemplates\x12+.io.clbs.openhes.models.common.ListSelector\x1a\x45.io.clbs.openhes.models.acquisition.ListOfDeviceConfigurationTemplate\x12\x7f\n\x1eGetDeviceConfigurationTemplate\x12\x1c.google.protobuf.StringValue\x1a?.io.clbs.openhes.models.acquisition.DeviceConfigurationTemplate\x12|\n!UpdateDeviceConfigurationTemplate\x12?.io.clbs.openhes.models.acquisition.DeviceConfigurationTemplate\x1a\x16.google.protobuf.Empty\x12Y\n!DeleteDeviceConfigurationTemplate\x12\x1c.google.protobuf.StringValue\x1a\x16.google.protobuf.Empty\x12\xbd\x01\n;AddDeviceConfigurationRegisterToDeviceConfigurationTemplate\x12\x66.io.clbs.openhes.models.acquisition.AddDeviceConfigurationRegisterToDeviceConfigurationTemplateRequest\x1a\x16.google.protobuf.Empty\x12\xc7\x01\n@RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplate\x12k.io.clbs.openhes.models.acquisition.RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplateRequest\x1a\x16.google.protobuf.Empty\x12\x64\n\x14ListFieldDescriptors\x12\x16.google.protobuf.Empty\x1a\x34.io.clbs.openhes.models.common.ListOfFieldDescriptor\x12k\n\x0f\x43reateProxyBulk\x12:.io.clbs.openhes.models.acquisition.CreateProxyBulkRequest\x1a\x1c.google.protobuf.StringValue\x12\x61\n\nCreateBulk\x12\x35.io.clbs.openhes.models.acquisition.CreateBulkRequest\x1a\x1c.google.protobuf.StringValue\x12h\n\tListBulks\x12+.io.clbs.openhes.models.common.ListSelector\x1a..io.clbs.openhes.models.acquisition.ListOfBulk\x12Q\n\x07GetBulk\x12\x1c.google.protobuf.StringValue\x1a(.io.clbs.openhes.models.acquisition.Bulk\x12\x42\n\nCancelBulk\x12\x1c.google.protobuf.StringValue\x1a\x16.google.protobuf.Empty\x12W\n\nGetBulkJob\x12\x1c.google.protobuf.StringValue\x1a+.io.clbs.openhes.models.acquisition.BulkJob\x12l\n\x0bListDrivers\x12+.io.clbs.openhes.models.common.ListSelector\x1a\x30.io.clbs.openhes.models.acquisition.ListOfDriver\x12U\n\tGetDriver\x12\x1c.google.protobuf.StringValue\x1a*.io.clbs.openhes.models.acquisition.Driver\x12{\n\x17\x43reateCommunicationUnit\x12\x42.io.clbs.openhes.models.acquisition.CreateCommunicationUnitRequest\x1a\x1c.google.protobuf.StringValue\x12\x82\x01\n\x16ListCommunicationUnits\x12+.io.clbs.openhes.models.common.ListSelector\x1a;.io.clbs.openhes.models.acquisition.ListOfCommunicationUnit\x12k\n\x14GetCommunicationUnit\x12\x1c.google.protobuf.StringValue\x1a\x35.io.clbs.openhes.models.acquisition.CommunicationUnit\x12y\n\x16\x43reateCommunicationBus\x12\x41.io.clbs.openhes.models.acquisition.CreateCommunicationBusRequest\x1a\x1c.google.protobuf.StringValue\x12\x81\x01\n\x16ListCommunicationBuses\x12+.io.clbs.openhes.models.common.ListSelector\x1a:.io.clbs.openhes.models.acquisition.ListOfCommunicationBus\x12\x95\x01\n\'AddCommunicationUnitsToCommunicationBus\x12R.io.clbs.openhes.models.acquisition.AddCommunicationUnitsToCommunicationBusRequest\x1a\x16.google.protobuf.Empty\x12\x9f\x01\n,RemoveCommunicationUnitsFromCommunicationBus\x12W.io.clbs.openhes.models.acquisition.RemoveCommunicationUnitsFromCommunicationBusRequest\x1a\x16.google.protobuf.Empty\x12\x65\n\x0c\x43reateDevice\x12\x37.io.clbs.openhes.models.acquisition.CreateDeviceRequest\x1a\x1c.google.protobuf.StringValue\x12l\n\x0bListDevices\x12+.io.clbs.openhes.models.common.ListSelector\x1a\x30.io.clbs.openhes.models.acquisition.ListOfDevice\x12U\n\tGetDevice\x12\x1c.google.protobuf.StringValue\x1a*.io.clbs.openhes.models.acquisition.Device\x12}\n\x1bSetDeviceCommunicationUnits\x12\x46.io.clbs.openhes.models.acquisition.SetDeviceCommunicationUnitsRequest\x1a\x16.google.protobuf.Empty\x12~\n\x1bGetDeviceCommunicationUnits\x12\x1c.google.protobuf.StringValue\x1a\x41.io.clbs.openhes.models.acquisition.ListOfDeviceCommunicationUnit\x12o\n\x11\x43reateDeviceGroup\x12<.io.clbs.openhes.models.acquisition.CreateDeviceGroupRequest\x1a\x1c.google.protobuf.StringValue\x12v\n\x10ListDeviceGroups\x12+.io.clbs.openhes.models.common.ListSelector\x1a\x35.io.clbs.openhes.models.acquisition.ListOfDeviceGroup\x12_\n\x0eGetDeviceGroup\x12\x1c.google.protobuf.StringValue\x1a/.io.clbs.openhes.models.acquisition.DeviceGroup\x12i\n\x11\x41\x64\x64\x44\x65vicesToGroup\x12<.io.clbs.openhes.models.acquisition.AddDevicesToGroupRequest\x1a\x16.google.protobuf.Empty\x12s\n\x16RemoveDevicesFromGroup\x12\x41.io.clbs.openhes.models.acquisition.RemoveDevicesFromGroupRequest\x1a\x16.google.protobuf.Empty\x12w\n\x16ListDeviceGroupDevices\x12+.io.clbs.openhes.models.common.ListSelector\x1a\x30.io.clbs.openhes.models.acquisition.ListOfDevice\x12r\n\x0eListModemPools\x12+.io.clbs.openhes.models.common.ListSelector\x1a\x33.io.clbs.openhes.models.acquisition.ListOfModemPool\x12[\n\x0cGetModemPool\x12\x1c.google.protobuf.StringValue\x1a-.io.clbs.openhes.models.acquisition.ModemPool\x12h\n\x0f\x43reateModemPool\x12\x37.io.clbs.openhes.models.acquisition.SetModemPoolRequest\x1a\x1c.google.protobuf.StringValue\x12\x62\n\x0fUpdateModemPool\x12\x37.io.clbs.openhes.models.acquisition.SetModemPoolRequest\x1a\x16.google.protobuf.Empty\x12G\n\x0f\x44\x65leteModemPool\x12\x1c.google.protobuf.StringValue\x1a\x16.google.protobuf.Empty\x12`\n\x0b\x43reateModem\x12\x33.io.clbs.openhes.models.acquisition.SetModemRequest\x1a\x1c.google.protobuf.StringValue\x12Z\n\x0bUpdateModem\x12\x33.io.clbs.openhes.models.acquisition.SetModemRequest\x1a\x16.google.protobuf.Empty\x12\x43\n\x0b\x44\x65leteModem\x12\x1c.google.protobuf.StringValue\x1a\x16.google.protobuf.Empty\x12P\n\tGetConfig\x12\x16.google.protobuf.Empty\x1a+.io.clbs.openhes.models.system.SystemConfig\x12P\n\tSetConfig\x12+.io.clbs.openhes.models.system.SystemConfig\x1a\x16.google.protobuf.EmptyB9Z7github.com/cybroslabs/hes-2-apis/gen/go/services/svcapib\x08\x65\x64itionsp\xe8\x07')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'services.svcapi.api_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z7github.com/cybroslabs/hes-2-apis/gen/go/services/svcapi'
  _globals['_APISERVICE']._serialized_start=234
  _globals['_APISERVICE']._serialized_end=5860
# @@protoc_insertion_point(module_scope)
