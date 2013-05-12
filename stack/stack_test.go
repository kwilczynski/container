package stack_test

import (
	"bytes"
	"fmt"
	. "github.com/kwilczynski/container/stack"
	"reflect"
	"testing"
)

func compareStrings(this, other string) bool {
	return bytes.Equal([]byte(this), []byte(other))
}

func TestNew(t *testing.T) {
	s := New()
	func(v interface{}) {
		if _, ok := v.(*Stack); !ok {
			t.Errorf("not a Stack type: %v", reflect.TypeOf(v).String())
		}
	}(s)
}

func TestString(t *testing.T) {
	s := New()
	s.Push(0)
	s.Push(1)

	v := fmt.Sprintf("Stack{%d}", s.Len())
	if !compareStrings(s.String(), v) {
		t.Errorf("value given \"%v\", want \"%v\"", v, s.String())
	}
}

func TestInit(t *testing.T) {
	s := New()
	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	s.Init()
	v, err := s.Pop()

	ok := compareStrings(err.Error(), ErrEmptyStack.Error())
	if v != nil || !ok {
		t.Errorf("value given {%v, %v}, want {%v, %v}",
			v, err.Error(), nil, ErrEmptyStack.Error())
	}

	if !s.Empty() {
		t.Errorf("value given %v, want %v", s.Empty(), true)
	}
}

func TestLen(t *testing.T) {
	s := New()
	last := 0
	for i := 0; i < 10; i++ {
		s.Push(i)
		last = i + 1
	}

	if s.Len() != last {
		t.Errorf("value given %v, want %v", s.Len(), last)
	}

	s.Push(0)
	s.Push(0)

	last += 2
	if s.Len() != last {
		t.Errorf("value given %v, want %v", s.Len(), last)
	}

	s.Init()

	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	for i := last; i > s.Len(); i-- {
		s.Pop()
		last = i - 1
	}

	if s.Len() != last {
		t.Errorf("value given %v, want %v", s.Len(), last)
	}
}

func TestEmpty(t *testing.T) {
	s := New()
	if s.Empty() != true {
		t.Errorf("value given %v, want %v", s.Empty(), true)
	}

	s.Push(0)
	s.Push(0)

	if s.Empty() != false {
		t.Errorf("value given %v, want %v", s.Empty(), false)
	}

	s.Init()

	if s.Empty() != true {
		t.Errorf("value given %v, want %v", s.Empty(), true)
	}
}

func TestPush(t *testing.T) {
	s := New()
	sum, check := 0, 0
	for i := 0; i < 10; i++ {
		s.Push(i)
		v, _ := s.Pop()

		if _, ok := v.(int); !ok {
			t.Errorf("type given %v, want %v", reflect.TypeOf(v).Kind(),
				reflect.TypeOf(i).Kind())
		}

		if v.(int) != i {
			t.Errorf("value given %v, want %v", v.(int), i)
		}
		check += v.(int)
		sum += i
	}

	s.Push(1)
	s.Push(1)
	sum += 2

	for !s.Empty() {
		v, _ := s.Pop()
		check += v.(int)
	}

	if sum != check {
		t.Errorf("value given %v, want %v", check, sum)
	}
}

func TestPeek(t *testing.T) {
	s := New()
	v, err := s.Peek()

	ok := compareStrings(err.Error(), ErrEmptyStack.Error())
	if v != nil || !ok {
		t.Errorf("value given {%v, %v}, want {%v, %v}",
			v, err.Error(), nil, ErrEmptyStack.Error())
	}

	if !s.Empty() {
		t.Errorf("value given %v, want %v", s.Empty(), true)
	}

	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	for i := s.Len(); i > 0; i-- {
		v, _ := s.Peek()
		if _, ok := v.(int); !ok {
			t.Errorf("type given %v, want %v", reflect.TypeOf(v).Kind(),
				reflect.TypeOf(i).Kind())
		}

		s.Pop()

		if v.(int) != i-1 {
			t.Errorf("value given %v, want %v", v.(int), i)
		}
	}

	s.Init()

	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	sum, check := 0, 0
	for !s.Empty() {
		peek, _ := s.Peek()
		pop, _ := s.Pop()
		check += peek.(int)
		sum += pop.(int)
	}

	if sum != check {
		t.Errorf("value given %v, want %v", check, sum)
	}

	if s.Empty() != true {
		t.Errorf("value given %v, want %v", s.Empty(), true)
	}
}

func TestPop(t *testing.T) {
	s := New()
	v, err := s.Pop()

	ok := compareStrings(err.Error(), ErrEmptyStack.Error())
	if v != nil || !ok {
		t.Errorf("value given {%v, %v}, want {%v, %v}",
			v, err.Error(), nil, ErrEmptyStack.Error())
	}

	if !s.Empty() {
		t.Errorf("value given %v, want %v", s.Empty(), true)
	}

	for i := 0; i < 10; i++ {
		s.Push(i)

		v, _ := s.Pop()
		if _, ok := v.(int); !ok {
			t.Errorf("type given %v, want %v", reflect.TypeOf(v).Kind(),
				reflect.TypeOf(i).Kind())
		}
		if v.(int) != i {
			t.Errorf("value given %v, want %v", v.(int), i)
		}
	}

	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	sum, check := 0, 0
	for !s.Empty() {
		peek, _ := s.Peek()
		pop, _ := s.Pop()
		check += pop.(int)
		sum += peek.(int)
	}

	if sum != check {
		t.Errorf("value given %v, want %v", check, sum)
	}

	if s.Empty() != true {
		t.Errorf("value given %v, want %v", s.Empty(), true)
	}
}

func TestSearch(t *testing.T) {
	s := New()

	for i := 1; i <= 10; i++ {
		s.Push(i)
	}

	b, v := s.Search(func(v interface{}) bool {
		return v == 1
	})

	if b != true || v != s.Len() {
		t.Errorf("value given {%v, %v}, want {%v, %v}",
			b, v, true, s.Len())
	}

	s.Pop()
	s.Pop()

	b, v = s.Search(func(v interface{}) bool {
		return v == 1
	})

	if b != true || v != s.Len() {
		t.Errorf("value given {%v, %v}, want {%v, %v}",
			b, v, true, s.Len())
	}

	s.Init()

	b, v = s.Search(func(v interface{}) bool {
		return v == 1
	})

	if b != false || v != s.Len() {
		t.Errorf("value given {%v, %v}, want {%v, %v}",
			b, v, false, s.Len())
	}

	defer func() {
		r := recover()
		if r == nil {
			t.Error("did not panic")
			return
		}
		if ok := compareStrings(r.(error).Error(), ErrNotAFunc.Error()); !ok {
			t.Errorf("value given \"%v\", want \"%v\"",
				r.(error).Error(), ErrNotAFunc.Error())
			return
		}
	}()

	s.Search(nil)
}
