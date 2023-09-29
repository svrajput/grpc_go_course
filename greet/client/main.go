package main

import (
	"log"

	pb "github.com/svrajput/grpc_go_course/greet/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var target string = "localhost:50051"

func main() {

	conn, err := grpc.Dial(target, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect grpc server %s \n", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	//doGreet(client)
	//doGreetManyTimes(client)

	// Many greet from client. stream from client
	//doLongGreet(client)

	// BiDirectional stream
	doGreetEveryOne(client)
}
