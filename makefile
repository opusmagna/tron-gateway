include META

GOPATH:=$(shell go env GOPATH)

.PHONY: proto
proto: 
	protoc \
		--proto_path=${GOPATH}/src:. \
		--go_out=. \
		--go_opt=paths=source_relative \
		--experimental_allow_proto3_optional=true \
		./core/*.proto ./api/*.proto
		

.PHONY: clean-proto
clean-proto: 
	rm -f ./core/*.pb.go && rm -f ./api/*.pb.go
		



.PHONY: tidy
tidy:
	@echo "[TIDY] Running go mod tidy"
	@go mod tidy
	
.PHONY: test
test:
	go test -v ./... -cover


.PHONY: push
push:
	@echo "[BUILD] Committing and pushing to remote repository"
	@echo " - Committing"
	@git add META
	@git commit -am "v$(VERSION)"
	@echo " - Tagging"
	@git tag v${VERSION}
	@echo " - Pushing"
	@git push --tags origin ${BRANCH}
