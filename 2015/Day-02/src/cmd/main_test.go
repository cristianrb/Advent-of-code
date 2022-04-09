package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var surfaceTests = []struct {
	length   int
	width    int
	height   int
	expected int
}{
	{2, 3, 4, 58},
	{1, 1, 10, 43},
}

func Test_FindSurface_SampleCases(t *testing.T) {
	for _, tt := range surfaceTests {
		actual := findSurface(tt.length, tt.width, tt.height)
		assert.Equal(t, tt.expected, actual)
	}
}
