package codegen

import (
	"bytes"
	"fmt"

	"github.com/darragh-downey/secure-c/pkg/parser"
)

// CodeGenerator holds the buffer for generated code.
type CodeGenerator struct {
	buffer bytes.Buffer
}

// NewCodeGenerator creates a new CodeGenerator instance.
func NewCodeGenerator() *CodeGenerator {
	return &CodeGenerator{}
}

// Generate traverses the AST and generates assembly code.
func (cg *CodeGenerator) Generate(ast *parser.ASTNode) (string, error) {
	if err := cg.generateNode(ast); err != nil {
		return "", err
	}
	return cg.buffer.String(), nil
}

// generateNode generates code for a given AST node.
func (cg *CodeGenerator) generateNode(node *parser.ASTNode) error {
	if node == nil {
		return nil
	}

	switch node.Type {
	case parser.Program:
		// Handle program node (e.g., generate assembly header)
		cg.buffer.WriteString(".globl main\n")
		cg.buffer.WriteString("main:\n")
		for _, child := range node.Children {
			if err := cg.generateNode(child); err != nil {
				return err
			}
		}
		// Handle program exit (e.g., return 0)
		cg.buffer.WriteString("  mov $0, %rax\n")
		cg.buffer.WriteString("  ret\n")

	case parser.ExpressionStatement:
		// Generate code for the expression
		if err := cg.generateNode(node.Children[0]); err != nil {
			return err
		}

	case parser.IntegerLiteral:
		// Generate code to load the integer into a register
		cg.buffer.WriteString(fmt.Sprintf("  mov $%s, %%rax\n", node.Value))

	// Add more cases for other AST node types (e.g., binary operators,
	// variable declarations, if statements, etc.)
	default:
		return fmt.Errorf("unsupported AST node type: %s", node.Type)
	}

	return nil
}
