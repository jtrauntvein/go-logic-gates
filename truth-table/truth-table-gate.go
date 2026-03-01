package truthtable

import (
	"errors"
	"math"

	gologicgates "github.com/jhtrauntvein/go-logic-gates"
)

// Defines errors that can be returned by the truth table gate
var (
	ErrTTInvalidIndex        = errors.New("invalid index for generating the inputs row")
	ErrTTNoOutputs           = errors.New("at least one output must be defined for a truth table row")
	ErrTTInconsistentOutputs = errors.New("all truth table output arrays must have the same length")
)

// Defines the output state of a logic device for a given combination of inputs.  The actual number
type TruthTableRow struct {
	// Specifies the bits for input values in this row
	Bits uint64

	// Specifies the maximum number of bits that will be used
	BitsCount int

	// Defines the outputs for this row
	Outputs []bool
}

// Construct a new truth table row
//
// Returns:
//   - *TruthTableRow: Returns the row on success
//   - error: Returns nil on success
//
// Parameters:
//   - bits: Specifies the inputs that should be matched for the row.  Each bit in this value will
//     represent the expected state of the input line.
//   - bits_count: Specifies the maximum number of input bits that should be used
//   - outputs: Specifies the output that matches the input pattern
func NewTruthTableRow(bits uint64, inputs_count int, outputs []bool) (*TruthTableRow, error) {
	// check the input parameters
	if len(outputs) == 0 {
		return nil, ErrTTNoOutputs
	}
	if float64(bits) >= math.Pow(2, float64(inputs_count)) {
		return nil, ErrTTInvalidIndex
	}

	// generate the truth table row
	rtn := &TruthTableRow{
		Outputs:   outputs,
		Bits:      bits,
		BitsCount: inputs_count,
	}
	return rtn, nil
}

// Defines a logic gate using truth tables to match patterns of input for a
// selected output.
type TruthTableGate struct {
	// Specifies the truth table used for this gate.
	TruthTable []TruthTableRow

	// Specifies the collection of input lines for the logic gate
	Inputs []*gologicgates.Line

	// Specifies the collection out output lines for the logic gate
	Outputs []*gologicgates.Line
}

// Generate a new truth table gate.
//
// Returns:
//   - *TruthTableGate: Returns the gate on success
//   - error: Returns nil on success
//
// Parameters:
//   - solutions: Specifies the array of output row values that will be used to generate
//     the indexed inputs.  Each array in this parameter must have the same number of elements.
//     The number of bits in the inputs will be equal to the length of the set of solutions.
func NewTruthTableGate(solutions [][]bool) (*TruthTableGate, error) {
	// we need to check that each of the solutions is the same size.
	var solution_len int
	inputs_len := int(math.Log2(float64(len(solutions))))
	if len(solutions) == 0 {
		return nil, ErrTTNoOutputs
	}
	solution_len = len(solutions[0])
	if solution_len == 0 {
		return nil, ErrTTNoOutputs
	}

	// we can now generate the truth table
	var rtn = &TruthTableGate{
		TruthTable: make([]TruthTableRow, len(solutions)),
		Inputs:     make([]*gologicgates.Line, inputs_len),
		Outputs:    make([]*gologicgates.Line, solution_len),
	}
	for i := 0; i < inputs_len; i++ {
		rtn.Inputs[i] = gologicgates.NewLine()
	}
	for i := 0; i < solution_len; i++ {
		rtn.Outputs[i] = gologicgates.NewLine()
	}
	for i, solution := range solutions {
		if len(solution) != solution_len {
			return nil, ErrTTInconsistentOutputs
		}
		rtn.TruthTable[i].Outputs = solution
		rtn.TruthTable[i].Bits = uint64(i)
		rtn.TruthTable[i].BitsCount = len(solutions)
	}

	// we need to set the outputs based upon current input states
	result := rtn.Evaluate()
	for i, value := range rtn.Outputs {
		value.Set(result[i])
	}
	return rtn, nil
}

func (g *TruthTableGate) SetInput(value bool, index int) ([]bool, error) {
	var rtn []bool
	if index >= len(g.TruthTable) {
		return nil, gologicgates.ErrSetInvalidLine
	}
	g.Inputs[index].Set(value)
	rtn = g.Evaluate()
	for j, output_line := range g.Outputs {
		output_line.Set(rtn[j])
	}
	return rtn, nil
}

func (g *TruthTableGate) Evaluate() []bool {
	var rtn []bool
	var index = 0
	for i, input_line := range g.Inputs {
		mask := 1 << i
		value := input_line.Measure()
		if value {
			index |= mask
		}
	}
	rtn = g.TruthTable[index].Outputs
	return rtn
}

func (g *TruthTableGate) ConnectInputProbe(probe gologicgates.LineProbe, index int, send_value ...bool) error {
	if index < 0 || index >= len(g.Inputs) {
		return gologicgates.ErrGetInvalidLine
	}
	g.Inputs[index].ConnectProbe(probe, send_value...)
	return nil
}

func (g *TruthTableGate) DisconnectInputProbe(probe gologicgates.LineProbe, index int) error {
	if index < 0 || index >= len(g.Inputs) {
		return gologicgates.ErrGetInvalidLine
	}
	g.Inputs[index].DisconnectProbe(probe)
	return nil
}

func (g *TruthTableGate) ConnectOutputProbe(probe gologicgates.LineProbe, index int, send_value ...bool) error {
	if index < 0 || index >= len(g.Outputs) {
		return gologicgates.ErrGetInvalidLine
	}
	g.Outputs[index].ConnectProbe(probe, send_value...)
	return nil
}

func (g *TruthTableGate) DisconnectOutputProbe(probe gologicgates.LineProbe, index int) error {
	if index < 0 || index >= len(g.Outputs) {
		return gologicgates.ErrGetInvalidLine
	}
	g.Outputs[index].DisconnectProbe(probe)
	return nil
}
