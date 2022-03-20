// collections package contains collection interfaces of generic types.
package collections

// Clonable is a interface for cloning.
type Clonable[T any] interface {
	Clone() T
}
