package truthtable

// Defines a single input gate that outputs the inverse of its input
type NotGate struct {
	TruthTableGate
}

// Constructs a new inverter object
//
// Returns:
//
//	*NotGate: returns the constructed gate
func NewNotGate() *NotGate {
	truth_table, err := NewTruthTableGate([][]bool{
		{true},
		{false},
	})
	if err != nil {
		panic(err)
	}
	rtn := &NotGate{
		TruthTableGate: *truth_table,
	}
	rtn.SetInput(false, 0)
	return rtn
}
