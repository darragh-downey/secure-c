package lexer

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	source   string
	start    int
	current  int
	line     int
	column   int
	keywords map[string]TokenType
}

func NewLexer(source string) *Lexer {
	return &Lexer{
		source: source,
		line:   1,
		column: 1,
		keywords: map[string]TokenType{
			"auto":           TOKEN_KEYWORD,
			"break":          TOKEN_KEYWORD,
			"case":           TOKEN_KEYWORD,
			"char":           TOKEN_KEYWORD,
			"const":          TOKEN_KEYWORD,
			"continue":       TOKEN_KEYWORD,
			"default":        TOKEN_KEYWORD,
			"do":             TOKEN_KEYWORD,
			"double":         TOKEN_KEYWORD,
			"else":           TOKEN_KEYWORD,
			"enum":           TOKEN_KEYWORD,
			"extern":         TOKEN_KEYWORD,
			"float":          TOKEN_KEYWORD,
			"for":            TOKEN_KEYWORD,
			"goto":           TOKEN_KEYWORD,
			"if":             TOKEN_KEYWORD,
			"inline":         TOKEN_KEYWORD,
			"int":            TOKEN_KEYWORD,
			"long":           TOKEN_KEYWORD,
			"register":       TOKEN_KEYWORD,
			"restrict":       TOKEN_KEYWORD,
			"return":         TOKEN_KEYWORD,
			"short":          TOKEN_KEYWORD,
			"signed":         TOKEN_KEYWORD,
			"sizeof":         TOKEN_KEYWORD,
			"static":         TOKEN_KEYWORD,
			"struct":         TOKEN_KEYWORD,
			"switch":         TOKEN_KEYWORD,
			"typedef":        TOKEN_KEYWORD,
			"union":          TOKEN_KEYWORD,
			"unsigned":       TOKEN_KEYWORD,
			"void":           TOKEN_KEYWORD,
			"volatile":       TOKEN_KEYWORD,
			"while":          TOKEN_KEYWORD,
			"_Alignas":       TOKEN_KEYWORD,
			"_Alignof":       TOKEN_KEYWORD,
			"_Atomic":        TOKEN_KEYWORD,
			"_Bool":          TOKEN_KEYWORD,
			"_Complex":       TOKEN_KEYWORD,
			"_Decimal128":    TOKEN_KEYWORD,
			"_Decimal32":     TOKEN_KEYWORD,
			"_Decimal64":     TOKEN_KEYWORD,
			"_Generic":       TOKEN_KEYWORD,
			"_Imaginary":     TOKEN_KEYWORD,
			"_Noreturn":      TOKEN_KEYWORD,
			"_Static_assert": TOKEN_KEYWORD,
			"_Thread_local":  TOKEN_KEYWORD,
		},
	}
}

func (l *Lexer) advance() rune {
	if l.current >= len(l.source) {
		return 0
	}
	r, size := utf8.DecodeRuneInString(l.source[l.current:])
	l.current += size
	l.column++
	return r
}

func (l *Lexer) addToken(tType TokenType) Token {
	text := l.source[l.start:l.current]
	token := Token{Type: tType, Lexeme: text, Line: l.line, Column: l.column - (l.current - l.start)}
	l.start = l.current
	return token
}

func (l *Lexer) Tokenize() []Token {
	var tokens []Token
	for {
		tok := l.scanToken()
		tokens = append(tokens, tok)
		if tok.Type == TOKEN_EOF {
			break
		}
	}
	return tokens
}

func (l *Lexer) scanToken() Token {
	l.skipWhitespace()
	l.start = l.current
	if l.current >= len(l.source) {
		return l.addToken(TOKEN_EOF)
	}

	c := l.advance()
	if unicode.IsLetter(c) || c == '_' {
		return l.identifier()
	}
	if unicode.IsDigit(c) {
		return l.number()
	}
	if strings.ContainsRune("+-*/=<>!&|", c) {
		return l.operator(c)
	}
	if c == '"' {
		return l.string()
	}
	if c == '#' {
		return l.preprocessor()
	}
	if c == '/' {
		if l.match('/') {
			return l.lineComment()
		} else if l.match('*') {
			return l.blockComment()
		}
	}
	if strings.ContainsRune("(),;{}", c) {
		return l.addToken(TOKEN_SEPARATOR)
	}
	return l.addToken(TOKEN_ERROR) // Unrecognized character
}

func (l *Lexer) skipWhitespace() {
	for {
		c := l.peek()
		switch c {
		case ' ', '\r', '\t':
			l.advance()
		case '\n':
			l.line++
			l.column = 0
			l.advance()
		default:
			return
		}
	}
}

func (l *Lexer) peek() rune {
	if l.current >= len(l.source) {
		return 0
	}
	r, _ := utf8.DecodeRuneInString(l.source[l.current:])
	return r
}

func (l *Lexer) match(expected rune) bool {
	if l.peek() == expected {
		l.advance()
		return true
	}
	return false
}

func (l *Lexer) identifier() Token {
	for unicode.IsLetter(l.peek()) || unicode.IsDigit(l.peek()) || l.peek() == '_' {
		l.advance()
	}
	lexeme := l.source[l.start:l.current]
	if tType, ok := l.keywords[lexeme]; ok {
		return l.addToken(tType)
	}
	return l.addToken(TOKEN_IDENTIFIER)
}

func (l *Lexer) number() Token {
	for unicode.IsDigit(l.peek()) {
		l.advance()
	}
	return l.addToken(TOKEN_NUMBER)
}

func (l *Lexer) operator(c rune) Token {
	for strings.ContainsRune("+-*/=<>!&|", l.peek()) {
		l.advance()
	}
	return l.addToken(TOKEN_OPERATOR)
}

func (l *Lexer) string() Token {
	for l.peek() != '"' && l.peek() != 0 {
		if l.peek() == '\n' {
			l.line++
		}
		l.advance()
	}
	if l.peek() == '"' {
		l.advance()
	}
	return l.addToken(TOKEN_STRING)
}

func (l *Lexer) preprocessor() Token {
	if strings.HasPrefix(l.source[l.start:], "#include") {
		l.advance() // skip '#'
		for l.peek() != '\n' && l.peek() != 0 {
			l.advance()
		}
		return l.addToken(TOKEN_INCLUDE)
	}
	for l.peek() != '\n' && l.peek() != 0 {
		l.advance()
	}
	return l.addToken(TOKEN_PREPROCESSOR)
}

func (l *Lexer) lineComment() Token {
	for l.peek() != '\n' && l.peek() != 0 {
		l.advance()
	}
	return l.addToken(TOKEN_COMMENT)
}

func (l *Lexer) blockComment() Token {
	for {
		if l.peek() == '*' && l.match('/') {
			break
		}
		if l.peek() == 0 {
			break
		}
		l.advance()
	}
	return l.addToken(TOKEN_COMMENT)
}
