package copyable

type Copyable[T any] interface {
	CopyTo(*T) error
}

type Clonable[T any] interface {
	Clone() T
}
