package result

// Map returns new Result by applying the given function to the old result.
func Map[T any, R any](param *Result[T], fn func(T) (R, error)) *Result[R] {
	if param.isOk {
		r, err := fn(param.Unwrap())
		if err != nil {
			return Err[R](err)
		}
		return OK(r)
	}
	return Err[R](param.Err())
}

// ToFunctor returns new function with Result by applying the given function to the old result.
func ToFunctor[T any, R any](fn func(T) (R, error)) func(*Result[T]) *Result[R] {
	return func(param *Result[T]) *Result[R] {
		if param.isOk {
			n, err := fn(param.Unwrap())
			if err != nil {
				return Err[R](err)
			}
			return OK(n)
		}
		return Err[R](param.Err())
	}
}
