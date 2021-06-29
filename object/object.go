// Package object -
package object

// ObjectType -
// ignore stutter warning
// nolint: revive
type ObjectType string

// Object -
type Object interface {
	Type() ObjectType
	Inspect() string
}
