package stack

import (
	"testing"

	"github.com/dterbah/gods/list/arraylist"
	comparator "github.com/dterbah/gods/utils"
	"github.com/stretchr/testify/assert"
)

func TestStackClear(t *testing.T) {
	assert := assert.New(t)
	stack := New(comparator.IntComparator)
	stack.Push(1, 2, 3)
	stack.Clear()
	assert.True(stack.IsEmpty())
}

func TestStackCopy(t *testing.T) {
	assert := assert.New(t)
	stack := New(comparator.IntComparator)
	stack.Push(1, 2, 3)

	newStack := stack.Copy()
	assert.Equal(3, newStack.Size())

	values := []int{3, 2, 1}

	for _, value := range values {
		currentValue, err := stack.Pop()
		assert.Nil(err)
		assert.Equal(value, currentValue)
	}
}

func TestStackFromIterable(t *testing.T) {
	assert := assert.New(t)
	list := arraylist.New(comparator.IntComparator, 1, 2, 3)
	stack := FromIterable(list, comparator.IntComparator)

	assert.Equal(3, stack.Size())

	values := []int{1, 2, 3}

	for _, value := range values {
		assert.True(stack.Contains(value))
	}
}

func TestStackPush(t *testing.T) {
	assert := assert.New(t)
	stack := New(comparator.IntComparator)
	stack.Push(1, 2, 3)
	assert.Equal(3, stack.Size())
	assert.False(stack.IsEmpty())
}

func TestStackContains(t *testing.T) {
	assert := assert.New(t)
	stack := New(comparator.IntComparator)
	stack.Push(1, 2, 3)
	values := []int{1, 2, 3}

	for _, value := range values {
		assert.True(stack.Contains(value))
	}

	assert.False(stack.Contains(4))
}

func TestStackIsEmpty(t *testing.T) {
	assert := assert.New(t)
	stack := New(comparator.IntComparator)

	assert.True(stack.IsEmpty())
	stack.Push(1)
	assert.False(stack.IsEmpty())
}

func TestStackPeek(t *testing.T) {
	assert := assert.New(t)
	stack := New(comparator.IntComparator)

	_, err := stack.Peek()
	assert.NotNil(err)

	stack.Push(1)
	value, err := stack.Peek()

	assert.False(stack.IsEmpty())
	assert.Equal(1, stack.Size())
	assert.Equal(1, value)
	assert.Nil(err)

	stack.Push(2, 3, 4)
	value, err = stack.Peek()

	assert.False(stack.IsEmpty())
	assert.Equal(4, stack.Size())
	assert.Equal(4, value)
	assert.Nil(err)
}

func TestStackPop(t *testing.T) {
	assert := assert.New(t)
	stack := New(comparator.IntComparator)

	_, err := stack.Pop()
	assert.NotNil(err)

	stack.Push(1)
	value, err := stack.Pop()

	assert.True(stack.IsEmpty())
	assert.Equal(0, stack.Size())
	assert.Equal(1, value)
	assert.Nil(err)

	stack.Push(2, 3, 4)
	value, err = stack.Pop()

	assert.False(stack.IsEmpty())
	assert.Equal(2, stack.Size())
	assert.Equal(4, value)
	assert.Nil(err)
}

func TestStackPrint(t *testing.T) {
	stack := New(comparator.IntComparator)
	stack.Push(1, 2, 3)
	stack.Print()
}

func TestStackSize(t *testing.T) {
	assert := assert.New(t)
	stack := New(comparator.IntComparator)
	assert.Equal(0, stack.Size())

	stack.Push(1, 2)
	assert.Equal(2, stack.Size())
}
