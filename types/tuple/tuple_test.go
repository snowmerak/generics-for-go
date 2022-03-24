package tuple_test

import (
	"fmt"
	"strings"

	"github.com/snowmerak/generics-for-go/v2/types/tuple"
)

func Example() {
	a := tuple.N5("Hello", "world", 1, 2, 3)
	fmt.Println(strings.Repeat(a.T1(), a.T3()))
	fmt.Println(strings.Repeat(a.T2(), a.T4()))
	fmt.Println(a.T5() * a.T5())
	// Output:
	// Hello
	// WorldWorld
	// 9
}
