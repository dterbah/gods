package stack

import (
	"testing"

	comparator "github.com/dterbah/gods/utils"
	"github.com/stretchr/testify/assert"
)

func TestStack_Clear(t *testing.T) {
	assert := assert.New(t)
	stack := New(comparator.IntComparator)
	stack.Push(1, 2, 3)
	stack.Clear()
	assert.True(stack.IsEmpty())
}

func TestStack_Push(t *testing.T) {
	assert := assert.New(t)
	stack := New(comparator.IntComparator)
	stack.Push(1, 2, 3)
	assert.Equal(3, stack.Size())
	assert.False(stack.IsEmpty())
}

func TestStack_Contains(t *testing.T) {
	assert := assert.New(t)
	stack := New(comparator.IntComparator)
	stack.Push(1, 2, 3)
	values := []int{1, 2, 3}

	for _, value := range values {
		assert.True(stack.Contains(value))
	}

	assert.False(stack.Contains(4))
}

func TestStack_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	stack := New(comparator.IntComparator)

	assert.True(stack.IsEmpty())
	stack.Push(1)
	assert.False(stack.IsEmpty())
}

func TestStack_Peek(t *testing.T) {
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

func TestStack_Pop(t *testing.T) {
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

func TestStack_Print(t *testing.T) {
	stack := New(comparator.IntComparator)
	stack.Push(1, 2, 3)
	stack.Print()
}

func TestStack_Size(t *testing.T) {
	assert := assert.New(t)
	stack := New(comparator.IntComparator)
	assert.Equal(0, stack.Size())

	stack.Push(1, 2)
	assert.Equal(2, stack.Size())
}
