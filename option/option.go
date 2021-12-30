package option

type Option[T any] struct {
	value any
}

func Some[T any](value T) *Option[T] {
	return &Option[T]{value: value}
}

func None[T any]() *Option[T] {
	return &Option[T]{value: nil}
}

func (o *Option[T]) Ok() bool {
	_, ok := o.value.(T)
	return ok
}

func (o *Option[T]) Unwrap() T {
	value, _ := o.value.(T)
	return value
}

func (o *Option[T]) UnwrapOr(defaultValue T) T {
	if o.Ok() {
		return o.Unwrap()
	}
	return defaultValue
}

func (o *Option[T]) Replace(fn func(T) (T, error)) {
	if o.Ok() {
		value, err := fn(o.Unwrap())
		if err != nil {
			o.value = nil
		} else {
			o.value = value
		}
	}
}

func (o *Option[T]) ReplaceOr(fn func(T) (T, error), defaultValue T) {
	if o.Ok() {
		value, err := fn(o.Unwrap())
		if err != nil {
			o.value = defaultValue
		} else {
			o.value = value
			return
		}
	}
	o.value = defaultValue
}

func (o *Option[T]) AndThen(fn func(T) (T, error)) *Option[T] {
	if o.Ok() {
		value, err := fn(o.Unwrap())
		if err != nil {
			return None[T]()
		}
		return Some(value)
	}
	return None[T]()
}
