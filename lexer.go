package main

type Lexer struct {
	line     int
	startPos int
	pos      int
	ch       byte
	input    []byte
}

func NewLexer(input []byte) *Lexer {
	return &Lexer{
		line:     1,
		ch:       input[0],
		startPos: 0,
		pos:      0,
		input:    input,
	}
}
func (l *Lexer) skip() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\f' {
		l.eat()
	}
}
func (l *Lexer) atEnd() bool {
	return l.pos >= len(l.input)
}
func (l *Lexer) eat() {
	if l.pos < len(l.input) {
		l.pos++
		l.ch = l.input[l.pos]
	}
}
func (l *Lexer) lexInteger() string {
	startPos := l.pos
	for !l.atEnd() && l.ch >= '0' && l.ch <= '9' {
		l.eat()
	}
	endPos := l.pos
	return string(l.input[startPos:endPos])
}
func (l *Lexer) Scan() Token {
	l.skip()
	if l.atEnd() {
		return Token{
			kind: Eof,
		}
	}
	switch l.ch {
	case '+':
		l.eat()
		return Token{kind: Plus}
	case '-':
		// l.eat()
		return Token{kind: Minus}
	case '*':
		// l.eat()
		return Token{kind: Star}
	case '/':
		// l.eat()
		return Token{kind: Slash}
	default:
		{
			if l.ch >= '0' && l.ch <= '9' {
				result := l.lexInteger()
				return Token{
					kind:  IntLit,
					value: result,
				}
			}

		}
	}

	return Token{kind: Eof}

}
