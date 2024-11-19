# jq-lsp

jq language server.

You probably want to use this via one of these:
- [vscode-jq](https://github.com/wader/vscode-jq)
- neovim using [nvim-lspconfig](https://github.com/neovim/nvim-lspconfig/blob/master/doc/server_configurations.md#jqls) and [mason.nvim](https://github.com/williamboman/mason.nvim)
- Emacs using [lsp-mode](https://github.com/emacs-lsp/lsp-mode)

It can currently do:
- Syntax checking.
- Error on missing function or binding.
- Goto definition of function or binding.
- Auto complete function and binding.
- Include/Import support.
- Hover definition of function.
- Hover documentation for builtin.
- Function symbols per document.
- Additional builtins using  `.jq-lsp.jq`.

## Install

```sh
# install directly
go install github.com/wader/jq-lsp@master
# copy binary to $PATH
cp "$(go env GOPATH)/bin/jq-lsp" /usr/local/bin

# build binary from cloned repo
go build -o jq-lsp .
```

## Additional builtins using `.jq-lsp.jq`

You can make jq-lsp aware of additional builtin function and variables by using a `.jq-lsp.jq` file. The file is normal jq file and will be included automatically in each file.

For example this `.jq-lsp.jq` file adds function `fetch` and the variable `$OPTIONS`:
```jq
# body will be ignored so can be any valid jq
def fetch: empty;
def $OPTIONS: empty;
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
