package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	handleError(err)
	openConnection := true

	for openConnection {
		// create buffer for reader
		reader := bufio.NewReader(os.Stdin)

		//get user input
		fmt.Print("\nYou:")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if userInput == "exit" {
			openConnection = false
		}

		// send this message to the server

		resp := SendMessageToGroq(userInput)
		fmt.Printf("Chani: %s", resp.Choices[0].Message.Content)

	}

	fmt.Print("Chani goes offline")
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
