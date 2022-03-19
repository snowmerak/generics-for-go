package iterable

// Slice is a iterable structure based on slice.
type Slice[V any] struct {
	slice []V
	index int
}

// FromSlice returns a new Slice with given slice.
func FromSlice[V any](inner []V) Iterable[int, V] {
	return &Slice[V]{inner, 0}
}

// Of returns a new Slice with given values.
func Of[V any](values ...V) Iterable[int, V] {
	return FromSlice(values)
}

// Next returns the next element in the Slice.
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

// HasNext returns true if the Slice has next element.
func (s *Slice[V]) HasNext() bool {
	return s.index < len(s.slice)
}

// Reset resets the index of Slice.
func (s *Slice[V]) Reset() {
	s.index = 0
}
