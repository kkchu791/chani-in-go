package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	registry *Registry
}

type MessageRequest struct {
	Message string `json:"message"`
	UserID  string `json:"userId"`
}

func NewServer() *Server {
	return &Server{
		registry: NewRegistry(),
	}
}

func (s *Server) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/stream", s.streamHandler)
	mux.HandleFunc("/message", s.messageHandler)
}

func (s *Server) streamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	userID := r.URL.Query().Get("user")
	cli := NewClient(userID)
	s.registry.Register(userID, cli)

	flusher, ok := w.(http.Flusher)

	if !ok {
		http.Error(w, "streaming not supported", http.StatusInternalServerError)
		return
	}

	for {
		select {
		case msg := <-cli.Stream:
			fmt.Fprintf(w, "data: %s\n\n", msg)
			flusher.Flush()
		case <-r.Context().Done():
			s.registry.Unregister(userID, cli)
			return
		}
	}
}

func (s *Server) messageHandler(w http.ResponseWriter, r *http.Request) {
	var data MessageRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	handleError(err)

	// get clients from registry

	clients := s.registry.GetClientsByUserId(data.UserID)

	// call SendMessagetoGroqStream
	chunkChannel := SendMessageToGroqStream(r.Context(), data.Message)

	for chunk := range chunkChannel {
		for _, cli := range clients {
			cli.Stream <- chunk
		}
	}

	// After the stream is fully consumed, send [DONE] once to all clients
	for _, cli := range clients {
		cli.Stream <- "[DONE]"
	}
}
