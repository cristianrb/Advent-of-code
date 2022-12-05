package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	input         map[int][]string
	instructions  []string
	preserveOrder bool
	want          string
}{
	{
		input:         generateInputTest(),
		instructions:  []string{"move 1 from 2 to 1", "move 3 from 1 to 3", "move 2 from 2 to 1", "move 1 from 1 to 2"},
		preserveOrder: false,
		want:          "CMZ",
	},
	{
		input:         generateInputTest(),
		instructions:  []string{"move 1 from 2 to 1", "move 3 from 1 to 3", "move 2 from 2 to 1", "move 1 from 1 to 2"},
		preserveOrder: true,
		want:          "MCD",
	},
}

func Test_TopCrates(t *testing.T) {
	t.Run("Top crates", func(t *testing.T) {
		for _, tt := range tests {
			actual := topCrates(tt.input, tt.instructions, tt.preserveOrder)
			assert.Equal(t, tt.want, actual)
		}
	})
}

func generateInputTest() map[int][]string {
	crates1 := []string{"Z", "N"}

	crates2 := []string{"M", "C", "D"}

	crates3 := []string{"P"}

	inputMap := map[int][]string{
		1: crates1,
		2: crates2,
		3: crates3,
	}

	return inputMap
}
