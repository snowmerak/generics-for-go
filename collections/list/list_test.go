package list_test

import (
	"fmt"

	"github.com/snowmerak/generics-for-go/v2/collections/list"
)

func Example() {
	list := list.New[int]()
	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Push(4)
	list.Foreach(func(i int) {
		fmt.Print(i, " ")
	})
	fmt.Println()
	// Output:
	// 1 2 3 4
}

func Example_stack() {
	list := list.Of(0)
	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Push(4)
	for !list.IsEmpty() {
		fmt.Print(list.Pop(), " ")
	}
	fmt.Println()
	// Output:
	// 4 3 2 1 0
}

func Example_queue() {
	list := list.Of(0)
	list.Push(1)
	list.Push(2)
	list.Push(3)
	list.Push(4)
	for !list.IsEmpty() {
		fmt.Print(list.Shift(), " ")
	}
	fmt.Println()
	// Output:
	// 0 1 2 3 4
}

func Example_reverse() {
	list := list.Of(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	list.Reverse()
	list.Foreach(func(i int) {
		fmt.Print(i, " ")
	})
	fmt.Println()
	// Output:
	// 10 9 8 7 6 5 4 3 2 1 0
}

func Example_squared() {
	list := list.Of(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	list.Map(func(i int) int {
		return i * i
	})
	list.Foreach(func(i int) {
		fmt.Print(i, " ")
	})
	fmt.Println()
	// Output:
	// 0 1 4 9 16 25 36 49 64 81 100
}
