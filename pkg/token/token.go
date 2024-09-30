package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT  = "IDENT"  // main, foo, bar, x, y, etc.
	STRING = "STRING" // "hello world"

	// Operators
	ASSIGN      = "="
	PLUS        = "+"
	MINUS       = "-"
	BANG        = "!"
	ASTERISK    = "*"
	SLASH       = "/"
	MODULO      = "%"
	BIT_AND     = "&"
	BIT_OR      = "|"
	BIT_XOR     = "^"
	BIT_NOT     = "~"
	SHIFT_LEFT  = "<<"
	SHIFT_RIGHT = ">>"

	// Extended assign operators
	P_ASSIGN   = "+="
	M_ASSIGN   = "-="
	A_ASSIGN   = "*="
	S_ASSIGN   = "/="
	MOD_ASSIGN = "%="
	AND_ASSIGN = "&="
	OR_ASSIGN  = "|="
	XOR_ASSIGN = "^="
	SL_ASSIGN  = "<<="
	SR_ASSIGN  = ">>="

	// Comparison operators
	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="
	LT_EQ  = "<="
	GT_EQ  = ">="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	DOT       = "."
	QUESTION  = "?"

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	// Preprocessor
	PREPROCESSOR = "PREPROCESSOR"

	// Keywords
	AUTO         = "AUTO"
	BREAK        = "BREAK"
	CASE         = "CASE"
	CHAR         = "CHAR"
	CONST        = "CONST"
	CONTINUE     = "CONTINUE"
	DEFAULT      = "DEFAULT"
	DO           = "DO"
	DOUBLE       = "DOUBLE"
	ELSE         = "ELSE"
	ENUM         = "ENUM"
	EXTERN       = "EXTERN"
	FUNCTION     = "FUNCTION"
	FLOAT        = "FLOAT"
	FOR          = "FOR"
	GOTO         = "GOTO"
	IF           = "IF"
	INLINE       = "INLINE"
	INT          = "INT"
	LONG         = "LONG"
	REGISTER     = "REGISTER"
	RESTRICT     = "RESTRICT"
	RETURN       = "RETURN"
	SHORT        = "SHORT"
	SIGNED       = "SIGNED"
	SIZEOF       = "SIZEOF"
	STATIC       = "STATIC"
	STRUCT       = "STRUCT"
	SWITCH       = "SWITCH"
	TYPEDEF      = "TYPEDEF"
	UNION        = "UNION"
	UNSIGNED     = "UNSIGNED"
	VOID         = "VOID"
	VOLATILE     = "VOLATILE"
	WHILE        = "WHILE"
	ALIGNAS      = "ALIGNAS"
	ALIGNOF      = "ALIGNOF"
	ATOMIC       = "ATOMIC"
	BOOL         = "BOOL"
	COMPLEX      = "COMPLEX"
	GENERIC      = "GENERIC"
	IMAGINARY    = "IMAGINARY"
	NORETURN     = "NORETURN"
	STATICASSERT = "STATICASSERT"
	THREADLOCAL  = "THREADLOCAL"
	TRUE         = "TRUE"
	FALSE        = "FALSE"

	// Preprocessor directives
	INCLUDE  = "#include"
	DEFINE   = "#define"
	IFDEF    = "#ifdef"
	IFNDEF   = "#ifndef"
	ENDIF    = "#endif"
	ELIF     = "#elif"
	PRE_ELSE = "#else"
	UNDEF    = "#undef"
	LINE     = "#line"
	ERROR    = "#error"
	PRAGMA   = "#pragma"
)

var keywords = map[string]TokenType{
	"auto":           AUTO,
	"bool":           BOOL,
	"true":           TRUE,
	"false":          FALSE,
	"break":          BREAK,
	"case":           CASE,
	"char":           CHAR,
	"const":          CONST,
	"continue":       CONTINUE,
	"default":        DEFAULT,
	"do":             DO,
	"double":         DOUBLE,
	"enum":           ENUM,
	"extern":         EXTERN,
	"float":          FLOAT,
	"for":            FOR,
	"goto":           GOTO,
	"if":             IF,
	"inline":         INLINE,
	"int":            INT,
	"long":           LONG,
	"register":       REGISTER,
	"restrict":       RESTRICT,
	"return":         RETURN,
	"short":          SHORT,
	"signed":         SIGNED,
	"sizeof":         SIZEOF,
	"static":         STATIC,
	"struct":         STRUCT,
	"switch":         SWITCH,
	"typedef":        TYPEDEF,
	"union":          UNION,
	"unsigned":       UNSIGNED,
	"void":           VOID,
	"volatile":       VOLATILE,
	"while":          WHILE,
	"_Alignas":       ALIGNAS,
	"_Alignof":       ALIGNOF,
	"_Atomic":        ATOMIC,
	"_Bool":          BOOL,
	"_Complex":       COMPLEX,
	"_Generic":       GENERIC,
	"_Imaginary":     IMAGINARY,
	"_Noreturn":      NORETURN,
	"_Static_assert": STATICASSERT,
	"_Thread_local":  THREADLOCAL,
	"include":        INCLUDE,
	"define":         DEFINE,
	"ifdef":          IFDEF,
	"ifndef":         IFNDEF,
	"endif":          ENDIF,
	"elif":           ELIF,
	"else":           ELSE,
	"undef":          UNDEF,
	"line":           LINE,
	"error":          ERROR,
	"pragma":         PRAGMA,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok // one of our languages keywords
	}
	return IDENT // user defined identifiers
}
