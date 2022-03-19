package iterable_test

import (
	"fmt"

	"github.com/snowmerak/generics-for-go/v2/iterable"
)

func Example_fromSlice() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	iter := iterable.FromSlice(a)
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
	// Output:
	// 0 1 true
	// 1 2 true
	// 2 3 true
	// 3 4 true
	// 4 5 true
	// 5 6 true
	// 6 7 true
	// 7 8 true
	// 8 9 true
	// 9 10 true
}

func Example_fromMap() {
	a := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
	}
	iter := iterable.FromMap(a)
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
	// Output:
	// 1 one true
	// 2 two true
	// 3 three true
	// 4 four true
	// 5 five true
}

func Example_generator() {
	generator := iterable.NewGenerator(0, 0, func(k int, v int) (int, int) {
		return k, v + 1
	})
	count := 0
	for generator.HasNext() && count < 10 {
		fmt.Println(generator.Next())
		count++
	}
	// Output:
	// 0 1 true
	// 0 2 true
	// 0 3 true
	// 0 4 true
	// 0 5 true
	// 0 6 true
	// 0 7 true
	// 0 8 true
	// 0 9 true
	// 0 10 true
}
