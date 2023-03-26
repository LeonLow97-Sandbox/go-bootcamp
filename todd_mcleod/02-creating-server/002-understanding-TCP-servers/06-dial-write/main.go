package main

import (
	"fmt"
	"log"
	"net"
)

// writes to a connection
// use 02-read file to read
func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you.")
}