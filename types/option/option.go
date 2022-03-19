package option

type Option[T any] struct {
	value  T
	isSome bool
}

func Some[T any](value T) *Option[T] {
	return &Option[T]{
		value:  value,
		isSome: true,
	}
}

func None[T any]() *Option[T] {
	return &Option[T]{
		isSome: false,
	}
}

func (o *Option[T]) IsSome() bool {
	return o.isSome
}

func (o *Option[T]) Unwrap() (rs T) {
	if o.isSome {
		return o.value
	}
	return rs
}

func (o *Option[T]) UnwrapOr(defaultValue T) T {
	if o == nil {
		return defaultValue
	}
	if o.isSome {
		return o.Unwrap()
	}
	return defaultValue
}

func (o *Option[T]) Map(fn func(T) (T, bool)) {
	if o.isSome {
		value, b := fn(o.Unwrap())
		if !b {
			o.isSome = false
			return
		}
		o.value = value
	}
}

func (o *Option[T]) MapOr(fn func(T) (T, bool), defaultValue T) {
	if o.isSome {
		value, b := fn(o.Unwrap())
		if !b {
			o.value = defaultValue
			return
		}
		o.value = value
		return
	}
	o.value = defaultValue
}

func (o *Option[T]) AndThen(fn func(T) (T, bool)) *Option[T] {
	if o.isSome {
		value, b := fn(o.Unwrap())
		if !b {
			return None[T]()
		}
		return Some(value)
	}
	return None[T]()
}
