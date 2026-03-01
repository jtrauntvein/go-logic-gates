package gologicgates

// Defines an object that implements a set of parallel lines as if it were a logic gate.
// the logic function is that the output is always equal to the input for any line.  This
// class exists to make it more convenient to connect high level components such as registers,
// alus, & etc.
type Bus struct {
	lines []Line
}

// @return:
//   - int: Returns the number of lines allocated for the bus
func (b *Bus) Size() int {
	return len(b.lines)
}

// Allocates a new bus with the given number of lines
// 
// @return:
//   - *Bus: returns the allocated bus
//
// Parameters:
//   - size: Specifies the number of lines to allocate for the bus.  This value must be greater than zero
//   or the function will panic
func NewBus(size int) *Bus {
	if size <= 0 {
		panic("bus size must be greater than zero")
	}
	rtn := &Bus{
		lines: make([]Line, size),
	}
	return rtn
}

func (b *Bus) SetInput(value bool, index int) ([]bool, error) {
	if index > len(b.lines) {
		return nil, ErrSetInvalidLine
	}
	b.lines[index].Set(value)
	return b.Evaluate(), nil
}

func (b *Bus) Evaluate() []bool {
	rtn := make([]bool, len(b.lines))
	for index, line := range b.lines {
		rtn[index] = line.value
	}
	return rtn
}

func (b *Bus) ConnectInputProbe(probe LineProbe, index int, send_value ...bool) error {
	if index > len(b.lines) {
		return ErrGetInvalidLine
	}
	b.lines[index].ConnectProbe(probe, send_value...)
	return nil
}

func (b *Bus) DisconnectInputProbe(probe LineProbe, index int) error {
	if index > len(b.lines) {
		return ErrGetInvalidLine
	}
	b.lines[index].DisconnectProbe(probe)
	return nil
}

func (b *Bus) ConnectOutputProbe(probe LineProbe, index int, send_value ...bool) error {
	if index > len(b.lines) {
		return ErrGetInvalidLine
	}
	b.lines[index].ConnectProbe(probe, send_value...)
	return nil
}

func (b *Bus) DisconnectOutputProbe(probe LineProbe, index int) error {
	if index > len(b.lines) {
		return ErrGetInvalidLine
	}
	b.lines[index].DisconnectProbe(probe)
	return nil
}
