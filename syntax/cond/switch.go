package cond

type SwitchCond[T comparable, R any] struct {
	conds        map[T]func(T) R
	defaultValue func(T) R
}

func Switch[T comparable, R any]() *SwitchCond[T, R] {
	return &SwitchCond[T, R]{
		conds:        make(map[T]func(T) R),
		defaultValue: nil,
	}
}

func (s *SwitchCond[T, R]) Case(value T, result func(T) R) *SwitchCond[T, R] {
	s.conds[value] = result
	return s
}

func (s *SwitchCond[T, R]) Default(defaultElse func(T) R) *SwitchCond[T, R] {
	s.defaultValue = defaultElse
	return s
}

func (s *SwitchCond[T, R]) Do(t T) (rs R) {
	if s.conds[t] != nil {
		return s.conds[t](t)
	}
	if s.defaultValue == nil {
		return rs
	}
	return s.defaultValue(t)
}
