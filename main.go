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
func isEven(foo *big.Int) bool {
	m := new(big.Int)
	foo.DivMod(foo, v2, m)
	return cmpToN(m, v0)
}

func stepAction(current *big.Int) {
	foo := new(big.Int)
	foo.Set(current)
	if isEven(foo) {
		current.Div(current, v2)
	} else {
		current.Mul(current, v3)
		current.Add(current, v1)
	}
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
	n.Exp(big.NewInt(2), big.NewInt(128), nil)
	current := new(big.Int)
	highStart := new(big.Int)
	var highScore int = 0
	for {
		current.Set(n)
		steps := 0
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
		if steps > highScore {
			highScore = steps
			highStart.Set(n)
			fmt.Printf("%s: Start=%s, Steps=%d, HighStart=%s, HighSteps=%d\n", timestamp(), n.Text(10), steps, highStart.Text(10), highScore)
		}
		n.Add(n, v1)
		//time.Sleep(1 * time.Second)
	}
}
