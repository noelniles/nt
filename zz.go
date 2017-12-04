package main

// zz is the package in NT (For now). It's just a bunch of functions over
// the integers. There will probably be other packages or modules or something
// that implement the real and complex stuff.

import (
	"fmt"
	"math"
)

// RGCD is the recurisive gcd algorithm. This shouldn't be used.
func RGCD(m int32, n int32) int32 {
	if n == 0 {
		return m
	}
	return RGCD(n, m%n)
}

// This is a slightly better GCD function, but it still
// shouldn't be used because of the recursion.
func RGCD2(a int32, b int32) int32 {
	if a == b {
		return a
	} else if a > b {
		return RGCD2(a-b, b)
	} else {
		return RGCD2(a, b-a)
	}
}

// GCD is the Greatest Common Denominator of a and b.
func GCD(a int32, b int32) int32 {
	var u int32
	var v int32
	var t int32
	var x int32

	if a < 0 && a < -math.MaxInt32 {
		fmt.Println("GCD: integer overflow")
		a = -a
	}
	if b < 0 && b < -math.MaxInt32 {
		fmt.Println("GCD: integer overflow")
		b = -b
	}
	if b == 0 {
		x = a
	} else {
		u = a
		v = b
		for v != 0 {
			t = u % v
			u = v
			v = t
		}
		x = u
	}
	return x
}

// XGCD is the extended Euclidean Algorithm. In addition to the GCD of a and b, XGCD computes the coefficients of
// Bézouts identity x, y such that ax + by = gcd(a, b).
func XGCD(a int64, b int64) (int64, int64, int64) {
	var u, v, u0, v0, u1, v1, u2, v2, q, r int64
	var aneg, bneg int64

	if a < 0 {
		if a < -math.MaxInt32 {
			fmt.Println("XGCD: integer overflow")
		}
		a = -a
		aneg = 1
	}

	if b < 0 {
		if b < -math.MaxInt32 {
			fmt.Println("XGCD: integer overflow")
		}
		b = -b
		bneg = 1
	}

	u1 = 1
	v1 = 0
	u2 = 0
	v2 = 1
	u = a
	v = b

	for v != 0 {
		q = u / v
		r = u % v
		u = v
		v = r
		u0 = u2
		v0 = v2
		u2 = u1 - q*u2
		v2 = v1 - q*v2
		u1 = u0
		v1 = v0
	}
	if aneg != 0 {
		u1 = -u1
	}
	if bneg != 0 {
		v1 = -v1
	}
	return u, u1, v1
}

// InvMod returns the multiplicative inverse of a mod n. Returns 0 if the inverse is undefined.
func InvMod(a int64, n int64) int64 {
	d, s, _ := XGCD(a, n)

	if d != 1 {
		fmt.Println("InvMod: inverse undefined")
		return 0
	}
	if s < 0 {
		return s + n
	} else {
		return s
	}
}

func main() {}
