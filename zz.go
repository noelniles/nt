// zz is the main package in NT (For now). It's just a bunch of functions over
// the integers. There will probably be other packages or modules or something
// that implement the real and complex stuff.
package main

import (
	"fmt"
	"math"
)

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

// XGCD is the extended Euclidean Algorithm.
func XGCD(a int32, b int32) (int32, int32, int32) {
	var u, v, u0, v0, u1, v1, u2, v2, q, r int32
	var aneg, bneg int32

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
