package assert

import "fmt"

// StrLenLess asserts that the length of the string is less than n.
func StrLenLess(s string, n int) error {
	if len(s) >= n {
		return fmt.Errorf("string length is not less than %d", n)
	}
	return nil
}

// StrLenLessEqual asserts that the length of the string is less than or equal to n.
func StrLenLessEqual(s string, n int) error {
	if len(s) > n {
		return fmt.Errorf("string length is not less than or equal to %d", n)
	}
	return nil
}

// StrLenGreater asserts that the length of the string is greater than n.
func StrLenGreater(s string, n int) error {
	if len(s) <= n {
		return fmt.Errorf("string length is not greater than %d", n)
	}
	return nil
}

// StrLenGreaterEqual asserts that the length of the string is greater than or equal to n.
func StrLenGreaterEqual(s string, n int) error {
	if len(s) < n {
		return fmt.Errorf("string length is not greater than or equal to %d", n)
	}
	return nil
}

// StrLenEqual asserts that the length of the string is equal to n.
func StrLenEqual(s string, n int) error {
	if len(s) != n {
		return fmt.Errorf("string length is not equal to %d", n)
	}
	return nil
}

// StrLenNotEqual asserts that the length of the string is not equal to n.
func StrLenNotEqual(s string, n int) error {
	if len(s) == n {
		return fmt.Errorf("string length is not not equal to %d", n)
	}
	return nil
}
