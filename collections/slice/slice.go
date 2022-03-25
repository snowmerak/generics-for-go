// slice is a package of Slice.
package slice

import (
	"github.com/snowmerak/generics-for-go/v2/interfaces/copyable"
)

var _ copyable.Clonable[Slice[int]] = Slice[int]{}

// Slice is a []T.
type Slice[T any] []T

// New creates a new Slice.
func New[T any]() Slice[T] {
	return make(Slice[T], 0)
}

// Of returns a Slice with the given elements.
func Of[T any](values ...T) Slice[T] {
	return Slice[T](values)
}

// Map returns a Slice with the result of mapping the elements.
func Map[T, R any](slice Slice[T], f func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}
