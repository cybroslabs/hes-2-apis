// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: services/svcapi/api.proto

package svcapi

import (
	acquisition "github.com/cybroslabs/hes-2-apis/gen/go/acquisition"
	_ "github.com/cybroslabs/hes-2-apis/gen/go/acquisition/timeofuse"
	common "github.com/cybroslabs/hes-2-apis/gen/go/common"
	system "github.com/cybroslabs/hes-2-apis/gen/go/system"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_services_svcapi_api_proto protoreflect.FileDescriptor

const file_services_svcapi_api_proto_rawDesc = "" +
	"\n" +
	"\x19services/svcapi/api.proto\x12\x1fio.clbs.openhes.services.svcapi\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x13common/fields.proto\x1a\x12common/types.proto\x1a\x16acquisition/main.proto\x1a\x18acquisition/shared.proto\x1a%acquisition/timeofuse/timeofuse.proto\x1a\x11system/main.proto2\xd2:\n" +
	"\n" +
	"ApiService\x12i\n" +
	"\x0eCreateVariable\x129.io.clbs.openhes.models.acquisition.CreateVariableRequest\x1a\x1c.google.protobuf.StringValue\x12p\n" +
	"\rListVariables\x12+.io.clbs.openhes.models.common.ListSelector\x1a2.io.clbs.openhes.models.acquisition.ListOfVariable\x12V\n" +
	"\x0eUpdateVariable\x12,.io.clbs.openhes.models.acquisition.Variable\x1a\x16.google.protobuf.Empty\x12F\n" +
	"\x0eDeleteVariable\x12\x1c.google.protobuf.StringValue\x1a\x16.google.protobuf.Empty\x12\x8f\x01\n" +
	"!CreateDeviceConfigurationRegister\x12L.io.clbs.openhes.models.acquisition.CreateDeviceConfigurationRegisterRequest\x1a\x1c.google.protobuf.StringValue\x12\x96\x01\n" +
	" ListDeviceConfigurationRegisters\x12+.io.clbs.openhes.models.common.ListSelector\x1aE.io.clbs.openhes.models.acquisition.ListOfDeviceConfigurationRegister\x12\x7f\n" +
	"\x1eGetDeviceConfigurationRegister\x12\x1c.google.protobuf.StringValue\x1a?.io.clbs.openhes.models.acquisition.DeviceConfigurationRegister\x12|\n" +
	"!UpdateDeviceConfigurationRegister\x12?.io.clbs.openhes.models.acquisition.DeviceConfigurationRegister\x1a\x16.google.protobuf.Empty\x12Y\n" +
	"!DeleteDeviceConfigurationRegister\x12\x1c.google.protobuf.StringValue\x1a\x16.google.protobuf.Empty\x12\x8f\x01\n" +
	"!CreateDeviceConfigurationTemplate\x12L.io.clbs.openhes.models.acquisition.CreateDeviceConfigurationTemplateRequest\x1a\x1c.google.protobuf.StringValue\x12\x96\x01\n" +
	" ListDeviceConfigurationTemplates\x12+.io.clbs.openhes.models.common.ListSelector\x1aE.io.clbs.openhes.models.acquisition.ListOfDeviceConfigurationTemplate\x12\x7f\n" +
	"\x1eGetDeviceConfigurationTemplate\x12\x1c.google.protobuf.StringValue\x1a?.io.clbs.openhes.models.acquisition.DeviceConfigurationTemplate\x12|\n" +
	"!UpdateDeviceConfigurationTemplate\x12?.io.clbs.openhes.models.acquisition.DeviceConfigurationTemplate\x1a\x16.google.protobuf.Empty\x12Y\n" +
	"!DeleteDeviceConfigurationTemplate\x12\x1c.google.protobuf.StringValue\x1a\x16.google.protobuf.Empty\x12\xbd\x01\n" +
	";AddDeviceConfigurationRegisterToDeviceConfigurationTemplate\x12f.io.clbs.openhes.models.acquisition.AddDeviceConfigurationRegisterToDeviceConfigurationTemplateRequest\x1a\x16.google.protobuf.Empty\x12\xc7\x01\n" +
	"@RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplate\x12k.io.clbs.openhes.models.acquisition.RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplateRequest\x1a\x16.google.protobuf.Empty\x12d\n" +
	"\x14ListFieldDescriptors\x12\x16.google.protobuf.Empty\x1a4.io.clbs.openhes.models.common.ListOfFieldDescriptor\x12h\n" +
	"\tListBulks\x12+.io.clbs.openhes.models.common.ListSelector\x1a..io.clbs.openhes.models.acquisition.ListOfBulk\x12z\n" +
	"\fListBulkJobs\x127.io.clbs.openhes.models.acquisition.ListBulkJobsRequest\x1a1.io.clbs.openhes.models.acquisition.ListOfBulkJob\x12W\n" +
	"\n" +
	"GetBulkJob\x12\x1c.google.protobuf.StringValue\x1a+.io.clbs.openhes.models.acquisition.BulkJob\x12B\n" +
	"\n" +
	"CancelBulk\x12\x1c.google.protobuf.StringValue\x1a\x16.google.protobuf.Empty\x12k\n" +
	"\x0fCreateProxyBulk\x12:.io.clbs.openhes.models.acquisition.CreateProxyBulkRequest\x1a\x1c.google.protobuf.StringValue\x12[\n" +
	"\fGetProxyBulk\x12\x1c.google.protobuf.StringValue\x1a-.io.clbs.openhes.models.acquisition.ProxyBulk\x12a\n" +
	"\n" +
	"CreateBulk\x125.io.clbs.openhes.models.acquisition.CreateBulkRequest\x1a\x1c.google.protobuf.StringValue\x12Q\n" +
	"\aGetBulk\x12\x1c.google.protobuf.StringValue\x1a(.io.clbs.openhes.models.acquisition.Bulk\x12l\n" +
	"\vListDrivers\x12+.io.clbs.openhes.models.common.ListSelector\x1a0.io.clbs.openhes.models.acquisition.ListOfDriver\x12U\n" +
	"\tGetDriver\x12\x1c.google.protobuf.StringValue\x1a*.io.clbs.openhes.models.acquisition.Driver\x12{\n" +
	"\x17CreateCommunicationUnit\x12B.io.clbs.openhes.models.acquisition.CreateCommunicationUnitRequest\x1a\x1c.google.protobuf.StringValue\x12\x82\x01\n" +
	"\x16ListCommunicationUnits\x12+.io.clbs.openhes.models.common.ListSelector\x1a;.io.clbs.openhes.models.acquisition.ListOfCommunicationUnit\x12k\n" +
	"\x14GetCommunicationUnit\x12\x1c.google.protobuf.StringValue\x1a5.io.clbs.openhes.models.acquisition.CommunicationUnit\x12y\n" +
	"\x16CreateCommunicationBus\x12A.io.clbs.openhes.models.acquisition.CreateCommunicationBusRequest\x1a\x1c.google.protobuf.StringValue\x12\x81\x01\n" +
	"\x16ListCommunicationBuses\x12+.io.clbs.openhes.models.common.ListSelector\x1a:.io.clbs.openhes.models.acquisition.ListOfCommunicationBus\x12\x95\x01\n" +
	"'AddCommunicationUnitsToCommunicationBus\x12R.io.clbs.openhes.models.acquisition.AddCommunicationUnitsToCommunicationBusRequest\x1a\x16.google.protobuf.Empty\x12\x9f\x01\n" +
	",RemoveCommunicationUnitsFromCommunicationBus\x12W.io.clbs.openhes.models.acquisition.RemoveCommunicationUnitsFromCommunicationBusRequest\x1a\x16.google.protobuf.Empty\x12e\n" +
	"\fCreateDevice\x127.io.clbs.openhes.models.acquisition.CreateDeviceRequest\x1a\x1c.google.protobuf.StringValue\x12R\n" +
	"\fUpdateDevice\x12*.io.clbs.openhes.models.acquisition.Device\x1a\x16.google.protobuf.Empty\x12l\n" +
	"\vListDevices\x12+.io.clbs.openhes.models.common.ListSelector\x1a0.io.clbs.openhes.models.acquisition.ListOfDevice\x12U\n" +
	"\tGetDevice\x12\x1c.google.protobuf.StringValue\x1a*.io.clbs.openhes.models.acquisition.Device\x12]\n" +
	"\rGetDeviceInfo\x12\x1c.google.protobuf.StringValue\x1a..io.clbs.openhes.models.acquisition.DeviceInfo\x12}\n" +
	"\x1bSetDeviceCommunicationUnits\x12F.io.clbs.openhes.models.acquisition.SetDeviceCommunicationUnitsRequest\x1a\x16.google.protobuf.Empty\x12~\n" +
	"\x1bGetDeviceCommunicationUnits\x12\x1c.google.protobuf.StringValue\x1aA.io.clbs.openhes.models.acquisition.ListOfDeviceCommunicationUnit\x12o\n" +
	"\x11CreateDeviceGroup\x12<.io.clbs.openhes.models.acquisition.CreateDeviceGroupRequest\x1a\x1c.google.protobuf.StringValue\x12v\n" +
	"\x10ListDeviceGroups\x12+.io.clbs.openhes.models.common.ListSelector\x1a5.io.clbs.openhes.models.acquisition.ListOfDeviceGroup\x12_\n" +
	"\x0eGetDeviceGroup\x12\x1c.google.protobuf.StringValue\x1a/.io.clbs.openhes.models.acquisition.DeviceGroup\x12i\n" +
	"\x11AddDevicesToGroup\x12<.io.clbs.openhes.models.acquisition.AddDevicesToGroupRequest\x1a\x16.google.protobuf.Empty\x12s\n" +
	"\x16RemoveDevicesFromGroup\x12A.io.clbs.openhes.models.acquisition.RemoveDevicesFromGroupRequest\x1a\x16.google.protobuf.Empty\x12\x8d\x01\n" +
	"\x16ListDeviceGroupDevices\x12A.io.clbs.openhes.models.acquisition.ListDeviceGroupDevicesRequest\x1a0.io.clbs.openhes.models.acquisition.ListOfDevice\x12r\n" +
	"\x0eListModemPools\x12+.io.clbs.openhes.models.common.ListSelector\x1a3.io.clbs.openhes.models.acquisition.ListOfModemPool\x12[\n" +
	"\fGetModemPool\x12\x1c.google.protobuf.StringValue\x1a-.io.clbs.openhes.models.acquisition.ModemPool\x12h\n" +
	"\x0fCreateModemPool\x127.io.clbs.openhes.models.acquisition.SetModemPoolRequest\x1a\x1c.google.protobuf.StringValue\x12b\n" +
	"\x0fUpdateModemPool\x127.io.clbs.openhes.models.acquisition.SetModemPoolRequest\x1a\x16.google.protobuf.Empty\x12G\n" +
	"\x0fDeleteModemPool\x12\x1c.google.protobuf.StringValue\x1a\x16.google.protobuf.Empty\x12`\n" +
	"\vCreateModem\x123.io.clbs.openhes.models.acquisition.SetModemRequest\x1a\x1c.google.protobuf.StringValue\x12Z\n" +
	"\vUpdateModem\x123.io.clbs.openhes.models.acquisition.SetModemRequest\x1a\x16.google.protobuf.Empty\x12C\n" +
	"\vDeleteModem\x12\x1c.google.protobuf.StringValue\x1a\x16.google.protobuf.Empty\x12P\n" +
	"\tGetConfig\x12\x16.google.protobuf.Empty\x1a+.io.clbs.openhes.models.system.SystemConfig\x12P\n" +
	"\tSetConfig\x12+.io.clbs.openhes.models.system.SystemConfig\x1a\x16.google.protobuf.Empty\x12\x86\x01\n" +
	"\x15GetMeterDataRegisters\x127.io.clbs.openhes.models.acquisition.GetMeterDataRequest\x1a2.io.clbs.openhes.models.acquisition.RegisterValues0\x01\x12\x84\x01\n" +
	"\x14GetMeterDataProfiles\x127.io.clbs.openhes.models.acquisition.GetMeterDataRequest\x1a1.io.clbs.openhes.models.acquisition.ProfileValues0\x01\x12\x96\x01\n" +
	"\x1dGetMeterDataIrregularProfiles\x127.io.clbs.openhes.models.acquisition.GetMeterDataRequest\x1a:.io.clbs.openhes.models.acquisition.IrregularProfileValues0\x01\x12\x7f\n" +
	"\x0eGetMeterEvents\x129.io.clbs.openhes.models.acquisition.GetMeterEventsRequest\x1a0.io.clbs.openhes.models.acquisition.EventRecords0\x01\x12u\n" +
	"\x14CreateTimeOfUseTable\x12?.io.clbs.openhes.models.acquisition.CreateTimeOfUseTableRequest\x1a\x1c.google.protobuf.StringValue\x12|\n" +
	"\x13ListTimeOfUseTables\x12+.io.clbs.openhes.models.common.ListSelector\x1a8.io.clbs.openhes.models.acquisition.ListOfTimeOfUseTable\x12e\n" +
	"\x11GetTimeOfUseTable\x12\x1c.google.protobuf.StringValue\x1a2.io.clbs.openhes.models.acquisition.TimeOfUseTable\x12b\n" +
	"\x14UpdateTimeOfUseTable\x122.io.clbs.openhes.models.acquisition.TimeOfUseTable\x1a\x16.google.protobuf.Empty\x12L\n" +
	"\x14DeleteTimeOfUseTable\x12\x1c.google.protobuf.StringValue\x1a\x16.google.protobuf.EmptyB9Z7github.com/cybroslabs/hes-2-apis/gen/go/services/svcapib\beditionsp\xe8\a"

var file_services_svcapi_api_proto_goTypes = []any{
	(*acquisition.CreateVariableRequest)(nil),                                                   // 0: io.clbs.openhes.models.acquisition.CreateVariableRequest
	(*common.ListSelector)(nil),                                                                 // 1: io.clbs.openhes.models.common.ListSelector
	(*acquisition.Variable)(nil),                                                                // 2: io.clbs.openhes.models.acquisition.Variable
	(*wrapperspb.StringValue)(nil),                                                              // 3: google.protobuf.StringValue
	(*acquisition.CreateDeviceConfigurationRegisterRequest)(nil),                                // 4: io.clbs.openhes.models.acquisition.CreateDeviceConfigurationRegisterRequest
	(*acquisition.DeviceConfigurationRegister)(nil),                                             // 5: io.clbs.openhes.models.acquisition.DeviceConfigurationRegister
	(*acquisition.CreateDeviceConfigurationTemplateRequest)(nil),                                // 6: io.clbs.openhes.models.acquisition.CreateDeviceConfigurationTemplateRequest
	(*acquisition.DeviceConfigurationTemplate)(nil),                                             // 7: io.clbs.openhes.models.acquisition.DeviceConfigurationTemplate
	(*acquisition.AddDeviceConfigurationRegisterToDeviceConfigurationTemplateRequest)(nil),      // 8: io.clbs.openhes.models.acquisition.AddDeviceConfigurationRegisterToDeviceConfigurationTemplateRequest
	(*acquisition.RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplateRequest)(nil), // 9: io.clbs.openhes.models.acquisition.RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplateRequest
	(*emptypb.Empty)(nil),                                                                       // 10: google.protobuf.Empty
	(*acquisition.ListBulkJobsRequest)(nil),                                                     // 11: io.clbs.openhes.models.acquisition.ListBulkJobsRequest
	(*acquisition.CreateProxyBulkRequest)(nil),                                                  // 12: io.clbs.openhes.models.acquisition.CreateProxyBulkRequest
	(*acquisition.CreateBulkRequest)(nil),                                                       // 13: io.clbs.openhes.models.acquisition.CreateBulkRequest
	(*acquisition.CreateCommunicationUnitRequest)(nil),                                          // 14: io.clbs.openhes.models.acquisition.CreateCommunicationUnitRequest
	(*acquisition.CreateCommunicationBusRequest)(nil),                                           // 15: io.clbs.openhes.models.acquisition.CreateCommunicationBusRequest
	(*acquisition.AddCommunicationUnitsToCommunicationBusRequest)(nil),                          // 16: io.clbs.openhes.models.acquisition.AddCommunicationUnitsToCommunicationBusRequest
	(*acquisition.RemoveCommunicationUnitsFromCommunicationBusRequest)(nil),                     // 17: io.clbs.openhes.models.acquisition.RemoveCommunicationUnitsFromCommunicationBusRequest
	(*acquisition.CreateDeviceRequest)(nil),                                                     // 18: io.clbs.openhes.models.acquisition.CreateDeviceRequest
	(*acquisition.Device)(nil),                                                                  // 19: io.clbs.openhes.models.acquisition.Device
	(*acquisition.SetDeviceCommunicationUnitsRequest)(nil),                                      // 20: io.clbs.openhes.models.acquisition.SetDeviceCommunicationUnitsRequest
	(*acquisition.CreateDeviceGroupRequest)(nil),                                                // 21: io.clbs.openhes.models.acquisition.CreateDeviceGroupRequest
	(*acquisition.AddDevicesToGroupRequest)(nil),                                                // 22: io.clbs.openhes.models.acquisition.AddDevicesToGroupRequest
	(*acquisition.RemoveDevicesFromGroupRequest)(nil),                                           // 23: io.clbs.openhes.models.acquisition.RemoveDevicesFromGroupRequest
	(*acquisition.ListDeviceGroupDevicesRequest)(nil),                                           // 24: io.clbs.openhes.models.acquisition.ListDeviceGroupDevicesRequest
	(*acquisition.SetModemPoolRequest)(nil),                                                     // 25: io.clbs.openhes.models.acquisition.SetModemPoolRequest
	(*acquisition.SetModemRequest)(nil),                                                         // 26: io.clbs.openhes.models.acquisition.SetModemRequest
	(*system.SystemConfig)(nil),                                                                 // 27: io.clbs.openhes.models.system.SystemConfig
	(*acquisition.GetMeterDataRequest)(nil),                                                     // 28: io.clbs.openhes.models.acquisition.GetMeterDataRequest
	(*acquisition.GetMeterEventsRequest)(nil),                                                   // 29: io.clbs.openhes.models.acquisition.GetMeterEventsRequest
	(*acquisition.CreateTimeOfUseTableRequest)(nil),                                             // 30: io.clbs.openhes.models.acquisition.CreateTimeOfUseTableRequest
	(*acquisition.TimeOfUseTable)(nil),                                                          // 31: io.clbs.openhes.models.acquisition.TimeOfUseTable
	(*acquisition.ListOfVariable)(nil),                                                          // 32: io.clbs.openhes.models.acquisition.ListOfVariable
	(*acquisition.ListOfDeviceConfigurationRegister)(nil),                                       // 33: io.clbs.openhes.models.acquisition.ListOfDeviceConfigurationRegister
	(*acquisition.ListOfDeviceConfigurationTemplate)(nil),                                       // 34: io.clbs.openhes.models.acquisition.ListOfDeviceConfigurationTemplate
	(*common.ListOfFieldDescriptor)(nil),                                                        // 35: io.clbs.openhes.models.common.ListOfFieldDescriptor
	(*acquisition.ListOfBulk)(nil),                                                              // 36: io.clbs.openhes.models.acquisition.ListOfBulk
	(*acquisition.ListOfBulkJob)(nil),                                                           // 37: io.clbs.openhes.models.acquisition.ListOfBulkJob
	(*acquisition.BulkJob)(nil),                                                                 // 38: io.clbs.openhes.models.acquisition.BulkJob
	(*acquisition.ProxyBulk)(nil),                                                               // 39: io.clbs.openhes.models.acquisition.ProxyBulk
	(*acquisition.Bulk)(nil),                                                                    // 40: io.clbs.openhes.models.acquisition.Bulk
	(*acquisition.ListOfDriver)(nil),                                                            // 41: io.clbs.openhes.models.acquisition.ListOfDriver
	(*acquisition.Driver)(nil),                                                                  // 42: io.clbs.openhes.models.acquisition.Driver
	(*acquisition.ListOfCommunicationUnit)(nil),                                                 // 43: io.clbs.openhes.models.acquisition.ListOfCommunicationUnit
	(*acquisition.CommunicationUnit)(nil),                                                       // 44: io.clbs.openhes.models.acquisition.CommunicationUnit
	(*acquisition.ListOfCommunicationBus)(nil),                                                  // 45: io.clbs.openhes.models.acquisition.ListOfCommunicationBus
	(*acquisition.ListOfDevice)(nil),                                                            // 46: io.clbs.openhes.models.acquisition.ListOfDevice
	(*acquisition.DeviceInfo)(nil),                                                              // 47: io.clbs.openhes.models.acquisition.DeviceInfo
	(*acquisition.ListOfDeviceCommunicationUnit)(nil),                                           // 48: io.clbs.openhes.models.acquisition.ListOfDeviceCommunicationUnit
	(*acquisition.ListOfDeviceGroup)(nil),                                                       // 49: io.clbs.openhes.models.acquisition.ListOfDeviceGroup
	(*acquisition.DeviceGroup)(nil),                                                             // 50: io.clbs.openhes.models.acquisition.DeviceGroup
	(*acquisition.ListOfModemPool)(nil),                                                         // 51: io.clbs.openhes.models.acquisition.ListOfModemPool
	(*acquisition.ModemPool)(nil),                                                               // 52: io.clbs.openhes.models.acquisition.ModemPool
	(*acquisition.RegisterValues)(nil),                                                          // 53: io.clbs.openhes.models.acquisition.RegisterValues
	(*acquisition.ProfileValues)(nil),                                                           // 54: io.clbs.openhes.models.acquisition.ProfileValues
	(*acquisition.IrregularProfileValues)(nil),                                                  // 55: io.clbs.openhes.models.acquisition.IrregularProfileValues
	(*acquisition.EventRecords)(nil),                                                            // 56: io.clbs.openhes.models.acquisition.EventRecords
	(*acquisition.ListOfTimeOfUseTable)(nil),                                                    // 57: io.clbs.openhes.models.acquisition.ListOfTimeOfUseTable
}
var file_services_svcapi_api_proto_depIdxs = []int32{
	0,  // 0: io.clbs.openhes.services.svcapi.ApiService.CreateVariable:input_type -> io.clbs.openhes.models.acquisition.CreateVariableRequest
	1,  // 1: io.clbs.openhes.services.svcapi.ApiService.ListVariables:input_type -> io.clbs.openhes.models.common.ListSelector
	2,  // 2: io.clbs.openhes.services.svcapi.ApiService.UpdateVariable:input_type -> io.clbs.openhes.models.acquisition.Variable
	3,  // 3: io.clbs.openhes.services.svcapi.ApiService.DeleteVariable:input_type -> google.protobuf.StringValue
	4,  // 4: io.clbs.openhes.services.svcapi.ApiService.CreateDeviceConfigurationRegister:input_type -> io.clbs.openhes.models.acquisition.CreateDeviceConfigurationRegisterRequest
	1,  // 5: io.clbs.openhes.services.svcapi.ApiService.ListDeviceConfigurationRegisters:input_type -> io.clbs.openhes.models.common.ListSelector
	3,  // 6: io.clbs.openhes.services.svcapi.ApiService.GetDeviceConfigurationRegister:input_type -> google.protobuf.StringValue
	5,  // 7: io.clbs.openhes.services.svcapi.ApiService.UpdateDeviceConfigurationRegister:input_type -> io.clbs.openhes.models.acquisition.DeviceConfigurationRegister
	3,  // 8: io.clbs.openhes.services.svcapi.ApiService.DeleteDeviceConfigurationRegister:input_type -> google.protobuf.StringValue
	6,  // 9: io.clbs.openhes.services.svcapi.ApiService.CreateDeviceConfigurationTemplate:input_type -> io.clbs.openhes.models.acquisition.CreateDeviceConfigurationTemplateRequest
	1,  // 10: io.clbs.openhes.services.svcapi.ApiService.ListDeviceConfigurationTemplates:input_type -> io.clbs.openhes.models.common.ListSelector
	3,  // 11: io.clbs.openhes.services.svcapi.ApiService.GetDeviceConfigurationTemplate:input_type -> google.protobuf.StringValue
	7,  // 12: io.clbs.openhes.services.svcapi.ApiService.UpdateDeviceConfigurationTemplate:input_type -> io.clbs.openhes.models.acquisition.DeviceConfigurationTemplate
	3,  // 13: io.clbs.openhes.services.svcapi.ApiService.DeleteDeviceConfigurationTemplate:input_type -> google.protobuf.StringValue
	8,  // 14: io.clbs.openhes.services.svcapi.ApiService.AddDeviceConfigurationRegisterToDeviceConfigurationTemplate:input_type -> io.clbs.openhes.models.acquisition.AddDeviceConfigurationRegisterToDeviceConfigurationTemplateRequest
	9,  // 15: io.clbs.openhes.services.svcapi.ApiService.RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplate:input_type -> io.clbs.openhes.models.acquisition.RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplateRequest
	10, // 16: io.clbs.openhes.services.svcapi.ApiService.ListFieldDescriptors:input_type -> google.protobuf.Empty
	1,  // 17: io.clbs.openhes.services.svcapi.ApiService.ListBulks:input_type -> io.clbs.openhes.models.common.ListSelector
	11, // 18: io.clbs.openhes.services.svcapi.ApiService.ListBulkJobs:input_type -> io.clbs.openhes.models.acquisition.ListBulkJobsRequest
	3,  // 19: io.clbs.openhes.services.svcapi.ApiService.GetBulkJob:input_type -> google.protobuf.StringValue
	3,  // 20: io.clbs.openhes.services.svcapi.ApiService.CancelBulk:input_type -> google.protobuf.StringValue
	12, // 21: io.clbs.openhes.services.svcapi.ApiService.CreateProxyBulk:input_type -> io.clbs.openhes.models.acquisition.CreateProxyBulkRequest
	3,  // 22: io.clbs.openhes.services.svcapi.ApiService.GetProxyBulk:input_type -> google.protobuf.StringValue
	13, // 23: io.clbs.openhes.services.svcapi.ApiService.CreateBulk:input_type -> io.clbs.openhes.models.acquisition.CreateBulkRequest
	3,  // 24: io.clbs.openhes.services.svcapi.ApiService.GetBulk:input_type -> google.protobuf.StringValue
	1,  // 25: io.clbs.openhes.services.svcapi.ApiService.ListDrivers:input_type -> io.clbs.openhes.models.common.ListSelector
	3,  // 26: io.clbs.openhes.services.svcapi.ApiService.GetDriver:input_type -> google.protobuf.StringValue
	14, // 27: io.clbs.openhes.services.svcapi.ApiService.CreateCommunicationUnit:input_type -> io.clbs.openhes.models.acquisition.CreateCommunicationUnitRequest
	1,  // 28: io.clbs.openhes.services.svcapi.ApiService.ListCommunicationUnits:input_type -> io.clbs.openhes.models.common.ListSelector
	3,  // 29: io.clbs.openhes.services.svcapi.ApiService.GetCommunicationUnit:input_type -> google.protobuf.StringValue
	15, // 30: io.clbs.openhes.services.svcapi.ApiService.CreateCommunicationBus:input_type -> io.clbs.openhes.models.acquisition.CreateCommunicationBusRequest
	1,  // 31: io.clbs.openhes.services.svcapi.ApiService.ListCommunicationBuses:input_type -> io.clbs.openhes.models.common.ListSelector
	16, // 32: io.clbs.openhes.services.svcapi.ApiService.AddCommunicationUnitsToCommunicationBus:input_type -> io.clbs.openhes.models.acquisition.AddCommunicationUnitsToCommunicationBusRequest
	17, // 33: io.clbs.openhes.services.svcapi.ApiService.RemoveCommunicationUnitsFromCommunicationBus:input_type -> io.clbs.openhes.models.acquisition.RemoveCommunicationUnitsFromCommunicationBusRequest
	18, // 34: io.clbs.openhes.services.svcapi.ApiService.CreateDevice:input_type -> io.clbs.openhes.models.acquisition.CreateDeviceRequest
	19, // 35: io.clbs.openhes.services.svcapi.ApiService.UpdateDevice:input_type -> io.clbs.openhes.models.acquisition.Device
	1,  // 36: io.clbs.openhes.services.svcapi.ApiService.ListDevices:input_type -> io.clbs.openhes.models.common.ListSelector
	3,  // 37: io.clbs.openhes.services.svcapi.ApiService.GetDevice:input_type -> google.protobuf.StringValue
	3,  // 38: io.clbs.openhes.services.svcapi.ApiService.GetDeviceInfo:input_type -> google.protobuf.StringValue
	20, // 39: io.clbs.openhes.services.svcapi.ApiService.SetDeviceCommunicationUnits:input_type -> io.clbs.openhes.models.acquisition.SetDeviceCommunicationUnitsRequest
	3,  // 40: io.clbs.openhes.services.svcapi.ApiService.GetDeviceCommunicationUnits:input_type -> google.protobuf.StringValue
	21, // 41: io.clbs.openhes.services.svcapi.ApiService.CreateDeviceGroup:input_type -> io.clbs.openhes.models.acquisition.CreateDeviceGroupRequest
	1,  // 42: io.clbs.openhes.services.svcapi.ApiService.ListDeviceGroups:input_type -> io.clbs.openhes.models.common.ListSelector
	3,  // 43: io.clbs.openhes.services.svcapi.ApiService.GetDeviceGroup:input_type -> google.protobuf.StringValue
	22, // 44: io.clbs.openhes.services.svcapi.ApiService.AddDevicesToGroup:input_type -> io.clbs.openhes.models.acquisition.AddDevicesToGroupRequest
	23, // 45: io.clbs.openhes.services.svcapi.ApiService.RemoveDevicesFromGroup:input_type -> io.clbs.openhes.models.acquisition.RemoveDevicesFromGroupRequest
	24, // 46: io.clbs.openhes.services.svcapi.ApiService.ListDeviceGroupDevices:input_type -> io.clbs.openhes.models.acquisition.ListDeviceGroupDevicesRequest
	1,  // 47: io.clbs.openhes.services.svcapi.ApiService.ListModemPools:input_type -> io.clbs.openhes.models.common.ListSelector
	3,  // 48: io.clbs.openhes.services.svcapi.ApiService.GetModemPool:input_type -> google.protobuf.StringValue
	25, // 49: io.clbs.openhes.services.svcapi.ApiService.CreateModemPool:input_type -> io.clbs.openhes.models.acquisition.SetModemPoolRequest
	25, // 50: io.clbs.openhes.services.svcapi.ApiService.UpdateModemPool:input_type -> io.clbs.openhes.models.acquisition.SetModemPoolRequest
	3,  // 51: io.clbs.openhes.services.svcapi.ApiService.DeleteModemPool:input_type -> google.protobuf.StringValue
	26, // 52: io.clbs.openhes.services.svcapi.ApiService.CreateModem:input_type -> io.clbs.openhes.models.acquisition.SetModemRequest
	26, // 53: io.clbs.openhes.services.svcapi.ApiService.UpdateModem:input_type -> io.clbs.openhes.models.acquisition.SetModemRequest
	3,  // 54: io.clbs.openhes.services.svcapi.ApiService.DeleteModem:input_type -> google.protobuf.StringValue
	10, // 55: io.clbs.openhes.services.svcapi.ApiService.GetConfig:input_type -> google.protobuf.Empty
	27, // 56: io.clbs.openhes.services.svcapi.ApiService.SetConfig:input_type -> io.clbs.openhes.models.system.SystemConfig
	28, // 57: io.clbs.openhes.services.svcapi.ApiService.GetMeterDataRegisters:input_type -> io.clbs.openhes.models.acquisition.GetMeterDataRequest
	28, // 58: io.clbs.openhes.services.svcapi.ApiService.GetMeterDataProfiles:input_type -> io.clbs.openhes.models.acquisition.GetMeterDataRequest
	28, // 59: io.clbs.openhes.services.svcapi.ApiService.GetMeterDataIrregularProfiles:input_type -> io.clbs.openhes.models.acquisition.GetMeterDataRequest
	29, // 60: io.clbs.openhes.services.svcapi.ApiService.GetMeterEvents:input_type -> io.clbs.openhes.models.acquisition.GetMeterEventsRequest
	30, // 61: io.clbs.openhes.services.svcapi.ApiService.CreateTimeOfUseTable:input_type -> io.clbs.openhes.models.acquisition.CreateTimeOfUseTableRequest
	1,  // 62: io.clbs.openhes.services.svcapi.ApiService.ListTimeOfUseTables:input_type -> io.clbs.openhes.models.common.ListSelector
	3,  // 63: io.clbs.openhes.services.svcapi.ApiService.GetTimeOfUseTable:input_type -> google.protobuf.StringValue
	31, // 64: io.clbs.openhes.services.svcapi.ApiService.UpdateTimeOfUseTable:input_type -> io.clbs.openhes.models.acquisition.TimeOfUseTable
	3,  // 65: io.clbs.openhes.services.svcapi.ApiService.DeleteTimeOfUseTable:input_type -> google.protobuf.StringValue
	3,  // 66: io.clbs.openhes.services.svcapi.ApiService.CreateVariable:output_type -> google.protobuf.StringValue
	32, // 67: io.clbs.openhes.services.svcapi.ApiService.ListVariables:output_type -> io.clbs.openhes.models.acquisition.ListOfVariable
	10, // 68: io.clbs.openhes.services.svcapi.ApiService.UpdateVariable:output_type -> google.protobuf.Empty
	10, // 69: io.clbs.openhes.services.svcapi.ApiService.DeleteVariable:output_type -> google.protobuf.Empty
	3,  // 70: io.clbs.openhes.services.svcapi.ApiService.CreateDeviceConfigurationRegister:output_type -> google.protobuf.StringValue
	33, // 71: io.clbs.openhes.services.svcapi.ApiService.ListDeviceConfigurationRegisters:output_type -> io.clbs.openhes.models.acquisition.ListOfDeviceConfigurationRegister
	5,  // 72: io.clbs.openhes.services.svcapi.ApiService.GetDeviceConfigurationRegister:output_type -> io.clbs.openhes.models.acquisition.DeviceConfigurationRegister
	10, // 73: io.clbs.openhes.services.svcapi.ApiService.UpdateDeviceConfigurationRegister:output_type -> google.protobuf.Empty
	10, // 74: io.clbs.openhes.services.svcapi.ApiService.DeleteDeviceConfigurationRegister:output_type -> google.protobuf.Empty
	3,  // 75: io.clbs.openhes.services.svcapi.ApiService.CreateDeviceConfigurationTemplate:output_type -> google.protobuf.StringValue
	34, // 76: io.clbs.openhes.services.svcapi.ApiService.ListDeviceConfigurationTemplates:output_type -> io.clbs.openhes.models.acquisition.ListOfDeviceConfigurationTemplate
	7,  // 77: io.clbs.openhes.services.svcapi.ApiService.GetDeviceConfigurationTemplate:output_type -> io.clbs.openhes.models.acquisition.DeviceConfigurationTemplate
	10, // 78: io.clbs.openhes.services.svcapi.ApiService.UpdateDeviceConfigurationTemplate:output_type -> google.protobuf.Empty
	10, // 79: io.clbs.openhes.services.svcapi.ApiService.DeleteDeviceConfigurationTemplate:output_type -> google.protobuf.Empty
	10, // 80: io.clbs.openhes.services.svcapi.ApiService.AddDeviceConfigurationRegisterToDeviceConfigurationTemplate:output_type -> google.protobuf.Empty
	10, // 81: io.clbs.openhes.services.svcapi.ApiService.RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplate:output_type -> google.protobuf.Empty
	35, // 82: io.clbs.openhes.services.svcapi.ApiService.ListFieldDescriptors:output_type -> io.clbs.openhes.models.common.ListOfFieldDescriptor
	36, // 83: io.clbs.openhes.services.svcapi.ApiService.ListBulks:output_type -> io.clbs.openhes.models.acquisition.ListOfBulk
	37, // 84: io.clbs.openhes.services.svcapi.ApiService.ListBulkJobs:output_type -> io.clbs.openhes.models.acquisition.ListOfBulkJob
	38, // 85: io.clbs.openhes.services.svcapi.ApiService.GetBulkJob:output_type -> io.clbs.openhes.models.acquisition.BulkJob
	10, // 86: io.clbs.openhes.services.svcapi.ApiService.CancelBulk:output_type -> google.protobuf.Empty
	3,  // 87: io.clbs.openhes.services.svcapi.ApiService.CreateProxyBulk:output_type -> google.protobuf.StringValue
	39, // 88: io.clbs.openhes.services.svcapi.ApiService.GetProxyBulk:output_type -> io.clbs.openhes.models.acquisition.ProxyBulk
	3,  // 89: io.clbs.openhes.services.svcapi.ApiService.CreateBulk:output_type -> google.protobuf.StringValue
	40, // 90: io.clbs.openhes.services.svcapi.ApiService.GetBulk:output_type -> io.clbs.openhes.models.acquisition.Bulk
	41, // 91: io.clbs.openhes.services.svcapi.ApiService.ListDrivers:output_type -> io.clbs.openhes.models.acquisition.ListOfDriver
	42, // 92: io.clbs.openhes.services.svcapi.ApiService.GetDriver:output_type -> io.clbs.openhes.models.acquisition.Driver
	3,  // 93: io.clbs.openhes.services.svcapi.ApiService.CreateCommunicationUnit:output_type -> google.protobuf.StringValue
	43, // 94: io.clbs.openhes.services.svcapi.ApiService.ListCommunicationUnits:output_type -> io.clbs.openhes.models.acquisition.ListOfCommunicationUnit
	44, // 95: io.clbs.openhes.services.svcapi.ApiService.GetCommunicationUnit:output_type -> io.clbs.openhes.models.acquisition.CommunicationUnit
	3,  // 96: io.clbs.openhes.services.svcapi.ApiService.CreateCommunicationBus:output_type -> google.protobuf.StringValue
	45, // 97: io.clbs.openhes.services.svcapi.ApiService.ListCommunicationBuses:output_type -> io.clbs.openhes.models.acquisition.ListOfCommunicationBus
	10, // 98: io.clbs.openhes.services.svcapi.ApiService.AddCommunicationUnitsToCommunicationBus:output_type -> google.protobuf.Empty
	10, // 99: io.clbs.openhes.services.svcapi.ApiService.RemoveCommunicationUnitsFromCommunicationBus:output_type -> google.protobuf.Empty
	3,  // 100: io.clbs.openhes.services.svcapi.ApiService.CreateDevice:output_type -> google.protobuf.StringValue
	10, // 101: io.clbs.openhes.services.svcapi.ApiService.UpdateDevice:output_type -> google.protobuf.Empty
	46, // 102: io.clbs.openhes.services.svcapi.ApiService.ListDevices:output_type -> io.clbs.openhes.models.acquisition.ListOfDevice
	19, // 103: io.clbs.openhes.services.svcapi.ApiService.GetDevice:output_type -> io.clbs.openhes.models.acquisition.Device
	47, // 104: io.clbs.openhes.services.svcapi.ApiService.GetDeviceInfo:output_type -> io.clbs.openhes.models.acquisition.DeviceInfo
	10, // 105: io.clbs.openhes.services.svcapi.ApiService.SetDeviceCommunicationUnits:output_type -> google.protobuf.Empty
	48, // 106: io.clbs.openhes.services.svcapi.ApiService.GetDeviceCommunicationUnits:output_type -> io.clbs.openhes.models.acquisition.ListOfDeviceCommunicationUnit
	3,  // 107: io.clbs.openhes.services.svcapi.ApiService.CreateDeviceGroup:output_type -> google.protobuf.StringValue
	49, // 108: io.clbs.openhes.services.svcapi.ApiService.ListDeviceGroups:output_type -> io.clbs.openhes.models.acquisition.ListOfDeviceGroup
	50, // 109: io.clbs.openhes.services.svcapi.ApiService.GetDeviceGroup:output_type -> io.clbs.openhes.models.acquisition.DeviceGroup
	10, // 110: io.clbs.openhes.services.svcapi.ApiService.AddDevicesToGroup:output_type -> google.protobuf.Empty
	10, // 111: io.clbs.openhes.services.svcapi.ApiService.RemoveDevicesFromGroup:output_type -> google.protobuf.Empty
	46, // 112: io.clbs.openhes.services.svcapi.ApiService.ListDeviceGroupDevices:output_type -> io.clbs.openhes.models.acquisition.ListOfDevice
	51, // 113: io.clbs.openhes.services.svcapi.ApiService.ListModemPools:output_type -> io.clbs.openhes.models.acquisition.ListOfModemPool
	52, // 114: io.clbs.openhes.services.svcapi.ApiService.GetModemPool:output_type -> io.clbs.openhes.models.acquisition.ModemPool
	3,  // 115: io.clbs.openhes.services.svcapi.ApiService.CreateModemPool:output_type -> google.protobuf.StringValue
	10, // 116: io.clbs.openhes.services.svcapi.ApiService.UpdateModemPool:output_type -> google.protobuf.Empty
	10, // 117: io.clbs.openhes.services.svcapi.ApiService.DeleteModemPool:output_type -> google.protobuf.Empty
	3,  // 118: io.clbs.openhes.services.svcapi.ApiService.CreateModem:output_type -> google.protobuf.StringValue
	10, // 119: io.clbs.openhes.services.svcapi.ApiService.UpdateModem:output_type -> google.protobuf.Empty
	10, // 120: io.clbs.openhes.services.svcapi.ApiService.DeleteModem:output_type -> google.protobuf.Empty
	27, // 121: io.clbs.openhes.services.svcapi.ApiService.GetConfig:output_type -> io.clbs.openhes.models.system.SystemConfig
	10, // 122: io.clbs.openhes.services.svcapi.ApiService.SetConfig:output_type -> google.protobuf.Empty
	53, // 123: io.clbs.openhes.services.svcapi.ApiService.GetMeterDataRegisters:output_type -> io.clbs.openhes.models.acquisition.RegisterValues
	54, // 124: io.clbs.openhes.services.svcapi.ApiService.GetMeterDataProfiles:output_type -> io.clbs.openhes.models.acquisition.ProfileValues
	55, // 125: io.clbs.openhes.services.svcapi.ApiService.GetMeterDataIrregularProfiles:output_type -> io.clbs.openhes.models.acquisition.IrregularProfileValues
	56, // 126: io.clbs.openhes.services.svcapi.ApiService.GetMeterEvents:output_type -> io.clbs.openhes.models.acquisition.EventRecords
	3,  // 127: io.clbs.openhes.services.svcapi.ApiService.CreateTimeOfUseTable:output_type -> google.protobuf.StringValue
	57, // 128: io.clbs.openhes.services.svcapi.ApiService.ListTimeOfUseTables:output_type -> io.clbs.openhes.models.acquisition.ListOfTimeOfUseTable
	31, // 129: io.clbs.openhes.services.svcapi.ApiService.GetTimeOfUseTable:output_type -> io.clbs.openhes.models.acquisition.TimeOfUseTable
	10, // 130: io.clbs.openhes.services.svcapi.ApiService.UpdateTimeOfUseTable:output_type -> google.protobuf.Empty
	10, // 131: io.clbs.openhes.services.svcapi.ApiService.DeleteTimeOfUseTable:output_type -> google.protobuf.Empty
	66, // [66:132] is the sub-list for method output_type
	0,  // [0:66] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_services_svcapi_api_proto_init() }
func file_services_svcapi_api_proto_init() {
	if File_services_svcapi_api_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_svcapi_api_proto_rawDesc), len(file_services_svcapi_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_svcapi_api_proto_goTypes,
		DependencyIndexes: file_services_svcapi_api_proto_depIdxs,
	}.Build()
	File_services_svcapi_api_proto = out.File
	file_services_svcapi_api_proto_goTypes = nil
	file_services_svcapi_api_proto_depIdxs = nil
}
