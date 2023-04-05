package doublyLinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoublyLinkList_PushBack(t *testing.T) {
	dl := New[int]()
	cases := []int{1, 2, 3, 4, 5, 6, 7, 8}
	want := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for _, v := range cases {
		dl.PushBack(v)
	}

	assert.Equal(t, len(want), dl.Size())

	get := dl.Slice()
	assert.Equal(t, want, get)
}

func TestDoublyLinkList_PushFront(t *testing.T) {
	dl := New[int]()
	cases := []int{1, 2, 3, 4, 5, 6, 7, 8}
	want := []int{8, 7, 6, 5, 4, 3, 2, 1}

	for _, v := range cases {
		dl.PushFront(v)
	}

	assert.Equal(t, len(want), dl.Size())

	get := dl.Slice()
	assert.Equal(t, want, get)
}

func TestDoublyLinkList_PopBack(t *testing.T) {
	dl := New[int]()
	cases := []int{1, 2, 3, 4, 5, 6, 7, 8}
	want := []int{8, 7, 6, 5}

	for _, v := range cases {
		dl.PushFront(v)
	}
	for i := 0; i < 4; i++ {
		dl.PopBack()
	}

	assert.Equal(t, len(want), dl.Size())

	get := dl.Slice()
	assert.Equal(t, want, get)
}

func TestDoublyLinkList_PopFront(t *testing.T) {
	dl := New[int]()
	cases := []int{1, 2, 3, 4, 5, 6, 7, 8}
	want := []int{4, 3, 2, 1}

	for _, v := range cases {
		dl.PushFront(v)
	}
	for i := 0; i < 4; i++ {
		dl.PopFront()
	}

	assert.Equal(t, len(want), dl.Size())

	get := dl.Slice()
	assert.Equal(t, want, get)
}
