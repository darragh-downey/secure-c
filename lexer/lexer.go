package lexer

import (
	"strings"
	"unicode"
)

type Lexer struct {
	source  string
	start   int
	current int
	line    int
	column  int
}

func NewLexer(source string) *Lexer {
	return &Lexer{source: source, line: 1, column: 1}
}

func (l *Lexer) advance() rune {
	if l.current >= len(l.source) {
		return 0
	}
	r := rune(l.source[l.current])
	l.current++
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
		tokens = append(tokens, l.scanToken())
		if tokens[len(tokens)-1].Type == TOKEN_EOF {
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
	if strings.ContainsRune("+-*/=", c) {
		return l.addToken(TOKEN_OPERATOR)
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
	return rune(l.source[l.current])
}

func (l *Lexer) identifier() Token {
	for unicode.IsLetter(l.peek()) || unicode.IsDigit(l.peek()) || l.peek() == '_' {
		l.advance()
	}
	return l.addToken(TOKEN_IDENTIFIER)
}

func (l *Lexer) number() Token {
	for unicode.IsDigit(l.peek()) {
		l.advance()
	}
	return l.addToken(TOKEN_NUMBER)
}
