package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var countHousesTests = []struct {
	input    string
	expected map[Coords]bool
}{
	{
		">", map[Coords]bool{
			Coords{0, 0}: true,
			Coords{1, 0}: true,
		},
	},
	{
		"^>v<", map[Coords]bool{
			Coords{0, 0}: true,
			Coords{1, 0}: true,
			Coords{0, 1}: true,
			Coords{1, 1}: true,
		},
	},
	{
		"^v^v^v^v^v", map[Coords]bool{
			Coords{0, 0}: true,
			Coords{0, 1}: true,
		},
	},
}

func Test_CountHouses_SampleCases(t *testing.T) {
	for _, tt := range countHousesTests {
		actual := countHouses(tt.input)
		assert.Equal(t, tt.expected, actual)
	}
}
