package assert

import "fmt"

// MapLenLess asserts that the length of the map is less than n.
func MapLenLess[K comparable, V any](m map[K]V, n int) error {
	if len(m) >= n {
		return fmt.Errorf("map length is not less than %d", n)
	}
	return nil
}

// SliceLenLessEqual asserts that the length of the slice is less than or equal to n.
func MapLenLessEqual[K comparable, V any](m map[K]V, n int) error {
	if len(m) > n {
		return fmt.Errorf("map length is not less than or equal to %d", n)
	}
	return nil
}

// SliceLenGreater asserts that the length of the slice is greater than n.
func MapLenGreater[K comparable, V any](m map[K]V, n int) error {
	if len(m) <= n {
		return fmt.Errorf("map length is not greater than %d", n)
	}
	return nil
}

// SliceLenGreaterEqual asserts that the length of the slice is greater than or equal to n.
func MapLenGreaterEqual[K comparable, V any](m map[K]V, n int) error {
	if len(m) < n {
		return fmt.Errorf("map length is not greater than or equal to %d", n)
	}
	return nil
}

// MapLenEqual asserts that the length of the map is equal to n.
func MapLenEqual[K comparable, V any](m map[K]V, n int) error {
	if len(m) != n {
		return fmt.Errorf("map length is not equal to %d", n)
	}
	return nil
}

// MapLenNotEqual asserts that the length of the map is not equal to n.
func MapLenNotEqual[K comparable, V any](m map[K]V, n int) error {
	if len(m) == n {
		return fmt.Errorf("map length is not not equal to %d", n)
	}
	return nil
}
