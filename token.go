package main

import "fmt"

type TokenKind int

const (
	Plus TokenKind = iota
	Minus
	Star
	Slash
	Int
)

type Token struct {
	kind         TokenKind
	integerValue uint
}

func (t Token) String() string {
	return fmt.Sprintf("(%d,%d)", t.kind, t.integerValue)
}
