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
		"9": 19,
		"97": 118,
		"871": 178,
		"6171": 261,
		"77031": 350,
		"837799": 524,
		"8400511": 685,
		"63728127": 949,
		"670617279": 986,
		"9780657630": 1132,
		"75128138247": 1228,
		"989345275647": 1348,
		"7887663552367": 1563,
		"80867137596217": 1662,
		"942488749153153": 1862,
		"7579309213675935": 1958,
		"93571393692802302": 2091,
		"931386509544713451": 2283,
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

func BenchmarkResolveN(b *testing.B) {
	tests := []string{
		"9",
		"97",
		"871",
		"6171",
		"77031",
		"837799",
		"8400511",
		"63728127",
		"670617279",
		"9780657630",
		"75128138247",
		"989345275647",
		"7887663552367",
		"80867137596217",
		"942488749153153",
		"7579309213675935",
		"93571393692802302",
		"931386509544713451",
	}
	b.ResetTimer()
	for _, test := range(tests) {
		x := new(big.Int)
		x.SetString(test, 10)
		resolveN(x)
	}
}
