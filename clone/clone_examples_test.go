package clone_test

import (
	"fmt"

	"github.com/go-extras/go-kit/clone"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) Clone() *Person {
	return &Person{Name: p.Name, Age: p.Age}
}

func ExampleMap() {
	// Create a map of Person values
	m1 := map[string]*Person{
		"person1": {Name: "Alice", Age: 30},
		"person2": {Name: "Bob", Age: 40},
	}

	// Clone the map
	m2 := clone.Map(m1)

	// Update the age of the Person value in m1
	p1 := m1["person1"]
	p1.Age = 31

	// Print the original and cloned maps
	fmt.Printf("m1[\"person1\"]=%+v\n", m1["person1"])
	fmt.Printf("m1[\"person2\"]=%+v\n", m1["person2"])
	fmt.Printf("m2[\"person1\"]=%+v\n", m2["person1"])
	fmt.Printf("m2[\"person2\"]=%+v\n", m2["person2"])

	// Output:
	// m1["person1"]=&{Name:Alice Age:31}
	// m1["person2"]=&{Name:Bob Age:40}
	// m2["person1"]=&{Name:Alice Age:30}
	// m2["person2"]=&{Name:Bob Age:40}
}
