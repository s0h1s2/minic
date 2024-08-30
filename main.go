package main

func main() {
	input := "1+2"
	lex := NewLexer([]byte(input))
	print(lex.Scan().String())
}
