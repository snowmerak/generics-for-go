package slice

import (
	"fmt"
	"math/rand"
	"sort"
)

// Clone returns a copy of the slice.
func (slice Slice[T]) Clone() Slice[T] {
	ns := make(Slice[T], len(slice))
	copy(ns, slice)
	return ns
}

// Reduce applies a function against an initial and each element in the slice.
func (slice Slice[T]) Reduce(f func(T, T) T, initial T) T {
	result := initial
	for _, v := range slice {
		result = f(result, v)
	}
	return result
}

// Any returns true if any element in the slice satisfies the predicate.
func (slice Slice[T]) Any(f func(T) bool) bool {
	for _, v := range slice {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns true if all elements in the slice satisfy the predicate.
func (slice Slice[T]) All(f func(T) bool) bool {
	for _, v := range slice {
		if !f(v) {
			return false
		}
	}
	return true
}

// BinarySearch returns the index of the element in the slice, or -1 if not found.
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

// Count returns the number of elements satisfy the predicate in the slice.
func (slice Slice[T]) Count(f func(T) bool) int {
	count := 0
	for _, v := range slice {
		if f(v) {
			count++
		}
	}
	return count
}

// Filter returns a new slice containing all elements satisfy the predicate in the slice.
func (slice Slice[T]) Filter(f func(T) bool) Slice[T] {
	result := make([]T, 0)
	for _, v := range slice {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

// FilterIndex returns a new slice containing all elements satisfy the predicate with index in the slice.
func (slice Slice[T]) FilterIndex(f func(int) bool) Slice[T] {
	result := make([]T, 0)
	for i, v := range slice {
		if f(i) {
			result = append(result, v)
		}
	}
	return result
}

// FirstOf returns the first element in the slice that satisfies the predicate.
func (slice Slice[T]) FirstOf(f func(T) bool) (r T) {
	for _, v := range slice {
		if f(v) {
			return v
		}
	}
	return r
}

// FirstIndexOf returns the index of the first element in the slice that satisfies the predicate.
func (slice Slice[T]) FirstIndexOf(f func(T) bool) int {
	for i, v := range slice {
		if f(v) {
			return i
		}
	}
	return -1
}

// LastOf returns the last element in the slice that satisfies the predicate.
func (slice Slice[T]) LastOf(f func(T) bool) (r T) {
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return slice[i]
		}
	}
	return r
}

// LastIndexOf returns the index of the last element in the slice that satisfies the predicate.
func (slice Slice[T]) LastIndexOf(f func(T) bool) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if f(slice[i]) {
			return i
		}
	}
	return -1
}

// Map applies a function to each element in the slice.
func (slice Slice[T]) Map(f func(int, T) T) Slice[T] {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[i] = f(i, v)
	}
	return result
}

// Reverse returns a new slice with elements in reverse order.
func (slice Slice[T]) Reverse() Slice[T] {
	result := make([]T, len(slice))
	for i, v := range slice {
		result[len(slice)-i-1] = v
	}
	return result
}

// Max returns the maximum element in the slice.
func (slice Slice[T]) Max(f func(T, T) int) (r T) {
	if len(slice) == 0 {
		return r
	}
	r = slice[0]
	for _, v := range slice[1:] {
		if f(v, r) == -1 {
			r = v
		}
	}
	return r
}

// Min returns the minimum element in the slice.
func (slice Slice[T]) Min(f func(T, T) int) (r T) {
	if len(slice) == 0 {
		return r
	}
	r = slice[0]
	for _, v := range slice[1:] {
		if f(v, r) == 1 {
			r = v
		}
	}
	return r
}

// Random returns a random element in the slice.
func (slice Slice[T]) Random() T {
	return slice[rand.Intn(len(slice))]
}

// Shuffle returns itself with elements in random order.
func (slice Slice[T]) Shuffle() Slice[T] {
	for i := len(slice) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

// Sort returns a new slice with elements in sorted order.
func (slice Slice[T]) Sort(f func(T, T) int) Slice[T] {
	result := make([]T, len(slice))
	copy(result, slice)
	sort.Slice(result, func(i, j int) bool {
		return f(result[i], result[j]) == -1
	})
	return result
}

//Chunk returns a slice of slices in length of size.
func (slice Slice[T]) Chunk(size int) []Slice[T] {
	result := make([]Slice[T], 0)
	for i := 0; i < len(slice); i += size {
		min := i + size
		if min > len(slice) {
			min = len(slice)
		}
		result = append(result, slice[i:min])
	}
	return result
}

// JoinToString returns a string with all elements joined by sep.
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

// Foreach calls the function f for each element in the slice.
func (slice Slice[T]) Foreach(f func(T)) {
	for _, v := range slice {
		f(v)
	}
}
