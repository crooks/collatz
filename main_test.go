package main

import (
	"math/big"
	"testing"
)

func TestIsEven(t *testing.T) {
	x := new(big.Int)
	x.SetString("1234567890", 10)
	if ! isEven(x) {
		t.Fatalf("Expected true response from isEven")
	}
	x.SetString("12345678901", 10)
	if isEven(x) {
		t.Fatalf("Expected false response from isEven")
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
	if ! cmpToN(x, y) {
		t.Fatalf("Expected: %s, Got: %s", y.Text(10), x.Text(10))
	}
}