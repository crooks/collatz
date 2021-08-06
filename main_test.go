package main

import (
	"math/big"
	"testing"
)

func TestIsEven(t *testing.T) {
	x := new(big.Int)
	y := new(big.Int)
	x.SetString("1234567890", 10)
	if ! isEven(x) {
		t.Fatalf("Expected true response from isEven")
	}
	x.SetString("12345678901", 10)
	y.Set(x)
	if isEven(x) {
		t.Fatalf("Expected false response from isEven")
	}
	if x.Cmp(y) != 0 {
		t.Fatalf("Input has been modified.  Expected=%s, Got=%s", y.Text(10), x.Text(10))
	}
}

func TestStepAction(t *testing.T) {
	x := new(big.Int)
	y := new(big.Int)
	x.SetString("22222222224444444444", 10)
	y.SetString("11111111112222222222", 10)
	stepAction(x)
	if ! cmpToN(x, y) {
		t.Fatalf("Expected: %s, Got: %s", y.Text(10), x.Text(10))
	}
	x.SetString("123456789012345", 10)
	y.SetString("370370367037036", 10)
	stepAction(x)
	if x.Cmp(y) != 0 {
		t.Fatalf("Expected: %s, Got: %s", y.Text(10), x.Text(10))
	}
}

func TestResolveN(t *testing.T) {
	tests := map[string]int{
		"1023": 62,
		"17": 12,
		"65537": 99,
		"2147575879": 649,
		"340282366920938463463374607431768220415": 1493,
	}
	for test := range(tests) {
		x := new(big.Int)
		y := new(big.Int)
		x.SetString(test, 10)
		y.Set(x)
		steps := resolveN(x)
		if steps != tests[test] {
			t.Fatalf("Expected: %d, Got: %d", tests[test], steps)
		}
		if x.Cmp(y) != 0 {
			t.Fatalf("Input has been modified.  Expected=%s, Got=%s", y.Text(10), x.Text(10))
		}
	}
}