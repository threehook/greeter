package main

import (
	"context"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	pb "github.com/threehook/greeter/srv/proto/hello"
	"log"
)

func main() {

	cmd.Init()

	// Create new greeter client
	client := pb.NewSayService(
		"go.micro.srv.greeter",
		microclient.DefaultClient,
	)

	r, err := client.Hello(context.TODO(), &pb.Request{Name: "Tonny"})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", r.Msg)
}
