#!/bin/bash

mkdir -p vendor/src/github.com/generated

for filename in protos/*.proto; do
    protoc protos/${filename##*/} --go_out=plugins=grpc:vendor/src/github.com/generated
done
