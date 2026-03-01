package gologicgates

import (
	"fmt"
	"slices"
)

// Defines an object that represents a logical connection similar to the way that a wire
// would work in an electrical circuit.  Because this is a simulation, values can only be
// sampled.  So this component will carry its current logical value.
type Line struct {
	// Specifies the last value set on this line
	value bool

	// Specifies the collection of probes objects that are connected to this wire.
	probes []LineProbe

	// Set to true if the line value has been set
	has_been_set bool
}

// Defines an interface that can receive notifications in the state of a logic line.
type LineProbe interface {
	// Called by the wire to register a change in the specified line's logic level
	//
	// Parameters
	//   - line: Specifies the line that has changed state
	//   - value: Specifies the new line state.
	OnLineChanged(line *Line, value bool)
}

// Defines a structure that represents a logic connection between the output of one gate and the input of another
type Connection struct {
	source, dest             BasicGate
	source_index, dest_index int
}

// Implements the change handler for the LineProbe interface
func (c *Connection) OnLineChanged(line *Line, value bool) {
	_, err := c.dest.SetInput(value, c.dest_index)
	if err != nil {
		panic(err)
	}
}

// Closes the connection between the source gate's output line and the dest gate's input line
//
// Returns:
//   - error: returns nil on success
func (c *Connection) Close() error {
	err := c.source.DisconnectOutputProbe(c, c.source_index)
	if err != nil {
		return fmt.Errorf("close connection failed: %v", err)
	}
	return nil
}

// Creates an object that represents the connection between a source output line and a dest input line
//
// Returns:
//   - *Connection: returns the connection on success
//   - error: returns nil on success
//
// Parameters:
//   - source: specifies the object with the output line
//   - source_index: specifies the index of the output line on the source gate
//   - dest: specifies the object with the input line
//   - dest_index: specifies the index of the input line
func NewConnection(source BasicGate, source_index int, dest BasicGate, dest_index int) (*Connection, error) {
	rtn := &Connection{
		source:       source,
		dest:         dest,
		source_index: source_index,
		dest_index:   dest_index,
	}
	err := source.ConnectOutputProbe(rtn, source_index, true)
	if err != nil {
		return nil, fmt.Errorf("connect gates failure: %v", err)
	}
	return rtn, nil
}

// Generates a connection object and panics if the is an error.
//
// Returns:
//   - *Connection: Returns the created connection
//
// Parameters:
//   - source_gate: specifies the source gate
//   - source_index: specifies the index number of the source gate output
//   - target_gate: specifies the target gate
//   - target_index: specifies the index of the target gate input
func MakeConnection(
	source_gate BasicGate, source_index int, target_gate BasicGate, target_index int,
) *Connection {
	rtn, err := NewConnection(source_gate, source_index, target_gate, target_index)
	if err != nil {
		panic(err)
	}
	return rtn
}

// Construct a new wire with an optional initial value.
//
// Returns:
//
//	*Line: Returns the new wire object
//
// Parameter:
//   - initial_value: Optionally specifies the initial logic value for the wire.  Defaults to a value of false.
func NewLine(initial_value ...bool) *Line {
	var rtn = &Line{value: false}
	if len(initial_value) > 0 {
		rtn.value = initial_value[0]
	}
	return rtn
}

// Polls the current logic value for the wire
//
// Returns:
//   - bool: returns the boolean value for the wire
func (w *Line) Value() bool {
	return w.value
}

// Sets the logical value for the line and invokes the Changed function of any connected probes
//
// Returns:
//   - bool: returns the new value for the wire
//
// Parameters:
//   - value: specifies the new logic level for the line
func (w *Line) Set(value bool) bool {
	if value != w.value || !w.has_been_set {
		w.has_been_set = true
		w.value = value
		for _, probe := range w.probes {
			probe.OnLineChanged(w, value)
		}
	}
	return w.value
}

// Reads the current logic level set for this line
//
// Returns:
//   - bool: returns the current logic level
func (w *Line) Measure() bool {
	return w.value
}

// Connect a probe for the given wire
//
// Parameters:
//   - probe: Specifies the probe object
//   - send_value: set to true if the probe's OnLineChanged() method should be immediately called
func (w *Line) ConnectProbe(probe LineProbe, send_value ...bool) {
	w.probes = append(w.probes, probe)
	if len(send_value) > 0 && send_value[0] {
		probe.OnLineChanged(w, w.value)
	}
}

// Disconnects the given probe from the line
//
// Parameters:
//   - probe: specifies the probe to disconnect
func (w *Line) DisconnectProbe(probe LineProbe) {
	// we need to find the position of the probe in the line's collection
	for pos, candidate := range w.probes {
		if candidate == probe {
			w.probes = slices.Delete(w.probes, pos, pos)
			break
		}
	}
}
