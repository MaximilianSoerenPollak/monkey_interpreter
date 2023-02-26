package parser

import (
	"monkey_interpreter/ast"
	"monkey_interpreter/token"
	"monkey_interpreter/lexer"
)

type Parser struct {
	l *lexer.Lexer //Pointer to instance of the lexer
	
	curToken token.Token
	peekToken token.Token 
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so curToken and peekToken can both be set.
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil 
}

// we read LET -> then advance token reader 
// -> check if next token is '='
// -> advance token reader to after = sign and parse that.

// ?? How do we know in 'parseOperatorExpression'  that the first literal will be an int.??
