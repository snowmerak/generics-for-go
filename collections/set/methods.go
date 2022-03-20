package set

// Clone returns a new Set with the same contents as this one.
func (s Set[T]) Clone() Set[T] {
	ns := New[T]()
	for k := range s {
		ns[k] = struct{}{}
	}
	return ns
}

// Add adds the given value to the Set.
func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

// Remove removes the given value from the Set.
func (s Set[T]) Remove(v T) {
	delete(s, v)
}

// Contains returns true if the given value is in the Set.
func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

// Union returns a new Set with the union of this Set and the given Set.
func (s1 Set[T]) Union(s2 Set[T]) Set[T] {
	union := Set[T]{}
	for k := range s1 {
		union[k] = struct{}{}
	}
	for k := range s2 {
		union[k] = struct{}{}
	}
	return union
}

// Intersect returns a new Set with the intersection of this Set and the given Set.
func (s1 Set[T]) Intersect(s2 Set[T]) Set[T] {
	intersection := Set[T]{}
	for k := range s1 {
		if _, ok := s2[k]; ok {
			intersection[k] = struct{}{}
		}
	}
	return intersection
}

// Subtract returns a new Set with the elements of this Set that are not in the given Set.
func (s1 Set[T]) Subtract(s2 Set[T]) Set[T] {
	subtract := Set[T]{}
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			subtract[k] = struct{}{}
		}
	}
	return subtract
}

// Equal returns true if this Set and the given Set are equal.
func (s1 Set[T]) Equal(s2 Set[T]) bool {
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

// IsSubset returns true if this Set is a subset of the given Set.
func (s1 Set[T]) IsSubset(s2 Set[T]) bool {
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			return false
		}
	}
	return true
}

// IsSuperset returns true if this Set is a superset of the given Set.
func (s1 Set[T]) IsSuperset(s2 Set[T]) bool {
	for k := range s2 {
		if _, ok := s1[k]; !ok {
			return false
		}
	}
	return true
}

// IsDisjoint returns true if this Set and the given Set have no elements in common.
func (s1 Set[T]) IsDisjoint(s2 Set[T]) bool {
	for k := range s1 {
		if _, ok := s2[k]; ok {
			return false
		}
	}
	return true
}

//Foreach calls the given function for each element in the Set.
func (s1 Set[T]) Foreach(f func(T)) {
	for k := range s1 {
		f(k)
	}
}
