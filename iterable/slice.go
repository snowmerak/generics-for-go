package iterable

type Slice[V any] struct {
	slice []V
	index int
}

func FromSlice[V any](inner []V) Iterable[int, V] {
	return &Slice[V]{inner, 0}
}

func (s *Slice[V]) Next() (key int, value V, exist bool) {
	if s.index < len(s.slice) {
		key = s.index
		value = s.slice[s.index]
		exist = true
		s.index++
		return
	}
	exist = false
	return
}

func (s *Slice[V]) HasNext() bool {
	return s.index < len(s.slice)
}

func (s *Slice[V]) Reset() {
	s.index = 0
}
