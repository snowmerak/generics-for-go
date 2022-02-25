package iterable

type Table[K comparable, V any] struct {
	data  map[K]V
	keys  []K
	index int
}

func FromMap[K comparable, V any](inner map[K]V) Iterable[K, V] {
	keys := make([]K, len(inner))
	i := 0
	for k := range inner {
		keys[i] = k
		i++
	}
	return &Table[K, V]{inner, keys, 0}
}

func FromMapWithKeyList[K comparable, V any](inner map[K]V, keys []K) Iterable[K, V] {
	return &Table[K, V]{inner, keys, 0}
}

func (m *Table[K, V]) Next() (key K, value V, exist bool) {
	if m.index < len(m.keys) {
		key = m.keys[m.index]
		value = m.data[key]
		exist = true
		m.index++
		return
	}
	exist = false
	return
}

func (m *Table[K, V]) HasNext() bool {
	return m.index < len(m.keys)
}

func (m *Table[K, V]) Reset() {
	m.index = 0
}
