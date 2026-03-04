package latches

import (
	gologicgates "github.com/jtrauntvein/go-logic-gates"
	truthtable "github.com/jtrauntvein/go-logic-gates/truth-table"
)

// Defines a two input latch device that implements the BasicGate interface by encapsulating
// two connected NOR gates.
type SrLatch struct {
	gates       []gologicgates.BasicGate
	connections []*gologicgates.Connection
}

// Defines the port numbers associatred with input names.
const (
	SrSet   = 0
	SrReset = 1
)

// Defines the port numbers associated with output line names
const (
	SrQ    = 0
	SrQNot = 1
)

// Generates a new SrLatch object
//
// Returns:
//   - *SrLatch: Returns the new object
func NewSrLatch() *SrLatch {
	gates := []gologicgates.BasicGate{
		truthtable.NewNorGate2(),
		truthtable.NewNorGate2(),
	}
	rtn := &SrLatch{
		gates: gates,
		connections: []*gologicgates.Connection{
			gologicgates.MakeConnection(gates[0], 0, gates[1], 0),
			gologicgates.MakeConnection(gates[1], 0, gates[0], 1),
		},
	}
	return rtn
}

func (d *SrLatch) SetInput(value bool, index int) ([]bool, error) {
	var err error
	switch index {
	case SrSet:
		_, err = d.gates[0].SetInput(value, 0)
		if err != nil {
			return nil, err
		}

	case SrReset:
		_, err = d.gates[1].SetInput(value, 1)
		if err != nil {
			return nil, err
		}

	default:
		return nil, gologicgates.ErrSetInvalidLine
	}
	return d.Evaluate(), nil
}

func (d *SrLatch) Evaluate() []bool {
	return []bool{
		d.gates[0].Evaluate()[0],
		d.gates[1].Evaluate()[0],
	}
}

func (d *SrLatch) ConnectInputProbe(probe gologicgates.LineProbe, index int, send_value ...bool) error {
	var err error
	switch index {
	case SrSet:
		err = d.gates[1].ConnectInputProbe(probe, 1)
		if err != nil {
			return err
		}

	case SrReset:
		err = d.gates[0].ConnectInputProbe(probe, 0)
		if err != nil {
			return err
		}

	default:
		return gologicgates.ErrGetInvalidLine
	}
	return nil
}

func (d *SrLatch) DisconnectInputProbe(probe gologicgates.LineProbe, index int) error {
	var err error
	switch index {
	case SrReset:
		err = d.gates[0].DisconnectInputProbe(probe, 0)
		if err != nil {
			return err
		}

	case SrSet:
		err = d.gates[1].DisconnectInputProbe(probe, 1)
		if err != nil {
			return err
		}

	default:
		return gologicgates.ErrGetInvalidLine
	}
	return nil
}

func (d *SrLatch) ConnectOutputProbe(probe gologicgates.LineProbe, index int, send_value ...bool) error {
	var err error
	switch index {
	case SrQ:
		err = d.gates[0].ConnectOutputProbe(probe, 0, send_value...)

	case SrQNot:
		err = d.gates[1].ConnectOutputProbe(probe, 0, send_value...)

	default:
		err = gologicgates.ErrGetInvalidLine
	}
	return err
}

func (d *SrLatch) DisconnectOutputProbe(probe gologicgates.LineProbe, index int) error {
	var err error
	switch index {
	case SrQ:
		err = d.gates[0].DisconnectOutputProbe(probe, 0)

	case SrQNot:
		err = d.gates[1].DisconnectOutputProbe(probe, 0)

	default:
		err = gologicgates.ErrGetInvalidLine
	}
	return err
}
