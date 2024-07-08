package tests

import (
	"testing"

	"github.com/darragh-downey/secure-c/lexer"
)

func TestLexer(t *testing.T) {
	testCases := []struct {
		name     string // name of the test case
		secure   bool   // true if the test case is in the secure directory and false if it is in the insecure directory
		caseID   string // directory name in secure/insecure directory
		filename string // name of file in case_id directory
		expected string // name of the XML file containing the expected parsed tokens
	}{
		{
			name:     "Hello World",
			secure:   true,
			filename: "hello_world.c",
			caseID:   "case_01",
			expected: "tokens.xml",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			source := loadTestCase(t, tc.secure, tc.caseID, tc.filename)
			expectedTokens := loadExpectedCase(t, tc.secure, tc.caseID, tc.expected).Tokens

			l := lexer.New(source)
			actualTokens := l.IterateTokens()

			if len(actualTokens) != len(expectedTokens) {
				t.Fatalf("Number of tokens mismatch: expected %d, got %d", len(expectedTokens), len(actualTokens))
			}

			for i, actualToken := range actualTokens {
				expectedToken := expectedTokens[i]
				if actualToken.Type != lexer.TokenType(expectedToken.Type) || actualToken.Literal != expectedToken.Literal {
					t.Errorf("Mismatch at token %d: expected %s, %s , got  %s, %s ",
						i, expectedToken.Type, expectedToken.Literal, actualToken.Type, actualToken.Literal)
				}
			}
		})
	}
}
