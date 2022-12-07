package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	input []string
	want1 int
	want2 int
}{
	{
		input: []string{
			"$ ls",
			"dir a",
			"14848514 b.txt",
			"8504156 c.dat",
			"dir d",
			"$ cd a",
			"$ ls",
			"dir e",
			"29116 f",
			"2557 g",
			"62596 h.lst",
			"$ cd e",
			"$ ls",
			"584 i",
			"$ cd ..",
			"$ cd ..",
			"$ cd d",
			"$ ls",
			"4060174 j",
			"8033020 d.log",
			"5626152 d.ext",
			"7214296 k",
		},
		want1: 95437,
		want2: 24933642,
	},
}

func Test_SumOfTotalSizes(t *testing.T) {
	t.Run("Sum of total sizes of directories", func(t *testing.T) {
		for _, tt := range tests {
			root := createFileSystem(tt.input)
			assert.Equal(t, tt.want1, findDirectoriesWithSizeLessThan(root, 100000))
			assert.Equal(t, tt.want2, findSmallestDirectoryToDelete(root, 30000000-(70000000-root.Size)))
		}
	})
}
