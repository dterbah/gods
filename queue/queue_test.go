package queue

import (
	"testing"

	comparator "github.com/dterbah/gods/utils"
	"github.com/stretchr/testify/assert"
)

func TestQueue_Clear(t *testing.T) {
	assert := assert.New(t)
	queue := New[int](comparator.IntComparator)

	queue.Enqueue(1, 2, 3)
	queue.Clear()

	assert.True(queue.IsEmpty())
}

func TestQueue_Clone(t *testing.T) {
	assert := assert.New(t)
	queue := New[int](comparator.IntComparator)

	expectedValues := []int{1, 2}
	queue.Enqueue(expectedValues...)

	copy := queue.Copy()
	copy.ForEach(func(element, index int) {
		assert.Equal(expectedValues[index], element)
	})
}

func TestQueue_Contains(t *testing.T) {
	assert := assert.New(t)
	queue := New[int](comparator.IntComparator)

	queue.Enqueue(1, 2, 3)
	assert.True(queue.Contains(1))
	assert.True(queue.Contains(2))
	assert.True(queue.Contains(3))
	assert.False(queue.Contains(4))

}

func TestQueue_Dequeue(t *testing.T) {
	assert := assert.New(t)
	queue := New[int](comparator.IntComparator)

	queue.Enqueue(1, 2, 3, 4)
	expectedValues := []int{1, 2, 3, 4}

	for _, value := range expectedValues {
		currentValue, err := queue.Dequeue()
		assert.Nil(err)
		assert.Equal(value, currentValue)
	}

	_, err := queue.Dequeue()
	assert.NotNil(err)
}

func TestQueue_Enqueue(t *testing.T) {
	assert := assert.New(t)
	queue := New[int](comparator.IntComparator)

	queue.Enqueue(1, 2, 3, 4)

	expectedValues := []int{1, 2, 3, 4}

	for _, value := range expectedValues {
		currentValue, err := queue.Dequeue()
		assert.Nil(err)
		assert.Equal(value, currentValue)
	}
}

func TestQueue_ForEach(t *testing.T) {
	assert := assert.New(t)
	queue := New[int](comparator.IntComparator)

	expectedValues := []int{1, 2}
	queue.Enqueue(expectedValues...)

	queue.ForEach(func(element, index int) {
		assert.Equal(expectedValues[index], element)
	})
}

func TestQueue_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	queue := New[int](comparator.IntComparator)
	assert.True(queue.IsEmpty())
	queue.Enqueue(1, 2)
	assert.False(queue.IsEmpty())
	queue.Dequeue()
	assert.False(queue.IsEmpty())
	queue.Dequeue()
	assert.True(queue.IsEmpty())
}

func TestQueue_Print(t *testing.T) {
	queue := New[int](comparator.IntComparator)
	queue.Enqueue(1, 2, 3)

	queue.Print()
}

func TestQueue_Peek(t *testing.T) {
	assert := assert.New(t)
	queue := New[int](comparator.IntComparator)

	_, err := queue.Peek()
	assert.NotNil(err)
	queue.Enqueue(1)
	value, err := queue.Peek()

	assert.Equal(1, value)
	assert.Nil(err)
}

func TestQueue_Size(t *testing.T) {
	assert := assert.New(t)
	queue := New[int](comparator.IntComparator)
	assert.Equal(0, queue.Size())
	queue.Enqueue(1, 2, 4)
	assert.Equal(3, queue.Size())
}
