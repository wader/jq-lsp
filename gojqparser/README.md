This is a modified version of gojq's parser https://github.com/itchyny/gojq which adds
- token position
- query marshal/unmarshal
- `def @name` support, used by jaq for custom @formats
- `def $name` support, used by jq-lsp to define own globals in .jq-lsp.jq
