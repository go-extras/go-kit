// Package clone provides functionality to clone data.
//
// License: MIT
// Copyright: 2023, Denis Voytyuk
package clone

// Clonable is an interface for types that can be cloned.
type Clonable[T any] interface {
	Clone() T
}

// Map returns a shallow copy of the given map.
// If the map is nil, it returns nil.
// If the map contains values that implement Clonable, Clone will invoke their
// Clone method and use the returned value in the new map. Otherwise,
// the original value will be used.
func Map[K comparable, V any](t map[K]V) map[K]V {
	if t == nil {
		return nil
	}

	// Create a new map with the same type as the original map
	n := make(map[K]V, len(t))

	// Iterate over the original map
	for k, v := range t {
		// Check if the value implements the Clonable interface
		switch v1 := (any)(v).(type) {
		case Clonable[V]:
			// If it does, clone the value and use the clone in the new map
			n[k] = v1.Clone()
		default:
			// If it doesn't, use the original value in the new map
			n[k] = v
		}
	}

	return n
}
