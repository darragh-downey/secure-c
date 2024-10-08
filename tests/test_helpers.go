package tests

import (
	"encoding/xml"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/darragh-downey/secure-c/pkg/token"

	"github.com/darragh-downey/secure-c/pkg/ast"
)

// loadTestCase reads the test case file which contains the C source code
func loadTestCase(t *testing.T, secure bool, case_id, filename string) string {
	t.Helper()
	var path string

	switch secure {
	case false:
		path = filepath.Join("..", "test_cases", "insecure", case_id, filename)
	default:
		path = filepath.Join("..", "test_cases", "secure", case_id, filename)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read test case file %s: %v", filename, err)
	}
	return string(content)
}

type Tokens struct {
	XMLName xml.Name `xml:"tokens"`
	Tokens  []Token  `xml:"token"`
}

type Token struct {
	Type    token.TokenType `xml:"type"`
	Literal string          `xml:"literal"`
}

// loadExpectedCase reads the expected tokens from the XML file
func loadExpectedCase(t *testing.T, secure bool, caseID, filename string) Tokens {
	t.Helper()
	var path string

	switch secure {
	case false:
		path = filepath.Join("..", "test_cases", "insecure", caseID, filename)
	default:
		path = filepath.Join("..", "test_cases", "secure", caseID, filename)
	}

	tokens, err := parseXML(path)
	if err != nil {
		t.Fatalf("Failed to read expected tokens file %s: %v", filename, err)
	}
	return tokens
}

// parseXML parses the XML file containing the expected tokens
func parseXML(filename string) (Tokens, error) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		return Tokens{}, err
	}
	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)

	var tokens Tokens
	xml.Unmarshal(byteValue, &tokens)

	return tokens, nil
}

// verifyTokens compares the lexer-generated tokens against the expected tokens from the XML file
func verifyTokens(t *testing.T, lexerTokens []token.Token, expectedTokens []Token) {
	if len(lexerTokens) != len(expectedTokens) {
		t.Fatalf("Number of tokens mismatch: expected %d, got %d", len(expectedTokens), len(lexerTokens))
	}

	for i, lexTok := range lexerTokens {
		expTok := expectedTokens[i]
		if lexTok.Type != expTok.Type || lexTok.Literal != expTok.Literal {
			t.Errorf("Mismatch at token %d: expected (%s, %s), got (%s, %s)",
				i, expTok.Type, expTok.Literal, lexTok.Type, lexTok.Literal)
		}
	}
}

// loadExpectedAST reads the expected AST from the XML file
func loadExpectedAST(t *testing.T, secure bool, caseID, filename string) *ast.Node {
	t.Helper()
	var path string

	switch secure {
	case false:
		path = filepath.Join("..", "test_cases", "insecure", caseID, filename)
	default:
		path = filepath.Join("..", "test_cases", "secure", caseID, filename)
	}

	xmlFile, err := os.Open(path)
	if err != nil {
		t.Fatalf("Failed to open expected AST file %s: %v", filename, err)
	}
	defer xmlFile.Close()

	byteValue, _ := io.ReadAll(xmlFile)
	var ast ast.Node
	xml.Unmarshal(byteValue, &ast)

	return &ast
}

// / CompareAST compares the actual AST generated by the parser with the expected AST from the XML file
func compareAST(actual, expected ast.Node) bool {
	if actual.TokenLiteral() != expected.TokenLiteral() || len(actual.Children) != len(expected.Children) {
		return false
	}

	for i := range actual.Children {
		if !compareAST(actual.Children[i], expected.Children[i]) {
			return false
		}
	}

	return true
}
