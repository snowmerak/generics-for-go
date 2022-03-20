// collections package contains collection interfaces of generic types.
package collections

type Clone[T any] interface {
	Clone() T
}
