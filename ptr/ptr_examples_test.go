package ptr_test

import (
	"fmt"

	"github.com/go-extras/go-kit/ptr"
)

func ExampleTo() {
	value := 42
	ptrValue := ptr.To(value)
	fmt.Println(*ptrValue)
	// Output: 42
}

func ExampleFrom() {
	ptrValue := ptr.To(42)
	value := ptr.From(ptrValue)
	fmt.Println(value)
	// Output: 42
}

func ExampleFromOrDefault() {
	var ptrValue *int
	value := ptr.FromOrDefault(ptrValue, 100)
	fmt.Println(value)
	// Output: 100
}
