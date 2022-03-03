package zx

import "math/rand"

func MapKey[K comparable, V any](m map[K]V, f func(K, V) K) map[K]V {
	r := make(map[K]V)
	for k, v := range m {
		r[f(k, v)] = v
	}
	return r
}

func MapValue[K comparable, V any](m map[K]V, f func(K, V) V) map[K]V {
	r := make(map[K]V)
	for k, v := range m {
		r[k] = f(k, v)
	}
	return r
}

func Reduce[K comparable, V any, R any](m map[K]V, f func(K, V, R) R, initial R) R {
	r := initial
	for k, v := range m {
		r = f(k, v, r)
	}
	return r
}

func Any[K comparable, V any](m map[K]V, f func(K, V) bool) bool {
	for k, v := range m {
		if f(k, v) {
			return true
		}
	}
	return false
}

func All[K comparable, V any](m map[K]V, f func(K, V) bool) bool {
	for k, v := range m {
		if !f(k, v) {
			return false
		}
	}
	return true
}

func Count[K comparable, V any](m map[K]V, f func(K, V) bool) int {
	r := 0
	for k, v := range m {
		if f(k, v) {
			r++
		}
	}
	return r
}

func Filter[K comparable, V any](m map[K]V, f func(K, V) bool) map[K]V {
	r := make(map[K]V)
	for k, v := range m {
		if f(k, v) {
			r[k] = v
		}
	}
	return r
}

func Foreach[K comparable, V any](m map[K]V, f func(K, V) (K, V)) map[K]V {
	r := make(map[K]V)
	for k, v := range m {
		k, v = f(k, v)
		r[k] = v
	}
	return r
}

func Random[K comparable, V any](m map[K]V) (rk K, rv V) {
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

func FromSlices[K comparable, V any](ks []K, vs []V) map[K]V {
	r := make(map[K]V)
	for i := 0; i < len(ks) && i < len(vs); i++ {
		r[ks[i]] = vs[i]
	}
	return r
}
