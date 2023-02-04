.DEFAULT_GOAL := check
check: lint vet test ## Check project

lint: ## Lint the files
	@golint -set_exit_status ./...

vet: ## Vet the files
	@go vet ./...

test: ## Run tests with data race detector
	@go test -race ./...

init:
	@go install golang.org/x/lint/golint@latest

goversion ?= "1.19"
test_version: ## Run tests inside Docker with given version (defaults to 1.19). Example for Go1.15: make test_version goversion=1.15
	@docker run --rm -it -v $(shell pwd):/project golang:$(goversion) /bin/sh -c "cd /project && make test"
