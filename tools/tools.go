//go:build tools
// +build tools

package tools

import (
	_ "github.com/goccmack/gocc"
	_ "github.com/kisielk/errcheck"
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/tools/cmd/goimports"
)
