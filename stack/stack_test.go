package stack

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	s := New()
	func(v interface{}) {
		if _, ok := v.(*Stack); !ok {
			t.Fatalf("not a Stack type: %s", reflect.TypeOf(v).String())
		}
	}(s)
}

func TestStack_String(t *testing.T) {
	s := New()
	s.Push(0)
	s.Push(1)

	v := fmt.Sprintf("Stack{%d}", s.Len())
	if ok := CompareStrings(s.String(), v); !ok {
		t.Errorf("value given \"%s\", want \"%s\"", s.String(), v)
	}
}

func TestStack_Init(t *testing.T) {
	s := New()
	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	s.Init()
	v, err := s.Pop()

	ok := CompareStrings(err.Error(), ErrEmptyStack.Error())
	if v != nil || !ok {
		t.Errorf("value given {%v, %s}, want {%v, %s}",
			v, err.Error(), nil, ErrEmptyStack.Error())
	}

	if ok := s.Empty(); !ok {
		t.Errorf("value given %v, want %v", ok, true)
	}
}

func TestStack_Len(t *testing.T) {
	s := New()

	last := 0
	for i := 0; i < 10; i++ {
		s.Push(i)
		last = i + 1
	}

	if s.Len() != last {
		t.Errorf("value given %d, want %d", s.Len(), last)
	}

	s.Push(0)
	s.Push(0)

	last += 2
	if s.Len() != last {
		t.Errorf("value given %d, want %d", s.Len(), last)
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
		t.Errorf("value given %d, want %d", s.Len(), last)
	}
}

func TestStack_Empty(t *testing.T) {
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

func TestStack_Push(t *testing.T) {
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
			t.Errorf("value given %d, want %d", v.(int), i)
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
		t.Errorf("value given %d, want %d", check, sum)
	}
}

func TestStack_Peek(t *testing.T) {
	s := New()
	v, err := s.Peek()

	ok := CompareStrings(err.Error(), ErrEmptyStack.Error())
	if v != nil || !ok {
		t.Errorf("value given {%v, %s}, want {%v, %s}",
			v, err.Error(), nil, ErrEmptyStack.Error())
	}

	if ok := s.Empty(); !ok {
		t.Errorf("value given %v, want %v", ok, true)
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
			t.Errorf("value given %d, want %d", v.(int), i)
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
		t.Errorf("value given %d, want %d", check, sum)
	}

	if ok := s.Empty(); !ok {
		t.Errorf("value given %v, want %v", ok, true)
	}
}

func TestStack_Pop(t *testing.T) {
	s := New()
	v, err := s.Pop()

	ok := CompareStrings(err.Error(), ErrEmptyStack.Error())
	if v != nil || !ok {
		t.Errorf("value given {%v, %s}, want {%v, %s}",
			v, err.Error(), nil, ErrEmptyStack.Error())
	}

	if ok := s.Empty(); !ok {
		t.Errorf("value given %v, want %v", ok, true)
	}

	for i := 0; i < 10; i++ {
		s.Push(i)

		v, _ := s.Pop()
		if _, ok := v.(int); !ok {
			t.Errorf("type given %v, want %v", reflect.TypeOf(v).Kind(),
				reflect.TypeOf(i).Kind())
		}
		if v.(int) != i {
			t.Errorf("value given %d, want %d", v.(int), i)
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
		t.Errorf("value given %d, want %d", check, sum)
	}

	if ok := s.Empty(); !ok {
		t.Errorf("value given %v, want %v", ok, true)
	}
}

func TestStack_Search(t *testing.T) {
	s := New()

	for i := 1; i <= 10; i++ {
		s.Push(i)
	}

	b, v, _ := s.Search(func(v interface{}) bool {
		return v == 1
	})

	if b != true || v != s.Len() {
		t.Errorf("value given {%v, %d}, want {%v, %d}",
			b, v, true, s.Len())
	}

	s.Pop()
	s.Pop()

	b, v, _ = s.Search(func(v interface{}) bool {
		return v == 1
	})

	if b != true || v != s.Len() {
		t.Errorf("value given {%v, %d}, want {%v, %d}",
			b, v, true, s.Len())
	}

	s.Init()

	b, v, _ = s.Search(func(v interface{}) bool {
		return v == 1
	})

	if b != false || v != s.Len() {
		t.Errorf("value given {%v, %d}, want {%v, %d}",
			b, v, false, s.Len())
	}

	s.Push(`a`)
	s.Push(`b`)
	s.Push(`c`)

	var e interface{}

	b, v, e = s.Search(func(v interface{}) bool {
		return v == `b`
	})

	if b != true || v != 2 || e != `b` {
		t.Errorf("value given {%v, %d, %v}, want {%v, %d, %v}",
			b, v, e, true, 2, `b`)
	}

	defer func() {
		r := recover()
		if r == nil {
			t.Error("did not panic")
			return
		}
		if ok := CompareStrings(r.(error).Error(), ErrNotAFunc.Error()); !ok {
			t.Errorf("value given \"%s\", want \"%s\"",
				r.(error).Error(), ErrNotAFunc.Error())
			return
		}
	}()

	s.Search(nil)
}
