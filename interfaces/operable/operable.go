// operable
package operable

// Arithmetic operations
// +, -, *, /, %, ^, sqrt
type Arithmetic[T any] interface {
	Add(T) T
	Sub(T) T
	Mul(T) T
	Div(T) T
	Mod(T) T
	Pow(T) T
	Sqrt(T) T
}

// Biswise operations without shift
// ~, &, |, ^
type Bitwise[T any] interface {
	Not(T) T
	And(T) T
	Or(T) T
	Xor(T) T
}

// Bitwise shift operations
// <<, >>, >>>
type Shift[T any] interface {
	LeftShift(T) T
	LogicalShift(T) T
	ArithmeticShift(T) T
}
