package truthtable

import "fmt"

// Defines a two input gate that implements the logic OR function
type OrGate2 struct {
	TruthTableGate
}

// Allocates a new two input OR gate
//
// Returns:
//   - *OrGate2: Returns the allocated gate
func NewOr2Gate() *OrGate2 {
	truth_table, err := NewTruthTableGate([][]bool{
		{false},
		{true},
		{true},
		{true},
	})
	if err != nil {
		panic(fmt.Errorf("failed to allocate truth table for or gate: %v", err))
	}
	return &OrGate2{
		TruthTableGate: *truth_table,
	}
}
