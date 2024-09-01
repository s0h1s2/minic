package main

import "fmt"

type Parser struct {
	currentToken Token
	lex          *Lexer
}

func NewParser(lex *Lexer) *Parser {
	return &Parser{
		lex: lex,
	}
}
func (p *Parser) nextToken() {
	p.currentToken = p.lex.Scan()
}
func (p *Parser) parsePrimary() Expr {
	if p.currentToken.kind == IntLit {
		result := &IntegerExpr{
			Value: p.currentToken.value,
		}
		p.nextToken()
		return result
	}
	panic(fmt.Sprintf("Syntax Error:%s", p.currentToken.kind.String()))
}
func (p *Parser) parseMultiply() Expr {
	left := p.parsePrimary()
	for p.currentToken.kind == Star || p.currentToken.kind == Slash {
		p.nextToken()
		left = p.parsePrimary()
	}
	return left

}

func (p *Parser) parseAddition() Expr {
	left := p.parseMultiply()
	for p.currentToken.kind == Star || p.currentToken.kind == Slash {
		p.nextToken()
		left = p.parseMultiply()
	}
	return left

	return p.parseMultiply()
}
func (p *Parser) Parse() Expr {
	p.nextToken()
	return p.parseAddition()
}
