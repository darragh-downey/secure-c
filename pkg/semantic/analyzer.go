package semantic

import (
	"fmt"

	"github.com/darragh-downey/secure-c/parser"
)

type Analyzer struct {
	unsafeFunctions map[string]bool
}

func NewAnalyzer() *Analyzer {
	return &Analyzer{
		unsafeFunctions: map[string]bool{
			"gets":   true,
			"strcpy": true,
			"system": true,
		},
	}
}

func (a *Analyzer) Analyze(ast *parser.ASTNode) error {
	return a.checkNode(ast)
}

func (a *Analyzer) checkNode(node *parser.ASTNode) error {
	if node == nil {
		return nil
	}

	// Check for unsafe function usage
	if a.unsafeFunctions[node.Value] {
		return fmt.Errorf("unsafe function usage: %s", node.Value)
	}

	// Check for format string vulnerabilities
	if node.Value == "snprintf" || node.Value == "printf" {
		if len(node.Children) > 1 && node.Children[1].Value == "%s" {
			return fmt.Errorf("format string vulnerability: %s", node.Value)
		}
	}

	// Check for potential buffer overflow
	if node.Value == "strcpy" {
		if len(node.Children) > 1 && len(node.Children[1].Value) > 10 {
			return fmt.Errorf("potential buffer overflow: %s", node.Value)
		}
	}

	// Check for integer overflow
	if node.Value == "+" {
		for _, child := range node.Children {
			if child.Value == "4294967295" {
				return fmt.Errorf("potential integer overflow: %s", node.Value)
			}
		}
	}

	for _, child := range node.Children {
		if err := a.checkNode(child); err != nil {
			return err
		}
	}
	return nil
}
