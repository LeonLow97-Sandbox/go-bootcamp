// create a multiplexer (mux, servemux, router, server, http mux, ...) so that your server responds to different URIs and METHODS.

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)

	// write response
	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			m := strings.Fields(ln)[0]
			url := strings.Fields(ln)[1]
			fmt.Println("***METHOD", m) // GET METHOD
			fmt.Println("URL:", url)    // URL

			if m == "GET" && url == "/print-my-name" {
				printMyName(conn)
			}
		}
		if ln == "" {
			// headers are done
			break
		}

		i++
	}
}

func printMyName(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong><div style="color:red">Leon</div></strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n") // status line
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func respond(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n") // status line
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
