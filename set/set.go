package set

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return make(map[T]struct{})
}

func Add[T comparable](s Set[T], v T) {
	s[v] = struct{}{}
}

func Remove[T comparable](s Set[T], v T) {
	delete(s, v)
}

func Contains[T comparable](s Set[T], v T) bool {
	_, ok := s[v]
	return ok
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
	slice := []T{}
	for k := range s {
		slice = append(slice, k)
	}
	return slice
}

func Union[T comparable](s1, s2 Set[T]) Set[T] {
	union := Set[T]{}
	for k := range s1 {
		union[k] = struct{}{}
	}
	for k := range s2 {
		union[k] = struct{}{}
	}
	return union
}

func Intersect[T comparable](s1, s2 Set[T]) Set[T] {
	intersection := Set[T]{}
	for k := range s1 {
		if _, ok := s2[k]; ok {
			intersection[k] = struct{}{}
		}
	}
	return intersection
}

func Subtract[T comparable](s1, s2 Set[T]) Set[T] {
	subtract := Set[T]{}
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			subtract[k] = struct{}{}
		}
	}
	return subtract
}

func Equal[T comparable](s1, s2 Set[T]) bool {
	if len(s1) != len(s2) {
		return false
	}
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			return false
		}
	}
	return true
}

func IsSubset[T comparable](s1, s2 Set[T]) bool {
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			return false
		}
	}
	return true
}

func IsSuperset[T comparable](s1, s2 Set[T]) bool {
	for k := range s2 {
		if _, ok := s1[k]; !ok {
			return false
		}
	}
	return true
}

func IsDisjoint[T comparable](s1, s2 Set[T]) bool {
	for k := range s1 {
		if _, ok := s2[k]; ok {
			return false
		}
	}
	return true
}
