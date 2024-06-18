package stack

import (
	"errors"
	"fmt"

	"github.com/dterbah/gods/iterable"
	comparator "github.com/dterbah/gods/utils"
)

/*
Struct that represents what is a Stack.
All the elements will be stored in an array
that will grow according to the needs
*/
type Stack[T any] struct {
	elements    []T
	size        int
	zeroElement T
	comparator  comparator.Comparator[T]
}

/*
Create a new Stack
*/
func New[T any](comparator comparator.Comparator[T]) *Stack[T] {
	var zero T
	return &Stack[T]{zeroElement: zero, comparator: comparator}
}

/*
Clear all the elements in the Stack
*/
func (stack *Stack[T]) Clear() {
	stack.elements = []T{}
	stack.size = 0
}

func (stack Stack[T]) Copy() *Stack[T] {
	newStack := New(stack.comparator)

	stack.ForEach(func(element T, index int) {
		newStack.Push(element)
	})

	return newStack
}

/*
Returns true if the stack contains the specified element,
else false
*/
func (stack Stack[T]) Contains(element T) bool {
	for _, value := range stack.elements[:stack.size] {
		if stack.comparator(value, element) == 0 {
			return true
		}
	}

	return false
}

/*
Call a function for each element in the queue
*/
func (stack Stack[T]) ForEach(callback func(element T, index int)) {
	for index, element := range stack.elements[:stack.size] {
		callback(element, index)
	}
}

/*
Create a stack from an iterable
*/
func FromIterable[T any](iterable iterable.Iterable[T], comparator comparator.Comparator[T]) *Stack[T] {
	stack := New(comparator)

	iterable.ForEach(func(element T, index int) {
		stack.Push(element)
	})

	return stack
}

/*
Return true if the stack is empty, else false
*/
func (stack Stack[T]) IsEmpty() bool {
	return stack.size == 0
}

/*
Return the last element of the stack without removing it. If the
stack is empty, this method will return an error
*/
func (stack Stack[T]) Peek() (T, error) {
	if stack.IsEmpty() {
		return stack.zeroElement, errors.New("empty stack")
	}

	return stack.elements[stack.size-1], nil
}

/*
Return the last element of the stack and remove it from the stack.
It will return an error if the stack is empty
*/
func (stack *Stack[T]) Pop() (T, error) {
	if stack.IsEmpty() {
		return stack.zeroElement, errors.New("empty stack")
	}

	element := stack.elements[stack.size-1]
	stack.shiftElements()
	return element, nil
}

func (stack Stack[T]) Print() {
	fmt.Print("[")

	stack.ForEach(func(element T, index int) {
		fmt.Print(element)
		if index < stack.size-1 {
			fmt.Print(", ")
		}
	})

	fmt.Println("]")
}

/*
Push elements i, the stack
*/
func (stack *Stack[T]) Push(elements ...T) {
	stack.growIfNeeded(len(elements))

	for _, element := range elements {
		stack.elements[stack.size] = element
		stack.size++
	}
}

/*
Return the number of elements in the stack
*/
func (stack Stack[T]) Size() int {
	return stack.size
}

// Private methods //
func (stack *Stack[T]) growIfNeeded(n int) {
	currentCapacity := cap(stack.elements)
	if currentCapacity <= stack.size+n {
		newCapacity := (currentCapacity + n) * 2
		newElements := make([]T, newCapacity)
		copy(newElements, stack.elements)
		stack.elements = newElements
	}
}

func (stack *Stack[T]) shiftElements() {
	stack.size--
}
