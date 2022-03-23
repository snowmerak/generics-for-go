package assert

import "fmt"

// MapEqual asserts that the value is equal to the expected value.
func MapEqual[K, V comparable](a, b map[K]V) error {
	if len(a) != len(b) {
		return fmt.Errorf("%v is not equal to %v", a, b)
	}
	for k, v := range a {
		if v != b[k] {
			return fmt.Errorf("%v is not equal to %v", a, b)
		}
	}
	return nil
}

// MapDeepEqual asserts that the value is equal to the expected value.
func MapNotEqual[K, V comparable](a, b map[K]V) error {
	if len(a) != len(b) {
		return nil
	}
	for k, v := range a {
		if v != b[k] {
			return nil
		}
	}
	return fmt.Errorf("%v is equal to %v", a, b)
}
