package result

type Result[T any] struct {
	value T
	err   error
	isOk  bool
}

func Ok[T any](data T) *Result[T] {
	return &Result[T]{
		value: data,
		err:   nil,
		isOk:  true,
	}
}

func Err[T any](err error) *Result[T] {
	return &Result[T]{
		err:  err,
		isOk: false,
	}
}

func (r *Result[T]) Map(fn func(T) (T, error)) {
	if r.isOk {
		v, err := fn(r.Unwrap())
		if err != nil {
			r.err = err
			r.isOk = false
			return
		}
		r.value = v
		return
	}
}

func (r *Result[T]) MapOr(fn func(T) (T, error), defaultValue T) {
	if r.isOk {
		v, err := fn(r.Unwrap())
		if err == nil {
			r.value = v
			return
		}
	}
	r.value = defaultValue
}

func (r *Result[T]) IsOk() bool {
	return r.isOk
}

func (r *Result[T]) Err() error {
	return r.err
}

func (r *Result[T]) AndThen(fn func(T) (T, error)) *Result[T] {
	if r.isOk {
		v, err := fn(r.Unwrap())
		if err != nil {
			return Err[T](err)
		}
		return Ok(v)
	}
	return Err[T](r.Err())
}

func (r *Result[T]) Unwrap() T {
	return r.value
}

func (r *Result[T]) UnwrapOr(defaultValue T) T {
	if r == nil {
		return defaultValue
	}
	if !r.isOk {
		return defaultValue
	}
	return r.value
}

func (r *Result[T]) UnwrapOrPanic() T {
	if r.isOk {
		return r.Unwrap()
	}
	panic(r.err)
}

func (r *Result[T]) UnwrapOrCustomPanic(err error) T {
	if r.isOk {
		return r.Unwrap()
	}
	panic(err)
}
