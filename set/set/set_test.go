package set

import (
	"testing"

	comparator "github.com/dterbah/gods/utils"
	"github.com/stretchr/testify/assert"
)

func TestSet_Add(t *testing.T) {
	assert := assert.New(t)

	set := New[int](comparator.IntComparator)

	set.Add(1, 2, 3)

	assert.Equal(3, set.Size())

	set.Add(1)

	assert.Equal(3, set.Size())
}

func TestSet_At(t *testing.T) {
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

func TestSet_Clear(t *testing.T) {
	assert := assert.New(t)
	set := New[int](comparator.IntComparator, 1, 3, 4)
	set.Clear()

	assert.Equal(0, set.Size())
}

func TestSet_Contains(t *testing.T) {
	assert := assert.New(t)
	set := New[int](comparator.IntComparator, 1, 3, 4)

	assert.True(set.Contains(1))
	assert.True(set.Contains(3))
	assert.True(set.Contains(4))
	assert.False(set.Contains(2))
}

func TestSet_Diff(t *testing.T) {
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

func TestSet_Intersection(t *testing.T) {
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

func TestSet_IsEmpty(t *testing.T) {
	assert := assert.New(t)

	set := New[int](comparator.IntComparator)

	assert.True(set.IsEmpty())

	set.Add(1, 2)

	assert.False(set.IsEmpty())
}

func TestSet_IsSubset(t *testing.T) {
	assert := assert.New(t)
	set := New(comparator.IntComparator, 1, 2, 3, 6, 9)
	otherSet := New(comparator.IntComparator, 1, 2)

	assert.True(set.IsSubset(otherSet))

	set = New(comparator.IntComparator, 1)
	assert.False(set.IsSubset(otherSet))

}

func TestSet_Union(t *testing.T) {
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
