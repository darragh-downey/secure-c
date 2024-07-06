package tests

import (
	"strings"
	"testing"

	"github.com/darragh-downey/secure-c/lexer"
	"github.com/darragh-downey/secure-c/parser"
	"github.com/darragh-downey/secure-c/semantic"
)

func TestSemanticAnalyzer(t *testing.T) {
	testCases := []struct {
		name           string
		secure         bool
		filename       string
		expectedErrors []string
	}{
		{
			name:     "Insecure",
			secure:   false,
			filename: "buffer_overflow.c",
			expectedErrors: []string{
				"unsafe function usage: gets",
				"unsafe function usage: strcpy",
				"unsafe function usage: system",
				"format string vulnerability: snprintf",
				"potential buffer overflow: strcpy",
				"potential integer overflow: +",
			},
		},
		{
			name:           "Basic",
			secure:         true,
			filename:       "basic.c",
			expectedErrors: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			source := loadTestCase(t, tc.secure, tc.filename)
			l := lexer.NewLexer(source)
			tokens := l.Tokenize()

			p := parser.NewParser(tokens)
			ast, err := p.Parse()
			if err != nil {
				t.Fatalf("Parse error: %v", err)
			}

			semanticAnalyzer := semantic.NewAnalyzer()
			semanticErr := semanticAnalyzer.Analyze(ast)
			if len(tc.expectedErrors) == 0 && semanticErr != nil {
				t.Fatalf("Expected no semantic errors, got: %v", semanticErr)
			}

			foundErrors := map[string]bool{}
			for _, err := range tc.expectedErrors {
				foundErrors[err] = false
			}

			if semanticErr != nil {
				errMsg := semanticErr.Error()
				for _, expectedError := range tc.expectedErrors {
					if strings.Contains(errMsg, expectedError) {
						foundErrors[expectedError] = true
					}
				}
			}

			for _, expectedError := range tc.expectedErrors {
				if !foundErrors[expectedError] {
					t.Errorf("Expected semantic error: %s", expectedError)
				}
			}
		})
	}
}
