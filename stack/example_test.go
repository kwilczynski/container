package stack_test

import (
	"fmt"
	"time"

	"github.com/kwilczynski/container/stack"
)

// This examples shows the basic usage of the package: Create a new
// Stack and push two items to it, then retrieve values from the Stack
// as long as the stack is not empty.
func ExampleStack_basic() {
	s := stack.New()
	s.Push("World")
	s.Push("Hello")

	for !s.Empty() {
		value, err := s.Pop()
		if err != nil {
			fmt.Printf("An error occurred: %v\n", err)
			return
		}

		fmt.Println(value)
	}
	// Output:
	// Hello
	// World
}

// This example shows how to use a closure: Search for a particular
// entry in the Stack, and if found display the result and the
// distance from the top of the Stack.
func ExampleStack_search() {
	s := stack.New()

	type Person struct {
		Name        string
		DateOfBirth time.Time // Date of birth of a person.
	}

	var t time.Time
	const dateForm = `2006-01-02`

	t, _ = time.Parse(dateForm, `1879-03-14`)
	s.Push(&Person{"Albert Einstein", t})

	t, _ = time.Parse(dateForm, `1856-07-10`)
	s.Push(&Person{"Nikola Tesla", t})

	t, _ = time.Parse(dateForm, `1887-08-12`)
	s.Push(&Person{"Erwin Schrodinger", t})

	// Note: For better performance it is advisable to re-use
	// closure by defining it first and then passing to Search.
	ok, distance, item := s.Search(func(v interface{}) bool {
		return v.(*Person).Name == "Nikola Tesla"
	})

	if ok {
		fmt.Printf("Found at the distance of %d from the top: %+v\n", distance, *item.(*Person))
		return
	}

	fmt.Println("Nothing found.")
	// Output:
	// Found at the distance of 2 from the top: {Name:Nikola Tesla DateOfBirth:1856-07-10 00:00:00 +0000 UTC}
}
