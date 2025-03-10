# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: acquisition/main.proto
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
    'acquisition/main.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from common import metadata_pb2 as common_dot_metadata__pb2
from common import fields_pb2 as common_dot_fields__pb2
from acquisition import shared_pb2 as acquisition_dot_shared__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x16\x61\x63quisition/main.proto\x12\"io.clbs.openhes.models.acquisition\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x15\x63ommon/metadata.proto\x1a\x13\x63ommon/fields.proto\x1a\x18\x61\x63quisition/shared.proto\"\xc1\x01\n\x1e\x43reateCommunicationUnitRequest\x12T\n\x04spec\x18\x01 \x01(\x0b\x32\x39.io.clbs.openhes.models.acquisition.CommunicationUnitSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"\x87\x01\n\x17ListOfCommunicationUnit\x12K\n\x05items\x18\x01 \x03(\x0b\x32\x35.io.clbs.openhes.models.acquisition.CommunicationUnitR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"p\n\x1d\x43reateCommunicationBusRequest\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadataJ\x04\x08\x01\x10\x02\"\x85\x01\n\x16ListOfCommunicationBus\x12J\n\x05items\x18\x01 \x03(\x0b\x32\x34.io.clbs.openhes.models.acquisition.CommunicationBusR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\x96\x01\n.AddCommunicationUnitsToCommunicationBusRequest\x12\x30\n\x14\x63ommunication_bus_id\x18\x01 \x01(\tR\x12\x63ommunicationBusId\x12\x32\n\x15\x63ommunication_unit_id\x18\x02 \x03(\tR\x13\x63ommunicationUnitId\"\x9b\x01\n3RemoveCommunicationUnitsFromCommunicationBusRequest\x12\x30\n\x14\x63ommunication_bus_id\x18\x01 \x01(\tR\x12\x63ommunicationBusId\x12\x32\n\x15\x63ommunication_unit_id\x18\x02 \x03(\tR\x13\x63ommunicationUnitId\"\xab\x01\n\x13\x43reateDeviceRequest\x12I\n\x04spec\x18\x01 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.DeviceSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"q\n\x0cListOfDevice\x12@\n\x05items\x18\x01 \x03(\x0b\x32*.io.clbs.openhes.models.acquisition.DeviceR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\xe8\x01\n\x06\x44\x65vice\x12I\n\x04spec\x18\x01 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.DeviceSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12H\n\x06status\x18\x02 \x01(\x0b\x32\x30.io.clbs.openhes.models.acquisition.DeviceStatusR\x06status\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"\xb5\x01\n\x18\x43reateDeviceGroupRequest\x12N\n\x04spec\x18\x01 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.DeviceGroupSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"{\n\x11ListOfDeviceGroup\x12\x45\n\x05items\x18\x01 \x03(\x0b\x32/.io.clbs.openhes.models.acquisition.DeviceGroupR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\x85\x02\n\x11StreamDeviceGroup\x12I\n\x04spec\x18\x01 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.DeviceGroupSpecH\x00R\x04spec\x12O\n\x06status\x18\x02 \x01(\x0b\x32\x35.io.clbs.openhes.models.acquisition.DeviceGroupStatusH\x00R\x06status\x12K\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsH\x00R\x08metadataB\x07\n\x05parts\"\xf7\x01\n\x0b\x44\x65viceGroup\x12N\n\x04spec\x18\x01 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.DeviceGroupSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12M\n\x06status\x18\x02 \x01(\x0b\x32\x35.io.clbs.openhes.models.acquisition.DeviceGroupStatusR\x06status\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"\x86\x01\n\x0f\x44\x65viceGroupSpec\x12\x1f\n\x0b\x65xternal_id\x18\x01 \x01(\tR\nexternalId\x12R\n\x0e\x64ynamic_filter\x18\x02 \x01(\x0b\x32+.io.clbs.openhes.models.common.ListSelectorR\rdynamicFilter\"\xea\x01\n\x11\x44\x65viceGroupStatus\x12\\\n\x07\x64\x65vices\x18\x04 \x03(\x0b\x32\x42.io.clbs.openhes.models.acquisition.DeviceGroupStatus.DevicesEntryR\x07\x64\x65vices\x1aw\n\x0c\x44\x65vicesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12Q\n\x05value\x18\x02 \x01(\x0b\x32;.io.clbs.openhes.models.acquisition.DeviceGroupStatusDeviceR\x05value:\x02\x38\x01\":\n\x17\x44\x65viceGroupStatusDevice\x12\x1f\n\x0b\x64river_type\x18\x01 \x01(\tR\ndriverType\"\xaf\x01\n\"SetDeviceCommunicationUnitsRequest\x12\x1b\n\tdevice_id\x18\x01 \x01(\tR\x08\x64\x65viceId\x12l\n\x13\x63ommunication_units\x18\x02 \x03(\x0b\x32;.io.clbs.openhes.models.acquisition.DeviceCommunicationUnitR\x12\x63ommunicationUnits\"\x93\x01\n\x1dListOfDeviceCommunicationUnit\x12Q\n\x05items\x18\x01 \x03(\x0b\x32;.io.clbs.openhes.models.acquisition.DeviceCommunicationUnitR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"R\n\x18\x41\x64\x64\x44\x65vicesToGroupRequest\x12\x19\n\x08group_id\x18\x01 \x01(\tR\x07groupId\x12\x1b\n\tdevice_id\x18\x02 \x03(\tR\x08\x64\x65viceId\"W\n\x1dRemoveDevicesFromGroupRequest\x12\x19\n\x08group_id\x18\x01 \x01(\tR\x07groupId\x12\x1b\n\tdevice_id\x18\x02 \x03(\tR\x08\x64\x65viceId\"\xa7\x01\n\x11\x43reateBulkRequest\x12G\n\x04spec\x18\x01 \x01(\x0b\x32,.io.clbs.openhes.models.acquisition.BulkSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"m\n\nListOfBulk\x12>\n\x05items\x18\x01 \x03(\x0b\x32(.io.clbs.openhes.models.acquisition.BulkR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\xe2\x01\n\x04\x42ulk\x12G\n\x04spec\x18\x01 \x01(\x0b\x32,.io.clbs.openhes.models.acquisition.BulkSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12\x46\n\x06status\x18\x02 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.BulkStatusR\x06status\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"\xef\x02\n\x08\x42ulkSpec\x12%\n\x0e\x63orrelation_id\x18\x01 \x01(\tR\rcorrelationId\x12Q\n\x07\x64\x65vices\x18\x02 \x01(\x0b\x32\x35.io.clbs.openhes.models.acquisition.ListOfJobDeviceIdH\x00R\x07\x64\x65vices\x12(\n\x0f\x64\x65vice_group_id\x18\x03 \x01(\tH\x00R\rdeviceGroupId\x12K\n\x08settings\x18\x04 \x01(\x0b\x32/.io.clbs.openhes.models.acquisition.JobSettingsR\x08settings\x12G\n\x07\x61\x63tions\x18\x05 \x03(\x0b\x32-.io.clbs.openhes.models.acquisition.JobActionR\x07\x61\x63tions\x12\x1f\n\x0bwebhook_url\x18\x06 \x01(\tR\nwebhookUrlB\x08\n\x06\x64\x65vice\"\x99\x01\n\nBulkStatus\x12J\n\x06status\x18\x01 \x01(\x0e\x32\x32.io.clbs.openhes.models.acquisition.BulkStatusCodeR\x06status\x12?\n\x04jobs\x18\x02 \x03(\x0b\x32+.io.clbs.openhes.models.acquisition.BulkJobR\x04jobs\"\\\n\x07\x42ulkJob\x12\x45\n\x06status\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.acquisition.JobStatusR\x06statusJ\x04\x08\x01\x10\x02J\x04\x08\x03\x10\x04\"\xb1\x01\n\x16\x43reateProxyBulkRequest\x12L\n\x04spec\x18\x01 \x01(\x0b\x32\x31.io.clbs.openhes.models.acquisition.ProxyBulkSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"\xdd\x02\n\rProxyBulkSpec\x12%\n\x0e\x63orrelation_id\x18\x01 \x01(\tR\rcorrelationId\x12\x1f\n\x0b\x64river_type\x18\x02 \x01(\tR\ndriverType\x12M\n\x07\x64\x65vices\x18\x03 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.ListOfJobDeviceR\x07\x64\x65vices\x12K\n\x08settings\x18\x04 \x01(\x0b\x32/.io.clbs.openhes.models.acquisition.JobSettingsR\x08settings\x12G\n\x07\x61\x63tions\x18\x05 \x03(\x0b\x32-.io.clbs.openhes.models.acquisition.JobActionR\x07\x61\x63tions\x12\x1f\n\x0bwebhook_url\x18\x06 \x01(\tR\nwebhookUrl\"\xb4\x01\n\x13SetModemPoolRequest\x12L\n\x04spec\x18\x01 \x01(\x0b\x32\x31.io.clbs.openhes.models.acquisition.ModemPoolSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadataJ\x04\x08\x02\x10\x03\"w\n\x0fListOfModemPool\x12\x43\n\x05items\x18\x01 \x03(\x0b\x32-.io.clbs.openhes.models.acquisition.ModemPoolR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\x0f\n\rModemPoolSpec\"X\n\x0fModemPoolStatus\x12\x45\n\x06modems\x18\x01 \x03(\x0b\x32-.io.clbs.openhes.models.acquisition.ModemInfoR\x06modems\"\xf1\x01\n\tModemPool\x12L\n\x04spec\x18\x01 \x01(\x0b\x32\x31.io.clbs.openhes.models.acquisition.ModemPoolSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12K\n\x06status\x18\x02 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.ModemPoolStatusR\x06status\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"o\n\x0fSetModemRequest\x12\x17\n\x07pool_id\x18\x01 \x01(\tR\x06poolId\x12\x43\n\x05modem\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.acquisition.ModemInfoR\x05modem\"b\n\tSetDriver\x12I\n\x04spec\x18\x01 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.DriverSpecB\x05\xaa\x01\x02\x08\x03R\x04specJ\x04\x08\x02\x10\x03J\x04\x08\x03\x10\x04\"q\n\x0cListOfDriver\x12@\n\x05items\x18\x01 \x03(\x0b\x32*.io.clbs.openhes.models.acquisition.DriverR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"_\n\x06\x44river\x12I\n\x04spec\x18\x01 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.DriverSpecB\x05\xaa\x01\x02\x08\x03R\x04specJ\x04\x08\x02\x10\x03J\x04\x08\x03\x10\x04\"\xec\x02\n\nDriverSpec\x12\x18\n\x07version\x18\x01 \x01(\tR\x07version\x12%\n\x0elistening_port\x18\x02 \x01(\rR\rlisteningPort\x12\x1f\n\x0b\x64river_type\x18\x03 \x01(\tR\ndriverType\x12.\n\x13max_concurrent_jobs\x18\x04 \x01(\x05R\x11maxConcurrentJobs\x12*\n\x11max_cascade_depth\x18\x05 \x01(\rR\x0fmaxCascadeDepth\x12*\n\x11typical_mem_usage\x18\x06 \x01(\x05R\x0ftypicalMemUsage\x12Q\n\ttemplates\x18\x07 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.DriverTemplatesR\ttemplates\x12!\n\x0c\x64isplay_name\x18\x08 \x01(\tR\x0b\x64isplayName\"\xbb\x01\n\x1b\x43reateDeviceRegisterRequest\x12Q\n\x04spec\x18\x01 \x01(\x0b\x32\x36.io.clbs.openhes.models.acquisition.DeviceRegisterSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"\x81\x01\n\x14ListOfDeviceRegister\x12H\n\x05items\x18\x01 \x03(\x0b\x32\x32.io.clbs.openhes.models.acquisition.DeviceRegisterR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\xb4\x01\n\x0e\x44\x65viceRegister\x12Q\n\x04spec\x18\x01 \x01(\x0b\x32\x36.io.clbs.openhes.models.acquisition.DeviceRegisterSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadataJ\x04\x08\x02\x10\x03\"\xe6\x02\n\x12\x44\x65viceRegisterSpec\x12&\n\x0b\x64river_type\x18\x01 \x01(\tB\x05\xaa\x01\x02\x08\x03R\ndriverType\x12V\n\x0b\x61\x63tion_type\x18\x02 \x01(\x0e\x32..io.clbs.openhes.models.acquisition.ActionTypeB\x05\xaa\x01\x02\x08\x03R\nactionType\x12\x66\n\nattributes\x18\x03 \x03(\x0b\x32\x46.io.clbs.openhes.models.acquisition.DeviceRegisterSpec.AttributesEntryR\nattributes\x1ah\n\x0f\x41ttributesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12?\n\x05value\x18\x02 \x01(\x0b\x32).io.clbs.openhes.models.common.FieldValueR\x05value:\x02\x38\x01\"\xd5\x01\n(CreateDeviceConfigurationTemplateRequest\x12^\n\x04spec\x18\x01 \x01(\x0b\x32\x43.io.clbs.openhes.models.acquisition.DeviceConfigurationTemplateSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"\x9b\x01\n!ListOfDeviceConfigurationTemplate\x12U\n\x05items\x18\x01 \x03(\x0b\x32?.io.clbs.openhes.models.acquisition.DeviceConfigurationTemplateR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\xce\x01\n\x1b\x44\x65viceConfigurationTemplate\x12^\n\x04spec\x18\x01 \x01(\x0b\x32\x43.io.clbs.openhes.models.acquisition.DeviceConfigurationTemplateSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadataJ\x04\x08\x02\x10\x03\"I\n\x1f\x44\x65viceConfigurationTemplateSpec\x12&\n\x0b\x64river_type\x18\x01 \x01(\tB\x05\xaa\x01\x02\x08\x03R\ndriverType\"s\n/AddRegisterToDeviceConfigurationTemplateRequest\x12\x1f\n\x0btemplate_id\x18\x01 \x01(\tR\ntemplateId\x12\x1f\n\x0bregister_id\x18\x02 \x01(\tR\nregisterId\"x\n4RemoveRegisterFromDeviceConfigurationTemplateRequest\x12\x1f\n\x0btemplate_id\x18\x01 \x01(\tR\ntemplateId\x12\x1f\n\x0bregister_id\x18\x02 \x01(\tR\nregisterIdB5Z3github.com/cybroslabs/hes-2-apis/gen/go/acquisitionb\x08\x65\x64itionsp\xe8\x07')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'acquisition.main_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z3github.com/cybroslabs/hes-2-apis/gen/go/acquisition'
  _globals['_CREATECOMMUNICATIONUNITREQUEST'].fields_by_name['spec']._loaded_options = None
  _globals['_CREATECOMMUNICATIONUNITREQUEST'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_CREATEDEVICEREQUEST'].fields_by_name['spec']._loaded_options = None
  _globals['_CREATEDEVICEREQUEST'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_DEVICE'].fields_by_name['spec']._loaded_options = None
  _globals['_DEVICE'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_CREATEDEVICEGROUPREQUEST'].fields_by_name['spec']._loaded_options = None
  _globals['_CREATEDEVICEGROUPREQUEST'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_DEVICEGROUP'].fields_by_name['spec']._loaded_options = None
  _globals['_DEVICEGROUP'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_DEVICEGROUPSTATUS_DEVICESENTRY']._loaded_options = None
  _globals['_DEVICEGROUPSTATUS_DEVICESENTRY']._serialized_options = b'8\001'
  _globals['_CREATEBULKREQUEST'].fields_by_name['spec']._loaded_options = None
  _globals['_CREATEBULKREQUEST'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_BULK'].fields_by_name['spec']._loaded_options = None
  _globals['_BULK'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_CREATEPROXYBULKREQUEST'].fields_by_name['spec']._loaded_options = None
  _globals['_CREATEPROXYBULKREQUEST'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_SETMODEMPOOLREQUEST'].fields_by_name['spec']._loaded_options = None
  _globals['_SETMODEMPOOLREQUEST'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_MODEMPOOL'].fields_by_name['spec']._loaded_options = None
  _globals['_MODEMPOOL'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_SETDRIVER'].fields_by_name['spec']._loaded_options = None
  _globals['_SETDRIVER'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_DRIVER'].fields_by_name['spec']._loaded_options = None
  _globals['_DRIVER'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_CREATEDEVICEREGISTERREQUEST'].fields_by_name['spec']._loaded_options = None
  _globals['_CREATEDEVICEREGISTERREQUEST'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_DEVICEREGISTER'].fields_by_name['spec']._loaded_options = None
  _globals['_DEVICEREGISTER'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_DEVICEREGISTERSPEC_ATTRIBUTESENTRY']._loaded_options = None
  _globals['_DEVICEREGISTERSPEC_ATTRIBUTESENTRY']._serialized_options = b'8\001'
  _globals['_DEVICEREGISTERSPEC'].fields_by_name['driver_type']._loaded_options = None
  _globals['_DEVICEREGISTERSPEC'].fields_by_name['driver_type']._serialized_options = b'\252\001\002\010\003'
  _globals['_DEVICEREGISTERSPEC'].fields_by_name['action_type']._loaded_options = None
  _globals['_DEVICEREGISTERSPEC'].fields_by_name['action_type']._serialized_options = b'\252\001\002\010\003'
  _globals['_CREATEDEVICECONFIGURATIONTEMPLATEREQUEST'].fields_by_name['spec']._loaded_options = None
  _globals['_CREATEDEVICECONFIGURATIONTEMPLATEREQUEST'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_DEVICECONFIGURATIONTEMPLATE'].fields_by_name['spec']._loaded_options = None
  _globals['_DEVICECONFIGURATIONTEMPLATE'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_DEVICECONFIGURATIONTEMPLATESPEC'].fields_by_name['driver_type']._loaded_options = None
  _globals['_DEVICECONFIGURATIONTEMPLATESPEC'].fields_by_name['driver_type']._serialized_options = b'\252\001\002\010\003'
  _globals['_CREATECOMMUNICATIONUNITREQUEST']._serialized_start=195
  _globals['_CREATECOMMUNICATIONUNITREQUEST']._serialized_end=388
  _globals['_LISTOFCOMMUNICATIONUNIT']._serialized_start=391
  _globals['_LISTOFCOMMUNICATIONUNIT']._serialized_end=526
  _globals['_CREATECOMMUNICATIONBUSREQUEST']._serialized_start=528
  _globals['_CREATECOMMUNICATIONBUSREQUEST']._serialized_end=640
  _globals['_LISTOFCOMMUNICATIONBUS']._serialized_start=643
  _globals['_LISTOFCOMMUNICATIONBUS']._serialized_end=776
  _globals['_ADDCOMMUNICATIONUNITSTOCOMMUNICATIONBUSREQUEST']._serialized_start=779
  _globals['_ADDCOMMUNICATIONUNITSTOCOMMUNICATIONBUSREQUEST']._serialized_end=929
  _globals['_REMOVECOMMUNICATIONUNITSFROMCOMMUNICATIONBUSREQUEST']._serialized_start=932
  _globals['_REMOVECOMMUNICATIONUNITSFROMCOMMUNICATIONBUSREQUEST']._serialized_end=1087
  _globals['_CREATEDEVICEREQUEST']._serialized_start=1090
  _globals['_CREATEDEVICEREQUEST']._serialized_end=1261
  _globals['_LISTOFDEVICE']._serialized_start=1263
  _globals['_LISTOFDEVICE']._serialized_end=1376
  _globals['_DEVICE']._serialized_start=1379
  _globals['_DEVICE']._serialized_end=1611
  _globals['_CREATEDEVICEGROUPREQUEST']._serialized_start=1614
  _globals['_CREATEDEVICEGROUPREQUEST']._serialized_end=1795
  _globals['_LISTOFDEVICEGROUP']._serialized_start=1797
  _globals['_LISTOFDEVICEGROUP']._serialized_end=1920
  _globals['_STREAMDEVICEGROUP']._serialized_start=1923
  _globals['_STREAMDEVICEGROUP']._serialized_end=2184
  _globals['_DEVICEGROUP']._serialized_start=2187
  _globals['_DEVICEGROUP']._serialized_end=2434
  _globals['_DEVICEGROUPSPEC']._serialized_start=2437
  _globals['_DEVICEGROUPSPEC']._serialized_end=2571
  _globals['_DEVICEGROUPSTATUS']._serialized_start=2574
  _globals['_DEVICEGROUPSTATUS']._serialized_end=2808
  _globals['_DEVICEGROUPSTATUS_DEVICESENTRY']._serialized_start=2689
  _globals['_DEVICEGROUPSTATUS_DEVICESENTRY']._serialized_end=2808
  _globals['_DEVICEGROUPSTATUSDEVICE']._serialized_start=2810
  _globals['_DEVICEGROUPSTATUSDEVICE']._serialized_end=2868
  _globals['_SETDEVICECOMMUNICATIONUNITSREQUEST']._serialized_start=2871
  _globals['_SETDEVICECOMMUNICATIONUNITSREQUEST']._serialized_end=3046
  _globals['_LISTOFDEVICECOMMUNICATIONUNIT']._serialized_start=3049
  _globals['_LISTOFDEVICECOMMUNICATIONUNIT']._serialized_end=3196
  _globals['_ADDDEVICESTOGROUPREQUEST']._serialized_start=3198
  _globals['_ADDDEVICESTOGROUPREQUEST']._serialized_end=3280
  _globals['_REMOVEDEVICESFROMGROUPREQUEST']._serialized_start=3282
  _globals['_REMOVEDEVICESFROMGROUPREQUEST']._serialized_end=3369
  _globals['_CREATEBULKREQUEST']._serialized_start=3372
  _globals['_CREATEBULKREQUEST']._serialized_end=3539
  _globals['_LISTOFBULK']._serialized_start=3541
  _globals['_LISTOFBULK']._serialized_end=3650
  _globals['_BULK']._serialized_start=3653
  _globals['_BULK']._serialized_end=3879
  _globals['_BULKSPEC']._serialized_start=3882
  _globals['_BULKSPEC']._serialized_end=4249
  _globals['_BULKSTATUS']._serialized_start=4252
  _globals['_BULKSTATUS']._serialized_end=4405
  _globals['_BULKJOB']._serialized_start=4407
  _globals['_BULKJOB']._serialized_end=4499
  _globals['_CREATEPROXYBULKREQUEST']._serialized_start=4502
  _globals['_CREATEPROXYBULKREQUEST']._serialized_end=4679
  _globals['_PROXYBULKSPEC']._serialized_start=4682
  _globals['_PROXYBULKSPEC']._serialized_end=5031
  _globals['_SETMODEMPOOLREQUEST']._serialized_start=5034
  _globals['_SETMODEMPOOLREQUEST']._serialized_end=5214
  _globals['_LISTOFMODEMPOOL']._serialized_start=5216
  _globals['_LISTOFMODEMPOOL']._serialized_end=5335
  _globals['_MODEMPOOLSPEC']._serialized_start=5337
  _globals['_MODEMPOOLSPEC']._serialized_end=5352
  _globals['_MODEMPOOLSTATUS']._serialized_start=5354
  _globals['_MODEMPOOLSTATUS']._serialized_end=5442
  _globals['_MODEMPOOL']._serialized_start=5445
  _globals['_MODEMPOOL']._serialized_end=5686
  _globals['_SETMODEMREQUEST']._serialized_start=5688
  _globals['_SETMODEMREQUEST']._serialized_end=5799
  _globals['_SETDRIVER']._serialized_start=5801
  _globals['_SETDRIVER']._serialized_end=5899
  _globals['_LISTOFDRIVER']._serialized_start=5901
  _globals['_LISTOFDRIVER']._serialized_end=6014
  _globals['_DRIVER']._serialized_start=6016
  _globals['_DRIVER']._serialized_end=6111
  _globals['_DRIVERSPEC']._serialized_start=6114
  _globals['_DRIVERSPEC']._serialized_end=6478
  _globals['_CREATEDEVICEREGISTERREQUEST']._serialized_start=6481
  _globals['_CREATEDEVICEREGISTERREQUEST']._serialized_end=6668
  _globals['_LISTOFDEVICEREGISTER']._serialized_start=6671
  _globals['_LISTOFDEVICEREGISTER']._serialized_end=6800
  _globals['_DEVICEREGISTER']._serialized_start=6803
  _globals['_DEVICEREGISTER']._serialized_end=6983
  _globals['_DEVICEREGISTERSPEC']._serialized_start=6986
  _globals['_DEVICEREGISTERSPEC']._serialized_end=7344
  _globals['_DEVICEREGISTERSPEC_ATTRIBUTESENTRY']._serialized_start=7240
  _globals['_DEVICEREGISTERSPEC_ATTRIBUTESENTRY']._serialized_end=7344
  _globals['_CREATEDEVICECONFIGURATIONTEMPLATEREQUEST']._serialized_start=7347
  _globals['_CREATEDEVICECONFIGURATIONTEMPLATEREQUEST']._serialized_end=7560
  _globals['_LISTOFDEVICECONFIGURATIONTEMPLATE']._serialized_start=7563
  _globals['_LISTOFDEVICECONFIGURATIONTEMPLATE']._serialized_end=7718
  _globals['_DEVICECONFIGURATIONTEMPLATE']._serialized_start=7721
  _globals['_DEVICECONFIGURATIONTEMPLATE']._serialized_end=7927
  _globals['_DEVICECONFIGURATIONTEMPLATESPEC']._serialized_start=7929
  _globals['_DEVICECONFIGURATIONTEMPLATESPEC']._serialized_end=8002
  _globals['_ADDREGISTERTODEVICECONFIGURATIONTEMPLATEREQUEST']._serialized_start=8004
  _globals['_ADDREGISTERTODEVICECONFIGURATIONTEMPLATEREQUEST']._serialized_end=8119
  _globals['_REMOVEREGISTERFROMDEVICECONFIGURATIONTEMPLATEREQUEST']._serialized_start=8121
  _globals['_REMOVEREGISTERFROMDEVICECONFIGURATIONTEMPLATEREQUEST']._serialized_end=8241
# @@protoc_insertion_point(module_scope)
