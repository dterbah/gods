package queue

import (
	"testing"

	"github.com/dterbah/gods/list/arraylist"
	comparator "github.com/dterbah/gods/utils"
	"github.com/stretchr/testify/assert"
)

func TestQueueClear(t *testing.T) {
	assert := assert.New(t)
	queue := New[int](comparator.IntComparator)

	queue.Enqueue(1, 2, 3)
	queue.Clear()

	assert.True(queue.IsEmpty())
}

func TestQueueClone(t *testing.T) {
	assert := assert.New(t)
	queue := New[int](comparator.IntComparator)

	expectedValues := []int{1, 2}
	queue.Enqueue(expectedValues...)

	copy := queue.Copy()
	copy.ForEach(func(element, index int) {
		assert.Equal(expectedValues[index], element)
	})
}

func TestQueueContains(t *testing.T) {
	assert := assert.New(t)
	queue := New[int](comparator.IntComparator)

	queue.Enqueue(1, 2, 3)
	assert.True(queue.Contains(1))
	assert.True(queue.Contains(2))
	assert.True(queue.Contains(3))
	assert.False(queue.Contains(4))

}

func TestQueueDequeue(t *testing.T) {
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

func TestQueueEnqueue(t *testing.T) {
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

func TestQueueForEach(t *testing.T) {
	assert := assert.New(t)
	queue := New[int](comparator.IntComparator)

	expectedValues := []int{1, 2}
	queue.Enqueue(expectedValues...)

	queue.ForEach(func(element, index int) {
		assert.Equal(expectedValues[index], element)
	})
}

func TestQueueFromIterable(t *testing.T) {
	assert := assert.New(t)
	values := []int{1, 2, 3}
	list := arraylist.New(comparator.IntComparator, values...)
	queue := FromIterable(list, comparator.IntComparator)

	assert.Equal(3, queue.Size())

	for _, value := range values {
		assert.True(queue.Contains(value))
	}
}

func TestQueueIsEmpty(t *testing.T) {
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

func TestQueuePrint(t *testing.T) {
	queue := New[int](comparator.IntComparator)
	queue.Enqueue(1, 2, 3)

	queue.Print()
}

func TestQueuePeek(t *testing.T) {
	assert := assert.New(t)
	queue := New[int](comparator.IntComparator)

	_, err := queue.Peek()
	assert.NotNil(err)
	queue.Enqueue(1)
	value, err := queue.Peek()

	assert.Equal(1, value)
	assert.Nil(err)
}

func TestQueueSize(t *testing.T) {
	assert := assert.New(t)
	queue := New[int](comparator.IntComparator)
	assert.Equal(0, queue.Size())
	queue.Enqueue(1, 2, 4)
	assert.Equal(3, queue.Size())
}
