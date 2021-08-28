// Package object -
package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/shanehowearth/interpreter/ast"
)

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
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
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

// Function -
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type -
func (f *Function) Type() ObjectType { return FUNCTION_OBJ }

// Inspect -
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())

	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

// String -
type String struct {
	Value string
}

// Type -
func (s *String) Type() ObjectType { return STRING_OBJ }

// Inspect -
func (s *String) Inspect() string { return s.Value }

// BuiltinFunction -
type BuiltinFunction func(args ...Object) Object

// Builtin -
type Builtin struct {
	Fn BuiltinFunction
}

// Type -
func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }

// Inspect -
func (b *Builtin) Inspect() string { return "builtin function" }

// Array -
type Array struct {
	Elements []Object
}

// Type -
func (ao *Array) Type() ObjectType { return ARRAY_OBJ }

// Inspect -
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
