package main

import (
	"log"
	"math/big"
	"time"

	"github.com/crooks/collatz/state"
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
func resolveN(n *big.Int) (steps uint64) {
	// Take a copy of N so we can modify it without losing our iteration placeholder.
	current := new(big.Int)
	current.Set(n)
	inProg := false
	for {
		stepAction(current)
		steps++
		if steps%1000 == 0 {
			inProg = true
			log.Printf("In progress: Start=%s, Steps=%d\n", n.Text(10), steps)
		}
		if cmpToN(current, v1) {
			break
		}
	}
	if inProg {
		log.Printf("Resolved:    Start=%s, Steps=%d\n", n.Text(10), steps)
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
	flags := state.ParseFlags()
	cfg, err := state.ParseState(flags.StateFile)
	if err != nil {
		log.Fatalf("Unable to parse %s: %v", flags.StateFile, err)
	}
	n := cfg.StartFrom()
	highInt := new(big.Int)
	highInt.SetString(cfg.HighInt, 10)
	// start and writeInterval are used to determine when to write the current state to file.
	start := time.Now()
	writeInterval := time.Duration(cfg.WriteInterval) * time.Second
	var iterationsPerWrite uint64 = 0
	// Being iterating candidate integers
	log.Printf("Starting from: %s", n.Text(10))
	for {
		steps := resolveN(n)
		iterationsPerWrite++
		if steps > cfg.HighSteps {
			cfg.HighSteps = steps
			highInt.Set(n)
			log.Printf("%s: Start=%s, Steps=%d, HighStart=%s, highSteps=%d\n", timestamp(), n.Text(10), steps, highInt.Text(10), cfg.HighSteps)
		}
		if time.Since(start) > time.Duration(writeInterval) {
			cfg.HighInt = highInt.Text(10)
			cfg.RestartInt = n.Text(10)
			cfg.WriteState(flags.StateFile)
			log.Printf("Current state written to %s.  Iterations per Second: %d", flags.StateFile, iterationsPerWrite/uint64(cfg.WriteInterval))
			iterationsPerWrite = 0
			start = time.Now()
		}
		n.Add(n, v1)
		//time.Sleep(1 * time.Second)
	}
}
