// Editions version of proto3 file

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: pbdriver.proto

package pbdriver

import (
	context "context"
	pbdrivermodels "github.com/cybroslabs/hes-2-apis/protobuf/pbdrivermodels"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	DriverService_StartJob_FullMethodName  = "/io.clbs.openhes.pbdriver.DriverService/StartJob"
	DriverService_CancelJob_FullMethodName = "/io.clbs.openhes.pbdriver.DriverService/CancelJob"
)

// DriverServiceClient is the client API for DriverService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The driver service definition.
// Those are the gRPC services that all drivers must implement to provide required control for the Taskmaster.
type DriverServiceClient interface {
	// The method called by the Taskmaster to start a new job. The parameter contains the job specification and the list of actions to be executed.
	StartJob(ctx context.Context, in *pbdrivermodels.StartJobsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[pbdrivermodels.ProgressUpdate], error)
	// The method called by the Taskmaster to cancel the job. The parameter contains the job identifier.
	CancelJob(ctx context.Context, in *pbdrivermodels.CancelJobRequest, opts ...grpc.CallOption) (*pbdrivermodels.CommonResponse, error)
}

type driverServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDriverServiceClient(cc grpc.ClientConnInterface) DriverServiceClient {
	return &driverServiceClient{cc}
}

func (c *driverServiceClient) StartJob(ctx context.Context, in *pbdrivermodels.StartJobsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[pbdrivermodels.ProgressUpdate], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &DriverService_ServiceDesc.Streams[0], DriverService_StartJob_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[pbdrivermodels.StartJobsRequest, pbdrivermodels.ProgressUpdate]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DriverService_StartJobClient = grpc.ServerStreamingClient[pbdrivermodels.ProgressUpdate]

func (c *driverServiceClient) CancelJob(ctx context.Context, in *pbdrivermodels.CancelJobRequest, opts ...grpc.CallOption) (*pbdrivermodels.CommonResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(pbdrivermodels.CommonResponse)
	err := c.cc.Invoke(ctx, DriverService_CancelJob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DriverServiceServer is the server API for DriverService service.
// All implementations must embed UnimplementedDriverServiceServer
// for forward compatibility.
//
// The driver service definition.
// Those are the gRPC services that all drivers must implement to provide required control for the Taskmaster.
type DriverServiceServer interface {
	// The method called by the Taskmaster to start a new job. The parameter contains the job specification and the list of actions to be executed.
	StartJob(*pbdrivermodels.StartJobsRequest, grpc.ServerStreamingServer[pbdrivermodels.ProgressUpdate]) error
	// The method called by the Taskmaster to cancel the job. The parameter contains the job identifier.
	CancelJob(context.Context, *pbdrivermodels.CancelJobRequest) (*pbdrivermodels.CommonResponse, error)
	mustEmbedUnimplementedDriverServiceServer()
}

// UnimplementedDriverServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDriverServiceServer struct{}

func (UnimplementedDriverServiceServer) StartJob(*pbdrivermodels.StartJobsRequest, grpc.ServerStreamingServer[pbdrivermodels.ProgressUpdate]) error {
	return status.Errorf(codes.Unimplemented, "method StartJob not implemented")
}
func (UnimplementedDriverServiceServer) CancelJob(context.Context, *pbdrivermodels.CancelJobRequest) (*pbdrivermodels.CommonResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelJob not implemented")
}
func (UnimplementedDriverServiceServer) mustEmbedUnimplementedDriverServiceServer() {}
func (UnimplementedDriverServiceServer) testEmbeddedByValue()                       {}

// UnsafeDriverServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DriverServiceServer will
// result in compilation errors.
type UnsafeDriverServiceServer interface {
	mustEmbedUnimplementedDriverServiceServer()
}

func RegisterDriverServiceServer(s grpc.ServiceRegistrar, srv DriverServiceServer) {
	// If the following call pancis, it indicates UnimplementedDriverServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DriverService_ServiceDesc, srv)
}

func _DriverService_StartJob_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(pbdrivermodels.StartJobsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DriverServiceServer).StartJob(m, &grpc.GenericServerStream[pbdrivermodels.StartJobsRequest, pbdrivermodels.ProgressUpdate]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type DriverService_StartJobServer = grpc.ServerStreamingServer[pbdrivermodels.ProgressUpdate]

func _DriverService_CancelJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(pbdrivermodels.CancelJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DriverServiceServer).CancelJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DriverService_CancelJob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DriverServiceServer).CancelJob(ctx, req.(*pbdrivermodels.CancelJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DriverService_ServiceDesc is the grpc.ServiceDesc for DriverService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DriverService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "io.clbs.openhes.pbdriver.DriverService",
	HandlerType: (*DriverServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CancelJob",
			Handler:    _DriverService_CancelJob_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StartJob",
			Handler:       _DriverService_StartJob_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pbdriver.proto",
}
