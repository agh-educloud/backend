#!/bin/bash

mkdir -p src/agh/educloud/generated

for filename in protos/*.proto; do
    protoc protos/${filename##*/} --go_out=plugins=grpc:src/agh/educloud/generated
done
