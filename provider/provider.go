// Package provider
package provider

import (
	"context"
	"fmt"

	"github.com/doodleEsc/CommitGPT/provider/message"
	"github.com/doodleEsc/CommitGPT/provider/openai"
	"github.com/doodleEsc/CommitGPT/provider/types"

	"github.com/spf13/viper"
)

// LLMProvider defines an interface for generative AI operations.
type LLMProvider interface {
	Completion(ctx context.Context, messages []message.Message) (*types.Response, error)
}

// NewProvider creates a new instance of Generative based on the provided configuration
func NewProvider(providerType types.ProviderType) (LLMProvider, error) {
	switch providerType {
	case types.OpenAIProvider:
		return newOpenAIClient()
	// return openai.NewProvider(config)
	// case AnthropicProvider:
	// 	return anthropic.NewProvider(config)
	// case GeminiProvider:
	// 	return gemini.NewProvider(config)
	default:
		return nil, fmt.Errorf("unsupported provider type: %s", providerType)
	}
}

func newOpenAIClient() (LLMProvider, error) {
	return openai.New(
		openai.WithBaseURL(viper.GetString("openai.base_url")),
		openai.WithAPIKey(viper.GetString("openai.api_key")),
		openai.WithModel(viper.GetString("openai.model")),
		openai.WithTemperature(float32(viper.GetFloat64("openai.temperature"))),
	)
}
