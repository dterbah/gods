package iterable

type Iterable[T any] interface {
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
