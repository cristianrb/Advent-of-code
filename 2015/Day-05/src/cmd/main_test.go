package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var isNiceTests = []struct {
	input  string
	isNice bool
}{
	{"ugknbfddgicrmopn", true},
	{"aaa", true},
	{"jchzalrnumimnmhp", false},
	{"haegwjzuvuyypxyu", false},
	{"dvszwmarrgswjxmb", false},
}

var isNiceTestsV2 = []struct {
	input  string
	isNice bool
}{
	{"qjhvhtzxzqqjkmpb", true},
	{"xxyxx", true},
	{"uurcxstgmygtbstg", false},
	{"ieodomkazucvgmuy", false},
	{"aaa", false},
}

func Test_IsNice_SampleCases(t *testing.T) {
	for _, tt := range isNiceTests {
		actual := isNice(tt.input)
		assert.Equal(t, tt.isNice, actual)
	}
}

func Test_IsNiceV2_SampleCases(t *testing.T) {
	for _, tt := range isNiceTestsV2 {
		actual := isNiceV2(tt.input)
		assert.Equal(t, tt.isNice, actual)
	}
}
