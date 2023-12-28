import asyncio
import logging

import grpc
from service import service_pb2
from service import service_pb2_grpc

class DemoService(service_pb2_grpc.DemoServicer):

    async def GetInfo(self, request, context=None):
        rsp = service_pb2.Info()
        if request.id == '0':
            rsp.message = 'test'
        else:
            rsp.message = 'other'
        return rsp

async def serve() -> None:
    server = grpc.aio.server()
    service_pb2_grpc.add_DemoServicer_to_server(DemoService(), server)
    listen_addr = '[::]:10001'
    server.add_insecure_port(listen_addr)
    logging.info("Starting server on %s", listen_addr)
    await server.start()
    await server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)
    asyncio.run(serve())
