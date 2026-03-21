package main

import (
	"net/http"
	"os"

	"github.com/magicx-ai/groq-go/groq"
)

func SendMessageToGroq(m string) *groq.ChatCompletionResponse {

	cli := groq.NewClient(os.Getenv("GROQ_API_KEY"), &http.Client{})
	req := groq.ChatCompletionRequest{
		Messages: []groq.Message{
			{
				Role:    "user",
				Content: m,
			},
		},
		Model:       "llama-3.3-70b-versatile", // Changed from 70B → 8B (cheaper & faster)
		MaxTokens:   150,
		Temperature: 0.7,
		TopP:        0.9,
		NumChoices:  1,
		Stream:      false,
	}

	resp, err := cli.CreateChatCompletion(req)

	handleError(err)

	return resp
}
