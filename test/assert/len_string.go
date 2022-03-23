package assert

import "fmt"

// StringLenLess asserts that the length of the string is less than n.
func StringLenLess(s string, n int) error {
	if len(s) >= n {
		return fmt.Errorf("string length is not less than %d", n)
	}
	return nil
}

// StringLenLessEqual asserts that the length of the string is less than or equal to n.
func StringLenLessEqual(s string, n int) error {
	if len(s) > n {
		return fmt.Errorf("string length is not less than or equal to %d", n)
	}
	return nil
}

// StringLenGreater asserts that the length of the string is greater than n.
func StringLenGreater(s string, n int) error {
	if len(s) <= n {
		return fmt.Errorf("string length is not greater than %d", n)
	}
	return nil
}

// StringLenGreaterEqual asserts that the length of the string is greater than or equal to n.
func StringLenGreaterEqual(s string, n int) error {
	if len(s) < n {
		return fmt.Errorf("string length is not greater than or equal to %d", n)
	}
	return nil
}

// StringLenEqual asserts that the length of the string is equal to n.
func StringLenEqual(s string, n int) error {
	if len(s) != n {
		return fmt.Errorf("string length is not equal to %d", n)
	}
	return nil
}

// StringLenNotEqual asserts that the length of the string is not equal to n.
func StringLenNotEqual(s string, n int) error {
	if len(s) == n {
		return fmt.Errorf("string length is not not equal to %d", n)
	}
	return nil
}
