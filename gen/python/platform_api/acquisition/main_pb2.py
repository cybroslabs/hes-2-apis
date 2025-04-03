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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x16\x61\x63quisition/main.proto\x12\"io.clbs.openhes.models.acquisition\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x15\x63ommon/metadata.proto\x1a\x13\x63ommon/fields.proto\x1a\x18\x61\x63quisition/shared.proto\"\xc1\x01\n\x1e\x43reateCommunicationUnitRequest\x12T\n\x04spec\x18\x01 \x01(\x0b\x32\x39.io.clbs.openhes.models.acquisition.CommunicationUnitSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"\x87\x01\n\x17ListOfCommunicationUnit\x12K\n\x05items\x18\x01 \x03(\x0b\x32\x35.io.clbs.openhes.models.acquisition.CommunicationUnitR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"p\n\x1d\x43reateCommunicationBusRequest\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadataJ\x04\x08\x01\x10\x02\"\x85\x01\n\x16ListOfCommunicationBus\x12J\n\x05items\x18\x01 \x03(\x0b\x32\x34.io.clbs.openhes.models.acquisition.CommunicationBusR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\x96\x01\n.AddCommunicationUnitsToCommunicationBusRequest\x12\x30\n\x14\x63ommunication_bus_id\x18\x01 \x01(\tR\x12\x63ommunicationBusId\x12\x32\n\x15\x63ommunication_unit_id\x18\x02 \x03(\tR\x13\x63ommunicationUnitId\"\x9b\x01\n3RemoveCommunicationUnitsFromCommunicationBusRequest\x12\x30\n\x14\x63ommunication_bus_id\x18\x01 \x01(\tR\x12\x63ommunicationBusId\x12\x32\n\x15\x63ommunication_unit_id\x18\x02 \x03(\tR\x13\x63ommunicationUnitId\"\xab\x01\n\x13\x43reateDeviceRequest\x12I\n\x04spec\x18\x01 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.DeviceSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"q\n\x0cListOfDevice\x12@\n\x05items\x18\x01 \x03(\x0b\x32*.io.clbs.openhes.models.acquisition.DeviceR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\xe8\x01\n\x06\x44\x65vice\x12I\n\x04spec\x18\x01 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.DeviceSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12H\n\x06status\x18\x02 \x01(\x0b\x32\x30.io.clbs.openhes.models.acquisition.DeviceStatusR\x06status\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"\xb5\x01\n\x18\x43reateDeviceGroupRequest\x12N\n\x04spec\x18\x01 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.DeviceGroupSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"{\n\x11ListOfDeviceGroup\x12\x45\n\x05items\x18\x01 \x03(\x0b\x32/.io.clbs.openhes.models.acquisition.DeviceGroupR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\x8b\x02\n\x11StreamDeviceGroup\x12I\n\x04spec\x18\x01 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.DeviceGroupSpecH\x00R\x04spec\x12U\n\x06status\x18\x02 \x01(\x0b\x32;.io.clbs.openhes.models.acquisition.StreamDeviceGroupStatusH\x00R\x06status\x12K\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsH\x00R\x08metadataB\x07\n\x05parts\"\xae\x01\n\x0b\x44\x65viceGroup\x12N\n\x04spec\x18\x01 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.DeviceGroupSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadataJ\x04\x08\x02\x10\x03\"2\n\x0f\x44\x65viceGroupSpec\x12\x1f\n\x0b\x65xternal_id\x18\x01 \x01(\tR\nexternalId\"\xf6\x01\n\x17StreamDeviceGroupStatus\x12\x62\n\x07\x64\x65vices\x18\x04 \x03(\x0b\x32H.io.clbs.openhes.models.acquisition.StreamDeviceGroupStatus.DevicesEntryR\x07\x64\x65vices\x1aw\n\x0c\x44\x65vicesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12Q\n\x05value\x18\x02 \x01(\x0b\x32;.io.clbs.openhes.models.acquisition.DeviceGroupStatusDeviceR\x05value:\x02\x38\x01\":\n\x17\x44\x65viceGroupStatusDevice\x12\x1f\n\x0b\x64river_type\x18\x01 \x01(\tR\ndriverType\"\xaf\x01\n\"SetDeviceCommunicationUnitsRequest\x12\x1b\n\tdevice_id\x18\x01 \x01(\tR\x08\x64\x65viceId\x12l\n\x13\x63ommunication_units\x18\x02 \x03(\x0b\x32;.io.clbs.openhes.models.acquisition.DeviceCommunicationUnitR\x12\x63ommunicationUnits\"\x93\x01\n\x1dListOfDeviceCommunicationUnit\x12Q\n\x05items\x18\x01 \x03(\x0b\x32;.io.clbs.openhes.models.acquisition.DeviceCommunicationUnitR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"R\n\x18\x41\x64\x64\x44\x65vicesToGroupRequest\x12\x19\n\x08group_id\x18\x01 \x01(\tR\x07groupId\x12\x1b\n\tdevice_id\x18\x02 \x03(\tR\x08\x64\x65viceId\"W\n\x1dRemoveDevicesFromGroupRequest\x12\x19\n\x08group_id\x18\x01 \x01(\tR\x07groupId\x12\x1b\n\tdevice_id\x18\x02 \x03(\tR\x08\x64\x65viceId\"\x83\x01\n\x1dListDeviceGroupDevicesRequest\x12\x19\n\x08group_id\x18\x01 \x01(\tR\x07groupId\x12G\n\x08selector\x18\x02 \x01(\x0b\x32+.io.clbs.openhes.models.common.ListSelectorR\x08selector\"\xa7\x01\n\x11\x43reateBulkRequest\x12G\n\x04spec\x18\x01 \x01(\x0b\x32,.io.clbs.openhes.models.acquisition.BulkSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"m\n\nListOfBulk\x12>\n\x05items\x18\x01 \x03(\x0b\x32(.io.clbs.openhes.models.acquisition.BulkR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\xe2\x01\n\x04\x42ulk\x12G\n\x04spec\x18\x01 \x01(\x0b\x32,.io.clbs.openhes.models.acquisition.BulkSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12\x46\n\x06status\x18\x02 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.BulkStatusR\x06status\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"\xf2\x02\n\x08\x42ulkSpec\x12%\n\x0e\x63orrelation_id\x18\x01 \x01(\tR\rcorrelationId\x12Q\n\x07\x64\x65vices\x18\x02 \x01(\x0b\x32\x35.io.clbs.openhes.models.acquisition.ListOfJobDeviceIdH\x00R\x07\x64\x65vices\x12(\n\x0f\x64\x65vice_group_id\x18\x03 \x01(\tH\x00R\rdeviceGroupId\x12K\n\x08settings\x18\x04 \x01(\x0b\x32/.io.clbs.openhes.models.acquisition.JobSettingsR\x08settings\x12J\n\x07\x61\x63tions\x18\x05 \x03(\x0b\x32\x30.io.clbs.openhes.models.acquisition.JobActionSetR\x07\x61\x63tions\x12\x1f\n\x0bwebhook_url\x18\x06 \x01(\tR\nwebhookUrlB\x08\n\x06\x64\x65vice\"\xf8\x02\n\nBulkStatus\x12J\n\x06status\x18\x01 \x01(\x0e\x32\x32.io.clbs.openhes.models.acquisition.BulkStatusCodeR\x06status\x12\x1d\n\njobs_count\x18\x02 \x01(\x05R\tjobsCount\x12#\n\rjobs_finished\x18\x03 \x01(\x05R\x0cjobsFinished\x12\'\n\x0fjobs_successful\x18\x04 \x01(\x05R\x0ejobsSuccessful\x12\x39\n\ncreated_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x39\n\nstarted_at\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tstartedAt\x12;\n\x0b\x66inished_at\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\nfinishedAt\"\xe0\x01\n\x07\x42ulkJob\x12\x43\n\x04spec\x18\x01 \x01(\x0b\x32/.io.clbs.openhes.models.acquisition.BulkJobSpecR\x04spec\x12\x45\n\x06status\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.acquisition.JobStatusR\x06status\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"T\n\x0b\x42ulkJobSpec\x12\x45\n\x06\x64\x65vice\x18\x01 \x01(\x0b\x32-.io.clbs.openhes.models.acquisition.JobDeviceR\x06\x64\x65vice\"w\n\x13ListBulkJobsRequest\x12\x17\n\x07\x62ulk_id\x18\x01 \x01(\tR\x06\x62ulkId\x12G\n\x08selector\x18\x02 \x01(\x0b\x32+.io.clbs.openhes.models.common.ListSelectorR\x08selector\"s\n\rListOfBulkJob\x12\x41\n\x05items\x18\x01 \x03(\x0b\x32+.io.clbs.openhes.models.acquisition.BulkJobR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\xb1\x01\n\x16\x43reateProxyBulkRequest\x12L\n\x04spec\x18\x01 \x01(\x0b\x32\x31.io.clbs.openhes.models.acquisition.ProxyBulkSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"\xe4\x02\n\rProxyBulkSpec\x12%\n\x0e\x63orrelation_id\x18\x01 \x01(\tR\rcorrelationId\x12&\n\x0b\x64river_type\x18\x02 \x01(\tB\x05\xaa\x01\x02\x08\x03R\ndriverType\x12M\n\x07\x64\x65vices\x18\x03 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.ListOfJobDeviceR\x07\x64\x65vices\x12K\n\x08settings\x18\x04 \x01(\x0b\x32/.io.clbs.openhes.models.acquisition.JobSettingsR\x08settings\x12G\n\x07\x61\x63tions\x18\x05 \x03(\x0b\x32-.io.clbs.openhes.models.acquisition.JobActionR\x07\x61\x63tions\x12\x1f\n\x0bwebhook_url\x18\x06 \x01(\tR\nwebhookUrl\"\xec\x01\n\tProxyBulk\x12L\n\x04spec\x18\x01 \x01(\x0b\x32\x31.io.clbs.openhes.models.acquisition.ProxyBulkSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12\x46\n\x06status\x18\x02 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.BulkStatusR\x06status\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"\xb4\x01\n\x13SetModemPoolRequest\x12L\n\x04spec\x18\x01 \x01(\x0b\x32\x31.io.clbs.openhes.models.acquisition.ModemPoolSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadataJ\x04\x08\x02\x10\x03\"w\n\x0fListOfModemPool\x12\x43\n\x05items\x18\x01 \x03(\x0b\x32-.io.clbs.openhes.models.acquisition.ModemPoolR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\x0f\n\rModemPoolSpec\"X\n\x0fModemPoolStatus\x12\x45\n\x06modems\x18\x01 \x03(\x0b\x32-.io.clbs.openhes.models.acquisition.ModemInfoR\x06modems\"\xf1\x01\n\tModemPool\x12L\n\x04spec\x18\x01 \x01(\x0b\x32\x31.io.clbs.openhes.models.acquisition.ModemPoolSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12K\n\x06status\x18\x02 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.ModemPoolStatusR\x06status\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"o\n\x0fSetModemRequest\x12\x17\n\x07pool_id\x18\x01 \x01(\tR\x06poolId\x12\x43\n\x05modem\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.acquisition.ModemInfoR\x05modem\"b\n\tSetDriver\x12I\n\x04spec\x18\x01 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.DriverSpecB\x05\xaa\x01\x02\x08\x03R\x04specJ\x04\x08\x02\x10\x03J\x04\x08\x03\x10\x04\"q\n\x0cListOfDriver\x12@\n\x05items\x18\x01 \x03(\x0b\x32*.io.clbs.openhes.models.acquisition.DriverR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"_\n\x06\x44river\x12I\n\x04spec\x18\x01 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.DriverSpecB\x05\xaa\x01\x02\x08\x03R\x04specJ\x04\x08\x02\x10\x03J\x04\x08\x03\x10\x04\"\xec\x02\n\nDriverSpec\x12\x18\n\x07version\x18\x01 \x01(\tR\x07version\x12%\n\x0elistening_port\x18\x02 \x01(\rR\rlisteningPort\x12\x1f\n\x0b\x64river_type\x18\x03 \x01(\tR\ndriverType\x12.\n\x13max_concurrent_jobs\x18\x04 \x01(\x05R\x11maxConcurrentJobs\x12*\n\x11max_cascade_depth\x18\x05 \x01(\rR\x0fmaxCascadeDepth\x12*\n\x11typical_mem_usage\x18\x06 \x01(\x05R\x0ftypicalMemUsage\x12Q\n\ttemplates\x18\x07 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.DriverTemplatesR\ttemplates\x12!\n\x0c\x64isplay_name\x18\x08 \x01(\tR\x0b\x64isplayName\"\xaf\x01\n\x15\x43reateVariableRequest\x12K\n\x04spec\x18\x01 \x01(\x0b\x32\x30.io.clbs.openhes.models.acquisition.VariableSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"u\n\x0eListOfVariable\x12\x42\n\x05items\x18\x01 \x03(\x0b\x32,.io.clbs.openhes.models.acquisition.VariableR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\xa8\x01\n\x08Variable\x12K\n\x04spec\x18\x01 \x01(\x0b\x32\x30.io.clbs.openhes.models.acquisition.VariableSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadataJ\x04\x08\x02\x10\x03\"/\n\x0cVariableSpec\x12\x1f\n\x0bregister_id\x18\x01 \x03(\tR\nregisterId\"\xd5\x01\n(CreateDeviceConfigurationRegisterRequest\x12^\n\x04spec\x18\x01 \x01(\x0b\x32\x43.io.clbs.openhes.models.acquisition.DeviceConfigurationRegisterSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"\x9b\x01\n!ListOfDeviceConfigurationRegister\x12U\n\x05items\x18\x01 \x03(\x0b\x32?.io.clbs.openhes.models.acquisition.DeviceConfigurationRegisterR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\xce\x01\n\x1b\x44\x65viceConfigurationRegister\x12^\n\x04spec\x18\x01 \x01(\x0b\x32\x43.io.clbs.openhes.models.acquisition.DeviceConfigurationRegisterSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadataJ\x04\x08\x02\x10\x03\"\xd5\x01\n(CreateDeviceConfigurationTemplateRequest\x12^\n\x04spec\x18\x01 \x01(\x0b\x32\x43.io.clbs.openhes.models.acquisition.DeviceConfigurationTemplateSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"\x9b\x01\n!ListOfDeviceConfigurationTemplate\x12U\n\x05items\x18\x01 \x03(\x0b\x32?.io.clbs.openhes.models.acquisition.DeviceConfigurationTemplateR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\xce\x01\n\x1b\x44\x65viceConfigurationTemplate\x12^\n\x04spec\x18\x01 \x01(\x0b\x32\x43.io.clbs.openhes.models.acquisition.DeviceConfigurationTemplateSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadataJ\x04\x08\x02\x10\x03\"j\n\x1f\x44\x65viceConfigurationTemplateSpec\x12&\n\x0b\x64river_type\x18\x01 \x01(\tB\x05\xaa\x01\x02\x08\x03R\ndriverType\x12\x1f\n\x0bregister_id\x18\x02 \x03(\tR\nregisterId\"|\nBAddDeviceConfigurationRegisterToDeviceConfigurationTemplateRequest\x12\x15\n\x06\x64\x63t_id\x18\x01 \x01(\tR\x05\x64\x63tId\x12\x1f\n\x0bregister_id\x18\x02 \x01(\tR\nregisterId\"\x81\x01\nGRemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplateRequest\x12\x15\n\x06\x64\x63t_id\x18\x01 \x01(\tR\x05\x64\x63tId\x12\x1f\n\x0bregister_id\x18\x02 \x01(\tR\nregisterId\"\xc1\x01\n\x13GetMeterDataRequest\x12\x35\n\x04\x66rom\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x05\xaa\x01\x02\x08\x03R\x04\x66rom\x12\x31\n\x02to\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x05\xaa\x01\x02\x08\x03R\x02to\x12\x1b\n\tdevice_id\x18\x03 \x03(\tR\x08\x64\x65viceId\x12#\n\rvariable_name\x18\x04 \x03(\tR\x0cvariableName\"`\n\x15StreamMeterDataPacket\x12G\n\x04part\x18\x01 \x03(\x0b\x32\x33.io.clbs.openhes.models.acquisition.StreamMeterDataR\x04part\"\x93\x01\n\x0fStreamMeterData\x12\x31\n\x02ts\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x05\xaa\x01\x02\x08\x03R\x02ts\x12M\n\x04\x64\x61ta\x18\x02 \x03(\x0b\x32\x39.io.clbs.openhes.models.acquisition.StreamMeterDataDeviceR\x04\x64\x61ta\"\x82\x01\n\x15StreamMeterDataDevice\x12\"\n\tdevice_id\x18\x01 \x01(\tB\x05\xaa\x01\x02\x08\x03R\x08\x64\x65viceId\x12\x45\n\x04\x64\x61ta\x18\x02 \x03(\x0b\x32\x31.io.clbs.openhes.models.acquisition.MeasuredValueR\x04\x64\x61taB5Z3github.com/cybroslabs/hes-2-apis/gen/go/acquisitionb\x08\x65\x64itionsp\xe8\x07')

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
  _globals['_STREAMDEVICEGROUPSTATUS_DEVICESENTRY']._loaded_options = None
  _globals['_STREAMDEVICEGROUPSTATUS_DEVICESENTRY']._serialized_options = b'8\001'
  _globals['_CREATEBULKREQUEST'].fields_by_name['spec']._loaded_options = None
  _globals['_CREATEBULKREQUEST'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_BULK'].fields_by_name['spec']._loaded_options = None
  _globals['_BULK'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_CREATEPROXYBULKREQUEST'].fields_by_name['spec']._loaded_options = None
  _globals['_CREATEPROXYBULKREQUEST'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_PROXYBULKSPEC'].fields_by_name['driver_type']._loaded_options = None
  _globals['_PROXYBULKSPEC'].fields_by_name['driver_type']._serialized_options = b'\252\001\002\010\003'
  _globals['_PROXYBULK'].fields_by_name['spec']._loaded_options = None
  _globals['_PROXYBULK'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_SETMODEMPOOLREQUEST'].fields_by_name['spec']._loaded_options = None
  _globals['_SETMODEMPOOLREQUEST'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_MODEMPOOL'].fields_by_name['spec']._loaded_options = None
  _globals['_MODEMPOOL'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_SETDRIVER'].fields_by_name['spec']._loaded_options = None
  _globals['_SETDRIVER'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_DRIVER'].fields_by_name['spec']._loaded_options = None
  _globals['_DRIVER'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_CREATEVARIABLEREQUEST'].fields_by_name['spec']._loaded_options = None
  _globals['_CREATEVARIABLEREQUEST'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_VARIABLE'].fields_by_name['spec']._loaded_options = None
  _globals['_VARIABLE'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_CREATEDEVICECONFIGURATIONREGISTERREQUEST'].fields_by_name['spec']._loaded_options = None
  _globals['_CREATEDEVICECONFIGURATIONREGISTERREQUEST'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_DEVICECONFIGURATIONREGISTER'].fields_by_name['spec']._loaded_options = None
  _globals['_DEVICECONFIGURATIONREGISTER'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_CREATEDEVICECONFIGURATIONTEMPLATEREQUEST'].fields_by_name['spec']._loaded_options = None
  _globals['_CREATEDEVICECONFIGURATIONTEMPLATEREQUEST'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_DEVICECONFIGURATIONTEMPLATE'].fields_by_name['spec']._loaded_options = None
  _globals['_DEVICECONFIGURATIONTEMPLATE'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_DEVICECONFIGURATIONTEMPLATESPEC'].fields_by_name['driver_type']._loaded_options = None
  _globals['_DEVICECONFIGURATIONTEMPLATESPEC'].fields_by_name['driver_type']._serialized_options = b'\252\001\002\010\003'
  _globals['_GETMETERDATAREQUEST'].fields_by_name['from']._loaded_options = None
  _globals['_GETMETERDATAREQUEST'].fields_by_name['from']._serialized_options = b'\252\001\002\010\003'
  _globals['_GETMETERDATAREQUEST'].fields_by_name['to']._loaded_options = None
  _globals['_GETMETERDATAREQUEST'].fields_by_name['to']._serialized_options = b'\252\001\002\010\003'
  _globals['_STREAMMETERDATA'].fields_by_name['ts']._loaded_options = None
  _globals['_STREAMMETERDATA'].fields_by_name['ts']._serialized_options = b'\252\001\002\010\003'
  _globals['_STREAMMETERDATADEVICE'].fields_by_name['device_id']._loaded_options = None
  _globals['_STREAMMETERDATADEVICE'].fields_by_name['device_id']._serialized_options = b'\252\001\002\010\003'
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
  _globals['_STREAMDEVICEGROUP']._serialized_end=2190
  _globals['_DEVICEGROUP']._serialized_start=2193
  _globals['_DEVICEGROUP']._serialized_end=2367
  _globals['_DEVICEGROUPSPEC']._serialized_start=2369
  _globals['_DEVICEGROUPSPEC']._serialized_end=2419
  _globals['_STREAMDEVICEGROUPSTATUS']._serialized_start=2422
  _globals['_STREAMDEVICEGROUPSTATUS']._serialized_end=2668
  _globals['_STREAMDEVICEGROUPSTATUS_DEVICESENTRY']._serialized_start=2549
  _globals['_STREAMDEVICEGROUPSTATUS_DEVICESENTRY']._serialized_end=2668
  _globals['_DEVICEGROUPSTATUSDEVICE']._serialized_start=2670
  _globals['_DEVICEGROUPSTATUSDEVICE']._serialized_end=2728
  _globals['_SETDEVICECOMMUNICATIONUNITSREQUEST']._serialized_start=2731
  _globals['_SETDEVICECOMMUNICATIONUNITSREQUEST']._serialized_end=2906
  _globals['_LISTOFDEVICECOMMUNICATIONUNIT']._serialized_start=2909
  _globals['_LISTOFDEVICECOMMUNICATIONUNIT']._serialized_end=3056
  _globals['_ADDDEVICESTOGROUPREQUEST']._serialized_start=3058
  _globals['_ADDDEVICESTOGROUPREQUEST']._serialized_end=3140
  _globals['_REMOVEDEVICESFROMGROUPREQUEST']._serialized_start=3142
  _globals['_REMOVEDEVICESFROMGROUPREQUEST']._serialized_end=3229
  _globals['_LISTDEVICEGROUPDEVICESREQUEST']._serialized_start=3232
  _globals['_LISTDEVICEGROUPDEVICESREQUEST']._serialized_end=3363
  _globals['_CREATEBULKREQUEST']._serialized_start=3366
  _globals['_CREATEBULKREQUEST']._serialized_end=3533
  _globals['_LISTOFBULK']._serialized_start=3535
  _globals['_LISTOFBULK']._serialized_end=3644
  _globals['_BULK']._serialized_start=3647
  _globals['_BULK']._serialized_end=3873
  _globals['_BULKSPEC']._serialized_start=3876
  _globals['_BULKSPEC']._serialized_end=4246
  _globals['_BULKSTATUS']._serialized_start=4249
  _globals['_BULKSTATUS']._serialized_end=4625
  _globals['_BULKJOB']._serialized_start=4628
  _globals['_BULKJOB']._serialized_end=4852
  _globals['_BULKJOBSPEC']._serialized_start=4854
  _globals['_BULKJOBSPEC']._serialized_end=4938
  _globals['_LISTBULKJOBSREQUEST']._serialized_start=4940
  _globals['_LISTBULKJOBSREQUEST']._serialized_end=5059
  _globals['_LISTOFBULKJOB']._serialized_start=5061
  _globals['_LISTOFBULKJOB']._serialized_end=5176
  _globals['_CREATEPROXYBULKREQUEST']._serialized_start=5179
  _globals['_CREATEPROXYBULKREQUEST']._serialized_end=5356
  _globals['_PROXYBULKSPEC']._serialized_start=5359
  _globals['_PROXYBULKSPEC']._serialized_end=5715
  _globals['_PROXYBULK']._serialized_start=5718
  _globals['_PROXYBULK']._serialized_end=5954
  _globals['_SETMODEMPOOLREQUEST']._serialized_start=5957
  _globals['_SETMODEMPOOLREQUEST']._serialized_end=6137
  _globals['_LISTOFMODEMPOOL']._serialized_start=6139
  _globals['_LISTOFMODEMPOOL']._serialized_end=6258
  _globals['_MODEMPOOLSPEC']._serialized_start=6260
  _globals['_MODEMPOOLSPEC']._serialized_end=6275
  _globals['_MODEMPOOLSTATUS']._serialized_start=6277
  _globals['_MODEMPOOLSTATUS']._serialized_end=6365
  _globals['_MODEMPOOL']._serialized_start=6368
  _globals['_MODEMPOOL']._serialized_end=6609
  _globals['_SETMODEMREQUEST']._serialized_start=6611
  _globals['_SETMODEMREQUEST']._serialized_end=6722
  _globals['_SETDRIVER']._serialized_start=6724
  _globals['_SETDRIVER']._serialized_end=6822
  _globals['_LISTOFDRIVER']._serialized_start=6824
  _globals['_LISTOFDRIVER']._serialized_end=6937
  _globals['_DRIVER']._serialized_start=6939
  _globals['_DRIVER']._serialized_end=7034
  _globals['_DRIVERSPEC']._serialized_start=7037
  _globals['_DRIVERSPEC']._serialized_end=7401
  _globals['_CREATEVARIABLEREQUEST']._serialized_start=7404
  _globals['_CREATEVARIABLEREQUEST']._serialized_end=7579
  _globals['_LISTOFVARIABLE']._serialized_start=7581
  _globals['_LISTOFVARIABLE']._serialized_end=7698
  _globals['_VARIABLE']._serialized_start=7701
  _globals['_VARIABLE']._serialized_end=7869
  _globals['_VARIABLESPEC']._serialized_start=7871
  _globals['_VARIABLESPEC']._serialized_end=7918
  _globals['_CREATEDEVICECONFIGURATIONREGISTERREQUEST']._serialized_start=7921
  _globals['_CREATEDEVICECONFIGURATIONREGISTERREQUEST']._serialized_end=8134
  _globals['_LISTOFDEVICECONFIGURATIONREGISTER']._serialized_start=8137
  _globals['_LISTOFDEVICECONFIGURATIONREGISTER']._serialized_end=8292
  _globals['_DEVICECONFIGURATIONREGISTER']._serialized_start=8295
  _globals['_DEVICECONFIGURATIONREGISTER']._serialized_end=8501
  _globals['_CREATEDEVICECONFIGURATIONTEMPLATEREQUEST']._serialized_start=8504
  _globals['_CREATEDEVICECONFIGURATIONTEMPLATEREQUEST']._serialized_end=8717
  _globals['_LISTOFDEVICECONFIGURATIONTEMPLATE']._serialized_start=8720
  _globals['_LISTOFDEVICECONFIGURATIONTEMPLATE']._serialized_end=8875
  _globals['_DEVICECONFIGURATIONTEMPLATE']._serialized_start=8878
  _globals['_DEVICECONFIGURATIONTEMPLATE']._serialized_end=9084
  _globals['_DEVICECONFIGURATIONTEMPLATESPEC']._serialized_start=9086
  _globals['_DEVICECONFIGURATIONTEMPLATESPEC']._serialized_end=9192
  _globals['_ADDDEVICECONFIGURATIONREGISTERTODEVICECONFIGURATIONTEMPLATEREQUEST']._serialized_start=9194
  _globals['_ADDDEVICECONFIGURATIONREGISTERTODEVICECONFIGURATIONTEMPLATEREQUEST']._serialized_end=9318
  _globals['_REMOVEDEVICECONFIGURATIONREGISTERFROMDEVICECONFIGURATIONTEMPLATEREQUEST']._serialized_start=9321
  _globals['_REMOVEDEVICECONFIGURATIONREGISTERFROMDEVICECONFIGURATIONTEMPLATEREQUEST']._serialized_end=9450
  _globals['_GETMETERDATAREQUEST']._serialized_start=9453
  _globals['_GETMETERDATAREQUEST']._serialized_end=9646
  _globals['_STREAMMETERDATAPACKET']._serialized_start=9648
  _globals['_STREAMMETERDATAPACKET']._serialized_end=9744
  _globals['_STREAMMETERDATA']._serialized_start=9747
  _globals['_STREAMMETERDATA']._serialized_end=9894
  _globals['_STREAMMETERDATADEVICE']._serialized_start=9897
  _globals['_STREAMMETERDATADEVICE']._serialized_end=10027
# @@protoc_insertion_point(module_scope)
