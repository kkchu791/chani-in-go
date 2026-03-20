package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	openConnection := true

	for openConnection {

		// create buffer for reader
		reader := bufio.NewReader(os.Stdin)

		//get user input
		fmt.Println("You:")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if userInput == "exit" {
			fmt.Print("here?")
			openConnection = false
		}

		fmt.Println("Chani in Go: Hey, ready to begin?")
		fmt.Printf("You wrote: %s", userInput)

	}

	fmt.Print("Chani goes offline")
}
