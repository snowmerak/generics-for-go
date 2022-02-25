package iterable

type Iterable[K comparable, V any] interface {
	Next() (K, V, bool)
	HasNext() bool
	Reset()
}

func ToSlice[V any](i Iterable[int, V]) []V {
	var slice []V
	for i.HasNext() {
		_, v, _ := i.Next()
		slice = append(slice, v)
	}
	return slice
}

func ToMap[K comparable, V any](i Iterable[K, V]) map[K]V {
	var m map[K]V
	for i.HasNext() {
		k, v, _ := i.Next()
		m[k] = v
	}
	return m
}

func Map[K, NK comparable, V, NV any](i Iterable[K, V], fn func(K, V) (NK, NV)) Iterable[NK, NV] {
	length := 0
	for i.HasNext() {
		i.Next()
		length++
	}
	i.Reset()
	m := make(map[NK]NV, length*4/3+1)
	for i.HasNext() {
		k, v, _ := i.Next()
		nk, nv := fn(k, v)
		m[nk] = nv
	}
	return FromMap(m)
}
