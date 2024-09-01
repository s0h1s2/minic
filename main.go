package main

func main() {
	input := "1+2\000"
	lex := NewLexer([]byte(input))
	// lex.Scan()
	parser := NewParser(lex)
	parser.Parse()
	// token := lex.Scan()
	// for token.kind != Eof {
	// 	print(token.String())
	// 	token = lex.Scan()
	// }
	// print(token.String())

}
