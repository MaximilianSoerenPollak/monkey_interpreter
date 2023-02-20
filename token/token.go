package token 

type TokenType string 

type Token struct {
	Type TokenType
	Literal string 
}



const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifiers + literals 
	IDENT = "INDENT" // add, foo, bar, x, y
	INT = "INT" // 12313123

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

	// Keywords
	FUNCTION = "FUNCTION"
	RETURN = "RETURN"
	LET = "LET"
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok{
		return tok
	}
	return IDENT
}
