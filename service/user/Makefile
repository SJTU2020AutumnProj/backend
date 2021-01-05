GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
    
	protoc --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/user/user.proto

.PHONY: build
build: proto

	go build -o user-srv main.go

.PHONY: test
test: 

	go test -v ./handler -coverprofile=size_coverage.out


.PHONY: print
print: 

	go tool cover -html=size_coverage.out