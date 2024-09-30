package semantic

import (
	"fmt"

	"github.com/darragh-downey/secure-c/pkg/ast"
	"github.com/darragh-downey/secure-c/pkg/parser"
)

type Analyzer struct {
	unsafeFunctions map[string]bool
	securityIssues  map[string]spec.SecurityIssue // Add this field
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
	// Load security specifications
	var err error
	a.securityIssues, err = spec.LoadSpecifications("specifications") // Path to your specs directory
	if err != nil {
		return fmt.Errorf("error loading specifications: %w", err)
	}

	return a.checkNode(ast, a.securityIssues)
}

func (a *Analyzer) checkNode(node *parser.ASTNode, securityIssues map[string]spec.SecurityIssue) error {
	if node == nil {
		return nil
	}

	switch node.Type {
	case parser.CallExpression:
		if err := a.checkUnsafeFunctionCall(node, securityIssues); err != nil {
			return err
		}
		if err := a.checkFormatStringVulnerability(node, securityIssues); err != nil {
			return err
		}
	case parser.InfixExpression:
		if err := a.checkIntegerOverflow(node, securityIssues); err != nil {
			return err
		}
	case parser.IndexExpression:
		if err := a.checkArrayAccess(node, securityIssues); err != nil {
			return err
		}
	// ... (add cases for other node types and rules)
	default:

	}

	for _, child := range node.Children {
		if err := a.checkNode(child, securityIssues); err != nil {
			return err
		}
	}
	return nil
}

// Dedicated function to check for unsafe function calls (FIO30-C)
func (a *Analyzer) checkUnsafeFunctionCall(node *parser.ASTNode, securityIssues map[string]spec.SecurityIssue) error {
	if a.unsafeFunctions[node.Value] {
		return reportIssue("FIO30-C", node, securityIssues)
	}
	return nil
}

// Dedicated function to check for format string vulnerabilities (FIO30-C)
func (a *Analyzer) checkFormatStringVulnerability(node *parser.ASTNode, securityIssues map[string]spec.SecurityIssue) error {
	if node.Value == "snprintf" || node.Value == "printf" {
		if len(node.Children) > 1 {
			formatStringNode := node.Children[1]

			if _, ok := formatStringNode.(*ast.StringLiteral); !ok {
				return reportIssue("FIO30-C", node, securityIssues)
			}

			// Additional checks (optional):
			// - Check if the string literal contains format specifiers (e.g., %s, %d)
			// if strings.Contains(formatStringNode.TokenLiteral(), "%") { ... }
			// - Check if the string literal is concatenated with other expressions
		}
	}
	return nil
}

func (a *Analyzer) checkIntegerOverflow(node *parser.ASTNode, securityIssues map[string]spec.SecurityIssue) error {
	if node.Operator != "+" && node.Operator != "-" && node.Operator != "*" {
		return nil // Not an arithmetic operation
	}

	// Check if both operands are signed integer literals (for simplicity)
	left, ok1 := node.Left.(*ast.IntegerLiteral)
	right, ok2 := node.Right.(*ast.IntegerLiteral)
	if !ok1 || !ok2 {
		return nil // Not integer literals, handle other cases later
	}

	// ... (Logic to check for potential overflow based on left.Value and right.Value)
	// You'll need to consider the maximum and minimum values for signed integers
	// and the operation being performed.

	if true { /* potential overflow detected */
		return reportIssue("INT32-C", node, securityIssues)
	}

	return nil
}

func (a *Analyzer) checkArrayAccess(node *parser.ASTNode, securityIssues map[string]spec.SecurityIssue) error {
	if node.Type != parser.IndexExpression {
		return nil // Not an array access
	}

	arrayNode, ok := node.Left.(*ast.ArrayLiteral) // Assuming the array is a literal
	if !ok {
		// Handle the case where the array is not a literal (e.g., a variable)
		// You'll need more sophisticated analysis to determine the bounds in this case.
		return nil
	}

	indexNode, ok := node.Index.(*ast.IntegerLiteral) // Assuming the index is a literal
	if !ok {
		// Handle the case where the index is not a literal (e.g., a variable)
		return nil
	}

	if indexNode.Value < 0 || indexNode.Value >= int64(arrayNode.Size) {
		return reportIssue("ARR30-C", node, securityIssues) // Array out of bounds
	}

	return nil
}

// ... (add more dedicated functions for other rules)

func reportIssue(ruleID string, node *parser.ASTNode, securityIssues map[string]spec.SecurityIssue) error {
	issue, ok := securityIssues[ruleID]
	if !ok {
		return fmt.Errorf("rule ID not found: %s", ruleID)
	}

	// Generate a warning using issue.Description, issue.Severity, issue.URL, etc.
	fmt.Printf("Warning: %s (line %d, col %d) - %s\n",
		issue.ID, node.Token.Line, node.Token.Column, issue.Description)
	fmt.Printf("  Recommendation: %s\n", issue.Recommendation)
	fmt.Printf("  See SEI CERT rule: %s\n", issue.URL)
	return nil
}
