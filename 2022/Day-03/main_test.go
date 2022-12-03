package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	input []string
	want  int
}{
	{
		input: []string{
			"vJrwpWtwJgWrhcsFMMfFFhFp",
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			"PmmdzqPrVvPwwTWBwg",
			"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			"ttgJtRGJQctTZtZT",
			"CrZsJsPPZsGzwwsLwLmpwMDw",
		},
		want: 157,
	},
}

var tests2 = []struct {
	input []string
	want  int
}{
	{
		input: []string{
			"vJrwpWtwJgWrhcsFMMfFFhFp",
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			"PmmdzqPrVvPwwTWBwg",
			"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			"ttgJtRGJQctTZtZT",
			"CrZsJsPPZsGzwwsLwLmpwMDw",
		},
		want: 70,
	},
}

func Test_RucksackPriorities(t *testing.T) {
	t.Run("Total Rucksack priorities", func(t *testing.T) {
		for _, tt := range tests {
			actual := rucksackPriorities(tt.input)
			assert.Equal(t, tt.want, actual)
		}
	})
	t.Run("Total elf badges", func(t *testing.T) {
		for _, tt := range tests2 {
			actual := elfBadges(tt.input)
			assert.Equal(t, tt.want, actual)
		}
	})
}
