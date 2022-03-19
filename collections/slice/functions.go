package slice

// GroupBy returns a map of the elements in the slice grouped by the result of the function.
func GroupBy[T any, K comparable](slice Slice[T], f func(T) K) map[K]Slice[T] {
	result := make(map[K]Slice[T])
	for _, v := range slice {
		key := f(v)
		result[key] = append(result[key], v)
	}
	return result
}

// Zip returns a Slice of elements with same index.
func Zip[T any](slices ...[]T) []Slice[T] {
	if len(slices) == 0 {
		return nil
	}
	min := len(slices[0])
	for _, v := range slices[1:] {
		if len(v) < min {
			min = len(v)
		}
	}
	result := make([]Slice[T], min)
	for i := 0; i < min; i++ {
		result[i] = make([]T, len(slices))
		for j, v := range slices {
			result[i][j] = v[i]
		}
	}
	return result
}
