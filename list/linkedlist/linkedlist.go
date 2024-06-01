package linkedlist

import (
	"errors"

	comparator "github.com/dterbah/gods/utils"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

func newNode[T any](value T) *Node[T] {
	return &Node[T]{value: value}
}

type LinkedList[T any] struct {
	head        *Node[T]
	tail        *Node[T]
	size        int
	zeroElement T
	comparator  comparator.Comparator[T]
}

func New[T any](comparator comparator.Comparator[T], elements ...T) *LinkedList[T] {
	var zero T
	list := &LinkedList[T]{head: nil, tail: nil, zeroElement: zero, comparator: comparator}
	for _, element := range elements {
		list.Add(element)
	}
	return list
}

/*
Add elements at the end of the list
*/
func (list *LinkedList[T]) Add(elements ...T) {
	for _, element := range elements {
		node := newNode(element)
		if list.head == nil {
			list.head = node
			list.tail = node
		} else {
			list.tail.next = node
			list.tail = node
		}
		list.size++
	}
}

func (list *LinkedList[T]) At(index int) (T, error) {
	if list.isOutOfBound(index) {
		return list.zeroElement, errors.New("index out of bounds")
	}

	if list.head == nil {
		return list.zeroElement, errors.New("The list doesn't contain any element")
	}

	currentNode := list.head

	for currentIndex := 0; currentIndex <= index; currentIndex++ {
		currentNode = currentNode.next
	}

	return currentNode.value, nil
}

/*
Return the size of the list
*/
func (list *LinkedList[T]) Size() int {
	return list.size
}

// Private functions
func (list *LinkedList[T]) isOutOfBound(index int) bool {
	return index < 0 || index > list.size
}
