package main

import (
	"fmt"
	"strings"
)

func traverseExpr(expr Expr, indent int) {

	switch e := expr.(type) {
	case *IntegerExpr:
		{
			println(e.Value)
		}
	case *BinaryExpr:
		{
			indent += 2
			fmt.Printf("%s\n", e.Kind.String())
			fmt.Printf("\\"strings.Repeat("", indent))
			fmt.Printf("/"strings.Repeat("", indent))
			indent -= 2
		}
	}
}
func main() {
	input := "1+2\000"
	lex := NewLexer([]byte(input))
	// lex.Scan()
	parser := NewParser(lex)
	expr := parser.Parse()
	traverseExpr(expr, 0)
	// token := lex.Scan()
	// for token.kind != Eof {
	// 	print(token.String())
	// 	token = lex.Scan()
	// }
	// print(token.String())

}
