package gologicgates

import "errors"

// Defines errors that can be thrown by a BasicGate implementation
var (
	ErrSetInvalidLine = errors.New("invalid index to set line")
	ErrGetInvalidLine = errors.New("invalid index to probe or measure line")
)

// Defines an interface for a device that manages a collection of input lines and a collection of
// output lines.  For instance, the logic AND type gate can have two or more inputs and a single output
type BasicGate interface {
	// Set the value of one of the gate's input lines.  This will invoke the OnLineChange() method
	// for all probes attached to the line if the line value has been changed.
	//
	// Returns:
	//   []bool: Returns the matched solution on success
	//   error: Returns non-nil if the given line index is not valid.
	//
	// Parameters:
	//   value: Specifies the new value for the line
	//   index: Specifies the index for the input line
	SetInput(value bool, index int) ([]bool, error)

	// Evaluates the intended logic function of the gate and returns an array of booleans that
	// represent the state of each output pin.
	//
	// Returns:
	//   - []bool: Returns the new values for each output pin on the device
	Evaluate() []bool

	// Connects the specified probe to the input pin identified by the index.
	//
	// Returns:
	//   - error: Returns nil on success
	//
	// Parameters:
	//   - probe: Specifies the probe that will be connected to the line
	//   - index: Specifies the index of the input line
	//   - send_value: Optionally specifies that the probe should be notified of the line's value when connected.
	ConnectInputProbe(probe LineProbe, index int, send_value ...bool) error

	// Disconnects the specified probe from the line identified by the index
	//
	// Returns:
	//   - error: Returns nil on success
	//
	// Parameters:
	//   - probe: Specifies the probe to disconnect
	//   - index: Specifies the index in the array of input pins
	DisconnectInputProbe(probe LineProbe, index int) error

	// Connect the given probe to the output pin identified by the index.
	//
	// Returns:
	//   - error: Returns nil on success
	//
	// Parameters:
	//   - probe: Specifies the probe to be connected
	//   - index: Specifies the index for the output line
	//   - send_value: Optionally set to true if the probe should be notified of the line's value on connection.
	ConnectOutputProbe(probe LineProbe, index int, send_value ...bool) error

	// Disconnect the given outprobe from the device outputs.
	//
	// Returns:
	//   - error: returns nil on success
	//
	// Parameters:
	//   - probe: specifies the probe object to remove
	DisconnectOutputProbe(probe LineProbe, index int) error
}
