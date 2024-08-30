package set

import (
	"github.com/ryqdev/golang_utils/util"
)

type Set interface {
	Add(v any) bool
	Remove(v any) bool
	IsExist(v any) bool
	IsEmpty() bool
	Size() int
	Clear()
}

type HashSet struct {
	items map[any]bool
}

func New() *HashSet {
	return &HashSet{
		items: make(map[any]bool),
	}
}

func (set *HashSet) Add(item any) {
	if _, exists := set.items[item]; exists {
		return
	}
	set.items[item] = true
}

func (set *HashSet) Remove(item any) {
	if _, exists := set.items[item]; !exists {
		return
	}
	delete(set.items, item)
}

func (set *HashSet) Size() int {
	return len(set.items)
}

func (set *HashSet) IsExist(item any) bool {
	_, exists := set.items[item]
	return util.Ternary(exists, true, false)
}

func (set *HashSet) IsEmpty() bool {
	return util.Ternary(set.Size() == 0, true, false)
}

func (set *HashSet) Clear() {
	set.items = make(map[any]bool)
}
