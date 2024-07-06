package tests

import (
	"testing"

	"github.com/darragh-downey/secure-c/lexer"
	"github.com/darragh-downey/secure-c/parser"
)

func TestParser(t *testing.T) {
	source := "int main() { return 0; }"
	l := lexer.NewLexer(source)
	tokens := l.Tokenize()

	p := parser.NewParser(tokens)
	ast, err := p.Parse()
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}

	if ast == nil {
		t.Fatal("Expected non-nil AST")
	}
}
