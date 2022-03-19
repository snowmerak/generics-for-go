package cond

// SwitchCond is a switch condition control flow structure.
type SwitchCond[T comparable, R any] struct {
	conds        map[T]func(T) R
	defaultValue func(T) R
}

// Switch returns a new SwitchCond.
func Switch[T comparable, R any]() *SwitchCond[T, R] {
	return &SwitchCond[T, R]{
		conds:        make(map[T]func(T) R),
		defaultValue: nil,
	}
}

// Case adds a new condition to the SwitchCond.
func (s *SwitchCond[T, R]) Case(value T, result func(T) R) *SwitchCond[T, R] {
	s.conds[value] = result
	return s
}

// Default adds a default result function to the SwitchCond.
func (s *SwitchCond[T, R]) Default(defaultElse func(T) R) *SwitchCond[T, R] {
	s.defaultValue = defaultElse
	return s
}

// Run runs the SwitchCond with given value.
func (s *SwitchCond[T, R]) Run(t T) (rs R) {
	if s.conds[t] != nil {
		return s.conds[t](t)
	}
	if s.defaultValue == nil {
		return rs
	}
	return s.defaultValue(t)
}
