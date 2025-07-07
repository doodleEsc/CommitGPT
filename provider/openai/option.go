// Package openai
package openai

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
