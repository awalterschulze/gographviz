regenerate:
	go install github.com/goccmack/gocc
	gocc -zip -o ./internal/ dot.bnf
	find . -type f -name '*.go' | xargs goimports -w

test:
	go test ./...
