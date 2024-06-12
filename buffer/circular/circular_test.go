package circular

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCircularBuffer_New(t *testing.T) {
	assert := assert.New(t)
	buffer := New[int](-1)

	assert.Nil(buffer)

	buffer = New[int](10)
	assert.NotNil(buffer)
}

func TestCircularBuffer_Enqueue(t *testing.T) {
	assert := assert.New(t)
	buffer := New[int](3)

	for i := 0; i < 3; i++ {
		assert.Nil(buffer.Enqueue(i))
	}

	assert.NotNil(buffer.Enqueue(4))
}

func TestCircularBuffer_Dequeue(t *testing.T) {
	assert := assert.New(t)
	buffer := New[int](3)

	for i := 0; i < 3; i++ {
		buffer.Enqueue(i)
	}

	for i := 0; i < 3; i++ {
		element, err := buffer.Dequeue()
		assert.Nil(err)
		assert.Equal(i, element)
	}

	_, err := buffer.Dequeue()
	assert.NotNil(err)
}

func TestCircularBuffer_IsEmpty(t *testing.T) {
	assert := assert.New(t)
	buffer := New[int](1)

	assert.True(buffer.IsEmpty())
	buffer.Enqueue(1)
	assert.False(buffer.IsEmpty())
}

func TestCircularBuffer_IsFull(t *testing.T) {
	assert := assert.New(t)
	buffer := New[int](1)
	assert.False(buffer.IsFull())

	buffer.Enqueue(1)
	assert.True(buffer.IsFull())
}
