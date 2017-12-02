package main

import (
	"math"
	"testing"
)

func TestGCD(t *testing.T) {
	if gcd := GCD(13, 13); gcd != 13 {
		t.Errorf("Expected GCD=13 but it was %d instead", gcd)
	}
	if gcd := GCD(37, 600); gcd != 1 {
		t.Errorf("Expected GCD=1 but it was %d instead", gcd)
	}
	if gcd := GCD(20, 100); gcd != 20 {
		t.Errorf("Expected GCD=6 but it was %d instead", gcd)
	}
	if gcd := GCD(624129, 2061517); gcd != 18913 {
		t.Errorf("Expected GCD=18913 but it was %d instead", gcd)
	}
	if gcd := GCD(math.MaxInt32, 0); gcd != math.MaxInt32 {
		t.Errorf("Expected GCD=0 but it was %d instead", gcd)
	}
}

func TestXGCD(t *testing.T) {
	if d, x, y := XGCD(4864, 3458); d != 38 && x != -45 && y != 38 {
		t.Errorf("Expected XGCD=32,-45,38 but it was %d,%d,%d instead", d, x, y)
	}

}
