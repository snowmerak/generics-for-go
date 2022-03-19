package pointer_test

import (
	"fmt"

	"github.com/snowmerak/generics-for-go/v2/types/pointer"
)

func Example_simpleNewAndUse() {
	i := 10
	a := pointer.New(&i)
	*a.Load() = 20
	fmt.Println(i)
	// Output:
	// 20
}

func Example_swapPointer() {
	i := 10
	a := pointer.New(&i)
	n := 20
	a.CompareAndSwap(&n, &i)
	fmt.Println(*a.Load())
	// Output:
	// 20
}
