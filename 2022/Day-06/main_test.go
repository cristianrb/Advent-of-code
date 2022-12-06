package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	input string
	n     int
	want  int
}{
	{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", n: 4, want: 7},
	{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", n: 4, want: 5},
	{input: "nppdvjthqldpwncqszvftbrmjlhg", n: 4, want: 6},
	{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", n: 4, want: 10},
	{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", n: 4, want: 11},

	{input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", n: 14, want: 19},
	{input: "bvwbjplbgvbhsrlpgdmjqwftvncz", n: 14, want: 23},
	{input: "nppdvjthqldpwncqszvftbrmjlhg", n: 14, want: 23},
	{input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", n: 14, want: 29},
	{input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", n: 14, want: 26},
}

func Test_FirstFourUniqueChars(t *testing.T) {
	t.Run("Find first four characters without duplicates", func(t *testing.T) {
		for _, tt := range tests {
			actual := firstNUniqueChars(tt.input, tt.n)
			assert.Equal(t, tt.want, actual)
		}
	})
}
