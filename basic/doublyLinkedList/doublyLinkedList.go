package doublyLinkedList

type Direction int

const (
	LEFT  Direction = -1
	RIGHT Direction = 1
)

type Node[T any] struct {
	prev, next *Node[T]
	val        T
}

type DoublyLinkedList[T any] struct {
	head, tail *Node[T]
	size       int
}

type Iterator[T any] struct {
	node *Node[T]
}

func CreateNode[T any](val T) *Node[T] {
	return &Node[T]{val: val}
}

func New[T any]() *DoublyLinkedList[T] {
	var dft T
	h, t := CreateNode(dft), CreateNode(dft)
	h.next, t.prev = t, h
	return &DoublyLinkedList[T]{head: h, tail: t}
}

func (dl *DoublyLinkedList[T]) Size() int {
	return dl.size
}

func (dl *DoublyLinkedList[T]) Empty() bool {
	return dl.size == 0
}

func (dl *DoublyLinkedList[T]) push(val T, dir Direction) *Node[T] {
	n := CreateNode[T](val)
	switch dir {
	case LEFT:
		dl.head.next.prev, n.next = n, dl.head.next
		dl.head.next, n.prev = n, dl.head
	case RIGHT:
		dl.tail.prev.next, n.prev = n, dl.tail.prev
		dl.tail.prev, n.next = n, dl.tail
	default:
		panic("invalid direction")
	}
	dl.size++
	return n
}

// PushBack pushes the value in the backend of the list
func (dl *DoublyLinkedList[T]) PushBack(val T) *Node[T] {
	return dl.push(val, RIGHT)
}

// PushFront pushes the value in the front of the list
func (dl *DoublyLinkedList[T]) PushFront(val T) *Node[T] {
	return dl.push(val, LEFT)
}

func (dl *DoublyLinkedList[T]) pop(dir Direction) bool {
	if dl.Empty() {
		return false
	}
	switch dir {
	case LEFT:
		dl.head.next, dl.head.next.next.prev = dl.head.next.next, dl.head
	case RIGHT:
		dl.tail.prev, dl.tail.prev.prev.next = dl.tail.prev.prev, dl.tail
	default:
		panic("invalid direction")
	}
	dl.size--
	return true
}

// PopBack pops the value in the backend of the list
func (dl *DoublyLinkedList[T]) PopBack() bool {
	return dl.pop(RIGHT)
}

// PopFront pops the value in the front of the list
func (dl *DoublyLinkedList[T]) PopFront() bool {
	return dl.pop(LEFT)
}

func (dl *DoublyLinkedList[T]) Back() (value T) {
	if dl.Empty() {
		return
	}
	return dl.tail.prev.val
}

func (dl *DoublyLinkedList[T]) Front() (value T) {
	if dl.Empty() {
		return
	}
	return dl.head.next.val
}

// RemoveNode can remove the node in O(1)
// But make sure that the node belongs to doubly linked list.
func (dl *DoublyLinkedList[T]) RemoveNode(node *Node[T]) bool {
	if dl.Empty() {
		return false
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	dl.size--
	return true
}

func (dl *DoublyLinkedList[T]) Slice() []T {
	array := make([]T, dl.size)
	for i, iter := 0, dl.Begin(); iter != dl.End(); i, iter = i+1, iter.Next() {
		array[i] = iter.node.val
	}
	return array
}

// Begin returns the iterator includes the first node
func (dl *DoublyLinkedList[T]) Begin() Iterator[T] {
	return Iterator[T]{node: dl.head.next}
}

// End returns the iterator includes the tail node
func (dl *DoublyLinkedList[T]) End() Iterator[T] {
	return Iterator[T]{node: dl.tail}
}

func (iter *Iterator[T]) Next() Iterator[T] {
	return Iterator[T]{node: iter.node.next}
}

func (iter *Iterator[T]) Value() T {
	return iter.node.val
}
