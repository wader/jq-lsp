module github.com/wader/jq-lsp

go 1.18

// GOPROXY=direct go get -d github.com/itchyny/gojq@main && go mod tidy

require (
	github.com/itchyny/gojq v0.12.10-0.20221114112817-6a683bb80e46
	github.com/pmezard/go-difflib v1.0.0
)

require github.com/itchyny/timefmt-go v0.1.4 // indirect
