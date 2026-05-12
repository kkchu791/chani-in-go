package groq

import (
	"context"
	"net/http"
	"os"

	"chani-in-go/internal/platform/errors"

	groqclient "github.com/magicx-ai/groq-go/groq"
)

func SendMessageToGroqStream(ctx context.Context, m string) <-chan string {
	chunkCh := make(chan string)

	go func() {
		cli := groqclient.NewClient(os.Getenv("GROQ_API_KEY"), &http.Client{})
		req := groqclient.ChatCompletionRequest{
			Messages: []groqclient.Message{
				{
					Role:    "user",
					Content: m,
				},
			},
			Model:       Model,
			MaxTokens:   MaxTokens,
			Temperature: Temperature,
			TopP:        TopP,
			NumChoices:  NumChoices,
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

func SendMessageToGroq(m string) *groqclient.ChatCompletionResponse {
	cli := groqclient.NewClient(os.Getenv("GROQ_API_KEY"), &http.Client{})
	req := groqclient.ChatCompletionRequest{
		Messages: []groqclient.Message{
			{
				Role:    "user",
				Content: m,
			},
		},
		Model:       Model,
		MaxTokens:   MaxTokens,
		Temperature: Temperature,
		TopP:        TopP,
		NumChoices:  NumChoices,
		Stream:      false,
	}

	resp, err := cli.CreateChatCompletion(req)
	errors.HandleError(err)

	return resp
}
