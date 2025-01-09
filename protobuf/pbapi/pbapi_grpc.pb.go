// Editions version of proto3 file

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: pbapi.proto

package pbapi

import (
	context "context"
	pbdataproxymodels "github.com/cybroslabs/hes-2-apis/protobuf/pbdataproxymodels"
	pbdrivermodels "github.com/cybroslabs/hes-2-apis/protobuf/pbdrivermodels"
	pbdriveroperatormodels "github.com/cybroslabs/hes-2-apis/protobuf/pbdriveroperatormodels"
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
	ApiService_CreateBulk_FullMethodName         = "/io.clbs.openhes.pbapi.ApiService/CreateBulk"
	ApiService_GetBulks_FullMethodName           = "/io.clbs.openhes.pbapi.ApiService/GetBulks"
	ApiService_GetBulk_FullMethodName            = "/io.clbs.openhes.pbapi.ApiService/GetBulk"
	ApiService_GetJobStatus_FullMethodName       = "/io.clbs.openhes.pbapi.ApiService/GetJobStatus"
	ApiService_GetDrivers_FullMethodName         = "/io.clbs.openhes.pbapi.ApiService/GetDrivers"
	ApiService_GetDriverTemplates_FullMethodName = "/io.clbs.openhes.pbapi.ApiService/GetDriverTemplates"
)

// ApiServiceClient is the client API for ApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The Dataproxy related service definition.
type ApiServiceClient interface {
	// Starts a new bulk of jobs.
	CreateBulk(ctx context.Context, in *PublicCreateBulkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Retrieves the list of bulks.
	GetBulks(ctx context.Context, in *pbdataproxymodels.GetBulksReuqest, opts ...grpc.CallOption) (*pbdataproxymodels.GetBulksResponse, error)
	// Retrieves the bulk info and status.
	GetBulk(ctx context.Context, in *pbdataproxymodels.GetBulkRequest, opts ...grpc.CallOption) (*pbdataproxymodels.GetBulkResponse, error)
	// Retrieves the job status.
	GetJobStatus(ctx context.Context, in *pbdataproxymodels.GetJobStatusRequest, opts ...grpc.CallOption) (*pbdataproxymodels.GetJobStatusResponse, error)
	// Retrieves the list of drivers.
	GetDrivers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*pbdriveroperatormodels.GetDriversResponse, error)
	// Retrieves the driver templates.
	GetDriverTemplates(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*pbdrivermodels.DriverTemplates, error)
}

type apiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServiceClient(cc grpc.ClientConnInterface) ApiServiceClient {
	return &apiServiceClient{cc}
}

func (c *apiServiceClient) CreateBulk(ctx context.Context, in *PublicCreateBulkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ApiService_CreateBulk_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetBulks(ctx context.Context, in *pbdataproxymodels.GetBulksReuqest, opts ...grpc.CallOption) (*pbdataproxymodels.GetBulksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pbdataproxymodels.GetBulksResponse)
	err := c.cc.Invoke(ctx, ApiService_GetBulks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetBulk(ctx context.Context, in *pbdataproxymodels.GetBulkRequest, opts ...grpc.CallOption) (*pbdataproxymodels.GetBulkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pbdataproxymodels.GetBulkResponse)
	err := c.cc.Invoke(ctx, ApiService_GetBulk_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetJobStatus(ctx context.Context, in *pbdataproxymodels.GetJobStatusRequest, opts ...grpc.CallOption) (*pbdataproxymodels.GetJobStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pbdataproxymodels.GetJobStatusResponse)
	err := c.cc.Invoke(ctx, ApiService_GetJobStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetDrivers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*pbdriveroperatormodels.GetDriversResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pbdriveroperatormodels.GetDriversResponse)
	err := c.cc.Invoke(ctx, ApiService_GetDrivers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServiceClient) GetDriverTemplates(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*pbdrivermodels.DriverTemplates, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pbdrivermodels.DriverTemplates)
	err := c.cc.Invoke(ctx, ApiService_GetDriverTemplates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServiceServer is the server API for ApiService service.
// All implementations must embed UnimplementedApiServiceServer
// for forward compatibility.
//
// The Dataproxy related service definition.
type ApiServiceServer interface {
	// Starts a new bulk of jobs.
	CreateBulk(context.Context, *PublicCreateBulkRequest) (*emptypb.Empty, error)
	// Retrieves the list of bulks.
	GetBulks(context.Context, *pbdataproxymodels.GetBulksReuqest) (*pbdataproxymodels.GetBulksResponse, error)
	// Retrieves the bulk info and status.
	GetBulk(context.Context, *pbdataproxymodels.GetBulkRequest) (*pbdataproxymodels.GetBulkResponse, error)
	// Retrieves the job status.
	GetJobStatus(context.Context, *pbdataproxymodels.GetJobStatusRequest) (*pbdataproxymodels.GetJobStatusResponse, error)
	// Retrieves the list of drivers.
	GetDrivers(context.Context, *emptypb.Empty) (*pbdriveroperatormodels.GetDriversResponse, error)
	// Retrieves the driver templates.
	GetDriverTemplates(context.Context, *wrapperspb.StringValue) (*pbdrivermodels.DriverTemplates, error)
	mustEmbedUnimplementedApiServiceServer()
}

// UnimplementedApiServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedApiServiceServer struct{}

func (UnimplementedApiServiceServer) CreateBulk(context.Context, *PublicCreateBulkRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBulk not implemented")
}
func (UnimplementedApiServiceServer) GetBulks(context.Context, *pbdataproxymodels.GetBulksReuqest) (*pbdataproxymodels.GetBulksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBulks not implemented")
}
func (UnimplementedApiServiceServer) GetBulk(context.Context, *pbdataproxymodels.GetBulkRequest) (*pbdataproxymodels.GetBulkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBulk not implemented")
}
func (UnimplementedApiServiceServer) GetJobStatus(context.Context, *pbdataproxymodels.GetJobStatusRequest) (*pbdataproxymodels.GetJobStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobStatus not implemented")
}
func (UnimplementedApiServiceServer) GetDrivers(context.Context, *emptypb.Empty) (*pbdriveroperatormodels.GetDriversResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDrivers not implemented")
}
func (UnimplementedApiServiceServer) GetDriverTemplates(context.Context, *wrapperspb.StringValue) (*pbdrivermodels.DriverTemplates, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDriverTemplates not implemented")
}
func (UnimplementedApiServiceServer) mustEmbedUnimplementedApiServiceServer() {}
func (UnimplementedApiServiceServer) testEmbeddedByValue()                    {}

// UnsafeApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServiceServer will
// result in compilation errors.
type UnsafeApiServiceServer interface {
	mustEmbedUnimplementedApiServiceServer()
}

func RegisterApiServiceServer(s grpc.ServiceRegistrar, srv ApiServiceServer) {
	// If the following call pancis, it indicates UnimplementedApiServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ApiService_ServiceDesc, srv)
}

func _ApiService_CreateBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicCreateBulkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).CreateBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiService_CreateBulk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).CreateBulk(ctx, req.(*PublicCreateBulkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetBulks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbdataproxymodels.GetBulksReuqest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetBulks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiService_GetBulks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetBulks(ctx, req.(*pbdataproxymodels.GetBulksReuqest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbdataproxymodels.GetBulkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiService_GetBulk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetBulk(ctx, req.(*pbdataproxymodels.GetBulkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetJobStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbdataproxymodels.GetJobStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetJobStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiService_GetJobStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetJobStatus(ctx, req.(*pbdataproxymodels.GetJobStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetDrivers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetDrivers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiService_GetDrivers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetDrivers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiService_GetDriverTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServiceServer).GetDriverTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiService_GetDriverTemplates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServiceServer).GetDriverTemplates(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiService_ServiceDesc is the grpc.ServiceDesc for ApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "io.clbs.openhes.pbapi.ApiService",
	HandlerType: (*ApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBulk",
			Handler:    _ApiService_CreateBulk_Handler,
		},
		{
			MethodName: "GetBulks",
			Handler:    _ApiService_GetBulks_Handler,
		},
		{
			MethodName: "GetBulk",
			Handler:    _ApiService_GetBulk_Handler,
		},
		{
			MethodName: "GetJobStatus",
			Handler:    _ApiService_GetJobStatus_Handler,
		},
		{
			MethodName: "GetDrivers",
			Handler:    _ApiService_GetDrivers_Handler,
		},
		{
			MethodName: "GetDriverTemplates",
			Handler:    _ApiService_GetDriverTemplates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pbapi.proto",
}
