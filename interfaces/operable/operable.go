// operable
package operable

// Add is a type that can be added and subtracted to another type.
type Add[T any] interface {
	Add(T) T
	Sub(T) T
}

// Multiply is a type that can be multiplied with another type.
type Multiply[T any] interface {
	Mul(T) T
	Div(T) T
	Mod(T) T
}

// Power is a type that can be raised to a power.
type Power[T any] interface {
	Pow(T) T
	Root(T) T
}

// Biswise operations without shift
// ~, &, |, ^
type Bitwise[T any] interface {
	Not() T
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
