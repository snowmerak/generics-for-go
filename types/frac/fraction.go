package frac

import "github.com/snowmerak/generics-for-go/v2/interfaces/comparable"

type Fraction[T uint8 | uint16 | uint32 | uint64] struct {
	numurator   T
	denominator T
	sign        bool
}

var _ comparable.Comparable[Fraction[uint8]] = Fraction[uint8]{}

func (f Fraction[T]) CompareTo(other Fraction[T]) comparable.Compared {
	if f.sign && !other.sign {
		return comparable.Greater
	}
	if !f.sign && other.sign {
		return comparable.Less
	}
	if f.sign && other.sign {
		if float64(f.numurator)/float64(f.denominator) > float64(other.numurator)/float64(other.denominator) {
			return comparable.Greater
		}
		if float64(f.numurator)/float64(f.denominator) < float64(other.numurator)/float64(other.denominator) {
			return comparable.Less
		}
	}
	if !f.sign && !other.sign {
		if float64(f.numurator)/float64(f.denominator) > float64(other.numurator)/float64(other.denominator) {
			return comparable.Less
		}
		if float64(f.numurator)/float64(f.denominator) < float64(other.numurator)/float64(other.denominator) {
			return comparable.Greater
		}
	}
	return comparable.Equal
}
