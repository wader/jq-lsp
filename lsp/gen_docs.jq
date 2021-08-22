#!/usr/bin/env fq -f
# ./lsp/gen_docs.jq <(gojq --yaml-input . ../jq/docs/content/manual/manual.yml)


# TODO: skip ^$
def map_titles:
  {
    "Convert to/from JSON": "`fromjson`, `tojson`",
    "`$ENV`, `env`": "`env`",
    # missing `
    "`sort, sort_by(path_expression)`": "`sort`, `sort_by(path_expression)`",
    "Dates": "`now`, `fromdateiso8601`, `todateiso8601`, `todate`, `fromdate`, `strptime(fmt)`, `strftime(fmt)`, `strflocaltime`, `mktime`, `gmtime`, `localtime`",
    "SQL-Style Operators": "`INDEX(stream; index_expression)`, `JOIN($idx; stream; idx_expr; join_expr)`, `JOIN($idx; stream; idx_expr)`, `JOIN($idx; idx_expr)`, `IN(s)`, `IN(source; s)`",
    # missing ,
    "`sub(regex; tostring)` `sub(regex; string; flags)`": "`sub(regex; tostring)`, `sub(regex; string; flags)`",
    # missing ,
    "`range(upto)`, `range(from;upto)` `range(from;upto;by)`": "`range(upto)`, `range(from;upto)`, `range(from;upto;by)`",
    "Math": "`acos`, `acosh`, `asin`, `asinh`, `atan`, `atanh`, `cbrt`, `ceil`, `cos`, `cosh`, `erf`, `erfc`, `exp`, `exp10`, `exp2`, `expm1`, `fabs`, `floor`, `gamma`, `j0`, `j1`, `lgamma`, `log`, `log10`, `log1p`, `log2`, `logb`, `nearbyint`, `pow10`, `rint`, `round`, `significand`, `sin`, `sinh`, `sqrt`, `tan`, `tanh`, `tgamma`, `trunc`, `y0`, `y1`, `atan2` `copysign` `drem`, `fdim`, `fmax`, `fmin`, `fmod`, `frexp`, `hypot`, `jn`, `ldexp`, `modf`, `nextafter`, `nexttoward`, `pow`, `remainder`, `scalb`, `scalbln`, `yn`, `fma`",
    # missing `,
    "`sort, sort_by(path_expression)`": "`sort`, `sort_by(path_expression)`",
    "`error(message)`": "`error`, `error(message)`",
    "'I/O'": "`stdout`"
  };

( ("def docs:\n" | stdout)
, ( [ ..
    | select(.title)?
    | .body as $doc
    | .title
    | if map_titles[.] then map_titles[.] else . end
    | select(test("^`\\w.*`$"))?
    | split(", ")[] | .[1:-1] | "\(.)"
    | {(split("(") | [.[0], "/", (.[1] | if . then split(";") | length else 0 end)] | join("")):
        ($doc | ltrimstr("\n") | rtrimstr("\n"))
      }
    ]
    | add
  )
, (";\n" | stdout)
)
