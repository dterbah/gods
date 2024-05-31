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
