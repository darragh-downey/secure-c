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
		node, err := p.parseDeclarationOrStatement()
		if err != nil {
			return nil, err
		}
		root.Children = append(root.Children, node)
	}
	return root, nil
}

func (p *Parser) parseDeclarationOrStatement() (*ASTNode, error) {
	if p.match(lexer.TOKEN_KEYWORD) {
		keyword := p.previous()
		if keyword.Lexeme == "int" || keyword.Lexeme == "char" || keyword.Lexeme == "void" {
			return p.parseFunctionOrVariableDeclaration(keyword)
		}
	}
	return nil, p.error(p.peek(), "expected declaration or statement")
}

func (p *Parser) parseFunctionOrVariableDeclaration(keyword lexer.Token) (*ASTNode, error) {
	identifier, err := p.consume(lexer.TOKEN_IDENTIFIER, "expected identifier")
	if err != nil {
		return nil, err
	}

	if p.match(lexer.TOKEN_SEPARATOR) && p.previous().Lexeme == ";" {
		return &ASTNode{
			Value: "variable_declaration",
			Children: []*ASTNode{
				{Value: keyword.Lexeme},
				{Value: identifier.Lexeme},
			}}, nil
	}

	if p.match(lexer.TOKEN_SEPARATOR) && p.previous().Lexeme == "(" {
		params, err := p.parseParameters()
		if err != nil {
			return nil, err
		}

		if _, err := p.consume(lexer.TOKEN_SEPARATOR, "expected ')'"); err != nil {
			return nil, err
		}

		body, err := p.parseBlock()
		if err != nil {
			return nil, err
		}

		return &ASTNode{
			Value: "function_declaration",
			Children: []*ASTNode{
				{Value: keyword.Lexeme},
				{Value: identifier.Lexeme},
				params,
				body,
			}}, nil
	}

	return nil, p.error(p.peek(), "expected ';' or '(' after identifier")
}

func (p *Parser) parseBlock() (*ASTNode, error) {
	if _, err := p.consume(lexer.TOKEN_SEPARATOR, "expected '{'"); err != nil {
		return nil, err
	}

	block := &ASTNode{Value: "block"}
	for !p.isAtEnd() && p.peek().Lexeme != "}" {
		node, err := p.parseDeclarationOrStatement()
		if err != nil {
			return nil, err
		}
		block.Children = append(block.Children, node)
	}

	if _, err := p.consume(lexer.TOKEN_SEPARATOR, "expected '}'"); err != nil {
		return nil, err
	}

	return block, nil
}

func (p *Parser) parseParameters() (*ASTNode, error) {
	params := &ASTNode{Value: "parameters"}
	
	for !p.isAtEnd() && p.peek().Lexeme != ")" {
		keyword, err := p.consume(lexer.TOKEN_KEYWORD, "expected type")
		if err != nil {
			return nil, err
		}

		identifier, err := p.consume(lexer.TOKEN_IDENTIFIER, "expected identifier")
		if err != nil {
			return nil, err
		}

		params.Children = append(params.Children, &ASTNode{
			Value: "parameter",
			Children: []*ASTNode{
				{Value: keyword.Lexeme},
				{Value: identifier.Lexeme},
			},
		})

		if p.match(lexer.TOKEN_SEPARATOR) && p.previous().Lexeme == "," {
			continue
		}

		break
	}
	return params, nil
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
