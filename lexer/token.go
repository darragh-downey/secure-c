package lexer

type TokenType string

const (
	TOKEN_IDENTIFIER TokenType = "IDENTIFIER"
	TOKEN_KEYWORD    TokenType = "KEYWORD"
	TOKEN_NUMBER     TokenType = "NUMBER"
	TOKEN_OPERATOR   TokenType = "OPERATOR"
	TOKEN_EOF        TokenType = "EOF"
	TOKEN_ERROR      TokenType = "ERROR"
)

type Token struct {
	Type   TokenType
	Lexeme string
	Line   int
	Column int
}
