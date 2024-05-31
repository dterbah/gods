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

func TestArrayList_Concat(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator)
	list2 := New[int](comparator.IntComparator)

	list.Add(1, 2, 3)
	list2.Add(4, 5, 6)

	expectedElements := []int{1, 2, 3, 4, 5, 6}

	newList := list.Concat(list2)

	assert.Equal(6, newList.Size())

	newList.ForEach(func(element, index int) {
		assert.Equal(expectedElements[index], element)
	})
}

func TestArrayList_Contains(t *testing.T) {
	assert := assert.New(t)

	list := New[int](comparator.IntComparator)

	list.Add(1, 2, 3)
	assert.True(list.Contains(1))
	assert.False(list.Contains(-190))
}

func TestArrayList_Filter(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator)

	list.Add(1, 2, 3, 6)

	newList := list.Filter(func(element int) bool {
		// event number filtering
		return element%2 == 0
	})

	assert.Equal(2, newList.Size())
	value, _ := newList.At(0)

	assert.Equal(2, value)

	value, _ = newList.At(1)
	assert.Equal(6, value)
}

func TestArrayList_IndexOf(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator)

	list.Add(1, 2, 3)

	index := list.IndexOf(3)
	expectedValue := 2
	assert.Equal(expectedValue, index)

	index = list.IndexOf(10)
	expectedValue = -1
	assert.Equal(expectedValue, index)
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

func TestArrayList_RemoveAt(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator)

	list.Add(1, 2, 3)

	isRemoved := list.RemoveAt(1)
	assert.True(isRemoved)
	assert.Equal(2, list.size)
	value, _ := list.At(1)
	assert.Equal(value, 3)

	isRemoved = list.RemoveAt(0)
	assert.True(isRemoved)
	assert.Equal(1, list.size)
	value, _ = list.At(0)
	assert.Equal(value, 3)

	// index out of bound cases
	isRemoved = list.RemoveAt(-1)
	assert.False(isRemoved)

	isRemoved = list.RemoveAt(1)
	assert.False(isRemoved)
}

func TestArrayList_ReplaceAt(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator)
	expectedValue := 90

	list.Add(1, 2, 3)
	isReplaced := list.ReplaceAt(0, expectedValue)
	value, _ := list.At(0)

	assert.True(isReplaced)
	assert.Equal(expectedValue, value)

	// check replace with invalid values (negative and > list size)
	isReplaced = list.ReplaceAt(-1, 0)
	assert.False(isReplaced)

	isReplaced = list.ReplaceAt(list.Size(), 90)
	assert.False(isReplaced)
}

func TestArrayList_Reverse(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator)

	list.Add(1, 2, 3)
	list.Reverse()

	expectedValues := []int{3, 2, 1}

	list.ForEach(func(element, index int) {
		value, _ := list.At(index)
		assert.Equal(expectedValues[index], value)
	})
}

func TestArrayList_Size(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator)
	list.Add(1, 2, 3)
	expectedSize := 3
	assert.Equal(expectedSize, list.Size())
}

func TestArrayList_Some(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator)
	list.Add(1, 2, 3)

	assert.True(list.Some(func(element, index int) bool {
		return element > 2
	}))

	assert.False(list.Some(func(element, index int) bool {
		return element < 0
	}))
}

func TestArrayList_Sort(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator)

	list.Add(3, 2, 1, 9, -1)
	list.Sort()

	expectedValues := []int{-1, 1, 2, 3, 9}

	list.ForEach(func(element, index int) {
		value, _ := list.At(index)
		assert.Equal(expectedValues[index], value)
	})
}
