package integer

import (
	"github.com/snowmerak/generics-for-go/v2/interfaces/comparable"
	"github.com/snowmerak/generics-for-go/v2/interfaces/copyable"
	"golang.org/x/exp/constraints"
)

// Integer is a type that represents an integer.
type Integer[T constraints.Integer] struct {
	Value T
}

// New creates a new Integer.
func New[T constraints.Integer](value T) Integer[T] {
	return Integer[T]{
		Value: value,
	}
}

var _ comparable.Comparable[Integer[int]] = Integer[int]{}

// CompareTo returns an comparable.Compared indicating the relative values of this and other.
func (i Integer[T]) CompareTo(other Integer[T]) comparable.Compared {
	if i.Value < other.Value {
		return comparable.Less
	} else if i.Value > other.Value {
		return comparable.Greater
	} else {
		return comparable.Equal
	}
}

var _ copyable.Clonable[Integer[int]] = Integer[int]{}

// Clone returns a copy of this.
func (i Integer[T]) Clone() Integer[T] {
	return Integer[T]{
		Value: i.Value,
	}
}
