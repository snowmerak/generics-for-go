// set is giving Set structure.
package set

// Set is a map[T]struct{} with additional methods.
type Set[T comparable] map[T]struct{}

// New returns a new Set.
func New[T comparable]() Set[T] {
	return make(map[T]struct{})
}

// FromSlice returns a new Set from a slice.
func FromSlice[T comparable](s []T) Set[T] {
	set := Set[T]{}
	for _, v := range s {
		set[v] = struct{}{}
	}
	return set
}

// Of returns a new Set with the given elements.
func Of[T comparable](values ...T) Set[T] {
	return FromSlice(values)
}

// FromMapKey returns a new Set with keys from a map.
func FromMapKey[K comparable, V any](m map[K]V) Set[K] {
	set := Set[K]{}
	for k := range m {
		set[k] = struct{}{}
	}
	return set
}

// FromMapValue returns a new Set with values from a map.
func FromMapValue[K comparable, V comparable](m map[K]V) Set[V] {
	set := Set[V]{}
	for _, v := range m {
		set[v] = struct{}{}
	}
	return set
}

// ToSlice returns a slice with the elements of the Set.
func ToSlice[T comparable](s Set[T]) []T {
	slice := make([]T, 0, len(s))
	for k := range s {
		slice = append(slice, k)
	}
	return slice
}
