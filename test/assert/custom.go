package assert

import "fmt"

// If is a custom assertion function.
func If(b bool, format string, a ...interface{}) error {
	if b {
		return nil
	}
	return fmt.Errorf(format, a...)
}
