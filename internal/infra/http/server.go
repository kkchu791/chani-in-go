package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"chani-in-go/internal/domain/model"
	"chani-in-go/internal/infra/groq"
	"chani-in-go/internal/platform/errors"
)

type Server struct {
	registry *Registry
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
	w.Header().Set("Content-Type", SSEContentType)
	w.Header().Set("Cache-Control", SSECacheControl)
	w.Header().Set("Connection", SSEConnection)

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
	var data model.MessageRequest
	err := json.NewDecoder(r.Body).Decode(&data)
	errors.HandleError(err)

	clients := s.registry.GetClientsByUserId(data.UserID)
	chunkChannel := groq.SendMessageToGroqStream(r.Context(), data.Message)

	for chunk := range chunkChannel {
		for _, cli := range clients {
			cli.Stream <- chunk
		}
	}

	for _, cli := range clients {
		cli.Stream <- SSEDone
	}
}
