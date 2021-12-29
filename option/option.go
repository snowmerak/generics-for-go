package option

type Option[T any] struct {
	value any
}

func (o Option[T]) Ok() bool {
	_, ok := o.value.(T)
	return ok
}

func (o Option[T]) Unwrap() T {
	value, _ := o.value.(T)
	return value
}

func Some[T any](value T) Option[T] {
	return Option[T]{value: value}
}

func None[T any]() Option[T] {
	return Option[T]{value: nil}
}
