package tests

import (
	"testing"

	"github.com/darragh-downey/secure-c/lexer"
)

func TestLexer(t *testing.T) {
	source := "int main() { return 0; }"
	l := lexer.NewLexer(source)
	tokens := l.Tokenize()

	expected := []lexer.TokenType{
		lexer.TOKEN_IDENTIFIER, lexer.TOKEN_IDENTIFIER, lexer.TOKEN_OPERATOR,
		lexer.TOKEN_IDENTIFIER, lexer.TOKEN_NUMBER, lexer.TOKEN_OPERATOR,
		lexer.TOKEN_EOF,
	}

	for i, token := range tokens {
		if token.Type != expected[i] {
			t.Errorf("Expected token %v, got %v", expected[i], token.Type)
		}
	}
}
