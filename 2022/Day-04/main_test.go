package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	input string
	want  bool
	want2 bool
}{
	{input: "2-4,6-8", want: false, want2: false},
	{input: "2-3,4-5", want: false, want2: false},
	{input: "5-7,7-9", want: false, want2: true},
	{input: "2-8,3-7", want: true, want2: true},
	{input: "6-6,4-6", want: true, want2: true},
	{input: "2-6,4-8", want: false, want2: true},
}

func Test_IsFullyContained(t *testing.T) {
	t.Run("Is fully contained", func(t *testing.T) {
		for _, tt := range tests {
			actual := isFullyContained(tt.input)
			assert.Equal(t, tt.want, actual)
		}
	})
	t.Run("Is contained", func(t *testing.T) {
		for _, tt := range tests {
			actual := isContained(tt.input)
			assert.Equal(t, tt.want2, actual)
		}
	})
}
