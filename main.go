package main

import (
	"os"

	"github.com/wader/jq-lsp/lsp"
	"github.com/wader/jq-lsp/profile"
)

func main() {
	defer profile.MaybeProfile()()

	err := lsp.Run(
		os.ReadFile,
		os.Stdin,
		os.Stdout,
		os.Stderr,
		os.Environ())
	if err != nil {
		panic(err)
	}
}
