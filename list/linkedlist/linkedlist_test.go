package linkedlist

import (
	"testing"

	"github.com/dterbah/gods/list/arraylist"
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

func TestLinkedList_AddAll(t *testing.T) {
	assert := assert.New(t)
	list := New(comparator.IntComparator, 1, 2, 3)
	list.AddAll(New(comparator.IntComparator, 4, 5, 6))

	assert.Equal(6, list.Size())

	expectedValues := []int{1, 2, 3, 4, 5, 6}

	for index, element := range expectedValues {
		value, err := list.At(index)
		assert.Nil(err)
		assert.Equal(element, value)
	}
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

	// Empty list
	list = New(comparator.IntComparator)
	_, err = list.At(0)
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

	// error case
	list = New(comparator.IntComparator)
	assert.False(list.Contains(10))
}

func TestLinkedList_Copy(t *testing.T) {
	assert := assert.New(t)
	list := New(comparator.IntComparator, 1, 2, 3)

	copy := list.Copy()
	expectedValues := []int{1, 2, 3}

	copy.ForEach(func(element, index int) {
		assert.Equal(expectedValues[index], element)
	})
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

func TestLinkedList_ForEach(t *testing.T) {
	assert := assert.New(t)
	list := New(comparator.IntComparator, 1, 2, 3)
	expectedValues := []int{1, 2, 3}

	list.ForEach(func(element, index int) {
		assert.Equal(expectedValues[index], element)
	})
}

func TestLinkedList_FromIterable(t *testing.T) {
	assert := assert.New(t)
	list := arraylist.New(comparator.IntComparator, 1, 2, 3)

	linkedList := FromIterable[int](list, comparator.IntComparator)

	assert.Equal(3, linkedList.Size())

	expectedValues := []int{1, 2, 3}

	for index, element := range expectedValues {
		value, err := linkedList.At(index)
		assert.Nil(err)
		assert.Equal(element, value)
	}
}

func TestLinkedList_HeadTest(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator)
	head := list.Head()
	assert.Equal(0, head)

	list.Add(1, 2)

	assert.Equal(1, list.Head())
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

func TestLinkedList_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator)

	assert.True(list.IsEmpty())

	list.Add(1)
	assert.False(list.IsEmpty())
}

func TestLinkedList_nodeAt(t *testing.T) {
	assert := assert.New(t)
	list := New(comparator.IntComparator, 1, 2, 3)

	expectedValues := []int{1, 2, 3}

	for index, element := range expectedValues {
		node := list.nodeAt(index)
		assert.Equal(element, node.value)
	}

	// Out of bounds use case
	node := list.nodeAt(-1)
	assert.Nil(node)
	node = list.nodeAt(5)
	assert.Nil(node)
}

func TestLinkedList_Print(t *testing.T) {
	list := New(comparator.IntComparator, 1, 2, 3)
	list.Print()
}

func TestLinkedList_Remove(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator)
	list.Remove(1)
	list.Add(1)
	list.Remove(1)
	assert.Equal(0, list.Size())

	list.Add(1, 2, 3)

	list.Remove(1)
	assert.Equal(2, list.size)

	value, err := list.At(0)
	assert.Nil(err)
	assert.Equal(2, value)

	value, err = list.At(1)
	assert.Nil(err)
	assert.Equal(3, value)

	list.Remove(90)
	assert.Equal(2, list.Size())

	list = New(comparator.IntComparator, 1, 2, 3, 4, 5)
	list.Remove(3)
	assert.Equal(4, list.Size())

	expectedValues := []int{1, 2, 4, 5}
	list.ForEach(func(element, index int) {
		assert.Equal(expectedValues[index], element)
	})
}

func TestLinkedTest_RemoveAt(t *testing.T) {
	assert := assert.New(t)

	list := New[int](comparator.IntComparator)
	assert.False(list.RemoveAt(0))
	list.Add(1, 2, 3)
	assert.True(list.RemoveAt(0))
	assert.Equal(2, list.Size())

	list = New(comparator.IntComparator, 1, 2, 3, 4, 5)
	assert.True(list.RemoveAt(2))
	assert.Equal(4, list.Size())

	expectedValues := []int{1, 2, 4, 5}
	list.ForEach(func(element, index int) {
		assert.Equal(expectedValues[index], element)
	})
}

func TestLinkedList_ReplaceAt(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator, 1, 2, 3, 1)

	assert.False(list.ReplaceAt(-1, 1))
	assert.True(list.ReplaceAt(0, 10))

	value, err := list.At(0)
	assert.Nil(err)
	assert.Equal(10, value)

	assert.True(list.ReplaceAt(2, 9))
	value, err = list.At(2)
	assert.Nil(err)
	assert.Equal(9, value)

	assert.True(list.ReplaceAt(3, 90))
	value, err = list.At(3)
	assert.Nil(err)
	assert.Equal(90, value)
}

func TestLinkedList_Reverse(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator, 1, 2, 3)

	list.Reverse()

	assert.Equal(3, list.size)
	expectedValues := []int{3, 2, 1}

	for index, element := range expectedValues {
		value, err := list.At(index)
		assert.Nil(err)
		assert.Equal(element, value)
	}

	// empty list
	list = New(comparator.IntComparator)
	list.Reverse()
	assert.Equal(0, list.Size())
}

func TestLinkedList_Some(t *testing.T) {
	assert := assert.New(t)
	list := New(comparator.IntComparator, 1, 2, 3, 4)

	hasEven := list.Some(func(element, index int) bool {
		return element%2 == 0
	})

	assert.True(hasEven)

	hasNegativeNumber := list.Some(func(element, index int) bool {
		return element < 0
	})

	assert.False(hasNegativeNumber)
}

func TestLinkedList_Sort(t *testing.T) {
	assert := assert.New(t)
	list := New(comparator.IntComparator, 1, 4, 3, 2, 7, -1)
	list.Sort()
	expectedValues := []int{-1, 1, 2, 3, 4, 7}
	list.ForEach(func(element, index int) {
		assert.Equal(expectedValues[index], element)
	})
}

func TestLinkedList_SubList(t *testing.T) {
	assert := assert.New(t)
	list := New(comparator.IntComparator, 1, 2, 3)

	// error use case
	subList := list.SubList(-1, -10)
	assert.Equal(list, subList)

	subList = list.SubList(1, 0)
	assert.Equal(list, subList)

	// normal cases
	subList = list.SubList(1, 2)
	assert.Equal(2, subList.Size())

	value, err := subList.At(0)
	assert.Nil(err)
	assert.Equal(2, value)

	value, err = subList.At(1)
	assert.Nil(err)
	assert.Equal(3, value)
}

func TestLinkedList_TailTest(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator, 1, 2, 3)
	assert.Equal(3, list.Tail())
	list.Add(5)
	assert.Equal(5, list.Tail())

	// error case
	list = New(comparator.IntComparator)
	assert.Equal(0, list.Tail())
}
