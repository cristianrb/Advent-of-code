package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var nextPasswordTests = []struct {
	input  string
	result string
}{
	{"abcdefgh", "abcdffaa"},
	{"ghijklmn", "ghjaabcc"},
}

func Test_NextPassword(t *testing.T) {
	for _, tt := range nextPasswordTests {
		actual := nextPassword(tt.input)
		assert.Equal(t, tt.result, actual)
	}
}
