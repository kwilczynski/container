package stack_test

import (
	"fmt"
	"github.com/kwilczynski/container/stack"
)

func ExampleStack() {
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

func ExampleStack_Search() {
	s := stack.New()

	type Person struct {
		Name string
		Year int // Year in which a Person was born.
	}

	s.Push(&Person{"Albert Einstein", 1879})
	s.Push(&Person{"Nikola Tesla", 1856})
	s.Push(&Person{"Erwin Schrodinger", 1887})

	found, distance := s.Search(func(v interface{}) bool {
		return v.(*Person).Name == "Nikola Tesla"
	})

	if found {
		fmt.Printf("Found at the distance of %d from the top.\n", distance)
	} else {
		fmt.Println("Nothing found.")
	}
	// Output:
	// Found at the distance of 2 from the top.
}
