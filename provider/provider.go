package provider

import (
	"context"
	"fmt"

	"github.com/doodleEsc/CommitGPT/provider/types"

	"github.com/doodleEsc/CommitGPT/provider/openai"
)

// Usage represents the token usage information
type Usage struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
}

// Response represents the response from the AI provider
type Response struct {
	Content string
	Usage   Usage
}

// Generative defines an interface for generative AI operations.
type Generative interface {
	Completion(ctx context.Context, content string) (*types.Response, error)
}

// NewProvider creates a new instance of Generative based on the provided configuration
func NewProvider(config types.Config) (Generative, error) {
	switch config.Type {
	case types.OpenAIProvider:
		return openai.NewProvider(config)
		// return openai.NewProvider(config)
	// case AnthropicProvider:
	// 	return anthropic.NewProvider(config)
	// case GeminiProvider:
	// 	return gemini.NewProvider(config)
	default:
		return nil, fmt.Errorf("unsupported provider type: %s", config.Type)
	}
}
