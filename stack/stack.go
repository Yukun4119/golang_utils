package stack

import (
	"github.com/ryqdev/golang_utils/util"
)

type Stack interface {
	Top() any
	Push(v any)
	Pop(v any)
	IsEmpty() bool
	Size() int
	Clear()
}

type GoStack struct {
	items []any
}

func New(initSize ...int) *GoStack {
	if len(initSize) > 1 {
		// TODO: is panic suitable?
		panic("Invalid argument")
	}
	size := util.IfThenElse(len(initSize) == 0, 0, initSize[0])
	return &GoStack{
		items: make([]any, size),
	}
}

func (s *GoStack) Top() any {
	if len(s.items) == 0 {
		panic("OOPS")
	}
	return s.items[len(s.items)-1]
}

func (s *GoStack) Push(item any) {
	s.items = append(s.items, item)
}

func (s *GoStack) Pop() {
	if len(s.items) == 0 {
		panic("OOPS")
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *GoStack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *GoStack) Clear() {
	s.items = make([]any, 0)
}

func (s *GoStack) Size() int {
	return len(s.items)
}
