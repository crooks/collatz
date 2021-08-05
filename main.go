package main

import (
	"fmt"
	"math/big"
)

func cmpToN(x, y *big.Int) bool {
	result := x.Cmp(y)
	if result == 0 {
		return true
	}
	return false
}

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

var (
	v0 *big.Int
	v1 *big.Int
	v2 *big.Int
	v3 *big.Int
)

func main() {
	v0 = big.NewInt(0)
	v1 = big.NewInt(1)
	v2 = big.NewInt(2)
	v3 = big.NewInt(3)

	n := new(big.Int)
	n.Exp(big.NewInt(10), big.NewInt(20), nil)
	current := new(big.Int)
	highStart := new(big.Int)
	var highScore int = 0
	for {
		current.Set(n)
		steps := 0
		for  {
			stepAction(current)
			steps++
			if steps % 2000 == 0 {
				fmt.Printf("In progress: Start=%s, Steps=%d\n", n.Text(10), steps)
			}
			if cmpToN(current, v1) {
				break
			}
		}
		if steps > highScore {
			highScore = steps
			highStart.Set(n)
			fmt.Printf("Start=%s, Steps=%d\n", n.Text(10), steps)
		}
		n.Add(n, v1)
	}
}