package iterable

type Iterable[T any] interface {
	/*
		Retrieve an element by its index
		The returned result is either the element at the index (if index < listSize), either nil
	*/
	At(index int) (T, error)

	/*
		Call a function for each element of the list
	*/
	ForEach(callback func(element T, index int))

	/*
		Return the index in the list of the element (if the element exists in the list)
		If the element is not present in the list, the method will return -1
	*/
	IndexOf(element T) int
}
