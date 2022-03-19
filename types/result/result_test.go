package result_test

import (
	"errors"
	"fmt"

	"github.com/snowmerak/generics-for-go/v2/types/result"
)

func Example_ok() {
	a := result.OK(10)
	fmt.Println(a.Unwrap())
	// Output:
	// 10
}

func Example_err() {
	a := result.Err[int](errors.New("error"))
	fmt.Println(a.Err())
	// Output:
	// error
}

func Example_switch() {
	a := result.OK(10)
	switch a.IsOK() {
	case true:
		fmt.Println(a.Unwrap())
	case false:
		fmt.Println(a.Err())
	}
	// Output:
	// 10
}
