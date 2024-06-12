package linkedlist

import (
	"errors"
	"fmt"

	"github.com/dterbah/gods/collection"
	"github.com/dterbah/gods/iterable"
	"github.com/dterbah/gods/list"
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
Create an arraylist from an iterable object
*/
func FromIterable[T any](iterable iterable.Iterable[T],
	comparator comparator.Comparator[T]) *LinkedList[T] {
	list := New(comparator)
	iterable.ForEach(func(element T, index int) {
		list.Add(element)
	})

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
	}

	list.size += len(elements)
}

func (list *LinkedList[T]) AddAll(elements collection.Collection[T]) {
	index := 0

	for index < elements.Size() {
		value, _ := elements.At(index)
		list.Add(value)
		index++
	}
}

/*
Retrieve an element by its index
If the index is negative or greater than the list size, the method will return an error
*/
func (list *LinkedList[T]) At(index int) (T, error) {
	if list.isOutOfBound(index) {
		return list.zeroElement, errors.New("index out of bounds")
	}

	currentNode := list.head

	for currentIndex := 0; currentIndex < index; currentIndex++ {
		currentNode = currentNode.next
	}

	return currentNode.value, nil
}

func (list LinkedList[T]) nodeAt(index int) *Node[T] {
	if list.isOutOfBound(index) {
		return nil
	}

	currentNode := list.head

	for currentIndex := 0; currentIndex < index; currentIndex++ {
		currentNode = currentNode.next
	}

	return currentNode
}

func (list *LinkedList[T]) Clear() {
	list.head = nil
	list.tail = nil
	list.size = 0
}

func (list LinkedList[T]) Contains(element T) bool {
	if list.head == nil {
		return false
	}

	for node := list.head; node != nil; node = node.next {
		if list.comparator(node.value, element) == 0 {
			return true
		}
	}

	return false
}

func (list LinkedList[T]) ContainsAll(collection collection.Collection[T]) bool {
	for index := 0; index < collection.Size(); index++ {
		value, _ := collection.At(index)
		if !list.Contains(value) {
			return false
		}
	}

	return true
}

func (list *LinkedList[T]) Copy() list.List[T] {
	newList := New[T](list.comparator)

	list.ForEach(func(element T, index int) {
		newList.Add(element)
	})

	return newList
}

/*
Return new list with elements matching the function passed in parameter
*/
func (list LinkedList[T]) Filter(callback func(element T) bool) list.List[T] {
	newList := New(list.comparator)

	for node := list.head; node != nil; node = node.next {
		if callback(node.value) {
			newList.Add(node.value)
		}
	}

	return newList
}

/*
Apply a function on each element of the list
*/
func (list LinkedList[T]) ForEach(callback func(element T, index int)) {
	if list.head != nil {
		index := 0
		for node := list.head; node != nil; node = node.next {
			callback(node.value, index)
			index++
		}
	}
}

/*
Return the value of the list's head
*/
func (list *LinkedList[T]) Head() T {
	if list.head == nil {
		return list.zeroElement
	} else {
		return list.head.value
	}
}

/*
Return the index of the specified element if present in the list.
If not, return -1
*/
func (list LinkedList[T]) IndexOf(element T) int {
	index := 0

	for node := list.head; node != nil; node = node.next {
		if list.comparator(element, node.value) == 0 {
			return index
		}

		index++
	}

	return -1
}

func (list LinkedList[T]) IsEmpty() bool {
	return list.head == nil && list.tail == nil
}

func (list LinkedList[T]) Print() {
	fmt.Print("[")
	for current := list.head; current != nil; current = current.next {
		fmt.Print(current.value)
		if current.next != nil {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

/*
Remove the first occurence of element in the list
*/
func (list *LinkedList[T]) Remove(element T) {
	if list.head == nil {
		return
	}

	if list.comparator(list.head.value, element) == 0 {
		list.head = list.head.next
		if list.head == nil {
			list.tail = nil
		}
		list.size--
		return
	}

	previous := list.head
	current := list.head.next

	for current != nil {
		if list.comparator(current.value, element) == 0 {
			previous.next = current.next
			if current == list.tail {
				list.tail = previous
			}
			list.size--
			return
		}
		previous = current
		current = current.next
	}
}

func (list *LinkedList[T]) RemoveAt(index int) bool {
	if list.isOutOfBound(index) || list.head == nil {
		return false
	}

	if index == 0 {
		list.head = list.head.next
		if list.head == nil {
			list.tail = nil
		}
		list.size--
		return true
	}

	previous := list.head
	current := list.head.next
	for currentIndex := 1; currentIndex < index; currentIndex++ {
		previous = current
		current = current.next
	}

	previous.next = current.next
	if current == list.tail {
		list.tail = previous
	}
	list.size--
	return true
}

func (list *LinkedList[T]) ReplaceAt(index int, element T) bool {
	if list.isOutOfBound(index) || list.head == nil {
		return false
	}

	currentIndex := 0
	node := list.head

	for ; currentIndex != index; node = node.next {
		currentIndex++
	}

	node.value = element

	return true
}

func (list *LinkedList[T]) Reverse() {
	if list.head == nil {
		return
	}

	newHead := newNode(list.tail.value)
	var currentNode *Node[T] = newHead

	for i := list.size - 2; i >= 0; i-- {
		value, _ := list.At(i)
		node := newNode(value)
		currentNode.next = node
		currentNode = node
	}

	list.tail = currentNode
	list.head = newHead
}

func (list LinkedList[T]) Some(callback func(element T, index int) bool) bool {
	result := false
	index := 0

	for node := list.head; node != nil; node = node.next {
		result = result || callback(node.value, index)
		index++
	}

	return result
}

func splitList[T any](head *Node[T]) (*Node[T], *Node[T]) {
	if head == nil || head.next == nil {
		return head, nil
	}

	slow := head
	fast := head
	var prev *Node[T]

	for fast != nil && fast.next != nil {
		prev = slow
		slow = slow.next
		fast = fast.next.next
	}

	prev.next = nil // Couper la liste en deux moitiés

	return head, slow
}

func mergeLists[T any](left *Node[T], right *Node[T], comparator comparator.Comparator[T]) *Node[T] {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	if comparator(left.value, right.value) == -1 {
		left.next = mergeLists(left.next, right, comparator)
		return left
	} else {
		right.next = mergeLists(left, right.next, comparator)
		return right
	}
}

func mergeSort[T any](head *Node[T], comparator comparator.Comparator[T]) *Node[T] {
	if head == nil || head.next == nil {
		return head
	}

	left, right := splitList(head)
	left = mergeSort(left, comparator)
	right = mergeSort(right, comparator)

	return mergeLists(left, right, comparator)
}

func (list *LinkedList[T]) Sort() {
	list.head = mergeSort[T](list.head, list.comparator)
}

func (list *LinkedList[T]) SubList(start, end int) list.List[T] {
	if (start < 0 || start > list.size) || (end < 0 || end > list.size) {
		return list
	}

	if start > end {
		return list
	}

	newList := New(list.comparator)
	index := 0

	for node := list.head; node != nil; node = node.next {
		if index >= start && index <= end {
			newList.Add(node.value)
		}
		index++
	}

	return newList
}

/*
Return the value of the list's tail (the last element)
*/
func (list *LinkedList[T]) Tail() T {
	if list.tail == nil {
		return list.zeroElement
	} else {
		return list.tail.value
	}
}

/*
Return the size of the list
*/
func (list *LinkedList[T]) Size() int {
	return list.size
}

// Private functions
func (list *LinkedList[T]) isOutOfBound(index int) bool {
	return index < 0 || index >= list.size
}
