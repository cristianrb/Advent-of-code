package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var lookAndSayTests = []struct {
	input  string
	result string
}{
	{"1", "11"},
	{"11", "21"},
	{"21", "1211"},
	{"1211", "111221"},
	{"111221", "312211"},
}

func Test_LookAndSay(t *testing.T) {
	for _, tt := range lookAndSayTests {
		actual := lookAndSay(tt.input)
		assert.Equal(t, tt.result, actual)
	}
}
