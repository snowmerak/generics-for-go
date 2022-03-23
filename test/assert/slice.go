package assert

import "fmt"

// SliceEqual asserts that the value is equal to the expected value.
func SliceEqual[T comparable](a []T, b ...T) error {
	if len(a) != len(b) {
		return fmt.Errorf("%v is not equal to %v", a, b)
	}
	for i, v := range a {
		if v != b[i] {
			return fmt.Errorf("%v is not equal to %v", a, b)
		}
	}
	return nil
}

// SliceNotEqual asserts that the value is not equal to the expected value.
func SliceNotEqual[T comparable](a []T, b ...T) error {
	if len(a) != len(b) {
		return nil
	}
	for i, v := range a {
		if v != b[i] {
			return nil
		}
	}
	return fmt.Errorf("%v is equal to %v", a, b)
}
