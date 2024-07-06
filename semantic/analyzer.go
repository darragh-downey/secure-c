package semantic

import (
	"fmt"

	"github.com/darragh-downey/secure-c/parser"
)

type Analyzer struct{}

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

func (a *Analyzer) Analyze(ast *parser.ASTNode) error {
	return a.checkNode(ast)
}

func (a *Analyzer) checkNode(node *parser.ASTNode) error {
	// Example: Check for undefined variables, type mismatches, etc.
	if node == nil {
		return nil
	}

	if node.Value == "undefined_variable" {
		return fmt.Errorf("undefined variable: %s", node.Value)
	}

	for _, child := range node.Children {
		if err := a.checkNode(child); err != nil {
			return err
		}
	}
	return nil
}
