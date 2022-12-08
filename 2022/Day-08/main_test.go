package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	input [][]int
	want  int
	want2 int
}{
	{
		input: [][]int{
			{3, 0, 3, 7, 3},
			{2, 5, 5, 1, 2},
			{6, 5, 3, 3, 2},
			{3, 3, 5, 4, 9},
			{3, 5, 3, 9, 0},
		},
		want:  21,
		want2: 8,
	},
}

func Test_VisibleTrees(t *testing.T) {
	t.Run("Number of visible trees outside the grid", func(t *testing.T) {
		for _, tt := range tests {
			actual1, actual2 := visibleTrees(tt.input)
			assert.Equal(t, tt.want, actual1)
			assert.Equal(t, tt.want2, actual2)
		}
	})
}
