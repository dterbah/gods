package arraylist

import (
	"errors"
	"sort"

	"github.com/dterbah/gods/list"
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
func New[T any](comparator comparator.Comparator[T], elements ...T) *ArrayList[T] {
	var zero T
	list := &ArrayList[T]{zeroElement: zero, comparator: comparator}
	for _, element := range elements {
		list.Add(element)
	}
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
Clear all the elements in the list. After a clear, the list is totally empty
*/
func (list *ArrayList[T]) Clear() {
	list.elements = make([]T, int(growCapacityFactor))
	list.size = 0
}

func (list *ArrayList[T]) Concat(elements list.List[T]) list.List[T] {
	newList := New[T](list.comparator)

	list.ForEach(func(element T, index int) {
		newList.Add(element)
	})

	elements.ForEach(func(element T, index int) {
		newList.Add(element)
	})

	return newList
}

/*
Return true if the list contains at least one occurence of the element, else false
*/
func (list ArrayList[T]) Contains(element T) bool {
	for _, currentElement := range list.elements[:list.size] {
		if list.comparator(currentElement, element) == 0 {
			return true
		}
	}

	return false
}

/*
*
Create a copy of the current list
*/
func (list *ArrayList[T]) Copy() list.List[T] {
	newList := New[T](list.comparator)

	list.ForEach(func(element T, _ int) {
		newList.Add(element)
	})

	return newList
}

/*
Filter the list according to the specified callback passed in parameter.
It will return a new List that match the filter
*/
func (list *ArrayList[T]) Filter(callback func(element T) bool) list.List[T] {
	newList := New[T](list.comparator)

	for _, element := range list.elements[:list.size] {
		if callback(element) {
			newList.Add(element)
		}
	}

	return newList
}

/*
Apply a function for each element of the list
*/
func (list *ArrayList[T]) ForEach(callback func(element T, index int)) {
	for index, element := range list.elements[:list.size] {
		callback(element, index)
	}
}

/*
Return the index in the list of the element (if the element exists in the list)
If the element is not present in the list, the method will return -1
*/
func (list ArrayList[T]) IndexOf(element T) int {
	for index, currentElement := range list.elements[:list.size] {
		if list.comparator(currentElement, element) == 0 {
			return index
		}
	}

	return -1
}

/*
Check if the list is empty or not. Return true if it is empty, otherwise false
*/
func (list ArrayList[T]) IsEmpty() bool {
	return list.size == 0
}

/*
Alias for the method Size and used for the sort. It shoud not be called directly
*/
func (list ArrayList[T]) Len() int {
	return list.size
}

/*
This method is used by sort.Sort to sort the list. It should not be call directly
*/
func (list ArrayList[T]) Less(i, j int) bool {
	return list.comparator(list.elements[i], list.elements[j]) == -1
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

func (list *ArrayList[T]) Reverse() {
	for i, j := 0, list.size-1; i < j; i, j = i+1, j-1 {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}

/*
Retrieve the list size
*/
func (list ArrayList[T]) Size() int {
	return list.size
}

func (list *ArrayList[T]) Some(callback func(element T, index int) bool) bool {
	result := false

	for index, element := range list.elements[:list.size] {
		result = result || callback(element, index)
	}

	return result
}

/*
Sort the list
*/
func (list *ArrayList[T]) Sort() {
	sort.Sort(list)
}

/*
Return a sublist according to the range [start:between].
It will return the same list if the start and end are out of bounds
(< 0 or >= list size), or if start > end
*/
func (list *ArrayList[T]) SubList(start, end int) list.List[T] {
	if (start < 0 || start > list.size) || (end < 0 || end > list.size) {
		return list
	}

	if start > end {
		return list
	}

	newList := New[T](list.comparator)
	for _, element := range list.elements[start:end] {
		newList.Add(element)
	}

	return newList
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

// Modules methods
func From() {

}
