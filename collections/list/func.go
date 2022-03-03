package list

func Map[T, R any](l *List[T], fun func(T) R) *List[R] {
	nl := New[R]()
	for c := l.head; c != nil; c = c.next {
		nl.Push(fun(c.value))
	}
	return nl
}
