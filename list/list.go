package list

/*
This interface defines which functions could be used in a Collection
*/
type List[T any] interface {
	/**
	Add an element in the current List
	@param element The element to add in the list
	*/
	Add(elements ...T)

	/**
	Retrieve an element by its index
	The returned result is either the element at the index (if index < listSize), either nil
	**/
	At(index int) (T, error)

	/**
	Clear all the content inside the list
	*/
	Clear()

	/**
	Retrieve the size of the list
	*/
	Size() int

	/**
	Return true if the list has no elements, otherwise false
	*/
	IsEmpty() bool

	/**
	Return true if the list contains at list one occurence of the element, otherwise false
	*/
	Contains(element T) bool
}
