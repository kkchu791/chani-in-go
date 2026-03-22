package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	handleError(err)

	fmt.Fprint(w, string(body))
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hey Kirk")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/kirk", handler2)
	http.ListenAndServe(":8080", nil)
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
