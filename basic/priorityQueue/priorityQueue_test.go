package priorityQueue

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestPriorityQueueDesc(t *testing.T) {
	cases1 := []int{1, 3, 2, 8, 4, 9}
	cases2 := []int{5, 11, 10, 16, 8}
	want1 := []int{9, 11, 11, 16, 16}
	want2 := append(cases2, cases1...)
	sort.Slice(want2, func(i, j int) bool {
		return want2[i] > want2[j]
	})
	desc := func(a, b int) int {
		return a - b
	}
	var get1 []int
	var get2 []int
	pq, ok := New[int](nil, desc)
	assert.Equal(t, true, ok)
	assert.Equal(t, true, pq.Empty())
	pq, ok = New[int](cases1, desc)
	assert.Equal(t, true, ok)
	for _, v := range cases2 {
		pq.Push(v)
		v1, ok := pq.Top()
		assert.Equal(t, true, ok)
		get1 = append(get1, v1)
	}

	assert.Equal(t, len(cases1)+len(want1), pq.Size())
	assert.Equal(t, want1, get1)

	for !pq.Empty() {
		v2, ok := pq.Top()
		assert.Equal(t, true, ok)
		get2 = append(get2, v2)
		pq.Pop()
	}

	assert.Equal(t, 0, pq.Size())
	assert.Equal(t, want2, get2)
}

func TestPriorityQueueAsc(t *testing.T) {
	cases1 := []int{1, 3, 2, 8, 4, 9}
	cases2 := []int{5, 11, 10, 16, 8}
	want1 := []int{1, 1, 1, 1, 1}
	want2 := append(cases2, cases1...)
	sort.Ints(want2)
	asc := func(a, b int) int {
		return b - a
	}
	var get1 []int
	var get2 []int
	pq, ok := New[int](nil, asc)
	assert.Equal(t, true, ok)
	assert.Equal(t, true, pq.Empty())
	pq, ok = New[int](cases1, asc)
	assert.Equal(t, true, ok)
	for _, v := range cases2 {
		pq.Push(v)
		v1, ok := pq.Top()
		assert.Equal(t, true, ok)
		get1 = append(get1, v1)
	}

	assert.Equal(t, len(cases1)+len(want1), pq.Size())
	assert.Equal(t, want1, get1)

	for !pq.Empty() {
		v2, ok := pq.Top()
		assert.Equal(t, true, ok)
		get2 = append(get2, v2)
		pq.Pop()
	}

	assert.Equal(t, 0, pq.Size())
	assert.Equal(t, want2, get2)
}
