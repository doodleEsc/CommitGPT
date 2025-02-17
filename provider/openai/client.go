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
}
