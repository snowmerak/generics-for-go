// assert is a assertion library for test.
package assert

import (
	"errors"
	"fmt"
	"strings"
)

// Assert is a assertion list.
type Assert struct {
	errs []error
}

// Of is a constructor of Assert.
func Of(assertions ...error) Assert {
	errs := make([]error, 0, len(assertions))
	for _, assertion := range assertions {
		if assertion != nil {
			errs = append(errs, assertion)
		}
	}
	return Assert{errs: errs}
}

// Result return the assertions result.
func (a Assert) Result() error {
	sb := strings.Builder{}
	for i, err := range a.errs {
		sb.WriteString(fmt.Sprintf("%d: %s\n", i, err))
	}
	return errors.New(sb.String())
}
