package slice

type Slice[T any] []T

func Map[T, R any](slice Slice[T], f func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}
