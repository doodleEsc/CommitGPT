package openai

import (
	"context"

	"github.com/doodleEsc/CommitGPT/provider/types"
)

type Provider struct {
	// client *Client
	config types.Config
}

func NewProvider(config types.Config) (*Provider, error) {
	// TODO: finish me
}

func (p *Provider) Completion(ctx context.Context, content string) (*types.Response, error) {
	// Implementation
	return nil, nil
}
