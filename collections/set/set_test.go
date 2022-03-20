package set_test

import (
	"fmt"

	"github.com/snowmerak/generics-for-go/v2/collections/set"
)

func Example_intersect() {
	a := set.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	b := set.Of(2, 4, 6, 8, 10)
	c := a.Intersect(b)
	c.Foreach(func(i int) {
		fmt.Print(i, " ")
	})
	fmt.Println()
	// Output:
	// 2 4 6 8 10
}

func Example_subtract() {
	a := set.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	b := set.Of(2, 4, 6, 8, 10)
	c := a.Subtract(b)
	c.Foreach(func(i int) {
		fmt.Print(i, " ")
	})
	fmt.Println()
	// Output:
	// 1 3 5 7 9
}

func Example_union() {
	a := set.Of(1, 3, 5, 7, 9)
	b := set.Of(2, 4, 6, 8, 10)
	c := a.Union(b)
	c.Foreach(func(i int) {
		fmt.Print(i, " ")
	})
	fmt.Println()
	// Output:
	// 1 2 3 4 5 6 7 8 9 10
}

func Example_addRemove() {
	a := set.Of(1, 3, 5, 7, 9)
	a.Add(2)
	a.Add(4)
	a.Remove(3)
	a.Remove(5)
	a.Foreach(func(i int) {
		fmt.Print(i, " ")
	})
	fmt.Println()
	// Output:
	// 1 2 4 7 9
}

func Example_contains() {
	a := set.Of(1, 3, 5, 7, 9)
	fmt.Println(a.Contains(4))
	// Output:
	// false
}
