package main

import (
	"context"
	"log"
	"time"

	pb "github.com/LeonLow97/grpc/proto"
)

type User struct {
	Username string
	Password string
}

func sendUserStructToServer(client pb.GreetServiceClient) {
	// simulate username and password already unmarshaled from REST JSON Frontend
	u := &User{
		Username: "Leon",
		Password: "password123",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SendUser(ctx, &pb.User{
		Username: u.Username,
		Password: u.Password,
	})
	if err != nil {
		log.Fatalf("Error sending user struct: %v", err)
	}

	log.Printf("Received response: %v", res)
}
