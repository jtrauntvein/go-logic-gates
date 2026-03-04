package registers

import gologicgates "github.com/jtrauntvein/go-logic-gates"

// Defines a memory register object that creates a single cell for each of the associated bus
// lines and that share a single load line as well as a single write enable line
type MemoryRegister struct {
	bus                *gologicgates.Bus
	cells              []*OneBitRegister
	input_connections  []*gologicgates.Connection
	output_connections []*gologicgates.Connection
}

// Allocates a new memory register object and connects all of the one bit register's inputs and outputs
// to the given bus.  The number of registers will be determined by the size of the bus
//
// Returns:
//   - *MemoryRegister: Returns the allocated register
//
// Parameters:
//   - bus: Specifies the bus to which the register is connected
func NewMemoryRegister(bus *gologicgates.Bus) *MemoryRegister {
	rtn := &MemoryRegister{
		bus:                bus,
		cells:              make([]*OneBitRegister, bus.Size()),
		input_connections:  make([]*gologicgates.Connection, bus.Size()),
		output_connections: make([]*gologicgates.Connection, bus.Size()),
	}
	for i := range bus.Size() {
		rtn.cells[i] = NewOneBitRegister()
		rtn.input_connections[i] = gologicgates.MakeConnection(bus, i, rtn.cells[i], DData)
		rtn.output_connections[i] = gologicgates.MakeConnection(rtn.cells[i], Dq, bus, i)
	}
	return rtn
}

// Returns:
//   - int: Returns the number of lines on the bus
func (m *MemoryRegister) Size() int {
	return m.bus.Size()
}

func (m *MemoryRegister) Evaluate() []bool {
	rtn := make([]bool, m.bus.Size())
	for i := range m.bus.Size() {
		rtn[i] = m.cells[i].Evaluate()[0]
	}
	return rtn
}

func (m *MemoryRegister) SetInput(value bool, idx int) ([]bool, error) {
	var err error
	switch idx {
	case DLoad:
		for _, cell := range m.cells {
			_, err = cell.SetInput(value, DLoad)
			if err != nil {
				return nil, err
			}
		}

	case DWEnable:
		for _, cell := range m.cells {
			_, err := cell.SetInput(value, DWEnable)
			if err != nil {
				return nil, err
			}
		}

	default:
		return nil, gologicgates.ErrSetInvalidLine
	}
	return m.Evaluate(), nil
}

func (m *MemoryRegister) ConnectInputProbe(probe gologicgates.LineProbe, idx int, send_value ...bool) error {
	var err error
	switch idx {
	case DLoad:
		err = m.cells[0].ConnectInputProbe(probe, DLoad, send_value...)

	case DWEnable:
		err = m.cells[0].ConnectInputProbe(probe, DWEnable, send_value...)

	default:
		err = gologicgates.ErrGetInvalidLine
	}
	return err
}

func (m *MemoryRegister) DisconnectInputProbe(probe gologicgates.LineProbe, idx int) error {
	var err error
	switch idx {
	case DLoad:
		err = m.cells[0].DisconnectInputProbe(probe, DLoad)

	case DWEnable:
		err = m.cells[0].DisconnectInputProbe(probe, DWEnable)

	default:
		err = gologicgates.ErrGetInvalidLine
	}
	return err
}

func (m *MemoryRegister) ConnectOutputProbe(probe gologicgates.LineProbe, idx int, send_value ...bool) error {
	var err error
	if idx >= 0 && idx < len(m.cells) {
		err = m.cells[idx].ConnectOutputProbe(probe, Dq, send_value...)
	}
	return err
}

func (m *MemoryRegister) DisconnectOutputProbe(probe gologicgates.LineProbe, idx int) error {
	var err error
	if idx >= 0 && idx <= len(m.cells) {
		err = m.cells[idx].DisconnectOutputProbe(probe, Dq)
	}
	return err
}
