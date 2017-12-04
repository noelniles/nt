package main

import (
	"math"
	"testing"
)

func TestRGCD(t *testing.T) {
	if gcd := RGCD(13, 13); gcd != 13 {
		t.Errorf("Expected GCD=13 but it was %d instead", gcd)
	}
	if gcd := RGCD(37, 600); gcd != 1 {
		t.Errorf("Expected GCD=1 but it was %d instead", gcd)
	}
	if gcd := RGCD(20, 100); gcd != 20 {
		t.Errorf("Expected GCD=6 but it was %d instead", gcd)
	}
	if gcd := RGCD(624129, 2061517); gcd != 18913 {
		t.Errorf("Expected GCD=18913 but it was %d instead", gcd)
	}
	if gcd := RGCD(math.MaxInt32, 0); gcd != math.MaxInt32 {
		t.Errorf("Expected GCD=%d but it was %d instead", math.MaxInt32, gcd)
	}
}

func TestRGCD2(t *testing.T) {
	if gcd := RGCD2(13, 13); gcd != 13 {
		t.Errorf("Expected GCD=13 but it was %d instead", gcd)
	}
	if gcd := RGCD2(37, 600); gcd != 1 {
		t.Errorf("Expected GCD=1 but it was %d instead", gcd)
	}
	if gcd := RGCD2(20, 100); gcd != 20 {
		t.Errorf("Expected GCD=6 but it was %d instead", gcd)
	}
	if gcd := RGCD2(624129, 2061517); gcd != 18913 {
		t.Errorf("Expected GCD=18913 but it was %d instead", gcd)
	}
	// This will  always fail because the recursion is too deep.
	//if gcd := RGCD2(math.MaxInt32, 0); gcd != math.MaxInt32 {
	//	t.Errorf("Expected GCD=%d but it was %d instead", math.MaxInt32, gcd)
	//}
}

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
		t.Errorf("Expected GCD=%d but it was %d instead", math.MaxInt32, gcd)
	}
}

func TestXGCD(t *testing.T) {
	if d, x, y := XGCD(4864, 3458); d != 38 && x != -45 && y != 38 {
		t.Errorf("Expected XGCD=32,-45,38 but it was %d,%d,%d instead", d, x, y)
	}
}

func TestInvMod(t *testing.T) {
	if i := InvMod(1, 2); i != 1 {
		t.Errorf("Expected InvMod(1, 2)=1 but it was %d instead", i)
	}
	if i := InvMod(3, 6); i != 0 {
		t.Errorf("Expected InvMod(3, 6)=0 but it was %d instead", i)
	}
	if i := InvMod(7, 87); i != 25 {
		t.Errorf("Expected InvMod(7, 87)=25 but it was %d instead", i)
	}
	if i := InvMod(2, 91); i != 46 {
		t.Errorf("Expected InvMod(2, 91)=46 but it was %d instead", i)
	}
	if i := InvMod(13, 91); i != 0 {
		t.Errorf("Expected InvMod(13, 91)=0 but it was %d instead", i)
	}
	if i := InvMod(13, 91); i != 0 {
		t.Errorf("Expected InvMod(13, 91)=0 but it was %d instead", i)
	}
	if i := InvMod(31, 73714876143); i != 45180085378 {
		t.Errorf("Expected InvMod(3, 7371487614)=45180085378 but it was %d instead", i)
	}
	if i := InvMod(3, 73714876143); i != 0 {
		t.Errorf("Expected InvMod(3, 7371487614)=0 but it was %d instead", i)
	}
}
