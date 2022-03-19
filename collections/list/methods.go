package list

// Clone returns a new List with the same contents as this one.
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

// Count returns the number of elements in the list.
func (l *List[T]) Count() int {
	count := 0
	for n := l.head; n != nil; n = n.next {
		count++
	}
	return count
}

// Push adds an element to the end of the list.
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

// Pop removes and returns the last element of the list.
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

// Unshift adds an element to the beginning of the list.
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

// Shift removes and returns the first element of the list.
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

// Reverse reverses the list.
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

// Map returns a new List with the result of mapping the elements.
func (l *List[T]) Map(fun func(T) T) {
	for n := l.head; n != nil; n = n.next {
		n.value = fun(n.value)
	}
}

// Filter returns a new List with the elements that satisfy the predicate.
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

// FoldLeft returns the result of folding the elements from first to last.
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

// FoldRight returns the result of folding the elements from last to first.
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

// Foreach calls the function for each element in the list.
func (l *List[T]) Foreach(fun func(T)) {
	for n := l.head; n != nil; n = n.next {
		fun(n.value)
	}
}

// Contains returns true if the list contains the element.
func (l *List[T]) Contains(comparer func(T, T) bool, value T) bool {
	for n := l.head; n != nil; n = n.next {
		if comparer(n.value, value) {
			return true
		}
	}
	return false
}

// ContainsAll returns true if the list contains all the elements.
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

// Split splits the list into two lists at the index from given function.
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

// Get returns the element at the given index.
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

// Set sets the element at the given index.
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
