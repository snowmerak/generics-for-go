package set

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return make(map[T]struct{})
}

func FromSlice[T comparable](s []T) Set[T] {
	set := Set[T]{}
	for _, v := range s {
		set[v] = struct{}{}
	}
	return set
}

func FromMapKey[K comparable, V any](m map[K]V) Set[K] {
	set := Set[K]{}
	for k := range m {
		set[k] = struct{}{}
	}
	return set
}

func FromMapValue[K comparable, V comparable](m map[K]V) Set[V] {
	set := Set[V]{}
	for _, v := range m {
		set[v] = struct{}{}
	}
	return set
}

func ToSlice[T comparable](s Set[T]) []T {
	slice := make([]T, 0, len(s))
	for k := range s {
		slice = append(slice, k)
	}
	return slice
}
