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
