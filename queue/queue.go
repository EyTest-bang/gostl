package queue

import list "github.com/EyTest-bang/goSTL/linkedList"

type Queue[T any] struct {
	list *list.List[T]
}

func New[T any](cmp func(T, T) int) *Queue[T] {
	return &Queue[T]{list: list.New[T](cmp)}
}

func (q *Queue[T]) Front() T {
	return q.list.Left()
}

func (q *Queue[T]) Back() T {
	return q.list.Right()
}

func (q *Queue[T]) Empty() bool {
	return q.list.Empty()
}
func (q *Queue[T]) Size() int {
	return q.list.Size()
}
func (q *Queue[T]) Push(val T) {
	q.list.InsertRight(val)
}
func (q *Queue[T]) Pop() {
	q.list.DeleteLeft()
}
