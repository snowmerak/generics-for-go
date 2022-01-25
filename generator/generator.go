package generator

type Generator[T any] struct {
	cur  T
	next func(T) T
}

func With[T any](init T, next func(T) T) Generator[T] {
	return Generator[T]{
		cur:  init,
		next: next,
	}
}

func (g Generator[T]) Next() Generator[T] {
	return Generator[T]{
		cur:  g.next(g.cur),
		next: g.next,
	}
}

func (g Generator[T]) Get() T {
	return g.cur
}

func (g Generator[T]) Take(n int) []T {
	l := make([]T, n)
	ng := g
	for i := 0; i < n; i++ {
		l[i] = ng.cur
		ng = ng.Next()
	}
	return l
}
