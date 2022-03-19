package iterable

// Generator is a structure that generates elements from a given function.
type Generator[K comparable, V any] struct {
	initKey    K
	initVal    V
	currentKey K
	currentVal V
	next       func(K, V) (K, V)
}

// Next returns the next elements in the generator.
func (g *Generator[K, V]) Next() (key K, value V, exist bool) {
	if g.next != nil {
		key, value = g.next(g.currentKey, g.currentVal)
		g.currentKey, g.currentVal = key, value
		exist = true
		return
	}
	exist = false
	return
}

// HasNext returns true if the generator has next element.
func (g *Generator[K, V]) HasNext() bool {
	return true
}

// Reset resets the generator.
func (g *Generator[K, V]) Reset() {
	g.currentKey = g.initKey
	g.currentVal = g.initVal
}

// NewGenerator returns a new generator with given key and value.
func NewGenerator[K comparable, V any](key K, value V, inner func(K, V) (K, V)) Iterable[K, V] {
	return &Generator[K, V]{key, value, key, value, inner}
}
