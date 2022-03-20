// collections package contains collection interfaces of generic types.
package collections

type Clonable[T any] interface {
	Clone() T
}
