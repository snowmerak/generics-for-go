package set

func (s Set[T]) Clone() Set[T] {
	ns := New[T]()
	for k := range s {
		ns[k] = struct{}{}
	}
	return ns
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Remove(v T) {
	delete(s, v)
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

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

func (s1 Set[T]) Intersect(s2 Set[T]) Set[T] {
	intersection := Set[T]{}
	for k := range s1 {
		if _, ok := s2[k]; ok {
			intersection[k] = struct{}{}
		}
	}
	return intersection
}

func (s1 Set[T]) Subtract(s2 Set[T]) Set[T] {
	subtract := Set[T]{}
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			subtract[k] = struct{}{}
		}
	}
	return subtract
}

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

func (s1 Set[T]) IsSubset(s2 Set[T]) bool {
	for k := range s1 {
		if _, ok := s2[k]; !ok {
			return false
		}
	}
	return true
}

func (s1 Set[T]) IsSuperset(s2 Set[T]) bool {
	for k := range s2 {
		if _, ok := s1[k]; !ok {
			return false
		}
	}
	return true
}

func (s1 Set[T]) IsDisjoint(s2 Set[T]) bool {
	for k := range s1 {
		if _, ok := s2[k]; ok {
			return false
		}
	}
	return true
}
