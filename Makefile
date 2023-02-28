.PHONY: help regenerate test dependencies build checkers action

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

dependencies: ## Grab necessary dependencies for checkers
	go version
	go mod download -x
	go install -x github.com/kisielk/errcheck golang.org/x/lint/golint

regenerate: ## Re-generate lexers and parsers
	go install github.com/goccmack/gocc
	gocc -zip -o ./internal/ dot.bnf

build: ## Perform build process
	go build .

checkers: ## Run all checkers (errcheck, gofmt and golint)
	errcheck -ignore 'fmt:[FS]?[Pp]rint*' ./...
	gofmt -l -s -w .
	golint -set_exit_status
	git diff --exit-code

test: ## Perform package tests
	go test ./...

action: dependencies regenerate build test checkers ## Run steps of github action
