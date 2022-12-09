package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	input   []string
	ropeLen int
	want    int
}{
	{
		input: []string{
			"R 4",
			"U 4",
			"L 3",
			"D 1",
			"R 4",
			"D 1",
			"L 5",
			"R 2",
		},
		ropeLen: 2,
		want:    13,
	},
	{
		input: []string{
			"R 5",
			"U 8",
			"L 8",
			"D 3",
			"R 17",
			"D 10",
			"L 25",
			"U 20",
		},
		ropeLen: 10,
		want:    36,
	},
}

func Test_TailDifferentPositions(t *testing.T) {
	t.Run("Number of positions the tail of the rope visit at least once", func(t *testing.T) {
		for _, tt := range tests {
			assert.Equal(t, tt.want, countTailPositions(tt.input, tt.ropeLen))
		}
	})
}
