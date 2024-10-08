package lexer

import (
	"fmt"

	"github.com/darragh-downey/secure-c/pkg/token"
)

type Lexer struct {
	input        string // user input to parse
	position     int    // current position in input (points to current char)
	readPosition int    // current reading position in input (after current char)
	ch           byte   // current char under examination
	line         int    // current line number
	column       int    // current column number
}

func New(input string) *Lexer {
	l := &Lexer{input: input, line: 1, column: 1} // Initialize line and column numbers
	l.readChar()
	return l
}

// TODO: delete - using for tracing a ']' bug
func (l *Lexer) Info() string {
	return l.input
}

// TODO: delete - using for tracing a ']' bug
func (l *Lexer) Busted() string {
	return fmt.Sprintf("position=%d readPostition=%d char=%q", l.position, l.readPosition, l.ch)
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1

	// Update line and column numbers
	if l.ch == '\n' {
		l.line++
		l.column = 1 // Reset column at the beginning of a new line
	} else {
		l.column++
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = l.newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = l.newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = l.newToken(token.COMMA, l.ch)
	case '(':
		tok = l.newToken(token.LPAREN, l.ch)
	case ')':
		tok = l.newToken(token.RPAREN, l.ch)
	case '{':
		tok = l.newToken(token.LBRACE, l.ch)
	case '}':
		tok = l.newToken(token.RBRACE, l.ch)
	case '[':
		tok = l.newToken(token.LBRACKET, l.ch)
	case ']':
		tok = l.newToken(token.RBRACKET, l.ch)
	case '+':
		// see '=' for how to implement +=
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.P_ASSIGN, Literal: literal}
		} else {
			tok = l.newToken(token.PLUS, l.ch)
		}
	case '-':
		// see '=' for how to implement -=
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.M_ASSIGN, Literal: literal}
		} else {
			tok = l.newToken(token.MINUS, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = l.newToken(token.BANG, l.ch)
		}
	case '/':
		// see '=' for how to implement /=
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.S_ASSIGN, Literal: literal}
		} else {
			tok = l.newToken(token.SLASH, l.ch)
		}
	case '*':
		// see '=' for how to implement *=
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.A_ASSIGN, Literal: literal}
		} else {
			tok = l.newToken(token.ASTERISK, l.ch)
		}
	case '<':
		// see '=' for how to implement <=
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.LT_EQ, Literal: literal}
		} else {
			tok = l.newToken(token.LT, l.ch)
		}
	case '>':
		// see '=' for how to implement >=
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.GT_EQ, Literal: literal}
		} else {
			tok = l.newToken(token.GT, l.ch)
		}
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()

	case ':':
		tok = l.newToken(token.COLON, l.ch)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = l.newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
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

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) newToken(tokenType token.TokenType, ch byte) token.Token {
	tok := token.Token{Type: tokenType, Literal: string(ch), Line: l.line, Column: l.column}
	l.readChar()
	return tok
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readString() string {
	pos := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 { // reads until closing double-quote or end of input
			break
		}
	}
	return l.input[pos:l.position]
}
