// pair is a pair of values of the two types.
package pair

// pair is a container of type T and type R.
type Pair[T, R any] struct {
	a T
	b R
}

// Of returns a new Pair with given values.
func Of[T, R any](a T, b R) Pair[T, R] {
	return Pair[T, R]{a, b}
}

// Get returns the values of the Pair.
func (p Pair[T, R]) Get() (T, R) {
	return p.a, p.b
}

// A returns the first value of the Pair.
func (p Pair[T, R]) A() T {
	return p.a
}

// B returns the second value of the Pair.
func (p Pair[T, R]) B() R {
	return p.b
}
