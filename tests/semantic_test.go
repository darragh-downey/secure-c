package tests

import (
	"strings"
	"testing"

	"github.com/darragh-downey/secure-c/pkg/lexer"
	"github.com/darragh-downey/secure-c/pkg/parser"
	"github.com/darragh-downey/secure-c/pkg/semantic"
)

func TestSemanticAnalyzer(t *testing.T) {
	testCases := []struct {
		name           string
		secure         bool
		caseID         string
		filename       string
		expectedErrors []string
	}{
		{
			name:           "secure",
			secure:         true,
			caseID:         "case_01",
			filename:       "hello_world.c",
			expectedErrors: []string{},
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
