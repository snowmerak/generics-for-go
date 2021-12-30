package result

type Result[T any] struct {
	value any
}

func Success[T any](data T) *Result[T] {
	return &Result[T]{value: data}
}

func Ok[T any](data T) *Result[T] {
	return &Result[T]{value: data}
}

func Failed[T any](err error) *Result[T] {
	return &Result[T]{value: err}
}

func Err[T any](err error) *Result[T] {
	return &Result[T]{value: err}
}

func (r *Result[T]) Replace(fn func(T) (T, error)) {
	if r.Ok() {
		v, err := fn(r.Unwrap())
		if err != nil {
			r.value = nil
			return
		}
		r.value = v
		return
	}
}

func (r *Result[T]) ReplaceOr(fn func(T) (T, error), defaultValue T) {
	if r.Ok() {
		v, err := fn(r.Unwrap())
		if err == nil {
			r.value = v
			return
		}
	}
	r.value = defaultValue
}

func (r *Result[T]) Ok() bool {
	_, ok := r.value.(error)
	return !ok
}

func (r *Result[T]) Err() error {
	err, ok := r.value.(error)
	if !ok {
		return nil
	}
	return err
}

func (r *Result[T]) AndThen(fn func(T) (T, error)) *Result[T] {
	if r.Ok() {
		v, err := fn(r.Unwrap())
		if err != nil {
			return Failed[T](err)
		}
		return Ok(v)
	}
	return Failed[T](r.Err())
}

func (r *Result[T]) Unwrap() T {
	value, _ := r.value.(T)
	return value
}

func (r *Result[T]) UnwrapOr(defaultValue T) T {
	value, ok := r.value.(T)
	if !ok {
		return defaultValue
	}
	return value
}

func (r *Result[T]) UnwrapOrPanic(err error) T {
	if r.Ok() {
		return r.Unwrap()
	}
	panic(err)
}
