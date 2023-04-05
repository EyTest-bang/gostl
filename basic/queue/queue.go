package queue

type node[T any] struct {
	next *node[T]
	val  T
}

type Queue[T any] struct {
	head, tail *node[T]
	size       int
}

func createNode[T any](val T) *node[T] {
	return &node[T]{next: nil, val: val}
}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (s *Queue[T]) Size() int {
	return s.size
}

func (s *Queue[T]) Empty() bool {
	return s.size == 0
}

func (s *Queue[T]) Push(val T) {
	n := createNode(val)
	if s.Empty() {
		s.head, s.tail = n, n
	} else {
		s.tail.next = n
		s.tail = n
	}
	s.size++
}

// Pop pops the top value in the stack,
// It returns false only when queue is empty.
func (s *Queue[T]) Pop() bool {
	if s.Empty() {
		return false
	}
	s.head = s.head.next
	if s.tail == nil {
		s.head = nil
	}
	s.size--
	return true
}

// Front returns the front value,
// if queue is empty, exist is false
func (s *Queue[T]) Front() (value T, exist bool) {
	if s.Empty() {
		return
	}
	value = s.head.val
	return value, true
}

// Back returns the back value,
// if queue is empty, exist is false
func (s *Queue[T]) Back() (value T, exist bool) {
	if s.Empty() {
		return
	}
	value = s.tail.val
	return value, true
}
