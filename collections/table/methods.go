package table

import "math/rand"

func (m Table[K, V]) Clone() map[K]V {
	n := make(map[K]V)
	for k, v := range m {
		n[k] = v
	}
	return n
}

func (m Table[K, V]) MapKey(f func(K, V) K) map[K]V {
	r := make(map[K]V)
	for k, v := range m {
		r[f(k, v)] = v
	}
	return r
}

func (m Table[K, V]) MapValue(f func(K, V) V) map[K]V {
	r := make(map[K]V)
	for k, v := range m {
		r[k] = f(k, v)
	}
	return r
}

func (m Table[K, V]) ReduceKey(f func(K, V) K, initial K) K {
	r := initial
	for k, v := range m {
		r = f(k, v)
	}
	return r
}

func (m Table[K, V]) ReduceVale(f func(K, V) V, initial V) V {
	r := initial
	for k, v := range m {
		r = f(k, v)
	}
	return r
}

func (m Table[K, V]) Any(f func(K, V) bool) bool {
	for k, v := range m {
		if f(k, v) {
			return true
		}
	}
	return false
}

func (m Table[K, V]) All(f func(K, V) bool) bool {
	for k, v := range m {
		if !f(k, v) {
			return false
		}
	}
	return true
}

func (m Table[K, V]) Count(f func(K, V) bool) int {
	r := 0
	for k, v := range m {
		if f(k, v) {
			r++
		}
	}
	return r
}

func (m Table[K, V]) Filter(f func(K, V) bool) map[K]V {
	r := make(map[K]V)
	for k, v := range m {
		if f(k, v) {
			r[k] = v
		}
	}
	return r
}

func (m Table[K, V]) Foreach(f func(K, V) (K, V)) map[K]V {
	r := make(map[K]V)
	for k, v := range m {
		k, v = f(k, v)
		r[k] = v
	}
	return r
}

func (m Table[K, V]) Random() (rk K, rv V) {
	l := rand.Intn(len(m))
	i := 0
	for k, v := range m {
		if i == l {
			return k, v
		}
		i++
	}
	return rk, rv
}
