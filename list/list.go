package list

import (
	"github.com/dterbah/gods/collection"
	"github.com/dterbah/gods/iterable"
)

/*
This interface defines which functions could be used in a List.
It defines features inspired by Javascript
*/
type List[T any] interface {
	collection.Collection[T]
	iterable.Iterable[T]

	/*
		Concat a list with the current one. The result is a new list with all elements
		of the current list and the one passed in parameter
	*/
	Concat(list List[T]) List[T]

	/*
		Create a copy of the current list. It is only a shallow copy
	*/
	Copy() List[T]

	/*
		Filter the list according to the specified callback passed in parameter
	*/
	Filter(callback func(element T) bool) List[T]

	/*
		Remove the element at the specified index in the list.
		If the element is correctly removed, it will return true.
		Otherwise, false
	*/
	RemoveAt(index int) bool

	/*
		Replace the element at the indice "index" with the new one.
		This method will return true if the previous element is correctly replaced, else false.
	*/
	ReplaceAt(index int, element T) bool

	/*
		Reverse the elements inside the list
	*/
	Reverse()

	/*
		Check if at least one element matchs with the callback in parameter.
	*/
	Some(callback func(element T, index int) bool) bool

	/*
		Sort the list
	*/
	Sort()

	/*
		Return a list with the elements between in the range of start:end
	*/
	SubList(start, end int) List[T]
}
