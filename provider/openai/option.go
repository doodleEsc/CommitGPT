// Package openai
package openai

import (
	"fmt"

	openai "github.com/sashabaranov/go-openai"
	"gopkg.in/yaml.v2"
)

const (
	defaultBaseURL    = "https://api.openai.com/v1"
	defaultAPIKey     = "you-api-key"
	defaultModel      = openai.GPT3Dot5Turbo
	defaultTemprature = 0.8
	defaultTopP       = 1.0
	defaultMaxTokens  = 0
)

type Option func(*Client)

func WithBaseURL(baseURL string) Option {
	return func(client *Client) {
		client.BaseURL = baseURL
	}
}

func WithAPIKey(apiKey string) Option {
	return func(client *Client) {
		client.APIKey = apiKey
	}
}

func WithModel(model string) Option {
	return func(client *Client) {
		client.Model = model
	}
}

func WithTemperature(temperature float32) Option {
	return func(client *Client) {
		client.Temperature = temperature
	}
}

func WithTopP(topP float32) Option {
	return func(client *Client) {
		client.TopP = topP
	}
}

func WithMaxTokens(maxTokens int) Option {
	return func(client *Client) {
		client.MaxTokens = maxTokens
	}
}

func GetDefaultConfigAsYAML() (string, error) {
	config := map[string]any{
		"openai": map[string]any{
			"base_url":    defaultBaseURL,
			"api_key":     defaultAPIKey,
			"model":       defaultModel,
			"temperature": defaultTemprature,
			"top_p":       defaultTopP,
			"max_tokens":  defaultMaxTokens,
		},
	}

	yamlBytes, err := yaml.Marshal(config)
	if err != nil {
		return "", fmt.Errorf("failed to marshal config to YAML: %w", err)
	}

	return string(yamlBytes), nil
}
