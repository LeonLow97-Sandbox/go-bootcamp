package main

import (
	"io"
	"log"

	pb "github.com/LeonLow97/grpc/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string

	// server accepting stream from client
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{
				Messages: messages,
			})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name %v", req.Name)
		messages = append(messages, "Hello ", req.Name)
	}
}
