#!/bin/bash

mkdir -p vendor/github.com/generated

for filename in protos/*.proto; do
    protoc protos/${filename##*/} --go_out=plugins=grpc:vendor/github.com/generated
done
