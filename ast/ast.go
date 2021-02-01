// Package ast -
package ast

import "github.com/shanehowearth/interpreter/token"

// Node -
type Node interface {
	TokenLiteral() string
}

// Statement - statements do not produce values
// let x = 5 assigns a value, but does not produce one
// return; does not produce a value
type Statement interface {
	Node
	statementNode()
}

// Expression - expressions produce values
// 5 produces the value 5
// ass(5, 5) produces another value
type Expression interface {
	Node
	expressionNode()
}

// Program - this will be the Root node of the AST
type Program struct {
	Statements []Statement
}

// TokenLiteral -
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement -
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral -
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier -
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral -
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// ReturnStatement -
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral -
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
