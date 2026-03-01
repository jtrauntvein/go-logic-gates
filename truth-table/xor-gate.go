package truthtable

// Defines a two input logic XOR gate implemented in terms of a truth table.
type XorGate2 struct {
	TruthTableGate
}

// Generates a new XOR two input logic gate
//
// Returns:
//   - *XorGate2: Returns the allocated gate
func NewXorGate2() *XorGate2 {
	truth_table, err := NewTruthTableGate([][]bool{
		{false},
		{true},
		{true},
		{false},
	})
	if err != nil {
		panic(err)
	}
	return &XorGate2{
		TruthTableGate: *truth_table,
	}
}
