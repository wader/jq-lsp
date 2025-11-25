package gojqparser

import (
	"strings"
)

// Parse a query string, and returns the query struct.
//
// If parsing failed, it returns an error of type [*ParseError], which has
// the byte offset and the invalid token. The byte offset is the scanned bytes
// when the error occurred. The token is empty if the error occurred after
// scanning the entire query string.
func Parse(src string) (*Query, error) {
	l := newLexer(src)
	if yyParse(l) > 0 {
		return nil, l.err
	}
	return l.result, nil
}

// Query represents the abstract syntax tree of a jq query.
type Query struct {
	Meta     *ConstObject `json:"meta,omitempty"`
	Imports  []*Import    `json:"imports,omitempty"`
	FuncDefs []*FuncDef   `json:"func_defs,omitempty"`
	Term     *Term        `json:"term,omitempty"`
	Left     *Query       `json:"left,omitempty"`
	Right    *Query       `json:"right,omitempty"`
	Patterns []*Pattern   `json:"patterns,omitempty"`
	Op       Operator     `json:"op,omitempty"`
}

func (e *Query) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *Query) writeTo(s *strings.Builder) {
	if e.Meta != nil {
		s.WriteString("module ")
		e.Meta.writeTo(s)
		s.WriteString(";\n")
	}
	for _, im := range e.Imports {
		im.writeTo(s)
	}
	for _, fd := range e.FuncDefs {
		fd.writeTo(s)
		s.WriteByte(' ')
	}
	if e.Term != nil {
		e.Term.writeTo(s)
	} else if e.Right != nil {
		e.Left.writeTo(s)
		if e.Op != OpComma {
			s.WriteByte(' ')
		}
		for i, p := range e.Patterns {
			if i == 0 {
				s.WriteString("as ")
			} else {
				s.WriteString("?// ")
			}
			p.writeTo(s)
			s.WriteByte(' ')
		}
		s.WriteString(e.Op.String())
		s.WriteByte(' ')
		e.Right.writeTo(s)
	}
}

// Import ...
type Import struct {
	ImportPath  *Token       `json:"import_path,omitempty"`
	ImportAlias *Token       `json:"import_alias,omitempty"`
	IncludePath *Token       `json:"include_path,omitempty"`
	Meta        *ConstObject `json:"meta,omitempty"`
}

func (e *Import) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *Import) writeTo(s *strings.Builder) {
	if e.ImportPath.Str != "" {
		s.WriteString("import ")
		jsonEncodeString(s, e.ImportPath.Str)
		s.WriteString(" as ")
		s.WriteString(e.ImportAlias.Str)
	} else {
		s.WriteString("include ")
		jsonEncodeString(s, e.IncludePath.Str)
	}
	if e.Meta != nil {
		s.WriteByte(' ')
		e.Meta.writeTo(s)
	}
	s.WriteString(";\n")
}

// FuncDef ...
type FuncDef struct {
	Name *Token   `json:"name,omitempty"`
	Args []*Token `json:"args,omitempty"`
	Body *Query   `json:"body,omitempty"`
}

func (e *FuncDef) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *FuncDef) writeTo(s *strings.Builder) {
	s.WriteString("def ")
	s.WriteString(e.Name.Str)
	if len(e.Args) > 0 {
		s.WriteByte('(')
		for i, e := range e.Args {
			if i > 0 {
				s.WriteString("; ")
			}
			s.WriteString(e.Str)
		}
		s.WriteByte(')')
	}
	s.WriteString(": ")
	e.Body.writeTo(s)
	s.WriteByte(';')
}

// Term ...
type Term struct {
	Type       TermType  `json:"type,omitempty"`
	Index      *Index    `json:"index,omitempty"`
	Func       *Func     `json:"func,omitempty"`
	Object     *Object   `json:"object,omitempty"`
	Array      *Array    `json:"array,omitempty"`
	Number     *Token    `json:"number,omitempty"`
	Unary      *Unary    `json:"unary,omitempty"`
	Format     *Token    `json:"format,omitempty"`
	Str        *String   `json:"str,omitempty"`
	If         *If       `json:"if,omitempty"`
	Try        *Try      `json:"try,omitempty"`
	Reduce     *Reduce   `json:"reduce,omitempty"`
	Foreach    *Foreach  `json:"foreach,omitempty"`
	Label      *Label    `json:"label,omitempty"`
	Break      *Token    `json:"break,omitempty"`
	Query      *Query    `json:"query,omitempty"`
	SuffixList []*Suffix `json:"suffix_list,omitempty"`
}

func (e *Term) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *Term) writeTo(s *strings.Builder) {
	switch e.Type {
	case TermTypeIdentity:
		s.WriteByte('.')
	case TermTypeRecurse:
		s.WriteString("..")
	case TermTypeNull:
		s.WriteString("null")
	case TermTypeTrue:
		s.WriteString("true")
	case TermTypeFalse:
		s.WriteString("false")
	case TermTypeIndex:
		e.Index.writeTo(s)
	case TermTypeFunc:
		e.Func.writeTo(s)
	case TermTypeObject:
		e.Object.writeTo(s)
	case TermTypeArray:
		e.Array.writeTo(s)
	case TermTypeNumber:
		s.WriteString(e.Number.Str)
	case TermTypeUnary:
		e.Unary.writeTo(s)
	case TermTypeFormat:
		s.WriteString(e.Format.Str)
		if e.Str != nil {
			s.WriteByte(' ')
			e.Str.writeTo(s)
		}
	case TermTypeString:
		e.Str.writeTo(s)
	case TermTypeIf:
		e.If.writeTo(s)
	case TermTypeTry:
		e.Try.writeTo(s)
	case TermTypeReduce:
		e.Reduce.writeTo(s)
	case TermTypeForeach:
		e.Foreach.writeTo(s)
	case TermTypeLabel:
		e.Label.writeTo(s)
	case TermTypeBreak:
		s.WriteString("break ")
		s.WriteString(e.Break.Str)
	case TermTypeQuery:
		s.WriteByte('(')
		e.Query.writeTo(s)
		s.WriteByte(')')
	}
	for _, e := range e.SuffixList {
		e.writeTo(s)
	}
}

// Unary ...
type Unary struct {
	Op   Operator `json:"op,omitempty"`
	Term *Term    `json:"term,omitempty"`
}

func (e *Unary) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *Unary) writeTo(s *strings.Builder) {
	s.WriteString(e.Op.String())
	e.Term.writeTo(s)
}

// Pattern ...
type Pattern struct {
	Name   *Token           `json:"name,omitempty"`
	Array  []*Pattern       `json:"array,omitempty"`
	Object []*PatternObject `json:"object,omitempty"`
}

func (e *Pattern) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *Pattern) writeTo(s *strings.Builder) {
	if e.Name.Str != "" {
		s.WriteString(e.Name.Str)
	} else if len(e.Array) > 0 {
		s.WriteByte('[')
		for i, e := range e.Array {
			if i > 0 {
				s.WriteString(", ")
			}
			e.writeTo(s)
		}
		s.WriteByte(']')
	} else if len(e.Object) > 0 {
		s.WriteByte('{')
		for i, e := range e.Object {
			if i > 0 {
				s.WriteString(", ")
			}
			e.writeTo(s)
		}
		s.WriteByte('}')
	}
}

// PatternObject ...
type PatternObject struct {
	Key       *Token   `json:"key,omitempty"`
	KeyString *String  `json:"key_string,omitempty"`
	KeyQuery  *Query   `json:"key_query,omitempty"`
	Val       *Pattern `json:"val,omitempty"`
}

func (e *PatternObject) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *PatternObject) writeTo(s *strings.Builder) {
	if e.Key.Str != "" {
		s.WriteString(e.Key.Str)
	} else if e.KeyString != nil {
		e.KeyString.writeTo(s)
	} else if e.KeyQuery != nil {
		s.WriteByte('(')
		e.KeyQuery.writeTo(s)
		s.WriteByte(')')
	}
	if e.Val != nil {
		s.WriteString(": ")
		e.Val.writeTo(s)
	}
}

// Index ...
type Index struct {
	Name    *Token  `json:"name,omitempty"`
	Str     *String `json:"str,omitempty"`
	Start   *Query  `json:"start,omitempty"`
	End     *Query  `json:"end,omitempty"`
	IsSlice bool    `json:"is_slice,omitempty"`
}

func (e *Index) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *Index) writeTo(s *strings.Builder) {
	if l := s.Len(); l > 0 {
		// ". .x" != "..x" and "0 .x" != "0.x"
		if c := s.String()[l-1]; c == '.' || '0' <= c && c <= '9' {
			s.WriteByte(' ')
		}
	}
	s.WriteByte('.')
	e.writeSuffixTo(s)
}

func (e *Index) writeSuffixTo(s *strings.Builder) {
	if e.Name.Str != "" {
		s.WriteString(e.Name.Str)
	} else if e.Str != nil {
		e.Str.writeTo(s)
	} else {
		s.WriteByte('[')
		if e.IsSlice {
			if e.Start != nil {
				e.Start.writeTo(s)
			}
			s.WriteByte(':')
			if e.End != nil {
				e.End.writeTo(s)
			}
		} else {
			e.Start.writeTo(s)
		}
		s.WriteByte(']')
	}
}

// Func ...
type Func struct {
	Name *Token   `json:"name,omitempty"`
	Args []*Query `json:"args,omitempty"`
}

func (e *Func) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *Func) writeTo(s *strings.Builder) {
	s.WriteString(e.Name.Str)
	if len(e.Args) > 0 {
		s.WriteByte('(')
		for i, e := range e.Args {
			if i > 0 {
				s.WriteString("; ")
			}
			e.writeTo(s)
		}
		s.WriteByte(')')
	}
}

// String ...
type String struct {
	Str     *Token   `json:"str,omitempty"`
	Queries []*Query `json:"queries,omitempty"`
}

func (e *String) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *String) writeTo(s *strings.Builder) {
	if e.Queries == nil {
		jsonEncodeString(s, e.Str.Str)
		return
	}
	s.WriteByte('"')
	for _, e := range e.Queries {
		if e.Term.Str == nil {
			s.WriteString(`\`)
			e.writeTo(s)
		} else {
			es := e.String()
			s.WriteString(es[1 : len(es)-1])
		}
	}
	s.WriteByte('"')
}

// Object ...
type Object struct {
	KeyVals []*ObjectKeyVal `json:"key_vals,omitempty"`
}

func (e *Object) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *Object) writeTo(s *strings.Builder) {
	if len(e.KeyVals) == 0 {
		s.WriteString("{}")
		return
	}
	s.WriteString("{ ")
	for i, kv := range e.KeyVals {
		if i > 0 {
			s.WriteString(", ")
		}
		kv.writeTo(s)
	}
	s.WriteString(" }")
}

// ObjectKeyVal ...
type ObjectKeyVal struct {
	Key       *Token  `json:"key,omitempty"`
	KeyString *String `json:"key_string,omitempty"`
	KeyQuery  *Query  `json:"key_query,omitempty"`
	Val       *Query  `json:"val,omitempty"`
}

func (e *ObjectKeyVal) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *ObjectKeyVal) writeTo(s *strings.Builder) {
	if e.Key.Str != "" {
		s.WriteString(e.Key.Str)
	} else if e.KeyString != nil {
		e.KeyString.writeTo(s)
	} else if e.KeyQuery != nil {
		s.WriteByte('(')
		e.KeyQuery.writeTo(s)
		s.WriteByte(')')
	}
	if e.Val != nil {
		s.WriteString(": ")
		e.Val.writeTo(s)
	}
}

// Array ...
type Array struct {
	Query *Query `json:"query,omitempty"`
}

func (e *Array) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *Array) writeTo(s *strings.Builder) {
	s.WriteByte('[')
	if e.Query != nil {
		e.Query.writeTo(s)
	}
	s.WriteByte(']')
}

// Suffix ...
type Suffix struct {
	Index    *Index `json:"index,omitempty"`
	Iter     bool   `json:"iter,omitempty"`
	Optional bool   `json:"optional,omitempty"`
}

func (e *Suffix) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *Suffix) writeTo(s *strings.Builder) {
	if e.Index != nil {
		if e.Index.Name.Str != "" || e.Index.Str != nil {
			e.Index.writeTo(s)
		} else {
			e.Index.writeSuffixTo(s)
		}
	} else if e.Iter {
		s.WriteString("[]")
	} else if e.Optional {
		s.WriteByte('?')
	}
}

// If ...
type If struct {
	Cond *Query    `json:"cond,omitempty"`
	Then *Query    `json:"then,omitempty"`
	Elif []*IfElif `json:"elif,omitempty"`
	Else *Query    `json:"else,omitempty"`
}

func (e *If) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *If) writeTo(s *strings.Builder) {
	s.WriteString("if ")
	e.Cond.writeTo(s)
	s.WriteString(" then ")
	e.Then.writeTo(s)
	for _, e := range e.Elif {
		s.WriteByte(' ')
		e.writeTo(s)
	}
	if e.Else != nil {
		s.WriteString(" else ")
		e.Else.writeTo(s)
	}
	s.WriteString(" end")
}

// IfElif ...
type IfElif struct {
	Cond *Query `json:"cond,omitempty"`
	Then *Query `json:"then,omitempty"`
}

func (e *IfElif) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *IfElif) writeTo(s *strings.Builder) {
	s.WriteString("elif ")
	e.Cond.writeTo(s)
	s.WriteString(" then ")
	e.Then.writeTo(s)
}

// Try ...
type Try struct {
	Body  *Query `json:"body,omitempty"`
	Catch *Query `json:"catch,omitempty"`
}

func (e *Try) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *Try) writeTo(s *strings.Builder) {
	s.WriteString("try ")
	e.Body.writeTo(s)
	if e.Catch != nil {
		s.WriteString(" catch ")
		e.Catch.writeTo(s)
	}
}

// Reduce ...
type Reduce struct {
	Query   *Query   `json:"query,omitempty"`
	Pattern *Pattern `json:"pattern,omitempty"`
	Start   *Query   `json:"start,omitempty"`
	Update  *Query   `json:"update,omitempty"`
}

func (e *Reduce) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *Reduce) writeTo(s *strings.Builder) {
	s.WriteString("reduce ")
	e.Query.writeTo(s)
	s.WriteString(" as ")
	e.Pattern.writeTo(s)
	s.WriteString(" (")
	e.Start.writeTo(s)
	s.WriteString("; ")
	e.Update.writeTo(s)
	s.WriteByte(')')
}

// Foreach ...
type Foreach struct {
	Query   *Query   `json:"query,omitempty"`
	Pattern *Pattern `json:"pattern,omitempty"`
	Start   *Query   `json:"start,omitempty"`
	Update  *Query   `json:"update,omitempty"`
	Extract *Query   `json:"extract,omitempty"`
}

func (e *Foreach) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *Foreach) writeTo(s *strings.Builder) {
	s.WriteString("foreach ")
	e.Query.writeTo(s)
	s.WriteString(" as ")
	e.Pattern.writeTo(s)
	s.WriteString(" (")
	e.Start.writeTo(s)
	s.WriteString("; ")
	e.Update.writeTo(s)
	if e.Extract != nil {
		s.WriteString("; ")
		e.Extract.writeTo(s)
	}
	s.WriteByte(')')
}

// Label ...
type Label struct {
	Ident *Token `json:"ident,omitempty"`
	Body  *Query `json:"body,omitempty"`
}

func (e *Label) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *Label) writeTo(s *strings.Builder) {
	s.WriteString("label ")
	s.WriteString(e.Ident.Str)
	s.WriteString(" | ")
	e.Body.writeTo(s)
}

// ConstTerm ...
type ConstTerm struct {
	Object *ConstObject `json:"object,omitempty"`
	Array  *ConstArray  `json:"array,omitempty"`
	Number *Token       `json:"number,omitempty"`
	Str    *Token       `json:"str,omitempty"`
	Null   bool         `json:"null,omitempty"`
	True   bool         `json:"true,omitempty"`
	False  bool         `json:"false,omitempty"`
}

func (e *ConstTerm) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *ConstTerm) writeTo(s *strings.Builder) {
	if e.Object != nil {
		e.Object.writeTo(s)
	} else if e.Array != nil {
		e.Array.writeTo(s)
	} else if e.Number.Str != "" {
		s.WriteString(e.Number.Str)
	} else if e.Null {
		s.WriteString("null")
	} else if e.True {
		s.WriteString("true")
	} else if e.False {
		s.WriteString("false")
	} else {
		jsonEncodeString(s, e.Str.Str)
	}
}

// ConstObject ...
type ConstObject struct {
	KeyVals []*ConstObjectKeyVal `json:"keyvals,omitempty"`
}

func (e *ConstObject) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *ConstObject) writeTo(s *strings.Builder) {
	if len(e.KeyVals) == 0 {
		s.WriteString("{}")
		return
	}
	s.WriteString("{ ")
	for i, kv := range e.KeyVals {
		if i > 0 {
			s.WriteString(", ")
		}
		kv.writeTo(s)
	}
	s.WriteString(" }")
}

// ConstObjectKeyVal ...
type ConstObjectKeyVal struct {
	Key       *Token     `json:"key,omitempty"`
	KeyString *Token     `json:"key_string,omitempty"`
	Val       *ConstTerm `json:"val,omitempty"`
}

func (e *ConstObjectKeyVal) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *ConstObjectKeyVal) writeTo(s *strings.Builder) {
	if e.Key.Str != "" {
		s.WriteString(e.Key.Str)
	} else {
		jsonEncodeString(s, e.KeyString.Str)
	}
	s.WriteString(": ")
	e.Val.writeTo(s)
}

// ConstArray ...
type ConstArray struct {
	Elems []*ConstTerm `json:"elems,omitempty"`
}

func (e *ConstArray) String() string {
	var s strings.Builder
	e.writeTo(&s)
	return s.String()
}

func (e *ConstArray) writeTo(s *strings.Builder) {
	s.WriteByte('[')
	for i, e := range e.Elems {
		if i > 0 {
			s.WriteString(", ")
		}
		e.writeTo(s)
	}
	s.WriteByte(']')
}
