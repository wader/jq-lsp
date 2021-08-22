module github.com/wader/jq-lsp

go 1.17

// uses fork of github.com/itchyny/gojq (jq-lsp branch) that provides token positions
// and json marshal/unmarshal support for queries
// GOPROXY=direct go get -d github.com/wader/gojq@jq-lsp && go mod tidy
require github.com/wader/gojq v0.12.1-0.20210910131636-e81b1c3eb707

require github.com/itchyny/timefmt-go v0.1.3 // indirect
