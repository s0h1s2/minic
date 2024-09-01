package main

import "fmt"

type TokenKind int

const (
	Start TokenKind = iota
	Plus
	Minus
	Star
	Slash
	IntLit
	Eof
)
const tokenRepresentationArraySize = (Eof - Start) + 1

var tokenRepresentation = [tokenRepresentationArraySize]string{
	Plus:   "+",
	Minus:  "-",
	Slash:  "/",
	Star:   "*",
	IntLit: "INT",
	Eof:    "EOF",
}

type Token struct {
	kind  TokenKind
	value string
}

func (kind TokenKind) String() string {
	return tokenRepresentation[kind]
}
func (t Token) String() string {
	return fmt.Sprintf("Token(kind='%s',value='%s')\n", t.kind.String(), t.value)
}
