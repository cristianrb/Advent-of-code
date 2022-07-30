package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var escapeTests = []struct {
	input         []string
	expectedCount int
}{
	{
		[]string{
			`""`,
			`"abc"`,
			`"aaa\"aaa"`,
			`"\x27"`,
		}, 12,
	},
}

func Test_EscapeCharacters(t *testing.T) {
	for _, tt := range escapeTests {
		characters := countTotalCharactersDecoded(tt.input)
		assert.Equal(t, tt.expectedCount, characters)
	}
}

var escapeTestsEncode = []struct {
	input         []string
	expectedCount int
}{
	{
		[]string{
			`""`,
			`"abc"`,
			`"aaa\"aaa"`,
			`"\x27"`,
		}, 19,
	},
}

func Test_EncodeCharacters(t *testing.T) {
	for _, tt := range escapeTestsEncode {
		characters := countTotalCharactersEncoded(tt.input)
		assert.Equal(t, tt.expectedCount, characters)
	}
}
