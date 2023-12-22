package main

import (
	"log"

	pb "github.com/LeonLow97/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const serverPort = ":8080"

func main() {
	conn, err := grpc.Dial("localhost"+serverPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("cannot connect to server %v", err)
	}
	defer conn.Close() // close connection between client and server once done

	client := pb.NewGreetServiceClient(conn)

	// names := &pb.NamesList{
	// 	Names: []string{"Leon", "JieWei", "Darrel"},
	// }

	callSayHello(client)
}
