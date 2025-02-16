package types

// ProviderType represents the type of AI provider
type ProviderType string

const (
	OpenAIProvider    ProviderType = "openai"
	AnthropicProvider ProviderType = "anthropic"
	GeminiProvider    ProviderType = "gemini"
)

// Config represents the configuration for AI providers
type Config struct {
	Type      ProviderType
	APIKey    string
	BaseURL   string
	MaxTokens int
	Model     string
}
