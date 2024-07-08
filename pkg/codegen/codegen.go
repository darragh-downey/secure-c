package codegen

import (
	"bytes"
	"fmt"

	"github.com/darragh-downey/secure-c/pkg/parser"
)

type CodeGenerator struct {
	buffer bytes.Buffer
}

func NewCodeGenerator() *CodeGenerator {
	return &CodeGenerator{}
}

func (cg *CodeGenerator) Generate(ast *parser.ASTNode) (string, error) {
	if err := cg.generateNode(ast); err != nil {
		return "", err
	}
	return cg.buffer.String(), nil
}

func (cg *CodeGenerator) generateNode(node *parser.ASTNode) error {
	if node == nil {
		return nil
	}

	// Example: Generate code for the node
	cg.buffer.WriteString(fmt.Sprintf("Node: %s\n", node.Value))
	for _, child := range node.Children {
		if err := cg.generateNode(child); err != nil {
			return err
		}
	}
	return nil
}
