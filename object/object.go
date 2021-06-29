// Package object -
package object

import "fmt"

// ObjectType -
// ignore stutter warning
// nolint: revive
type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
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
