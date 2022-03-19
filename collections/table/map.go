package table

type Table[K comparable, V any] map[K]V

func New[K comparable, V any]() map[K]V {
	return make(map[K]V)
}

func FromSlices[K comparable, V any](ks []K, vs []V) map[K]V {
	r := make(map[K]V)
	for i := 0; i < len(ks) && i < len(vs); i++ {
		r[ks[i]] = vs[i]
	}
	return r
}
