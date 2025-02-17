package openai

type Client struct {
	BaseUrl     string
	ApiKey      string
	Model       string
	Temperature float32
}
