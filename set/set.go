package set

import (
	"github.com/Yukun4119/golang_utils/util"
)

type Set interface {
	Add(v interface{}) bool
	Remove(v interface{}) bool
	IsExist(v interface{}) bool
	IsEmpty() bool
	Size() int
	Clear()
}

type HashSet struct {
	items map[interface{}]bool
}

func New() *HashSet {
	return &HashSet{
		items: make(map[interface{}]bool),
	}
}

func (set *HashSet) Add(item interface{}) {
	if _, exists := set.items[item]; exists {
		return
	}
	set.items[item] = true
}

func (set *HashSet) Remove(item interface{}) {
	if _, exists := set.items[item]; !exists {
		return
	}
	delete(set.items, item)
}

func (set *HashSet) Size() int {
	return len(set.items)
}

func (set *HashSet) IsExist(item interface{}) bool {
	_, exists := set.items[item]
	return util.IfThenElse(exists, true, false)
}

func (set *HashSet) IsEmpty(item interface{}) bool {
	return util.IfThenElse(set.Size() == 0, true, false)
}

func (set *HashSet) Clear() {
	set.items = make(map[interface{}]bool)
}
