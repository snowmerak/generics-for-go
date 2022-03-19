package sequence_test

import (
	"fmt"

	"github.com/snowmerak/generics-for-go/v2/iterable"
	"github.com/snowmerak/generics-for-go/v2/iterable/sequence"
)

func Example_simple() {
	seq := sequence.New(iterable.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	for i := 0; i < 10; i++ {
		seq.Take(1)
		fmt.Println(seq.FirstValue())
		seq.Drop(1)
	}
	// Output:
	// 1 true
	// 2 true
	// 3 true
	// 4 true
	// 5 true
	// 6 true
	// 7 true
	// 8 true
	// 9 true
	// 10 true
}

func Example_prime() {
	seq := sequence.New(iterable.NewGenerator(0, 1, func(_, v int) (int, int) {
		return 0, v + 2
	}))
	primes := []int{2}
	seq.Filter(func(_, v int) bool {
		return v%2 == 0
	})
	for i := 0; i < 100; i++ {
		value, _ := seq.Take(1).FirstValue()
		primes = append(primes, value)
		seq.Drop(1).Filter(func(_, v int) bool {
			t := value
			return v%t == 0
		})
	}
	fmt.Println(primes)
	// Output:
	// [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97 101 103 107 109 113 127 131 137 139 149 151 157 163 167 173 179 181 191 193 197 199 211 223 227 229 233 239 241 251 257 263 269 271 277 281 283 293 307 311 313 317 331 337 347 349 353 359 367 373 379 383 389 397 401 409 419 421 431 433 439 443 449 457 461 463 467 479 487 491 499 503 509 521 523 541 547]
}

func Example_time2() {
	seq := sequence.New(iterable.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	seq.Map(func(i1, i2 int) (int, int) {
		return i1 * 2, i2 * 2
	})
	for i := 0; i < 10; i++ {
		seq.Take(1)
		fmt.Println(seq.FirstValue())
		seq.Drop(1)
	}
	// Output:
	// 2 true
	// 4 true
	// 6 true
	// 8 true
	// 10 true
	// 12 true
	// 14 true
	// 16 true
	// 18 true
	// 20 true
}
