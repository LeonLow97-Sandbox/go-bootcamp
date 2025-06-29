package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// Listening on Port 8080
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	// Constantly Listening
	for {
		// Accepts calls to render a connection
		// Conn has reader and writer interface
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}
	defer conn.Close()

	fmt.Println("Code got here...")
}
