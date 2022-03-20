package table_test

import (
	"fmt"

	"github.com/snowmerak/generics-for-go/v2/collections/table"
)

func Example() {
	m := table.New[int, string]()
	m[0] = "zero"
	m[1] = "one"
	m[2] = "two"
	fmt.Println(m)
	// Output:
	// map[0:zero 1:one 2:two]
}

func Example_fromSlice() {
	m := table.FromSlices([]int{0, 1, 2}, []string{"zero", "one", "two"})
	fmt.Println(m)
	// Output:
	// map[0:zero 1:one 2:two]
}

func Example_allAny() {
	m := table.FromSlices([]int{0, 1, 2}, []string{"zero", "one", "two"})
	r := m.All(func(k int, v string) bool {
		return k == 1
	})
	fmt.Println(r)
	r = m.Any(func(k int, v string) bool {
		return k == 3
	})
	fmt.Println(r)
	// Output:
	// false
	// false
}

func Example_filter() {
	m := table.FromSlices([]int{0, 1, 2}, []string{"zero", "one", "two"})
	r := m.Filter(func(i int, s string) bool {
		return i == 1
	})
	fmt.Println(r)
	// Output:
	// map[1:one]
}

func Example_map() {
	m := table.FromSlices([]int{0, 1, 2}, []string{"zero", "one", "two"})
	r := m.Map(func(i int, s string) (int, string) {
		return i * i, s + s
	})
	fmt.Println(r)
	// Output:
	// map[0:zerozero 1:oneone 4:twotwo]
}

func Example_reduceKey() {
	m := table.FromSlices([]int{0, 1, 2}, []string{"zero", "one", "two"})
	sumKey := m.ReduceKey(func(i int, v int) int {
		return i + v
	}, 0)
	fmt.Println(sumKey)
	// Output:
	// 3
}

func Example_reduceValue() {
	m := table.FromSlices([]int{0, 1, 2}, []string{"zero", "one", "two"})
	sumKey := m.ReduceValue(func(i string, v string) string {
		return i + v
	}, "")
	fmt.Println(sumKey)
	// Output:
	// zeroonetwo
}
