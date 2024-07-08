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
	ILLEGAL = "ILLEGAL" // Illegal token
	EOF     = "EOF"     // End of file token
	IDENT   = "IDENT"   // Identifier token

	STRING       = "STRING"       // String literal token
	PREPROCESSOR = "PREPROCESSOR" // Preprocessor directive token

	// Operators
	ASSIGN   = "ASSIGN"   // Assignment operator token
	PLUS     = "PLUS"     // Addition operator token
	MINUS    = "MINUS"    // Subtraction operator token
	BANG     = "BANG"     // Negation operator token
	ASTERISK = "ASTERISK" // Multiplication operator token
	SLASH    = "SLASH"    // Division operator token
	MOD      = "MOD"      // Modulo operator token
	AND      = "AND"      // Logical AND operator token
	OR       = "OR"       // Logical OR operator token
	NOT      = "NOT"      // Logical NOT operator token

	LT     = "LT"     // Less than operator token
	GT     = "GT"     // Greater than operator token
	LE     = "LE"     // Less than or equal to operator token
	GE     = "GE"     // Greater than or equal to operator token
	EQ     = "EQ"     // Equality operator token
	NOT_EQ = "NOT_EQ" // Not equal to operator token

	BIT_AND = "BIT_AND" // Bitwise AND operator token
	BIT_OR  = "BIT_OR"  // Bitwise OR operator token
	BIT_XOR = "BIT_XOR" // Bitwise XOR operator token
	BIT_NOT = "BIT_NOT" // Bitwise NOT operator token
	SHL     = "SHL"     // Left shift operator token
	SHR     = "SHR"     // Right shift operator token

	PLUS_EQ  = "PLUS_EQ"  // Addition assignment operator token
	MINUS_EQ = "MINUS_EQ" // Subtraction assignment operator token
	MUL_EQ   = "MUL_EQ"   // Multiplication assignment operator token
	DIV_EQ   = "DIV_EQ"   // Division assignment operator token
	MOD_EQ   = "MOD_EQ"   // Modulo assignment operator token
	AND_EQ   = "AND_EQ"   // Bitwise AND assignment operator token
	OR_EQ    = "OR_EQ"    // Bitwise OR assignment operator token
	XOR_EQ   = "XOR_EQ"   // Bitwise XOR assignment operator token
	SHL_EQ   = "SHL_EQ"   // Left shift assignment operator token
	SHR_EQ   = "SHR_EQ"   // Right shift assignment operator token

	INC = "INC" // Increment operator token
	DEC = "DEC" // Decrement operator token

	PTR = "PTR" // Pointer operator token
	DOT = "DOT" // Dot operator token

	// Delimiters
	COMMA     = "COMMA"     // Comma delimiter token
	SEMICOLON = "SEMICOLON" // Semicolon delimiter token
	COLON     = "COLON"     // Colon delimiter token
	QUESTION  = "QUESTION"  // Question mark delimiter token
	LPAREN    = "LPAREN"    // Left parenthesis delimiter token
	RPAREN    = "RPAREN"    // Right parenthesis delimiter token
	LBRACE    = "LBRACE"    // Left brace delimiter token
	RBRACE    = "RBRACE"    // Right brace delimiter token
	LBRACK    = "LBRACK"    // Left bracket delimiter token
	RBRACK    = "RBRACK"    // Right bracket delimiter token

	// Keywords
	FUNCTION = "FUNCTION" // Function keyword token
	LET      = "LET"      // Let keyword token
	AUTO     = "AUTO"     // Auto keyword token
	BREAK    = "BREAK"    // Break keyword token
	CASE     = "CASE"     // Case keyword token
	CHAR     = "CHAR"     // Char keyword token
	CONST    = "CONST"    // Const keyword token
	CONTINUE = "CONTINUE" // Continue keyword token
	DEFAULT  = "DEFAULT"  // Default keyword token
	DO       = "DO"       // Do keyword token
	DOUBLE   = "DOUBLE"   // Double keyword token
	ELSE     = "ELSE"     // Else keyword token
	ENUM     = "ENUM"     // Enum keyword token
	EXTERN   = "EXTERN"   // Extern keyword token
	FLOAT    = "FLOAT"    // Float keyword token
	FOR      = "FOR"      // For keyword token
	GOTO     = "GOTO"     // Goto keyword token
	IF       = "IF"       // If keyword token
	INLINE   = "INLINE"   // Inline keyword token
	INT      = "INT"      // Int keyword token
	LONG     = "LONG"     // Long keyword token
	REGISTER = "REGISTER" // Register keyword token
	RESTRICT = "RESTRICT" // Restrict keyword token
	RETURN   = "RETURN"   // Return keyword token
	SHORT    = "SHORT"    // Short keyword token
	SIGNED   = "SIGNED"   // Signed keyword token
	SIZEOF   = "SIZEOF"   // Sizeof keyword token
	STATIC   = "STATIC"   // Static keyword token
	STRUCT   = "STRUCT"   // Struct keyword token
	SWITCH   = "SWITCH"   // Switch keyword token
	TYPEDEF  = "TYPEDEF"  // Typedef keyword token
	UNION    = "UNION"    // Union keyword token
	UNSIGNED = "UNSIGNED" // Unsigned keyword token
	VOID     = "VOID"     // Void keyword token
	VOLATILE = "VOLATILE" // Volatile keyword token
	WHILE    = "WHILE"    // While keyword token
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
