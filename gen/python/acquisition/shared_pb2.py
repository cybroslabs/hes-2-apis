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


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x18\x61\x63quisition/shared.proto\x12\"io.clbs.openhes.models.acquisition\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x13\x63ommon/fields.proto\x1a\x15\x63ommon/metadata.proto\x1a\x12\x63ommon/types.proto\"\x96\x02\n\x0bJobSettings\x12!\n\x0cmax_duration\x18\x01 \x01(\x03R\x0bmaxDuration\x12K\n\x08priority\x18\x02 \x01(\x0e\x32/.io.clbs.openhes.models.acquisition.JobPriorityR\x08priority\x12\x1a\n\x08\x61ttempts\x18\x03 \x03(\x05R\x08\x61ttempts\x12\x1f\n\x0bretry_delay\x18\x04 \x01(\x03R\nretryDelay\x12\x1f\n\x0b\x64\x65\x66\x65r_start\x18\x05 \x01(\x04R\ndeferStart\x12\x39\n\nexpires_at\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\texpiresAt\"\xe6\x0b\n\tJobAction\x12\x1b\n\taction_id\x18\x01 \x01(\tR\x08\x61\x63tionId\x12]\n\nattributes\x18\x02 \x03(\x0b\x32=.io.clbs.openhes.models.acquisition.JobAction.AttributesEntryR\nattributes\x12Z\n\x0cget_register\x18\x03 \x01(\x0b\x32\x35.io.clbs.openhes.models.acquisition.ActionGetRegisterH\x00R\x0bgetRegister\x12v\n\x16get_periodical_profile\x18\x04 \x01(\x0b\x32>.io.clbs.openhes.models.acquisition.ActionGetPeriodicalProfileH\x00R\x14getPeriodicalProfile\x12s\n\x15get_irregular_profile\x18\x05 \x01(\x0b\x32=.io.clbs.openhes.models.acquisition.ActionGetIrregularProfileH\x00R\x13getIrregularProfile\x12T\n\nget_events\x18\x06 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.ActionGetEventsH\x00R\tgetEvents\x12\x61\n\x0fget_device_info\x18\x07 \x01(\x0b\x32\x37.io.clbs.openhes.models.acquisition.ActionGetDeviceInfoH\x00R\rgetDeviceInfo\x12T\n\nsync_clock\x18\x08 \x01(\x0b\x32\x33.io.clbs.openhes.models.acquisition.ActionSyncClockH\x00R\tsyncClock\x12\x61\n\x0fset_relay_state\x18\n \x01(\x0b\x32\x37.io.clbs.openhes.models.acquisition.ActionSetRelayStateH\x00R\rsetRelayState\x12v\n\x16set_disconnector_state\x18\x0c \x01(\x0b\x32>.io.clbs.openhes.models.acquisition.ActionSetDisconnectorStateH\x00R\x14setDisconnectorState\x12K\n\x07get_tou\x18\r \x01(\x0b\x32\x30.io.clbs.openhes.models.acquisition.ActionGetTouH\x00R\x06getTou\x12K\n\x07set_tou\x18\x0e \x01(\x0b\x32\x30.io.clbs.openhes.models.acquisition.ActionSetTouH\x00R\x06setTou\x12W\n\x0bset_limiter\x18\x10 \x01(\x0b\x32\x34.io.clbs.openhes.models.acquisition.ActionSetLimiterH\x00R\nsetLimiter\x12p\n\x14reset_billing_period\x18\x11 \x01(\x0b\x32<.io.clbs.openhes.models.acquisition.ActionResetBillingPeriodH\x00R\x12resetBillingPeriod\x12Q\n\tfw_update\x18\x12 \x01(\x0b\x32\x32.io.clbs.openhes.models.acquisition.ActionFwUpdateH\x00R\x08\x66wUpdate\x1ah\n\x0f\x41ttributesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12?\n\x05value\x18\x02 \x01(\x0b\x32).io.clbs.openhes.models.common.FieldValueR\x05value:\x02\x38\x01\x42\x08\n\x06\x61\x63tion\"T\n\x0fListOfJobDevice\x12\x41\n\x04list\x18\x05 \x03(\x0b\x32-.io.clbs.openhes.models.acquisition.JobDeviceR\x04list\"X\n\x11ListOfJobDeviceId\x12\x43\n\x04list\x18\x05 \x03(\x0b\x32/.io.clbs.openhes.models.acquisition.JobDeviceIdR\x04list\"A\n\x0bJobDeviceId\x12\x15\n\x06job_id\x18\x01 \x01(\tR\x05jobId\x12\x1b\n\tdevice_id\x18\x02 \x01(\tR\x08\x64\x65viceId\"\x97\x04\n\tJobDevice\x12\x15\n\x06job_id\x18\x01 \x01(\tR\x05jobId\x12\x1b\n\tdevice_id\x18\x02 \x01(\tR\x08\x64\x65viceId\x12\x1f\n\x0b\x65xternal_id\x18\x03 \x01(\tR\nexternalId\x12p\n\x11\x64\x65vice_attributes\x18\x04 \x03(\x0b\x32\x43.io.clbs.openhes.models.acquisition.JobDevice.DeviceAttributesEntryR\x10\x64\x65viceAttributes\x12[\n\x0f\x63onnection_info\x18\x05 \x03(\x0b\x32\x32.io.clbs.openhes.models.acquisition.ConnectionInfoR\x0e\x63onnectionInfo\x12Z\n\x0c\x61pp_protocol\x18\x06 \x01(\x0e\x32\x37.io.clbs.openhes.models.acquisition.ApplicationProtocolR\x0b\x61ppProtocol\x12\x1a\n\x08timezone\x18\x07 \x01(\tR\x08timezone\x1an\n\x15\x44\x65viceAttributesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12?\n\x05value\x18\x02 \x01(\x0b\x32).io.clbs.openhes.models.common.FieldValueR\x05value:\x02\x38\x01\"\xf9\x03\n\tModemInfo\x12\x19\n\x08modem_id\x18\x01 \x01(\tR\x07modemId\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\x12\x17\n\x07\x61t_init\x18\x03 \x01(\tR\x06\x61tInit\x12\x17\n\x07\x61t_dial\x18\x04 \x01(\tR\x06\x61tDial\x12\x1b\n\tat_hangup\x18\x05 \x01(\tR\x08\x61tHangup\x12\x1b\n\tat_escape\x18\x06 \x01(\tR\x08\x61tEscape\x12\'\n\x0f\x63onnect_timeout\x18\x07 \x01(\rR\x0e\x63onnectTimeout\x12\'\n\x0f\x63ommand_timeout\x18\x08 \x01(\rR\x0e\x63ommandTimeout\x12U\n\x05tcpip\x18\t \x01(\x0b\x32=.io.clbs.openhes.models.acquisition.ConnectionTypeDirectTcpIpH\x00R\x05tcpip\x12j\n\x0eserial_over_ip\x18\n \x01(\x0b\x32\x42.io.clbs.openhes.models.acquisition.ConnectionTypeControlledSerialH\x00R\x0cserialOverIp\x12(\n\x10serial_baud_rate\x18\x0b \x01(\rR\x0eserialBaudRateB\x12\n\x10modem_connection\"\xe3\x02\n\x0cSerialConfig\x12U\n\tbaud_rate\x18\x01 \x01(\x0e\x32\x38.io.clbs.openhes.models.acquisition.SerialConfigBaudRateR\x08\x62\x61udRate\x12N\n\x06parity\x18\x02 \x01(\x0e\x32\x36.io.clbs.openhes.models.acquisition.SerialConfigParityR\x06parity\x12U\n\tdata_bits\x18\x03 \x01(\x0e\x32\x38.io.clbs.openhes.models.acquisition.SerialConfigDataBitsR\x08\x64\x61taBits\x12U\n\tstop_bits\x18\x04 \x01(\x0e\x32\x38.io.clbs.openhes.models.acquisition.SerialConfigStopBitsR\x08stopBits\"\xff\x04\n\x0e\x43onnectionInfo\x12U\n\x05tcpip\x18\x01 \x01(\x0b\x32=.io.clbs.openhes.models.acquisition.ConnectionTypeDirectTcpIpH\x00R\x05tcpip\x12\\\n\nmodem_pool\x18\x02 \x01(\x0b\x32;.io.clbs.openhes.models.acquisition.ConnectionTypeModemPoolH\x00R\tmodemPool\x12j\n\x0eserial_over_ip\x18\x03 \x01(\x0b\x32\x42.io.clbs.openhes.models.acquisition.ConnectionTypeControlledSerialH\x00R\x0cserialOverIp\x12Y\n\rlink_protocol\x18\x04 \x01(\x0e\x32\x34.io.clbs.openhes.models.acquisition.DataLinkProtocolR\x0clinkProtocol\x12\x15\n\x06\x62us_id\x18\x05 \x01(\tR\x05\x62usId\x12\x62\n\nattributes\x18\x06 \x03(\x0b\x32\x42.io.clbs.openhes.models.acquisition.ConnectionInfo.AttributesEntryR\nattributes\x1ah\n\x0f\x41ttributesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12?\n\x05value\x18\x02 \x01(\x0b\x32).io.clbs.openhes.models.common.FieldValueR\x05value:\x02\x38\x01\x42\x0c\n\nconnection\"]\n\x19\x43onnectionTypeDirectTcpIp\x12\x12\n\x04host\x18\x01 \x01(\tR\x04host\x12\x12\n\x04port\x18\x02 \x01(\rR\x04port\x12\x18\n\x07timeout\x18\x03 \x01(\x05R\x07timeout\"\x8f\x01\n\x17\x43onnectionTypeModemPool\x12\x16\n\x06number\x18\x01 \x01(\tR\x06number\x12\x17\n\x07pool_id\x18\x02 \x01(\tR\x06poolId\x12\x43\n\x05modem\x18\x04 \x01(\x0b\x32-.io.clbs.openhes.models.acquisition.ModemInfoR\x05modem\"^\n\x1a\x43onnectionTypeSerialDirect\x12\x12\n\x04host\x18\x01 \x01(\tR\x04host\x12\x12\n\x04port\x18\x02 \x01(\rR\x04port\x12\x18\n\x07timeout\x18\x03 \x01(\x05R\x07timeout\"\x87\x01\n\x18\x43onnectionTypeSerialMoxa\x12\x12\n\x04host\x18\x01 \x01(\tR\x04host\x12\x1a\n\x08\x64\x61taPort\x18\x02 \x01(\rR\x08\x64\x61taPort\x12!\n\x0c\x63ommand_port\x18\x03 \x01(\rR\x0b\x63ommandPort\x12\x18\n\x07timeout\x18\x04 \x01(\x05R\x07timeout\"_\n\x1b\x43onnectionTypeSerialRfc2217\x12\x12\n\x04host\x18\x01 \x01(\tR\x04host\x12\x12\n\x04port\x18\x02 \x01(\rR\x04port\x12\x18\n\x07timeout\x18\x03 \x01(\x05R\x07timeout\"\xd2\x01\n\x1b\x41pplicationProtocolTemplate\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12S\n\x08protocol\x18\x02 \x01(\x0e\x32\x37.io.clbs.openhes.models.acquisition.ApplicationProtocolR\x08protocol\x12N\n\nattributes\x18\x03 \x03(\x0b\x32..io.clbs.openhes.models.common.FieldDescriptorR\nattributes\"\xa2\x02\n\x10\x44\x61taLinkTemplate\x12Y\n\rlink_protocol\x18\x01 \x01(\x0e\x32\x34.io.clbs.openhes.models.acquisition.DataLinkProtocolR\x0clinkProtocol\x12\x63\n\x11\x61pp_protocol_refs\x18\x02 \x03(\x0e\x32\x37.io.clbs.openhes.models.acquisition.ApplicationProtocolR\x0f\x61ppProtocolRefs\x12N\n\nattributes\x18\x03 \x03(\x0b\x32..io.clbs.openhes.models.common.FieldDescriptorR\nattributes\"\xb6\x01\n\x15\x43ommunicationTemplate\x12I\n\x04type\x18\x01 \x01(\x0e\x32\x35.io.clbs.openhes.models.acquisition.CommunicationTypeR\x04type\x12R\n\tdatalinks\x18\x02 \x03(\x0b\x32\x34.io.clbs.openhes.models.acquisition.DataLinkTemplateR\tdatalinks\"9\n\x13\x41\x63\x63\x65ssLevelTemplate\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\"\x9d\x04\n\x0f\x44riverTemplates\x12r\n\x17\x63ommunication_templates\x18\x01 \x03(\x0b\x32\x39.io.clbs.openhes.models.acquisition.CommunicationTemplateR\x16\x63ommunicationTemplates\x12\x64\n\rapp_protocols\x18\x02 \x03(\x0b\x32?.io.clbs.openhes.models.acquisition.ApplicationProtocolTemplateR\x0c\x61ppProtocols\x12\x64\n\x11\x61\x63tion_attributes\x18\x03 \x03(\x0b\x32\x37.io.clbs.openhes.models.acquisition.JobActionAttributesR\x10\x61\x63tionAttributes\x12\x62\n\x10\x61\x63\x63\x65ss_templates\x18\x04 \x03(\x0b\x32\x37.io.clbs.openhes.models.acquisition.AccessLevelTemplateR\x0f\x61\x63\x63\x65ssTemplates\x12\x66\n\x12\x61\x63tion_constraints\x18\x05 \x01(\x0b\x32\x37.io.clbs.openhes.models.acquisition.JobActionContraintsR\x11\x61\x63tionConstraints\"\xd8\x01\n\x14\x41\x63tionProgressUpdate\x12\x15\n\x06job_id\x18\x01 \x01(\tR\x05jobId\x12\x1b\n\taction_id\x18\x02 \x01(\tR\x08\x61\x63tionId\x12H\n\x04\x63ode\x18\x03 \x01(\x0e\x32\x34.io.clbs.openhes.models.acquisition.ActionResultCodeR\x04\x63ode\x12\x42\n\x04\x64\x61ta\x18\x04 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.ActionDataR\x04\x64\x61ta\"\x8c\x01\n\x11JobProgressUpdate\x12\x15\n\x06job_id\x18\x01 \x01(\tR\x05jobId\x12\x44\n\x04\x63ode\x18\x02 \x01(\x0e\x32\x30.io.clbs.openhes.models.acquisition.JobErrorCodeR\x04\x63ode\x12\x1a\n\x08\x64uration\x18\x03 \x01(\x03R\x08\x64uration\"\xa4\x03\n\nActionData\x12\x30\n\x06nodata\x18\x01 \x01(\x0b\x32\x16.google.protobuf.EmptyH\x00R\x06nodata\x12O\n\x08\x62illings\x18\x02 \x01(\x0b\x32\x31.io.clbs.openhes.models.acquisition.BillingValuesH\x00R\x08\x62illings\x12M\n\x07profile\x18\x03 \x01(\x0b\x32\x31.io.clbs.openhes.models.acquisition.ProfileValuesH\x00R\x07profile\x12i\n\x11irregular_profile\x18\x04 \x01(\x0b\x32:.io.clbs.openhes.models.acquisition.IrregularProfileValuesH\x00R\x10irregularProfile\x12Q\n\x0b\x64\x65vice_info\x18\x05 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.DeviceInfoH\x00R\ndeviceInfoB\x06\n\x04\x64\x61ta\"\xa3\x03\n\nDeviceInfo\x12\x41\n\x0einfo_timestamp\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\rinfoTimestamp\x12<\n\x1amanufacturer_serial_number\x18\x02 \x01(\tR\x18manufacturerSerialNumber\x12\x30\n\x14\x64\x65vice_serial_number\x18\x03 \x01(\tR\x12\x64\x65viceSerialNumber\x12)\n\x10\x66irmware_version\x18\x04 \x01(\tR\x0f\x66irmwareVersion\x12\x1f\n\x0b\x63lock_delta\x18\x05 \x01(\x01R\nclockDelta\x12!\n\x0c\x64\x65vice_model\x18\x06 \x01(\tR\x0b\x64\x65viceModel\x12%\n\x0e\x65rror_register\x18\x07 \x01(\rR\rerrorRegister\x12!\n\x0crelay_states\x18\x08 \x03(\x08R\x0brelayStates\x12)\n\x10\x63onnection_state\x18\t \x01(\x08R\x0f\x63onnectionState\"\x84\x01\n\rProfileValues\x12\x16\n\x06period\x18\x01 \x01(\x05R\x06period\x12\x12\n\x04unit\x18\x02 \x01(\tR\x04unit\x12G\n\x06\x62locks\x18\x03 \x03(\x0b\x32/.io.clbs.openhes.models.acquisition.ProfileBlokR\x06\x62locks\"x\n\x16IrregularProfileValues\x12\x12\n\x04unit\x18\x01 \x01(\tR\x04unit\x12J\n\x06values\x18\x02 \x03(\x0b\x32\x32.io.clbs.openhes.models.acquisition.IrregularValueR\x06values\"\x93\x01\n\x0eIrregularValue\x12\x38\n\ttimestamp\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\ttimestamp\x12G\n\x05value\x18\x02 \x01(\x0b\x32\x31.io.clbs.openhes.models.acquisition.MeasuredValueR\x05value\"\x9d\x01\n\x0bProfileBlok\x12\x43\n\x0fstart_timestamp\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x0estartTimestamp\x12I\n\x06values\x18\x02 \x03(\x0b\x32\x31.io.clbs.openhes.models.acquisition.MeasuredValueR\x06values\"Y\n\rBillingValues\x12H\n\x06values\x18\x01 \x03(\x0b\x32\x30.io.clbs.openhes.models.acquisition.BillingValueR\x06values\"\xa5\x01\n\x0c\x42illingValue\x12\x38\n\ttimestamp\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\ttimestamp\x12\x12\n\x04unit\x18\x02 \x01(\tR\x04unit\x12G\n\x05value\x18\x03 \x01(\x0b\x32\x31.io.clbs.openhes.models.acquisition.MeasuredValueR\x05value\"\xd4\x02\n\rMeasuredValue\x12\x16\n\x06status\x18\x01 \x01(\x03R\x06status\x12\x1a\n\x08\x65xponent\x18\x02 \x01(\x05R\x08\x65xponent\x12#\n\x0c\x64ouble_value\x18\x03 \x01(\x01H\x00R\x0b\x64oubleValue\x12%\n\rinteger_value\x18\x04 \x01(\x03H\x00R\x0cintegerValue\x12#\n\x0cstring_value\x18\x05 \x01(\tH\x00R\x0bstringValue\x12\x45\n\x0ftimestamp_value\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampH\x00R\x0etimestampValue\x12.\n\x12timestamp_tz_value\x18\x07 \x01(\tH\x00R\x10timestampTzValue\x12\x1f\n\nbool_value\x18\x08 \x01(\x08H\x00R\tboolValueB\x06\n\x04kind\"\xa9\x01\n\x13JobActionAttributes\x12\x42\n\x04type\x18\x01 \x01(\x0e\x32..io.clbs.openhes.models.acquisition.ActionTypeR\x04type\x12N\n\nattributes\x18\x02 \x03(\x0b\x32..io.clbs.openhes.models.common.FieldDescriptorR\nattributes\"\xb8\x02\n\x1e\x43onnectionTypeControlledSerial\x12X\n\x06\x64irect\x18\x01 \x01(\x0b\x32>.io.clbs.openhes.models.acquisition.ConnectionTypeSerialDirectH\x00R\x06\x64irect\x12R\n\x04moxa\x18\x02 \x01(\x0b\x32<.io.clbs.openhes.models.acquisition.ConnectionTypeSerialMoxaH\x00R\x04moxa\x12[\n\x07rfc2217\x18\x03 \x01(\x0b\x32?.io.clbs.openhes.models.acquisition.ConnectionTypeSerialRfc2217H\x00R\x07rfc2217B\x0b\n\tconverter\"\x13\n\x11\x41\x63tionGetRegister\"x\n\x1a\x41\x63tionGetPeriodicalProfile\x12.\n\x04\x66rom\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x04\x66rom\x12*\n\x02to\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x02to\"w\n\x19\x41\x63tionGetIrregularProfile\x12.\n\x04\x66rom\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x04\x66rom\x12*\n\x02to\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x02to\"m\n\x0f\x41\x63tionGetEvents\x12.\n\x04\x66rom\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x04\x66rom\x12*\n\x02to\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x02to\"\x15\n\x13\x41\x63tionGetDeviceInfo\"\x11\n\x0f\x41\x63tionSyncClock\"\x15\n\x13\x41\x63tionSetRelayState\"\x1c\n\x1a\x41\x63tionSetDisconnectorState\"\x0e\n\x0c\x41\x63tionGetTou\"\x0e\n\x0c\x41\x63tionSetTou\"\x12\n\x10\x41\x63tionSetLimiter\"\x1a\n\x18\x41\x63tionResetBillingPeriod\"\x10\n\x0e\x41\x63tionFwUpdate\"\xfa\x03\n\x13JobActionContraints\x12\x85\x01\n\x16get_register_type_name\x18\x01 \x03(\x0b\x32P.io.clbs.openhes.models.acquisition.JobActionContraints.GetRegisterTypeNameEntryR\x13getRegisterTypeName\x12\x97\x01\n\x1cget_register_type_attributes\x18\x02 \x03(\x0b\x32V.io.clbs.openhes.models.acquisition.JobActionContraints.GetRegisterTypeAttributesEntryR\x19getRegisterTypeAttributes\x1a\x46\n\x18GetRegisterTypeNameEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\x1ay\n\x1eGetRegisterTypeAttributesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x41\n\x05value\x18\x02 \x01(\x0b\x32+.io.clbs.openhes.models.common.ListOfStringR\x05value:\x02\x38\x01\"g\n\x07\x42ulkJob\x12\x15\n\x06job_id\x18\x01 \x01(\tR\x05jobId\x12\x45\n\x06status\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.acquisition.JobStatusR\x06status\"\x91\x04\n\tJobStatus\x12I\n\x06status\x18\x01 \x01(\x0e\x32\x31.io.clbs.openhes.models.acquisition.JobStatusCodeR\x06status\x12\x44\n\x04\x63ode\x18\x02 \x01(\x0e\x32\x30.io.clbs.openhes.models.acquisition.JobErrorCodeR\x04\x63ode\x12J\n\x07results\x18\x03 \x03(\x0b\x32\x30.io.clbs.openhes.models.acquisition.ActionResultR\x07results\x12\x39\n\ncreated_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x39\n\nstarted_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tstartedAt\x12;\n\x0b\x66inished_at\x18\x06 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\nfinishedAt\x12#\n\rattempts_done\x18\x07 \x01(\x05R\x0c\x61ttemptsDone\x12O\n\x0b\x64\x65vice_info\x18\x08 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.DeviceInfoR\ndeviceInfo\"\xa6\x04\n\x0cStartJobData\x12s\n\x11\x64\x65vice_attributes\x18\x01 \x03(\x0b\x32\x46.io.clbs.openhes.models.acquisition.StartJobData.DeviceAttributesEntryR\x10\x64\x65viceAttributes\x12\x15\n\x06job_id\x18\x02 \x01(\tR\x05jobId\x12R\n\x0cjob_settings\x18\x03 \x01(\x0b\x32/.io.clbs.openhes.models.acquisition.JobSettingsR\x0bjobSettings\x12N\n\x0bjob_actions\x18\x04 \x03(\x0b\x32-.io.clbs.openhes.models.acquisition.JobActionR\njobActions\x12Z\n\x0c\x61pp_protocol\x18\x05 \x01(\x0e\x32\x37.io.clbs.openhes.models.acquisition.ApplicationProtocolR\x0b\x61ppProtocol\x12\x1a\n\x08timezone\x18\x06 \x01(\tR\x08timezone\x1an\n\x15\x44\x65viceAttributesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12?\n\x05value\x18\x02 \x01(\x0b\x32).io.clbs.openhes.models.common.FieldValueR\x05value:\x02\x38\x01\")\n\x10\x43\x61ncelJobRequest\x12\x15\n\x06job_id\x18\x01 \x01(\tR\x05jobId\"\xc2\x03\n\x14\x44\x65viceConnectionInfo\x12\x61\n\x12\x63ommunication_unit\x18\x01 \x01(\x0b\x32\x32.io.clbs.openhes.models.acquisition.ConnectionInfoR\x11\x63ommunicationUnit\x12Z\n\x0c\x61pp_protocol\x18\x02 \x01(\x0e\x32\x37.io.clbs.openhes.models.acquisition.ApplicationProtocolR\x0b\x61ppProtocol\x12{\n\x11\x64\x65vice_attributes\x18\x03 \x03(\x0b\x32N.io.clbs.openhes.models.acquisition.DeviceConnectionInfo.DeviceAttributesEntryR\x10\x64\x65viceAttributes\x1an\n\x15\x44\x65viceAttributesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12?\n\x05value\x18\x02 \x01(\x0b\x32).io.clbs.openhes.models.common.FieldValueR\x05value:\x02\x38\x01\"V\n\x0fListOfModemInfo\x12\x43\n\x05items\x18\x01 \x03(\x0b\x32-.io.clbs.openhes.models.acquisition.ModemInfoR\x05items\"f\n\x14ListOfConnectionInfo\x12N\n\x05items\x18\x01 \x03(\x0b\x32\x38.io.clbs.openhes.models.acquisition.DeviceConnectionInfoR\x05items\"\xbd\x01\n\x0c\x41\x63tionResult\x12\x1b\n\taction_id\x18\x01 \x01(\tR\x08\x61\x63tionId\x12L\n\x06status\x18\x02 \x01(\x0e\x32\x34.io.clbs.openhes.models.acquisition.ActionResultCodeR\x06status\x12\x42\n\x04\x64\x61ta\x18\x03 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.ActionDataR\x04\x64\x61ta\"%\n\x0cJobEventData\x12\x15\n\x06job_id\x18\x01 \x01(\x0cR\x05jobId\"\x88\x03\n\nDeviceSpec\x12\x1f\n\x0b\x65xternal_id\x18\x02 \x01(\tR\nexternalId\x12^\n\nattributes\x18\x04 \x03(\x0b\x32>.io.clbs.openhes.models.acquisition.DeviceSpec.AttributesEntryR\nattributes\x12s\n\x17\x63ommunication_unit_link\x18\x05 \x03(\x0b\x32;.io.clbs.openhes.models.acquisition.DeviceCommunicationUnitR\x15\x63ommunicationUnitLink\x12\x1a\n\x08timezone\x18\x06 \x01(\tR\x08timezone\x1ah\n\x0f\x41ttributesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12?\n\x05value\x18\x02 \x01(\x0b\x32).io.clbs.openhes.models.common.FieldValueR\x05value:\x02\x38\x01\"R\n\x0c\x44\x65viceStatus\x12\x42\n\x04info\x18\x01 \x01(\x0b\x32..io.clbs.openhes.models.acquisition.DeviceInfoR\x04info\"\xa9\x01\n\x17\x44\x65viceCommunicationUnit\x12\x32\n\x15\x63ommunication_unit_id\x18\x01 \x01(\tR\x13\x63ommunicationUnitId\x12Z\n\x0c\x61pp_protocol\x18\x02 \x01(\x0e\x32\x37.io.clbs.openhes.models.acquisition.ApplicationProtocolR\x0b\x61ppProtocol\"\x95\x02\n\x07JobSpec\x12\x45\n\x06\x64\x65vice\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.acquisition.JobDeviceR\x06\x64\x65vice\x12R\n\x0cjob_settings\x18\x03 \x01(\x0b\x32/.io.clbs.openhes.models.acquisition.JobSettingsR\x0bjobSettings\x12\x1f\n\x0b\x64river_type\x18\x04 \x01(\tR\ndriverType\x12N\n\x0bjob_actions\x18\x05 \x03(\x0b\x32-.io.clbs.openhes.models.acquisition.JobActionR\njobActions\"\x9d\x01\n\x13JobDoneNotification\x12?\n\x04spec\x18\x01 \x01(\x0b\x32+.io.clbs.openhes.models.acquisition.JobSpecR\x04spec\x12\x45\n\x06status\x18\x02 \x01(\x0b\x32-.io.clbs.openhes.models.acquisition.JobStatusR\x06status\"G\n\nDriverInfo\x12\x1f\n\x0b\x64river_type\x18\x01 \x01(\tR\ndriverType\x12\x18\n\x07version\x18\x02 \x01(\tR\x07version\"\xba\x01\n\x11\x43ommunicationUnit\x12T\n\x04spec\x18\x01 \x01(\x0b\x32\x39.io.clbs.openhes.models.acquisition.CommunicationUnitSpecB\x05\xaa\x01\x02\x08\x03R\x04spec\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadataJ\x04\x08\x02\x10\x03\"\x95\x01\n\x15\x43ommunicationUnitSpec\x12\x1f\n\x0b\x65xternal_id\x18\x01 \x01(\tR\nexternalId\x12[\n\x0f\x63onnection_info\x18\x02 \x01(\x0b\x32\x32.io.clbs.openhes.models.acquisition.ConnectionInfoR\x0e\x63onnectionInfo\"\xbf\x01\n\x14\x43ommunicationUnitBus\x12V\n\x06status\x18\x02 \x01(\x0b\x32>.io.clbs.openhes.models.acquisition.CommunicationUnitBusStatusR\x06status\x12I\n\x08metadata\x18\x03 \x01(\x0b\x32-.io.clbs.openhes.models.common.MetadataFieldsR\x08metadataJ\x04\x08\x01\x10\x02\"P\n\x1a\x43ommunicationUnitBusStatus\x12\x32\n\x15\x63ommunication_unit_id\x18\x01 \x03(\tR\x13\x63ommunicationUnitId*\xc1\x01\n\x0bJobPriority\x12\x12\n\x0eJOB_PRIORITY_0\x10\x00\x12\x12\n\x0eJOB_PRIORITY_1\x10\x01\x12\x12\n\x0eJOB_PRIORITY_2\x10\x02\x12\x12\n\x0eJOB_PRIORITY_3\x10\x03\x12\x12\n\x0eJOB_PRIORITY_4\x10\x04\x12\x12\n\x0eJOB_PRIORITY_5\x10\x05\x12\x12\n\x0eJOB_PRIORITY_6\x10\x06\x12\x12\n\x0eJOB_PRIORITY_7\x10\x07\x12\x12\n\x0eJOB_PRIORITY_8\x10\x08*i\n\x12SerialConfigParity\x12\x0f\n\x0bPARITY_NONE\x10\x00\x12\x0f\n\x0bPARITY_EVEN\x10\x01\x12\x0e\n\nPARITY_ODD\x10\x02\x12\x0f\n\x0bPARITY_MARK\x10\x03\x12\x10\n\x0cPARITY_SPACE\x10\x04*\xcb\x02\n\x14SerialConfigBaudRate\x12\x11\n\rBAUD_RATE_110\x10\x00\x12\x11\n\rBAUD_RATE_300\x10\x01\x12\x11\n\rBAUD_RATE_600\x10\x02\x12\x12\n\x0e\x42\x41UD_RATE_1200\x10\x03\x12\x12\n\x0e\x42\x41UD_RATE_2400\x10\x04\x12\x12\n\x0e\x42\x41UD_RATE_4800\x10\x05\x12\x12\n\x0e\x42\x41UD_RATE_9600\x10\x06\x12\x13\n\x0f\x42\x41UD_RATE_14400\x10\x07\x12\x13\n\x0f\x42\x41UD_RATE_19200\x10\x08\x12\x13\n\x0f\x42\x41UD_RATE_38400\x10\t\x12\x13\n\x0f\x42\x41UD_RATE_57600\x10\n\x12\x14\n\x10\x42\x41UD_RATE_115200\x10\x0b\x12\x14\n\x10\x42\x41UD_RATE_230400\x10\x0c\x12\x14\n\x10\x42\x41UD_RATE_460800\x10\r\x12\x14\n\x10\x42\x41UD_RATE_921600\x10\x0e*Z\n\x14SerialConfigDataBits\x12\x0f\n\x0b\x44\x41TA_BITS_5\x10\x00\x12\x0f\n\x0b\x44\x41TA_BITS_6\x10\x01\x12\x0f\n\x0b\x44\x41TA_BITS_7\x10\x02\x12\x0f\n\x0b\x44\x41TA_BITS_8\x10\x03*K\n\x14SerialConfigStopBits\x12\x0f\n\x0bSTOP_BITS_1\x10\x00\x12\x11\n\rSTOP_BITS_1_5\x10\x01\x12\x0f\n\x0bSTOP_BITS_2\x10\x02*\xca\x01\n\x11\x43ommunicationType\x12\x1c\n\x18\x43OMMUNICATION_TYPE_TCPIP\x10\x00\x12!\n\x1d\x43OMMUNICATION_TYPE_MODEM_POOL\x10\x01\x12)\n%COMMUNICATION_TYPE_SERIAL_LINE_DIRECT\x10\x02\x12\'\n#COMMUNICATION_TYPE_SERIAL_LINE_MOXA\x10\x03\x12 \n\x1c\x43OMMUNICATION_TYPE_LISTENING\x10\x63*\xa7\x01\n\x10\x44\x61taLinkProtocol\x12\x1a\n\x16LINKPROTO_IEC_62056_21\x10\x00\x12\x12\n\x0eLINKPROTO_HDLC\x10\x01\x12\x1b\n\x17LINKPROTO_COSEM_WRAPPER\x10\x02\x12\x14\n\x10LINKPROTO_MODBUS\x10\x03\x12\x12\n\x0eLINKPROTO_MBUS\x10\x04\x12\x1c\n\x18LINKPROTO_NOT_APPLICABLE\x10\x63*\xae\x01\n\x13\x41pplicationProtocol\x12\x19\n\x15\x41PPPROTO_IEC_62056_21\x10\x00\x12\x14\n\x10\x41PPPROTO_DLMS_SN\x10\x01\x12\x14\n\x10\x41PPPROTO_DLMS_LN\x10\x02\x12\x11\n\rAPPPROTO_SCTM\x10\x03\x12\x13\n\x0f\x41PPPROTO_LIS200\x10\x04\x12\x15\n\x11\x41PPPROTO_ANSI_C12\x10\x05\x12\x11\n\rAPPPROTO_MQTT\x10\x06*\xd3\x03\n\nActionType\x12\x1c\n\x18\x41\x43TION_TYPE_GET_REGISTER\x10\x00\x12&\n\"ACTION_TYPE_GET_PERIODICAL_PROFILE\x10\x01\x12%\n!ACTION_TYPE_GET_IRREGULAR_PROFILE\x10\x02\x12\x1a\n\x16\x41\x43TION_TYPE_GET_EVENTS\x10\x03\x12\x1f\n\x1b\x41\x43TION_TYPE_GET_DEVICE_INFO\x10\n\x12\x1a\n\x16\x41\x43TION_TYPE_SYNC_CLOCK\x10\x0b\x12\x1f\n\x1b\x41\x43TION_TYPE_SET_RELAY_STATE\x10\x15\x12&\n\"ACTION_TYPE_GET_DISCONNECTOR_STATE\x10\x16\x12&\n\"ACTION_TYPE_SET_DISCONNECTOR_STATE\x10\x17\x12\x17\n\x13\x41\x43TION_TYPE_GET_TOU\x10\x18\x12\x17\n\x13\x41\x43TION_TYPE_SET_TOU\x10\x19\x12\x1b\n\x17\x41\x43TION_TYPE_SET_LIMITER\x10\x1b\x12$\n ACTION_TYPE_RESET_BILLING_PERIOD\x10(\x12\x19\n\x15\x41\x43TION_TYPE_FW_UPDATE\x10\x32*\x8b\x01\n\x10\x41\x63tionResultCode\x12\x18\n\x14\x45RROR_CODE_ACTION_OK\x10\x00\x12!\n\x1d\x45RROR_CODE_ACTION_UNSUPPORTED\x10\x01\x12\x1d\n\x19\x45RROR_CODE_ACTION_PENDING\x10\x03\x12\x1b\n\x17\x45RROR_CODE_ACTION_ERROR\x10\x05*\x97\x01\n\x0cJobErrorCode\x12\x17\n\x13JOB_ERROR_CODE_NONE\x10\x00\x12\x17\n\x13JOB_ERROR_CODE_BUSY\x10\x01\x12\x18\n\x14JOB_ERROR_CODE_ERROR\x10\x05\x12!\n\x1dJOB_ERROR_CODE_ALREADY_EXISTS\x10\x08\x12\x18\n\x14JOB_ERROR_CODE_FATAL\x10\t*\x90\x01\n\x0e\x42ulkStatusCode\x12\x16\n\x12\x42ULK_STATUS_QUEUED\x10\x00\x12\x17\n\x13\x42ULK_STATUS_RUNNING\x10\x01\x12\x19\n\x15\x42ULK_STATUS_COMPLETED\x10\x02\x12\x19\n\x15\x42ULK_STATUS_CANCELLED\x10\x03\x12\x17\n\x13\x42ULK_STATUS_EXPIRED\x10\x04*\xa1\x01\n\rJobStatusCode\x12\x15\n\x11JOB_STATUS_QUEUED\x10\x00\x12\x16\n\x12JOB_STATUS_RUNNING\x10\x01\x12\x18\n\x14JOB_STATUS_COMPLETED\x10\x02\x12\x15\n\x11JOB_STATUS_FAILED\x10\x03\x12\x18\n\x14JOB_STATUS_CANCELLED\x10\x04\x12\x16\n\x12JOB_STATUS_EXPIRED\x10\x05\x42\x35Z3github.com/cybroslabs/hes-2-apis/gen/go/acquisitionb\x08\x65\x64itionsp\xe8\x07')

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
  _globals['_CONNECTIONINFO_ATTRIBUTESENTRY']._loaded_options = None
  _globals['_CONNECTIONINFO_ATTRIBUTESENTRY']._serialized_options = b'8\001'
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
  _globals['_JOBPRIORITY']._serialized_start=14000
  _globals['_JOBPRIORITY']._serialized_end=14193
  _globals['_SERIALCONFIGPARITY']._serialized_start=14195
  _globals['_SERIALCONFIGPARITY']._serialized_end=14300
  _globals['_SERIALCONFIGBAUDRATE']._serialized_start=14303
  _globals['_SERIALCONFIGBAUDRATE']._serialized_end=14634
  _globals['_SERIALCONFIGDATABITS']._serialized_start=14636
  _globals['_SERIALCONFIGDATABITS']._serialized_end=14726
  _globals['_SERIALCONFIGSTOPBITS']._serialized_start=14728
  _globals['_SERIALCONFIGSTOPBITS']._serialized_end=14803
  _globals['_COMMUNICATIONTYPE']._serialized_start=14806
  _globals['_COMMUNICATIONTYPE']._serialized_end=15008
  _globals['_DATALINKPROTOCOL']._serialized_start=15011
  _globals['_DATALINKPROTOCOL']._serialized_end=15178
  _globals['_APPLICATIONPROTOCOL']._serialized_start=15181
  _globals['_APPLICATIONPROTOCOL']._serialized_end=15355
  _globals['_ACTIONTYPE']._serialized_start=15358
  _globals['_ACTIONTYPE']._serialized_end=15825
  _globals['_ACTIONRESULTCODE']._serialized_start=15828
  _globals['_ACTIONRESULTCODE']._serialized_end=15967
  _globals['_JOBERRORCODE']._serialized_start=15970
  _globals['_JOBERRORCODE']._serialized_end=16121
  _globals['_BULKSTATUSCODE']._serialized_start=16124
  _globals['_BULKSTATUSCODE']._serialized_end=16268
  _globals['_JOBSTATUSCODE']._serialized_start=16271
  _globals['_JOBSTATUSCODE']._serialized_end=16432
  _globals['_JOBSETTINGS']._serialized_start=221
  _globals['_JOBSETTINGS']._serialized_end=499
  _globals['_JOBACTION']._serialized_start=502
  _globals['_JOBACTION']._serialized_end=2012
  _globals['_JOBACTION_ATTRIBUTESENTRY']._serialized_start=1898
  _globals['_JOBACTION_ATTRIBUTESENTRY']._serialized_end=2002
  _globals['_LISTOFJOBDEVICE']._serialized_start=2014
  _globals['_LISTOFJOBDEVICE']._serialized_end=2098
  _globals['_LISTOFJOBDEVICEID']._serialized_start=2100
  _globals['_LISTOFJOBDEVICEID']._serialized_end=2188
  _globals['_JOBDEVICEID']._serialized_start=2190
  _globals['_JOBDEVICEID']._serialized_end=2255
  _globals['_JOBDEVICE']._serialized_start=2258
  _globals['_JOBDEVICE']._serialized_end=2793
  _globals['_JOBDEVICE_DEVICEATTRIBUTESENTRY']._serialized_start=2683
  _globals['_JOBDEVICE_DEVICEATTRIBUTESENTRY']._serialized_end=2793
  _globals['_MODEMINFO']._serialized_start=2796
  _globals['_MODEMINFO']._serialized_end=3301
  _globals['_SERIALCONFIG']._serialized_start=3304
  _globals['_SERIALCONFIG']._serialized_end=3659
  _globals['_CONNECTIONINFO']._serialized_start=3662
  _globals['_CONNECTIONINFO']._serialized_end=4301
  _globals['_CONNECTIONINFO_ATTRIBUTESENTRY']._serialized_start=1898
  _globals['_CONNECTIONINFO_ATTRIBUTESENTRY']._serialized_end=2002
  _globals['_CONNECTIONTYPEDIRECTTCPIP']._serialized_start=4303
  _globals['_CONNECTIONTYPEDIRECTTCPIP']._serialized_end=4396
  _globals['_CONNECTIONTYPEMODEMPOOL']._serialized_start=4399
  _globals['_CONNECTIONTYPEMODEMPOOL']._serialized_end=4542
  _globals['_CONNECTIONTYPESERIALDIRECT']._serialized_start=4544
  _globals['_CONNECTIONTYPESERIALDIRECT']._serialized_end=4638
  _globals['_CONNECTIONTYPESERIALMOXA']._serialized_start=4641
  _globals['_CONNECTIONTYPESERIALMOXA']._serialized_end=4776
  _globals['_CONNECTIONTYPESERIALRFC2217']._serialized_start=4778
  _globals['_CONNECTIONTYPESERIALRFC2217']._serialized_end=4873
  _globals['_APPLICATIONPROTOCOLTEMPLATE']._serialized_start=4876
  _globals['_APPLICATIONPROTOCOLTEMPLATE']._serialized_end=5086
  _globals['_DATALINKTEMPLATE']._serialized_start=5089
  _globals['_DATALINKTEMPLATE']._serialized_end=5379
  _globals['_COMMUNICATIONTEMPLATE']._serialized_start=5382
  _globals['_COMMUNICATIONTEMPLATE']._serialized_end=5564
  _globals['_ACCESSLEVELTEMPLATE']._serialized_start=5566
  _globals['_ACCESSLEVELTEMPLATE']._serialized_end=5623
  _globals['_DRIVERTEMPLATES']._serialized_start=5626
  _globals['_DRIVERTEMPLATES']._serialized_end=6167
  _globals['_ACTIONPROGRESSUPDATE']._serialized_start=6170
  _globals['_ACTIONPROGRESSUPDATE']._serialized_end=6386
  _globals['_JOBPROGRESSUPDATE']._serialized_start=6389
  _globals['_JOBPROGRESSUPDATE']._serialized_end=6529
  _globals['_ACTIONDATA']._serialized_start=6532
  _globals['_ACTIONDATA']._serialized_end=6952
  _globals['_DEVICEINFO']._serialized_start=6955
  _globals['_DEVICEINFO']._serialized_end=7374
  _globals['_PROFILEVALUES']._serialized_start=7377
  _globals['_PROFILEVALUES']._serialized_end=7509
  _globals['_IRREGULARPROFILEVALUES']._serialized_start=7511
  _globals['_IRREGULARPROFILEVALUES']._serialized_end=7631
  _globals['_IRREGULARVALUE']._serialized_start=7634
  _globals['_IRREGULARVALUE']._serialized_end=7781
  _globals['_PROFILEBLOK']._serialized_start=7784
  _globals['_PROFILEBLOK']._serialized_end=7941
  _globals['_BILLINGVALUES']._serialized_start=7943
  _globals['_BILLINGVALUES']._serialized_end=8032
  _globals['_BILLINGVALUE']._serialized_start=8035
  _globals['_BILLINGVALUE']._serialized_end=8200
  _globals['_MEASUREDVALUE']._serialized_start=8203
  _globals['_MEASUREDVALUE']._serialized_end=8543
  _globals['_JOBACTIONATTRIBUTES']._serialized_start=8546
  _globals['_JOBACTIONATTRIBUTES']._serialized_end=8715
  _globals['_CONNECTIONTYPECONTROLLEDSERIAL']._serialized_start=8718
  _globals['_CONNECTIONTYPECONTROLLEDSERIAL']._serialized_end=9030
  _globals['_ACTIONGETREGISTER']._serialized_start=9032
  _globals['_ACTIONGETREGISTER']._serialized_end=9051
  _globals['_ACTIONGETPERIODICALPROFILE']._serialized_start=9053
  _globals['_ACTIONGETPERIODICALPROFILE']._serialized_end=9173
  _globals['_ACTIONGETIRREGULARPROFILE']._serialized_start=9175
  _globals['_ACTIONGETIRREGULARPROFILE']._serialized_end=9294
  _globals['_ACTIONGETEVENTS']._serialized_start=9296
  _globals['_ACTIONGETEVENTS']._serialized_end=9405
  _globals['_ACTIONGETDEVICEINFO']._serialized_start=9407
  _globals['_ACTIONGETDEVICEINFO']._serialized_end=9428
  _globals['_ACTIONSYNCCLOCK']._serialized_start=9430
  _globals['_ACTIONSYNCCLOCK']._serialized_end=9447
  _globals['_ACTIONSETRELAYSTATE']._serialized_start=9449
  _globals['_ACTIONSETRELAYSTATE']._serialized_end=9470
  _globals['_ACTIONSETDISCONNECTORSTATE']._serialized_start=9472
  _globals['_ACTIONSETDISCONNECTORSTATE']._serialized_end=9500
  _globals['_ACTIONGETTOU']._serialized_start=9502
  _globals['_ACTIONGETTOU']._serialized_end=9516
  _globals['_ACTIONSETTOU']._serialized_start=9518
  _globals['_ACTIONSETTOU']._serialized_end=9532
  _globals['_ACTIONSETLIMITER']._serialized_start=9534
  _globals['_ACTIONSETLIMITER']._serialized_end=9552
  _globals['_ACTIONRESETBILLINGPERIOD']._serialized_start=9554
  _globals['_ACTIONRESETBILLINGPERIOD']._serialized_end=9580
  _globals['_ACTIONFWUPDATE']._serialized_start=9582
  _globals['_ACTIONFWUPDATE']._serialized_end=9598
  _globals['_JOBACTIONCONTRAINTS']._serialized_start=9601
  _globals['_JOBACTIONCONTRAINTS']._serialized_end=10107
  _globals['_JOBACTIONCONTRAINTS_GETREGISTERTYPENAMEENTRY']._serialized_start=9914
  _globals['_JOBACTIONCONTRAINTS_GETREGISTERTYPENAMEENTRY']._serialized_end=9984
  _globals['_JOBACTIONCONTRAINTS_GETREGISTERTYPEATTRIBUTESENTRY']._serialized_start=9986
  _globals['_JOBACTIONCONTRAINTS_GETREGISTERTYPEATTRIBUTESENTRY']._serialized_end=10107
  _globals['_BULKJOB']._serialized_start=10109
  _globals['_BULKJOB']._serialized_end=10212
  _globals['_JOBSTATUS']._serialized_start=10215
  _globals['_JOBSTATUS']._serialized_end=10744
  _globals['_STARTJOBDATA']._serialized_start=10747
  _globals['_STARTJOBDATA']._serialized_end=11297
  _globals['_STARTJOBDATA_DEVICEATTRIBUTESENTRY']._serialized_start=2683
  _globals['_STARTJOBDATA_DEVICEATTRIBUTESENTRY']._serialized_end=2793
  _globals['_CANCELJOBREQUEST']._serialized_start=11299
  _globals['_CANCELJOBREQUEST']._serialized_end=11340
  _globals['_DEVICECONNECTIONINFO']._serialized_start=11343
  _globals['_DEVICECONNECTIONINFO']._serialized_end=11793
  _globals['_DEVICECONNECTIONINFO_DEVICEATTRIBUTESENTRY']._serialized_start=2683
  _globals['_DEVICECONNECTIONINFO_DEVICEATTRIBUTESENTRY']._serialized_end=2793
  _globals['_LISTOFMODEMINFO']._serialized_start=11795
  _globals['_LISTOFMODEMINFO']._serialized_end=11881
  _globals['_LISTOFCONNECTIONINFO']._serialized_start=11883
  _globals['_LISTOFCONNECTIONINFO']._serialized_end=11985
  _globals['_ACTIONRESULT']._serialized_start=11988
  _globals['_ACTIONRESULT']._serialized_end=12177
  _globals['_JOBEVENTDATA']._serialized_start=12179
  _globals['_JOBEVENTDATA']._serialized_end=12216
  _globals['_DEVICESPEC']._serialized_start=12219
  _globals['_DEVICESPEC']._serialized_end=12611
  _globals['_DEVICESPEC_ATTRIBUTESENTRY']._serialized_start=1898
  _globals['_DEVICESPEC_ATTRIBUTESENTRY']._serialized_end=2002
  _globals['_DEVICESTATUS']._serialized_start=12613
  _globals['_DEVICESTATUS']._serialized_end=12695
  _globals['_DEVICECOMMUNICATIONUNIT']._serialized_start=12698
  _globals['_DEVICECOMMUNICATIONUNIT']._serialized_end=12867
  _globals['_JOBSPEC']._serialized_start=12870
  _globals['_JOBSPEC']._serialized_end=13147
  _globals['_JOBDONENOTIFICATION']._serialized_start=13150
  _globals['_JOBDONENOTIFICATION']._serialized_end=13307
  _globals['_DRIVERINFO']._serialized_start=13309
  _globals['_DRIVERINFO']._serialized_end=13380
  _globals['_COMMUNICATIONUNIT']._serialized_start=13383
  _globals['_COMMUNICATIONUNIT']._serialized_end=13569
  _globals['_COMMUNICATIONUNITSPEC']._serialized_start=13572
  _globals['_COMMUNICATIONUNITSPEC']._serialized_end=13721
  _globals['_COMMUNICATIONUNITBUS']._serialized_start=13724
  _globals['_COMMUNICATIONUNITBUS']._serialized_end=13915
  _globals['_COMMUNICATIONUNITBUSSTATUS']._serialized_start=13917
  _globals['_COMMUNICATIONUNITBUSSTATUS']._serialized_end=13997
# @@protoc_insertion_point(module_scope)
