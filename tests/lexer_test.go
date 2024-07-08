package tests

import (
	"testing"

	"github.com/darragh-downey/secure-c/pkg/token"

	"github.com/darragh-downey/secure-c/pkg/lexer"
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
			var tokens []token.Token
			for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
				tokens = append(tokens, tok)
			}

			verifyTokens(t, tokens, expectedTokens)
		})
	}
}
