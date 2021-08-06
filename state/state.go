package state

import (
	"flag"
	"io/ioutil"
	"math/big"
	"os"

	"gopkg.in/yaml.v2"
)

type StateInfo struct {
	// These should never be changed, except by the user.
	StartInt      string `yaml:"start_integer"`
	WriteInterval int    `yaml:"write_interval"`
	// All the following will be written during operation.
	HighInt    string `yaml:"high_integer"`
	HighSteps  uint64 `yaml:"high_steps"`
	RestartInt string `yaml:"restart_integer"`
}

type Flags struct {
	StateFile string
}

func ParseFlags() *Flags {
	f := new(Flags)
	flag.StringVar(
		&f.StateFile,
		"config",
		"collatz.yml",
		"Path to the Collatz state file",
	)
	flag.Parse()
	return f
}

func (s *StateInfo) StartFrom() *big.Int {
	start := new(big.Int)
	start.SetString(s.StartInt, 10)
	restart := new(big.Int)
	restart.SetString(s.RestartInt, 10)
	if restart.Cmp(start) == 1 {
		return restart
	}
	return start
}

func ParseState(filename string) (*StateInfo, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	y := yaml.NewDecoder(file)
	state := new(StateInfo)
	if err := y.Decode(&state); err != nil {
		return nil, err
	}
	return state, nil
}

func (s *StateInfo) WriteState(filename string) error {
	data, err := yaml.Marshal(s)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
