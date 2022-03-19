// option is a type that represents optional values.
package option

// Option is a container for a value of type T.
type Option[T any] struct {
	value  T
	isSome bool
}

// Some returns a new Option with the given value.
func Some[T any](value T) *Option[T] {
	return &Option[T]{
		value:  value,
		isSome: true,
	}
}

// None returns a new Option with no value.
func None[T any]() *Option[T] {
	return &Option[T]{
		isSome: false,
	}
}

// IsSome returns true if the Option has a value.
func (o *Option[T]) IsSome() bool {
	return o.isSome
}

// Unwrap returns the value of the Option.
// if Option does not have a value, returns zero value of T.
func (o *Option[T]) Unwrap() (rs T) {
	if o.isSome {
		return o.value
	}
	return rs
}

// UnwrapOr returns the value of the Option or the given value.
func (o *Option[T]) UnwrapOr(defaultValue T) T {
	if o == nil {
		return defaultValue
	}
	if o.isSome {
		return o.Unwrap()
	}
	return defaultValue
}

// Map transforms the Option by applying the given function to the value.
// if Option does not have a value, do not applying.
// if given function returns false, Option will be changed to None.
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

// MapOr transforms the Option by applying the given function to the value.
// if Option does not have a value or given function returns false, Option will be have defaultValue.
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

// MapOrElse returns new Option by applying the given function to old Option.
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
