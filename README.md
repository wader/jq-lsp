# jq-lsp

jq language server.

You probably want to use this via one of these:
- [vscode-jq](https://github.com/wader/vscode-jq)
- neovim using [nvim-lspconfig](https://github.com/neovim/nvim-lspconfig/blob/master/doc/configs.md#jqls) and [mason.nvim](https://github.com/williamboman/mason.nvim)
- [Emacs lsp-mode](https://github.com/emacs-lsp/lsp-mode)
- [helix](https://github.com/helix-editor/helix) (since 25.01)
- [IntelliJ with LSP4IJ](https://github.com/redhat-developer/lsp4ij/blob/main/docs/user-defined-ls/jq-lsp.md)

It can currently do:
- Syntax checking
- Error on missing function and binding
- Goto definition of function and binding
- Auto complete function and binding
- Include/Import support
- Hover definition of function
- Hover documentation for builtin
- Function symbols per document
- Additional builtins using  `.jq-lsp.jq`

## Install

```sh
# install latest release
go install github.com/wader/jq-lsp@latest
# install master
go install github.com/wader/jq-lsp@master

# copy binary to $PATH
cp "$(go env GOPATH)/bin/jq-lsp" /usr/local/bin

# build binary from cloned repo
go build -o jq-lsp .
```

## Additional builtins using `.jq-lsp.jq`

To make jq-lsp aware of additional builtin function and variables you can use a `.jq-lsp.jq` file. The file is a normal jq file that is included automatically in each file.

This `.jq-lsp.jq` file add a function `fetch` and the variable `$OPTIONS`:
```jq
# function body is ignored but has to be valid jq
def fetch: empty;    # adds function fetch/0
def $OPTIONS: empty; # adds variable $OPTIONS
```

## Development

```sh
# run all tests
go test -v ./...
# run tests and update responses
go test -v ./... -update
```

## Thanks

jq-lsp uses a modified version of
[itchyny](https://github.com/itchyny)'s [gojq](https://github.com/itchyny/gojq) parser.

builtins documentation is based on https://github.com/stedolan/jq/blob/master/docs/content/manual/manual.yml
and is licensed under https://github.com/stedolan/jq/blob/master/COPYING

## TODO and ideas

- Signature help
- Semantic tokens
- Own parser or modified gojq parser to be able to recover and give more useful errors
- Server loop and https://github.com/itchyny/gojq/issues/86
- Warn about unused functions and bindings
- Better at handling broken syntax while typing
   - `$<cursor>`/`@<cursor>` auto complete, add temp name?
   - `a | n<cursor> f`, add temp name/try fix syntax?
- Transitive include behavior (jq/gojq behaves differently?)
- Source formatting
    - Requires whitespace/comment support in lexer/parser
- Goto definition for builtin functions somehow
- Input completion. How to indicate input and to do safe eval?
- REPL or some kind of eval support?
    - See https://github.com/ailisp/commonlisp-vscode
