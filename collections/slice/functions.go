package slice

func GroupBy[T any, K comparable](slice Slice[T], f func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, v := range slice {
		key := f(v)
		result[key] = append(result[key], v)
	}
	return result
}

func Zip[T any](slices ...[]T) [][]T {
	if len(slices) == 0 {
		return nil
	}
	min := len(slices[0])
	for _, v := range slices[1:] {
		if len(v) < min {
			min = len(v)
		}
	}
	result := make([][]T, min)
	for i := 0; i < min; i++ {
		result[i] = make([]T, len(slices))
		for j, v := range slices {
			result[i][j] = v[i]
		}
	}
	return result
}
