package assert

import (
	"fmt"
	"math"
	"reflect"

	"golang.org/x/exp/constraints"
)

// Positive asserts that the value is positive.
func Positive[T constraints.Signed | constraints.Float](a T) error {
	if a < 0 {
		return fmt.Errorf("%v is not positive", a)
	}
	return nil
}

// Nagative asserts that the value is negative.
func Nagative[T constraints.Signed | constraints.Float](a T) error {
	if a > 0 {
		return fmt.Errorf("%v is not negative", a)
	}
	return nil
}

// Between asserts that the value is between the expected values.
// The expected values are inclusive.
func Between[T constraints.Ordered](a T, min, max T) error {
	if a < min || a > max {
		return fmt.Errorf("%v is not between %v and %v", a, min, max)
	}
	return nil
}

// CloseTo asserts that the value is close to the expected value.
// The expected value is within epsilon of the actual value.
func CloseTo[T constraints.Integer | constraints.Float](a T, b T, epsilon T) error {
	if a-b > epsilon || b-a > epsilon {
		return fmt.Errorf("%v is not close to %v", a, b)
	}
	return nil
}

// NotCloseTo asserts that the value is not close to the expected value.
// The expected value is within epsilon of the actual value.
func NotCloseTo[T constraints.Integer | constraints.Float](a T, b T, epsilon T) error {
	if a-b < epsilon && b-a < epsilon {
		return fmt.Errorf("%v is close to %v", a, b)
	}
	return nil
}

// Max asserts that the value is the maximum value.
func Max[T constraints.Integer | constraints.Float](a T) error {
	switch reflect.TypeOf(a).Kind() {
	case reflect.Int:
		if int(a) != math.MaxInt {
			return fmt.Errorf("%v is not max int", a)
		}
	case reflect.Int8:
		if int8(a) != math.MaxInt8 {
			return fmt.Errorf("%v is not max int8", a)
		}
	case reflect.Int16:
		if int16(a) != math.MaxInt16 {
			return fmt.Errorf("%v is not max int16", a)
		}
	case reflect.Int32:
		if int32(a) != math.MaxInt32 {
			return fmt.Errorf("%v is not max int32", a)
		}
	case reflect.Int64:
		if int64(a) != math.MaxInt64 {
			return fmt.Errorf("%v is not max int64", a)
		}
	case reflect.Uint:
		if uint(a) != math.MaxUint {
			return fmt.Errorf("%v is not max uint", a)
		}
	case reflect.Uint8:
		if uint8(a) != math.MaxUint8 {
			return fmt.Errorf("%v is not max uint8", a)
		}
	case reflect.Uint16:
		if uint16(a) != math.MaxUint16 {
			return fmt.Errorf("%v is not max uint16", a)
		}
	case reflect.Uint32:
		if uint32(a) != math.MaxUint32 {
			return fmt.Errorf("%v is not max uint32", a)
		}
	case reflect.Uint64:
		if uint64(a) != math.MaxUint64 {
			return fmt.Errorf("%v is not max uint64", a)
		}
	case reflect.Float32:
		if float32(a) != math.MaxFloat32 {
			return fmt.Errorf("%v is not max float32", a)
		}
	case reflect.Float64:
		if float64(a) != math.MaxFloat64 {
			return fmt.Errorf("%v is not max float64", a)
		}
	}
	return nil
}

// Min asserts that the value is the minimum value.
func Min[T constraints.Integer | constraints.Float](a T) error {
	switch reflect.TypeOf(a).Kind() {
	case reflect.Int:
		if int(a) != math.MinInt {
			return fmt.Errorf("%v is not max int", a)
		}
	case reflect.Int8:
		if int8(a) != math.MinInt8 {
			return fmt.Errorf("%v is not max int8", a)
		}
	case reflect.Int16:
		if int16(a) != math.MinInt16 {
			return fmt.Errorf("%v is not max int16", a)
		}
	case reflect.Int32:
		if int32(a) != math.MinInt32 {
			return fmt.Errorf("%v is not max int32", a)
		}
	case reflect.Int64:
		if int64(a) != math.MinInt64 {
			return fmt.Errorf("%v is not max int64", a)
		}
	case reflect.Uint:
		if uint(a) != 0 {
			return fmt.Errorf("%v is not max uint", a)
		}
	case reflect.Uint8:
		if uint8(a) != 0 {
			return fmt.Errorf("%v is not max uint8", a)
		}
	case reflect.Uint16:
		if uint16(a) != 0 {
			return fmt.Errorf("%v is not max uint16", a)
		}
	case reflect.Uint32:
		if uint32(a) != 0 {
			return fmt.Errorf("%v is not max uint32", a)
		}
	case reflect.Uint64:
		if uint64(a) != 0 {
			return fmt.Errorf("%v is not max uint64", a)
		}
	case reflect.Float32:
		if float32(a) != -math.MaxFloat32 {
			return fmt.Errorf("%v is not max float32", a)
		}
	case reflect.Float64:
		if float64(a) != -math.MaxFloat64 {
			return fmt.Errorf("%v is not max float64", a)
		}
	}
	return nil
}
