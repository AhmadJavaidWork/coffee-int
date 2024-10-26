package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	SLASH    = "/"
	ASTERISK = "*"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}
