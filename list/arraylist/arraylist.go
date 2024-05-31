package arraylist

import "errors"

/*
Struct that defines what is an ArrayList.
It will store all the elements inside a slice.
The available methods of this structures wi
*/
type ArrayList[T any] struct {
	elements []T
	size     int
	empty    T
}

// Constants
const growCapacityFactor = float32(2.0)

// Public methods
func New[T any](values ...T) *ArrayList[T] {
	var zero T
	list := &ArrayList[T]{empty: zero}

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
		return list.empty, errors.New("index out of bound")
	}

	return list.elements[index], nil
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
