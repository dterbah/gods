package tree

import (
	"errors"

	comparator "github.com/dterbah/gods/utils"
)

type Node[T any] struct {
	left  *Node[T]
	right *Node[T]
	value T
}

/*
Struct that represents what is a BinaryTree
*/
type BinaryTree[T any] struct {
	root       *Node[T]
	comparator comparator.Comparator[T]
	zeroValue  T
}

/*
Create a new node for the tree
*/
func newNode[T any](value T) *Node[T] {
	return &Node[T]{value: value}
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
func (node *Node[T]) insertValue(value T, comparator comparator.Comparator[T]) {
	diff := comparator(node.value, value)

	if diff < 0 {
		// node.value < value, so create a node on the right
		if node.right == nil {
			node.right = newNode(value)
		} else {
			node.right.insertValue(value, comparator)
		}
	} else if diff > 0 {
		// node.value > value, so create a node on the left
		if node.left == nil {
			node.left = newNode(value)
		} else {
			node.left.insertValue(value, comparator)
		}
	}
}

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
			tree.root = newNode(value)
		} else {
			tree.root.insertValue(value, tree.comparator)
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
