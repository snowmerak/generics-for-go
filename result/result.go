package result

type Result[T any] struct {
	value any
}

func (r Result[T]) Ok() bool {
	_, ok := r.value.(error)
	return !ok
}

func (r Result[T]) Err() error {
	err, ok := r.value.(error)
	if !ok {
		return nil
	}
	return err
}

func (r Result[T]) Unwrap() T {
	value, _ := r.value.(T)
	return value
}

func Success[T any](data T) Result[T] {
	return Result[T]{value: data}
}

func Ok[T any](data T) Result[T] {
	return Result[T]{value: data}
}

func Failed[T any](err error) Result[T] {
	return Result[T]{value: err}
}
