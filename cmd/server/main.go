package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

// holds connected clients
type ClientManager struct {
	clients map[chan string]bool
	mutex   sync.Mutex
}

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	handleError(err)

	fmt.Fprint(w, string(body))
}

func streamHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	clientChan := make(chan string)
	flusher, ok := w.(http.Flusher)
	
	if !ok {
		http.Error(w, "streaming not supporqqted", http.StatusInternalServerError)
		return
	}

}

func main() {
	manager := &ClientManager{
		clients: make(map[chan string]bool),
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/stream", streamHandler)
	http.ListenAndServe(":8080", nil)
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
