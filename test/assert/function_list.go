package assert

import (
	"errors"
	"fmt"
	"strings"
)

type Assert struct {
	errs []error
}

func Of(assertions ...error) Assert {
	errs := make([]error, len(assertions))
	for _, assertion := range assertions {
		if assertion != nil {
			errs = append(errs, assertion)
		}
	}
	return Assert{errs: errs}
}

func (a Assert) Result() error {
	sb := strings.Builder{}
	for i, err := range a.errs {
		sb.WriteString(fmt.Sprintf("%d: %s\n", i, err))
	}
	return errors.New(sb.String())
}
