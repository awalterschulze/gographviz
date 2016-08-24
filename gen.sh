#!/bin/bash
gocc dot.bnf
find . -type f -name '*.go' | xargs goimports -w
