package main

import (
	"context"
	"io"
	"log"

	pb "github.com/svrajput/grpc_go_course/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {

	log.Println("doGreetManyTimes was involved")

	req := &pb.GreetRequest{
		FirstName: "Sunil ",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v\n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Response)
	}

}
