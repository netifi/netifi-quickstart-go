package main

import (
	"context"
	"fmt"
	"github.com/netifi/netifi-go"
	service "github.com/netifi/netifi-go-quickstart"
	"github.com/netifi/netifi-go/tags"
)

const accessKey = 9007199254740991
const accessToken = "kTBDVtfRBO4tHOnZzSyY5ym2kfY="

func main() {
	client, e := netifi.
		New().
		Group("helloGoClient").
		Uri("tcp://localhost:8001").
		AccessKey(accessKey).
		AccessTokenBase64(accessToken).
		Build()

	if e != nil {
		panic(e)
	}

	rs := client.GroupServiceSocket("helloGoService", tags.Empty())

	hsc := service.NewHelloServiceClient(rs, nil, nil)
	request := &service.HelloRequest{
		Name: "go",
	}

	res, err := hsc.SayHello(context.Background(), request)
loop:
	for {
		select {
		case r, o := <-res:
			if o {
				fmt.Println(r)
			} else {
				break loop
			}
		case e := <-err:
			if e != nil {
				fmt.Println(e)
				panic(e)
			}
		}
	}
}
