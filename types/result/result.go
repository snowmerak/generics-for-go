// result is a helper for handling errors.
package result

// Result is a container for values, T and error, bool.
type Result[T any] struct {
	value T
	err   error
	isOk  bool
}

// OK returns a new Result with the given value.
func OK[T any](data T) *Result[T] {
	return &Result[T]{
		value: data,
		err:   nil,
		isOk:  true,
	}
}

// Err returns a new Result with the given error.
func Err[T any](err error) *Result[T] {
	return &Result[T]{
		err:  err,
		isOk: false,
	}
}

// Map transforms the Result by applying the given function to the value.
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

// MapOr transforms the Result by applying the given function to the value.
// if given function returns false, Result will be changed to OK(defaultValue).
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

// IsOK returns true if the Result has a value.
func (r *Result[T]) IsOK() bool {
	return r.isOk
}

// Err returns the error of the Result.
func (r *Result[T]) Err() error {
	return r.err
}

// AndThen returns new Result by applying the given function to old Result.
func (r *Result[T]) AndThen(fn func(T) (T, error)) *Result[T] {
	if r.isOk {
		v, err := fn(r.Unwrap())
		if err != nil {
			return Err[T](err)
		}
		return OK(v)
	}
	return Err[T](r.Err())
}

// Unwrap returns the value of the Result.
func (r *Result[T]) Unwrap() T {
	return r.value
}

// UnwrapOr returns the value of the Result or the given value.
func (r *Result[T]) UnwrapOr(defaultValue T) T {
	if r == nil {
		return defaultValue
	}
	if !r.isOk {
		return defaultValue
	}
	return r.value
}

// UnwrapOrPanic returns the value of the Result.
// if Result does not have a value, panic with result's error.
func (r *Result[T]) UnwrapOrPanic() T {
	if r.isOk {
		return r.Unwrap()
	}
	panic(r.err)
}

// UnwrapOrCustomPanic returns the value of the Result.
// if Result does not have a value, panic with custom message.
func (r *Result[T]) UnwrapOrCustomPanic(err error) T {
	if r.isOk {
		return r.Unwrap()
	}
	panic(err)
}
