#!/bin/bash

mkdir -p generated/protos/grpc

for filename in protos/grpc/*.proto; do
    protoc protos/${filename##*/} --go_out=plugins=grpc:generated
done
