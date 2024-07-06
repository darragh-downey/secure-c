package tests

import (
	"testing"

	"github.com/darragh-downey/secure-c/optimizer"
	"github.com/darragh-downey/secure-c/parser"
)

func TestOptimizer(t *testing.T) {
	ast := &parser.ASTNode{
		Value: "main",
		Children: []*parser.ASTNode{
			{Value: "return", Children: []*parser.ASTNode{{Value: "0"}}},
		},
	}

	opt := optimizer.NewOptimizer()
	err := opt.Optimize(ast)
	if err != nil {
		t.Fatalf("Optimization error: %v", err)
	}
}
