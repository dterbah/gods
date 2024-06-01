package linkedlist

import (
	"testing"

	comparator "github.com/dterbah/gods/utils"
	"github.com/stretchr/testify/assert"
)

func TestLinkedList_AddTest(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator)

	list.Add(1, 2)
	assert.Equal(2, list.Size())

	assert.Equal(1, list.head.value)
	assert.Equal(2, list.tail.value)
}

func TestLinkedList_AtTest(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator)

	list.Add(1, 2, 3, 4, 5)

	value, err := list.At(0)
	assert.Equal(1, value)
	assert.NotNil(err)

	value, err = list.At(3)
	assert.Equal(4, value)
	assert.NotNil(err)

	// Index out of bound cases
	_, err = list.At(-1)
	assert.Nil(err)

	_, err = list.At(5)
	assert.Nil(err)
}
