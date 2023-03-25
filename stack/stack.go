package stack

import list "github.com/EyTest-bang/goSTL/linkedList"

type Stack[T any] struct {
	list *list.List[T]
}

func New[T any](cmp func(T, T) int) *Stack[T] {
	return &Stack[T]{list: list.New[T](cmp)}
}

func (st *Stack[T]) Top() T {
	return st.list.Right()
}

func (st *Stack[T]) Empty() bool {
	return st.list.Empty()
}

func (st *Stack[T]) Size() int {
	return st.list.Size()
}

func (st *Stack[T]) Push(val T) {
	st.list.InsertRight(val)
}
