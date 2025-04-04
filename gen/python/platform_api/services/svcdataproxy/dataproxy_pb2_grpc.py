# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from acquisition import main_pb2 as acquisition_dot_main__pb2
from acquisition import shared_pb2 as acquisition_dot_shared__pb2
from common import fields_pb2 as common_dot_fields__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2
from system import main_pb2 as system_dot_main__pb2


class DataproxyServiceStub(object):
    """The Dataproxy related service definition.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ListBulks = channel.unary_unary(
                '/io.clbs.openhes.services.svcdataproxy.DataproxyService/ListBulks',
                request_serializer=common_dot_fields__pb2.ListSelector.SerializeToString,
                response_deserializer=acquisition_dot_main__pb2.ListOfBulk.FromString,
                _registered_method=True)
        self.ListBulkJobs = channel.unary_unary(
                '/io.clbs.openhes.services.svcdataproxy.DataproxyService/ListBulkJobs',
                request_serializer=acquisition_dot_main__pb2.ListBulkJobsRequest.SerializeToString,
                response_deserializer=acquisition_dot_main__pb2.ListOfBulkJob.FromString,
                _registered_method=True)
        self.GetBulkJob = channel.unary_unary(
                '/io.clbs.openhes.services.svcdataproxy.DataproxyService/GetBulkJob',
                request_serializer=google_dot_protobuf_dot_wrappers__pb2.StringValue.SerializeToString,
                response_deserializer=acquisition_dot_main__pb2.BulkJob.FromString,
                _registered_method=True)
        self.CancelBulk = channel.unary_unary(
                '/io.clbs.openhes.services.svcdataproxy.DataproxyService/CancelBulk',
                request_serializer=google_dot_protobuf_dot_wrappers__pb2.StringValue.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.CreateProxyBulk = channel.unary_unary(
                '/io.clbs.openhes.services.svcdataproxy.DataproxyService/CreateProxyBulk',
                request_serializer=acquisition_dot_main__pb2.CreateProxyBulkRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_wrappers__pb2.StringValue.FromString,
                _registered_method=True)
        self.GetProxyBulk = channel.unary_unary(
                '/io.clbs.openhes.services.svcdataproxy.DataproxyService/GetProxyBulk',
                request_serializer=google_dot_protobuf_dot_wrappers__pb2.StringValue.SerializeToString,
                response_deserializer=acquisition_dot_main__pb2.ProxyBulk.FromString,
                _registered_method=True)
        self.CreateBulk = channel.unary_unary(
                '/io.clbs.openhes.services.svcdataproxy.DataproxyService/CreateBulk',
                request_serializer=acquisition_dot_main__pb2.CreateBulkRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_wrappers__pb2.StringValue.FromString,
                _registered_method=True)
        self.GetBulk = channel.unary_unary(
                '/io.clbs.openhes.services.svcdataproxy.DataproxyService/GetBulk',
                request_serializer=google_dot_protobuf_dot_wrappers__pb2.StringValue.SerializeToString,
                response_deserializer=acquisition_dot_main__pb2.Bulk.FromString,
                _registered_method=True)
        self.GetConfig = channel.unary_unary(
                '/io.clbs.openhes.services.svcdataproxy.DataproxyService/GetConfig',
                request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
                response_deserializer=system_dot_main__pb2.SystemConfig.FromString,
                _registered_method=True)
        self.SetConfig = channel.unary_unary(
                '/io.clbs.openhes.services.svcdataproxy.DataproxyService/SetConfig',
                request_serializer=system_dot_main__pb2.SystemConfig.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.GetMeterDataRegisters = channel.unary_stream(
                '/io.clbs.openhes.services.svcdataproxy.DataproxyService/GetMeterDataRegisters',
                request_serializer=acquisition_dot_main__pb2.GetMeterDataRequest.SerializeToString,
                response_deserializer=acquisition_dot_shared__pb2.RegisterValues.FromString,
                _registered_method=True)
        self.GetMeterDataProfiles = channel.unary_stream(
                '/io.clbs.openhes.services.svcdataproxy.DataproxyService/GetMeterDataProfiles',
                request_serializer=acquisition_dot_main__pb2.GetMeterDataRequest.SerializeToString,
                response_deserializer=acquisition_dot_shared__pb2.ProfileValues.FromString,
                _registered_method=True)
        self.GetMeterDataIrregularProfiles = channel.unary_stream(
                '/io.clbs.openhes.services.svcdataproxy.DataproxyService/GetMeterDataIrregularProfiles',
                request_serializer=acquisition_dot_main__pb2.GetMeterDataRequest.SerializeToString,
                response_deserializer=acquisition_dot_shared__pb2.IrregularProfileValues.FromString,
                _registered_method=True)
        self.GetMeterEvents = channel.unary_stream(
                '/io.clbs.openhes.services.svcdataproxy.DataproxyService/GetMeterEvents',
                request_serializer=acquisition_dot_main__pb2.GetMeterEventsRequest.SerializeToString,
                response_deserializer=acquisition_dot_shared__pb2.EventRecords.FromString,
                _registered_method=True)


class DataproxyServiceServicer(object):
    """The Dataproxy related service definition.
    """

    def ListBulks(self, request, context):
        """@group: Bulks
        Retrieves the list of bulks. The list of bulks is paginated. The page size is defined in the request. The page number is 0-based.
        The list contains both the proxy bulks and the regular bulks.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ListBulkJobs(self, request, context):
        """@group: Bulks
        Retrieves the list of jobs. The list of jobs is paginated. The page size is defined in the request. The page number is 0-based.
        The listing can be used for both proxy bulks and regular bulks.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetBulkJob(self, request, context):
        """@group: Bulks
        Retrieves the job status. It can be used for jobs related to both proxy and regular bulks.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CancelBulk(self, request, context):
        """@group: Bulks
        Cancels the bulk of jobs. It can be used for both proxy and regular bulks.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateProxyBulk(self, request, context):
        """@group: Bulks
        Starts a new proxy bulk. The proxy bolk is a collection of jobs where each job represents a single device. Devices must be fully defined in the request.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetProxyBulk(self, request, context):
        """@group: Bulks
        Retrieves the proxy bulk info and status.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateBulk(self, request, context):
        """@group: Bulks
        Starts a new bulk. The bulk is a collection of jobs where each jobs represents a single device. Devices that are part of the bulk are identified either as a list of registered device identifiers or as a group identifier.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetBulk(self, request, context):
        """@group: Bulks
        Retrieves the bulk info and status.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetConfig(self, request, context):
        """// Cancels the job(s) identified by the job identifier(s).
        rpc CancelJobs(google.protobuf.ListValue) returns (google.protobuf.Empty);

        The method called by the RestApi to get the system configuration.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetConfig(self, request, context):
        """The method called by the RestApi to set the system configuration.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetMeterDataRegisters(self, request, context):
        """@group: Meter Data
        The method to stream out register-typed meter data.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetMeterDataProfiles(self, request, context):
        """@group: Meter Data
        The method to stream out profile-typed meter data.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetMeterDataIrregularProfiles(self, request, context):
        """@group: Meter Data
        The method to stream out profile-typed meter data.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetMeterEvents(self, request, context):
        """@group: Meter Events
        The method to stream out profile-typed meter data.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DataproxyServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ListBulks': grpc.unary_unary_rpc_method_handler(
                    servicer.ListBulks,
                    request_deserializer=common_dot_fields__pb2.ListSelector.FromString,
                    response_serializer=acquisition_dot_main__pb2.ListOfBulk.SerializeToString,
            ),
            'ListBulkJobs': grpc.unary_unary_rpc_method_handler(
                    servicer.ListBulkJobs,
                    request_deserializer=acquisition_dot_main__pb2.ListBulkJobsRequest.FromString,
                    response_serializer=acquisition_dot_main__pb2.ListOfBulkJob.SerializeToString,
            ),
            'GetBulkJob': grpc.unary_unary_rpc_method_handler(
                    servicer.GetBulkJob,
                    request_deserializer=google_dot_protobuf_dot_wrappers__pb2.StringValue.FromString,
                    response_serializer=acquisition_dot_main__pb2.BulkJob.SerializeToString,
            ),
            'CancelBulk': grpc.unary_unary_rpc_method_handler(
                    servicer.CancelBulk,
                    request_deserializer=google_dot_protobuf_dot_wrappers__pb2.StringValue.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'CreateProxyBulk': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateProxyBulk,
                    request_deserializer=acquisition_dot_main__pb2.CreateProxyBulkRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_wrappers__pb2.StringValue.SerializeToString,
            ),
            'GetProxyBulk': grpc.unary_unary_rpc_method_handler(
                    servicer.GetProxyBulk,
                    request_deserializer=google_dot_protobuf_dot_wrappers__pb2.StringValue.FromString,
                    response_serializer=acquisition_dot_main__pb2.ProxyBulk.SerializeToString,
            ),
            'CreateBulk': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateBulk,
                    request_deserializer=acquisition_dot_main__pb2.CreateBulkRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_wrappers__pb2.StringValue.SerializeToString,
            ),
            'GetBulk': grpc.unary_unary_rpc_method_handler(
                    servicer.GetBulk,
                    request_deserializer=google_dot_protobuf_dot_wrappers__pb2.StringValue.FromString,
                    response_serializer=acquisition_dot_main__pb2.Bulk.SerializeToString,
            ),
            'GetConfig': grpc.unary_unary_rpc_method_handler(
                    servicer.GetConfig,
                    request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                    response_serializer=system_dot_main__pb2.SystemConfig.SerializeToString,
            ),
            'SetConfig': grpc.unary_unary_rpc_method_handler(
                    servicer.SetConfig,
                    request_deserializer=system_dot_main__pb2.SystemConfig.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'GetMeterDataRegisters': grpc.unary_stream_rpc_method_handler(
                    servicer.GetMeterDataRegisters,
                    request_deserializer=acquisition_dot_main__pb2.GetMeterDataRequest.FromString,
                    response_serializer=acquisition_dot_shared__pb2.RegisterValues.SerializeToString,
            ),
            'GetMeterDataProfiles': grpc.unary_stream_rpc_method_handler(
                    servicer.GetMeterDataProfiles,
                    request_deserializer=acquisition_dot_main__pb2.GetMeterDataRequest.FromString,
                    response_serializer=acquisition_dot_shared__pb2.ProfileValues.SerializeToString,
            ),
            'GetMeterDataIrregularProfiles': grpc.unary_stream_rpc_method_handler(
                    servicer.GetMeterDataIrregularProfiles,
                    request_deserializer=acquisition_dot_main__pb2.GetMeterDataRequest.FromString,
                    response_serializer=acquisition_dot_shared__pb2.IrregularProfileValues.SerializeToString,
            ),
            'GetMeterEvents': grpc.unary_stream_rpc_method_handler(
                    servicer.GetMeterEvents,
                    request_deserializer=acquisition_dot_main__pb2.GetMeterEventsRequest.FromString,
                    response_serializer=acquisition_dot_shared__pb2.EventRecords.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'io.clbs.openhes.services.svcdataproxy.DataproxyService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('io.clbs.openhes.services.svcdataproxy.DataproxyService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class DataproxyService(object):
    """The Dataproxy related service definition.
    """

    @staticmethod
    def ListBulks(request,
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
            '/io.clbs.openhes.services.svcdataproxy.DataproxyService/ListBulks',
            common_dot_fields__pb2.ListSelector.SerializeToString,
            acquisition_dot_main__pb2.ListOfBulk.FromString,
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
    def ListBulkJobs(request,
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
            '/io.clbs.openhes.services.svcdataproxy.DataproxyService/ListBulkJobs',
            acquisition_dot_main__pb2.ListBulkJobsRequest.SerializeToString,
            acquisition_dot_main__pb2.ListOfBulkJob.FromString,
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
    def GetBulkJob(request,
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
            '/io.clbs.openhes.services.svcdataproxy.DataproxyService/GetBulkJob',
            google_dot_protobuf_dot_wrappers__pb2.StringValue.SerializeToString,
            acquisition_dot_main__pb2.BulkJob.FromString,
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
    def CancelBulk(request,
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
            '/io.clbs.openhes.services.svcdataproxy.DataproxyService/CancelBulk',
            google_dot_protobuf_dot_wrappers__pb2.StringValue.SerializeToString,
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

    @staticmethod
    def CreateProxyBulk(request,
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
            '/io.clbs.openhes.services.svcdataproxy.DataproxyService/CreateProxyBulk',
            acquisition_dot_main__pb2.CreateProxyBulkRequest.SerializeToString,
            google_dot_protobuf_dot_wrappers__pb2.StringValue.FromString,
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
    def GetProxyBulk(request,
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
            '/io.clbs.openhes.services.svcdataproxy.DataproxyService/GetProxyBulk',
            google_dot_protobuf_dot_wrappers__pb2.StringValue.SerializeToString,
            acquisition_dot_main__pb2.ProxyBulk.FromString,
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
    def CreateBulk(request,
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
            '/io.clbs.openhes.services.svcdataproxy.DataproxyService/CreateBulk',
            acquisition_dot_main__pb2.CreateBulkRequest.SerializeToString,
            google_dot_protobuf_dot_wrappers__pb2.StringValue.FromString,
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
    def GetBulk(request,
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
            '/io.clbs.openhes.services.svcdataproxy.DataproxyService/GetBulk',
            google_dot_protobuf_dot_wrappers__pb2.StringValue.SerializeToString,
            acquisition_dot_main__pb2.Bulk.FromString,
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
    def GetConfig(request,
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
            '/io.clbs.openhes.services.svcdataproxy.DataproxyService/GetConfig',
            google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            system_dot_main__pb2.SystemConfig.FromString,
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
    def SetConfig(request,
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
            '/io.clbs.openhes.services.svcdataproxy.DataproxyService/SetConfig',
            system_dot_main__pb2.SystemConfig.SerializeToString,
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

    @staticmethod
    def GetMeterDataRegisters(request,
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
            '/io.clbs.openhes.services.svcdataproxy.DataproxyService/GetMeterDataRegisters',
            acquisition_dot_main__pb2.GetMeterDataRequest.SerializeToString,
            acquisition_dot_shared__pb2.RegisterValues.FromString,
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
    def GetMeterDataProfiles(request,
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
            '/io.clbs.openhes.services.svcdataproxy.DataproxyService/GetMeterDataProfiles',
            acquisition_dot_main__pb2.GetMeterDataRequest.SerializeToString,
            acquisition_dot_shared__pb2.ProfileValues.FromString,
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
    def GetMeterDataIrregularProfiles(request,
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
            '/io.clbs.openhes.services.svcdataproxy.DataproxyService/GetMeterDataIrregularProfiles',
            acquisition_dot_main__pb2.GetMeterDataRequest.SerializeToString,
            acquisition_dot_shared__pb2.IrregularProfileValues.FromString,
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
    def GetMeterEvents(request,
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
            '/io.clbs.openhes.services.svcdataproxy.DataproxyService/GetMeterEvents',
            acquisition_dot_main__pb2.GetMeterEventsRequest.SerializeToString,
            acquisition_dot_shared__pb2.EventRecords.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
