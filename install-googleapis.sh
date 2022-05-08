#!/bin/sh
set -e

go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger@latest
go install github.com/golang/protobuf/protoc-gen-go@latest

wget http://central.maven.org/maven2/com/google/api/grpc/googleapis-common-protos/0.0.3/googleapis-common-protos-0.0.3.jar
jar xvf googleapis-common-protos-0.0.3.jar
cp -r google/ $HOME/protobuf/include/
ls -l



