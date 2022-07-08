package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var circuitsTests = []struct {
	input         []string
	expectedWireA uint16
}{
	{
		[]string{
			"x LSHIFT 2 -> f",
			"NOT y -> a",
			"x OR y -> e",
			"123 -> x",
			"x AND y -> d",
			"456 -> y",
			"NOT x -> h",
			"y RSHIFT 2 -> g",
		}, 65079,
	},
}

func Test_Circuits_SampleCases(t *testing.T) {
	for _, tt := range circuitsTests {
		wires := runCircuit(tt.input)
		assert.Equal(t, tt.expectedWireA, wires["a"])
	}
}
