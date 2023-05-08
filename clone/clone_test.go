package clone_test

import (
	"testing"

	qt "github.com/frankban/quicktest"

	"github.com/go-extras/go-kit/clone"
)

type myStruct struct {
	ID   int
	Name string
}

func (s *myStruct) Clone() *myStruct {
	return &myStruct{
		ID:   s.ID,
		Name: s.Name,
	}
}

func TestCloneMap(t *testing.T) {
	c := qt.New(t)

	var nilMap map[string]int

	// Test that passing a nil map returns nil.
	c.Assert(clone.Map(nilMap), qt.IsNil)

	// Test that a map with non-Clonable values is cloned correctly.
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := clone.Map(m1)
	c.Assert(m2, qt.DeepEquals, m1)

	// Test that a map with Clonable values is cloned correctly.
	m3 := map[string]*myStruct{
		"a": {ID: 1, Name: "Alice"},
		"b": {ID: 2, Name: "Bob"},
	}
	m4 := clone.Map(m3)
	c.Assert(m4["a"], qt.Not(qt.Equals), m3["a"])
	c.Assert(m4["a"].ID, qt.Equals, m3["a"].ID)
	c.Assert(m4["a"].Name, qt.Equals, m3["a"].Name)
	c.Assert(m4["b"], qt.Not(qt.Equals), m3["b"])
	c.Assert(m4["b"].ID, qt.Equals, m3["b"].ID)
	c.Assert(m4["b"].Name, qt.Equals, m3["b"].Name)
}
