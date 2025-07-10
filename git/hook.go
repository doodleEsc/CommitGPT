// Package git
package git

import (
	"embed"
)

//go:embed templates/*
var files embed.FS

const (
	HookPrepareCommitMessageTemplate = "prepare-commit-msg"
)

func GetHookFileContent(filename string) ([]byte, error) {
	content, err := files.ReadFile("templates/prepare-commit-msg")
	if err != nil {
		return nil, err
	}

	return content, nil
}
