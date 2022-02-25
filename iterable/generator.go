package iterable

type Generator[K comparable, V any] struct {
	initKey    K
	initVal    V
	currentKey K
	currentVal V
	next       func(K, V) (K, V)
}

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

func (g *Generator[K, V]) HasNext() bool {
	return true
}

func (g *Generator[K, V]) Reset() {
	g.currentKey = g.initKey
	g.currentVal = g.initVal
}

func NewGenerator[K comparable, V any](key K, value V, inner func(K, V) (K, V)) Iterable[K, V] {
	return &Generator[K, V]{key, value, key, value, inner}
}
