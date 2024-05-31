package arraylist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayList_Add(t *testing.T) {
	assert := assert.New(t)

	list := New[int]()

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

	list := New[int]()

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

func TestArrayList_Size(t *testing.T) {
	assert := assert.New(t)
	list := New[int]()
	list.Add(1, 2, 3)
	expectedSize := 3
	assert.Equal(expectedSize, list.Size())
}

func TestArrayList_IsEmpty(t *testing.T) {
	assert := assert.New(t)

	list := New[int]()
	assert.True(list.IsEmpty())
	list.Add(1)
	assert.False(list.IsEmpty())
	list.Add(2, 3, 4)
	assert.False(list.IsEmpty())
}
