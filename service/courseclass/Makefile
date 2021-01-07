
GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
    
	protoc --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/courseclass/courseclass.proto
    

.PHONY: build
build: proto

	go build -o courseclass-service *.go


.PHONY: docker
docker:
	docker build . -t courseclass-service:latest

	.PHONY: test
test: 

	go test -v ./handler -coverprofile=size_coverage.out


.PHONY: print
print: 

	go tool cover -html=size_coverage.out
