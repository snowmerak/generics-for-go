package assert

import "fmt"

// A is a custom assertion function.
func A(b bool, format string, a ...interface{}) error {
	if b {
		return nil
	}
	return fmt.Errorf(format, a...)
}
