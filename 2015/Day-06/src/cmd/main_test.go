package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var lightsTests = []struct {
	input              []string
	onLights           int
	expectedBrightness int
}{
	{
		[]string{
			"turn on 0,0 through 999,999",
			"toggle 0,0 through 999,0",
			"turn off 499,499 through 500,500",
		}, 998996, 1001996,
	},
}

func Test_Lights_SampleCases(t *testing.T) {
	for _, tt := range lightsTests {
		actualLights, actualBrightness := doOnLights(tt.input)
		assert.Equal(t, tt.onLights, actualLights)
		assert.Equal(t, tt.expectedBrightness, actualBrightness)
	}
}
