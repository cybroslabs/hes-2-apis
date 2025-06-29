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
from acquisition.timeofuse import timeofuse_pb2 as acquisition_dot_timeofuse_dot_timeofuse__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x16\x61\x63quisition/main.proto\x12\"io.clbs.openhes.models.acquisition\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x15\x63ommon/metadata.proto\x1a\x13\x63ommon/fields.proto\x1a\x18\x61\x63quisition/shared.proto\x1a%acquisition/timeofuse/timeofuse.proto\"\xba\x01\n\x1e\x43reateCommunicationUnitRequest\x12M\n\x04spec\x18\x01 \x01(\x0b\x32\x39.io.clbs.openhes.models.acquisition.CommunicationUnitSpecR\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"\x87\x01\n\x17ListOfCommunicationUnit\x12K\n\x05items\x18\x01 \x03(\x0b\x32\x35.io.clbs.openhes.models.acquisition.CommunicationUnitR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"p\n\x1d\x43reateCommunicationBusRequest\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadataJ\x04\x08\x01\x10\x02\"\x85\x01\n\x16ListOfCommunicationBus\x12J\n\x05items\x18\x01 \x03(\x0b\x32\x34.io.clbs.openhes.models.acquisition.CommunicationBusR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\x96\x01\n.AddCommunicationUnitsToCommunicationBusRequest\x12\x30\n\x14\x63ommunication_bus_id\x18\x01 \x01(\tR\x12\x63ommunicationBusId\x12\x32\n\x15\x63ommunication_unit_id\x18\x02 \x03(\tR\x13\x63ommunicationUnitId\"\x9b\x01\n3RemoveCommunicationUnitsFromCommunicationBusRequest\x12\x30\n\x14\x63ommunication_bus_id\x18\x01 \x01(\tR\x12\x63ommunicationBusId\x12\x32\n\x15\x63ommunication_unit_id\x18\x02 \x03(\tR\x13\x63ommunicationUnitId\"\xa4\x01\n\x13\x43reateDeviceRequest\x12\x42\n\x04spec\x18\x01 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.DeviceSpecR\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"q\n\x0cListOfDevice\x12@\n\x05items\x18\x01 \x03(\x0b\x32*.io.clbs.openhes.models.acquisition.DeviceR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\xe1\x01\n\x06\x44\x65vice\x12\x42\n\x04spec\x18\x01 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.DeviceSpecR\x04spec\x12H\n\x06status\x18\x02 \x01(\x0b\x32\x30.io.clbs.openhes.models.acquisition.DeviceStatusR\x06status\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\">\n\x1fStreamDevicesDriverTypesRequest\x12\x1b\n\tdevice_id\x18\x01 \x03(\tR\x08\x64\x65viceId\"\xbf\x01\n StreamDevicesDriverTypesResponse\x12\x62\n\x04\x64\x61ta\x18\x01 \x03(\x0b\x32N.io.clbs.openhes.models.acquisition.StreamDevicesDriverTypesResponse.DataEntryR\x04\x64\x61ta\x1a\x37\n\tDataEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\"\xb1\x01\n\rUnknownDevice\x12O\n\x06status\x18\x02 \x01(\x0b\x32\x37.io.clbs.openhes.models.acquisition.UnknownDeviceStatusR\x06status\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadataJ\x04\x08\x01\x10\x02\"j\n\x13UnknownDeviceStatus\x12\x1f\n\x0b\x64river_type\x18\x01 \x01(\tR\ndriverType\x12\x32\n\x15\x63ommunication_unit_id\x18\x02 \x01(\tR\x13\x63ommunicationUnitId\"\xae\x01\n\x18\x43reateDeviceGroupRequest\x12G\n\x04spec\x18\x01 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.DeviceGroupSpecR\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"{\n\x11ListOfDeviceGroup\x12\x45\n\x05items\x18\x01 \x03(\x0b\x32/.io.clbs.openhes.models.acquisition.DeviceGroupR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\x8b\x02\n\x11StreamDeviceGroup\x12I\n\x04spec\x18\x01 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.DeviceGroupSpecH\x00R\x04spec\x12U\n\x06status\x18\x02 \x01(\x0b\x32;.io.clbs.openhes.models.acquisition.StreamDeviceGroupStatusH\x00R\x06status\x12K\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsH\x00R\x08metadataB\x07\n\x05parts\"\xa7\x01\n\x0b\x44\x65viceGroup\x12G\n\x04spec\x18\x01 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.DeviceGroupSpecR\x04spec\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadataJ\x04\x08\x02\x10\x03\"2\n\x0f\x44\x65viceGroupSpec\x12\x1f\n\x0b\x65xternal_id\x18\x01 \x01(\tR\nexternalId\"\xf6\x01\n\x17StreamDeviceGroupStatus\x12\x62\n\x07\x64\x65vices\x18\x04 \x03(\x0b\x32H.io.clbs.openhes.models.acquisition.StreamDeviceGroupStatus.DevicesEntryR\x07\x64\x65vices\x1aw\n\x0c\x44\x65vicesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12Q\n\x05value\x18\x02 \x01(\x0b\x32;.io.clbs.openhes.models.acquisition.DeviceGroupStatusDeviceR\x05value:\x02\x38\x01\":\n\x17\x44\x65viceGroupStatusDevice\x12\x1f\n\x0b\x64river_type\x18\x01 \x01(\tR\ndriverType\"\xaf\x01\n\"SetDeviceCommunicationUnitsRequest\x12\x1b\n\tdevice_id\x18\x01 \x01(\tR\x08\x64\x65viceId\x12l\n\x13\x63ommunication_units\x18\x02 \x03(\x0b\x32;.io.clbs.openhes.models.acquisition.DeviceCommunicationUnitR\x12\x63ommunicationUnits\"\x93\x01\n\x1dListOfDeviceCommunicationUnit\x12Q\n\x05items\x18\x01 \x03(\x0b\x32;.io.clbs.openhes.models.acquisition.DeviceCommunicationUnitR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"R\n\x18\x41\x64\x64\x44\x65vicesToGroupRequest\x12\x19\n\x08group_id\x18\x01 \x01(\tR\x07groupId\x12\x1b\n\tdevice_id\x18\x02 \x03(\tR\x08\x64\x65viceId\"W\n\x1dRemoveDevicesFromGroupRequest\x12\x19\n\x08group_id\x18\x01 \x01(\tR\x07groupId\x12\x1b\n\tdevice_id\x18\x02 \x03(\tR\x08\x64\x65viceId\"\x83\x01\n\x1dListDeviceGroupDevicesRequest\x12\x19\n\x08group_id\x18\x01 \x01(\tR\x07groupId\x12G\n\x08selector\x18\x02 \x01(\x0b\x32+.io.clbs.openhes.models.common.ListSelectorR\x08selector\"\xa0\x01\n\x11\x43reateBulkRequest\x12@\n\x04spec\x18\x01 \x01(\x0b\x32,.io.clbs.openhes.models.acquisition.BulkSpecR\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"m\n\nListOfBulk\x12>\n\x05items\x18\x01 \x03(\x0b\x32(.io.clbs.openhes.models.acquisition.BulkR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\xdb\x01\n\x04\x42ulk\x12@\n\x04spec\x18\x01 \x01(\x0b\x32,.io.clbs.openhes.models.acquisition.BulkSpecR\x04spec\x12\x46\n\x06status\x18\x02 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.BulkStatusR\x06status\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"\xf2\x02\n\x08\x42ulkSpec\x12%\n\x0e\x63orrelation_id\x18\x01 \x01(\tR\rcorrelationId\x12Q\n\x07\x64\x65vices\x18\x02 \x01(\x0b\x32\x35.io.clbs.openhes.models.acquisition.ListOfJobDeviceIdH\x00R\x07\x64\x65vices\x12(\n\x0f\x64\x65vice_group_id\x18\x03 \x01(\tH\x00R\rdeviceGroupId\x12K\n\x08settings\x18\x04 \x01(\x0b\x32/.io.clbs.openhes.models.acquisition.JobSettingsR\x08settings\x12J\n\x07\x61\x63tions\x18\x05 \x03(\x0b\x32\x30.io.clbs.openhes.models.acquisition.JobActionSetR\x07\x61\x63tions\x12\x1f\n\x0bwebhook_url\x18\x06 \x01(\tR\nwebhookUrlB\x08\n\x06\x64\x65vice\"\xf8\x02\n\nBulkStatus\x12J\n\x06status\x18\x01 \x01(\x0e\x32\x32.io.clbs.openhes.models.acquisition.BulkStatusCodeR\x06status\x12\x1d\n\njobs_count\x18\x02 \x01(\x05R\tjobsCount\x12#\n\rjobs_finished\x18\x03 \x01(\x05R\x0cjobsFinished\x12\'\n\x0fjobs_successful\x18\x04 \x01(\x05R\x0ejobsSuccessful\x12\x39\n\ncreated_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x39\n\nstarted_at\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tstartedAt\x12;\n\x0b\x66inished_at\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\nfinishedAt\"\xe0\x01\n\x07\x42ulkJob\x12\x43\n\x04spec\x18\x01 \x01(\x0b\x32/.io.clbs.openhes.models.acquisition.BulkJobSpecR\x04spec\x12\x45\n\x06status\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.acquisition.JobStatusR\x06status\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"T\n\x0b\x42ulkJobSpec\x12\x45\n\x06\x64\x65vice\x18\x01 \x01(\x0b\x32-.io.clbs.openhes.models.acquisition.JobDeviceR\x06\x64\x65vice\"w\n\x13ListBulkJobsRequest\x12\x17\n\x07\x62ulk_id\x18\x01 \x01(\tR\x06\x62ulkId\x12G\n\x08selector\x18\x02 \x01(\x0b\x32+.io.clbs.openhes.models.common.ListSelectorR\x08selector\"s\n\rListOfBulkJob\x12\x41\n\x05items\x18\x01 \x03(\x0b\x32+.io.clbs.openhes.models.acquisition.BulkJobR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\xaa\x01\n\x16\x43reateProxyBulkRequest\x12\x45\n\x04spec\x18\x01 \x01(\x0b\x32\x31.io.clbs.openhes.models.acquisition.ProxyBulkSpecR\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"\xdd\x02\n\rProxyBulkSpec\x12%\n\x0e\x63orrelation_id\x18\x01 \x01(\tR\rcorrelationId\x12\x1f\n\x0b\x64river_type\x18\x02 \x01(\tR\ndriverType\x12M\n\x07\x64\x65vices\x18\x03 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.ListOfJobDeviceR\x07\x64\x65vices\x12K\n\x08settings\x18\x04 \x01(\x0b\x32/.io.clbs.openhes.models.acquisition.JobSettingsR\x08settings\x12G\n\x07\x61\x63tions\x18\x05 \x03(\x0b\x32-.io.clbs.openhes.models.acquisition.JobActionR\x07\x61\x63tions\x12\x1f\n\x0bwebhook_url\x18\x06 \x01(\tR\nwebhookUrl\"\xe5\x01\n\tProxyBulk\x12\x45\n\x04spec\x18\x01 \x01(\x0b\x32\x31.io.clbs.openhes.models.acquisition.ProxyBulkSpecR\x04spec\x12\x46\n\x06status\x18\x02 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.BulkStatusR\x06status\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"\xad\x01\n\x13SetModemPoolRequest\x12\x45\n\x04spec\x18\x01 \x01(\x0b\x32\x31.io.clbs.openhes.models.acquisition.ModemPoolSpecR\x04spec\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadataJ\x04\x08\x02\x10\x03\"w\n\x0fListOfModemPool\x12\x43\n\x05items\x18\x01 \x03(\x0b\x32-.io.clbs.openhes.models.acquisition.ModemPoolR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\x0f\n\rModemPoolSpec\"X\n\x0fModemPoolStatus\x12\x45\n\x06modems\x18\x01 \x03(\x0b\x32-.io.clbs.openhes.models.acquisition.ModemInfoR\x06modems\"\xea\x01\n\tModemPool\x12\x45\n\x04spec\x18\x01 \x01(\x0b\x32\x31.io.clbs.openhes.models.acquisition.ModemPoolSpecR\x04spec\x12K\n\x06status\x18\x02 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.ModemPoolStatusR\x06status\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"o\n\x0fSetModemRequest\x12\x17\n\x07pool_id\x18\x01 \x01(\tR\x06poolId\x12\x43\n\x05modem\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.acquisition.ModemInfoR\x05modem\"[\n\tSetDriver\x12\x42\n\x04spec\x18\x01 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.DriverSpecR\x04specJ\x04\x08\x02\x10\x03J\x04\x08\x03\x10\x04\"q\n\x0cListOfDriver\x12@\n\x05items\x18\x01 \x03(\x0b\x32*.io.clbs.openhes.models.acquisition.DriverR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\x9c\x01\n\x06\x44river\x12\x42\n\x04spec\x18\x01 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.DriverSpecR\x04spec\x12H\n\x06status\x18\x02 \x01(\x0b\x32\x30.io.clbs.openhes.models.acquisition.DriverStatusR\x06statusJ\x04\x08\x03\x10\x04\"\xec\x02\n\nDriverSpec\x12\x18\n\x07version\x18\x01 \x01(\tR\x07version\x12%\n\x0elistening_port\x18\x02 \x01(\rR\rlisteningPort\x12\x1f\n\x0b\x64river_type\x18\x03 \x01(\tR\ndriverType\x12.\n\x13max_concurrent_jobs\x18\x04 \x01(\x05R\x11maxConcurrentJobs\x12*\n\x11max_cascade_depth\x18\x05 \x01(\rR\x0fmaxCascadeDepth\x12*\n\x11typical_mem_usage\x18\x06 \x01(\x05R\x0ftypicalMemUsage\x12Q\n\ttemplates\x18\x07 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.DriverTemplatesR\ttemplates\x12!\n\x0c\x64isplay_name\x18\x08 \x01(\tR\x0b\x64isplayName\"T\n\x0c\x44riverStatus\x12\x1b\n\tis_latest\x18\x01 \x01(\x08R\x08isLatest\x12\'\n\x0fupdate_finished\x18\x02 \x01(\x08R\x0eupdateFinished\"\xa8\x01\n\x15\x43reateVariableRequest\x12\x44\n\x04spec\x18\x01 \x01(\x0b\x32\x30.io.clbs.openhes.models.acquisition.VariableSpecR\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"u\n\x0eListOfVariable\x12\x42\n\x05items\x18\x01 \x03(\x0b\x32,.io.clbs.openhes.models.acquisition.VariableR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\xa1\x01\n\x08Variable\x12\x44\n\x04spec\x18\x01 \x01(\x0b\x32\x30.io.clbs.openhes.models.acquisition.VariableSpecR\x04spec\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadataJ\x04\x08\x02\x10\x03\"\xa6\x01\n\x0cVariableSpec\x12\x1f\n\x0bregister_id\x18\x01 \x03(\tR\nregisterId\x12I\n\tdata_type\x18\x02 \x01(\x0e\x32,.io.clbs.openhes.models.common.FieldDataTypeR\x08\x64\x61taType\x12*\n\x11\x65xclude_data_from\x18\x03 \x01(\x08R\x0f\x65xcludeDataFrom\"`\n\x1c\x41\x64\x64RegisterToVariableRequest\x12\x1f\n\x0bvariable_id\x18\x01 \x01(\tR\nvariableId\x12\x1f\n\x0bregister_id\x18\x02 \x03(\tR\nregisterId\"e\n!RemoveRegisterFromVariableRequest\x12\x1f\n\x0bvariable_id\x18\x01 \x01(\tR\nvariableId\x12\x1f\n\x0bregister_id\x18\x02 \x03(\tR\nregisterId\"\xce\x01\n(CreateDeviceConfigurationRegisterRequest\x12W\n\x04spec\x18\x01 \x01(\x0b\x32\x43.io.clbs.openhes.models.acquisition.DeviceConfigurationRegisterSpecR\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"\x9b\x01\n!ListOfDeviceConfigurationRegister\x12U\n\x05items\x18\x01 \x03(\x0b\x32?.io.clbs.openhes.models.acquisition.DeviceConfigurationRegisterR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\xc7\x01\n\x1b\x44\x65viceConfigurationRegister\x12W\n\x04spec\x18\x01 \x01(\x0b\x32\x43.io.clbs.openhes.models.acquisition.DeviceConfigurationRegisterSpecR\x04spec\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadataJ\x04\x08\x02\x10\x03\"\xce\x01\n(CreateDeviceConfigurationTemplateRequest\x12W\n\x04spec\x18\x01 \x01(\x0b\x32\x43.io.clbs.openhes.models.acquisition.DeviceConfigurationTemplateSpecR\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"\x9b\x01\n!ListOfDeviceConfigurationTemplate\x12U\n\x05items\x18\x01 \x03(\x0b\x32?.io.clbs.openhes.models.acquisition.DeviceConfigurationTemplateR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\xc7\x01\n\x1b\x44\x65viceConfigurationTemplate\x12W\n\x04spec\x18\x01 \x01(\x0b\x32\x43.io.clbs.openhes.models.acquisition.DeviceConfigurationTemplateSpecR\x04spec\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadataJ\x04\x08\x02\x10\x03\"c\n\x1f\x44\x65viceConfigurationTemplateSpec\x12\x1f\n\x0b\x64river_type\x18\x01 \x01(\tR\ndriverType\x12\x1f\n\x0bregister_id\x18\x02 \x03(\tR\nregisterId\"|\nBAddDeviceConfigurationRegisterToDeviceConfigurationTemplateRequest\x12\x15\n\x06\x64\x63t_id\x18\x01 \x01(\tR\x05\x64\x63tId\x12\x1f\n\x0bregister_id\x18\x02 \x03(\tR\nregisterId\"\x81\x01\nGRemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplateRequest\x12\x15\n\x06\x64\x63t_id\x18\x01 \x01(\tR\x05\x64\x63tId\x12\x1f\n\x0bregister_id\x18\x02 \x03(\tR\nregisterId\"\xd0\x02\n\x14GetDeviceDataRequest\x12.\n\x04\x66rom\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x04\x66rom\x12*\n\x02to\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x02to\x12\x1b\n\tdevice_id\x18\x03 \x03(\tR\x08\x64\x65viceId\x12\x1f\n\x0bvariable_id\x18\x04 \x03(\tR\nvariableId\x12\x32\n\x15\x66ilter_include_status\x18\x05 \x01(\x03R\x13\x66ilterIncludeStatus\x12\x32\n\x15\x66ilter_exclude_status\x18\x06 \x01(\x03R\x13\x66ilterExcludeStatus\x12\x36\n\x08snapshot\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x08snapshot\"\\\n\nDeviceData\x12N\n\x07\x64\x65vices\x18\x01 \x03(\x0b\x32\x34.io.clbs.openhes.models.acquisition.DeviceDeviceDataR\x07\x64\x65vices\"{\n\x10\x44\x65viceDeviceData\x12\x1b\n\tdevice_id\x18\x01 \x01(\tR\x08\x64\x65viceId\x12J\n\x04\x64\x61ta\x18\x02 \x03(\x0b\x32\x36.io.clbs.openhes.models.acquisition.VariableDeviceDataR\x04\x64\x61ta\"\xce\x01\n\x12VariableDeviceData\x12\x1f\n\x0bvariable_id\x18\x01 \x01(\tR\nvariableId\x12:\n\ntimestamps\x18\x02 \x03(\x0b\x32\x1a.google.protobuf.TimestampR\ntimestamps\x12\x12\n\x04unit\x18\x03 \x03(\tR\x04unit\x12G\n\x05value\x18\x04 \x03(\x0b\x32\x31.io.clbs.openhes.models.acquisition.MeasuredValueR\x05value\"t\n\x16GetDeviceEventsRequest\x12.\n\x04\x66rom\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x04\x66rom\x12*\n\x02to\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x02to\"\xbe\x01\n\x1b\x43reateTimeOfUseTableRequest\x12T\n\x04spec\x18\x01 \x01(\x0b\x32@.io.clbs.openhes.models.acquisition.timeofuse.TimeOfUseTableSpecR\x04spec\x12I\n\x08metadata\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadata\"\x81\x01\n\x14ListOfTimeOfUseTable\x12H\n\x05items\x18\x01 \x03(\x0b\x32\x32.io.clbs.openhes.models.acquisition.TimeOfUseTableR\x05items\x12\x1f\n\x0btotal_count\x18\x02 \x01(\x05R\ntotalCount\"\xb7\x01\n\x0eTimeOfUseTable\x12T\n\x04spec\x18\x01 \x01(\x0b\x32@.io.clbs.openhes.models.acquisition.timeofuse.TimeOfUseTableSpecR\x04spec\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadataJ\x04\x08\x02\x10\x03\x42:Z8github.com/cybroslabs/ouro-api-shared/gen/go/acquisitionb\x08\x65\x64itionsp\xe8\x07')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'acquisition.main_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z8github.com/cybroslabs/ouro-api-shared/gen/go/acquisition'
  _globals['_STREAMDEVICESDRIVERTYPESRESPONSE_DATAENTRY']._loaded_options = None
  _globals['_STREAMDEVICESDRIVERTYPESRESPONSE_DATAENTRY']._serialized_options = b'8\001'
  _globals['_STREAMDEVICEGROUPSTATUS_DEVICESENTRY']._loaded_options = None
  _globals['_STREAMDEVICEGROUPSTATUS_DEVICESENTRY']._serialized_options = b'8\001'
  _globals['_CREATECOMMUNICATIONUNITREQUEST']._serialized_start=234
  _globals['_CREATECOMMUNICATIONUNITREQUEST']._serialized_end=420
  _globals['_LISTOFCOMMUNICATIONUNIT']._serialized_start=423
  _globals['_LISTOFCOMMUNICATIONUNIT']._serialized_end=558
  _globals['_CREATECOMMUNICATIONBUSREQUEST']._serialized_start=560
  _globals['_CREATECOMMUNICATIONBUSREQUEST']._serialized_end=672
  _globals['_LISTOFCOMMUNICATIONBUS']._serialized_start=675
  _globals['_LISTOFCOMMUNICATIONBUS']._serialized_end=808
  _globals['_ADDCOMMUNICATIONUNITSTOCOMMUNICATIONBUSREQUEST']._serialized_start=811
  _globals['_ADDCOMMUNICATIONUNITSTOCOMMUNICATIONBUSREQUEST']._serialized_end=961
  _globals['_REMOVECOMMUNICATIONUNITSFROMCOMMUNICATIONBUSREQUEST']._serialized_start=964
  _globals['_REMOVECOMMUNICATIONUNITSFROMCOMMUNICATIONBUSREQUEST']._serialized_end=1119
  _globals['_CREATEDEVICEREQUEST']._serialized_start=1122
  _globals['_CREATEDEVICEREQUEST']._serialized_end=1286
  _globals['_LISTOFDEVICE']._serialized_start=1288
  _globals['_LISTOFDEVICE']._serialized_end=1401
  _globals['_DEVICE']._serialized_start=1404
  _globals['_DEVICE']._serialized_end=1629
  _globals['_STREAMDEVICESDRIVERTYPESREQUEST']._serialized_start=1631
  _globals['_STREAMDEVICESDRIVERTYPESREQUEST']._serialized_end=1693
  _globals['_STREAMDEVICESDRIVERTYPESRESPONSE']._serialized_start=1696
  _globals['_STREAMDEVICESDRIVERTYPESRESPONSE']._serialized_end=1887
  _globals['_STREAMDEVICESDRIVERTYPESRESPONSE_DATAENTRY']._serialized_start=1832
  _globals['_STREAMDEVICESDRIVERTYPESRESPONSE_DATAENTRY']._serialized_end=1887
  _globals['_UNKNOWNDEVICE']._serialized_start=1890
  _globals['_UNKNOWNDEVICE']._serialized_end=2067
  _globals['_UNKNOWNDEVICESTATUS']._serialized_start=2069
  _globals['_UNKNOWNDEVICESTATUS']._serialized_end=2175
  _globals['_CREATEDEVICEGROUPREQUEST']._serialized_start=2178
  _globals['_CREATEDEVICEGROUPREQUEST']._serialized_end=2352
  _globals['_LISTOFDEVICEGROUP']._serialized_start=2354
  _globals['_LISTOFDEVICEGROUP']._serialized_end=2477
  _globals['_STREAMDEVICEGROUP']._serialized_start=2480
  _globals['_STREAMDEVICEGROUP']._serialized_end=2747
  _globals['_DEVICEGROUP']._serialized_start=2750
  _globals['_DEVICEGROUP']._serialized_end=2917
  _globals['_DEVICEGROUPSPEC']._serialized_start=2919
  _globals['_DEVICEGROUPSPEC']._serialized_end=2969
  _globals['_STREAMDEVICEGROUPSTATUS']._serialized_start=2972
  _globals['_STREAMDEVICEGROUPSTATUS']._serialized_end=3218
  _globals['_STREAMDEVICEGROUPSTATUS_DEVICESENTRY']._serialized_start=3099
  _globals['_STREAMDEVICEGROUPSTATUS_DEVICESENTRY']._serialized_end=3218
  _globals['_DEVICEGROUPSTATUSDEVICE']._serialized_start=3220
  _globals['_DEVICEGROUPSTATUSDEVICE']._serialized_end=3278
  _globals['_SETDEVICECOMMUNICATIONUNITSREQUEST']._serialized_start=3281
  _globals['_SETDEVICECOMMUNICATIONUNITSREQUEST']._serialized_end=3456
  _globals['_LISTOFDEVICECOMMUNICATIONUNIT']._serialized_start=3459
  _globals['_LISTOFDEVICECOMMUNICATIONUNIT']._serialized_end=3606
  _globals['_ADDDEVICESTOGROUPREQUEST']._serialized_start=3608
  _globals['_ADDDEVICESTOGROUPREQUEST']._serialized_end=3690
  _globals['_REMOVEDEVICESFROMGROUPREQUEST']._serialized_start=3692
  _globals['_REMOVEDEVICESFROMGROUPREQUEST']._serialized_end=3779
  _globals['_LISTDEVICEGROUPDEVICESREQUEST']._serialized_start=3782
  _globals['_LISTDEVICEGROUPDEVICESREQUEST']._serialized_end=3913
  _globals['_CREATEBULKREQUEST']._serialized_start=3916
  _globals['_CREATEBULKREQUEST']._serialized_end=4076
  _globals['_LISTOFBULK']._serialized_start=4078
  _globals['_LISTOFBULK']._serialized_end=4187
  _globals['_BULK']._serialized_start=4190
  _globals['_BULK']._serialized_end=4409
  _globals['_BULKSPEC']._serialized_start=4412
  _globals['_BULKSPEC']._serialized_end=4782
  _globals['_BULKSTATUS']._serialized_start=4785
  _globals['_BULKSTATUS']._serialized_end=5161
  _globals['_BULKJOB']._serialized_start=5164
  _globals['_BULKJOB']._serialized_end=5388
  _globals['_BULKJOBSPEC']._serialized_start=5390
  _globals['_BULKJOBSPEC']._serialized_end=5474
  _globals['_LISTBULKJOBSREQUEST']._serialized_start=5476
  _globals['_LISTBULKJOBSREQUEST']._serialized_end=5595
  _globals['_LISTOFBULKJOB']._serialized_start=5597
  _globals['_LISTOFBULKJOB']._serialized_end=5712
  _globals['_CREATEPROXYBULKREQUEST']._serialized_start=5715
  _globals['_CREATEPROXYBULKREQUEST']._serialized_end=5885
  _globals['_PROXYBULKSPEC']._serialized_start=5888
  _globals['_PROXYBULKSPEC']._serialized_end=6237
  _globals['_PROXYBULK']._serialized_start=6240
  _globals['_PROXYBULK']._serialized_end=6469
  _globals['_SETMODEMPOOLREQUEST']._serialized_start=6472
  _globals['_SETMODEMPOOLREQUEST']._serialized_end=6645
  _globals['_LISTOFMODEMPOOL']._serialized_start=6647
  _globals['_LISTOFMODEMPOOL']._serialized_end=6766
  _globals['_MODEMPOOLSPEC']._serialized_start=6768
  _globals['_MODEMPOOLSPEC']._serialized_end=6783
  _globals['_MODEMPOOLSTATUS']._serialized_start=6785
  _globals['_MODEMPOOLSTATUS']._serialized_end=6873
  _globals['_MODEMPOOL']._serialized_start=6876
  _globals['_MODEMPOOL']._serialized_end=7110
  _globals['_SETMODEMREQUEST']._serialized_start=7112
  _globals['_SETMODEMREQUEST']._serialized_end=7223
  _globals['_SETDRIVER']._serialized_start=7225
  _globals['_SETDRIVER']._serialized_end=7316
  _globals['_LISTOFDRIVER']._serialized_start=7318
  _globals['_LISTOFDRIVER']._serialized_end=7431
  _globals['_DRIVER']._serialized_start=7434
  _globals['_DRIVER']._serialized_end=7590
  _globals['_DRIVERSPEC']._serialized_start=7593
  _globals['_DRIVERSPEC']._serialized_end=7957
  _globals['_DRIVERSTATUS']._serialized_start=7959
  _globals['_DRIVERSTATUS']._serialized_end=8043
  _globals['_CREATEVARIABLEREQUEST']._serialized_start=8046
  _globals['_CREATEVARIABLEREQUEST']._serialized_end=8214
  _globals['_LISTOFVARIABLE']._serialized_start=8216
  _globals['_LISTOFVARIABLE']._serialized_end=8333
  _globals['_VARIABLE']._serialized_start=8336
  _globals['_VARIABLE']._serialized_end=8497
  _globals['_VARIABLESPEC']._serialized_start=8500
  _globals['_VARIABLESPEC']._serialized_end=8666
  _globals['_ADDREGISTERTOVARIABLEREQUEST']._serialized_start=8668
  _globals['_ADDREGISTERTOVARIABLEREQUEST']._serialized_end=8764
  _globals['_REMOVEREGISTERFROMVARIABLEREQUEST']._serialized_start=8766
  _globals['_REMOVEREGISTERFROMVARIABLEREQUEST']._serialized_end=8867
  _globals['_CREATEDEVICECONFIGURATIONREGISTERREQUEST']._serialized_start=8870
  _globals['_CREATEDEVICECONFIGURATIONREGISTERREQUEST']._serialized_end=9076
  _globals['_LISTOFDEVICECONFIGURATIONREGISTER']._serialized_start=9079
  _globals['_LISTOFDEVICECONFIGURATIONREGISTER']._serialized_end=9234
  _globals['_DEVICECONFIGURATIONREGISTER']._serialized_start=9237
  _globals['_DEVICECONFIGURATIONREGISTER']._serialized_end=9436
  _globals['_CREATEDEVICECONFIGURATIONTEMPLATEREQUEST']._serialized_start=9439
  _globals['_CREATEDEVICECONFIGURATIONTEMPLATEREQUEST']._serialized_end=9645
  _globals['_LISTOFDEVICECONFIGURATIONTEMPLATE']._serialized_start=9648
  _globals['_LISTOFDEVICECONFIGURATIONTEMPLATE']._serialized_end=9803
  _globals['_DEVICECONFIGURATIONTEMPLATE']._serialized_start=9806
  _globals['_DEVICECONFIGURATIONTEMPLATE']._serialized_end=10005
  _globals['_DEVICECONFIGURATIONTEMPLATESPEC']._serialized_start=10007
  _globals['_DEVICECONFIGURATIONTEMPLATESPEC']._serialized_end=10106
  _globals['_ADDDEVICECONFIGURATIONREGISTERTODEVICECONFIGURATIONTEMPLATEREQUEST']._serialized_start=10108
  _globals['_ADDDEVICECONFIGURATIONREGISTERTODEVICECONFIGURATIONTEMPLATEREQUEST']._serialized_end=10232
  _globals['_REMOVEDEVICECONFIGURATIONREGISTERFROMDEVICECONFIGURATIONTEMPLATEREQUEST']._serialized_start=10235
  _globals['_REMOVEDEVICECONFIGURATIONREGISTERFROMDEVICECONFIGURATIONTEMPLATEREQUEST']._serialized_end=10364
  _globals['_GETDEVICEDATAREQUEST']._serialized_start=10367
  _globals['_GETDEVICEDATAREQUEST']._serialized_end=10703
  _globals['_DEVICEDATA']._serialized_start=10705
  _globals['_DEVICEDATA']._serialized_end=10797
  _globals['_DEVICEDEVICEDATA']._serialized_start=10799
  _globals['_DEVICEDEVICEDATA']._serialized_end=10922
  _globals['_VARIABLEDEVICEDATA']._serialized_start=10925
  _globals['_VARIABLEDEVICEDATA']._serialized_end=11131
  _globals['_GETDEVICEEVENTSREQUEST']._serialized_start=11133
  _globals['_GETDEVICEEVENTSREQUEST']._serialized_end=11249
  _globals['_CREATETIMEOFUSETABLEREQUEST']._serialized_start=11252
  _globals['_CREATETIMEOFUSETABLEREQUEST']._serialized_end=11442
  _globals['_LISTOFTIMEOFUSETABLE']._serialized_start=11445
  _globals['_LISTOFTIMEOFUSETABLE']._serialized_end=11574
  _globals['_TIMEOFUSETABLE']._serialized_start=11577
  _globals['_TIMEOFUSETABLE']._serialized_end=11760
# @@protoc_insertion_point(module_scope)
