package must_test

import (
	"errors"
	"fmt"

	"github.com/go-extras/go-kit/must"
)

func ExampleMust() {
	// Example 1: Successful operation
	result1 := must.Must(add(2, 3))
	fmt.Println(result1)

	// Example 2: Unsuccessful operation
	func() {
		// Example 2: Unsuccessful operation
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("panic:", r)
			}
		}()
		result2 := must.Must(divide(4, 0))
		fmt.Println(result2)
	}()

	// Output:
	// 5
	// panic: cannot divide by zero
}

func add(a, b int) (int, error) {
	return a + b, nil
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}
