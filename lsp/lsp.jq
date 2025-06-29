include "builtin_env";
include "docs";

def _cond(cond; f):
  if cond then
    ( (f | empty)
    , .
    )
  end;

def stdlog(f):
  ( (f | stdlog)
  , .
  );

def debug(f):
  ( ( (["DEBUG", f] | tojson)
    , "\n"
    | stderr
    )
  , .
  );
def debug: debug(.);

def loop(f):
  def _loop: f | _loop;
  _loop;

def repeat_break(f):
  try repeat(f)
  catch
    if . == "break" then empty
    else error end;

def readline:
  ( [ repeat_break(
        ( stdin(1)
        | if . == "\n" then error("break") end
        )
      )
    ]
  | join("")
  | rtrimstr("\r")
  );

def jsonrpc_read:
  def _header:
    [ repeat_break(
        ( readline
        | if . == "" then error("break") end
        | [splits(":\\s?")]
        | {(.[0] | ascii_downcase | gsub("-"; "_")): .[1]}
        )
      )
    ] | add;
  ( stdin(_header.content_length | tonumber)
  | fromjson
  );

def jsonrpc_write:
  ( tojson
  | ["Content-Length: \(utf8bytelength)\r\n\r\n", .]
  | join("")
  | stdout
  );

# line/character to position
def lc_to_byte_pos($l; $c):
  ( split("\n")
  | .[0:$l]
  | map(utf8bytelength+1)
  | add
  | . + $c
  );

# pos to line/character
def byte_pos_to_lc($pos):
  ( split("\n")
  | map(utf8bytelength+1)
  | . as $lens
  | length as $nr_lines
  | [ {i: 0, p: $pos}
    | while(
        .p >= 0 and .i < $nr_lines;
        ( .p -= $lens[.i]
        | .i += 1
        )
      )
    ]
  | length as $lines
  | { line: ($lines-1),
      character: ($pos - (($lens[0:$lines-1] | add) // 0))
    }
  );

def in_range($start; $stop): . >= $start and . <= $stop;


def TextFormatSnippet: 2;

def CompletionItemKindFunction: 3;
def CompletionItemKindVariable: 6;

def TextDocumentSyncFull: 1;
def MarkupKindsMarkdown: "markdown";

def SymbolKindFile: 1;
def SymbolKindModule: 2;
def SymbolKindNamespace: 3;
def SymbolKindPackage: 4;
def SymbolKindClass: 5;
def SymbolKindMethod: 6;
def SymbolKindProperty: 7;
def SymbolKindField: 8;
def SymbolKindConstructor: 9;
def SymbolKindEnum: 10;
def SymbolKindInterface: 11;
def SymbolKindFunction: 12;
def SymbolKindVariable: 13;
def SymbolKindConstant: 14;
def SymbolKindString: 15;
def SymbolKindNumber: 16;
def SymbolKindBoolean: 17;
def SymbolKindArray: 18;
def SymbolKindObject: 19;
def SymbolKindKey: 20;
def SymbolKindNull: 21;
def SymbolKindEnumMember: 22;
def SymbolKindStruct: 23;
def SymbolKindEvent: 24;
def SymbolKindOperator: 25;
def SymbolKindTypeParameter: 26;

def SemanticTokenTypeNamespace: "namespace";
def SemanticTokenTypeType: "type";
def SemanticTokenTypeClass: "class";
def SemanticTokenTypeEnum: "enum";
def SemanticTokenTypeInterface: "interface";
def SemanticTokenTypeStruct: "struct";
def SemanticTokenTypeTypeParameter: "typeParameter";
def SemanticTokenTypeParameter: "parameter";
def SemanticTokenTypeVariable: "variable";
def SemanticTokenTypeProperty: "property";
def SemanticTokenTypeEnumMember: "enumMember";
def SemanticTokenTypeEvent: "event";
def SemanticTokenTypeFunction: "function";
def SemanticTokenTypeMethod: "method";
def SemanticTokenTypeMacro: "macro";
def SemanticTokenTypeKeyword: "keyword";
def SemanticTokenTypeModifier: "modifier";
def SemanticTokenTypeComment: "comment";
def SemanticTokenTypeString: "string";
def SemanticTokenTypeNumber: "number";
def SemanticTokenTypeRegexp: "regexp";
def SemanticTokenTypeOperator: "operator";
def SemanticTokenTypeDecorator: "decorator";

def SemanticTokenTypeNamespaceNr: 0;
def SemanticTokenTypeTypeNr: 1;
def SemanticTokenTypeClassNr: 2;
def SemanticTokenTypeEnumNr: 3;
def SemanticTokenTypeInterfaceNr: 4;
def SemanticTokenTypeStructNr: 5;
def SemanticTokenTypeTypeParameterNr: 6;
def SemanticTokenTypeParameterNr: 7;
def SemanticTokenTypeVariableNr: 8;
def SemanticTokenTypePropertyNr: 9;
def SemanticTokenTypeEnumMemberNr: 10;
def SemanticTokenTypeEventNr: 11;
def SemanticTokenTypeFunctionNr: 12;
def SemanticTokenTypeMethodNr: 13;
def SemanticTokenTypeMacroNr: 14;
def SemanticTokenTypeKeywordNr: 15;
def SemanticTokenTypeModifierNr: 16;
def SemanticTokenTypeCommentNr: 17;
def SemanticTokenTypeStringNr: 18;
def SemanticTokenTypeNumberNr: 19;
def SemanticTokenTypeRegexpNr: 20;
def SemanticTokenTypeOperatorNr: 21;
def SemanticTokenTypeDecoratorNr: 22;

def SemanticTokenModifierDeclaration: "declaration";
def SemanticTokenModifierDefinition: "definition";
def SemanticTokenModifierReadonly: "readonly";
def SemanticTokenModifierStatic: "static";
def SemanticTokenModifierDeprecated: "deprecated";
def SemanticTokenModifierAbstract: "abstract";
def SemanticTokenModifierAsync: "async";
def SemanticTokenModifierModification: "modification";
def SemanticTokenModifierDocumentation: "documentation";
def SemanticTokenModifierDefaultLibrary: "defaultLibrary";

def SemanticTokenModifierDeclarationBit: 1;
def SemanticTokenModifierDefinitionBit: 2;
def SemanticTokenModifierReadonlyBit: 4;
def SemanticTokenModifierStaticBit: 8;
def SemanticTokenModifierDeprecatedBit: 16;
def SemanticTokenModifierAbstractBit: 32;
def SemanticTokenModifierAsyncBit: 64;
def SemanticTokenModifierModificationBit: 128;
def SemanticTokenModifierDocumentationBit: 256;
def SemanticTokenModifierDefaultLibraryBit: 512;

def env_iter_entries:
  ( reverse
  | .[]
  | keys[] as $k
  | {key: $k, value: .[$k]}
  );

def query_token:
  ( .term.func.name
  // .term.format
  // .term.break
  // .name
  // if .str then . else null end
  );

def query_args:
  if .str then [] # TODO: should assume var?
  elif .term.func then (.term.func.args // [])
  elif .term.format then []
  else null
  end;

# file:///a/b/c rel "d" -> file:///a/b/d
def uri_resolve($rel):
  ( [.[0:rindex("/")+1], $rel]
  | join("")
  );

def file_uri_to_local: ltrimstr("file://");

def func_def_signature:
  ( [ .name.str
    , if (.args | length) > 0 then
        ( "("
        , (.args | map(.str) | join("; "))
        , ")"
        )
      else empty
      end
    ]
  | join("")
  );

def func_term_name:
  if .term.func then
    ( .term.func
    | [ .name.str
      , if .args and (.args | length) > 0 then
          ( "/"
          , (.args | length)
          )
        else empty
        end
      ]
    | join("")
    )
  elif .term.format then
    .term.format.str
  end;

def env_func_name:
  ( [ .str
    , "/"
    , (.args | if . then length else 0 end)
    ]
  | join("")
  );

def env_func_signature:
  ( [ .str
    , if (.args | length) > 0 then
        ( "("
        , (.args | join("; "))
        , ")"
        )
      else empty
      end
    ]
  | join("")
  );

def env_func_markdown:
  ( docs[env_func_name] as $doc
  | [ "```jq"
    , "def \(env_func_signature):"
    , "```"
    , if $doc then $doc else empty end
    ]
  | join("\n")
  );

def query_walk($uri; $start_env; f):
  def _t($start_env):
    def _pattern_env:
      def _f:
        ( if .name then
            {(.name.str): (.name + {type: "binding", uri: $uri, args: []})}
          elif .array then
            .array[] | _f
          elif .object then
            ( .object[] |
              ( if .key then
                  { (.key.str):
                    (.key + {type: "binding", uri: $uri, args: []})
                  }
                else empty
                end
              , if .val then .val | _f else empty end
              )
            )
          else error("_pattern_env unreachable")
          end
        );
      ( [_f]
      | add
      );
    def _pattern_traverse($env):
      ( if .name then .
        elif .array then
          ( .array[]
          )
        elif .object then
          ( .object[]
          | ( (.key_query // empty)
            , (.key // empty)
            , (.val // empty)
            )
          )
        else error("_pattern_traverse unreachable")
        end
      | _t($env)
      );

    def _func_args: map(.str);

    def _func_def_env($uri; $import_alias):
      ( (.args // []) as $args
      | ( if $import_alias then
            { name: "\($import_alias.str)::\(.name.str)"
            , start: .name.start
            , stop: .name.stop
            }
          else
            { name: .name.str
            , start: .name.start
            , stop: .name.stop
            }
          end
        ) as $f
      # {"(prefix::)?func/2":  name token (has .str etc) + {args: #}
      | { "\($f.name)/\($args | length)":
            { str: $f.name
            , start: $f.start
            , stop: $f.stop
            , type: "function"
            , uri: $uri
            , args: ($args | _func_args)
            }
        }
      );
    def _func_def_env($uri): _func_def_env($uri; null);

    def _import_env:
      # TODO: import failure
      # TODO: transitive include, max depth
      ( . #debug({import: .})
      | [ (.include_path.str // .import_path.str) as $path
        | if .meta.keyvals then
            # [key: {str: "a"}, val: {str: "b"}] => {a: [b]}
            # [key: {str: "a"}, val: [{array: {elems: [{str: "b"}]}}] => {a: [b]}
            ( .meta.keyvals
            | map(
                ( .key = .key.str
                | .value =
                    ( .val
                    | if .str then [.str.str]
                      else
                        ( .array.elems
                        | map(.str.str)
                        )
                      end
                    )
                )
              )
            # {search: ["a","b"]} => "a/path", "b/path
            | from_entries
            | .search[]?
            | "\(.)/\($path)"
            )
          else $path
          end
        ] as $paths
      | .import_alias as $import_alias
      | [ $paths[] as $path
        | ($uri | uri_resolve($path)) + ".jq"
        ] as $include_uris
      | if $import_alias | (. == null | not) and (.str | startswith("$")) then
          # import "f" as $name
          ( $import_alias
          | { (.str):
                { str: .str
                , start: .start
                , stop: .stop
                , type: "binding"
                , uri: $uri
                , args: []
                }
            }
          )
        else
          ( first(
              # include "f"
              # include "f" {search: "path"}
              # import "f" as name
              # import "f" as name {search: "path"}
              ( $include_uris[]
              | try
                  ( . as $include_uri
                  | file_uri_to_local
                  | readfile
                  | [query_fromstring, $include_uri]
                  )
                catch empty
              )
            ) as [$query, $uri]
          | $query.func_defs[]?
          | _func_def_env($uri; $import_alias)
          )
        end
      );

    def _term_traverse($env):
      ( if .suffix_list then
          ( .suffix_list[]
          | if .bind then
              ( ( (.bind.patterns // [])
                | map(_pattern_env)
                ) as $bindenvs
              | ( .bind.patterns[]?
                | _pattern_traverse($env + $bindenvs)
                )
              , ( .bind.body
                | _t($env + $bindenvs)
                )
              )
            else empty
            end
          , if .index then
              ( .index.str.queries[]?
              , (.index.start // empty)
              , (.index.end // empty)
              | _t($env)
              )
            else empty
            end
          )
        else empty
        end
      , if .index then
          ( .index.str.queries[]?
          , (.index.start // empty)
          , (.index.end // empty)
          | _t($env)
          )
        elif .func then
          ( .func.args[]?
          | _t($env)
          )
        elif .object then
          ( .object.key_vals[]?
          | ( if .key | . != null and (.str | startswith("$")) then
                # rewrite {$name} shorthand to {name: $name}
                # note key is {"key":{"start":1209,"stop":1210,"str":"$name"}
                {term: {func: {name: .key}, type: "TermTypeFunc"}}
              else .key // empty
              end
            , (.key_query // empty)
            , (.val // empty)
            )
          | _t($env)
          )
        elif .array then
          ( .array[]
          | _t($env)
          )
        elif .unary then
          ( .unary
          | _t($env)
          )
        elif .try then
          ( .try.body
          , .try.catch
          | _t($env)
          )
        elif .reduce then
          ( ( .reduce.pattern
            | [_pattern_env]
            ) as $penv
          | ( .reduce.pattern
            | _pattern_traverse($env)
            )
          , ( .reduce.query
            , .reduce.start
            , .reduce.update
            | _t($env + $penv)
            )
          )
        elif .foreach then
          ( ( .foreach.pattern
            | [_pattern_env]
            ) as $penv
          | ( .foreach.pattern
            | _pattern_traverse($env)
            )
          , ( .foreach.query
            , .foreach.start
            , .foreach.update
            , .foreach.extract
            | _t($env + $penv)
            )
          )
        elif .label then
          ( [{(.label.ident.str): (.label.ident + {type: "label", uri: $uri})}]
            as $labelenv
          | .label.body
          | _t($env + $labelenv)
          )
        elif .str then
          ( .str.queries[]?
          | _t($env)
          )
        elif .if then
          ( .if.cond
          , .if.then
          , (.if.elif[]? | .cond,  .then)
          , (.if.else // empty)
          | _t($env)
          )
        elif .query then
          .query | _t($env)
        else empty
        end
      );

    ( # inject #include ".jq-lsp" to allow adding additional builtins
      ( [{include_path: {str: ".jq-lsp"}}]
      | map(_import_env)
      ) as $dotjqlsp_envs
    | ( (.imports // [])
      | map(_import_env)
      ) as $imports_envs
    | ( (.func_defs // [])
      | map(_func_def_env($uri))
      ) as $funcdef_envs
    | ($start_env + $dotjqlsp_envs + $imports_envs + $funcdef_envs) as $env
    # | debug({_t: .})
    # | debug({$start_env})
    # | debug({$imports_envs})
    # | debug({$funcdef_envs})
    | if .func_defs then
        ( range(.func_defs | length) as $i
        | .func_defs[$i]
        | ( .args // []
          | map(
              {(.str): (
                . + {
                  type: (if (.str | startswith("$")) then "binding" else "function" end),
                  uri: $uri, args: []
                }
              )}
            )
          ) as $argsenv
        | .body
        # should only see previously defined functions plus arguments
        | _t($start_env + $imports_envs + $funcdef_envs[0:$i+1] + $argsenv)
        )
      else empty
      end
    , if .op then
        ( .left
        , .right
        | _t($env)
        )
      elif .term then
        ( .term
        | _term_traverse($env)
        )
      else empty
      end
    , if f then {q: ., env: ($env)}
      else empty
      end
    );
  _t($start_env);


def handle($state):
  def _readfile_uri($state; $uri):
    ( $state.files[$uri]
    | if (. | not) then
        ( $uri
        | file_uri_to_local
        | try
            ( readfile
            | {text: ., query: query_fromstring}
            )
          catch null
        )
      end
    );

  ( . as {$id, $method, $params}
  | ( def null_result:
        error({response: [{id: $id, result: null}]});
      def result(f):
        { response: [{
            id: $id,
            result: f
          }]
        };
      def qe_from_params(f):
        ( $params.textDocument.uri as $uri
        | $params.position as $pos
        | _readfile_uri($state; $uri) as $def_file
        | if ($def_file | not) then null_result end
        | ($def_file.text | lc_to_byte_pos($pos.line; $pos.character)) as $file_pos
        | ( $def_file.query
          | first(query_walk(
              $uri;
              builtin_env;
              ( query_token as $t
              | $t != null and
                ($uri == $t.uri or $t.uri == null) and
                ($file_pos | in_range($t.start; $t.stop))
              ) and f
            )) // null_result
          )
        );
      def def_from_env(f):
        ( ( first(
              ( env_iter_entries
              | select(.value | f)
              )
            ) // null_result
          )
        | .value
        );

      if $method == "initialize" then
        { response: [{
            id: $id,
            result: {
              capabilities: {
                textDocumentSync: TextDocumentSyncFull,
                definitionProvider: true,
                completionProvider: {
                  #resolveProvider: true
                  # completionItem: {
                  #   labelDetailsSupport: true
                  # }
                },
                hoverProvider: true,
                publishDiagnostics: {},
                documentSymbolProvider: true,
                workspace: {
                  workspaceFolders: {
                    supported: true,
                    # changeNotifications: true
                  }
                },
                semanticTokensProvider: {
                  legend: {
                    tokenTypes: [
                      SemanticTokenTypeNamespace,
                      SemanticTokenTypeType,
                      SemanticTokenTypeClass,
                      SemanticTokenTypeEnum,
                      SemanticTokenTypeInterface,
                      SemanticTokenTypeStruct,
                      SemanticTokenTypeTypeParameter,
                      SemanticTokenTypeParameter,
                      SemanticTokenTypeVariable,
                      SemanticTokenTypeProperty,
                      SemanticTokenTypeEnumMember,
                      SemanticTokenTypeEvent,
                      SemanticTokenTypeFunction,
                      SemanticTokenTypeMethod,
                      SemanticTokenTypeMacro,
                      SemanticTokenTypeKeyword,
                      SemanticTokenTypeModifier,
                      SemanticTokenTypeComment,
                      SemanticTokenTypeString,
                      SemanticTokenTypeNumber,
                      SemanticTokenTypeRegexp,
                      SemanticTokenTypeOperator,
                      SemanticTokenTypeDecorator
                    ],
                    tokenModifiers: [
                      SemanticTokenModifierDeclaration,
                      SemanticTokenModifierDefinition,
                      SemanticTokenModifierReadonly,
                      SemanticTokenModifierStatic,
                      SemanticTokenModifierDeprecated,
                      SemanticTokenModifierAbstract,
                      SemanticTokenModifierAsync,
                      SemanticTokenModifierModification,
                      SemanticTokenModifierDocumentation,
                      SemanticTokenModifierDefaultLibrary
                    ]
                  },
                  full: true
                }
              }
            }
          }
        ]}
      elif $method == "initialized" then
        { response: [{
            method: "client/registerCapability",
            params: {
              registrations:
                [
                  {
                    "id": "79eee87c-c409-4664-8102-e03263673f6f",
                    "method": "workspace/didChangeConfiguration",
                    "registerOptions": {
                      "documentSelector": [
                        { "language": "jq" }
                      ]
                    }
                  }
                ]
            }
          },
          { method: "workspace/configuration",
            params: {
                items: [
                  {scopeURI: "resource", section: "jqlsp"}
                ]
            }
          }
        ]}
      # TODO: exit
      elif $method == "shutdown" then null_result
      elif (
          $method == "textDocument/didOpen" or
          $method == "textDocument/didChange"
        ) then
        # textDocument/didOpen:
        # "params": {
        #   "textDocument": {
        #     "languageId": "plaintext",
        #     "text": "abc",
        #     "uri": "file:///a",
        #     "version": 1
        #   }
        # }
        # textDocument/didChange:
        # when textDocumentSync full i guess always get full text?
        # "params": {
        #   "contentChanges": [{"text": "abc"}],
        #   "textDocument": {
        #     "uri": "file:///a",
        #     "version": 2
        #   }
        # }
        ( $params.textDocument as $doc
        | ($doc.text // $params.contentChanges[0].text // "") as $text
        | try
            ( ($text | query_fromstring) as $text_query
            | { state:
                  ( $state
                  | .files[$doc.uri] = {text: $text, query: $text_query}
                  ),
                response:
                  [ { method: "textDocument/publishDiagnostics",
                      params: {
                        uri: $doc.uri,
                        diagnostics:
                          [ $text_query
                          | query_walk($doc.uri; builtin_env; .term.func or .term.format) as {$env, $q}
                          | ($q | query_token) as $token
                          | ($q | query_args) as $args
                          | if isempty(
                                ( $env
                                | env_iter_entries
                                | .value
                                | select(
                                    # TODO refactor share with definition
                                    .str == $token.str and
                                    ( ( .args == null and $args == null) or
                                      ( .args != null and $args != null and
                                        (.args | length) == ($args | length)
                                      )
                                    )
                                  )
                                )
                              )
                            then
                              { range: {
                                  start: ($text | byte_pos_to_lc($token.start)),
                                  end: ($text | byte_pos_to_lc($token.stop))
                                },
                                message: "\($q | func_term_name) not found"
                              }
                            else empty
                            end
                          ]
                      }
                  } ]
              }
            )
          catch
            ( . as $err
            | { response: [{
                  method: "textDocument/publishDiagnostics",
                  params:
                    { uri: $doc.uri,
                      diagnostics:
                        [ { range: {
                              start: ($text | byte_pos_to_lc($err.offset)),
                              end: ($text | byte_pos_to_lc($err.offset))
                            },
                            message: $err.error
                          }
                        ]
                    }
                }]
              }
            )
        )
      elif $method == "textDocument/documentSymbol" then
        ( $params.textDocument.uri as $uri
        | _readfile_uri($state; $uri) as $file
        | if $file == null then null_result end
        | result(
            # TODO: just traverse, no env
            [ $file.query
            | query_walk($uri; []; .func_defs)
            | .q.func_defs[] as $f
            | { name : ($f | func_def_signature),
                kind: SymbolKindFunction,
                location:
                  { uri: $uri,
                    range:
                      { start: ($file.text | byte_pos_to_lc($f.name.start)),
                        end: ($file.text | byte_pos_to_lc($f.name.stop))
                      }
                  },
              }
            ]
          )
        )
      elif $method == "textDocument/didSave" then
        # textDocument/didSave:
        # "params": {
        #   "textDocument": {
        #     "uri": "file:///a"
        #   }
        # }
        {state: ($state | del($state.files[$params.textDocument.uri]))}
      elif $method == "textDocument/completion" then
        # "method": "textDocument/completion",
        # "params": {
        #   "position": {"character": 4, "line": 7},
        #   "textDocument": {
        #     "uri": "file:///a/b/c.jq",
        #   }
        # }
        ( qe_from_params(true) as {$env, $q}
        | ($q | query_token.str) as $prefix
        | result(
            [ $env
            | env_iter_entries
            | .value
            | select(.str | startswith($prefix))
            | . as $func
            | if .args and (.args | length) > 0 then
                { label: env_func_signature,
                  insertText: "\(.str)($0)",
                  insertTextFormat: TextFormatSnippet,
                }
              else
                {label: .str}
              end
            | if $func.str | startswith("$") then
                .kind = CompletionItemKindVariable
              else
                ( .kind = CompletionItemKindFunction
                | .documentation =
                    { value: ($func | env_func_markdown),
                      kind: MarkupKindsMarkdown
                    }
                )
              end
            ]
          )
        )
      elif $method == "textDocument/definition" then
        # "params": {
        #   "context": {
        #     "triggerKind": 1
        #   },
        #   "position": {
        #     "character": 34,
        #     "line": 1
        #   },
        #   "textDocument": {
        #     "uri": "file:///Users/wader/src/test/lsp-test/test.jq"
        #   }
        # }
        ( qe_from_params(true) as {$env, $q}
        | ($q | query_token.str) as $name
        | ($q | query_args) as $args
        | $env
        | def_from_env(
            .str == $name and
            ( ( .args == null and $args == null) or
              ( .args != null and $args != null and
                (.args | length) == ($args | length)
              )
            ) and
            .start and
            .stop and
            .uri
          ) as $def
        | _readfile_uri($state; $def.uri) as $def_file
        | if ($def_file | not) then null_result end
        | result({
            uri: $def.uri,
            range: {
              start: ($def_file.text | byte_pos_to_lc($def.start)),
              end: ($def_file.text | byte_pos_to_lc($def.stop))
            }
          })
        )
      elif $method == "textDocument/hover" then
        # "params": {
        #   "position": {
        #     "character": 55,
        #     "line":44
        #   },
        #   "textDocument": {
        #     "uri": "file:///a"
        #   }
        # }
        ( qe_from_params(.term.func or .term.format) as {$env, $q}
        | $env
        | def_from_env(
            .type != "binding" and
            ( ( $q.term.func and
                .str == $q.term.func.name.str and
                (.args | length) == (($q.term.func.args // []) | length)
              ) or
              ( $q.term.format and
                .str == $q.term.format.str
              )
            )
          )
        | docs[env_func_name] as $doc
        | result({
          contents: env_func_markdown
          })
        )
      elif $method == "textDocument/semanticTokens/full" then
         # "params": {"textDocument":{"uri":"file:///a"}}
        ( $params.textDocument.uri as $uri
        | _readfile_uri($state; $uri) as $file
        | if $file == null then null_result end
        | result(
            { data:
                [ $file.query
                | query_walk($uri; []; true)
                | .q
                | select(.term.func)
                | debug
                | .term.func as $t
                # line, startChar, length, tokenType, tokenModifiers
                | (( $file.text | byte_pos_to_lc($t.name.start)) | .line, .character)
                  , ($t.name.stop - $t.name.start)
                  , SemanticTokenTypeFunctionNr
                  , 0
                # | .q.func_defs[] as $f
                # | { name : ($f | func_def_signature),
                #     kind: SymbolKindFunction,
                #     location:
                #       { uri: $uri,
                #         range:
                #           { start: ($file.text | byte_pos_to_lc($f.name.start)),
                #             end: ($file.text | byte_pos_to_lc($f.name.stop))
                #           }
                #       },
                #   }
                ]
            }
            | debug
            # TODO: just traverse, no env
          )
        )
      else
        null
      end
    )
  );

def serve:
  ( . as $state
  | jsonrpc_read as $request
  | debug({$request})
  | $request
  | try handle($state)
    catch
      if (type != "object" or .response or .state | not) then error end
  | ( .response[]?
    #| debug({response: .})
    | jsonrpc_write
    )
  , .state // $state
  #| debug({state: .})
  );

# # TODO: not used atm, see comment in lsp.go
# def main:
#   ( {}
#   | loop(serve)
#   );
