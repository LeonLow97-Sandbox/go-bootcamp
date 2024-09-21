package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

type Server struct {
	// conns is a map to keep track of connected websocket connections.
	// The keys are pointers to websocket connections (*websocket.Conn),
	// and the values are boolean flags indicating connection status.
	conns map[*websocket.Conn]bool
}

// NewServer creates and returns a new instance of Server.
func NewServer() *Server {
	return &Server{
		// Initializes the conns map to keep track of websocket connections.
		conns: make(map[*websocket.Conn]bool),
	}
}

// People who subscribe to this OrderBook will receive
// real-time feeds of the order books. These order books could
// come from a database
func (s *Server) handleWSOrderBook(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client to OrderBook Feed:", ws.RemoteAddr())

	for {
		payload := fmt.Sprintf("OrderBook data -> %d\n", time.Now().UnixNano())
		ws.Write([]byte(payload))
		time.Sleep(time.Second * 2)
	}
}

// handleWS handles a new WebSocket connection by adding it to the server's connections
func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client:", ws.RemoteAddr())

	s.conns[ws] = true

	// for each connection, listen to messages sent from client
	s.readLoop(ws)
}

// readLoop continuously reads messages from a WebSocket connection
func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := ws.Read(buf)

		if err != nil {
			// connection on the client closed itself
			if err == io.EOF {
				break
			}
			fmt.Println("read error:", err)
			// we don't break because we want to keep the connection alive
			// between the client and the server
			continue
		}
		msg := buf[:n]

		// let everyone know there is a new message from client
		s.broadcast(msg)
	}
}

// broadcast sends a message to all connected WebSocket clients
func (s *Server) broadcast(b []byte) {
	for ws := range s.conns {
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				// probably log to Logstash or Grafana
				fmt.Println("write error:", err)
			}
		}(ws)
	}
}

func main() {
	server := NewServer()
	http.Handle("/ws", websocket.Handler(server.handleWS))
	http.Handle("/orderbookfeed", websocket.Handler(server.handleWSOrderBook))

	fmt.Println("Server listening on port 3000!")
	http.ListenAndServe(":3000", nil)
}
