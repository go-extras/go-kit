// Package ptr provides utility functions for working with pointers in Go.
// It includes functions to convert values to pointers, dereference pointers,
// and handle nil pointers gracefully.
//
// This package is useful for simplifying code that deals with pointers,
// especially when working with optional values or when interfacing with
// APIs that expect pointers.
//
// License: MIT
// Copyright: 2025, Denis Voytyuk
package ptr

// To pointerizes the given value.
func To[T any](v T) *T {
	return &v
}

// From dereferences the given pointer.
func From[T any](v *T) (ret T) {
	if v == nil {
		return ret
	}
	return *v
}

// FromOrDefault dereferences the given pointer or returns the default value if the pointer is nil.
func FromOrDefault[T any](v *T, defaultValue T) T {
	if v == nil {
		return defaultValue
	}
	return *v
}
