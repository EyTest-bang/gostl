package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
	q := New[int]()
	for i := 0; i < 10; i++ {
		q.Push(i)
		v, ok := q.Back()
		assert.Equal(t, true, ok)
		assert.Equal(t, i, v)
		v, ok = q.Front()
		assert.Equal(t, true, ok)
		assert.Equal(t, 0, v)
	}
}

func TestStack_Empty(t *testing.T) {
	q := New[int]()
	_, ok := q.Front()
	assert.Equal(t, false, ok)
	_, ok = q.Back()
	assert.Equal(t, false, ok)
	assert.Equal(t, true, q.Empty())
	q.Push(1)
	assert.Equal(t, false, q.Empty())
}

func TestStack_Pop(t *testing.T) {
	q := New[int]()
	assert.Equal(t, 0, q.Size())
	for i := 0; i < 10; i++ {
		q.Push(i)
		assert.Equal(t, i+1, q.Size())
	}
	for i := 1; i <= 10; i++ {
		assert.Equal(t, true, q.Pop())
		assert.Equal(t, 10-i, q.Size())
		v, ok := q.Front()
		if i < 10 {
			assert.Equal(t, true, ok)
			assert.Equal(t, i, v)
		} else {
			assert.Equal(t, false, ok)
		}
	}
}
