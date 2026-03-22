package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("You: ")

	for scanner.Scan() {

		userInput := scanner.Text()

		if userInput == "exit" {
			break
		}

		resp, err := http.Post(
			"http://localhost:8080",      //url
			"text/plain",                 // content type
			strings.NewReader(userInput), // io.Reader reads the input into a stream of bytes to place in buffer
		)

		handleError(err)

		body, err := io.ReadAll(resp.Body) // reads the resp.Body to an array of bytes

		handleError(err)
		resp.Body.Close()

		fmt.Printf("Chani in Go: %s is not a word. Lets begin. \n", string(body))

		fmt.Print("You: ")
	}

	fmt.Println("Chani goes offline")
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
