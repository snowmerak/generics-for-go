package list

func (l *List[T]) Clone() *List[T] {
	nl := New[T]()
	if l == nil {
		return nil
	}

	if l.Count() == 0 {
		return nil
	}

	tail := &node[T]{
		value: l.head.value,
		prev:  nil,
		next:  nil,
	}
	nl.head = tail
	for n := l.head; n != nil; n = n.next {
		tail.next = &node[T]{
			value: n.value,
			prev:  tail,
			next:  nil,
		}
		tail = tail.next
	}
	nl.tail = tail
	return nl
}

func (l *List[T]) Count() int {
	count := 0
	for n := l.head; n != nil; n = n.next {
		count++
	}
	return count
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

func (l *List[T]) Foreach(fun func(T)) {
	for n := l.head; n != nil; n = n.next {
		fun(n.value)
	}
}

func (l *List[T]) Contains(comparer func(T, T) bool, value T) bool {
	for n := l.head; n != nil; n = n.next {
		if comparer(n.value, value) {
			return true
		}
	}
	return false
}

func (l *List[T]) ContainsAll(comparer func(T, T) bool, values ...T) bool {
	for n := l.head; n != nil; n = n.next {
		for _, v := range values {
			if !comparer(n.value, v) {
				return false
			}
		}
	}
	return true
}

func (l *List[T]) Split(fun func(T) bool) (left, right *List[T]) {
	left = New[T]()
	right = New[T]()
	for n := l.head; n != nil; n = n.next {
		if fun(n.value) {
			left.Push(n.value)
		} else {
			right.Push(n.value)
		}
	}
	return
}

func (l *List[T]) Get(index int) *T {
	if index < 0 {
		return nil
	}
	for n := l.head; n != nil; n = n.next {
		if index == 0 {
			return &n.value
		}
		index--
	}
	return nil
}

func (l *List[T]) Set(index int, value T) {
	if index < 0 {
		return
	}
	for n := l.head; n != nil; n = n.next {
		if index == 0 {
			n.value = value
			return
		}
		index--
	}
}
