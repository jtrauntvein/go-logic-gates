package truthtable

// Defines a two input logic NAND gate
type NandGate2 struct {
	TruthTableGate
}

// Generates a new two input Nand gate
//
// Returns:
//   - *NandGate2: Returns the initialised gate object
func NewNandGate2() *NandGate2 {
	truth_table, err := NewTruthTableGate([][]bool{
		{true},
		{true},
		{true},
		{false},
	})
	if err != nil {
		panic(err)
	}
	gate := &NandGate2{
		TruthTableGate: *truth_table,
	}
	gate.SetInput(false, 0)
	gate.SetInput(false, 1)
	return gate
}
