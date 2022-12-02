package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	input    []string
	outcomes map[string]int
	want     int
}{
	{
		input: []string{"A Y", "B X", "C Z"},
		outcomes: map[string]int{
			"A X": 3 + 1,
			"A Y": 6 + 2,
			"A Z": 3 + 0,
			"B X": 1 + 0,
			"B Y": 2 + 3,
			"B Z": 3 + 6,
			"C X": 1 + 6,
			"C Y": 2 + 0,
			"C Z": 3 + 3,
		},
		want: 15,
	},
	{
		input: []string{"A Y", "B X", "C Z"},
		outcomes: map[string]int{
			"A X": 3 + 0,
			"A Y": 1 + 3,
			"A Z": 2 + 6,
			"B X": 1 + 0,
			"B Y": 2 + 3,
			"B Z": 6 + 3,
			"C X": 2 + 0,
			"C Y": 3 + 3,
			"C Z": 6 + 1,
		},
		want: 12,
	},
}

func Test_RPSGames(t *testing.T) {
	t.Run("Total RPS score", func(t *testing.T) {
		for _, tt := range tests {
			actual := calcTotalScore(tt.input, tt.outcomes)
			assert.Equal(t, tt.want, actual)
		}
	})
}
