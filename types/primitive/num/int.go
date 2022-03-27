// num package implements the interfaces from the operable package to integers.
//
// Int and Uint type is a wrapper for int and uint types.
// Int and Uint type implements Comparable, Clonable, Add, Multiply, Power, Bitwise, BitShift.
//
// Float type is a wrapper for float32, float64 type.
// Float type implements Comparable, Clonable, Add, Multiply, Power.
//
package num

import (
	"fmt"
	"strconv"

	"github.com/snowmerak/generics-for-go/v2/interfaces/comparable"
	"github.com/snowmerak/generics-for-go/v2/interfaces/copyable"
	"github.com/snowmerak/generics-for-go/v2/interfaces/operable"
	"github.com/snowmerak/generics-for-go/v2/math"
)

type Int int
type Int8 int8
type Int16 int16
type Int32 int32
type Int64 int64
type Uint uint
type Uint8 uint8
type Uint16 uint16
type Uint32 uint32
type Uint64 uint64

var _ fmt.Stringer = Int(0)
var _ fmt.Stringer = Int8(0)
var _ fmt.Stringer = Int16(0)
var _ fmt.Stringer = Int32(0)
var _ fmt.Stringer = Int64(0)
var _ fmt.Stringer = Uint(0)
var _ fmt.Stringer = Uint8(0)
var _ fmt.Stringer = Uint16(0)
var _ fmt.Stringer = Uint32(0)
var _ fmt.Stringer = Uint64(0)

func (i Int) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int8) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int16) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int32) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (i Int64) String() string {
	return strconv.FormatInt(int64(i), 10)
}

func (u Uint) String() string {
	return strconv.FormatUint(uint64(u), 10)
}

func (u Uint8) String() string {
	return strconv.FormatUint(uint64(u), 10)
}

func (u Uint16) String() string {
	return strconv.FormatUint(uint64(u), 10)
}

func (u Uint32) String() string {
	return strconv.FormatUint(uint64(u), 10)
}

func (u Uint64) String() string {
	return strconv.FormatUint(uint64(u), 10)
}

var _ comparable.Comparable[Int] = Int(0)
var _ comparable.Comparable[Int8] = Int8(0)
var _ comparable.Comparable[Int16] = Int16(0)
var _ comparable.Comparable[Int32] = Int32(0)
var _ comparable.Comparable[Int64] = Int64(0)
var _ comparable.Comparable[Uint] = Uint(0)
var _ comparable.Comparable[Uint8] = Uint8(0)
var _ comparable.Comparable[Uint16] = Uint16(0)
var _ comparable.Comparable[Uint32] = Uint32(0)
var _ comparable.Comparable[Uint64] = Uint64(0)

func (i Int) CompareTo(other Int) comparable.Compared {
	if i < other {
		return comparable.Less
	}
	if i > other {
		return comparable.Greater
	}
	return comparable.Equal
}

func (i Int8) CompareTo(other Int8) comparable.Compared {
	if i < other {
		return comparable.Less
	}
	if i > other {
		return comparable.Greater
	}
	return comparable.Equal
}

func (i Int16) CompareTo(other Int16) comparable.Compared {
	if i < other {
		return comparable.Less
	}
	if i > other {
		return comparable.Greater
	}
	return comparable.Equal
}

func (i Int32) CompareTo(other Int32) comparable.Compared {
	if i < other {
		return comparable.Less
	}
	if i > other {
		return comparable.Greater
	}
	return comparable.Equal
}

func (i Int64) CompareTo(other Int64) comparable.Compared {
	if i < other {
		return comparable.Less
	}
	if i > other {
		return comparable.Greater
	}
	return comparable.Equal
}

func (u Uint) CompareTo(other Uint) comparable.Compared {
	if u < other {
		return comparable.Less
	}
	if u > other {
		return comparable.Greater
	}
	return comparable.Equal
}

func (u Uint8) CompareTo(other Uint8) comparable.Compared {
	if u < other {
		return comparable.Less
	}
	if u > other {
		return comparable.Greater
	}
	return comparable.Equal
}

func (u Uint16) CompareTo(other Uint16) comparable.Compared {
	if u < other {
		return comparable.Less
	}
	if u > other {
		return comparable.Greater
	}
	return comparable.Equal
}

func (u Uint32) CompareTo(other Uint32) comparable.Compared {
	if u < other {
		return comparable.Less
	}
	if u > other {
		return comparable.Greater
	}
	return comparable.Equal
}

func (u Uint64) CompareTo(other Uint64) comparable.Compared {
	if u < other {
		return comparable.Less
	}
	if u > other {
		return comparable.Greater
	}
	return comparable.Equal
}

var _ copyable.Clonable[Int] = Int(0)
var _ copyable.Clonable[Int8] = Int8(0)
var _ copyable.Clonable[Int16] = Int16(0)
var _ copyable.Clonable[Int32] = Int32(0)
var _ copyable.Clonable[Int64] = Int64(0)
var _ copyable.Clonable[Uint] = Uint(0)
var _ copyable.Clonable[Uint8] = Uint8(0)
var _ copyable.Clonable[Uint16] = Uint16(0)
var _ copyable.Clonable[Uint32] = Uint32(0)
var _ copyable.Clonable[Uint64] = Uint64(0)

func (i Int) Clone() Int {
	return i
}

func (i Int8) Clone() Int8 {
	return i
}

func (i Int16) Clone() Int16 {
	return i
}

func (i Int32) Clone() Int32 {
	return i
}

func (i Int64) Clone() Int64 {
	return i
}

func (u Uint) Clone() Uint {
	return u
}

func (u Uint8) Clone() Uint8 {
	return u
}

func (u Uint16) Clone() Uint16 {
	return u
}

func (u Uint32) Clone() Uint32 {
	return u
}

func (u Uint64) Clone() Uint64 {
	return u
}

var _ operable.Add[Int] = Int(0)
var _ operable.Add[Int8] = Int8(0)
var _ operable.Add[Int16] = Int16(0)
var _ operable.Add[Int32] = Int32(0)
var _ operable.Add[Int64] = Int64(0)
var _ operable.Add[Uint] = Uint(0)
var _ operable.Add[Uint8] = Uint8(0)
var _ operable.Add[Uint16] = Uint16(0)
var _ operable.Add[Uint32] = Uint32(0)
var _ operable.Add[Uint64] = Uint64(0)

func (i Int) Add(other Int) Int {
	return i + other
}

func (i Int) Sub(other Int) Int {
	return i - other
}

func (i Int8) Add(other Int8) Int8 {
	return i + other
}

func (i Int8) Sub(other Int8) Int8 {
	return i - other
}

func (i Int16) Add(other Int16) Int16 {
	return i + other
}

func (i Int16) Sub(other Int16) Int16 {
	return i - other
}

func (i Int32) Add(other Int32) Int32 {
	return i + other
}

func (i Int32) Sub(other Int32) Int32 {
	return i - other
}

func (i Int64) Add(other Int64) Int64 {
	return i + other
}

func (i Int64) Sub(other Int64) Int64 {
	return i - other
}

func (u Uint) Add(other Uint) Uint {
	return u + other
}

func (u Uint) Sub(other Uint) Uint {
	return u - other
}

func (u Uint8) Add(other Uint8) Uint8 {
	return u + other
}

func (u Uint8) Sub(other Uint8) Uint8 {
	return u - other
}

func (u Uint16) Add(other Uint16) Uint16 {
	return u + other
}

func (u Uint16) Sub(other Uint16) Uint16 {
	return u - other
}

func (u Uint32) Add(other Uint32) Uint32 {
	return u + other
}

func (u Uint32) Sub(other Uint32) Uint32 {
	return u - other
}

func (u Uint64) Add(other Uint64) Uint64 {
	return u + other
}

func (u Uint64) Sub(other Uint64) Uint64 {
	return u - other
}

var _ operable.Multiply[Int] = Int(0)
var _ operable.Multiply[Int8] = Int8(0)
var _ operable.Multiply[Int16] = Int16(0)
var _ operable.Multiply[Int32] = Int32(0)
var _ operable.Multiply[Int64] = Int64(0)
var _ operable.Multiply[Uint] = Uint(0)
var _ operable.Multiply[Uint8] = Uint8(0)
var _ operable.Multiply[Uint16] = Uint16(0)
var _ operable.Multiply[Uint32] = Uint32(0)
var _ operable.Multiply[Uint64] = Uint64(0)

func (i Int) Mul(other Int) Int {
	return i * other
}

func (i Int) Div(other Int) Int {
	return i / other
}

func (i Int) Mod(other Int) Int {
	return i % other
}

func (i Int8) Mul(other Int8) Int8 {
	return i * other
}

func (i Int8) Div(other Int8) Int8 {
	return i / other
}

func (i Int8) Mod(other Int8) Int8 {
	return i % other
}

func (i Int16) Mul(other Int16) Int16 {
	return i * other
}

func (i Int16) Div(other Int16) Int16 {
	return i / other
}

func (i Int16) Mod(other Int16) Int16 {
	return i % other
}

func (i Int32) Mul(other Int32) Int32 {
	return i * other
}

func (i Int32) Div(other Int32) Int32 {
	return i / other
}

func (i Int32) Mod(other Int32) Int32 {
	return i % other
}

func (i Int64) Mul(other Int64) Int64 {
	return i * other
}

func (i Int64) Div(other Int64) Int64 {
	return i / other
}

func (i Int64) Mod(other Int64) Int64 {
	return i % other
}

func (u Uint) Mul(other Uint) Uint {
	return u * other
}

func (u Uint) Div(other Uint) Uint {
	return u / other
}

func (u Uint) Mod(other Uint) Uint {
	return u % other
}

func (u Uint8) Mul(other Uint8) Uint8 {
	return u * other
}

func (u Uint8) Div(other Uint8) Uint8 {
	return u / other
}

func (u Uint8) Mod(other Uint8) Uint8 {
	return u % other
}

func (u Uint16) Mul(other Uint16) Uint16 {
	return u * other
}

func (u Uint16) Div(other Uint16) Uint16 {
	return u / other
}

func (u Uint16) Mod(other Uint16) Uint16 {
	return u % other
}

func (u Uint32) Mul(other Uint32) Uint32 {
	return u * other
}

func (u Uint32) Div(other Uint32) Uint32 {
	return u / other
}

func (u Uint32) Mod(other Uint32) Uint32 {
	return u % other
}

func (u Uint64) Mul(other Uint64) Uint64 {
	return u * other
}

func (u Uint64) Div(other Uint64) Uint64 {
	return u / other
}

func (u Uint64) Mod(other Uint64) Uint64 {
	return u % other
}

var _ operable.Power[Int] = Int(0)
var _ operable.Power[Int8] = Int8(0)
var _ operable.Power[Int16] = Int16(0)
var _ operable.Power[Int32] = Int32(0)
var _ operable.Power[Int64] = Int64(0)
var _ operable.Power[Uint] = Uint(0)
var _ operable.Power[Uint8] = Uint8(0)
var _ operable.Power[Uint16] = Uint16(0)
var _ operable.Power[Uint32] = Uint32(0)
var _ operable.Power[Uint64] = Uint64(0)

func (i Int) Pow(other Int) Int {
	return math.PowInt(i, other)
}

func (i Int) Root(other Int) Int {
	return math.Root(i, other)
}

func (i Int8) Pow(other Int8) Int8 {
	return math.PowInt(i, other)
}

func (i Int8) Root(other Int8) Int8 {
	return math.Root(i, other)
}

func (i Int16) Pow(other Int16) Int16 {
	return math.PowInt(i, other)
}

func (i Int16) Root(other Int16) Int16 {
	return math.Root(i, other)
}

func (i Int32) Pow(other Int32) Int32 {
	return math.PowInt(i, other)
}

func (i Int32) Root(other Int32) Int32 {
	return math.Root(i, other)
}

func (i Int64) Pow(other Int64) Int64 {
	return math.PowInt(i, other)
}

func (i Int64) Root(other Int64) Int64 {
	return math.Root(i, other)
}

func (u Uint) Pow(other Uint) Uint {
	return math.PowInt(u, other)
}

func (u Uint) Root(other Uint) Uint {
	return math.Root(u, other)
}

func (u Uint8) Pow(other Uint8) Uint8 {
	return math.PowInt(u, other)
}

func (u Uint8) Root(other Uint8) Uint8 {
	return math.Root(u, other)
}

func (u Uint16) Pow(other Uint16) Uint16 {
	return math.PowInt(u, other)
}

func (u Uint16) Root(other Uint16) Uint16 {
	return math.Root(u, other)
}

func (u Uint32) Pow(other Uint32) Uint32 {
	return math.PowInt(u, other)
}

func (u Uint32) Root(other Uint32) Uint32 {
	return math.Root(u, other)
}

func (u Uint64) Pow(other Uint64) Uint64 {
	return math.PowInt(u, other)
}

func (u Uint64) Root(other Uint64) Uint64 {
	return math.Root(u, other)
}

var _ operable.Bitwise[Int] = Int(0)
var _ operable.Bitwise[Int8] = Int8(0)
var _ operable.Bitwise[Int16] = Int16(0)
var _ operable.Bitwise[Int32] = Int32(0)
var _ operable.Bitwise[Int64] = Int64(0)
var _ operable.Bitwise[Uint] = Uint(0)
var _ operable.Bitwise[Uint8] = Uint8(0)
var _ operable.Bitwise[Uint16] = Uint16(0)
var _ operable.Bitwise[Uint32] = Uint32(0)
var _ operable.Bitwise[Uint64] = Uint64(0)

func (i Int) And(other Int) Int {
	return i & other
}

func (i Int) Or(other Int) Int {
	return i | other
}

func (i Int) Xor(other Int) Int {
	return i ^ other
}

func (i Int) Not() Int {
	return ^i
}

func (i Int8) And(other Int8) Int8 {
	return i & other
}

func (i Int8) Or(other Int8) Int8 {
	return i | other
}

func (i Int8) Xor(other Int8) Int8 {
	return i ^ other
}

func (i Int8) Not() Int8 {
	return ^i
}

func (i Int16) And(other Int16) Int16 {
	return i & other
}

func (i Int16) Or(other Int16) Int16 {
	return i | other
}

func (i Int16) Xor(other Int16) Int16 {
	return i ^ other
}

func (i Int16) Not() Int16 {
	return ^i
}

func (i Int32) And(other Int32) Int32 {
	return i & other
}

func (i Int32) Or(other Int32) Int32 {
	return i | other
}

func (i Int32) Xor(other Int32) Int32 {
	return i ^ other
}

func (i Int32) Not() Int32 {
	return ^i
}

func (i Int64) And(other Int64) Int64 {
	return i & other
}

func (i Int64) Or(other Int64) Int64 {
	return i | other
}

func (i Int64) Xor(other Int64) Int64 {
	return i ^ other
}

func (i Int64) Not() Int64 {
	return ^i
}

func (u Uint) And(other Uint) Uint {
	return u & other
}

func (u Uint) Or(other Uint) Uint {
	return u | other
}

func (u Uint) Xor(other Uint) Uint {
	return u ^ other
}

func (u Uint) Not() Uint {
	return ^u
}

func (u Uint8) And(other Uint8) Uint8 {
	return u & other
}

func (u Uint8) Or(other Uint8) Uint8 {
	return u | other
}

func (u Uint8) Xor(other Uint8) Uint8 {
	return u ^ other
}

func (u Uint8) Not() Uint8 {
	return ^u
}

func (u Uint16) And(other Uint16) Uint16 {
	return u & other
}

func (u Uint16) Or(other Uint16) Uint16 {
	return u | other
}

func (u Uint16) Xor(other Uint16) Uint16 {
	return u ^ other
}

func (u Uint16) Not() Uint16 {
	return ^u
}

func (u Uint32) And(other Uint32) Uint32 {
	return u & other
}

func (u Uint32) Or(other Uint32) Uint32 {
	return u | other
}

func (u Uint32) Xor(other Uint32) Uint32 {
	return u ^ other
}

func (u Uint32) Not() Uint32 {
	return ^u
}

func (u Uint64) And(other Uint64) Uint64 {
	return u & other
}

func (u Uint64) Or(other Uint64) Uint64 {
	return u | other
}

func (u Uint64) Xor(other Uint64) Uint64 {
	return u ^ other
}

func (u Uint64) Not() Uint64 {
	return ^u
}

var _ operable.BitShift[Int] = Int(0)
var _ operable.BitShift[Int8] = Int8(0)
var _ operable.BitShift[Int16] = Int16(0)
var _ operable.BitShift[Int32] = Int32(0)
var _ operable.BitShift[Int64] = Int64(0)
var _ operable.BitShift[Uint] = Uint(0)
var _ operable.BitShift[Uint8] = Uint8(0)
var _ operable.BitShift[Uint16] = Uint16(0)
var _ operable.BitShift[Uint32] = Uint32(0)
var _ operable.BitShift[Uint64] = Uint64(0)

func (i Int) LeftShift(other Int) Int {
	return i << other
}

func (i Int) LogicalShift(other Int) Int {
	return Int(uint(i) >> other)
}

func (i Int) ArithmeticShift(other Int) Int {
	return i >> other
}

func (i Int8) LeftShift(other Int8) Int8 {
	return i << other
}

func (i Int8) LogicalShift(other Int8) Int8 {
	return Int8(uint8(i) >> other)
}

func (i Int8) ArithmeticShift(other Int8) Int8 {
	return i >> other
}

func (i Int16) LeftShift(other Int16) Int16 {
	return i << other
}

func (i Int16) LogicalShift(other Int16) Int16 {
	return Int16(uint16(i) >> other)
}

func (i Int16) ArithmeticShift(other Int16) Int16 {
	return i >> other
}

func (i Int32) LeftShift(other Int32) Int32 {
	return i << other
}

func (i Int32) LogicalShift(other Int32) Int32 {
	return Int32(uint32(i) >> other)
}

func (i Int32) ArithmeticShift(other Int32) Int32 {
	return i >> other
}

func (i Int64) LeftShift(other Int64) Int64 {
	return i << other
}

func (i Int64) LogicalShift(other Int64) Int64 {
	return Int64(uint64(i) >> other)
}

func (i Int64) ArithmeticShift(other Int64) Int64 {
	return i >> other
}

func (u Uint) LeftShift(other Uint) Uint {
	return u << other
}

func (u Uint) LogicalShift(other Uint) Uint {
	return u >> other
}

func (u Uint) ArithmeticShift(other Uint) Uint {
	return u >> other
}

func (u Uint8) LeftShift(other Uint8) Uint8 {
	return u << other
}

func (u Uint8) LogicalShift(other Uint8) Uint8 {
	return u >> other
}

func (u Uint8) ArithmeticShift(other Uint8) Uint8 {
	return u >> other
}

func (u Uint16) LeftShift(other Uint16) Uint16 {
	return u << other
}

func (u Uint16) LogicalShift(other Uint16) Uint16 {
	return u >> other
}

func (u Uint16) ArithmeticShift(other Uint16) Uint16 {
	return u >> other
}

func (u Uint32) LeftShift(other Uint32) Uint32 {
	return u << other
}

func (u Uint32) LogicalShift(other Uint32) Uint32 {
	return u >> other
}

func (u Uint32) ArithmeticShift(other Uint32) Uint32 {
	return u >> other
}

func (u Uint64) LeftShift(other Uint64) Uint64 {
	return u << other
}

func (u Uint64) LogicalShift(other Uint64) Uint64 {
	return u >> other
}

func (u Uint64) ArithmeticShift(other Uint64) Uint64 {
	return u >> other
}
