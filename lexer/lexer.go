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

	l.skipWhiteSpace()

	switch l.ch {
	case 0:
		t.Type = token.EOF
		t.Literal = ""
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			t.Type = token.EQ
			t.Literal = string(ch) + string(l.ch)
		} else {
			t.Type = token.ASSIGN
			t.Literal = "="
		}
	case '+':
		t.Type = token.PLUS
		t.Literal = "+"
	case '-':
		t.Type = token.MINUS
		t.Literal = "-"
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			t.Type = token.NOT_EQ
			t.Literal = string(ch) + string(l.ch)
		} else {
			t.Type = token.BANG
			t.Literal = "!"
		}
	case '/':
		t.Type = token.SLASH
		t.Literal = "/"
	case '*':
		t.Type = token.ASTERISK
		t.Literal = "*"
	case '<':
		t.Type = token.LT
		t.Literal = "<"
	case '>':
		t.Type = token.GT
		t.Literal = ">"
	case ',':
		t.Type = token.COMMA
		t.Literal = ","
	case ';':
		t.Type = token.SEMICOLON
		t.Literal = ";"
	case '(':
		t.Type = token.LPAREN
		t.Literal = "("
	case ')':
		t.Type = token.RPAREN
		t.Literal = ")"
	case '{':
		t.Type = token.LBRACE
		t.Literal = "{"
	case '}':
		t.Type = token.RBRACE
		t.Literal = "}"
	case '"':
		t.Type = token.STRING
		t.Literal = l.readString()
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier(isLetter)
			t.Type = token.LookupIdent(t.Literal)
			return t
		} else if isDigit(l.ch) {
			t.Type = token.INT
			t.Literal = l.readIdentifier(isDigit)
			return t
		} else {
			t = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return t
}

func (l *Lexer) readIdentifier(f func(byte) bool) string {
	pos := l.position
	for f(l.ch) {
		l.readChar()
	}
	return l.input[pos:l.position]
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

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}
