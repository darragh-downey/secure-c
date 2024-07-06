package parser

type ASTNode struct {
	Value    string
	Children []*ASTNode
}
