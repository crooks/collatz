package main

import (
	"fmt"
	"math/big"
	"time"
)

// cmpToN compares two bigints.  If they match, it returns true.
func cmpToN(x, y *big.Int) bool {
	return x.Cmp(y) == 0
}

// isEven tests if a given bigint is odd or even
func isEven(x *big.Int) bool {
	m := new(big.Int)
	m.DivMod(x, v2, m)
	return cmpToN(m, v0)
}

func stepAction(current *big.Int) {
	if isEven(current) {
		current.Div(current, v2)
		} else {
			current.Mul(current, v3)
			current.Add(current, v1)
		}
	}
	
// resolveN attempts to resolve a given bigint N to 1 using Collatz Conjecture
func resolveN(n *big.Int) (steps int) {
	// Take a copy of N so we can modify it without losing our iteration placeholder.
	current := new(big.Int)
	current.Set(n)
	inProg := false
	for {
		stepAction(current)
		steps++
		if steps%1500 == 0 {
			inProg = true
			fmt.Printf("In progress: Start=%s, Steps=%d\n", n.Text(10), steps)
		}
		if cmpToN(current, v1) {
			break
		}
	}
	if inProg {
		fmt.Printf("Resolved:    Start=%s, Steps=%d\n", n.Text(10), steps)
	}
	return
}

func timestamp() string {
	t := time.Now()
	return t.Format("2006-01-02 15:04:05")
}

var (
	v0 *big.Int
	v1 *big.Int
	v2 *big.Int
	v3 *big.Int
)

func init() {
	v0 = big.NewInt(0)
	v1 = big.NewInt(1)
	v2 = big.NewInt(2)
	v3 = big.NewInt(3)
}

func main() {
	n := new(big.Int)
	n.Exp(big.NewInt(2), big.NewInt(68), nil)
	highStart := new(big.Int)
	var highScore int = 0
	// Being interating candidate 
	for {
		steps := resolveN(n)
		if steps > highScore {
			highScore = steps
			highStart.Set(n)
			fmt.Printf("%s: Start=%s, Steps=%d, HighStart=%s, HighSteps=%d\n", timestamp(), n.Text(10), steps, highStart.Text(10), highScore)
		}
		n.Add(n, v1)
		//time.Sleep(1 * time.Second)
	}
}
