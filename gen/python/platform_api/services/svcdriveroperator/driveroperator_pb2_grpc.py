# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from acquisition import internal_pb2 as acquisition_dot_internal__pb2
from acquisition import main_pb2 as acquisition_dot_main__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.protobuf import wrappers_pb2 as google_dot_protobuf_dot_wrappers__pb2


class DriverOperatorServiceStub(object):
    """The Driver Operator service definition.
    Those are the gRPC services that the Driver Operator provides for other components.
    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ListDrivers = channel.unary_unary(
                '/io.clbs.openhes.services.svcdriveroperator.DriverOperatorService/ListDrivers',
                request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
                response_deserializer=acquisition_dot_main__pb2.ListOfDriver.FromString,
                _registered_method=True)
        self.SetDriver = channel.unary_unary(
                '/io.clbs.openhes.services.svcdriveroperator.DriverOperatorService/SetDriver',
                request_serializer=acquisition_dot_main__pb2.Driver.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.GetDriver = channel.unary_unary(
                '/io.clbs.openhes.services.svcdriveroperator.DriverOperatorService/GetDriver',
                request_serializer=google_dot_protobuf_dot_wrappers__pb2.StringValue.SerializeToString,
                response_deserializer=acquisition_dot_main__pb2.Driver.FromString,
                _registered_method=True)
        self.SetDriverScale = channel.unary_unary(
                '/io.clbs.openhes.services.svcdriveroperator.DriverOperatorService/SetDriverScale',
                request_serializer=acquisition_dot_internal__pb2.SetDriverScaleRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)
        self.GetDriverScale = channel.unary_unary(
                '/io.clbs.openhes.services.svcdriveroperator.DriverOperatorService/GetDriverScale',
                request_serializer=acquisition_dot_internal__pb2.GetDriverScaleRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_wrappers__pb2.UInt32Value.FromString,
                _registered_method=True)
        self.StartUpgrade = channel.unary_unary(
                '/io.clbs.openhes.services.svcdriveroperator.DriverOperatorService/StartUpgrade',
                request_serializer=acquisition_dot_internal__pb2.StartUpgradeRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                _registered_method=True)


class DriverOperatorServiceServicer(object):
    """The Driver Operator service definition.
    Those are the gRPC services that the Driver Operator provides for other components.
    """

    def ListDrivers(self, request, context):
        """The method called by the RestApi to get the list of drivers.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetDriver(self, request, context):
        """The method called by the Driver to set the driver templates. The parameter contains the driver templates.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetDriver(self, request, context):
        """The method called by the RestApi to get the driver templates.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SetDriverScale(self, request, context):
        """The method called by the Taskmaster to set the driver scale.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetDriverScale(self, request, context):
        """The method called by the Taskmaster to get the driver scale.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def StartUpgrade(self, request, context):
        """The method called by the DeviceRegistry to start the driver in upgrade mode. It will provide structure upgrade between the driver versions.
        The driver is started as Kubernetes job and ends when all the structures are upgraded; which is controlled by the DeviceRegistry.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DriverOperatorServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ListDrivers': grpc.unary_unary_rpc_method_handler(
                    servicer.ListDrivers,
                    request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                    response_serializer=acquisition_dot_main__pb2.ListOfDriver.SerializeToString,
            ),
            'SetDriver': grpc.unary_unary_rpc_method_handler(
                    servicer.SetDriver,
                    request_deserializer=acquisition_dot_main__pb2.Driver.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'GetDriver': grpc.unary_unary_rpc_method_handler(
                    servicer.GetDriver,
                    request_deserializer=google_dot_protobuf_dot_wrappers__pb2.StringValue.FromString,
                    response_serializer=acquisition_dot_main__pb2.Driver.SerializeToString,
            ),
            'SetDriverScale': grpc.unary_unary_rpc_method_handler(
                    servicer.SetDriverScale,
                    request_deserializer=acquisition_dot_internal__pb2.SetDriverScaleRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'GetDriverScale': grpc.unary_unary_rpc_method_handler(
                    servicer.GetDriverScale,
                    request_deserializer=acquisition_dot_internal__pb2.GetDriverScaleRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_wrappers__pb2.UInt32Value.SerializeToString,
            ),
            'StartUpgrade': grpc.unary_unary_rpc_method_handler(
                    servicer.StartUpgrade,
                    request_deserializer=acquisition_dot_internal__pb2.StartUpgradeRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'io.clbs.openhes.services.svcdriveroperator.DriverOperatorService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('io.clbs.openhes.services.svcdriveroperator.DriverOperatorService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class DriverOperatorService(object):
    """The Driver Operator service definition.
    Those are the gRPC services that the Driver Operator provides for other components.
    """

    @staticmethod
    def ListDrivers(request,
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
            '/io.clbs.openhes.services.svcdriveroperator.DriverOperatorService/ListDrivers',
            google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            acquisition_dot_main__pb2.ListOfDriver.FromString,
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
    def SetDriver(request,
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
            '/io.clbs.openhes.services.svcdriveroperator.DriverOperatorService/SetDriver',
            acquisition_dot_main__pb2.Driver.SerializeToString,
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
    def GetDriver(request,
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
            '/io.clbs.openhes.services.svcdriveroperator.DriverOperatorService/GetDriver',
            google_dot_protobuf_dot_wrappers__pb2.StringValue.SerializeToString,
            acquisition_dot_main__pb2.Driver.FromString,
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
    def SetDriverScale(request,
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
            '/io.clbs.openhes.services.svcdriveroperator.DriverOperatorService/SetDriverScale',
            acquisition_dot_internal__pb2.SetDriverScaleRequest.SerializeToString,
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
    def GetDriverScale(request,
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
            '/io.clbs.openhes.services.svcdriveroperator.DriverOperatorService/GetDriverScale',
            acquisition_dot_internal__pb2.GetDriverScaleRequest.SerializeToString,
            google_dot_protobuf_dot_wrappers__pb2.UInt32Value.FromString,
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
    def StartUpgrade(request,
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
            '/io.clbs.openhes.services.svcdriveroperator.DriverOperatorService/StartUpgrade',
            acquisition_dot_internal__pb2.StartUpgradeRequest.SerializeToString,
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
