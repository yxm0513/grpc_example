import grpc
from service import service_pb2
from service import service_pb2_grpc

def run():
    with grpc.insecure_channel('localhost:10001') as channel:
        stub = service_pb2_grpc.DemoStub(channel)
        response = stub.GetInfo(service_pb2.Request(id='0'))
        print("{}".format(response))
        response = stub.GetInfo(service_pb2.Request(id='1'))
        print("{}".format(response))

if __name__ == '__main__':
    run()
