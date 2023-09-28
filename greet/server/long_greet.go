package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/svrajput/grpc_go_course/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	fmt.Printf(" LogGreet is invoked at server ..")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Response: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading the client stream %v\n", err)
		}
		res += fmt.Sprintf("Hello %s\n", req.FirstName)
		log.Printf("Received Request %s \n", req.FirstName)
	}

}
