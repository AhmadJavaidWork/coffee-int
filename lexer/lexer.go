package lexer

import "github.com/ahmadjavaidwork/coffee-int/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	switch l.ch {
	case 0:
		t.Type = token.EOF
		t.Literal = ""
	case '=':
		t.Type = token.ASSIGN
		t.Literal = "="
	case '+':
		t.Type = token.PLUS
		t.Literal = "+"
	case '-':
		t.Type = token.MINUS
		t.Literal = "-"
	case '!':
		t.Type = token.BANG
		t.Literal = "!"
	case '/':
		t.Type = token.SLASH
		t.Literal = "/"
	case '*':
		t.Type = token.ASTERISK
		t.Literal = "*"
	default:
		t = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return t
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}
