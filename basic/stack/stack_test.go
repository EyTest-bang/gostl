package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
	s := New[int]()
	for i := 0; i < 10; i++ {
		s.Push(i)
		v, ok := s.Top()
		assert.Equal(t, true, ok)
		assert.Equal(t, i, v)
	}
}

func TestStack_Empty(t *testing.T) {
	s := New[int]()
	_, ok := s.Top()
	assert.Equal(t, false, ok)
	assert.Equal(t, true, s.Empty())
	s.Push(1)
	assert.Equal(t, false, s.Empty())
}

func TestStack_Size(t *testing.T) {
	s := New[int]()
	assert.Equal(t, 0, s.Size())
	for i := 0; i < 10; i++ {
		s.Push(i)
		assert.Equal(t, i+1, s.Size())
	}
	for i := 10; i > 0; i-- {
		assert.Equal(t, true, s.Pop())
		assert.Equal(t, i-1, s.Size())
	}
}
