package api

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/rs/cors"

	server "chani-in-go/internal/infra/http"
	"chani-in-go/internal/platform/errors"
)

func Run() {
	err := godotenv.Load()
	errors.HandleError(err)

	srv := server.NewServer()

	mux := http.NewServeMux()
	srv.SetupRoutes(mux)

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
