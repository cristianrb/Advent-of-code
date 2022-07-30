package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var jsonTests = []struct {
	input      string
	doNotDoRed bool
	result     float64
}{
	{`[1,2,3]`, false, 6},
	{`{"a":2,"b":4}`, false, 6},
	{`[[[3]]]`, false, 3},
	{`{"a":{"b":4},"c":-1}`, false, 3},
	{`{"a":[-1,1]}`, false, 0},
	{`[-1,{"a":1}]`, false, 0},
	{`{}`, false, 0},
	{`[]`, false, 0},
	{`[1,2,3]`, true, 6},
	{`[1,{"c":"red","b":2},3]`, true, 4},
	{`{"d":"red","e":[1,2,3,4],"f":5}`, true, 0},
	{`[1,"red",5]`, true, 6},
}

func Test_JsonTests(t *testing.T) {
	for _, tt := range jsonTests {
		var jsonUnmarshalled interface{}
		json.Unmarshal([]byte(tt.input), &jsonUnmarshalled)
		actual := walkJson(jsonUnmarshalled, tt.doNotDoRed)
		assert.Equal(t, tt.result, actual)
	}
}
