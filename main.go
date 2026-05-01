package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"net/http"

	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	handleError(err)

	server := NewServer()

	//setup routes
	mux := http.NewServeMux()
	server.SetupRoutes(mux)

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
