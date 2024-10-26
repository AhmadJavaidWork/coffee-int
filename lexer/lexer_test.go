package lexer

import (
	"testing"

	"github.com/ahmadjavaidwork/coffee-int/token"
)

func TestNextToken(t *testing.T) {
	input := `=+-!/* `
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.BANG, "!"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.ILLEGAL, " "},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
