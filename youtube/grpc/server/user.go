package main

import (
	"context"
	"net/http"

	pb "github.com/LeonLow97/grpc/proto"
)

func (s *helloServer) SendUser(ctx context.Context, user *pb.User) (*pb.AuthenticatedResponse, error) {
	return &pb.AuthenticatedResponse{
		Status:  http.StatusOK,
		Message: "Authenticated for user " + user.Username,
	}, nil
}
