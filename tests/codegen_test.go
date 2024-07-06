package tests

import (
	"testing"

	"github.com/darragh-downey/secure-c/codegen"
	"github.com/darragh-downey/secure-c/parser"
)

func TestCodeGenerator(t *testing.T) {
	ast := &parser.ASTNode{
		Value: "main",
		Children: []*parser.ASTNode{
			{Value: "return", Children: []*parser.ASTNode{{Value: "0"}}},
		},
	}

	codeGen := codegen.NewCodeGenerator()
	code, err := codeGen.Generate(ast)
	if err != nil {
		t.Fatalf("Code generation error: %v", err)
	}

	expected := "Node: main\nNode: return\nNode: 0\n"
	if code != expected {
		t.Errorf("Expected code:\n%s\nGot:\n%s", expected, code)
	}
}
