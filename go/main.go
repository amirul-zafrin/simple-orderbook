package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWS(ws *websocket.Conn) {
	fmt.Println("new incoming connection from client: ", ws.RemoteAddr())

	for {
		payload := fmt.Sprintf("data --> %d\n", time.Now().UnixNano())
		ws.Write([]byte(payload))
		time.Sleep(time.Second * 2)
	}
}

func main() {
	server := NewServer()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Welcome to Simple Orderbook!")
		io.WriteString(w, "Welcome to Simple Orderbook!")
	})

	http.Handle("/orderbook/ws", websocket.Handler(server.handleWS))

	log.Fatal(http.ListenAndServe(":3000", nil))
}
