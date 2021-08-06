package state

import (
	"io/ioutil"
	"math/big"
	"os"
	"testing"
)

func TestFlags(t *testing.T) {
	expected_result := "collatz.yml"
	f := ParseFlags()
	if f.StateFile != expected_result {
		t.Fatalf("Expected statefile to contain \"%s\" but got \"%s\".", expected_result, f.StateFile)
	}
}

func TestState(t *testing.T) {
	expectedStartInt := "0"
	expectedHighInt := "0"
	expectedHighSteps := 0
	state, err := ParseState("collatz.yml")
	if err != nil {
		t.Fatalf("Failed to read state file: %v", err)
	}
	if state.StartInt != "0" {
		t.Fatalf("Expected state.StartInt to contain \"%s\" but got \"%s\".", expectedStartInt, state.StartInt)
	}
	if state.HighInt != "0" {
		t.Fatalf("Expected state.HighInt to contain \"%s\" but got \"%s\".", expectedHighInt, state.HighInt)
	}
	if state.HighSteps != 0 {
		t.Fatalf("Expected state.HighSteps to contain \"%d\" but got \"%d\".", expectedHighSteps, state.HighSteps)
	}
}

func TestStateWrite(t *testing.T) {
	testFile, err := ioutil.TempFile("/tmp", "collatz")
	if err != nil {
		t.Fatalf("Unable to create TempFile: %v", err)
	}
	defer os.Remove(testFile.Name())
	state := new(StateInfo)
	state.StartInt = "340282366920938463463374607431768211456"
	state.RestartInt = "0"
	state.HighInt = "340282366920938463463374607431768211457"
	state.HighSteps = 100
	err = state.WriteState(testFile.Name())
	if err != nil {
		t.Fatalf("Failed to write state file: %v", err)
	}
	rState, err := ParseState(testFile.Name())
	if err != nil {
		t.Fatalf("Failed to read state file: %v", err)
	}
	if rState.StartInt != state.StartInt {
		t.Fatalf("Expected state.StartInt to contain \"%s\" but got \"%s\".", state.StartInt, rState.StartInt)
	}
	if rState.HighInt != state.HighInt {
		t.Fatalf("Expected state.HighInt to contain \"%s\" but got \"%s\".", state.HighInt, rState.HighInt)
	}
	if rState.HighSteps != state.HighSteps {
		t.Fatalf("Expected state.HighSteps to contain \"%d\" but got \"%d\".", state.HighSteps, rState.HighSteps)
	}
}

func TestStartFrom(t *testing.T) {
	state := new(StateInfo)
	// Restart is greater than Start so the result should be Restart
	state.StartInt = "340282366920938463463374607431768211456"
	state.RestartInt = "340282366920938463463374607431768211457"
	restart := new(big.Int)
	restart.SetString(state.RestartInt, 10)
	testStart := state.StartFrom()
	if testStart.Cmp(restart) != 0 {
		t.Fatalf("Expected testStart to contain \"%s\" but got \"%s\".", restart.Text(10), testStart.Text(10))
	}
	// Make Start greater than Restart.  Now StartFrom should return StartInt (as a bigInt).
	state.StartInt = "340282366920938463463374607431768211458"
	start := new(big.Int)
	start.SetString(state.StartInt, 10)
	testStart = state.StartFrom()
	if testStart.Cmp(start) != 0 {
		t.Fatalf("Expected testStart to contain \"%s\" but got \"%s\".", start.Text(10), testStart.Text(10))
	}
}
