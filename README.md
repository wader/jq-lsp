## jq-lsp

jq language server.

**NOTICE** This project is currently an experiment so beware of rough edges.

You probably don't want to use this directly but instead use it thru an IDE extension:

- [vscode-jq](https://github.com/wader/vscode-jq)

It can currently do:
- Syntax checking
- Goto definition of functions and variables
- Auto complete functions and variables
- Include/Import support
- Hover definitions of functions
- Hover documentation for builtins
- Function symbols per document

### Install

```sh
# build from cloned repo
go build -o jq-lsp main.go

# install directly
go install github.com/wader/jq-lsp@latest
# copy binary to $PATH
cp $(go env GOPATH)/bin/jq-lsp /usr/local/bin
```

### Thanks

jq-lsp uses a [fork](https://github.com/wader/gojq/tree/jq-lsp) of
[itchyny](https://github.com/itchyny)'s [gojq](github.com/itchyny/gojq).

### Development

```sh
URI="$PWD" tests/test.jq | go run main.go
```

### TODO

- Shutdown correctly?
- Own parser or modified gojq parser to be able to recover and give more useful errors
- Server loop and https://github.com/itchyny/gojq/issues/86
- Warn about unused/missing functions and variables
- Better at handling broken syntax while typing
   - `$<cursor>` auto complete, add temp name?
   - `a | n<cursor> f`, add temp name?
- Transitive include behavior (jq/gojq behaves differently?)
- Source formatting
    - Requires whitespace/comment support in lexer/parser
- Goto definition for builtin functions somehow
- REPL or some kind of eval support?
    - See https://github.com/ailisp/commonlisp-vscode
