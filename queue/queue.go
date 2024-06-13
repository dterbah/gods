package queue

import (
	"errors"
	"fmt"

	comparator "github.com/dterbah/gods/utils"
)

/*
Struct that represents whats is a Queue.
It stores the elements in an array that grow when
needed.
*/
type Queue[T any] struct {
	elements    []T
	size        int
	zeroElement T
	comparator  comparator.Comparator[T]
}

/*
Create a new Queue
*/
func New[T any](comparator comparator.Comparator[T]) *Queue[T] {
	var zero T
	return &Queue[T]{elements: []T{}, zeroElement: zero, comparator: comparator}
}

/*
Remove all the elements of the queue
*/
func (queue *Queue[T]) Clear() {
	queue.size = 0
	queue.elements = []T{}
}

/*
Create a copy of the current Queue
*/
func (queue Queue[T]) Copy() *Queue[T] {
	newQueue := New[T](queue.comparator)

	queue.ForEach(func(element T, index int) {
		newQueue.Enqueue(element)
	})

	return newQueue
}

/*
Return true if the element is present inside the Queue, else false
*/
func (queue Queue[T]) Contains(element T) bool {
	for _, value := range queue.elements[:queue.size] {
		if queue.comparator(value, element) == 0 {
			return true
		}
	}

	return false
}

/*
Enqueue an element inside the Queue
*/
func (queue *Queue[T]) Enqueue(elements ...T) {
	queue.growIfNeeded(len(elements))
	for _, element := range elements {
		queue.elements[queue.size] = element
		queue.size++
	}
}

/*
Dequeue the first element of the Queue. If she is empty, it returns an error
*/
func (queue *Queue[T]) Dequeue() (T, error) {
	if queue.IsEmpty() {
		return queue.zeroElement, errors.New("queue empty")
	}

	head := queue.elements[0]
	queue.shiftElements()

	return head, nil
}

/*
Call a function for each element in the queue
*/
func (queue Queue[T]) ForEach(callback func(element T, index int)) {
	for index, element := range queue.elements[:queue.size] {
		callback(element, index)
	}
}

/*
Return true if no element is present in the Queue, else false
*/
func (queue Queue[T]) IsEmpty() bool {
	return queue.size == 0
}

func (queue Queue[T]) Peek() (T, error) {
	if queue.IsEmpty() {
		return queue.zeroElement, errors.New("queue empty")
	}

	return queue.elements[0], nil
}

func (queue Queue[T]) Print() {
	fmt.Print("[")

	queue.ForEach(func(element T, index int) {
		fmt.Print(element)
		if index < queue.size-1 {
			fmt.Print(", ")
		}
	})

	fmt.Println("]")
}

/*
Return the current size of the Queue
*/
func (queue Queue[T]) Size() int {
	return queue.size
}

// Private methods //
func (queue *Queue[T]) growIfNeeded(n int) {
	currentCapacity := cap(queue.elements)
	if currentCapacity <= queue.size+n {
		newCapacity := (currentCapacity + n) * 2
		newElements := make([]T, newCapacity)
		copy(newElements, queue.elements)
		queue.elements = newElements
	}
}

func (queue *Queue[T]) shiftElements() {
	for index := range queue.elements[:queue.size] {
		queue.elements[index] = queue.elements[index+1]
	}

	queue.size--
}
