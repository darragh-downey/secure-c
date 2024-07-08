package tests

import (
	"testing"

	"github.com/darragh-downey/secure-c/pkg/lexer"
	"github.com/darragh-downey/secure-c/pkg/parser"
)

func TestParser(t *testing.T) {
	testCases := []struct {
		name     string
		secure   bool
		caseID   string
		filename string
		expected string
	}{
		{
			name:     "HelloWorld",
			secure:   true,
			filename: "hello_world.c",
			caseID:   "case_01",
			expected: "ast.xml",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			source := loadTestCase(t, tc.secure, tc.caseID, tc.filename)
			l := lexer.New(source)
			tokens := l.IterateTokens()

			p := parser.NewParser(tokens)
			ast, err := p.Parse()
			if err != nil {
				t.Fatalf("Parser error: %v", err)
			}

			// Compare the generated AST with the expected AST from the XML file
			expectedAST := loadExpectedAST(t, tc.secure, tc.caseID, tc.expected)
			if !compareAST(ast, expectedAST) {
				t.Errorf("AST does not match expected AST")
			}
		})
	}
}
