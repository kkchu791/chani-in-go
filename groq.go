package main

import (
	"context"
	"net/http"
	"os"

	"github.com/magicx-ai/groq-go/groq"
)

func SendMessageToGroqStream(ctx context.Context, m string) <-chan string {

	chunkCh := make(chan string)

	go func() {
		cli := groq.NewClient(os.Getenv("GROQ_API_KEY"), &http.Client{})
		req := groq.ChatCompletionRequest{
			Messages: []groq.Message{
				{
					Role:    "user",
					Content: m,
				},
			},
			Model:       "llama-3.3-70b-versatile",
			MaxTokens:   150,
			Temperature: 0.7,
			TopP:        0.9,
			NumChoices:  1,
			Stream:      true,
		}

		respCh, cleanup, err := cli.CreateChatCompletionStream(ctx, req)
		if err != nil {
			close(chunkCh)
			return
		}

		defer cleanup()

		for res := range respCh {
			if res.Error != nil {
				break
			}
			chunkCh <- res.Response.Choices[0].Delta.Content
		}
		close(chunkCh)
	}()

	return chunkCh
}

func SendMessageToGroq(m string) *groq.ChatCompletionResponse {
	cli := groq.NewClient(os.Getenv("GROQ_API_KEY"), &http.Client{})
	req := groq.ChatCompletionRequest{
		Messages: []groq.Message{
			{
				Role:    "user",
				Content: m,
			},
		},
		Model:       "llama-3.3-70b-versatile",
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
