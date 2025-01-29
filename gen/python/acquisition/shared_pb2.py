# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: acquisition/shared.proto
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
    'acquisition/shared.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from common import fields_pb2 as common_dot_fields__pb2
from common import metadata_pb2 as common_dot_metadata__pb2
from common import types_pb2 as common_dot_types__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x18\x61\x63quisition/shared.proto\x12\"io.clbs.openhes.models.acquisition\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x13\x63ommon/fields.proto\x1a\x15\x63ommon/metadata.proto\x1a\x12\x63ommon/types.proto\"\x96\x02\n\x0bJobSettings\x12!\n\x0cmax_duration\x18\x01 \x01(\x03R\x0bmaxDuration\x12K\n\x08priority\x18\x02 \x01(\x0e\x32/.io.clbs.openhes.models.acquisition.JobPriorityR\x08priority\x12\x1a\n\x08\x61ttempts\x18\x03 \x03(\x05R\x08\x61ttempts\x12\x1f\n\x0bretry_delay\x18\x04 \x01(\x03R\nretryDelay\x12\x1f\n\x0b\x64\x65\x66\x65r_start\x18\x05 \x01(\x04R\ndeferStart\x12\x39\n\nexpires_at\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\texpiresAt\"\x9a\x0e\n\tJobAction\x12\x1b\n\taction_id\x18\x01 \x01(\tR\x08\x61\x63tionId\x12]\n\nattributes\x18\x02 \x03(\x0b\x32=.io.clbs.openhes.models.acquisition.JobAction.AttributesEntryR\nattributes\x12Z\n\x0cget_register\x18\x03 \x01(\x0b\x32\x35.io.clbs.openhes.models.acquisition.ActionGetRegisterH\x00R\x0bgetRegister\x12v\n\x16get_periodical_profile\x18\x04 \x01(\x0b\x32>.io.clbs.openhes.models.acquisition.ActionGetPeriodicalProfileH\x00R\x14getPeriodicalProfile\x12s\n\x15get_irregular_profile\x18\x05 \x01(\x0b\x32=.io.clbs.openhes.models.acquisition.ActionGetIrregularProfileH\x00R\x13getIrregularProfile\x12T\n\nget_events\x18\x06 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.ActionGetEventsH\x00R\tgetEvents\x12\x61\n\x0fget_device_info\x18\x07 \x01(\x0b\x32\x37.io.clbs.openhes.models.acquisition.ActionGetDeviceInfoH\x00R\rgetDeviceInfo\x12T\n\nsync_clock\x18\x08 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.ActionSyncClockH\x00R\tsyncClock\x12\x61\n\x0fget_relay_state\x18\t \x01(\x0b\x32\x37.io.clbs.openhes.models.acquisition.ActionGetRelayStateH\x00R\rgetRelayState\x12\x61\n\x0fset_relay_state\x18\n \x01(\x0b\x32\x37.io.clbs.openhes.models.acquisition.ActionSetRelayStateH\x00R\rsetRelayState\x12v\n\x16get_disconnector_state\x18\x0b \x01(\x0b\x32>.io.clbs.openhes.models.acquisition.ActionGetDisconnectorStateH\x00R\x14getDisconnectorState\x12v\n\x16set_disconnector_state\x18\x0c \x01(\x0b\x32>.io.clbs.openhes.models.acquisition.ActionSetDisconnectorStateH\x00R\x14setDisconnectorState\x12K\n\x07get_tou\x18\r \x01(\x0b\x32\x30.io.clbs.openhes.models.acquisition.ActionGetTouH\x00R\x06getTou\x12K\n\x07set_tou\x18\x0e \x01(\x0b\x32\x30.io.clbs.openhes.models.acquisition.ActionSetTouH\x00R\x06setTou\x12W\n\x0bget_limiter\x18\x0f \x01(\x0b\x32\x34.io.clbs.openhes.models.acquisition.ActionGetLimiterH\x00R\ngetLimiter\x12W\n\x0bset_limiter\x18\x10 \x01(\x0b\x32\x34.io.clbs.openhes.models.acquisition.ActionSetLimiterH\x00R\nsetLimiter\x12p\n\x14reset_billing_period\x18\x11 \x01(\x0b\x32<.io.clbs.openhes.models.acquisition.ActionResetBillingPeriodH\x00R\x12resetBillingPeriod\x12Q\n\tfw_update\x18\x12 \x01(\x0b\x32\x32.io.clbs.openhes.models.acquisition.ActionFwUpdateH\x00R\x08\x66wUpdate\x1ah\n\x0f\x41ttributesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12?\n\x05value\x18\x02 \x01(\x0b\x32).io.clbs.openhes.models.common.FieldValueR\x05value:\x02\x38\x01\x42\x08\n\x06\x61\x63tion\"T\n\x0fListOfJobDevice\x12\x41\n\x04list\x18\x05 \x03(\x0b\x32-.io.clbs.openhes.models.acquisition.JobDeviceR\x04list\"X\n\x11ListOfJobDeviceId\x12\x43\n\x04list\x18\x05 \x03(\x0b\x32/.io.clbs.openhes.models.acquisition.JobDeviceIdR\x04list\"A\n\x0bJobDeviceId\x12\x15\n\x06job_id\x18\x01 \x01(\tR\x05jobId\x12\x1b\n\tdevice_id\x18\x02 \x01(\tR\x08\x64\x65viceId\"\x97\x04\n\tJobDevice\x12\x15\n\x06job_id\x18\x01 \x01(\tR\x05jobId\x12\x1b\n\tdevice_id\x18\x02 \x01(\tR\x08\x64\x65viceId\x12\x1f\n\x0b\x65xternal_id\x18\x03 \x01(\tR\nexternalId\x12p\n\x11\x64\x65vice_attributes\x18\x04 \x03(\x0b\x32\x43.io.clbs.openhes.models.acquisition.JobDevice.DeviceAttributesEntryR\x10\x64\x65viceAttributes\x12[\n\x0f\x63onnection_info\x18\x05 \x03(\x0b\x32\x32.io.clbs.openhes.models.acquisition.ConnectionInfoR\x0e\x63onnectionInfo\x12Z\n\x0c\x61pp_protocol\x18\x06 \x01(\x0e\x32\x37.io.clbs.openhes.models.acquisition.ApplicationProtocolR\x0b\x61ppProtocol\x12\x1a\n\x08timezone\x18\x07 \x01(\tR\x08timezone\x1an\n\x15\x44\x65viceAttributesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12?\n\x05value\x18\x02 \x01(\x0b\x32).io.clbs.openhes.models.common.FieldValueR\x05value:\x02\x38\x01\"\x87\x03\n\tModemInfo\x12\x19\n\x08modem_id\x18\x01 \x01(\tR\x07modemId\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\x12\x17\n\x07\x61t_init\x18\x03 \x01(\tR\x06\x61tInit\x12\x17\n\x07\x61t_test\x18\x04 \x01(\tR\x06\x61tTest\x12\x1b\n\tat_config\x18\x05 \x01(\tR\x08\x61tConfig\x12\x17\n\x07\x61t_dial\x18\x06 \x01(\tR\x06\x61tDial\x12\x1b\n\tat_hangup\x18\x07 \x01(\tR\x08\x61tHangup\x12\x1b\n\tat_escape\x18\x08 \x01(\tR\x08\x61tEscape\x12\x15\n\x06\x61t_dsr\x18\t \x01(\x08R\x05\x61tDsr\x12\'\n\x0f\x63onnect_timeout\x18\n \x01(\rR\x0e\x63onnectTimeout\x12U\n\x05tcpip\x18\x0b \x01(\x0b\x32=.io.clbs.openhes.models.acquisition.ConnectionTypeDirectTcpIpH\x00R\x05tcpipB\x12\n\x10modem_connection\"\xe3\x02\n\x0cSerialConfig\x12U\n\tbaud_rate\x18\x01 \x01(\x0e\x32\x38.io.clbs.openhes.models.acquisition.SerialConfigBaudRateR\x08\x62\x61udRate\x12N\n\x06parity\x18\x02 \x01(\x0e\x32\x36.io.clbs.openhes.models.acquisition.SerialConfigParityR\x06parity\x12U\n\tdata_bits\x18\x03 \x01(\x0e\x32\x38.io.clbs.openhes.models.acquisition.SerialConfigDataBitsR\x08\x64\x61taBits\x12U\n\tstop_bits\x18\x04 \x01(\x0e\x32\x38.io.clbs.openhes.models.acquisition.SerialConfigStopBitsR\x08stopBits\"\xc8\x03\n\x0e\x43onnectionInfo\x12U\n\x05tcpip\x18\x01 \x01(\x0b\x32=.io.clbs.openhes.models.acquisition.ConnectionTypeDirectTcpIpH\x00R\x05tcpip\x12\\\n\nmodem_pool\x18\x02 \x01(\x0b\x32;.io.clbs.openhes.models.acquisition.ConnectionTypeModemPoolH\x00R\tmodemPool\x12j\n\x0eserial_over_ip\x18\x03 \x01(\x0b\x32\x42.io.clbs.openhes.models.acquisition.ConnectionTypeControlledSerialH\x00R\x0cserialOverIp\x12Y\n\rlink_protocol\x18\x04 \x01(\x0e\x32\x34.io.clbs.openhes.models.acquisition.DataLinkProtocolR\x0clinkProtocol\x12,\n\x12\x63ustom_grouping_id\x18\x05 \x01(\tR\x10\x63ustomGroupingIdB\x0c\n\nconnection\"]\n\x19\x43onnectionTypeDirectTcpIp\x12\x12\n\x04host\x18\x01 \x01(\tR\x04host\x12\x12\n\x04port\x18\x02 \x01(\rR\x04port\x12\x18\n\x07timeout\x18\x03 \x01(\x05R\x07timeout\"\x8f\x01\n\x17\x43onnectionTypeModemPool\x12\x16\n\x06number\x18\x01 \x01(\tR\x06number\x12\x17\n\x07pool_id\x18\x02 \x01(\tR\x06poolId\x12\x43\n\x05modem\x18\x04 \x01(\x0b\x32-.io.clbs.openhes.models.acquisition.ModemInfoR\x05modem\"D\n\x1a\x43onnectionTypeSerialDirect\x12\x12\n\x04host\x18\x01 \x01(\tR\x04host\x12\x12\n\x04port\x18\x02 \x01(\rR\x04port\"l\n\x18\x43onnectionTypeSerialMoxa\x12\x12\n\x04host\x18\x01 \x01(\tR\x04host\x12\x1a\n\x08\x64\x61taPort\x18\x02 \x01(\rR\x08\x64\x61taPort\x12 \n\x0b\x63ommandPort\x18\x03 \x01(\rR\x0b\x63ommandPort\"\xd2\x01\n\x1b\x41pplicationProtocolTemplate\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12S\n\x08protocol\x18\x02 \x01(\x0e\x32\x37.io.clbs.openhes.models.acquisition.ApplicationProtocolR\x08protocol\x12N\n\nattributes\x18\x03 \x03(\x0b\x32..io.clbs.openhes.models.common.FieldDescriptorR\nattributes\"\xa2\x02\n\x10\x44\x61taLinkTemplate\x12Y\n\rlink_protocol\x18\x01 \x01(\x0e\x32\x34.io.clbs.openhes.models.acquisition.DataLinkProtocolR\x0clinkProtocol\x12\x63\n\x11\x61pp_protocol_refs\x18\x02 \x03(\x0e\x32\x37.io.clbs.openhes.models.acquisition.ApplicationProtocolR\x0f\x61ppProtocolRefs\x12N\n\nattributes\x18\x03 \x03(\x0b\x32..io.clbs.openhes.models.common.FieldDescriptorR\nattributes\"\xb6\x01\n\x15\x43ommunicationTemplate\x12I\n\x04type\x18\x01 \x01(\x0e\x32\x35.io.clbs.openhes.models.acquisition.CommunicationTypeR\x04type\x12R\n\tdatalinks\x18\x02 \x03(\x0b\x32\x34.io.clbs.openhes.models.acquisition.DataLinkTemplateR\tdatalinks\"9\n\x13\x41\x63\x63\x65ssLevelTemplate\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\"\x9d\x04\n\x0f\x44riverTemplates\x12r\n\x17\x63ommunication_templates\x18\x01 \x03(\x0b\x32\x39.io.clbs.openhes.models.acquisition.CommunicationTemplateR\x16\x63ommunicationTemplates\x12\x64\n\rapp_protocols\x18\x02 \x03(\x0b\x32?.io.clbs.openhes.models.acquisition.ApplicationProtocolTemplateR\x0c\x61ppProtocols\x12\x64\n\x11\x61\x63tion_attributes\x18\x03 \x03(\x0b\x32\x37.io.clbs.openhes.models.acquisition.JobActionAttributesR\x10\x61\x63tionAttributes\x12\x62\n\x10\x61\x63\x63\x65ss_templates\x18\x04 \x03(\x0b\x32\x37.io.clbs.openhes.models.acquisition.AccessLevelTemplateR\x0f\x61\x63\x63\x65ssTemplates\x12\x66\n\x12\x61\x63tion_constraints\x18\x05 \x01(\x0b\x32\x37.io.clbs.openhes.models.acquisition.JobActionContraintsR\x11\x61\x63tionConstraints\"\xd8\x01\n\x14\x41\x63tionProgressUpdate\x12\x15\n\x06job_id\x18\x01 \x01(\tR\x05jobId\x12\x1b\n\taction_id\x18\x02 \x01(\tR\x08\x61\x63tionId\x12H\n\x04\x63ode\x18\x03 \x01(\x0e\x32\x34.io.clbs.openhes.models.acquisition.ActionResultCodeR\x04\x63ode\x12\x42\n\x04\x64\x61ta\x18\x04 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.ActionDataR\x04\x64\x61ta\"\x8c\x01\n\x11JobProgressUpdate\x12\x15\n\x06job_id\x18\x01 \x01(\tR\x05jobId\x12\x44\n\x04\x63ode\x18\x02 \x01(\x0e\x32\x30.io.clbs.openhes.models.acquisition.JobErrorCodeR\x04\x63ode\x12\x1a\n\x08\x64uration\x18\x03 \x01(\x03R\x08\x64uration\"\xc8\x02\n\nActionData\x12\x30\n\x06nodata\x18\x01 \x01(\x0b\x32\x16.google.protobuf.EmptyH\x00R\x06nodata\x12O\n\x08\x62illings\x18\x02 \x01(\x0b\x32\x31.io.clbs.openhes.models.acquisition.BillingValuesH\x00R\x08\x62illings\x12M\n\x07profile\x18\x03 \x01(\x0b\x32\x31.io.clbs.openhes.models.acquisition.ProfileValuesH\x00R\x07profile\x12`\n\x11irregular_profile\x18\x04 \x01(\x0b\x32\x31.io.clbs.openhes.models.acquisition.BillingValuesH\x00R\x10irregularProfileB\x06\n\x04\x64\x61ta\"\x84\x01\n\rProfileValues\x12\x16\n\x06period\x18\x01 \x01(\x05R\x06period\x12\x12\n\x04unit\x18\x02 \x01(\tR\x04unit\x12G\n\x06\x62locks\x18\x03 \x03(\x0b\x32/.io.clbs.openhes.models.acquisition.ProfileBlokR\x06\x62locks\"\x9d\x01\n\x0bProfileBlok\x12\x43\n\x0fstart_timestamp\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x0estartTimestamp\x12I\n\x06values\x18\x02 \x03(\x0b\x32\x31.io.clbs.openhes.models.acquisition.MeasuredValueR\x06values\"Y\n\rBillingValues\x12H\n\x06values\x18\x01 \x03(\x0b\x32\x30.io.clbs.openhes.models.acquisition.BillingValueR\x06values\"\xa5\x01\n\x0c\x42illingValue\x12\x38\n\ttimestamp\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\ttimestamp\x12\x12\n\x04unit\x18\x02 \x01(\tR\x04unit\x12G\n\x05value\x18\x03 \x01(\x0b\x32\x31.io.clbs.openhes.models.acquisition.MeasuredValueR\x05value\"\xd4\x02\n\rMeasuredValue\x12\x16\n\x06status\x18\x01 \x01(\x03R\x06status\x12\x1a\n\x08\x65xponent\x18\x02 \x01(\x05R\x08\x65xponent\x12#\n\x0c\x64ouble_value\x18\x03 \x01(\x01H\x00R\x0b\x64oubleValue\x12%\n\rinteger_value\x18\x04 \x01(\x03H\x00R\x0cintegerValue\x12#\n\x0cstring_value\x18\x05 \x01(\tH\x00R\x0bstringValue\x12\x45\n\x0ftimestamp_value\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampH\x00R\x0etimestampValue\x12.\n\x12timestamp_tz_value\x18\x07 \x01(\tH\x00R\x10timestampTzValue\x12\x1f\n\nbool_value\x18\x08 \x01(\x08H\x00R\tboolValueB\x06\n\x04kind\"\xa9\x01\n\x13JobActionAttributes\x12\x42\n\x04type\x18\x01 \x01(\x0e\x32..io.clbs.openhes.models.acquisition.ActionTypeR\x04type\x12N\n\nattributes\x18\x02 \x03(\x0b\x32..io.clbs.openhes.models.common.FieldDescriptorR\nattributes\"\xdb\x01\n\x1e\x43onnectionTypeControlledSerial\x12X\n\x06\x64irect\x18\x01 \x01(\x0b\x32>.io.clbs.openhes.models.acquisition.ConnectionTypeSerialDirectH\x00R\x06\x64irect\x12R\n\x04moxa\x18\x02 \x01(\x0b\x32<.io.clbs.openhes.models.acquisition.ConnectionTypeSerialMoxaH\x00R\x04moxaB\x0b\n\tconverter\"\x13\n\x11\x41\x63tionGetRegister\"x\n\x1a\x41\x63tionGetPeriodicalProfile\x12.\n\x04\x66rom\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x04\x66rom\x12*\n\x02to\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x02to\"w\n\x19\x41\x63tionGetIrregularProfile\x12.\n\x04\x66rom\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x04\x66rom\x12*\n\x02to\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x02to\"m\n\x0f\x41\x63tionGetEvents\x12.\n\x04\x66rom\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x04\x66rom\x12*\n\x02to\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x02to\"\xde\x02\n\x13\x41\x63tionGetDeviceInfo\x12\x41\n\x0einfo_timestamp\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\rinfoTimestamp\x12<\n\x1amanufacturer_serial_number\x18\x02 \x01(\tR\x18manufacturerSerialNumber\x12\x30\n\x14\x64\x65vice_serial_number\x18\x03 \x01(\tR\x12\x64\x65viceSerialNumber\x12)\n\x10\x66irmware_version\x18\x04 \x01(\tR\x0f\x66irmwareVersion\x12\x1f\n\x0b\x63lock_delta\x18\x05 \x01(\x01R\nclockDelta\x12!\n\x0c\x64\x65vice_model\x18\x06 \x01(\tR\x0b\x64\x65viceModel\x12%\n\x0e\x65rror_register\x18\x07 \x01(\rR\rerrorRegister\"\x11\n\x0f\x41\x63tionSyncClock\"\x15\n\x13\x41\x63tionGetRelayState\"\x15\n\x13\x41\x63tionSetRelayState\"\x1c\n\x1a\x41\x63tionGetDisconnectorState\"\x1c\n\x1a\x41\x63tionSetDisconnectorState\"\x0e\n\x0c\x41\x63tionGetTou\"\x0e\n\x0c\x41\x63tionSetTou\"\x12\n\x10\x41\x63tionGetLimiter\"\x12\n\x10\x41\x63tionSetLimiter\"\x1a\n\x18\x41\x63tionResetBillingPeriod\"\x10\n\x0e\x41\x63tionFwUpdate\"\xfa\x03\n\x13JobActionContraints\x12\x85\x01\n\x16get_register_type_name\x18\x01 \x03(\x0b\x32P.io.clbs.openhes.models.acquisition.JobActionContraints.GetRegisterTypeNameEntryR\x13getRegisterTypeName\x12\x97\x01\n\x1cget_register_type_attributes\x18\x02 \x03(\x0b\x32V.io.clbs.openhes.models.acquisition.JobActionContraints.GetRegisterTypeAttributesEntryR\x19getRegisterTypeAttributes\x1a\x46\n\x18GetRegisterTypeNameEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\x1ay\n\x1eGetRegisterTypeAttributesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x41\n\x05value\x18\x02 \x01(\x0b\x32+.io.clbs.openhes.models.common.ListOfStringR\x05value:\x02\x38\x01\"g\n\x07\x42ulkJob\x12\x15\n\x06job_id\x18\x01 \x01(\tR\x05jobId\x12\x45\n\x06status\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.acquisition.JobStatusR\x06status\"\xc0\x03\n\tJobStatus\x12I\n\x06status\x18\x01 \x01(\x0e\x32\x31.io.clbs.openhes.models.acquisition.JobStatusCodeR\x06status\x12\x44\n\x04\x63ode\x18\x02 \x01(\x0e\x32\x30.io.clbs.openhes.models.acquisition.JobErrorCodeR\x04\x63ode\x12J\n\x07results\x18\x03 \x03(\x0b\x32\x30.io.clbs.openhes.models.acquisition.ActionResultR\x07results\x12\x39\n\ncreated_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x39\n\nstarted_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tstartedAt\x12;\n\x0b\x66inished_at\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\nfinishedAt\x12#\n\rattempts_done\x18\x07 \x01(\x05R\x0c\x61ttemptsDone\"\xa6\x04\n\x0cStartJobData\x12s\n\x11\x64\x65vice_attributes\x18\x01 \x03(\x0b\x32\x46.io.clbs.openhes.models.acquisition.StartJobData.DeviceAttributesEntryR\x10\x64\x65viceAttributes\x12\x15\n\x06job_id\x18\x02 \x01(\tR\x05jobId\x12R\n\x0cjob_settings\x18\x03 \x01(\x0b\x32/.io.clbs.openhes.models.acquisition.JobSettingsR\x0bjobSettings\x12N\n\x0bjob_actions\x18\x04 \x03(\x0b\x32-.io.clbs.openhes.models.acquisition.JobActionR\njobActions\x12Z\n\x0c\x61pp_protocol\x18\x05 \x01(\x0e\x32\x37.io.clbs.openhes.models.acquisition.ApplicationProtocolR\x0b\x61ppProtocol\x12\x1a\n\x08timezone\x18\x06 \x01(\tR\x08timezone\x1an\n\x15\x44\x65viceAttributesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12?\n\x05value\x18\x02 \x01(\x0b\x32).io.clbs.openhes.models.common.FieldValueR\x05value:\x02\x38\x01\")\n\x10\x43\x61ncelJobRequest\x12\x15\n\x06job_id\x18\x01 \x01(\tR\x05jobId\"\xc2\x03\n\x14\x44\x65viceConnectionInfo\x12\x61\n\x12\x63ommunication_unit\x18\x01 \x01(\x0b\x32\x32.io.clbs.openhes.models.acquisition.ConnectionInfoR\x11\x63ommunicationUnit\x12Z\n\x0c\x61pp_protocol\x18\x02 \x01(\x0e\x32\x37.io.clbs.openhes.models.acquisition.ApplicationProtocolR\x0b\x61ppProtocol\x12{\n\x11\x64\x65vice_attributes\x18\x03 \x03(\x0b\x32N.io.clbs.openhes.models.acquisition.DeviceConnectionInfo.DeviceAttributesEntryR\x10\x64\x65viceAttributes\x1an\n\x15\x44\x65viceAttributesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12?\n\x05value\x18\x02 \x01(\x0b\x32).io.clbs.openhes.models.common.FieldValueR\x05value:\x02\x38\x01\"V\n\x0fListOfModemInfo\x12\x43\n\x05items\x18\x01 \x03(\x0b\x32-.io.clbs.openhes.models.acquisition.ModemInfoR\x05items\"f\n\x14ListOfConnectionInfo\x12N\n\x05items\x18\x01 \x03(\x0b\x32\x38.io.clbs.openhes.models.acquisition.DeviceConnectionInfoR\x05items\"\xbd\x01\n\x0c\x41\x63tionResult\x12\x1b\n\taction_id\x18\x01 \x01(\tR\x08\x61\x63tionId\x12L\n\x06status\x18\x02 \x01(\x0e\x32\x34.io.clbs.openhes.models.acquisition.ActionResultCodeR\x06status\x12\x42\n\x04\x64\x61ta\x18\x03 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.ActionDataR\x04\x64\x61ta\"%\n\x0cJobEventData\x12\x15\n\x06job_id\x18\x01 \x01(\x0cR\x05jobId\"\x88\x03\n\nDeviceSpec\x12\x1f\n\x0b\x65xternal_id\x18\x02 \x01(\tR\nexternalId\x12^\n\nattributes\x18\x04 \x03(\x0b\x32>.io.clbs.openhes.models.acquisition.DeviceSpec.AttributesEntryR\nattributes\x12s\n\x17\x63ommunication_unit_link\x18\x05 \x03(\x0b\x32;.io.clbs.openhes.models.acquisition.DeviceCommunicationUnitR\x15\x63ommunicationUnitLink\x12\x1a\n\x08timezone\x18\x06 \x01(\tR\x08timezone\x1ah\n\x0f\x41ttributesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12?\n\x05value\x18\x02 \x01(\x0b\x32).io.clbs.openhes.models.common.FieldValueR\x05value:\x02\x38\x01\"\xa9\x01\n\x17\x44\x65viceCommunicationUnit\x12\x32\n\x15\x63ommunication_unit_id\x18\x01 \x01(\tR\x13\x63ommunicationUnitId\x12Z\n\x0c\x61pp_protocol\x18\x02 \x01(\x0e\x32\x37.io.clbs.openhes.models.acquisition.ApplicationProtocolR\x0b\x61ppProtocol\"\x95\x02\n\x07JobSpec\x12\x45\n\x06\x64\x65vice\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.acquisition.JobDeviceR\x06\x64\x65vice\x12R\n\x0cjob_settings\x18\x03 \x01(\x0b\x32/.io.clbs.openhes.models.acquisition.JobSettingsR\x0bjobSettings\x12\x1f\n\x0b\x64river_type\x18\x04 \x01(\tR\ndriverType\x12N\n\x0bjob_actions\x18\x05 \x03(\x0b\x32-.io.clbs.openhes.models.acquisition.JobActionR\njobActions\"\x9d\x01\n\x13JobDoneNotification\x12?\n\x04spec\x18\x01 \x01(\x0b\x32+.io.clbs.openhes.models.acquisition.JobSpecR\x04spec\x12\x45\n\x06status\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.acquisition.JobStatusR\x06status\"G\n\nDriverInfo\x12\x1f\n\x0b\x64river_type\x18\x01 \x01(\tR\ndriverType\x12\x18\n\x07version\x18\x02 \x01(\tR\x07version\"\xba\x01\n\x11\x43ommunicationUnit\x12T\n\x04spec\x18\x01 \x01(\x0b\x32\x39.io.clbs.openhes.models.acquisition.CommunicationUnitSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadataJ\x04\x08\x02\x10\x03\"\x95\x01\n\x15\x43ommunicationUnitSpec\x12\x1f\n\x0b\x65xternal_id\x18\x01 \x01(\tR\nexternalId\x12[\n\x0f\x63onnection_info\x18\x02 \x01(\x0b\x32\x32.io.clbs.openhes.models.acquisition.ConnectionInfoR\x0e\x63onnectionInfo*\xc1\x01\n\x0bJobPriority\x12\x12\n\x0eJOB_PRIORITY_0\x10\x00\x12\x12\n\x0eJOB_PRIORITY_1\x10\x01\x12\x12\n\x0eJOB_PRIORITY_2\x10\x02\x12\x12\n\x0eJOB_PRIORITY_3\x10\x03\x12\x12\n\x0eJOB_PRIORITY_4\x10\x04\x12\x12\n\x0eJOB_PRIORITY_5\x10\x05\x12\x12\n\x0eJOB_PRIORITY_6\x10\x06\x12\x12\n\x0eJOB_PRIORITY_7\x10\x07\x12\x12\n\x0eJOB_PRIORITY_8\x10\x08*i\n\x12SerialConfigParity\x12\x0f\n\x0bPARITY_NONE\x10\x00\x12\x0f\n\x0bPARITY_EVEN\x10\x01\x12\x0e\n\nPARITY_ODD\x10\x02\x12\x0f\n\x0bPARITY_MARK\x10\x03\x12\x10\n\x0cPARITY_SPACE\x10\x04*\xcb\x02\n\x14SerialConfigBaudRate\x12\x11\n\rBAUD_RATE_110\x10\x00\x12\x11\n\rBAUD_RATE_300\x10\x01\x12\x11\n\rBAUD_RATE_600\x10\x02\x12\x12\n\x0e\x42\x41UD_RATE_1200\x10\x03\x12\x12\n\x0e\x42\x41UD_RATE_2400\x10\x04\x12\x12\n\x0e\x42\x41UD_RATE_4800\x10\x05\x12\x12\n\x0e\x42\x41UD_RATE_9600\x10\x06\x12\x13\n\x0f\x42\x41UD_RATE_14400\x10\x07\x12\x13\n\x0f\x42\x41UD_RATE_19200\x10\x08\x12\x13\n\x0f\x42\x41UD_RATE_38400\x10\t\x12\x13\n\x0f\x42\x41UD_RATE_57600\x10\n\x12\x14\n\x10\x42\x41UD_RATE_115200\x10\x0b\x12\x14\n\x10\x42\x41UD_RATE_230400\x10\x0c\x12\x14\n\x10\x42\x41UD_RATE_460800\x10\r\x12\x14\n\x10\x42\x41UD_RATE_921600\x10\x0e*Z\n\x14SerialConfigDataBits\x12\x0f\n\x0b\x44\x41TA_BITS_5\x10\x00\x12\x0f\n\x0b\x44\x41TA_BITS_6\x10\x01\x12\x0f\n\x0b\x44\x41TA_BITS_7\x10\x02\x12\x0f\n\x0b\x44\x41TA_BITS_8\x10\x03*K\n\x14SerialConfigStopBits\x12\x0f\n\x0bSTOP_BITS_1\x10\x00\x12\x11\n\rSTOP_BITS_1_5\x10\x01\x12\x0f\n\x0bSTOP_BITS_2\x10\x02*\xca\x01\n\x11\x43ommunicationType\x12\x1c\n\x18\x43OMMUNICATION_TYPE_TCPIP\x10\x00\x12!\n\x1d\x43OMMUNICATION_TYPE_MODEM_POOL\x10\x01\x12)\n%COMMUNICATION_TYPE_SERIAL_LINE_DIRECT\x10\x02\x12\'\n#COMMUNICATION_TYPE_SERIAL_LINE_MOXA\x10\x03\x12 \n\x1c\x43OMMUNICATION_TYPE_LISTENING\x10\x63*\xa7\x01\n\x10\x44\x61taLinkProtocol\x12\x1a\n\x16LINKPROTO_IEC_62056_21\x10\x00\x12\x12\n\x0eLINKPROTO_HDLC\x10\x01\x12\x1b\n\x17LINKPROTO_COSEM_WRAPPER\x10\x02\x12\x14\n\x10LINKPROTO_MODBUS\x10\x03\x12\x12\n\x0eLINKPROTO_MBUS\x10\x04\x12\x1c\n\x18LINKPROTO_NOT_APPLICABLE\x10\x63*\xae\x01\n\x13\x41pplicationProtocol\x12\x19\n\x15\x41PPPROTO_IEC_62056_21\x10\x00\x12\x14\n\x10\x41PPPROTO_DLMS_SN\x10\x01\x12\x14\n\x10\x41PPPROTO_DLMS_LN\x10\x02\x12\x11\n\rAPPPROTO_SCTM\x10\x03\x12\x13\n\x0f\x41PPPROTO_LIS200\x10\x04\x12\x15\n\x11\x41PPPROTO_ANSI_C12\x10\x05\x12\x11\n\rAPPPROTO_MQTT\x10\x06*\x91\x04\n\nActionType\x12\x1c\n\x18\x41\x43TION_TYPE_GET_REGISTER\x10\x00\x12&\n\"ACTION_TYPE_GET_PERIODICAL_PROFILE\x10\x01\x12%\n!ACTION_TYPE_GET_IRREGULAR_PROFILE\x10\x02\x12\x1a\n\x16\x41\x43TION_TYPE_GET_EVENTS\x10\x03\x12\x1f\n\x1b\x41\x43TION_TYPE_GET_DEVICE_INFO\x10\n\x12\x1a\n\x16\x41\x43TION_TYPE_SYNC_CLOCK\x10\x0b\x12\x1f\n\x1b\x41\x43TION_TYPE_GET_RELAY_STATE\x10\x14\x12\x1f\n\x1b\x41\x43TION_TYPE_SET_RELAY_STATE\x10\x15\x12&\n\"ACTION_TYPE_GET_DISCONNECTOR_STATE\x10\x16\x12&\n\"ACTION_TYPE_SET_DISCONNECTOR_STATE\x10\x17\x12\x17\n\x13\x41\x43TION_TYPE_GET_TOU\x10\x18\x12\x17\n\x13\x41\x43TION_TYPE_SET_TOU\x10\x19\x12\x1b\n\x17\x41\x43TION_TYPE_GET_LIMITER\x10\x1a\x12\x1b\n\x17\x41\x43TION_TYPE_SET_LIMITER\x10\x1b\x12$\n ACTION_TYPE_RESET_BILLING_PERIOD\x10(\x12\x19\n\x15\x41\x43TION_TYPE_FW_UPDATE\x10\x32*\x8b\x01\n\x10\x41\x63tionResultCode\x12\x18\n\x14\x45RROR_CODE_ACTION_OK\x10\x00\x12!\n\x1d\x45RROR_CODE_ACTION_UNSUPPORTED\x10\x01\x12\x1d\n\x19\x45RROR_CODE_ACTION_PENDING\x10\x03\x12\x1b\n\x17\x45RROR_CODE_ACTION_ERROR\x10\x05*\x97\x01\n\x0cJobErrorCode\x12\x17\n\x13JOB_ERROR_CODE_NONE\x10\x00\x12\x17\n\x13JOB_ERROR_CODE_BUSY\x10\x01\x12\x18\n\x14JOB_ERROR_CODE_ERROR\x10\x05\x12!\n\x1dJOB_ERROR_CODE_ALREADY_EXISTS\x10\x08\x12\x18\n\x14JOB_ERROR_CODE_FATAL\x10\t*\x90\x01\n\x0e\x42ulkStatusCode\x12\x16\n\x12\x42ULK_STATUS_QUEUED\x10\x00\x12\x17\n\x13\x42ULK_STATUS_RUNNING\x10\x01\x12\x19\n\x15\x42ULK_STATUS_COMPLETED\x10\x02\x12\x19\n\x15\x42ULK_STATUS_CANCELLED\x10\x03\x12\x17\n\x13\x42ULK_STATUS_EXPIRED\x10\x04*\xa1\x01\n\rJobStatusCode\x12\x15\n\x11JOB_STATUS_QUEUED\x10\x00\x12\x16\n\x12JOB_STATUS_RUNNING\x10\x01\x12\x18\n\x14JOB_STATUS_COMPLETED\x10\x02\x12\x15\n\x11JOB_STATUS_FAILED\x10\x03\x12\x18\n\x14JOB_STATUS_CANCELLED\x10\x04\x12\x16\n\x12JOB_STATUS_EXPIRED\x10\x05\x42\x35Z3github.com/cybroslabs/hes-2-apis/gen/go/acquisitionb\x08\x65\x64itionsp\xe8\x07')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'acquisition.shared_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z3github.com/cybroslabs/hes-2-apis/gen/go/acquisition'
  _globals['_JOBACTION_ATTRIBUTESENTRY']._loaded_options = None
  _globals['_JOBACTION_ATTRIBUTESENTRY']._serialized_options = b'8\001'
  _globals['_JOBDEVICE_DEVICEATTRIBUTESENTRY']._loaded_options = None
  _globals['_JOBDEVICE_DEVICEATTRIBUTESENTRY']._serialized_options = b'8\001'
  _globals['_JOBACTIONCONTRAINTS_GETREGISTERTYPENAMEENTRY']._loaded_options = None
  _globals['_JOBACTIONCONTRAINTS_GETREGISTERTYPENAMEENTRY']._serialized_options = b'8\001'
  _globals['_JOBACTIONCONTRAINTS_GETREGISTERTYPEATTRIBUTESENTRY']._loaded_options = None
  _globals['_JOBACTIONCONTRAINTS_GETREGISTERTYPEATTRIBUTESENTRY']._serialized_options = b'8\001'
  _globals['_STARTJOBDATA_DEVICEATTRIBUTESENTRY']._loaded_options = None
  _globals['_STARTJOBDATA_DEVICEATTRIBUTESENTRY']._serialized_options = b'8\001'
  _globals['_DEVICECONNECTIONINFO_DEVICEATTRIBUTESENTRY']._loaded_options = None
  _globals['_DEVICECONNECTIONINFO_DEVICEATTRIBUTESENTRY']._serialized_options = b'8\001'
  _globals['_DEVICESPEC_ATTRIBUTESENTRY']._loaded_options = None
  _globals['_DEVICESPEC_ATTRIBUTESENTRY']._serialized_options = b'8\001'
  _globals['_COMMUNICATIONUNIT'].fields_by_name['spec']._loaded_options = None
  _globals['_COMMUNICATIONUNIT'].fields_by_name['spec']._serialized_options = b'\252\001\002\010\003'
  _globals['_JOBPRIORITY']._serialized_start=12943
  _globals['_JOBPRIORITY']._serialized_end=13136
  _globals['_SERIALCONFIGPARITY']._serialized_start=13138
  _globals['_SERIALCONFIGPARITY']._serialized_end=13243
  _globals['_SERIALCONFIGBAUDRATE']._serialized_start=13246
  _globals['_SERIALCONFIGBAUDRATE']._serialized_end=13577
  _globals['_SERIALCONFIGDATABITS']._serialized_start=13579
  _globals['_SERIALCONFIGDATABITS']._serialized_end=13669
  _globals['_SERIALCONFIGSTOPBITS']._serialized_start=13671
  _globals['_SERIALCONFIGSTOPBITS']._serialized_end=13746
  _globals['_COMMUNICATIONTYPE']._serialized_start=13749
  _globals['_COMMUNICATIONTYPE']._serialized_end=13951
  _globals['_DATALINKPROTOCOL']._serialized_start=13954
  _globals['_DATALINKPROTOCOL']._serialized_end=14121
  _globals['_APPLICATIONPROTOCOL']._serialized_start=14124
  _globals['_APPLICATIONPROTOCOL']._serialized_end=14298
  _globals['_ACTIONTYPE']._serialized_start=14301
  _globals['_ACTIONTYPE']._serialized_end=14830
  _globals['_ACTIONRESULTCODE']._serialized_start=14833
  _globals['_ACTIONRESULTCODE']._serialized_end=14972
  _globals['_JOBERRORCODE']._serialized_start=14975
  _globals['_JOBERRORCODE']._serialized_end=15126
  _globals['_BULKSTATUSCODE']._serialized_start=15129
  _globals['_BULKSTATUSCODE']._serialized_end=15273
  _globals['_JOBSTATUSCODE']._serialized_start=15276
  _globals['_JOBSTATUSCODE']._serialized_end=15437
  _globals['_JOBSETTINGS']._serialized_start=221
  _globals['_JOBSETTINGS']._serialized_end=499
  _globals['_JOBACTION']._serialized_start=502
  _globals['_JOBACTION']._serialized_end=2320
  _globals['_JOBACTION_ATTRIBUTESENTRY']._serialized_start=2206
  _globals['_JOBACTION_ATTRIBUTESENTRY']._serialized_end=2310
  _globals['_LISTOFJOBDEVICE']._serialized_start=2322
  _globals['_LISTOFJOBDEVICE']._serialized_end=2406
  _globals['_LISTOFJOBDEVICEID']._serialized_start=2408
  _globals['_LISTOFJOBDEVICEID']._serialized_end=2496
  _globals['_JOBDEVICEID']._serialized_start=2498
  _globals['_JOBDEVICEID']._serialized_end=2563
  _globals['_JOBDEVICE']._serialized_start=2566
  _globals['_JOBDEVICE']._serialized_end=3101
  _globals['_JOBDEVICE_DEVICEATTRIBUTESENTRY']._serialized_start=2991
  _globals['_JOBDEVICE_DEVICEATTRIBUTESENTRY']._serialized_end=3101
  _globals['_MODEMINFO']._serialized_start=3104
  _globals['_MODEMINFO']._serialized_end=3495
  _globals['_SERIALCONFIG']._serialized_start=3498
  _globals['_SERIALCONFIG']._serialized_end=3853
  _globals['_CONNECTIONINFO']._serialized_start=3856
  _globals['_CONNECTIONINFO']._serialized_end=4312
  _globals['_CONNECTIONTYPEDIRECTTCPIP']._serialized_start=4314
  _globals['_CONNECTIONTYPEDIRECTTCPIP']._serialized_end=4407
  _globals['_CONNECTIONTYPEMODEMPOOL']._serialized_start=4410
  _globals['_CONNECTIONTYPEMODEMPOOL']._serialized_end=4553
  _globals['_CONNECTIONTYPESERIALDIRECT']._serialized_start=4555
  _globals['_CONNECTIONTYPESERIALDIRECT']._serialized_end=4623
  _globals['_CONNECTIONTYPESERIALMOXA']._serialized_start=4625
  _globals['_CONNECTIONTYPESERIALMOXA']._serialized_end=4733
  _globals['_APPLICATIONPROTOCOLTEMPLATE']._serialized_start=4736
  _globals['_APPLICATIONPROTOCOLTEMPLATE']._serialized_end=4946
  _globals['_DATALINKTEMPLATE']._serialized_start=4949
  _globals['_DATALINKTEMPLATE']._serialized_end=5239
  _globals['_COMMUNICATIONTEMPLATE']._serialized_start=5242
  _globals['_COMMUNICATIONTEMPLATE']._serialized_end=5424
  _globals['_ACCESSLEVELTEMPLATE']._serialized_start=5426
  _globals['_ACCESSLEVELTEMPLATE']._serialized_end=5483
  _globals['_DRIVERTEMPLATES']._serialized_start=5486
  _globals['_DRIVERTEMPLATES']._serialized_end=6027
  _globals['_ACTIONPROGRESSUPDATE']._serialized_start=6030
  _globals['_ACTIONPROGRESSUPDATE']._serialized_end=6246
  _globals['_JOBPROGRESSUPDATE']._serialized_start=6249
  _globals['_JOBPROGRESSUPDATE']._serialized_end=6389
  _globals['_ACTIONDATA']._serialized_start=6392
  _globals['_ACTIONDATA']._serialized_end=6720
  _globals['_PROFILEVALUES']._serialized_start=6723
  _globals['_PROFILEVALUES']._serialized_end=6855
  _globals['_PROFILEBLOK']._serialized_start=6858
  _globals['_PROFILEBLOK']._serialized_end=7015
  _globals['_BILLINGVALUES']._serialized_start=7017
  _globals['_BILLINGVALUES']._serialized_end=7106
  _globals['_BILLINGVALUE']._serialized_start=7109
  _globals['_BILLINGVALUE']._serialized_end=7274
  _globals['_MEASUREDVALUE']._serialized_start=7277
  _globals['_MEASUREDVALUE']._serialized_end=7617
  _globals['_JOBACTIONATTRIBUTES']._serialized_start=7620
  _globals['_JOBACTIONATTRIBUTES']._serialized_end=7789
  _globals['_CONNECTIONTYPECONTROLLEDSERIAL']._serialized_start=7792
  _globals['_CONNECTIONTYPECONTROLLEDSERIAL']._serialized_end=8011
  _globals['_ACTIONGETREGISTER']._serialized_start=8013
  _globals['_ACTIONGETREGISTER']._serialized_end=8032
  _globals['_ACTIONGETPERIODICALPROFILE']._serialized_start=8034
  _globals['_ACTIONGETPERIODICALPROFILE']._serialized_end=8154
  _globals['_ACTIONGETIRREGULARPROFILE']._serialized_start=8156
  _globals['_ACTIONGETIRREGULARPROFILE']._serialized_end=8275
  _globals['_ACTIONGETEVENTS']._serialized_start=8277
  _globals['_ACTIONGETEVENTS']._serialized_end=8386
  _globals['_ACTIONGETDEVICEINFO']._serialized_start=8389
  _globals['_ACTIONGETDEVICEINFO']._serialized_end=8739
  _globals['_ACTIONSYNCCLOCK']._serialized_start=8741
  _globals['_ACTIONSYNCCLOCK']._serialized_end=8758
  _globals['_ACTIONGETRELAYSTATE']._serialized_start=8760
  _globals['_ACTIONGETRELAYSTATE']._serialized_end=8781
  _globals['_ACTIONSETRELAYSTATE']._serialized_start=8783
  _globals['_ACTIONSETRELAYSTATE']._serialized_end=8804
  _globals['_ACTIONGETDISCONNECTORSTATE']._serialized_start=8806
  _globals['_ACTIONGETDISCONNECTORSTATE']._serialized_end=8834
  _globals['_ACTIONSETDISCONNECTORSTATE']._serialized_start=8836
  _globals['_ACTIONSETDISCONNECTORSTATE']._serialized_end=8864
  _globals['_ACTIONGETTOU']._serialized_start=8866
  _globals['_ACTIONGETTOU']._serialized_end=8880
  _globals['_ACTIONSETTOU']._serialized_start=8882
  _globals['_ACTIONSETTOU']._serialized_end=8896
  _globals['_ACTIONGETLIMITER']._serialized_start=8898
  _globals['_ACTIONGETLIMITER']._serialized_end=8916
  _globals['_ACTIONSETLIMITER']._serialized_start=8918
  _globals['_ACTIONSETLIMITER']._serialized_end=8936
  _globals['_ACTIONRESETBILLINGPERIOD']._serialized_start=8938
  _globals['_ACTIONRESETBILLINGPERIOD']._serialized_end=8964
  _globals['_ACTIONFWUPDATE']._serialized_start=8966
  _globals['_ACTIONFWUPDATE']._serialized_end=8982
  _globals['_JOBACTIONCONTRAINTS']._serialized_start=8985
  _globals['_JOBACTIONCONTRAINTS']._serialized_end=9491
  _globals['_JOBACTIONCONTRAINTS_GETREGISTERTYPENAMEENTRY']._serialized_start=9298
  _globals['_JOBACTIONCONTRAINTS_GETREGISTERTYPENAMEENTRY']._serialized_end=9368
  _globals['_JOBACTIONCONTRAINTS_GETREGISTERTYPEATTRIBUTESENTRY']._serialized_start=9370
  _globals['_JOBACTIONCONTRAINTS_GETREGISTERTYPEATTRIBUTESENTRY']._serialized_end=9491
  _globals['_BULKJOB']._serialized_start=9493
  _globals['_BULKJOB']._serialized_end=9596
  _globals['_JOBSTATUS']._serialized_start=9599
  _globals['_JOBSTATUS']._serialized_end=10047
  _globals['_STARTJOBDATA']._serialized_start=10050
  _globals['_STARTJOBDATA']._serialized_end=10600
  _globals['_STARTJOBDATA_DEVICEATTRIBUTESENTRY']._serialized_start=2991
  _globals['_STARTJOBDATA_DEVICEATTRIBUTESENTRY']._serialized_end=3101
  _globals['_CANCELJOBREQUEST']._serialized_start=10602
  _globals['_CANCELJOBREQUEST']._serialized_end=10643
  _globals['_DEVICECONNECTIONINFO']._serialized_start=10646
  _globals['_DEVICECONNECTIONINFO']._serialized_end=11096
  _globals['_DEVICECONNECTIONINFO_DEVICEATTRIBUTESENTRY']._serialized_start=2991
  _globals['_DEVICECONNECTIONINFO_DEVICEATTRIBUTESENTRY']._serialized_end=3101
  _globals['_LISTOFMODEMINFO']._serialized_start=11098
  _globals['_LISTOFMODEMINFO']._serialized_end=11184
  _globals['_LISTOFCONNECTIONINFO']._serialized_start=11186
  _globals['_LISTOFCONNECTIONINFO']._serialized_end=11288
  _globals['_ACTIONRESULT']._serialized_start=11291
  _globals['_ACTIONRESULT']._serialized_end=11480
  _globals['_JOBEVENTDATA']._serialized_start=11482
  _globals['_JOBEVENTDATA']._serialized_end=11519
  _globals['_DEVICESPEC']._serialized_start=11522
  _globals['_DEVICESPEC']._serialized_end=11914
  _globals['_DEVICESPEC_ATTRIBUTESENTRY']._serialized_start=2206
  _globals['_DEVICESPEC_ATTRIBUTESENTRY']._serialized_end=2310
  _globals['_DEVICECOMMUNICATIONUNIT']._serialized_start=11917
  _globals['_DEVICECOMMUNICATIONUNIT']._serialized_end=12086
  _globals['_JOBSPEC']._serialized_start=12089
  _globals['_JOBSPEC']._serialized_end=12366
  _globals['_JOBDONENOTIFICATION']._serialized_start=12369
  _globals['_JOBDONENOTIFICATION']._serialized_end=12526
  _globals['_DRIVERINFO']._serialized_start=12528
  _globals['_DRIVERINFO']._serialized_end=12599
  _globals['_COMMUNICATIONUNIT']._serialized_start=12602
  _globals['_COMMUNICATIONUNIT']._serialized_end=12788
  _globals['_COMMUNICATIONUNITSPEC']._serialized_start=12791
  _globals['_COMMUNICATIONUNITSPEC']._serialized_end=12940
# @@protoc_insertion_point(module_scope)
