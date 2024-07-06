package parser

import (
	"fmt"

	"github.com/darragh-downey/secure-c/lexer"
)

type Parser struct {
	tokens  []lexer.Token
	current int
}

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) Parse() (*ASTNode, error) {
	root := &ASTNode{Value: "root"}
	for !p.isAtEnd() {
		node, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		root.Children = append(root.Children, node)
	}
	return root, nil
}

func (p *Parser) parseExpression() (*ASTNode, error) {
	token := p.advance()
	node := &ASTNode{Value: token.Lexeme}

	if token.Type == lexer.TOKEN_IDENTIFIER || token.Type == lexer.TOKEN_NUMBER {
		if p.match(lexer.TOKEN_OPERATOR) {
			operator := p.previous()
			right, err := p.parseExpression()
			if err != nil {
				return nil, err
			}
			node = &ASTNode{Value: operator.Lexeme, Children: []*ASTNode{node, right}}
		}
		return node, nil
	}

	return nil, fmt.Errorf("unexpected token: %s", token.Lexeme)
}

func (p *Parser) advance() lexer.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) match(tokenType lexer.TokenType) bool {
	if p.check(tokenType) {
		p.advance()
		return true
	}
	return false
}

func (p *Parser) check(tokenType lexer.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

func (p *Parser) peek() lexer.Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() lexer.Token {
	return p.tokens[p.current-1]
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == lexer.TOKEN_EOF
}
