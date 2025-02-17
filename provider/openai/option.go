package openai

// type Client struct {
// 	BaseUrl     string
// 	ApiKey      string
// 	Model       string
// 	Temperature float32
// }

type Option func(*Client)

func WithBaseUrl(baseUrl string) Option {
	return func(client *Client) {
		client.BaseUrl = baseUrl
	}
}
