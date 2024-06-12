package circular

import (
	"errors"
)

/*
Struct used to defines what is a CircularBuffer
*/
type CircularBuffer[T any] struct {
	size         int
	readPointer  int
	writePointer int
	elements     []T
	full         bool
	zeroElement  T
}

/*
Create a new CircularBuffer with a specified size
*/
func New[T any](size int) *CircularBuffer[T] {
	if size < 0 {
		return nil
	}

	var zero T

	return &CircularBuffer[T]{
		size:        size,
		elements:    make([]T, size),
		full:        false,
		zeroElement: zero,
	}
}

/*
Enqueue new element inside the buffer. If the buffer is already full,
this method will return an error
*/
func (buffer *CircularBuffer[T]) Enqueue(element T) error {
	if buffer.full {
		return errors.New("circular buffer full")
	}

	buffer.elements[buffer.writePointer] = element
	buffer.writePointer = (buffer.writePointer + 1) % buffer.size

	if buffer.writePointer == buffer.readPointer {
		buffer.full = true
	}

	return nil
}

/*
Dequeue an element. It returns the element and nil if the buffer is not empty.
Otherwise, it returns a zero element and an error
*/
func (buffer *CircularBuffer[T]) Dequeue() (T, error) {
	if !buffer.full && buffer.readPointer == buffer.writePointer {
		return buffer.zeroElement, errors.New("buffer empty")
	}

	element := buffer.elements[buffer.readPointer]
	buffer.readPointer = (buffer.readPointer + 1) % buffer.size
	buffer.full = false

	return element, nil
}

/*
Return true if the buffer is empty, else false
*/
func (buffer CircularBuffer[T]) IsEmpty() bool {
	return !buffer.full && buffer.readPointer == buffer.writePointer
}

/*
Return true if the buffer is full, else false
*/
func (buffer CircularBuffer[T]) IsFull() bool {
	return buffer.full
}
