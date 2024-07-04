package set

import (
	"testing"

	"github.com/dterbah/gods/collection"
	"github.com/dterbah/gods/list/arraylist"
	comparator "github.com/dterbah/gods/utils"
	"github.com/stretchr/testify/assert"
)

func TestSetAdd(t *testing.T) {
	assert := assert.New(t)

	set := New[int](comparator.IntComparator)

	set.Add(1, 2, 3)

	assert.Equal(3, set.Size())

	set.Add(1)

	assert.Equal(3, set.Size())
}

func TestSetAddAll(t *testing.T) {
	assert := assert.New(t)
	var collection collection.Collection[int] = arraylist.New(comparator.IntComparator, 5, 6)
	set := New[int](comparator.IntComparator)
	set.AddAll(collection)

	assert.Equal(2, set.Size())

	expectedElements := []int{5, 6}

	for index, element := range expectedElements {
		value, err := set.At(index)
		assert.Nil(err)
		assert.Equal(element, value)
	}
}

func TestSetAt(t *testing.T) {
	assert := assert.New(t)
	set := New[int](comparator.IntComparator, 1, 2, 4)

	value, err := set.At(0)
	assert.Equal(1, value)
	assert.Nil(err)

	_, err = set.At(-1)
	assert.NotNil(err)

	_, err = set.At(3)
	assert.NotNil(err)
}

func TestSetClear(t *testing.T) {
	assert := assert.New(t)
	set := New[int](comparator.IntComparator, 1, 3, 4)
	set.Clear()

	assert.Equal(0, set.Size())
}

func TestSetContains(t *testing.T) {
	assert := assert.New(t)
	set := New[int](comparator.IntComparator, 1, 3, 4)

	assert.True(set.Contains(1))
	assert.True(set.Contains(3))
	assert.True(set.Contains(4))
	assert.False(set.Contains(2))
}

func TestSetContainsAll(t *testing.T) {
	assert := assert.New(t)
	set := New(comparator.IntComparator, 1, 2, 3, 4)
	collection := arraylist.New(comparator.IntComparator, 1, 2, 5)

	assert.False(set.ContainsAll(collection))
	assert.True(set.ContainsAll(set))

	collection = arraylist.New(comparator.IntComparator, 1, 2)
	assert.True(set.ContainsAll(collection))
}

func TestSetCopy(t *testing.T) {
	assert := assert.New(t)
	set := New(comparator.IntComparator, 1, 3)
	copy := set.Copy()

	assert.True(set.IsSubset(copy))
	set.ForEach(func(element, index int) {
		value, err := copy.At(index)
		assert.Nil(err)
		assert.Equal(element, value)
	})
}

func TestSetDiff(t *testing.T) {
	assert := assert.New(t)
	set := New(comparator.IntComparator, 1, 2, 3, 6, 9)
	otherSet := New(comparator.IntComparator, 1, 2, 5, 6)

	result := set.Diff(otherSet)
	assert.Equal(2, result.Size())

	expectedValues := []int{3, 9}

	for _, element := range expectedValues {
		assert.True(result.Contains(element))
	}
}

func TestSetIndexOf(t *testing.T) {
	assert := assert.New(t)
	set := New(comparator.IntComparator, 1, 2, 3, 4)

	assert.Equal(0, set.IndexOf(1))
	assert.Equal(3, set.IndexOf(4))
	assert.Equal(-1, set.IndexOf(9))
}

func TestSetIntersection(t *testing.T) {
	assert := assert.New(t)
	set := New(comparator.IntComparator, 1, 2, 3, 6, 9)
	otherSet := New(comparator.IntComparator, 1, 2, 5, 6)

	result := set.Intersection(otherSet)
	expectedValues := []int{1, 2, 6}

	assert.Equal(3, result.Size())

	for _, element := range expectedValues {
		assert.True(result.Contains(element))
	}
}

func TestSetIsEmpty(t *testing.T) {
	assert := assert.New(t)

	set := New[int](comparator.IntComparator)

	assert.True(set.IsEmpty())

	set.Add(1, 2)

	assert.False(set.IsEmpty())
}

func TestSetIsSubset(t *testing.T) {
	assert := assert.New(t)
	set := New(comparator.IntComparator, 1, 2, 3, 6, 9)
	otherSet := New(comparator.IntComparator, 1, 2)

	assert.True(set.IsSubset(otherSet))

	set = New(comparator.IntComparator, 1)
	assert.False(set.IsSubset(otherSet))

}

func TestSetPrint(t *testing.T) {
	set := New(comparator.IntComparator, 1, 2, 3)
	set.Print()
}

func TestSetRemove(t *testing.T) {
	assert := assert.New(t)
	set := New(comparator.IntComparator, 1, 2, 3, 6, 9)

	set.Remove(1)

	assert.Equal(4, set.Size())

	set.Remove(10)

	assert.Equal(4, set.Size())
}

func TestSetToArray(t *testing.T) {
	assert := assert.New(t)
	set := New[int](comparator.IntComparator, 1, 2, 4)

	elements := set.ToArray()
	expectedElements := []int{1, 2, 4}

	assert.Equal(3, len(elements))

	for index, el := range expectedElements {
		assert.Equal(el, elements[index])
	}
}

func TestSetUnion(t *testing.T) {
	assert := assert.New(t)
	set := New(comparator.IntComparator, 1, 2, 3, 6, 9)
	otherSet := New(comparator.IntComparator, 1, 2, 5, 6)

	result := set.Union(otherSet)
	expectedValues := []int{3, 5, 9}

	assert.Equal(3, result.Size())

	for _, element := range expectedValues {
		assert.True(result.Contains(element))
	}
}
