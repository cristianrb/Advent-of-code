package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var countHousesTests = []struct {
	input    string
	expected int
}{
	{">", 2},
	{"^>v<", 4},
	{"^v^v^v^v^v", 2},
}

func Test_CountHouses_SampleCases(t *testing.T) {
	for _, tt := range countHousesTests {
		actual := countHouses(tt.input)
		assert.Equal(t, tt.expected, actual)
	}
}
