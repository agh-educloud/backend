#!/usr/bin/env bash

go get -u github.com/golang/dep/cmd/dep
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip
apt-get update
apt-get install unzip
unzip protoc-3.6.1-linux-x86_64.zip -d protoc3
mv protoc3/bin/* /usr/local/bin/
mv protoc3/include/* /usr/local/include/
rm -rf protoc-3.6.1-linux-x86_64.zip -d protoc3
go get -u github.com/golang/protobuf/protoc-gen-go
