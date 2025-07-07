// Package openai
package openai

import (
	"context"

	"github.com/doodleEsc/CommitGPT/provider/types"
	openai "github.com/sashabaranov/go-openai"
)

const (
	defaultBaseURL    = "https://api.openai.com/v1"
	defaultModel      = openai.GPT3Dot5Turbo
	defaultTemprature = 0.8
)

type Client struct {
	BaseURL     string
	APIKey      string
	Model       string
	Temperature float32
}

func New(opts ...Option) (*Client, error) {
	client := &Client{
		BaseURL:     defaultBaseURL,
		APIKey:      "",
		Model:       defaultModel,
		Temperature: defaultTemprature,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client, nil
}

func (c *Client) Completion(ctx context.Context, content string) (*types.Response, error) {
	config := openai.DefaultConfig(c.APIKey)
	config.BaseURL = c.BaseURL

	client := openai.NewClientWithConfig(config)

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:       c.Model,
			Temperature: c.Temperature,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a helpful assistant.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	respContent := resp.Choices[0].Message.Content

	response := &types.Response{
		Content: respContent,
		Usage: types.Usage{
			PromptTokens:     resp.Usage.PromptTokens,
			CompletionTokens: resp.Usage.CompletionTokens,
			TotalTokens:      resp.Usage.TotalTokens,
		},
	}

	return response, nil
}
