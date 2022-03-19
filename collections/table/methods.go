package table

import "math/rand"

// Clone returns a new Table with the same contents as this one.
func (m Table[K, V]) Clone() map[K]V {
	n := make(map[K]V)
	for k, v := range m {
		n[k] = v
	}
	return n
}

// MapKey returns a new Table with the result of applying the function to each key.
func (m Table[K, V]) MapKey(f func(K, V) K) map[K]V {
	r := make(map[K]V)
	for k, v := range m {
		r[f(k, v)] = v
	}
	return r
}

// MapValue returns a new Table with the result of applying the function to each value.
func (m Table[K, V]) MapValue(f func(K, V) V) map[K]V {
	r := make(map[K]V)
	for k, v := range m {
		r[k] = f(k, v)
	}
	return r
}

// ReduceKey applies a function against an initial and each key in the Table.
func (m Table[K, V]) ReduceKey(f func(K, V) K, initial K) K {
	r := initial
	for k, v := range m {
		r = f(k, v)
	}
	return r
}

// ReduceValue applies a function against an initial and each value in the Table.
func (m Table[K, V]) ReduceVale(f func(K, V) V, initial V) V {
	r := initial
	for k, v := range m {
		r = f(k, v)
	}
	return r
}

// Any returns true if any element in the Table satisfies the predicate.
func (m Table[K, V]) Any(f func(K, V) bool) bool {
	for k, v := range m {
		if f(k, v) {
			return true
		}
	}
	return false
}

// All returns true if all elements in the Table satisfy the predicate.
func (m Table[K, V]) All(f func(K, V) bool) bool {
	for k, v := range m {
		if !f(k, v) {
			return false
		}
	}
	return true
}

// Count returns the number of elements satisfy the predicate in the Table.
func (m Table[K, V]) Count(f func(K, V) bool) int {
	r := 0
	for k, v := range m {
		if f(k, v) {
			r++
		}
	}
	return r
}

// Filter returns a new Table with all elements satisfy the predicate.
func (m Table[K, V]) Filter(f func(K, V) bool) map[K]V {
	r := make(map[K]V)
	for k, v := range m {
		if f(k, v) {
			r[k] = v
		}
	}
	return r
}

// Map applies the function to each element in the Table.
func (m Table[K, V]) Map(f func(K, V) (K, V)) map[K]V {
	r := make(map[K]V)
	for k, v := range m {
		k, v = f(k, v)
		r[k] = v
	}
	return r
}

// Random returns a random element from the Table.
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
