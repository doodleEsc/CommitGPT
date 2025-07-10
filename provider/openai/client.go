// Package openai
package openai

import (
	"context"

	"github.com/doodleEsc/CommitGPT/provider/message"
	"github.com/doodleEsc/CommitGPT/provider/types"

	openai "github.com/sashabaranov/go-openai"
)

type Client struct {
	BaseURL     string
	APIKey      string
	Model       string
	Temperature float32
	TopP        float32
	MaxTokens   int
}

func New(opts ...Option) (*Client, error) {
	client := &Client{
		BaseURL:     defaultBaseURL,
		APIKey:      "",
		Model:       defaultModel,
		Temperature: defaultTemprature,
		TopP:        defaultTopP,
		MaxTokens:   defaultMaxTokens,
	}

	for _, opt := range opts {
		opt(client)
	}

	return client, nil
}

func convertMessages(messages []message.Message) []openai.ChatCompletionMessage {
	var result []openai.ChatCompletionMessage
	for _, msg := range messages {
		result = append(result, openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: msg.Content,
		})
	}
	return result
}

func (c *Client) Completion(ctx context.Context, messages []message.Message) (*types.Response, error) {
	config := openai.DefaultConfig(c.APIKey)
	config.BaseURL = c.BaseURL

	client := openai.NewClientWithConfig(config)

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:       c.Model,
			Temperature: c.Temperature,
			Messages:    convertMessages(messages),
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
