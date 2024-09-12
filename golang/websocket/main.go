package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Server struct {
	conns map[*websocket.Conn]bool
}

type Message struct {
	Name    string `json:"name,omitempty"`
	Content string `json:"content,omitempty"`
}

func NewServer() *Server {
	return &Server{
		make(map[*websocket.Conn]bool),
	}
}

func (s *Server) handleWs(conn *websocket.Conn) {
	fmt.Println("new incomming connection", conn.LocalAddr().String())
	s.conns[conn] = true
	s.readLoop(conn)
}

func (s *Server) readLoop(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("can not read message", err)
			return
		}

		message := Message{}
		err = json.Unmarshal(p, &message)
		if err != nil {
			log.Println("can not unmarshal message", err)
			return
		}

		msg, err := json.Marshal(message)
		if err != nil {
			log.Println("can not marshal the message data ", message)
		}

		s.broadcast(msg, messageType)
	}
}

func (s *Server) broadcast(p []byte, messageType int) {
	for ws := range s.conns {
		go func(ws *websocket.Conn, messageType int) {
			if err := ws.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				return
			}
		}(ws, messageType)
	}
}

func main() {
	listener := http.NewServeMux()
	server := NewServer()
	listener.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("error upgrade request from http to ws")
		}
		server.handleWs(ws)
	})

	listener.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	add := "localhost:8080"
	http.ListenAndServe(add, listener)
}
