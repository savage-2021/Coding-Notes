package arraylist

import (
	"lists/lists"
)

func assertListImplmementation() {
	var _ lists.List = (*List)(nil)
}

type List struct {
	elements []interface{}
	size int
}

const (
	growthFactor = float32(2.0)
	shrinkFactor = float32(0.25)
)

// New instantiates a new list and adds teh passed values, if any to the list
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values)
	}
	return list
}

// Add appends a value at the end of a list
func (list *List) Add(values ...interface{}){
	list.growBy(len(values))
	for _ , value := range values {
		list.elements[list.size] = value
		list.size++
	}
}

// Get returns the elemet at index
func(list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index){
		return nil, false
	}

	return list.elements[index], true
}

func (list *List) Remove(index int){
	if !list.withinRange(index){
		return 
	}
	list.elements[index] = nil
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.shrink()
}
func(list *List) withinRange(index int) bool {
	return index >=0 && index < list.size
}