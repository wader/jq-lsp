package gojqparser

import (
	"encoding/json"
	"unicode/utf8"
)

type lexer struct {
	source    string
	offset    int
	result    *Query
	token     string
	tokenType int
	inString  bool
	err       error
}

type Token struct {
	Str   string `json:"str"`
	Start int    `json:"start"`
	Stop  int    `json:"stop"`
}

func newLexer(src string) *lexer {
	return &lexer{source: src}
}

const eof = -1

var keywords = map[string]int{
	"or":      tokOrOp,
	"and":     tokAndOp,
	"module":  tokModule,
	"import":  tokImport,
	"include": tokInclude,
	"def":     tokDef,
	"as":      tokAs,
	"label":   tokLabel,
	"break":   tokBreak,
	// jq-lsp: don't treat as keyword, jq, jaq etc treat as idents
	// "null":    tokNull,
	// "true":    tokTrue,
	// "false":   tokFalse,
	"if":      tokIf,
	"then":    tokThen,
	"elif":    tokElif,
	"else":    tokElse,
	"end":     tokEnd,
	"try":     tokTry,
	"catch":   tokCatch,
	"reduce":  tokReduce,
	"foreach": tokForeach,
}

func (l *lexer) Lex(lval *yySymType) (tokenType int) {
	lval.token = &Token{}
	defer func() {
		l.tokenType = tokenType
		lval.token.Stop = l.offset
	}()
	if len(l.source) == l.offset {
		l.token = ""
		return eof
	}
	if l.inString {
		tok, str := l.scanString(l.offset)
		lval.token.Str = str
		return tok
	}
	ch, iseof := l.next()
	lval.token.Start = l.offset - 1
	if iseof {
		l.token = ""
		return eof
	}
	switch {
	case isIdent(ch, false):
		i := l.offset - 1
		j, isModule := l.scanIdentOrModule()
		l.token = l.source[i:j]
		lval.token.Str = l.token
		if isModule {
			return tokModuleIdent
		}
		if tok, ok := keywords[l.token]; ok {
			return tok
		}
		return tokIdent
	case isNumber(ch):
		if ch == '0' {
			base := l.peek()
			switch base {
			case 'b', 'o', 'x':
				i := l.offset - 1
				l.offset++
				for isInteger(base, l.peek()) {
					l.offset++
				}
				l.token = string(l.source[i:l.offset])
				lval.token.Str = l.token
				return tokNumber
			}
		}

		i := l.offset - 1
		j := l.scanNumber(numberStateLead)
		if j < 0 {
			l.token = l.source[i:-j]
			return tokInvalid
		}
		l.token = l.source[i:j]
		lval.token.Str = l.token
		return tokNumber
	}
	switch ch {
	case '.':
		ch := l.peek()
		switch {
		case ch == '.':
			l.offset++
			l.token = ".."
			return tokRecurse
		case isIdent(ch, false):
			l.token = l.source[l.offset-1 : l.scanIdent()]
			lval.token.Str = l.token[1:]
			return tokIndex
		case isNumber(ch):
			i := l.offset - 1
			j := l.scanNumber(numberStateFloat)
			if j < 0 {
				l.token = l.source[i:-j]
				return tokInvalid
			}
			l.token = l.source[i:j]
			lval.token.Str = l.token
			return tokNumber
		default:
			return '.'
		}
	case '$':
		if isIdent(l.peek(), false) {
			i := l.offset - 1
			j, isModule := l.scanIdentOrModule()
			l.token = l.source[i:j]
			lval.token.Str = l.token
			if isModule {
				return tokModuleVariable
			}
			return tokVariable
		}
	case '|':
		if l.peek() == '=' {
			l.offset++
			l.token = "|="
			lval.operator = OpModify
			return tokUpdateOp
		}
	case '?':
		if l.peek() == '/' {
			l.offset++
			if l.peek() == '/' {
				l.offset++
				l.token = "?//"
				return tokDestAltOp
			}
			l.offset--
		}
	case '+':
		if l.peek() == '=' {
			l.offset++
			l.token = "+="
			lval.operator = OpUpdateAdd
			return tokUpdateOp
		}
	case '-':
		if l.peek() == '=' {
			l.offset++
			l.token = "-="
			lval.operator = OpUpdateSub
			return tokUpdateOp
		}
	case '*':
		if l.peek() == '=' {
			l.offset++
			l.token = "*="
			lval.operator = OpUpdateMul
			return tokUpdateOp
		}
	case '/':
		switch l.peek() {
		case '=':
			l.offset++
			l.token = "/="
			lval.operator = OpUpdateDiv
			return tokUpdateOp
		case '/':
			l.offset++
			if l.peek() == '=' {
				l.offset++
				l.token = "//="
				lval.operator = OpUpdateAlt
				return tokUpdateOp
			}
			l.token = "//"
			lval.operator = OpAlt
			return tokAltOp
		}
	case '%':
		if l.peek() == '=' {
			l.offset++
			l.token = "%="
			lval.operator = OpUpdateMod
			return tokUpdateOp
		}
	case '=':
		if l.peek() == '=' {
			l.offset++
			l.token = "=="
			lval.operator = OpEq
			return tokCompareOp
		}
		l.token = "="
		lval.operator = OpAssign
		return tokUpdateOp
	case '!':
		if l.peek() == '=' {
			l.offset++
			l.token = "!="
			lval.operator = OpNe
			return tokCompareOp
		}
	case '>':
		if l.peek() == '=' {
			l.offset++
			l.token = ">="
			lval.operator = OpGe
			return tokCompareOp
		}
		l.token = ">"
		lval.operator = OpGt
		return tokCompareOp
	case '<':
		if l.peek() == '=' {
			l.offset++
			l.token = "<="
			lval.operator = OpLe
			return tokCompareOp
		}
		l.token = "<"
		lval.operator = OpLt
		return tokCompareOp
	case '@':
		if isIdent(l.peek(), true) {
			l.token = l.source[l.offset-1 : l.scanIdent()]
			lval.token.Str = l.token
			return tokFormat
		}
	case '"':
		tok, str := l.scanString(l.offset - 1)
		lval.token.Str = str
		return tok
	case '`':
		tok, str := l.scanRawString(l.offset - 1)
		lval.token.Str = str
		return tok
	default:
		if ch >= utf8.RuneSelf {
			r, size := utf8.DecodeRuneInString(l.source[l.offset-1:])
			// -1 to adjust for first byte consumed by next()
			l.offset += size - 1
			l.token = string(r)
		}
	}
	return int(ch)
}

func (l *lexer) next() (byte, bool) {
	for {
		ch := l.source[l.offset]
		l.offset++
		if ch == '#' {
			if l.skipComment() {
				return 0, true
			}
		} else if !isWhite(ch) {
			return ch, false
		} else if len(l.source) == l.offset {
			return 0, true
		}
	}
}

func (l *lexer) skipComment() bool {
	for {
		switch l.peek() {
		case 0:
			return true
		case '\\':
			switch l.offset++; l.peek() {
			case '\\', '\n':
				l.offset++
			case '\r':
				if l.offset++; l.peek() == '\n' {
					l.offset++
				}
			}
		case '\n', '\r':
			return false
		default:
			l.offset++
		}
	}
}

func (l *lexer) peek() byte {
	if len(l.source) == l.offset {
		return 0
	}
	return l.source[l.offset]
}

func (l *lexer) scanIdent() int {
	for isIdent(l.peek(), true) {
		l.offset++
	}
	return l.offset
}

func (l *lexer) scanIdentOrModule() (int, bool) {
	index := l.scanIdent()
	var isModule bool
	if l.peek() == ':' {
		l.offset++
		if l.peek() == ':' {
			l.offset++
			if isIdent(l.peek(), false) {
				l.offset++
				index = l.scanIdent()
				isModule = true
			} else {
				l.offset -= 2
			}
		} else {
			l.offset--
		}
	}
	return index, isModule
}

const (
	numberStateLead = iota
	numberStateFloat
	numberStateExpLead
	numberStateExp
)

func (l *lexer) scanNumber(state int) int {
	for {
		switch state {
		case numberStateLead, numberStateFloat:
			if ch := l.peek(); isNumber(ch) {
				l.offset++
			} else {
				switch ch {
				case '.':
					if state != numberStateLead {
						l.offset++
						return -l.offset
					}
					l.offset++
					state = numberStateFloat
				case 'e', 'E':
					l.offset++
					switch l.peek() {
					case '-', '+':
						l.offset++
					}
					state = numberStateExpLead
				default:
					if isIdent(ch, false) {
						l.offset++
						return -l.offset
					}
					return l.offset
				}
			}
		case numberStateExpLead, numberStateExp:
			if ch := l.peek(); !isNumber(ch) {
				if isIdent(ch, false) {
					l.offset++
					return -l.offset
				}
				if state == numberStateExpLead {
					return -l.offset
				}
				return l.offset
			}
			l.offset++
			state = numberStateExp
		default:
			panic(state)
		}
	}
}

func (l *lexer) scanString(start int) (int, string) {
	var decode bool
	var controls int
	unquote := func(src string, quote bool) (string, error) {
		if !decode {
			if quote {
				return src, nil
			}
			return src[1 : len(src)-1], nil
		}
		var buf []byte
		if !quote && controls == 0 {
			buf = []byte(src)
		} else {
			buf = quoteAndEscape(src, quote, controls)
		}
		if err := json.Unmarshal(buf, &src); err != nil {
			return "", err
		}
		return src, nil
	}
	for i := l.offset; i < len(l.source); i++ {
		ch := l.source[i]
		switch ch {
		case '\\':
			if i++; i >= len(l.source) {
				break
			}
			switch l.source[i] {
			case 'u':
				for j := 1; j <= 4; j++ {
					if i+j >= len(l.source) || !isHex(l.source[i+j]) {
						l.offset = i + j
						l.token = l.source[i-1 : l.offset]
						return tokInvalidEscapeSequence, ""
					}
				}
				i += 4
				fallthrough
			case '"', '/', '\\', 'b', 'f', 'n', 'r', 't':
				decode = true
			case '(':
				if !l.inString {
					l.inString = true
					return tokStringStart, ""
				}
				if i == l.offset+1 {
					l.offset += 2
					l.inString = false
					return tokStringQuery, ""
				}
				l.offset = i - 1
				l.token = l.source[start:l.offset]
				str, err := unquote(l.token, true)
				if err != nil {
					return tokInvalid, ""
				}
				return tokString, str
			default:
				l.offset = i + 1
				l.token = l.source[l.offset-2 : l.offset]
				return tokInvalidEscapeSequence, ""
			}
		case '"':
			if !l.inString {
				l.offset = i + 1
				l.token = l.source[start:l.offset]
				str, err := unquote(l.token, false)
				if err != nil {
					return tokInvalid, ""
				}
				return tokString, str
			}
			if i > l.offset {
				l.offset = i
				l.token = l.source[start:l.offset]
				str, err := unquote(l.token, true)
				if err != nil {
					return tokInvalid, ""
				}
				return tokString, str
			}
			l.inString = false
			l.offset = i + 1
			return tokStringEnd, ""
		default:
			if !decode {
				decode = ch > '~'
			}
			if ch < ' ' { // ref: unquoteBytes in encoding/json
				controls++
			}
		}
	}
	l.offset = len(l.source)
	l.token = ""
	return tokUnterminatedString, ""
}

func (l *lexer) scanRawString(start int) (int, string) {
	for i := l.offset; i < len(l.source); i++ {
		ch := l.source[i]
		switch ch {
		case '`':
			l.offset = i + 1
			l.token = l.source[start:l.offset]
			return tokString, l.token[1 : len(l.token)-1]
		}
	}
	l.offset = len(l.source)
	l.token = ""
	return tokUnterminatedString, ""
}

func quoteAndEscape(src string, quote bool, controls int) []byte {
	size := len(src) + controls*5
	if quote {
		size += 2
	}
	buf := make([]byte, size)
	var j int
	if quote {
		buf[0] = '"'
		buf[len(buf)-1] = '"'
		j++
	}
	for i := 0; i < len(src); i++ {
		if ch := src[i]; ch < ' ' {
			const hex = "0123456789abcdef"
			copy(buf[j:], `\u00`)
			buf[j+4] = hex[ch>>4]
			buf[j+5] = hex[ch&0xF]
			j += 6
		} else {
			buf[j] = ch
			j++
		}
	}
	return buf
}

// ParseError represents a description of a query parsing error.
type ParseError struct {
	Offset    int    // the error occurred after reading Offset bytes
	Token     string // the Token that caused the error (may be empty)
	tokenType int
}

func (err *ParseError) Error() string {
	switch err.tokenType {
	case eof:
		return "unexpected EOF"
	case tokInvalid:
		return "invalid token " + jsonMarshal(err.Token)
	case tokInvalidEscapeSequence:
		return `invalid escape sequence "` + err.Token + `" in string literal`
	case tokUnterminatedString:
		return "unterminated string literal"
	default:
		return "unexpected token " + jsonMarshal(err.Token)
	}
}

func (l *lexer) Error(string) {
	offset, token := l.offset, l.token
	if l.tokenType != eof && l.tokenType < utf8.RuneSelf {
		token = string(rune(l.tokenType))
	}
	l.err = &ParseError{offset, token, l.tokenType}
}

func isWhite(ch byte) bool {
	switch ch {
	case '\t', '\n', '\r', ' ':
		return true
	default:
		return false
	}
}

func isIdent(ch byte, tail bool) bool {
	return 'a' <= ch && ch <= 'z' ||
		'A' <= ch && ch <= 'Z' || ch == '_' ||
		tail && isNumber(ch)
}

func isHex(ch byte) bool {
	return 'a' <= ch && ch <= 'f' ||
		'A' <= ch && ch <= 'F' ||
		isNumber(ch)
}

func isBinary(ch byte) bool {
	return '0' == ch || ch == '1'
}

func isOctal(ch byte) bool {
	return '0' <= ch && ch <= '7'
}

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isInteger(base byte, ch byte) bool {
	if ch == '_' {
		return true
	}
	switch base {
	case 'b':
		return isBinary(ch)
	case 'o':
		return isOctal(ch)
	case 'x':
		return isHex(ch)
	default:
		panic("unreachable")
	}
}
