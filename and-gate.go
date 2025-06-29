package gologicgates

import "fmt"

// Defines a two input logic AND gate
type AndGate2 struct {
	TruthTableGate
}

// Generates a new two input AND gate
//
// Returns:
//   - *AndGate2: returns the and gate
func NewAndGate2() *AndGate2 {
	truth_table, err := NewTruthTableGate([][]bool{
		{ false },
		{ false },
		{ false },
		{ true },
	})
	if err != nil {
		panic(fmt.Errorf("AND2 failed to allocate truth table: %v", err))
	}
	rtn := &AndGate2{
		TruthTableGate: *truth_table,
	}
	return rtn
}
