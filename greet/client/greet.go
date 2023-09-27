package main

import (
	"context"
	"log"

	pb "github.com/svrajput/grpc_go_course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was involved")

	response, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Sunil Rajput"})

	if err != nil {
		log.Fatalf(" could not greet %v \n ", err)
	}

	log.Printf("Greeting : %s\n", response)
}
