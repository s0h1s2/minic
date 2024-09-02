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
		op := p.currentToken
		p.nextToken()
		left = &BinaryExpr{
			Left:  left,
			Right: p.parsePrimary(),
			Kind:  op.kind,
		}
	}
	return left

}

func (p *Parser) parseAddition() Expr {
	left := p.parseMultiply()
	for p.currentToken.kind == Plus || p.currentToken.kind == Minus {
		op := p.currentToken
		p.nextToken()
		left = &BinaryExpr{
			Left:  left,
			Right: p.parseMultiply(),
			Kind:  op.kind,
		}
	}
	return left
}
func (p *Parser) Parse() Expr {
	// Fetch first token
	p.nextToken()
	return p.parseAddition()
}
