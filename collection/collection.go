package collection

/*
Interface that define what is a Collection
*/
type Collection[T any] interface {
	/*
		Add an element in the current List
		@param element The element to add in the list
	*/
	Add(elements ...T)

	/*
		Concat a list with the current one. The result is a new list with all elements
		of the current list and the one passed in parameter
	*/
	AddAll(list Collection[T])

	/*
		Retrieve an element by its index
		The returned result is either the element at the index (if index < listSize), either nil
	*/
	At(index int) (T, error)

	/*
		Clear all the content inside the list
	*/
	Clear()

	/*
		Return true if the list contains at list one occurence of the element, otherwise false
	*/
	Contains(element T) bool

	/*
		Return true if the list has no elements, otherwise false
	*/
	IsEmpty() bool

	/*
		Remove specified element if it exists
	*/
	Remove(element T)

	/*
		Retrieve the size of the list
	*/
	Size() int
}
