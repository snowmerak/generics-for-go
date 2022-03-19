// tripple is a tripple of values with other types.
package tripple

// Tripple is a container of 3 values of the other types.
type Tripple[A, B, C any] struct {
	a A
	b B
	c C
}

// Of returns a new Tripple with given values.
func Of[A, B, C any](a A, b B, c C) Tripple[A, B, C] {
	return Tripple[A, B, C]{a, b, c}
}

// Get returns the values of the Tripple.
func (t Tripple[A, B, C]) Get() (A, B, C) {
	return t.a, t.b, t.c
}

// GetA returns the value of the first element.
func (t Tripple[A, B, C]) GetA() A {
	return t.a
}

// GetB returns the value of the second element.
func (t Tripple[A, B, C]) GetB() B {
	return t.b
}

// GetC returns the value of the third element.
func (t Tripple[A, B, C]) GetC() C {
	return t.c
}
