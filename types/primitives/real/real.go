package real

import (
	"github.com/snowmerak/generics-for-go/v2/interfaces/comparable"
	"github.com/snowmerak/generics-for-go/v2/interfaces/copyable"
	"golang.org/x/exp/constraints"
)

// Real is a type that represents a floating point number.
type Real[T constraints.Float] struct {
	Value T
}

// New creates a new Real.
func New[T constraints.Float](value T) Real[T] {
	return Real[T]{
		Value: value,
	}
}

var _ comparable.Comparable[Real[float64]] = Real[float64]{}

// CompareTo returns an comparable.Compared indicating the relative values of this and other.
func (r Real[T]) CompareTo(other Real[T]) comparable.Compared {
	if r.Value < other.Value {
		return comparable.Less
	} else if r.Value > other.Value {
		return comparable.Greater
	} else {
		return comparable.Equal
	}
}

var _ copyable.Clonable[Real[float64]] = Real[float64]{}

// Clone returns a copy of this.
func (r Real[T]) Clone() Real[T] {
	return Real[T]{
		Value: r.Value,
	}
}
