package lexer

import (
	"testing"

	"github.com/shanehowearth/interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := map[string]struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		"Assign":    {token.ASSIGN, "="},
		"Plus":      {token.PLUS, "+"},
		"LParen":    {token.LPAREN, "("},
		"RParen":    {token.RPAREN, ")"},
		"LBrace":    {token.LBRACE, "{"},
		"RBrace":    {token.RBRACE, "}"},
		"Comma":     {token.COMMA, ","},
		"Semicolon": {token.SEMICOLON, ";"},
		"EOF":       {token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("%s - TokenType wrong, expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("%s - Literal wrong, expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
