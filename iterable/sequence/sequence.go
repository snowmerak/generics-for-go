package sequence

import "github.com/snowmerak/generics-for-go/iterable"

type Sequence[K comparable, V any] struct {
	iter         iterable.Iterable[K, V]
	firstIndex   int
	currentIndex int

	data map[K]V
	keys []K

	filters []func(K, V) bool
	maps    []func(K, V) (K, V)
	orders  []uint8
}

const (
	filterOrder = 0
	mapOrder    = 1
)

func New[K comparable, V any](iter iterable.Iterable[K, V]) *Sequence[K, V] {
	return &Sequence[K, V]{iter, 0, 0, map[K]V{}, []K{}, []func(K, V) bool{}, []func(K, V) (K, V){}, []uint8{}}
}

func (s *Sequence[K, V]) Take(number int) *Sequence[K, V] {
loop:
	for i := 0; i < number && s.iter.HasNext(); {
		k, v, _ := s.iter.Next()
		filterIndex, mapIndex := 0, 0
		for _, order := range s.orders {
			switch order {
			case filterOrder:
				if s.filters[filterIndex](k, v) {
					continue loop
				}
				filterIndex++
			case mapOrder:
				k, v = s.maps[mapIndex](k, v)
				mapIndex++
			}
		}
		s.data[k] = v
		s.keys = append(s.keys, k)
		i++
	}
	return s
}

func (s *Sequence[K, V]) Drop(number int) *Sequence[K, V] {
	currentLength := len(s.keys)
	if currentLength > 0 {
		l := []K(nil)
		if number >= currentLength {
			number -= currentLength
			l = s.keys
			s.keys = []K{}
		} else {
			l = s.keys[:number]
			s.keys = s.keys[number:]
		}
		for _, k := range l {
			delete(s.data, k)
		}
	}
	for i := 0; i < number; i++ {
		k, v, b := s.iter.Next()
		if !b {
			break
		}
		s.keys = append(s.keys, k)
		s.data[k] = v
	}
	return s
}

func (s *Sequence[K, V]) Map(fn func(K, V) (K, V)) *Sequence[K, V] {
	s.maps = append(s.maps, fn)
	s.orders = append(s.orders, mapOrder)
	keys := make([]K, 0, len(s.keys))
	for i, k := range s.keys {
		nk, nv := fn(k, s.data[k])
		s.data[nk] = nv
		delete(s.data, k)
		keys[i] = nk
	}
	s.keys = keys
	return s
}

func (s *Sequence[K, V]) Filter(fn func(K, V) bool) *Sequence[K, V] {
	s.filters = append(s.filters, fn)
	s.orders = append(s.orders, filterOrder)
	keys := make([]K, len(s.keys))
	i := 0
	for _, k := range s.keys {
		if !fn(k, s.data[k]) {
			keys[i] = k
			i++
		} else {
			delete(s.data, k)
		}
	}
	s.keys = keys[:i]
	return s
}

func (s *Sequence[K, V]) FirstKey() (key K, exist bool) {
	if len(s.keys) < 1 {
		exist = false
		return
	}
	key = s.keys[0]
	exist = true
	return
}

func (s *Sequence[K, V]) FirstValue() (value V, exist bool) {
	if len(s.keys) < 1 {
		exist = false
		return
	}
	value = s.data[s.keys[0]]
	exist = true
	return
}

func (s *Sequence[K, V]) ToKeySlice() []K {
	keys := make([]K, len(s.keys))
	copy(keys, s.keys)
	return keys
}

func (s *Sequence[K, V]) ToValueSlice() []V {
	values := make([]V, len(s.keys))
	for i, k := range s.keys {
		values[i] = s.data[k]
	}
	return values
}

func (s *Sequence[K, V]) ToMap() map[K]V {
	m := make(map[K]V, len(s.keys)*4/3+1)
	for _, k := range s.keys {
		m[k] = s.data[k]
	}
	return m
}
