package prompt

import (
	"embed"
	"log"

	"github.com/doodleEsc/CommitGPT/utils"
)

//go:embed templates/*
var templatesFS embed.FS

// Template file names
const (
	SystemTemplate = "system.toml"
	CommitTemplate = "commit.toml"
)

// Initializes the prompt package by loading the templates from the embedded file system.
func init() { //nolint:gochecknoinits
	if err := utils.LoadTemplates(templatesFS); err != nil {
		log.Fatal(err)
	}
}

// GetRawData returns the raw data of the template with the given name.
func GetRawData(name string) ([]byte, error) {
	key := "templates/" + name
	return templatesFS.ReadFile(key)
}
