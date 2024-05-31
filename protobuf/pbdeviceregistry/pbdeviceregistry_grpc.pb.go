// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: pbdeviceregistry.proto

package pbdeviceregistry

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DeviceRegistryService_CreateCommunicationUnit_FullMethodName     = "/io.clbs.openhes.pbdeviceregistry.DeviceRegistryService/CreateCommunicationUnit"
	DeviceRegistryService_GetCommunicationUnits_FullMethodName       = "/io.clbs.openhes.pbdeviceregistry.DeviceRegistryService/GetCommunicationUnits"
	DeviceRegistryService_CreateDevice_FullMethodName                = "/io.clbs.openhes.pbdeviceregistry.DeviceRegistryService/CreateDevice"
	DeviceRegistryService_GetDevices_FullMethodName                  = "/io.clbs.openhes.pbdeviceregistry.DeviceRegistryService/GetDevices"
	DeviceRegistryService_SetDeviceCommunicationUnits_FullMethodName = "/io.clbs.openhes.pbdeviceregistry.DeviceRegistryService/SetDeviceCommunicationUnits"
	DeviceRegistryService_CreateDeviceGroup_FullMethodName           = "/io.clbs.openhes.pbdeviceregistry.DeviceRegistryService/CreateDeviceGroup"
	DeviceRegistryService_GetDeviceGroups_FullMethodName             = "/io.clbs.openhes.pbdeviceregistry.DeviceRegistryService/GetDeviceGroups"
	DeviceRegistryService_AddDevicesToGroup_FullMethodName           = "/io.clbs.openhes.pbdeviceregistry.DeviceRegistryService/AddDevicesToGroup"
	DeviceRegistryService_RemoveDevicesFromGroup_FullMethodName      = "/io.clbs.openhes.pbdeviceregistry.DeviceRegistryService/RemoveDevicesFromGroup"
)

// DeviceRegistryServiceClient is the client API for DeviceRegistryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DeviceRegistryServiceClient interface {
	// The method called by the RestAPI to register a new communication unit. The parameter contains the communication unit specification.
	CreateCommunicationUnit(ctx context.Context, in *CreateCommunicationUnitRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// The method called by the RestAPI to get the information about the communication unit. The parameter contains the search criteria.
	GetCommunicationUnits(ctx context.Context, in *GetCommunicationUnitsRequest, opts ...grpc.CallOption) (*GetCommunicationUnitsResponse, error)
	// The method called by the RestAPI to register a new device. The parameter contains the device specification.
	CreateDevice(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// The method called by the RestAPI to get the information about the device. The parameter contains the search criteria.
	GetDevices(ctx context.Context, in *GetDevicesRequest, opts ...grpc.CallOption) (*GetDevicesResponse, error)
	// The method called by the RestAPI to replace ordered set of linked communication units.
	SetDeviceCommunicationUnits(ctx context.Context, in *SetDeviceCommunicationUnitsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// The method called by the RestAPI to create a new device group. The parameter contains the device group specification.
	CreateDeviceGroup(ctx context.Context, in *CreateDeviceGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// The method called by the RestAPI to get the information about the device group. The parameter contains the search criteria.
	GetDeviceGroups(ctx context.Context, in *GetDeviceGroupsRequest, opts ...grpc.CallOption) (*GetDeviceGroupsResponse, error)
	// The method called by the RestAPI to add a new device to the device group. The parameter contains the device group specification.
	AddDevicesToGroup(ctx context.Context, in *AddDevicesToGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// The method called by the RestAPI to remove a device from the device group. The parameter contains the device group specification.
	RemoveDevicesFromGroup(ctx context.Context, in *RemoveDevicesFromGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type deviceRegistryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceRegistryServiceClient(cc grpc.ClientConnInterface) DeviceRegistryServiceClient {
	return &deviceRegistryServiceClient{cc}
}

func (c *deviceRegistryServiceClient) CreateCommunicationUnit(ctx context.Context, in *CreateCommunicationUnitRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DeviceRegistryService_CreateCommunicationUnit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceRegistryServiceClient) GetCommunicationUnits(ctx context.Context, in *GetCommunicationUnitsRequest, opts ...grpc.CallOption) (*GetCommunicationUnitsResponse, error) {
	out := new(GetCommunicationUnitsResponse)
	err := c.cc.Invoke(ctx, DeviceRegistryService_GetCommunicationUnits_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceRegistryServiceClient) CreateDevice(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DeviceRegistryService_CreateDevice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceRegistryServiceClient) GetDevices(ctx context.Context, in *GetDevicesRequest, opts ...grpc.CallOption) (*GetDevicesResponse, error) {
	out := new(GetDevicesResponse)
	err := c.cc.Invoke(ctx, DeviceRegistryService_GetDevices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceRegistryServiceClient) SetDeviceCommunicationUnits(ctx context.Context, in *SetDeviceCommunicationUnitsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DeviceRegistryService_SetDeviceCommunicationUnits_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceRegistryServiceClient) CreateDeviceGroup(ctx context.Context, in *CreateDeviceGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DeviceRegistryService_CreateDeviceGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceRegistryServiceClient) GetDeviceGroups(ctx context.Context, in *GetDeviceGroupsRequest, opts ...grpc.CallOption) (*GetDeviceGroupsResponse, error) {
	out := new(GetDeviceGroupsResponse)
	err := c.cc.Invoke(ctx, DeviceRegistryService_GetDeviceGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceRegistryServiceClient) AddDevicesToGroup(ctx context.Context, in *AddDevicesToGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DeviceRegistryService_AddDevicesToGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceRegistryServiceClient) RemoveDevicesFromGroup(ctx context.Context, in *RemoveDevicesFromGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DeviceRegistryService_RemoveDevicesFromGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceRegistryServiceServer is the server API for DeviceRegistryService service.
// All implementations must embed UnimplementedDeviceRegistryServiceServer
// for forward compatibility
type DeviceRegistryServiceServer interface {
	// The method called by the RestAPI to register a new communication unit. The parameter contains the communication unit specification.
	CreateCommunicationUnit(context.Context, *CreateCommunicationUnitRequest) (*emptypb.Empty, error)
	// The method called by the RestAPI to get the information about the communication unit. The parameter contains the search criteria.
	GetCommunicationUnits(context.Context, *GetCommunicationUnitsRequest) (*GetCommunicationUnitsResponse, error)
	// The method called by the RestAPI to register a new device. The parameter contains the device specification.
	CreateDevice(context.Context, *CreateDeviceRequest) (*emptypb.Empty, error)
	// The method called by the RestAPI to get the information about the device. The parameter contains the search criteria.
	GetDevices(context.Context, *GetDevicesRequest) (*GetDevicesResponse, error)
	// The method called by the RestAPI to replace ordered set of linked communication units.
	SetDeviceCommunicationUnits(context.Context, *SetDeviceCommunicationUnitsRequest) (*emptypb.Empty, error)
	// The method called by the RestAPI to create a new device group. The parameter contains the device group specification.
	CreateDeviceGroup(context.Context, *CreateDeviceGroupRequest) (*emptypb.Empty, error)
	// The method called by the RestAPI to get the information about the device group. The parameter contains the search criteria.
	GetDeviceGroups(context.Context, *GetDeviceGroupsRequest) (*GetDeviceGroupsResponse, error)
	// The method called by the RestAPI to add a new device to the device group. The parameter contains the device group specification.
	AddDevicesToGroup(context.Context, *AddDevicesToGroupRequest) (*emptypb.Empty, error)
	// The method called by the RestAPI to remove a device from the device group. The parameter contains the device group specification.
	RemoveDevicesFromGroup(context.Context, *RemoveDevicesFromGroupRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedDeviceRegistryServiceServer()
}

// UnimplementedDeviceRegistryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDeviceRegistryServiceServer struct {
}

func (UnimplementedDeviceRegistryServiceServer) CreateCommunicationUnit(context.Context, *CreateCommunicationUnitRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommunicationUnit not implemented")
}
func (UnimplementedDeviceRegistryServiceServer) GetCommunicationUnits(context.Context, *GetCommunicationUnitsRequest) (*GetCommunicationUnitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommunicationUnits not implemented")
}
func (UnimplementedDeviceRegistryServiceServer) CreateDevice(context.Context, *CreateDeviceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDevice not implemented")
}
func (UnimplementedDeviceRegistryServiceServer) GetDevices(context.Context, *GetDevicesRequest) (*GetDevicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDevices not implemented")
}
func (UnimplementedDeviceRegistryServiceServer) SetDeviceCommunicationUnits(context.Context, *SetDeviceCommunicationUnitsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDeviceCommunicationUnits not implemented")
}
func (UnimplementedDeviceRegistryServiceServer) CreateDeviceGroup(context.Context, *CreateDeviceGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeviceGroup not implemented")
}
func (UnimplementedDeviceRegistryServiceServer) GetDeviceGroups(context.Context, *GetDeviceGroupsRequest) (*GetDeviceGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceGroups not implemented")
}
func (UnimplementedDeviceRegistryServiceServer) AddDevicesToGroup(context.Context, *AddDevicesToGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDevicesToGroup not implemented")
}
func (UnimplementedDeviceRegistryServiceServer) RemoveDevicesFromGroup(context.Context, *RemoveDevicesFromGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDevicesFromGroup not implemented")
}
func (UnimplementedDeviceRegistryServiceServer) mustEmbedUnimplementedDeviceRegistryServiceServer() {}

// UnsafeDeviceRegistryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceRegistryServiceServer will
// result in compilation errors.
type UnsafeDeviceRegistryServiceServer interface {
	mustEmbedUnimplementedDeviceRegistryServiceServer()
}

func RegisterDeviceRegistryServiceServer(s grpc.ServiceRegistrar, srv DeviceRegistryServiceServer) {
	s.RegisterService(&DeviceRegistryService_ServiceDesc, srv)
}

func _DeviceRegistryService_CreateCommunicationUnit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommunicationUnitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceRegistryServiceServer).CreateCommunicationUnit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceRegistryService_CreateCommunicationUnit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceRegistryServiceServer).CreateCommunicationUnit(ctx, req.(*CreateCommunicationUnitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceRegistryService_GetCommunicationUnits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommunicationUnitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceRegistryServiceServer).GetCommunicationUnits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceRegistryService_GetCommunicationUnits_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceRegistryServiceServer).GetCommunicationUnits(ctx, req.(*GetCommunicationUnitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceRegistryService_CreateDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceRegistryServiceServer).CreateDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceRegistryService_CreateDevice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceRegistryServiceServer).CreateDevice(ctx, req.(*CreateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceRegistryService_GetDevices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceRegistryServiceServer).GetDevices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceRegistryService_GetDevices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceRegistryServiceServer).GetDevices(ctx, req.(*GetDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceRegistryService_SetDeviceCommunicationUnits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDeviceCommunicationUnitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceRegistryServiceServer).SetDeviceCommunicationUnits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceRegistryService_SetDeviceCommunicationUnits_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceRegistryServiceServer).SetDeviceCommunicationUnits(ctx, req.(*SetDeviceCommunicationUnitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceRegistryService_CreateDeviceGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceRegistryServiceServer).CreateDeviceGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceRegistryService_CreateDeviceGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceRegistryServiceServer).CreateDeviceGroup(ctx, req.(*CreateDeviceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceRegistryService_GetDeviceGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceRegistryServiceServer).GetDeviceGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceRegistryService_GetDeviceGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceRegistryServiceServer).GetDeviceGroups(ctx, req.(*GetDeviceGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceRegistryService_AddDevicesToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDevicesToGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceRegistryServiceServer).AddDevicesToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceRegistryService_AddDevicesToGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceRegistryServiceServer).AddDevicesToGroup(ctx, req.(*AddDevicesToGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceRegistryService_RemoveDevicesFromGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDevicesFromGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceRegistryServiceServer).RemoveDevicesFromGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceRegistryService_RemoveDevicesFromGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceRegistryServiceServer).RemoveDevicesFromGroup(ctx, req.(*RemoveDevicesFromGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceRegistryService_ServiceDesc is the grpc.ServiceDesc for DeviceRegistryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceRegistryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "io.clbs.openhes.pbdeviceregistry.DeviceRegistryService",
	HandlerType: (*DeviceRegistryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCommunicationUnit",
			Handler:    _DeviceRegistryService_CreateCommunicationUnit_Handler,
		},
		{
			MethodName: "GetCommunicationUnits",
			Handler:    _DeviceRegistryService_GetCommunicationUnits_Handler,
		},
		{
			MethodName: "CreateDevice",
			Handler:    _DeviceRegistryService_CreateDevice_Handler,
		},
		{
			MethodName: "GetDevices",
			Handler:    _DeviceRegistryService_GetDevices_Handler,
		},
		{
			MethodName: "SetDeviceCommunicationUnits",
			Handler:    _DeviceRegistryService_SetDeviceCommunicationUnits_Handler,
		},
		{
			MethodName: "CreateDeviceGroup",
			Handler:    _DeviceRegistryService_CreateDeviceGroup_Handler,
		},
		{
			MethodName: "GetDeviceGroups",
			Handler:    _DeviceRegistryService_GetDeviceGroups_Handler,
		},
		{
			MethodName: "AddDevicesToGroup",
			Handler:    _DeviceRegistryService_AddDevicesToGroup_Handler,
		},
		{
			MethodName: "RemoveDevicesFromGroup",
			Handler:    _DeviceRegistryService_RemoveDevicesFromGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pbdeviceregistry.proto",
}
