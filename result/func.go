package result

func Map[T any, R any](param *Result[T], fn func(T) *Result[R]) *Result[R] {
	if param.Ok() {
		return fn(param.Unwrap())
	}
	return Err[R](param.Err())
}

func ToFunctor[T any, R any](fn func(T) (R, error)) func(*Result[T]) *Result[R] {
	return func(param *Result[T]) *Result[R] {
		if param.Ok() {
			n, err := fn(param.Unwrap())
			if err != nil {
				return Err[R](err)
			}
			return Ok(n)
		}
		return Err[R](param.Err())
	}
}
