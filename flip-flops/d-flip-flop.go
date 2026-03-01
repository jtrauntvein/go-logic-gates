package flipflops

import (
	gologicgates "github.com/jhtrauntvein/go-logic-gates"
	latches "github.com/jhtrauntvein/go-logic-gates/latches"
	truthtable "github.com/jhtrauntvein/go-logic-gates/truth-table"
)

// defines the indices for the input lines
const (
	DData  = 0
	DClock = 1
)

// defines the indices for the output lines
const (
	Dq    = latches.Dq
	DqNot = latches.DqNot
)

// Defines an object that represents a D-Flip flop which toggles its output based upon
// a sample of the data input on the rising edge of the clock line.
type DFlipFlop struct {
	master_latch *latches.DLatch
	slave_latch  *latches.DLatch
	inverters    []*truthtable.NotGate
	connections  []*gologicgates.Connection
}

// Construct a new d-flip-flop
//
// Returns
//   - *DFlipFlop: Returns an initialised D flip flop
func NewDFlipFlop() *DFlipFlop {
	master_latch := latches.NewDLatch()
	slave_latch := latches.NewDLatch()
	inverters := []*truthtable.NotGate{
		truthtable.NewNotGate(),
		truthtable.NewNotGate(),
	}
	rtn := &DFlipFlop{
		master_latch: master_latch,
		slave_latch:  slave_latch,
		inverters:    inverters,
		connections: []*gologicgates.Connection{
			gologicgates.MakeConnection(inverters[0], 0, master_latch, latches.DEnable),
			gologicgates.MakeConnection(inverters[1], 0, slave_latch, latches.DEnable),
			gologicgates.MakeConnection(inverters[0], 0, inverters[1], 0),
			gologicgates.MakeConnection(master_latch, latches.Dq, slave_latch, latches.DData),
		},
	}
	return rtn
}

func (d *DFlipFlop) SetInput(value bool, index int) ([]bool, error) {
	var err error
	switch index {
	case DData:
		_, err = d.master_latch.SetInput(value, latches.DData)
		if err != nil {
			return nil, err
		}

	case DClock:
		_, err = d.inverters[0].SetInput(value, 0)
		if err != nil {
			return nil, err
		}

	default:
		return nil, gologicgates.ErrSetInvalidLine
	}
	return d.Evaluate(), nil
}

func (d *DFlipFlop) Evaluate() []bool {
	return d.slave_latch.Evaluate()
}

func (d *DFlipFlop) ConnectInputProbe(probe gologicgates.LineProbe, index int, send_value ...bool) error {
	var err error
	switch index {
	case DData:
		err = d.master_latch.ConnectInputProbe(probe, latches.DData, send_value...)

	case DClock:
		err = d.inverters[0].ConnectInputProbe(probe, 0, send_value...)

	default:
		err = gologicgates.ErrGetInvalidLine
	}
	return err
}

func (d *DFlipFlop) DisconnectInputProbe(probe gologicgates.LineProbe, index int) error {
	var err error
	switch index {
	case DData:
		err = d.master_latch.DisconnectInputProbe(probe, latches.DData)

	case DClock:
		err = d.inverters[0].DisconnectInputProbe(probe, 0)

	default:
		err = gologicgates.ErrGetInvalidLine
	}
	return err
}

func (d *DFlipFlop) ConnectOutputProbe(probe gologicgates.LineProbe, index int, send_value ...bool) error {
	var err error
	switch index {
	case Dq:
		err = d.slave_latch.ConnectOutputProbe(probe, latches.Dq, send_value...)

	case DqNot:
		err = d.slave_latch.ConnectOutputProbe(probe, latches.DqNot, send_value...)

	default:
		err = gologicgates.ErrGetInvalidLine
	}
	return err
}

func (d *DFlipFlop) DisconnectOutputProbe(probe gologicgates.LineProbe, index int) error {
	var err error
	switch index {
	case Dq:
		err = d.slave_latch.DisconnectOutputProbe(probe, latches.Dq)

	case DqNot:
		err = d.slave_latch.DisconnectOutputProbe(probe, latches.DqNot)

	default:
		err = gologicgates.ErrGetInvalidLine
	}
	return err
}

// Pulses the clock to true and then back to false
func (d *DFlipFlop) PulseClock() {
	d.SetInput(true, DClock)
	d.SetInput(false, DClock)
}
