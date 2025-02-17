package openai

type Option func(*Client)

func WithBaseUrl(baseUrl string) Option {
	return func(client *Client) {
		client.BaseUrl = baseUrl
	}
}

func WithApiKey(apiKey string) Option {
	return func(client *Client) {
		client.ApiKey = apiKey
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
