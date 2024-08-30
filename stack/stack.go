package stack

import (
	"github.com/ryqdev/golang_utils/util"
)

type StackType interface {
	int | float32 | float64 | string
}

type Stack[T StackType] interface {
	Top() T
	Push(v T)
	Pop(v T)
	IsEmpty() bool
	Size() int
	Clear()
}

type GoStack[T StackType] struct {
	items []T
}

func New[T StackType](initSize ...int) *GoStack[T] {
	if len(initSize) > 1 {
		// TODO: is panic suitable?
		panic("Invalid argument")
	}
	size := util.IfThenElse(len(initSize) == 0, 0, initSize[0])
	return &GoStack[T]{
		items: make([]T, size),
	}
}

func (s *GoStack[T]) Top() T {
	if len(s.items) == 0 {
		panic("OOPS")
	}
	return s.items[len(s.items)-1]
}

func (s *GoStack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *GoStack[T]) Pop() {
	if len(s.items) == 0 {
		panic("OOPS")
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *GoStack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *GoStack[T]) Clear() {
	s.items = make([]T, 0)
}

func (s *GoStack[T]) Size() int {
	return len(s.items)
}
