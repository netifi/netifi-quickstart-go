all:	proto build

build:
	go build ./...
	go build cmd/client/client.go
	go build cmd/service/service.go

proto:
	protoc --go_out=plugins=rrpc:./ ./service.proto


