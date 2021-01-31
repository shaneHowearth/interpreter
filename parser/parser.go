// Package parser -
package parser

import (
	"github.com/shanehowearth/interpreter/ast"
	"github.com/shanehowearth/interpreter/lexer"
	"github.com/shanehowearth/interpreter/token"
)

// Parser -
type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

// New -
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens, so both curToken and peekToken are set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram -
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
