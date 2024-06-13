package queue

import (
	"errors"
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
}

/*
Create a new Queue
*/
func New[T any]() *Queue[T] {
	var zero T
	return &Queue[T]{elements: []T{}, zeroElement: zero}
}

/*
Remove all the elements of the queue
*/
func (queue *Queue[T]) Clear() {
	queue.size = 0
	queue.elements = []T{}
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
