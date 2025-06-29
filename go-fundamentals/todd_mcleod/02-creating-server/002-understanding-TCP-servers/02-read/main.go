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
	// func NewScanner(r io.Reader) *Scanner
	// returns a scanner to read from r
	scanner := bufio.NewScanner(conn)
	// func (s *Scanner) Scan() bool
	// It returns false when the scan stops, 
	// either by reaching the end of the input or an error
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	// we never get herre
	// we have an open stream connection
	// how does the above reader know when it's done?
	fmt.Println("Code got here.")
}
