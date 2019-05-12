#!/bin/bash

mkdir -p generated

for filename in protos/*.proto; do
    protoc protos/${filename##*/} --go_out=plugins=grpc:generated
done
