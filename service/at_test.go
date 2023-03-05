package service

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/sashabaranov/go-openai"
)

func TestChatWithContext(t *testing.T) {
	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
