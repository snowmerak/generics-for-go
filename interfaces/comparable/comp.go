package comparable

// Compared is alias for int8.
type Compared int8

const (
	// Less is alias for 0.
	Less = Compared(iota)
	// Equal is alias for 1.
	Equal
	// Greater is alias for 2.
	Greater
)

// Comparable is interface for comparable objects.
type Comparable[T any] interface {
	CompareTo(Comparable[T]) Compared
}
