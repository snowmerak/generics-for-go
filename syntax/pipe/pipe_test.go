package pipe_test

import (
	"fmt"
	"strconv"

	"github.com/snowmerak/generics-for-go/v2/syntax/pipe"
)

func Example() {
	funs := pipe.Link[int](func(i int) int {
		return i + 1
	}, func(i int) (int, int) {
		return i, i * i
	}, func(a, b int) string {
		return fmt.Sprintf("%d%d", a, b)
	}, func(s string) int {
		v, err := strconv.Atoi(s)
		if err != nil {
			return 0
		}
		return v
	})
	fmt.Println(funs(10))
	// Output:
	// 11121
}

func Example_withError() {
	funs := pipe.Link[struct {
		Result int
		Err    error
	}](func(i int) int {
		return i + 1
	}, func(i int) (int, int) {
		return i, i * i
	}, func(a, b int) string {
		return fmt.Sprintf("%d%d!", a, b)
	}, func(s string) (int, error) {
		v, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		return v, nil
	})
	fmt.Println(funs(10))
	// Output:
	// {0 strconv.Atoi: parsing "11121!": invalid syntax}
}
