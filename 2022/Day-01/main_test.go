package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	input []string
	want  []int
}{
	{
		input: []string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"},
		want:  []int{24000, 11000, 10000, 6000, 4000},
	},
}

func Test_CountCalories(t *testing.T) {
	t.Run("Count calories top 1", func(t *testing.T) {
		for _, tt := range tests {
			actual := mapCaloriesToIntOrdered(tt.input)
			assert.Equal(t, tt.want, actual)
		}
	})
}
