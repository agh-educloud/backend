#!/bin/bash

protoc protos/user.proto  --go_out=plugins=grpc:generated
protoc protos/chat.proto  --go_out=plugins=grpc:generated
protoc protos/status.proto  --go_out=plugins=grpc:generated
protoc protos/class.proto  --go_out=plugins=grpc:generated