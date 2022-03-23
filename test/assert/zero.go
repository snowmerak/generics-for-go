package assert

import (
	"fmt"
	"reflect"
)

// Zero is check what a is zero value.
func Zero[T any](a T) error {
	if reflect.ValueOf(a).IsZero() {
		return nil
	}
	return fmt.Errorf("%v is not zero value", a)
}

// NotZero is check what a is not zero value.
func NotZero[T any](a T) error {
	if !reflect.ValueOf(a).IsZero() {
		return nil
	}
	return fmt.Errorf("%v is zero value", a)
}
