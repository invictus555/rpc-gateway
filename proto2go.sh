#!/bin/bash

protoc --proto_path=$GOPATH/pkg/mod/github.com/googleapis:proto:. \
--go_out=. \
--go-grpc_out=. \
--grpc-gateway_out=. \
proto/*.proto
