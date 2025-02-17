# Go parameters
GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get
BINARY_NAME := commitgpt
PKG := ./...

help: ## Prints help for targets with comments
	@cat $(MAKEFILE_LIST) | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build-linux: ## build dist for linux platform
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./bin/linux/$(BINARY_NAME) -v

build-windows: ## build dist for windows platform
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o ./bin/windows/$(BINARY_NAME).exe -v

build-darwin-intel: ## build dist for macos platform
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o ./bin/darwin/x86_64/$(BINARY_NAME) -v

build-darwin-arm: ## build dist for macos platform
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GOBUILD) -o ./bin/darawin/arm64/$(BINARY_NAME) -v

build: ## build CommitGPT
	make build-linux
	make build-windows
	make build-darwin-intel
	make build-darwin-arm

.PHONY: build test clean deps
