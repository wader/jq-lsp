#!/usr/bin/env fq -d bytes -nf
# ./lsp/gen_builtin_env.jq ../gojq/builtin.jq <(jq -rn 'builtins | tojson') <(gojq --yaml-input . ../jq/docs/content/manual/manual.yml)

def gen_args:
  [range(.)] | map(.+97) | map([.] | implode);

def _builtins_json:
  [ fromjson[]
  | split("/") as [$name,$args]
  | {"\($name)/\($args)": {
          str: $name,
          args: ($args | tonumber | gen_args),
      }
    }
  ] | add;

def _builtins_jq:
  [ _query_fromstring
  | .func_defs
  | map(select(.name | startswith("_") | not))
  | sort_by(.name)
  | .[]
  | (.args // []) as $args
  | {("\(.name)/\(.args | length)"): {str: .name, args: $args}}
  ] | add;

def _builtins_doc:
  ( ..
  | select(.title | test("^`\\w+`"))?
  | .body as $doc
  | .title
  | split(", ")[]
  | .[1:-1]
  | {(.): ($doc | ltrimstr("\n") | rtrimstr("\n"))}
  );

( ("def builtin_env:\n[{\n" | print)
, ( ( ( input | tobytes | tostring | _builtins_json)
    + ( input | tobytes | tostring | _builtins_jq)
    | to_entries
    | sort_by(.key)
    | .[]
    | [(.key | tojson), ": ", (.value | tojson), ","]
    | join("")
    | println
    )
  )
, ("}];\n" | print)
)