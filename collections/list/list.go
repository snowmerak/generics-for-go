// list is giving double linked list.
package list

// node a node of list.
type node[T any] struct {
	value T
	prev  *node[T]
	next  *node[T]
}

// List is a double linked list.
type List[T any] struct {
	head *node[T]
	tail *node[T]
}

// New returns a new List.
func New[T any]() *List[T] {
	return &List[T]{}
}

// Of returns a new List with the given elements.
func Of[T any](values ...T) *List[T] {
	l := New[T]()
	if len(values) == 0 {
		return l
	}
	tail := &node[T]{
		value: values[0],
		prev:  nil,
		next:  nil,
	}
	l.head = tail
	for _, v := range values[1:] {
		tail.next = &node[T]{
			value: v,
			prev:  tail,
			next:  nil,
		}
		tail = tail.next
	}
	l.tail = tail
	return l
}
