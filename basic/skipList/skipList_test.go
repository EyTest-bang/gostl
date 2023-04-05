package skipList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	asc := func(a, b int) int {
		return a - b
	}
	sl := New[int](asc)
	assert.Equal(t, true, sl.Empty())
	assert.Equal(t, uint64(0), sl.Size())
}

func TestSkipList_Insert(t *testing.T) {
	cases := []int{1, 3, 2, 8, 4, 9}
	want := []int{1, 2, 3, 4, 8, 9}
	asc := func(a, b int) int {
		return a - b
	}
	sl := New[int](asc)
	for i, v := range cases {
		sl.Insert(v)
		assert.Equal(t, uint64(i+1), sl.Size())
	}
	get := sl.Slice()
	assert.Equal(t, want, get)
}

func TestSkipList_Delete(t *testing.T) {
	cases := []int{1, 3, 2, 8, 4, 9, 2, 2, 5, 10}
	deletes := []int{11, 10, 1, 7, 6, 2, 2, 9}
	want := []int{2, 3, 4, 5, 8}
	asc := func(a, b int) int {
		return a - b
	}
	sl := New[int](asc)
	for _, v := range cases {
		sl.Insert(v)
	}
	for _, v := range deletes {
		sl.Delete(v)
	}

	get := sl.Slice()
	assert.Equal(t, uint64(len(want)), sl.Size())
	assert.Equal(t, want, get)
}
