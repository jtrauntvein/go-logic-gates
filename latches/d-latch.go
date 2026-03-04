package latches

import (
	gologicgates "github.com/jtrauntvein/go-logic-gates"
	truthtable "github.com/jtrauntvein/go-logic-gates/truth-table"
)

// Defines an object that works as a D-Latch
type DLatch struct {
	input_gates  []gologicgates.BasicGate
	output_gates []gologicgates.BasicGate
	inverter     gologicgates.BasicGate
	connections  []*gologicgates.Connection
}

// Defines the indices for the input pins
const (
	DData   = 0
	DEnable = 1
)

// defines the indices for the outputs
const (
	Dq    = 0
	DqNot = 1
)

// Construct a new d-latch
//
// Returns:
//   - *DLatch: returns the constructed gate
func NewDLatch() *DLatch {
	input_gates := []gologicgates.BasicGate{
		truthtable.NewAndGate2(),
		truthtable.NewAndGate2(),
	}
	output_gates := []gologicgates.BasicGate{
		truthtable.NewNorGate2(),
		truthtable.NewNorGate2(),
	}
	inverter := truthtable.NewNotGate()
	rtn := &DLatch{
		input_gates:  input_gates,
		output_gates: output_gates,
		inverter:     inverter,
		connections: []*gologicgates.Connection{
			gologicgates.MakeConnection(inverter, 0, input_gates[0], 0),
			gologicgates.MakeConnection(input_gates[0], 0, output_gates[0], 0),
			gologicgates.MakeConnection(input_gates[1], 0, output_gates[1], 1),
			gologicgates.MakeConnection(output_gates[0], 0, output_gates[1], 0),
			gologicgates.MakeConnection(output_gates[1], 0, output_gates[0], 1),
		},
	}

	// latch a false state in before returning
	rtn.SetInput(false, DData)
	rtn.SetInput(true, DEnable)
	rtn.SetInput(false, DEnable)
	return rtn
}

func (d *DLatch) SetInput(value bool, index int) ([]bool, error) {
	var err error
	switch index {
	case DData:
		_, err = d.inverter.SetInput(value, 0)
		if err != nil {
			return nil, err
		}
		_, err = d.input_gates[1].SetInput(value, 1)
		if err != nil {
			return nil, err
		}

	case DEnable:
		_, err = d.input_gates[0].SetInput(value, 1)
		if err != nil {
			return nil, err
		}
		_, err = d.input_gates[1].SetInput(value, 0)
		if err != nil {
			return nil, err
		}

	default:
		return nil, gologicgates.ErrSetInvalidLine
	}
	return d.Evaluate(), nil
}

func (d *DLatch) Evaluate() []bool {
	return []bool{d.output_gates[0].Evaluate()[0], d.output_gates[1].Evaluate()[0]}
}

func (d *DLatch) ConnectInputProbe(probe gologicgates.LineProbe, index int, send_value ...bool) error {
	var err error
	switch index {
	case DData:
		err = d.input_gates[0].ConnectInputProbe(probe, 0, send_value...)

	case DEnable:
		err = d.input_gates[1].ConnectInputProbe(probe, 1, send_value...)

	default:
		err = gologicgates.ErrGetInvalidLine
	}
	return err
}

func (d *DLatch) DisconnectInputProbe(probe gologicgates.LineProbe, index int) error {
	var err error
	switch index {
	case DData:
		err = d.input_gates[0].DisconnectInputProbe(probe, 0)

	case DEnable:
		err = d.input_gates[1].DisconnectInputProbe(probe, 1)

	default:
		err = gologicgates.ErrGetInvalidLine
	}
	return err
}

func (d *DLatch) ConnectOutputProbe(probe gologicgates.LineProbe, index int, send_value ...bool) error {
	var err error
	switch index {
	case Dq:
		err = d.output_gates[0].ConnectOutputProbe(probe, 0, send_value...)
		if err != nil {
			return gologicgates.ErrGetInvalidLine
		}

	case DqNot:
		err = d.output_gates[1].ConnectOutputProbe(probe, 0, send_value...)
		if err != nil {
			return gologicgates.ErrGetInvalidLine
		}

	default:
		return gologicgates.ErrGetInvalidLine
	}
	return nil
}

func (d *DLatch) DisconnectOutputProbe(probe gologicgates.LineProbe, index int) error {
	switch index {
	case Dq:
		err := d.output_gates[0].DisconnectOutputProbe(probe, 0)
		if err != nil {
			return gologicgates.ErrGetInvalidLine
		}

	case DqNot:
		err := d.output_gates[1].DisconnectOutputProbe(probe, 0)
		if err != nil {
			return gologicgates.ErrGetInvalidLine
		}

	default:
		return gologicgates.ErrGetInvalidLine
	}
	return nil
}
