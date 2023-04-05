package lru

import (
	"github.com/EyTest-bang/gostl/basic/doublyLinkedList"
)

type LRU[T comparable] struct {
	cache    map[T]*doublyLinkedList.Node[T]
	list     *doublyLinkedList.DoublyLinkList[T]
	capacity int
}

func New[T comparable](capacity int) *LRU[T] {
	return &LRU[T]{
		cache:    make(map[T]*doublyLinkedList.Node[T]),
		list:     doublyLinkedList.New[T](),
		capacity: capacity,
	}
}

func (lru *LRU[T]) Size() int {
	return lru.list.Size()
}

func (lru *LRU[T]) Empty() bool {
	return lru.list.Size() == 0
}

func (lru *LRU[T]) Full() bool {
	return lru.list.Size() > lru.capacity
}

func (lru *LRU[T]) Exist(val T) bool {
	_, exist := lru.cache[val]
	return exist
}

func (lru *LRU[T]) Eliminate() {
	if lru.Empty() {
		return
	}
	val := lru.list.Back()
	lru.list.PopBack()
	delete(lru.cache, val)
}

func (lru *LRU[T]) InsertOrUpdate(val T) {
	if !lru.Exist(val) {
		node := lru.list.PushFront(val)
		lru.cache[val] = node
		if lru.Full() {
			lru.Eliminate()
		}
	} else {
		node := lru.cache[val]
		lru.list.RemoveNode(node)
		lru.list.PushFront(val)
	}
}

func (lru *LRU[T]) Slice() []T {
	return lru.list.Slice()
}
