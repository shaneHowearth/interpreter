// Package ast -
package ast

import (
	"bytes"

	"github.com/shanehowearth/interpreter/token"
)

// Node -
type Node interface {
	TokenLiteral() string
	String() string
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

// String -
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

// String -
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
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

// String -
func (i *Identifier) String() string {
	return i.Value
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

// String -
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// ExpressionStatement -
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral -
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String -
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
