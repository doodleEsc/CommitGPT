// Package types
package types

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
