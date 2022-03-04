package deque

type node[T any] struct {
	values [16]T
	start  int
	end    int
	prev   *node[T]
	next   *node[T]
}

type Deque[T any] struct {
	head *node[T]
	tail *node[T]
}
