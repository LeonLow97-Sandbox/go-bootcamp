package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// func Listen(network, address string) (Listener, error)
	Listener, _ := net.Listen("tcp", ":8080")

	defer Listener.Close()

	Listener.Accept()

	_, err := io.WriteString(os.Stdout, "I see you connected")
	if err != nil {
		log.Fatal(err)
	}
}
