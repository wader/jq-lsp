module github.com/wader/jq-lsp

go 1.18

// GOPROXY=direct go get -d github.com/itchyny/gojq@main && go mod tidy

require (
	github.com/itchyny/gojq v0.12.14
	github.com/pmezard/go-difflib v1.0.0
)

require github.com/itchyny/timefmt-go v0.1.5 // indirect
