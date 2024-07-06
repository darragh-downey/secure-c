package tests

import (
	"testing"

	"github.com/darragh-downey/secure-c/lexer"
)

func TestLexer(t *testing.T) {
	testCases := []struct {
		name     string
		secure   bool
		filename string
	}{
		{name: "BufferOverflow", secure: false, filename: "buffer_overflow.c"},
		{name: "Basic", secure: true, filename: "basic.c"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			source := loadTestCase(t, tc.secure, tc.filename)
			l := lexer.NewLexer(source)
			tokens := l.Tokenize()

			if len(tokens) == 0 {
				t.Fatalf("Expected tokens, got none")
			}

			expectedTypes := []lexer.TokenType{
				lexer.TOKEN_PREPROCESSOR, lexer.TOKEN_PREPROCESSOR, lexer.TOKEN_PREPROCESSOR,
				lexer.TOKEN_KEYWORD, lexer.TOKEN_IDENTIFIER, lexer.TOKEN_SEPARATOR, lexer.TOKEN_KEYWORD, lexer.TOKEN_IDENTIFIER, lexer.TOKEN_SEPARATOR,
				lexer.TOKEN_KEYWORD, lexer.TOKEN_IDENTIFIER, lexer.TOKEN_SEPARATOR, lexer.TOKEN_OPERATOR,
				lexer.TOKEN_IDENTIFIER, lexer.TOKEN_SEPARATOR, lexer.TOKEN_STRING, lexer.TOKEN_SEPARATOR,
				lexer.TOKEN_KEYWORD, lexer.TOKEN_IDENTIFIER, lexer.TOKEN_SEPARATOR, lexer.TOKEN_OPERATOR,
				lexer.TOKEN_IDENTIFIER, lexer.TOKEN_SEPARATOR, lexer.TOKEN_IDENTIFIER, lexer.TOKEN_OPERATOR,
				lexer.TOKEN_STRING, lexer.TOKEN_SEPARATOR, lexer.TOKEN_SEPARATOR, lexer.TOKEN_SEPARATOR,
				// Add more expected token types as needed
			}

			for i, token := range tokens {
				if i < len(expectedTypes) && token.Type != expectedTypes[i] {
					t.Errorf("Expected token type %v, got %v", expectedTypes[i], token.Type)
				}
			}
		})
	}
}
