// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: services/svcdeviceregistry/deviceregistry.proto

package svcdeviceregistry

import (
	acquisition "github.com/cybroslabs/hes-2-apis/gen/go/acquisition"
	common "github.com/cybroslabs/hes-2-apis/gen/go/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
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

var File_services_svcdeviceregistry_deviceregistry_proto protoreflect.FileDescriptor

const file_services_svcdeviceregistry_deviceregistry_proto_rawDesc = "" +
	"\n" +
	"/services/svcdeviceregistry/deviceregistry.proto\x12*io.clbs.openhes.services.svcdeviceregistry\x1a\x1bgoogle/protobuf/empty.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x16acquisition/main.proto\x1a\x1aacquisition/internal.proto\x1a\x18acquisition/shared.proto\x1a\x13common/fields.proto2\xea+\n" +
	"\x15DeviceRegistryService\x12i\n" +
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
	"@RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplate\x12k.io.clbs.openhes.models.acquisition.RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplateRequest\x1a\x16.google.protobuf.Empty\x12W\n" +
	"\vListDrivers\x12\x16.google.protobuf.Empty\x1a0.io.clbs.openhes.models.acquisition.ListOfDriver\x12U\n" +
	"\fCreateDriver\x12-.io.clbs.openhes.models.acquisition.SetDriver\x1a\x16.google.protobuf.Empty\x12U\n" +
	"\tGetDriver\x12\x1c.google.protobuf.StringValue\x1a*.io.clbs.openhes.models.acquisition.Driver\x12{\n" +
	"\x17CreateCommunicationUnit\x12B.io.clbs.openhes.models.acquisition.CreateCommunicationUnitRequest\x1a\x1c.google.protobuf.StringValue\x12\x82\x01\n" +
	"\x16ListCommunicationUnits\x12+.io.clbs.openhes.models.common.ListSelector\x1a;.io.clbs.openhes.models.acquisition.ListOfCommunicationUnit\x12k\n" +
	"\x14GetCommunicationUnit\x12\x1c.google.protobuf.StringValue\x1a5.io.clbs.openhes.models.acquisition.CommunicationUnit\x12y\n" +
	"\x16CreateCommunicationBus\x12A.io.clbs.openhes.models.acquisition.CreateCommunicationBusRequest\x1a\x1c.google.protobuf.StringValue\x12\x81\x01\n" +
	"\x16ListCommunicationBuses\x12+.io.clbs.openhes.models.common.ListSelector\x1a:.io.clbs.openhes.models.acquisition.ListOfCommunicationBus\x12\x95\x01\n" +
	"'AddCommunicationUnitsToCommunicationBus\x12R.io.clbs.openhes.models.acquisition.AddCommunicationUnitsToCommunicationBusRequest\x1a\x16.google.protobuf.Empty\x12\x9f\x01\n" +
	",RemoveCommunicationUnitsFromCommunicationBus\x12W.io.clbs.openhes.models.acquisition.RemoveCommunicationUnitsFromCommunicationBusRequest\x1a\x16.google.protobuf.Empty\x12e\n" +
	"\fCreateDevice\x127.io.clbs.openhes.models.acquisition.CreateDeviceRequest\x1a\x1c.google.protobuf.StringValue\x12l\n" +
	"\vListDevices\x12+.io.clbs.openhes.models.common.ListSelector\x1a0.io.clbs.openhes.models.acquisition.ListOfDevice\x12U\n" +
	"\tGetDevice\x12\x1c.google.protobuf.StringValue\x1a*.io.clbs.openhes.models.acquisition.Device\x12}\n" +
	"\x1bSetDeviceCommunicationUnits\x12F.io.clbs.openhes.models.acquisition.SetDeviceCommunicationUnitsRequest\x1a\x16.google.protobuf.Empty\x12~\n" +
	"\x1bGetDeviceCommunicationUnits\x12\x1c.google.protobuf.StringValue\x1aA.io.clbs.openhes.models.acquisition.ListOfDeviceCommunicationUnit\x12r\n" +
	"\x17GetDeviceConnectionInfo\x12\x1a.google.protobuf.ListValue\x1a;.io.clbs.openhes.models.acquisition.MapDeviceConnectionInfo\x12a\n" +
	"\rSetDeviceInfo\x128.io.clbs.openhes.models.acquisition.SetDeviceInfoRequest\x1a\x16.google.protobuf.Empty\x12o\n" +
	"\x11CreateDeviceGroup\x12<.io.clbs.openhes.models.acquisition.CreateDeviceGroupRequest\x1a\x1c.google.protobuf.StringValue\x12v\n" +
	"\x10ListDeviceGroups\x12+.io.clbs.openhes.models.common.ListSelector\x1a5.io.clbs.openhes.models.acquisition.ListOfDeviceGroup\x12_\n" +
	"\x0eGetDeviceGroup\x12\x1c.google.protobuf.StringValue\x1a/.io.clbs.openhes.models.acquisition.DeviceGroup\x12j\n" +
	"\x11StreamDeviceGroup\x12\x1c.google.protobuf.StringValue\x1a5.io.clbs.openhes.models.acquisition.StreamDeviceGroup0\x01\x12i\n" +
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
	"\vDeleteModem\x12\x1c.google.protobuf.StringValue\x1a\x16.google.protobuf.EmptyBDZBgithub.com/cybroslabs/hes-2-apis/gen/go/services/svcdeviceregistryb\beditionsp\xe8\a"

var file_services_svcdeviceregistry_deviceregistry_proto_goTypes = []any{
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
	(*acquisition.SetDriver)(nil),                                                               // 11: io.clbs.openhes.models.acquisition.SetDriver
	(*acquisition.CreateCommunicationUnitRequest)(nil),                                          // 12: io.clbs.openhes.models.acquisition.CreateCommunicationUnitRequest
	(*acquisition.CreateCommunicationBusRequest)(nil),                                           // 13: io.clbs.openhes.models.acquisition.CreateCommunicationBusRequest
	(*acquisition.AddCommunicationUnitsToCommunicationBusRequest)(nil),                          // 14: io.clbs.openhes.models.acquisition.AddCommunicationUnitsToCommunicationBusRequest
	(*acquisition.RemoveCommunicationUnitsFromCommunicationBusRequest)(nil),                     // 15: io.clbs.openhes.models.acquisition.RemoveCommunicationUnitsFromCommunicationBusRequest
	(*acquisition.CreateDeviceRequest)(nil),                                                     // 16: io.clbs.openhes.models.acquisition.CreateDeviceRequest
	(*acquisition.SetDeviceCommunicationUnitsRequest)(nil),                                      // 17: io.clbs.openhes.models.acquisition.SetDeviceCommunicationUnitsRequest
	(*structpb.ListValue)(nil),                                                                  // 18: google.protobuf.ListValue
	(*acquisition.SetDeviceInfoRequest)(nil),                                                    // 19: io.clbs.openhes.models.acquisition.SetDeviceInfoRequest
	(*acquisition.CreateDeviceGroupRequest)(nil),                                                // 20: io.clbs.openhes.models.acquisition.CreateDeviceGroupRequest
	(*acquisition.AddDevicesToGroupRequest)(nil),                                                // 21: io.clbs.openhes.models.acquisition.AddDevicesToGroupRequest
	(*acquisition.RemoveDevicesFromGroupRequest)(nil),                                           // 22: io.clbs.openhes.models.acquisition.RemoveDevicesFromGroupRequest
	(*acquisition.ListDeviceGroupDevicesRequest)(nil),                                           // 23: io.clbs.openhes.models.acquisition.ListDeviceGroupDevicesRequest
	(*acquisition.SetModemPoolRequest)(nil),                                                     // 24: io.clbs.openhes.models.acquisition.SetModemPoolRequest
	(*acquisition.SetModemRequest)(nil),                                                         // 25: io.clbs.openhes.models.acquisition.SetModemRequest
	(*acquisition.ListOfVariable)(nil),                                                          // 26: io.clbs.openhes.models.acquisition.ListOfVariable
	(*acquisition.ListOfDeviceConfigurationRegister)(nil),                                       // 27: io.clbs.openhes.models.acquisition.ListOfDeviceConfigurationRegister
	(*acquisition.ListOfDeviceConfigurationTemplate)(nil),                                       // 28: io.clbs.openhes.models.acquisition.ListOfDeviceConfigurationTemplate
	(*acquisition.ListOfDriver)(nil),                                                            // 29: io.clbs.openhes.models.acquisition.ListOfDriver
	(*acquisition.Driver)(nil),                                                                  // 30: io.clbs.openhes.models.acquisition.Driver
	(*acquisition.ListOfCommunicationUnit)(nil),                                                 // 31: io.clbs.openhes.models.acquisition.ListOfCommunicationUnit
	(*acquisition.CommunicationUnit)(nil),                                                       // 32: io.clbs.openhes.models.acquisition.CommunicationUnit
	(*acquisition.ListOfCommunicationBus)(nil),                                                  // 33: io.clbs.openhes.models.acquisition.ListOfCommunicationBus
	(*acquisition.ListOfDevice)(nil),                                                            // 34: io.clbs.openhes.models.acquisition.ListOfDevice
	(*acquisition.Device)(nil),                                                                  // 35: io.clbs.openhes.models.acquisition.Device
	(*acquisition.ListOfDeviceCommunicationUnit)(nil),                                           // 36: io.clbs.openhes.models.acquisition.ListOfDeviceCommunicationUnit
	(*acquisition.MapDeviceConnectionInfo)(nil),                                                 // 37: io.clbs.openhes.models.acquisition.MapDeviceConnectionInfo
	(*acquisition.ListOfDeviceGroup)(nil),                                                       // 38: io.clbs.openhes.models.acquisition.ListOfDeviceGroup
	(*acquisition.DeviceGroup)(nil),                                                             // 39: io.clbs.openhes.models.acquisition.DeviceGroup
	(*acquisition.StreamDeviceGroup)(nil),                                                       // 40: io.clbs.openhes.models.acquisition.StreamDeviceGroup
	(*acquisition.ListOfModemPool)(nil),                                                         // 41: io.clbs.openhes.models.acquisition.ListOfModemPool
	(*acquisition.ModemPool)(nil),                                                               // 42: io.clbs.openhes.models.acquisition.ModemPool
}
var file_services_svcdeviceregistry_deviceregistry_proto_depIdxs = []int32{
	0,  // 0: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateVariable:input_type -> io.clbs.openhes.models.acquisition.CreateVariableRequest
	1,  // 1: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListVariables:input_type -> io.clbs.openhes.models.common.ListSelector
	2,  // 2: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.UpdateVariable:input_type -> io.clbs.openhes.models.acquisition.Variable
	3,  // 3: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.DeleteVariable:input_type -> google.protobuf.StringValue
	4,  // 4: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateDeviceConfigurationRegister:input_type -> io.clbs.openhes.models.acquisition.CreateDeviceConfigurationRegisterRequest
	1,  // 5: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListDeviceConfigurationRegisters:input_type -> io.clbs.openhes.models.common.ListSelector
	3,  // 6: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetDeviceConfigurationRegister:input_type -> google.protobuf.StringValue
	5,  // 7: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.UpdateDeviceConfigurationRegister:input_type -> io.clbs.openhes.models.acquisition.DeviceConfigurationRegister
	3,  // 8: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.DeleteDeviceConfigurationRegister:input_type -> google.protobuf.StringValue
	6,  // 9: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateDeviceConfigurationTemplate:input_type -> io.clbs.openhes.models.acquisition.CreateDeviceConfigurationTemplateRequest
	1,  // 10: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListDeviceConfigurationTemplates:input_type -> io.clbs.openhes.models.common.ListSelector
	3,  // 11: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetDeviceConfigurationTemplate:input_type -> google.protobuf.StringValue
	7,  // 12: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.UpdateDeviceConfigurationTemplate:input_type -> io.clbs.openhes.models.acquisition.DeviceConfigurationTemplate
	3,  // 13: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.DeleteDeviceConfigurationTemplate:input_type -> google.protobuf.StringValue
	8,  // 14: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.AddDeviceConfigurationRegisterToDeviceConfigurationTemplate:input_type -> io.clbs.openhes.models.acquisition.AddDeviceConfigurationRegisterToDeviceConfigurationTemplateRequest
	9,  // 15: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplate:input_type -> io.clbs.openhes.models.acquisition.RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplateRequest
	10, // 16: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListDrivers:input_type -> google.protobuf.Empty
	11, // 17: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateDriver:input_type -> io.clbs.openhes.models.acquisition.SetDriver
	3,  // 18: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetDriver:input_type -> google.protobuf.StringValue
	12, // 19: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateCommunicationUnit:input_type -> io.clbs.openhes.models.acquisition.CreateCommunicationUnitRequest
	1,  // 20: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListCommunicationUnits:input_type -> io.clbs.openhes.models.common.ListSelector
	3,  // 21: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetCommunicationUnit:input_type -> google.protobuf.StringValue
	13, // 22: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateCommunicationBus:input_type -> io.clbs.openhes.models.acquisition.CreateCommunicationBusRequest
	1,  // 23: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListCommunicationBuses:input_type -> io.clbs.openhes.models.common.ListSelector
	14, // 24: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.AddCommunicationUnitsToCommunicationBus:input_type -> io.clbs.openhes.models.acquisition.AddCommunicationUnitsToCommunicationBusRequest
	15, // 25: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.RemoveCommunicationUnitsFromCommunicationBus:input_type -> io.clbs.openhes.models.acquisition.RemoveCommunicationUnitsFromCommunicationBusRequest
	16, // 26: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateDevice:input_type -> io.clbs.openhes.models.acquisition.CreateDeviceRequest
	1,  // 27: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListDevices:input_type -> io.clbs.openhes.models.common.ListSelector
	3,  // 28: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetDevice:input_type -> google.protobuf.StringValue
	17, // 29: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.SetDeviceCommunicationUnits:input_type -> io.clbs.openhes.models.acquisition.SetDeviceCommunicationUnitsRequest
	3,  // 30: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetDeviceCommunicationUnits:input_type -> google.protobuf.StringValue
	18, // 31: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetDeviceConnectionInfo:input_type -> google.protobuf.ListValue
	19, // 32: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.SetDeviceInfo:input_type -> io.clbs.openhes.models.acquisition.SetDeviceInfoRequest
	20, // 33: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateDeviceGroup:input_type -> io.clbs.openhes.models.acquisition.CreateDeviceGroupRequest
	1,  // 34: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListDeviceGroups:input_type -> io.clbs.openhes.models.common.ListSelector
	3,  // 35: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetDeviceGroup:input_type -> google.protobuf.StringValue
	3,  // 36: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.StreamDeviceGroup:input_type -> google.protobuf.StringValue
	21, // 37: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.AddDevicesToGroup:input_type -> io.clbs.openhes.models.acquisition.AddDevicesToGroupRequest
	22, // 38: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.RemoveDevicesFromGroup:input_type -> io.clbs.openhes.models.acquisition.RemoveDevicesFromGroupRequest
	23, // 39: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListDeviceGroupDevices:input_type -> io.clbs.openhes.models.acquisition.ListDeviceGroupDevicesRequest
	1,  // 40: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListModemPools:input_type -> io.clbs.openhes.models.common.ListSelector
	3,  // 41: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetModemPool:input_type -> google.protobuf.StringValue
	24, // 42: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateModemPool:input_type -> io.clbs.openhes.models.acquisition.SetModemPoolRequest
	24, // 43: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.UpdateModemPool:input_type -> io.clbs.openhes.models.acquisition.SetModemPoolRequest
	3,  // 44: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.DeleteModemPool:input_type -> google.protobuf.StringValue
	25, // 45: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateModem:input_type -> io.clbs.openhes.models.acquisition.SetModemRequest
	25, // 46: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.UpdateModem:input_type -> io.clbs.openhes.models.acquisition.SetModemRequest
	3,  // 47: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.DeleteModem:input_type -> google.protobuf.StringValue
	3,  // 48: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateVariable:output_type -> google.protobuf.StringValue
	26, // 49: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListVariables:output_type -> io.clbs.openhes.models.acquisition.ListOfVariable
	10, // 50: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.UpdateVariable:output_type -> google.protobuf.Empty
	10, // 51: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.DeleteVariable:output_type -> google.protobuf.Empty
	3,  // 52: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateDeviceConfigurationRegister:output_type -> google.protobuf.StringValue
	27, // 53: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListDeviceConfigurationRegisters:output_type -> io.clbs.openhes.models.acquisition.ListOfDeviceConfigurationRegister
	5,  // 54: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetDeviceConfigurationRegister:output_type -> io.clbs.openhes.models.acquisition.DeviceConfigurationRegister
	10, // 55: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.UpdateDeviceConfigurationRegister:output_type -> google.protobuf.Empty
	10, // 56: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.DeleteDeviceConfigurationRegister:output_type -> google.protobuf.Empty
	3,  // 57: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateDeviceConfigurationTemplate:output_type -> google.protobuf.StringValue
	28, // 58: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListDeviceConfigurationTemplates:output_type -> io.clbs.openhes.models.acquisition.ListOfDeviceConfigurationTemplate
	7,  // 59: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetDeviceConfigurationTemplate:output_type -> io.clbs.openhes.models.acquisition.DeviceConfigurationTemplate
	10, // 60: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.UpdateDeviceConfigurationTemplate:output_type -> google.protobuf.Empty
	10, // 61: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.DeleteDeviceConfigurationTemplate:output_type -> google.protobuf.Empty
	10, // 62: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.AddDeviceConfigurationRegisterToDeviceConfigurationTemplate:output_type -> google.protobuf.Empty
	10, // 63: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.RemoveDeviceConfigurationRegisterFromDeviceConfigurationTemplate:output_type -> google.protobuf.Empty
	29, // 64: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListDrivers:output_type -> io.clbs.openhes.models.acquisition.ListOfDriver
	10, // 65: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateDriver:output_type -> google.protobuf.Empty
	30, // 66: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetDriver:output_type -> io.clbs.openhes.models.acquisition.Driver
	3,  // 67: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateCommunicationUnit:output_type -> google.protobuf.StringValue
	31, // 68: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListCommunicationUnits:output_type -> io.clbs.openhes.models.acquisition.ListOfCommunicationUnit
	32, // 69: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetCommunicationUnit:output_type -> io.clbs.openhes.models.acquisition.CommunicationUnit
	3,  // 70: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateCommunicationBus:output_type -> google.protobuf.StringValue
	33, // 71: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListCommunicationBuses:output_type -> io.clbs.openhes.models.acquisition.ListOfCommunicationBus
	10, // 72: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.AddCommunicationUnitsToCommunicationBus:output_type -> google.protobuf.Empty
	10, // 73: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.RemoveCommunicationUnitsFromCommunicationBus:output_type -> google.protobuf.Empty
	3,  // 74: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateDevice:output_type -> google.protobuf.StringValue
	34, // 75: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListDevices:output_type -> io.clbs.openhes.models.acquisition.ListOfDevice
	35, // 76: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetDevice:output_type -> io.clbs.openhes.models.acquisition.Device
	10, // 77: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.SetDeviceCommunicationUnits:output_type -> google.protobuf.Empty
	36, // 78: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetDeviceCommunicationUnits:output_type -> io.clbs.openhes.models.acquisition.ListOfDeviceCommunicationUnit
	37, // 79: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetDeviceConnectionInfo:output_type -> io.clbs.openhes.models.acquisition.MapDeviceConnectionInfo
	10, // 80: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.SetDeviceInfo:output_type -> google.protobuf.Empty
	3,  // 81: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateDeviceGroup:output_type -> google.protobuf.StringValue
	38, // 82: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListDeviceGroups:output_type -> io.clbs.openhes.models.acquisition.ListOfDeviceGroup
	39, // 83: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetDeviceGroup:output_type -> io.clbs.openhes.models.acquisition.DeviceGroup
	40, // 84: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.StreamDeviceGroup:output_type -> io.clbs.openhes.models.acquisition.StreamDeviceGroup
	10, // 85: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.AddDevicesToGroup:output_type -> google.protobuf.Empty
	10, // 86: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.RemoveDevicesFromGroup:output_type -> google.protobuf.Empty
	34, // 87: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListDeviceGroupDevices:output_type -> io.clbs.openhes.models.acquisition.ListOfDevice
	41, // 88: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.ListModemPools:output_type -> io.clbs.openhes.models.acquisition.ListOfModemPool
	42, // 89: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.GetModemPool:output_type -> io.clbs.openhes.models.acquisition.ModemPool
	3,  // 90: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateModemPool:output_type -> google.protobuf.StringValue
	10, // 91: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.UpdateModemPool:output_type -> google.protobuf.Empty
	10, // 92: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.DeleteModemPool:output_type -> google.protobuf.Empty
	3,  // 93: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.CreateModem:output_type -> google.protobuf.StringValue
	10, // 94: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.UpdateModem:output_type -> google.protobuf.Empty
	10, // 95: io.clbs.openhes.services.svcdeviceregistry.DeviceRegistryService.DeleteModem:output_type -> google.protobuf.Empty
	48, // [48:96] is the sub-list for method output_type
	0,  // [0:48] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_services_svcdeviceregistry_deviceregistry_proto_init() }
func file_services_svcdeviceregistry_deviceregistry_proto_init() {
	if File_services_svcdeviceregistry_deviceregistry_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_services_svcdeviceregistry_deviceregistry_proto_rawDesc), len(file_services_svcdeviceregistry_deviceregistry_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_svcdeviceregistry_deviceregistry_proto_goTypes,
		DependencyIndexes: file_services_svcdeviceregistry_deviceregistry_proto_depIdxs,
	}.Build()
	File_services_svcdeviceregistry_deviceregistry_proto = out.File
	file_services_svcdeviceregistry_deviceregistry_proto_goTypes = nil
	file_services_svcdeviceregistry_deviceregistry_proto_depIdxs = nil
}
