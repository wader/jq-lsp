all: parser.go

%.go: %.go.y
	go run golang.org/x/tools/cmd/goyacc@latest -o $@ $<
