package registers

import gates "github.com/jtrauntvein/go-logic-gates"

const (
	TriBufferData   = 0
	TriBufferEnable = 1
)

// Defines buffered data line that will only report outputs when the write-enable input is set
type TriStateBuffer struct {
	input  *gates.Line
	output *gates.Line
	enable *gates.Line
}

func (b *TriStateBuffer) SetInput(value bool, index int) ([]bool, error) {
	var err error
	switch index {
	case TriBufferData:
		b.input.Set(value)
		if b.enable.Value() {
			b.output.Set(value)
		}

	case TriBufferEnable:
		b.enable.Set(value)
		if value {
			b.output.Set(b.input.Value())
		}

	default:
		err = gates.ErrSetInvalidLine
	}
	return b.Evaluate(), err
}

func (b *TriStateBuffer) Evaluate() []bool {
	return []bool{b.output.Value()}
}

func (b *TriStateBuffer) ConnectInputProbe(probe gates.LineProbe, index int, send_value ...bool) error {
	var err error
	switch index {
	case TriBufferData:
		b.input.ConnectProbe(probe, send_value...)

	case TriBufferEnable:
		b.enable.ConnectProbe(probe, send_value...)

	default:
		err = gates.ErrGetInvalidLine
	}
	return err
}

func (b *TriStateBuffer) DisconnectInputProbe(probe gates.LineProbe, index int) error {
	var err error
	switch index {
	case TriBufferData:
		b.input.DisconnectProbe(probe)

	case TriBufferEnable:
		b.enable.DisconnectProbe(probe)

	default:
		err = gates.ErrGetInvalidLine
	}
	return err
}

func (b *TriStateBuffer) ConnectOutputProbe(probe gates.LineProbe, index int, send_value ...bool) error {
	var err error
	switch index {
	case TriBufferData:
		b.output.ConnectProbe(probe)

	default:
		err = gates.ErrGetInvalidLine
	}
	return err
}

func (b *TriStateBuffer) DisconnectOutputProbe(probe gates.LineProbe, index int) error {
	var err error
	switch index {
	case TriBufferData:
		b.output.DisconnectProbe(probe)

	default:
		err = gates.ErrGetInvalidLine
	}
	return err
}

// Return:
//   - *TriStateBuffer: returns a new tri state buffer
func NewTriStateBuffer() *TriStateBuffer {
	return &TriStateBuffer{
		input:  gates.NewLine(),
		output: gates.NewLine(),
		enable: gates.NewLine(),
	}
}
