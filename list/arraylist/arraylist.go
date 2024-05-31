package arraylist

import (
	"errors"

	comparator "github.com/dterbah/gods/utils"
)

/*
Struct that defines what is an ArrayList.
It will store all the elements inside a slice.
The available methods of this structures wi
*/
type ArrayList[T any] struct {
	elements    []T
	size        int
	zeroElement T
	comparator  comparator.Comparator[T]
}

// Constants
const growCapacityFactor = float32(2.0)

// Public methods
func New[T any](comparator comparator.Comparator[T]) *ArrayList[T] {
	var zero T
	list := &ArrayList[T]{zeroElement: zero, comparator: comparator}

	return list
}

/*
*

	Add an element inside the list
*/
func (list *ArrayList[T]) Add(elements ...T) {
	list.growIfNeeded(len(elements))

	for _, element := range elements {
		list.elements[list.size] = element
		list.size++
	}
}

/*
*

	Retrieve an element by its index
	If the index is negative or greater than the list size, the method will return an error
*/
func (list *ArrayList[T]) At(index int) (T, error) {
	if index >= list.size || index < 0 {
		return list.zeroElement, errors.New("index out of bound")
	}

	return list.elements[index], nil
}

/*
Check if the list is empty or not. Return true if it is empty, otherwise false
*/
func (list *ArrayList[T]) IsEmpty() bool {
	return list.size == 0
}

/*
Clear all the elements in the list. After a clear, the list is totally empty
*/
func (list *ArrayList[T]) Clear() {
	list.elements = make([]T, int(growCapacityFactor))
	list.size = 0
}

/*
Return true if the list contains at least one occurence of the element, else false
*/
func (list *ArrayList[T]) Contains(element T) bool {
	for _, currentElement := range list.elements[:list.size] {
		if list.comparator(currentElement, element) {
			return true
		}
	}

	return false
}

/*
*

	Retrieve the list size
*/
func (list *ArrayList[T]) Size() int {
	return list.size
}

// Private methods
func (list *ArrayList[T]) resize(cap int) {
	newElements := make([]T, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

func (list *ArrayList[T]) growIfNeeded(n int) {
	currentCapacity := cap(list.elements)
	if currentCapacity <= (list.size + n) {
		// We need to grow the capacity of the list
		newCapacity := int(growCapacityFactor * float32(currentCapacity+n))
		list.resize(newCapacity)
	}
}
