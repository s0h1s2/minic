package main

type Expr interface {
	expr()
}

type IntegerExpr struct {
	Value int
}
type BinaryExpr struct {
	Kind  TokenKind
	Left  Expr
	Right Expr
}

func (e *IntegerExpr) expr() {}
func (e *BinaryExpr) expr()  {}
