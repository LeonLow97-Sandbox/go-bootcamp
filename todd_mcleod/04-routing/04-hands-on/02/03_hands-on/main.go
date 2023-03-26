package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	// func Listen(network, address string) (Listener, error)
	Listener, _ := net.Listen("tcp", ":8080")

	// Close the listener
	defer Listener.Close()

	for {
		// returns Accept() (Conn, error)
		// Conn implements the reader and writer interaace
		c, err := Listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		bufio.NewScanner(c)

		fmt.Println("Code got here.")
		_, err = io.WriteString(c, "I see you connected")
		if err != nil {
			log.Fatal(err)
		}
	}
}
