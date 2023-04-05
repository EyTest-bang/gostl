package priorityQueue

type PriorityQueue[T any] struct {
	element []T
	comp    func(a, b T) int
}

func New[T any](ele []T, comp func(a, b T) int) (*PriorityQueue[T], bool) {
	if comp == nil {
		return nil, false
	}
	var (
		element []T
		lc, rc  int
	)
	if ele == nil || len(ele) == 0 {
		element = make([]T, 1)
	} else {
		l := len(ele) + 1 // 1 for element[0]
		element = make([]T, l)
		copy(element[1:], ele)
		l--
		for pos := l / 2; pos >= 1; pos-- {
			lc = pos * 2
			if comp(element[pos], element[lc]) < 0 {
				element[pos], element[lc] = element[lc], element[pos]
			}
			if rc = lc + 1; rc <= l && comp(element[pos], element[rc]) < 0 {
				element[pos], element[rc] = element[rc], element[pos]
			}
		}
	}
	return &PriorityQueue[T]{
		element: element,
		comp:    comp,
	}, true
}

func (pq *PriorityQueue[T]) Size() int {
	return len(pq.element) - 1
}

func (pq *PriorityQueue[T]) Empty() bool {
	return pq.Size() == 0
}

func (pq *PriorityQueue[T]) Push(val T) {
	pos := len(pq.element)
	pq.element = append(pq.element, val)
	for ; pos > 1 && pq.comp(val, pq.element[pos/2]) > 0; pos /= 2 {
		pq.element[pos], pq.element[pos/2] = pq.element[pos/2], pq.element[pos]
	}
}

func (pq *PriorityQueue[T]) Pop() bool {
	if pq.Empty() {
		return false
	}
	size := pq.Size()
	pq.element[1] = pq.element[size]
	pq.element = pq.element[:size]
	size--
	for parent := 1; 2*parent <= size; {
		lc := 2 * parent
		if lc+1 <= size && pq.comp(pq.element[lc], pq.element[lc+1]) < 0 {
			lc = lc + 1
		}
		if pq.comp(pq.element[parent], pq.element[lc]) < 0 {
			pq.element[parent], pq.element[lc] = pq.element[lc], pq.element[parent]
			parent = lc
		} else {
			break
		}
	}
	return true
}

func (pq *PriorityQueue[T]) Top() (value T, exist bool) {
	if pq.Empty() {
		return
	}
	value = pq.element[1]
	return value, true
}
