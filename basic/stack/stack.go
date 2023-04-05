package stack

type node[T any] struct {
	next *node[T]
	val  T
}

type Stack[T any] struct {
	head, tail *node[T]
	size       int
}

func createNode[T any](val T) *node[T] {
	return &node[T]{next: nil, val: val}
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Size() int {
	return s.size
}

func (s *Stack[T]) Empty() bool {
	return s.size == 0
}

func (s *Stack[T]) Push(val T) {
	n := createNode(val)
	if s.Empty() {
		s.head, s.tail = n, n
	} else {
		n.next = s.tail
		s.tail = n
	}
	s.size++
}

// Pop pops the top value in the stack,
// It returns false only when stack is empty.
func (s *Stack[T]) Pop() bool {
	if s.Empty() {
		return false
	}
	s.tail = s.tail.next
	if s.tail == nil {
		s.head = nil
	}
	s.size--
	return true
}

// Top returns the top value,
// if stack is empty, exist is false
func (s *Stack[T]) Top() (value T, exist bool) {
	if s.Empty() {
		return
	}
	value = s.tail.val
	return value, true
}
