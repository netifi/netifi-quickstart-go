all:	build

build:
	go build

proto:
    protoc --go_out=plugins=rrpc:./hello_service ./service.proto

test:
	go test -v ./...

clean:
	go clean ./...

nuke:
	go clean -i ./...
