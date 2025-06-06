# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from acquisition import internal_pb2 as acquisition_dot_internal__pb2
from common import types_pb2 as common_dot_types__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


class DriverServiceStub(object):
    """The driver service definition.
    Those are the gRPC services that all drivers must implement to provide required control for the Taskmaster.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.StartJob = channel.unary_stream(
                '/io.clbs.openhes.services.svcdriver.DriverService/StartJob',
                request_serializer=acquisition_dot_internal__pb2.StartJobsRequest.SerializeToString,
                response_deserializer=acquisition_dot_internal__pb2.ProgressUpdate.FromString,
                _registered_method=True)
        self.CancelJob = channel.unary_unary(
                '/io.clbs.openhes.services.svcdriver.DriverService/CancelJob',
                request_serializer=common_dot_types__pb2.ListOfId.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)


class DriverServiceServicer(object):
    """The driver service definition.
    Those are the gRPC services that all drivers must implement to provide required control for the Taskmaster.
    """

    def StartJob(self, request, context):
        """The method called by the Taskmaster to start a new job. The parameter contains the job specification and the list of actions to be executed.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CancelJob(self, request, context):
        """The method called by the Taskmaster to cancel the job. The parameter contains the job identifier.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DriverServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'StartJob': grpc.unary_stream_rpc_method_handler(
                    servicer.StartJob,
                    request_deserializer=acquisition_dot_internal__pb2.StartJobsRequest.FromString,
                    response_serializer=acquisition_dot_internal__pb2.ProgressUpdate.SerializeToString,
            ),
            'CancelJob': grpc.unary_unary_rpc_method_handler(
                    servicer.CancelJob,
                    request_deserializer=common_dot_types__pb2.ListOfId.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'io.clbs.openhes.services.svcdriver.DriverService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('io.clbs.openhes.services.svcdriver.DriverService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class DriverService(object):
    """The driver service definition.
    Those are the gRPC services that all drivers must implement to provide required control for the Taskmaster.
    """

    @staticmethod
    def StartJob(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(
            request,
            target,
            '/io.clbs.openhes.services.svcdriver.DriverService/StartJob',
            acquisition_dot_internal__pb2.StartJobsRequest.SerializeToString,
            acquisition_dot_internal__pb2.ProgressUpdate.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def CancelJob(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/io.clbs.openhes.services.svcdriver.DriverService/CancelJob',
            common_dot_types__pb2.ListOfId.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
