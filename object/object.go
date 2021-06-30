// Package object -
package object

import "fmt"

// ObjectType -
// ignore stutter warning
// nolint: revive
type ObjectType string

// nolint: revive
const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
)

// Object -
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer -
type Integer struct {
	Value int64
}

// Inspect -
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Type -
func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

// Boolean -
type Boolean struct {
	Value bool
}

// Type -
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

// Inspect -
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Null -
type Null struct{}

// Type -
func (n *Null) Type() ObjectType { return NULL_OBJ }

// Inspect -
func (n *Null) Inspect() string { return "null" }

// ReturnValue -
type ReturnValue struct {
	Value Object
}

// Type -
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

// Inspect -
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

// Error -
type Error struct {
	Message string
}

// Type -
func (e *Error) Type() ObjectType { return ERROR_OBJ }

// Inspect -
func (e *Error) Inspect() string { return "ERROR: " + e.Message }
