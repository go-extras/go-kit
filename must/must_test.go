package must_test

import (
	"errors"
	"testing"

	qt "github.com/frankban/quicktest"

	"github.com/go-extras/go-kit/must"
)

func TestMust(t *testing.T) {
	c := qt.New(t)

	// Test case 1: no error
	result1 := must.Must[string]("Hello, world!", nil)
	c.Assert(result1, qt.Equals, "Hello, world!")

	// Test case 2: error
	err2 := errors.New("something went wrong")
	c.Assert(func() { must.Must[any](nil, err2) }, qt.PanicMatches, err2.Error())
}
