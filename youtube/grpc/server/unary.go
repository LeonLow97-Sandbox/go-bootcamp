package main

import (
	"context"

	pb "github.com/LeonLow97/grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello from the server side!!!",
	}, nil
}

