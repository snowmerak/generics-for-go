// slice is a package of Slice.
package slice

// Slice is a []T.
type Slice[T any] []T

// Map returns a Slice with the result of mapping the elements.
func Map[T, R any](slice Slice[T], f func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}
