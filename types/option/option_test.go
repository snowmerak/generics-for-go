package option_test

import (
	"fmt"

	"github.com/snowmerak/generics-for-go/v2/types/option"
)

func ExampleSome_int() {
	a := option.Some(1)
	b := a.Unwrap()
	fmt.Println(b)
	// Output:
	// 1
}

func ExampleSome_string() {
	a := option.Some("example string")
	b := a.Unwrap()
	fmt.Println(b)
	// Output:
	// example string
}

func ExampleNone_int() {
	a := option.None[int]()
	b := a.Unwrap()
	fmt.Println(b)
	// Output:
	// 0
}

func ExampleNone_string() {
	a := option.None[string]()
	b := a.Unwrap()
	fmt.Println(b)
	// Output:
	// ""
}

func ExampleOption_AndThen_true() {
	a := option.Some(1)
	b := a.AndThen(func(i int) (int, bool) {
		return i + 1, true
	})
	c := b.Unwrap()
	fmt.Println(c)
	// Output:
	// 2
}

func ExampleOption_AndThen_false() {
	a := option.Some(1)
	b := a.AndThen(func(i int) (int, bool) {
		return i + 1, false
	})
	c := b.Unwrap()
	fmt.Println(c)
	// Output:
	// 0
}

func ExampleOption_IsSome_some() {
	a := option.Some(1)
	b := a.IsSome()
	fmt.Println(b)
	// Output:
	// true
}

func ExampleOption_IsSome_none() {
	a := option.None[int]()
	b := a.IsSome()
	fmt.Println(b)
	// Output:
	// false
}

func ExampleOption_Map_someTrue() {
	a := option.Some(1)
	a.Map(func(i int) (int, bool) {
		return i + 1, true
	})
	c := a.Unwrap()
	fmt.Println(c)
	// Output:
	// 2
}

func ExampleOption_Map_someFalse() {
	a := option.Some(1)
	a.Map(func(i int) (int, bool) {
		return i + 1, false
	})
	c := a.Unwrap()
	fmt.Println(c)
	// Output:
	// 0
}

func ExampleOption_Map_noneTrue() {
	a := option.None[int]()
	a.Map(func(i int) (int, bool) {
		return i + 1, true
	})
	c := a.Unwrap()
	fmt.Println(c)
	// Output:
	// 1
}

func ExampleOption_Map_noneFalse() {
	a := option.None[int]()
	a.Map(func(i int) (int, bool) {
		return i + 1, false
	})
	c := a.Unwrap()
	fmt.Println(c)
	// Output:
	// 0
}

func ExampleOption_MapOr_true() {
	a := option.Some(1)
	a.MapOr(func(i int) (int, bool) {
		return i + 1, true
	}, -99)
	fmt.Println(a)
	// Output:
	// 2
}

func ExampleOption_MapOr_false() {
	a := option.Some(1)
	a.MapOr(func(i int) (int, bool) {
		return i + 1, false
	}, -99)
	fmt.Println(a)
	// Output:
	// -99
}

func ExampleOption_Unwrap_true() {
	a := option.Some(1)
	b := a.Unwrap()
	fmt.Println(b)
	// Output:
	// 1
}

func ExampleOption_Unwrap_false() {
	a := option.None[int]()
	b := a.Unwrap()
	fmt.Println(b)
	// Output:
	// 0
}

func ExampleOption_UnwrapOr_true() {
	a := option.Some(1)
	b := a.UnwrapOr(0)
	fmt.Println(b)
	// Output:
	// 1
}

func ExampleOption_UnwrapOr_false() {
	a := option.None[int]()
	b := a.UnwrapOr(99)
	fmt.Println(b)
	// Output:
	// 99
}

func Example_defaultParameter() {
	add := func(a, b *option.Option[int]) int {
		return a.UnwrapOr(0) + b.UnwrapOr(0)
	}
	r := add(option.Some(4), option.None[int]())
	fmt.Println(r)
	// Output:
	// 4
}

func Example_iAmEven() {
	a := option.Some(4).AndThen(func(i int) (int, bool) {
		if i%2 == 0 {
			return i, true
		}
		return 0, false
	})
	fmt.Println(a)
	// Output:
	// 4
}

func Example_iAmOdd() {
	a := option.Some(5).AndThen(func(i int) (int, bool) {
		if i%2 == 0 {
			return 0, false
		}
		return i, true
	})
	fmt.Println(a)
	// Output:
	// 0
}
