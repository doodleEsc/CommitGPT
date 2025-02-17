package openai

import (
	"context"

	"github.com/doodleEsc/CommitGPT/provider/types"
	openai "github.com/sashabaranov/go-openai"
)

const (
	defaultBaseUrl    = "https://api.openai.com/v1"
	defaultModel      = openai.GPT3Dot5Turbo
	defaultTemprature = 0.8
)

type Client struct {
	BaseUrl     string
	ApiKey      string
	Model       string
	Temperature float32
}

func New(opts ...Option) (*Client, error) {
	client := &Client{
		BaseUrl:     defaultBaseUrl,
		ApiKey:      "",
		Model:       defaultModel,
		Temperature: defaultTemprature,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client, nil
}

func (c *Client) Completion(ctx context.Context, content string) (*types.Response, error) {
	config := openai.DefaultConfig(c.ApiKey)
	config.BaseURL = c.BaseUrl

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

	resp_content := resp.Choices[0].Message.Content

	response := &types.Response{
		Content: resp_content,
		Usage: types.Usage{
			PromptTokens:     resp.Usage.PromptTokens,
			CompletionTokens: resp.Usage.CompletionTokens,
			TotalTokens:      resp.Usage.TotalTokens,
		},
	}

	return response, nil
}
