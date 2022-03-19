// cond gives condition control flow
package cond

// ifCond is a pair of condition function and result function.
type ifCond[T, R any] struct {
	fun    func(T) bool
	result func(T) R
}

// IfCond is a if condition control flow structure.
type IfCond[T, R any] struct {
	conds        []ifCond[T, R]
	defaultValue func(T) R
}

// If returns a new IfCond with given condition.
func If[T, R any](condition func(T) bool, result func(T) R) *IfCond[T, R] {
	c := new(IfCond[T, R])
	c.conds = []ifCond[T, R]{
		{
			fun:    condition,
			result: result,
		},
	}
	return c
}

// ElseIf adds a new condition to the IfCond.
func (c *IfCond[T, R]) ElseIf(condition func(T) bool, result func(T) R) *IfCond[T, R] {
	c.conds = append(c.conds, ifCond[T, R]{
		fun:    condition,
		result: result,
	})
	return c
}

// Else adds a default result function to the IfCond.
func (c *IfCond[T, R]) Else(defaultElse func(T) R) *IfCond[T, R] {
	c.defaultValue = defaultElse
	return c
}

// Run runs the IfCond with given value.
func (c *IfCond[T, R]) Run(t T) (rs R) {
	for _, cond := range c.conds {
		ok := cond.fun(t)
		if ok {
			return cond.result(t)
		}
	}
	if c.defaultValue == nil {
		return rs
	}
	return c.defaultValue(t)
}
