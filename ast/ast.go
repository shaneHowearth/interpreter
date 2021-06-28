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

// IntegerLiteral -
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

// TokenLiteral -
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) expressionNode() {}

// String -
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// PrefixExpression -
type PrefixExpression struct {
	Token    token.Token // The prefix token, e.g. !
	Operator string
	Right    Expression
}

// TokenLiteral -
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

func (pe *PrefixExpression) expressionNode() {}

// String -
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")
	return out.String()

}

// InfixExpression -
type InfixExpression struct {
	Token    token.Token // The operator token, eg. +
	Left     Expression
	Operator string
	Right    Expression
}

func (oe *InfixExpression) expressionNode() {}

// TokenLiteral -
func (oe *InfixExpression) TokenLiteral() string { return oe.Token.Literal }

// String -
func (oe *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

// Boolean -
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}

// TokenLiteral -
func (b *Boolean) TokenLiteral() string { return b.Token.Literal }

// String -
func (b *Boolean) String() string { return b.Token.Literal }

// IfExpression -
type IfExpression struct {
	Token       token.Token // The 'if' token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}

// TokenLiteral -
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }

// String -
func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}
	return out.String()
}

// BlockStatement -
type BlockStatement struct {
	Token      token.Token // the { token
	Statements []Statement
}

func (bs *BlockStatement) expressionNode() {}

// TokenLiteral -
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }

// String -
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
