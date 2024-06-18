package tree

import (
	"errors"

	comparator "github.com/dterbah/gods/utils"
)

type Node[T any] struct {
	left   *Node[T]
	right  *Node[T]
	parent *Node[T]
	value  T
}

/*
Struct that represents what is a BinaryTree
*/
type BinaryTree[T any] struct {
	root       *Node[T]
	comparator comparator.Comparator[T]
	zeroValue  T
}

type BinaryTreeIterator[T any] struct {
	zeroValue   T
	currentNode *Node[T]
}

// ---- BinaryTreeIterator API ---- //
func newIterator[T any](tree BinaryTree[T]) *BinaryTreeIterator[T] {
	return &BinaryTreeIterator[T]{currentNode: tree.root, zeroValue: tree.zeroValue}
}

/*
Return true if the current node of the iterator has a right node, else false
*/
func (iterator BinaryTreeIterator[T]) HasRight() bool {
	return iterator.currentNode != nil && iterator.currentNode.right != nil
}

/*
Return true if the current node of the iterator has a left node, else false
*/
func (iterator BinaryTreeIterator[T]) HasLeft() bool {
	return iterator.currentNode != nil && iterator.currentNode.left != nil
}

/*
Return true if the current node of the iterator has a prent, else false
*/
func (iterator BinaryTreeIterator[T]) HasParent() bool {
	return iterator.currentNode != nil && iterator.currentNode.parent != nil
}

/*
Return the left value of the current node of the iterator. The iterator
will automatically move to the left element
*/
func (iterator *BinaryTreeIterator[T]) Left() (T, error) {
	if iterator.currentNode == nil || !iterator.HasLeft() {
		return iterator.zeroValue, errors.New("no left value available")
	}

	iterator.currentNode = iterator.currentNode.left

	return iterator.currentNode.value, nil
}

/*
Return the right value of the current node of the iterator. The iterator
will automatically move to the right element
*/
func (iterator *BinaryTreeIterator[T]) Right() (T, error) {
	if iterator.currentNode == nil || !iterator.HasRight() {
		return iterator.zeroValue, errors.New("no right value available")
	}

	iterator.currentNode = iterator.currentNode.right

	return iterator.currentNode.value, nil
}

/*
Return the parent value of the current node of the iterator. The iterator
will automatically move to the parent element
*/
func (iterator *BinaryTreeIterator[T]) Parent() (T, error) {
	if iterator.currentNode == nil || !iterator.HasParent() {
		return iterator.zeroValue, errors.New("no parent value available")
	}

	iterator.currentNode = iterator.currentNode.parent

	return iterator.currentNode.value, nil
}

/*
Return the current value of the current node of the iterator.
*/
func (iterator *BinaryTreeIterator[T]) Current() (T, error) {
	if iterator.currentNode == nil {
		return iterator.zeroValue, errors.New("no current value available")
	}

	return iterator.currentNode.value, nil
}

// ---- Node API ---- //

/*
Create a new node for the tree
*/
func newNode[T any](value T, parent *Node[T]) *Node[T] {
	return &Node[T]{value: value, parent: parent}
}

func (node Node[T]) hasValue(value T, comparator comparator.Comparator[T]) bool {
	if comparator(node.value, value) == 0 {
		return true
	}

	res := false

	if node.left != nil {
		res = res || node.left.hasValue(value, comparator)
	}

	if node.right != nil {
		res = res || node.right.hasValue(value, comparator)
	}

	return res
}

/*
Insert a value in a node
*/
func (node *Node[T]) insertValue(value T, comparator comparator.Comparator[T], parent *Node[T]) {
	diff := comparator(node.value, value)

	if diff < 0 {
		// node.value < value, so create a node on the right
		if node.right == nil {
			node.right = newNode(value, node)
		} else {
			node.right.insertValue(value, comparator, node)
		}
	} else if diff > 0 {
		// node.value > value, so create a node on the left
		if node.left == nil {
			node.left = newNode(value, node)
		} else {
			node.left.insertValue(value, comparator, node)
		}
	}
}

// ---- BinaryTree API ---- //

/*
Create a new Binary tree
*/
func New[T any](comparator comparator.Comparator[T]) *BinaryTree[T] {
	var zero T
	return &BinaryTree[T]{comparator: comparator, zeroValue: zero}
}

/*
Add values in the Tree
*/
func (tree *BinaryTree[T]) Add(values ...T) {
	for _, value := range values {
		if tree.root == nil {
			tree.root = newNode(value, nil)
		} else {
			tree.root.insertValue(value, tree.comparator, tree.root)
		}
	}
}

/*
Check if a value is present in the tree. Return true if the value is present,
else false
*/
func (tree BinaryTree[T]) Has(value T) bool {
	return tree.root.hasValue(value, tree.comparator)
}

func (tree BinaryTree[T]) Iterator() *BinaryTreeIterator[T] {
	return newIterator[T](tree)
}

/*
Find the maximum value present in the tree
*/
func (tree BinaryTree[T]) Max() (T, error) {
	if tree.root == nil {
		return tree.zeroValue, errors.New("empty tree")
	}

	max := tree.root.value

	// visit all the right nodes

	currentNode := tree.root.right

	for currentNode != nil {
		max = currentNode.value
		currentNode = currentNode.right
	}

	return max, nil
}

/*
Find the minimum value present in the tree
*/
func (tree BinaryTree[T]) Min() (T, error) {
	if tree.root == nil {
		return tree.zeroValue, errors.New("empty tree")
	}

	min := tree.root.value

	// visit all the left nodes
	currentNode := tree.root.left

	for currentNode != nil {
		min = currentNode.value
		currentNode = currentNode.left
	}

	return min, nil
}

// MAP .??????
