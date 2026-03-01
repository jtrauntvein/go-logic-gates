package gologicgates_test

import (
	gologicgates "github.com/jhtrauntvein/go-logic-gates"
	truthtable "github.com/jhtrauntvein/go-logic-gates/truth-table"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type AndGate4 struct {
	input_layer          []gologicgates.BasicGate
	internal_connections []*gologicgates.Connection
	output_gate          gologicgates.BasicGate
}

func NewAndGate4() *AndGate4 {
	input_layer := []gologicgates.BasicGate{
		truthtable.NewAndGate2(),
		truthtable.NewAndGate2(),
	}
	output_gate := truthtable.NewAndGate2()
	rtn := &AndGate4{
		input_layer: input_layer,
		output_gate: output_gate,
		internal_connections: []*gologicgates.Connection{
			gologicgates.MakeConnection(input_layer[0], 0, output_gate, 0),
			gologicgates.MakeConnection(input_layer[1], 0, output_gate, 1),
		},
	}
	return rtn
}

func (g *AndGate4) SetInput(val bool, index int) ([]bool, error) {
	var err error
	if index < 2 {
		_, err = g.input_layer[0].SetInput(val, index%2)
	} else {
		_, err = g.input_layer[1].SetInput(val, index%2)
	}
	if err != nil {
		return nil, err
	}
	return g.output_gate.Evaluate(), nil
}

func (g *AndGate4) Evaluate() []bool {
	return g.output_gate.Evaluate()
}

func (g *AndGate4) ConnectInputProbe(probe gologicgates.LineProbe, index int, send_value ...bool) error {
	return g.input_layer[index%2].ConnectInputProbe(probe, index%2, send_value...)
}

func (g *AndGate4) DisconnectInputProbe(probe gologicgates.LineProbe, index int) error {
	return g.input_layer[index%2].DisconnectInputProbe(probe, index%2)
}

func (g *AndGate4) ConnectOutputProbe(probe gologicgates.LineProbe, index int, send_value ...bool) error {
	return g.output_gate.ConnectOutputProbe(probe, index, send_value...)
}

var _ = Describe("line connection unit tests", func() {
	It("can connect two gates", func() {
		// we should be able to run through the truth table for the four input configuration
		var err error
		gate := NewAndGate4()
		var last_output []bool
		expected := [][]bool{
			{false, false, false, false, false},
			{false, false, false, true, false},
			{false, false, true, false, false},
			{false, false, true, true, false},
			{false, true, false, false, false},
			{false, true, false, true, false},
			{false, true, true, false, false},
			{false, true, true, true, false},
			{true, false, false, false, false},
			{true, false, false, true, false},
			{true, false, true, false, false},
			{true, false, true, true, false},
			{true, true, false, false, false},
			{true, true, false, true, false},
			{true, true, true, false, false},
			{true, true, true, true, true},
		}
		for _, test := range expected {
			for j, value := range test {
				if j < 4 {
					last_output, err = gate.SetInput(value, j)
					Expect(err).ToNot(HaveOccurred())
				}
			}
			Expect(last_output).To(HaveLen(1))
			Expect(last_output[0]).To(Equal(test[4]))
		}
	})
})
