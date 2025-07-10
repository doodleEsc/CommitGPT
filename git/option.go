// Package git
package git

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

const (
	defaultDiffUnified = 3
	defaultIsAmend     = false
)

// Option is a function that configures a Command.
type Option func(*Command)

// WithDiffUnified is a function that generate diffs with <n> lines of context instead of the usual three.
func WithDiffUnified(val int) Option {
	return func(c *Command) {
		c.diffUnified = val
	}
}

// WithExcludeList returns an Option that sets the excludeList field of a config object to the given value.
func WithExcludeList(val []string) Option {
	return func(c *Command) {
		// If the given value is empty, do nothing.
		if len(val) == 0 {
			return
		}
		c.excludeList = val
	}
}

// WithEnableAmend returns an Option that sets the isAmend field of a config object to the given value.
func WithEnableAmend(val bool) Option {
	return func(c *Command) {
		c.isAmend = val
	}
}

func GetDefaultConfigAsYAML() (string, error) {
	config := map[string]any{
		"git": map[string]any{
			"diff_unified": defaultDiffUnified,
			"exclude":      []string{},
			"amend":        defaultIsAmend,
		},
	}

	yamlBytes, err := yaml.Marshal(config)
	if err != nil {
		return "", fmt.Errorf("failed to marshal config to YAML: %w", err)
	}

	return string(yamlBytes), nil
}
