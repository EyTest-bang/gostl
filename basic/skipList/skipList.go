package skipList

import "math/rand"

const (
	MAXLEVEL = 32
	MAXP     = 0.25
)

type skipListLevel[T any] struct {
	forward *skipListNode[T]
	span    uint64
}

type skipListNode[T any] struct {
	element  T
	backward *skipListNode[T]
	level    []skipListLevel[T]
}

type SkipList[T any] struct {
	head, tail *skipListNode[T]
	level      int
	length     uint64
	cmp        func(a, b T) int
}

type Iterator[T any] struct {
	node *skipListNode[T]
}

func createSkipListNode[T any](element T, level int) *skipListNode[T] {
	return &skipListNode[T]{element: element, backward: nil, level: make([]skipListLevel[T], level)}
}

func New[T any](cmp func(a, b T) int) *SkipList[T] {
	var element T
	return &SkipList[T]{
		head:  createSkipListNode[T](element, MAXLEVEL),
		level: 1,
		cmp:   cmp,
	}
}

func (sl *SkipList[T]) Size() uint64 {
	return sl.length
}

func (sl *SkipList[T]) Empty() bool {
	return sl.length == 0
}

func randomLevel() int {
	newLevel := 1
	for rand.Float64() < MAXP {
		newLevel++
	}
	return newLevel
}

// Insert inserts a new element into the list.
// If there are some equal elements, make sure that cmp works in right way
func (sl *SkipList[T]) Insert(element T) {
	var (
		update [MAXLEVEL]*skipListNode[T]
		rank   [MAXLEVEL]uint64
		x      *skipListNode[T]
		level  int
	)
	x = sl.head
	for i := sl.level - 1; i >= 0; i-- {
		if i == sl.level-1 {
			rank[i] = 0
		} else {
			rank[i] = rank[i+1]
		}
		for x.level[i].forward != nil && sl.cmp(x.level[i].forward.element, element) < 0 {
			x = x.level[i].forward
			rank[i] += x.level[i].span
		}
		update[i] = x
	}

	level = randomLevel()
	if level > sl.level {
		for i := sl.level; i < level; i++ {
			rank[i] = 0
			update[i] = sl.head
			update[i].level[i].span = sl.length
		}
		sl.level = level
	}

	x = createSkipListNode[T](element, level)

	for i := 0; i < level; i++ {
		x.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = x

		x.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
		update[i].level[i].span = rank[0] - rank[i] + 1
	}

	for i := level; i < sl.level; i++ {
		update[i].level[i].span++
	}

	if update[0] == sl.head {
		x.backward = nil
	} else {
		x.backward = update[0]
	}

	if x.level[0].forward != nil {
		x.level[0].forward.backward = x
	} else {
		sl.tail = x
	}

	sl.length++
}

func (sl *SkipList[T]) deleteNode(x *skipListNode[T], update [MAXLEVEL]*skipListNode[T]) {
	for i := 0; i < sl.level; i++ {
		if update[i].level[i].forward == x {
			update[i].level[i].forward = x.level[i].forward
			update[i].level[i].span += x.level[i].span - 1
		} else {
			update[i].level[i].span--
		}
	}

	if x.level[0].forward != nil {
		x.level[0].forward.backward = x.backward
	} else {
		sl.tail = x.backward
	}

	for sl.level > 1 && sl.head.level[sl.level-1].forward == nil {
		sl.level--
	}

	sl.length--
}

// Delete deletes element from the list.
// It returns true if the element exists, returns false if not.
func (sl *SkipList[T]) Delete(element T) bool {
	var (
		update [MAXLEVEL]*skipListNode[T]
		x      *skipListNode[T]
	)

	x = sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && sl.cmp(x.level[i].forward.element, element) < 0 {
			x = x.level[i].forward
		}
		update[i] = x
	}

	x = x.level[0].forward
	if x != nil && sl.cmp(x.element, element) == 0 {
		sl.deleteNode(x, update)
		return true
	}
	return false
}

func (sl *SkipList[T]) Slice() []T {
	array := make([]T, sl.length)
	for i, iter := 0, sl.Begin(); iter != sl.End(); iter = iter.Next() {
		array[i] = iter.node.element
		i++
	}
	return array
}

func (sl *SkipList[T]) Begin() Iterator[T] {
	return Iterator[T]{node: sl.head.level[0].forward}
}

func (sl *SkipList[T]) End() Iterator[T] {
	return Iterator[T]{node: nil}
}

func (iter *Iterator[T]) Next() Iterator[T] {
	return Iterator[T]{node: iter.node.level[0].forward}
}

func (iter *Iterator[T]) Value() T {
	return iter.node.element
}
