package token

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, . . .
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS = "+"

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
)


type TokenType string

type Token struct {
	Type TokenType
	Literal string
}