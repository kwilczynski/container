/*
 * example_test.go
 *
 * Copyright 2013-2016 Krzysztof Wilczynski
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package stack_test

import (
	"fmt"

	"github.com/kwilczynski/container/stack"
)

func Example_Basic() {
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

func Example_Search() {
	s := stack.New()

	type Person struct {
		Name string
		Year int // Year in which a Person was born.
	}

	s.Push(&Person{"Albert Einstein", 1879})
	s.Push(&Person{"Nikola Tesla", 1856})
	s.Push(&Person{"Erwin Schrodinger", 1887})

	// Note: For better performance it is advisable to re-use
	// closure by defining it first and then passing to Search.
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
