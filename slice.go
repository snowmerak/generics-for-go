package zx

import (
	"errors"
	"math/rand"
	"sort"
)

func Map[T, R any](slice []T, f func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

func Reduce[T, R any](slice []T, f func(R, T) R, initial R) R {
	result := initial
	for _, v := range slice {
		result = f(result, v)
	}
	return result
}

func Any[T any](slice []T, f func(T) bool) bool {
	for _, v := range slice {
		if f(v) {
			return true
		}
	}
	return false
}

func All[T any](slice []T, f func(T) bool) bool {
	for _, v := range slice {
		if !f(v) {
			return false
		}
	}
	return true
}

func BinarySearch[T any](slice []T, value T, f func(T, T) int) int {
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

func Count[T any](slice []T, f func(T) bool) int {
	count := 0
	for _, v := range slice {
		if f(v) {
			count++
		}
	}
	return count
}

func Filter[T any](slice []T, f func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func FilterIndex[T any](slice []T, f func(int) bool) []T {
	result := make([]T, 0)
	for i, v := range slice {
		if f(i) {
			result = append(result, v)
		}
	}
	return result
}

func FirstOf[T any](slice []T, f func(T) bool) (r T, err error) {
	for _, v := range slice {
		if f(v) {
			return v, nil
		}
	}
	return r, errors.New("First: not found")
}

func FirstIndexOf[T any](slice []T, f func(T) bool) (int, error) {
	for i, v := range slice {
		if f(v) {
			return i, nil
		}
	}
	return -1, errors.New("First: not found")
}

func LastOf[T any](slice []T, f func(T) bool) (r T, err error) {
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return slice[i], nil
		}
	}
	return r, errors.New("Last: not found")
}

func LastIndexOf[T any](slice []T, f func(T) bool) (int, error) {
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i, nil
		}
	}
	return -1, errors.New("Last: not found")
}

func Reverse[T any](slice []T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[len(slice)-i-1] = v
	}
	return result
}

func Foreach[T any](slice []T, f func(int, T) T) []T {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[i] = f(i, v)
	}
	return result
}

func GroupBy[T any, K comparable](slice []T, f func(T) K) map[K][]T {
	result := make(map[K][]T)
	for _, v := range slice {
		key := f(v)
		result[key] = append(result[key], v)
	}
	return result
}

func Max[T any](slice []T, f func(T, T) int) (r T, err error) {
	if len(slice) == 0 {
		return r, errors.New("Max: empty slice")
	}
	r = slice[0]
	for _, v := range slice {
		if f(v, r) == -1 {
			r = v
		}
	}
	return r, err
}

func Min[T any](slice []T, f func(T, T) int) (r T, err error) {
	if len(slice) == 0 {
		return r, errors.New("Min: empty slice")
	}
	r = slice[0]
	for _, v := range slice {
		if f(v, r) == 1 {
			r = v
		}
	}
	return r, err
}

func Random[T any](slice []T) T {
	return slice[rand.Intn(len(slice))]
}

func Shuffle[T any](slice []T) []T {
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func Sort[T any](slice []T, f func(T, T) int) []T {
	result := make([]T, len(slice))
	copy(result, slice)
	sort.Slice(result, func(i, j int) bool {
		return f(result[i], result[j]) == -1
	})
	return result
}

func Chunk[T any](slice []T, size int) [][]T {
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

func Zip[T any](sliceA, sliceB []T) [][]T {
	min := len(sliceA)
	if len(sliceB) < min {
		min = len(sliceB)
	}
	result := make([][]T, min)
	for i := 0; i < min; i++ {
		result[i] = []T{sliceA[i], sliceB[i]}
	}
	return result
}
