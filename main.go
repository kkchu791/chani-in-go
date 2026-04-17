package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"net/http"

	"github.com/rs/cors"
)

type MessageRequest struct {
	Message string `json:"message"`
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	var data MessageRequest

	// Decode the request body into the struct
	err := json.NewDecoder(r.Body).Decode(&data)
	handleError(err)

	// Use the decoded data...
	resp := SendMessageToGroq(data.Message)

	chaniResp := resp.Choices[0].Message.Content

	w.Write([]byte(chaniResp))

}

func main() {
	err := godotenv.Load()
	handleError(err)

	mux := http.NewServeMux()

	mux.HandleFunc("/message", messageHandler)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
		Debug:            true,
	})

	handler := c.Handler(mux)

	http.ListenAndServe(":9080", handler)
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
