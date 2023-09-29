package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/svrajput/grpc_go_course/greet/proto"
)

func doGreetEveryOne(c pb.GreetServiceClient) {
	log.Printf("doGreetEveryOne is invoked ... \n ")

	reqs := []*pb.GreetRequest{
		{FirstName: "Sunil"},
		{FirstName: "Arjeet"},
		{FirstName: "Sam"},
		{FirstName: "Albert"},
	}

	stream, err := c.GreetEveryOne(context.Background())

	if err != nil {
		log.Printf("Error while invoking GreetEveryOne from client ..")
	}

	ch := make(chan struct{})

	go func() {
		for _, req := range reqs {
			err := stream.Send(req)

			if err != nil {
				log.Printf(" Error while sending value : %s", req)
			}

			log.Printf("Send Msg %s \n", req.FirstName)

			time.Sleep(1 * time.Second)

		}
		stream.CloseSend()
	}()

	go func() {
		for {
			msg, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf(" Error while receving msg from server %s\n", err)
				break
			}

			log.Printf("Received MSG from server : %s", msg.Response)
		}

		close(ch)

	}()

	<-ch
}
