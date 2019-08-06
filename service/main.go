package main

import (
	"context"
	"github.com/netifi/netifi-go"
	service "github.com/netifi/netifi-go-quickstart"
)

const accessKey = 9007199254740991
const accessToken = "kTBDVtfRBO4tHOnZzSyY5ym2kfY="

type helloService struct {
}

func (helloService) SayHello(c context.Context, s *service.HelloRequest, b []byte) (<-chan *service.HelloResponse, <-chan error) {
	r := make(chan *service.HelloResponse, 1)
	e := make(chan error)

	defer func() {
		close(r)
		close(e)
	}()
	response := &service.HelloResponse{
		Message: "Hello, " + s.Name,
	}

	r <- response
	return r, e
}

func main() {
	client, e := netifi.
		New().
		Group("helloGoService").
		Uri("tcp://localhost:8001").
		AccessKey(accessKey).
		AccessTokenBase64(accessToken).
		Build()

	if e != nil {
		panic(e)
	}

	server := service.NewHelloServiceServer(&helloService{})
	e = client.AddService(server)
	if e != nil {
		panic(e)
	}

	<-context.Background().Done()
}
