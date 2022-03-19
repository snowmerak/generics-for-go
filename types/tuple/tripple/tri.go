package tripple

type Tripple[A, B, C any] struct {
	a A
	b B
	c C
}

func Of[A, B, C any](a A, b B, c C) Tripple[A, B, C] {
	return Tripple[A, B, C]{a, b, c}
}

func (t Tripple[A, B, C]) Get() (A, B, C) {
	return t.a, t.b, t.c
}

func (t Tripple[A, B, C]) GetA() A {
	return t.a
}

func (t Tripple[A, B, C]) GetB() B {
	return t.b
}

func (t Tripple[A, B, C]) GetC() C {
	return t.c
}
