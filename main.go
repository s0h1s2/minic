package main

func main() {
	input := "11+2\000"
	lex := NewLexer([]byte(input))
	token := lex.Scan()
	for token.kind != Eof {
		print(token.String())
		token = lex.Scan()
	}
	print(token.String())
}
