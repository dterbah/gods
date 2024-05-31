package set

import (
	"github.com/dterbah/gods/list"
	"github.com/dterbah/gods/list/arraylist"
	comparator "github.com/dterbah/gods/utils"
)

// @todo: Union, intersection, etc

/*
Struct that defines what a Set is. The Set can only store on occurence
of a specific value
*/
type Set[T any] struct {
	elements list.List[T]
}

/*
Create a new Set
*/
func New[T any](comparator comparator.Comparator[T], elements ...T) *Set[T] {
	list := arraylist.New[T](comparator)
	set := &Set[T]{elements: list}
	for _, element := range elements {
		set.Add(element)
	}
	return set
}

/*
Add elements in the Set. If some elements are already
present in the Set, they won't be include a second time
*/
func (set *Set[T]) Add(elements ...T) {
	for _, element := range elements {
		if !set.elements.Contains(element) {
			set.elements.Add(element)
		}
	}
}

/*
Retrieve an element by its index
If the index is negative or greater than the set size, the method will return an error
*/
func (set *Set[T]) At(index int) (T, error) {
	return set.elements.At(index)
}

/*
Clear all the elements in the set. After a clear, the set is totally empty
*/
func (set *Set[T]) Clear() {
	set.elements.Clear()
}

/*
Return the size of the set
*/
func (set Set[T]) Size() int {
	return set.elements.Size()
}
