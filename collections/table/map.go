// table is giving Table structure.
package table

// Table is a map[K]V.
type Table[K comparable, V any] map[K]V

// new creates a new Table.
func New[K comparable, V any]() Table[K, V] {
	return make(Table[K, V])
}

// FromSlices creates a new Table from the given slices.
// key is the index of the slice.
func FromSlices[K comparable, V any](ks []K, vs []V) map[K]V {
	r := make(map[K]V)
	for i := 0; i < len(ks) && i < len(vs); i++ {
		r[ks[i]] = vs[i]
	}
	return r
}
