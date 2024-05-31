package arraylist

import (
	"testing"

	comparator "github.com/dterbah/gods/utils"
	"github.com/stretchr/testify/assert"
)

func TestArrayList_Add(t *testing.T) {
	assert := assert.New(t)

	list := New[int](comparator.IntComparator)

	list.Add(1, 2, 3)

	assert.Equal(3, list.size, "Incorrect list size. Expected 3, Received ", list.size)

	for i, element := range list.elements[:list.size] {
		expectedValue := i + 1

		assert.Equal(expectedValue, element, "Incorrect")

		if expectedValue != element {
			assert.Equal(expectedValue, element, "Incorrect value at index ", i)
		}
	}
}

func TestArrayList_At(t *testing.T) {
	assert := assert.New(t)

	list := New[int](comparator.IntComparator)

	list.Add(1, 2, 3)

	expectedValue := 2
	value, err := list.At(1)

	assert.Equal(err, nil, "The list shouldn't return any error")
	assert.Equal(expectedValue, value)

	// Test a value out of bounds
	_, err = list.At(3)

	assert.NotEqual(err, nil)

	// Test negative index
	_, err = list.At(-1)

	assert.NotEqual(err, nil)
}

func TestArrayList_Clear(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator)

	list.Add(10, 20, 50)
	assert.False(list.IsEmpty())
	list.Clear()
	assert.True(list.IsEmpty())
}

func TestArrayListy_Contains(t *testing.T) {
	assert := assert.New(t)

	list := New[int](comparator.IntComparator)

	list.Add(1, 2, 3)
	assert.True(list.Contains(1))
	assert.False(list.Contains(-190))
}

func TestArrayList_Size(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator)
	list.Add(1, 2, 3)
	expectedSize := 3
	assert.Equal(expectedSize, list.Size())
}

func TestArrayList_IsEmpty(t *testing.T) {
	assert := assert.New(t)

	list := New[int](comparator.IntComparator)
	assert.True(list.IsEmpty())
	list.Add(1)
	assert.False(list.IsEmpty())
	list.Add(2, 3, 4)
	assert.False(list.IsEmpty())
}
