package arraylist

import (
	"errors"
	"sort"

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
Retrieve an element by its index
If the index is negative or greater than the list size, the method will return an error
*/
func (list *ArrayList[T]) At(index int) (T, error) {
	if list.isOutOfBounds(index) {
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
		if list.comparator(currentElement, element) == 0 {
			return true
		}
	}

	return false
}

/*
Filter the list according to the specified callback passed in parameter.
It will return a new List that match the filter
*/
func (list *ArrayList[T]) Filter(callback func(element T) bool) *ArrayList[T] {
	newList := New[T](list.comparator)

	for _, element := range list.elements[:list.size] {
		if callback(element) {
			newList.Add(element)
		}
	}

	return newList
}

func (list *ArrayList[T]) ForEach(callback func(element T, index int)) {
	for index, element := range list.elements[:list.size] {
		callback(element, index)
	}
}

/*
Alias for the method Size and used for the sort. It shoud not be called directly
*/
func (list *ArrayList[T]) Len() int {
	return list.size
}

/*
This method is used by sort.Sort to sort the list. It should not be call directly
*/
func (list *ArrayList[T]) Less(i, j int) bool {
	return list.comparator(list.elements[i], list.elements[j]) == -1
}

/*
Return the index in the list of the element (if the element exists in the list)
If the element is not present in the list, the method will return -1
*/
func (list *ArrayList[T]) IndexOf(element T) int {
	for index, currentElement := range list.elements[:list.size] {
		if list.comparator(currentElement, element) == 0 {
			return index
		}
	}

	return -1
}

/*
Remove the element at the specified index in the list.
If the element is correctly removed, it will return true.
Otherwise, false
*/
func (list *ArrayList[T]) RemoveAt(index int) bool {
	if list.isOutOfBounds(index) {
		return false
	}

	for rangeIndex := range list.elements[index:list.size] {
		currentIndex := rangeIndex + index
		list.elements[currentIndex] = list.elements[currentIndex+1]
	}

	list.size--

	return true
}

func (list *ArrayList[T]) ReplaceAt(index int, element T) bool {
	if list.isOutOfBounds(index) {
		return false
	}

	list.elements[index] = element

	return true
}

/*
Retrieve the list size
*/
func (list *ArrayList[T]) Size() int {
	return list.size
}

func (list *ArrayList[T]) Sort() {
	sort.Sort(list)
}

/*
Swap two elements in the list. This method is used by the Sort method, it should not be called directly
*/
func (list *ArrayList[T]) Swap(i, j int) {
	list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
}

// Private methods

// Resize the size of the list
func (list *ArrayList[T]) resize(cap int) {
	newElements := make([]T, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

// Grow the list
func (list *ArrayList[T]) growIfNeeded(n int) {
	currentCapacity := cap(list.elements)
	if currentCapacity <= (list.size + n) {
		// We need to grow the capacity of the list
		newCapacity := int(growCapacityFactor * float32(currentCapacity+n))
		list.resize(newCapacity)
	}
}

/*
Method used to know if an index is out of bounds the range of the list.
To be true, the index should be < 0 or >= list size
*/
func (list *ArrayList[T]) isOutOfBounds(index int) bool {
	return index < 0 || index >= list.size
}
