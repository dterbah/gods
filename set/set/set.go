package set

import (
	"fmt"

	"github.com/dterbah/gods/collection"
	"github.com/dterbah/gods/list"
	"github.com/dterbah/gods/list/arraylist"
	"github.com/dterbah/gods/set"
	comparator "github.com/dterbah/gods/utils"
)

// @todo: Union, intersection, etc

/*
Struct that defines what a Set is. The Set can only store on occurence
of a specific value
*/
type Set[T any] struct {
	elements   list.List[T]
	comparator comparator.Comparator[T]
}

/*
Create a new Set
*/
func New[T any](comparator comparator.Comparator[T], elements ...T) *Set[T] {
	list := arraylist.New[T](comparator)
	set := &Set[T]{elements: list, comparator: comparator}
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
Add all elements present in the collection
*/
func (set *Set[T]) AddAll(elements collection.Collection[T]) {
	for i := 0; i < elements.Size(); i++ {
		element, _ := elements.At(i)
		set.Add(element)
	}
}

/*
Retrieve an element by its index
If the index is negative or greater than the set size, the method will return an error
*/
func (set Set[T]) At(index int) (T, error) {
	return set.elements.At(index)
}

/*
Clear all the elements in the set. After a clear, the set is totally empty
*/
func (set *Set[T]) Clear() {
	set.elements.Clear()
}

/*
Return true if the set contains at least one occurence of the element, else false
*/
func (set Set[T]) Contains(element T) bool {
	return set.elements.Contains(element)
}

func (set Set[T]) ContainsAll(collection collection.Collection[T]) bool {
	for index := 0; index < collection.Size(); index++ {
		value, _ := collection.At(index)
		if !set.Contains(value) {
			return false
		}
	}

	return true
}

/*
*
Create a copy of the current list
*/
func (set *Set[T]) Copy() set.BasicSet[T] {
	newSet := New[T](set.comparator)

	set.elements.ForEach(func(element T, index int) {
		newSet.elements.Add(element)
	})

	return newSet
}

func (set *Set[T]) Diff(otherSet set.BasicSet[T]) set.BasicSet[T] {
	newSet := New(set.comparator)

	set.ForEach(func(element T, index int) {
		if !otherSet.Contains(element) {
			newSet.Add(element)
		}
	})

	return newSet
}

/*
Apply a function for each element of the set
*/
func (set *Set[T]) ForEach(callback func(element T, index int)) {
	set.elements.ForEach(callback)
}

/*
Return the index in the set of the element (if the element exists in the set)
If the element is not present in the set, the method will return -1
*/
func (set Set[T]) IndexOf(element T) int {
	return set.elements.IndexOf(element)
}

func (set *Set[T]) IsSubset(otherSet set.BasicSet[T]) bool {
	for i := 0; i < otherSet.Size(); i++ {
		value, _ := otherSet.At(i)
		if !set.Contains(value) {
			return false
		}
	}

	return true
}

/*
Compute the intersection between the current set and the one passed in parameter.
The result is equivalent of A ∩ B
*/
func (set *Set[T]) Intersection(otherSet set.BasicSet[T]) set.BasicSet[T] {
	newSet := New(set.comparator)

	set.ForEach(func(element T, index int) {
		if otherSet.Contains(element) {
			newSet.Add(element)
		}
	})

	return newSet
}

/*
Check if the set is empty or not. Return true if it is empty, otherwise false
*/
func (set *Set[T]) IsEmpty() bool {
	return set.elements.IsEmpty()
}

func (set Set[T]) Print() {
	fmt.Print("{")

	set.elements.ForEach(func(element T, index int) {
		fmt.Print(element)
		if index < set.Size()-1 {
			fmt.Print(", ")
		}
	})

	fmt.Println("}")
}

func (set *Set[T]) Remove(element T) {
	set.elements.Remove(element)
}

/*
Compute the union between the current set and the one passed in parameter.
The result is equivalent of A ∪ B
*/
func (set *Set[T]) Union(otherSet set.BasicSet[T]) set.BasicSet[T] {
	newSet := New(set.comparator)

	set.ForEach(func(element T, index int) {
		if !otherSet.Contains(element) {
			newSet.Add(element)
		}
	})

	otherSet.ForEach(func(element T, index int) {
		if !set.Contains(element) {
			newSet.Add(element)
		}
	})

	return newSet
}

/*
Return the size of the set
*/
func (set Set[T]) Size() int {
	return set.elements.Size()
}
