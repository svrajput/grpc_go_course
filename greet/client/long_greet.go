package main

import (
	"context"
	"log"
	"time"

	pb "github.com/svrajput/grpc_go_course/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Printf(" doLongGreet is invoked ... \n ")

	reqs := []*pb.GreetRequest{
		{FirstName: "Sunil"},
		{FirstName: "Sam"},
		{FirstName: "Rajesh"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Printf(" Error while invoking longGreet from client...")
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from long greet \n ")
	}

	log.Printf("long greet completed %s \n", res.Response)
}
