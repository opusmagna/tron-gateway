include META

GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto: 
	protoc \
		--proto_path=${GOPATH}/src:. \
		--go_out=. \
		--go_opt=paths=source_relative \
		--experimental_allow_proto3_optional=true \
		./*/*.proto
		


	
.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t jac-service:latest
