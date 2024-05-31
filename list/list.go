package list

/*
This interface defines which functions could be used in a List.
It defines features inspired by Javascript
*/
type List[T any] interface {
	/*
		Add an element in the current List
		@param element The element to add in the list
	*/
	Add(elements ...T)

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
		Call a function for each element of the list
	*/
	ForEach(callback func(element T, index int))

	/*
		Return the index in the list of the element (if the element exists in the list)
		If the element is not present in the list, the method will return -1
	*/
	IndexOf(element T) int

	/*
		Return true if the list has no elements, otherwise false
	*/
	IsEmpty() bool

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
		Retrieve the size of the list
	*/
	Size() int

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
