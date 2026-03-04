package registers

import (
	gologicgates "github.com/jtrauntvein/go-logic-gates"
	dflipflops "github.com/jtrauntvein/go-logic-gates/flip-flops"
	truthtable "github.com/jtrauntvein/go-logic-gates/truth-table"
)

// Defines the indices for lines for this device
const (
	DData    = 0
	DClock   = 1
	DWEnable = 2
	DLoad    = 3
	Dq       = 0
)

// Defines a register object that stores one bit of data with read-enable and write-enable
type OneBitRegister struct {
	cell           *dflipflops.DFlipFlop
	output_buff    *TriStateBuffer
	load_not       *truthtable.NotGate
	load_and_gates []*truthtable.AndGate2
	load_or_gate   *truthtable.OrGate2
	connections    []*gologicgates.Connection
}

// @return:
//   - OneBitRegister: Allocates a new one-bit register
func NewOneBitRegister() *OneBitRegister {
	cell := dflipflops.NewDFlipFlop()
	load_not := truthtable.NewNotGate()
	load_and_gates := []*truthtable.AndGate2{
		truthtable.NewAndGate2(),
		truthtable.NewAndGate2(),
	}
	load_or_gate := truthtable.NewOr2Gate()
	output := NewTriStateBuffer()
	rtn := &OneBitRegister{
		load_and_gates: load_and_gates,
		load_not:       load_not,
		load_or_gate:   load_or_gate,
		cell:           cell,
		output_buff:    output,
		connections: []*gologicgates.Connection{
			gologicgates.MakeConnection(load_not, 0, load_and_gates[0], 0),
			gologicgates.MakeConnection(load_and_gates[0], 0, load_or_gate, 0),
			gologicgates.MakeConnection(load_and_gates[1], 0, load_or_gate, 1),
			gologicgates.MakeConnection(load_or_gate, 0, cell, dflipflops.DData),
			gologicgates.MakeConnection(cell, dflipflops.Dq, output, TriBufferData),
			gologicgates.MakeConnection(cell, dflipflops.Dq, load_and_gates[0], 0),
		},
	}
	return rtn
}

func (r *OneBitRegister) SetInput(value bool, index int) ([]bool, error) {
	switch index {
	case DData:
		r.load_and_gates[1].SetInput(value, 1)

	case DClock:
		r.cell.SetInput(value, dflipflops.DClock)

	case DWEnable:
		r.output_buff.SetInput(value, TriBufferEnable)

	case DLoad:
		r.load_not.SetInput(value, 0)
		r.load_and_gates[1].SetInput(value, 0)

	default:
		return nil, gologicgates.ErrSetInvalidLine
	}
	return r.output_buff.Evaluate(), nil
}

func (r *OneBitRegister) Evaluate() []bool {
	return r.output_buff.Evaluate()
}

func (r *OneBitRegister) ConnectInputProbe(probe gologicgates.LineProbe, index int, send_value ...bool) error {
	var err error
	switch index {
	case DData:
		err = r.load_and_gates[1].ConnectInputProbe(probe, 1, send_value...)

	case DClock:
		err = r.cell.ConnectInputProbe(probe, dflipflops.DClock, send_value...)

	case DLoad:
		err = r.load_not.ConnectInputProbe(probe, 0, send_value...)

	case DWEnable:
		err = r.output_buff.ConnectInputProbe(probe, TriBufferEnable, send_value...)

	default:
		err = gologicgates.ErrGetInvalidLine
	}
	return err
}

func (r *OneBitRegister) DisconnectInputProbe(probe gologicgates.LineProbe, index int) error {
	var err error
	switch index {
	case DData:
		err = r.load_and_gates[1].DisconnectInputProbe(probe, 1)

	case DClock:
		err = r.cell.DisconnectInputProbe(probe, dflipflops.DClock)

	case DLoad:
		err = r.load_not.DisconnectInputProbe(probe, 0)

	case DWEnable:
		err = r.output_buff.DisconnectInputProbe(probe, TriBufferEnable)

	default:
		err = gologicgates.ErrGetInvalidLine
	}
	return err
}

func (r *OneBitRegister) ConnectOutputProbe(probe gologicgates.LineProbe, index int, send_value ...bool) error {
	var err error
	switch index {
	case Dq:
		err = r.output_buff.ConnectOutputProbe(probe, TriBufferData, send_value...)

	default:
		err = gologicgates.ErrGetInvalidLine
	}
	return err
}

func (r *OneBitRegister) DisconnectOutputProbe(probe gologicgates.LineProbe, index int) error {
	var err error
	switch index {
	case Dq:
		err = r.output_buff.DisconnectOutputProbe(probe, TriBufferData)

	default:
		err = gologicgates.ErrGetInvalidLine
	}
	return err
}

func (r *OneBitRegister) PulseClock() {
	r.cell.PulseClock()
}
