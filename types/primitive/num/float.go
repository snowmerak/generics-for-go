package num

import (
	"fmt"
	"strconv"

	"github.com/snowmerak/generics-for-go/v2/interfaces/comparable"
	"github.com/snowmerak/generics-for-go/v2/interfaces/operable"
	"github.com/snowmerak/generics-for-go/v2/math"
)

type Float32 float32
type Float64 float64

var _ fmt.Stringer = Float32(0)
var _ fmt.Stringer = Float32(0)

func (f Float32) String() string {
	return strconv.FormatFloat(float64(f), 'f', -1, 32)
}

func (f Float64) String() string {
	return strconv.FormatFloat(float64(f), 'f', -1, 64)
}

var _ comparable.Comparable[Float32] = Float32(0)
var _ comparable.Comparable[Float64] = Float64(0)

func (f Float32) CompareTo(other Float32) comparable.Compared {
	if f < other {
		return comparable.Less
	}
	if f > other {
		return comparable.Greater
	}
	return comparable.Equal
}

func (f Float64) CompareTo(other Float64) comparable.Compared {
	if f < other {
		return comparable.Less
	}
	if f > other {
		return comparable.Greater
	}
	return comparable.Equal
}

var _ operable.Add[Float32] = Float32(0)
var _ operable.Add[Float64] = Float64(0)

func (f Float32) Add(other Float32) Float32 {
	return f + other
}

func (f Float32) Sub(other Float32) Float32 {
	return f - other
}

func (f Float64) Add(other Float64) Float64 {
	return f + other
}

func (f Float64) Sub(other Float64) Float64 {
	return f - other
}

var _ operable.Multiply[Float32] = Float32(0)
var _ operable.Multiply[Float64] = Float64(0)

func (f Float32) Mul(other Float32) Float32 {
	return f * other
}

func (f Float32) Div(other Float32) Float32 {
	return f / other
}

func (f Float32) Mod(other Float32) Float32 {
	return 0
}

func (f Float64) Mul(other Float64) Float64 {
	return f * other
}

func (f Float64) Div(other Float64) Float64 {
	return f / other
}

func (f Float64) Mod(other Float64) Float64 {
	return 0
}

var _ operable.Power[Float32] = Float32(0)
var _ operable.Power[Float64] = Float64(0)

func (f Float32) Pow(other Float32) Float32 {
	return math.PowFloat(f, other)
}

func (f Float32) Root(other Float32) Float32 {
	return math.Root(f, other)
}

func (f Float64) Pow(other Float64) Float64 {
	return math.PowFloat(f, other)
}

func (f Float64) Root(other Float64) Float64 {
	return math.Root(f, other)
}
