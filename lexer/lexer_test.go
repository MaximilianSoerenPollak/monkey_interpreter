package lexer 

import (
	"testing"

	"monkey_interpreter/token"

)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		excpectedType	token.TokenType
		excpectedLiteral	string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		
		if tok.Type != tt.excpectedType {
			t.Fatalf("tests[%d] - tokentype wrong. excpected=%q, got=%q", i, tt.excpectedType, tok.Type)
		}
		if tok.Literal != tt.excpectedLiteral {
			t.Fatalf("tests[%d] - tokentype wrong. excpected=%q, got=%q", i, tt.excpectedLiteral, tok.Literal)
		}
	}
}
