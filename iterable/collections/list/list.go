package list

type node[T any] struct {
	value T
	prev  *node[T]
	next  *node[T]
}

type List[T any] struct {
	head *node[T]
	tail *node[T]
}

func New[T any]() *List[T] {
	return &List[T]{}
}

func Of[T any](values ...T) *List[T] {
	l := New[T]()
	for _, v := range values {
		l.Push(v)
	}
	return l
}

func (l *List[T]) Push(value T) {
	n := &node[T]{
		value: value,
		next:  nil,
	}
	if l.head == nil {
		l.head = n
		l.tail = n
		return
	}
	n.prev = l.tail
	l.tail.next = n
	l.tail = n
}

func (l *List[T]) Pop() *T {
	if l.tail == nil {
		return nil
	}
	n := l.tail
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return &n.value
	}
	l.tail = n.prev
	l.tail.next = nil
	n.prev = nil
	n.next = nil
	return &n.value
}

func (l *List[T]) Unshift(value T) {
	n := &node[T]{
		value: value,
		next:  nil,
	}
	if l.tail == nil {
		l.head = n
		l.tail = n
		return
	}
	n.next = l.head
	l.head.prev = n
	l.head = n
}

func (l *List[T]) Shift() *T {
	if l.head == nil {
		return nil
	}
	n := l.head
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		return &n.value
	}
	l.head = n.next
	l.head.prev = nil
	n.next = nil
	n.prev = nil
	return &n.value
}

func (l *List[T]) Reverse() {
	if l.head == nil {
		return
	}
	for n := l.head; n != nil; {
		next := n.next
		n.prev, n.next = n.next, n.prev
		n = next
	}
	l.head, l.tail = l.tail, l.head
}

func (l *List[T]) Map(fun func(T) T) {
	for n := l.head; n != nil; n = n.next {
		n.value = fun(n.value)
	}
}

func (l *List[T]) Filter(fun func(T) bool) {
	for n := l.head; n != nil; n = n.next {
		if fun(n.value) {
			continue
		}
		if n.prev != nil {
			n.prev.next = n.next
		}
		if n.next != nil {
			n.next.prev = n.prev
		}
		if n.prev == nil {
			l.head = n.next
		}
		if n.next == nil {
			l.tail = n.prev
		}
	}
}

func (l *List[T]) FoldLeft(fun func(T, T) T, initial T) T {
	var acc T
	if l.head == nil {
		return initial
	}
	acc = initial
	for n := l.head; n != nil; n = n.next {
		acc = fun(acc, n.value)
	}
	return acc
}

func (l *List[T]) FoldRight(fun func(T, T) T, initial T) T {
	var acc T
	if l.tail == nil {
		return initial
	}
	acc = initial
	for n := l.tail; n != nil; n = n.prev {
		acc = fun(n.value, acc)
	}
	return acc
}
