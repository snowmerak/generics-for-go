package pair

type Pair[T, R any] struct {
	a T
	b R
}

func Of[T, R any](a T, b R) Pair[T, R] {
	return Pair[T, R]{a, b}
}

func (p Pair[T, R]) GetA() T {
	return p.a
}

func (p Pair[T, R]) GetB() R {
	return p.b
}
