package linkedlist

import (
	"testing"

	"github.com/dterbah/gods/list/arraylist"
	comparator "github.com/dterbah/gods/utils"
	"github.com/stretchr/testify/assert"
)

func TestLinkedListAddTest(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator)

	list.Add(1, 2)
	assert.Equal(2, list.Size())

	assert.Equal(1, list.head.value)
	assert.Equal(2, list.tail.value)
}

func TestLinkedListAddAll(t *testing.T) {
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

func TestLinkedListAt(t *testing.T) {
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

func TestLinkedListClear(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator, 1, 2, 3)

	list.Clear()

	assert.Equal(0, list.Size())
}

func TestLinkedListContains(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator, 1, 2, 3)

	assert.True(list.Contains(1))

	assert.False(list.Contains(4))

	// error case
	list = New(comparator.IntComparator)
	assert.False(list.Contains(10))
}

func TestLinkedListContainsAll(t *testing.T) {
	assert := assert.New(t)
	list := New(comparator.IntComparator, 1, 2, 3)
	collection := New(comparator.IntComparator, 1, 2)

	assert.True(list.ContainsAll(collection))
	collection.Add(4, 5)
	assert.False(list.ContainsAll(collection))
}

func TestLinkedListCopy(t *testing.T) {
	assert := assert.New(t)
	list := New(comparator.IntComparator, 1, 2, 3)

	copy := list.Copy()
	expectedValues := []int{1, 2, 3}

	copy.ForEach(func(element, index int) {
		assert.Equal(expectedValues[index], element)
	})
}

func TestLinkedListFilter(t *testing.T) {
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

func TestLinkedListForEach(t *testing.T) {
	assert := assert.New(t)
	list := New(comparator.IntComparator, 1, 2, 3)
	expectedValues := []int{1, 2, 3}

	list.ForEach(func(element, index int) {
		assert.Equal(expectedValues[index], element)
	})
}

func TestLinkedListFromIterable(t *testing.T) {
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

func TestLinkedListHead(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator)
	head := list.Head()
	assert.Equal(0, head)

	list.Add(1, 2)

	assert.Equal(1, list.Head())
}

func TestLinkedListIndexOf(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator)

	assert.Equal(list.IndexOf(1), -1)

	list.Add(1, 2)

	assert.Equal(0, list.IndexOf(1))
	assert.Equal(1, list.IndexOf(2))
	assert.Equal(-1, list.IndexOf(4))
}

func TestLinkedListIsEmpty(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator)

	assert.True(list.IsEmpty())

	list.Add(1)
	assert.False(list.IsEmpty())
}

func TestLinkedListNodeAt(t *testing.T) {
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

func TestLinkedListPrint(t *testing.T) {
	list := New(comparator.IntComparator, 1, 2, 3)
	list.Print()
}

func TestLinkedListRemove(t *testing.T) {
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

func TestLinkedTestRemoveAt(t *testing.T) {
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

func TestLinkedListReplaceAt(t *testing.T) {
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

func TestLinkedListReverse(t *testing.T) {
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

func TestLinkedListSome(t *testing.T) {
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

func TestLinkedListSort(t *testing.T) {
	assert := assert.New(t)
	list := New(comparator.IntComparator, 1, 4, 3, 2, 7, -1)
	list.Sort()
	expectedValues := []int{-1, 1, 2, 3, 4, 7}
	list.ForEach(func(element, index int) {
		assert.Equal(expectedValues[index], element)
	})
}

func TestLinkedListSubList(t *testing.T) {
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

func TestLinkedListTailTest(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator, 1, 2, 3)
	assert.Equal(3, list.Tail())
	list.Add(5)
	assert.Equal(5, list.Tail())

	// error case
	list = New(comparator.IntComparator)
	assert.Equal(0, list.Tail())
}

func TestLinkedListToArray(t *testing.T) {
	assert := assert.New(t)
	list := New[int](comparator.IntComparator, 1, 2, 3)

	elements := list.ToArray()

	assert.Equal(3, len(elements))
	expectedElements := []int{1, 2, 3}

	for index, el := range expectedElements {
		assert.Equal(el, elements[index])
	}
}
