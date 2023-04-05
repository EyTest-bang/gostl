package lru

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLRU_InsertOrUpdate(t *testing.T) {
	capacity := 5
	lru := New[int](capacity)
	cases := []int{1, 2, 3, 4, 5, 6, 7}
	want1 := []int{7, 6, 5, 4, 3}
	want2 := []int{6, 7, 5, 4, 3}

	assert.Equal(t, true, lru.Empty())
	for i := 0; i < len(cases); i++ {
		lru.InsertOrUpdate(cases[i])
		if i < capacity {
			assert.Equal(t, i+1, lru.Size())
		} else {
			assert.Equal(t, capacity, lru.Size())
		}
	}
	assert.Equal(t, false, lru.Empty())

	get1 := lru.Slice()
	assert.Equal(t, want1, get1)

	lru.InsertOrUpdate(6)
	assert.Equal(t, capacity, lru.Size())
	get2 := lru.Slice()
	assert.Equal(t, want2, get2)
}
