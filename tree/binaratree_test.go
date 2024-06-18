package tree

import (
	"testing"

	comparator "github.com/dterbah/gods/utils"
	"github.com/stretchr/testify/assert"
)

func TestBinaryTree_Add(t *testing.T) {
	assert := assert.New(t)
	tree := New(comparator.IntComparator)

	tree.Add(1)
	tree.Add(2)

	assert.True(tree.Has(1))
	assert.True(tree.Has(2))

	assert.False(tree.Has(4))
}

func TestBinaryTree_Min(t *testing.T) {
	assert := assert.New(t)
	tree := New(comparator.IntComparator)

	_, err := tree.Min()
	assert.NotNil(err)
	tree.Add(1, 2, 3, 4)

	min, err := tree.Min()
	assert.Nil(err)
	assert.Equal(1, min)

	tree.Add(-10)
	min, err = tree.Min()

	assert.Nil(err)
	assert.Equal(-10, min)
}

func TestBinaryTree_Max(t *testing.T) {
	assert := assert.New(t)
	tree := New(comparator.IntComparator)

	_, err := tree.Max()
	assert.NotNil(err)
	tree.Add(-1, 0, 1, 2)
	max, err := tree.Max()

	assert.Nil(err)
	assert.Equal(2, max)
	tree.Add(10)
	max, err = tree.Max()
	assert.Nil(err)
	assert.Equal(10, max)
}
