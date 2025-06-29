package gologicgates_test

import (
	"testing"

	gologicgates "github.com/jhtrauntvein/go-logic-gates"
)

func TestAndGate(t *testing.T) {
	test_cases := [][]bool{
		{false, false, false},
		{false, true, false},
		{true, false, false},
		{true, true, true},
	}
	gate := gologicgates.NewAndGate2()
	default_value := gate.Evaluate()
	if len(default_value) != 1 {
		t.Errorf("expect and gate output to be an array of size 1")
	}
	if default_value[0] != false {
		t.Errorf("expect and gate default value to be false")
	}
	for _, test := range test_cases {
		_, err := gate.SetInput(test[0], 0)
		if err != nil {
			t.Errorf("failed to evaluate first input: %v", err)
		}
		results, err := gate.SetInput(test[1], 1)
		if err != nil {
			t.Errorf("failed to evaluate second input: %v", err)
		}
		if results[0] != test[2] {
			t.Errorf("expected %[1]t and got %[2]t", test[2], results[0])
		}
	}
}
