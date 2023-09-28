package main

import (
	"fmt"
	"log"

	pb "github.com/svrajput/grpc_go_course/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("Greet Many times function was invoked with %v \n ", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s , number %d", in.FirstName, i)
		stream.Send(&pb.GreetResponse{Response: res})
	}

	return nil
}
