// Package ast -
package ast

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
