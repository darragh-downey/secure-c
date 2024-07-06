package tests

import (
	"testing"

	"github.com/darragh-downey/secure-c/parser"
	"github.com/darragh-downey/secure-c/semantic"
)

func TestSemanticAnalyzer(t *testing.T) {
	ast := &parser.ASTNode{
		Value: "main",
		Children: []*parser.ASTNode{
			{Value: "return", Children: []*parser.ASTNode{{Value: "0"}}},
		},
	}

	analyzer := semantic.NewAnalyzer()
	err := analyzer.Analyze(ast)
	if err != nil {
		t.Fatalf("Semantic error: %v", err)
	}
}
