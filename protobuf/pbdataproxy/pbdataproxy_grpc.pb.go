// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: pbdataproxy.proto

package pbdataproxy

import (
	context "context"
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
	DataproxyService_CreateBulk_FullMethodName   = "/io.clbs.openhes.pbdataproxy.DataproxyService/CreateBulk"
	DataproxyService_GetBulks_FullMethodName     = "/io.clbs.openhes.pbdataproxy.DataproxyService/GetBulks"
	DataproxyService_GetBulk_FullMethodName      = "/io.clbs.openhes.pbdataproxy.DataproxyService/GetBulk"
	DataproxyService_GetJobStatus_FullMethodName = "/io.clbs.openhes.pbdataproxy.DataproxyService/GetJobStatus"
)

// DataproxyServiceClient is the client API for DataproxyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The Dataproxy related service definition.
type DataproxyServiceClient interface {
	// The method called by the RestApi to start a new bulk of jobs.
	CreateBulk(ctx context.Context, in *CreateBulkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// The method called by the RestApi to retrieve stored bulks.
	GetBulks(ctx context.Context, in *GetBulksReuqest, opts ...grpc.CallOption) (*GetBulksResponse, error)
	// The method called by the RestApi to retrieve single bulk.
	GetBulk(ctx context.Context, in *GetBulkRequest, opts ...grpc.CallOption) (*GetBulkResponse, error)
	// The method called by the RestApi to get the job status. The parameter contains the job identifier.
	GetJobStatus(ctx context.Context, in *GetJobStatusRequest, opts ...grpc.CallOption) (*GetJobStatusResponse, error)
}

type dataproxyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataproxyServiceClient(cc grpc.ClientConnInterface) DataproxyServiceClient {
	return &dataproxyServiceClient{cc}
}

func (c *dataproxyServiceClient) CreateBulk(ctx context.Context, in *CreateBulkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DataproxyService_CreateBulk_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataproxyServiceClient) GetBulks(ctx context.Context, in *GetBulksReuqest, opts ...grpc.CallOption) (*GetBulksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBulksResponse)
	err := c.cc.Invoke(ctx, DataproxyService_GetBulks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataproxyServiceClient) GetBulk(ctx context.Context, in *GetBulkRequest, opts ...grpc.CallOption) (*GetBulkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBulkResponse)
	err := c.cc.Invoke(ctx, DataproxyService_GetBulk_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataproxyServiceClient) GetJobStatus(ctx context.Context, in *GetJobStatusRequest, opts ...grpc.CallOption) (*GetJobStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetJobStatusResponse)
	err := c.cc.Invoke(ctx, DataproxyService_GetJobStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataproxyServiceServer is the server API for DataproxyService service.
// All implementations must embed UnimplementedDataproxyServiceServer
// for forward compatibility.
//
// The Dataproxy related service definition.
type DataproxyServiceServer interface {
	// The method called by the RestApi to start a new bulk of jobs.
	CreateBulk(context.Context, *CreateBulkRequest) (*emptypb.Empty, error)
	// The method called by the RestApi to retrieve stored bulks.
	GetBulks(context.Context, *GetBulksReuqest) (*GetBulksResponse, error)
	// The method called by the RestApi to retrieve single bulk.
	GetBulk(context.Context, *GetBulkRequest) (*GetBulkResponse, error)
	// The method called by the RestApi to get the job status. The parameter contains the job identifier.
	GetJobStatus(context.Context, *GetJobStatusRequest) (*GetJobStatusResponse, error)
	mustEmbedUnimplementedDataproxyServiceServer()
}

// UnimplementedDataproxyServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDataproxyServiceServer struct{}

func (UnimplementedDataproxyServiceServer) CreateBulk(context.Context, *CreateBulkRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBulk not implemented")
}
func (UnimplementedDataproxyServiceServer) GetBulks(context.Context, *GetBulksReuqest) (*GetBulksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBulks not implemented")
}
func (UnimplementedDataproxyServiceServer) GetBulk(context.Context, *GetBulkRequest) (*GetBulkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBulk not implemented")
}
func (UnimplementedDataproxyServiceServer) GetJobStatus(context.Context, *GetJobStatusRequest) (*GetJobStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJobStatus not implemented")
}
func (UnimplementedDataproxyServiceServer) mustEmbedUnimplementedDataproxyServiceServer() {}
func (UnimplementedDataproxyServiceServer) testEmbeddedByValue()                          {}

// UnsafeDataproxyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataproxyServiceServer will
// result in compilation errors.
type UnsafeDataproxyServiceServer interface {
	mustEmbedUnimplementedDataproxyServiceServer()
}

func RegisterDataproxyServiceServer(s grpc.ServiceRegistrar, srv DataproxyServiceServer) {
	// If the following call pancis, it indicates UnimplementedDataproxyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DataproxyService_ServiceDesc, srv)
}

func _DataproxyService_CreateBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBulkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServiceServer).CreateBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataproxyService_CreateBulk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServiceServer).CreateBulk(ctx, req.(*CreateBulkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataproxyService_GetBulks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBulksReuqest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServiceServer).GetBulks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataproxyService_GetBulks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServiceServer).GetBulks(ctx, req.(*GetBulksReuqest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataproxyService_GetBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBulkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServiceServer).GetBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataproxyService_GetBulk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServiceServer).GetBulk(ctx, req.(*GetBulkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataproxyService_GetJobStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataproxyServiceServer).GetJobStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataproxyService_GetJobStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataproxyServiceServer).GetJobStatus(ctx, req.(*GetJobStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataproxyService_ServiceDesc is the grpc.ServiceDesc for DataproxyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataproxyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "io.clbs.openhes.pbdataproxy.DataproxyService",
	HandlerType: (*DataproxyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBulk",
			Handler:    _DataproxyService_CreateBulk_Handler,
		},
		{
			MethodName: "GetBulks",
			Handler:    _DataproxyService_GetBulks_Handler,
		},
		{
			MethodName: "GetBulk",
			Handler:    _DataproxyService_GetBulk_Handler,
		},
		{
			MethodName: "GetJobStatus",
			Handler:    _DataproxyService_GetJobStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pbdataproxy.proto",
}
