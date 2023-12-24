## jq-lsp

jq language server.

You probably want to use this via one of these:
- [vscode-jq](https://github.com/wader/vscode-jq)
- neovim using [nvim-lspconfig](https://github.com/neovim/nvim-lspconfig/blob/master/doc/server_configurations.md#jqls) and [mason.nvim](https://github.com/williamboman/mason.nvim)
- Emacs using [lsp-mode](https://github.com/emacs-lsp/lsp-mode)

It can currently do:
- Syntax checking
- Error on missing function and binding
- Goto definition of function and binding
- Auto complete functions and binding
- Include/Import support
- Hover definitions of functions
- Hover documentation for builtins
- Function symbols per document
- Read additional builtins from  `.jq-lsp.jq`

### Install

```sh
# build from cloned repo
go build -o jq-lsp .

# install directly
go install github.com/wader/jq-lsp@master
# copy binary to $PATH
cp "$(go env GOPATH)/bin/jq-lsp" /usr/local/bin
```

### Development

```sh
# run all tests
$ go test -v ./...
# run with debug output
DEBUG=1 go run main.go
```

### Thanks

jq-lsp uses a [fork](https://github.com/wader/gojq/tree/jq-lsp) of
[itchyny](https://github.com/itchyny)'s [gojq](https://github.com/itchyny/gojq).

jq documentation is based on https://github.com/stedolan/jq/blob/master/docs/content/manual/manual.yml
and is licensed under https://github.com/stedolan/jq/blob/master/COPYING

### TODO

- Semantic tokens
- Own parser or modified gojq parser to be able to recover and give more useful errors
- Server loop and https://github.com/itchyny/gojq/issues/86
- Warn about unused functions and bindings
- Better at handling broken syntax while typing
   - `$<cursor>` auto complete, add temp name?
   - `a | n<cursor> f`, add temp name?
- Transitive include behavior (jq/gojq behaves differently?)
- Source formatting
    - Requires whitespace/comment support in lexer/parser
- Goto definition for builtin functions somehow
- REPL or some kind of eval support?
    - See https://github.com/ailisp/commonlisp-vscode
