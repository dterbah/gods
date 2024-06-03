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
	assert.Nil(err)

	value, err = list.At(3)
	assert.Equal(4, value)
	assert.Nil(err)

	// Index out of bound cases
	_, err = list.At(-1)
	assert.NotNil(err)

	_, err = list.At(5)
	assert.NotNil(err)
}

func TestLinkedList_ClearTest(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator, 1, 2, 3)

	list.Clear()

	assert.Equal(0, list.Size())
}

func TestLinkedList_Contains(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator, 1, 2, 3)

	assert.True(list.Contains(1))

	assert.False(list.Contains(4))
}

func TestLinkedList_Filter(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator, 1, 2, 3)

	newList := list.Filter(func(element int) bool {
		return element%2 == 0
	})

	assert.Equal(1, newList.Size())

	value, err := newList.At(0)
	assert.Nil(err)
	assert.Equal(2, value)
}

func TestLinkedList_HeadTest(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator, 1, 2, 3)

	index := list.IndexOf(1)
	assert.Equal(0, index)
	index = list.IndexOf(3)
	assert.Equal(2, index)

	index = list.IndexOf(10)
	assert.Equal(-1, index)

	list = New(comparator.IntComparator)
	index = list.IndexOf(9)
	assert.Equal(-1, index)
}

func TestLinkedList_IndexOf(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator)

	assert.Equal(list.IndexOf(1), -1)

	list.Add(1, 2)

	assert.Equal(0, list.IndexOf(1))
	assert.Equal(1, list.IndexOf(2))
	assert.Equal(-1, list.IndexOf(4))
}

func TestLinkedList_TailTest(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator, 1, 2, 3)
	assert.Equal(3, list.Tail())
	list.Add(5)
	assert.Equal(5, list.Tail())
}
