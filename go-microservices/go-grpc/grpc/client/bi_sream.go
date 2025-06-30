package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/LeonLow97/grpc/proto"
)

func callHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bidirectional Streaming Started!")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("could not send names: %v", err)
	}

	// waitc is a channel used for synchronization
	waitc := make(chan struct{})

	go func() {
		for {
			// receiving stream from server
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		// close the channel when receiving is completed
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()

	// wait for the receiving goroutine to finish
	<-waitc
	log.Printf("Bidirectional streaming finished")
}
