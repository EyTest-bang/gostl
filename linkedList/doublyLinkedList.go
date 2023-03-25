package linkedList

import "fmt"

type direction int

const (
	LEFT  direction = -1
	RIGHT direction = 1
)

type node[T any] struct {
	prev, next *node[T]
	val        T
}

type List[T any] struct {
	head, tail *node[T]
	size       int
	cmp        func(T, T) int
}

type iterator[T any] struct {
	cur *node[T]
	cmp func(T, T) int
}

func New[T any](cmp func(T, T) int) *List[T] {
	return &List[T]{cmp: cmp}
}

func (l *List[T]) insert(val T, dir direction) {
	newNode := &node[T]{val: val}
	if l.size == 0 {
		l.head, l.tail = newNode, newNode
	} else {
		switch dir {
		case LEFT:
			newNode.next, newNode.prev = l.head, l.head.prev
			l.head.prev, l.head = newNode, newNode
		case RIGHT:
			newNode.next, newNode.prev = l.tail.next, l.tail
			l.tail.next, l.tail = newNode, newNode
		default:
			panic(fmt.Sprintf("Invalid direction: %d\n", dir))
		}
	}

	l.size++
}

func (l *List[T]) InsertRight(val T) {
	l.insert(val, RIGHT)
}

func (l *List[T]) InsertLeft(val T) {
	l.insert(val, LEFT)
}

func (l *List[T]) delete(dir direction) {
	if l.size == 0 {
		return
	}
	if l.size == 1 {
		l.head = nil
		l.tail = nil
	} else {
		switch dir {
		case RIGHT:
			l.tail = l.tail.prev
			l.tail.next = nil
		case LEFT:
			l.head = l.head.next
			l.head.prev = nil
		default:
			panic(fmt.Sprintf("Invalid direction: %d\n", dir))
		}
	}

	l.size--
}

func (l *List[T]) DeleteRight() {
	l.delete(RIGHT)
}

func (l *List[T]) DeleteLeft() {
	l.delete(LEFT)
}

func (l *List[T]) get(dir direction) T {
	var value T
	if l.size != 0 {
		switch dir {
		case RIGHT:
			value = l.tail.val
		case LEFT:
			value = l.head.val
		default:
			panic(fmt.Sprintf("Invalid direction: %d\n", dir))
		}
	}
	return value
}
func (l *List[T]) Right() T {
	return l.get(RIGHT)
}

func (l *List[T]) Left() T {
	return l.get(LEFT)
}

func (l *List[T]) Empty() bool {
	return l.size == 0
}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) Begin() *iterator[T] {
	return &iterator[T]{cur: l.head, cmp: l.cmp}
}

func (l *List[T]) End() *iterator[T] {
	return &iterator[T]{cur: nil, cmp: l.cmp}
}

func (iter *iterator[T]) equals(iter1 *iterator[T]) bool {
	return iter == iter1
}

func (iter *iterator[T]) Next() *iterator[T] {
	return &iterator[T]{cur: iter.cur.next, cmp: iter.cmp}
}

func (l *List[T]) Exist(val T) bool {
	for iter := l.Begin(); !iter.equals(l.End()); iter = iter.Next() {
		if iter.cmp != nil && iter.cmp(iter.cur.val, val) == 0 {
			return true
		}
	}
	return false
}
