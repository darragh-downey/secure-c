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
		node, err := p.parseDeclaration()
		if err != nil {
			return nil, err
		}
		root.Children = append(root.Children, node)
	}
	return root, nil
}

func (p *Parser) parseDeclaration() (*ASTNode, error) {
	if p.match(lexer.TOKEN_KEYWORD) {
		keyword := p.previous()
		identifier, err := p.consume(lexer.TOKEN_IDENTIFIER, "expected identifier after keyword")
		if err != nil {
			return nil, err
		}
		if _, err := p.consume(lexer.TOKEN_SEPARATOR, "expected ';' after declaration"); err != nil {
			return nil, err
		}
		return &ASTNode{Value: keyword.Lexeme, Children: []*ASTNode{
			{Value: identifier.Lexeme},
		}}, nil
	}
	return nil, p.error(p.peek(), "expected declaration")
}

func (p *Parser) match(tokenTypes ...lexer.TokenType) bool {
	for _, tokenType := range tokenTypes {
		if p.check(tokenType) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) check(tokenType lexer.TokenType) bool {
	if p.isAtEnd() {
		return false
	}
	return p.peek().Type == tokenType
}

func (p *Parser) advance() lexer.Token {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == lexer.TOKEN_EOF
}

func (p *Parser) peek() lexer.Token {
	return p.tokens[p.current]
}

func (p *Parser) previous() lexer.Token {
	return p.tokens[p.current-1]
}

func (p *Parser) consume(tokenType lexer.TokenType, message string) (lexer.Token, error) {
	if p.check(tokenType) {
		return p.advance(), nil
	}
	return lexer.Token{}, p.error(p.peek(), message)
}

func (p *Parser) error(token lexer.Token, message string) error {
	return fmt.Errorf("[line %d] Error at '%s': %s", token.Line, token.Lexeme, message)
}
