package tree

import (
	"testing"

	comparator "github.com/dterbah/gods/utils"
	"github.com/stretchr/testify/assert"
)

func TestBinaryTreeAdd(t *testing.T) {
	assert := assert.New(t)
	tree := New(comparator.IntComparator)

	tree.Add(1)
	tree.Add(2)

	assert.True(tree.Has(1))
	assert.True(tree.Has(2))

	assert.False(tree.Has(4))
}

func TestBinaryTreeHas(t *testing.T) {
	assert := assert.New(t)
	tree := New(comparator.IntComparator)

	tree.Add(1)
	tree.Add(-1)
	tree.Add(-2)
	tree.Add(2)

	assert.True(tree.Has(1))
	assert.False(tree.Has(4))
	assert.True(tree.Has(-2))
}

func TestBinaryTreeMin(t *testing.T) {
	assert := assert.New(t)
	tree := New(comparator.IntComparator)

	_, err := tree.Min()
	assert.NotNil(err)
	tree.Add(1, 2, 3, 4)

	min, err := tree.Min()
	assert.Nil(err)
	assert.Equal(1, min)

	tree.Add(-10)
	min, err = tree.Min()

	assert.Nil(err)
	assert.Equal(-10, min)
}

func TestBinaryTreeMax(t *testing.T) {
	assert := assert.New(t)
	tree := New(comparator.IntComparator)

	_, err := tree.Max()
	assert.NotNil(err)
	tree.Add(-1, 0, 1, 2)
	max, err := tree.Max()

	assert.Nil(err)
	assert.Equal(2, max)
	tree.Add(10)
	max, err = tree.Max()
	assert.Nil(err)
	assert.Equal(10, max)
}

func TestBinaryTreeIteratorRight(t *testing.T) {
	assert := assert.New(t)
	tree := New(comparator.IntComparator)

	iterator := tree.Iterator()
	assert.False(iterator.HasRight())
	_, err := iterator.Right()
	assert.NotNil(err)

	tree.Add(1, 2, 3, -1)

	iterator = tree.Iterator()
	expectedValues := []int{2, 3}
	index := 0

	for iterator.HasRight() {
		currentValue, err := iterator.Right()
		assert.Nil(err)
		assert.Equal(expectedValues[index], currentValue)
		index++
	}

	_, err = iterator.Right()
	assert.NotNil(err)
}

func TestBinaryTreeIteratorLeft(t *testing.T) {
	assert := assert.New(t)
	tree := New(comparator.IntComparator)

	iterator := tree.Iterator()
	assert.False(iterator.HasLeft())
	_, err := iterator.Left()
	assert.NotNil(err)

	tree.Add(1, 2, 3, -1)

	expectedValues := []int{-1}
	index := 0

	iterator = tree.Iterator()

	for iterator.HasLeft() {
		currentValue, err := iterator.Left()
		assert.Nil(err)
		assert.Equal(expectedValues[index], currentValue)
		index++
	}
}

func TestBinaryTreeIteratorCurrent(t *testing.T) {
	assert := assert.New(t)
	tree := New(comparator.IntComparator)
	iterator := tree.Iterator()
	_, err := iterator.Current()
	assert.NotNil(err)
	tree.Add(1, 2, 3, -1)

	iterator = tree.Iterator()

	currentValue, err := iterator.Current()
	assert.Nil(err)
	assert.Equal(1, currentValue)

	iterator.Left()
	currentValue, err = iterator.Current()

	assert.Nil(err)
	assert.Equal(-1, currentValue)

	iterator.Right()
	currentValue, err = iterator.Current()
	assert.Nil(err)
	assert.Equal(-1, currentValue)
}

func TestBinaryTreeIteratorParent(t *testing.T) {
	assert := assert.New(t)
	tree := New(comparator.IntComparator)

	iterator := tree.Iterator()
	assert.False(iterator.HasParent())
	_, err := iterator.Parent()
	assert.NotNil(err)

	tree.Add(1)

	iterator = tree.Iterator()
	assert.False(iterator.HasParent())

	tree.Add(2, 3)

	for iterator.HasRight() {
		iterator.Right()
	}

	expectedValue := []int{2, 1}
	index := 0

	for iterator.HasParent() {
		currentValue, err := iterator.Parent()
		assert.Nil(err)
		assert.Equal(expectedValue[index], currentValue)
		index++
	}
}
