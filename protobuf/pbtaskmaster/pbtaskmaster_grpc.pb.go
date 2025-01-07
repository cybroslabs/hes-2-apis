// Editions version of proto3 file

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: pbtaskmaster.proto

package pbtaskmaster

import (
	context "context"
	pbdrivermodels "github.com/cybroslabs/hes-2-apis/protobuf/pbdrivermodels"
	pbtaskmastermodels "github.com/cybroslabs/hes-2-apis/protobuf/pbtaskmastermodels"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TaskmasterService_QueueJobs_FullMethodName      = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/QueueJobs"
	TaskmasterService_GetJob_FullMethodName         = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/GetJob"
	TaskmasterService_PurgeJobs_FullMethodName      = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/PurgeJobs"
	TaskmasterService_CancelJobs_FullMethodName     = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/CancelJobs"
	TaskmasterService_Subscribe_FullMethodName      = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/Subscribe"
	TaskmasterService_NegotiateStart_FullMethodName = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/NegotiateStart"
	TaskmasterService_CacheSet_FullMethodName       = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/CacheSet"
	TaskmasterService_CacheGet_FullMethodName       = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/CacheGet"
	TaskmasterService_GetConfig_FullMethodName      = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/GetConfig"
	TaskmasterService_SetConfig_FullMethodName      = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/SetConfig"
)

// TaskmasterServiceClient is the client API for TaskmasterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The Taskmaster service definition.
// Those are the gRPC services that the Taskmaster provides for other components.
type TaskmasterServiceClient interface {
	QueueJobs(ctx context.Context, in *pbtaskmastermodels.QueueJobsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// The method called by the RestApi to get the job status. The parameter contains the job identifier.
	GetJob(ctx context.Context, in *pbtaskmastermodels.GetJobRequest, opts ...grpc.CallOption) (*pbtaskmastermodels.GetJobResponse, error)
	// The method called by the RestApi to purge all jobs (queued/running/finished) from the system.
	PurgeJobs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// The method called by the RestApi to cancel the job.
	CancelJobs(ctx context.Context, in *pbtaskmastermodels.CancelJobsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Subcribe for notification stream
	Subscribe(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[pbtaskmastermodels.StreamEventsData], error)
	// The method called by the driver to inform Taskmaster about the instance existence. The parameter contains the driver version, the listening port, the meter type, the maximum number of concurrent jobs, the typical memory usage, the connection attributes template, and the job action templates.
	NegotiateStart(ctx context.Context, in *pbdrivermodels.NegotiateRequest, opts ...grpc.CallOption) (*pbdrivermodels.CommonResponse, error)
	// The method called by the driver to store the cache entry. The parameter contains the cache key and the cache value. The key is unique within the driver type.
	CacheSet(ctx context.Context, in *pbtaskmastermodels.CacheSetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// The method called by the driver to retrieve the cache entry. The parameter contains the cache key. The key is unique within the driver type.
	CacheGet(ctx context.Context, in *pbtaskmastermodels.CacheGetRequest, opts ...grpc.CallOption) (*pbtaskmastermodels.CacheGetResponse, error)
	// The method called by the RestApi to get the system configuration.
	GetConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*pbtaskmastermodels.SystemConfigResponse, error)
	// The method called by the RestApi to set the system configuration.
	SetConfig(ctx context.Context, in *pbtaskmastermodels.SystemConfig, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type taskmasterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskmasterServiceClient(cc grpc.ClientConnInterface) TaskmasterServiceClient {
	return &taskmasterServiceClient{cc}
}

func (c *taskmasterServiceClient) QueueJobs(ctx context.Context, in *pbtaskmastermodels.QueueJobsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TaskmasterService_QueueJobs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) GetJob(ctx context.Context, in *pbtaskmastermodels.GetJobRequest, opts ...grpc.CallOption) (*pbtaskmastermodels.GetJobResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pbtaskmastermodels.GetJobResponse)
	err := c.cc.Invoke(ctx, TaskmasterService_GetJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) PurgeJobs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TaskmasterService_PurgeJobs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) CancelJobs(ctx context.Context, in *pbtaskmastermodels.CancelJobsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TaskmasterService_CancelJobs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) Subscribe(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[pbtaskmastermodels.StreamEventsData], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &TaskmasterService_ServiceDesc.Streams[0], TaskmasterService_Subscribe_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[emptypb.Empty, pbtaskmastermodels.StreamEventsData]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TaskmasterService_SubscribeClient = grpc.ServerStreamingClient[pbtaskmastermodels.StreamEventsData]

func (c *taskmasterServiceClient) NegotiateStart(ctx context.Context, in *pbdrivermodels.NegotiateRequest, opts ...grpc.CallOption) (*pbdrivermodels.CommonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pbdrivermodels.CommonResponse)
	err := c.cc.Invoke(ctx, TaskmasterService_NegotiateStart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) CacheSet(ctx context.Context, in *pbtaskmastermodels.CacheSetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TaskmasterService_CacheSet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) CacheGet(ctx context.Context, in *pbtaskmastermodels.CacheGetRequest, opts ...grpc.CallOption) (*pbtaskmastermodels.CacheGetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pbtaskmastermodels.CacheGetResponse)
	err := c.cc.Invoke(ctx, TaskmasterService_CacheGet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) GetConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*pbtaskmastermodels.SystemConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pbtaskmastermodels.SystemConfigResponse)
	err := c.cc.Invoke(ctx, TaskmasterService_GetConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) SetConfig(ctx context.Context, in *pbtaskmastermodels.SystemConfig, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TaskmasterService_SetConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskmasterServiceServer is the server API for TaskmasterService service.
// All implementations must embed UnimplementedTaskmasterServiceServer
// for forward compatibility.
//
// The Taskmaster service definition.
// Those are the gRPC services that the Taskmaster provides for other components.
type TaskmasterServiceServer interface {
	QueueJobs(context.Context, *pbtaskmastermodels.QueueJobsRequest) (*emptypb.Empty, error)
	// The method called by the RestApi to get the job status. The parameter contains the job identifier.
	GetJob(context.Context, *pbtaskmastermodels.GetJobRequest) (*pbtaskmastermodels.GetJobResponse, error)
	// The method called by the RestApi to purge all jobs (queued/running/finished) from the system.
	PurgeJobs(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// The method called by the RestApi to cancel the job.
	CancelJobs(context.Context, *pbtaskmastermodels.CancelJobsRequest) (*emptypb.Empty, error)
	// Subcribe for notification stream
	Subscribe(*emptypb.Empty, grpc.ServerStreamingServer[pbtaskmastermodels.StreamEventsData]) error
	// The method called by the driver to inform Taskmaster about the instance existence. The parameter contains the driver version, the listening port, the meter type, the maximum number of concurrent jobs, the typical memory usage, the connection attributes template, and the job action templates.
	NegotiateStart(context.Context, *pbdrivermodels.NegotiateRequest) (*pbdrivermodels.CommonResponse, error)
	// The method called by the driver to store the cache entry. The parameter contains the cache key and the cache value. The key is unique within the driver type.
	CacheSet(context.Context, *pbtaskmastermodels.CacheSetRequest) (*emptypb.Empty, error)
	// The method called by the driver to retrieve the cache entry. The parameter contains the cache key. The key is unique within the driver type.
	CacheGet(context.Context, *pbtaskmastermodels.CacheGetRequest) (*pbtaskmastermodels.CacheGetResponse, error)
	// The method called by the RestApi to get the system configuration.
	GetConfig(context.Context, *emptypb.Empty) (*pbtaskmastermodels.SystemConfigResponse, error)
	// The method called by the RestApi to set the system configuration.
	SetConfig(context.Context, *pbtaskmastermodels.SystemConfig) (*emptypb.Empty, error)
	mustEmbedUnimplementedTaskmasterServiceServer()
}

// UnimplementedTaskmasterServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTaskmasterServiceServer struct{}

func (UnimplementedTaskmasterServiceServer) QueueJobs(context.Context, *pbtaskmastermodels.QueueJobsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueueJobs not implemented")
}
func (UnimplementedTaskmasterServiceServer) GetJob(context.Context, *pbtaskmastermodels.GetJobRequest) (*pbtaskmastermodels.GetJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJob not implemented")
}
func (UnimplementedTaskmasterServiceServer) PurgeJobs(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurgeJobs not implemented")
}
func (UnimplementedTaskmasterServiceServer) CancelJobs(context.Context, *pbtaskmastermodels.CancelJobsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelJobs not implemented")
}
func (UnimplementedTaskmasterServiceServer) Subscribe(*emptypb.Empty, grpc.ServerStreamingServer[pbtaskmastermodels.StreamEventsData]) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedTaskmasterServiceServer) NegotiateStart(context.Context, *pbdrivermodels.NegotiateRequest) (*pbdrivermodels.CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NegotiateStart not implemented")
}
func (UnimplementedTaskmasterServiceServer) CacheSet(context.Context, *pbtaskmastermodels.CacheSetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CacheSet not implemented")
}
func (UnimplementedTaskmasterServiceServer) CacheGet(context.Context, *pbtaskmastermodels.CacheGetRequest) (*pbtaskmastermodels.CacheGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CacheGet not implemented")
}
func (UnimplementedTaskmasterServiceServer) GetConfig(context.Context, *emptypb.Empty) (*pbtaskmastermodels.SystemConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedTaskmasterServiceServer) SetConfig(context.Context, *pbtaskmastermodels.SystemConfig) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConfig not implemented")
}
func (UnimplementedTaskmasterServiceServer) mustEmbedUnimplementedTaskmasterServiceServer() {}
func (UnimplementedTaskmasterServiceServer) testEmbeddedByValue()                           {}

// UnsafeTaskmasterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskmasterServiceServer will
// result in compilation errors.
type UnsafeTaskmasterServiceServer interface {
	mustEmbedUnimplementedTaskmasterServiceServer()
}

func RegisterTaskmasterServiceServer(s grpc.ServiceRegistrar, srv TaskmasterServiceServer) {
	// If the following call pancis, it indicates UnimplementedTaskmasterServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TaskmasterService_ServiceDesc, srv)
}

func _TaskmasterService_QueueJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbtaskmastermodels.QueueJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskmasterServiceServer).QueueJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskmasterService_QueueJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskmasterServiceServer).QueueJobs(ctx, req.(*pbtaskmastermodels.QueueJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskmasterService_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbtaskmastermodels.GetJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskmasterServiceServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskmasterService_GetJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskmasterServiceServer).GetJob(ctx, req.(*pbtaskmastermodels.GetJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskmasterService_PurgeJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskmasterServiceServer).PurgeJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskmasterService_PurgeJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskmasterServiceServer).PurgeJobs(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskmasterService_CancelJobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbtaskmastermodels.CancelJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskmasterServiceServer).CancelJobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskmasterService_CancelJobs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskmasterServiceServer).CancelJobs(ctx, req.(*pbtaskmastermodels.CancelJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskmasterService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskmasterServiceServer).Subscribe(m, &grpc.GenericServerStream[emptypb.Empty, pbtaskmastermodels.StreamEventsData]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type TaskmasterService_SubscribeServer = grpc.ServerStreamingServer[pbtaskmastermodels.StreamEventsData]

func _TaskmasterService_NegotiateStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbdrivermodels.NegotiateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskmasterServiceServer).NegotiateStart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskmasterService_NegotiateStart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskmasterServiceServer).NegotiateStart(ctx, req.(*pbdrivermodels.NegotiateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskmasterService_CacheSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbtaskmastermodels.CacheSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskmasterServiceServer).CacheSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskmasterService_CacheSet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskmasterServiceServer).CacheSet(ctx, req.(*pbtaskmastermodels.CacheSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskmasterService_CacheGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbtaskmastermodels.CacheGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskmasterServiceServer).CacheGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskmasterService_CacheGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskmasterServiceServer).CacheGet(ctx, req.(*pbtaskmastermodels.CacheGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskmasterService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskmasterServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskmasterService_GetConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskmasterServiceServer).GetConfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskmasterService_SetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbtaskmastermodels.SystemConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskmasterServiceServer).SetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskmasterService_SetConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskmasterServiceServer).SetConfig(ctx, req.(*pbtaskmastermodels.SystemConfig))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskmasterService_ServiceDesc is the grpc.ServiceDesc for TaskmasterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskmasterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "io.clbs.openhes.pbtaskmaster.TaskmasterService",
	HandlerType: (*TaskmasterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueueJobs",
			Handler:    _TaskmasterService_QueueJobs_Handler,
		},
		{
			MethodName: "GetJob",
			Handler:    _TaskmasterService_GetJob_Handler,
		},
		{
			MethodName: "PurgeJobs",
			Handler:    _TaskmasterService_PurgeJobs_Handler,
		},
		{
			MethodName: "CancelJobs",
			Handler:    _TaskmasterService_CancelJobs_Handler,
		},
		{
			MethodName: "NegotiateStart",
			Handler:    _TaskmasterService_NegotiateStart_Handler,
		},
		{
			MethodName: "CacheSet",
			Handler:    _TaskmasterService_CacheSet_Handler,
		},
		{
			MethodName: "CacheGet",
			Handler:    _TaskmasterService_CacheGet_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _TaskmasterService_GetConfig_Handler,
		},
		{
			MethodName: "SetConfig",
			Handler:    _TaskmasterService_SetConfig_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _TaskmasterService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pbtaskmaster.proto",
}
