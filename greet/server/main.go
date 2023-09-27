package main

import (
	"log"
	"net"

	pb "github.com/svrajput/grpc_go_course/greet/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetServiceServer
}

var addr string = "0.0.0.0:50051"

func main() {
	handle, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen address %v", err)
	}

	log.Printf("Listing on %s \n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	err = s.Serve(handle)

	if err != nil {
		log.Fatalf("Failed to server %v \n ", err)
	}

}
