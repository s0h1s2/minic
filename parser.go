package main

import "fmt"

type Parser struct {
	currentToken Token
	lex          *Lexer
}

func NewParser(lex *Lexer) *Parser {
	return &Parser{}
}
func (p *Parser) nextToken() {
	p.currentToken = p.lex.Scan()
}
func (p *Parser) parsePrimary() Expr {
	if p.currentToken.kind == IntLit {
		return &IntegerExpr{
			Value: 1,
		}
	}
	panic(fmt.Sprintf("Syntax Error:%s",p.currentToken.))
}
func (p *Parser) parseMultiply() Expr {
	left := p.parsePrimary()
  
	return left
}

func (p *Parser) parseAddition() Expr {

}
func (p *Parser) Parse() Expr {
	return p.parseAddition()
}
