package lexer

import (
	"fmt"
)

// Lexer represents a lexical scanner.
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	line         int  // current line number
}

// New creates a new instance of Lexer for the input string.
func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1}
	l.readChar()
	return l
}

// readChar gives us the next character and advances our position in the input string.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	if l.ch == '\n' {
		l.line++
	}

	l.position = l.readPosition
	l.readPosition++
}

// NextToken returns the next token from the input.
func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '#':
		tok.Type = PREPROCESSOR
		tok.Literal = l.readPreprocessorDirective()
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(ASSIGN, l.ch, l.line)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(BANG, l.ch, l.line)
		}
	case '+':
		if l.peekChar() == '+' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: INC, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: PLUS_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(PLUS, l.ch, l.line)
		}
	case '-':
		if l.peekChar() == '-' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: DEC, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: MINUS_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(MINUS, l.ch, l.line)
		}
	case '/':
		if l.peekChar() == '/' {
			for l.ch != '\n' && l.ch != 0 {
				l.readChar()
			}
		} else if l.peekChar() == '*' {
			l.readChar()
			l.readChar()
			for !(l.ch == '*' && l.peekChar() == '/') {
				l.readChar()
			}
			l.readChar()
			l.readChar()
		} else {
			tok = newToken(SLASH, l.ch, l.line)
		}
	case '*':
		tok = newToken(ASTERISK, l.ch, l.line)
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: LE, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == '<' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: SHL, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(LT, l.ch, l.line)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: GE, Literal: string(ch) + string(l.ch)}
		} else if l.peekChar() == '>' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: SHR, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(GT, l.ch, l.line)
		}
	case ';':
		tok = newToken(SEMICOLON, l.ch, l.line)
	case ',':
		tok = newToken(COMMA, l.ch, l.line)
	case '(':
		tok = newToken(LPAREN, l.ch, l.line)
	case ')':
		tok = newToken(RPAREN, l.ch, l.line)
	case '{':
		tok = newToken(LBRACE, l.ch, l.line)
	case '}':
		tok = newToken(RBRACE, l.ch, l.line)
	case '"':
		tok.Type = STRING
		tok.Literal = l.readString()
	case 0:
		tok.Literal = ""
		tok.Type = EOF
		tok.Line = l.line
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = lookupIdent(tok.Literal)
			tok.Line = l.line
			return tok
		} else if isDigit(l.ch) {
			tok.Type = INT
			tok.Literal = l.readNumber()
			tok.Line = l.line
			return tok
		} else {
			tok = newToken(ILLEGAL, l.ch, l.line)
			fmt.Printf("Illegal character: %c on line %d\n", l.ch, l.line)
		}
	}

	l.readChar()
	return tok
}

// IterateTokens iterates over all tokens produced by the lexer
func (l *Lexer) IterateTokens() []Token {
	var tokens []Token
	for tok := l.NextToken(); tok.Type != EOF; tok = l.NextToken() {
		tokens = append(tokens, tok)
	}
	tokens = append(tokens, Token{Type: EOF, Literal: ""})
	return tokens
}

func newToken(tokenType TokenType, ch byte, line int) Token {
	return Token{Type: tokenType, Literal: string(ch), Line: line}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) readPreprocessorDirective() string {
	position := l.position
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
