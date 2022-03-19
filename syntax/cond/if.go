package cond

type ifCond[T, R any] struct {
	fun    func(T) bool
	result func(T) R
}

type IfCond[T, R any] struct {
	conds        []ifCond[T, R]
	defaultValue func(T) R
}

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

func (c *IfCond[T, R]) ElseIf(condition func(T) bool, result func(T) R) *IfCond[T, R] {
	c.conds = append(c.conds, ifCond[T, R]{
		fun:    condition,
		result: result,
	})
	return c
}

func (c *IfCond[T, R]) Else(defaultElse func(T) R) *IfCond[T, R] {
	c.defaultValue = defaultElse
	return c
}

func (c *IfCond[T, R]) Do(t T) (rs R) {
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
