package truthtable

// Defines a two input NOR logic gate based on a truth table.
type NorGate2 struct {
	TruthTableGate
}

// Generates a a two input logic NOR gate
//
// Returns:
//   - *NorGate2: returns the allocated gate
func NewNorGate2() *NorGate2 {
	truth_table, err := NewTruthTableGate([][]bool{
		{true},
		{false},
		{false},
		{false},
	})
	if err != nil {
		panic(err)
	}
	gate := &NorGate2{
		TruthTableGate: *truth_table,
	}
	gate.SetInput(false, 0)
	gate.SetInput(false, 1)
	return gate
}
