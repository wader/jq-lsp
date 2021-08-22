#!/usr/bin/env jq -rnf

[
  { "id": 0,
    "jsonrpc": "2.0",
    "method": "initialize"
  },
  {
    "id": 1,
    "jsonrpc": "2.0",
    "method": "textDocument/didOpen",
    "params": {
      "textDocument": {
        "languageId": "plaintext",
        "text": "def aaa:111; aaa",
        "uri": "file:///a/b/c.jq",
        "version": 1
      }
    }
  },
  {
    "id": 2,
    "jsonrpc": "2.0",
    "method": "textDocument/didChange",
    "params": {
      "contentChanges": [{"text": "include \"a\";\naa"}],
      "textDocument": {
        "uri": "file://\(env.URI)/tests/b.jq",
        "version": 2
      }
    }
  },
  {
    "id": 3,
    "jsonrpc": "2.0",
    "method": "textDocument/definition",
    "params": {
      "position": {"character": 2, "line": 1},
      "textDocument": {
        "uri": "file://\(env.URI)/tests/b.jq"
      }
    }
  }

][]
| tojson
# +1 for jq string output newline
# can't use -0, not in jq 1.6
| ["Content-Length: \(utf8bytelength+1)\r\n\r\n", .]
| join("")