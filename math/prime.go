package math

import (
	"golang.org/x/exp/constraints"
)

// Gcd returns the greatest common divisor of a and b.
func Gcd[T constraints.Integer](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Lcm returns the least common multiple of a and b.
func Lcm[T constraints.Integer](a, b T) T {
	g := Gcd(a, b)
	return a * b / g
}

// powermod returns a^b mod c.
func powermod[T constraints.Integer](x, y, m T) T {
	z := T(1)
	for y > 0 {
		if y&1 == 1 {
			z = z * x % m
		}
		y = y >> 1
		x = x * x % m
	}
	return z
}

// MillerRabin returns true if n is probably prime.
func MillerRabin[T constraints.Integer](n T) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

loop:
	for _, c := range []T{2, 3, 5, 7, 11, 13, 17} {
		if c >= n {
			continue loop
		}
		d := n - 1
		for d%2 == 0 {
			if powermod(c, d, n) == n-1 {
				continue loop
			}
			d /= 2
		}
		t := powermod(c, d, n)
		if t == 1 || t == n-1 {
			continue loop
		}
		return false
	}
	return true
}

// IsPrime returns true if n is must have prime.
func IsPrime[T constraints.Integer](n T) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	l := []T{2}
	for i := T(3); i*i <= n; i += 2 {
		isPrime := true
		for _, v := range l {
			if i%v == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			l = append(l, i)
		}
	}
	for _, v := range l {
		if n%v == 0 {
			return false
		}
	}
	return true
}

// SieveOfSundaram returns the prime numbers from 2 to n.
func SieveOfSundaram[T constraints.Integer](n T) []T {
	m := n/2 - 1
	set := map[T]struct{}{}
	for i := T(1); i <= m; i++ {
		for j := i; j <= m; j++ {
			k := i + j + 2*i*j
			if k <= n {
				set[k] = struct{}{}
				continue
			}
			break
		}
	}
	l := make([]T, 0, int(n)-len(set)+1)
	for i := T(1); i <= m; i++ {
		if _, ok := set[i]; !ok {
			l = append(l, 2*i+1)
		}
	}
	return l
}

// PrimesTo returns the prime numbers from 2 to n.
func PrimesTo[T constraints.Integer](n T) []T {
	l := []T{2}
	for i := T(3); i <= n; i += 2 {
		isPrime := true
		for _, v := range l {
			if i%v == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			l = append(l, i)
		}
	}
	return l
}
