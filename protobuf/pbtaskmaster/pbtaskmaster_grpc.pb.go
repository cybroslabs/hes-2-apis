// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: pbtaskmaster.proto

package pbtaskmaster

import (
	context "context"
	pbdriver "github.com/cybroslabs/hes-2-apis/protobuf/pbdriver"
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
	TaskmasterService_CreateBulk_FullMethodName         = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/CreateBulk"
	TaskmasterService_GetBulks_FullMethodName           = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/GetBulks"
	TaskmasterService_GetBulk_FullMethodName            = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/GetBulk"
	TaskmasterService_GetJob_FullMethodName             = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/GetJob"
	TaskmasterService_PurgeJobs_FullMethodName          = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/PurgeJobs"
	TaskmasterService_CancelJob_FullMethodName          = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/CancelJob"
	TaskmasterService_NegotiateStart_FullMethodName     = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/NegotiateStart"
	TaskmasterService_CacheSet_FullMethodName           = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/CacheSet"
	TaskmasterService_CacheGet_FullMethodName           = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/CacheGet"
	TaskmasterService_RegisterDriver_FullMethodName     = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/RegisterDriver"
	TaskmasterService_GetDrivers_FullMethodName         = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/GetDrivers"
	TaskmasterService_UnregisterDriver_FullMethodName   = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/UnregisterDriver"
	TaskmasterService_GetDriverTemplates_FullMethodName = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/GetDriverTemplates"
	TaskmasterService_GetConfig_FullMethodName          = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/GetConfig"
	TaskmasterService_SetConfig_FullMethodName          = "/io.clbs.openhes.pbtaskmaster.TaskmasterService/SetConfig"
)

// TaskmasterServiceClient is the client API for TaskmasterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskmasterServiceClient interface {
	// The method called by the RestApi to start a new bulk of jobs.
	CreateBulk(ctx context.Context, in *CreateBulksRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// The method called by the RestApi to retrieve stored bulks.
	GetBulks(ctx context.Context, in *GetBulksReuqest, opts ...grpc.CallOption) (*GetBulksResponse, error)
	// The method called by the RestApi to retrieve single bulk.
	GetBulk(ctx context.Context, in *GetBulkRequest, opts ...grpc.CallOption) (*GetBulkResponse, error)
	// The method called by the RestApi to get the job status. The parameter contains the job identifier.
	GetJob(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*GetJobResponse, error)
	// The method called by the RestApi to purge all jobs (queued/running/finished) from the system.
	PurgeJobs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CommonResponse, error)
	// The method called by the RestApi to cancel the job.
	CancelJob(ctx context.Context, in *CancelJobsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// The method called by the driver to inform Taskmaster about the instance existence. The parameter contains the driver version, the listening port, the meter type, the maximum number of concurrent jobs, the typical memory usage, the connection attributes template, and the job action templates.
	NegotiateStart(ctx context.Context, in *pbdriver.NegotiateRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	// The method called by the driver to store the cache entry. The parameter contains the cache key and the cache value. The key is unique within the driver type.
	CacheSet(ctx context.Context, in *CacheSetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// The method called by the driver to retrieve the cache entry. The parameter contains the cache key. The key is unique within the driver type.
	CacheGet(ctx context.Context, in *CacheGetRequest, opts ...grpc.CallOption) (*CacheGetResponse, error)
	// The method called by the RestApi to register the new driver to the Kubernetes. The parameter contains the driver type and the docker image.
	RegisterDriver(ctx context.Context, in *RegisterDriverRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	// The method called by the RestApi to get the list of drivers.
	GetDrivers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDriversResponse, error)
	// The method called by the RestApi to unregister driver from the system. The parameter contains the driver type.
	UnregisterDriver(ctx context.Context, in *UnregisterDriverRequest, opts ...grpc.CallOption) (*CommonResponse, error)
	// The method called by the RestApi to get the driver templates.
	GetDriverTemplates(ctx context.Context, in *GetDriverTemplatesRequest, opts ...grpc.CallOption) (*pbdriver.DriverTemplates, error)
	// The method called by the RestApi to get the system configuration.
	GetConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SystemConfigResponse, error)
	// The method called by the RestApi to set the system configuration.
	SetConfig(ctx context.Context, in *SystemConfig, opts ...grpc.CallOption) (*CommonResponse, error)
}

type taskmasterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskmasterServiceClient(cc grpc.ClientConnInterface) TaskmasterServiceClient {
	return &taskmasterServiceClient{cc}
}

func (c *taskmasterServiceClient) CreateBulk(ctx context.Context, in *CreateBulksRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TaskmasterService_CreateBulk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) GetBulks(ctx context.Context, in *GetBulksReuqest, opts ...grpc.CallOption) (*GetBulksResponse, error) {
	out := new(GetBulksResponse)
	err := c.cc.Invoke(ctx, TaskmasterService_GetBulks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) GetBulk(ctx context.Context, in *GetBulkRequest, opts ...grpc.CallOption) (*GetBulkResponse, error) {
	out := new(GetBulkResponse)
	err := c.cc.Invoke(ctx, TaskmasterService_GetBulk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) GetJob(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*GetJobResponse, error) {
	out := new(GetJobResponse)
	err := c.cc.Invoke(ctx, TaskmasterService_GetJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) PurgeJobs(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, TaskmasterService_PurgeJobs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) CancelJob(ctx context.Context, in *CancelJobsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TaskmasterService_CancelJob_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) NegotiateStart(ctx context.Context, in *pbdriver.NegotiateRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, TaskmasterService_NegotiateStart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) CacheSet(ctx context.Context, in *CacheSetRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, TaskmasterService_CacheSet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) CacheGet(ctx context.Context, in *CacheGetRequest, opts ...grpc.CallOption) (*CacheGetResponse, error) {
	out := new(CacheGetResponse)
	err := c.cc.Invoke(ctx, TaskmasterService_CacheGet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) RegisterDriver(ctx context.Context, in *RegisterDriverRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, TaskmasterService_RegisterDriver_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) GetDrivers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetDriversResponse, error) {
	out := new(GetDriversResponse)
	err := c.cc.Invoke(ctx, TaskmasterService_GetDrivers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) UnregisterDriver(ctx context.Context, in *UnregisterDriverRequest, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, TaskmasterService_UnregisterDriver_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) GetDriverTemplates(ctx context.Context, in *GetDriverTemplatesRequest, opts ...grpc.CallOption) (*pbdriver.DriverTemplates, error) {
	out := new(pbdriver.DriverTemplates)
	err := c.cc.Invoke(ctx, TaskmasterService_GetDriverTemplates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) GetConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*SystemConfigResponse, error) {
	out := new(SystemConfigResponse)
	err := c.cc.Invoke(ctx, TaskmasterService_GetConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskmasterServiceClient) SetConfig(ctx context.Context, in *SystemConfig, opts ...grpc.CallOption) (*CommonResponse, error) {
	out := new(CommonResponse)
	err := c.cc.Invoke(ctx, TaskmasterService_SetConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskmasterServiceServer is the server API for TaskmasterService service.
// All implementations must embed UnimplementedTaskmasterServiceServer
// for forward compatibility
type TaskmasterServiceServer interface {
	// The method called by the RestApi to start a new bulk of jobs.
	CreateBulk(context.Context, *CreateBulksRequest) (*emptypb.Empty, error)
	// The method called by the RestApi to retrieve stored bulks.
	GetBulks(context.Context, *GetBulksReuqest) (*GetBulksResponse, error)
	// The method called by the RestApi to retrieve single bulk.
	GetBulk(context.Context, *GetBulkRequest) (*GetBulkResponse, error)
	// The method called by the RestApi to get the job status. The parameter contains the job identifier.
	GetJob(context.Context, *GetJobRequest) (*GetJobResponse, error)
	// The method called by the RestApi to purge all jobs (queued/running/finished) from the system.
	PurgeJobs(context.Context, *emptypb.Empty) (*CommonResponse, error)
	// The method called by the RestApi to cancel the job.
	CancelJob(context.Context, *CancelJobsRequest) (*emptypb.Empty, error)
	// The method called by the driver to inform Taskmaster about the instance existence. The parameter contains the driver version, the listening port, the meter type, the maximum number of concurrent jobs, the typical memory usage, the connection attributes template, and the job action templates.
	NegotiateStart(context.Context, *pbdriver.NegotiateRequest) (*CommonResponse, error)
	// The method called by the driver to store the cache entry. The parameter contains the cache key and the cache value. The key is unique within the driver type.
	CacheSet(context.Context, *CacheSetRequest) (*emptypb.Empty, error)
	// The method called by the driver to retrieve the cache entry. The parameter contains the cache key. The key is unique within the driver type.
	CacheGet(context.Context, *CacheGetRequest) (*CacheGetResponse, error)
	// The method called by the RestApi to register the new driver to the Kubernetes. The parameter contains the driver type and the docker image.
	RegisterDriver(context.Context, *RegisterDriverRequest) (*CommonResponse, error)
	// The method called by the RestApi to get the list of drivers.
	GetDrivers(context.Context, *emptypb.Empty) (*GetDriversResponse, error)
	// The method called by the RestApi to unregister driver from the system. The parameter contains the driver type.
	UnregisterDriver(context.Context, *UnregisterDriverRequest) (*CommonResponse, error)
	// The method called by the RestApi to get the driver templates.
	GetDriverTemplates(context.Context, *GetDriverTemplatesRequest) (*pbdriver.DriverTemplates, error)
	// The method called by the RestApi to get the system configuration.
	GetConfig(context.Context, *emptypb.Empty) (*SystemConfigResponse, error)
	// The method called by the RestApi to set the system configuration.
	SetConfig(context.Context, *SystemConfig) (*CommonResponse, error)
	mustEmbedUnimplementedTaskmasterServiceServer()
}

// UnimplementedTaskmasterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTaskmasterServiceServer struct {
}

func (UnimplementedTaskmasterServiceServer) CreateBulk(context.Context, *CreateBulksRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBulk not implemented")
}
func (UnimplementedTaskmasterServiceServer) GetBulks(context.Context, *GetBulksReuqest) (*GetBulksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBulks not implemented")
}
func (UnimplementedTaskmasterServiceServer) GetBulk(context.Context, *GetBulkRequest) (*GetBulkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBulk not implemented")
}
func (UnimplementedTaskmasterServiceServer) GetJob(context.Context, *GetJobRequest) (*GetJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJob not implemented")
}
func (UnimplementedTaskmasterServiceServer) PurgeJobs(context.Context, *emptypb.Empty) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurgeJobs not implemented")
}
func (UnimplementedTaskmasterServiceServer) CancelJob(context.Context, *CancelJobsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelJob not implemented")
}
func (UnimplementedTaskmasterServiceServer) NegotiateStart(context.Context, *pbdriver.NegotiateRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NegotiateStart not implemented")
}
func (UnimplementedTaskmasterServiceServer) CacheSet(context.Context, *CacheSetRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CacheSet not implemented")
}
func (UnimplementedTaskmasterServiceServer) CacheGet(context.Context, *CacheGetRequest) (*CacheGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CacheGet not implemented")
}
func (UnimplementedTaskmasterServiceServer) RegisterDriver(context.Context, *RegisterDriverRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDriver not implemented")
}
func (UnimplementedTaskmasterServiceServer) GetDrivers(context.Context, *emptypb.Empty) (*GetDriversResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDrivers not implemented")
}
func (UnimplementedTaskmasterServiceServer) UnregisterDriver(context.Context, *UnregisterDriverRequest) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnregisterDriver not implemented")
}
func (UnimplementedTaskmasterServiceServer) GetDriverTemplates(context.Context, *GetDriverTemplatesRequest) (*pbdriver.DriverTemplates, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDriverTemplates not implemented")
}
func (UnimplementedTaskmasterServiceServer) GetConfig(context.Context, *emptypb.Empty) (*SystemConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedTaskmasterServiceServer) SetConfig(context.Context, *SystemConfig) (*CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConfig not implemented")
}
func (UnimplementedTaskmasterServiceServer) mustEmbedUnimplementedTaskmasterServiceServer() {}

// UnsafeTaskmasterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskmasterServiceServer will
// result in compilation errors.
type UnsafeTaskmasterServiceServer interface {
	mustEmbedUnimplementedTaskmasterServiceServer()
}

func RegisterTaskmasterServiceServer(s grpc.ServiceRegistrar, srv TaskmasterServiceServer) {
	s.RegisterService(&TaskmasterService_ServiceDesc, srv)
}

func _TaskmasterService_CreateBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBulksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskmasterServiceServer).CreateBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskmasterService_CreateBulk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskmasterServiceServer).CreateBulk(ctx, req.(*CreateBulksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskmasterService_GetBulks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBulksReuqest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskmasterServiceServer).GetBulks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskmasterService_GetBulks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskmasterServiceServer).GetBulks(ctx, req.(*GetBulksReuqest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskmasterService_GetBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBulkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskmasterServiceServer).GetBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskmasterService_GetBulk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskmasterServiceServer).GetBulk(ctx, req.(*GetBulkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskmasterService_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobRequest)
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
		return srv.(TaskmasterServiceServer).GetJob(ctx, req.(*GetJobRequest))
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

func _TaskmasterService_CancelJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskmasterServiceServer).CancelJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskmasterService_CancelJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskmasterServiceServer).CancelJob(ctx, req.(*CancelJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskmasterService_NegotiateStart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbdriver.NegotiateRequest)
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
		return srv.(TaskmasterServiceServer).NegotiateStart(ctx, req.(*pbdriver.NegotiateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskmasterService_CacheSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CacheSetRequest)
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
		return srv.(TaskmasterServiceServer).CacheSet(ctx, req.(*CacheSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskmasterService_CacheGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CacheGetRequest)
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
		return srv.(TaskmasterServiceServer).CacheGet(ctx, req.(*CacheGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskmasterService_RegisterDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDriverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskmasterServiceServer).RegisterDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskmasterService_RegisterDriver_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskmasterServiceServer).RegisterDriver(ctx, req.(*RegisterDriverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskmasterService_GetDrivers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskmasterServiceServer).GetDrivers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskmasterService_GetDrivers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskmasterServiceServer).GetDrivers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskmasterService_UnregisterDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnregisterDriverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskmasterServiceServer).UnregisterDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskmasterService_UnregisterDriver_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskmasterServiceServer).UnregisterDriver(ctx, req.(*UnregisterDriverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskmasterService_GetDriverTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDriverTemplatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskmasterServiceServer).GetDriverTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskmasterService_GetDriverTemplates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskmasterServiceServer).GetDriverTemplates(ctx, req.(*GetDriverTemplatesRequest))
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
	in := new(SystemConfig)
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
		return srv.(TaskmasterServiceServer).SetConfig(ctx, req.(*SystemConfig))
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
			MethodName: "CreateBulk",
			Handler:    _TaskmasterService_CreateBulk_Handler,
		},
		{
			MethodName: "GetBulks",
			Handler:    _TaskmasterService_GetBulks_Handler,
		},
		{
			MethodName: "GetBulk",
			Handler:    _TaskmasterService_GetBulk_Handler,
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
			MethodName: "CancelJob",
			Handler:    _TaskmasterService_CancelJob_Handler,
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
			MethodName: "RegisterDriver",
			Handler:    _TaskmasterService_RegisterDriver_Handler,
		},
		{
			MethodName: "GetDrivers",
			Handler:    _TaskmasterService_GetDrivers_Handler,
		},
		{
			MethodName: "UnregisterDriver",
			Handler:    _TaskmasterService_UnregisterDriver_Handler,
		},
		{
			MethodName: "GetDriverTemplates",
			Handler:    _TaskmasterService_GetDriverTemplates_Handler,
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
	Streams:  []grpc.StreamDesc{},
	Metadata: "pbtaskmaster.proto",
}