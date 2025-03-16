package ptr_test

import (
	"reflect"
	"testing"

	"github.com/go-extras/go-kit/ptr"
)

func TestTo(t *testing.T) {
	type testcase[T any] struct {
		name string
		arg  T
		want *T
	}
	type testscases[T any] []testcase[T]

	t.Run("int", func(t *testing.T) {
		tests := testscases[int]{
			{
				name: "test1",
				arg:  42,
				want: (func() *int {
					v := 42
					return &v
				})(),
			},
			{
				name: "test2",
				arg:  -1,
				want: (func() *int {
					v := -1
					return &v
				})(),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.To(tt.arg); *got != *tt.want {
					t.Errorf("To() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("string", func(t *testing.T) {
		tests := testscases[string]{
			{
				name: "test1",
				arg:  "hello",
				want: (func() *string {
					v := "hello"
					return &v
				})(),
			},
			{
				name: "test2",
				arg:  "",
				want: (func() *string {
					v := ""
					return &v
				})(),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.To(tt.arg); *got != *tt.want {
					t.Errorf("To() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("bool", func(t *testing.T) {
		tests := testscases[bool]{
			{
				name: "test1",
				arg:  true,
				want: (func() *bool {
					v := true
					return &v
				})(),
			},
			{
				name: "test2",
				arg:  false,
				want: (func() *bool {
					v := false
					return &v
				})(),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.To(tt.arg); *got != *tt.want {
					t.Errorf("To() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("float64", func(t *testing.T) {
		tests := testscases[float64]{
			{
				name: "test1",
				arg:  42.0,
				want: (func() *float64 {
					v := 42.0
					return &v
				})(),
			},
			{
				name: "test2",
				arg:  0.0,
				want: (func() *float64 {
					v := 0.0
					return &v
				})(),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.To(tt.arg); *got != *tt.want {
					t.Errorf("To() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("slice", func(t *testing.T) {
		tests := testscases[[]int]{
			{
				name: "test1",
				arg:  []int{1, 2, 3},
				want: (func() *[]int {
					v := []int{1, 2, 3}
					return &v
				})(),
			},
			{
				name: "test2",
				arg:  nil,
				want: (func() *[]int {
					v := []int(nil)
					return &v
				})(),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.To(tt.arg); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("To() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("map", func(t *testing.T) {
		tests := testscases[map[string]int]{
			{
				name: "test1",
				arg:  map[string]int{"a": 1, "b": 2},
				want: (func() *map[string]int {
					v := map[string]int{"a": 1, "b": 2}
					return &v
				})(),
			},
			{
				name: "test2",
				arg:  nil,
				want: (func() *map[string]int {
					v := map[string]int(nil)
					return &v
				})(),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.To(tt.arg); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("To() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("struct", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		// Test case for struct
		tests := testscases[Person]{
			{
				name: "test1",
				arg:  Person{Name: "Alice", Age: 25},
				want: (func() *Person {
					v := Person{Name: "Alice", Age: 25}
					return &v
				})(),
			},
			{
				name: "test2",
				arg:  Person{},
				want: (func() *Person {
					v := Person{}
					return &v
				})(),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.To(tt.arg); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("To() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("empty struct", func(t *testing.T) {
		tests := testscases[struct{}]{
			{
				name: "empty",
				arg:  struct{}{},
				want: (func() *struct{} {
					v := struct{}{}
					return &v
				})(),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.To(tt.arg); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("To() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func TestFrom(t *testing.T) {
	type testcase[T any] struct {
		name string
		arg  *T
		want T
	}
	type testscases[T any] []testcase[T]

	t.Run("int", func(t *testing.T) {
		tests := testscases[int]{
			{
				name: "test1",
				arg:  (func() *int { v := 42; return &v })(),
				want: 42,
			},
			{
				name: "test2",
				arg:  nil,
				want: 0,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.From(tt.arg); got != tt.want {
					t.Errorf("From() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("string", func(t *testing.T) {
		tests := testscases[string]{
			{
				name: "test1",
				arg:  (func() *string { v := "hello"; return &v })(),
				want: "hello",
			},
			{
				name: "test2",
				arg:  nil,
				want: "",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.From(tt.arg); got != tt.want {
					t.Errorf("From() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("bool", func(t *testing.T) {
		tests := testscases[bool]{
			{
				name: "test1",
				arg:  (func() *bool { v := true; return &v })(),
				want: true,
			},
			{
				name: "test2",
				arg:  nil,
				want: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.From(tt.arg); got != tt.want {
					t.Errorf("From() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("float64", func(t *testing.T) {
		tests := testscases[float64]{
			{
				name: "test1",
				arg:  (func() *float64 { v := 42.0; return &v })(),
				want: 42.0,
			},
			{
				name: "test2",
				arg:  nil,
				want: 0.0,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.From(tt.arg); got != tt.want {
					t.Errorf("From() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("slice", func(t *testing.T) {
		tests := testscases[[]int]{
			{
				name: "test1",
				arg:  (func() *[]int { v := []int{1, 2, 3}; return &v })(),
				want: []int{1, 2, 3},
			},
			{
				name: "test2",
				arg:  nil,
				want: nil,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.From(tt.arg); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("From() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("map", func(t *testing.T) {
		tests := testscases[map[string]int]{
			{
				name: "test1",
				arg:  (func() *map[string]int { v := map[string]int{"a": 1, "b": 2}; return &v })(),
				want: map[string]int{"a": 1, "b": 2},
			},
			{
				name: "test2",
				arg:  nil,
				want: nil,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.From(tt.arg); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("From() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("struct", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		tests := testscases[Person]{
			{
				name: "test1",
				arg:  (func() *Person { v := Person{Name: "Alice", Age: 25}; return &v })(),
				want: Person{Name: "Alice", Age: 25},
			},
			{
				name: "test2",
				arg:  nil,
				want: Person{},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.From(tt.arg); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("From() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("nested pointer", func(t *testing.T) {
		v := 42
		ptr1 := &v
		ptr2 := &ptr1

		if got := ptr.From(ptr2); !reflect.DeepEqual(got, ptr1) {
			t.Errorf("From() = %v, want %v", got, 42)
		}
	})
}

func TestFromPtrOrDefault(t *testing.T) {
	type testcase[T any] struct {
		name         string
		arg          *T
		defaultValue T
		want         T
	}
	type testscases[T any] []testcase[T]

	t.Run("int", func(t *testing.T) {
		tests := testscases[int]{
			{
				name:         "test1",
				arg:          (func() *int { v := 42; return &v })(),
				defaultValue: 0,
				want:         42,
			},
			{
				name:         "test2",
				arg:          nil,
				defaultValue: 100,
				want:         100,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.FromOrDefault(tt.arg, tt.defaultValue); got != tt.want {
					t.Errorf("FromOrDefault() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("string", func(t *testing.T) {
		tests := testscases[string]{
			{
				name:         "test1",
				arg:          (func() *string { v := "hello"; return &v })(),
				defaultValue: "",
				want:         "hello",
			},
			{
				name:         "test2",
				arg:          nil,
				defaultValue: "default",
				want:         "default",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.FromOrDefault(tt.arg, tt.defaultValue); got != tt.want {
					t.Errorf("FromOrDefault() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("bool", func(t *testing.T) {
		tests := testscases[bool]{
			{
				name:         "test1",
				arg:          (func() *bool { v := true; return &v })(),
				defaultValue: false,
				want:         true,
			},
			{
				name:         "test2",
				arg:          nil,
				defaultValue: true,
				want:         true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.FromOrDefault(tt.arg, tt.defaultValue); got != tt.want {
					t.Errorf("FromOrDefault() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("float64", func(t *testing.T) {
		tests := testscases[float64]{
			{
				name:         "test1",
				arg:          (func() *float64 { v := 42.0; return &v })(),
				defaultValue: 0.0,
				want:         42.0,
			},
			{
				name:         "test2",
				arg:          nil,
				defaultValue: 3.14,
				want:         3.14,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.FromOrDefault(tt.arg, tt.defaultValue); got != tt.want {
					t.Errorf("FromOrDefault() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("slice", func(t *testing.T) {
		tests := testscases[[]int]{
			{
				name:         "test1",
				arg:          (func() *[]int { v := []int{1, 2, 3}; return &v })(),
				defaultValue: []int{0},
				want:         []int{1, 2, 3},
			},
			{
				name:         "test2",
				arg:          nil,
				defaultValue: []int{0},
				want:         []int{0},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.FromOrDefault(tt.arg, tt.defaultValue); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("FromOrDefault() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("map", func(t *testing.T) {
		tests := testscases[map[string]int]{
			{
				name:         "test1",
				arg:          (func() *map[string]int { v := map[string]int{"a": 1, "b": 2}; return &v })(),
				defaultValue: map[string]int{"default": 0},
				want:         map[string]int{"a": 1, "b": 2},
			},
			{
				name:         "test2",
				arg:          nil,
				defaultValue: map[string]int{"default": 0},
				want:         map[string]int{"default": 0},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.FromOrDefault(tt.arg, tt.defaultValue); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("FromOrDefault() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("struct", func(t *testing.T) {
		type Person struct {
			Name string
			Age  int
		}

		tests := testscases[Person]{
			{
				name:         "test1",
				arg:          (func() *Person { v := Person{Name: "Alice", Age: 25}; return &v })(),
				defaultValue: Person{Name: "Bob", Age: 30},
				want:         Person{Name: "Alice", Age: 25},
			},
			{
				name:         "test2",
				arg:          nil,
				defaultValue: Person{Name: "Bob", Age: 30},
				want:         Person{Name: "Bob", Age: 30},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ptr.FromOrDefault(tt.arg, tt.defaultValue); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("FromOrDefault() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("nested pointer", func(t *testing.T) {
		v := 42
		ptr1 := &v
		ptr2 := (**int)(nil)

		if got := ptr.FromOrDefault(ptr2, ptr1); !reflect.DeepEqual(got, ptr1) {
			t.Errorf("FromOrDefault() = %v, want %v", got, 42)
		}
	})
}
