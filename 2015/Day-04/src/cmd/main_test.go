package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var countHousesTests = []struct {
	input    string
	nZeros   int
	expected int
}{
	{"abcdef", 5, 609043},
	{"pqrstuv", 5, 1048970},
}

func Test_LowestMd5Number_SampleCases(t *testing.T) {
	for _, tt := range countHousesTests {
		actual := lowestMd5Number(tt.input, tt.nZeros)
		assert.Equal(t, tt.expected, actual)
	}
}
