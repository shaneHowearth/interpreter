// Package object -
package object

import "fmt"

// ObjectType -
// ignore stutter warning
// nolint: revive
type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
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
