package main

type Lexer struct {
	line  int
	pos   int
	ch    byte
	input []byte
}

func NewLexer(input []byte) *Lexer {
	return &Lexer{
		line:  1,
		ch:    0,
		pos:   0,
		input: input,
	}
}
func (l *Lexer) skip() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\f' {
		l.pos++
	}
}
func (l *Lexer) atEnd() bool {
	return l.pos >= len(l.input)
}
func (l *Lexer) Scan() Token {
	l.skip()
	if l.atEnd() {
		return Token{
			kind: Eof,
		}
	}
	l.ch = l.input[l.pos]
	switch l.ch {
	case '+':
		return Token{kind: Plus}
	case '-':
		return Token{kind: Minus}
	case '*':
		return Token{kind: Star}
	case '/':
		return Token{kind: Slash}
	default:
		{
			if l.ch >= '0' && l.ch <= '9' {
				return Token{
					kind:  IntLit,
					value: string(l.ch),
				}
			}
		}
	}

	return Token{kind: Eof}

}
