// trie
package trie

import "golang.org/x/exp/constraints"

type node[T constraints.Ordered] struct {
	value   T
	less    *node[T]
	equal   *node[T]
	greater *node[T]
}

type Trie[T constraints.Ordered] struct {
	root *node[T]
}

func New[T constraints.Ordered]() *Trie[T] {
	return new(Trie[T])
}

func (t *Trie[T]) Insert(x ...T) error {
	return nil
}
