package main

import (
	"fmt"
	"os"

	"github.com/wader/jq-lsp/lsp"
	"github.com/wader/jq-lsp/profile"
)

// set by release build
var version string = "dev"

func main() {
	defer profile.MaybeProfile()()

	if err := lsp.Run(lsp.Env{
		Version:  version,
		ReadFile: os.ReadFile,
		Stdin:    os.Stdin,
		Stdout:   os.Stdout,
		Stderr:   os.Stderr,
		Args:     os.Args,
		Environ:  os.Environ(),
	}); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
