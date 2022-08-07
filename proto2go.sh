#!/bin/bash

protoc --proto_path=$GOPATH/pkg/mod/github.com/googleapis:. \
--go_out=src/ \
--go-grpc_out=src/ \
--grpc-gateway_out=src/ \
src/proto/*.proto
