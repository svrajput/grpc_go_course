package main

import (
	"io"
	"log"

	pb "github.com/svrajput/grpc_go_course/greet/proto"
)

func (s *Server) GreetEveryOne(stream pb.GreetService_GreetEveryOneServer) error {
	log.Printf("Greet Everyone at server invoked..\n")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v \n ", err)
		}

		res := "Hello " + req.FirstName + " ! "

		err = stream.Send(&pb.GreetResponse{
			Response: res,
		})

		if err != nil {
			log.Fatalf("Error while sending message from server %v \n", res)
		}

	}

}
