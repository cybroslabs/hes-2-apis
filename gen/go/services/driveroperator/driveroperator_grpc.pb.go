// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: services/driveroperator/driveroperator.proto

package driveroperator

import (
	context "context"
	acquisition "github.com/cybroslabs/hes-2-apis/gen/go/acquisition"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DriverOperatorService_ListDrivers_FullMethodName    = "/io.clbs.openhes.services.driveroperator.DriverOperatorService/ListDrivers"
	DriverOperatorService_SetDriver_FullMethodName      = "/io.clbs.openhes.services.driveroperator.DriverOperatorService/SetDriver"
	DriverOperatorService_GetDriver_FullMethodName      = "/io.clbs.openhes.services.driveroperator.DriverOperatorService/GetDriver"
	DriverOperatorService_SetDriverScale_FullMethodName = "/io.clbs.openhes.services.driveroperator.DriverOperatorService/SetDriverScale"
	DriverOperatorService_GetDriverScale_FullMethodName = "/io.clbs.openhes.services.driveroperator.DriverOperatorService/GetDriverScale"
	DriverOperatorService_StartUpgrade_FullMethodName   = "/io.clbs.openhes.services.driveroperator.DriverOperatorService/StartUpgrade"
)

// DriverOperatorServiceClient is the client API for DriverOperatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The Driver Operator service definition.
// Those are the gRPC services that the Driver Operator provides for other components.
type DriverOperatorServiceClient interface {
	// The method called by the RestApi to get the list of drivers.
	ListDrivers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*acquisition.ListOfDriver, error)
	// The method called by the Driver to set the driver templates. The parameter contains the driver templates.
	SetDriver(ctx context.Context, in *acquisition.Driver, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// The method called by the RestApi to get the driver templates.
	GetDriver(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*acquisition.Driver, error)
	// The method called by the Taskmaster to set the driver scale.
	SetDriverScale(ctx context.Context, in *acquisition.SetDriverScaleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// The method called by the Taskmaster to get the driver scale.
	GetDriverScale(ctx context.Context, in *acquisition.GetDriverScaleRequest, opts ...grpc.CallOption) (*wrapperspb.UInt32Value, error)
	// The method called by the DeviceRegistry to start the driver in upgrade mode. It will provide structure upgrade between the driver versions.
	// The driver is started as Kubernetes job and ends when all the structures are upgraded; which is controlled by the DeviceRegistry.
	StartUpgrade(ctx context.Context, in *acquisition.StartUpgradeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type driverOperatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverOperatorServiceClient(cc grpc.ClientConnInterface) DriverOperatorServiceClient {
	return &driverOperatorServiceClient{cc}
}

func (c *driverOperatorServiceClient) ListDrivers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*acquisition.ListOfDriver, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(acquisition.ListOfDriver)
	err := c.cc.Invoke(ctx, DriverOperatorService_ListDrivers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverOperatorServiceClient) SetDriver(ctx context.Context, in *acquisition.Driver, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DriverOperatorService_SetDriver_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverOperatorServiceClient) GetDriver(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*acquisition.Driver, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(acquisition.Driver)
	err := c.cc.Invoke(ctx, DriverOperatorService_GetDriver_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverOperatorServiceClient) SetDriverScale(ctx context.Context, in *acquisition.SetDriverScaleRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DriverOperatorService_SetDriverScale_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverOperatorServiceClient) GetDriverScale(ctx context.Context, in *acquisition.GetDriverScaleRequest, opts ...grpc.CallOption) (*wrapperspb.UInt32Value, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(wrapperspb.UInt32Value)
	err := c.cc.Invoke(ctx, DriverOperatorService_GetDriverScale_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *driverOperatorServiceClient) StartUpgrade(ctx context.Context, in *acquisition.StartUpgradeRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DriverOperatorService_StartUpgrade_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverOperatorServiceServer is the server API for DriverOperatorService service.
// All implementations must embed UnimplementedDriverOperatorServiceServer
// for forward compatibility.
//
// The Driver Operator service definition.
// Those are the gRPC services that the Driver Operator provides for other components.
type DriverOperatorServiceServer interface {
	// The method called by the RestApi to get the list of drivers.
	ListDrivers(context.Context, *emptypb.Empty) (*acquisition.ListOfDriver, error)
	// The method called by the Driver to set the driver templates. The parameter contains the driver templates.
	SetDriver(context.Context, *acquisition.Driver) (*emptypb.Empty, error)
	// The method called by the RestApi to get the driver templates.
	GetDriver(context.Context, *wrapperspb.StringValue) (*acquisition.Driver, error)
	// The method called by the Taskmaster to set the driver scale.
	SetDriverScale(context.Context, *acquisition.SetDriverScaleRequest) (*emptypb.Empty, error)
	// The method called by the Taskmaster to get the driver scale.
	GetDriverScale(context.Context, *acquisition.GetDriverScaleRequest) (*wrapperspb.UInt32Value, error)
	// The method called by the DeviceRegistry to start the driver in upgrade mode. It will provide structure upgrade between the driver versions.
	// The driver is started as Kubernetes job and ends when all the structures are upgraded; which is controlled by the DeviceRegistry.
	StartUpgrade(context.Context, *acquisition.StartUpgradeRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedDriverOperatorServiceServer()
}

// UnimplementedDriverOperatorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDriverOperatorServiceServer struct{}

func (UnimplementedDriverOperatorServiceServer) ListDrivers(context.Context, *emptypb.Empty) (*acquisition.ListOfDriver, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDrivers not implemented")
}
func (UnimplementedDriverOperatorServiceServer) SetDriver(context.Context, *acquisition.Driver) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDriver not implemented")
}
func (UnimplementedDriverOperatorServiceServer) GetDriver(context.Context, *wrapperspb.StringValue) (*acquisition.Driver, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDriver not implemented")
}
func (UnimplementedDriverOperatorServiceServer) SetDriverScale(context.Context, *acquisition.SetDriverScaleRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDriverScale not implemented")
}
func (UnimplementedDriverOperatorServiceServer) GetDriverScale(context.Context, *acquisition.GetDriverScaleRequest) (*wrapperspb.UInt32Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDriverScale not implemented")
}
func (UnimplementedDriverOperatorServiceServer) StartUpgrade(context.Context, *acquisition.StartUpgradeRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartUpgrade not implemented")
}
func (UnimplementedDriverOperatorServiceServer) mustEmbedUnimplementedDriverOperatorServiceServer() {}
func (UnimplementedDriverOperatorServiceServer) testEmbeddedByValue()                               {}

// UnsafeDriverOperatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DriverOperatorServiceServer will
// result in compilation errors.
type UnsafeDriverOperatorServiceServer interface {
	mustEmbedUnimplementedDriverOperatorServiceServer()
}

func RegisterDriverOperatorServiceServer(s grpc.ServiceRegistrar, srv DriverOperatorServiceServer) {
	// If the following call pancis, it indicates UnimplementedDriverOperatorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DriverOperatorService_ServiceDesc, srv)
}

func _DriverOperatorService_ListDrivers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverOperatorServiceServer).ListDrivers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverOperatorService_ListDrivers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverOperatorServiceServer).ListDrivers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverOperatorService_SetDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(acquisition.Driver)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverOperatorServiceServer).SetDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverOperatorService_SetDriver_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverOperatorServiceServer).SetDriver(ctx, req.(*acquisition.Driver))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverOperatorService_GetDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverOperatorServiceServer).GetDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverOperatorService_GetDriver_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverOperatorServiceServer).GetDriver(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverOperatorService_SetDriverScale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(acquisition.SetDriverScaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverOperatorServiceServer).SetDriverScale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverOperatorService_SetDriverScale_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverOperatorServiceServer).SetDriverScale(ctx, req.(*acquisition.SetDriverScaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverOperatorService_GetDriverScale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(acquisition.GetDriverScaleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverOperatorServiceServer).GetDriverScale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverOperatorService_GetDriverScale_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverOperatorServiceServer).GetDriverScale(ctx, req.(*acquisition.GetDriverScaleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DriverOperatorService_StartUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(acquisition.StartUpgradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverOperatorServiceServer).StartUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverOperatorService_StartUpgrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverOperatorServiceServer).StartUpgrade(ctx, req.(*acquisition.StartUpgradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DriverOperatorService_ServiceDesc is the grpc.ServiceDesc for DriverOperatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DriverOperatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "io.clbs.openhes.services.driveroperator.DriverOperatorService",
	HandlerType: (*DriverOperatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDrivers",
			Handler:    _DriverOperatorService_ListDrivers_Handler,
		},
		{
			MethodName: "SetDriver",
			Handler:    _DriverOperatorService_SetDriver_Handler,
		},
		{
			MethodName: "GetDriver",
			Handler:    _DriverOperatorService_GetDriver_Handler,
		},
		{
			MethodName: "SetDriverScale",
			Handler:    _DriverOperatorService_SetDriverScale_Handler,
		},
		{
			MethodName: "GetDriverScale",
			Handler:    _DriverOperatorService_GetDriverScale_Handler,
		},
		{
			MethodName: "StartUpgrade",
			Handler:    _DriverOperatorService_StartUpgrade_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/driveroperator/driveroperator.proto",
}
