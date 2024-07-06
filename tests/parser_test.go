package tests

import (
	"testing"

	"github.com/darragh-downey/secure-c/lexer"
	"github.com/darragh-downey/secure-c/parser"
)

func TestParser(t *testing.T) {
	testCases := []struct {
		name     string
		secure   bool
		filename string
	}{
		{name: "Insecure", secure: false, filename: "buffer_overflow.c"},
		{name: "Basic", secure: true, filename: "basic.c"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			source := loadTestCase(t, tc.secure, tc.filename)
			preprocessedSource := lexer.Preprocess(source)
			l := lexer.NewLexer(preprocessedSource)
			tokens := l.Tokenize()

			p := parser.NewParser(tokens)
			ast, err := p.Parse()
			if err != nil {
				t.Fatalf("Parse error: %v", err)
			}

			if ast == nil {
				t.Fatal("Expected non-nil AST")
			}
		})
	}
}
