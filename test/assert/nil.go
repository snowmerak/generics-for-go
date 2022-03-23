package assert

import "fmt"

// Nil asserts that the value is nil.
func Nil[T any](a *T) error {
	if a == nil {
		return nil
	}
	return fmt.Errorf("%v is not nil", a)
}

// NotNil asserts that the value is not nil.
func NotNil[T any](a *T) error {
	if a != nil {
		return nil
	}
	return fmt.Errorf("%v is nil", a)
}
