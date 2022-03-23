package assert

import "fmt"

// SliceLenLess asserts that the length of the slice is less than n.
func SliceLenLess[T any](slice []T, n int) error {
	if len(slice) >= n {
		return fmt.Errorf("slice length is not less than %d", n)
	}
	return nil
}

// SliceLenLessEqual asserts that the length of the slice is less than or equal to n.
func SliceLenLessEqual[T any](slice []T, n int) error {
	if len(slice) > n {
		return fmt.Errorf("slice length is not less than or equal to %d", n)
	}
	return nil
}

// SliceLenGreater asserts that the length of the slice is greater than n.
func SliceLenGreater[T any](slice []T, n int) error {
	if len(slice) <= n {
		return fmt.Errorf("slice length is not greater than %d", n)
	}
	return nil
}

// SliceLenGreaterEqual asserts that the length of the slice is greater than or equal to n.
func SliceLenGreaterEqual[T any](slice []T, n int) error {
	if len(slice) < n {
		return fmt.Errorf("slice length is not greater than or equal to %d", n)
	}
	return nil
}

// SliceLenEqual asserts that the length of the slice is equal to n.
func SliceLenEqual[T any](slice []T, n int) error {
	if len(slice) != n {
		return fmt.Errorf("slice length is not equal to %d", n)
	}
	return nil
}

// SliceLenNotEqual asserts that the length of the slice is not equal to n.
func SliceLenNotEqual[T any](slice []T, n int) error {
	if len(slice) == n {
		return fmt.Errorf("slice length is not not equal to %d", n)
	}
	return nil
}
