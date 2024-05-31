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
		Retrieve the size of the list
	*/
	Size() int
}
