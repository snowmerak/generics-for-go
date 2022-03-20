package slice_test

import (
	"fmt"

	"github.com/snowmerak/generics-for-go/v2/collections/slice"
)

func Example_allAny() {
	s := slice.Of(1, 2, 3, 4, 5)
	b := s.All(func(i int) bool {
		return i <= 5
	})
	fmt.Println(b)
	b = s.Any(func(i int) bool {
		return i == 3
	})
	fmt.Println(b)
	// Output:
	// true
	// true
}

func Example_binarySearch() {
	s := slice.Of(1, 3, 5, 7, 9, 11, 13, 15, 17, 19)
	found := s.BinarySearch(3, func(i1, i2 int) int {
		if i1 < i2 {
			return -1
		} else if i1 > i2 {
			return 1
		}
		return 0
	})
	fmt.Println(found)
	notFound := s.BinarySearch(8, func(i1, i2 int) int {
		if i1 < i2 {
			return -1
		} else if i1 > i2 {
			return 1
		}
		return 0
	})
	fmt.Println(notFound)
	// Output:
	// 1
	// -1
}

func Example_chunk() {
	s := slice.Of(1, 3, 5, 7, 9, 11, 13, 15, 17, 19)
	result := s.Chunk(2)
	slice.Of(result...).Foreach(func(s slice.Slice[int]) {
		fmt.Println(s.JoinToString(", "))
	})
	// Output:
	// 1, 3
	// 5, 7
	// 9, 11
	// 13, 15
	// 17, 19
}

func Example_count() {
	s := slice.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	evens := s.Count(func(i int) bool {
		return i%2 == 0
	})
	fmt.Println(evens)
	// Output:
	// 5
}

func Example_filter() {
	s := slice.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	r := s.Filter(func(i int) bool {
		return i > 5
	}).JoinToString(", ")
	fmt.Println(r)
	// Output:
	// 6, 7, 8, 9, 10, 11
}

func Example_firstOf() {
	s := slice.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	r := s.FirstOf(func(i int) bool {
		return i > 5
	})
	fmt.Println(r)
	// Output:
	// 6
}

func Example_map() {
	s := slice.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	r := s.Map(func(index, value int) int {
		return value * 2
	})
	fmt.Println(r.JoinToString(", "))
	// Output:
	// 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22
}

func Example_max() {
	s := slice.Of(1, 11, 2, 10, 3, 9, 4, 8, 5, 7, 6)
	r := s.Max(func(i1, i2 int) int {
		if i1 > i2 {
			return -1
		} else if i1 < i2 {
			return 1
		}
		return 0
	})
	fmt.Println(r)
	// Output:
	// 11
}

func Example_min() {
	s := slice.Of(1, 11, 2, 10, 3, 9, 4, 8, 5, 7, 6)
	r := s.Min(func(i1, i2 int) int {
		if i1 > i2 {
			return -1
		} else if i1 < i2 {
			return 1
		}
		return 0
	})
	fmt.Println(r)
	// Output:
	// 1
}

func Example_random() {
	s := slice.Of(1, 11, 2, 10, 3, 9, 4, 8, 5, 7, 6)
	r := s.Random()
	fmt.Println(r)
	// Output:
	// random value of 1, 11, 2, 10, 3, 9, 4, 8, 5, 7, 6
}

func Example_reduce() {
	s := slice.Of(1, 11, 2, 10, 3, 9, 4, 8, 5, 7, 6)
	sum := s.Reduce(func(i1, i2 int) int {
		return i1 + i2
	}, 0)
	fmt.Println(sum)
	// Output:
	// 66
}

func Example_sort() {
	s := slice.Of(1, 11, 2, 10, 3, 9, 4, 8, 5, 7, 6)
	r := s.Sort(func(i1, i2 int) int {
		if i1 < i2 {
			return -1
		} else if i1 > i2 {
			return 1
		}
		return 0
	})
	fmt.Println(r.JoinToString(", "))
	// Output:
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11
}
