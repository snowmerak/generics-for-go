package math

import (
	"math"

	"golang.org/x/exp/constraints"
)

// PascalTriangle returns the Pascal's triangle for the given n.
func PascalsTriangle(n int) []int {
	l := []int{1}
	for i := 1; i < n; i++ {
		l = append(l, 1)
		for j := len(l) - 2; j > 0; j-- {
			l[j] += l[j-1]
		}
	}
	return l
}

// PowInt returns x^y of int.
func PowInt[T constraints.Integer](x, y T) T {
	z := T(1)
	for y > 0 {
		if y&1 == 1 {
			z = z * x
		}
		y = y >> 1
		x = x * x
	}
	return z
}

// PowFloat returns x^y of float.
func PowFloat[T constraints.Float](x, y T) T {
	return T(math.Pow(float64(x), float64(y)))
}

// Divide returns x/y.
func Divide[T constraints.Integer | constraints.Float](x, y T) T {
	return x / y
}

// DevRem returns the division and remainder of x/y.
func DevRem[T constraints.Integer](x, y T) (T, T) {
	d := T(0)
	r := T(0)
	for x >= y {
		d++
		r = x - y
		x = r
	}
	return d, r
}

// PerfectNumber returns perfect numbers to n.
func PerfectNumber[T constraints.Integer](n T) T {
	return PowInt(T(2), n-1) * (PowInt(T(2), n) - 1)
}

// Fibonacci returns the nth Fibonacci number.
func Fibonacci[T constraints.Integer](n T) T {
	return T((1 / math.Sqrt(5)) * (PowFloat((1.0+math.Sqrt(5))/2, float64(n)) - PowFloat((1.0-math.Sqrt(5))/2, float64(n))))
}
