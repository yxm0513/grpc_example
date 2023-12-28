#!/usr/bin/env bash
protoc -I=simple/ --python_out=simple/ simple/simple.proto
protoc -I=enums/ --python_out=enums/ enums/enum_example.proto
protoc -I=complex/ --python_out=complex/ complex/complex.proto
python3 -m grpc_tools.protoc -I service/ --python_out=service/ --grpc_python_out=service/ service/service.proto
