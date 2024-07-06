package lexer

type TokenType string

const (
	TOKEN_IDENTIFIER   TokenType = "IDENTIFIER"
	TOKEN_KEYWORD      TokenType = "KEYWORD"
	TOKEN_NUMBER       TokenType = "NUMBER"
	TOKEN_OPERATOR     TokenType = "OPERATOR"
	TOKEN_EOF          TokenType = "EOF"
	TOKEN_ERROR        TokenType = "ERROR"
	TOKEN_STRING       TokenType = "STRING"
	TOKEN_SEPARATOR    TokenType = "SEPARATOR"
	TOKEN_PREPROCESSOR TokenType = "PREPROCESSOR"
	TOKEN_INCLUDE      TokenType = "INCLUDE"
)

type Token struct {
	Type   TokenType
	Lexeme string
	Line   int
	Column int
}
