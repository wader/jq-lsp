package gojq

import (
	"encoding/json"
	"fmt"
)

// Operator ...
type Operator int

// Operators ...
const (
	OpPipe Operator = iota + 1
	OpComma
	OpAdd
	OpSub
	OpMul
	OpDiv
	OpMod
	OpEq
	OpNe
	OpGt
	OpLt
	OpGe
	OpLe
	OpAnd
	OpOr
	OpAlt
	OpAssign
	OpModify
	OpUpdateAdd
	OpUpdateSub
	OpUpdateMul
	OpUpdateDiv
	OpUpdateMod
	OpUpdateAlt
)

// OperatorFromString convert string to Operator or zero if not possible
func OperatorFromString(s string) Operator {
	switch s {
	case "|":
		return OpPipe
	case ",":
		return OpComma
	case "+":
		return OpAdd
	case "-":
		return OpSub
	case "*":
		return OpMul
	case "/":
		return OpDiv
	case "%":
		return OpMod
	case "==":
		return OpEq
	case "!=":
		return OpNe
	case ">":
		return OpGt
	case "<":
		return OpLt
	case ">=":
		return OpGe
	case "<=":
		return OpLe
	case "and":
		return OpAnd
	case "or":
		return OpOr
	case "//":
		return OpAlt
	case "=":
		return OpAssign
	case "|=":
		return OpModify
	case "+=":
		return OpUpdateAdd
	case "-=":
		return OpUpdateSub
	case "*=":
		return OpUpdateMul
	case "/=":
		return OpUpdateDiv
	case "%=":
		return OpUpdateMod
	case "//=":
		return OpUpdateAlt
	default:
		return 0
	}
}

// String implements Stringer.
func (op Operator) String() string {
	switch op {
	case OpPipe:
		return "|"
	case OpComma:
		return ","
	case OpAdd:
		return "+"
	case OpSub:
		return "-"
	case OpMul:
		return "*"
	case OpDiv:
		return "/"
	case OpMod:
		return "%"
	case OpEq:
		return "=="
	case OpNe:
		return "!="
	case OpGt:
		return ">"
	case OpLt:
		return "<"
	case OpGe:
		return ">="
	case OpLe:
		return "<="
	case OpAnd:
		return "and"
	case OpOr:
		return "or"
	case OpAlt:
		return "//"
	case OpAssign:
		return "="
	case OpModify:
		return "|="
	case OpUpdateAdd:
		return "+="
	case OpUpdateSub:
		return "-="
	case OpUpdateMul:
		return "*="
	case OpUpdateDiv:
		return "/="
	case OpUpdateMod:
		return "%="
	case OpUpdateAlt:
		return "//="
	default:
		return ""
	}
}

func (op Operator) MarshalJSON() ([]byte, error) {
	if op == 0 {
		return json.Marshal(nil)
	}
	return json.Marshal(op.String())
}

func (op *Operator) UnmarshalJSON(text []byte) error {
	var s string
	err := json.Unmarshal(text, &s)
	if s == "" || err != nil {
		*op = 0
		return nil
	}
	*op = OperatorFromString(s)
	if *op == 0 {
		return fmt.Errorf("unknown operator %v", s)
	}
	return nil
}
