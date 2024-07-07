package lexer

import "strings"

// Token represents a lexical token.
type Token struct {
	Type    TokenType
	Literal string
	Line    int
}

// TokenType identifies the type of lexical tokens.
type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT" // main, foo, bar, etc.

	STRING       = "STRING"       // "foo"
	PREPROCESSOR = "PREPROCESSOR" // #include <stdio.h>

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	MOD      = "%"
	AND      = "&&"
	OR       = "||"
	NOT      = "!"

	LT     = "<"
	GT     = ">"
	LE     = "<="
	GE     = ">="
	EQ     = "=="
	NOT_EQ = "!="

	BIT_AND = "&"
	BIT_OR  = "|"
	BIT_XOR = "^"
	BIT_NOT = "~"
	SHL     = "<<"
	SHR     = ">>"

	PLUS_EQ  = "+="
	MINUS_EQ = "-="
	MUL_EQ   = "*="
	DIV_EQ   = "/="
	MOD_EQ   = "%="
	AND_EQ   = "&="
	OR_EQ    = "|="
	XOR_EQ   = "^="
	SHL_EQ   = "<<="
	SHR_EQ   = ">>="

	INC = "++"
	DEC = "--"

	PTR = "->"
	DOT = "."

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	QUESTION  = "?"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACK    = "["
	RBRACK    = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	AUTO     = "AUTO"
	BREAK    = "BREAK"
	CASE     = "CASE"
	CHAR     = "CHAR"
	CONST    = "CONST"
	CONTINUE = "CONTINUE"
	DEFAULT  = "DEFAULT"
	DO       = "DO"
	DOUBLE   = "DOUBLE"
	ELSE     = "ELSE"
	ENUM     = "ENUM"
	EXTERN   = "EXTERN"
	FLOAT    = "FLOAT"
	FOR      = "FOR"
	GOTO     = "GOTO"
	IF       = "IF"
	INLINE   = "INLINE"
	INT      = "INT"
	LONG     = "LONG"
	REGISTER = "REGISTER"
	RESTRICT = "RESTRICT"
	RETURN   = "RETURN"
	SHORT    = "SHORT"
	SIGNED   = "SIGNED"
	SIZEOF   = "SIZEOF"
	STATIC   = "STATIC"
	STRUCT   = "STRUCT"
	SWITCH   = "SWITCH"
	TYPEDEF  = "TYPEDEF"
	UNION    = "UNION"
	UNSIGNED = "UNSIGNED"
	VOID     = "VOID"
	VOLATILE = "VOLATILE"
	WHILE    = "WHILE"
)

func lookupIdent(ident string) TokenType {
	switch strings.ToLower(ident) {
	case "fn":
		return FUNCTION
	case "let":
		return LET
	case "auto":
		return AUTO
	case "break":
		return BREAK
	case "case":
		return CASE
	case "char":
		return CHAR
	case "const":
		return CONST
	case "continue":
		return CONTINUE
	case "default":
		return DEFAULT
	case "do":
		return DO
	case "double":
		return DOUBLE
	case "else":
		return ELSE
	case "enum":
		return ENUM
	case "extern":
		return EXTERN
	case "float":
		return FLOAT
	case "for":
		return FOR
	case "goto":
		return GOTO
	case "if":
		return IF
	case "inline":
		return INLINE
	case "int":
		return INT
	case "long":
		return LONG
	case "register":
		return REGISTER
	case "restrict":
		return RESTRICT
	case "return":
		return RETURN
	case "short":
		return SHORT
	case "signed":
		return SIGNED
	case "sizeof":
		return SIZEOF
	case "static":
		return STATIC
	case "struct":
		return STRUCT
	case "switch":
		return SWITCH
	case "typedef":
		return TYPEDEF
	case "union":
		return UNION
	case "unsigned":
		return UNSIGNED
	case "void":
		return VOID
	case "volatile":
		return VOLATILE
	case "while":
		return WHILE
	default:
		return IDENT
	}
}
