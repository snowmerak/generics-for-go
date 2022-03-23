package assert

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Less is check what a is less than b.
func Less[T constraints.Ordered](a, b T) error {
	if a < b {
		return nil
	}
	return fmt.Errorf("%v is not less than %v", a, b)
}

// LessEqual is check what a is less than or equal to b.
func LessEqual[T constraints.Ordered](a, b T) error {
	if a <= b {
		return nil
	}
	return fmt.Errorf("%v is not less than or equal to %v", a, b)
}

// Greater is check what a is greater than b.
func Greater[T constraints.Ordered](a, b T) error {
	if a > b {
		return nil
	}
	return fmt.Errorf("%v is not greater than %v", a, b)
}

// GreaterEqual is check what a is greater than or equal to b.
func GreaterEqual[T constraints.Ordered](a, b T) error {
	if a >= b {
		return nil
	}
	return fmt.Errorf("%v is not greater than or equal to %v", a, b)
}

// Equal is check what a is equal to b.
func Equal[T comparable](a, b T) error {
	if a == b {
		return nil
	}
	return fmt.Errorf("%v is not equal to %v", a, b)
}
