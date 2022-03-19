package slice

import (
	"fmt"
	"math/rand"
	"sort"
)

func (slice Slice[T]) Reduce(f func(T, T) T, initial T) T {
	result := initial
	for _, v := range slice {
		result = f(result, v)
	}
	return result
}

func (slice Slice[T]) Any(f func(T) bool) bool {
	for _, v := range slice {
		if f(v) {
			return true
		}
	}
	return false
}

func (slice Slice[T]) All(f func(T) bool) bool {
	for _, v := range slice {
		if !f(v) {
			return false
		}
	}
	return true
}

func (slice Slice[T]) BinarySearch(value T, f func(T, T) int) int {
	low := 0
	high := len(slice)
	for low < high {
		mid := (low + high) / 2
		switch f(slice[mid], value) {
		case 0:
			return mid
		case 1:
			high = mid
		case -1:
			low = mid + 1
		}
	}
	return -1
}

func (slice Slice[T]) Count(f func(T) bool) int {
	count := 0
	for _, v := range slice {
		if f(v) {
			count++
		}
	}
	return count
}

func (slice Slice[T]) Filter(f func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func (slice Slice[T]) FilterIndex(f func(int) bool) []T {
	result := make([]T, 0)
	for i, v := range slice {
		if f(i) {
			result = append(result, v)
		}
	}
	return result
}

func (slice Slice[T]) FirstOf(f func(T) bool) (r T) {
	for _, v := range slice {
		if f(v) {
			return v
		}
	}
	return r
}

func (slice Slice[T]) FirstIndexOf(f func(T) bool) int {
	for i, v := range slice {
		if f(v) {
			return i
		}
	}
	return -1
}

func (slice Slice[T]) LastOf(f func(T) bool) (r T) {
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return slice[i]
		}
	}
	return r
}

func (slice Slice[T]) LastIndexOf(f func(T) bool) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

func (slice Slice[T]) Foreach(f func(int, T) T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[i] = f(i, v)
	}
	return result
}

func (slice Slice[T]) Reverse() []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[len(slice)-i-1] = v
	}
	return result
}

func (slice Slice[T]) Max(f func(T, T) int) (r T) {
	if len(slice) == 0 {
		return r
	}
	r = slice[0]
	for _, v := range slice {
		if f(v, r) == -1 {
			r = v
		}
	}
	return r
}

func (slice Slice[T]) Min(f func(T, T) int) (r T) {
	if len(slice) == 0 {
		return r
	}
	r = slice[0]
	for _, v := range slice {
		if f(v, r) == 1 {
			r = v
		}
	}
	return r
}

func (slice Slice[T]) Random() T {
	return slice[rand.Intn(len(slice))]
}

func Shuffle[T any](slice Slice[T]) []T {
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func (slice Slice[T]) Sort(f func(T, T) int) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	sort.Slice(result, func(i, j int) bool {
		return f(result[i], result[j]) == -1
	})
	return result
}

func (slice Slice[T]) Chunk(size int) [][]T {
	result := make([][]T, 0)
	for i := 0; i < len(slice); i += size {
		min := i + size
		if min > len(slice) {
			min = len(slice)
		}
		result = append(result, slice[i:min])
	}
	return result
}

func (slice Slice[T]) JoinToString(sep string) string {
	result := ""
	for i, v := range slice {
		if i != 0 {
			result += sep
		}
		result += fmt.Sprintf("%v", v)
	}
	return result
}
