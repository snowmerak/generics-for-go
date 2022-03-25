package copyable

// Copyable copy to given object.
type Copyable[T any] interface {
	CopyTo(*T) error
}

// Clonable returns clone object.
type Clonable[T any] interface {
	Clone() T
}
