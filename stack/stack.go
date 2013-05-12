/*
 * stack.go
 *
 * Copyright 2013 Krzysztof Wilczynski
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

// Package stack implements a Linked List based stack, a LIFO (Last In,
// Fist Out) data structure.
//
// For more information about Stack and similar Abstract Data Types (ADT)
// please refer to: http://en.wikipedia.org/wiki/Abstract_data_type
package stack

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	// Pop and Peek returns this error when the Stack is empty.
	ErrEmptyStack = errors.New("stack: stack is empty")
	// Search will panic with this error when given function is invalid.
	ErrNotAFunc = errors.New("stack: not a function or nil pointer")
)

// Stack represents a LIFO (Last In, First Out) data structure.
type Stack struct {
	top  *Element
	size int
}

// Element is an element in the Linked List.
type Element struct {
	value interface{}
	next  *Element
}

// Init initialises or clears a Stack.
func (s *Stack) Init() *Stack {
	s.top = nil
	s.size = 0
	return s
}

// New returns an initialised Stack.
func New() *Stack {
	return &Stack{
		top:  nil,
		size: 0,
	}
}

// String returns a string representation a Stack.
func (s *Stack) String() string {
	return fmt.Sprintf("Stack{%d}", s.size)
}

// Len returns the number of elements on the Stack.
func (s *Stack) Len() int {
	return s.size
}

// Empty checks whether Stack is empty. See Len().
func (s *Stack) Empty() bool {
	if s.size == 0 {
		return true
	}

	return false
}

// Push pushes an element onto the top of the Stack.
func (s *Stack) Push(value interface{}) {
	s.top = &Element{
		value: value,
		next:  s.top,
	}
	s.size++
}

// Peek looks at the element on the top of the Stack without removing it,
// or returns an error in a case of empty Stack.
func (s *Stack) Peek() (interface{}, error) {
	if s.Empty() {
		return nil, ErrEmptyStack
	}

	return s.top.value, nil
}

// Pop removes an element at the top of the Stack and returns it,
// or returns an error in a case of empty Stack.
func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, ErrEmptyStack
	}

	value := s.top.value
	s.top = s.top.next

	s.size--

	return value, nil
}

// Search checks whether an item exists on the Stack and returns the 1-based
// position where an element is on the Stack.
//
// A top-most element on the stack is considered to have distance of 1, whereas
// distance of 0 indicates that an element is not on the Stack.
//
// If an element occurs more than once in the Stack then only occurrence nearest
// the top of the Stack is taken into the account.
//
// Search takes a function that will be used to determine whether a given item
// exists on the Stack or not. Such function should return a single boolean value.
//
// Search will call given function for every element on the Stack proceeding in
// a descending order passing an element taken from the Stack as its argument.
func (s *Stack) Search(f func(element interface{}) bool) (bool, int) {
	if f == nil || reflect.TypeOf(f).Kind() != reflect.Func {
		panic(ErrNotAFunc)
	}

	distance := 0
	for e := s.top; e != nil; e = e.next {
		distance++
		if ok := f(e.value); ok {
			return true, distance
		}
	}

	return false, distance
}
