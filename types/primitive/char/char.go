// char package implements the wrapper of Byte, Rune, String type.
package char

import (
	"fmt"
	"strings"

	"github.com/snowmerak/generics-for-go/v2/collections/slice"
	"github.com/snowmerak/generics-for-go/v2/interfaces/comparable"
	"github.com/snowmerak/generics-for-go/v2/interfaces/copyable"
	"github.com/snowmerak/generics-for-go/v2/interfaces/operable"
	"github.com/snowmerak/generics-for-go/v2/types/primitive/num"
)

type Byte byte
type Rune rune
type String string

var _ fmt.Stringer = (Byte)(' ')
var _ fmt.Stringer = (Rune)(' ')
var _ fmt.Stringer = (String)(" ")

func (b Byte) String() string {
	return string(b)
}

func (r Rune) String() string {
	return string(r)
}

func (s String) String() string {
	return string(s)
}

var _ comparable.Comparable[Byte] = (Byte)(' ')
var _ comparable.Comparable[Rune] = (Rune)(' ')
var _ comparable.Comparable[String] = (String)(" ")

func (b Byte) CompareTo(other Byte) comparable.Compared {
	if b < other {
		return comparable.Less
	}
	if b > other {
		return comparable.Greater
	}
	return comparable.Equal
}

func (r Rune) CompareTo(other Rune) comparable.Compared {
	if r < other {
		return comparable.Less
	}
	if r > other {
		return comparable.Greater
	}
	return comparable.Equal
}

func (s String) CompareTo(other String) comparable.Compared {
	if s < other {
		return comparable.Less
	}
	if s > other {
		return comparable.Greater
	}
	return comparable.Equal
}

var _ copyable.Clonable[Byte] = (Byte)(' ')
var _ copyable.Clonable[Rune] = (Rune)(' ')
var _ copyable.Clonable[String] = (String)(" ")

func (b Byte) Clone() Byte {
	return b
}

func (r Rune) Clone() Rune {
	return r
}

func (s String) Clone() String {
	return s
}

var _ operable.Add[Byte] = (Byte)(' ')
var _ operable.Add[Rune] = (Rune)(' ')
var _ operable.Add[String] = (String)(" ")

func (b Byte) Add(other Byte) Byte {
	return b + other
}

func (b Byte) Sub(other Byte) Byte {
	return b - other
}

func (r Rune) Add(other Rune) Rune {
	return r + other
}

func (r Rune) Sub(other Rune) Rune {
	return r - other
}

func (s String) Add(other String) String {
	return s + other
}

func (s String) Sub(other String) String {
	return String(strings.Replace(string(s), string(other), "", 1))
}

var _ operable.Multiply[Byte] = (Byte)(' ')
var _ operable.Multiply[Rune] = (Rune)(' ')
var _ operable.Multiply[String] = (String)(" ")

func (b Byte) Mul(other Byte) Byte {
	return b * other
}

func (b Byte) Div(other Byte) Byte {
	return b / other
}

func (b Byte) Mod(other Byte) Byte {
	return b % other
}

func (r Rune) Mul(other Rune) Rune {
	return r * other
}

func (r Rune) Div(other Rune) Rune {
	return r / other
}

func (r Rune) Mod(other Rune) Rune {
	return r % other
}

func (s String) Mul(other String) String {
	return s + other
}

func (s String) Div(other String) String {
	return String(strings.Repeat(string(other), strings.Count(string(s), string(other))))
}

func (s String) Mod(other String) String {
	return String(strings.Replace(string(s), string(other), "", -1))
}

// ToUint8 returns the uint8 value of the byte.
func (b Byte) ToUint8() num.Uint8 {
	return num.Uint8(b)
}

// ToInt returns the int value of the rune.
func (r Rune) ToInt() num.Int {
	return num.Int(r)
}

// ToBytes returns the byte array of the string.
func (s String) ToBytes() slice.Slice[Byte] {
	return []Byte(s)
}

// ToRunes returns the rune array of the string.
func (s String) ToRunes() slice.Slice[Rune] {
	return []Rune(s)
}
