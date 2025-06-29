package gologicgates

import "slices"

// Defines an object that represents a logical connection similar to the way that a wire
// would work in an electrical circuit.  Because this is a simulation, values can only be
// sampled.  So this component will carry its current logical value.
type Line struct {
	// Specifies the last value set on this line
	value bool

	// Specifies the collection of probes objects that are connected to this wire.
	probes []LineProbe
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

// Construct a new wire with an optional initial value.
//
// Returns:
//   *Line: Returns the new wire object
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
	if value != w.value {
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
