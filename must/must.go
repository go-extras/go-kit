// Package must offers a convenient approach for transforming a two-value function
// into a single-value function by throwing a panic if an error is returned as the second value
// in the original function.
//
// It should be used sparingly and only in cases where it's unlikely for an error to occur
// or where the error handling would be identical in every case.
//
// License: MIT
// Copyright: 2023, Denis Voytyuk
package must

// Assert takes an error err. If err is not nil, Assert panics with the error message.
func Assert(err error) {
	if err != nil {
		panic(err)
	}
}

// Must takes a value v of any type T and an error err. If err is not nil, Must panics
// with the error message. Otherwise, it returns v.
//
// Example usage:
//
//	f, err := os.Open("file.txt")
//	content := make([]byte, 100)
//	n := must.Must(f.Read(content))
//	// handle the content
//	defer f.Close()
//
// This function should be used judiciously, as panics can have unintended consequences and
// make debugging more difficult. Use it only in cases where it simplifies the code and the
// error handling is the same in every case.
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// Must1 is an exact copy of Must, but with a different name.
func Must1[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// Must2 is the same as Must, but for functions that return two values.
func Must2[T any, U any](v1 T, v2 U, err error) (T, U) {
	if err != nil {
		panic(err)
	}
	return v1, v2
}

// Must3 is the same as Must, but for functions that return three values.
func Must3[T any, U any, V any](v1 T, v2 U, v3 V, err error) (T, U, V) {
	if err != nil {
		panic(err)
	}
	return v1, v2, v3
}
