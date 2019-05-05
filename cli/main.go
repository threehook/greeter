package main

import (
	"fmt"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	pb "github.com/threehook/greeter/srv/proto/hello"
	"context"
)

func main() {

	cmd.Init()



	//service := grpc.NewService()
	//service.Init()

	// use the generated client stub
	//cl := hello.NewSayService("go.micro.srv.greeter", service.Client())

	// Create new greeter client
	client := pb.NewSayService("go.micro.srv.greeter", microclient.DefaultClient)

	// Set arbitrary headers in context
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	rsp, err := cl.Hello(ctx, &hello.Request{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(rsp.Msg)
}
